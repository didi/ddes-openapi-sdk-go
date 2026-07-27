package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didi/ddes-openapi-sdk-go/core"
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

// --- LeaderItem Builder 测试 ---

func TestLeaderItemBuilder(t *testing.T) {
	item := NewLeaderItemBuilder().
		LeaderId("1125922289295589").
		LeaderName("张三").
		LeaderType("major").
		Build()

	if item.LeaderId == nil || *item.LeaderId != "1125922289295589" {
		t.Errorf("LeaderId = %v, want 1125922289295589", item.LeaderId)
	}
	if item.LeaderName == nil || *item.LeaderName != "张三" {
		t.Errorf("LeaderName = %v, want 张三", item.LeaderName)
	}
	if item.LeaderType == nil || *item.LeaderType != "major" {
		t.Errorf("LeaderType = %v, want major", item.LeaderType)
	}

	// 部分设置
	item2 := NewLeaderItemBuilder().
		LeaderId("1001").
		Build()

	if item2.LeaderId == nil || *item2.LeaderId != "1001" {
		t.Errorf("LeaderId = %v, want 1001", item2.LeaderId)
	}
	if item2.LeaderName != nil {
		t.Errorf("LeaderName = %v, want nil", item2.LeaderName)
	}
	if item2.LeaderType != nil {
		t.Errorf("LeaderType = %v, want nil", item2.LeaderType)
	}
}

// --- OutTravelerItem Builder 测试 ---

func TestOutTravelerItemBuilder(t *testing.T) {
	item := NewOutTravelerItemBuilder().
		OutTravelerId("OUT001").
		Id(1125922289295589).
		RelatedEmployees([]RelatedEmployeeItem{
			*NewRelatedEmployeeItemBuilder().
				RelatedEmployeeId("1125922289295589").
				EmployeeNumber("D0001").
				Phone("13800000001").
				Email("test@example.com").
				Build(),
		}).
		Build()

	if item.OutTravelerId == nil || *item.OutTravelerId != "OUT001" {
		t.Errorf("OutTravelerId = %v, want OUT001", item.OutTravelerId)
	}
	if item.Id == nil || *item.Id != 1125922289295589 {
		t.Errorf("Id = %v, want 1125922289295589", item.Id)
	}
	if len(item.RelatedEmployees) != 1 {
		t.Fatalf("RelatedEmployees len = %d, want 1", len(item.RelatedEmployees))
	}
	if item.RelatedEmployees[0].RelatedEmployeeId == nil || *item.RelatedEmployees[0].RelatedEmployeeId != "1125922289295589" {
		t.Errorf("RelatedEmployees[0].RelatedEmployeeId = %v, want 1125922289295589", item.RelatedEmployees[0].RelatedEmployeeId)
	}

	// 部分设置
	item2 := NewOutTravelerItemBuilder().
		OutTravelerId("OUT002").
		Build()

	if item2.OutTravelerId == nil || *item2.OutTravelerId != "OUT002" {
		t.Errorf("OutTravelerId = %v, want OUT002", item2.OutTravelerId)
	}
	if item2.Id != nil {
		t.Errorf("Id = %v, want nil", item2.Id)
	}
	if item2.RelatedEmployees != nil {
		t.Errorf("RelatedEmployees = %v, want nil", item2.RelatedEmployees)
	}
}

// --- RelatedEmployeeItem Builder 测试 ---

func TestRelatedEmployeeItemBuilder(t *testing.T) {
	item := NewRelatedEmployeeItemBuilder().
		RelatedEmployeeId("1125922289295589").
		EmployeeNumber("D0001").
		Phone("13800000001").
		Email("test@example.com").
		Build()

	if item.RelatedEmployeeId == nil || *item.RelatedEmployeeId != "1125922289295589" {
		t.Errorf("RelatedEmployeeId = %v, want 1125922289295589", item.RelatedEmployeeId)
	}
	if item.EmployeeNumber == nil || *item.EmployeeNumber != "D0001" {
		t.Errorf("EmployeeNumber = %v, want D0001", item.EmployeeNumber)
	}
	if item.Phone == nil || *item.Phone != "13800000001" {
		t.Errorf("Phone = %v, want 13800000001", item.Phone)
	}
	if item.Email == nil || *item.Email != "test@example.com" {
		t.Errorf("Email = %v, want test@example.com", item.Email)
	}

	// 部分设置
	item2 := NewRelatedEmployeeItemBuilder().
		Phone("13800000002").
		Build()

	if item2.Phone == nil || *item2.Phone != "13800000002" {
		t.Errorf("Phone = %v, want 13800000002", item2.Phone)
	}
	if item2.RelatedEmployeeId != nil {
		t.Errorf("RelatedEmployeeId = %v, want nil", item2.RelatedEmployeeId)
	}
	if item2.EmployeeNumber != nil {
		t.Errorf("EmployeeNumber = %v, want nil", item2.EmployeeNumber)
	}
	if item2.Email != nil {
		t.Errorf("Email = %v, want nil", item2.Email)
	}
}

// --- CreateBudgetCenterRequestBuilder 全量参数测试 ---

func TestCreateBudgetCenterRequestBuilder_AllFields(t *testing.T) {
	extraInfoObj := map[string]string{"key1": "value1"}
	req := NewCreateBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Type(2).
		Name("测试项目").
		BudgetCycle(1).
		TotalQuota("10000.00").
		OutBudgetId("OUT001").
		LeaderId("1125922289295589").
		LeaderEmployeeId("[111,2222]").
		ParentId("1125904357323169").
		OutParentId("OUT_PARENT_001").
		OutParentName("上级项目").
		MemberUsed(1).
		StartDate("2026-01-01").
		ExpiryDate("2026-12-31").
		LegalEntityId("LE001,LE002").
		BudgetExtraInfo(`{"field":"value"}`).
		BudgetExtraInfoObj(extraInfoObj).
		DepartmentId("D001,D002").
		OutDepartmentId("OD001").
		Scope("include_sub").
		ExtendField(`[{"id":1}]`).
		ExtendFieldObj([]ExtendFieldItem{*NewExtendFieldItemBuilder().Id(1).Code("custom").Value("v").Build()}).
		PoiList(`[{"city":"北京"}]`).
		PoiListObj([]PoiItem{*NewPoiItemBuilder().City("北京").CityId(1).Build()}).
		OutTravelers(`[{"out_traveler_id":"OUT001"}]`).
		OutTravelersObj([]OutTravelerItem{*NewOutTravelerItemBuilder().OutTravelerId("OUT001").Build()}).
		BelongEnterpriseName("测试企业").
		TaxpayerNo("91110000MA001").
		OutLegalEntityId("LE001").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.Type == nil || *req.Type != 2 {
		t.Errorf("Type = %v, want 2", req.Type)
	}
	if req.Name == nil || *req.Name != "测试项目" {
		t.Errorf("Name = %v, want 测试项目", req.Name)
	}
	if req.BudgetCycle == nil || *req.BudgetCycle != 1 {
		t.Errorf("BudgetCycle = %v, want 1", req.BudgetCycle)
	}
	if req.TotalQuota == nil || *req.TotalQuota != "10000.00" {
		t.Errorf("TotalQuota = %v, want 10000.00", req.TotalQuota)
	}
	if req.OutBudgetId == nil || *req.OutBudgetId != "OUT001" {
		t.Errorf("OutBudgetId = %v, want OUT001", req.OutBudgetId)
	}
	if req.LeaderId == nil || *req.LeaderId != "1125922289295589" {
		t.Errorf("LeaderId = %v, want 1125922289295589", req.LeaderId)
	}
	if req.LeaderEmployeeId == nil || *req.LeaderEmployeeId != "[111,2222]" {
		t.Errorf("LeaderEmployeeId = %v, want [111,2222]", req.LeaderEmployeeId)
	}
	if req.ParentId == nil || *req.ParentId != "1125904357323169" {
		t.Errorf("ParentId = %v, want 1125904357323169", req.ParentId)
	}
	if req.OutParentId == nil || *req.OutParentId != "OUT_PARENT_001" {
		t.Errorf("OutParentId = %v, want OUT_PARENT_001", req.OutParentId)
	}
	if req.OutParentName == nil || *req.OutParentName != "上级项目" {
		t.Errorf("OutParentName = %v, want 上级项目", req.OutParentName)
	}
	if req.MemberUsed == nil || *req.MemberUsed != 1 {
		t.Errorf("MemberUsed = %v, want 1", req.MemberUsed)
	}
	if req.StartDate == nil || *req.StartDate != "2026-01-01" {
		t.Errorf("StartDate = %v, want 2026-01-01", req.StartDate)
	}
	if req.ExpiryDate == nil || *req.ExpiryDate != "2026-12-31" {
		t.Errorf("ExpiryDate = %v, want 2026-12-31", req.ExpiryDate)
	}
	if req.LegalEntityId == nil || *req.LegalEntityId != "LE001,LE002" {
		t.Errorf("LegalEntityId = %v, want LE001,LE002", req.LegalEntityId)
	}
	if req.BudgetExtraInfo == nil || *req.BudgetExtraInfo != `{"field":"value"}` {
		t.Errorf("BudgetExtraInfo = %v, want {{\"field\":\"value\"}}", req.BudgetExtraInfo)
	}
	if req.BudgetExtraInfoObj == nil || (*req.BudgetExtraInfoObj)["key1"] != "value1" {
		t.Errorf("BudgetExtraInfoObj = %v, want key1=value1", req.BudgetExtraInfoObj)
	}
	if req.DepartmentId == nil || *req.DepartmentId != "D001,D002" {
		t.Errorf("DepartmentId = %v, want D001,D002", req.DepartmentId)
	}
	if req.OutDepartmentId == nil || *req.OutDepartmentId != "OD001" {
		t.Errorf("OutDepartmentId = %v, want OD001", req.OutDepartmentId)
	}
	if req.Scope == nil || *req.Scope != "include_sub" {
		t.Errorf("Scope = %v, want include_sub", req.Scope)
	}
	if req.ExtendField == nil || *req.ExtendField != `[{"id":1}]` {
		t.Errorf("ExtendField = %v, want [{\"id\":1}]", req.ExtendField)
	}
	if len(req.ExtendFieldObj) != 1 {
		t.Errorf("ExtendFieldObj len = %d, want 1", len(req.ExtendFieldObj))
	}
	if req.PoiList == nil || *req.PoiList != `[{"city":"北京"}]` {
		t.Errorf("PoiList = %v, want [{\"city\":\"北京\"}]", req.PoiList)
	}
	if len(req.PoiListObj) != 1 {
		t.Errorf("PoiListObj len = %d, want 1", len(req.PoiListObj))
	}
	if req.OutTravelers == nil || *req.OutTravelers != `[{"out_traveler_id":"OUT001"}]` {
		t.Errorf("OutTravelers = %v, want [{\"out_traveler_id\":\"OUT001\"}]", req.OutTravelers)
	}
	if len(req.OutTravelersObj) != 1 {
		t.Errorf("OutTravelersObj len = %d, want 1", len(req.OutTravelersObj))
	}
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

func TestCreateBudgetCenterRequestBuilder_PartialFields(t *testing.T) {
	req := NewCreateBudgetCenterRequestBuilder().
		ClientId("test_client").
		Type(1).
		Name("技术部").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Type == nil || *req.Type != 1 {
		t.Errorf("Type = %v, want 1", req.Type)
	}
	if req.Name == nil || *req.Name != "技术部" {
		t.Errorf("Name = %v, want 技术部", req.Name)
	}
	// 未设置字段应为 nil
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.BudgetCycle != nil {
		t.Errorf("BudgetCycle = %v, want nil", req.BudgetCycle)
	}
	if req.TotalQuota != nil {
		t.Errorf("TotalQuota = %v, want nil", req.TotalQuota)
	}
	if req.BelongEnterpriseName != nil {
		t.Errorf("BelongEnterpriseName = %v, want nil", req.BelongEnterpriseName)
	}
}

func TestCreateBudgetCenterRequestBuilder_ZeroInt32(t *testing.T) {
	req := NewCreateBudgetCenterRequestBuilder().
		ClientId("test_client").
		Type(0).
		BudgetCycle(0).
		MemberUsed(0).
		Timestamp(0).
		Build()

	// int32 零值也通过 Builder 设置（有 Set 标记）
	if req.Type == nil || *req.Type != 0 {
		t.Errorf("Type = %v, want 0", req.Type)
	}
	if req.BudgetCycle == nil || *req.BudgetCycle != 0 {
		t.Errorf("BudgetCycle = %v, want 0", req.BudgetCycle)
	}
	if req.MemberUsed == nil || *req.MemberUsed != 0 {
		t.Errorf("MemberUsed = %v, want 0", req.MemberUsed)
	}
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestCreateBudgetCenterRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewCreateBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	// 业务参数不应存在
	if req.Type != nil {
		t.Errorf("Type = %v, want nil", req.Type)
	}
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
	if req.OutBudgetId != nil {
		t.Errorf("OutBudgetId = %v, want nil", req.OutBudgetId)
	}
}

// --- CreateBudgetCenterApiReply 反序列化测试 ---

func TestCreateBudgetCenterApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_001"}`

	var reply CreateBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Id != "1125904357323169" {
		t.Errorf("Data.Id = %v, want 1125904357323169", reply.Data.Id)
	}
}

func TestCreateBudgetCenterApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply CreateBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestCreateBudgetCenterApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","request_id":"req_nodata"}`

	var reply CreateBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- GetBudgetCenterApiReqBuilder 测试 ---

func TestGetBudgetCenterApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Id("1125904357323169").
		OutBudgetId("OUT001").
		Type(1).
		IsExactName(1).
		Name("技术部").
		Offset(0).
		Length(10).
		IsNeedLimitRule(1).
		IsGetPoi(1).
		IsGetExtendFields(1).
		Build()

	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"id", "1125904357323169"},
		{"out_budget_id", "OUT001"},
		{"type", "1"},
		{"is_exact_name", "1"},
		{"name", "技术部"},
		{"offset", "0"},
		{"length", "10"},
		{"is_need_limit_rule", "1"},
		{"is_get_poi", "1"},
		{"is_get_extend_fields", "1"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetBudgetCenterApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Id("1125904357323169").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("id") != "1125904357323169" {
		t.Errorf("id mismatch")
	}
	if req.apiReq.QueryParams.Get("out_budget_id") != "" {
		t.Errorf("out_budget_id should be empty")
	}
	if req.apiReq.QueryParams.Get("name") != "" {
		t.Errorf("name should be empty")
	}
	if req.apiReq.QueryParams.Get("type") != "" {
		t.Errorf("type should be empty")
	}
}

func TestGetBudgetCenterApiReqBuilder_ZeroInt32(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Type(0).
		IsExactName(0).
		Offset(0).
		Length(0).
		IsNeedLimitRule(0).
		IsGetPoi(0).
		IsGetExtendFields(0).
		Build()

	if req.apiReq.QueryParams.Get("type") != "0" {
		t.Errorf("type = %q, want \"0\"", req.apiReq.QueryParams.Get("type"))
	}
	if req.apiReq.QueryParams.Get("is_exact_name") != "0" {
		t.Errorf("is_exact_name = %q, want \"0\"", req.apiReq.QueryParams.Get("is_exact_name"))
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
	if req.apiReq.QueryParams.Get("length") != "0" {
		t.Errorf("length = %q, want \"0\"", req.apiReq.QueryParams.Get("length"))
	}
	if req.apiReq.QueryParams.Get("is_need_limit_rule") != "0" {
		t.Errorf("is_need_limit_rule = %q, want \"0\"", req.apiReq.QueryParams.Get("is_need_limit_rule"))
	}
}

func TestGetBudgetCenterApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	for _, key := range []string{"id", "out_budget_id", "type", "is_exact_name", "name", "offset", "length", "is_need_limit_rule", "is_get_poi", "is_get_extend_fields"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- GetBudgetCenterApiReply 反序列化测试 ---

func TestGetBudgetCenterApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": "2",
			"records": [
				{
					"id": "1125904357323169",
					"name": "技术部",
					"type": "1",
					"budget_cycle": 1,
					"out_budget_id": "OUT001",
					"total_quota": "10000.00",
					"is_limit_quota": 1,
					"member_num": 50,
					"available_quota": "5000.00",
					"freeze_quota": "200.00",
					"leader_id": "1125922289295589",
					"parent_id": "1125904357323160",
					"member_used": 0,
					"status": "1",
					"limit_rule_list": [{"rule_name":"月限额","budget_cycle":1,"total_quota":10000.00,"available_quota":"5000.00","freeze_quota":"200.00"}],
					"extend_field": [{"id":1,"code":"custom","value":"test"}],
					"poi_list": [{"city":"北京","city_id":1,"city_adcode":"110100","flat":39.9,"flng":116.4,"poi_range":500,"label":"国贸"}],
					"leader_item_list": [{"leader_id":"1125922289295589","leader_name":"张三","leader_type":"major"}]
				}
			]
		},
		"request_id": "req_001"
	}`

	var reply GetBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != "2" {
		t.Errorf("Total = %v, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
	}
	record := reply.Data.Records[0]
	if record.Id == nil || *record.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", record.Id)
	}
	if record.Name == nil || *record.Name != "技术部" {
		t.Errorf("Name = %v, want 技术部", record.Name)
	}
	if record.Status == nil || *record.Status != "1" {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	if len(record.LimitRuleList) != 1 {
		t.Errorf("LimitRuleList len = %d, want 1", len(record.LimitRuleList))
	}
	if len(record.ExtendField) != 1 {
		t.Errorf("ExtendField len = %d, want 1", len(record.ExtendField))
	}
	if len(record.PoiList) != 1 {
		t.Errorf("PoiList len = %d, want 1", len(record.PoiList))
	}
	if len(record.LeaderItemList) != 1 {
		t.Errorf("LeaderItemList len = %d, want 1", len(record.LeaderItemList))
	}
}

func TestGetBudgetCenterApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply GetBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestGetBudgetCenterApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": "3",
			"records": [
				{"id": "1001", "name": "部门A", "status": "1"},
				{"id": "1002", "name": "部门B"},
				{"id": "1003"}
			]
		},
		"request_id": "req_multi"
	}`

	var reply GetBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 3 {
		t.Fatalf("Records len = %d, want 3", len(reply.Data.Records))
	}
	if reply.Data.Records[0].Status == nil || *reply.Data.Records[0].Status != "1" {
		t.Errorf("Records[0].Status = %v, want 1", reply.Data.Records[0].Status)
	}
	if reply.Data.Records[1].Status != nil {
		t.Errorf("Records[1].Status = %v, want nil", reply.Data.Records[1].Status)
	}
	if reply.Data.Records[2].Name != nil {
		t.Errorf("Records[2].Name = %v, want nil", reply.Data.Records[2].Name)
	}
}

func TestGetBudgetCenterApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":"0","records":[]},"request_id":"req_empty"}`

	var reply GetBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestGetBudgetCenterApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- DelBudgetCenterRequestBuilder 测试 ---

func TestDelBudgetCenterRequestBuilder_AllFields(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Type(1).
		Id("1125904357323169").
		Name("技术部").
		OutBudgetId("OUT001").
		BelongEnterpriseName("测试企业").
		TaxpayerNo("91110000MA001").
		OutLegalEntityId("LE001").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.Type == nil || *req.Type != 1 {
		t.Errorf("Type = %v, want 1", req.Type)
	}
	if req.Id == nil || *req.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", req.Id)
	}
	if req.Name == nil || *req.Name != "技术部" {
		t.Errorf("Name = %v, want 技术部", req.Name)
	}
	if req.OutBudgetId == nil || *req.OutBudgetId != "OUT001" {
		t.Errorf("OutBudgetId = %v, want OUT001", req.OutBudgetId)
	}
}

func TestDelBudgetCenterRequestBuilder_PartialFields(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		Id("1125904357323169").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Id == nil || *req.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", req.Id)
	}
	if req.Type != nil {
		t.Errorf("Type = %v, want nil", req.Type)
	}
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
	if req.OutBudgetId != nil {
		t.Errorf("OutBudgetId = %v, want nil", req.OutBudgetId)
	}
}

func TestDelBudgetCenterRequestBuilder_ZeroInt32(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		Type(0).
		Timestamp(0).
		Build()

	if req.Type == nil || *req.Type != 0 {
		t.Errorf("Type = %v, want 0", req.Type)
	}
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestDelBudgetCenterRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewDelBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Type != nil {
		t.Errorf("Type = %v, want nil", req.Type)
	}
	if req.Id != nil {
		t.Errorf("Id = %v, want nil", req.Id)
	}
	if req.OutBudgetId != nil {
		t.Errorf("OutBudgetId = %v, want nil", req.OutBudgetId)
	}
}

// --- DelBudgetCenterApiReply 反序列化测试 ---

func TestDelBudgetCenterApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`

	var reply DelBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
}

func TestDelBudgetCenterApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply DelBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

// --- UpdateBudgetCenterRequestBuilder 测试 ---

func TestUpdateBudgetCenterRequestBuilder_AllFields(t *testing.T) {
	extraInfoObj := map[string]string{"key1": "value1"}
	req := NewUpdateBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Type(2).
		BudgetCycle(1).
		TotalQuota("20000.00").
		Id("1125904357323169").
		Name("更新项目").
		OutBudgetId("OUT001").
		LeaderId("1125922289295589").
		LeaderEmployeeId("[111,2222]").
		ParentId("1125904357323160").
		OutParentId("OUT_PARENT_001").
		OutParentName("上级项目").
		MemberUsed(1).
		StartDate("2026-01-01").
		ExpiryDate("2026-12-31").
		LegalEntityId("LE001").
		BudgetExtraInfo(`{"field":"value"}`).
		BudgetExtraInfoObj(extraInfoObj).
		DepartmentId("D001").
		OutDepartmentId("OD001").
		Scope("current_only").
		ExtendField(`[{"id":1}]`).
		ExtendFieldObj([]ExtendFieldItem{*NewExtendFieldItemBuilder().Id(1).Code("c").Value("v").Build()}).
		PoiList(`[{"city":"上海"}]`).
		PoiListObj([]PoiItem{*NewPoiItemBuilder().City("上海").CityId(2).Build()}).
		OutTravelers(`[{"out_traveler_id":"OUT002"}]`).
		OutTravelersObj([]OutTravelerItem{*NewOutTravelerItemBuilder().OutTravelerId("OUT002").Build()}).
		OperateType("cover").
		BelongEnterpriseName("测试企业").
		TaxpayerNo("91110000MA001").
		OutLegalEntityId("LE001").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Type == nil || *req.Type != 2 {
		t.Errorf("Type = %v, want 2", req.Type)
	}
	if req.Id == nil || *req.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", req.Id)
	}
	if req.Name == nil || *req.Name != "更新项目" {
		t.Errorf("Name = %v, want 更新项目", req.Name)
	}
	if req.OperateType == nil || *req.OperateType != "cover" {
		t.Errorf("OperateType = %v, want cover", req.OperateType)
	}
	if req.BelongEnterpriseName == nil || *req.BelongEnterpriseName != "测试企业" {
		t.Errorf("BelongEnterpriseName = %v, want 测试企业", req.BelongEnterpriseName)
	}
}

func TestUpdateBudgetCenterRequestBuilder_PartialFields(t *testing.T) {
	req := NewUpdateBudgetCenterRequestBuilder().
		ClientId("test_client").
		Type(1).
		Id("1125904357323169").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Type == nil || *req.Type != 1 {
		t.Errorf("Type = %v, want 1", req.Type)
	}
	if req.Id == nil || *req.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", req.Id)
	}
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
	if req.OperateType != nil {
		t.Errorf("OperateType = %v, want nil", req.OperateType)
	}
}

func TestUpdateBudgetCenterRequestBuilder_ZeroInt32(t *testing.T) {
	req := NewUpdateBudgetCenterRequestBuilder().
		ClientId("test_client").
		Type(0).
		BudgetCycle(0).
		MemberUsed(0).
		Timestamp(0).
		Build()

	if req.Type == nil || *req.Type != 0 {
		t.Errorf("Type = %v, want 0", req.Type)
	}
	if req.BudgetCycle == nil || *req.BudgetCycle != 0 {
		t.Errorf("BudgetCycle = %v, want 0", req.BudgetCycle)
	}
	if req.MemberUsed == nil || *req.MemberUsed != 0 {
		t.Errorf("MemberUsed = %v, want 0", req.MemberUsed)
	}
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestUpdateBudgetCenterRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewUpdateBudgetCenterRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Type != nil {
		t.Errorf("Type = %v, want nil", req.Type)
	}
	if req.Id != nil {
		t.Errorf("Id = %v, want nil", req.Id)
	}
	if req.Name != nil {
		t.Errorf("Name = %v, want nil", req.Name)
	}
}

// --- UpdateBudgetCenterApiReply 反序列化测试 ---

func TestUpdateBudgetCenterApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`

	var reply UpdateBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
}

func TestUpdateBudgetCenterApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply UpdateBudgetCenterApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

// --- BudgetCenterRecord Builder 全量参数测试 ---

func TestBudgetCenterRecordBuilder_AllFields(t *testing.T) {
	record := NewBudgetCenterRecordBuilder().
		Id("1125904357323169").
		Name("技术部").
		Type("1").
		BudgetCycle(1).
		OutBudgetId("OUT001").
		TotalQuota("10000.00").
		IsLimitQuota(1).
		MemberNum(50).
		AvailableQuota("5000.00").
		FreezeQuota("200.00").
		LeaderId("1125922289295589").
		LeaderItemList([]LeaderItem{*NewLeaderItemBuilder().LeaderId("1125922289295589").LeaderName("张三").LeaderType("major").Build()}).
		ParentId("1125904357323160").
		MemberUsed(0).
		StartDate("2026-01-01").
		ExpiryDate("2026-12-31").
		LegalEntityId("LE001").
		BudgetExtraInfo(`{"field":"value"}`).
		Status("1").
		OutParentId("OUT_PARENT_001").
		LimitRuleList([]LimitRuleItem{*NewLimitRuleItemBuilder().RuleName("月限额").BudgetCycle(1).Build()}).
		ExtendField([]ExtendFieldItem{*NewExtendFieldItemBuilder().Id(1).Code("custom").Value("v").Build()}).
		PoiList([]PoiItem{*NewPoiItemBuilder().City("北京").CityId(1).Build()}).
		OutLegalEntityId("LE001").
		DepartmentId("D001").
		OutDepartmentId("OD001").
		Scope("include_sub").
		Build()

	if record.Id == nil || *record.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", record.Id)
	}
	if record.Name == nil || *record.Name != "技术部" {
		t.Errorf("Name = %v, want 技术部", record.Name)
	}
	if record.Type == nil || *record.Type != "1" {
		t.Errorf("Type = %v, want 1", record.Type)
	}
	if record.BudgetCycle == nil || *record.BudgetCycle != 1 {
		t.Errorf("BudgetCycle = %v, want 1", record.BudgetCycle)
	}
	if record.IsLimitQuota == nil || *record.IsLimitQuota != 1 {
		t.Errorf("IsLimitQuota = %v, want 1", record.IsLimitQuota)
	}
	if record.MemberNum == nil || *record.MemberNum != 50 {
		t.Errorf("MemberNum = %v, want 50", record.MemberNum)
	}
	if record.MemberUsed == nil || *record.MemberUsed != 0 {
		t.Errorf("MemberUsed = %v, want 0", record.MemberUsed)
	}
	if len(record.LeaderItemList) != 1 {
		t.Errorf("LeaderItemList len = %d, want 1", len(record.LeaderItemList))
	}
}

// ==================== 资源方法测试 ====================

func newBudgetCenterTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// --- CreateBudgetCenter 资源方法测试 ---

func TestCreateBudgetCenter_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/BudgetCenter/add" {
			t.Errorf("expected path /river/BudgetCenter/add, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().
			ClientId("test_client").
			Type(2).
			Name("测试项目").
			Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Errno == nil || *resp.CreateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateBudgetCenterApiReply.Errno)
	}
	if resp.CreateBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data.Id != "1125904357323169" {
		t.Errorf("Data.Id = %v, want 1125904357323169", resp.CreateBudgetCenterApiReply.Data.Id)
	}
}

func TestCreateBudgetCenter_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":""},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data.Id != "" {
		t.Errorf("Data.Id = %v, want empty", resp.CreateBudgetCenterApiReply.Data.Id)
	}
}

func TestCreateBudgetCenter_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Errno == nil || *resp.CreateBudgetCenterApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.CreateBudgetCenterApiReply.Errno)
	}
}

func TestCreateBudgetCenter_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.CreateBudgetCenterApiReply != nil {
		t.Errorf("CreateBudgetCenterApiReply should be nil for non-200 response")
	}
}

func TestCreateBudgetCenter_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Errno == nil || *resp.CreateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateBudgetCenterApiReply.Errno)
	}
	if resp.CreateBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data.Id != "1125904357323169" {
		t.Errorf("Data.Id = %v, want 1125904357323169", resp.CreateBudgetCenterApiReply.Data.Id)
	}
}

func TestCreateBudgetCenter_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data.Id != "1125904357323169" {
		t.Errorf("Data.Id = %v, want 1125904357323169", resp.CreateBudgetCenterApiReply.Data.Id)
	}
}

func TestCreateBudgetCenter_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	bc := &budgetCenter{option: option}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
	if resp.CreateBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestCreateBudgetCenter_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":"1125904357323169"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(NewCreateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.CreateBudgetCenter(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateBudgetCenter() error = %v", err)
	}
	if resp.CreateBudgetCenterApiReply == nil {
		t.Fatal("CreateBudgetCenterApiReply is nil")
	}
}

// --- DelBudgetCenter 资源方法测试 ---

func TestDelBudgetCenter_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/BudgetCenter/del" {
			t.Errorf("expected path /river/BudgetCenter/del, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().
			ClientId("test_client").
			Type(1).
			Id("1125904357323169").
			Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
	if resp.DelBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelBudgetCenterApiReply.Errno)
	}
}

func TestDelBudgetCenter_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
	if resp.DelBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelBudgetCenterApiReply.Errno)
	}
}

func TestDelBudgetCenter_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
	if resp.DelBudgetCenterApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.DelBudgetCenterApiReply.Errno)
	}
}

func TestDelBudgetCenter_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.DelBudgetCenterApiReply != nil {
		t.Errorf("DelBudgetCenterApiReply should be nil for non-200 response")
	}
}

func TestDelBudgetCenter_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
	if resp.DelBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelBudgetCenterApiReply.Errno)
	}
}

func TestDelBudgetCenter_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
	if resp.DelBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelBudgetCenterApiReply.Errno)
	}
}

func TestDelBudgetCenter_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	bc := &budgetCenter{option: option}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
}

func TestDelBudgetCenter_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(NewDelBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.DelBudgetCenter(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("DelBudgetCenter() error = %v", err)
	}
	if resp.DelBudgetCenterApiReply == nil {
		t.Fatal("DelBudgetCenterApiReply is nil")
	}
}

// --- GetBudgetCenter 资源方法测试 ---

func TestGetBudgetCenter_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/BudgetCenter/get" {
			t.Errorf("expected path /river/BudgetCenter/get, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":"1","records":[{"id":"1125904357323169","name":"技术部","type":"1","status":"1","limit_rule_list":[{"rule_name":"月限额","budget_cycle":1,"total_quota":10000.00,"available_quota":"5000.00","freeze_quota":"200.00"}]}]},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Id("1125904357323169").
		Offset(0).
		Length(10).
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if resp.GetBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetBudgetCenterApiReply.Errno)
	}
	if resp.GetBudgetCenterApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetBudgetCenterApiReply.Data.Total != "1" {
		t.Errorf("Total = %v, want 1", resp.GetBudgetCenterApiReply.Data.Total)
	}
	if len(resp.GetBudgetCenterApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetBudgetCenterApiReply.Data.Records))
	}
	record := resp.GetBudgetCenterApiReply.Data.Records[0]
	if record.Id == nil || *record.Id != "1125904357323169" {
		t.Errorf("Id = %v, want 1125904357323169", record.Id)
	}
	if record.Name == nil || *record.Name != "技术部" {
		t.Errorf("Name = %v, want 技术部", record.Name)
	}
	if record.Status == nil || *record.Status != "1" {
		t.Errorf("Status = %v, want 1", record.Status)
	}
}

func TestGetBudgetCenter_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":"0","records":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if len(resp.GetBudgetCenterApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.GetBudgetCenterApiReply.Data.Records))
	}
}

func TestGetBudgetCenter_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if resp.GetBudgetCenterApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.GetBudgetCenterApiReply.Errno)
	}
}

func TestGetBudgetCenter_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.GetBudgetCenterApiReply != nil {
		t.Errorf("GetBudgetCenterApiReply should be nil for non-200 response")
	}
}

func TestGetBudgetCenter_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":"1","records":[{"id":"1125904357323169","name":"技术部","status":"1"}]},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if resp.GetBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetBudgetCenterApiReply.Errno)
	}
	if len(resp.GetBudgetCenterApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetBudgetCenterApiReply.Data.Records))
	}
}

func TestGetBudgetCenter_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":"1","records":[{"id":"1125904357323169","name":"技术部"}]},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if len(resp.GetBudgetCenterApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetBudgetCenterApiReply.Data.Records))
	}
}

func TestGetBudgetCenter_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":"1","records":[{"id":"1125904357323169"}]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	bc := &budgetCenter{option: option}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
	if len(resp.GetBudgetCenterApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.GetBudgetCenterApiReply.Data.Records))
	}
}

func TestGetBudgetCenter_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":"0","records":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewGetBudgetCenterApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := bc.GetBudgetCenter(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetBudgetCenter() error = %v", err)
	}
	if resp.GetBudgetCenterApiReply == nil {
		t.Fatal("GetBudgetCenterApiReply is nil")
	}
}

// --- UpdateBudgetCenter 资源方法测试 ---

func TestUpdateBudgetCenter_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/BudgetCenter/edit" {
			t.Errorf("expected path /river/BudgetCenter/edit, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().
			ClientId("test_client").
			Type(2).
			Id("1125904357323169").
			Name("更新项目").
			Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
	if resp.UpdateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateBudgetCenterApiReply.Errno)
	}
}

func TestUpdateBudgetCenter_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
	if resp.UpdateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateBudgetCenterApiReply.Errno)
	}
}

func TestUpdateBudgetCenter_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
	if resp.UpdateBudgetCenterApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.UpdateBudgetCenterApiReply.Errno)
	}
}

func TestUpdateBudgetCenter_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateBudgetCenterApiReply != nil {
		t.Errorf("UpdateBudgetCenterApiReply should be nil for non-200 response")
	}
}

func TestUpdateBudgetCenter_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
	if resp.UpdateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateBudgetCenterApiReply.Errno)
	}
}

func TestUpdateBudgetCenter_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
	if resp.UpdateBudgetCenterApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateBudgetCenterApiReply.Errno)
	}
}

func TestUpdateBudgetCenter_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	bc := &budgetCenter{option: option}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
}

func TestUpdateBudgetCenter_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBudgetCenterTestOption(testServer.URL)
	bc := &budgetCenter{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(NewUpdateBudgetCenterRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := bc.UpdateBudgetCenter(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateBudgetCenter() error = %v", err)
	}
	if resp.UpdateBudgetCenterApiReply == nil {
		t.Fatal("UpdateBudgetCenterApiReply is nil")
	}
}
