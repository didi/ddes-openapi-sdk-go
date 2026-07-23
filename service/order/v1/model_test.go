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

// ============================================================
// 通用 helper
// ============================================================

func newOrderTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// encryptResp 构造加密响应体
func encryptResp(plaintext string, key []byte, useURL bool) string {
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		panic(err)
	}
	var enc string
	if useURL {
		enc = base64.URLEncoding.EncodeToString(encrypted)
	} else {
		enc = base64.StdEncoding.EncodeToString(encrypted)
	}
	return `{"encrypt_data":"` + enc + `"}`
}

// ============================================================
// 数据模型 Builder 测试
// ============================================================

func TestPriceInfoBuilder(t *testing.T) {
	info := NewPriceInfoBuilder().
		Name("车费").
		Amount("25.50").
		Type("1").
		Build()
	if info.Name == nil || *info.Name != "车费" {
		t.Errorf("Name = %v, want 车费", info.Name)
	}
	if info.Amount == nil || *info.Amount != "25.50" {
		t.Errorf("Amount = %v, want 25.50", info.Amount)
	}
	if info.Type == nil || *info.Type != "1" {
		t.Errorf("Type = %v, want 1", info.Type)
	}
	// 部分设置
	info2 := NewPriceInfoBuilder().Name("夜间费").Build()
	if info2.Name == nil || *info2.Name != "夜间费" {
		t.Errorf("Name = %v, want 夜间费", info2.Name)
	}
	if info2.Amount != nil {
		t.Errorf("Amount = %v, want nil", info2.Amount)
	}
}

func TestStopoverPointBuilder(t *testing.T) {
	info := NewStopoverPointBuilder().
		Cityid("1").
		Address("朝阳区望京SOHO").
		StopId(1).
		City("北京").
		Lng("116.481").
		Lat("39.996").
		Name("望京SOHO").
		Status(0).
		Build()
	if info.Cityid == nil || *info.Cityid != "1" {
		t.Errorf("Cityid = %v, want 1", info.Cityid)
	}
	if info.StopId == nil || *info.StopId != 1 {
		t.Errorf("StopId = %v, want 1", info.StopId)
	}
	if info.Name == nil || *info.Name != "望京SOHO" {
		t.Errorf("Name = %v, want 望京SOHO", info.Name)
	}
	// 部分设置
	info2 := NewStopoverPointBuilder().City("上海").Build()
	if info2.City == nil || *info2.City != "上海" {
		t.Errorf("City = %v, want 上海", info2.City)
	}
	if info2.StopId != nil {
		t.Errorf("StopId = %v, want nil", info2.StopId)
	}
}

func TestBudgetCenterListItemBuilder(t *testing.T) {
	info := NewBudgetCenterListItemBuilder().
		AppName("成本中心1").
		Sequence(1).
		Id("1125922289295589").
		Value("市场部").
		Code("BC001").
		Build()
	if info.AppName == nil || *info.AppName != "成本中心1" {
		t.Errorf("AppName = %v, want 成本中心1", info.AppName)
	}
	if info.Sequence == nil || *info.Sequence != 1 {
		t.Errorf("Sequence = %v, want 1", info.Sequence)
	}
	if info.Id == nil || *info.Id != "1125922289295589" {
		t.Errorf("Id = %v, want 1125922289295589", info.Id)
	}
	if info.Value == nil || *info.Value != "市场部" {
		t.Errorf("Value = %v, want 市场部", info.Value)
	}
	if info.Code == nil || *info.Code != "BC001" {
		t.Errorf("Code = %v, want BC001", info.Code)
	}
	// 部分设置
	info2 := NewBudgetCenterListItemBuilder().Id("1001").Build()
	if info2.Id == nil || *info2.Id != "1001" {
		t.Errorf("Id = %v, want 1001", info2.Id)
	}
	if info2.AppName != nil {
		t.Errorf("AppName = %v, want nil", info2.AppName)
	}
}

func TestHistoryDestinationBuilder(t *testing.T) {
	info := NewHistoryDestinationBuilder().
		Address("北京南站").
		Lng("116.378").
		Lat("39.865").
		Name("北京南站").
		Sequence(1).
		Build()
	if info.Address == nil || *info.Address != "北京南站" {
		t.Errorf("Address = %v, want 北京南站", info.Address)
	}
	if info.Sequence == nil || *info.Sequence != 1 {
		t.Errorf("Sequence = %v, want 1", info.Sequence)
	}
	// 部分设置
	info2 := NewHistoryDestinationBuilder().Name("上海虹桥").Build()
	if info2.Name == nil || *info2.Name != "上海虹桥" {
		t.Errorf("Name = %v, want 上海虹桥", info2.Name)
	}
	if info2.Address != nil {
		t.Errorf("Address = %v, want nil", info2.Address)
	}
}

func TestLegalEntityInfoBuilder(t *testing.T) {
	info := NewLegalEntityInfoBuilder().
		LegalEntityId("1125922289295589").
		OutLegalEntityId("OUT001").
		LegalEntityName("滴滴出行").
		Build()
	if info.LegalEntityId == nil || *info.LegalEntityId != "1125922289295589" {
		t.Errorf("LegalEntityId = %v, want 1125922289295589", info.LegalEntityId)
	}
	if info.OutLegalEntityId == nil || *info.OutLegalEntityId != "OUT001" {
		t.Errorf("OutLegalEntityId = %v, want OUT001", info.OutLegalEntityId)
	}
	if info.LegalEntityName == nil || *info.LegalEntityName != "滴滴出行" {
		t.Errorf("LegalEntityName = %v, want 滴滴出行", info.LegalEntityName)
	}
	// 部分设置
	info2 := NewLegalEntityInfoBuilder().LegalEntityId("1001").Build()
	if info2.LegalEntityId == nil || *info2.LegalEntityId != "1001" {
		t.Errorf("LegalEntityId = %v, want 1001", info2.LegalEntityId)
	}
	if info2.LegalEntityName != nil {
		t.Errorf("LegalEntityName = %v, want nil", info2.LegalEntityName)
	}
}

func TestOrderPassengerInfoBuilder(t *testing.T) {
	cs := NewCostShareBuilder().
		ShareRatio(1.0).
		CompanyPay("25.50").
		Build()
	info := NewOrderPassengerInfoBuilder().
		PassengerName("张三").
		PassengerPhone("13800000001").
		PassengerTravelerId("1125922289295589").
		PassengerEmployeeId("D0001").
		OutTravelerId("out_001").
		IsTraveler(0).
		CostShare(*cs).
		Build()
	if info.PassengerName == nil || *info.PassengerName != "张三" {
		t.Errorf("PassengerName = %v, want 张三", info.PassengerName)
	}
	if info.PassengerTravelerId == nil || *info.PassengerTravelerId != "1125922289295589" {
		t.Errorf("PassengerTravelerId = %v, want 1125922289295589", info.PassengerTravelerId)
	}
	if info.IsTraveler == nil || *info.IsTraveler != 0 {
		t.Errorf("IsTraveler = %v, want 0", info.IsTraveler)
	}
	if info.CostShare == nil {
		t.Fatal("CostShare is nil")
	}
	if info.CostShare.ShareRatio == nil || *info.CostShare.ShareRatio != 1.0 {
		t.Errorf("CostShare.ShareRatio = %v, want 1.0", info.CostShare.ShareRatio)
	}
	// 部分设置
	info2 := NewOrderPassengerInfoBuilder().PassengerName("李四").Build()
	if info2.PassengerName == nil || *info2.PassengerName != "李四" {
		t.Errorf("PassengerName = %v, want 李四", info2.PassengerName)
	}
	if info2.CostShare != nil {
		t.Errorf("CostShare = %v, want nil", info2.CostShare)
	}
}

func TestCostShareBuilder(t *testing.T) {
	info := NewCostShareBuilder().
		ShareRatio(0.8).
		ShareCompanyRatio(0.6).
		SharePersonalRatio(0.4).
		CompanyPay("100.00").
		PersonalPay("40.00").
		CompanyRealPay("100.00").
		PersonalRealPay("40.00").
		CompanyRealRefund("0").
		PersonalRealRefund("0").
		Build()
	if info.ShareRatio == nil || *info.ShareRatio != 0.8 {
		t.Errorf("ShareRatio = %v, want 0.8", info.ShareRatio)
	}
	if info.CompanyPay == nil || *info.CompanyPay != "100.00" {
		t.Errorf("CompanyPay = %v, want 100.00", info.CompanyPay)
	}
	if info.PersonalPay == nil || *info.PersonalPay != "40.00" {
		t.Errorf("PersonalPay = %v, want 40.00", info.PersonalPay)
	}
	// 部分设置
	info2 := NewCostShareBuilder().CompanyPay("50.00").Build()
	if info2.CompanyPay == nil || *info2.CompanyPay != "50.00" {
		t.Errorf("CompanyPay = %v, want 50.00", info2.CompanyPay)
	}
	if info2.ShareRatio != nil {
		t.Errorf("ShareRatio = %v, want nil", info2.ShareRatio)
	}
}

func TestRcListItemBuilder(t *testing.T) {
	rc := NewRcInfoBuilder().
		RcType("car_type").
		RcCode("OVER_LIMIT").
		RcReason("超出车型限制").
		RcRemark("主管审批通过").
		Build()
	info := NewRcListItemBuilder().RcInfo(*rc).Build()
	if info.RcInfo == nil {
		t.Fatal("RcInfo is nil")
	}
	if info.RcInfo.RcType == nil || *info.RcInfo.RcType != "car_type" {
		t.Errorf("RcType = %v, want car_type", info.RcInfo.RcType)
	}
	if info.RcInfo.RcReason == nil || *info.RcInfo.RcReason != "超出车型限制" {
		t.Errorf("RcReason = %v, want 超出车型限制", info.RcInfo.RcReason)
	}
	// 部分设置
	info2 := NewRcListItemBuilder().Build()
	if info2.RcInfo != nil {
		t.Errorf("RcInfo = %v, want nil", info2.RcInfo)
	}
}

func TestRcInfoBuilder(t *testing.T) {
	info := NewRcInfoBuilder().
		RcType("budget").
		RcCode("OVER_BUDGET").
		RcReason("超出预算").
		RcRemark("业务需要").
		Build()
	if info.RcType == nil || *info.RcType != "budget" {
		t.Errorf("RcType = %v, want budget", info.RcType)
	}
	if info.RcCode == nil || *info.RcCode != "OVER_BUDGET" {
		t.Errorf("RcCode = %v, want OVER_BUDGET", info.RcCode)
	}
	if info.RcRemark == nil || *info.RcRemark != "业务需要" {
		t.Errorf("RcRemark = %v, want 业务需要", info.RcRemark)
	}
	// 部分设置
	info2 := NewRcInfoBuilder().RcCode("TEST").Build()
	if info2.RcCode == nil || *info2.RcCode != "TEST" {
		t.Errorf("RcCode = %v, want TEST", info2.RcCode)
	}
	if info2.RcType != nil {
		t.Errorf("RcType = %v, want nil", info2.RcType)
	}
}

func TestPageInfoBuilder(t *testing.T) {
	info := NewPageInfoBuilder().
		CurPage(1).
		Limit(20).
		Total(100).
		Build()
	if info.CurPage == nil || *info.CurPage != 1 {
		t.Errorf("CurPage = %v, want 1", info.CurPage)
	}
	if info.Limit == nil || *info.Limit != 20 {
		t.Errorf("Limit = %v, want 20", info.Limit)
	}
	if info.Total == nil || *info.Total != 100 {
		t.Errorf("Total = %v, want 100", info.Total)
	}
	// 部分设置
	info2 := NewPageInfoBuilder().Total(50).Build()
	if info2.Total == nil || *info2.Total != 50 {
		t.Errorf("Total = %v, want 50", info2.Total)
	}
	if info2.CurPage != nil {
		t.Errorf("CurPage = %v, want nil", info2.CurPage)
	}
}

func TestAirlineInfoBuilder(t *testing.T) {
	info := NewAirlineInfoBuilder().
		AirlineName("中国东方航空股份有限公司").
		AirlineSimpleName("东方航空").
		AirlineVerySimpleName("东航").
		FlightNumber("MU5100").
		Build()
	if info.AirlineName == nil || *info.AirlineName != "中国东方航空股份有限公司" {
		t.Errorf("AirlineName = %v, want 中国东方航空股份有限公司", info.AirlineName)
	}
	if info.FlightNumber == nil || *info.FlightNumber != "MU5100" {
		t.Errorf("FlightNumber = %v, want MU5100", info.FlightNumber)
	}
	// 部分设置
	info2 := NewAirlineInfoBuilder().FlightNumber("CA1815").Build()
	if info2.FlightNumber == nil || *info2.FlightNumber != "CA1815" {
		t.Errorf("FlightNumber = %v, want CA1815", info2.FlightNumber)
	}
	if info2.AirlineName != nil {
		t.Errorf("AirlineName = %v, want nil", info2.AirlineName)
	}
}

func TestFlightRoutePriceBuilder(t *testing.T) {
	info := NewFlightRoutePriceBuilder().
		FlightNumber("MU5100").
		EstimateFirst(150000).
		EstimateBiz(100000).
		EstimateEco(50000).
		Build()
	if info.FlightNumber == nil || *info.FlightNumber != "MU5100" {
		t.Errorf("FlightNumber = %v, want MU5100", info.FlightNumber)
	}
	if info.EstimateFirst == nil || *info.EstimateFirst != 150000 {
		t.Errorf("EstimateFirst = %v, want 150000", info.EstimateFirst)
	}
	if info.EstimateEco == nil || *info.EstimateEco != 50000 {
		t.Errorf("EstimateEco = %v, want 50000", info.EstimateEco)
	}
	// 部分设置
	info2 := NewFlightRoutePriceBuilder().EstimateEco(0).Build()
	if info2.EstimateEco == nil || *info2.EstimateEco != 0 {
		t.Errorf("EstimateEco = %v, want 0", info2.EstimateEco)
	}
	if info2.FlightNumber != nil {
		t.Errorf("FlightNumber = %v, want nil", info2.FlightNumber)
	}
}

func TestFlightInfoBuilder(t *testing.T) {
	route := NewRoutesBuilder().Build()
	info := NewFlightInfoBuilder().
		Routes([]Routes{*route}).
		Build()
	if len(info.Routes) != 1 {
		t.Errorf("Routes len = %d, want 1", len(info.Routes))
	}
	// 部分设置
	info2 := NewFlightInfoBuilder().Build()
	if len(info2.Routes) != 0 {
		t.Errorf("Routes len = %d, want 0", len(info2.Routes))
	}
}

func TestFlightListBuilder(t *testing.T) {
	fi := NewFlightInfoBuilder().Build()
	info := NewFlightListBuilder().
		FlightInfo(*fi).
		FlightRoutePrice([]FlightRoutePrice{*NewFlightRoutePriceBuilder().FlightNumber("MU5100").Build()}).
		Build()
	if info.FlightInfo == nil {
		t.Error("FlightInfo is nil")
	}
	if len(info.FlightRoutePrice) != 1 {
		t.Errorf("FlightRoutePrice len = %d, want 1", len(info.FlightRoutePrice))
	}
	// 部分设置
	info2 := NewFlightListBuilder().Build()
	if info2.FlightInfo != nil {
		t.Error("FlightInfo should be nil")
	}
	if len(info2.FlightRoutePrice) != 0 {
		t.Errorf("FlightRoutePrice len = %d, want 0", len(info2.FlightRoutePrice))
	}
}

func TestRoutesBuilder(t *testing.T) {
	dep := NewDepartureInfoBuilder().Build()
	arr := NewArrivalInfoBuilder().Build()
	air := NewAirlineInfoBuilder().FlightNumber("MU5100").Build()
	info := NewRoutesBuilder().
		DepartureInfo(*dep).
		ArrivalInfo(*arr).
		AirlineInfo(*air).
		Build()
	if info.DepartureInfo == nil {
		t.Error("DepartureInfo is nil")
	}
	if info.ArrivalInfo == nil {
		t.Error("ArrivalInfo is nil")
	}
	if info.AirlineInfo == nil || info.AirlineInfo.FlightNumber == nil || *info.AirlineInfo.FlightNumber != "MU5100" {
		t.Errorf("AirlineInfo.FlightNumber mismatch")
	}
	// 部分设置
	info2 := NewRoutesBuilder().AirlineInfo(*air).Build()
	if info2.AirlineInfo == nil {
		t.Error("AirlineInfo is nil")
	}
	if info2.DepartureInfo != nil {
		t.Error("DepartureInfo should be nil")
	}
}

func TestSelectConditionBuilder(t *testing.T) {
	info := NewSelectConditionBuilder().
		TransferCityList([]string{"北京", "上海"}).
		ToStationList([]string{"北京南"}).
		FromStationList([]string{"上海虹桥"}).
		IsShowCross(1).
		IsShowLocal(0).
		Build()
	if len(info.TransferCityList) != 2 {
		t.Errorf("TransferCityList len = %d, want 2", len(info.TransferCityList))
	}
	if info.IsShowCross == nil || *info.IsShowCross != 1 {
		t.Errorf("IsShowCross = %v, want 1", info.IsShowCross)
	}
	// 部分设置
	info2 := NewSelectConditionBuilder().IsShowLocal(1).Build()
	if info2.IsShowLocal == nil || *info2.IsShowLocal != 1 {
		t.Errorf("IsShowLocal = %v, want 1", info2.IsShowLocal)
	}
	if info2.TransferCityList != nil {
		t.Errorf("TransferCityList = %v, want nil", info2.TransferCityList)
	}
}

func TestDomesticFlightOrderListItemBuilder(t *testing.T) {
	info := NewDomesticFlightOrderListItemBuilder().Build()
	if info.OrderInfo != nil {
		t.Error("OrderInfo should be nil")
	}
	// 带字段构建
	oi := NewDomesticFlightOrderInfoBuilder().Build()
	info2 := NewDomesticFlightOrderListItemBuilder().
		OrderInfo(*oi).
		Build()
	if info2.OrderInfo == nil {
		t.Error("OrderInfo is nil")
	}
}

func TestTrainOrderListItemBuilder(t *testing.T) {
	info := NewTrainOrderListItemBuilder().Build()
	if info.OrderInfo != nil {
		t.Error("OrderInfo should be nil")
	}
}

func TestDomesticHotelOrderListItemBuilder(t *testing.T) {
	info := NewDomesticHotelOrderListItemBuilder().Build()
	if info.OrderInfo != nil {
		t.Error("OrderInfo should be nil")
	}
}

func TestDomesticFlightDataBuilder(t *testing.T) {
	info := NewDomesticFlightDataBuilder().
		OrderList([]DomesticFlightOrderListItem{*NewDomesticFlightOrderListItemBuilder().Build()}).
		Build()
	if len(info.OrderList) != 1 {
		t.Errorf("OrderList len = %d, want 1", len(info.OrderList))
	}
	// 部分设置
	info2 := NewDomesticFlightDataBuilder().Build()
	if len(info2.OrderList) != 0 {
		t.Errorf("OrderList len = %d, want 0", len(info2.OrderList))
	}
}

func TestDomesticTrainDataBuilder(t *testing.T) {
	info := NewDomesticTrainDataBuilder().
		OrderList([]TrainOrderListItem{*NewTrainOrderListItemBuilder().Build()}).
		Build()
	if len(info.OrderList) != 1 {
		t.Errorf("OrderList len = %d, want 1", len(info.OrderList))
	}
}

func TestDomesticFlightIdDataBuilder(t *testing.T) {
	pg := NewPageInfoBuilder().Total(10).Build()
	info := NewDomesticFlightIdDataBuilder().
		OrderIds("O1,O2").
		Page(*pg).
		Build()
	if info.OrderIds == nil || *info.OrderIds != "O1,O2" {
		t.Errorf("OrderIds = %v, want O1,O2", info.OrderIds)
	}
	if info.Page == nil || info.Page.Total == nil || *info.Page.Total != 10 {
		t.Errorf("Page.Total mismatch")
	}
	// 部分设置
	info2 := NewDomesticFlightIdDataBuilder().OrderIds("O3").Build()
	if info2.OrderIds == nil || *info2.OrderIds != "O3" {
		t.Errorf("OrderIds = %v, want O3", info2.OrderIds)
	}
	if info2.Page != nil {
		t.Errorf("Page = %v, want nil", info2.Page)
	}
}

func TestCarDataBuilder(t *testing.T) {
	pg := NewPageInfoBuilder().CurPage(1).Build()
	info := NewCarDataBuilder().
		OrderIds("C1,C2").
		Page(*pg).
		Build()
	if info.OrderIds == nil || *info.OrderIds != "C1,C2" {
		t.Errorf("OrderIds = %v, want C1,C2", info.OrderIds)
	}
	if info.Page == nil || info.Page.CurPage == nil || *info.Page.CurPage != 1 {
		t.Errorf("Page.CurPage mismatch")
	}
}

func TestTrainDataBuilder(t *testing.T) {
	info := NewTrainDataBuilder().OrderIds("T1,T2").Build()
	if info.OrderIds == nil || *info.OrderIds != "T1,T2" {
		t.Errorf("OrderIds = %v, want T1,T2", info.OrderIds)
	}
	if info.Page != nil {
		t.Errorf("Page = %v, want nil", info.Page)
	}
}

// ============================================================
// Request Builder 测试（POST 型）
// ============================================================

func TestGetFlightEstimatePriceRequestBuilder_AllFields(t *testing.T) {
	req := NewGetFlightEstimatePriceRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		DepartureCityId("BJS").
		ArrivalCityId("SHA").
		Date("2024-01-31").
		SearchType(1).
		Build()
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.DepartureCityId == nil || *req.DepartureCityId != "BJS" {
		t.Errorf("DepartureCityId = %v, want BJS", req.DepartureCityId)
	}
	if req.SearchType == nil || *req.SearchType != 1 {
		t.Errorf("SearchType = %v, want 1", req.SearchType)
	}
	// JSON 序列化验证
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if m["client_id"] != "test_client" {
		t.Errorf("json client_id = %v, want test_client", m["client_id"])
	}
	if m["search_type"].(float64) != 1 {
		t.Errorf("json search_type = %v, want 1", m["search_type"])
	}
}

func TestGetFlightEstimatePriceRequestBuilder_PartialFields(t *testing.T) {
	req := NewGetFlightEstimatePriceRequestBuilder().
		ClientId("test_client").
		DepartureCityId("BJS").
		Build()
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestGetFlightOrderDetailRequestBuilder_AllFields(t *testing.T) {
	req := NewGetFlightOrderDetailRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ProductType(1).
		OrderIds("O1,O2").
		Build()
	if req.ProductType == nil || *req.ProductType != 1 {
		t.Errorf("ProductType = %v, want 1", req.ProductType)
	}
	if req.OrderIds == nil || *req.OrderIds != "O1,O2" {
		t.Errorf("OrderIds = %v, want O1,O2", req.OrderIds)
	}
}

func TestGetHotelOrderDetailRequestBuilder_AllFields(t *testing.T) {
	req := NewGetHotelOrderDetailRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ProductType(2).
		OrderIds("H1,H2").
		Build()
	if req.ProductType == nil || *req.ProductType != 2 {
		t.Errorf("ProductType = %v, want 2", req.ProductType)
	}
	if req.OrderIds == nil || *req.OrderIds != "H1,H2" {
		t.Errorf("OrderIds = %v, want H1,H2", req.OrderIds)
	}
}

func TestGetTrainOrderDetailRequestBuilder_AllFields(t *testing.T) {
	req := NewGetTrainOrderDetailRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		OrderIds("T1,T2").
		Build()
	if req.OrderIds == nil || *req.OrderIds != "T1,T2" {
		t.Errorf("OrderIds = %v, want T1,T2", req.OrderIds)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
}

func TestListOrderRequestBuilder_AllFields(t *testing.T) {
	req := NewListOrderRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson(`{"order_type":"car"}`).
		ParamJsonObj(*NewParamJsonObjBuilder().OrderType("car").CurPage(1).Limit(20).Build()).
		Build()
	if req.ParamJson == nil || *req.ParamJson != `{"order_type":"car"}` {
		t.Errorf("ParamJson = %v, want {\"order_type\":\"car\"}", req.ParamJson)
	}
	if req.ParamJsonObj == nil {
		t.Fatal("ParamJsonObj is nil")
	}
	if req.ParamJsonObj.OrderType == nil || *req.ParamJsonObj.OrderType != "car" {
		t.Errorf("ParamJsonObj.OrderType = %v, want car", req.ParamJsonObj.OrderType)
	}
	if req.ParamJsonObj.CurPage == nil || *req.ParamJsonObj.CurPage != 1 {
		t.Errorf("ParamJsonObj.CurPage = %v, want 1", req.ParamJsonObj.CurPage)
	}
}

func TestListTrainLeftTicketRequestBuilder_AllFields(t *testing.T) {
	req := NewListTrainLeftTicketRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		TrainDate("2024-01-31").
		FromStationName("北京南").
		ToStationName("上海虹桥").
		StartTime("08:00:00").
		EndTime("23:59:59").
		Build()
	if req.TrainDate == nil || *req.TrainDate != "2024-01-31" {
		t.Errorf("TrainDate = %v, want 2024-01-31", req.TrainDate)
	}
	if req.FromStationName == nil || *req.FromStationName != "北京南" {
		t.Errorf("FromStationName = %v, want 北京南", req.FromStationName)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
}

func TestListTransferTrainTicketRequestBuilder_AllFields(t *testing.T) {
	req := NewListTransferTrainTicketRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		TrainDate("2024-01-31").
		FromStationName("北京南").
		ToStationName("上海虹桥").
		CurPage(1).
		Build()
	if req.CurPage == nil || *req.CurPage != 1 {
		t.Errorf("CurPage = %v, want 1", req.CurPage)
	}
	if req.TrainDate == nil || *req.TrainDate != "2024-01-31" {
		t.Errorf("TrainDate = %v, want 2024-01-31", req.TrainDate)
	}
}

func TestParamJsonObjBuilder_AllFields(t *testing.T) {
	st := NewSearchTimeBuilder().Build()
	info := NewParamJsonObjBuilder().
		OrderType("car").
		SearchTime(*st).
		CurPage(1).
		Limit(20).
		Build()
	if info.OrderType == nil || *info.OrderType != "car" {
		t.Errorf("OrderType = %v, want car", info.OrderType)
	}
	if info.CurPage == nil || *info.CurPage != 1 {
		t.Errorf("CurPage = %v, want 1", info.CurPage)
	}
	if info.Limit == nil || *info.Limit != 20 {
		t.Errorf("Limit = %v, want 20", info.Limit)
	}
	// 部分设置
	info2 := NewParamJsonObjBuilder().OrderType("train").Build()
	if info2.OrderType == nil || *info2.OrderType != "train" {
		t.Errorf("OrderType = %v, want train", info2.OrderType)
	}
	if info2.CurPage != nil {
		t.Errorf("CurPage = %v, want nil", info2.CurPage)
	}
}

// ============================================================
// GET 型 Request Builder 测试
// ============================================================

func TestGetCarOrderDetailApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetCarOrderDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		OrderId("1125922289295589").
		NeedAbnormalMsg(1).
		NeedRuleInfo(1).
		NeedCallEmployeeNumber(1).
		Sign("test_sign").
		Build()
	tests := []struct {
		key, want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"order_id", "1125922289295589"},
		{"need_abnormal_msg", "1"},
		{"need_rule_info", "1"},
		{"need_call_employee_number", "1"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetCarOrderDetailApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetCarOrderDetailApiReqBuilder().
		ClientId("test_client").
		OrderId("1125922289295589").
		Build()
	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("order_id") != "1125922289295589" {
		t.Errorf("order_id mismatch")
	}
	if req.apiReq.QueryParams.Get("need_abnormal_msg") != "" {
		t.Errorf("need_abnormal_msg should be empty")
	}
}

func TestGetCarOrderDetailApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetCarOrderDetailApiReqBuilder().
		ClientId("test_client").
		NeedAbnormalMsg(0).
		NeedRuleInfo(0).
		Build()
	if req.apiReq.QueryParams.Get("need_abnormal_msg") != "0" {
		t.Errorf("need_abnormal_msg = %q, want \"0\"", req.apiReq.QueryParams.Get("need_abnormal_msg"))
	}
	if req.apiReq.QueryParams.Get("need_rule_info") != "0" {
		t.Errorf("need_rule_info = %q, want \"0\"", req.apiReq.QueryParams.Get("need_rule_info"))
	}
}

func TestGetCarOrderDetailApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetCarOrderDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()
	for _, key := range []string{"order_id", "need_abnormal_msg", "need_rule_info", "need_call_employee_number"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

func TestGetOrderApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetOrderApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		CallPhone("13800000001").
		Phone("13800000002").
		StartDate("2024-01-01").
		EndDate("2024-01-31").
		StartTime("00:00:00").
		EndTime("23:59:59").
		UseCarType(2).
		PayType(0).
		IsInvoice(1).
		BudgetCenterId("BC001").
		OutBudgetId("OB001").
		Name("张三").
		Offset(0).
		Length(10).
		NeedApprovalId(1).
		NeedRuleInfo(1).
		NeedAbnormalMsg(1).
		NeedProjectInfo(1).
		NeedCallEmployeeNumber(1).
		Sign("test_sign").
		Build()
	tests := []struct {
		key, want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"call_phone", "13800000001"},
		{"phone", "13800000002"},
		{"start_date", "2024-01-01"},
		{"end_date", "2024-01-31"},
		{"use_car_type", "2"},
		{"pay_type", "0"},
		{"is_invoice", "1"},
		{"budget_center_id", "BC001"},
		{"out_budget_id", "OB001"},
		{"name", "张三"},
		{"offset", "0"},
		{"length", "10"},
		{"need_approval_id", "1"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetOrderApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetOrderApiReqBuilder().
		ClientId("test_client").
		Phone("13800000001").
		Build()
	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("phone") != "13800000001" {
		t.Errorf("phone mismatch")
	}
	if req.apiReq.QueryParams.Get("call_phone") != "" {
		t.Errorf("call_phone should be empty")
	}
}

func TestGetOrderApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetOrderApiReqBuilder().
		ClientId("test_client").
		UseCarType(0).
		PayType(0).
		Offset(0).
		Length(0).
		Build()
	if req.apiReq.QueryParams.Get("use_car_type") != "0" {
		t.Errorf("use_car_type = %q, want \"0\"", req.apiReq.QueryParams.Get("use_car_type"))
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
}

func TestGetOrderApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetOrderApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()
	for _, key := range []string{"call_phone", "phone", "use_car_type", "pay_type", "offset", "length"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

// ============================================================
// ApiReply 反序列化测试
// ============================================================

func TestGetCarOrderDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"order_id": "1125922289295589",
			"company_id": "1125915646135311",
			"call_phone": "13800000001",
			"passenger_name": "张三",
			"city_name": "北京",
			"start_name": "望京SOHO",
			"end_name": "北京南站",
			"status": "2",
			"pay_type": "0",
			"total_price": "25.50",
			"actual_price": "25.50",
			"company_pay": "25.50",
			"personal_pay": "0",
			"member_id": 1125922289295589,
			"regulation_id": 1125920826148759,
			"price": [{"name":"车费","amount":"25.50","type":"1"}],
			"budget_center_list": [{"app_name":"成本中心1","sequence":1,"id":"1125922289295589","value":"市场部"}],
			"stopover_points": [{"cityid":"1","name":"途经点","stop_id":1,"status":0}],
			"history_destinations": [{"address":"北京南站","name":"北京南站","sequence":1}],
			"passenger_list": [{"passenger_name":"张三","passenger_phone":"13800000001","is_traveler":0}],
			"rc_list": [{"rc_info":{"rc_type":"car_type","rc_code":"OVER_LIMIT","rc_reason":"超标"}}]
		},
		"request_id": "req_001"
	}`
	var reply GetCarOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.OrderId == nil || *reply.Data.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", reply.Data.OrderId)
	}
	if reply.Data.MemberId == nil || *reply.Data.MemberId != 1125922289295589 {
		t.Errorf("MemberId = %v, want 1125922289295589", reply.Data.MemberId)
	}
	if reply.Data.RegulationId == nil || *reply.Data.RegulationId != 1125920826148759 {
		t.Errorf("RegulationId = %v, want 1125920826148759", reply.Data.RegulationId)
	}
	if len(reply.Data.Price) != 1 {
		t.Fatalf("Price len = %d, want 1", len(reply.Data.Price))
	}
	if reply.Data.Price[0].Amount == nil || *reply.Data.Price[0].Amount != "25.50" {
		t.Errorf("Price[0].Amount = %v, want 25.50", reply.Data.Price[0].Amount)
	}
	if len(reply.Data.BudgetCenterList) != 1 {
		t.Fatalf("BudgetCenterList len = %d, want 1", len(reply.Data.BudgetCenterList))
	}
	if reply.Data.BudgetCenterList[0].Sequence == nil || *reply.Data.BudgetCenterList[0].Sequence != 1 {
		t.Errorf("BudgetCenterList[0].Sequence = %v, want 1", reply.Data.BudgetCenterList[0].Sequence)
	}
	if len(reply.Data.StopoverPoints) != 1 {
		t.Fatalf("StopoverPoints len = %d, want 1", len(reply.Data.StopoverPoints))
	}
	if reply.Data.StopoverPoints[0].StopId == nil || *reply.Data.StopoverPoints[0].StopId != 1 {
		t.Errorf("StopoverPoints[0].StopId = %v, want 1", reply.Data.StopoverPoints[0].StopId)
	}
	if len(reply.Data.HistoryDestinations) != 1 {
		t.Fatalf("HistoryDestinations len = %d, want 1", len(reply.Data.HistoryDestinations))
	}
	if len(reply.Data.PassengerList) != 1 {
		t.Fatalf("PassengerList len = %d, want 1", len(reply.Data.PassengerList))
	}
	if reply.Data.PassengerList[0].IsTraveler == nil || *reply.Data.PassengerList[0].IsTraveler != 0 {
		t.Errorf("PassengerList[0].IsTraveler = %v, want 0", reply.Data.PassengerList[0].IsTraveler)
	}
	if len(reply.Data.RcList) != 1 {
		t.Fatalf("RcList len = %d, want 1", len(reply.Data.RcList))
	}
	if reply.Data.RcList[0].RcInfo == nil || reply.Data.RcList[0].RcInfo.RcCode == nil || *reply.Data.RcList[0].RcInfo.RcCode != "OVER_LIMIT" {
		t.Errorf("RcList[0].RcInfo.RcCode mismatch")
	}
}

func TestGetCarOrderDetailApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
	var reply GetCarOrderDetailApiReply
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

func TestGetCarOrderDetailApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589","price":[],"budget_center_list":[],"stopover_points":[],"passenger_list":[]},"request_id":"req_empty"}`
	var reply GetCarOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Price) != 0 {
		t.Errorf("Price len = %d, want 0", len(reply.Data.Price))
	}
	if len(reply.Data.BudgetCenterList) != 0 {
		t.Errorf("BudgetCenterList len = %d, want 0", len(reply.Data.BudgetCenterList))
	}
}

func TestGetCarOrderDetailApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`
	var reply GetCarOrderDetailApiReply
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

func TestGetFlightEstimatePriceApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"flight_list": [
				{
					"flight_info": {
						"routes": [
							{
								"departure_info": {"airport_name":"北京首都机场"},
								"arrival_info": {"airport_name":"上海虹桥机场"},
								"airline_info": {"airline_name":"东方航空","flight_number":"MU5100"}
							}
						]
					},
					"flight_route_price": [
						{"flight_number":"MU5100","estimate_first":150000,"estimate_biz":100000,"estimate_eco":50000}
					]
				}
			]
		},
		"request_id": "req_001"
	}`
	var reply GetFlightEstimatePriceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.FlightList) != 1 {
		t.Fatalf("FlightList len = %d, want 1", len(reply.Data.FlightList))
	}
	if reply.Data.FlightList[0].FlightInfo == nil {
		t.Fatal("FlightInfo is nil")
	}
	if len(reply.Data.FlightList[0].FlightInfo.Routes) != 1 {
		t.Fatalf("Routes len = %d, want 1", len(reply.Data.FlightList[0].FlightInfo.Routes))
	}
	if reply.Data.FlightList[0].FlightInfo.Routes[0].AirlineInfo == nil {
		t.Fatal("AirlineInfo is nil")
	}
	if reply.Data.FlightList[0].FlightInfo.Routes[0].AirlineInfo.FlightNumber == nil || *reply.Data.FlightList[0].FlightInfo.Routes[0].AirlineInfo.FlightNumber != "MU5100" {
		t.Errorf("FlightNumber mismatch")
	}
	if len(reply.Data.FlightList[0].FlightRoutePrice) != 1 {
		t.Fatalf("FlightRoutePrice len = %d, want 1", len(reply.Data.FlightList[0].FlightRoutePrice))
	}
	if reply.Data.FlightList[0].FlightRoutePrice[0].EstimateEco == nil || *reply.Data.FlightList[0].FlightRoutePrice[0].EstimateEco != 50000 {
		t.Errorf("EstimateEco mismatch")
	}
}

func TestGetFlightEstimatePriceApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[]},"request_id":"req_empty"}`
	var reply GetFlightEstimatePriceApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.FlightList) != 0 {
		t.Errorf("FlightList len = %d, want 0", len(reply.Data.FlightList))
	}
}

func TestGetFlightEstimatePriceApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
	var reply GetFlightEstimatePriceApiReply
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

func TestGetFlightOrderDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 1,
			"domesticflight_data": {
				"order_list": [
					{"order_info": {"order_id":"FO001"}}
				]
			}
		},
		"request_id": "req_001"
	}`
	var reply GetFlightOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	// Data 是值类型（非指针），total 默认 0，但 JSON 提供 1
	if reply.Data.Total != 1 {
		t.Errorf("Total = %d, want 1", reply.Data.Total)
	}
	if reply.Data.DomesticflightData == nil {
		t.Fatal("DomesticflightData is nil")
	}
	if len(reply.Data.DomesticflightData.OrderList) != 1 {
		t.Fatalf("OrderList len = %d, want 1", len(reply.Data.DomesticflightData.OrderList))
	}
}

func TestGetFlightOrderDetailApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"domesticflight_data":{"order_list":[]}},"request_id":"req_empty"}`
	var reply GetFlightOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.Total != 0 {
		t.Errorf("Total = %d, want 0", reply.Data.Total)
	}
	if len(reply.Data.DomesticflightData.OrderList) != 0 {
		t.Errorf("OrderList len = %d, want 0", len(reply.Data.DomesticflightData.OrderList))
	}
}

func TestGetHotelOrderDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 1,
			"domestichotel_data": {
				"order_list": [
					{"order_info": {"order_id":"HO001"}}
				]
			}
		},
		"request_id": "req_001"
	}`
	var reply GetHotelOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data.Total != 1 {
		t.Errorf("Total = %d, want 1", reply.Data.Total)
	}
	if reply.Data.DomestichotelData == nil {
		t.Fatal("DomestichotelData is nil")
	}
	if len(reply.Data.DomestichotelData.OrderList) != 1 {
		t.Fatalf("OrderList len = %d, want 1", len(reply.Data.DomestichotelData.OrderList))
	}
}

func TestGetOrderApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{"order_id":"1125922289295589","call_phone":"13800000001","passenger_name":"张三","status":"2","pay_type":"0","total_price":"25.50"},
				{"order_id":"1125922289295590","passenger_name":"李四","status":"3"}
			]
		},
		"request_id": "req_001"
	}`
	var reply GetOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total == nil || *reply.Data.Total != 2 {
		t.Errorf("Total = %v, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 2 {
		t.Fatalf("Records len = %d, want 2", len(reply.Data.Records))
	}
	if reply.Data.Records[0].OrderId == nil || *reply.Data.Records[0].OrderId != "1125922289295589" {
		t.Errorf("Records[0].OrderId = %v, want 1125922289295589", reply.Data.Records[0].OrderId)
	}
	if reply.Data.Records[1].PassengerName == nil || *reply.Data.Records[1].PassengerName != "李四" {
		t.Errorf("Records[1].PassengerName = %v, want 李四", reply.Data.Records[1].PassengerName)
	}
}

func TestGetOrderApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_empty"}`
	var reply GetOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestGetOrderApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`
	var reply GetOrderApiReply
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

func TestGetTrainOrderDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 1,
			"domestictrain_data": {
				"order_list": [
					{"order_info": {"order_id":"TO001"}}
				]
			}
		},
		"request_id": "req_001"
	}`
	var reply GetTrainOrderDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != 1 {
		t.Errorf("Total = %d, want 1", reply.Data.Total)
	}
	if reply.Data.DomestictrainData == nil {
		t.Fatal("DomestictrainData is nil")
	}
	if len(reply.Data.DomestictrainData.OrderList) != 1 {
		t.Fatalf("OrderList len = %d, want 1", len(reply.Data.DomestictrainData.OrderList))
	}
}

func TestListOrderApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"car_data": {"order_ids":"C1,C2","page":{"cur_page":1,"limit":20,"total":2}},
			"train_data": {"order_ids":"T1"},
			"domesticflight_data": {"order_ids":"F1"}
		},
		"request_id": "req_001"
	}`
	var reply ListOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.CarData == nil || reply.Data.CarData.OrderIds == nil || *reply.Data.CarData.OrderIds != "C1,C2" {
		t.Errorf("CarData.OrderIds mismatch")
	}
	if reply.Data.CarData.Page == nil || reply.Data.CarData.Page.Total == nil || *reply.Data.CarData.Page.Total != 2 {
		t.Errorf("CarData.Page.Total mismatch")
	}
	if reply.Data.TrainData == nil || reply.Data.TrainData.OrderIds == nil || *reply.Data.TrainData.OrderIds != "T1" {
		t.Errorf("TrainData.OrderIds mismatch")
	}
}

func TestListOrderApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_empty"}`
	var reply ListOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.CarData != nil {
		t.Errorf("CarData = %v, want nil", reply.Data.CarData)
	}
}

func TestListTrainLeftTicketApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"list": [
				{"train_no":"G1","train_code":"G1","from_station_name":"北京南","to_station_name":"上海虹桥","start_time":"2024-01-31 09:00:00","arrive_time":"2024-01-31 13:28:00","travel_time":268,"sale_flag":0,"is_support_card":1,"day_difference":0,"ticket_data":[{"seat_type":"O","seat_type_name":"二等座","ticket_num":"有","ticket_price":55300}]}
			]
		},
		"request_id": "req_001"
	}`
	var reply ListTrainLeftTicketApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.List) != 1 {
		t.Fatalf("List len = %d, want 1", len(reply.Data.List))
	}
	if reply.Data.List[0].TrainNo == nil || *reply.Data.List[0].TrainNo != "G1" {
		t.Errorf("TrainNo = %v, want G1", reply.Data.List[0].TrainNo)
	}
	if reply.Data.List[0].TravelTime == nil || *reply.Data.List[0].TravelTime != 268 {
		t.Errorf("TravelTime = %v, want 268", reply.Data.List[0].TravelTime)
	}
	if len(reply.Data.List[0].TicketData) != 1 {
		t.Fatalf("TicketData len = %d, want 1", len(reply.Data.List[0].TicketData))
	}
	if reply.Data.List[0].TicketData[0].TicketPrice == nil || *reply.Data.List[0].TicketData[0].TicketPrice != 55300 {
		t.Errorf("TicketPrice = %v, want 55300", reply.Data.List[0].TicketData[0].TicketPrice)
	}
}

func TestListTrainLeftTicketApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_empty"}`
	var reply ListTrainLeftTicketApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.List) != 0 {
		t.Errorf("List len = %d, want 0", len(reply.Data.List))
	}
}

func TestListTransferTrainTicketApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"list": [
				{"supplier_id":1,"from_station_name":"北京南","to_station_name":"上海虹桥","transfer_station_name":"南京南","transfer_station_type":1,"transfer_city_name":"南京","total_runtime":"06:30","day_difference":0,"transfer_stop_time":"01:00","segment_items":[{"sequence":1,"train_no":"G1","from_station_name":"北京南","to_station_name":"南京南"}]}
			],
			"has_more": 0,
			"trace": "trace_001",
			"cur_page": 1,
			"page_size": 10
		},
		"request_id": "req_001"
	}`
	var reply ListTransferTrainTicketApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if len(reply.Data.List) != 1 {
		t.Fatalf("List len = %d, want 1", len(reply.Data.List))
	}
	if reply.Data.List[0].TransferStationType == nil || *reply.Data.List[0].TransferStationType != 1 {
		t.Errorf("TransferStationType = %v, want 1", reply.Data.List[0].TransferStationType)
	}
	if len(reply.Data.List[0].SegmentItems) != 1 {
		t.Fatalf("SegmentItems len = %d, want 1", len(reply.Data.List[0].SegmentItems))
	}
	if reply.Data.List[0].SegmentItems[0].Sequence == nil || *reply.Data.List[0].SegmentItems[0].Sequence != 1 {
		t.Errorf("SegmentItems[0].Sequence = %v, want 1", reply.Data.List[0].SegmentItems[0].Sequence)
	}
	if reply.Data.HasMore == nil || *reply.Data.HasMore != 0 {
		t.Errorf("HasMore = %v, want 0", reply.Data.HasMore)
	}
	if reply.Data.Trace == nil || *reply.Data.Trace != "trace_001" {
		t.Errorf("Trace = %v, want trace_001", reply.Data.Trace)
	}
}

func TestListTransferTrainTicketApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_empty"}`
	var reply ListTransferTrainTicketApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.List) != 0 {
		t.Errorf("List len = %d, want 0", len(reply.Data.List))
	}
}

// ============================================================
// 资源方法测试（9 方法 x 8 场景，表驱动）
// ============================================================

// orderMethodCase 描述一个资源方法测试场景
type orderMethodCase struct {
	name       string
	method     string // GET / POST
	path       string
	respBody   string
	respStatus int
	encKey     []byte // 非 nil 表示启用加密
	encUseURL  bool
	encNoData  bool // 启用加密但无 encrypt_data
	checkReq   func(t *testing.T, r *http.Request)
	wantErr    bool
}

// runOrderMethodCase 表驱动执行：构造 httptest，调用对应方法，校验返回
func runOrderMethodCase(t *testing.T, c orderMethodCase, call func(svc *order, reqOption *core.ReqOption) bool) {
	t.Helper()
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != c.method {
			t.Errorf("expected %s, got %s", c.method, r.Method)
		}
		if r.URL.Path != c.path {
			t.Errorf("expected path %s, got %s", c.path, r.URL.Path)
		}
		if c.checkReq != nil {
			c.checkReq(t, r)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(c.respStatus)
		w.Write([]byte(c.respBody))
	}
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	option := newOrderTestOption(testServer.URL)
	if c.encKey != nil {
		option.EnableEncryption = true
		option.EncryptionOption = &core.EncryptionOption{
			Ent: map[bool]int{false: 1, true: 2}[c.encUseURL],
			Key: string(c.encKey),
		}
	}
	// encNoData 场景：启用加密但响应无 encrypt_data
	if c.encNoData {
		option.EnableEncryption = true
		option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string([]byte("16byte-key-12345"))}
	}

	var reqOption *core.ReqOption
	if c.name == "WithReqOption" {
		customHeader := http.Header{}
		customHeader.Set("X-Custom-Header", "custom-value")
		reqOption = &core.ReqOption{Header: customHeader}
	}

	svc := &order{option: option}
	hasReply := call(svc, reqOption)
	if c.wantErr {
		return
	}
	if !hasReply && c.respStatus == http.StatusOK {
		t.Errorf("%s: expected reply not nil, got nil", c.name)
	}
	if hasReply && c.respStatus != http.StatusOK {
		t.Errorf("%s: expected reply nil for non-200, got not nil", c.name)
	}
}

// --- GetCarOrderDetail (GET) ---

func TestGetCarOrderDetail_AllScenarios(t *testing.T) {
	cases := []orderMethodCase{
		{
			name: "Success", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			respBody: `{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589","passenger_name":"张三","status":"2","total_price":"25.50"},"request_id":"req_001"}`,
			checkReq: func(t *testing.T, r *http.Request) {
				if r.URL.Query().Get("order_id") != "1125922289295589" {
					t.Errorf("order_id = %q", r.URL.Query().Get("order_id"))
				}
			},
		},
		{name: "EmptyData", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			respBody: `{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589","price":[]},"request_id":"req_002"}`},
		{name: "ApiError", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusInternalServerError,
			respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(`{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589","status":"2"},"request_id":"req_enc"}`, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			encKey:    []byte("32byte-key-1234567890abcdefghijk"),
			encUseURL: true,
			respBody:  encryptResp(`{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589"},"request_id":"req_enc256"}`, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			encNoData: true,
			respBody:  `{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589"},"request_id":"req_noenc"}`},
		{name: "WithReqOption", method: http.MethodGet, path: "/river/Order/detail", respStatus: http.StatusOK,
			respBody: `{"errno":0,"errmsg":"SUCCESS","data":{"order_id":"1125922289295589"},"request_id":"req_opt"}`,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q, want custom-value", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetCarOrderDetailApiReqBuilder().
					ClientId("test_client").
					AccessToken("test_token").
					CompanyId("test_company").
					Timestamp(1583484681).
					OrderId("1125922289295589").
					Sign("test_sign").
					Build()
				resp, err := svc.GetCarOrderDetail(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetCarOrderDetail() error = %v", err)
				}
				return resp.GetCarOrderDetailApiReply != nil
			})
		})
	}
}

// --- GetFlightEstimatePrice (POST) ---

func TestGetFlightEstimatePrice_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[{"flight_info":{"routes":[{"airline_info":{"flight_number":"MU5100"}}]},"flight_route_price":[{"estimate_eco":50000}]}]},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[]},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(`{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[{"flight_info":{"routes":[{"airline_info":{"flight_number":"MU5100"}}]}}]},"request_id":"req_enc"}`, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(`{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[]},"request_id":"req_enc256"}`, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK,
			encNoData: true, respBody: `{"errno":0,"errmsg":"SUCCESS","data":{"flight_list":[]},"request_id":"req_noenc"}`},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/g/flight/info/estimatePrice", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetFlightEstimatePriceApiReqBuilder().
					GetFlightEstimatePriceRequest(NewGetFlightEstimatePriceRequestBuilder().
						ClientId("test_client").
						DepartureCityId("BJS").
						ArrivalCityId("SHA").
						Date("2024-01-31").
						SearchType(1).
						Build()).
					Build()
				resp, err := svc.GetFlightEstimatePrice(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetFlightEstimatePrice() error = %v", err)
				}
				return resp.GetFlightEstimatePriceApiReply != nil
			})
		})
	}
}

// --- GetFlightOrderDetail (POST) ---

func TestGetFlightOrderDetail_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"domesticflight_data":{"order_list":[{"order_info":{"order_id":"FO001"}}]}},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"domesticflight_data":{"order_list":[]}},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/g/flight/orderDetail", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetFlightOrderDetailApiReqBuilder().
					GetFlightOrderDetailRequest(NewGetFlightOrderDetailRequestBuilder().
						ClientId("test_client").
						ProductType(1).
						OrderIds("FO001").
						Build()).
					Build()
				resp, err := svc.GetFlightOrderDetail(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetFlightOrderDetail() error = %v", err)
				}
				return resp.GetFlightOrderDetailApiReply != nil
			})
		})
	}
}

// --- GetHotelOrderDetail (POST) ---

func TestGetHotelOrderDetail_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"domestichotel_data":{"order_list":[{"order_info":{"order_id":"HO001"}}]}},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"domestichotel_data":{"order_list":[]}},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/g/hotel/orderDetail", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetHotelOrderDetailApiReqBuilder().
					GetHotelOrderDetailRequest(NewGetHotelOrderDetailRequestBuilder().
						ClientId("test_client").
						ProductType(1).
						OrderIds("HO001").
						Build()).
					Build()
				resp, err := svc.GetHotelOrderDetail(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetHotelOrderDetail() error = %v", err)
				}
				return resp.GetHotelOrderDetailApiReply != nil
			})
		})
	}
}

// --- GetOrder (GET) ---

func TestGetOrder_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"order_id":"1125922289295589","passenger_name":"张三","status":"2","total_price":"25.50"}]},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK, respBody: successBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if v := r.URL.Query().Get("offset"); v != "0" {
					t.Errorf("offset = %q, want 0", v)
				}
			}},
		{name: "EmptyData", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(`{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"order_id":"1125922289295589"}]},"request_id":"req_enc"}`, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodGet, path: "/river/Order/get", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetOrderApiReqBuilder().
					ClientId("test_client").
					AccessToken("test_token").
					CompanyId("test_company").
					Timestamp("1583484681").
					Offset(0).
					Length(10).
					Sign("test_sign").
					Build()
				resp, err := svc.GetOrder(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetOrder() error = %v", err)
				}
				return resp.GetOrderApiReply != nil
			})
		})
	}
}

// --- GetTrainOrderDetail (POST) ---

func TestGetTrainOrderDetail_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"domestictrain_data":{"order_list":[{"order_info":{"order_id":"TO001"}}]}},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"domestictrain_data":{"order_list":[]}},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/g/train/orderDetail", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewGetTrainOrderDetailApiReqBuilder().
					GetTrainOrderDetailRequest(NewGetTrainOrderDetailRequestBuilder().
						ClientId("test_client").
						OrderIds("TO001").
						Build()).
					Build()
				resp, err := svc.GetTrainOrderDetail(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("GetTrainOrderDetail() error = %v", err)
				}
				return resp.GetTrainOrderDetailApiReply != nil
			})
		})
	}
}

// --- ListOrder (POST) ---

func TestListOrder_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"car_data":{"order_ids":"C1,C2","page":{"cur_page":1,"limit":20,"total":2}}},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/open-apis/v1/order/list", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewListOrderApiReqBuilder().
					ListOrderRequest(NewListOrderRequestBuilder().
						ClientId("test_client").
						ParamJson(`{"order_type":"car"}`).
						Build()).
					Build()
				resp, err := svc.ListOrder(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("ListOrder() error = %v", err)
				}
				return resp.ListOrderApiReply != nil
			})
		})
	}
}

// --- ListTrainLeftTicket (POST) ---

func TestListTrainLeftTicket_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[{"train_no":"G1","train_code":"G1","from_station_name":"北京南","to_station_name":"上海虹桥","start_time":"2024-01-31 09:00:00","arrive_time":"2024-01-31 13:28:00","travel_time":268,"sale_flag":0,"ticket_data":[{"seat_type":"O","seat_type_name":"二等座","ticket_num":"有","ticket_price":55300}]}]},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/train/queryLeftTicket", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewListTrainLeftTicketApiReqBuilder().
					ListTrainLeftTicketRequest(NewListTrainLeftTicketRequestBuilder().
						ClientId("test_client").
						TrainDate("2024-01-31").
						FromStationName("北京南").
						ToStationName("上海虹桥").
						Build()).
					Build()
				resp, err := svc.ListTrainLeftTicket(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("ListTrainLeftTicket() error = %v", err)
				}
				return resp.ListTrainLeftTicketApiReply != nil
			})
		})
	}
}

// --- ListTransferTrainTicket (POST) ---

func TestListTransferTrainTicket_AllScenarios(t *testing.T) {
	successBody := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[{"supplier_id":1,"from_station_name":"北京南","to_station_name":"上海虹桥","transfer_station_name":"南京南","transfer_station_type":1,"transfer_city_name":"南京","total_runtime":"06:30","day_difference":0,"transfer_stop_time":"01:00","segment_items":[{"sequence":1,"train_no":"G1","from_station_name":"北京南","to_station_name":"南京南"}]}],"has_more":0,"trace":"trace_001","cur_page":1,"page_size":10},"request_id":"req_001"}`
	emptyBody := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_002"}`
	cases := []orderMethodCase{
		{name: "Success", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK, respBody: successBody},
		{name: "EmptyData", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK, respBody: emptyBody},
		{name: "ApiError", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK, respBody: `{"errno":10003,"errmsg":"param error","request_id":"req_003"}`},
		{name: "HttpError", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusInternalServerError, respBody: ``},
		{name: "WithEncryption_AES128", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK,
			encKey:   []byte("16byte-key-12345"),
			respBody: encryptResp(successBody, []byte("16byte-key-12345"), false)},
		{name: "WithEncryption_AES256", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK,
			encKey: []byte("32byte-key-1234567890abcdefghijk"), encUseURL: true,
			respBody: encryptResp(emptyBody, []byte("32byte-key-1234567890abcdefghijk"), true)},
		{name: "EncryptionNoEncryptData", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK,
			encNoData: true, respBody: emptyBody},
		{name: "WithReqOption", method: http.MethodPost, path: "/api-gateway/g/train/transfer/queryLeftTicket", respStatus: http.StatusOK, respBody: emptyBody,
			checkReq: func(t *testing.T, r *http.Request) {
				if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
					t.Errorf("X-Custom-Header = %q", h)
				}
			}},
	}
	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			runOrderMethodCase(t, tc, func(svc *order, reqOption *core.ReqOption) bool {
				req := NewListTransferTrainTicketApiReqBuilder().
					ListTransferTrainTicketRequest(NewListTransferTrainTicketRequestBuilder().
						ClientId("test_client").
						TrainDate("2024-01-31").
						FromStationName("北京南").
						ToStationName("上海虹桥").
						CurPage(1).
						Build()).
					Build()
				resp, err := svc.ListTransferTrainTicket(context.Background(), req, reqOption)
				if err != nil {
					t.Fatalf("ListTransferTrainTicket() error = %v", err)
				}
				return resp.ListTransferTrainTicketApiReply != nil
			})
		})
	}
}
