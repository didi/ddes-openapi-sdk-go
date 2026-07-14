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

func TestGetProjectDetailApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		ProjectId("1125904357323169").
		ProjectName("测试项目").
		ProjectCode("CODE001").
		Offset(0).
		Lenth(20).
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
		{"project_id", "1125904357323169"},
		{"project_name", "测试项目"},
		{"project_code", "CODE001"},
		{"offset", "0"},
		{"lenth", "20"},
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

func TestGetProjectDetailApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Offset(0).
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("project_id") != "1125904357323169" {
		t.Errorf("project_id mismatch")
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("project_name") != "" {
		t.Errorf("project_name should be empty, got %q", req.apiReq.QueryParams.Get("project_name"))
	}
	if req.apiReq.QueryParams.Get("belong_enterprise_name") != "" {
		t.Errorf("belong_enterprise_name should be empty, got %q", req.apiReq.QueryParams.Get("belong_enterprise_name"))
	}
}

func TestGetProjectDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
				"member_id": "1125923579325",
				"phone": "13800000001",
				"employee_number": "D0001",
				"email": "test@example.com"
			}
		],
		"request_id": "test_request_id"
	}`

	var reply GetProjectDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}

	member := reply.Data[0]
	if member.MemberId == nil || *member.MemberId != "1125923579325" {
		t.Errorf("MemberId = %v, want 1125923579325", member.MemberId)
	}
	if member.Phone == nil || *member.Phone != "13800000001" {
		t.Errorf("Phone = %v, want 13800000001", member.Phone)
	}
	if member.EmployeeNumber == nil || *member.EmployeeNumber != "D0001" {
		t.Errorf("EmployeeNumber = %v, want D0001", member.EmployeeNumber)
	}
	if member.Email == nil || *member.Email != "test@example.com" {
		t.Errorf("Email = %v, want test@example.com", member.Email)
	}
}

func TestGetProjectDetailApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply GetProjectDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestProjectDetailMemberInfoBuilder(t *testing.T) {
	info := NewProjectDetailMemberInfoBuilder().
		MemberId("1125923579325").
		Phone("13800000001").
		EmployeeNumber("D0001").
		Email("test@example.com").
		Build()

	if info.MemberId == nil || *info.MemberId != "1125923579325" {
		t.Errorf("MemberId = %v, want 1125923579325", info.MemberId)
	}
	if info.Phone == nil || *info.Phone != "13800000001" {
		t.Errorf("Phone = %v, want 13800000001", info.Phone)
	}
	if info.EmployeeNumber == nil || *info.EmployeeNumber != "D0001" {
		t.Errorf("EmployeeNumber = %v, want D0001", info.EmployeeNumber)
	}
	if info.Email == nil || *info.Email != "test@example.com" {
		t.Errorf("Email = %v, want test@example.com", info.Email)
	}

	// 部分设置
	info2 := NewProjectDetailMemberInfoBuilder().
		MemberId("1125923579325").
		Build()

	if info2.MemberId == nil || *info2.MemberId != "1125923579325" {
		t.Errorf("MemberId = %v, want 1125923579325", info2.MemberId)
	}
	if info2.Phone != nil {
		t.Errorf("Phone = %v, want nil", info2.Phone)
	}
}

// --- GetProjectDetail 资源方法测试 ---

func newTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestGetProjectDetail_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/Project/detail" {
			t.Errorf("expected path /river/Project/detail, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
				{"member_id": "1125923579325", "phone": "13800000001", "employee_number": "D0001", "email": "test@example.com"},
				{"member_id": "1125923579326", "phone": "13800000002", "employee_number": "D0002", "email": "test2@example.com"}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		ProjectId("1125904357323169").
		Offset(0).
		Lenth(20).
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if resp.GetProjectDetailApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetProjectDetailApiReply.Errno)
	}
	if len(resp.GetProjectDetailApiReply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(resp.GetProjectDetailApiReply.Data))
	}
	if resp.GetProjectDetailApiReply.Data[0].MemberId == nil || *resp.GetProjectDetailApiReply.Data[0].MemberId != "1125923579325" {
		t.Errorf("Data[0].MemberId = %v, want 1125923579325", resp.GetProjectDetailApiReply.Data[0].MemberId)
	}
	if resp.GetProjectDetailApiReply.Data[1].EmployeeNumber == nil || *resp.GetProjectDetailApiReply.Data[1].EmployeeNumber != "D0002" {
		t.Errorf("Data[1].EmployeeNumber = %v, want D0002", resp.GetProjectDetailApiReply.Data[1].EmployeeNumber)
	}
}

func TestGetProjectDetail_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if len(resp.GetProjectDetailApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.GetProjectDetailApiReply.Data))
	}
}

func TestGetProjectDetail_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if resp.GetProjectDetailApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.GetProjectDetailApiReply.Errno)
	}
}

func TestGetProjectDetail_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.GetProjectDetailApiReply != nil {
		t.Errorf("GetProjectDetailApiReply should be nil for non-200 response")
	}
}

func TestGetProjectDetail_WithEncryption_AES128(t *testing.T) {
	// 准备加密数据
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"member_id":"1125923579325","phone":"13800000001","employee_number":"D0001","email":"test@example.com"}],"request_id":"req_enc"}`
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

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if resp.GetProjectDetailApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetProjectDetailApiReply.Errno)
	}
	if len(resp.GetProjectDetailApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.GetProjectDetailApiReply.Data))
	}
	if resp.GetProjectDetailApiReply.Data[0].MemberId == nil || *resp.GetProjectDetailApiReply.Data[0].MemberId != "1125923579325" {
		t.Errorf("Data[0].MemberId = %v, want 1125923579325", resp.GetProjectDetailApiReply.Data[0].MemberId)
	}
}

func TestGetProjectDetail_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"member_id":"1125923579325","phone":"13800000001"}],"request_id":"req_enc256"}`
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

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if len(resp.GetProjectDetailApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.GetProjectDetailApiReply.Data))
	}
	if resp.GetProjectDetailApiReply.Data[0].Phone == nil || *resp.GetProjectDetailApiReply.Data[0].Phone != "13800000001" {
		t.Errorf("Data[0].Phone = %v, want 13800000001", resp.GetProjectDetailApiReply.Data[0].Phone)
	}
}

func TestGetProjectDetail_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"member_id":"1125923579325"}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	p := &project{option: option}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
	if len(resp.GetProjectDetailApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.GetProjectDetailApiReply.Data))
	}
}

func TestGetProjectDetail_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := p.GetProjectDetail(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("GetProjectDetail() error = %v", err)
	}
	if resp.GetProjectDetailApiReply == nil {
		t.Fatal("GetProjectDetailApiReply is nil")
	}
}

// --- 更多反序列化边界场景 ---

func TestGetProjectDetailApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"member_id": "1001", "phone": "13800000001", "employee_number": "D0001", "email": "a@test.com"},
			{"member_id": "1002", "phone": "13800000002", "employee_number": "D0002", "email": "b@test.com"},
			{"member_id": "1003", "phone": "13800000003"}
		],
		"request_id": "req_multi"
	}`

	var reply GetProjectDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 3 {
		t.Fatalf("Data len = %d, want 3", len(reply.Data))
	}
	// 第3条缺少 email 和 employee_number，应为 nil
	if reply.Data[2].Email != nil {
		t.Errorf("Data[2].Email = %v, want nil", reply.Data[2].Email)
	}
	if reply.Data[2].EmployeeNumber != nil {
		t.Errorf("Data[2].EmployeeNumber = %v, want nil", reply.Data[2].EmployeeNumber)
	}
	if reply.Data[2].MemberId == nil || *reply.Data[2].MemberId != "1003" {
		t.Errorf("Data[2].MemberId = %v, want 1003", reply.Data[2].MemberId)
	}
}

func TestGetProjectDetailApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply GetProjectDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestGetProjectDetailApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply GetProjectDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestGetProjectDetailApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		CompanyId("test_company").
		Offset(0).
		Lenth(0).
		Build()

	// int32 零值也应该被设置到 query params 中
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
	if req.apiReq.QueryParams.Get("lenth") != "0" {
		t.Errorf("lenth = %q, want \"0\"", req.apiReq.QueryParams.Get("lenth"))
	}
}

func TestGetProjectDetailApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetProjectDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
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
	for _, key := range []string{"project_id", "project_name", "project_code", "offset", "lenth", "belong_enterprise_name", "taxpayer_no"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// --- UpdateMember 模型层测试 ---

func TestUpdateMemberRequestBuilder(t *testing.T) {
	req := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		ProjectId("1125904357323169").
		ProjectCode("CODE001").
		ProjectName("测试项目").
		MemberIds("emp001,emp002").
		MemberType(1).
		MemberValues("E001,E002").
		BelongEnterpriseName("子公司A").
		TaxpayerNo("91110000xxx").
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
	if req.ProjectId == nil || *req.ProjectId != "1125904357323169" {
		t.Errorf("ProjectId = %v, want 1125904357323169", req.ProjectId)
	}
	if req.ProjectCode == nil || *req.ProjectCode != "CODE001" {
		t.Errorf("ProjectCode = %v, want CODE001", req.ProjectCode)
	}
	if req.ProjectName == nil || *req.ProjectName != "测试项目" {
		t.Errorf("ProjectName = %v, want 测试项目", req.ProjectName)
	}
	if req.MemberIds == nil || *req.MemberIds != "emp001,emp002" {
		t.Errorf("MemberIds = %v, want emp001,emp002", req.MemberIds)
	}
	if req.MemberType == nil || *req.MemberType != 1 {
		t.Errorf("MemberType = %v, want 1", req.MemberType)
	}
	if req.MemberValues == nil || *req.MemberValues != "E001,E002" {
		t.Errorf("MemberValues = %v, want E001,E002", req.MemberValues)
	}
	if req.BelongEnterpriseName == nil || *req.BelongEnterpriseName != "子公司A" {
		t.Errorf("BelongEnterpriseName = %v, want 子公司A", req.BelongEnterpriseName)
	}
	if req.TaxpayerNo == nil || *req.TaxpayerNo != "91110000xxx" {
		t.Errorf("TaxpayerNo = %v, want 91110000xxx", req.TaxpayerNo)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
}

func TestUpdateMemberRequestBuilder_PartialParams(t *testing.T) {
	req := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		MemberIds("emp001").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ProjectId == nil || *req.ProjectId != "1125904357323169" {
		t.Errorf("ProjectId = %v, want 1125904357323169", req.ProjectId)
	}
	// 未设置字段应为 nil
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.MemberType != nil {
		t.Errorf("MemberType = %v, want nil", req.MemberType)
	}
	if req.BelongEnterpriseName != nil {
		t.Errorf("BelongEnterpriseName = %v, want nil", req.BelongEnterpriseName)
	}
}

func TestUpdateMemberErrorInfoBuilder(t *testing.T) {
	info := NewUpdateMemberErrorInfoBuilder().
		ErrorMsg("员工不在该企业中").
		ErrorMemberIds([]string{"emp001", "emp002"}).
		ErrorMemberValues([]string{"E001", "E002"}).
		Build()

	if info.ErrorMsg == nil || *info.ErrorMsg != "员工不在该企业中" {
		t.Errorf("ErrorMsg = %v, want 员工不在该企业中", info.ErrorMsg)
	}
	if len(info.ErrorMemberIds) != 2 || info.ErrorMemberIds[0] != "emp001" {
		t.Errorf("ErrorMemberIds = %v, want [emp001 emp002]", info.ErrorMemberIds)
	}
	if len(info.ErrorMemberValues) != 2 || info.ErrorMemberValues[0] != "E001" {
		t.Errorf("ErrorMemberValues = %v, want [E001 E002]", info.ErrorMemberValues)
	}

	// 部分设置
	info2 := NewUpdateMemberErrorInfoBuilder().
		ErrorMsg("该员工已经在这个项目中").
		Build()

	if info2.ErrorMsg == nil || *info2.ErrorMsg != "该员工已经在这个项目中" {
		t.Errorf("ErrorMsg = %v, want 该员工已经在这个项目中", info2.ErrorMsg)
	}
	if info2.ErrorMemberIds != nil {
		t.Errorf("ErrorMemberIds = %v, want nil", info2.ErrorMemberIds)
	}
	if info2.ErrorMemberValues != nil {
		t.Errorf("ErrorMemberValues = %v, want nil", info2.ErrorMemberValues)
	}
}

func TestUpdateMemberApiReqBuilder_FullParams(t *testing.T) {
	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		ProjectId("1125904357323169").
		MemberIds("emp001,emp002").
		Sign("test_sign").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	// 验证 Body 被设置
	if req.apiReq.Body == nil {
		t.Fatal("Body should not be nil")
	}
	body, ok := req.apiReq.Body.(*UpdateMemberRequest)
	if !ok {
		t.Fatal("Body should be *UpdateMemberRequest")
	}
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	if body.ProjectId == nil || *body.ProjectId != "1125904357323169" {
		t.Errorf("Body.ProjectId = %v, want 1125904357323169", body.ProjectId)
	}
}

func TestUpdateMemberApiReqBuilder_PartialParams(t *testing.T) {
	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*UpdateMemberRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 未设置的字段应为 nil
	if body.AccessToken != nil {
		t.Errorf("Body.AccessToken = %v, want nil", body.AccessToken)
	}
}

func TestUpdateMemberApiReqBuilder_OnlyCommonParams(t *testing.T) {
	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	body := req.apiReq.Body.(*UpdateMemberRequest)
	if body.ClientId == nil || *body.ClientId != "test_client" {
		t.Errorf("Body.ClientId = %v, want test_client", body.ClientId)
	}
	// 业务参数不应存在
	if body.ProjectId != nil {
		t.Errorf("Body.ProjectId = %v, want nil", body.ProjectId)
	}
	if body.MemberIds != nil {
		t.Errorf("Body.MemberIds = %v, want nil", body.MemberIds)
	}
	if body.MemberType != nil {
		t.Errorf("Body.MemberType = %v, want nil", body.MemberType)
	}
}

func TestUpdateMemberApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"success_data": ["emp001", "emp002"],
			"error_data": [
				{
					"error_msg": "员工不在该企业中",
					"error_member_ids": ["emp003"],
					"error_member_values": ["E003"]
				}
			],
			"success_member_values": ["E001", "E002"]
		},
		"request_id": "test_request_id"
	}`

	var reply UpdateMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %q, want test_request_id", reply.RequestId)
	}
	if len(reply.Data.SuccessData) != 2 {
		t.Fatalf("SuccessData len = %d, want 2", len(reply.Data.SuccessData))
	}
	if reply.Data.SuccessData[0] != "emp001" {
		t.Errorf("SuccessData[0] = %q, want emp001", reply.Data.SuccessData[0])
	}
	if len(reply.Data.ErrorData) != 1 {
		t.Fatalf("ErrorData len = %d, want 1", len(reply.Data.ErrorData))
	}
	if reply.Data.ErrorData[0].ErrorMsg == nil || *reply.Data.ErrorData[0].ErrorMsg != "员工不在该企业中" {
		t.Errorf("ErrorData[0].ErrorMsg = %v, want 员工不在该企业中", reply.Data.ErrorData[0].ErrorMsg)
	}
	if len(reply.Data.ErrorData[0].ErrorMemberIds) != 1 || reply.Data.ErrorData[0].ErrorMemberIds[0] != "emp003" {
		t.Errorf("ErrorData[0].ErrorMemberIds = %v, want [emp003]", reply.Data.ErrorData[0].ErrorMemberIds)
	}
	if len(reply.Data.SuccessMemberValues) != 2 {
		t.Errorf("SuccessMemberValues len = %d, want 2", len(reply.Data.SuccessMemberValues))
	}
}

func TestUpdateMemberApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 70000,
		"errmsg": "project 非项目 id",
		"request_id": "test_request_id"
	}`

	var reply UpdateMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 70000 {
		t.Errorf("Errno = %d, want 70000", reply.Errno)
	}
	if reply.Errmsg != "project 非项目 id" {
		t.Errorf("Errmsg = %q, want project 非项目 id", reply.Errmsg)
	}
}

func TestUpdateMemberApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"success_data": ["emp001", "emp002", "emp003"],
			"error_data": [
				{
					"error_msg": "员工不在该企业中",
					"error_member_ids": ["emp004", "emp005"]
				},
				{
					"error_msg": "该员工已经在这个项目中",
					"error_member_ids": ["emp006"]
				}
			]
		},
		"request_id": "req_multi"
	}`

	var reply UpdateMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.SuccessData) != 3 {
		t.Errorf("SuccessData len = %d, want 3", len(reply.Data.SuccessData))
	}
	if len(reply.Data.ErrorData) != 2 {
		t.Fatalf("ErrorData len = %d, want 2", len(reply.Data.ErrorData))
	}
	if reply.Data.ErrorData[1].ErrorMsg == nil || *reply.Data.ErrorData[1].ErrorMsg != "该员工已经在这个项目中" {
		t.Errorf("ErrorData[1].ErrorMsg = %v, want 该员工已经在这个项目中", reply.Data.ErrorData[1].ErrorMsg)
	}
}

func TestUpdateMemberApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"success_data":[],"error_data":[]},"request_id":"req_empty"}`

	var reply UpdateMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.SuccessData) != 0 {
		t.Errorf("SuccessData len = %d, want 0", len(reply.Data.SuccessData))
	}
	if len(reply.Data.ErrorData) != 0 {
		t.Errorf("ErrorData len = %d, want 0", len(reply.Data.ErrorData))
	}
}

func TestUpdateMemberApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":70000,"errmsg":"project 非项目 id","request_id":"req_nodata"}`

	var reply UpdateMemberApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 70000 {
		t.Errorf("Errno = %d, want 70000", reply.Errno)
	}
}

// --- UpdateMember 资源方法测试 ---

func TestUpdateMember_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/Project/updateMember" {
			t.Errorf("expected path /river/Project/updateMember, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"success_data": ["emp001", "emp002"],
				"error_data": [
					{
						"error_msg": "员工不在该企业中",
						"error_member_ids": ["emp003"]
					}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		ProjectId("1125904357323169").
		MemberIds("emp001,emp002,emp003").
		Sign("test_sign").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if resp.UpdateMemberApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateMemberApiReply.Errno)
	}
	if len(resp.UpdateMemberApiReply.Data.SuccessData) != 2 {
		t.Fatalf("SuccessData len = %d, want 2", len(resp.UpdateMemberApiReply.Data.SuccessData))
	}
	if resp.UpdateMemberApiReply.Data.SuccessData[0] != "emp001" {
		t.Errorf("SuccessData[0] = %q, want emp001", resp.UpdateMemberApiReply.Data.SuccessData[0])
	}
	if len(resp.UpdateMemberApiReply.Data.ErrorData) != 1 {
		t.Fatalf("ErrorData len = %d, want 1", len(resp.UpdateMemberApiReply.Data.ErrorData))
	}
	if resp.UpdateMemberApiReply.Data.ErrorData[0].ErrorMsg == nil || *resp.UpdateMemberApiReply.Data.ErrorData[0].ErrorMsg != "员工不在该企业中" {
		t.Errorf("ErrorData[0].ErrorMsg = %v, want 员工不在该企业中", resp.UpdateMemberApiReply.Data.ErrorData[0].ErrorMsg)
	}
}

func TestUpdateMember_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"success_data":[],"error_data":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if len(resp.UpdateMemberApiReply.Data.SuccessData) != 0 {
		t.Errorf("SuccessData len = %d, want 0", len(resp.UpdateMemberApiReply.Data.SuccessData))
	}
}

func TestUpdateMember_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":70000,"errmsg":"project 非项目 id","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if resp.UpdateMemberApiReply.Errno != 70000 {
		t.Errorf("Errno = %d, want 70000", resp.UpdateMemberApiReply.Errno)
	}
}

func TestUpdateMember_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.UpdateMemberApiReply != nil {
		t.Errorf("UpdateMemberApiReply should be nil for non-200 response")
	}
}

func TestUpdateMember_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"success_data":["emp001"],"error_data":[]},"request_id":"req_enc"}`
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

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if resp.UpdateMemberApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateMemberApiReply.Errno)
	}
	if len(resp.UpdateMemberApiReply.Data.SuccessData) != 1 {
		t.Fatalf("SuccessData len = %d, want 1", len(resp.UpdateMemberApiReply.Data.SuccessData))
	}
	if resp.UpdateMemberApiReply.Data.SuccessData[0] != "emp001" {
		t.Errorf("SuccessData[0] = %q, want emp001", resp.UpdateMemberApiReply.Data.SuccessData[0])
	}
}

func TestUpdateMember_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"success_data":["emp001"],"error_data":[]},"request_id":"req_enc256"}`
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

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		ProjectId("1125904357323169").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if len(resp.UpdateMemberApiReply.Data.SuccessData) != 1 {
		t.Fatalf("SuccessData len = %d, want 1", len(resp.UpdateMemberApiReply.Data.SuccessData))
	}
}

func TestUpdateMember_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"success_data":["emp001"],"error_data":[]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	p := &project{option: option}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
	if len(resp.UpdateMemberApiReply.Data.SuccessData) != 1 {
		t.Fatalf("SuccessData len = %d, want 1", len(resp.UpdateMemberApiReply.Data.SuccessData))
	}
}

func TestUpdateMember_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"success_data":[],"error_data":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newTestOption(testServer.URL)
	p := &project{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	requestBody := NewUpdateMemberRequestBuilder().
		ClientId("test_client").
		Build()

	req := NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(requestBody).
		Build()

	resp, err := p.UpdateMember(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateMember() error = %v", err)
	}
	if resp.UpdateMemberApiReply == nil {
		t.Fatal("UpdateMemberApiReply is nil")
	}
}
