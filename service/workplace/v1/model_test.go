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

// --- WorkplaceInfo Builder 测试 ---

func TestWorkplaceInfoBuilder_FullFields(t *testing.T) {
	info := NewWorkplaceInfoBuilder().
		Id("1125922289295589").
		CityId(1).
		OutAddressId("out_addr_001").
		Address("朝阳区望京SOHO").
		Lng("116.481444").
		Lat("39.998756").
		IsWorkplace(1).
		PointRange(500).
		Remark("总部职场").
		Build()

	if info.Id == nil || *info.Id != "1125922289295589" {
		t.Errorf("Id = %v, want 1125922289295589", info.Id)
	}
	if info.CityId == nil || *info.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info.CityId)
	}
	if info.OutAddressId == nil || *info.OutAddressId != "out_addr_001" {
		t.Errorf("OutAddressId = %v, want out_addr_001", info.OutAddressId)
	}
	if info.Address == nil || *info.Address != "朝阳区望京SOHO" {
		t.Errorf("Address = %v, want 朝阳区望京SOHO", info.Address)
	}
	if info.Lng == nil || *info.Lng != "116.481444" {
		t.Errorf("Lng = %v, want 116.481444", info.Lng)
	}
	if info.Lat == nil || *info.Lat != "39.998756" {
		t.Errorf("Lat = %v, want 39.998756", info.Lat)
	}
	if info.IsWorkplace == nil || *info.IsWorkplace != 1 {
		t.Errorf("IsWorkplace = %v, want 1", info.IsWorkplace)
	}
	if info.PointRange == nil || *info.PointRange != 500 {
		t.Errorf("PointRange = %v, want 500", info.PointRange)
	}
	if info.Remark == nil || *info.Remark != "总部职场" {
		t.Errorf("Remark = %v, want 总部职场", info.Remark)
	}
}

func TestWorkplaceInfoBuilder_PartialFields(t *testing.T) {
	info := NewWorkplaceInfoBuilder().
		Id("1125922289295589").
		Address("朝阳区望京SOHO").
		Build()

	if info.Id == nil || *info.Id != "1125922289295589" {
		t.Errorf("Id = %v, want 1125922289295589", info.Id)
	}
	if info.Address == nil || *info.Address != "朝阳区望京SOHO" {
		t.Errorf("Address = %v, want 朝阳区望京SOHO", info.Address)
	}
	// 未设置的字段应为 nil
	if info.CityId != nil {
		t.Errorf("CityId = %v, want nil", info.CityId)
	}
	if info.OutAddressId != nil {
		t.Errorf("OutAddressId = %v, want nil", info.OutAddressId)
	}
	if info.Lng != nil {
		t.Errorf("Lng = %v, want nil", info.Lng)
	}
	if info.Lat != nil {
		t.Errorf("Lat = %v, want nil", info.Lat)
	}
	if info.IsWorkplace != nil {
		t.Errorf("IsWorkplace = %v, want nil", info.IsWorkplace)
	}
	if info.PointRange != nil {
		t.Errorf("PointRange = %v, want nil", info.PointRange)
	}
	if info.Remark != nil {
		t.Errorf("Remark = %v, want nil", info.Remark)
	}
}

func TestWorkplaceInfoBuilder_ZeroIntValues(t *testing.T) {
	info := NewWorkplaceInfoBuilder().
		CityId(0).
		IsWorkplace(0).
		PointRange(0).
		Build()

	// int32 零值也应被设置
	if info.CityId == nil || *info.CityId != 0 {
		t.Errorf("CityId = %v, want 0", info.CityId)
	}
	if info.IsWorkplace == nil || *info.IsWorkplace != 0 {
		t.Errorf("IsWorkplace = %v, want 0", info.IsWorkplace)
	}
	if info.PointRange == nil || *info.PointRange != 0 {
		t.Errorf("PointRange = %v, want 0", info.PointRange)
	}
}

// --- CreateWorkplaceRequestBuilder 测试 ---

func TestCreateWorkplaceRequestBuilder_FullFields(t *testing.T) {
	req := NewCreateWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"city_id":1,"address":"望京SOHO"}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"city_id":1,"address":"望京SOHO"}` {
		t.Errorf("ParamJson = %v, want json", req.ParamJson)
	}
}

func TestCreateWorkplaceRequestBuilder_PartialFields(t *testing.T) {
	req := NewCreateWorkplaceRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"city_id":1}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"city_id":1}` {
		t.Errorf("ParamJson = %v, want {\"city_id\":1}", req.ParamJson)
	}
	// 未设置的字段应为 nil
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", req.CompanyId)
	}
	if req.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", req.Timestamp)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestCreateWorkplaceRequestBuilder_ZeroTimestamp(t *testing.T) {
	req := NewCreateWorkplaceRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	// int64 零值也应被设置
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestCreateWorkplaceRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewCreateWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	// 业务参数不应存在
	if req.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", req.ParamJson)
	}
}

func TestCreateWorkplaceRequestBuilder_JSON(t *testing.T) {
	req := NewCreateWorkplaceRequestBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		ParamJson(`{"city_id":1}`).
		Build()

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal to map failed: %v", err)
	}

	if v, ok := m["client_id"].(string); !ok || v != "test_client" {
		t.Errorf("client_id = %v, want test_client", m["client_id"])
	}
	if v, ok := m["company_id"].(string); !ok || v != "test_company" {
		t.Errorf("company_id = %v, want test_company", m["company_id"])
	}
	if v, ok := m["param_json"].(string); !ok || v != `{"city_id":1}` {
		t.Errorf("param_json = %v, want {\"city_id\":1}", m["param_json"])
	}
	// 未设置的字段不应出现（omitempty）
	if _, ok := m["access_token"]; ok {
		t.Errorf("access_token should be omitted, got %v", m["access_token"])
	}
	if _, ok := m["sign"]; ok {
		t.Errorf("sign should be omitted, got %v", m["sign"])
	}
}

// --- CreateWorkplaceApiReply 反序列化测试 ---

func TestCreateWorkplaceApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {"id": "1125922289295589"},
		"request_id": "test_request_id"
	}`

	var reply CreateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %q, want test_request_id", reply.RequestId)
	}
	if reply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", reply.Data.Id)
	}
}

func TestCreateWorkplaceApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply CreateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", reply.Errmsg)
	}
}

func TestCreateWorkplaceApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_empty"}`

	var reply CreateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", reply.Data.Id)
	}
}

func TestCreateWorkplaceApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

// --- DeleteWorkplaceRequestBuilder 测试 ---

func TestDeleteWorkplaceRequestBuilder_FullFields(t *testing.T) {
	req := NewDeleteWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"id":"1125922289295589"}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"id":"1125922289295589"}` {
		t.Errorf("ParamJson = %v, want json", req.ParamJson)
	}
}

func TestDeleteWorkplaceRequestBuilder_PartialFields(t *testing.T) {
	req := NewDeleteWorkplaceRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"id":"1125922289295589"}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"id":"1125922289295589"}` {
		t.Errorf("ParamJson = %v, want {\"id\":\"1125922289295589\"}", req.ParamJson)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", req.CompanyId)
	}
	if req.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", req.Timestamp)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestDeleteWorkplaceRequestBuilder_ZeroTimestamp(t *testing.T) {
	req := NewDeleteWorkplaceRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestDeleteWorkplaceRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewDeleteWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", req.ParamJson)
	}
}

// --- DeleteWorkplaceApiReply 反序列化测试 ---

func TestDeleteWorkplaceApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": null,
		"request_id": "test_request_id"
	}`

	var reply DeleteWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %v, want test_request_id", reply.RequestId)
	}
}

func TestDeleteWorkplaceApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply DeleteWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", reply.Errmsg)
	}
}

func TestDeleteWorkplaceApiReply_MissingRequestId(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null}`

	var reply DeleteWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.RequestId != nil {
		t.Errorf("RequestId = %v, want nil", reply.RequestId)
	}
}

// --- UpdateWorkplaceRequestBuilder 测试 ---

func TestUpdateWorkplaceRequestBuilder_FullFields(t *testing.T) {
	req := NewUpdateWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"id":"1125922289295589","address":"望京SOHO"}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"id":"1125922289295589","address":"望京SOHO"}` {
		t.Errorf("ParamJson = %v, want json", req.ParamJson)
	}
}

func TestUpdateWorkplaceRequestBuilder_PartialFields(t *testing.T) {
	req := NewUpdateWorkplaceRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"id":"1125922289295589"}`).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ParamJson == nil || *req.ParamJson != `{"id":"1125922289295589"}` {
		t.Errorf("ParamJson = %v, want {\"id\":\"1125922289295589\"}", req.ParamJson)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", req.CompanyId)
	}
	if req.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", req.Timestamp)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestUpdateWorkplaceRequestBuilder_ZeroTimestamp(t *testing.T) {
	req := NewUpdateWorkplaceRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestUpdateWorkplaceRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewUpdateWorkplaceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", req.ParamJson)
	}
}

// --- UpdateWorkplaceApiReply 反序列化测试 ---

func TestUpdateWorkplaceApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": null,
		"request_id": "test_request_id"
	}`

	var reply UpdateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %q, want test_request_id", reply.RequestId)
	}
}

func TestUpdateWorkplaceApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply UpdateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", reply.Errmsg)
	}
}

func TestUpdateWorkplaceApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`

	var reply UpdateWorkplaceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
}

// --- 资源方法测试辅助函数 ---

func newWorkplaceTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// --- CreateWorkplace 资源方法测试 ---

func TestCreateWorkplace_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/workplace/create" {
			t.Errorf("expected path /open-apis/v1/workplace/create, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			ParamJson(`{"city_id":1,"address":"望京SOHO"}`).
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreateWorkplaceApiReply.Errno)
	}
	if resp.CreateWorkplaceApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.CreateWorkplaceApiReply.Errmsg)
	}
	if resp.CreateWorkplaceApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateWorkplaceApiReply.Data.Id)
	}
}

func TestCreateWorkplace_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", resp.CreateWorkplaceApiReply.Data.Id)
	}
}

func TestCreateWorkplace_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.CreateWorkplaceApiReply.Errno)
	}
	if resp.CreateWorkplaceApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.CreateWorkplaceApiReply.Errmsg)
	}
}

func TestCreateWorkplace_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.CreateWorkplaceApiReply != nil {
		t.Errorf("CreateWorkplaceApiReply should be nil for non-200 response")
	}
}

func TestCreateWorkplace_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_enc"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreateWorkplaceApiReply.Errno)
	}
	if resp.CreateWorkplaceApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateWorkplaceApiReply.Data.Id)
	}
}

func TestCreateWorkplace_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_enc256"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateWorkplaceApiReply.Data.Id)
	}
}

func TestCreateWorkplace_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	wp := &workplace{option: option}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
	if resp.CreateWorkplaceApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateWorkplaceApiReply.Data.Id)
	}
}

func TestCreateWorkplace_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(NewCreateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.CreateWorkplace(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateWorkplace() error = %v", err)
	}
	if resp.CreateWorkplaceApiReply == nil {
		t.Fatal("CreateWorkplaceApiReply is nil")
	}
}

// --- DeleteWorkplace 资源方法测试 ---

func TestDeleteWorkplace_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/workplace/del" {
			t.Errorf("expected path /open-apis/v1/workplace/del, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			ParamJson(`{"id":"1125922289295589"}`).
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DeleteWorkplaceApiReply.Errno)
	}
	if resp.DeleteWorkplaceApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.DeleteWorkplaceApiReply.Errmsg)
	}
	if resp.DeleteWorkplaceApiReply.RequestId == nil || *resp.DeleteWorkplaceApiReply.RequestId != "req_001" {
		t.Errorf("RequestId = %v, want req_001", resp.DeleteWorkplaceApiReply.RequestId)
	}
}

func TestDeleteWorkplace_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DeleteWorkplaceApiReply.Errno)
	}
}

func TestDeleteWorkplace_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.DeleteWorkplaceApiReply.Errno)
	}
	if resp.DeleteWorkplaceApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.DeleteWorkplaceApiReply.Errmsg)
	}
}

func TestDeleteWorkplace_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.DeleteWorkplaceApiReply != nil {
		t.Errorf("DeleteWorkplaceApiReply should be nil for non-200 response")
	}
}

func TestDeleteWorkplace_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DeleteWorkplaceApiReply.Errno)
	}
}

func TestDeleteWorkplace_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc256"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DeleteWorkplaceApiReply.Errno)
	}
}

func TestDeleteWorkplace_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	wp := &workplace{option: option}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
	if resp.DeleteWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DeleteWorkplaceApiReply.Errno)
	}
}

func TestDeleteWorkplace_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(NewDeleteWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.DeleteWorkplace(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("DeleteWorkplace() error = %v", err)
	}
	if resp.DeleteWorkplaceApiReply == nil {
		t.Fatal("DeleteWorkplaceApiReply is nil")
	}
}

// --- UpdateWorkplace 资源方法测试 ---

func TestUpdateWorkplace_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/workplace/update" {
			t.Errorf("expected path /open-apis/v1/workplace/update, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			ParamJson(`{"id":"1125922289295589","address":"望京SOHO"}`).
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateWorkplaceApiReply.Errno)
	}
	if resp.UpdateWorkplaceApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.UpdateWorkplaceApiReply.Errmsg)
	}
	if resp.UpdateWorkplaceApiReply.RequestId != "req_001" {
		t.Errorf("RequestId = %q, want req_001", resp.UpdateWorkplaceApiReply.RequestId)
	}
}

func TestUpdateWorkplace_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateWorkplaceApiReply.Errno)
	}
}

func TestUpdateWorkplace_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.UpdateWorkplaceApiReply.Errno)
	}
	if resp.UpdateWorkplaceApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.UpdateWorkplaceApiReply.Errmsg)
	}
}

func TestUpdateWorkplace_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateWorkplaceApiReply != nil {
		t.Errorf("UpdateWorkplaceApiReply should be nil for non-200 response")
	}
}

func TestUpdateWorkplace_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateWorkplaceApiReply.Errno)
	}
}

func TestUpdateWorkplace_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc256"}`
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

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateWorkplaceApiReply.Errno)
	}
}

func TestUpdateWorkplace_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	wp := &workplace{option: option}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
	if resp.UpdateWorkplaceApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateWorkplaceApiReply.Errno)
	}
}

func TestUpdateWorkplace_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newWorkplaceTestOption(testServer.URL)
	wp := &workplace{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(NewUpdateWorkplaceRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := wp.UpdateWorkplace(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateWorkplace() error = %v", err)
	}
	if resp.UpdateWorkplaceApiReply == nil {
		t.Fatal("UpdateWorkplaceApiReply is nil")
	}
}
