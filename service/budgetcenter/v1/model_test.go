package v1

import (
	"encoding/json"
	"testing"
)

func TestGetBudgetCenterApiReqBuilder_IsNeedLimitRule(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		IsNeedLimitRule(1).
		Build()

	val := req.apiReq.QueryParams.Get("is_need_limit_rule")
	if val != "1" {
		t.Errorf("is_need_limit_rule = %q, want %q", val, "1")
	}
}

func TestGetBudgetCenterApiReqBuilder_IsGetPoi(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		IsGetPoi(1).
		Build()

	val := req.apiReq.QueryParams.Get("is_get_poi")
	if val != "1" {
		t.Errorf("is_get_poi = %q, want %q", val, "1")
	}
}

func TestGetBudgetCenterApiReqBuilder_IsGetExtendFields(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		IsGetExtendFields(1).
		Build()

	val := req.apiReq.QueryParams.Get("is_get_extend_fields")
	if val != "1" {
		t.Errorf("is_get_extend_fields = %q, want %q", val, "1")
	}
}

func TestGetBudgetCenterApiReqBuilder_ExistingFieldsStillWork(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Offset(0).
		Length(100).
		IsNeedLimitRule(1).
		IsGetPoi(1).
		IsGetExtendFields(1).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("access_token") != "test_token" {
		t.Errorf("access_token mismatch")
	}
	if req.apiReq.QueryParams.Get("company_id") != "test_company" {
		t.Errorf("company_id mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	if req.apiReq.QueryParams.Get("length") != "100" {
		t.Errorf("length mismatch")
	}
	if req.apiReq.QueryParams.Get("is_need_limit_rule") != "1" {
		t.Errorf("is_need_limit_rule mismatch")
	}
	if req.apiReq.QueryParams.Get("is_get_poi") != "1" {
		t.Errorf("is_get_poi mismatch")
	}
	if req.apiReq.QueryParams.Get("is_get_extend_fields") != "1" {
		t.Errorf("is_get_extend_fields mismatch")
	}
}

func TestDelBudgetCenterRequestBuilder_GroupAccountFields(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Type(1).
		Id("1125904357323169").
		BelongEnterpriseName("测试企业").
		TaxpayerNo("91110000MA001").
		OutLegalEntityId("LE001").
		Build()

	// 原有字段
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Type == nil || *req.Type != 1 {
		t.Errorf("Type = %v, want 1", req.Type)
	}
	if req.Id == nil || *req.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", req.Id)
	}
	// 集团账户字段
	if req.BelongEnterpriseName == nil || *req.BelongEnterpriseName != "测试企业" {
		t.Errorf("BelongEnterpriseName = %v, want 测试企业", req.BelongEnterpriseName)
	}
	if req.TaxpayerNo == nil || *req.TaxpayerNo != "91110000MA001" {
		t.Errorf("TaxpayerNo = %v, want 91110000MA001", req.TaxpayerNo)
	}
	if req.OutLegalEntityId == nil || *req.OutLegalEntityId != "LE001" {
		t.Errorf("OutLegalEntityId = %v, want LE001", req.OutLegalEntityId)
	}
}

func TestDelBudgetCenterRequestBuilder_GroupAccountFields_OmitEmpty(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		Build()

	// 未设置的字段应为 nil
	if req.BelongEnterpriseName != nil {
		t.Errorf("BelongEnterpriseName = %v, want nil", req.BelongEnterpriseName)
	}
	if req.TaxpayerNo != nil {
		t.Errorf("TaxpayerNo = %v, want nil", req.TaxpayerNo)
	}
	if req.OutLegalEntityId != nil {
		t.Errorf("OutLegalEntityId = %v, want nil", req.OutLegalEntityId)
	}
}

func TestDelBudgetCenterRequestBuilder_GroupAccountFields_JSON(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		BelongEnterpriseName("测试企业").
		TaxpayerNo("91110000MA001").
		OutLegalEntityId("LE001").
		Build()

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal to map failed: %v", err)
	}

	if v, ok := m["belong_enterprise_name"].(string); !ok || v != "测试企业" {
		t.Errorf("belong_enterprise_name = %v, want 测试企业", m["belong_enterprise_name"])
	}
	if v, ok := m["taxpayer_no"].(string); !ok || v != "91110000MA001" {
		t.Errorf("taxpayer_no = %v, want 91110000MA001", m["taxpayer_no"])
	}
	if v, ok := m["out_legal_entity_id"].(string); !ok || v != "LE001" {
		t.Errorf("out_legal_entity_id = %v, want LE001", m["out_legal_entity_id"])
	}
}

func TestDelBudgetCenterRequestBuilder_GroupAccountFields_OmitEmpty_JSON(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		Build()

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal to map failed: %v", err)
	}

	// 未设置的集团账户字段不应出现在 JSON 中（omitempty）
	if _, ok := m["belong_enterprise_name"]; ok {
		t.Errorf("belong_enterprise_name should be omitted, got %v", m["belong_enterprise_name"])
	}
	if _, ok := m["taxpayer_no"]; ok {
		t.Errorf("taxpayer_no should be omitted, got %v", m["taxpayer_no"])
	}
	if _, ok := m["out_legal_entity_id"]; ok {
		t.Errorf("out_legal_entity_id should be omitted, got %v", m["out_legal_entity_id"])
	}
}
