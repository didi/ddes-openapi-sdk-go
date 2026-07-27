package v1

import (
	"encoding/json"
	"testing"
)

func TestBudgetCenterRecord_NewFields_JSONUnmarshal(t *testing.T) {
	jsonStr := `{
		"id": "123",
		"name": "技术部",
		"type": "1",
		"status": "1",
		"out_parent_id": "EXT001",
		"out_legal_entity_id": "LE001",
		"department_id": "D001,D002",
		"out_department_id": "OD001",
		"scope": "include_sub",
		"limit_rule_list": [{"rule_name":"月限额","budget_cycle":1,"is_accumulative":0,"total_quota":10000.00,"limit_management_scope":0,"available_quota":"5000.00","freeze_quota":"200.00"}],
		"extend_field": [{"id":1,"code":"custom","value":"test"}],
		"poi_list": [{"city":"北京","city_id":1,"city_adcode":"110000","flat":39.9,"flng":116.4,"poi_range":500,"label":"国贸"}]
	}`

	var record BudgetCenterRecord
	if err := json.Unmarshal([]byte(jsonStr), &record); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	// status
	if record.Status == nil || *record.Status != "1" {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	// out_parent_id
	if record.OutParentId == nil || *record.OutParentId != "EXT001" {
		t.Errorf("OutParentId = %v, want EXT001", record.OutParentId)
	}
	// out_legal_entity_id
	if record.OutLegalEntityId == nil || *record.OutLegalEntityId != "LE001" {
		t.Errorf("OutLegalEntityId = %v, want LE001", record.OutLegalEntityId)
	}
	// department_id
	if record.DepartmentId == nil || *record.DepartmentId != "D001,D002" {
		t.Errorf("DepartmentId = %v, want D001,D002", record.DepartmentId)
	}
	// out_department_id
	if record.OutDepartmentId == nil || *record.OutDepartmentId != "OD001" {
		t.Errorf("OutDepartmentId = %v, want OD001", record.OutDepartmentId)
	}
	// scope
	if record.Scope == nil || *record.Scope != "include_sub" {
		t.Errorf("Scope = %v, want include_sub", record.Scope)
	}
	// limit_rule_list
	if len(record.LimitRuleList) != 1 {
		t.Fatalf("LimitRuleList len = %v, want 1", len(record.LimitRuleList))
	}
	if record.LimitRuleList[0].RuleName == nil || *record.LimitRuleList[0].RuleName != "月限额" {
		t.Errorf("LimitRuleList[0].RuleName = %v, want 月限额", record.LimitRuleList[0].RuleName)
	}
	// extend_field
	if len(record.ExtendField) != 1 {
		t.Fatalf("ExtendField len = %v, want 1", len(record.ExtendField))
	}
	if record.ExtendField[0].Code == nil || *record.ExtendField[0].Code != "custom" {
		t.Errorf("ExtendField[0].Code = %v, want custom", record.ExtendField[0].Code)
	}
	// poi_list
	if len(record.PoiList) != 1 {
		t.Fatalf("PoiList len = %v, want 1", len(record.PoiList))
	}
	if record.PoiList[0].City == nil || *record.PoiList[0].City != "北京" {
		t.Errorf("PoiList[0].City = %v, want 北京", record.PoiList[0].City)
	}
}

func TestBudgetCenterRecord_NewFields_OmitEmpty(t *testing.T) {
	jsonStr := `{"id": "123"}`
	var record BudgetCenterRecord
	if err := json.Unmarshal([]byte(jsonStr), &record); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	// 新增字段在 JSON 中不存在时应为零值
	if record.Status != nil {
		t.Errorf("Status = %v, want nil", record.Status)
	}
	if record.OutParentId != nil {
		t.Errorf("OutParentId = %v, want nil", record.OutParentId)
	}
	if record.LimitRuleList != nil {
		t.Errorf("LimitRuleList = %v, want nil", record.LimitRuleList)
	}
	if record.ExtendField != nil {
		t.Errorf("ExtendField = %v, want nil", record.ExtendField)
	}
	if record.PoiList != nil {
		t.Errorf("PoiList = %v, want nil", record.PoiList)
	}
	if record.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", record.OutLegalEntityId)
	}
	if record.DepartmentId != nil {
		t.Errorf("DepartmentId = %v, want nil", record.DepartmentId)
	}
	if record.OutDepartmentId != nil {
		t.Errorf("OutDepartmentId = %v, want nil", record.OutDepartmentId)
	}
	if record.Scope != nil {
		t.Errorf("Scope = %v, want nil", record.Scope)
	}
}

func TestBudgetCenterRecordBuilder_NewFields(t *testing.T) {
	record := NewBudgetCenterRecordBuilder().
		Id("123").
		Status("1").
		OutParentId("EXT001").
		OutLegalEntityId("LE001").
		DepartmentId("D001,D002").
		OutDepartmentId("OD001").
		Scope("include_sub").
		Build()

	if record.Id == nil || *record.Id != "123" {
		t.Errorf("Id = %v, want 123", record.Id)
	}
	if record.Status == nil || *record.Status != "1" {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	if record.OutParentId == nil || *record.OutParentId != "EXT001" {
		t.Errorf("OutParentId = %v, want EXT001", record.OutParentId)
	}
	if record.OutLegalEntityId == nil || *record.OutLegalEntityId != "LE001" {
		t.Errorf("OutLegalEntityId = %v, want LE001", record.OutLegalEntityId)
	}
	if record.DepartmentId == nil || *record.DepartmentId != "D001,D002" {
		t.Errorf("DepartmentId = %v, want D001,D002", record.DepartmentId)
	}
	if record.OutDepartmentId == nil || *record.OutDepartmentId != "OD001" {
		t.Errorf("OutDepartmentId = %v, want OD001", record.OutDepartmentId)
	}
	if record.Scope == nil || *record.Scope != "include_sub" {
		t.Errorf("Scope = %v, want include_sub", record.Scope)
	}
}
