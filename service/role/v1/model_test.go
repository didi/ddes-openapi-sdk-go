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

// --- ListRoleApiReqBuilder 测试 ---

func TestListRoleApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListRoleApiReqBuilder().
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

func TestListRoleApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Sign("test_sign").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("sign") != "test_sign" {
		t.Errorf("sign mismatch")
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
}

func TestListRoleApiReqBuilder_OnlyCommonParams(t *testing.T) {
	// ListRoleApiReqBuilder 仅含 5 个通用参数，无业务参数
	req := NewListRoleApiReqBuilder().
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

// --- ListRoleApiReply 反序列化测试 ---

func TestListRoleApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"id": "1125915646090887", "name": "管理员", "alias": "admin"},
			{"id": "1125915646107623", "name": "财务", "alias": "finance"}
		],
		"request_id": "test_request_id"
	}`

	var reply ListRoleApiReply
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
	if len(reply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(reply.Data))
	}

	role0 := reply.Data[0]
	if role0.Id == nil || *role0.Id != "1125915646090887" {
		t.Errorf("Data[0].Id = %v, want 1125915646090887", role0.Id)
	}
	if role0.Name == nil || *role0.Name != "管理员" {
		t.Errorf("Data[0].Name = %v, want 管理员", role0.Name)
	}
	if role0.Alias == nil || *role0.Alias != "admin" {
		t.Errorf("Data[0].Alias = %v, want admin", role0.Alias)
	}

	role1 := reply.Data[1]
	if role1.Id == nil || *role1.Id != "1125915646107623" {
		t.Errorf("Data[1].Id = %v, want 1125915646107623", role1.Id)
	}
	if role1.Name == nil || *role1.Name != "财务" {
		t.Errorf("Data[1].Name = %v, want 财务", role1.Name)
	}
	if role1.Alias == nil || *role1.Alias != "finance" {
		t.Errorf("Data[1].Alias = %v, want finance", role1.Alias)
	}
}

func TestListRoleApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply ListRoleApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %v, want param error", reply.Errmsg)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestListRoleApiReply_MultipleItems(t *testing.T) {
	// 多条数据 + 部分字段缺失
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"id": "1125915646090887", "name": "管理员", "alias": "admin"},
			{"id": "1125915646107623", "name": "财务"},
			{"id": "1125915646123456"}
		],
		"request_id": "req_multi"
	}`

	var reply ListRoleApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 3 {
		t.Fatalf("Data len = %d, want 3", len(reply.Data))
	}
	// 第1条全字段
	if reply.Data[0].Alias == nil || *reply.Data[0].Alias != "admin" {
		t.Errorf("Data[0].Alias = %v, want admin", reply.Data[0].Alias)
	}
	// 第2条缺 alias
	if reply.Data[1].Name == nil || *reply.Data[1].Name != "财务" {
		t.Errorf("Data[1].Name = %v, want 财务", reply.Data[1].Name)
	}
	if reply.Data[1].Alias != nil {
		t.Errorf("Data[1].Alias = %v, want nil", reply.Data[1].Alias)
	}
	// 第3条仅 id
	if reply.Data[2].Id == nil || *reply.Data[2].Id != "1125915646123456" {
		t.Errorf("Data[2].Id = %v, want 1125915646123456", reply.Data[2].Id)
	}
	if reply.Data[2].Name != nil {
		t.Errorf("Data[2].Name = %v, want nil", reply.Data[2].Name)
	}
	if reply.Data[2].Alias != nil {
		t.Errorf("Data[2].Alias = %v, want nil", reply.Data[2].Alias)
	}
}

func TestListRoleApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListRoleApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListRoleApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListRoleApiReply
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

// --- ListRoleReply 反序列化测试（无 Builder，直接验证字段）---

func TestListRoleReply_Deserialization(t *testing.T) {
	jsonData := `{"id":"1125915646090887","name":"管理员","alias":"admin"}`

	var role ListRoleReply
	if err := json.Unmarshal([]byte(jsonData), &role); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if role.Id == nil || *role.Id != "1125915646090887" {
		t.Errorf("Id = %v, want 1125915646090887", role.Id)
	}
	if role.Name == nil || *role.Name != "管理员" {
		t.Errorf("Name = %v, want 管理员", role.Name)
	}
	if role.Alias == nil || *role.Alias != "admin" {
		t.Errorf("Alias = %v, want admin", role.Alias)
	}

	// 部分字段缺失
	jsonData2 := `{"id":"1125915646107623"}`
	var role2 ListRoleReply
	if err := json.Unmarshal([]byte(jsonData2), &role2); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if role2.Id == nil || *role2.Id != "1125915646107623" {
		t.Errorf("Id = %v, want 1125915646107623", role2.Id)
	}
	if role2.Name != nil {
		t.Errorf("Name = %v, want nil", role2.Name)
	}
	if role2.Alias != nil {
		t.Errorf("Alias = %v, want nil", role2.Alias)
	}
}

// --- ListRole 资源方法测试 ---

func newRoleTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestListRole_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Role/get" {
			t.Errorf("expected path /river/Role/get, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
				{"id": "1125915646090887", "name": "管理员", "alias": "admin"},
				{"id": "1125915646107623", "name": "财务", "alias": "finance"}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if resp.ListRoleApiReply.Errno == nil || *resp.ListRoleApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListRoleApiReply.Errno)
	}
	if len(resp.ListRoleApiReply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(resp.ListRoleApiReply.Data))
	}
	if resp.ListRoleApiReply.Data[0].Id == nil || *resp.ListRoleApiReply.Data[0].Id != "1125915646090887" {
		t.Errorf("Data[0].Id = %v, want 1125915646090887", resp.ListRoleApiReply.Data[0].Id)
	}
	if resp.ListRoleApiReply.Data[0].Name == nil || *resp.ListRoleApiReply.Data[0].Name != "管理员" {
		t.Errorf("Data[0].Name = %v, want 管理员", resp.ListRoleApiReply.Data[0].Name)
	}
	if resp.ListRoleApiReply.Data[0].Alias == nil || *resp.ListRoleApiReply.Data[0].Alias != "admin" {
		t.Errorf("Data[0].Alias = %v, want admin", resp.ListRoleApiReply.Data[0].Alias)
	}
}

func TestListRole_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if len(resp.ListRoleApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListRoleApiReply.Data))
	}
}

func TestListRole_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if resp.ListRoleApiReply.Errno == nil || *resp.ListRoleApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.ListRoleApiReply.Errno)
	}
}

func TestListRole_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.ListRoleApiReply != nil {
		t.Errorf("ListRoleApiReply should be nil for non-200 response")
	}
}

func TestListRole_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"id":"1125915646090887","name":"管理员","alias":"admin"}],"request_id":"req_enc"}`
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

	option := newRoleTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if resp.ListRoleApiReply.Errno == nil || *resp.ListRoleApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListRoleApiReply.Errno)
	}
	if len(resp.ListRoleApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRoleApiReply.Data))
	}
	if resp.ListRoleApiReply.Data[0].Id == nil || *resp.ListRoleApiReply.Data[0].Id != "1125915646090887" {
		t.Errorf("Data[0].Id = %v, want 1125915646090887", resp.ListRoleApiReply.Data[0].Id)
	}
	if resp.ListRoleApiReply.Data[0].Name == nil || *resp.ListRoleApiReply.Data[0].Name != "管理员" {
		t.Errorf("Data[0].Name = %v, want 管理员", resp.ListRoleApiReply.Data[0].Name)
	}
}

func TestListRole_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"id":"1125915646090887","name":"管理员"}],"request_id":"req_enc256"}`
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

	option := newRoleTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if len(resp.ListRoleApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRoleApiReply.Data))
	}
	if resp.ListRoleApiReply.Data[0].Id == nil || *resp.ListRoleApiReply.Data[0].Id != "1125915646090887" {
		t.Errorf("Data[0].Id = %v, want 1125915646090887", resp.ListRoleApiReply.Data[0].Id)
	}
}

func TestListRole_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"id":"1125915646090887","name":"管理员"}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	ro := &role{option: option}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if len(resp.ListRoleApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRoleApiReply.Data))
	}
}

func TestListRole_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"id":"1125915646090887","name":"管理员"}],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newRoleTestOption(testServer.URL)
	ro := &role{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewListRoleApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := ro.ListRole(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListRole() error = %v", err)
	}
	if resp.ListRoleApiReply == nil {
		t.Fatal("ListRoleApiReply is nil")
	}
	if len(resp.ListRoleApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListRoleApiReply.Data))
	}
}
