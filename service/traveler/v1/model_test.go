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

// --- 数据模型 Builder 测试 ---

func TestTravelerInfoBuilder(t *testing.T) {
	card := NewTravelCardInfoBuilder().
		CardType(1).
		CardNo("110101199001011234").
		ExpireDate("2050-01-01").
		Build()

	info := NewTravelerInfoBuilder().
		TravelerId("1125922289295589").
		Phone("13800000001").
		Name("张三").
		EnglishSurname("Zhang").
		EnglishName("San").
		Remark("VIP客户").
		Sex(1).
		OutTravelerId("out_001").
		BirthDate("1990-01-01").
		CardList([]TravelCardInfo{*card}).
		ForceClearFields([]string{"card_list"}).
		Build()

	if info.TravelerId == nil || *info.TravelerId != "1125922289295589" {
		t.Errorf("TravelerId = %v, want 1125922289295589", info.TravelerId)
	}
	if info.Phone == nil || *info.Phone != "13800000001" {
		t.Errorf("Phone = %v, want 13800000001", info.Phone)
	}
	if info.Name == nil || *info.Name != "张三" {
		t.Errorf("Name = %v, want 张三", info.Name)
	}
	if info.EnglishSurname == nil || *info.EnglishSurname != "Zhang" {
		t.Errorf("EnglishSurname = %v, want Zhang", info.EnglishSurname)
	}
	if info.EnglishName == nil || *info.EnglishName != "San" {
		t.Errorf("EnglishName = %v, want San", info.EnglishName)
	}
	if info.Remark == nil || *info.Remark != "VIP客户" {
		t.Errorf("Remark = %v, want VIP客户", info.Remark)
	}
	if info.Sex == nil || *info.Sex != 1 {
		t.Errorf("Sex = %v, want 1", info.Sex)
	}
	if info.OutTravelerId == nil || *info.OutTravelerId != "out_001" {
		t.Errorf("OutTravelerId = %v, want out_001", info.OutTravelerId)
	}
	if info.BirthDate == nil || *info.BirthDate != "1990-01-01" {
		t.Errorf("BirthDate = %v, want 1990-01-01", info.BirthDate)
	}
	if len(info.CardList) != 1 {
		t.Fatalf("CardList len = %d, want 1", len(info.CardList))
	}
	if info.CardList[0].CardType == nil || *info.CardList[0].CardType != 1 {
		t.Errorf("CardList[0].CardType = %v, want 1", info.CardList[0].CardType)
	}
	if info.CardList[0].CardNo == nil || *info.CardList[0].CardNo != "110101199001011234" {
		t.Errorf("CardList[0].CardNo = %v, want 110101199001011234", info.CardList[0].CardNo)
	}
	if info.CardList[0].ExpireDate == nil || *info.CardList[0].ExpireDate != "2050-01-01" {
		t.Errorf("CardList[0].ExpireDate = %v, want 2050-01-01", info.CardList[0].ExpireDate)
	}
	if len(info.ForceClearFields) != 1 || info.ForceClearFields[0] != "card_list" {
		t.Errorf("ForceClearFields = %v, want [card_list]", info.ForceClearFields)
	}

	// 部分设置 + int32 零值（Sex=0 仍应写入）
	info2 := NewTravelerInfoBuilder().
		Name("李四").
		Sex(0).
		Build()

	if info2.Name == nil || *info2.Name != "李四" {
		t.Errorf("Name = %v, want 李四", info2.Name)
	}
	if info2.Sex == nil || *info2.Sex != 0 {
		t.Errorf("Sex = %v, want 0", info2.Sex)
	}
	if info2.TravelerId != nil {
		t.Errorf("TravelerId = %v, want nil", info2.TravelerId)
	}
	if info2.Phone != nil {
		t.Errorf("Phone = %v, want nil", info2.Phone)
	}
	if info2.CardList != nil {
		t.Errorf("CardList = %v, want nil", info2.CardList)
	}
	if info2.ForceClearFields != nil {
		t.Errorf("ForceClearFields = %v, want nil", info2.ForceClearFields)
	}
}

func TestTravelCardInfoBuilder(t *testing.T) {
	card := NewTravelCardInfoBuilder().
		CardType(2).
		CardNo("EE12345678CN").
		ExpireDate("2030-12-31").
		Build()

	if card.CardType == nil || *card.CardType != 2 {
		t.Errorf("CardType = %v, want 2", card.CardType)
	}
	if card.CardNo == nil || *card.CardNo != "EE12345678CN" {
		t.Errorf("CardNo = %v, want EE12345678CN", card.CardNo)
	}
	if card.ExpireDate == nil || *card.ExpireDate != "2030-12-31" {
		t.Errorf("ExpireDate = %v, want 2030-12-31", card.ExpireDate)
	}

	// 部分设置 + int32 零值
	card2 := NewTravelCardInfoBuilder().
		CardType(0).
		Build()

	if card2.CardType == nil || *card2.CardType != 0 {
		t.Errorf("CardType = %v, want 0", card2.CardType)
	}
	if card2.CardNo != nil {
		t.Errorf("CardNo = %v, want nil", card2.CardNo)
	}
	if card2.ExpireDate != nil {
		t.Errorf("ExpireDate = %v, want nil", card2.ExpireDate)
	}
}

// --- CreateTraveler 模型层测试 ---

func TestCreateTravelerRequestBuilder_FullParams(t *testing.T) {
	travelerInfo := NewTravelerInfoBuilder().
		Name("张三").
		Phone("13800000001").
		Build()

	request := NewCreateTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson("{\"name\":\"张三\",\"phone\":\"13800000001\"}").
		ParamJsonObj(*travelerInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"name\":\"张三\",\"phone\":\"13800000001\"}" {
		t.Errorf("ParamJson = %v, want json string", request.ParamJson)
	}
	if request.ParamJsonObj == nil {
		t.Fatal("ParamJsonObj is nil")
	}
	if request.ParamJsonObj.Name == nil || *request.ParamJsonObj.Name != "张三" {
		t.Errorf("ParamJsonObj.Name = %v, want 张三", request.ParamJsonObj.Name)
	}
}

func TestCreateTravelerRequestBuilder_PartialParams(t *testing.T) {
	request := NewCreateTravelerRequestBuilder().
		ClientId("test_client").
		ParamJson("{\"name\":\"张三\"}").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"name\":\"张三\"}" {
		t.Errorf("ParamJson = %v, want {\"name\":\"张三\"}", request.ParamJson)
	}
	// 未设置的参数应为 nil
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestCreateTravelerRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewCreateTravelerRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	// int64 零值也应该被设置到 Request 字段中
	if request.Timestamp == nil {
		t.Fatal("Timestamp is nil, want 0")
	}
	if *request.Timestamp != 0 {
		t.Errorf("Timestamp = %d, want 0", *request.Timestamp)
	}
}

func TestCreateTravelerRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewCreateTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	// 业务参数不应存在
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestCreateTravelerApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": 1125922289295589
		},
		"request_id": "test_request_id"
	}`

	var reply CreateTravelerApiReply
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
	if reply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", reply.Data.Id)
	}
}

func TestCreateTravelerApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply CreateTravelerApiReply
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

func TestCreateTravelerApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`

	var reply CreateTravelerApiReply
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

func TestCreateTravelerApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreateTravelerApiReply
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

// --- DelTraveler 模型层测试 ---

func TestDelTravelerRequestBuilder_FullParams(t *testing.T) {
	travelerInfo := NewTravelerInfoBuilder().
		TravelerId("1125922289295589").
		Build()

	request := NewDelTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson("{\"traveler_id\":\"1125922289295589\"}").
		ParamJsonObj(*travelerInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"traveler_id\":\"1125922289295589\"}" {
		t.Errorf("ParamJson = %v, want json string", request.ParamJson)
	}
	if request.ParamJsonObj == nil {
		t.Fatal("ParamJsonObj is nil")
	}
	if request.ParamJsonObj.TravelerId == nil || *request.ParamJsonObj.TravelerId != "1125922289295589" {
		t.Errorf("ParamJsonObj.TravelerId = %v, want 1125922289295589", request.ParamJsonObj.TravelerId)
	}
}

func TestDelTravelerRequestBuilder_PartialParams(t *testing.T) {
	request := NewDelTravelerRequestBuilder().
		ClientId("test_client").
		ParamJson("{\"traveler_id\":\"1125922289295589\"}").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"traveler_id\":\"1125922289295589\"}" {
		t.Errorf("ParamJson = %v, want json string", request.ParamJson)
	}
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestDelTravelerRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewDelTravelerRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if request.Timestamp == nil {
		t.Fatal("Timestamp is nil, want 0")
	}
	if *request.Timestamp != 0 {
		t.Errorf("Timestamp = %d, want 0", *request.Timestamp)
	}
}

func TestDelTravelerRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewDelTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestDelTravelerApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": 1125922289295589
		},
		"request_id": "test_request_id"
	}`

	var reply DelTravelerApiReply
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
	if reply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", reply.Data.Id)
	}
}

func TestDelTravelerApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply DelTravelerApiReply
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

func TestDelTravelerApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`

	var reply DelTravelerApiReply
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

func TestDelTravelerApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply DelTravelerApiReply
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

// --- UpdateTraveler 模型层测试 ---

func TestUpdateTravelerRequestBuilder_FullParams(t *testing.T) {
	travelerInfo := NewTravelerInfoBuilder().
		TravelerId("1125922289295589").
		Name("李四").
		Phone("13800000002").
		Build()

	request := NewUpdateTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson("{\"traveler_id\":\"1125922289295589\",\"name\":\"李四\"}").
		ParamJsonObj(*travelerInfo).
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"traveler_id\":\"1125922289295589\",\"name\":\"李四\"}" {
		t.Errorf("ParamJson = %v, want json string", request.ParamJson)
	}
	if request.ParamJsonObj == nil {
		t.Fatal("ParamJsonObj is nil")
	}
	if request.ParamJsonObj.TravelerId == nil || *request.ParamJsonObj.TravelerId != "1125922289295589" {
		t.Errorf("ParamJsonObj.TravelerId = %v, want 1125922289295589", request.ParamJsonObj.TravelerId)
	}
	if request.ParamJsonObj.Name == nil || *request.ParamJsonObj.Name != "李四" {
		t.Errorf("ParamJsonObj.Name = %v, want 李四", request.ParamJsonObj.Name)
	}
}

func TestUpdateTravelerRequestBuilder_PartialParams(t *testing.T) {
	request := NewUpdateTravelerRequestBuilder().
		ClientId("test_client").
		ParamJson("{\"traveler_id\":\"1125922289295589\"}").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.ParamJson == nil || *request.ParamJson != "{\"traveler_id\":\"1125922289295589\"}" {
		t.Errorf("ParamJson = %v, want json string", request.ParamJson)
	}
	if request.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", request.AccessToken)
	}
	if request.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", request.CompanyId)
	}
	if request.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", request.Timestamp)
	}
	if request.Sign != nil {
		t.Errorf("Sign = %v, want nil", request.Sign)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestUpdateTravelerRequestBuilder_ZeroIntValues(t *testing.T) {
	request := NewUpdateTravelerRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if request.Timestamp == nil {
		t.Fatal("Timestamp is nil, want 0")
	}
	if *request.Timestamp != 0 {
		t.Errorf("Timestamp = %d, want 0", *request.Timestamp)
	}
}

func TestUpdateTravelerRequestBuilder_OnlyCommonParams(t *testing.T) {
	request := NewUpdateTravelerRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if request.ClientId == nil || *request.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", request.ClientId)
	}
	if request.AccessToken == nil || *request.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", request.AccessToken)
	}
	if request.CompanyId == nil || *request.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", request.CompanyId)
	}
	if request.Timestamp == nil || *request.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", request.Timestamp)
	}
	if request.Sign == nil || *request.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", request.Sign)
	}
	if request.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", request.ParamJson)
	}
	if request.ParamJsonObj != nil {
		t.Errorf("ParamJsonObj = %v, want nil", request.ParamJsonObj)
	}
}

func TestUpdateTravelerApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"id": 1125922289295589
		},
		"request_id": "test_request_id"
	}`

	var reply UpdateTravelerApiReply
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
	if reply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", reply.Data.Id)
	}
}

func TestUpdateTravelerApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{
		"errno": 10003,
		"errmsg": "param error",
		"request_id": "test_request_id"
	}`

	var reply UpdateTravelerApiReply
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

func TestUpdateTravelerApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`

	var reply UpdateTravelerApiReply
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

func TestUpdateTravelerApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply UpdateTravelerApiReply
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

// --- 资源方法测试 ---

func newTravelerTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// --- CreateTraveler 资源方法测试 ---

func TestCreateTraveler_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/traveler/create" {
			t.Errorf("expected path /open-apis/v1/traveler/create, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Errno == nil || *resp.CreateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateTravelerApiReply.Errno)
	}
	if resp.CreateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.CreateTravelerApiReply.Data.Id)
	}
}

func TestCreateTraveler_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Errno == nil || *resp.CreateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateTravelerApiReply.Errno)
	}
	if resp.CreateTravelerApiReply.Data != nil {
		t.Errorf("Data = %v, want nil", resp.CreateTravelerApiReply.Data)
	}
}

func TestCreateTraveler_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Errno == nil || *resp.CreateTravelerApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.CreateTravelerApiReply.Errno)
	}
}

func TestCreateTraveler_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	// 非 200 响应不应设置 ApiReply
	if resp.CreateTravelerApiReply != nil {
		t.Errorf("CreateTravelerApiReply should be nil for non-200 response")
	}
}

func TestCreateTraveler_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Errno == nil || *resp.CreateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateTravelerApiReply.Errno)
	}
	if resp.CreateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.CreateTravelerApiReply.Data.Id)
	}
}

func TestCreateTraveler_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc256"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.CreateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.CreateTravelerApiReply.Data.Id)
	}
}

func TestCreateTraveler_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &traveler{option: option}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
	if resp.CreateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestCreateTraveler_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(NewCreateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.CreateTraveler(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateTraveler() error = %v", err)
	}
	if resp.CreateTravelerApiReply == nil {
		t.Fatal("CreateTravelerApiReply is nil")
	}
}

// --- DelTraveler 资源方法测试 ---

func TestDelTraveler_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/traveler/del" {
			t.Errorf("expected path /open-apis/v1/traveler/del, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Errno == nil || *resp.DelTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelTravelerApiReply.Errno)
	}
	if resp.DelTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.DelTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.DelTravelerApiReply.Data.Id)
	}
}

func TestDelTraveler_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Errno == nil || *resp.DelTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelTravelerApiReply.Errno)
	}
	if resp.DelTravelerApiReply.Data != nil {
		t.Errorf("Data = %v, want nil", resp.DelTravelerApiReply.Data)
	}
}

func TestDelTraveler_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Errno == nil || *resp.DelTravelerApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.DelTravelerApiReply.Errno)
	}
}

func TestDelTraveler_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.DelTravelerApiReply != nil {
		t.Errorf("DelTravelerApiReply should be nil for non-200 response")
	}
}

func TestDelTraveler_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Errno == nil || *resp.DelTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.DelTravelerApiReply.Errno)
	}
	if resp.DelTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.DelTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.DelTravelerApiReply.Data.Id)
	}
}

func TestDelTraveler_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc256"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.DelTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.DelTravelerApiReply.Data.Id)
	}
}

func TestDelTraveler_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &traveler{option: option}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
	if resp.DelTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestDelTraveler_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewDelTravelerApiReqBuilder().
		DelTravelerRequest(NewDelTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.DelTraveler(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("DelTraveler() error = %v", err)
	}
	if resp.DelTravelerApiReply == nil {
		t.Fatal("DelTravelerApiReply is nil")
	}
}

// --- UpdateTraveler 资源方法测试 ---

func TestUpdateTraveler_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/traveler/update" {
			t.Errorf("expected path /open-apis/v1/traveler/update, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Errno == nil || *resp.UpdateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateTravelerApiReply.Errno)
	}
	if resp.UpdateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.UpdateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.UpdateTravelerApiReply.Data.Id)
	}
}

func TestUpdateTraveler_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Errno == nil || *resp.UpdateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateTravelerApiReply.Errno)
	}
	if resp.UpdateTravelerApiReply.Data != nil {
		t.Errorf("Data = %v, want nil", resp.UpdateTravelerApiReply.Data)
	}
}

func TestUpdateTraveler_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Errno == nil || *resp.UpdateTravelerApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.UpdateTravelerApiReply.Errno)
	}
}

func TestUpdateTraveler_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateTravelerApiReply != nil {
		t.Errorf("UpdateTravelerApiReply should be nil for non-200 response")
	}
}

func TestUpdateTraveler_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Errno == nil || *resp.UpdateTravelerApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.UpdateTravelerApiReply.Errno)
	}
	if resp.UpdateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.UpdateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.UpdateTravelerApiReply.Data.Id)
	}
}

func TestUpdateTraveler_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_enc256"}`
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

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.UpdateTravelerApiReply.Data.Id != 1125922289295589 {
		t.Errorf("Id = %d, want 1125922289295589", resp.UpdateTravelerApiReply.Data.Id)
	}
}

func TestUpdateTraveler_EncryptionNoEncryptData(t *testing.T) {
	// 启用加密但响应无 encrypt_data 字段，应直接反序列化
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &traveler{option: option}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
	if resp.UpdateTravelerApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
}

func TestUpdateTraveler_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"id":1125922289295589},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newTravelerTestOption(testServer.URL)
	svc := &traveler{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(NewUpdateTravelerRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := svc.UpdateTraveler(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateTraveler() error = %v", err)
	}
	if resp.UpdateTravelerApiReply == nil {
		t.Fatal("UpdateTravelerApiReply is nil")
	}
}
