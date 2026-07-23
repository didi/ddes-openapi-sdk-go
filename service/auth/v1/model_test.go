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

// --- AuthorizeApiReply 反序列化测试 ---

func TestAuthorizeApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"access_token": "accesstoken_demo_1125922289295589",
		"expires_in": 7200,
		"token_type": "Bearer",
		"scope": "all",
		"request_id": "req_001"
	}`

	var reply AuthorizeApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.AccessToken == nil || *reply.AccessToken != "accesstoken_demo_1125922289295589" {
		t.Errorf("AccessToken = %v, want accesstoken_demo_1125922289295589", reply.AccessToken)
	}
	if reply.ExpiresIn == nil || *reply.ExpiresIn != 7200 {
		t.Errorf("ExpiresIn = %v, want 7200", reply.ExpiresIn)
	}
	if reply.TokenType == nil || *reply.TokenType != "Bearer" {
		t.Errorf("TokenType = %v, want Bearer", reply.TokenType)
	}
	if reply.Scope == nil || *reply.Scope != "all" {
		t.Errorf("Scope = %v, want all", reply.Scope)
	}
	if reply.RequestId == nil || *reply.RequestId != "req_001" {
		t.Errorf("RequestId = %v, want req_001", reply.RequestId)
	}
}

func TestAuthorizeApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply AuthorizeApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %v, want param error", reply.Errmsg)
	}
	// 错误响应不应有 access_token
	if reply.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", reply.AccessToken)
	}
}

func TestAuthorizeApiReply_PartialFields(t *testing.T) {
	// 多条数据 + 部分字段缺失场景：仅返回 access_token，缺失 token_type/scope
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"access_token": "token_partial",
		"expires_in": 3600,
		"request_id": "req_partial"
	}`

	var reply AuthorizeApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.AccessToken == nil || *reply.AccessToken != "token_partial" {
		t.Errorf("AccessToken = %v, want token_partial", reply.AccessToken)
	}
	if reply.ExpiresIn == nil || *reply.ExpiresIn != 3600 {
		t.Errorf("ExpiresIn = %v, want 3600", reply.ExpiresIn)
	}
	// 缺失字段应为 nil
	if reply.TokenType != nil {
		t.Errorf("TokenType = %v, want nil", reply.TokenType)
	}
	if reply.Scope != nil {
		t.Errorf("Scope = %v, want nil", reply.Scope)
	}
}

func TestAuthorizeApiReply_EmptyData(t *testing.T) {
	// 空 data 场景：仅返回 errno/errmsg，无业务数据
	jsonData := `{"errno":0,"errmsg":"SUCCESS","request_id":"req_empty"}`

	var reply AuthorizeApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", reply.AccessToken)
	}
	if reply.ExpiresIn != nil {
		t.Errorf("ExpiresIn = %v, want nil", reply.ExpiresIn)
	}
}

func TestAuthorizeApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply AuthorizeApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", reply.AccessToken)
	}
}

// --- AuthorizeRequestBuilder 测试 ---

func TestAuthorizeRequestBuilder_FullParams(t *testing.T) {
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		ClientSecret("test_secret").
		GrantType("client_credentials").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ClientSecret == nil || *request.ClientSecret != "test_secret" {
		t.Errorf("ClientSecret = %v, want test_secret", request.ClientSecret)
	}
	if request.GrantType == nil || *request.GrantType != "client_credentials" {
		t.Errorf("GrantType = %v, want client_credentials", request.GrantType)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}

	// 验证 JSON marshal（POST 型业务参数序列化到 Body）
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &m); err != nil {
		t.Fatalf("Unmarshal map failed: %v", err)
	}
	if m["client_id"] != "test_client" {
		t.Errorf("json client_id = %v, want test_client", m["client_id"])
	}
	if m["client_secret"] != "test_secret" {
		t.Errorf("json client_secret = %v, want test_secret", m["client_secret"])
	}
	if m["grant_type"] != "client_credentials" {
		t.Errorf("json grant_type = %v, want client_credentials", m["grant_type"])
	}
	if m["timestamp"] != float64(1583484681) {
		t.Errorf("json timestamp = %v, want 1583484681", m["timestamp"])
	}
	if m["sign"] != "test_sign" {
		t.Errorf("json sign = %v, want test_sign", m["sign"])
	}
}

func TestAuthorizeRequestBuilder_PartialParams(t *testing.T) {
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		GrantType("client_credentials").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.GrantType == nil || *request.GrantType != "client_credentials" {
		t.Errorf("GrantType = %v, want client_credentials", request.GrantType)
	}
	// 未设置的字段应为 nil
	if request.ClientSecret != nil {
		t.Errorf("ClientSecret = %v, want nil", request.ClientSecret)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
}

func TestAuthorizeRequestBuilder_ZeroInt64Value(t *testing.T) {
	// int64 零值也应被写入 Request 字段（非 nil）
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if request.Timestamp == nil {
		t.Fatal("Timestamp = nil, want pointer to 0")
	}
	if *request.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", *request.Timestamp)
	}

	// 验证零值也序列化到 JSON
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &m); err != nil {
		t.Fatalf("Unmarshal map failed: %v", err)
	}
	if ts, ok := m["timestamp"]; !ok || ts != float64(0) {
		t.Errorf("json timestamp = %v, want 0", ts)
	}
}

func TestAuthorizeRequestBuilder_OnlyCommonParams(t *testing.T) {
	// auth 域的 Request 只有通用参数（client_id/client_secret/grant_type/timestamp/sign），无额外业务参数
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		ClientSecret("test_secret").
		GrantType("client_credentials").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	// 全部通用参数均设置
	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ClientSecret == nil || *request.ClientSecret != "test_secret" {
		t.Errorf("ClientSecret = %v, want test_secret", request.ClientSecret)
	}
	if request.GrantType == nil || *request.GrantType != "client_credentials" {
		t.Errorf("GrantType = %v, want client_credentials", request.GrantType)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
}

// --- AuthorizeApiReqBuilder 测试 ---

func TestAuthorizeApiReqBuilder_BodyAndQueryParams(t *testing.T) {
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		ClientSecret("test_secret").
		GrantType("client_credentials").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(request).
		Build()

	// POST 型：业务参数在 Body（即 AuthorizeRequest），QueryParams 为空
	if req.apiReq.Body == nil {
		t.Fatal("apiReq.Body is nil")
	}
	bodyReq, ok := req.apiReq.Body.(*AuthorizeRequest)
	if !ok {
		t.Fatalf("apiReq.Body type = %T, want *AuthorizeRequest", req.apiReq.Body)
	}
	if bodyReq.ClientId == nil || *bodyReq.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", bodyReq.ClientId)
	}
	if bodyReq.GrantType == nil || *bodyReq.GrantType != "client_credentials" {
		t.Errorf("Body.GrantType = %v, want client_credentials", bodyReq.GrantType)
	}
	// QueryParams 为空 url.Values，不含业务参数
	if v := req.apiReq.QueryParams.Get("client_id"); v != "" {
		t.Errorf("QueryParams[client_id] = %q, want empty (POST 型业务参数走 Body)", v)
	}
}

func TestAuthorizeApiReqBuilder_PartialParams(t *testing.T) {
	request := NewAuthorizeRequestBuilder().
		ClientId("test_client").
		Build()

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(request).
		Build()

	bodyReq, ok := req.apiReq.Body.(*AuthorizeRequest)
	if !ok {
		t.Fatalf("apiReq.Body type = %T, want *AuthorizeRequest", req.apiReq.Body)
	}
	if bodyReq.ClientId == nil || *bodyReq.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", bodyReq.ClientId)
	}
	// 未设置的字段应为 nil
	if bodyReq.ClientSecret != nil {
		t.Errorf("Body.ClientSecret = %v, want nil", bodyReq.ClientSecret)
	}
	if bodyReq.Timestamp != nil {
		t.Errorf("Body.Timestamp = %v, want nil", bodyReq.Timestamp)
	}
}

// --- Authorize 资源方法测试 ---

func newAuthTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestAuthorize_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/Auth/authorize" {
			t.Errorf("expected path /river/Auth/authorize, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"access_token": "access_token_demo",
			"expires_in": 7200,
			"token_type": "Bearer",
			"scope": "all",
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().
			ClientId("test_client").
			ClientSecret("test_secret").
			GrantType("client_credentials").
			Timestamp(1583484681).
			Sign("test_sign").
			Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.Errno == nil || *resp.AuthorizeApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.AuthorizeApiReply.Errno)
	}
	if resp.AuthorizeApiReply.AccessToken == nil || *resp.AuthorizeApiReply.AccessToken != "access_token_demo" {
		t.Errorf("AccessToken = %v, want access_token_demo", resp.AuthorizeApiReply.AccessToken)
	}
	if resp.AuthorizeApiReply.ExpiresIn == nil || *resp.AuthorizeApiReply.ExpiresIn != 7200 {
		t.Errorf("ExpiresIn = %v, want 7200", resp.AuthorizeApiReply.ExpiresIn)
	}
	if resp.AuthorizeApiReply.TokenType == nil || *resp.AuthorizeApiReply.TokenType != "Bearer" {
		t.Errorf("TokenType = %v, want Bearer", resp.AuthorizeApiReply.TokenType)
	}
	if resp.AuthorizeApiReply.Scope == nil || *resp.AuthorizeApiReply.Scope != "all" {
		t.Errorf("Scope = %v, want all", resp.AuthorizeApiReply.Scope)
	}
}

func TestAuthorize_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// 空 data 场景：仅返回 errno/errmsg，无 access_token 等业务字段
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.Errno == nil || *resp.AuthorizeApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.AuthorizeApiReply.Errno)
	}
	// 无业务数据
	if resp.AuthorizeApiReply.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", resp.AuthorizeApiReply.AccessToken)
	}
	if resp.AuthorizeApiReply.ExpiresIn != nil {
		t.Errorf("ExpiresIn = %v, want nil", resp.AuthorizeApiReply.ExpiresIn)
	}
}

func TestAuthorize_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.Errno == nil || *resp.AuthorizeApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.AuthorizeApiReply.Errno)
	}
}

func TestAuthorize_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.AuthorizeApiReply != nil {
		t.Errorf("AuthorizeApiReply should be nil for non-200 response")
	}
}

func TestAuthorize_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","access_token":"token_aes128","expires_in":7200,"token_type":"Bearer","scope":"all","request_id":"req_enc"}`
	key := []byte("16byte-key-12345") // 必须 16 字节
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted) // AES128 用 StdEncoding

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.Errno == nil || *resp.AuthorizeApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.AuthorizeApiReply.Errno)
	}
	if resp.AuthorizeApiReply.AccessToken == nil || *resp.AuthorizeApiReply.AccessToken != "token_aes128" {
		t.Errorf("AccessToken = %v, want token_aes128", resp.AuthorizeApiReply.AccessToken)
	}
	if resp.AuthorizeApiReply.ExpiresIn == nil || *resp.AuthorizeApiReply.ExpiresIn != 7200 {
		t.Errorf("ExpiresIn = %v, want 7200", resp.AuthorizeApiReply.ExpiresIn)
	}
}

func TestAuthorize_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","access_token":"token_aes256","expires_in":3600,"token_type":"Bearer","request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk") // 必须 32 字节
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted) // AES256 用 URLEncoding

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.AccessToken == nil || *resp.AuthorizeApiReply.AccessToken != "token_aes256" {
		t.Errorf("AccessToken = %v, want token_aes256", resp.AuthorizeApiReply.AccessToken)
	}
	if resp.AuthorizeApiReply.ExpiresIn == nil || *resp.AuthorizeApiReply.ExpiresIn != 3600 {
		t.Errorf("ExpiresIn = %v, want 3600", resp.AuthorizeApiReply.ExpiresIn)
	}
}

func TestAuthorize_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","access_token":"token_noenc","request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	a := &auth{option: option}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.AccessToken == nil || *resp.AuthorizeApiReply.AccessToken != "token_noenc" {
		t.Errorf("AccessToken = %v, want token_noenc", resp.AuthorizeApiReply.AccessToken)
	}
}

func TestAuthorize_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","access_token":"token_opt","request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newAuthTestOption(testServer.URL)
	a := &auth{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewAuthorizeApiReqBuilder().
		AuthorizeRequest(NewAuthorizeRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := a.Authorize(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("Authorize() error = %v", err)
	}
	if resp.AuthorizeApiReply == nil {
		t.Fatal("AuthorizeApiReply is nil")
	}
	if resp.AuthorizeApiReply.AccessToken == nil || *resp.AuthorizeApiReply.AccessToken != "token_opt" {
		t.Errorf("AccessToken = %v, want token_opt", resp.AuthorizeApiReply.AccessToken)
	}
}
