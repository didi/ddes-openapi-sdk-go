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

// --- RankInfo Builder 测试 ---

func TestRankInfoBuilder(t *testing.T) {
	info := NewRankInfoBuilder().
		RankId("1125922289295589").
		OutRankId("OUT_RANK_001").
		Name("P7").
		Build()

	if info.RankId == nil || *info.RankId != "1125922289295589" {
		t.Errorf("RankId = %v, want 1125922289295589", info.RankId)
	}
	if info.OutRankId == nil || *info.OutRankId != "OUT_RANK_001" {
		t.Errorf("OutRankId = %v, want OUT_RANK_001", info.OutRankId)
	}
	if info.Name == nil || *info.Name != "P7" {
		t.Errorf("Name = %v, want P7", info.Name)
	}

	// 部分设置
	info2 := NewRankInfoBuilder().
		Name("P6").
		Build()

	if info2.Name == nil || *info2.Name != "P6" {
		t.Errorf("Name = %v, want P6", info2.Name)
	}
	if info2.RankId != nil {
		t.Errorf("RankId = %v, want nil", info2.RankId)
	}
	if info2.OutRankId != nil {
		t.Errorf("OutRankId = %v, want nil", info2.OutRankId)
	}
}

// --- CreateRankRequestBuilder 测试 ---

func TestCreateRankRequestBuilder_FullParams(t *testing.T) {
	rankInfo := NewRankInfoBuilder().
		RankId("1125922289295589").
		OutRankId("OUT_RANK_001").
		Name("P7").
		Build()

	request := NewCreateRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"name":"P7"}`).
		ParamJsonObj(*rankInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"name":"P7"}` {
		t.Errorf("ParamJson = %v, want {\"name\":\"P7\"}", request.ParamJson)
	}
	if request.ParamJsonObj == nil || request.ParamJsonObj.Name == nil || *request.ParamJsonObj.Name != "P7" {
		t.Errorf("ParamJsonObj.Name = %v, want P7", request.ParamJsonObj)
	}
}

func TestCreateRankRequestBuilder_PartialParams(t *testing.T) {
	request := NewCreateRankRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"name":"P6"}`).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"name":"P6"}` {
		t.Errorf("ParamJson = %v, want {\"name\":\"P6\"}", request.ParamJson)
	}
	// 未设置的参数应为 nil
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestCreateRankRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewCreateRankRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	// int64 零值也应该被设置
	if request.Timestamp == nil || *request.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", request.Timestamp)
	}
}

func TestCreateRankRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewCreateRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	// 业务参数不应存在
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

// --- CreateRankApiReply 反序列化测试 ---

func TestCreateRankApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": "1125922289295589"
		},
		"request_id": "req_001"
	}`

	var reply CreateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId != "req_001" {
		t.Errorf("RequestId = %q, want req_001", reply.RequestId)
	}
	if reply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", reply.Data.Id)
	}
}

func TestCreateRankApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply CreateRankApiReply
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

func TestCreateRankApiReply_PartialFields(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_partial"}`

	var reply CreateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", reply.Data.Id)
	}
}

func TestCreateRankApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_empty"}`

	var reply CreateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", reply.Data.Id)
	}
}

func TestCreateRankApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	// data 字段缺失，Id 为零值
	if reply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", reply.Data.Id)
	}
}

// --- DelRankRequestBuilder 测试 ---

func TestDelRankRequestBuilder_FullParams(t *testing.T) {
	rankInfo := NewRankInfoBuilder().
		RankId("1125922289295589").
		Build()

	request := NewDelRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		ParamJson(`{"rank_id":"1125922289295589"}`).
		Sign("test_sign").
		ParamJsonObj(*rankInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"rank_id":"1125922289295589"}` {
		t.Errorf("ParamJson = %v, want {\"rank_id\":\"1125922289295589\"}", request.ParamJson)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJsonObj == nil || request.ParamJsonObj.RankId == nil || *request.ParamJsonObj.RankId != "1125922289295589" {
		t.Errorf("ParamJsonObj.RankId = %v, want 1125922289295589", request.ParamJsonObj)
	}
}

func TestDelRankRequestBuilder_PartialParams(t *testing.T) {
	request := NewDelRankRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"rank_id":"1125922289295589"}`).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"rank_id":"1125922289295589"}` {
		t.Errorf("ParamJson = %v, want {\"rank_id\":\"1125922289295589\"}", request.ParamJson)
	}
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
}

func TestDelRankRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewDelRankRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if request.Timestamp == nil || *request.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", request.Timestamp)
	}
}

func TestDelRankRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewDelRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

// --- DelRankApiReply 反序列化测试 ---

func TestDelRankApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": "1125922289295589"
		},
		"request_id": "req_001"
	}`

	var reply DelRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "req_001" {
		t.Errorf("RequestId = %v, want req_001", reply.RequestId)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Id == nil || *reply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", reply.Data.Id)
	}
}

func TestDelRankApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply DelRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestDelRankApiReply_PartialFields(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_partial"}`

	var reply DelRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	// id 缺失应为 nil
	if reply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", reply.Data.Id)
	}
}

func TestDelRankApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_empty"}`

	var reply DelRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data != nil && reply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", reply.Data.Id)
	}
}

func TestDelRankApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply DelRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- ListRankApiReqBuilder 测试 (GET 型) ---

func TestListRankApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Offset(0).
		Length(10).
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
		{"offset", "0"},
		{"length", "10"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListRankApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	if req.apiReq.QueryParams.Get("length") != "10" {
		t.Errorf("length mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty, got %q", req.apiReq.QueryParams.Get("access_token"))
	}
	if req.apiReq.QueryParams.Get("company_id") != "" {
		t.Errorf("company_id should be empty, got %q", req.apiReq.QueryParams.Get("company_id"))
	}
	if req.apiReq.QueryParams.Get("sign") != "" {
		t.Errorf("sign should be empty, got %q", req.apiReq.QueryParams.Get("sign"))
	}
}

func TestListRankApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
	if req.apiReq.QueryParams.Get("length") != "0" {
		t.Errorf("length = %q, want \"0\"", req.apiReq.QueryParams.Get("length"))
	}
}

func TestListRankApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListRankApiReqBuilder().
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
	// 业务参数不应存在
	for _, key := range []string{"offset", "length"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- ListRankApiReply 反序列化测试 ---

func TestListRankApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{"rank_id": "1125922289295589", "out_rank_id": "OUT_001", "name": "P7"},
				{"rank_id": "1125922289295590", "out_rank_id": "OUT_002", "name": "P6"}
			]
		},
		"request_id": "req_001"
	}`

	var reply ListRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.Data.Total == nil || *reply.Data.Total != 2 {
		t.Errorf("Total = %v, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 2 {
		t.Fatalf("Records len = %d, want 2", len(reply.Data.Records))
	}
	if reply.Data.Records[0].RankId == nil || *reply.Data.Records[0].RankId != "1125922289295589" {
		t.Errorf("Records[0].RankId = %v, want 1125922289295589", reply.Data.Records[0].RankId)
	}
	if reply.Data.Records[0].Name == nil || *reply.Data.Records[0].Name != "P7" {
		t.Errorf("Records[0].Name = %v, want P7", reply.Data.Records[0].Name)
	}
}

func TestListRankApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListRankApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 3,
			"records": [
				{"rank_id": "1001", "name": "P7"},
				{"out_rank_id": "OUT_002"},
				{"rank_id": "1003", "out_rank_id": "OUT_003", "name": "P5"}
			]
		},
		"request_id": "req_multi"
	}`

	var reply ListRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 3 {
		t.Fatalf("Records len = %d, want 3", len(reply.Data.Records))
	}
	// 第1条有 name
	if reply.Data.Records[0].Name == nil || *reply.Data.Records[0].Name != "P7" {
		t.Errorf("Records[0].Name = %v, want P7", reply.Data.Records[0].Name)
	}
	// 第2条缺少 name 和 rank_id
	if reply.Data.Records[1].Name != nil {
		t.Errorf("Records[1].Name = %v, want nil", reply.Data.Records[1].Name)
	}
	if reply.Data.Records[1].RankId != nil {
		t.Errorf("Records[1].RankId = %v, want nil", reply.Data.Records[1].RankId)
	}
	// 第3条全字段
	if reply.Data.Records[2].RankId == nil || *reply.Data.Records[2].RankId != "1003" {
		t.Errorf("Records[2].RankId = %v, want 1003", reply.Data.Records[2].RankId)
	}
}

func TestListRankApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_empty"}`

	var reply ListRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestListRankApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data.Records != nil {
		t.Errorf("Records = %v, want nil", reply.Data.Records)
	}
	if reply.Data.Total != nil {
		t.Errorf("Total = %v, want nil", reply.Data.Total)
	}
}

// --- UpdateRankRequestBuilder 测试 ---

func TestUpdateRankRequestBuilder_FullParams(t *testing.T) {
	rankInfo := NewRankInfoBuilder().
		RankId("1125922289295589").
		Name("P7").
		Build()

	request := NewUpdateRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"rank_id":"1125922289295589","name":"P7"}`).
		ParamJsonObj(*rankInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"rank_id":"1125922289295589","name":"P7"}` {
		t.Errorf("ParamJson = %v, want {\"rank_id\":\"1125922289295589\",\"name\":\"P7\"}", request.ParamJson)
	}
	if request.ParamJsonObj == nil || request.ParamJsonObj.RankId == nil || *request.ParamJsonObj.RankId != "1125922289295589" {
		t.Errorf("ParamJsonObj.RankId = %v, want 1125922289295589", request.ParamJsonObj)
	}
}

func TestUpdateRankRequestBuilder_PartialParams(t *testing.T) {
	request := NewUpdateRankRequestBuilder().
		ClientId("test_client").
		ParamJson(`{"name":"P6"}`).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != `{"name":"P6"}` {
		t.Errorf("ParamJson = %v, want {\"name\":\"P6\"}", request.ParamJson)
	}
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
}

func TestUpdateRankRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewUpdateRankRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if request.Timestamp == nil || *request.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", request.Timestamp)
	}
}

func TestUpdateRankRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewUpdateRankRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

// --- UpdateRankApiReply 反序列化测试 ---

func TestUpdateRankApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": "1125922289295589"
		},
		"request_id": "req_001"
	}`

	var reply UpdateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "req_001" {
		t.Errorf("RequestId = %v, want req_001", reply.RequestId)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Id == nil || *reply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", reply.Data.Id)
	}
}

func TestUpdateRankApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply UpdateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestUpdateRankApiReply_PartialFields(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_partial"}`

	var reply UpdateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", reply.Data.Id)
	}
}

func TestUpdateRankApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_empty"}`

	var reply UpdateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data != nil && reply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", reply.Data.Id)
	}
}

func TestUpdateRankApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply UpdateRankApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- 资源方法测试 ---

func newRankTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// --- CreateRank 资源方法测试 ---

func TestCreateRank_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/rank/create" {
			t.Errorf("expected path /open-apis/v1/rank/create, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreateRankApiReply.Errno)
	}
	if resp.CreateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateRankApiReply.Data.Id)
	}
}

func TestCreateRank_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Data.Id != "" {
		t.Errorf("Data.Id = %q, want empty", resp.CreateRankApiReply.Data.Id)
	}
}

func TestCreateRank_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.CreateRankApiReply.Errno)
	}
}

func TestCreateRank_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.CreateRankApiReply != nil {
		t.Errorf("CreateRankApiReply should be nil for non-200 response")
	}
}

func TestCreateRank_WithEncryption_AES128(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreateRankApiReply.Errno)
	}
	if resp.CreateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateRankApiReply.Data.Id)
	}
}

func TestCreateRank_WithEncryption_AES256(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateRankApiReply.Data.Id)
	}
}

func TestCreateRank_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	rk := &rank{option: option}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
	if resp.CreateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %q, want 1125922289295589", resp.CreateRankApiReply.Data.Id)
	}
}

func TestCreateRank_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreateRankApiReqBuilder().
		CreateRankRequest(NewCreateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.CreateRank(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateRank() error = %v", err)
	}
	if resp.CreateRankApiReply == nil {
		t.Fatal("CreateRankApiReply is nil")
	}
}

// --- DelRank 资源方法测试 ---

func TestDelRank_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/rank/del" {
			t.Errorf("expected path /open-apis/v1/rank/del, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DelRankApiReply.Errno)
	}
	if resp.DelRankApiReply.Data == nil || resp.DelRankApiReply.Data.Id == nil || *resp.DelRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.DelRankApiReply.Data)
	}
}

func TestDelRank_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Data != nil && resp.DelRankApiReply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", resp.DelRankApiReply.Data.Id)
	}
}

func TestDelRank_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.DelRankApiReply.Errno)
	}
}

func TestDelRank_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.DelRankApiReply != nil {
		t.Errorf("DelRankApiReply should be nil for non-200 response")
	}
}

func TestDelRank_WithEncryption_AES128(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.DelRankApiReply.Errno)
	}
	if resp.DelRankApiReply.Data == nil || resp.DelRankApiReply.Data.Id == nil || *resp.DelRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.DelRankApiReply.Data)
	}
}

func TestDelRank_WithEncryption_AES256(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Data == nil || resp.DelRankApiReply.Data.Id == nil || *resp.DelRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.DelRankApiReply.Data)
	}
}

func TestDelRank_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	rk := &rank{option: option}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
	if resp.DelRankApiReply.Data == nil || resp.DelRankApiReply.Data.Id == nil || *resp.DelRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.DelRankApiReply.Data)
	}
}

func TestDelRank_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewDelRankApiReqBuilder().
		DelRankRequest(NewDelRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.DelRank(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("DelRank() error = %v", err)
	}
	if resp.DelRankApiReply == nil {
		t.Fatal("DelRankApiReply is nil")
	}
}

// --- ListRank 资源方法测试 ---

func TestListRank_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Rank/getRanks" {
			t.Errorf("expected path /river/Rank/getRanks, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":2,"records":[{"rank_id":"1125922289295589","out_rank_id":"OUT_001","name":"P7"},{"rank_id":"1125922289295590","out_rank_id":"OUT_002","name":"P6"}]},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Offset(0).
		Length(10).
		Sign("test_sign").
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if resp.ListRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListRankApiReply.Errno)
	}
	if resp.ListRankApiReply.Data.Total == nil || *resp.ListRankApiReply.Data.Total != 2 {
		t.Errorf("Total = %v, want 2", resp.ListRankApiReply.Data.Total)
	}
	if len(resp.ListRankApiReply.Data.Records) != 2 {
		t.Fatalf("Records len = %d, want 2", len(resp.ListRankApiReply.Data.Records))
	}
	if resp.ListRankApiReply.Data.Records[0].RankId == nil || *resp.ListRankApiReply.Data.Records[0].RankId != "1125922289295589" {
		t.Errorf("Records[0].RankId = %v, want 1125922289295589", resp.ListRankApiReply.Data.Records[0].RankId)
	}
}

func TestListRank_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if len(resp.ListRankApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.ListRankApiReply.Data.Records))
	}
}

func TestListRank_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if resp.ListRankApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListRankApiReply.Errno)
	}
}

func TestListRank_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListRankApiReply != nil {
		t.Errorf("ListRankApiReply should be nil for non-200 response")
	}
}

func TestListRank_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"rank_id":"1125922289295589","name":"P7"}]},"request_id":"req_enc"}`
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if resp.ListRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListRankApiReply.Errno)
	}
	if len(resp.ListRankApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListRankApiReply.Data.Records))
	}
}

func TestListRank_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"rank_id":"1125922289295589"}]},"request_id":"req_enc256"}`
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if len(resp.ListRankApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListRankApiReply.Data.Records))
	}
}

func TestListRank_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"rank_id":"1125922289295589"}]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	rk := &rank{option: option}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := rk.ListRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
	if len(resp.ListRankApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListRankApiReply.Data.Records))
	}
}

func TestListRank_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewListRankApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := rk.ListRank(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListRank() error = %v", err)
	}
	if resp.ListRankApiReply == nil {
		t.Fatal("ListRankApiReply is nil")
	}
}

// --- UpdateRank 资源方法测试 ---

func TestUpdateRank_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/rank/update" {
			t.Errorf("expected path /open-apis/v1/rank/update, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateRankApiReply.Errno)
	}
	if resp.UpdateRankApiReply.Data == nil || resp.UpdateRankApiReply.Data.Id == nil || *resp.UpdateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.UpdateRankApiReply.Data)
	}
}

func TestUpdateRank_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Data != nil && resp.UpdateRankApiReply.Data.Id != nil {
		t.Errorf("Data.Id = %v, want nil", resp.UpdateRankApiReply.Data.Id)
	}
}

func TestUpdateRank_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.UpdateRankApiReply.Errno)
	}
}

func TestUpdateRank_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateRankApiReply != nil {
		t.Errorf("UpdateRankApiReply should be nil for non-200 response")
	}
}

func TestUpdateRank_WithEncryption_AES128(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateRankApiReply.Errno)
	}
	if resp.UpdateRankApiReply.Data == nil || resp.UpdateRankApiReply.Data.Id == nil || *resp.UpdateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.UpdateRankApiReply.Data)
	}
}

func TestUpdateRank_WithEncryption_AES256(t *testing.T) {
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

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Data == nil || resp.UpdateRankApiReply.Data.Id == nil || *resp.UpdateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.UpdateRankApiReply.Data)
	}
}

func TestUpdateRank_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	rk := &rank{option: option}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
	if resp.UpdateRankApiReply.Data == nil || resp.UpdateRankApiReply.Data.Id == nil || *resp.UpdateRankApiReply.Data.Id != "1125922289295589" {
		t.Errorf("Data.Id = %v, want 1125922289295589", resp.UpdateRankApiReply.Data)
	}
}

func TestUpdateRank_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125922289295589"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRankTestOption(testServer.URL)
	rk := &rank{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateRankApiReqBuilder().
		UpdateRankRequest(NewUpdateRankRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := rk.UpdateRank(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateRank() error = %v", err)
	}
	if resp.UpdateRankApiReply == nil {
		t.Fatal("UpdateRankApiReply is nil")
	}
}
