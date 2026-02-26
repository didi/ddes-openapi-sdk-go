package core

import (
	"net/url"
	"testing"
)

func TestSigner_SignQueryParams(t *testing.T) {
	type args struct {
		params url.Values
		option *Option
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple query params with MD5",
			args: args{
				params: url.Values{
					"key1": []string{"value1"},
					"key2": []string{"value2"},
				},
				option: &Option{
					signKey:    "test_sign_key",
					SignMethod: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "query params with sign field should be removed",
			args: args{
				params: url.Values{
					"key1":     []string{"value1"},
					"sign":     []string{"old_sign"},
					"sign_key": []string{"test_key"},
				},
				option: &Option{
					signKey:    "test_sign_key",
					SignMethod: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "query params with multiple values",
			args: args{
				params: url.Values{
					"key1": []string{"value1", "value2", "value3"},
				},
				option: &Option{
					signKey:    "test_sign_key",
					SignMethod: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "query params with SHA256",
			args: args{
				params: url.Values{
					"key1": []string{"value1"},
					"key2": []string{"value2"},
				},
				option: &Option{
					signKey:    "test_sign_key",
					SignMethod: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "empty params",
			args: args{
				params: url.Values{},
				option: &Option{
					signKey:    "test_sign_key",
					SignMethod: 1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signer := &Signer{}
			got, err := signer.SignQueryParams(tt.args.params, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.SignQueryParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("Signer.SignQueryParams() returned empty string")
			}
		})
	}
}

func TestSigner_SignPostBody(t *testing.T) {
	tests := []struct {
		name    string
		option  *Option
		m       map[string]interface{}
		wantLen int
	}{
		{
			name: "simple body with MD5",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			wantLen: 3,
		},
		{
			name: "body with various types",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m: map[string]interface{}{
				"string_key":  "string_value",
				"int_key":     123,
				"int64_key":   int64(456),
				"float64_key": 78.9,
			},
			wantLen: 5,
		},
		{
			name: "body with array",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m: map[string]interface{}{
				"key1":      "value1",
				"array_key": []interface{}{"a", "b", "c"},
			},
			wantLen: 3,
		},
		{
			name: "body with object suffix",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m: map[string]interface{}{
				"key1":       "value1",
				"obj__obj__": map[string]interface{}{"sub": "value"},
			},
			wantLen: 3,
		},
		{
			name: "body with SHA256",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 2,
			},
			m: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			wantLen: 3,
		},
		{
			name: "body with existing sign",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m: map[string]interface{}{
				"key1": "value1",
				"sign": "old_sign",
			},
			wantLen: 2,
		},
		{
			name: "body with empty map",
			option: &Option{
				signKey:    "test_sign_key",
				SignMethod: 1,
			},
			m:       map[string]interface{}{},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signer := &Signer{}
			got := signer.SignPostBody(tt.option, tt.m)
			if got == nil {
				t.Errorf("Signer.SignPostBody() returned nil")
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Signer.SignPostBody() returned map with length = %v, want %v", len(got), tt.wantLen)
			}
			if _, exists := got["sign"]; !exists {
				t.Errorf("Signer.SignPostBody() returned map without sign field")
			}
			if _, exists := got["sign_key"]; exists {
				t.Errorf("Signer.SignPostBody() returned map with sign_key field (should be removed)")
			}
		})
	}
}

func TestSignRequest(t *testing.T) {
	type args struct {
		signKey string
		body    interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple body",
			args: args{
				signKey: "test_sign_key",
				body: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			wantErr: false,
		},
		{
			name: "body with existing sign",
			args: args{
				signKey: "test_sign_key",
				body: map[string]interface{}{
					"key1": "value1",
					"sign": "old_sign",
				},
			},
			wantErr: false,
		},
		{
			name: "body with various types",
			args: args{
				signKey: "test_sign_key",
				body: map[string]interface{}{
					"string_key":  "string_value",
					"int_key":     123,
					"int64_key":   int64(456),
					"float64_key": 78.9,
				},
			},
			wantErr: false,
		},
		{
			name: "struct body",
			args: args{
				signKey: "test_sign_key",
				body: struct {
					Key1 string `json:"key1"`
					Key2 string `json:"key2"`
				}{
					Key1: "value1",
					Key2: "value2",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid body",
			args: args{
				signKey: "test_sign_key",
				body:    func() {}, // invalid type
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignRequest(tt.args.signKey, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("SignRequest() returned empty string")
			}
		})
	}
}

func BenchmarkSigner_SignQueryParams(b *testing.B) {
	signer := &Signer{}
	params := url.Values{
		"key1": []string{"value1"},
		"key2": []string{"value2"},
		"key3": []string{"value3"},
	}
	option := &Option{
		signKey:    "test_sign_key",
		SignMethod: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = signer.SignQueryParams(params, option)
	}
}

func BenchmarkSigner_SignPostBody(b *testing.B) {
	signer := &Signer{}
	option := &Option{
		signKey:    "test_sign_key",
		SignMethod: 1,
	}
	m := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = signer.SignPostBody(option, m)
	}
}
