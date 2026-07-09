package v1

import (
	"encoding/json"
	"testing"
)

func TestExtendFieldItem_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		wantId    int32
		wantCode  string
		wantValue string
	}{
		{
			name:      "full fields",
			jsonStr:   `{"id":1,"code":"custom_field","value":"hello"}`,
			wantId:    1,
			wantCode:  "custom_field",
			wantValue: "hello",
		},
		{
			name:      "empty object",
			jsonStr:   `{}`,
			wantId:    0,
			wantCode:  "",
			wantValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item ExtendFieldItem
			if err := json.Unmarshal([]byte(tt.jsonStr), &item); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if tt.wantId != 0 {
				if item.Id == nil || *item.Id != tt.wantId {
					t.Errorf("Id = %v, want %v", item.Id, tt.wantId)
				}
			} else {
				if item.Id != nil {
					t.Errorf("Id = %v, want nil", item.Id)
				}
			}
			if tt.wantCode != "" {
				if item.Code == nil || *item.Code != tt.wantCode {
					t.Errorf("Code = %v, want %v", item.Code, tt.wantCode)
				}
			} else {
				if item.Code != nil {
					t.Errorf("Code = %v, want nil", item.Code)
				}
			}
			if tt.wantValue != "" {
				if item.Value == nil || *item.Value != tt.wantValue {
					t.Errorf("Value = %v, want %v", item.Value, tt.wantValue)
				}
			} else {
				if item.Value != nil {
					t.Errorf("Value = %v, want nil", item.Value)
				}
			}
		})
	}
}

func TestExtendFieldItemBuilder(t *testing.T) {
	item := NewExtendFieldItemBuilder().
		Id(1).
		Code("custom_field").
		Value("hello").
		Build()

	if item.Id == nil || *item.Id != 1 {
		t.Errorf("Id = %v, want 1", item.Id)
	}
	if item.Code == nil || *item.Code != "custom_field" {
		t.Errorf("Code = %v, want custom_field", item.Code)
	}
	if item.Value == nil || *item.Value != "hello" {
		t.Errorf("Value = %v, want hello", item.Value)
	}
}
