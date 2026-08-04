package v1

import (
	"encoding/json"
	"testing"
)

func TestBillListItemOfWangYCLegacyTypesDecodeNumberAndString(t *testing.T) {
	for _, body := range []string{
		`{"member_id":1125922289295589,"personal_real_pay":20.5}`,
		`{"member_id":"1125922289295589","personal_real_pay":"20.5"}`,
	} {
		var item BillListItemOfWangYC
		if err := json.Unmarshal([]byte(body), &item); err != nil {
			t.Fatalf("json.Unmarshal(%s): %v", body, err)
		}
		if item.MemberId == nil || *item.MemberId != 1125922289295589 {
			t.Errorf("MemberId = %v, want 1125922289295589", item.MemberId)
		}
		if item.PersonalRealPay == nil || *item.PersonalRealPay != float32(20.5) {
			t.Errorf("PersonalRealPay = %v, want 20.5", item.PersonalRealPay)
		}
	}
}

func TestNotGenBDOfWangYCItemLegacyIsSensitiveDecodeNumberAndString(t *testing.T) {
	for _, body := range []string{
		`{"is_sensitive":1}`,
		`{"is_sensitive":"1"}`,
	} {
		var item NotGenBDOfWangYCItem
		if err := json.Unmarshal([]byte(body), &item); err != nil {
			t.Fatalf("json.Unmarshal(%s): %v", body, err)
		}
		if item.IsSensitive == nil || *item.IsSensitive != 1 {
			t.Errorf("IsSensitive = %v, want 1", item.IsSensitive)
		}
	}
}
