package v1

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// --- ResidentsListInfo Builder 测试 ---

func TestResidentsListInfoBuilder(t *testing.T) {
	info := NewResidentsListInfoBuilder().
		Id("4").
		Name("青岛").
		Adcode("370200").
		Build()

	if info.Id == nil || *info.Id != "4" {
		t.Errorf("Id = %v, want 4", info.Id)
	}
	if info.Name == nil || *info.Name != "青岛" {
		t.Errorf("Name = %v, want 青岛", info.Name)
	}
	if info.Adcode == nil || *info.Adcode != "370200" {
		t.Errorf("Adcode = %v, want 370200", info.Adcode)
	}

	// 部分设置
	info2 := NewResidentsListInfoBuilder().
		Id("5").
		Build()

	if info2.Id == nil || *info2.Id != "5" {
		t.Errorf("Id = %v, want 5", info2.Id)
	}
	if info2.Name != nil {
		t.Errorf("Name = %v, want nil", info2.Name)
	}
	if info2.Adcode != nil {
		t.Errorf("Adcode = %v, want nil", info2.Adcode)
	}
}

// --- LimitRuleInfo Builder 测试 ---

func TestLimitRuleInfoBuilder(t *testing.T) {
	info := NewLimitRuleInfoBuilder().
		RuleName("员工个人限额（默认）").
		BudgetCycle(1).
		IsAccumulative(0).
		TotalQuota(50000).
		AvailableQuota(50000).
		FreezeQuota(0).
		LimitManagementScope(0).
		Build()

	if info.RuleName == nil || *info.RuleName != "员工个人限额（默认）" {
		t.Errorf("RuleName = %v, want 员工个人限额（默认）", info.RuleName)
	}
	if info.BudgetCycle == nil || *info.BudgetCycle != 1 {
		t.Errorf("BudgetCycle = %v, want 1", info.BudgetCycle)
	}
	if info.IsAccumulative == nil || *info.IsAccumulative != 0 {
		t.Errorf("IsAccumulative = %v, want 0", info.IsAccumulative)
	}
	if info.TotalQuota == nil || *info.TotalQuota != 50000 {
		t.Errorf("TotalQuota = %v, want 50000", info.TotalQuota)
	}
	if info.AvailableQuota == nil || *info.AvailableQuota != 50000 {
		t.Errorf("AvailableQuota = %v, want 50000", info.AvailableQuota)
	}
	if info.FreezeQuota == nil || *info.FreezeQuota != 0 {
		t.Errorf("FreezeQuota = %v, want 0", info.FreezeQuota)
	}
	if info.LimitManagementScope == nil || *info.LimitManagementScope != 0 {
		t.Errorf("LimitManagementScope = %v, want 0", info.LimitManagementScope)
	}

	// 部分设置
	info2 := NewLimitRuleInfoBuilder().
		RuleName("test").
		Build()

	if info2.RuleName == nil || *info2.RuleName != "test" {
		t.Errorf("RuleName = %v, want test", info2.RuleName)
	}
	if info2.BudgetCycle != nil {
		t.Errorf("BudgetCycle = %v, want nil", info2.BudgetCycle)
	}
	if info2.TotalQuota != nil {
		t.Errorf("TotalQuota = %v, want nil", info2.TotalQuota)
	}
}

// --- HomeAddressInfo Builder 测试 ---

func TestHomeAddressInfoBuilder(t *testing.T) {
	info := NewHomeAddressInfoBuilder().
		City("北京").
		CityId(1).
		CityAdcode("110100").
		AddressName("朝阳区望京SOHO").
		Build()

	if info.City == nil || *info.City != "北京" {
		t.Errorf("City = %v, want 北京", info.City)
	}
	if info.CityId == nil || *info.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info.CityId)
	}
	if info.CityAdcode == nil || *info.CityAdcode != "110100" {
		t.Errorf("CityAdcode = %v, want 110100", info.CityAdcode)
	}
	if info.AddressName == nil || *info.AddressName != "朝阳区望京SOHO" {
		t.Errorf("AddressName = %v, want 朝阳区望京SOHO", info.AddressName)
	}

	// 部分设置
	info2 := NewHomeAddressInfoBuilder().
		City("上海").
		Build()

	if info2.City == nil || *info2.City != "上海" {
		t.Errorf("City = %v, want 上海", info2.City)
	}
	if info2.CityId != nil {
		t.Errorf("CityId = %v, want nil", info2.CityId)
	}
	if info2.CityAdcode != nil {
		t.Errorf("CityAdcode = %v, want nil", info2.CityAdcode)
	}
}

// --- ListMemberApiReqBuilder 测试 ---

func TestListMemberApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Email("test@example.com").
		EmployeeNumber("D0001").
		Phone("13800000001").
		Realname("张三").
		Status("1,4").
		Offset(0).
		Length(10).
		LastId("0").
		BelongEnterpriseName("子公司A").
		TaxpayerNo("91110000xxx").
		QuerySubCompany("true").
		NextToken("next_token_value").
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
		{"email", "test@example.com"},
		{"employee_number", "D0001"},
		{"phone", "13800000001"},
		{"realname", "张三"},
		{"status", "1,4"},
		{"offset", "0"},
		{"length", "10"},
		{"last_id", "0"},
		{"belong_enterprise_name", "子公司A"},
		{"taxpayer_no", "91110000xxx"},
		{"query_sub_company", "true"},
		{"next_token", "next_token_value"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListMemberApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Offset(0).
		Length(10).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("phone") != "13800000001" {
		t.Errorf("phone mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	if req.apiReq.QueryParams.Get("length") != "10" {
		t.Errorf("length mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("email") != "" {
		t.Errorf("email should be empty, got %q", req.apiReq.QueryParams.Get("email"))
	}
	if req.apiReq.QueryParams.Get("belong_enterprise_name") != "" {
		t.Errorf("belong_enterprise_name should be empty, got %q", req.apiReq.QueryParams.Get("belong_enterprise_name"))
	}
	if req.apiReq.QueryParams.Get("taxpayer_no") != "" {
		t.Errorf("taxpayer_no should be empty, got %q", req.apiReq.QueryParams.Get("taxpayer_no"))
	}
	if req.apiReq.QueryParams.Get("next_token") != "" {
		t.Errorf("next_token should be empty, got %q", req.apiReq.QueryParams.Get("next_token"))
	}
}

func TestListMemberApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		Offset(0).
		Length(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
	if req.apiReq.QueryParams.Get("length") != "0" {
		t.Errorf("length = %q, want \"0\"", req.apiReq.QueryParams.Get("length"))
	}
}

func TestListMemberApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	// 仅通用参数
	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
	// 业务参数不应存在
	for _, key := range []string{"email", "employee_number", "phone", "realname", "status", "offset", "length", "last_id", "belong_enterprise_name", "taxpayer_no", "query_sub_company", "next_token"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- ListMemberApiReply 反序列化测试 ---

func TestListMemberApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"records": [
				{
					"available_quota": "0.00",
					"birth_date": "",
					"budget_center_id": "1125915646135311",
					"card_list": [],
					"con_department_ids": ["1125920823148798", "1125928765383351"],
					"email": "",
					"employee_number": "",
					"english_name": "",
					"english_surname": "",
					"id": "1125922289295589",
					"immediate_superior_phone": "",
					"is_remark": "0",
					"legal_entity_id": "",
					"nickname": "",
					"phone": "00016189857",
					"rank_id": "",
					"realname": "ZHOUZH",
					"regulation_id": ["1125920826148759", "1125920865383421"],
					"residents_list": [
						{"id": "13", "name": "青岛", "adcode": "370200"}
					],
					"role_ids": "1125915646090887_1125915646107623",
					"set_dismiss_time": "",
					"sex": 0,
					"status": 1,
					"system_role": 2,
					"total_quota": "0.00",
					"use_car_config": ["1125920826148759"],
					"use_company_money": 1,
					"limit_rule_list": [
						{
							"rule_name": "员工个人限额（默认）",
							"budget_cycle": 1,
							"is_accumulative": 0,
							"total_quota": 50000,
							"limit_management_scope": 0,
							"available_quota": 50000,
							"freeze_quota": 0
						}
					],
					"cert_realname": "张三",
					"cert_english_surname": "Zhang",
					"cert_english_name": "San",
					"home_address": [
						{"city": "北京", "city_id": 1, "city_adcode": "110100", "address_name": "望京SOHO"}
					]
				}
			],
			"total": 78
		},
		"errmsg": "SUCCESS",
		"errno": 0,
		"request_id": "test_request_id"
	}`

	var reply ListMemberApiReply
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
	if reply.Data.Total == nil || *reply.Data.Total != 78 {
		t.Errorf("Total = %v, want 78", reply.Data.Total)
	}
	if len(reply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
	}

	record := reply.Data.Records[0]
	// 新增字段验证
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	if len(record.ResidentsList) != 1 {
		t.Fatalf("ResidentsList len = %d, want 1", len(record.ResidentsList))
	}
	if record.ResidentsList[0].Id == nil || *record.ResidentsList[0].Id != "13" {
		t.Errorf("ResidentsList[0].Id = %v, want 13", record.ResidentsList[0].Id)
	}
	if record.ResidentsList[0].Name == nil || *record.ResidentsList[0].Name != "青岛" {
		t.Errorf("ResidentsList[0].Name = %v, want 青岛", record.ResidentsList[0].Name)
	}
	if record.ResidentsList[0].Adcode == nil || *record.ResidentsList[0].Adcode != "370200" {
		t.Errorf("ResidentsList[0].Adcode = %v, want 370200", record.ResidentsList[0].Adcode)
	}
	if len(record.LimitRuleList) != 1 {
		t.Fatalf("LimitRuleList len = %d, want 1", len(record.LimitRuleList))
	}
	if record.LimitRuleList[0].RuleName == nil || *record.LimitRuleList[0].RuleName != "员工个人限额（默认）" {
		t.Errorf("LimitRuleList[0].RuleName = %v, want 员工个人限额（默认）", record.LimitRuleList[0].RuleName)
	}
	if record.LimitRuleList[0].BudgetCycle == nil || *record.LimitRuleList[0].BudgetCycle != 1 {
		t.Errorf("LimitRuleList[0].BudgetCycle = %v, want 1", record.LimitRuleList[0].BudgetCycle)
	}
	if record.CertRealname == nil || *record.CertRealname != "张三" {
		t.Errorf("CertRealname = %v, want 张三", record.CertRealname)
	}
	if record.CertEnglishSurname == nil || *record.CertEnglishSurname != "Zhang" {
		t.Errorf("CertEnglishSurname = %v, want Zhang", record.CertEnglishSurname)
	}
	if record.CertEnglishName == nil || *record.CertEnglishName != "San" {
		t.Errorf("CertEnglishName = %v, want San", record.CertEnglishName)
	}
	if len(record.HomeAddress) != 1 {
		t.Fatalf("HomeAddress len = %d, want 1", len(record.HomeAddress))
	}
	if record.HomeAddress[0].City == nil || *record.HomeAddress[0].City != "北京" {
		t.Errorf("HomeAddress[0].City = %v, want 北京", record.HomeAddress[0].City)
	}
	if record.HomeAddress[0].CityId == nil || *record.HomeAddress[0].CityId != 1 {
		t.Errorf("HomeAddress[0].CityId = %v, want 1", record.HomeAddress[0].CityId)
	}
	if record.HomeAddress[0].CityAdcode == nil || *record.HomeAddress[0].CityAdcode != "110100" {
		t.Errorf("HomeAddress[0].CityAdcode = %v, want 110100", record.HomeAddress[0].CityAdcode)
	}
	if record.HomeAddress[0].AddressName == nil || *record.HomeAddress[0].AddressName != "望京SOHO" {
		t.Errorf("HomeAddress[0].AddressName = %v, want 望京SOHO", record.HomeAddress[0].AddressName)
	}
}

func TestListMemberApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply ListMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestListMemberApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"records": [
				{"id": "1001", "phone": "13800000001", "status": 1, "residents_list": [{"id": "13", "name": "青岛", "adcode": "370200"}]},
				{"id": "1002", "phone": "13800000002", "status": 4},
				{"id": "1003", "phone": "13800000003"}
			],
			"total": 3
		},
		"request_id": "req_multi"
	}`

	var reply ListMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 3 {
		t.Fatalf("Records len = %d, want 3", len(reply.Data.Records))
	}
	// 第1条有 residents_list
	if len(reply.Data.Records[0].ResidentsList) != 1 {
		t.Errorf("Records[0].ResidentsList len = %d, want 1", len(reply.Data.Records[0].ResidentsList))
	}
	// 第2条有 status 但无 residents_list
	if reply.Data.Records[1].Status == nil || *reply.Data.Records[1].Status != 4 {
		t.Errorf("Records[1].Status = %v, want 4", reply.Data.Records[1].Status)
	}
	// 第3条缺少 status，应为 nil
	if reply.Data.Records[2].Status != nil {
		t.Errorf("Records[2].Status = %v, want nil", reply.Data.Records[2].Status)
	}
}

func TestListMemberApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_empty"}`

	var reply ListMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestListMemberApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListMemberApiReply
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

// --- MemberRecord Builder 测试（新增字段）---

func TestMemberRecordBuilder_NewFields(t *testing.T) {
	record := NewMemberRecordBuilder().
		Id("1125922289295589").
		Phone("13800000001").
		Status(1).
		CertRealname("张三").
		CertEnglishSurname("Zhang").
		CertEnglishName("San").
		ResidentsList([]ResidentsListInfo{
			*NewResidentsListInfoBuilder().Id("4").Name("青岛").Adcode("370200").Build(),
		}).
		LimitRuleList([]LimitRuleInfo{
			*NewLimitRuleInfoBuilder().RuleName("默认").BudgetCycle(1).Build(),
		}).
		HomeAddress([]HomeAddressInfo{
			*NewHomeAddressInfoBuilder().City("北京").CityId(1).Build(),
		}).
		Build()

	if record.Id == nil || *record.Id != "1125922289295589" {
		t.Errorf("Id = %v, want 1125922289295589", record.Id)
	}
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	if record.CertRealname == nil || *record.CertRealname != "张三" {
		t.Errorf("CertRealname = %v, want 张三", record.CertRealname)
	}
	if record.CertEnglishSurname == nil || *record.CertEnglishSurname != "Zhang" {
		t.Errorf("CertEnglishSurname = %v, want Zhang", record.CertEnglishSurname)
	}
	if record.CertEnglishName == nil || *record.CertEnglishName != "San" {
		t.Errorf("CertEnglishName = %v, want San", record.CertEnglishName)
	}
	if len(record.ResidentsList) != 1 {
		t.Errorf("ResidentsList len = %d, want 1", len(record.ResidentsList))
	}
	if len(record.LimitRuleList) != 1 {
		t.Errorf("LimitRuleList len = %d, want 1", len(record.LimitRuleList))
	}
	if len(record.HomeAddress) != 1 {
		t.Errorf("HomeAddress len = %d, want 1", len(record.HomeAddress))
	}

	// 部分设置：仅设置部分新字段
	record2 := NewMemberRecordBuilder().
		Id("1001").
		Status(4).
		Build()

	if record2.Id == nil || *record2.Id != "1001" {
		t.Errorf("Id = %v, want 1001", record2.Id)
	}
	if record2.Status == nil || *record2.Status != 4 {
		t.Errorf("Status = %v, want 4", record2.Status)
	}
	if record2.CertRealname != nil {
		t.Errorf("CertRealname = %v, want nil", record2.CertRealname)
	}
	if record2.ResidentsList != nil {
		t.Errorf("ResidentsList = %v, want nil", record2.ResidentsList)
	}
	if record2.HomeAddress != nil {
		t.Errorf("HomeAddress = %v, want nil", record2.HomeAddress)
	}
}

// --- ListMember 资源方法测试 ---

func newMemberTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestListMember_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Member/get" {
			t.Errorf("expected path /river/Member/get, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"records": [
					{
						"id": "1125922289295589",
						"phone": "13800000001",
						"realname": "张三",
						"status": 1,
						"residents_list": [{"id": "4", "name": "青岛", "adcode": "370200"}],
						"limit_rule_list": [{"rule_name": "默认", "budget_cycle": 1, "total_quota": 50000, "available_quota": 50000, "freeze_quota": 0}],
						"cert_realname": "张三",
						"cert_english_surname": "Zhang",
						"cert_english_name": "San",
						"home_address": [{"city": "北京", "city_id": 1, "city_adcode": "110100", "address_name": "望京SOHO"}]
					}
				],
				"total": 1
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Offset(0).
		Length(10).
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if resp.ListMemberApiReply.Errno == nil || *resp.ListMemberApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListMemberApiReply.Errno)
	}
	if len(resp.ListMemberApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListMemberApiReply.Data.Records))
	}
	record := resp.ListMemberApiReply.Data.Records[0]
	if record.Status == nil || *record.Status != 1 {
		t.Errorf("Status = %v, want 1", record.Status)
	}
	if record.CertRealname == nil || *record.CertRealname != "张三" {
		t.Errorf("CertRealname = %v, want 张三", record.CertRealname)
	}
}

func TestListMember_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if len(resp.ListMemberApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.ListMemberApiReply.Data.Records))
	}
}

func TestListMember_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if resp.ListMemberApiReply.Errno == nil || *resp.ListMemberApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.ListMemberApiReply.Errno)
	}
}

func TestListMember_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.ListMemberApiReply != nil {
		t.Errorf("ListMemberApiReply should be nil for non-200 response")
	}
}

func TestListMember_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"id":"1125922289295589","phone":"13800000001","status":1}],"total":1},"request_id":"req_enc"}`
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

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Offset(0).
		Length(10).
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if resp.ListMemberApiReply.Errno == nil || *resp.ListMemberApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListMemberApiReply.Errno)
	}
	if len(resp.ListMemberApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListMemberApiReply.Data.Records))
	}
}

func TestListMember_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"id":"1125922289295589","phone":"13800000001"}],"total":1},"request_id":"req_enc256"}`
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

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if len(resp.ListMemberApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListMemberApiReply.Data.Records))
	}
}

func TestListMember_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"id":"1125922289295589"}],"total":1},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	m := &member{option: option}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.ListMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
	if len(resp.ListMemberApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListMemberApiReply.Data.Records))
	}
}

func TestListMember_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"records":[],"total":0},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewListMemberApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.ListMember(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListMember() error = %v", err)
	}
	if resp.ListMemberApiReply == nil {
		t.Fatal("ListMemberApiReply is nil")
	}
}

// --- GetMemberDetailApiReqBuilder 测试 ---

func TestGetMemberDetailApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		MemberId("1125922289295589").
		EmployeeNumber("D0001").
		Phone("13800000001").
		BelongEnterpriseName("子公司A").
		TaxpayerNo("91110000xxx").
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
		{"member_id", "1125922289295589"},
		{"employee_number", "D0001"},
		{"phone", "13800000001"},
		{"belong_enterprise_name", "子公司A"},
		{"taxpayer_no", "91110000xxx"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetMemberDetailApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		MemberId("1125922289295589").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("member_id") != "1125922289295589" {
		t.Errorf("member_id mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("employee_number") != "" {
		t.Errorf("employee_number should be empty, got %q", req.apiReq.QueryParams.Get("employee_number"))
	}
	if req.apiReq.QueryParams.Get("phone") != "" {
		t.Errorf("phone should be empty, got %q", req.apiReq.QueryParams.Get("phone"))
	}
	if req.apiReq.QueryParams.Get("belong_enterprise_name") != "" {
		t.Errorf("belong_enterprise_name should be empty, got %q", req.apiReq.QueryParams.Get("belong_enterprise_name"))
	}
	if req.apiReq.QueryParams.Get("taxpayer_no") != "" {
		t.Errorf("taxpayer_no should be empty, got %q", req.apiReq.QueryParams.Get("taxpayer_no"))
	}
}

func TestGetMemberDetailApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		Timestamp(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("timestamp") != "0" {
		t.Errorf("timestamp = %q, want \"0\"", req.apiReq.QueryParams.Get("timestamp"))
	}
}

func TestGetMemberDetailApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	// 仅通用参数
	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
	// 业务参数不应存在
	for _, key := range []string{"member_id", "employee_number", "phone", "belong_enterprise_name", "taxpayer_no"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- GetMemberDetailApiReply 反序列化测试 ---

func TestGetMemberDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"member_id": "1125916593600206",
			"phone": "13500007300",
			"realname": "翁舒嵘",
			"employee_number": "66824",
			"email": "",
			"system_role": 0,
			"role_ids": "1125901629154595",
			"immediate_superior_phone": "",
			"immediate_superior_eid": "",
			"residentsname": "福州",
			"use_company_money": 1,
			"total_quota": "0.00",
			"available_quota": "0.00",
			"limit_rule_list": [
				{
					"rule_name": "员工个人限额（默认）",
					"budget_cycle": 1,
					"is_accumulative": 0,
					"total_quota": 50000,
					"available_quota": 50000,
					"freeze_quota": 0
				}
			],
			"is_remark": 2,
			"budget_center_id": "1125901629164662",
			"con_department_ids": [],
			"regulation_id": ["1125934264707830"],
			"project_ids": "",
			"set_dismiss_time": "",
			"dismiss_time": "",
			"invoice_info": "",
			"legal_entity_id": "",
			"out_legal_entity_id": "",
			"rank_id": "",
			"out_rank_id": "",
			"english_surname": "",
			"english_name": "",
			"nickname": "",
			"sex": 0,
			"birth_date": "",
			"card_list": [],
			"source": "4",
			"status": 1,
			"residents_list": [
				{"id": "34", "name": "福州", "adcode": "350100"}
			],
			"cert_realname": "翁舒嵘",
			"cert_english_surname": "",
			"cert_english_name": "",
			"home_address": [
				{"city": "北京", "city_id": 1, "city_adcode": "110100", "address_name": "望京SOHO"}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply GetMemberDetailApiReply
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

	data := reply.Data
	// 基础字段
	if data.MemberId == nil || *data.MemberId != "1125916593600206" {
		t.Errorf("MemberId = %v, want 1125916593600206", data.MemberId)
	}
	if data.Phone == nil || *data.Phone != "13500007300" {
		t.Errorf("Phone = %v, want 13500007300", data.Phone)
	}
	if data.Realname == nil || *data.Realname != "翁舒嵘" {
		t.Errorf("Realname = %v, want 翁舒嵘", data.Realname)
	}
	// 新增字段验证
	if data.Status == nil || *data.Status != 1 {
		t.Errorf("Status = %v, want 1", data.Status)
	}
	if data.Source == nil || *data.Source != "4" {
		t.Errorf("Source = %v, want 4", data.Source)
	}
	if data.IsRemark == nil || *data.IsRemark != 2 {
		t.Errorf("IsRemark = %v, want 2", data.IsRemark)
	}
	// residents_list
	if len(data.ResidentsList) != 1 {
		t.Fatalf("ResidentsList len = %d, want 1", len(data.ResidentsList))
	}
	if data.ResidentsList[0].Id == nil || *data.ResidentsList[0].Id != "34" {
		t.Errorf("ResidentsList[0].Id = %v, want 34", data.ResidentsList[0].Id)
	}
	if data.ResidentsList[0].Name == nil || *data.ResidentsList[0].Name != "福州" {
		t.Errorf("ResidentsList[0].Name = %v, want 福州", data.ResidentsList[0].Name)
	}
	if data.ResidentsList[0].Adcode == nil || *data.ResidentsList[0].Adcode != "350100" {
		t.Errorf("ResidentsList[0].Adcode = %v, want 350100", data.ResidentsList[0].Adcode)
	}
	// limit_rule_list
	if len(data.LimitRuleList) != 1 {
		t.Fatalf("LimitRuleList len = %d, want 1", len(data.LimitRuleList))
	}
	if data.LimitRuleList[0].RuleName == nil || *data.LimitRuleList[0].RuleName != "员工个人限额（默认）" {
		t.Errorf("LimitRuleList[0].RuleName = %v, want 员工个人限额（默认）", data.LimitRuleList[0].RuleName)
	}
	if data.LimitRuleList[0].BudgetCycle == nil || *data.LimitRuleList[0].BudgetCycle != 1 {
		t.Errorf("LimitRuleList[0].BudgetCycle = %v, want 1", data.LimitRuleList[0].BudgetCycle)
	}
	// cert 字段
	if data.CertRealname == nil || *data.CertRealname != "翁舒嵘" {
		t.Errorf("CertRealname = %v, want 翁舒嵘", data.CertRealname)
	}
	// home_address
	if len(data.HomeAddress) != 1 {
		t.Fatalf("HomeAddress len = %d, want 1", len(data.HomeAddress))
	}
	if data.HomeAddress[0].City == nil || *data.HomeAddress[0].City != "北京" {
		t.Errorf("HomeAddress[0].City = %v, want 北京", data.HomeAddress[0].City)
	}
	if data.HomeAddress[0].CityId == nil || *data.HomeAddress[0].CityId != 1 {
		t.Errorf("HomeAddress[0].CityId = %v, want 1", data.HomeAddress[0].CityId)
	}
	if data.HomeAddress[0].CityAdcode == nil || *data.HomeAddress[0].CityAdcode != "110100" {
		t.Errorf("HomeAddress[0].CityAdcode = %v, want 110100", data.HomeAddress[0].CityAdcode)
	}
	if data.HomeAddress[0].AddressName == nil || *data.HomeAddress[0].AddressName != "望京SOHO" {
		t.Errorf("HomeAddress[0].AddressName = %v, want 望京SOHO", data.HomeAddress[0].AddressName)
	}
}

// 一人多企业场景：other_companies_list 元素=完整明细对象+belong_enterprise_name（依据服务端实现，文档无结构表）
func TestGetMemberDetailReply_OtherCompaniesList(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"member_id": "1125916593600206",
			"realname": "翁舒嵘",
			"status": 1,
			"belong_enterprise_name": "主公司A",
			"other_companies_list": [
				{
					"member_id": "1125916593600999",
					"realname": "翁舒嵘",
					"status": 4,
					"legal_entity_id": "112689540238402",
					"belong_enterprise_name": "子公司B",
					"card_list": [
						{"card_type": "1", "card_no": "encrypted_no", "expire_date": "encrypted_date"}
					]
				},
				{
					"member_id": "1125916593600888",
					"realname": "翁舒嵘",
					"status": 1,
					"belong_enterprise_name": "子公司C"
				}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply GetMemberDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	data := reply.Data
	if data == nil {
		t.Fatal("Data is nil")
	}
	// 主对象 belong_enterprise_name
	if data.BelongEnterpriseName == nil || *data.BelongEnterpriseName != "主公司A" {
		t.Errorf("BelongEnterpriseName = %v, want 主公司A", data.BelongEnterpriseName)
	}
	// other_companies_list
	if len(data.OtherCompaniesList) != 2 {
		t.Fatalf("OtherCompaniesList len = %d, want 2", len(data.OtherCompaniesList))
	}
	other0 := data.OtherCompaniesList[0]
	if other0.MemberId == nil || *other0.MemberId != "1125916593600999" {
		t.Errorf("OtherCompaniesList[0].MemberId = %v, want 1125916593600999", other0.MemberId)
	}
	if other0.Status == nil || *other0.Status != 4 {
		t.Errorf("OtherCompaniesList[0].Status = %v, want 4", other0.Status)
	}
	if other0.BelongEnterpriseName == nil || *other0.BelongEnterpriseName != "子公司B" {
		t.Errorf("OtherCompaniesList[0].BelongEnterpriseName = %v, want 子公司B", other0.BelongEnterpriseName)
	}
	// 嵌套结构与明细对象一致：card_list
	if len(other0.CardList) != 1 || other0.CardList[0].CardType == nil || *other0.CardList[0].CardType != "1" {
		t.Errorf("OtherCompaniesList[0].CardList parse failed: %v", other0.CardList)
	}
	other1 := data.OtherCompaniesList[1]
	if other1.MemberId == nil || *other1.MemberId != "1125916593600888" {
		t.Errorf("OtherCompaniesList[1].MemberId = %v, want 1125916593600888", other1.MemberId)
	}
	if other1.BelongEnterpriseName == nil || *other1.BelongEnterpriseName != "子公司C" {
		t.Errorf("OtherCompaniesList[1].BelongEnterpriseName = %v, want 子公司C", other1.BelongEnterpriseName)
	}
}

// 普通场景（非一人多企业）：无 belong_enterprise_name/other_companies_list 时解析不报错且为 nil
func TestGetMemberDetailReply_WithoutOtherCompaniesList(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206","realname":"张三"},"request_id":"r"}`
	var reply GetMemberDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.BelongEnterpriseName != nil {
		t.Errorf("BelongEnterpriseName should be nil, got %v", reply.Data.BelongEnterpriseName)
	}
	if reply.Data.OtherCompaniesList != nil {
		t.Errorf("OtherCompaniesList should be nil, got %v", reply.Data.OtherCompaniesList)
	}
}

func TestGetMemberDetailApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply GetMemberDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestGetMemberDetailApiReply_MultipleItems(t *testing.T) {
	// 测试部分字段缺失的情况
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"member_id": "1125916593600206",
			"phone": "13500007300",
			"status": 1,
			"residents_list": [{"id": "34", "name": "福州"}]
		},
		"request_id": "req_partial"
	}`

	var reply GetMemberDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	// 有值的字段
	if reply.Data.Status == nil || *reply.Data.Status != 1 {
		t.Errorf("Status = %v, want 1", reply.Data.Status)
	}
	if len(reply.Data.ResidentsList) != 1 {
		t.Fatalf("ResidentsList len = %d, want 1", len(reply.Data.ResidentsList))
	}
	// adcode 缺失应为 nil
	if reply.Data.ResidentsList[0].Adcode != nil {
		t.Errorf("ResidentsList[0].Adcode = %v, want nil", reply.Data.ResidentsList[0].Adcode)
	}
	// 缺失的新增字段应为 nil
	if reply.Data.CertRealname != nil {
		t.Errorf("CertRealname = %v, want nil", reply.Data.CertRealname)
	}
	if reply.Data.HomeAddress != nil {
		t.Errorf("HomeAddress = %v, want nil", reply.Data.HomeAddress)
	}
	if reply.Data.LimitRuleList != nil {
		t.Errorf("LimitRuleList = %v, want nil", reply.Data.LimitRuleList)
	}
}

func TestGetMemberDetailApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"member_id": "1125916593600206",
			"residents_list": [],
			"limit_rule_list": [],
			"home_address": [],
			"card_list": [],
			"con_department_ids": []
		},
		"request_id": "req_empty"
	}`

	var reply GetMemberDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.ResidentsList) != 0 {
		t.Errorf("ResidentsList len = %d, want 0", len(reply.Data.ResidentsList))
	}
	if len(reply.Data.LimitRuleList) != 0 {
		t.Errorf("LimitRuleList len = %d, want 0", len(reply.Data.LimitRuleList))
	}
	if len(reply.Data.HomeAddress) != 0 {
		t.Errorf("HomeAddress len = %d, want 0", len(reply.Data.HomeAddress))
	}
}

func TestGetMemberDetailApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetMemberDetailApiReply
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

// --- GetMemberDetail 资源方法测试 ---

func TestGetMemberDetail_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Member/detail" {
			t.Errorf("expected path /river/Member/detail, got %s", r.URL.Path)
		}
		// 验证请求参数
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("member_id") != "1125916593600206" {
			t.Errorf("member_id = %q, want 1125916593600206", r.URL.Query().Get("member_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"member_id": "1125916593600206",
				"phone": "13500007300",
				"realname": "翁舒嵘",
				"status": 1,
				"source": "4",
				"residents_list": [{"id": "34", "name": "福州", "adcode": "350100"}],
				"limit_rule_list": [{"rule_name": "默认", "budget_cycle": 1, "total_quota": 50000, "available_quota": 50000, "freeze_quota": 0}],
				"cert_realname": "翁舒嵘",
				"cert_english_surname": "",
				"cert_english_name": "",
				"home_address": [{"city": "北京", "city_id": 1, "city_adcode": "110100", "address_name": "望京SOHO"}]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		MemberId("1125916593600206").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if resp.GetMemberDetailApiReply.Errno == nil || *resp.GetMemberDetailApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetMemberDetailApiReply.Errno)
	}
	if resp.GetMemberDetailApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetMemberDetailApiReply.Data.Status == nil || *resp.GetMemberDetailApiReply.Data.Status != 1 {
		t.Errorf("Status = %v, want 1", resp.GetMemberDetailApiReply.Data.Status)
	}
	if resp.GetMemberDetailApiReply.Data.CertRealname == nil || *resp.GetMemberDetailApiReply.Data.CertRealname != "翁舒嵘" {
		t.Errorf("CertRealname = %v, want 翁舒嵘", resp.GetMemberDetailApiReply.Data.CertRealname)
	}
}

func TestGetMemberDetail_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206","residents_list":[],"limit_rule_list":[],"home_address":[],"card_list":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		MemberId("1125916593600206").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if len(resp.GetMemberDetailApiReply.Data.ResidentsList) != 0 {
		t.Errorf("ResidentsList len = %d, want 0", len(resp.GetMemberDetailApiReply.Data.ResidentsList))
	}
	if len(resp.GetMemberDetailApiReply.Data.LimitRuleList) != 0 {
		t.Errorf("LimitRuleList len = %d, want 0", len(resp.GetMemberDetailApiReply.Data.LimitRuleList))
	}
	if len(resp.GetMemberDetailApiReply.Data.HomeAddress) != 0 {
		t.Errorf("HomeAddress len = %d, want 0", len(resp.GetMemberDetailApiReply.Data.HomeAddress))
	}
}

func TestGetMemberDetail_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if resp.GetMemberDetailApiReply.Errno == nil || *resp.GetMemberDetailApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.GetMemberDetailApiReply.Errno)
	}
}

func TestGetMemberDetail_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetMemberDetailApiReply != nil {
		t.Errorf("GetMemberDetailApiReply should be nil for non-200 response")
	}
}

func TestGetMemberDetail_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206","status":1,"residents_list":[{"id":"34","name":"福州","adcode":"350100"}]},"request_id":"req_enc"}`
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

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		MemberId("1125916593600206").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if resp.GetMemberDetailApiReply.Errno == nil || *resp.GetMemberDetailApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.GetMemberDetailApiReply.Errno)
	}
	if resp.GetMemberDetailApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetMemberDetailApiReply.Data.Status == nil || *resp.GetMemberDetailApiReply.Data.Status != 1 {
		t.Errorf("Status = %v, want 1", resp.GetMemberDetailApiReply.Data.Status)
	}
	if len(resp.GetMemberDetailApiReply.Data.ResidentsList) != 1 {
		t.Fatalf("ResidentsList len = %d, want 1", len(resp.GetMemberDetailApiReply.Data.ResidentsList))
	}
	if resp.GetMemberDetailApiReply.Data.ResidentsList[0].Id == nil || *resp.GetMemberDetailApiReply.Data.ResidentsList[0].Id != "34" {
		t.Errorf("ResidentsList[0].Id = %v, want 34", resp.GetMemberDetailApiReply.Data.ResidentsList[0].Id)
	}
}

func TestGetMemberDetail_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206","phone":"13500007300"},"request_id":"req_enc256"}`
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

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if resp.GetMemberDetailApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.GetMemberDetailApiReply.Data.MemberId == nil || *resp.GetMemberDetailApiReply.Data.MemberId != "1125916593600206" {
		t.Errorf("MemberId = %v, want 1125916593600206", resp.GetMemberDetailApiReply.Data.MemberId)
	}
}

func TestGetMemberDetail_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	m := &member{option: option}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
	if resp.GetMemberDetailApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestGetMemberDetail_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"member_id":"1125916593600206"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetMemberDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := m.GetMemberDetail(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetMemberDetail() error = %v", err)
	}
	if resp.GetMemberDetailApiReply == nil {
		t.Fatal("GetMemberDetailApiReply is nil")
	}
}

func TestMemberInfoResidentsAdcodeAndHomeAddress(t *testing.T) {
	info := NewMemberInfoBuilder().
		ResidentsAdcode("110000,110100").
		HomeAddress([]HomeAddressInfo{
			*NewHomeAddressInfoBuilder().Build(),
		}).
		Build()

	if info.ResidentsAdcode == nil || *info.ResidentsAdcode != "110000,110100" {
		t.Errorf("ResidentsAdcode = %v, want 110000,110100", info.ResidentsAdcode)
	}
	if len(info.HomeAddress) != 1 {
		t.Errorf("HomeAddress len = %d, want 1", len(info.HomeAddress))
	}

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	if !bytes.Contains(data, []byte("residents_adcode")) {
		t.Error("residents_adcode should be serialized")
	}
	if !bytes.Contains(data, []byte("home_address")) {
		t.Error("home_address should be serialized")
	}

	// 部分字段：不设置为空
	partial := NewMemberInfoBuilder().Realname("王五").Build()
	if partial.ResidentsAdcode != nil {
		t.Errorf("ResidentsAdcode = %v, want nil", partial.ResidentsAdcode)
	}
	if partial.HomeAddress != nil {
		t.Errorf("HomeAddress = %v, want nil", partial.HomeAddress)
	}
}

func TestDelMemberRequestMultiIdentifierAndGroupParams(t *testing.T) {
	// 仅手机号
	req := NewDelMemberRequestBuilder().
		ClientId("test_client").
		EmployeePhone(`["+86 13800001234"]`).
		Build()
	data, _ := json.Marshal(req)
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	if m["employee_phone"] != `["+86 13800001234"]` {
		t.Errorf("employee_phone = %v", m["employee_phone"])
	}
	if _, ok := m["member_id"]; ok {
		t.Error("member_id should be omitted")
	}

	// 仅邮箱
	req2 := NewDelMemberRequestBuilder().EmployeeEmail(`["a@example.com"]`).Build()
	data2, _ := json.Marshal(req2)
	if !bytes.Contains(data2, []byte(`"employee_email":"[\"a@example.com\"]"`)) {
		t.Errorf("employee_email not serialized: %s", data2)
	}

	// 与 member_id 共存 + 集团参数
	req3 := NewDelMemberRequestBuilder().
		MemberId("123_456").
		EmployeePhone(`["+86 13800005678"]`).
		BelongEnterpriseName("某某集团").
		TaxpayerNo("91110000XXXXXXXXXX").
		Build()
	data3, _ := json.Marshal(req3)
	m3 := map[string]interface{}{}
	json.Unmarshal(data3, &m3)
	for _, key := range []string{"member_id", "employee_phone", "belong_enterprise_name", "taxpayer_no"} {
		if _, ok := m3[key]; !ok {
			t.Errorf("%s should be serialized", key)
		}
	}
	if m3["belong_enterprise_name"] != "某某集团" {
		t.Errorf("belong_enterprise_name = %v", m3["belong_enterprise_name"])
	}
}

func TestMemberCreateUpdateRequestGroupParams(t *testing.T) {
	// create 顶层
	createReq := NewCreateMemberRequestBuilder().
		ClientId("test_client").
		BelongEnterpriseName("某某集团").
		TaxpayerNo("91110000XXXXXXXXXX").
		Build()
	if createReq.BelongEnterpriseName == nil || *createReq.BelongEnterpriseName != "某某集团" {
		t.Errorf("CreateRequest BelongEnterpriseName = %v", createReq.BelongEnterpriseName)
	}
	if createReq.TaxpayerNo == nil || *createReq.TaxpayerNo != "91110000XXXXXXXXXX" {
		t.Errorf("CreateRequest TaxpayerNo = %v", createReq.TaxpayerNo)
	}

	// update 顶层 + 序列化
	updateReq := NewUpdateMemberRequestBuilder().
		BelongEnterpriseName("某某集团").
		TaxpayerNo("91110000XXXXXXXXXX").
		Build()
	data, _ := json.Marshal(updateReq)
	if !bytes.Contains(data, []byte("belong_enterprise_name")) || !bytes.Contains(data, []byte("taxpayer_no")) {
		t.Errorf("group params not serialized: %s", data)
	}

	// 不设置时省略
	empty := NewCreateMemberRequestBuilder().ClientId("test_client").Build()
	data2, _ := json.Marshal(empty)
	if bytes.Contains(data2, []byte("belong_enterprise_name")) {
		t.Error("belong_enterprise_name should be omitted")
	}
}

func TestGetMemberQuotaApiReqBuilder_GroupAndMonthlyLimit(t *testing.T) {
	req := NewGetMemberQuotaApiReqBuilder().
		ClientId("test_client").
		MemberIds("123,456").
		MonthlyLimitOnlyQuery(2).
		BelongEnterpriseName("某某集团").
		TaxpayerNo("91110000XXXXXXXXXX").
		Build()

	if got := req.apiReq.QueryParams.Get("monthly_limit_only_query"); got != "2" {
		t.Errorf("monthly_limit_only_query = %q, want 2", got)
	}
	if got := req.apiReq.QueryParams.Get("belong_enterprise_name"); got != "某某集团" {
		t.Errorf("belong_enterprise_name = %q, want 某某集团", got)
	}
	if got := req.apiReq.QueryParams.Get("taxpayer_no"); got != "91110000XXXXXXXXXX" {
		t.Errorf("taxpayer_no = %q", got)
	}

	// int32 零值仍写入
	req0 := NewGetMemberQuotaApiReqBuilder().MonthlyLimitOnlyQuery(0).Build()
	if got := req0.apiReq.QueryParams.Get("monthly_limit_only_query"); got != "0" {
		t.Errorf("monthly_limit_only_query(0) = %q, want 0", got)
	}

	// 不设置时不存在
	reqN := NewGetMemberQuotaApiReqBuilder().ClientId("c").Build()
	if _, ok := reqN.apiReq.QueryParams["monthly_limit_only_query"]; ok {
		t.Error("monthly_limit_only_query should not exist")
	}
	if _, ok := reqN.apiReq.QueryParams["belong_enterprise_name"]; ok {
		t.Error("belong_enterprise_name should not exist")
	}
}

// --- ListMemberV2（集团账户 v2）测试 ---

func TestListMemberV2ApiReqBuilder_FullParams(t *testing.T) {
	req := NewListMemberV2ApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Email("a@example.com").
		EmployeeNumber("D1001").
		Phone("+86 13800001234").
		Status("1,4").
		Length(100).
		NextToken("TOKEN_1").
		QuerySubCompany("1").
		BelongEnterpriseName("某某集团").
		TaxpayerNo("91110000XXXXXXXXXX").
		Build()

	wantKeys := map[string]string{
		"client_id": "test_client", "email": "a@example.com",
		"status": "1,4", "length": "100", "next_token": "TOKEN_1",
		"query_sub_company":      "1",
		"belong_enterprise_name": "某某集团", "taxpayer_no": "91110000XXXXXXXXXX",
	}
	for k, v := range wantKeys {
		if got := req.apiReq.QueryParams.Get(k); got != v {
			t.Errorf("query[%s] = %q, want %q", k, got, v)
		}
	}
}

func TestListMemberV2ApiReqBuilder_PartialAndZeroInt(t *testing.T) {
	// 部分参数 + int 零值/空值仍写入
	req := NewListMemberV2ApiReqBuilder().
		ClientId("c").
		Length(0).
		NextToken("").
		Build()
	if got := req.apiReq.QueryParams.Get("length"); got != "0" {
		t.Errorf("length = %q, want 0", got)
	}
	if got := req.apiReq.QueryParams.Get("next_token"); got != "" {
		t.Errorf("next_token = %q, want empty", got)
	}
	// 未设置的不存在
	for _, k := range []string{"email", "phone", "belong_enterprise_name", "taxpayer_no", "query_sub_company"} {
		if _, ok := req.apiReq.QueryParams[k]; ok {
			t.Errorf("%s should not exist", k)
		}
	}
}

func TestListMemberV2Reply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0, "errmsg": "SUCCESS",
		"data": {
			"next_token": "NEXT_2",
			"records": [{
				"id": "1125922289295589",
				"set_dismiss_time": "2026-01-01",
				"belong_enterprise_name": "某某集团",
				"status": 1,
				"con_department_ids": ["1","2"]
			}],
			"progress": 1,
			"company_total": 3,
			"company_list_has_change": false
		},
		"request_id": "req_v2"
	}`
	var reply ListMemberV2ApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil || len(reply.Data.Records) != 1 {
		t.Fatalf("Data.Records parse failed")
	}
	rec := reply.Data.Records[0]
	if rec.Id == nil || *rec.Id != "1125922289295589" {
		t.Errorf("Id = %v", rec.Id)
	}
	if rec.SetDismissTime == nil || *rec.SetDismissTime != "2026-01-01" {
		t.Errorf("SetDismissTime = %v", rec.SetDismissTime)
	}
	if rec.BelongEnterpriseName == nil || *rec.BelongEnterpriseName != "某某集团" {
		t.Errorf("BelongEnterpriseName = %v", rec.BelongEnterpriseName)
	}
	if reply.Data.NextToken == nil || *reply.Data.NextToken != "NEXT_2" {
		t.Errorf("NextToken = %v", reply.Data.NextToken)
	}
	if reply.Data.Progress == nil || *reply.Data.Progress != 1 {
		t.Errorf("Progress = %v", reply.Data.Progress)
	}
	if reply.Data.CompanyTotal == nil || *reply.Data.CompanyTotal != 3 {
		t.Errorf("CompanyTotal = %v", reply.Data.CompanyTotal)
	}
	if reply.Data.CompanyListHasChange == nil || *reply.Data.CompanyListHasChange {
		t.Errorf("CompanyListHasChange = %v", reply.Data.CompanyListHasChange)
	}
}

func TestListMemberV2Reply_EmptyAndError(t *testing.T) {
	// 翻页结束：next_token 空字符串
	var r1 ListMemberV2ApiReply
	if err := json.Unmarshal([]byte(`{"errno":0,"data":{"next_token":"","records":[]}}`), &r1); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if r1.Data.NextToken == nil || *r1.Data.NextToken != "" {
		t.Errorf("NextToken = %v, want empty string pointer", r1.Data.NextToken)
	}
	// errno 非 0
	var r2 ListMemberV2ApiReply
	if err := json.Unmarshal([]byte(`{"errno":10003,"errmsg":"param error"}`), &r2); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if r2.Errno == nil || *r2.Errno != 10003 {
		t.Errorf("Errno = %v", r2.Errno)
	}
	if r2.Data != nil {
		t.Errorf("Data = %v, want nil", r2.Data)
	}
}

func TestListMemberV2_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v2/member/get" {
			t.Errorf("expected path /open-apis/v2/member/get, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"next_token":"N1","records":[{"id":"1","realname":"张三","belong_enterprise_name":"某某集团"}],"progress":0.5,"company_total":2,"company_list_has_change":false},"request_id":"r"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}

	req := NewListMemberV2ApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		Length(10).
		QuerySubCompany("1").
		BelongEnterpriseName("某某集团").
		Build()

	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMemberV2() error = %v", err)
	}
	if resp.ListMemberV2ApiReply == nil {
		t.Fatal("ListMemberV2ApiReply is nil")
	}
	if resp.ListMemberV2ApiReply.Data == nil || len(resp.ListMemberV2ApiReply.Data.Records) != 1 {
		t.Fatal("records parse failed")
	}
	if rec := resp.ListMemberV2ApiReply.Data.Records[0]; rec.BelongEnterpriseName == nil || *rec.BelongEnterpriseName != "某某集团" {
		t.Errorf("BelongEnterpriseName = %v", rec.BelongEnterpriseName)
	}
}

func TestListMemberV2_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error"}`))
	}))
	defer testServer.Close()
	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}
	req := NewListMemberV2ApiReqBuilder().ClientId("c").Build()
	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.ListMemberV2ApiReply.Data != nil {
		t.Errorf("Data should be nil on errno!=0")
	}
}

func TestListMemberV2_AES128Encrypted(t *testing.T) {
	plain := `{"errno":0,"errmsg":"SUCCESS","data":{"next_token":"N","records":[{"id":"9","set_dismiss_time":"2026-01-01"}]},"request_id":"r"}`
	key := make([]byte, 16)
	for i := range key {
		key[i] = byte(i + 1)
	}
	cipher, err := core.AESEncryptECB([]byte(plain), key)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(cipher)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	m := &member{option: option}

	req := NewListMemberV2ApiReqBuilder().ClientId("c").Build()
	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMemberV2() error = %v", err)
	}
	if resp.ListMemberV2ApiReply.Data == nil || len(resp.ListMemberV2ApiReply.Data.Records) != 1 {
		t.Fatal("decrypt+parse failed")
	}
	if rec := resp.ListMemberV2ApiReply.Data.Records[0]; rec.SetDismissTime == nil || *rec.SetDismissTime != "2026-01-01" {
		t.Errorf("SetDismissTime = %v", rec.SetDismissTime)
	}
}

func TestListMemberV2_AES256Encrypted(t *testing.T) {
	plain := `{"errno":0,"data":{"next_token":"","records":[{"id":"7"}]}}`
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 2)
	}
	cipher, err := core.AESEncryptECB([]byte(plain), key)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(cipher)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	m := &member{option: option}

	req := NewListMemberV2ApiReqBuilder().ClientId("c").Build()
	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMemberV2() error = %v", err)
	}
	if resp.ListMemberV2ApiReply.Data == nil || len(resp.ListMemberV2ApiReply.Data.Records) != 1 {
		t.Fatal("decrypt+parse failed")
	}
}

func TestListMemberV2_EncryptedNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"data":{"records":[{"id":"5"}]}}`))
	}))
	defer testServer.Close()

	key := make([]byte, 16)
	option := newMemberTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	m := &member{option: option}

	req := NewListMemberV2ApiReqBuilder().ClientId("c").Build()
	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListMemberV2() error = %v", err)
	}
	if resp.ListMemberV2ApiReply.Data == nil || len(resp.ListMemberV2ApiReply.Data.Records) != 1 {
		t.Fatal("plain fallback parse failed")
	}
}

func TestListMemberV2_HttpNot200(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`internal error`))
	}))
	defer testServer.Close()
	option := newMemberTestOption(testServer.URL)
	m := &member{option: option}
	req := NewListMemberV2ApiReqBuilder().ClientId("c").Build()
	resp, err := m.ListMemberV2(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.ListMemberV2ApiReply != nil {
		t.Error("ApiReply should not be set on non-200")
	}
}

func TestMemberV2RecordBuilder_AllFields(t *testing.T) {
	rec := NewMemberV2RecordBuilder().
		Id("1125922289295589").
		Phone("+86 13800001234").
		Realname("张三").
		EmployeeNumber("D1001").
		ThirdUserId("ext_1").
		Email("a@example.com").
		Department("研发").
		BranchName("北京").
		SystemRole(2).
		RoleIds("1_2").
		Residentsname("青岛").
		UseCompanyMoney(1).
		TotalQuota("937.70").
		AvailableQuota("100.00").
		IsRemark("1").
		BudgetCenterId("2001").
		SetDismissTime("2026-01-01").
		DismissTime("").
		ImmediateSuperiorPhone("+86 13900005678").
		InvoiceInfo("备注").
		LegalEntityId("3001").
		RankId("4001").
		OutRankId("P7").
		EnglishSurname("Zhang").
		EnglishName("San").
		Nickname("昵称").
		Sex(1).
		BirthDate("1990-01-01").
		Source("4").
		ImmediateSuperiorEid("D1000").
		OutLegalEntityId("OUT_3001").
		CertRealname("张三").
		CertEnglishSurname("Zhang").
		CertEnglishName("San").
		Status(1).
		MonthQuota("937.70").
		BelongEnterpriseName("某某集团").
		ConDepartmentIds([]string{"1", "2"}).
		ResidentsList([]MemberV2ResidentsListInfo{*NewMemberV2ResidentsListInfoBuilder().Build()}).
		LimitRuleList([]MemberV2LimitRuleInfo{*NewMemberV2LimitRuleInfoBuilder().Build()}).
		UseCarConfig([]string{"r1"}).
		RegulationId([]string{"g1"}).
		GuestCarRight(0).
		CardList([]MemberV2CardInfo{*NewMemberV2CardInfoBuilder().Build()}).
		Build()

	// 关键字段断言（覆盖各类型）
	if rec.Id == nil || *rec.Id != "1125922289295589" {
		t.Errorf("Id = %v", rec.Id)
	}
	if rec.SystemRole == nil || *rec.SystemRole != 2 {
		t.Errorf("SystemRole = %v", rec.SystemRole)
	}
	if rec.TotalQuota == nil || *rec.TotalQuota != "937.70" {
		t.Errorf("TotalQuota = %v", rec.TotalQuota)
	}
	if rec.BelongEnterpriseName == nil || *rec.BelongEnterpriseName != "某某集团" {
		t.Errorf("BelongEnterpriseName = %v", rec.BelongEnterpriseName)
	}
	if rec.SetDismissTime == nil || *rec.SetDismissTime != "2026-01-01" {
		t.Errorf("SetDismissTime = %v", rec.SetDismissTime)
	}
	if len(rec.ConDepartmentIds) != 2 || len(rec.ResidentsList) != 1 || len(rec.LimitRuleList) != 1 ||
		len(rec.UseCarConfig) != 1 || len(rec.RegulationId) != 1 || len(rec.CardList) != 1 {
		t.Error("slice fields should be set")
	}

	// 部分字段：未设置应为 nil/空
	partial := NewMemberV2RecordBuilder().Id("1").Build()
	if partial.Phone != nil || partial.Realname != nil || partial.BelongEnterpriseName != nil {
		t.Error("unset fields should be nil")
	}
	if partial.ConDepartmentIds != nil || partial.CardList != nil {
		t.Error("unset slices should be nil")
	}
}

func TestListMemberV2Reply_MultipleRecordsPartialFields(t *testing.T) {
	jsonData := `{"errno":0,"data":{"records":[
		{"id":"1","realname":"张三","belong_enterprise_name":"某某集团"},
		{"id":"2"},
		{"id":"3","status":4,"month_quota":"50.00"}
	]}}`
	var reply ListMemberV2ApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 3 {
		t.Fatalf("Records len = %d, want 3", len(reply.Data.Records))
	}
	if reply.Data.Records[1].Realname != nil {
		t.Error("Records[1].Realname should be nil")
	}
	if reply.Data.Records[2].Status == nil || *reply.Data.Records[2].Status != 4 {
		t.Errorf("Records[2].Status = %v", reply.Data.Records[2].Status)
	}
}

func TestListMemberV2Reply_MissingDataField(t *testing.T) {
	var reply ListMemberV2ApiReply
	if err := json.Unmarshal([]byte(`{"errno":10003,"errmsg":"param error","request_id":"r"}`), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestListMemberV2_EmptyRecords(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"next_token":"","records":[],"progress":1,"company_total":0},"request_id":"r"}`))
	}))
	defer testServer.Close()
	m := &member{option: newMemberTestOption(testServer.URL)}
	resp, err := m.ListMemberV2(context.Background(), NewListMemberV2ApiReqBuilder().ClientId("c").Build(), nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.ListMemberV2ApiReply.Data == nil || len(resp.ListMemberV2ApiReply.Data.Records) != 0 {
		t.Fatal("empty records parse failed")
	}
}

func TestListMemberV2_CustomHeader(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", r.Header.Get("X-Custom-Header"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"data":{"records":[]},"request_id":"r"}`))
	}))
	defer testServer.Close()
	m := &member{option: newMemberTestOption(testServer.URL)}

	reqOption := &core.ReqOption{Header: http.Header{}}
	reqOption.Header.Set("X-Custom-Header", "custom-value")
	_, err := m.ListMemberV2(context.Background(), NewListMemberV2ApiReqBuilder().ClientId("c").Build(), reqOption)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
}

// --- CreateMemberReply 失败场景字段测试 ---

func TestCreateMemberReplyBuilder(t *testing.T) {
	reply := NewCreateMemberReplyBuilder().
		Id(1125921312180375).
		Phone("13800000001").
		MemberId("1125921312180376").
		Status(4).
		Build()

	if reply.Id != 1125921312180375 {
		t.Errorf("Id = %d, want 1125921312180375", reply.Id)
	}
	if reply.Phone == nil || *reply.Phone != "13800000001" {
		t.Errorf("Phone = %v, want 13800000001", reply.Phone)
	}
	if reply.MemberId == nil || *reply.MemberId != "1125921312180376" {
		t.Errorf("MemberId = %v, want 1125921312180376", reply.MemberId)
	}
	if reply.Status == nil || *reply.Status != 4 {
		t.Errorf("Status = %v, want 4", reply.Status)
	}

	// 部分字段
	partial := NewCreateMemberReplyBuilder().Id(123).Build()
	if partial.Phone != nil || partial.MemberId != nil || partial.Status != nil {
		t.Errorf("partial build should leave Phone/MemberId/Status nil, got %+v", partial)
	}
}

func TestCreateMemberReply_Deserialization(t *testing.T) {
	// 新增成功场景：id/phone，无 member_id/status
	success := `{"id": 1125921312180375, "phone": "13800000001"}`
	var okReply CreateMemberReply
	if err := json.Unmarshal([]byte(success), &okReply); err != nil {
		t.Fatalf("Unmarshal success case failed: %v", err)
	}
	if okReply.Id != 1125921312180375 {
		t.Errorf("Id = %d, want 1125921312180375", okReply.Id)
	}
	if okReply.MemberId != nil || okReply.Status != nil {
		t.Errorf("success case MemberId/Status should be nil, got %+v", okReply)
	}

	// 员工已存在场景：member_id 为 number（文档 int64），须容错转 string
	exists := `{"member_id": 1125921312180376, "phone": "13800000001", "status": 4}`
	var exReply CreateMemberReply
	if err := json.Unmarshal([]byte(exists), &exReply); err != nil {
		t.Fatalf("Unmarshal exists case failed: %v", err)
	}
	if exReply.MemberId == nil || *exReply.MemberId != "1125921312180376" {
		t.Errorf("MemberId = %v, want 1125921312180376", exReply.MemberId)
	}
	if exReply.Status == nil || *exReply.Status != 4 {
		t.Errorf("Status = %v, want 4", exReply.Status)
	}

	// member_id 为 string 形式（服务端历史行为不确定，两种形式均须兼容）
	existsStr := `{"member_id": "1125921312180377", "status": 6}`
	var strReply CreateMemberReply
	if err := json.Unmarshal([]byte(existsStr), &strReply); err != nil {
		t.Fatalf("Unmarshal string member_id case failed: %v", err)
	}
	if strReply.MemberId == nil || *strReply.MemberId != "1125921312180377" {
		t.Errorf("MemberId = %v, want 1125921312180377", strReply.MemberId)
	}
}

// --- v2 独立嵌套类型 Builder 测试（禁止与 /river 类型复用） ---

func TestMemberV2NestedTypesBuilders(t *testing.T) {
	res := NewMemberV2ResidentsListInfoBuilder().Id("4").Name("青岛").Adcode("370200").Build()
	if res.Id == nil || *res.Id != "4" || res.Name == nil || *res.Name != "青岛" || res.Adcode == nil || *res.Adcode != "370200" {
		t.Errorf("MemberV2ResidentsListInfo build failed: %v", res)
	}
	// 容错反序列化：Id 收 number
	var resParsed MemberV2ResidentsListInfo
	if err := json.Unmarshal([]byte(`{"id":14,"name":"大连"}`), &resParsed); err != nil {
		t.Fatalf("residents SmartDecode failed: %v", err)
	}
	if resParsed.Id == nil || *resParsed.Id != "14" {
		t.Errorf("residents Id = %v, want 14", resParsed.Id)
	}

	limit := NewMemberV2LimitRuleInfoBuilder().RuleName("默认").BudgetCycle(1).TotalQuota(50000).AvailableQuota(50000).FreezeQuota(0).LimitManagementScope(0).Build()
	if limit.RuleName == nil || *limit.RuleName != "默认" || limit.AvailableQuota == nil || *limit.AvailableQuota != 50000 {
		t.Errorf("MemberV2LimitRuleInfo build failed: %v", limit)
	}

	card := NewMemberV2CardInfoBuilder().CardType("1").CardNo("enc_no").ExpireDate("enc_date").Build()
	if card.CardType == nil || *card.CardType != "1" || card.CardNo == nil || *card.CardNo != "enc_no" {
		t.Errorf("MemberV2CardInfo build failed: %v", card)
	}
	partial := NewMemberV2CardInfoBuilder().CardType("2").Build()
	if partial.CardNo != nil || partial.ExpireDate != nil {
		t.Error("unset fields should be nil")
	}
}
