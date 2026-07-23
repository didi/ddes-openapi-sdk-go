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

// --- UpdateOutApprovalStatusRequestBuilder 测试 ---

func TestUpdateOutApprovalStatusRequestBuilder_FullParams(t *testing.T) {
	req := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		OutId("out_approval_001").
		Status(1).
		ApprovalType(2).
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
	if req.OutId == nil || *req.OutId != "out_approval_001" {
		t.Errorf("OutId = %v, want out_approval_001", req.OutId)
	}
	if req.Status == nil || *req.Status != 1 {
		t.Errorf("Status = %v, want 1", req.Status)
	}
	if req.ApprovalType == nil || *req.ApprovalType != 2 {
		t.Errorf("ApprovalType = %v, want 2", req.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusRequestBuilder_PartialParams(t *testing.T) {
	req := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		OutId("out_approval_001").
		Status(1).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.OutId == nil || *req.OutId != "out_approval_001" {
		t.Errorf("OutId = %v, want out_approval_001", req.OutId)
	}
	if req.Status == nil || *req.Status != 1 {
		t.Errorf("Status = %v, want 1", req.Status)
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
	if req.ApprovalType != nil {
		t.Errorf("ApprovalType = %v, want nil", req.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusRequestBuilder_ZeroIntValues(t *testing.T) {
	// int32 零值也应该被设置
	req := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		Status(0).
		ApprovalType(0).
		Build()

	if req.Status == nil || *req.Status != 0 {
		t.Errorf("Status = %v, want 0", req.Status)
	}
	if req.ApprovalType == nil || *req.ApprovalType != 0 {
		t.Errorf("ApprovalType = %v, want 0", req.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewUpdateOutApprovalStatusRequestBuilder().
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
	if req.OutId != nil {
		t.Errorf("OutId = %v, want nil", req.OutId)
	}
	if req.Status != nil {
		t.Errorf("Status = %v, want nil", req.Status)
	}
	if req.ApprovalType != nil {
		t.Errorf("ApprovalType = %v, want nil", req.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusRequestBuilder_JSONMarshal(t *testing.T) {
	req := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		OutId("out_approval_001").
		Status(1).
		ApprovalType(2).
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
	if v, ok := m["out_id"].(string); !ok || v != "out_approval_001" {
		t.Errorf("out_id = %v, want out_approval_001", m["out_id"])
	}
	if v, ok := m["status"].(float64); !ok || v != 1 {
		t.Errorf("status = %v, want 1", m["status"])
	}
	if v, ok := m["approval_type"].(float64); !ok || v != 2 {
		t.Errorf("approval_type = %v, want 2", m["approval_type"])
	}
}

func TestUpdateOutApprovalStatusRequestBuilder_OmitEmpty_JSON(t *testing.T) {
	req := NewUpdateOutApprovalStatusRequestBuilder().
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
	if _, ok := m["out_id"]; ok {
		t.Errorf("out_id should be omitted, got %v", m["out_id"])
	}
	if _, ok := m["status"]; ok {
		t.Errorf("status should be omitted, got %v", m["status"])
	}
	if _, ok := m["approval_type"]; ok {
		t.Errorf("approval_type should be omitted, got %v", m["approval_type"])
	}
}

// --- UpdateOutApprovalStatusApiReqBuilder 测试 ---

func TestUpdateOutApprovalStatusApiReqBuilder_FullParams(t *testing.T) {
	requestBody := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		OutId("out_approval_001").
		Status(1).
		ApprovalType(2).
		Build()

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(requestBody).
		Build()

	// 验证 Body 被设置
	if req.apiReq.Body == nil {
		t.Fatal("Body should not be nil")
	}
	body, ok := req.apiReq.Body.(*UpdateOutApprovalStatusRequest)
	if !ok {
		t.Fatal("Body should be *UpdateOutApprovalStatusRequest")
	}
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.OutId == nil || *body.OutId != "out_approval_001" {
		t.Errorf("Body.OutId = %v, want out_approval_001", body.OutId)
	}
	if body.Status == nil || *body.Status != 1 {
		t.Errorf("Body.Status = %v, want 1", body.Status)
	}
	if body.ApprovalType == nil || *body.ApprovalType != 2 {
		t.Errorf("Body.ApprovalType = %v, want 2", body.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusApiReqBuilder_PartialParams(t *testing.T) {
	requestBody := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		OutId("out_approval_001").
		Build()

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*UpdateOutApprovalStatusRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 未设置的字段应为 nil
	if body.AccessToken != nil {
		t.Errorf("Body.AccessToken = %v, want nil", body.AccessToken)
	}
	if body.Status != nil {
		t.Errorf("Body.Status = %v, want nil", body.Status)
	}
	if body.ApprovalType != nil {
		t.Errorf("Body.ApprovalType = %v, want nil", body.ApprovalType)
	}
}

func TestUpdateOutApprovalStatusApiReqBuilder_OnlyCommonParams(t *testing.T) {
	requestBody := NewUpdateOutApprovalStatusRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*UpdateOutApprovalStatusRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 业务参数不应存在
	if body.OutId != nil {
		t.Errorf("Body.OutId = %v, want nil", body.OutId)
	}
	if body.Status != nil {
		t.Errorf("Body.Status = %v, want nil", body.Status)
	}
	if body.ApprovalType != nil {
		t.Errorf("Body.ApprovalType = %v, want nil", body.ApprovalType)
	}
}

// --- UpdateOutApprovalStatusApiReply 反序列化测试 ---

func TestUpdateOutApprovalStatusApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {"result": "ok"},
		"request_id": "test_request_id"
	}`

	var reply UpdateOutApprovalStatusApiReply
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
	if reply.Data == nil {
		t.Error("Data should not be nil")
	}
}

func TestUpdateOutApprovalStatusApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply UpdateOutApprovalStatusApiReply
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

func TestUpdateOutApprovalStatusApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply UpdateOutApprovalStatusApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
}

func TestUpdateOutApprovalStatusApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply UpdateOutApprovalStatusApiReply
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

// --- UpdateOutApprovalStatusReply Builder 测试（空结构体）---

func TestUpdateOutApprovalStatusReply(t *testing.T) {
	// UpdateOutApprovalStatusReply 为空结构体，验证可正常构造
	reply := UpdateOutApprovalStatusReply{}
	if reply != (UpdateOutApprovalStatusReply{}) {
		t.Errorf("UpdateOutApprovalStatusReply should be zero value struct")
	}
}

// --- UpdateOutApprovalStatus 资源方法测试 ---

func newOutApprovalTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestUpdateOutApprovalStatus_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/OutApproval/Status" {
			t.Errorf("expected path /river/OutApproval/Status, got %s", r.URL.Path)
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

	option := newOutApprovalTestOption(testServer.URL)
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			OutId("out_approval_001").
			Status(1).
			ApprovalType(2).
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
	if resp.UpdateOutApprovalStatusApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.UpdateOutApprovalStatusApiReply.Errmsg)
	}
	if resp.UpdateOutApprovalStatusApiReply.RequestId != "req_001" {
		t.Errorf("RequestId = %q, want req_001", resp.UpdateOutApprovalStatusApiReply.RequestId)
	}
}

func TestUpdateOutApprovalStatus_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newOutApprovalTestOption(testServer.URL)
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			OutId("out_approval_001").
			Status(1).
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
}

func TestUpdateOutApprovalStatus_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newOutApprovalTestOption(testServer.URL)
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
	if resp.UpdateOutApprovalStatusApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.UpdateOutApprovalStatusApiReply.Errmsg)
	}
}

func TestUpdateOutApprovalStatus_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newOutApprovalTestOption(testServer.URL)
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.UpdateOutApprovalStatusApiReply != nil {
		t.Errorf("UpdateOutApprovalStatusApiReply should be nil for non-200 response")
	}
}

func TestUpdateOutApprovalStatus_WithEncryption_AES128(t *testing.T) {
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

	option := newOutApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			OutId("out_approval_001").
			Status(1).
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
	if resp.UpdateOutApprovalStatusApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %q, want req_enc", resp.UpdateOutApprovalStatusApiReply.RequestId)
	}
}

func TestUpdateOutApprovalStatus_WithEncryption_AES256(t *testing.T) {
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

	option := newOutApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
	if resp.UpdateOutApprovalStatusApiReply.RequestId != "req_enc256" {
		t.Errorf("RequestId = %q, want req_enc256", resp.UpdateOutApprovalStatusApiReply.RequestId)
	}
}

func TestUpdateOutApprovalStatus_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newOutApprovalTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &outApproval{option: option}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
	if resp.UpdateOutApprovalStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateOutApprovalStatusApiReply.Errno)
	}
}

func TestUpdateOutApprovalStatus_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newOutApprovalTestOption(testServer.URL)
	svc := &outApproval{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(NewUpdateOutApprovalStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateOutApprovalStatus(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateOutApprovalStatus() error = %v", err)
	}
	if resp.UpdateOutApprovalStatusApiReply == nil {
		t.Fatal("UpdateOutApprovalStatusApiReply is nil")
	}
}
