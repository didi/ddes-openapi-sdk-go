package core

import "testing"

func TestSmartDecodePreservesLargeInt64(t *testing.T) {
	const want int64 = 9007199254740993

	var decoded struct {
		Id *int64 `json:"id"`
	}
	if err := SmartDecode([]byte(`{"id":9007199254740993}`), &decoded); err != nil {
		t.Fatalf("SmartDecode() error = %v", err)
	}
	if decoded.Id == nil || *decoded.Id != want {
		t.Fatalf("decoded Id = %v, want %d", decoded.Id, want)
	}
}

func TestSmartDecodeAcceptsJsonNumberAndNumericString(t *testing.T) {
	for _, body := range []string{
		`{"id":9007199254740993}`,
		`{"id":"9007199254740993"}`,
	} {
		var decoded struct {
			Id *int64 `json:"id"`
		}
		if err := SmartDecode([]byte(body), &decoded); err != nil {
			t.Fatalf("SmartDecode(%s) error = %v", body, err)
		}
		if decoded.Id == nil || *decoded.Id != 9007199254740993 {
			t.Errorf("decoded Id = %v, want 9007199254740993", decoded.Id)
		}
	}
}

func TestSmartDecoderDecodeKeepsStandardExactInt64Path(t *testing.T) {
	var decoded struct {
		Id *int64 `json:"id"`
	}
	if err := NewSmartDecoder().Decode([]byte(`{"id":9007199254740993}`), &decoded); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if decoded.Id == nil || *decoded.Id != 9007199254740993 {
		t.Fatalf("decoded Id = %v, want 9007199254740993", decoded.Id)
	}
}

func TestSmartDecodeKeepsNullSliceNil(t *testing.T) {
	var decoded struct {
		Items []string `json:"items"`
	}
	if err := SmartDecode([]byte(`{"items":null}`), &decoded); err != nil {
		t.Fatalf("SmartDecode() error = %v", err)
	}
	if decoded.Items != nil {
		t.Fatalf("decoded Items = %#v, want nil", decoded.Items)
	}
}

func TestSmartDecodeKeepsEmptySliceNonNil(t *testing.T) {
	var decoded struct {
		Items []string `json:"items"`
	}
	if err := SmartDecode([]byte(`{"items":[]}`), &decoded); err != nil {
		t.Fatalf("SmartDecode() error = %v", err)
	}
	if decoded.Items == nil || len(decoded.Items) != 0 {
		t.Fatalf("decoded Items = %#v, want non-nil empty slice", decoded.Items)
	}
}
