package v1

import (
	"encoding/json"
	"testing"
)

func TestOutTravelerListApiReqBuilder_QueryParams(t *testing.T) {
	req := NewOutTravelerListApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		ProjectId("1125904357323169").
		OutBudgetId("budget_001").
		Page(1).
		PageSize(20).
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
		{"project_id", "1125904357323169"},
		{"out_budget_id", "budget_001"},
		{"page", "1"},
		{"page_size", "20"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestOutTravelerListApiReqBuilder_PartialParams(t *testing.T) {
	req := NewOutTravelerListApiReqBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Page(1).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("project_id") != "1125904357323169" {
		t.Errorf("project_id mismatch")
	}
	if req.apiReq.QueryParams.Get("page") != "1" {
		t.Errorf("page mismatch")
	}
	// 未设置的 out_budget_id 应为空
	if req.apiReq.QueryParams.Get("out_budget_id") != "" {
		t.Errorf("out_budget_id should be empty, got %q", req.apiReq.QueryParams.Get("out_budget_id"))
	}
}

func TestOutTravelerListApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"out_travelers": [
				{
					"traveler_id": "4503599690839935",
					"out_traveler_id": "out_traveler_id_11",
					"phone": "+852 000****8060",
					"name": "TYPE_3_外部_11",
					"english_surname": "Wai",
					"english_name": "BuCuXingRen",
					"remark": "备注",
					"sex": 1,
					"birth_date": "2000-01-01",
					"card_list": [
						{
							"card_no": "110101199001011234",
							"card_type": 1,
							"expire_date": "2030-12-31"
						}
					],
					"related_employees": [
						{
							"related_employee_id": "emp_001",
							"related_employee_phone": "13800000001",
							"related_employee_employee_number": "E001",
							"related_employee_email": "emp@test.com"
						}
					],
					"project_id": "4503600193081934",
					"project_name": "项目_test_8",
					"out_budget_id": "xm_test_8"
				}
			],
			"total": 1
		},
		"request_id": "test_request_id"
	}`

	var reply OutTravelerListApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.Data.Total == nil || *reply.Data.Total != 1 {
		t.Errorf("Total = %v, want 1", reply.Data.Total)
	}
	if len(reply.Data.OutTravelers) != 1 {
		t.Fatalf("OutTravelers len = %d, want 1", len(reply.Data.OutTravelers))
	}

	traveler := reply.Data.OutTravelers[0]
	if traveler.TravelerId == nil || *traveler.TravelerId != "4503599690839935" {
		t.Errorf("TravelerId = %v, want 4503599690839935", traveler.TravelerId)
	}
	if traveler.OutTravelerId == nil || *traveler.OutTravelerId != "out_traveler_id_11" {
		t.Errorf("OutTravelerId = %v, want out_traveler_id_11", traveler.OutTravelerId)
	}
	if traveler.Phone == nil || *traveler.Phone != "+852 000****8060" {
		t.Errorf("Phone = %v, want +852 000****8060", traveler.Phone)
	}
	if traveler.Name == nil || *traveler.Name != "TYPE_3_外部_11" {
		t.Errorf("Name = %v, want TYPE_3_外部_11", traveler.Name)
	}
	if traveler.Sex == nil || *traveler.Sex != 1 {
		t.Errorf("Sex = %v, want 1", traveler.Sex)
	}
	if traveler.ProjectId == nil || *traveler.ProjectId != "4503600193081934" {
		t.Errorf("ProjectId = %v, want 4503600193081934", traveler.ProjectId)
	}
	if traveler.ProjectName == nil || *traveler.ProjectName != "项目_test_8" {
		t.Errorf("ProjectName = %v, want 项目_test_8", traveler.ProjectName)
	}

	// CardList
	if len(traveler.CardList) != 1 {
		t.Fatalf("CardList len = %d, want 1", len(traveler.CardList))
	}
	card := traveler.CardList[0]
	if card.CardNo == nil || *card.CardNo != "110101199001011234" {
		t.Errorf("CardNo = %v, want 110101199001011234", card.CardNo)
	}
	if card.CardType == nil || *card.CardType != 1 {
		t.Errorf("CardType = %v, want 1", card.CardType)
	}

	// RelatedEmployees
	if len(traveler.RelatedEmployees) != 1 {
		t.Fatalf("RelatedEmployees len = %d, want 1", len(traveler.RelatedEmployees))
	}
	emp := traveler.RelatedEmployees[0]
	if emp.RelatedEmployeeId == nil || *emp.RelatedEmployeeId != "emp_001" {
		t.Errorf("RelatedEmployeeId = %v, want emp_001", emp.RelatedEmployeeId)
	}
	if emp.RelatedEmployeeEmail == nil || *emp.RelatedEmployeeEmail != "emp@test.com" {
		t.Errorf("RelatedEmployeeEmail = %v, want emp@test.com", emp.RelatedEmployeeEmail)
	}
}

func TestOutTravelerListApiReply_NullData(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "未找到对应的项目",
		"request_id": "test_request_id"
	}`

	var reply OutTravelerListApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestOutTravelerInfoBuilder(t *testing.T) {
	info := NewOutTravelerInfoBuilder().
		TravelerId("4503599690839935").
		OutTravelerId("out_001").
		Phone("+852 12345678").
		Name("测试").
		Sex(1).
		Build()

	if info.TravelerId == nil || *info.TravelerId != "4503599690839935" {
		t.Errorf("TravelerId = %v, want 4503599690839935", info.TravelerId)
	}
	if info.Sex == nil || *info.Sex != 1 {
		t.Errorf("Sex = %v, want 1", info.Sex)
	}
	// 未设置字段应为 nil
	if info.EnglishSurname != nil {
		t.Errorf("EnglishSurname = %v, want nil", info.EnglishSurname)
	}
	if info.BirthDate != nil {
		t.Errorf("BirthDate = %v, want nil", info.BirthDate)
	}
}

func TestCardInfoBuilder(t *testing.T) {
	card := NewCardInfoBuilder().
		CardNo("110101199001011234").
		CardType(1).
		ExpireDate("2030-12-31").
		Build()

	if card.CardNo == nil || *card.CardNo != "110101199001011234" {
		t.Errorf("CardNo = %v, want 110101199001011234", card.CardNo)
	}
	if card.CardType == nil || *card.CardType != 1 {
		t.Errorf("CardType = %v, want 1", card.CardType)
	}
}

func TestRelatedEmployeeInfoBuilder(t *testing.T) {
	emp := NewRelatedEmployeeInfoBuilder().
		RelatedEmployeeId("emp_001").
		RelatedEmployeePhone("13800000001").
		RelatedEmployeeEmployeeNumber("E001").
		RelatedEmployeeEmail("emp@test.com").
		Build()

	if emp.RelatedEmployeeId == nil || *emp.RelatedEmployeeId != "emp_001" {
		t.Errorf("RelatedEmployeeId = %v, want emp_001", emp.RelatedEmployeeId)
	}
	if emp.RelatedEmployeeEmployeeNumber == nil || *emp.RelatedEmployeeEmployeeNumber != "E001" {
		t.Errorf("RelatedEmployeeEmployeeNumber = %v, want E001", emp.RelatedEmployeeEmployeeNumber)
	}
}
