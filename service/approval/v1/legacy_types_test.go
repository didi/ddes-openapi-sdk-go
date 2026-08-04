package v1

import (
	"encoding/json"
	"testing"
)

func TestApprovalOrderRecordLegacyTypesDecodeNumberAndString(t *testing.T) {
	for _, body := range []string{
		`{"car_level":100,"order_status":2,"pay_type":0,"is_invoice":1}`,
		`{"car_level":"100","order_status":"2","pay_type":"0","is_invoice":"1"}`,
	} {
		var record ApprovalOrderRecord
		if err := json.Unmarshal([]byte(body), &record); err != nil {
			t.Fatalf("json.Unmarshal(%s): %v", body, err)
		}
		if record.CarLevel == nil || *record.CarLevel != 100 {
			t.Errorf("CarLevel = %v, want 100", record.CarLevel)
		}
		if record.OrderStatus == nil || *record.OrderStatus != 2 {
			t.Errorf("OrderStatus = %v, want 2", record.OrderStatus)
		}
		if record.PayType == nil || *record.PayType != 0 {
			t.Errorf("PayType = %v, want 0", record.PayType)
		}
		if record.IsInvoice == nil || *record.IsInvoice != 1 {
			t.Errorf("IsInvoice = %v, want 1", record.IsInvoice)
		}
	}
}
