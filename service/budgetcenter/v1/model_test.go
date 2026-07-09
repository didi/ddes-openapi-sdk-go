package v1

import (
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
