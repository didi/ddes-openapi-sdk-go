package core

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestRequestTranslator_translate(t *testing.T) {
	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        "http://example.com",
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	translator := &RequestTranslator{}

	tests := []struct {
		name    string
		apiReq  *ApiReq
		wantErr bool
	}{
		{
			name: "GET request with query params",
			apiReq: &ApiReq{
				HttpMethod: http.MethodGet,
				ApiPath:    "/test",
				QueryParams: url.Values{
					"key1": []string{"value1"},
				},
			},
			wantErr: false,
		},
		{
			name: "POST request with body",
			apiReq: &ApiReq{
				HttpMethod: http.MethodPost,
				ApiPath:    "/test",
				Body: map[string]interface{}{
					"key1": "value1",
				},
			},
			wantErr: false,
		},
		{
			name: "PUT request with body",
			apiReq: &ApiReq{
				HttpMethod: http.MethodPut,
				ApiPath:    "/test",
				Body: map[string]interface{}{
					"key1": "value1",
				},
			},
			wantErr: false,
		},
		{
			name: "DELETE request with query params",
			apiReq: &ApiReq{
				HttpMethod: http.MethodDelete,
				ApiPath:    "/test",
				QueryParams: url.Values{
					"key1": []string{"value1"},
				},
			},
			wantErr: false,
		},
		{
			name: "PATCH request with body",
			apiReq: &ApiReq{
				HttpMethod: http.MethodPatch,
				ApiPath:    "/test",
				Body: map[string]interface{}{
					"key1": "value1",
				},
			},
			wantErr: false,
		},
		{
			name: "request with path params",
			apiReq: &ApiReq{
				HttpMethod: http.MethodGet,
				ApiPath:    "/test/{id}",
				PathParams: map[string]string{
					"id": "123",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := translator.translate(ctx, tt.apiReq, option, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestTranslator.translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("RequestTranslator.translate() returned nil request")
				}
				if got.Method != tt.apiReq.HttpMethod {
					t.Errorf("RequestTranslator.translate() method = %v, want %v", got.Method, tt.apiReq.HttpMethod)
				}
				if got.Header.Get("Content-Type") != "application/json" {
					t.Errorf("RequestTranslator.translate() Content-Type = %v, want application/json", got.Header.Get("Content-Type"))
				}
				if got.Header.Get("Accept") != "application/json" {
					t.Errorf("RequestTranslator.translate() Accept = %v, want application/json", got.Header.Get("Accept"))
				}
			}
		})
	}
}

func TestAssemblyBodyMapToEncryptionDataMap(t *testing.T) {
	type args struct {
		apiReq  *ApiReq
		dataMap map[string]interface{}
		option  *Option
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "encrypt body data with AES128",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/test",
				},
				dataMap: map[string]interface{}{
					"client_id":    "test_client",
					"access_token": "test_token",
					"data":         "sensitive_data",
				},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
		{
			name: "non-encryption API should not encrypt",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/river/Auth/authorize",
				},
				dataMap: map[string]interface{}{
					"client_id":     "test_client",
					"client_secret": "test_secret",
				},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := &RequestTranslator{}
			got, err := translator.AssemblyBodyMapToEncryptionDataMap(tt.args.apiReq, tt.args.dataMap, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestTranslator.AssemblyBodyMapToEncryptionDataMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("RequestTranslator.AssemblyBodyMapToEncryptionDataMap() returned nil")
				}
			}
		})
	}
}

func TestAssemblyQueryParamsToEncryptionUrlValue(t *testing.T) {
	type args struct {
		apiReq      *ApiReq
		queryParams url.Values
		option      *Option
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "encrypt query params",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/test",
				},
				queryParams: url.Values{
					"client_id":    []string{"test_client"},
					"access_token": []string{"test_token"},
					"data":         []string{"sensitive_data"},
				},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
		{
			name: "non-encryption API should not encrypt",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/river/Auth/authorize",
				},
				queryParams: url.Values{
					"client_id":     []string{"test_client"},
					"client_secret": []string{"test_secret"},
				},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
		{
			name: "query params with multiple values",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/test",
				},
				queryParams: url.Values{
					"key1": []string{"value1", "value2"},
				},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
		{
			name: "empty query params",
			args: args{
				apiReq: &ApiReq{
					ApiPath: "/test",
				},
				queryParams: url.Values{},
				option: &Option{
					EncryptionOption: &EncryptionOption{
						Ent: 1,
						Key: "16byte-key-12345", // 16字节密钥
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := &RequestTranslator{}
			got, err := translator.AssemblyQueryParamsToEncryptionUrlValue(tt.args.apiReq, tt.args.queryParams, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestTranslator.AssemblyQueryParamsToEncryptionUrlValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("RequestTranslator.AssemblyQueryParamsToEncryptionUrlValue() returned nil")
				}
			}
		})
	}
}

func TestRequestTranslator_WithEncryption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer testServer.Close()

	option := &Option{
		clientId:         "test_client_id",
		clientSecret:     "test_client_secret",
		signKey:          "test_sign_key",
		BaseUrl:          testServer.URL,
		RequestTimeOut:   5 * time.Second,
		SignMethod:       1,
		EnableEncryption: true,
		EncryptionOption: &EncryptionOption{
			Ent: 1,
			Key: "16byte-key-12345", // 16字节密钥
		},
		Logger: NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/test",
		Body: map[string]interface{}{
			"client_id": "test_client",
			"data":      "sensitive_data",
		},
	}

	ctx := context.Background()
	_, err := Request(ctx, apiReq, option, nil)
	if err != nil {
		t.Errorf("Request() with encryption error = %v", err)
	}
}

func TestRequestTranslator_WithSHA256(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer testServer.Close()

	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        testServer.URL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     2, // SHA256
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/test",
		Body: map[string]interface{}{
			"key1": "value1",
		},
	}

	ctx := context.Background()
	_, err := Request(ctx, apiReq, option, nil)
	if err != nil {
		t.Errorf("Request() with SHA256 error = %v", err)
	}
}

func TestRequestTranslator_EncryptionResultFormat(t *testing.T) {
	translator := &RequestTranslator{}
	option := &Option{
		EncryptionOption: &EncryptionOption{
			Ent: 1,
			Key: "16byte-key-12345", // 16字节密钥
		},
	}

	apiReq := &ApiReq{
		ApiPath: "/test",
	}

	dataMap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	result, err := translator.AssemblyBodyMapToEncryptionDataMap(apiReq, dataMap, option)
	if err != nil {
		t.Errorf("AssemblyBodyMapToEncryptionDataMap() error = %v", err)
		return
	}

	// 验证返回的map包含ent字段
	if ent, exists := result["ent"]; !exists {
		t.Errorf("Expected 'ent' field in result")
	} else if entInt, ok := ent.(int); !ok || entInt != 1 {
		t.Errorf("Expected ent=1, got %v", ent)
	}

	// 验证返回的map包含encrypt_content字段
	if encContent, exists := result["encrypt_content"]; !exists {
		t.Errorf("Expected 'encrypt_content' field in result")
	} else if encStr, ok := encContent.(string); !ok {
		t.Errorf("Expected encrypt_content to be string, got %T", encContent)
	} else {
		// 验证是有效的base64
		_, err := base64.StdEncoding.DecodeString(encStr)
		if err != nil {
			t.Errorf("encrypt_content is not valid base64: %v", err)
		}
	}
}

func TestRequestTranslator_NotEncryptionKeys(t *testing.T) {
	translator := &RequestTranslator{}
	option := &Option{
		EncryptionOption: &EncryptionOption{
			Ent: 1,
			Key: "16byte-key-12345", // 16字节密钥
		},
	}

	apiReq := &ApiReq{
		ApiPath: "/test",
	}

	// 测试不加密的字段
	dataMap := map[string]interface{}{
		"client_id":    "test_client",
		"access_token": "test_token",
		"sensitive":    "sensitive_data",
	}

	result, err := translator.AssemblyBodyMapToEncryptionDataMap(apiReq, dataMap, option)
	if err != nil {
		t.Errorf("AssemblyBodyMapToEncryptionDataMap() error = %v", err)
		return
	}

	// 验证不加密的字段仍然存在
	if clientId, exists := result["client_id"]; !exists {
		t.Errorf("Expected 'client_id' field to not be encrypted")
	} else if clientId.(string) != "test_client" {
		t.Errorf("client_id should not be encrypted, got %v", clientId)
	}

	if accessToken, exists := result["access_token"]; !exists {
		t.Errorf("Expected 'access_token' field to not be encrypted")
	} else if accessToken.(string) != "test_token" {
		t.Errorf("access_token should not be encrypted, got %v", accessToken)
	}

	// 验证敏感字段被加密了
	if _, exists := result["sensitive"]; exists {
		t.Errorf("Expected 'sensitive' field to be encrypted (should not exist in plain form)")
	}
}
