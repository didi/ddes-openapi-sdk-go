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

// =====================================================================
// 公共 helper
// =====================================================================

func newApprovalTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// =====================================================================
// 数据模型 Builder 测试
// =====================================================================

func TestErrorInfoListItemBuilder(t *testing.T) {
	info := NewErrorInfoListItemBuilder().ErrNo(10003).ErrMsg("参数错误").Build()
	if info.ErrNo == nil || *info.ErrNo != 10003 {
		t.Errorf("ErrNo = %v, want 10003", info.ErrNo)
	}
	if info.ErrMsg == nil || *info.ErrMsg != "参数错误" {
		t.Errorf("ErrMsg = %v, want 参数错误", info.ErrMsg)
	}
	info2 := NewErrorInfoListItemBuilder().ErrNo(0).Build()
	if info2.ErrNo == nil || *info2.ErrNo != 0 {
		t.Errorf("ErrNo = %v, want 0", info2.ErrNo)
	}
	if info2.ErrMsg != nil {
		t.Errorf("ErrMsg = %v, want nil", info2.ErrMsg)
	}
}

func TestCurAppRoverBuilder(t *testing.T) {
	info := NewCurAppRoverBuilder().Type("phone").Value("13800000001").Build()
	if info.Type == nil || *info.Type != "phone" {
		t.Errorf("Type = %v, want phone", info.Type)
	}
	if info.Value == nil || *info.Value != "13800000001" {
		t.Errorf("Value = %v, want 13800000001", info.Value)
	}
	info2 := NewCurAppRoverBuilder().Type("email").Build()
	if info2.Type == nil || *info2.Type != "email" {
		t.Errorf("Type = %v, want email", info2.Type)
	}
	if info2.Value != nil {
		t.Errorf("Value = %v, want nil", info2.Value)
	}
}

func TestBudgetCenterListItemBuilder(t *testing.T) {
	info := NewBudgetCenterListItemBuilder().Sequence(1).Id("1125920020961744").Value("研发部").Code("DEPT001").Build()
	if info.Sequence == nil || *info.Sequence != 1 {
		t.Errorf("Sequence = %v, want 1", info.Sequence)
	}
	if info.Id == nil || *info.Id != "1125920020961744" {
		t.Errorf("Id = %v, want 1125920020961744", info.Id)
	}
	if info.Value == nil || *info.Value != "研发部" {
		t.Errorf("Value = %v, want 研发部", info.Value)
	}
	if info.Code == nil || *info.Code != "DEPT001" {
		t.Errorf("Code = %v, want DEPT001", info.Code)
	}
	info2 := NewBudgetCenterListItemBuilder().Sequence(2).Build()
	if info2.Sequence == nil || *info2.Sequence != 2 {
		t.Errorf("Sequence = %v, want 2", info2.Sequence)
	}
	if info2.Id != nil {
		t.Errorf("Id = %v, want nil", info2.Id)
	}
}

func TestPassengerInfoBuilder(t *testing.T) {
	info := NewPassengerInfoBuilder().
		Name("张三").Phone("13800000001").PassengerType(0).
		EmployeeNumber("D0001").Email("zhangsan@example.com").
		BudgetCenterList([]BudgetCenterListItem{*NewBudgetCenterListItemBuilder().Sequence(1).Build()}).
		Build()
	if info.Name == nil || *info.Name != "张三" {
		t.Errorf("Name = %v, want 张三", info.Name)
	}
	if info.PassengerType == nil || *info.PassengerType != 0 {
		t.Errorf("PassengerType = %v, want 0", info.PassengerType)
	}
	if len(info.BudgetCenterList) != 1 {
		t.Errorf("BudgetCenterList len = %d, want 1", len(info.BudgetCenterList))
	}
	info2 := NewPassengerInfoBuilder().Name("李四").Build()
	if info2.Name == nil || *info2.Name != "李四" {
		t.Errorf("Name = %v, want 李四", info2.Name)
	}
	if info2.Phone != nil {
		t.Errorf("Phone = %v, want nil", info2.Phone)
	}
}

func TestTravelerBuilder(t *testing.T) {
	info := NewTravelerBuilder().
		Name("张三").Phone("13800000001").PassengerType("0").
		PassengerNumber("PN001").Email("zhangsan@example.com").Build()
	if info.Name == nil || *info.Name != "张三" {
		t.Errorf("Name = %v, want 张三", info.Name)
	}
	if info.PassengerType == nil || *info.PassengerType != "0" {
		t.Errorf("PassengerType = %v, want 0", info.PassengerType)
	}
	info2 := NewTravelerBuilder().Name("李四").Build()
	if info2.Phone != nil {
		t.Errorf("Phone = %v, want nil", info2.Phone)
	}
}

func TestTravelBudgetBuilder(t *testing.T) {
	info := NewTravelBudgetBuilder().BudgetAmount(50000).BudgetType(1).BudgetShare([]int32{1, 2, 301}).Build()
	if info.BudgetAmount == nil || *info.BudgetAmount != 50000 {
		t.Errorf("BudgetAmount = %v, want 50000", info.BudgetAmount)
	}
	if info.BudgetType == nil || *info.BudgetType != 1 {
		t.Errorf("BudgetType = %v, want 1", info.BudgetType)
	}
	if len(info.BudgetShare) != 3 {
		t.Errorf("BudgetShare len = %d, want 3", len(info.BudgetShare))
	}
	info2 := NewTravelBudgetBuilder().BudgetType(1).Build()
	if info2.BudgetAmount != nil {
		t.Errorf("BudgetAmount = %v, want nil", info2.BudgetAmount)
	}
}

func TestTravelManagementBuilder(t *testing.T) {
	info := NewTravelManagementBuilder().
		DailyAmountControl([]DailyAmountControl{*NewDailyAmountControlBuilder().DailyAmount(10000).Build()}).
		Build()
	if len(info.DailyAmountControl) != 1 {
		t.Errorf("DailyAmountControl len = %d, want 1", len(info.DailyAmountControl))
	}
	info2 := NewTravelManagementBuilder().Build()
	if info2.DailyAmountControl != nil {
		t.Errorf("DailyAmountControl = %v, want nil", info2.DailyAmountControl)
	}
}

func TestOutTripInfoBuilder(t *testing.T) {
	info := NewOutTripInfoBuilder().BeginTime(1675602713).EndTime(1675689113).TravelPurpose("北京出差").Build()
	if info.BeginTime == nil || *info.BeginTime != 1675602713 {
		t.Errorf("BeginTime = %v, want 1675602713", info.BeginTime)
	}
	if info.TravelPurpose == nil || *info.TravelPurpose != "北京出差" {
		t.Errorf("TravelPurpose = %v, want 北京出差", info.TravelPurpose)
	}
	info2 := NewOutTripInfoBuilder().BeginTime(0).Build()
	if info2.EndTime != nil {
		t.Errorf("EndTime = %v, want nil", info2.EndTime)
	}
}

func TestExtendFieldListBuilder(t *testing.T) {
	info := NewExtendFieldListBuilder().ExtendField01("field1").ExtendField02("field2").ExtendField03("field3").Build()
	if info.ExtendField01 == nil || *info.ExtendField01 != "field1" {
		t.Errorf("ExtendField01 = %v, want field1", info.ExtendField01)
	}
	if info.ExtendField03 == nil || *info.ExtendField03 != "field3" {
		t.Errorf("ExtendField03 = %v, want field3", info.ExtendField03)
	}
	info2 := NewExtendFieldListBuilder().ExtendField01("only").Build()
	if info2.ExtendField02 != nil {
		t.Errorf("ExtendField02 = %v, want nil", info2.ExtendField02)
	}
}

func TestTravelDetailBuilder(t *testing.T) {
	info := NewTravelDetailBuilder().
		StartDate("2026-01-01").EndDate("2026-01-14").
		Trips([]Trip{*NewTripBuilder().DepartureCity("北京").Build()}).
		StartCityRule(1).EndCityRule(0).
		TrainTotalCount(2).FlightTotalCount(1).HotelTotalCount(3).PickupTotalCount(2).
		CategoryControl([]int32{3, 4, 6}).Build()
	if info.StartDate == nil || *info.StartDate != "2026-01-01" {
		t.Errorf("StartDate = %v, want 2026-01-01", info.StartDate)
	}
	if info.StartCityRule == nil || *info.StartCityRule != 1 {
		t.Errorf("StartCityRule = %v, want 1", info.StartCityRule)
	}
	if len(info.Trips) != 1 {
		t.Errorf("Trips len = %d, want 1", len(info.Trips))
	}
	if len(info.CategoryControl) != 3 {
		t.Errorf("CategoryControl len = %d, want 3", len(info.CategoryControl))
	}
	info2 := NewTravelDetailBuilder().StartDate("2026-02-01").Build()
	if info2.EndDate != nil {
		t.Errorf("EndDate = %v, want nil", info2.EndDate)
	}
}

func TestTripPassengerBuilder(t *testing.T) {
	info := NewTripPassengerBuilder().
		PassengerType(0).PassengerName("张三").MemberType(0).
		PassengerPhone("13800000001").EmployeeNumber("D0001").Email("zhangsan@example.com").
		TravelerId("T001").PassengerEnglishSurname("Zhang").PassengerEnglishName("San").
		OutTravelerId("OT001").Build()
	if info.PassengerType == nil || *info.PassengerType != 0 {
		t.Errorf("PassengerType = %v, want 0", info.PassengerType)
	}
	if info.PassengerName == nil || *info.PassengerName != "张三" {
		t.Errorf("PassengerName = %v, want 张三", info.PassengerName)
	}
	if info.OutTravelerId == nil || *info.OutTravelerId != "OT001" {
		t.Errorf("OutTravelerId = %v, want OT001", info.OutTravelerId)
	}
	info2 := NewTripPassengerBuilder().PassengerName("李四").Build()
	if info2.PassengerPhone != nil {
		t.Errorf("PassengerPhone = %v, want nil", info2.PassengerPhone)
	}
}

func TestTripBuilder(t *testing.T) {
	info := NewTripBuilder().
		DepartureCity("北京").DepartureCityId(1).
		DestinationCity("上海").DestinationCityId(2).
		StartDate("2026-01-01").EndDate("2026-01-03").
		TripType("1,2").IsReturn(1).
		DepartureAddressDimension(0).DepartureCountryId(1).DepartureCountryName("中国").
		DepartureProvinceId(11).DepartureProvinceName("北京市").
		DestinationAddressDimension(0).DestinationCountryId(1).DestinationCountryName("中国").
		DestinationProvinceId(31).DestinationProvinceName("上海市").
		ToCitys([]TravelCity{*NewTravelCityBuilder().Id(1).Name("北京").Build()}).
		Build()
	if info.DepartureCity == nil || *info.DepartureCity != "北京" {
		t.Errorf("DepartureCity = %v, want 北京", info.DepartureCity)
	}
	if info.DepartureCityId == nil || *info.DepartureCityId != 1 {
		t.Errorf("DepartureCityId = %v, want 1", info.DepartureCityId)
	}
	if info.IsReturn == nil || *info.IsReturn != 1 {
		t.Errorf("IsReturn = %v, want 1", info.IsReturn)
	}
	if len(info.ToCitys) != 1 {
		t.Errorf("ToCitys len = %d, want 1", len(info.ToCitys))
	}
	info2 := NewTripBuilder().DepartureCity("广州").Build()
	if info2.DestinationCity != nil {
		t.Errorf("DestinationCity = %v, want nil", info2.DestinationCity)
	}
}

func TestMeetingTripBuilder(t *testing.T) {
	info := NewMeetingTripBuilder().
		MeetingCity([]BusinessCity{*NewBusinessCityBuilder().CityId(1).City("北京").Build()}).Build()
	if len(info.MeetingCity) != 1 {
		t.Errorf("MeetingCity len = %d, want 1", len(info.MeetingCity))
	}
	info2 := NewMeetingTripBuilder().Build()
	if info2.MeetingCity != nil {
		t.Errorf("MeetingCity = %v, want nil", info2.MeetingCity)
	}
}

func TestBusinessTripDetailByTimesBuilder(t *testing.T) {
	info := NewBusinessTripDetailByTimesBuilder().
		StartTime("2026-01-01 09:00:00").EndTime("2026-01-01 18:00:00").
		DepartureCityId(1).DepartureCity("北京").
		DestinationCityId(2).DestinationCity("上海").
		StartName("望京SOHO").StartAddress("朝阳区望京SOHO").
		Flat(39.99).Flng(116.48).
		EndName("陆家嘴").EndAddress("浦东新区陆家嘴").
		Tlat(31.24).Tlng(121.50).
		IsReturn(0).TripTimes(1).
		PerorderMoneyQuota(10000).TotalMoneyQuota(50000).Build()
	if info.StartTime == nil || *info.StartTime != "2026-01-01 09:00:00" {
		t.Errorf("StartTime = %v, want 2026-01-01 09:00:00", info.StartTime)
	}
	if info.Flat == nil || *info.Flat != 39.99 {
		t.Errorf("Flat = %v, want 39.99", info.Flat)
	}
	if info.TripTimes == nil || *info.TripTimes != 1 {
		t.Errorf("TripTimes = %v, want 1", info.TripTimes)
	}
	info2 := NewBusinessTripDetailByTimesBuilder().StartTime("2026-02-01 09:00:00").Build()
	if info2.EndTime != nil {
		t.Errorf("EndTime = %v, want nil", info2.EndTime)
	}
}

func TestBusinessTripDetailByDateBuilder(t *testing.T) {
	info := NewBusinessTripDetailByDateBuilder().
		StartTime("2026-01-01 09:00:00").EndTime("2026-01-03 18:00:00").
		Trips([]BusinessCity{*NewBusinessCityBuilder().CityId(1).City("北京").Build()}).
		TripAmount(50000).TripTimes(3).PerorderMoneyQuota(10000).Build()
	if info.StartTime == nil || *info.StartTime != "2026-01-01 09:00:00" {
		t.Errorf("StartTime = %v, want 2026-01-01 09:00:00", info.StartTime)
	}
	if info.TripAmount == nil || *info.TripAmount != 50000 {
		t.Errorf("TripAmount = %v, want 50000", info.TripAmount)
	}
	if len(info.Trips) != 1 {
		t.Errorf("Trips len = %d, want 1", len(info.Trips))
	}
	info2 := NewBusinessTripDetailByDateBuilder().StartTime("2026-02-01 09:00:00").Build()
	if info2.TripAmount != nil {
		t.Errorf("TripAmount = %v, want nil", info2.TripAmount)
	}
}

func TestDailyAmountControlBuilder(t *testing.T) {
	info := NewDailyAmountControlBuilder().DailyAmount(10000).ControlType(1).ControlProduct([]int32{1}).Build()
	if info.DailyAmount == nil || *info.DailyAmount != 10000 {
		t.Errorf("DailyAmount = %v, want 10000", info.DailyAmount)
	}
	if info.ControlType == nil || *info.ControlType != 1 {
		t.Errorf("ControlType = %v, want 1", info.ControlType)
	}
	if len(info.ControlProduct) != 1 {
		t.Errorf("ControlProduct len = %d, want 1", len(info.ControlProduct))
	}
	info2 := NewDailyAmountControlBuilder().DailyAmount(0).Build()
	if info2.ControlType != nil {
		t.Errorf("ControlType = %v, want nil", info2.ControlType)
	}
}

func TestCarRuleBuilder(t *testing.T) {
	info := NewCarRuleBuilder().
		RuleId("R001").RuleName("市内用车规则").RuleStatus(1).
		CityId("1").CityName("北京").UseScene(1).
		TotalCount(10).AvailableCount(8).
		TotalMoney(500.50).AvailableMoney(400.25).PerorderMoney(100.00).
		StartTime("2026-01-01 00:00:00").EndTime("2026-01-31 23:59:59").Build()
	if info.RuleId == nil || *info.RuleId != "R001" {
		t.Errorf("RuleId = %v, want R001", info.RuleId)
	}
	if info.RuleStatus == nil || *info.RuleStatus != 1 {
		t.Errorf("RuleStatus = %v, want 1", info.RuleStatus)
	}
	if info.TotalMoney == nil || *info.TotalMoney != 500.50 {
		t.Errorf("TotalMoney = %v, want 500.50", info.TotalMoney)
	}
	info2 := NewCarRuleBuilder().RuleId("R002").Build()
	if info2.RuleName != nil {
		t.Errorf("RuleName = %v, want nil", info2.RuleName)
	}
}

func TestHotelRuleBuilder(t *testing.T) {
	info := NewHotelRuleBuilder().
		RuleId("H001").RuleName("酒店规则").RuleStatus("1").
		CityList([]TravelCity{*NewTravelCityBuilder().Id(1).Name("北京").Build()}).
		TotalCount(3).AvailableCount(2).
		StartTime("2026-01-01").EndTime("2026-01-03").Build()
	if info.RuleId == nil || *info.RuleId != "H001" {
		t.Errorf("RuleId = %v, want H001", info.RuleId)
	}
	if info.RuleStatus == nil || *info.RuleStatus != "1" {
		t.Errorf("RuleStatus = %v, want 1", info.RuleStatus)
	}
	if len(info.CityList) != 1 {
		t.Errorf("CityList len = %d, want 1", len(info.CityList))
	}
	info2 := NewHotelRuleBuilder().RuleId("H002").Build()
	if info2.CityList != nil {
		t.Errorf("CityList = %v, want nil", info2.CityList)
	}
}

func TestFlightRuleBuilder(t *testing.T) {
	info := NewFlightRuleBuilder().
		RuleId("F001").RuleName("航班规则").RuleStatus("1").
		TotalCount(2).AvailableCount(1).
		StartTime("2026-01-01").EndTime("2026-01-03").
		CityLine([]CityLine{*NewCityLineBuilder().StartId("1").StartName("北京").EndId("2").EndName("上海").Build()}).Build()
	if info.RuleId == nil || *info.RuleId != "F001" {
		t.Errorf("RuleId = %v, want F001", info.RuleId)
	}
	if len(info.CityLine) != 1 {
		t.Errorf("CityLine len = %d, want 1", len(info.CityLine))
	}
	info2 := NewFlightRuleBuilder().RuleId("F002").Build()
	if info2.CityLine != nil {
		t.Errorf("CityLine = %v, want nil", info2.CityLine)
	}
}

func TestTrainRuleBuilder(t *testing.T) {
	info := NewTrainRuleBuilder().
		RuleId("T001").RuleName("火车规则").RuleStatus("1").
		TotalCount(2).AvailableCount(1).
		StartTime("2026-01-01").EndTime("2026-01-03").
		CityLine([]CityLine{*NewCityLineBuilder().StartId("1").StartName("北京").EndId("2").EndName("上海").Build()}).Build()
	if info.RuleId == nil || *info.RuleId != "T001" {
		t.Errorf("RuleId = %v, want T001", info.RuleId)
	}
	if len(info.CityLine) != 1 {
		t.Errorf("CityLine len = %d, want 1", len(info.CityLine))
	}
}

func TestBusinessCityBuilder(t *testing.T) {
	info := NewBusinessCityBuilder().CityId(1).City("北京").Build()
	if info.CityId == nil || *info.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info.CityId)
	}
	if info.City == nil || *info.City != "北京" {
		t.Errorf("City = %v, want 北京", info.City)
	}
	info2 := NewBusinessCityBuilder().CityId(2).Build()
	if info2.City != nil {
		t.Errorf("City = %v, want nil", info2.City)
	}
}

func TestTravelCityBuilder(t *testing.T) {
	info := NewTravelCityBuilder().
		Id(1).Name("北京").AddressDimension(0).
		CountryId(1).CountryName("中国").
		ProvinceId(11).ProvinceName("北京市").Build()
	if info.Id == nil || *info.Id != 1 {
		t.Errorf("Id = %v, want 1", info.Id)
	}
	if info.Name == nil || *info.Name != "北京" {
		t.Errorf("Name = %v, want 北京", info.Name)
	}
	if info.AddressDimension == nil || *info.AddressDimension != 0 {
		t.Errorf("AddressDimension = %v, want 0", info.AddressDimension)
	}
	info2 := NewTravelCityBuilder().Id(2).Build()
	if info2.Name != nil {
		t.Errorf("Name = %v, want nil", info2.Name)
	}
}

func TestCityLineBuilder(t *testing.T) {
	info := NewCityLineBuilder().StartId("1").StartName("北京").EndId("2").EndName("上海").Build()
	if info.StartId == nil || *info.StartId != "1" {
		t.Errorf("StartId = %v, want 1", info.StartId)
	}
	if info.EndName == nil || *info.EndName != "上海" {
		t.Errorf("EndName = %v, want 上海", info.EndName)
	}
	info2 := NewCityLineBuilder().StartId("3").Build()
	if info2.EndId != nil {
		t.Errorf("EndId = %v, want nil", info2.EndId)
	}
}

func TestApprovalDetailBuilder(t *testing.T) {
	info := NewApprovalDetailBuilder().
		ApprovalId("AP001").OutApprovalId("TA_001").ApprovalType(1).Type(0).
		BudgetCenterId("1125920020961744").Reason("北京出差").ExtraInfo("{}").
		StartDate("2026-01-01").EndDate("2026-01-14").ApprovalStatus(1).
		SceneId(1).CityType(1).
		TravelBudget(*NewTravelBudgetBuilder().BudgetAmount(50000).Build()).
		PassengerList([]PassengerInfo{*NewPassengerInfoBuilder().Name("张三").Build()}).
		TravelerList([]Traveler{*NewTravelerBuilder().Name("李四").Build()}).
		Applicant(*NewPassengerInfoBuilder().Name("申请人").Build()).
		TravelDetail("{}").BusinessTripDetail("{}").
		CarRule([]CarRule{*NewCarRuleBuilder().RuleId("R001").Build()}).
		HotelRule([]HotelRule{*NewHotelRuleBuilder().RuleId("H001").Build()}).
		FlightRule([]FlightRule{*NewFlightRuleBuilder().RuleId("F001").Build()}).
		TrainRule([]TrainRule{*NewTrainRuleBuilder().RuleId("T001").Build()}).
		BudgetCenterList([]BudgetCenterListItem{*NewBudgetCenterListItemBuilder().Sequence(1).Build()}).
		Build()
	if info.ApprovalId == nil || *info.ApprovalId != "AP001" {
		t.Errorf("ApprovalId = %v, want AP001", info.ApprovalId)
	}
	if info.ApprovalType == nil || *info.ApprovalType != 1 {
		t.Errorf("ApprovalType = %v, want 1", info.ApprovalType)
	}
	if len(info.PassengerList) != 1 {
		t.Errorf("PassengerList len = %d, want 1", len(info.PassengerList))
	}
	if info.TravelBudget == nil {
		t.Error("TravelBudget is nil")
	}
	info2 := NewApprovalDetailBuilder().ApprovalId("AP002").Build()
	if info2.OutApprovalId != nil {
		t.Errorf("OutApprovalId = %v, want nil", info2.OutApprovalId)
	}
}

func TestApprovalOrderRecordBuilder(t *testing.T) {
	info := NewApprovalOrderRecordBuilder().
		OrderId("1125922289295589").ApprovalId("AP001").OutApprovalId("TA_001").
		RuleId("R001").RegulationId("RG001").SceneType("2").
		OrderCreateTime("1675602713").BeginChargeTime("1675582413").
		FinishTime("1675583857").DepartureTime("1675582103").
		UseCarType(2).CarLevel(100).CityName("北京").
		StartName("望京SOHO").EndName("陆家嘴").
		ActualStartName("望京SOHO实际").ActualEndName("陆家嘴实际").
		ActualFlat("39.99").ActualFlng("116.48").
		ActualTlat("31.24").ActualTlng("121.50").
		PayTime("1675600000").OrderStatus(2).PayType(0).IsInvoice(0).
		CallPhone("13800000001").PassengerPhone("13800000002").
		TotalPrice("100.50").ActualPrice(80.00).RefundPrice("0").
		CompanyPay("80.00").PersonalPay("20.50").
		CompanyRealPay("80.00").PersonalRealPay("20.50").
		CompanyRefund("0").PersonalRefund("0").
		BudgetCenterId("1125920020961744").ExtraInfo("{}").
		SceneId("102").PreTotalFee("100.00").
		FixedDiscountFee("10.00").DiscountFee("10.00").
		BudgetCenterList([]BudgetCenterListItem{*NewBudgetCenterListItemBuilder().Sequence(1).Build()}).
		Build()
	if info.OrderId == nil || *info.OrderId != "1125922289295589" {
		t.Errorf("OrderId = %v, want 1125922289295589", info.OrderId)
	}
	if info.UseCarType == nil || *info.UseCarType != 2 {
		t.Errorf("UseCarType = %v, want 2", info.UseCarType)
	}
	if info.ActualPrice == nil || *info.ActualPrice != 80.00 {
		t.Errorf("ActualPrice = %v, want 80.00", info.ActualPrice)
	}
	if len(info.BudgetCenterList) != 1 {
		t.Errorf("BudgetCenterList len = %d, want 1", len(info.BudgetCenterList))
	}
	info2 := NewApprovalOrderRecordBuilder().OrderId("O002").Build()
	if info2.ApprovalId != nil {
		t.Errorf("ApprovalId = %v, want nil", info2.ApprovalId)
	}
}

// =====================================================================
// POST 型 Request Builder 测试
// =====================================================================

func TestApprovalPassRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewApprovalPassRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp(1583484681).Sign("test_sign").
			ObjectType(2).ObjectId(1125922289295589).ObjectApprovalType(21).
			IsPass(1).CurApprover(`{"type":"phone","value":"13800000001"}`).
			CurApproverObj(*NewCurAppRoverBuilder().Type("phone").Value("13800000001").Build()).
			Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId = %v, want test_client", req.ClientId)
		}
		if req.Timestamp == nil || *req.Timestamp != 1583484681 {
			t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
		}
		if req.ObjectType == nil || *req.ObjectType != 2 {
			t.Errorf("ObjectType = %v, want 2", req.ObjectType)
		}
		if req.ObjectId == nil || *req.ObjectId != 1125922289295589 {
			t.Errorf("ObjectId = %v, want 1125922289295589", req.ObjectId)
		}
		if req.IsPass == nil || *req.IsPass != 1 {
			t.Errorf("IsPass = %v, want 1", req.IsPass)
		}
		if req.CurApproverObj == nil {
			t.Error("CurApproverObj is nil")
		}
	})
	t.Run("PartialParams", func(t *testing.T) {
		req := NewApprovalPassRequestBuilder().
			ClientId("test_client").ObjectId(100).Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId = %v, want test_client", req.ClientId)
		}
		if req.ObjectId == nil || *req.ObjectId != 100 {
			t.Errorf("ObjectId = %v, want 100", req.ObjectId)
		}
		if req.AccessToken != nil {
			t.Errorf("AccessToken = %v, want nil", req.AccessToken)
		}
		if req.IsPass != nil {
			t.Errorf("IsPass = %v, want nil", req.IsPass)
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewApprovalPassRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp(1583484681).Sign("test_sign").Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId mismatch")
		}
		if req.ObjectType != nil {
			t.Errorf("ObjectType = %v, want nil", req.ObjectType)
		}
	})
}

func TestCancelApprovalRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewCancelApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			ApprovalId("AP001").IsForce(1).OutApprovalId("TA_001").Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId = %v, want test_client", req.ClientId)
		}
		if req.ApprovalId == nil || *req.ApprovalId != "AP001" {
			t.Errorf("ApprovalId = %v, want AP001", req.ApprovalId)
		}
		if req.IsForce == nil || *req.IsForce != 1 {
			t.Errorf("IsForce = %v, want 1", req.IsForce)
		}
		if req.OutApprovalId == nil || *req.OutApprovalId != "TA_001" {
			t.Errorf("OutApprovalId = %v, want TA_001", req.OutApprovalId)
		}
	})
	t.Run("PartialParams", func(t *testing.T) {
		req := NewCancelApprovalRequestBuilder().
			ClientId("test_client").ApprovalId("AP001").Build()
		if req.ApprovalId == nil || *req.ApprovalId != "AP001" {
			t.Errorf("ApprovalId mismatch")
		}
		if req.IsForce != nil {
			t.Errorf("IsForce = %v, want nil", req.IsForce)
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewCancelApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").Build()
		if req.ApprovalId != nil {
			t.Errorf("ApprovalId = %v, want nil", req.ApprovalId)
		}
	})
}

func TestCreateApprovalRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewCreateApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			OutTripId("TRIP001").ApprovalType(1).RegulationId("RG001").
			PolicyType(3).PolicyTypeValue("13800000001").
			OutApprovalId("TA_001").BudgetCenterId("1125920020961744").
			Name("研发部").OutBudgetId("DEPT001").Reason("北京出差").
			MemberType(0).Phone("13800000001").EmployeeNumber("D0001").
			Email("zhangsan@example.com").ExecutiveRegulationType(0).
			ExecutiveRegulationId("RG002").ExecutiveRegulationMemberType(0).
			ExecutiveRegulationMember("13800000002").OutTripInfo("{}").
			ExtraInfo("{}").ExtendFieldList("[]").TravelDetail("{}").
			TravelBudget("{}").TravelManagement("{}").PassengerList("[]").
			BudgetCenterList("[]").Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId mismatch")
		}
		if req.ApprovalType == nil || *req.ApprovalType != 1 {
			t.Errorf("ApprovalType = %v, want 1", req.ApprovalType)
		}
		if req.Phone == nil || *req.Phone != "13800000001" {
			t.Errorf("Phone mismatch")
		}
		if req.MemberType == nil || *req.MemberType != 0 {
			t.Errorf("MemberType = %v, want 0", req.MemberType)
		}
	})
	t.Run("PartialParams", func(t *testing.T) {
		req := NewCreateApprovalRequestBuilder().
			ClientId("test_client").ApprovalType(2).Phone("13800000001").Build()
		if req.ApprovalType == nil || *req.ApprovalType != 2 {
			t.Errorf("ApprovalType mismatch")
		}
		if req.Email != nil {
			t.Errorf("Email = %v, want nil", req.Email)
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewCreateApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").Build()
		if req.ApprovalType != nil {
			t.Errorf("ApprovalType = %v, want nil", req.ApprovalType)
		}
		if req.Phone != nil {
			t.Errorf("Phone = %v, want nil", req.Phone)
		}
	})
}

func TestCreateApprovalBusinessByTimesRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewCreateApprovalBusinessByTimesRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			OutTripId("TRIP001").ApprovalType(2).RegulationId("RG001").
			PolicyType(3).PolicyTypeValue("13800000001").
			OutApprovalId("TA_001").BudgetCenterId("1125920020961744").
			Name("研发部").OutBudgetId("DEPT001").Reason("北京出差").
			MemberType(0).Phone("13800000001").EmployeeNumber("D0001").
			Email("zhangsan@example.com").ExecutiveRegulationType(0).
			ExecutiveRegulationId("RG002").ExecutiveRegulationMemberType(0).
			ExecutiveRegulationMember("13800000002").OutTripInfo("{}").
			ExtraInfo("{}").ExtendFieldList("[]").
			BusinessTripDetail("{}").TravelBudget("{}").
			TravelManagement("{}").PassengerList("[]").BudgetCenterList("[]").Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId mismatch")
		}
		if req.ApprovalType == nil || *req.ApprovalType != 2 {
			t.Errorf("ApprovalType = %v, want 2", req.ApprovalType)
		}
		if req.BusinessTripDetail == nil || *req.BusinessTripDetail != "{}" {
			t.Errorf("BusinessTripDetail mismatch")
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewCreateApprovalBusinessByTimesRequestBuilder().
			ClientId("test_client").Sign("test_sign").Build()
		if req.BusinessTripDetail != nil {
			t.Errorf("BusinessTripDetail = %v, want nil", req.BusinessTripDetail)
		}
	})
}

func TestCreateApprovalBusinessByDateRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewCreateApprovalBusinessByDateRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			OutTripId("TRIP001").ApprovalType(3).RegulationId("RG001").
			PolicyType(3).PolicyTypeValue("13800000001").
			OutApprovalId("TA_001").BudgetCenterId("1125920020961744").
			Name("研发部").OutBudgetId("DEPT001").Reason("北京出差").
			MemberType(0).Phone("13800000001").EmployeeNumber("D0001").
			Email("zhangsan@example.com").ExecutiveRegulationType(0).
			ExecutiveRegulationId("RG002").ExecutiveRegulationMemberType(0).
			ExecutiveRegulationMember("13800000002").OutTripInfo("{}").
			ExtraInfo("{}").ExtendFieldList("[]").
			BusinessTripDetail("{}").TravelBudget("{}").
			TravelManagement("{}").PassengerList("[]").BudgetCenterList("[]").Build()
		if req.ApprovalType == nil || *req.ApprovalType != 3 {
			t.Errorf("ApprovalType = %v, want 3", req.ApprovalType)
		}
		if req.BusinessTripDetail == nil || *req.BusinessTripDetail != "{}" {
			t.Errorf("BusinessTripDetail mismatch")
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewCreateApprovalBusinessByDateRequestBuilder().
			ClientId("test_client").Sign("test_sign").Build()
		if req.BusinessTripDetail != nil {
			t.Errorf("BusinessTripDetail = %v, want nil", req.BusinessTripDetail)
		}
	})
}

func TestUpdateApprovalRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewUpdateApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			ApprovalType(1).ApprovalId("AP001").OutApprovalId("TA_001").
			Reason("北京出差").Name("研发部").OutBudgetId("DEPT001").
			TravelBudget("{}").TravelManagement("{}").ExtraInfo("{}").
			TravelDetail("{}").PassengerList("[]").BudgetCenterList("[]").
			ExtendFieldList("[]").BudgetCenterId("1125920020961744").Build()
		if req.ClientId == nil || *req.ClientId != "test_client" {
			t.Errorf("ClientId mismatch")
		}
		if req.ApprovalType == nil || *req.ApprovalType != 1 {
			t.Errorf("ApprovalType = %v, want 1", req.ApprovalType)
		}
		if req.ApprovalId == nil || *req.ApprovalId != "AP001" {
			t.Errorf("ApprovalId mismatch")
		}
		if req.BudgetCenterId == nil || *req.BudgetCenterId != "1125920020961744" {
			t.Errorf("BudgetCenterId mismatch")
		}
	})
	t.Run("PartialParams", func(t *testing.T) {
		req := NewUpdateApprovalRequestBuilder().
			ClientId("test_client").ApprovalId("AP001").Reason("修改原因").Build()
		if req.ApprovalId == nil || *req.ApprovalId != "AP001" {
			t.Errorf("ApprovalId mismatch")
		}
		if req.OutApprovalId != nil {
			t.Errorf("OutApprovalId = %v, want nil", req.OutApprovalId)
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewUpdateApprovalRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").Build()
		if req.ApprovalId != nil {
			t.Errorf("ApprovalId = %v, want nil", req.ApprovalId)
		}
	})
}

func TestUpdateApprovalBusinessByTimesRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewUpdateApprovalBusinessByTimesRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			ApprovalType(2).ApprovalId("AP001").OutApprovalId("TA_001").
			Reason("北京出差").Name("研发部").OutBudgetId("DEPT001").
			TravelBudget("{}").TravelManagement("{}").ExtraInfo("{}").
			BusinessTripDetail("{}").PassengerList("[]").BudgetCenterList("[]").
			ExtendFieldList("[]").BudgetCenterId("1125920020961744").Build()
		if req.ApprovalType == nil || *req.ApprovalType != 2 {
			t.Errorf("ApprovalType = %v, want 2", req.ApprovalType)
		}
		if req.BusinessTripDetail == nil || *req.BusinessTripDetail != "{}" {
			t.Errorf("BusinessTripDetail mismatch")
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewUpdateApprovalBusinessByTimesRequestBuilder().
			ClientId("test_client").Sign("test_sign").Build()
		if req.BusinessTripDetail != nil {
			t.Errorf("BusinessTripDetail = %v, want nil", req.BusinessTripDetail)
		}
	})
}

func TestUpdateApprovalBusinessByDateRequestBuilder(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewUpdateApprovalBusinessByDateRequestBuilder().
			ClientId("test_client").AccessToken("test_token").
			Timestamp(1583484681).CompanyId("test_company").Sign("test_sign").
			ApprovalType(3).ApprovalId("AP001").OutApprovalId("TA_001").
			Reason("北京出差").Name("研发部").OutBudgetId("DEPT001").
			TravelBudget("{}").TravelManagement("{}").ExtraInfo("{}").
			BusinessTripDetail("{}").PassengerList("[]").BudgetCenterList("[]").
			ExtendFieldList("[]").BudgetCenterId("1125920020961744").Build()
		if req.ApprovalType == nil || *req.ApprovalType != 3 {
			t.Errorf("ApprovalType = %v, want 3", req.ApprovalType)
		}
		if req.BusinessTripDetail == nil || *req.BusinessTripDetail != "{}" {
			t.Errorf("BusinessTripDetail mismatch")
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewUpdateApprovalBusinessByDateRequestBuilder().
			ClientId("test_client").Sign("test_sign").Build()
		if req.BusinessTripDetail != nil {
			t.Errorf("BusinessTripDetail = %v, want nil", req.BusinessTripDetail)
		}
	})
}

// =====================================================================
// GET 型 ApiReqBuilder QueryParams 测试
// =====================================================================

func TestGetApprovalDetailApiReqBuilder_QueryParams(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewGetApprovalDetailApiReqBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp("1583484681").Sign("test_sign").
			ApprovalId("AP001").OutApprovalId("TA_001").Build()
		tests := []struct{ key, want string }{
			{"client_id", "test_client"},
			{"access_token", "test_token"},
			{"company_id", "test_company"},
			{"timestamp", "1583484681"},
			{"sign", "test_sign"},
			{"approval_id", "AP001"},
			{"out_approval_id", "TA_001"},
		}
		for _, tt := range tests {
			if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
				t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
			}
		}
	})
	t.Run("PartialParams", func(t *testing.T) {
		req := NewGetApprovalDetailApiReqBuilder().
			ClientId("test_client").ApprovalId("AP001").Build()
		if req.apiReq.QueryParams.Get("client_id") != "test_client" {
			t.Errorf("client_id mismatch")
		}
		if req.apiReq.QueryParams.Get("approval_id") != "AP001" {
			t.Errorf("approval_id mismatch")
		}
		if req.apiReq.QueryParams.Get("out_approval_id") != "" {
			t.Errorf("out_approval_id should be empty")
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewGetApprovalDetailApiReqBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp("1583484681").Sign("test_sign").Build()
		for _, key := range []string{"approval_id", "out_approval_id"} {
			if v := req.apiReq.QueryParams.Get(key); v != "" {
				t.Errorf("QueryParams[%s] = %q, want empty", key, v)
			}
		}
	})
}

func TestListApprovalOrderApiReqBuilder_QueryParams(t *testing.T) {
	t.Run("AllParams", func(t *testing.T) {
		req := NewListApprovalOrderApiReqBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp("1583484681").Sign("test_sign").
			ApprovalId("AP001").Offset(0).Length(10).Build()
		tests := []struct{ key, want string }{
			{"client_id", "test_client"},
			{"access_token", "test_token"},
			{"company_id", "test_company"},
			{"timestamp", "1583484681"},
			{"sign", "test_sign"},
			{"approval_id", "AP001"},
			{"offset", "0"},
			{"length", "10"},
		}
		for _, tt := range tests {
			if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
				t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
			}
		}
	})
	t.Run("ZeroIntValues", func(t *testing.T) {
		req := NewListApprovalOrderApiReqBuilder().
			ClientId("test_client").Offset(0).Length(0).Build()
		if req.apiReq.QueryParams.Get("offset") != "0" {
			t.Errorf("offset = %q, want 0", req.apiReq.QueryParams.Get("offset"))
		}
		if req.apiReq.QueryParams.Get("length") != "0" {
			t.Errorf("length = %q, want 0", req.apiReq.QueryParams.Get("length"))
		}
	})
	t.Run("OnlyCommonParams", func(t *testing.T) {
		req := NewListApprovalOrderApiReqBuilder().
			ClientId("test_client").AccessToken("test_token").
			CompanyId("test_company").Timestamp("1583484681").Sign("test_sign").Build()
		for _, key := range []string{"approval_id", "offset", "length"} {
			if v := req.apiReq.QueryParams.Get(key); v != "" {
				t.Errorf("QueryParams[%s] = %q, want empty", key, v)
			}
		}
	})
}

// =====================================================================
// Reply 反序列化测试
// =====================================================================

func TestApprovalPassApiReply_Deserialization(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001","error_info_list":[{"err_no":0,"err_msg":"ok"}]},"request_id":"req_001"}`
		var reply ApprovalPassApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno == nil || *reply.Errno != 0 {
			t.Errorf("Errno = %v, want 0", reply.Errno)
		}
		if reply.Data == nil {
			t.Fatal("Data is nil")
		}
		if reply.Data.ApprovalId == nil || *reply.Data.ApprovalId != "AP001" {
			t.Errorf("ApprovalId = %v, want AP001", reply.Data.ApprovalId)
		}
		if len(reply.Data.ErrorInfoList) != 1 {
			t.Fatalf("ErrorInfoList len = %d, want 1", len(reply.Data.ErrorInfoList))
		}
		if reply.Data.ErrorInfoList[0].ErrNo == nil || *reply.Data.ErrorInfoList[0].ErrNo != 0 {
			t.Errorf("ErrorInfoList[0].ErrNo = %v, want 0", reply.Data.ErrorInfoList[0].ErrNo)
		}
	})
	t.Run("ApiError", func(t *testing.T) {
		jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
		var reply ApprovalPassApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno == nil || *reply.Errno != 10003 {
			t.Errorf("Errno = %v, want 10003", reply.Errno)
		}
		if reply.Data != nil {
			t.Errorf("Data = %v, want nil", reply.Data)
		}
	})
	t.Run("EmptyData", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`
		var reply ApprovalPassApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Data != nil {
			t.Errorf("Data = %v, want nil", reply.Data)
		}
	})
	t.Run("MissingDataField", func(t *testing.T) {
		jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`
		var reply ApprovalPassApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno == nil || *reply.Errno != 10003 {
			t.Errorf("Errno = %v, want 10003", reply.Errno)
		}
		if reply.Data != nil {
			t.Errorf("Data = %v, want nil", reply.Data)
		}
	})
}

func TestCreateApprovalApiReply_Deserialization(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001"},"request_id":"req_001"}`
		var reply CreateApprovalApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 0 {
			t.Errorf("Errno = %d, want 0", reply.Errno)
		}
		if reply.Data.ApprovalId == nil || *reply.Data.ApprovalId != "AP001" {
			t.Errorf("ApprovalId = %v, want AP001", reply.Data.ApprovalId)
		}
	})
	t.Run("ApiError", func(t *testing.T) {
		jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
		var reply CreateApprovalApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 10003 {
			t.Errorf("Errno = %d, want 10003", reply.Errno)
		}
	})
}

func TestGetApprovalDetailApiReply_Deserialization(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001","out_approval_id":"TA_001","approval_type":1,"type":0,"budget_center_id":"1125920020961744","reason":"北京出差","approval_status":1,"scene_id":1,"city_type":1,"passenger_list":[{"name":"张三","phone":"13800000001","passenger_type":0}],"traveler_list":[{"name":"李四","phone":"13800000002"}],"applicant":{"name":"申请人","phone":"13800000003"},"car_rule":[{"rule_id":"R001","rule_name":"市内用车","rule_status":1}],"hotel_rule":[{"rule_id":"H001","rule_status":"1"}],"budget_center_list":[{"sequence":1,"id":"1125920020961744"}]},"request_id":"req_001"}`
		var reply GetApprovalDetailApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 0 {
			t.Errorf("Errno = %d, want 0", reply.Errno)
		}
		if reply.Data.ApprovalId == nil || *reply.Data.ApprovalId != "AP001" {
			t.Errorf("ApprovalId = %v, want AP001", reply.Data.ApprovalId)
		}
		if reply.Data.ApprovalType == nil || *reply.Data.ApprovalType != 1 {
			t.Errorf("ApprovalType = %v, want 1", reply.Data.ApprovalType)
		}
		if len(reply.Data.PassengerList) != 1 {
			t.Fatalf("PassengerList len = %d, want 1", len(reply.Data.PassengerList))
		}
		if reply.Data.PassengerList[0].Name == nil || *reply.Data.PassengerList[0].Name != "张三" {
			t.Errorf("PassengerList[0].Name = %v, want 张三", reply.Data.PassengerList[0].Name)
		}
		if reply.Data.PassengerList[0].PassengerType == nil || *reply.Data.PassengerList[0].PassengerType != 0 {
			t.Errorf("PassengerList[0].PassengerType = %v, want 0", reply.Data.PassengerList[0].PassengerType)
		}
		if len(reply.Data.CarRule) != 1 {
			t.Fatalf("CarRule len = %d, want 1", len(reply.Data.CarRule))
		}
		if reply.Data.CarRule[0].RuleStatus == nil || *reply.Data.CarRule[0].RuleStatus != 1 {
			t.Errorf("CarRule[0].RuleStatus = %v, want 1", reply.Data.CarRule[0].RuleStatus)
		}
		if len(reply.Data.BudgetCenterList) != 1 {
			t.Fatalf("BudgetCenterList len = %d, want 1", len(reply.Data.BudgetCenterList))
		}
	})
	t.Run("PartialFields", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP002","approval_status":1},"request_id":"req_partial"}`
		var reply GetApprovalDetailApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Data.ApprovalStatus == nil || *reply.Data.ApprovalStatus != 1 {
			t.Errorf("ApprovalStatus = %v, want 1", reply.Data.ApprovalStatus)
		}
		if reply.Data.Reason != nil {
			t.Errorf("Reason = %v, want nil", reply.Data.Reason)
		}
		if reply.Data.PassengerList != nil {
			t.Errorf("PassengerList = %v, want nil", reply.Data.PassengerList)
		}
	})
	t.Run("ApiError", func(t *testing.T) {
		jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
		var reply GetApprovalDetailApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 10003 {
			t.Errorf("Errno = %d, want 10003", reply.Errno)
		}
	})
}

func TestListApprovalOrderApiReply_Deserialization(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"order_id":"1125922289295589","approval_id":"AP001","use_car_type":2,"car_level":100,"order_status":2,"actual_price":80.00,"total_price":"100.50","city_name":"北京","start_name":"望京SOHO","end_name":"陆家嘴","budget_center_list":[{"sequence":1}]}]},"request_id":"req_001"}`
		var reply ListApprovalOrderApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 0 {
			t.Errorf("Errno = %d, want 0", reply.Errno)
		}
		if len(reply.Data.Records) != 1 {
			t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
		}
		rec := reply.Data.Records[0]
		if rec.OrderId == nil || *rec.OrderId != "1125922289295589" {
			t.Errorf("OrderId = %v, want 1125922289295589", rec.OrderId)
		}
		if rec.UseCarType == nil || *rec.UseCarType != 2 {
			t.Errorf("UseCarType = %v, want 2", rec.UseCarType)
		}
		if rec.ActualPrice == nil || *rec.ActualPrice != 80.00 {
			t.Errorf("ActualPrice = %v, want 80.00", rec.ActualPrice)
		}
		if rec.TotalPrice == nil || *rec.TotalPrice != "100.50" {
			t.Errorf("TotalPrice = %v, want 100.50", rec.TotalPrice)
		}
		if len(rec.BudgetCenterList) != 1 {
			t.Errorf("BudgetCenterList len = %d, want 1", len(rec.BudgetCenterList))
		}
	})
	t.Run("MultipleItems", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"order_id":"O001","order_status":2},{"order_id":"O002"}]},"request_id":"req_multi"}`
		var reply ListApprovalOrderApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if len(reply.Data.Records) != 2 {
			t.Fatalf("Records len = %d, want 2", len(reply.Data.Records))
		}
		if reply.Data.Records[0].OrderStatus == nil || *reply.Data.Records[0].OrderStatus != 2 {
			t.Errorf("Records[0].OrderStatus = %v, want 2", reply.Data.Records[0].OrderStatus)
		}
		if reply.Data.Records[1].OrderStatus != nil {
			t.Errorf("Records[1].OrderStatus = %v, want nil", reply.Data.Records[1].OrderStatus)
		}
	})
	t.Run("EmptyData", func(t *testing.T) {
		jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"records":[]},"request_id":"req_empty"}`
		var reply ListApprovalOrderApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if len(reply.Data.Records) != 0 {
			t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
		}
	})
	t.Run("ApiError", func(t *testing.T) {
		jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
		var reply ListApprovalOrderApiReply
		if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if reply.Errno != 10003 {
			t.Errorf("Errno = %d, want 10003", reply.Errno)
		}
	})
}

// =====================================================================
// 资源方法表驱动测试（10 方法 × 8 场景）
// =====================================================================

type approvalMethod struct {
	name          string
	httpMethod    string
	path          string
	buildReq      func() interface{}
	invoke        func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error)
	hasReply      func(resp interface{}) bool
	successBody   string
	emptyBody     string
	verifySuccess func(t *testing.T, resp interface{})
}

func runApprovalScenarios(t *testing.T, m approvalMethod) {
	t.Helper()

	t.Run("Success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m.httpMethod {
				t.Errorf("expected %s, got %s", m.httpMethod, r.Method)
			}
			if r.URL.Path != m.path {
				t.Errorf("expected path %s, got %s", m.path, r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		a := &approval{option: newApprovalTestOption(ts.URL)}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
		if m.verifySuccess != nil {
			m.verifySuccess(t, resp)
		}
	})

	t.Run("EmptyData", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.emptyBody))
		}))
		defer ts.Close()
		a := &approval{option: newApprovalTestOption(ts.URL)}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("ApiError", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_err"}`))
		}))
		defer ts.Close()
		a := &approval{option: newApprovalTestOption(ts.URL)}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("HttpError", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()
		a := &approval{option: newApprovalTestOption(ts.URL)}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if m.hasReply(resp) {
			t.Errorf("reply should be nil for non-200 response")
		}
	})

	t.Run("WithEncryption_AES128", func(t *testing.T) {
		plaintext := m.successBody
		key := []byte("16byte-key-12345")
		encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
		if err != nil {
			t.Fatalf("AESEncryptECB() error = %v", err)
		}
		encryptData := base64.StdEncoding.EncodeToString(encrypted)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
		}))
		defer ts.Close()
		opt := newApprovalTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
		a := &approval{option: opt}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("decryption failed: reply is nil")
		}
	})

	t.Run("WithEncryption_AES256", func(t *testing.T) {
		plaintext := m.successBody
		key := []byte("32byte-key-1234567890abcdefghijk")
		encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
		if err != nil {
			t.Fatalf("AESEncryptECB() error = %v", err)
		}
		encryptData := base64.URLEncoding.EncodeToString(encrypted)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
		}))
		defer ts.Close()
		opt := newApprovalTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
		a := &approval{option: opt}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("decryption failed: reply is nil")
		}
	})

	t.Run("EncryptionNoEncryptData", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		opt := newApprovalTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
		a := &approval{option: opt}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("WithReqOption", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
				t.Errorf("expected X-Custom-Header, got %s", h)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		a := &approval{option: newApprovalTestOption(ts.URL)}
		customHeader := http.Header{}
		customHeader.Set("X-Custom-Header", "custom-value")
		reqOption := &core.ReqOption{Header: customHeader}
		resp, err := m.invoke(a, context.Background(), m.buildReq(), reqOption)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})
}

// mock JSON 常量
const approvalPassSuccessBody = `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001","error_info_list":[{"err_no":0,"err_msg":"ok"}]},"request_id":"req_001"}`
const approvalPassEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`
const cancelApprovalBody = `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`
const createApprovalSuccessBody = `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001"},"request_id":"req_001"}`
const createApprovalEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":""},"request_id":"req_empty"}`
const getApprovalDetailSuccessBody = `{"errno":0,"errmsg":"SUCCESS","data":{"approval_id":"AP001","out_approval_id":"TA_001","approval_type":1,"approval_status":1,"passenger_list":[{"name":"张三","phone":"13800000001","passenger_type":0}],"car_rule":[{"rule_id":"R001","rule_status":1}],"budget_center_list":[{"sequence":1,"id":"1125920020961744"}]},"request_id":"req_001"}`
const getApprovalDetailEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":{},"request_id":"req_empty"}`
const listApprovalOrderSuccessBody = `{"errno":0,"errmsg":"SUCCESS","data":{"records":[{"order_id":"1125922289295589","approval_id":"AP001","use_car_type":2,"order_status":2,"actual_price":80.00,"total_price":"100.50","budget_center_list":[{"sequence":1}]}]},"request_id":"req_001"}`
const listApprovalOrderEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":{"records":[]},"request_id":"req_empty"}`
const updateApprovalBody = `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_001"}`

func TestApprovalPass_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/pass",
		buildReq: func() interface{} {
			return NewApprovalPassApiReqBuilder().
				ApprovalPassRequest(NewApprovalPassRequestBuilder().ClientId("test_client").Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.ApprovalPass(ctx, req.(*ApprovalPassApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*ApprovalPassApiResp).ApprovalPassApiReply != nil },
		successBody: approvalPassSuccessBody,
		emptyBody:   approvalPassEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*ApprovalPassApiResp).ApprovalPassApiReply
			if r.Errno == nil || *r.Errno != 0 {
				t.Errorf("Errno = %v, want 0", r.Errno)
			}
			if r.Data == nil || r.Data.ApprovalId == nil || *r.Data.ApprovalId != "AP001" {
				t.Errorf("Data.ApprovalId mismatch")
			}
		},
	})
}

func TestCancelApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/cancel",
		buildReq: func() interface{} {
			return NewCancelApprovalApiReqBuilder().
				CancelApprovalRequest(NewCancelApprovalRequestBuilder().ClientId("test_client").ApprovalId("AP001").Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.CancelApproval(ctx, req.(*CancelApprovalApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*CancelApprovalApiResp).CancelApprovalApiReply != nil },
		successBody: cancelApprovalBody,
		emptyBody:   cancelApprovalBody,
	})
}

func TestCreateTravelApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/create",
		buildReq: func() interface{} {
			return NewCreateTravelApprovalApiReqBuilder().
				CreateTravelApprovalRequest(NewCreateApprovalRequestBuilder().ClientId("test_client").ApprovalType(1).Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.CreateTravelApproval(ctx, req.(*CreateTravelApprovalApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*CreateTravelApprovalApiResp).CreateApprovalApiReply != nil },
		successBody: createApprovalSuccessBody,
		emptyBody:   createApprovalEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*CreateTravelApprovalApiResp).CreateApprovalApiReply
			if r.Errno != 0 {
				t.Errorf("Errno = %d, want 0", r.Errno)
			}
			if r.Data.ApprovalId == nil || *r.Data.ApprovalId != "AP001" {
				t.Errorf("Data.ApprovalId mismatch")
			}
		},
	})
}

func TestCreateBusinessByTimesApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/create",
		buildReq: func() interface{} {
			return NewCreateBusinessByTimesApprovalApiReqBuilder().
				CreateBusinessByTimesApprovalRequest(NewCreateApprovalBusinessByTimesRequestBuilder().ClientId("test_client").ApprovalType(2).Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.CreateBusinessByTimesApproval(ctx, req.(*CreateBusinessByTimesApprovalApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*CreateBusinessByTimesApprovalApiResp).CreateApprovalApiReply != nil
		},
		successBody: createApprovalSuccessBody,
		emptyBody:   createApprovalEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*CreateBusinessByTimesApprovalApiResp).CreateApprovalApiReply
			if r.Data.ApprovalId == nil || *r.Data.ApprovalId != "AP001" {
				t.Errorf("Data.ApprovalId mismatch")
			}
		},
	})
}

func TestCreateBusinessByDateApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/create",
		buildReq: func() interface{} {
			return NewCreateBusinessByDateApprovalApiReqBuilder().
				CreateBusinessByDateApprovalRequest(NewCreateApprovalBusinessByDateRequestBuilder().ClientId("test_client").ApprovalType(3).Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.CreateBusinessByDateApproval(ctx, req.(*CreateBusinessByDateApprovalApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*CreateBusinessByDateApprovalApiResp).CreateApprovalApiReply != nil
		},
		successBody: createApprovalSuccessBody,
		emptyBody:   createApprovalEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*CreateBusinessByDateApprovalApiResp).CreateApprovalApiReply
			if r.Data.ApprovalId == nil || *r.Data.ApprovalId != "AP001" {
				t.Errorf("Data.ApprovalId mismatch")
			}
		},
	})
}

func TestGetApprovalDetail_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodGet,
		path:       "/open-apis/v1/approval/detail",
		buildReq: func() interface{} {
			return NewGetApprovalDetailApiReqBuilder().
				ClientId("test_client").ApprovalId("AP001").Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.GetApprovalDetail(ctx, req.(*GetApprovalDetailApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*GetApprovalDetailApiResp).GetApprovalDetailApiReply != nil },
		successBody: getApprovalDetailSuccessBody,
		emptyBody:   getApprovalDetailEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*GetApprovalDetailApiResp).GetApprovalDetailApiReply
			if r.Errno != 0 {
				t.Errorf("Errno = %d, want 0", r.Errno)
			}
			if r.Data.ApprovalId == nil || *r.Data.ApprovalId != "AP001" {
				t.Errorf("Data.ApprovalId mismatch")
			}
			if r.Data.ApprovalType == nil || *r.Data.ApprovalType != 1 {
				t.Errorf("Data.ApprovalType = %v, want 1", r.Data.ApprovalType)
			}
			if len(r.Data.PassengerList) != 1 {
				t.Errorf("Data.PassengerList len = %d, want 1", len(r.Data.PassengerList))
			}
		},
	})
}

func TestListApprovalOrder_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodGet,
		path:       "/river/Approval/getOrder",
		buildReq: func() interface{} {
			return NewListApprovalOrderApiReqBuilder().
				ClientId("test_client").ApprovalId("AP001").Offset(0).Length(10).Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.ListApprovalOrder(ctx, req.(*ListApprovalOrderApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*ListApprovalOrderApiResp).ListApprovalOrderApiReply != nil },
		successBody: listApprovalOrderSuccessBody,
		emptyBody:   listApprovalOrderEmptyBody,
		verifySuccess: func(t *testing.T, resp interface{}) {
			r := resp.(*ListApprovalOrderApiResp).ListApprovalOrderApiReply
			if r.Errno != 0 {
				t.Errorf("Errno = %d, want 0", r.Errno)
			}
			if len(r.Data.Records) != 1 {
				t.Fatalf("Records len = %d, want 1", len(r.Data.Records))
			}
			rec := r.Data.Records[0]
			if rec.OrderId == nil || *rec.OrderId != "1125922289295589" {
				t.Errorf("OrderId mismatch")
			}
			if rec.UseCarType == nil || *rec.UseCarType != 2 {
				t.Errorf("UseCarType mismatch")
			}
		},
	})
}

func TestUpdateApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/update",
		buildReq: func() interface{} {
			return NewUpdateApprovalApiReqBuilder().
				UpdateApprovalRequest(NewUpdateApprovalRequestBuilder().ClientId("test_client").ApprovalId("AP001").Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.UpdateApproval(ctx, req.(*UpdateApprovalApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*UpdateApprovalApiResp).UpdateApprovalApiReply != nil },
		successBody: updateApprovalBody,
		emptyBody:   updateApprovalBody,
	})
}

func TestUpdateBusinessByTimesApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/update",
		buildReq: func() interface{} {
			return NewUpdateBusinessByTimesApprovalApiReqBuilder().
				UpdateBusinessByTimesApprovalRequest(NewUpdateApprovalBusinessByTimesRequestBuilder().ClientId("test_client").ApprovalId("AP001").Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.UpdateBusinessByTimesApproval(ctx, req.(*UpdateBusinessByTimesApprovalApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*UpdateBusinessByTimesApprovalApiResp).UpdateApprovalApiReply != nil
		},
		successBody: updateApprovalBody,
		emptyBody:   updateApprovalBody,
	})
}

func TestUpdateBusinessByDateApproval_AllScenarios(t *testing.T) {
	runApprovalScenarios(t, approvalMethod{
		httpMethod: http.MethodPost,
		path:       "/river/Approval/update",
		buildReq: func() interface{} {
			return NewUpdateBusinessByDateApprovalApiReqBuilder().
				UpdateBusinessByDateApprovalRequest(NewUpdateApprovalBusinessByDateRequestBuilder().ClientId("test_client").ApprovalId("AP001").Build()).
				Build()
		},
		invoke: func(a *approval, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return a.UpdateBusinessByDateApproval(ctx, req.(*UpdateBusinessByDateApprovalApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*UpdateBusinessByDateApprovalApiResp).UpdateApprovalApiReply != nil
		},
		successBody: updateApprovalBody,
		emptyBody:   updateApprovalBody,
	})
}
