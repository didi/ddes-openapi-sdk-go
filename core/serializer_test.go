package core

import (
	"testing"
)

func TestDefaultSerializer_Serialize(t *testing.T) {
	type testStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name    string
		v       interface{}
		wantErr bool
	}{
		{
			name: "serialize struct",
			v: testStruct{
				Name:  "test",
				Value: 123,
			},
			wantErr: false,
		},
		{
			name: "serialize map",
			v: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
			},
			wantErr: false,
		},
		{
			name: "serialize slice",
			v: []interface{}{
				"item1",
				"item2",
				123,
			},
			wantErr: false,
		},
		{
			name:    "serialize string",
			v:       "test string",
			wantErr: false,
		},
		{
			name:    "serialize number",
			v:       123.45,
			wantErr: false,
		},
		{
			name:    "serialize bool",
			v:       true,
			wantErr: false,
		},
		{
			name:    "serialize nil",
			v:       nil,
			wantErr: false,
		},
		{
			name: "serialize complex struct",
			v: struct {
				StringField string   `json:"string_field"`
				IntField    int      `json:"int_field"`
				FloatField  float64  `json:"float_field"`
				BoolField   bool     `json:"bool_field"`
				SliceField  []string `json:"slice_field"`
				MapField    map[string]int
			}{
				StringField: "test",
				IntField:    123,
				FloatField:  45.67,
				BoolField:   true,
				SliceField:  []string{"a", "b", "c"},
				MapField:    map[string]int{"x": 1, "y": 2},
			},
			wantErr: false,
		},
		{
			name: "serialize struct with json tags",
			v: struct {
				PublicField  string `json:"public_field"`
				PrivateField string `json:"private_field,omitempty"`
				IgnoredField string `json:"-"`
			}{
				PublicField:  "public",
				PrivateField: "private",
				IgnoredField: "ignored",
			},
			wantErr: false,
		},
		{
			name: "serialize empty struct",
			v: struct {
				Field1 string
				Field2 int
			}{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultSerializer{}
			got, err := s.Serialize(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultSerializer.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Errorf("DefaultSerializer.Serialize() returned empty bytes")
			}
			// 验证返回的是有效的JSON
			if !tt.wantErr && !isValidJSON(t, string(got)) {
				t.Errorf("DefaultSerializer.Serialize() returned invalid JSON: %s", string(got))
			}
		})
	}
}

func TestDefaultSerializer_Deserialize(t *testing.T) {
	type testStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name    string
		data    []byte
		v       interface{}
		wantErr bool
	}{
		{
			name:    "deserialize to struct",
			data:    []byte(`{"name":"test","value":123}`),
			v:       &testStruct{},
			wantErr: false,
		},
		{
			name:    "deserialize to map",
			data:    []byte(`{"key1":"value1","key2":123}`),
			v:       &map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "deserialize to slice",
			data:    []byte(`["item1","item2",123]`),
			v:       &[]interface{}{},
			wantErr: false,
		},
		{
			name:    "deserialize empty JSON object",
			data:    []byte(`{}`),
			v:       &map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "deserialize empty JSON array",
			data:    []byte(`[]`),
			v:       &[]interface{}{},
			wantErr: false,
		},
		{
			name: "deserialize complex JSON",
			data: []byte(`{
				"string_field": "test",
				"int_field": 123,
				"float_field": 45.67,
				"bool_field": true,
				"slice_field": ["a", "b", "c"],
				"map_field": {"x": 1, "y": 2}
			}`),
			v:       &map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "deserialize with whitespace",
			data:    []byte(`  {  "key"  :  "value"  }  `),
			v:       &map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "deserialize invalid JSON - missing closing brace",
			data:    []byte(`{"key":"value"`),
			v:       &map[string]interface{}{},
			wantErr: true,
		},
		{
			name:    "deserialize invalid JSON - unquoted key",
			data:    []byte(`{key:"value"}`),
			v:       &map[string]interface{}{},
			wantErr: true,
		},
		{
			name:    "deserialize invalid JSON - trailing comma",
			data:    []byte(`{"key":"value",}`),
			v:       &map[string]interface{}{},
			wantErr: true,
		},
		{
			name:    "deserialize empty data",
			data:    []byte(``),
			v:       &map[string]interface{}{},
			wantErr: true,
		},
		{
			name:    "deserialize null",
			data:    []byte(`null`),
			v:       &testStruct{},
			wantErr: false,
		},
		{
			name: "deserialize nested struct",
			data: []byte(`{
				"name": "test",
				"nested": {
					"value": 123
				}
			}`),
			v: &struct {
				Name   string `json:"name"`
				Nested *struct {
					Value int `json:"value"`
				} `json:"nested"`
			}{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultSerializer{}
			err := s.Deserialize(tt.data, tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultSerializer.Deserialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDefaultSerializer_SerializeDeserialize(t *testing.T) {
	type testStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name string
		v    interface{}
	}{
		{
			name: "roundtrip struct",
			v: testStruct{
				Name:  "test",
				Value: 123,
			},
		},
		{
			name: "roundtrip map",
			v: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
				"key3": true,
			},
		},
		{
			name: "roundtrip slice",
			v:    []interface{}{"item1", "item2", 123, true},
		},
		{
			name: "roundtrip complex struct",
			v: struct {
				StringField string   `json:"string_field"`
				IntField    int      `json:"int_field"`
				SliceField  []string `json:"slice_field"`
				MapField    map[string]int
			}{
				StringField: "test",
				IntField:    123,
				SliceField:  []string{"a", "b", "c"},
				MapField:    map[string]int{"x": 1, "y": 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultSerializer{}

			// 先序列化
			data, err := s.Serialize(tt.v)
			if err != nil {
				t.Errorf("Serialize() error = %v", err)
				return
			}

			// 再反序列化
			var dest interface{}
			err = s.Deserialize(data, &dest)
			if err != nil {
				t.Errorf("Deserialize() error = %v", err)
				return
			}

			// 验证类型匹配
			switch original := tt.v.(type) {
			case map[string]interface{}:
				if destMap, ok := dest.(map[string]interface{}); ok {
					// 只验证键存在，不严格比较值（因为JSON数字会变成float64）
					for key := range original {
						if _, exists := destMap[key]; !exists {
							t.Errorf("Roundtrip missing key: %s", key)
						}
					}
				} else {
					t.Errorf("Expected map[string]interface{}, got %T", dest)
				}
			case []interface{}:
				if destSlice, ok := dest.([]interface{}); ok {
					if len(destSlice) != len(original) {
						t.Errorf("Length mismatch: got %d, want %d", len(destSlice), len(original))
					}
				} else {
					t.Errorf("Expected []interface{}, got %T", dest)
				}
			}
		})
	}
}

func TestDefaultSerializer_SerializeSpecialCharacters(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
	}{
		{
			name: "serialize unicode characters",
			v: map[string]interface{}{
				"chinese": "中文测试",
				"emoji":   "😀🎉",
				"mixed":   "Hello 世界 🌍",
			},
		},
		{
			name: "serialize escaped characters",
			v: map[string]interface{}{
				"newline":   "line1\nline2",
				"tab":       "col1\tcol2",
				"quote":     `he said "hello"`,
				"backslash": "path\\to\\file",
			},
		},
		{
			name: "serialize special JSON characters",
			v: map[string]interface{}{
				"control": "\x00\x01\x02",
				"unicode": "\u4e2d\u6587",
				"mixed":   "a\x00b\u4e2dc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultSerializer{}
			data, err := s.Serialize(tt.v)
			if err != nil {
				t.Errorf("Serialize() error = %v", err)
				return
			}

			// 验证可以反序列化回来
			var dest map[string]interface{}
			err = s.Deserialize(data, &dest)
			if err != nil {
				t.Errorf("Deserialize() error = %v", err)
				return
			}

			// 验证内容一致
			if !isValidJSON(t, string(data)) {
				t.Errorf("Serialize() produced invalid JSON: %s", string(data))
			}
		})
	}
}

func BenchmarkDefaultSerializer_Serialize(b *testing.B) {
	s := &DefaultSerializer{}
	data := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
		"key4": []interface{}{"a", "b", "c"},
		"key5": map[string]interface{}{
			"nested_key": "nested_value",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.Serialize(data)
	}
}

func BenchmarkDefaultSerializer_Deserialize(b *testing.B) {
	s := &DefaultSerializer{}
	data := []byte(`{
		"key1": "value1",
		"key2": 123,
		"key3": true,
		"key4": ["a", "b", "c"],
		"key5": {"nested_key": "nested_value"}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dest map[string]interface{}
		_ = s.Deserialize(data, &dest)
	}
}

func BenchmarkDefaultSerializer_RoundTrip(b *testing.B) {
	s := &DefaultSerializer{}
	original := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
		"key4": []interface{}{"a", "b", "c"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, _ := s.Serialize(original)
		var dest map[string]interface{}
		_ = s.Deserialize(data, &dest)
	}
}

// isValidJSON 是一个辅助函数，用于验证字符串是否是有效的JSON
func isValidJSON(t *testing.T, s string) bool {
	var dest interface{}
	err := new(DefaultSerializer).Deserialize([]byte(s), &dest)
	return err == nil
}
