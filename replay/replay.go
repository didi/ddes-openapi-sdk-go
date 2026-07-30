package replay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
)

// replayFixture 是单条回放样本（与 tools/replaygen 输出格式一致）。
type replayFixture struct {
	URI      string                 `json:"uri"`
	Category string                 `json:"category"`
	In       map[string]interface{} `json:"in"`
	Out      map[string]interface{} `json:"out"`
}

// replayEntry 是 uri → 接口元数据 的映射表条目。
//
// build: 从 fixture.in 构造 *XxxApiReq（river 走逐字段反射，open-apis 走 param_json 注入）。
// call: 调用 resource 方法，返回 *XxxApiResp（其内嵌 ApiReply 的 Errno 用于断言）。
// reply: 从 *XxxApiResp 提取内嵌 *XxxApiReply 的 (errno, data)（字段覆盖对比的目标）。
type replayEntry struct {
	family string
	build  func(in map[string]interface{}) (apiReq interface{}, err error)
	call   func(client *didi.Client, apiReq interface{}) (resp interface{}, err error)
	reply  func(resp interface{}) (errno int32, data interface{})
}

// loadReplayFixtures 从 dir 加载全部 fixture 文件，返回 uri -> []fixture。
// dir 为空且默认 testdata/replay 不存在时返回 nil（调用方据此跳过）。
func loadReplayFixtures(dir string) map[string][]replayFixture {
	if dir == "" {
		dir = "testdata/replay"
		if env := os.Getenv("REPLAY_FIXTURES_DIR"); env != "" {
			dir = env
		}
	}
	if _, err := os.Stat(dir); err != nil {
		return nil
	}
	out := map[string][]replayFixture{}
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		var fixtures []replayFixture
		if err := json.Unmarshal(raw, &fixtures); err != nil {
			return nil
		}
		for _, f := range fixtures {
			out[f.URI] = append(out[f.URI], f)
		}
		return nil
	})
	return out
}

// replayServer 是可动态切换响应的 mock 服务端：setOut 更新某 uri 的当前 out，
// 逐条回放时每条 fixture 切换响应，使 SDK 真实反序列化该条 out。
type replayServer struct {
	*httptest.Server
	out map[string]map[string]interface{}
}

func newDynamicReplayServer() *replayServer {
	out := map[string]map[string]interface{}{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, ok := out[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		body, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}))
	return &replayServer{Server: srv, out: out}
}

// setOut 设置 uri 的当前响应 out。
func (s *replayServer) setOut(uri string, out map[string]interface{}) {
	s.out[uri] = out
}

// newReplayClient 构造指向 mock server 的测试 Client。
// signKey 用占位值——回放不校验 sign。BaseUrl 经导出字段 *core.Option 注入。
func newReplayClient(baseURL string) (*didi.Client, error) {
	return didi.NewClientWithOption(
		"test_client_id", "test_client_secret", "test_sign_key_for_replay_only",
		&core.Option{BaseUrl: baseURL},
	)
}

// fieldCoverage 对比 out.data 的顶层 key 与 reply Data struct 的 json tag，
// 返回「真实返回但 SDK 未接住」的字段清单。
func fieldCoverage(outData interface{}, replyData interface{}) []string {
	dataMap, ok := outData.(map[string]interface{})
	if !ok || len(dataMap) == 0 {
		return nil
	}
	tags := jsonTagsOfStruct(replyData)
	if len(tags) == 0 {
		return nil
	}
	tagSet := map[string]struct{}{}
	for _, t := range tags {
		tagSet[t] = struct{}{}
	}
	var missing []string
	for k := range dataMap {
		if _, ok := tagSet[k]; !ok {
			missing = append(missing, k)
		}
	}
	sort.Strings(missing)
	return missing
}

// jsonTagsOfStruct 反射收集 v（或其指针指向的 struct）顶层字段的 json tag。
func jsonTagsOfStruct(v interface{}) []string {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	rt := rv.Type()
	var tags []string
	for i := 0; i < rt.NumField(); i++ {
		tag := rt.Field(i).Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			continue
		}
		tags = append(tags, name)
	}
	return tags
}

// buildRiverApiReq 通用 /river/ 回填：剔除通用字段后，按 snake_case→CamelCase
// 反射调用 Builder 方法 + 类型适配。builder 是 *XxxApiReqBuilder，返回其 Build()。
func buildRiverApiReq(builder interface{}, in map[string]interface{}) (interface{}, error) {
	bv := reflect.ValueOf(builder)
	common := map[string]struct{}{
		"client_id": {}, "client_secret": {}, "access_token": {},
		"company_id": {}, "sign": {}, "timestamp": {}, "sign_key": {},
	}
	for k, v := range in {
		if _, ok := common[k]; ok {
			continue
		}
		methodName := snakeToCamel(k)
		mv := bv.MethodByName(methodName)
		if !mv.IsValid() {
			// 字段无对应 Builder 方法，跳过（记录由调用方处理）
			continue
		}
		arg, err := adaptArg(v, mv.Type())
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", k, err)
		}
		mv.Call([]reflect.Value{arg})
	}
	build := bv.MethodByName("Build")
	if !build.IsValid() {
		return nil, fmt.Errorf("builder 缺 Build 方法")
	}
	return build.Call(nil)[0].Interface(), nil
}

// adaptArg 将 in 字段值适配为 Builder 方法参数类型（string/int32/int64/float64 等）。
func adaptArg(v interface{}, methodType reflect.Type) (reflect.Value, error) {
	if methodType.NumIn() != 1 {
		return reflect.Value{}, fmt.Errorf("方法参数数量不为 1")
	}
	pt := methodType.In(0)
	switch pt.Kind() {
	case reflect.String:
		s, err := toString(v)
		return reflect.ValueOf(s), err
	case reflect.Int32:
		i, err := toInt64(v)
		return reflect.ValueOf(int32(i)), err
	case reflect.Int64:
		i, err := toInt64(v)
		return reflect.ValueOf(i), err
	case reflect.Int:
		i, err := toInt64(v)
		return reflect.ValueOf(int(i)), err
	case reflect.Float64:
		f, err := toFloat64(v)
		return reflect.ValueOf(f), err
	case reflect.Float32:
		f, err := toFloat64(v)
		return reflect.ValueOf(float32(f)), err
	case reflect.Bool:
		b, err := toBool(v)
		return reflect.ValueOf(b), err
	default:
		return reflect.Value{}, fmt.Errorf("不支持的参数类型 %s", pt.Kind())
	}
}

func toString(v interface{}) (string, error) {
	switch x := v.(type) {
	case string:
		return x, nil
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(x), nil
	case nil:
		return "", nil
	default:
		return fmt.Sprintf("%v", x), nil
	}
}

func toInt64(v interface{}) (int64, error) {
	switch x := v.(type) {
	case float64:
		return int64(x), nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	case bool:
		if x {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("无法转 int: %T", v)
	}
}

func toFloat64(v interface{}) (float64, error) {
	switch x := v.(type) {
	case float64:
		return x, nil
	case string:
		return strconv.ParseFloat(x, 64)
	default:
		return 0, fmt.Errorf("无法转 float: %T", v)
	}
}

func toBool(v interface{}) (bool, error) {
	switch x := v.(type) {
	case bool:
		return x, nil
	case string:
		return strconv.ParseBool(x)
	case float64:
		return x != 0, nil
	default:
		return false, fmt.Errorf("无法转 bool: %T", v)
	}
}

// snakeToCamel: order_id -> OrderId, need_rule_info -> NeedRuleInfo。
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

// extractReply 反射从 *XxxApiResp 提取内嵌 *XxxApiReply 的 (errno, data)。
// ApiResp 有一个类型为 *XxxApiReply 的字段；ApiReply 有 Errno int32 与 Data *XxxReply。
func extractReply(resp interface{}) (errno int32, data interface{}, ok bool) {
	if resp == nil {
		return 0, nil, false
	}
	rv := reflect.ValueOf(resp)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return 0, nil, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return 0, nil, false
	}
	// 找类型名以 ApiReply 结尾的字段
	rt := rv.Type()
	var replyField reflect.Value
	for i := 0; i < rt.NumField(); i++ {
		if strings.HasSuffix(rt.Field(i).Type.String(), "ApiReply") {
			replyField = rv.Field(i)
			break
		}
	}
	if !replyField.IsValid() || replyField.IsNil() {
		return 0, nil, false
	}
	reply := replyField.Elem()
	replyT := reply.Type()
	errnoV := reply.FieldByName("Errno")
	if errnoV.IsValid() {
		ev := errnoV
		for ev.Kind() == reflect.Ptr {
			if ev.IsNil() {
				break
			}
			ev = ev.Elem()
		}
		if ev.Kind() >= reflect.Int && ev.Kind() <= reflect.Int64 {
			errno = int32(ev.Int())
		}
	}
	dataV := reply.FieldByName("Data")
	if dataV.IsValid() {
		// Data 可能是 *XxxReply（指针）或 XxxReply（值结构体）
		for dataV.Kind() == reflect.Ptr {
			if dataV.IsNil() {
				break
			}
			dataV = dataV.Elem()
		}
		if dataV.Kind() != reflect.Ptr || !dataV.IsNil() {
			data = dataV.Interface()
		}
	}
	_ = replyT
	return errno, data, true
}

// buildRequestApiReq 通用 POST-body 回填：构造 RequestBuilder（param_json 或直接字段），
// Build 后注入 ApiReqBuilder 的 XxxRequest(req) 方法。
// - param_json 接口（open-apis）：in 含 param_json 字符串，调 RequestBuilder.ParamJson。
// - 直接字段接口（river POST-body）：in 为平铺字段，反射调 RequestBuilder 字段方法。
func buildRequestApiReq(apiReqBuilder interface{}, requestBuilder interface{}, in map[string]interface{}) (interface{}, error) {
	rb := reflect.ValueOf(requestBuilder)
	if pj, ok := in["param_json"]; ok {
		pjStr, _ := toString(pj)
		if mv := rb.MethodByName("ParamJson"); mv.IsValid() {
			mv.Call([]reflect.Value{reflect.ValueOf(pjStr)})
		}
	} else {
		// 直接字段：复用 river 平铺反射，但作用在 RequestBuilder 上
		if _, err := buildRiverApiReq(requestBuilder, in); err != nil {
			return nil, err
		}
	}
	req := rb.MethodByName("Build").Call(nil)[0]
	// 找 ApiReqBuilder 上接收 *XxxRequest 的方法
	ab := reflect.ValueOf(apiReqBuilder)
	abT := ab.Type()
	for i := 0; i < abT.NumMethod(); i++ {
		m := abT.Method(i)
		if m.Type.NumIn() == 2 && m.Type.In(1) == req.Type() {
			ab.Method(i).Call([]reflect.Value{req})
			break
		}
	}
	build := ab.MethodByName("Build")
	if !build.IsValid() {
		return nil, fmt.Errorf("apiReqBuilder 缺 Build 方法")
	}
	return build.Call(nil)[0].Interface(), nil
}
