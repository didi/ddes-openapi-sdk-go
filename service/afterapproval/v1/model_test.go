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

// --- CreatePersonalReceiptRequestBuilder 测试 ---

func TestCreatePersonalReceiptRequestBuilder_FullParams(t *testing.T) {
	req := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		OrderId("1125922289295589").
		IsPass(1).
		ApproverPhone("13800000001").
		Remark("审批通过").
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
	if req.OrderId == nil || *req.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", req.OrderId)
	}
	if req.IsPass == nil || *req.IsPass != 1 {
		t.Errorf("IsPass = %v, want 1", req.IsPass)
	}
	if req.ApproverPhone == nil || *req.ApproverPhone != "13800000001" {
		t.Errorf("ApproverPhone = %v, want 13800000001", req.ApproverPhone)
	}
	if req.Remark == nil || *req.Remark != "审批通过" {
		t.Errorf("Remark = %v, want 审批通过", req.Remark)
	}
}

func TestCreatePersonalReceiptRequestBuilder_PartialParams(t *testing.T) {
	req := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		OrderId("1125922289295589").
		IsPass(1).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.OrderId == nil || *req.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", req.OrderId)
	}
	if req.IsPass == nil || *req.IsPass != 1 {
		t.Errorf("IsPass = %v, want 1", req.IsPass)
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
	if req.ApproverPhone != nil {
		t.Errorf("ApproverPhone = %v, want nil", req.ApproverPhone)
	}
	if req.Remark != nil {
		t.Errorf("Remark = %v, want nil", req.Remark)
	}
}

func TestCreatePersonalReceiptRequestBuilder_ZeroIntValues(t *testing.T) {
	// int32 零值也应该被设置
	req := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		IsPass(0).
		Build()

	if req.IsPass == nil || *req.IsPass != 0 {
		t.Errorf("IsPass = %v, want 0", req.IsPass)
	}
}

func TestCreatePersonalReceiptRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewCreatePersonalReceiptRequestBuilder().
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
	if req.OrderId != nil {
		t.Errorf("OrderId = %v, want nil", req.OrderId)
	}
	if req.IsPass != nil {
		t.Errorf("IsPass = %v, want nil", req.IsPass)
	}
	if req.ApproverPhone != nil {
		t.Errorf("ApproverPhone = %v, want nil", req.ApproverPhone)
	}
	if req.Remark != nil {
		t.Errorf("Remark = %v, want nil", req.Remark)
	}
}

func TestCreatePersonalReceiptRequestBuilder_JSONMarshal(t *testing.T) {
	req := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		OrderId("1125922289295589").
		IsPass(1).
		ApproverPhone("13800000001").
		Remark("审批通过").
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
	if v, ok := m["order_id"].(string); !ok || v != "1125922289295589" {
		t.Errorf("order_id = %v, want 1125922289295589", m["order_id"])
	}
	if v, ok := m["is_pass"].(float64); !ok || v != 1 {
		t.Errorf("is_pass = %v, want 1", m["is_pass"])
	}
	if v, ok := m["approver_phone"].(string); !ok || v != "13800000001" {
		t.Errorf("approver_phone = %v, want 13800000001", m["approver_phone"])
	}
	if v, ok := m["remark"].(string); !ok || v != "审批通过" {
		t.Errorf("remark = %v, want 审批通过", m["remark"])
	}
}

func TestCreatePersonalReceiptRequestBuilder_OmitEmpty_JSON(t *testing.T) {
	req := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		Build()

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal to map failed: %v", err)
	}

	// 未设置的业务字段不应出现在 JSON 中（omitempty）
	if _, ok := m["order_id"]; ok {
		t.Errorf("order_id should be omitted, got %v", m["order_id"])
	}
	if _, ok := m["is_pass"]; ok {
		t.Errorf("is_pass should be omitted, got %v", m["is_pass"])
	}
	if _, ok := m["approver_phone"]; ok {
		t.Errorf("approver_phone should be omitted, got %v", m["approver_phone"])
	}
	if _, ok := m["remark"]; ok {
		t.Errorf("remark should be omitted, got %v", m["remark"])
	}
}

// --- CreatePersonalReceiptApiReqBuilder 测试 ---

func TestCreatePersonalReceiptApiReqBuilder_FullParams(t *testing.T) {
	requestBody := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		OrderId("1125922289295589").
		IsPass(1).
		ApproverPhone("13800000001").
		Remark("审批通过").
		Build()

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(requestBody).
		Build()

	// 验证 Body 被设置
	if req.apiReq.Body == nil {
		t.Fatal("Body should not be nil")
	}
	body, ok := req.apiReq.Body.(*CreatePersonalReceiptRequest)
	if !ok {
		t.Fatal("Body should be *CreatePersonalReceiptRequest")
	}
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.OrderId == nil || *body.OrderId != "1125922289295589" {
		t.Errorf("Body.OrderId = %v, want 1125922289295589", body.OrderId)
	}
	if body.IsPass == nil || *body.IsPass != 1 {
		t.Errorf("Body.IsPass = %v, want 1", body.IsPass)
	}
	if body.ApproverPhone == nil || *body.ApproverPhone != "13800000001" {
		t.Errorf("Body.ApproverPhone = %v, want 13800000001", body.ApproverPhone)
	}
	if body.Remark == nil || *body.Remark != "审批通过" {
		t.Errorf("Body.Remark = %v, want 审批通过", body.Remark)
	}
}

func TestCreatePersonalReceiptApiReqBuilder_PartialParams(t *testing.T) {
	requestBody := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		OrderId("1125922289295589").
		Build()

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*CreatePersonalReceiptRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.OrderId == nil || *body.OrderId != "1125922289295589" {
		t.Errorf("Body.OrderId = %v, want 1125922289295589", body.OrderId)
	}
	// 未设置的字段应为 nil
	if body.AccessToken != nil {
		t.Errorf("Body.AccessToken = %v, want nil", body.AccessToken)
	}
	if body.IsPass != nil {
		t.Errorf("Body.IsPass = %v, want nil", body.IsPass)
	}
	if body.ApproverPhone != nil {
		t.Errorf("Body.ApproverPhone = %v, want nil", body.ApproverPhone)
	}
}

func TestCreatePersonalReceiptApiReqBuilder_OnlyCommonParams(t *testing.T) {
	requestBody := NewCreatePersonalReceiptRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*CreatePersonalReceiptRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 业务参数不应存在
	if body.OrderId != nil {
		t.Errorf("Body.OrderId = %v, want nil", body.OrderId)
	}
	if body.IsPass != nil {
		t.Errorf("Body.IsPass = %v, want nil", body.IsPass)
	}
	if body.ApproverPhone != nil {
		t.Errorf("Body.ApproverPhone = %v, want nil", body.ApproverPhone)
	}
	if body.Remark != nil {
		t.Errorf("Body.Remark = %v, want nil", body.Remark)
	}
}

// --- CreatePersonalReceiptApiReply 反序列化测试 ---

func TestCreatePersonalReceiptApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {"result": "ok"},
		"request_id": "test_request_id"
	}`

	var reply CreatePersonalReceiptApiReply
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
	if reply.Data == nil {
		t.Error("Data should not be nil")
	}
}

func TestCreatePersonalReceiptApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply CreatePersonalReceiptApiReply
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

func TestCreatePersonalReceiptApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply CreatePersonalReceiptApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
}

func TestCreatePersonalReceiptApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreatePersonalReceiptApiReply
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

// --- CreatePersonalReceiptReply 测试（空结构体）---

func TestCreatePersonalReceiptReply(t *testing.T) {
	// CreatePersonalReceiptReply 为空结构体，验证可正常构造
	reply := CreatePersonalReceiptReply{}
	if reply != (CreatePersonalReceiptReply{}) {
		t.Errorf("CreatePersonalReceiptReply should be zero value struct")
	}
}

// --- CreatePersonalReceipt 资源方法测试 ---

func newAfterApprovalTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestCreatePersonalReceipt_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/AfterApproval/createPersonalReceipt" {
			t.Errorf("expected path /river/AfterApproval/createPersonalReceipt, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {"result": "ok"},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			OrderId("1125922289295589").
			IsPass(1).
			ApproverPhone("13800000001").
			Remark("审批通过").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreatePersonalReceiptApiReply.Errno)
	}
	if resp.CreatePersonalReceiptApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.CreatePersonalReceiptApiReply.Errmsg)
	}
	if resp.CreatePersonalReceiptApiReply.RequestId == nil || *resp.CreatePersonalReceiptApiReply.RequestId != "req_001" {
		t.Errorf("RequestId = %v, want req_001", resp.CreatePersonalReceiptApiReply.RequestId)
	}
}

func TestCreatePersonalReceipt_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			OrderId("1125922289295589").
			IsPass(1).
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreatePersonalReceiptApiReply.Errno)
	}
}

func TestCreatePersonalReceipt_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.CreatePersonalReceiptApiReply.Errno)
	}
	if resp.CreatePersonalReceiptApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.CreatePersonalReceiptApiReply.Errmsg)
	}
}

func TestCreatePersonalReceipt_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.CreatePersonalReceiptApiReply != nil {
		t.Errorf("CreatePersonalReceiptApiReply should be nil for non-200 response")
	}
}

func TestCreatePersonalReceipt_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_enc"}`
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

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			OrderId("1125922289295589").
			IsPass(1).
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreatePersonalReceiptApiReply.Errno)
	}
	if resp.CreatePersonalReceiptApiReply.RequestId == nil || *resp.CreatePersonalReceiptApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %v, want req_enc", resp.CreatePersonalReceiptApiReply.RequestId)
	}
}

func TestCreatePersonalReceipt_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_enc256"}`
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

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreatePersonalReceiptApiReply.Errno)
	}
	if resp.CreatePersonalReceiptApiReply.RequestId == nil || *resp.CreatePersonalReceiptApiReply.RequestId != "req_enc256" {
		t.Errorf("RequestId = %v, want req_enc256", resp.CreatePersonalReceiptApiReply.RequestId)
	}
}

func TestCreatePersonalReceipt_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &afterApproval{option: option}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
	if resp.CreatePersonalReceiptApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.CreatePersonalReceiptApiReply.Errno)
	}
}

func TestCreatePersonalReceipt_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(NewCreatePersonalReceiptRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreatePersonalReceipt(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreatePersonalReceipt() error = %v", err)
	}
	if resp.CreatePersonalReceiptApiReply == nil {
		t.Fatal("CreatePersonalReceiptApiReply is nil")
	}
}

// --- GetPersonalReceiptOrderApiReqBuilder 测试 ---

func TestGetPersonalReceiptOrderApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		QueryTimeType(1).
		StartDate("2024-01-01").
		EndDate("2024-01-31").
		Status(1).
		Type(2).
		Phone("13800000001").
		Offset("0").
		Length("10").
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
		{"query_time_type", "1"},
		{"start_date", "2024-01-01"},
		{"end_date", "2024-01-31"},
		{"status", "1"},
		{"type", "2"},
		{"phone", "13800000001"},
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

func TestGetPersonalReceiptOrderApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Offset("0").
		Length("10").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("phone") != "13800000001" {
		t.Errorf("phone mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	if req.apiReq.QueryParams.Get("length") != "10" {
		t.Errorf("length mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("query_time_type") != "" {
		t.Errorf("query_time_type should be empty, got %q", req.apiReq.QueryParams.Get("query_time_type"))
	}
	if req.apiReq.QueryParams.Get("start_date") != "" {
		t.Errorf("start_date should be empty, got %q", req.apiReq.QueryParams.Get("start_date"))
	}
	if req.apiReq.QueryParams.Get("end_date") != "" {
		t.Errorf("end_date should be empty, got %q", req.apiReq.QueryParams.Get("end_date"))
	}
	if req.apiReq.QueryParams.Get("status") != "" {
		t.Errorf("status should be empty, got %q", req.apiReq.QueryParams.Get("status"))
	}
	if req.apiReq.QueryParams.Get("type") != "" {
		t.Errorf("type should be empty, got %q", req.apiReq.QueryParams.Get("type"))
	}
}

func TestGetPersonalReceiptOrderApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		QueryTimeType(0).
		Status(0).
		Type(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("query_time_type") != "0" {
		t.Errorf("query_time_type = %q, want \"0\"", req.apiReq.QueryParams.Get("query_time_type"))
	}
	if req.apiReq.QueryParams.Get("status") != "0" {
		t.Errorf("status = %q, want \"0\"", req.apiReq.QueryParams.Get("status"))
	}
	if req.apiReq.QueryParams.Get("type") != "0" {
		t.Errorf("type = %q, want \"0\"", req.apiReq.QueryParams.Get("type"))
	}
}

func TestGetPersonalReceiptOrderApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetPersonalReceiptOrderApiReqBuilder().
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
	for _, key := range []string{"query_time_type", "start_date", "end_date", "status", "type", "phone", "offset", "length"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- GetPersonalReceiptOrderApiReply 反序列化测试 ---

func TestGetPersonalReceiptOrderApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{
					"order_id": "1125922289295589",
					"approval_id": "1125922289295590"
				}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply GetPersonalReceiptOrderApiReply
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
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != 2 {
		t.Errorf("Total = %d, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
	}
	record := reply.Data.Records[0]
	if record.OrderId == nil || *record.OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", record.OrderId)
	}
	if record.ApprovalId == nil || *record.ApprovalId != "1125922289295590" {
		t.Errorf("Records[0].ApprovalId = %v, want 1125922289295590", record.ApprovalId)
	}
}

func TestGetPersonalReceiptOrderApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply GetPersonalReceiptOrderApiReply
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

func TestGetPersonalReceiptOrderApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 3,
			"records": [
				{"order_id": "1125922289295589", "approval_id": "1125922289295590"},
				{"order_id": "1125922289295591"},
				{}
			]
		},
		"request_id": "req_multi"
	}`

	var reply GetPersonalReceiptOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != 3 {
		t.Errorf("Total = %d, want 3", reply.Data.Total)
	}
	if len(reply.Data.Records) != 3 {
		t.Fatalf("Records len = %d, want 3", len(reply.Data.Records))
	}
	// 第1条有 order_id 和 approval_id
	if reply.Data.Records[0].OrderId == nil || *reply.Data.Records[0].OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", reply.Data.Records[0].OrderId)
	}
	if reply.Data.Records[0].ApprovalId == nil || *reply.Data.Records[0].ApprovalId != "1125922289295590" {
		t.Errorf("Records[0].ApprovalId = %v, want 1125922289295590", reply.Data.Records[0].ApprovalId)
	}
	// 第2条有 order_id 但无 approval_id
	if reply.Data.Records[1].OrderId == nil || *reply.Data.Records[1].OrderId != "1125922289295591" {
		t.Errorf("Records[1].OrderId = %v, want 1125922289295591", reply.Data.Records[1].OrderId)
	}
	if reply.Data.Records[1].ApprovalId != nil {
		t.Errorf("Records[1].ApprovalId = %v, want nil", reply.Data.Records[1].ApprovalId)
	}
	// 第3条全空
	if reply.Data.Records[2].OrderId != nil {
		t.Errorf("Records[2].OrderId = %v, want nil", reply.Data.Records[2].OrderId)
	}
	if reply.Data.Records[2].ApprovalId != nil {
		t.Errorf("Records[2].ApprovalId = %v, want nil", reply.Data.Records[2].ApprovalId)
	}
}

func TestGetPersonalReceiptOrderApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_empty"}`

	var reply GetPersonalReceiptOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != 0 {
		t.Errorf("Total = %d, want 0", reply.Data.Total)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestGetPersonalReceiptOrderApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetPersonalReceiptOrderApiReply
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

// --- GetPersonalReceiptOrderRecord Builder 测试 ---

func TestGetPersonalReceiptOrderRecordBuilder(t *testing.T) {
	record := NewGetPersonalReceiptOrderRecordBuilder().
		OrderId("1125922289295589").
		ApprovalId("1125922289295590").
		Build()

	if record.OrderId == nil || *record.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", record.OrderId)
	}
	if record.ApprovalId == nil || *record.ApprovalId != "1125922289295590" {
		t.Errorf("ApprovalId = %v, want 1125922289295590", record.ApprovalId)
	}

	// 部分设置
	record2 := NewGetPersonalReceiptOrderRecordBuilder().
		OrderId("1125922289295591").
		Build()

	if record2.OrderId == nil || *record2.OrderId != "1125922289295591" {
		t.Errorf("OrderId = %v, want 1125922289295591", record2.OrderId)
	}
	if record2.ApprovalId != nil {
		t.Errorf("ApprovalId = %v, want nil", record2.ApprovalId)
	}
}

// --- GetPersonalReceiptOrder 资源方法测试 ---

func TestGetPersonalReceiptOrder_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/AfterApproval/getPersonalReceiptOrder" {
			t.Errorf("expected path /river/AfterApproval/getPersonalReceiptOrder, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("phone") != "13800000001" {
			t.Errorf("phone = %q, want 13800000001", r.URL.Query().Get("phone"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"total": 1,
				"records": [
					{"order_id": "1125922289295589", "approval_id": "1125922289295590"}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Phone("13800000001").
		Offset("0").
		Length("10").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetPersonalReceiptOrderApiReply.Errno)
	}
	if resp.GetPersonalReceiptOrderApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Data.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.GetPersonalReceiptOrderApiReply.Data.Total)
	}
	if len(resp.GetPersonalReceiptOrderApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetPersonalReceiptOrderApiReply.Data.Records))
	}
	record := resp.GetPersonalReceiptOrderApiReply.Data.Records[0]
	if record.OrderId == nil || *record.OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", record.OrderId)
	}
	if record.ApprovalId == nil || *record.ApprovalId != "1125922289295590" {
		t.Errorf("Records[0].ApprovalId = %v, want 1125922289295590", record.ApprovalId)
	}
}

func TestGetPersonalReceiptOrder_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Offset("0").
		Length("10").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(resp.GetPersonalReceiptOrderApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.GetPersonalReceiptOrderApiReply.Data.Records))
	}
}

func TestGetPersonalReceiptOrder_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.GetPersonalReceiptOrderApiReply.Errno)
	}
	if resp.GetPersonalReceiptOrderApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.GetPersonalReceiptOrderApiReply.Errmsg)
	}
}

func TestGetPersonalReceiptOrder_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetPersonalReceiptOrderApiReply != nil {
		t.Errorf("GetPersonalReceiptOrderApiReply should be nil for non-200 response")
	}
}

func TestGetPersonalReceiptOrder_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"order_id":"1125922289295589","approval_id":"1125922289295590"}]},"request_id":"req_enc"}`
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

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Offset("0").
		Length("10").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetPersonalReceiptOrderApiReply.Errno)
	}
	if resp.GetPersonalReceiptOrderApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(resp.GetPersonalReceiptOrderApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetPersonalReceiptOrderApiReply.Data.Records))
	}
	if resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId == nil || *resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId)
	}
}

func TestGetPersonalReceiptOrder_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"order_id":"1125922289295589"}]},"request_id":"req_enc256"}`
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

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(resp.GetPersonalReceiptOrderApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetPersonalReceiptOrderApiReply.Data.Records))
	}
	if resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId == nil || *resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", resp.GetPersonalReceiptOrderApiReply.Data.Records[0].OrderId)
	}
}

func TestGetPersonalReceiptOrder_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"order_id":"1125922289295589"}]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &afterApproval{option: option}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
	if resp.GetPersonalReceiptOrderApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(resp.GetPersonalReceiptOrderApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetPersonalReceiptOrderApiReply.Data.Records))
	}
}

func TestGetPersonalReceiptOrder_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newAfterApprovalTestOption(testServer.URL)
	svc := &afterApproval{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetPersonalReceiptOrderApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.GetPersonalReceiptOrder(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetPersonalReceiptOrder() error = %v", err)
	}
	if resp.GetPersonalReceiptOrderApiReply == nil {
		t.Fatal("GetPersonalReceiptOrderApiReply is nil")
	}
}
