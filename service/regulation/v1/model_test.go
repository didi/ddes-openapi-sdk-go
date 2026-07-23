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

// --- RegulationInfo Builder 测试 ---

func TestRegulationInfoBuilder(t *testing.T) {
	info := NewRegulationInfoBuilder().
		RegulationId("1125920826148759").
		RegulationName("差旅制度").
		RegulationStatus("1").
		Source(3).
		SceneType("2").
		IsUseQuota(1).
		IsApprove(2).
		InstitutionId("1125900000000001").
		CityType(2).
		ApprovalType(1).
		Build()

	if info.RegulationId == nil || *info.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", info.RegulationId)
	}
	if info.RegulationName == nil || *info.RegulationName != "差旅制度" {
		t.Errorf("RegulationName = %v, want 差旅制度", info.RegulationName)
	}
	if info.RegulationStatus == nil || *info.RegulationStatus != "1" {
		t.Errorf("RegulationStatus = %v, want 1", info.RegulationStatus)
	}
	if info.Source == nil || *info.Source != 3 {
		t.Errorf("Source = %v, want 3", info.Source)
	}
	if info.SceneType == nil || *info.SceneType != "2" {
		t.Errorf("SceneType = %v, want 2", info.SceneType)
	}
	if info.IsUseQuota == nil || *info.IsUseQuota != 1 {
		t.Errorf("IsUseQuota = %v, want 1", info.IsUseQuota)
	}
	if info.IsApprove == nil || *info.IsApprove != 2 {
		t.Errorf("IsApprove = %v, want 2", info.IsApprove)
	}
	if info.InstitutionId == nil || *info.InstitutionId != "1125900000000001" {
		t.Errorf("InstitutionId = %v, want 1125900000000001", info.InstitutionId)
	}
	if info.CityType == nil || *info.CityType != 2 {
		t.Errorf("CityType = %v, want 2", info.CityType)
	}
	if info.ApprovalType == nil || *info.ApprovalType != 1 {
		t.Errorf("ApprovalType = %v, want 1", info.ApprovalType)
	}

	// 部分设置
	info2 := NewRegulationInfoBuilder().
		RegulationId("1001").
		Source(0).
		Build()

	if info2.RegulationId == nil || *info2.RegulationId != "1001" {
		t.Errorf("RegulationId = %v, want 1001", info2.RegulationId)
	}
	// int32 零值也应当被设置
	if info2.Source == nil || *info2.Source != 0 {
		t.Errorf("Source = %v, want 0", info2.Source)
	}
	// 未设置的字段应为 nil
	if info2.RegulationName != nil {
		t.Errorf("RegulationName = %v, want nil", info2.RegulationName)
	}
	if info2.SceneType != nil {
		t.Errorf("SceneType = %v, want nil", info2.SceneType)
	}
	if info2.IsUseQuota != nil {
		t.Errorf("IsUseQuota = %v, want nil", info2.IsUseQuota)
	}
	if info2.IsApprove != nil {
		t.Errorf("IsApprove = %v, want nil", info2.IsApprove)
	}
	if info2.InstitutionId != nil {
		t.Errorf("InstitutionId = %v, want nil", info2.InstitutionId)
	}
	if info2.CityType != nil {
		t.Errorf("CityType = %v, want nil", info2.CityType)
	}
	if info2.ApprovalType != nil {
		t.Errorf("ApprovalType = %v, want nil", info2.ApprovalType)
	}
	if info2.RegulationStatus != nil {
		t.Errorf("RegulationStatus = %v, want nil", info2.RegulationStatus)
	}
}

// --- GetRegulationApiReqBuilder 测试 ---

func TestGetRegulationApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		RegulationId("1125920826148759").
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
		{"regulation_id", "1125920826148759"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetRegulationApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		RegulationId("1125920826148759").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("regulation_id") != "1125920826148759" {
		t.Errorf("regulation_id mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty, got %q", req.apiReq.QueryParams.Get("access_token"))
	}
	if req.apiReq.QueryParams.Get("company_id") != "" {
		t.Errorf("company_id should be empty, got %q", req.apiReq.QueryParams.Get("company_id"))
	}
	if req.apiReq.QueryParams.Get("timestamp") != "" {
		t.Errorf("timestamp should be empty, got %q", req.apiReq.QueryParams.Get("timestamp"))
	}
	if req.apiReq.QueryParams.Get("sign") != "" {
		t.Errorf("sign should be empty, got %q", req.apiReq.QueryParams.Get("sign"))
	}
}

func TestGetRegulationApiReqBuilder_EmptyRegulationId(t *testing.T) {
	// regulation_id 设为空字符串也应写入 query params
	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		RegulationId("").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("regulation_id") != "" {
		t.Errorf("regulation_id = %q, want empty", req.apiReq.QueryParams.Get("regulation_id"))
	}
}

func TestGetRegulationApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetRegulationApiReqBuilder().
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
	if v := req.apiReq.QueryParams.Get("regulation_id"); v != "" {
		t.Errorf("QueryParams[regulation_id] = %q, want empty", v)
	}
}

// --- GetRegulationApiReply 反序列化测试 ---

func TestGetRegulationApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"regulation_id": "1125920826148759",
			"regulation_name": "差旅制度",
			"regulation_status": "1",
			"source": 3,
			"scene_type": "2",
			"is_use_quota": 1,
			"is_approve": 2,
			"institution_id": "1125900000000001",
			"city_type": 2,
			"approval_type": 1
		},
		"request_id": "test_request_id"
	}`

	var reply GetRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}

	data := reply.Data
	if data.RegulationId == nil || *data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", data.RegulationId)
	}
	if data.RegulationName == nil || *data.RegulationName != "差旅制度" {
		t.Errorf("RegulationName = %v, want 差旅制度", data.RegulationName)
	}
	if data.RegulationStatus == nil || *data.RegulationStatus != "1" {
		t.Errorf("RegulationStatus = %v, want 1", data.RegulationStatus)
	}
	if data.Source == nil || *data.Source != 3 {
		t.Errorf("Source = %v, want 3", data.Source)
	}
	if data.SceneType == nil || *data.SceneType != "2" {
		t.Errorf("SceneType = %v, want 2", data.SceneType)
	}
	if data.IsUseQuota == nil || *data.IsUseQuota != 1 {
		t.Errorf("IsUseQuota = %v, want 1", data.IsUseQuota)
	}
	if data.IsApprove == nil || *data.IsApprove != 2 {
		t.Errorf("IsApprove = %v, want 2", data.IsApprove)
	}
	if data.InstitutionId == nil || *data.InstitutionId != "1125900000000001" {
		t.Errorf("InstitutionId = %v, want 1125900000000001", data.InstitutionId)
	}
	if data.CityType == nil || *data.CityType != 2 {
		t.Errorf("CityType = %v, want 2", data.CityType)
	}
	if data.ApprovalType == nil || *data.ApprovalType != 1 {
		t.Errorf("ApprovalType = %v, want 1", data.ApprovalType)
	}
}

func TestGetRegulationApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply GetRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestGetRegulationApiReply_PartialFields(t *testing.T) {
	// 部分字段缺失
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"regulation_id": "1125920826148759",
			"regulation_name": "差旅制度",
			"source": 0
		},
		"request_id": "req_partial"
	}`

	var reply GetRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.RegulationId == nil || *reply.Data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", reply.Data.RegulationId)
	}
	if reply.Data.RegulationName == nil || *reply.Data.RegulationName != "差旅制度" {
		t.Errorf("RegulationName = %v, want 差旅制度", reply.Data.RegulationName)
	}
	// source 为 0，应被设置
	if reply.Data.Source == nil || *reply.Data.Source != 0 {
		t.Errorf("Source = %v, want 0", reply.Data.Source)
	}
	// 缺失字段应为 nil
	if reply.Data.RegulationStatus != nil {
		t.Errorf("RegulationStatus = %v, want nil", reply.Data.RegulationStatus)
	}
	if reply.Data.SceneType != nil {
		t.Errorf("SceneType = %v, want nil", reply.Data.SceneType)
	}
	if reply.Data.IsUseQuota != nil {
		t.Errorf("IsUseQuota = %v, want nil", reply.Data.IsUseQuota)
	}
	if reply.Data.IsApprove != nil {
		t.Errorf("IsApprove = %v, want nil", reply.Data.IsApprove)
	}
	if reply.Data.InstitutionId != nil {
		t.Errorf("InstitutionId = %v, want nil", reply.Data.InstitutionId)
	}
	if reply.Data.CityType != nil {
		t.Errorf("CityType = %v, want nil", reply.Data.CityType)
	}
	if reply.Data.ApprovalType != nil {
		t.Errorf("ApprovalType = %v, want nil", reply.Data.ApprovalType)
	}
}

func TestGetRegulationApiReply_EmptyData(t *testing.T) {
	// data 为空对象
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {},
		"request_id": "req_empty"
	}`

	var reply GetRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.RegulationId != nil {
		t.Errorf("RegulationId = %v, want nil", reply.Data.RegulationId)
	}
	if reply.Data.RegulationName != nil {
		t.Errorf("RegulationName = %v, want nil", reply.Data.RegulationName)
	}
	if reply.Data.Source != nil {
		t.Errorf("Source = %v, want nil", reply.Data.Source)
	}
}

func TestGetRegulationApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

// --- ListRegulationApiReqBuilder 测试 ---

func TestListRegulationApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
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
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListRegulationApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty, got %q", req.apiReq.QueryParams.Get("access_token"))
	}
	if req.apiReq.QueryParams.Get("company_id") != "" {
		t.Errorf("company_id should be empty, got %q", req.apiReq.QueryParams.Get("company_id"))
	}
	if req.apiReq.QueryParams.Get("timestamp") != "" {
		t.Errorf("timestamp should be empty, got %q", req.apiReq.QueryParams.Get("timestamp"))
	}
	if req.apiReq.QueryParams.Get("sign") != "" {
		t.Errorf("sign should be empty, got %q", req.apiReq.QueryParams.Get("sign"))
	}
}

// --- ListRegulationApiReply 反序列化测试 ---

func TestListRegulationApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
				"regulation_id": "1125920826148759",
				"regulation_name": "差旅制度",
				"regulation_status": "1",
				"source": 3,
				"scene_type": "2",
				"is_use_quota": 1,
				"is_approve": 2,
				"institution_id": "1125900000000001",
				"city_type": 2,
				"approval_type": 1
			}
		],
		"request_id": "test_request_id"
	}`

	var reply ListRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}

	item := reply.Data[0]
	if item.RegulationId == nil || *item.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", item.RegulationId)
	}
	if item.RegulationName == nil || *item.RegulationName != "差旅制度" {
		t.Errorf("RegulationName = %v, want 差旅制度", item.RegulationName)
	}
	if item.RegulationStatus == nil || *item.RegulationStatus != "1" {
		t.Errorf("RegulationStatus = %v, want 1", item.RegulationStatus)
	}
	if item.Source == nil || *item.Source != 3 {
		t.Errorf("Source = %v, want 3", item.Source)
	}
	if item.SceneType == nil || *item.SceneType != "2" {
		t.Errorf("SceneType = %v, want 2", item.SceneType)
	}
	if item.IsUseQuota == nil || *item.IsUseQuota != 1 {
		t.Errorf("IsUseQuota = %v, want 1", item.IsUseQuota)
	}
	if item.IsApprove == nil || *item.IsApprove != 2 {
		t.Errorf("IsApprove = %v, want 2", item.IsApprove)
	}
	if item.InstitutionId == nil || *item.InstitutionId != "1125900000000001" {
		t.Errorf("InstitutionId = %v, want 1125900000000001", item.InstitutionId)
	}
	if item.CityType == nil || *item.CityType != 2 {
		t.Errorf("CityType = %v, want 2", item.CityType)
	}
	if item.ApprovalType == nil || *item.ApprovalType != 1 {
		t.Errorf("ApprovalType = %v, want 1", item.ApprovalType)
	}
}

func TestListRegulationApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply ListRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestListRegulationApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"regulation_id": "1001", "regulation_name": "差旅制度", "source": 3},
			{"regulation_id": "1002", "regulation_name": "通用规则"},
			{"regulation_id": "1003"}
		],
		"request_id": "req_multi"
	}`

	var reply ListRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 3 {
		t.Fatalf("Data len = %d, want 3", len(reply.Data))
	}
	// 第1条有 source
	if reply.Data[0].Source == nil || *reply.Data[0].Source != 3 {
		t.Errorf("Data[0].Source = %v, want 3", reply.Data[0].Source)
	}
	// 第2条无 source
	if reply.Data[1].Source != nil {
		t.Errorf("Data[1].Source = %v, want nil", reply.Data[1].Source)
	}
	// 第3条只有 regulation_id
	if reply.Data[2].RegulationId == nil || *reply.Data[2].RegulationId != "1003" {
		t.Errorf("Data[2].RegulationId = %v, want 1003", reply.Data[2].RegulationId)
	}
	if reply.Data[2].RegulationName != nil {
		t.Errorf("Data[2].RegulationName = %v, want nil", reply.Data[2].RegulationName)
	}
}

func TestListRegulationApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListRegulationApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListRegulationApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- GetRegulation 资源方法测试 ---

func newRegulationTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestGetRegulation_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Regulation/detail" {
			t.Errorf("expected path /river/Regulation/detail, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("regulation_id") != "1125920826148759" {
			t.Errorf("regulation_id = %q, want 1125920826148759", r.URL.Query().Get("regulation_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"regulation_id": "1125920826148759",
				"regulation_name": "差旅制度",
				"regulation_status": "1",
				"source": 3,
				"scene_type": "2",
				"is_use_quota": 1,
				"is_approve": 2,
				"institution_id": "1125900000000001",
				"city_type": 2,
				"approval_type": 1
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		RegulationId("1125920826148759").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetRegulationApiReply.Errno)
	}
	if resp.GetRegulationApiReply.Data.RegulationId == nil || *resp.GetRegulationApiReply.Data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", resp.GetRegulationApiReply.Data.RegulationId)
	}
	if resp.GetRegulationApiReply.Data.RegulationName == nil || *resp.GetRegulationApiReply.Data.RegulationName != "差旅制度" {
		t.Errorf("RegulationName = %v, want 差旅制度", resp.GetRegulationApiReply.Data.RegulationName)
	}
}

func TestGetRegulation_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		RegulationId("1125920826148759").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Data.RegulationId != nil {
		t.Errorf("RegulationId = %v, want nil", resp.GetRegulationApiReply.Data.RegulationId)
	}
}

func TestGetRegulation_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.GetRegulationApiReply.Errno)
	}
}

func TestGetRegulation_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetRegulationApiReply != nil {
		t.Errorf("GetRegulationApiReply should be nil for non-200 response")
	}
}

func TestGetRegulation_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"regulation_id":"1125920826148759","regulation_name":"差旅制度","source":3},"request_id":"req_enc"}`
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

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		RegulationId("1125920826148759").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetRegulationApiReply.Errno)
	}
	if resp.GetRegulationApiReply.Data.RegulationId == nil || *resp.GetRegulationApiReply.Data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", resp.GetRegulationApiReply.Data.RegulationId)
	}
	if resp.GetRegulationApiReply.Data.Source == nil || *resp.GetRegulationApiReply.Data.Source != 3 {
		t.Errorf("Source = %v, want 3", resp.GetRegulationApiReply.Data.Source)
	}
}

func TestGetRegulation_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"regulation_id":"1125920826148759","regulation_name":"差旅制度"},"request_id":"req_enc256"}`
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

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Data.RegulationId == nil || *resp.GetRegulationApiReply.Data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", resp.GetRegulationApiReply.Data.RegulationId)
	}
}

func TestGetRegulation_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"regulation_id":"1125920826148759"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	reg := &regulation{option: option}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
	if resp.GetRegulationApiReply.Data.RegulationId == nil || *resp.GetRegulationApiReply.Data.RegulationId != "1125920826148759" {
		t.Errorf("RegulationId = %v, want 1125920826148759", resp.GetRegulationApiReply.Data.RegulationId)
	}
}

func TestGetRegulation_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"regulation_id":"1125920826148759"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.GetRegulation(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetRegulation() error = %v", err)
	}
	if resp.GetRegulationApiReply == nil {
		t.Fatal("GetRegulationApiReply is nil")
	}
}

// --- ListRegulation 资源方法测试 ---

func TestListRegulation_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Regulation/get" {
			t.Errorf("expected path /river/Regulation/get, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
				{
					"regulation_id": "1125920826148759",
					"regulation_name": "差旅制度",
					"regulation_status": "1",
					"source": 3,
					"scene_type": "2",
					"is_use_quota": 1,
					"is_approve": 2,
					"institution_id": "1125900000000001",
					"city_type": 2,
					"approval_type": 1
				},
				{
					"regulation_id": "1125920826148760",
					"regulation_name": "通用规则",
					"source": 1
				}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if resp.ListRegulationApiReply.Errno == nil || *resp.ListRegulationApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListRegulationApiReply.Errno)
	}
	if len(resp.ListRegulationApiReply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(resp.ListRegulationApiReply.Data))
	}
	if resp.ListRegulationApiReply.Data[0].RegulationId == nil || *resp.ListRegulationApiReply.Data[0].RegulationId != "1125920826148759" {
		t.Errorf("Data[0].RegulationId = %v, want 1125920826148759", resp.ListRegulationApiReply.Data[0].RegulationId)
	}
	if resp.ListRegulationApiReply.Data[1].RegulationName == nil || *resp.ListRegulationApiReply.Data[1].RegulationName != "通用规则" {
		t.Errorf("Data[1].RegulationName = %v, want 通用规则", resp.ListRegulationApiReply.Data[1].RegulationName)
	}
}

func TestListRegulation_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if len(resp.ListRegulationApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListRegulationApiReply.Data))
	}
}

func TestListRegulation_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if resp.ListRegulationApiReply.Errno == nil || *resp.ListRegulationApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.ListRegulationApiReply.Errno)
	}
}

func TestListRegulation_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.ListRegulationApiReply != nil {
		t.Errorf("ListRegulationApiReply should be nil for non-200 response")
	}
}

func TestListRegulation_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"regulation_id":"1125920826148759","regulation_name":"差旅制度","source":3}],"request_id":"req_enc"}`
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

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if resp.ListRegulationApiReply.Errno == nil || *resp.ListRegulationApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListRegulationApiReply.Errno)
	}
	if len(resp.ListRegulationApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRegulationApiReply.Data))
	}
	if resp.ListRegulationApiReply.Data[0].RegulationId == nil || *resp.ListRegulationApiReply.Data[0].RegulationId != "1125920826148759" {
		t.Errorf("Data[0].RegulationId = %v, want 1125920826148759", resp.ListRegulationApiReply.Data[0].RegulationId)
	}
}

func TestListRegulation_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"regulation_id":"1125920826148759","regulation_name":"差旅制度"}],"request_id":"req_enc256"}`
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

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if len(resp.ListRegulationApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRegulationApiReply.Data))
	}
	if resp.ListRegulationApiReply.Data[0].RegulationId == nil || *resp.ListRegulationApiReply.Data[0].RegulationId != "1125920826148759" {
		t.Errorf("Data[0].RegulationId = %v, want 1125920826148759", resp.ListRegulationApiReply.Data[0].RegulationId)
	}
}

func TestListRegulation_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"regulation_id":"1125920826148759"}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	reg := &regulation{option: option}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
	if len(resp.ListRegulationApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRegulationApiReply.Data))
	}
}

func TestListRegulation_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"regulation_id":"1125920826148759"}],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRegulationTestOption(testServer.URL)
	reg := &regulation{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewListRegulationApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := reg.ListRegulation(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListRegulation() error = %v", err)
	}
	if resp.ListRegulationApiReply == nil {
		t.Fatal("ListRegulationApiReply is nil")
	}
}
