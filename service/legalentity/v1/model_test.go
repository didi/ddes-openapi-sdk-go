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

// --- LegalEntityRecord Builder 测试 ---

func TestLegalEntityRecordBuilder_FullFields(t *testing.T) {
	record := NewLegalEntityRecordBuilder().
		Address("北京市朝阳区望京SOHO").
		BankCardNo("6222020200112345678").
		EndTime("2099-12-31 23:59:59").
		LegalEntityId(1125922289295589).
		Name("滴滴出行科技有限公司").
		OpenBank("招商银行北京分行").
		OutLegalEntityId("LE_001").
		ParentId(1125915646135311).
		StartTime("2020-01-01 00:00:00").
		TaxpayerNo("91110000MA0012345X").
		Telephone("010-12345678").
		Status(1).
		Build()

	if record.Address == nil || *record.Address != "北京市朝阳区望京SOHO" {
		t.Errorf("Address = %v, want 北京市朝阳区望京SOHO", record.Address)
	}
	if record.BankCardNo == nil || *record.BankCardNo != "6222020200112345678" {
		t.Errorf("BankCardNo = %v, want 6222020200112345678", record.BankCardNo)
	}
	if record.EndTime == nil || *record.EndTime != "2099-12-31 23:59:59" {
		t.Errorf("EndTime = %v, want 2099-12-31 23:59:59", record.EndTime)
	}
	if record.LegalEntityId == nil || *record.LegalEntityId != 1125922289295589 {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", record.LegalEntityId)
	}
	if record.Name == nil || *record.Name != "滴滴出行科技有限公司" {
		t.Errorf("Name = %v, want 滴滴出行科技有限公司", record.Name)
	}
	if record.OpenBank == nil || *record.OpenBank != "招商银行北京分行" {
		t.Errorf("OpenBank = %v, want 招商银行北京分行", record.OpenBank)
	}
	if record.OutLegalEntityId == nil || *record.OutLegalEntityId != "LE_001" {
		t.Errorf("OutLegalEntityId = %v, want LE_001", record.OutLegalEntityId)
	}
	if record.ParentId == nil || *record.ParentId != 1125915646135311 {
		t.Errorf("ParentId = %v, want 1125915646135311", record.ParentId)
	}
	if record.StartTime == nil || *record.StartTime != "2020-01-01 00:00:00" {
		t.Errorf("StartTime = %v, want 2020-01-01 00:00:00", record.StartTime)
	}
	if record.TaxpayerNo == nil || *record.TaxpayerNo != "91110000MA0012345X" {
		t.Errorf("TaxpayerNo = %v, want 91110000MA0012345X", record.TaxpayerNo)
	}
	if record.Telephone == nil || *record.Telephone != "010-12345678" {
		t.Errorf("Telephone = %v, want 010-12345678", record.Telephone)
	}
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Status = %v, want 1", record.Status)
	}
}

func TestLegalEntityRecordBuilder_PartialFields(t *testing.T) {
	record := NewLegalEntityRecordBuilder().
		LegalEntityId(1001).
		Name("测试公司").
		Status(0).
		Build()

	if record.LegalEntityId == nil || *record.LegalEntityId != 1001 {
		t.Errorf("LegalEntityId = %v, want 1001", record.LegalEntityId)
	}
	if record.Name == nil || *record.Name != "测试公司" {
		t.Errorf("Name = %v, want 测试公司", record.Name)
	}
	if record.Status == nil || *record.Status != 0 {
		t.Errorf("Status = %v, want 0", record.Status)
	}
	// 未设置的字段应为 nil
	if record.Address != nil {
		t.Errorf("Address = %v, want nil", record.Address)
	}
	if record.BankCardNo != nil {
		t.Errorf("BankCardNo = %v, want nil", record.BankCardNo)
	}
	if record.EndTime != nil {
		t.Errorf("EndTime = %v, want nil", record.EndTime)
	}
	if record.OpenBank != nil {
		t.Errorf("OpenBank = %v, want nil", record.OpenBank)
	}
	if record.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", record.OutLegalEntityId)
	}
	if record.ParentId != nil {
		t.Errorf("ParentId = %v, want nil", record.ParentId)
	}
	if record.StartTime != nil {
		t.Errorf("StartTime = %v, want nil", record.StartTime)
	}
	if record.TaxpayerNo != nil {
		t.Errorf("TaxpayerNo = %v, want nil", record.TaxpayerNo)
	}
	if record.Telephone != nil {
		t.Errorf("Telephone = %v, want nil", record.Telephone)
	}
}

// --- CreateLegalEntityRequestBuilder 测试 ---

func TestCreateLegalEntityRequestBuilder_FullParams(t *testing.T) {
	req := NewCreateLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Name("滴滴出行科技有限公司").
		OutLegalEntityId("LE_001").
		ParentId(1001).
		StartTime("2020-01-01 00:00:00").
		EndTime("2099-12-31 23:59:59").
		TaxpayerNo("91110000MA0012345X").
		Address("北京市朝阳区望京SOHO").
		Telephone("010-12345678").
		OpenBank("招商银行北京分行").
		BankCardNo("6222020200112345678").
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
	if req.Name == nil || *req.Name != "滴滴出行科技有限公司" {
		t.Errorf("Name = %v, want 滴滴出行科技有限公司", req.Name)
	}
	if req.OutLegalEntityId == nil || *req.OutLegalEntityId != "LE_001" {
		t.Errorf("OutLegalEntityId = %v, want LE_001", req.OutLegalEntityId)
	}
	if req.ParentId == nil || *req.ParentId != 1001 {
		t.Errorf("ParentId = %v, want 1001", req.ParentId)
	}
	if req.StartTime == nil || *req.StartTime != "2020-01-01 00:00:00" {
		t.Errorf("StartTime = %v, want 2020-01-01 00:00:00", req.StartTime)
	}
	if req.EndTime == nil || *req.EndTime != "2099-12-31 23:59:59" {
		t.Errorf("EndTime = %v, want 2099-12-31 23:59:59", req.EndTime)
	}
	if req.TaxpayerNo == nil || *req.TaxpayerNo != "91110000MA0012345X" {
		t.Errorf("TaxpayerNo = %v, want 91110000MA0012345X", req.TaxpayerNo)
	}
	if req.Address == nil || *req.Address != "北京市朝阳区望京SOHO" {
		t.Errorf("Address = %v, want 北京市朝阳区望京SOHO", req.Address)
	}
	if req.Telephone == nil || *req.Telephone != "010-12345678" {
		t.Errorf("Telephone = %v, want 010-12345678", req.Telephone)
	}
	if req.OpenBank == nil || *req.OpenBank != "招商银行北京分行" {
		t.Errorf("OpenBank = %v, want 招商银行北京分行", req.OpenBank)
	}
	if req.BankCardNo == nil || *req.BankCardNo != "6222020200112345678" {
		t.Errorf("BankCardNo = %v, want 6222020200112345678", req.BankCardNo)
	}
}

func TestCreateLegalEntityRequestBuilder_PartialParams(t *testing.T) {
	req := NewCreateLegalEntityRequestBuilder().
		ClientId("test_client").
		Name("测试公司").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Name == nil || *req.Name != "测试公司" {
		t.Errorf("Name = %v, want 测试公司", req.Name)
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
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
	if req.ParentId != nil {
		t.Errorf("ParentId = %v, want nil", req.ParentId)
	}
	if req.StartTime != nil {
		t.Errorf("StartTime = %v, want nil", req.StartTime)
	}
	if req.EndTime != nil {
		t.Errorf("EndTime = %v, want nil", req.EndTime)
	}
	if req.TaxpayerNo != nil {
		t.Errorf("TaxpayerNo = %v, want nil", req.TaxpayerNo)
	}
	if req.Address != nil {
		t.Errorf("Address = %v, want nil", req.Address)
	}
	if req.Telephone != nil {
		t.Errorf("Telephone = %v, want nil", req.Telephone)
	}
	if req.OpenBank != nil {
		t.Errorf("OpenBank = %v, want nil", req.OpenBank)
	}
	if req.BankCardNo != nil {
		t.Errorf("BankCardNo = %v, want nil", req.BankCardNo)
	}
}

func TestCreateLegalEntityRequestBuilder_ZeroIntValues(t *testing.T) {
	// int32 零值也应该被设置
	req := NewCreateLegalEntityRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		ParentId(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.ParentId == nil || *req.ParentId != 0 {
		t.Errorf("ParentId = %v, want 0", req.ParentId)
	}
}

func TestCreateLegalEntityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewCreateLegalEntityRequestBuilder().
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
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
	if req.ParentId != nil {
		t.Errorf("ParentId = %v, want nil", req.ParentId)
	}
}

// --- CreateLegalEntityApiReqBuilder 测试 ---

func TestCreateLegalEntityApiReqBuilder_FullParams(t *testing.T) {
	requestBody := NewCreateLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Name("测试公司").
		OutLegalEntityId("LE_001").
		Build()

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(requestBody).
		Build()

	// 验证 Body 被设置
	if req.apiReq.Body == nil {
		t.Fatal("Body should not be nil")
	}
	body, ok := req.apiReq.Body.(*CreateLegalEntityRequest)
	if !ok {
		t.Fatal("Body should be *CreateLegalEntityRequest")
	}
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.Name == nil || *body.Name != "测试公司" {
		t.Errorf("Body.Name = %v, want 测试公司", body.Name)
	}
	if body.OutLegalEntityId == nil || *body.OutLegalEntityId != "LE_001" {
		t.Errorf("Body.OutLegalEntityId = %v, want LE_001", body.OutLegalEntityId)
	}
}

func TestCreateLegalEntityApiReqBuilder_OnlyCommonParams(t *testing.T) {
	requestBody := NewCreateLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*CreateLegalEntityRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 业务参数不应存在
	if body.Name != nil {
		t.Errorf("Body.Name = %v, want nil", body.Name)
	}
	if body.OutLegalEntityId != nil {
		t.Errorf("Body.OutLegalEntityId = %v, want nil", body.OutLegalEntityId)
	}
}

// --- CreateLegalEntityApiReply 反序列化测试 ---

func TestCreateLegalEntityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"legal_entity_id": "1125922289295589"
		},
		"request_id": "test_request_id"
	}`

	var reply CreateLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %v, want test_request_id", reply.RequestId)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.LegalEntityId == nil || *reply.Data.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", reply.Data.LegalEntityId)
	}
}

func TestCreateLegalEntityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply CreateLegalEntityApiReply
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

func TestCreateLegalEntityApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"legal_entity_id":""},"request_id":"req_empty"}`

	var reply CreateLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	// JSON 中 legal_entity_id 为空字符串，反序列化后为指向空字符串的指针
	if reply.Data.LegalEntityId == nil || *reply.Data.LegalEntityId != "" {
		t.Errorf("LegalEntityId = %v, want empty string", reply.Data.LegalEntityId)
	}
}

func TestCreateLegalEntityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreateLegalEntityApiReply
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

// --- DelLegalEntityRequestBuilder 测试 ---

func TestDelLegalEntityRequestBuilder_FullParams(t *testing.T) {
	req := NewDelLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		LegalEntityId("1125922289295589").
		OutLegalEntityId("LE_001").
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
	if req.LegalEntityId == nil || *req.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", req.LegalEntityId)
	}
	if req.OutLegalEntityId == nil || *req.OutLegalEntityId != "LE_001" {
		t.Errorf("OutLegalEntityId = %v, want LE_001", req.OutLegalEntityId)
	}
}

func TestDelLegalEntityRequestBuilder_PartialParams(t *testing.T) {
	req := NewDelLegalEntityRequestBuilder().
		ClientId("test_client").
		LegalEntityId("1125922289295589").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.LegalEntityId == nil || *req.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", req.LegalEntityId)
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
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
}

func TestDelLegalEntityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewDelLegalEntityRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestDelLegalEntityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewDelLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	// 业务参数不应存在
	if req.LegalEntityId != nil {
		t.Errorf("LegalEntityId = %v, want nil", req.LegalEntityId)
	}
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
}

// --- DelLegalEntityApiReply 反序列化测试 ---

func TestDelLegalEntityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {},
		"request_id": "test_request_id"
	}`

	var reply DelLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %v, want test_request_id", reply.RequestId)
	}
}

func TestDelLegalEntityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply DelLegalEntityApiReply
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

func TestDelLegalEntityApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply DelLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
}

func TestDelLegalEntityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply DelLegalEntityApiReply
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

// --- GetLegalEntityApiReqBuilder 测试 (GET 型) ---

func TestGetLegalEntityApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Keyword("滴滴").
		LegalEntityId("1125922289295589").
		Offset(0).
		Length(10).
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
		{"keyword", "滴滴"},
		{"legal_entity_id", "1125922289295589"},
		{"offset", "0"},
		{"length", "10"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetLegalEntityApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Keyword("测试").
		Offset(0).
		Length(20).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("keyword") != "测试" {
		t.Errorf("keyword mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	if req.apiReq.QueryParams.Get("length") != "20" {
		t.Errorf("length mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("legal_entity_id") != "" {
		t.Errorf("legal_entity_id should be empty, got %q", req.apiReq.QueryParams.Get("legal_entity_id"))
	}
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty, got %q", req.apiReq.QueryParams.Get("access_token"))
	}
}

func TestGetLegalEntityApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetLegalEntityApiReqBuilder().
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

func TestGetLegalEntityApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetLegalEntityApiReqBuilder().
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
	for _, key := range []string{"keyword", "legal_entity_id", "offset", "length"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- GetLegalEntityApiReply 反序列化测试 ---

func TestGetLegalEntityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{
					"address": "北京市朝阳区望京SOHO",
					"bank_card_no": "6222020200112345678",
					"end_time": "2099-12-31 23:59:59",
					"legal_entity_id": 1125922289295589,
					"name": "滴滴出行科技有限公司",
					"open_bank": "招商银行北京分行",
					"out_legal_entity_id": "LE_001",
					"parent_id": 1125915646135311,
					"start_time": "2020-01-01 00:00:00",
					"taxpayer_no": "91110000MA0012345X",
					"telephone": "010-12345678",
					"status": 1
				},
				{
					"legal_entity_id": 1125915646135311,
					"name": "滴滴集团总部",
					"status": 0
				}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply GetLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total == nil || *reply.Data.Total != 2 {
		t.Errorf("Total = %v, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 2 {
		t.Fatalf("Records len = %d, want 2", len(reply.Data.Records))
	}

	record := reply.Data.Records[0]
	if record.LegalEntityId == nil || *record.LegalEntityId != 1125922289295589 {
		t.Errorf("Records[0].LegalEntityId = %v, want 1125922289295589", record.LegalEntityId)
	}
	if record.Name == nil || *record.Name != "滴滴出行科技有限公司" {
		t.Errorf("Records[0].Name = %v, want 滴滴出行科技有限公司", record.Name)
	}
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Records[0].Status = %v, want 1", record.Status)
	}
	if record.Address == nil || *record.Address != "北京市朝阳区望京SOHO" {
		t.Errorf("Records[0].Address = %v, want 北京市朝阳区望京SOHO", record.Address)
	}
	if record.ParentId == nil || *record.ParentId != 1125915646135311 {
		t.Errorf("Records[0].ParentId = %v, want 1125915646135311", record.ParentId)
	}

	// 第2条部分字段缺失
	record2 := reply.Data.Records[1]
	if record2.LegalEntityId == nil || *record2.LegalEntityId != 1125915646135311 {
		t.Errorf("Records[1].LegalEntityId = %v, want 1125915646135311", record2.LegalEntityId)
	}
	if record2.Status == nil || *record2.Status != 0 {
		t.Errorf("Records[1].Status = %v, want 0", record2.Status)
	}
	// 缺失的字段应为 nil
	if record2.Address != nil {
		t.Errorf("Records[1].Address = %v, want nil", record2.Address)
	}
	if record2.TaxpayerNo != nil {
		t.Errorf("Records[1].TaxpayerNo = %v, want nil", record2.TaxpayerNo)
	}
}

func TestGetLegalEntityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply GetLegalEntityApiReply
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

func TestGetLegalEntityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_empty"}`

	var reply GetLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
	if reply.Data.Total == nil || *reply.Data.Total != 0 {
		t.Errorf("Total = %v, want 0", reply.Data.Total)
	}
}

func TestGetLegalEntityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetLegalEntityApiReply
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

// --- UpdateLegalEntityRequestBuilder 测试 ---

func TestUpdateLegalEntityRequestBuilder_FullParams(t *testing.T) {
	req := NewUpdateLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Name("滴滴出行科技有限公司").
		LegalEntityId("1125922289295589").
		OutLegalEntityId("LE_001").
		ParentId(1001).
		StartTime("2020-01-01 00:00:00").
		EndTime("2099-12-31 23:59:59").
		TaxpayerNo("91110000MA0012345X").
		Address("北京市朝阳区望京SOHO").
		Telephone("010-12345678").
		OpenBank("招商银行北京分行").
		BankCardNo("6222020200112345678").
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
	if req.Name == nil || *req.Name != "滴滴出行科技有限公司" {
		t.Errorf("Name = %v, want 滴滴出行科技有限公司", req.Name)
	}
	if req.LegalEntityId == nil || *req.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", req.LegalEntityId)
	}
	if req.OutLegalEntityId == nil || *req.OutLegalEntityId != "LE_001" {
		t.Errorf("OutLegalEntityId = %v, want LE_001", req.OutLegalEntityId)
	}
	if req.ParentId == nil || *req.ParentId != 1001 {
		t.Errorf("ParentId = %v, want 1001", req.ParentId)
	}
	if req.StartTime == nil || *req.StartTime != "2020-01-01 00:00:00" {
		t.Errorf("StartTime = %v, want 2020-01-01 00:00:00", req.StartTime)
	}
	if req.EndTime == nil || *req.EndTime != "2099-12-31 23:59:59" {
		t.Errorf("EndTime = %v, want 2099-12-31 23:59:59", req.EndTime)
	}
	if req.TaxpayerNo == nil || *req.TaxpayerNo != "91110000MA0012345X" {
		t.Errorf("TaxpayerNo = %v, want 91110000MA0012345X", req.TaxpayerNo)
	}
	if req.Address == nil || *req.Address != "北京市朝阳区望京SOHO" {
		t.Errorf("Address = %v, want 北京市朝阳区望京SOHO", req.Address)
	}
	if req.Telephone == nil || *req.Telephone != "010-12345678" {
		t.Errorf("Telephone = %v, want 010-12345678", req.Telephone)
	}
	if req.OpenBank == nil || *req.OpenBank != "招商银行北京分行" {
		t.Errorf("OpenBank = %v, want 招商银行北京分行", req.OpenBank)
	}
	if req.BankCardNo == nil || *req.BankCardNo != "6222020200112345678" {
		t.Errorf("BankCardNo = %v, want 6222020200112345678", req.BankCardNo)
	}
}

func TestUpdateLegalEntityRequestBuilder_PartialParams(t *testing.T) {
	req := NewUpdateLegalEntityRequestBuilder().
		ClientId("test_client").
		LegalEntityId("1125922289295589").
		Name("更新后的公司名").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.LegalEntityId == nil || *req.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", req.LegalEntityId)
	}
	if req.Name == nil || *req.Name != "更新后的公司名" {
		t.Errorf("Name = %v, want 更新后的公司名", req.Name)
	}
	// 未设置的字段应为 nil
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
	if req.ParentId != nil {
		t.Errorf("ParentId = %v, want nil", req.ParentId)
	}
	if req.TaxpayerNo != nil {
		t.Errorf("TaxpayerNo = %v, want nil", req.TaxpayerNo)
	}
}

func TestUpdateLegalEntityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewUpdateLegalEntityRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		ParentId(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.ParentId == nil || *req.ParentId != 0 {
		t.Errorf("ParentId = %v, want 0", req.ParentId)
	}
}

func TestUpdateLegalEntityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewUpdateLegalEntityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	// 业务参数不应存在
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
	if req.LegalEntityId != nil {
		t.Errorf("LegalEntityId = %v, want nil", req.LegalEntityId)
	}
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
}

// --- UpdateLegalEntityApiReqBuilder 测试 ---

func TestUpdateLegalEntityApiReqBuilder_FullParams(t *testing.T) {
	requestBody := NewUpdateLegalEntityRequestBuilder().
		ClientId("test_client").
		LegalEntityId("1125922289295589").
		Name("更新公司").
		Build()

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(requestBody).
		Build()

	if req.apiReq.Body == nil {
		t.Fatal("Body should not be nil")
	}
	body, ok := req.apiReq.Body.(*UpdateLegalEntityRequest)
	if !ok {
		t.Fatal("Body should be *UpdateLegalEntityRequest")
	}
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.LegalEntityId == nil || *body.LegalEntityId != "1125922289295589" {
		t.Errorf("Body.LegalEntityId = %v, want 1125922289295589", body.LegalEntityId)
	}
	if body.Name == nil || *body.Name != "更新公司" {
		t.Errorf("Body.Name = %v, want 更新公司", body.Name)
	}
}

// --- UpdateLegalEntityApiReply 反序列化测试 ---

func TestUpdateLegalEntityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {},
		"request_id": "test_request_id"
	}`

	var reply UpdateLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %v, want test_request_id", reply.RequestId)
	}
}

func TestUpdateLegalEntityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply UpdateLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestUpdateLegalEntityApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply UpdateLegalEntityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
}

func TestUpdateLegalEntityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply UpdateLegalEntityApiReply
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

// --- 资源方法测试 ---

func newLegalEntityTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// === CreateLegalEntity 资源方法测试 ===

func TestCreateLegalEntity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/LegalEntity/add" {
			t.Errorf("expected path /river/LegalEntity/add, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"legal_entity_id": "1125922289295589"
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			Name("滴滴出行科技有限公司").
			OutLegalEntityId("LE_001").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Errno == nil || *resp.CreateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateLegalEntityApiReply.Errno)
	}
	if resp.CreateLegalEntityApiReply.Errmsg == nil || *resp.CreateLegalEntityApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", resp.CreateLegalEntityApiReply.Errmsg)
	}
	if resp.CreateLegalEntityApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateLegalEntityApiReply.Data.LegalEntityId == nil || *resp.CreateLegalEntityApiReply.Data.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", resp.CreateLegalEntityApiReply.Data.LegalEntityId)
	}
}

func TestCreateLegalEntity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Name("测试公司").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Errno == nil || *resp.CreateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateLegalEntityApiReply.Errno)
	}
}

func TestCreateLegalEntity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Errno == nil || *resp.CreateLegalEntityApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.CreateLegalEntityApiReply.Errno)
	}
}

func TestCreateLegalEntity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.CreateLegalEntityApiReply != nil {
		t.Errorf("CreateLegalEntityApiReply should be nil for non-200 response")
	}
}

func TestCreateLegalEntity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"legal_entity_id":"1125922289295589"},"request_id":"req_enc"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Name("测试公司").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Errno == nil || *resp.CreateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateLegalEntityApiReply.Errno)
	}
	if resp.CreateLegalEntityApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateLegalEntityApiReply.Data.LegalEntityId == nil || *resp.CreateLegalEntityApiReply.Data.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", resp.CreateLegalEntityApiReply.Data.LegalEntityId)
	}
}

func TestCreateLegalEntity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"legal_entity_id":"1125922289295589"},"request_id":"req_enc256"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Errno == nil || *resp.CreateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateLegalEntityApiReply.Errno)
	}
	if resp.CreateLegalEntityApiReply.RequestId == nil || *resp.CreateLegalEntityApiReply.RequestId != "req_enc256" {
		t.Errorf("RequestId = %v, want req_enc256", resp.CreateLegalEntityApiReply.RequestId)
	}
}

func TestCreateLegalEntity_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"legal_entity_id":"1125922289295589"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &legalEntity{option: option}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
	if resp.CreateLegalEntityApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestCreateLegalEntity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"legal_entity_id":"1125922289295589"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(NewCreateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateLegalEntity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateLegalEntity() error = %v", err)
	}
	if resp.CreateLegalEntityApiReply == nil {
		t.Fatal("CreateLegalEntityApiReply is nil")
	}
}

// === DelLegalEntity 资源方法测试 ===

func TestDelLegalEntity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/LegalEntity/del" {
			t.Errorf("expected path /river/LegalEntity/del, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			LegalEntityId("1125922289295589").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelLegalEntityApiReply.Errno)
	}
	if resp.DelLegalEntityApiReply.Errmsg == nil || *resp.DelLegalEntityApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", resp.DelLegalEntityApiReply.Errmsg)
	}
}

func TestDelLegalEntity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			LegalEntityId("1125922289295589").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelLegalEntityApiReply.Errno)
	}
}

func TestDelLegalEntity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.DelLegalEntityApiReply.Errno)
	}
}

func TestDelLegalEntity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.DelLegalEntityApiReply != nil {
		t.Errorf("DelLegalEntityApiReply should be nil for non-200 response")
	}
}

func TestDelLegalEntity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_enc"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			LegalEntityId("1125922289295589").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelLegalEntityApiReply.Errno)
	}
	if resp.DelLegalEntityApiReply.RequestId == nil || *resp.DelLegalEntityApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %v, want req_enc", resp.DelLegalEntityApiReply.RequestId)
	}
}

func TestDelLegalEntity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_enc256"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelLegalEntityApiReply.Errno)
	}
}

func TestDelLegalEntity_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &legalEntity{option: option}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
	if resp.DelLegalEntityApiReply.Errno == nil || *resp.DelLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelLegalEntityApiReply.Errno)
	}
}

func TestDelLegalEntity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(NewDelLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.DelLegalEntity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("DelLegalEntity() error = %v", err)
	}
	if resp.DelLegalEntityApiReply == nil {
		t.Fatal("DelLegalEntityApiReply is nil")
	}
}

// === GetLegalEntity 资源方法测试 ===

func TestGetLegalEntity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/LegalEntity/get" {
			t.Errorf("expected path /river/LegalEntity/get, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("keyword") != "滴滴" {
			t.Errorf("keyword = %q, want 滴滴", r.URL.Query().Get("keyword"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"total": 1,
				"records": [
					{
						"legal_entity_id": 1125922289295589,
						"name": "滴滴出行科技有限公司",
						"status": 1,
						"address": "北京市朝阳区望京SOHO",
						"taxpayer_no": "91110000MA0012345X",
						"parent_id": 1125915646135311
					}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Keyword("滴滴").
		Offset(0).
		Length(10).
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if resp.GetLegalEntityApiReply.Errno == nil || *resp.GetLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetLegalEntityApiReply.Errno)
	}
	if resp.GetLegalEntityApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetLegalEntityApiReply.Data.Total == nil || *resp.GetLegalEntityApiReply.Data.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.GetLegalEntityApiReply.Data.Total)
	}
	if len(resp.GetLegalEntityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetLegalEntityApiReply.Data.Records))
	}
	record := resp.GetLegalEntityApiReply.Data.Records[0]
	if record.LegalEntityId == nil || *record.LegalEntityId != 1125922289295589 {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", record.LegalEntityId)
	}
	if record.Name == nil || *record.Name != "滴滴出行科技有限公司" {
		t.Errorf("Name = %v, want 滴滴出行科技有限公司", record.Name)
	}
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Status = %v, want 1", record.Status)
	}
}

func TestGetLegalEntity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if len(resp.GetLegalEntityApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.GetLegalEntityApiReply.Data.Records))
	}
}

func TestGetLegalEntity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if resp.GetLegalEntityApiReply.Errno == nil || *resp.GetLegalEntityApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.GetLegalEntityApiReply.Errno)
	}
}

func TestGetLegalEntity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetLegalEntityApiReply != nil {
		t.Errorf("GetLegalEntityApiReply should be nil for non-200 response")
	}
}

func TestGetLegalEntity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"legal_entity_id":1125922289295589,"name":"滴滴出行","status":1}]},"request_id":"req_enc"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if resp.GetLegalEntityApiReply.Errno == nil || *resp.GetLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetLegalEntityApiReply.Errno)
	}
	if len(resp.GetLegalEntityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetLegalEntityApiReply.Data.Records))
	}
	if resp.GetLegalEntityApiReply.Data.Records[0].LegalEntityId == nil || *resp.GetLegalEntityApiReply.Data.Records[0].LegalEntityId != 1125922289295589 {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", resp.GetLegalEntityApiReply.Data.Records[0].LegalEntityId)
	}
}

func TestGetLegalEntity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"legal_entity_id":1125922289295589,"name":"滴滴出行"}]},"request_id":"req_enc256"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if len(resp.GetLegalEntityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetLegalEntityApiReply.Data.Records))
	}
	if resp.GetLegalEntityApiReply.Data.Records[0].Name == nil || *resp.GetLegalEntityApiReply.Data.Records[0].Name != "滴滴出行" {
		t.Errorf("Name = %v, want 滴滴出行", resp.GetLegalEntityApiReply.Data.Records[0].Name)
	}
}

func TestGetLegalEntity_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"legal_entity_id":1125922289295589}]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &legalEntity{option: option}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
	if len(resp.GetLegalEntityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetLegalEntityApiReply.Data.Records))
	}
}

func TestGetLegalEntity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetLegalEntityApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetLegalEntity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetLegalEntity() error = %v", err)
	}
	if resp.GetLegalEntityApiReply == nil {
		t.Fatal("GetLegalEntityApiReply is nil")
	}
}

// === UpdateLegalEntity 资源方法测试 ===

func TestUpdateLegalEntity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/LegalEntity/edit" {
			t.Errorf("expected path /river/LegalEntity/edit, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			LegalEntityId("1125922289295589").
			Name("更新后的公司名").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateLegalEntityApiReply.Errno)
	}
	if resp.UpdateLegalEntityApiReply.Errmsg == nil || *resp.UpdateLegalEntityApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", resp.UpdateLegalEntityApiReply.Errmsg)
	}
}

func TestUpdateLegalEntity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			LegalEntityId("1125922289295589").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateLegalEntityApiReply.Errno)
	}
}

func TestUpdateLegalEntity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.UpdateLegalEntityApiReply.Errno)
	}
}

func TestUpdateLegalEntity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.UpdateLegalEntityApiReply != nil {
		t.Errorf("UpdateLegalEntityApiReply should be nil for non-200 response")
	}
}

func TestUpdateLegalEntity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_enc"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			LegalEntityId("1125922289295589").
			Name("更新公司").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateLegalEntityApiReply.Errno)
	}
	if resp.UpdateLegalEntityApiReply.RequestId == nil || *resp.UpdateLegalEntityApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %v, want req_enc", resp.UpdateLegalEntityApiReply.RequestId)
	}
}

func TestUpdateLegalEntity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_enc256"}`
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

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateLegalEntityApiReply.Errno)
	}
}

func TestUpdateLegalEntity_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &legalEntity{option: option}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
	if resp.UpdateLegalEntityApiReply.Errno == nil || *resp.UpdateLegalEntityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateLegalEntityApiReply.Errno)
	}
}

func TestUpdateLegalEntity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newLegalEntityTestOption(testServer.URL)
	svc := &legalEntity{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(NewUpdateLegalEntityRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateLegalEntity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateLegalEntity() error = %v", err)
	}
	if resp.UpdateLegalEntityApiReply == nil {
		t.Fatal("UpdateLegalEntityApiReply is nil")
	}
}
