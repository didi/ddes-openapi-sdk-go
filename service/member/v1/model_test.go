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
