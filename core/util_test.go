package core

import (
	"testing"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty string",
			input:    []byte(""),
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:     "simple string",
			input:    []byte("hello"),
			expected: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:     "string with special chars",
			input:    []byte("hello world!"),
			expected: "fc3ff98e8c6a0d3087d515c0473f8677",
		},
		{
			name:     "unicode string",
			input:    []byte("你好世界"),
			expected: "65396ee4aad0b4f17aacd1c6112ee364",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MD5(tt.input)
			if result != tt.expected {
				t.Errorf("MD5() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMD5Base64(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty string",
			input:    []byte(""),
			expected: "1B2M2Y8AsgTpgAmY7PhCfg==",
		},
		{
			name:     "simple string",
			input:    []byte("hello"),
			expected: "XUFAKrxLKna5cZ2REBfFkg==",
		},
		{
			name:     "string with special chars",
			input:    []byte("hello world!"),
			expected: "/D/5joxqDTCH1RXARz+Gdw==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MD5Base64(tt.input)
			if result != tt.expected {
				t.Errorf("MD5Base64() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSha256(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty string",
			input:    []byte(""),
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "simple string",
			input:    []byte("hello"),
			expected: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name:     "string with special chars",
			input:    []byte("hello world!"),
			expected: "7509e5bda0c762d2bac7f90d758b5b2263fa01ccbc542ab5e3df163be08e6ca9",
		},
		{
			name:     "unicode string",
			input:    []byte("你好世界"),
			expected: "beca6335b20ff57ccc47403ef4d9e0b8fccb4442b3151c2e7d50050673d43172",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha256(tt.input)
			if result != tt.expected {
				t.Errorf("Sha256() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGenMD5EncryptionKey(t *testing.T) {
	tests := []struct {
		name    string
		company string
	}{
		{
			name:    "test company",
			company: "test123",
		},
		{
			name:    "empty company",
			company: "",
		},
		{
			name:    "company with special chars",
			company: "test-company@2023",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenMD5EncryptionKey(tt.company)
			if result == "" {
				t.Errorf("GenMD5EncryptionKey() returned empty string")
			}
			t.Logf("GenMD5EncryptionKey(%q) = %q", tt.company, result)
		})
	}
}

func TestGenSha256EncryptionKey(t *testing.T) {
	tests := []struct {
		name    string
		company string
	}{
		{
			name:    "test company",
			company: "test123",
		},
		{
			name:    "empty company",
			company: "",
		},
		{
			name:    "company with special chars",
			company: "test-company@2023",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenSha256EncryptionKey(tt.company)
			if result == "" {
				t.Errorf("GenSha256EncryptionKey() returned empty string")
			}
			t.Logf("GenSha256EncryptionKey(%q) = %q", tt.company, result)
		})
	}
}

func TestUnescapeUnicode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple unicode",
			input:    "hello\\u4e16\\u754c",
			expected: "hello世界",
		},
		{
			name:     "mixed unicode and text",
			input:    "\\u4f60\\u597dworld\\uff01",
			expected: "你好world！",
		},
		{
			name:     "no unicode",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "incomplete unicode sequence",
			input:    "hello\\u4e1",
			expected: "hello\\u4e1",
		},
		{
			name:     "multiple unicode chars",
			input:    "\\u4e2d\\u6587\\u6d4b\\u8bd5",
			expected: "中文测试",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnescapeUnicode(tt.input)
			if result != tt.expected {
				t.Errorf("UnescapeUnicode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
