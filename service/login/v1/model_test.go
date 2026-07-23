package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// --- GetLoginEncryptStrApiReqBuilder 测试 ---

func TestGetLoginEncryptStrApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		AppType("1").
		Phone("13800000001").
		Email("test@example.com").
		EmployeeNumber("D0001").
		ProductType(2).
		Build()

	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"app_type", "1"},
		{"phone", "13800000001"},
		{"email", "test@example.com"},
		{"employee_number", "D0001"},
		{"product_type", "2"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetLoginEncryptStrApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("phone") != "13800000001" {
		t.Errorf("phone mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("app_type") != "" {
		t.Errorf("app_type should be empty, got %q", req.apiReq.QueryParams.Get("app_type"))
	}
	if req.apiReq.QueryParams.Get("email") != "" {
		t.Errorf("email should be empty, got %q", req.apiReq.QueryParams.Get("email"))
	}
	if req.apiReq.QueryParams.Get("employee_number") != "" {
		t.Errorf("employee_number should be empty, got %q", req.apiReq.QueryParams.Get("employee_number"))
	}
	if req.apiReq.QueryParams.Get("product_type") != "" {
		t.Errorf("product_type should be empty, got %q", req.apiReq.QueryParams.Get("product_type"))
	}
}

func TestGetLoginEncryptStrApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		ProductType(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("product_type") != "0" {
		t.Errorf("product_type = %q, want \"0\"", req.apiReq.QueryParams.Get("product_type"))
	}
}

func TestGetLoginEncryptStrApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()

	// 仅通用参数
	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
	// 业务参数不应存在
	for _, key := range []string{"app_type", "phone", "email", "employee_number", "product_type"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- GetLoginEncryptStrApiReply 反序列化测试 ---

func TestGetLoginEncryptStrApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"encrypt_str": "https://page.es.xiaojukeji.com/login?token=xxx",
			"callCarNow": {
				"rule_id": "1125922289295589",
				"jumpPage": "callCarNow",
				"passenger_phone": "13800000001"
			},
			"pcLogin": {
				"jumpPage": "bill"
			}
		},
		"request_id": "test_request_id"
	}`

	var reply GetLoginEncryptStrApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.Data.EncryptStr == nil || *reply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=xxx" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=xxx", reply.Data.EncryptStr)
	}
	if reply.Data.CallCarNow == nil {
		t.Fatal("CallCarNow is nil")
	}
	if reply.Data.CallCarNow.RuleId == nil || *reply.Data.CallCarNow.RuleId != "1125922289295589" {
		t.Errorf("CallCarNow.RuleId = %v, want 1125922289295589", reply.Data.CallCarNow.RuleId)
	}
	if reply.Data.CallCarNow.JumpPage == nil || *reply.Data.CallCarNow.JumpPage != "callCarNow" {
		t.Errorf("CallCarNow.JumpPage = %v, want callCarNow", reply.Data.CallCarNow.JumpPage)
	}
	if reply.Data.PcLogin == nil {
		t.Fatal("PcLogin is nil")
	}
	if reply.Data.PcLogin.JumpPage == nil || *reply.Data.PcLogin.JumpPage != "bill" {
		t.Errorf("PcLogin.JumpPage = %v, want bill", reply.Data.PcLogin.JumpPage)
	}
}

func TestGetLoginEncryptStrApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply GetLoginEncryptStrApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	// 错误响应 data 应为空
	if reply.Data.EncryptStr != nil {
		t.Errorf("Data.EncryptStr = %v, want nil", reply.Data.EncryptStr)
	}
}

func TestGetLoginEncryptStrApiReply_PartialFields(t *testing.T) {
	// 部分字段缺失
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"encrypt_str": "https://page.es.xiaojukeji.com/login?token=yyy",
			"orderDetail": {
				"order_id": "1125922289295589"
			}
		},
		"request_id": "req_partial"
	}`

	var reply GetLoginEncryptStrApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.EncryptStr == nil || *reply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=yyy" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=yyy", reply.Data.EncryptStr)
	}
	if reply.Data.OrderDetail == nil {
		t.Fatal("OrderDetail is nil")
	}
	if reply.Data.OrderDetail.OrderId == nil || *reply.Data.OrderDetail.OrderId != "1125922289295589" {
		t.Errorf("OrderDetail.OrderId = %v, want 1125922289295589", reply.Data.OrderDetail.OrderId)
	}
	// 未设置的子模型应为 nil
	if reply.Data.CallCarNow != nil {
		t.Errorf("CallCarNow = %v, want nil", reply.Data.CallCarNow)
	}
	if reply.Data.PcLogin != nil {
		t.Errorf("PcLogin = %v, want nil", reply.Data.PcLogin)
	}
}

func TestGetLoginEncryptStrApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":""},"request_id":"req_empty"}`

	var reply GetLoginEncryptStrApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	// encrypt_str 为空字符串时 omitempty 不生效（json 反序列化空字符串赋给指针为 nil）
	// 验证无 panic 即可
}

func TestGetLoginEncryptStrApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetLoginEncryptStrApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	// 无 data 字段，所有指针字段应为 nil
	if reply.Data.EncryptStr != nil {
		t.Errorf("Data.EncryptStr = %v, want nil", reply.Data.EncryptStr)
	}
	if reply.Data.CallCarNow != nil {
		t.Errorf("Data.CallCarNow = %v, want nil", reply.Data.CallCarNow)
	}
}

// --- 数据模型 Builder 测试 ---

func TestPCLoginBuilder(t *testing.T) {
	// 全字段
	login := NewPCLoginBuilder().
		JumpPage("bill").
		Build()
	if login.JumpPage == nil || *login.JumpPage != "bill" {
		t.Errorf("JumpPage = %v, want bill", login.JumpPage)
	}

	// 部分字段
	login2 := NewPCLoginBuilder().Build()
	if login2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", login2.JumpPage)
	}
}

func TestH5CallCarNowBuilder(t *testing.T) {
	// 全字段
	h5 := NewH5CallCarNowBuilder().
		RuleId("1125922289295589").
		JumpPage("callCarNow").
		ApprovalId("1125916593600206").
		LatFrom("39.984728").
		LngFrom("116.307585").
		PoiFromName("望京SOHO").
		ToCityId("1").
		LatTo("40.0572").
		LngTo("116.3045").
		PoiToName("首都机场").
		CityId("1").
		PassengerPhone("13800000001").
		PassengerName("张三").
		RestrictPoiFlag("1").
		RestrictPassenger("1").
		RequireLevelList("express,comfort").
		AppendCar("0").
		CallbackInfo("custom_info").
		HideEstimatePriceFlag("0").
		Build()

	if h5.RuleId == nil || *h5.RuleId != "1125922289295589" {
		t.Errorf("RuleId = %v, want 1125922289295589", h5.RuleId)
	}
	if h5.JumpPage == nil || *h5.JumpPage != "callCarNow" {
		t.Errorf("JumpPage = %v, want callCarNow", h5.JumpPage)
	}
	if h5.ApprovalId == nil || *h5.ApprovalId != "1125916593600206" {
		t.Errorf("ApprovalId = %v, want 1125916593600206", h5.ApprovalId)
	}
	if h5.LatFrom == nil || *h5.LatFrom != "39.984728" {
		t.Errorf("LatFrom = %v, want 39.984728", h5.LatFrom)
	}
	if h5.LngFrom == nil || *h5.LngFrom != "116.307585" {
		t.Errorf("LngFrom = %v, want 116.307585", h5.LngFrom)
	}
	if h5.PoiFromName == nil || *h5.PoiFromName != "望京SOHO" {
		t.Errorf("PoiFromName = %v, want 望京SOHO", h5.PoiFromName)
	}
	if h5.ToCityId == nil || *h5.ToCityId != "1" {
		t.Errorf("ToCityId = %v, want 1", h5.ToCityId)
	}
	if h5.LatTo == nil || *h5.LatTo != "40.0572" {
		t.Errorf("LatTo = %v, want 40.0572", h5.LatTo)
	}
	if h5.LngTo == nil || *h5.LngTo != "116.3045" {
		t.Errorf("LngTo = %v, want 116.3045", h5.LngTo)
	}
	if h5.PoiToName == nil || *h5.PoiToName != "首都机场" {
		t.Errorf("PoiToName = %v, want 首都机场", h5.PoiToName)
	}
	if h5.CityId == nil || *h5.CityId != "1" {
		t.Errorf("CityId = %v, want 1", h5.CityId)
	}
	if h5.PassengerPhone == nil || *h5.PassengerPhone != "13800000001" {
		t.Errorf("PassengerPhone = %v, want 13800000001", h5.PassengerPhone)
	}
	if h5.PassengerName == nil || *h5.PassengerName != "张三" {
		t.Errorf("PassengerName = %v, want 张三", h5.PassengerName)
	}
	if h5.RestrictPoiFlag == nil || *h5.RestrictPoiFlag != "1" {
		t.Errorf("RestrictPoiFlag = %v, want 1", h5.RestrictPoiFlag)
	}
	if h5.RestrictPassenger == nil || *h5.RestrictPassenger != "1" {
		t.Errorf("RestrictPassenger = %v, want 1", h5.RestrictPassenger)
	}
	if h5.RequireLevelList == nil || *h5.RequireLevelList != "express,comfort" {
		t.Errorf("RequireLevelList = %v, want express,comfort", h5.RequireLevelList)
	}
	if h5.AppendCar == nil || *h5.AppendCar != "0" {
		t.Errorf("AppendCar = %v, want 0", h5.AppendCar)
	}
	if h5.CallbackInfo == nil || *h5.CallbackInfo != "custom_info" {
		t.Errorf("CallbackInfo = %v, want custom_info", h5.CallbackInfo)
	}
	if h5.HideEstimatePriceFlag == nil || *h5.HideEstimatePriceFlag != "0" {
		t.Errorf("HideEstimatePriceFlag = %v, want 0", h5.HideEstimatePriceFlag)
	}

	// 部分字段
	h5_2 := NewH5CallCarNowBuilder().
		RuleId("1001").
		Build()
	if h5_2.RuleId == nil || *h5_2.RuleId != "1001" {
		t.Errorf("RuleId = %v, want 1001", h5_2.RuleId)
	}
	if h5_2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", h5_2.JumpPage)
	}
	if h5_2.PassengerPhone != nil {
		t.Errorf("PassengerPhone = %v, want nil", h5_2.PassengerPhone)
	}
}

func TestH5OrderDetailBuilder(t *testing.T) {
	detail := NewH5OrderDetailBuilder().
		OrderId("1125922289295589").
		Build()
	if detail.OrderId == nil || *detail.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", detail.OrderId)
	}

	// 部分字段
	detail2 := NewH5OrderDetailBuilder().Build()
	if detail2.OrderId != nil {
		t.Errorf("OrderId = %v, want nil", detail2.OrderId)
	}
}

func TestH5MyWalletBuilder(t *testing.T) {
	wallet := NewH5MyWalletBuilder().
		JumpPage("mywallet").
		Build()
	if wallet.JumpPage == nil || *wallet.JumpPage != "mywallet" {
		t.Errorf("JumpPage = %v, want mywallet", wallet.JumpPage)
	}

	wallet2 := NewH5MyWalletBuilder().Build()
	if wallet2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", wallet2.JumpPage)
	}
}

func TestH5CallCenterBuilder(t *testing.T) {
	cc := NewH5CallCenterBuilder().
		JumpPage("callCenter").
		Build()
	if cc.JumpPage == nil || *cc.JumpPage != "callCenter" {
		t.Errorf("JumpPage = %v, want callCenter", cc.JumpPage)
	}

	cc2 := NewH5CallCenterBuilder().Build()
	if cc2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", cc2.JumpPage)
	}
}

func TestH5HomeAddressBuilder(t *testing.T) {
	addr := NewH5HomeAddressBuilder().
		JumpPage("homeAddress").
		Build()
	if addr.JumpPage == nil || *addr.JumpPage != "homeAddress" {
		t.Errorf("JumpPage = %v, want homeAddress", addr.JumpPage)
	}

	addr2 := NewH5HomeAddressBuilder().Build()
	if addr2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", addr2.JumpPage)
	}
}

func TestH5InvoiceBuilder(t *testing.T) {
	inv := NewH5InvoiceBuilder().
		JumpPage("invoice").
		Build()
	if inv.JumpPage == nil || *inv.JumpPage != "invoice" {
		t.Errorf("JumpPage = %v, want invoice", inv.JumpPage)
	}

	inv2 := NewH5InvoiceBuilder().Build()
	if inv2.JumpPage != nil {
		t.Errorf("JumpPage = %v, want nil", inv2.JumpPage)
	}
}

// --- GetLoginEncryptStr 资源方法测试 ---

func newLoginTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestGetLoginEncryptStr_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Login/getLoginEncryptStr" {
			t.Errorf("expected path /river/Login/getLoginEncryptStr, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("phone") != "13800000001" {
			t.Errorf("phone = %q, want 13800000001", r.URL.Query().Get("phone"))
		}
		if r.URL.Query().Get("product_type") != "2" {
			t.Errorf("product_type = %q, want 2", r.URL.Query().Get("product_type"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"encrypt_str": "https://page.es.xiaojukeji.com/login?token=abc123",
				"callCarNow": {
					"rule_id": "1125922289295589",
					"jumpPage": "callCarNow"
				},
				"pcLogin": {
					"jumpPage": "bill"
				}
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Phone("13800000001").
		ProductType(2).
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetLoginEncryptStrApiReply.Errno)
	}
	if resp.GetLoginEncryptStrApiReply.Data.EncryptStr == nil ||
		*resp.GetLoginEncryptStrApiReply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=abc123" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=abc123",
			resp.GetLoginEncryptStrApiReply.Data.EncryptStr)
	}
	if resp.GetLoginEncryptStrApiReply.Data.CallCarNow == nil {
		t.Fatal("CallCarNow is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Data.CallCarNow.RuleId == nil ||
		*resp.GetLoginEncryptStrApiReply.Data.CallCarNow.RuleId != "1125922289295589" {
		t.Errorf("CallCarNow.RuleId = %v, want 1125922289295589",
			resp.GetLoginEncryptStrApiReply.Data.CallCarNow.RuleId)
	}
}

func TestGetLoginEncryptStr_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":""},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetLoginEncryptStrApiReply.Errno)
	}
}

func TestGetLoginEncryptStr_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.GetLoginEncryptStrApiReply.Errno)
	}
}

func TestGetLoginEncryptStr_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetLoginEncryptStrApiReply != nil {
		t.Errorf("GetLoginEncryptStrApiReply should be nil for non-200 response")
	}
}

func TestGetLoginEncryptStr_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":"https://page.es.xiaojukeji.com/login?token=enc128","pcLogin":{"jumpPage":"bill"}},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetLoginEncryptStrApiReply.Errno)
	}
	if resp.GetLoginEncryptStrApiReply.Data.EncryptStr == nil ||
		*resp.GetLoginEncryptStrApiReply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=enc128" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=enc128",
			resp.GetLoginEncryptStrApiReply.Data.EncryptStr)
	}
}

func TestGetLoginEncryptStr_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":"https://page.es.xiaojukeji.com/login?token=enc256"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Data.EncryptStr == nil ||
		*resp.GetLoginEncryptStrApiReply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=enc256" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=enc256",
			resp.GetLoginEncryptStrApiReply.Data.EncryptStr)
	}
}

func TestGetLoginEncryptStr_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":"https://page.es.xiaojukeji.com/login?token=noenc"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	l := &login{option: option}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
	if resp.GetLoginEncryptStrApiReply.Data.EncryptStr == nil ||
		*resp.GetLoginEncryptStrApiReply.Data.EncryptStr != "https://page.es.xiaojukeji.com/login?token=noenc" {
		t.Errorf("EncryptStr = %v, want https://page.es.xiaojukeji.com/login?token=noenc",
			resp.GetLoginEncryptStrApiReply.Data.EncryptStr)
	}
}

func TestGetLoginEncryptStr_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"encrypt_str":"https://page.es.xiaojukeji.com/login?token=opt"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newLoginTestOption(testServer.URL)
	l := &login{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetLoginEncryptStrApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := l.GetLoginEncryptStr(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetLoginEncryptStr() error = %v", err)
	}
	if resp.GetLoginEncryptStrApiReply == nil {
		t.Fatal("GetLoginEncryptStrApiReply is nil")
	}
}
