package core

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLocalTokenManager_SetToken(t *testing.T) {
	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}
	manager := NewLocalTokenManager(mockCache)

	tests := []struct {
		name    string
		token   string
		ttl     time.Duration
		wantErr bool
	}{
		{
			name:    "set token with TTL",
			token:   "test_token_123",
			ttl:     1 * time.Hour,
			wantErr: false,
		},
		{
			name:    "set empty token",
			token:   "",
			ttl:     1 * time.Hour,
			wantErr: false,
		},
		{
			name:    "set token with zero TTL",
			token:   "test_token_456",
			ttl:     0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager.SetToken(tt.token, tt.ttl)
			// 不返回错误，只是验证没有panic
		})
	}
}

func TestLocalTokenManager_GetAccessToken(t *testing.T) {
	tests := []struct {
		name       string
		option     *Option
		mockCache  *mockTokenCache
		serverResp *http.Response
		serverErr  error
		wantErr    bool
		setup      func(*mockTokenCache)
	}{
		{
			name: "get token from cache",
			option: &Option{
				clientId:     "test_client",
				clientSecret: "test_secret",
				BaseUrl:      "http://example.com",
			},
			mockCache: &mockTokenCache{
				tokens: make(map[string]string),
			},
			setup: func(m *mockTokenCache) {
				m.tokens["cached"] = "cached_token_123"
			},
			wantErr: false,
		},
		{
			name:   "nil option",
			option: nil,
			mockCache: &mockTokenCache{
				tokens: make(map[string]string),
			},
			wantErr: true,
		},
		{
			name: "nil cache",
			option: &Option{
				clientId:     "test_client",
				clientSecret: "test_secret",
			},
			mockCache: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &LocalTokenManager{
				tokenCache: tt.mockCache,
			}

			if tt.setup != nil && tt.mockCache != nil {
				tt.setup(tt.mockCache)
			}

			got, err := manager.GetAccessToken(tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalTokenManager.GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.setup != nil && got == "" {
				t.Errorf("LocalTokenManager.GetAccessToken() returned empty token")
			}
		})
	}
}

func TestLocalTokenManager_GetAccessToken_FromServer(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/river/Auth/authorize" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"access_token": "new_token_123",
				"expires_in": 3600,
				"token_type": "Bearer"
			}`))
		}
	}))
	defer testServer.Close()

	option := &Option{
		clientId:     "test_client",
		clientSecret: "test_secret",
		signKey:      "test_sign_key",
		BaseUrl:      testServer.URL,
		SignMethod:   1,
		Logger:       NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}

	// 使用NewLocalTokenManager设置全局tokenManager
	manager := NewLocalTokenManager(mockCache)

	token, err := manager.GetAccessToken(option)
	if err != nil {
		t.Errorf("LocalTokenManager.GetAccessToken() error = %v", err)
		return
	}

	if token != "new_token_123" {
		t.Errorf("LocalTokenManager.GetAccessToken() = %v, want %v", token, "new_token_123")
	}

	// 验证token被缓存了
	cachedToken, _ := mockCache.GetToken()
	if cachedToken != "new_token_123" {
		t.Errorf("Token not cached properly, got %v", cachedToken)
	}

	// 重置全局状态
	tokenManager.tokenCache = tokenCache
}

func TestLocalTokenManager_GetAccessToken_ServerError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/river/Auth/authorize" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{
				"errno": 401,
				"errmsg": "unauthorized"
			}`))
		}
	}))
	defer testServer.Close()

	option := &Option{
		clientId:     "test_client",
		clientSecret: "test_secret",
		signKey:      "test_sign_key",
		BaseUrl:      testServer.URL,
		SignMethod:   1,
		Logger:       NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}

	manager := &LocalTokenManager{
		tokenCache: mockCache,
	}

	_, err := manager.GetAccessToken(option)
	if err == nil {
		t.Errorf("LocalTokenManager.GetAccessToken() expected error, got nil")
	}
}

func TestLocalTokenManager_GetAccessToken_InvalidResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/river/Auth/authorize" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"access_token": "token123",
				"expires_in": "invalid"
			}`))
		}
	}))
	defer testServer.Close()

	option := &Option{
		clientId:     "test_client",
		clientSecret: "test_secret",
		signKey:      "test_sign_key",
		BaseUrl:      testServer.URL,
		SignMethod:   1,
		Logger:       NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}

	manager := &LocalTokenManager{
		tokenCache: mockCache,
	}

	_, err := manager.GetAccessToken(option)
	if err == nil {
		t.Errorf("LocalTokenManager.GetAccessToken() expected error for invalid response, got nil")
	}
}

func TestApiReqParamsAssembly(t *testing.T) {
	type args struct {
		apiReq  *ApiReq
		option  *Option
		wantErr bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "assemble params for POST request",
			args: args{
				apiReq: &ApiReq{
					HttpMethod: http.MethodPost,
					ApiPath:    "/test",
					Body: map[string]interface{}{
						"data": "test",
					},
				},
				option: &Option{
					clientId:     "test_client",
					clientSecret: "test_secret",
				},
				wantErr: false,
			},
		},
		{
			name: "assemble params for GET request",
			args: args{
				apiReq: &ApiReq{
					HttpMethod:  http.MethodGet,
					ApiPath:     "/test",
					QueryParams: map[string][]string{},
				},
				option: &Option{
					clientId:     "test_client",
					clientSecret: "test_secret",
				},
				wantErr: false,
			},
		},
		{
			name: "nil apiReq",
			args: args{
				apiReq:  nil,
				option:  &Option{},
				wantErr: false,
			},
		},
		{
			name: "empty api path",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "",
				},
				option:  &Option{},
				wantErr: false,
			},
		},
		{
			name: "nil option",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/test",
				},
				option:  nil,
				wantErr: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ApiReqParamsAssembly(tt.args.apiReq, tt.args.option)
			if (err != nil) != tt.args.wantErr {
				t.Errorf("ApiReqParamsAssembly() error = %v, wantErr %v", err, tt.args.wantErr)
			}
		})
	}
}

func TestApiReqParamsAssembly_WithTokenCache(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/river/Auth/authorize" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"access_token": "test_token_123",
				"expires_in": 3600
			}`))
		}
	}))
	defer testServer.Close()

	tokenCache := NewLocalTokenCache()

	option := &Option{
		clientId:         "test_client",
		clientSecret:     "test_secret",
		signKey:          "test_sign_key",
		BaseUrl:          testServer.URL,
		SignMethod:       1,
		EnableTokenCache: true,
		TokenCache:       tokenCache,
		Logger:           NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	// 设置全局tokenManager使用这个cache
	NewLocalTokenManager(tokenCache)

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/test",
		Body: map[string]interface{}{
			"data": "test",
		},
	}

	err := ApiReqParamsAssembly(apiReq, option)
	if err != nil {
		t.Errorf("ApiReqParamsAssembly() error = %v", err)
	}

	// 验证access_token被添加到body
	if bodyMap, ok := apiReq.Body.(map[string]interface{}); ok {
		if token, exists := bodyMap["access_token"]; !exists {
			t.Errorf("access_token not added to body")
		} else if token.(string) != "test_token_123" {
			t.Errorf("access_token = %v, want %v", token, "test_token_123")
		}
	} else {
		t.Errorf("Body is not map[string]interface{}")
	}

	// 重置全局状态
	tokenManager.tokenCache = tokenCache
}

func TestApiReqParamsAssembly_AuthorizeEndpoint(t *testing.T) {
	option := &Option{
		clientId:     "test_client",
		clientSecret: "test_secret",
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/river/Auth/authorize",
		Body: map[string]interface{}{
			"grant_type": "client_credentials",
		},
	}

	err := ApiReqParamsAssembly(apiReq, option)
	if err != nil {
		t.Errorf("ApiReqParamsAssembly() error = %v", err)
	}

	// 验证client_id和client_secret都被添加到body
	bodyMap, ok := apiReq.Body.(map[string]interface{})
	if !ok {
		t.Errorf("Body is not map[string]interface{}")
		return
	}

	if clientId, exists := bodyMap["client_id"]; !exists {
		t.Errorf("client_id not added to body")
	} else if clientId.(string) != "test_client" {
		t.Errorf("client_id = %v, want %v", clientId, "test_client")
	}

	if clientSecret, exists := bodyMap["client_secret"]; !exists {
		t.Errorf("client_secret not added to body")
	} else if clientSecret.(string) != "test_secret" {
		t.Errorf("client_secret = %v, want %v", clientSecret, "test_secret")
	}
}

// mockTokenCache 是一个用于测试的mock缓存实现
type mockTokenCache struct {
	tokens map[string]string
}

func (m *mockTokenCache) SetToken(token string, ttl time.Duration) {
	if m == nil {
		return
	}
	if m.tokens == nil {
		m.tokens = make(map[string]string)
	}
	m.tokens["cached"] = token
}

func (m *mockTokenCache) GetToken() (string, error) {
	if m == nil {
		return "", errors.New("cache is nil")
	}
	if m.tokens == nil {
		return "", nil // 返回空字符串而不是错误，模拟localTokenCache的行为
	}
	token, exists := m.tokens["cached"]
	if !exists {
		return "", nil // 返回空字符串而不是错误，模拟localTokenCache的行为
	}
	return token, nil
}

func BenchmarkLocalTokenManager_SetToken(b *testing.B) {
	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}
	manager := NewLocalTokenManager(mockCache)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.SetToken("test_token", 1*time.Hour)
	}
}

func BenchmarkLocalTokenManager_GetToken(b *testing.B) {
	mockCache := &mockTokenCache{
		tokens: make(map[string]string),
	}
	manager := NewLocalTokenManager(mockCache)
	manager.SetToken("test_token", 1*time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = mockCache.GetToken()
	}
}
