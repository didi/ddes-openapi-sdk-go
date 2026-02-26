package core

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAssemblyFullUrl(t *testing.T) {
	type args struct {
		baseUrl     string
		apiPath     string
		pathParams  map[string]string
		queryParams map[string][]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full url with path and query params",
			args: args{
				baseUrl: "http://127.0.0.1:8080",
				apiPath: "/river/auth/{name}/authorize/{age}",
				pathParams: map[string]string{
					"name": "test",
					"age":  "100",
				},
				queryParams: map[string][]string{
					"city":   {"chengdu", "chongqing"},
					"street": {"tianfuStreet"},
				},
			},
			want: "http://127.0.0.1:8080/river/auth/test/authorize/100?city=chengdu&city=chongqing&street=tianfuStreet",
		},
		{
			name: "base url with trailing slash",
			args: args{
				baseUrl:     "http://127.0.0.1:8080/",
				apiPath:     "/river/auth/authorize",
				pathParams:  nil,
				queryParams: nil,
			},
			want: "http://127.0.0.1:8080/river/auth/authorize",
		},
		{
			name: "url with special characters in path params",
			args: args{
				baseUrl: "http://example.com",
				apiPath: "/api/{id}/details",
				pathParams: map[string]string{
					"id": "test/user@123",
				},
				queryParams: nil,
			},
			want: "http://example.com/api/test%2Fuser@123/details",
		},
		{
			name: "empty path params and query params",
			args: args{
				baseUrl:     "http://example.com",
				apiPath:     "/api/test",
				pathParams:  nil,
				queryParams: nil,
			},
			want: "http://example.com/api/test",
		},
		{
			name: "only query params",
			args: args{
				baseUrl:    "http://example.com",
				apiPath:    "/api/test",
				pathParams: nil,
				queryParams: map[string][]string{
					"key1": {"value1"},
					"key2": {"value2"},
				},
			},
			want: "http://example.com/api/test?key1=value1&key2=value2",
		},
		{
			name: "only path params",
			args: args{
				baseUrl: "http://example.com",
				apiPath: "/api/{id}",
				pathParams: map[string]string{
					"id": "12345",
				},
				queryParams: nil,
			},
			want: "http://example.com/api/12345",
		},
		{
			name: "query params with special characters",
			args: args{
				baseUrl:    "http://example.com",
				apiPath:    "/api/test",
				pathParams: nil,
				queryParams: map[string][]string{
					"search": {"hello world"},
					"filter": {"status=active"},
				},
			},
			want: "http://example.com/api/test?filter=status%3Dactive&search=hello+world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssemblyFullUrl(tt.args.baseUrl, tt.args.apiPath, tt.args.pathParams, tt.args.queryParams); got != tt.want {
				t.Errorf("AssemblyFullUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest(t *testing.T) {
	// 创建一个测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查请求方法
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		// 检查请求头
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", contentType)
		}

		// 返回响应
		w.Header().Set("Content-Type", "application/json")
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
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/test",
		Body: map[string]interface{}{
			"test_key": "test_value",
		},
	}

	ctx := context.Background()
	resp, err := Request(ctx, apiReq, option, nil)
	if err != nil {
		t.Errorf("Request() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Request() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if string(resp.Body) != `{"result":"success"}` {
		t.Errorf("Request() body = %v, want %v", string(resp.Body), `{"result":"success"}`)
	}
}

func TestRequestWithContext(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer testServer.Close()

	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        testServer.URL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    "/test",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := Request(ctx, apiReq, option, nil)
	if err != nil {
		t.Errorf("Request() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Request() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRequestWithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查自定义header
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
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
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")

	reqOption := &ReqOption{
		Header: customHeader,
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    "/test",
		Body:       map[string]interface{}{},
	}

	ctx := context.Background()
	resp, err := Request(ctx, apiReq, option, reqOption)
	if err != nil {
		t.Errorf("Request() error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Request() status = %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRequestTimeout(t *testing.T) {
	// 创建一个会延迟响应的测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        testServer.URL,
		RequestTimeOut: 1 * time.Second,
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    "/test",
	}

	ctx := context.Background()
	_, err := Request(ctx, apiReq, option, nil)
	if err == nil {
		t.Errorf("Request() expected timeout error, got nil")
	}
}

func TestRequestErrorResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
	}))
	defer testServer.Close()

	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        testServer.URL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    "/test",
	}

	ctx := context.Background()
	resp, err := Request(ctx, apiReq, option, nil)
	if err != nil {
		t.Errorf("Request() unexpected error = %v", err)
		return
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Request() status = %v, want %v", resp.StatusCode, http.StatusBadRequest)
	}

	if string(resp.Body) != `{"error":"bad request"}` {
		t.Errorf("Request() body = %v, want %v", string(resp.Body), `{"error":"bad request"}`)
	}
}

func TestRequestNetworkError(t *testing.T) {
	option := &Option{
		clientId:       "test_client_id",
		clientSecret:   "test_client_secret",
		signKey:        "test_sign_key",
		BaseUrl:        "http://localhost:9999", // 假设没有服务监听此端口
		RequestTimeOut: 1 * time.Second,
		SignMethod:     1,
		Logger:         NewLoggerImpl(LogLevelInfo, NewDefaultLogger(LogLevelInfo)),
	}

	apiReq := &ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    "/test",
	}

	ctx := context.Background()
	_, err := Request(ctx, apiReq, option, nil)
	if err == nil {
		t.Errorf("Request() expected network error, got nil")
	}
}
