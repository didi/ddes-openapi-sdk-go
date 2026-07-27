package v1

import (
	"encoding/json"
	"testing"
)

func TestLimitRuleItem_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		wantRule  string
		wantCycle int32
		wantQuota float64
	}{
		{
			name:      "full fields",
			jsonStr:   `{"rule_name":"月限额","budget_cycle":1,"is_accumulative":0,"total_quota":10000.00,"limit_management_scope":0,"available_quota":"5000.00","freeze_quota":"200.00"}`,
			wantRule:  "月限额",
			wantCycle: 1,
			wantQuota: 10000.00,
		},
		{
			name:      "empty object",
			jsonStr:   `{}`,
			wantRule:  "",
			wantCycle: 0,
			wantQuota: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item LimitRuleItem
			if err := json.Unmarshal([]byte(tt.jsonStr), &item); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if tt.wantRule != "" {
				if item.RuleName == nil || *item.RuleName != tt.wantRule {
					t.Errorf("RuleName = %v, want %v", item.RuleName, tt.wantRule)
				}
			} else {
				if item.RuleName != nil {
					t.Errorf("RuleName = %v, want nil", item.RuleName)
				}
			}
			if tt.wantCycle != 0 {
				if item.BudgetCycle == nil || *item.BudgetCycle != tt.wantCycle {
					t.Errorf("BudgetCycle = %v, want %v", item.BudgetCycle, tt.wantCycle)
				}
			}
			if tt.wantQuota != 0 {
				if item.TotalQuota == nil || *item.TotalQuota != tt.wantQuota {
					t.Errorf("TotalQuota = %v, want %v", item.TotalQuota, tt.wantQuota)
				}
			}
		})
	}
}

func TestLimitRuleItemBuilder(t *testing.T) {
	item := NewLimitRuleItemBuilder().
		RuleName("月限额").
		BudgetCycle(1).
		IsAccumulative(0).
		TotalQuota(10000.00).
		LimitManagementScope(0).
		AvailableQuota("5000.00").
		FreezeQuota("200.00").
		Build()

	if item.RuleName == nil || *item.RuleName != "月限额" {
		t.Errorf("RuleName = %v, want 月限额", item.RuleName)
	}
	if item.BudgetCycle == nil || *item.BudgetCycle != 1 {
		t.Errorf("BudgetCycle = %v, want 1", item.BudgetCycle)
	}
	if item.TotalQuota == nil || *item.TotalQuota != 10000.00 {
		t.Errorf("TotalQuota = %v, want 10000.00", item.TotalQuota)
	}
	if item.IsAccumulative == nil || *item.IsAccumulative != 0 {
		t.Errorf("IsAccumulative = %v, want 0", item.IsAccumulative)
	}
	if item.LimitManagementScope == nil || *item.LimitManagementScope != 0 {
		t.Errorf("LimitManagementScope = %v, want 0", item.LimitManagementScope)
	}
	if item.AvailableQuota == nil || *item.AvailableQuota != "5000.00" {
		t.Errorf("AvailableQuota = %v, want 5000.00", item.AvailableQuota)
	}
	if item.FreezeQuota == nil || *item.FreezeQuota != "200.00" {
		t.Errorf("FreezeQuota = %v, want 200.00", item.FreezeQuota)
	}
}
