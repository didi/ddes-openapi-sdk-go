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

// --- 共享测试辅助 ---

func newCityTestOption(serverURL string) *core.Option {
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

// --- AirCityFlightStationBuilder ---

func TestAirCityFlightStationBuilder(t *testing.T) {
	info := NewAirCityFlightStationBuilder().
		AirportNameCn("首都国际机场").
		AirportNameEn("Beijing Capital International Airport").
		AirportCode("PEK").
		Build()

	if info.AirportNameCn == nil || *info.AirportNameCn != "首都国际机场" {
		t.Errorf("AirportNameCn = %v, want 首都国际机场", info.AirportNameCn)
	}
	if info.AirportNameEn == nil || *info.AirportNameEn != "Beijing Capital International Airport" {
		t.Errorf("AirportNameEn = %v, want Beijing Capital International Airport", info.AirportNameEn)
	}
	if info.AirportCode == nil || *info.AirportCode != "PEK" {
		t.Errorf("AirportCode = %v, want PEK", info.AirportCode)
	}

	// 部分设置
	info2 := NewAirCityFlightStationBuilder().
		AirportCode("SHA").
		Build()
	if info2.AirportCode == nil || *info2.AirportCode != "SHA" {
		t.Errorf("AirportCode = %v, want SHA", info2.AirportCode)
	}
	if info2.AirportNameCn != nil {
		t.Errorf("AirportNameCn = %v, want nil", info2.AirportNameCn)
	}
}

// --- FlightStationBuilder ---

func TestFlightStationBuilder(t *testing.T) {
	info := NewFlightStationBuilder().
		AirportId(1001).
		AirportNameCn("浦东国际机场").
		AirportNameEn("Pudong International Airport").
		AirportCode("PVG").
		Build()

	if info.AirportId == nil || *info.AirportId != 1001 {
		t.Errorf("AirportId = %v, want 1001", info.AirportId)
	}
	if info.AirportNameCn == nil || *info.AirportNameCn != "浦东国际机场" {
		t.Errorf("AirportNameCn = %v, want 浦东国际机场", info.AirportNameCn)
	}
	if info.AirportNameEn == nil || *info.AirportNameEn != "Pudong International Airport" {
		t.Errorf("AirportNameEn = %v, want Pudong International Airport", info.AirportNameEn)
	}
	if info.AirportCode == nil || *info.AirportCode != "PVG" {
		t.Errorf("AirportCode = %v, want PVG", info.AirportCode)
	}

	// 部分设置
	info2 := NewFlightStationBuilder().
		AirportId(1002).
		Build()
	if info2.AirportId == nil || *info2.AirportId != 1002 {
		t.Errorf("AirportId = %v, want 1002", info2.AirportId)
	}
	if info2.AirportNameCn != nil {
		t.Errorf("AirportNameCn = %v, want nil", info2.AirportNameCn)
	}
}

// --- TrainStationBuilder ---

func TestTrainStationBuilder(t *testing.T) {
	info := NewTrainStationBuilder().
		StationNameCn("北京南站").
		StationNameEn("Beijing South Railway Station").
		StationName("北京南").
		StationId(2001).
		Build()

	if info.StationNameCn == nil || *info.StationNameCn != "北京南站" {
		t.Errorf("StationNameCn = %v, want 北京南站", info.StationNameCn)
	}
	if info.StationNameEn == nil || *info.StationNameEn != "Beijing South Railway Station" {
		t.Errorf("StationNameEn = %v, want Beijing South Railway Station", info.StationNameEn)
	}
	if info.StationName == nil || *info.StationName != "北京南" {
		t.Errorf("StationName = %v, want 北京南", info.StationName)
	}
	if info.StationId == nil || *info.StationId != 2001 {
		t.Errorf("StationId = %v, want 2001", info.StationId)
	}

	// 部分设置
	info2 := NewTrainStationBuilder().
		StationId(2002).
		Build()
	if info2.StationId == nil || *info2.StationId != 2002 {
		t.Errorf("StationId = %v, want 2002", info2.StationId)
	}
	if info2.StationNameCn != nil {
		t.Errorf("StationNameCn = %v, want nil", info2.StationNameCn)
	}
}

// --- CountyRecordBuilder ---

func TestCountyRecordBuilder(t *testing.T) {
	info := NewCountyRecordBuilder().
		CountyId(3001).
		CountyName("朝阳区").
		Build()

	if info.CountyId == nil || *info.CountyId != 3001 {
		t.Errorf("CountyId = %v, want 3001", info.CountyId)
	}
	if info.CountyName == nil || *info.CountyName != "朝阳区" {
		t.Errorf("CountyName = %v, want 朝阳区", info.CountyName)
	}

	// 部分设置
	info2 := NewCountyRecordBuilder().
		CountyName("海淀区").
		Build()
	if info2.CountyName == nil || *info2.CountyName != "海淀区" {
		t.Errorf("CountyName = %v, want 海淀区", info2.CountyName)
	}
	if info2.CountyId != nil {
		t.Errorf("CountyId = %v, want nil", info2.CountyId)
	}
}

// --- CityInfoBuilder ---

func TestCityInfoBuilder(t *testing.T) {
	info := NewCityInfoBuilder().
		CityId(1).
		ProductType([]int32{10, 20, 30}).
		CityNameCn("北京").
		CityShortName("京").
		CityPathId("1-1").
		CityPathCn("北京市 - 北京").
		CityPathEn("Beijing - Beijing").
		CityNameEn("Beijing").
		FlightStation([]FlightStation{
			*NewFlightStationBuilder().AirportId(1001).AirportCode("PEK").Build(),
		}).
		TrainStation([]TrainStation{
			*NewTrainStationBuilder().StationId(2001).StationNameCn("北京南站").Build(),
		}).
		Build()

	if info.CityId == nil || *info.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info.CityId)
	}
	if len(info.ProductType) != 3 || info.ProductType[0] != 10 {
		t.Errorf("ProductType = %v, want [10,20,30]", info.ProductType)
	}
	if info.CityNameCn == nil || *info.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", info.CityNameCn)
	}
	if info.CityShortName == nil || *info.CityShortName != "京" {
		t.Errorf("CityShortName = %v, want 京", info.CityShortName)
	}
	if info.CityPathId == nil || *info.CityPathId != "1-1" {
		t.Errorf("CityPathId = %v, want 1-1", info.CityPathId)
	}
	if info.CityPathCn == nil || *info.CityPathCn != "北京市 - 北京" {
		t.Errorf("CityPathCn = %v, want 北京市 - 北京", info.CityPathCn)
	}
	if info.CityPathEn == nil || *info.CityPathEn != "Beijing - Beijing" {
		t.Errorf("CityPathEn = %v, want Beijing - Beijing", info.CityPathEn)
	}
	if info.CityNameEn == nil || *info.CityNameEn != "Beijing" {
		t.Errorf("CityNameEn = %v, want Beijing", info.CityNameEn)
	}
	if len(info.FlightStation) != 1 {
		t.Errorf("FlightStation len = %d, want 1", len(info.FlightStation))
	}
	if len(info.TrainStation) != 1 {
		t.Errorf("TrainStation len = %d, want 1", len(info.TrainStation))
	}

	// 部分设置
	info2 := NewCityInfoBuilder().
		CityId(2).
		CityNameCn("上海").
		Build()
	if info2.CityId == nil || *info2.CityId != 2 {
		t.Errorf("CityId = %v, want 2", info2.CityId)
	}
	if info2.CityNameCn == nil || *info2.CityNameCn != "上海" {
		t.Errorf("CityNameCn = %v, want 上海", info2.CityNameCn)
	}
	if info2.ProductType != nil {
		t.Errorf("ProductType = %v, want nil", info2.ProductType)
	}
	if info2.FlightStation != nil {
		t.Errorf("FlightStation = %v, want nil", info2.FlightStation)
	}
}

// --- CityRecordBuilder ---

func TestCityRecordBuilder(t *testing.T) {
	info := NewCityRecordBuilder().
		ProvinceId(1).
		ProvinceNameCn("北京市").
		ProvinceNameEn("Beijing").
		CityList([]CityInfo{
			*NewCityInfoBuilder().CityId(1).CityNameCn("北京").Build(),
		}).
		Build()

	if info.ProvinceId == nil || *info.ProvinceId != 1 {
		t.Errorf("ProvinceId = %v, want 1", info.ProvinceId)
	}
	if info.ProvinceNameCn == nil || *info.ProvinceNameCn != "北京市" {
		t.Errorf("ProvinceNameCn = %v, want 北京市", info.ProvinceNameCn)
	}
	if info.ProvinceNameEn == nil || *info.ProvinceNameEn != "Beijing" {
		t.Errorf("ProvinceNameEn = %v, want Beijing", info.ProvinceNameEn)
	}
	if len(info.CityList) != 1 {
		t.Errorf("CityList len = %d, want 1", len(info.CityList))
	}

	// 部分设置
	info2 := NewCityRecordBuilder().
		ProvinceId(2).
		Build()
	if info2.ProvinceId == nil || *info2.ProvinceId != 2 {
		t.Errorf("ProvinceId = %v, want 2", info2.ProvinceId)
	}
	if info2.CityList != nil {
		t.Errorf("CityList = %v, want nil", info2.CityList)
	}
}

// --- ListCityParamObjBuilder ---

func TestListCityParamObjBuilder(t *testing.T) {
	info := NewListCityParamObjBuilder().
		CountryId(1).
		ProvinceId(1).
		CityId(1).
		ProductType("10,20,30").
		Build()

	if info.CountryId == nil || *info.CountryId != 1 {
		t.Errorf("CountryId = %v, want 1", info.CountryId)
	}
	if info.ProvinceId == nil || *info.ProvinceId != 1 {
		t.Errorf("ProvinceId = %v, want 1", info.ProvinceId)
	}
	if info.CityId == nil || *info.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info.CityId)
	}
	if info.ProductType == nil || *info.ProductType != "10,20,30" {
		t.Errorf("ProductType = %v, want 10,20,30", info.ProductType)
	}

	// 部分设置
	info2 := NewListCityParamObjBuilder().
		CountryId(2).
		Build()
	if info2.CountryId == nil || *info2.CountryId != 2 {
		t.Errorf("CountryId = %v, want 2", info2.CountryId)
	}
	if info2.CityId != nil {
		t.Errorf("CityId = %v, want nil", info2.CityId)
	}
}

// --- HotelCityInfoBuilder ---

func TestHotelCityInfoBuilder(t *testing.T) {
	info := NewHotelCityInfoBuilder().
		CityId("1001").
		DescriptionCn("北京城市描述").
		DescriptionEn("Beijing city description").
		CityNameCn("北京").
		CityNameEn("Beijing").
		ProvinceId("1").
		ProvinceNameCn("北京市").
		ProvinceNameEn("Beijing").
		CountryId("1").
		CanonicalCountryCode("CN").
		CountryCode("CHN").
		CountryNameCn("中国").
		CountryNameEn("China").
		Build()

	if info.CityId == nil || *info.CityId != "1001" {
		t.Errorf("CityId = %v, want 1001", info.CityId)
	}
	if info.DescriptionCn == nil || *info.DescriptionCn != "北京城市描述" {
		t.Errorf("DescriptionCn = %v, want 北京城市描述", info.DescriptionCn)
	}
	if info.DescriptionEn == nil || *info.DescriptionEn != "Beijing city description" {
		t.Errorf("DescriptionEn = %v, want Beijing city description", info.DescriptionEn)
	}
	if info.CityNameCn == nil || *info.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", info.CityNameCn)
	}
	if info.CityNameEn == nil || *info.CityNameEn != "Beijing" {
		t.Errorf("CityNameEn = %v, want Beijing", info.CityNameEn)
	}
	if info.ProvinceId == nil || *info.ProvinceId != "1" {
		t.Errorf("ProvinceId = %v, want 1", info.ProvinceId)
	}
	if info.ProvinceNameCn == nil || *info.ProvinceNameCn != "北京市" {
		t.Errorf("ProvinceNameCn = %v, want 北京市", info.ProvinceNameCn)
	}
	if info.ProvinceNameEn == nil || *info.ProvinceNameEn != "Beijing" {
		t.Errorf("ProvinceNameEn = %v, want Beijing", info.ProvinceNameEn)
	}
	if info.CountryId == nil || *info.CountryId != "1" {
		t.Errorf("CountryId = %v, want 1", info.CountryId)
	}
	if info.CanonicalCountryCode == nil || *info.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", info.CanonicalCountryCode)
	}
	if info.CountryCode == nil || *info.CountryCode != "CHN" {
		t.Errorf("CountryCode = %v, want CHN", info.CountryCode)
	}
	if info.CountryNameCn == nil || *info.CountryNameCn != "中国" {
		t.Errorf("CountryNameCn = %v, want 中国", info.CountryNameCn)
	}
	if info.CountryNameEn == nil || *info.CountryNameEn != "China" {
		t.Errorf("CountryNameEn = %v, want China", info.CountryNameEn)
	}

	// 部分设置
	info2 := NewHotelCityInfoBuilder().
		CityId("1002").
		CityNameCn("上海").
		Build()
	if info2.CityId == nil || *info2.CityId != "1002" {
		t.Errorf("CityId = %v, want 1002", info2.CityId)
	}
	if info2.CityNameCn == nil || *info2.CityNameCn != "上海" {
		t.Errorf("CityNameCn = %v, want 上海", info2.CityNameCn)
	}
	if info2.CountryId != nil {
		t.Errorf("CountryId = %v, want nil", info2.CountryId)
	}
}

// --- TrainCityInfoBuilder ---

func TestTrainCityInfoBuilder(t *testing.T) {
	info := NewTrainCityInfoBuilder().
		CityId("2001").
		CityNameCn("北京").
		CityNameEn("Beijing").
		ProvinceId("1").
		ProvinceNameCn("北京市").
		ProvinceNameEn("Beijing").
		CountryId("1").
		CanonicalCountryCode("CN").
		CountryCode("CHN").
		CountryNameCn("中国").
		TrainStation([]TrainStation{
			*NewTrainStationBuilder().StationId(2001).StationNameCn("北京南站").Build(),
		}).
		Build()

	if info.CityId == nil || *info.CityId != "2001" {
		t.Errorf("CityId = %v, want 2001", info.CityId)
	}
	if info.CityNameCn == nil || *info.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", info.CityNameCn)
	}
	if info.CityNameEn == nil || *info.CityNameEn != "Beijing" {
		t.Errorf("CityNameEn = %v, want Beijing", info.CityNameEn)
	}
	if info.ProvinceId == nil || *info.ProvinceId != "1" {
		t.Errorf("ProvinceId = %v, want 1", info.ProvinceId)
	}
	if info.ProvinceNameCn == nil || *info.ProvinceNameCn != "北京市" {
		t.Errorf("ProvinceNameCn = %v, want 北京市", info.ProvinceNameCn)
	}
	if info.ProvinceNameEn == nil || *info.ProvinceNameEn != "Beijing" {
		t.Errorf("ProvinceNameEn = %v, want Beijing", info.ProvinceNameEn)
	}
	if info.CountryId == nil || *info.CountryId != "1" {
		t.Errorf("CountryId = %v, want 1", info.CountryId)
	}
	if info.CanonicalCountryCode == nil || *info.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", info.CanonicalCountryCode)
	}
	if info.CountryCode == nil || *info.CountryCode != "CHN" {
		t.Errorf("CountryCode = %v, want CHN", info.CountryCode)
	}
	if info.CountryNameCn == nil || *info.CountryNameCn != "中国" {
		t.Errorf("CountryNameCn = %v, want 中国", info.CountryNameCn)
	}
	if len(info.TrainStation) != 1 {
		t.Errorf("TrainStation len = %d, want 1", len(info.TrainStation))
	}

	// 部分设置
	info2 := NewTrainCityInfoBuilder().
		CityId("2002").
		Build()
	if info2.CityId == nil || *info2.CityId != "2002" {
		t.Errorf("CityId = %v, want 2002", info2.CityId)
	}
	if info2.TrainStation != nil {
		t.Errorf("TrainStation = %v, want nil", info2.TrainStation)
	}
}

// --- CarCityInfoBuilder ---

func TestCarCityInfoBuilder(t *testing.T) {
	info := NewCarCityInfoBuilder().
		CityId(2).
		CityName("深圳").
		CountyList([]CountyRecord{
			*NewCountyRecordBuilder().CountyId(3001).CountyName("南山区").Build(),
		}).
		Build()

	if info.CityId == nil || *info.CityId != 2 {
		t.Errorf("CityId = %v, want 2", info.CityId)
	}
	if info.CityName == nil || *info.CityName != "深圳" {
		t.Errorf("CityName = %v, want 深圳", info.CityName)
	}
	if len(info.CountyList) != 1 {
		t.Errorf("CountyList len = %d, want 1", len(info.CountyList))
	}

	// 部分设置
	info2 := NewCarCityInfoBuilder().
		CityId(1).
		Build()
	if info2.CityId == nil || *info2.CityId != 1 {
		t.Errorf("CityId = %v, want 1", info2.CityId)
	}
	if info2.CountyList != nil {
		t.Errorf("CountyList = %v, want nil", info2.CountyList)
	}
}

// =====================================================================
// ListAirportCity 模型层测试
// =====================================================================

func TestListAirportCityRequestBuilder_FullParams(t *testing.T) {
	req := NewListAirportCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp(1583484681).
		CompanyId("test_company").
		Sign("test_sign").
		CountryLevel(1).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.CountryLevel == nil || *req.CountryLevel != 1 {
		t.Errorf("CountryLevel = %v, want 1", req.CountryLevel)
	}
}

func TestListAirportCityRequestBuilder_PartialParams(t *testing.T) {
	req := NewListAirportCityRequestBuilder().
		ClientId("test_client").
		CountryLevel(2).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.CountryLevel == nil || *req.CountryLevel != 2 {
		t.Errorf("CountryLevel = %v, want 2", req.CountryLevel)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestListAirportCityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewListAirportCityRequestBuilder().
		ClientId("test_client").
		CountryLevel(0).
		Timestamp(0).
		Build()

	// int32 零值也应该被设置
	if req.CountryLevel == nil || *req.CountryLevel != 0 {
		t.Errorf("CountryLevel = %v, want 0", req.CountryLevel)
	}
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestListAirportCityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListAirportCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	// 业务参数不应存在
	if req.CountryLevel != nil {
		t.Errorf("CountryLevel = %v, want nil", req.CountryLevel)
	}
}

func TestListAirportCityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
				"city_id": 1,
				"city_name_cn": "北京",
				"city_name_en": "Beijing",
				"province_id": 1,
				"province_name_cn": "北京市",
				"province_name_en": "Beijing",
				"country_id": 1,
				"canonical_country_code": "CN",
				"country_code": "CHN",
				"country_name_cn": "中国",
				"country_name_en": "China",
				"flight_station": [
					{"airport_name_cn": "首都国际机场", "airport_name_en": "Beijing Capital", "airport_code": "PEK"}
				]
			}
		],
		"request_id": "req_001"
	}`

	var reply ListAirportCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}
	record := reply.Data[0]
	if record.CityId == nil || *record.CityId != 1 {
		t.Errorf("CityId = %v, want 1", record.CityId)
	}
	if record.CityNameCn == nil || *record.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", record.CityNameCn)
	}
	if record.ProvinceId == nil || *record.ProvinceId != 1 {
		t.Errorf("ProvinceId = %v, want 1", record.ProvinceId)
	}
	if record.CountryId == nil || *record.CountryId != 1 {
		t.Errorf("CountryId = %v, want 1", record.CountryId)
	}
	if record.CanonicalCountryCode == nil || *record.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", record.CanonicalCountryCode)
	}
	if len(record.FlightStation) != 1 {
		t.Fatalf("FlightStation len = %d, want 1", len(record.FlightStation))
	}
	if record.FlightStation[0].AirportCode == nil || *record.FlightStation[0].AirportCode != "PEK" {
		t.Errorf("FlightStation[0].AirportCode = %v, want PEK", record.FlightStation[0].AirportCode)
	}
}

func TestListAirportCityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListAirportCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListAirportCityApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"city_id": 1, "city_name_cn": "北京"},
			{"city_id": 2, "city_name_cn": "上海"}
		],
		"request_id": "req_multi"
	}`

	var reply ListAirportCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(reply.Data))
	}
	// 第2条部分字段缺失
	if reply.Data[1].ProvinceId != nil {
		t.Errorf("Data[1].ProvinceId = %v, want nil", reply.Data[1].ProvinceId)
	}
}

func TestListAirportCityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListAirportCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListAirportCityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListAirportCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

// =====================================================================
// ListAirportCity 资源方法测试
// =====================================================================

func TestListAirportCity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/DemeterAres/AirportCity/index" {
			t.Errorf("expected path /river/DemeterAres/AirportCity/index, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
				{
					"city_id": 1,
					"city_name_cn": "北京",
					"province_id": 1,
					"country_id": 1,
					"flight_station": [{"airport_name_cn": "首都国际机场", "airport_code": "PEK"}]
				}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().
			ClientId("test_client").
			CountryLevel(1).
			Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if resp.ListAirportCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListAirportCityApiReply.Errno)
	}
	if len(resp.ListAirportCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListAirportCityApiReply.Data))
	}
	if resp.ListAirportCityApiReply.Data[0].CityId == nil || *resp.ListAirportCityApiReply.Data[0].CityId != 1 {
		t.Errorf("Data[0].CityId = %v, want 1", resp.ListAirportCityApiReply.Data[0].CityId)
	}
}

func TestListAirportCity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if len(resp.ListAirportCityApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListAirportCityApiReply.Data))
	}
}

func TestListAirportCity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if resp.ListAirportCityApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListAirportCityApiReply.Errno)
	}
}

func TestListAirportCity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListAirportCityApiReply != nil {
		t.Errorf("ListAirportCityApiReply should be nil for non-200 response")
	}
}

func TestListAirportCity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1,"city_name_cn":"北京"}],"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if resp.ListAirportCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListAirportCityApiReply.Errno)
	}
	if len(resp.ListAirportCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListAirportCityApiReply.Data))
	}
}

func TestListAirportCity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1,"city_name_cn":"北京"}],"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if len(resp.ListAirportCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListAirportCityApiReply.Data))
	}
}

func TestListAirportCity_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
	if len(resp.ListAirportCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListAirportCityApiReply.Data))
	}
}

func TestListAirportCity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListAirportCityApiReqBuilder().
		ListAirportCityRequest(NewListAirportCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListAirportCity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListAirportCity() error = %v", err)
	}
	if resp.ListAirportCityApiReply == nil {
		t.Fatal("ListAirportCityApiReply is nil")
	}
}

// =====================================================================
// ListCarCity 模型层测试（GET 型，Builder 直接设 QueryParams）
// =====================================================================

func TestListCarCityApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListCarCityApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp("1583484681").
		CompanyId("test_company").
		Sign("test_sign").
		Build()

	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"timestamp", "1583484681"},
		{"company_id", "test_company"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListCarCityApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListCarCityApiReqBuilder().
		ClientId("test_client").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	// 未设置的参数应为空
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty")
	}
	if req.apiReq.QueryParams.Get("company_id") != "" {
		t.Errorf("company_id should be empty")
	}
}

func TestListCarCityApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListCarCityApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()

	// ListCarCity 无业务参数，仅有通用参数
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
}

func TestListCarCityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
				"city_id": 1,
				"city_name": "北京",
				"county_list": [
					{"county_id": 3001, "county_name": "朝阳区"},
					{"county_id": 3002, "county_name": "海淀区"}
				]
			}
		],
		"request_id": "req_001"
	}`

	var reply ListCarCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}
	if reply.Data[0].CityId == nil || *reply.Data[0].CityId != 1 {
		t.Errorf("Data[0].CityId = %v, want 1", reply.Data[0].CityId)
	}
	if len(reply.Data[0].CountyList) != 2 {
		t.Fatalf("CountyList len = %d, want 2", len(reply.Data[0].CountyList))
	}
	if reply.Data[0].CountyList[0].CountyId == nil || *reply.Data[0].CountyList[0].CountyId != 3001 {
		t.Errorf("CountyList[0].CountyId = %v, want 3001", reply.Data[0].CountyList[0].CountyId)
	}
}

func TestListCarCityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListCarCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestListCarCityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListCarCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListCarCityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListCarCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

// =====================================================================
// ListCarCity 资源方法测试（GET 型）
// =====================================================================

func TestListCarCity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/City/get" {
			t.Errorf("expected path /river/City/get, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
				{"city_id": 1, "city_name": "北京", "county_list": [{"county_id": 3001, "county_name": "朝阳区"}]}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if resp.ListCarCityApiReply.Errno == nil || *resp.ListCarCityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListCarCityApiReply.Errno)
	}
	if len(resp.ListCarCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCarCityApiReply.Data))
	}
	if resp.ListCarCityApiReply.Data[0].CityId == nil || *resp.ListCarCityApiReply.Data[0].CityId != 1 {
		t.Errorf("Data[0].CityId = %v, want 1", resp.ListCarCityApiReply.Data[0].CityId)
	}
}

func TestListCarCity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if len(resp.ListCarCityApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListCarCityApiReply.Data))
	}
}

func TestListCarCity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if resp.ListCarCityApiReply.Errno == nil || *resp.ListCarCityApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.ListCarCityApiReply.Errno)
	}
}

func TestListCarCity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListCarCityApiReply != nil {
		t.Errorf("ListCarCityApiReply should be nil for non-200 response")
	}
}

func TestListCarCity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1,"city_name":"北京"}],"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if resp.ListCarCityApiReply.Errno == nil || *resp.ListCarCityApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListCarCityApiReply.Errno)
	}
	if len(resp.ListCarCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCarCityApiReply.Data))
	}
}

func TestListCarCity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1,"city_name":"北京"}],"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if len(resp.ListCarCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCarCityApiReply.Data))
	}
}

func TestListCarCity_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":1}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
	if len(resp.ListCarCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCarCityApiReply.Data))
	}
}

func TestListCarCity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListCarCityApiReqBuilder().ClientId("test_client").Build()

	resp, err := c.ListCarCity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListCarCity() error = %v", err)
	}
	if resp.ListCarCityApiReply == nil {
		t.Fatal("ListCarCityApiReply is nil")
	}
}

// =====================================================================
// ListCity 模型层测试（POST 型，Data 为值类型 ListCityReply）
// =====================================================================

func TestListCityRequestBuilder_FullParams(t *testing.T) {
	req := NewListCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		ParamJson("{\"country_id\":1}").
		ParamJsonObj(*NewListCityParamObjBuilder().
			CountryId(1).
			ProvinceId(1).
			CityId(1).
			ProductType("10,20,30").
			Build()).
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
	if req.ParamJson == nil || *req.ParamJson != "{\"country_id\":1}" {
		t.Errorf("ParamJson = %v, want {\"country_id\":1}", req.ParamJson)
	}
	if req.ParamJsonObj.CountryId == nil || *req.ParamJsonObj.CountryId != 1 {
		t.Errorf("ParamJsonObj.CountryId = %v, want 1", req.ParamJsonObj.CountryId)
	}
	if req.ParamJsonObj.CityId == nil || *req.ParamJsonObj.CityId != 1 {
		t.Errorf("ParamJsonObj.CityId = %v, want 1", req.ParamJsonObj.CityId)
	}
}

func TestListCityRequestBuilder_PartialParams(t *testing.T) {
	req := NewListCityRequestBuilder().
		ClientId("test_client").
		ParamJson("{\"city_id\":1}").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ParamJson == nil || *req.ParamJson != "{\"city_id\":1}" {
		t.Errorf("ParamJson = %v, want {\"city_id\":1}", req.ParamJson)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestListCityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewListCityRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestListCityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.ParamJson != nil {
		t.Errorf("ParamJson = %v, want nil", req.ParamJson)
	}
}

func TestListCityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{
					"province_id": 1,
					"province_name_cn": "北京市",
					"province_name_en": "Beijing",
					"city_list": [
						{
							"city_id": 1,
							"product_type": [10, 20, 30],
							"city_name_cn": "北京",
							"city_short_name": "京",
							"city_path_id": "1-1",
							"city_path_cn": "北京市 - 北京",
							"city_path_en": "Beijing - Beijing",
							"city_name_en": "Beijing",
							"flight_station": [{"airport_id": 1001, "airport_code": "PEK"}],
							"train_station": [{"station_id": 2001, "station_name_cn": "北京南站"}]
						}
					]
				}
			]
		},
		"request_id": "req_001"
	}`

	var reply ListCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	// Data 是值类型，不会为 nil
	if reply.Data.Total == nil || *reply.Data.Total != 2 {
		t.Errorf("Total = %v, want 2", reply.Data.Total)
	}
	if len(reply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
	}
	record := reply.Data.Records[0]
	if record.ProvinceId == nil || *record.ProvinceId != 1 {
		t.Errorf("ProvinceId = %v, want 1", record.ProvinceId)
	}
	if record.ProvinceNameCn == nil || *record.ProvinceNameCn != "北京市" {
		t.Errorf("ProvinceNameCn = %v, want 北京市", record.ProvinceNameCn)
	}
	if len(record.CityList) != 1 {
		t.Fatalf("CityList len = %d, want 1", len(record.CityList))
	}
	city := record.CityList[0]
	if city.CityId == nil || *city.CityId != 1 {
		t.Errorf("CityId = %v, want 1", city.CityId)
	}
	if len(city.ProductType) != 3 || city.ProductType[0] != 10 {
		t.Errorf("ProductType = %v, want [10,20,30]", city.ProductType)
	}
	if city.CityShortName == nil || *city.CityShortName != "京" {
		t.Errorf("CityShortName = %v, want 京", city.CityShortName)
	}
	if len(city.FlightStation) != 1 {
		t.Fatalf("FlightStation len = %d, want 1", len(city.FlightStation))
	}
	if city.FlightStation[0].AirportId == nil || *city.FlightStation[0].AirportId != 1001 {
		t.Errorf("FlightStation[0].AirportId = %v, want 1001", city.FlightStation[0].AirportId)
	}
	if len(city.TrainStation) != 1 {
		t.Fatalf("TrainStation len = %d, want 1", len(city.TrainStation))
	}
	if city.TrainStation[0].StationId == nil || *city.TrainStation[0].StationId != 2001 {
		t.Errorf("TrainStation[0].StationId = %v, want 2001", city.TrainStation[0].StationId)
	}
}

func TestListCityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListCityApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"total": 2,
			"records": [
				{"province_id": 1, "province_name_cn": "北京市"},
				{"province_id": 2, "province_name_cn": "上海市"}
			]
		},
		"request_id": "req_multi"
	}`

	var reply ListCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 2 {
		t.Fatalf("Records len = %d, want 2", len(reply.Data.Records))
	}
	// 第2条部分字段缺失
	if reply.Data.Records[1].CityList != nil {
		t.Errorf("Records[1].CityList = %v, want nil", reply.Data.Records[1].CityList)
	}
}

func TestListCityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_empty"}`

	var reply ListCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestListCityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	// Data 是值类型，Records 应为 nil（空）
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

// =====================================================================
// ListCity 资源方法测试
// =====================================================================

func TestListCity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/open-apis/v1/city/list" {
			t.Errorf("expected path /open-apis/v1/city/list, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"total": 1,
				"records": [
					{
						"province_id": 1,
						"province_name_cn": "北京市",
						"city_list": [
							{"city_id": 1, "city_name_cn": "北京", "product_type": [10, 20]}
						]
					}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().
			ClientId("test_client").
			ParamJson("{\"country_id\":1}").
			Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if resp.ListCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListCityApiReply.Errno)
	}
	if resp.ListCityApiReply.Data.Total == nil || *resp.ListCityApiReply.Data.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.ListCityApiReply.Data.Total)
	}
	if len(resp.ListCityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListCityApiReply.Data.Records))
	}
	if len(resp.ListCityApiReply.Data.Records[0].CityList) != 1 {
		t.Fatalf("CityList len = %d, want 1", len(resp.ListCityApiReply.Data.Records[0].CityList))
	}
}

func TestListCity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if len(resp.ListCityApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.ListCityApiReply.Data.Records))
	}
}

func TestListCity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if resp.ListCityApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListCityApiReply.Errno)
	}
}

func TestListCity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListCityApiReply != nil {
		t.Errorf("ListCityApiReply should be nil for non-200 response")
	}
}

func TestListCity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"province_id":1,"province_name_cn":"北京市"}]},"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if resp.ListCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListCityApiReply.Errno)
	}
	if len(resp.ListCityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListCityApiReply.Data.Records))
	}
}

func TestListCity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"province_id":1,"province_name_cn":"北京市"}]},"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if len(resp.ListCityApiReply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(resp.ListCityApiReply.Data.Records))
	}
}

func TestListCity_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
	if len(resp.ListCityApiReply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(resp.ListCityApiReply.Data.Records))
	}
}

func TestListCity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListCityApiReqBuilder().
		ListCityRequest(NewListCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListCity() error = %v", err)
	}
	if resp.ListCityApiReply == nil {
		t.Fatal("ListCityApiReply is nil")
	}
}

// =====================================================================
// ListCountry 模型层测试（POST 型，无业务参数）
// =====================================================================

func TestListCountryRequestBuilder_FullParams(t *testing.T) {
	req := NewListCountryRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp(1583484681).
		CompanyId("test_company").
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
}

func TestListCountryRequestBuilder_PartialParams(t *testing.T) {
	req := NewListCountryRequestBuilder().
		ClientId("test_client").
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

func TestListCountryRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewListCountryRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestListCountryRequestBuilder_OnlyCommonParams(t *testing.T) {
	// ListCountry 无业务参数，仅通用参数
	req := NewListCountryRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
}

func TestListCountryApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
				"country_id": 1,
				"canonical_country_code": "CN",
				"country_code": "CHN",
				"country_name_cn": "中国",
				"country_name_en": "China",
				"continent_id": 1,
				"continent_name_cn": "亚洲",
				"continent_name_en": "Asia"
			}
		],
		"request_id": "req_001"
	}`

	var reply ListCountryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}
	record := reply.Data[0]
	if record.CountryId == nil || *record.CountryId != 1 {
		t.Errorf("CountryId = %v, want 1", record.CountryId)
	}
	if record.CanonicalCountryCode == nil || *record.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", record.CanonicalCountryCode)
	}
	if record.CountryCode == nil || *record.CountryCode != "CHN" {
		t.Errorf("CountryCode = %v, want CHN", record.CountryCode)
	}
	if record.CountryNameCn == nil || *record.CountryNameCn != "中国" {
		t.Errorf("CountryNameCn = %v, want 中国", record.CountryNameCn)
	}
	if record.ContinentId == nil || *record.ContinentId != 1 {
		t.Errorf("ContinentId = %v, want 1", record.ContinentId)
	}
	if record.ContinentNameCn == nil || *record.ContinentNameCn != "亚洲" {
		t.Errorf("ContinentNameCn = %v, want 亚洲", record.ContinentNameCn)
	}
}

func TestListCountryApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListCountryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListCountryApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"country_id": 1, "country_name_cn": "中国"},
			{"country_id": 2, "country_name_cn": "美国"}
		],
		"request_id": "req_multi"
	}`

	var reply ListCountryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(reply.Data))
	}
	// 第2条部分字段缺失
	if reply.Data[1].ContinentId != nil {
		t.Errorf("Data[1].ContinentId = %v, want nil", reply.Data[1].ContinentId)
	}
}

func TestListCountryApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListCountryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListCountryApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListCountryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

// =====================================================================
// ListCountry 资源方法测试
// =====================================================================

func TestListCountry_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/DemeterAres/Country/index" {
			t.Errorf("expected path /river/DemeterAres/Country/index, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
			{"country_id": 1, "country_name_cn": "中国", "country_code": "CHN"}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if resp.ListCountryApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListCountryApiReply.Errno)
	}
	if len(resp.ListCountryApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCountryApiReply.Data))
	}
	if resp.ListCountryApiReply.Data[0].CountryId == nil || *resp.ListCountryApiReply.Data[0].CountryId != 1 {
		t.Errorf("Data[0].CountryId = %v, want 1", resp.ListCountryApiReply.Data[0].CountryId)
	}
}

func TestListCountry_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if len(resp.ListCountryApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListCountryApiReply.Data))
	}
}

func TestListCountry_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if resp.ListCountryApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListCountryApiReply.Errno)
	}
}

func TestListCountry_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListCountryApiReply != nil {
		t.Errorf("ListCountryApiReply should be nil for non-200 response")
	}
}

func TestListCountry_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"country_id":1,"country_name_cn":"中国"}],"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if resp.ListCountryApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListCountryApiReply.Errno)
	}
	if len(resp.ListCountryApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCountryApiReply.Data))
	}
}

func TestListCountry_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"country_id":1,"country_name_cn":"中国"}],"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if len(resp.ListCountryApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCountryApiReply.Data))
	}
}

func TestListCountry_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"country_id":1}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
	if len(resp.ListCountryApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListCountryApiReply.Data))
	}
}

func TestListCountry_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListCountryApiReqBuilder().
		ListCountryRequest(NewListCountryRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListCountry(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListCountry() error = %v", err)
	}
	if resp.ListCountryApiReply == nil {
		t.Fatal("ListCountryApiReply is nil")
	}
}

// =====================================================================
// ListHotelCity 模型层测试（POST 型，CityId 为 *string）
// =====================================================================

func TestListHotelCityRequestBuilder_FullParams(t *testing.T) {
	req := NewListHotelCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp(1583484681).
		CompanyId("test_company").
		Sign("test_sign").
		CountryId(1).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.CountryId == nil || *req.CountryId != 1 {
		t.Errorf("CountryId = %v, want 1", req.CountryId)
	}
}

func TestListHotelCityRequestBuilder_PartialParams(t *testing.T) {
	req := NewListHotelCityRequestBuilder().
		ClientId("test_client").
		CountryId(2).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.CountryId == nil || *req.CountryId != 2 {
		t.Errorf("CountryId = %v, want 2", req.CountryId)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
}

func TestListHotelCityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewListHotelCityRequestBuilder().
		ClientId("test_client").
		CountryId(0).
		Timestamp(0).
		Build()

	if req.CountryId == nil || *req.CountryId != 0 {
		t.Errorf("CountryId = %v, want 0", req.CountryId)
	}
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestListHotelCityRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListHotelCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.CountryId != nil {
		t.Errorf("CountryId = %v, want nil", req.CountryId)
	}
}

func TestListHotelCityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
			"city_id": "1001",
				"description_cn": "北京城市描述",
				"description_en": "Beijing city description",
				"city_name_cn": "北京",
				"city_name_en": "Beijing",
				"province_id": "1",
				"province_name_cn": "北京市",
				"province_name_en": "Beijing",
				"country_id": "1",
				"canonical_country_code": "CN",
				"country_code": "CHN",
				"country_name_cn": "中国",
				"country_name_en": "China"
			}
		],
		"request_id": "req_001"
	}`

	var reply ListHotelCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}
	record := reply.Data[0]
	if record.CityId == nil || *record.CityId != "1001" {
		t.Errorf("CityId = %v, want 1001", record.CityId)
	}
	if record.DescriptionCn == nil || *record.DescriptionCn != "北京城市描述" {
		t.Errorf("DescriptionCn = %v, want 北京城市描述", record.DescriptionCn)
	}
	if record.CityNameCn == nil || *record.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", record.CityNameCn)
	}
	if record.ProvinceId == nil || *record.ProvinceId != "1" {
		t.Errorf("ProvinceId = %v, want 1", record.ProvinceId)
	}
	if record.CountryId == nil || *record.CountryId != "1" {
		t.Errorf("CountryId = %v, want 1", record.CountryId)
	}
	if record.CanonicalCountryCode == nil || *record.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", record.CanonicalCountryCode)
	}
}

func TestListHotelCityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListHotelCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListHotelCityApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"city_id": "1001", "city_name_cn": "北京"},
			{"city_id": "1002", "city_name_cn": "上海"}
		],
		"request_id": "req_multi"
	}`

	var reply ListHotelCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(reply.Data))
	}
	// 第2条部分字段缺失
	if reply.Data[1].CountryId != nil {
		t.Errorf("Data[1].CountryId = %v, want nil", reply.Data[1].CountryId)
	}
}

func TestListHotelCityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListHotelCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListHotelCityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListHotelCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

// =====================================================================
// ListHotelCity 资源方法测试
// =====================================================================

func TestListHotelCity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/DemeterAres/HotelCity/index" {
			t.Errorf("expected path /river/DemeterAres/HotelCity/index, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
			{"city_id": "1001", "city_name_cn": "北京", "country_id": "1"}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().
			ClientId("test_client").
			CountryId(1).
			Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if resp.ListHotelCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListHotelCityApiReply.Errno)
	}
	if len(resp.ListHotelCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListHotelCityApiReply.Data))
	}
	if resp.ListHotelCityApiReply.Data[0].CityId == nil || *resp.ListHotelCityApiReply.Data[0].CityId != "1001" {
		t.Errorf("Data[0].CityId = %v, want 1001", resp.ListHotelCityApiReply.Data[0].CityId)
	}
}

func TestListHotelCity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if len(resp.ListHotelCityApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListHotelCityApiReply.Data))
	}
}

func TestListHotelCity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if resp.ListHotelCityApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListHotelCityApiReply.Errno)
	}
}

func TestListHotelCity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListHotelCityApiReply != nil {
		t.Errorf("ListHotelCityApiReply should be nil for non-200 response")
	}
}

func TestListHotelCity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"1001","city_name_cn":"北京"}],"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if resp.ListHotelCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListHotelCityApiReply.Errno)
	}
	if len(resp.ListHotelCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListHotelCityApiReply.Data))
	}
}

func TestListHotelCity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"1001","city_name_cn":"北京"}],"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if len(resp.ListHotelCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListHotelCityApiReply.Data))
	}
}

func TestListHotelCity_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"1001"}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
	if len(resp.ListHotelCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListHotelCityApiReply.Data))
	}
}

func TestListHotelCity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListHotelCityApiReqBuilder().
		ListHotelCityRequest(NewListHotelCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListHotelCity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListHotelCity() error = %v", err)
	}
	if resp.ListHotelCityApiReply == nil {
		t.Fatal("ListHotelCityApiReply is nil")
	}
}

// =====================================================================
// ListTrainCity 模型层测试（POST 型，无业务参数，CityId 为 *string）
// =====================================================================

func TestListTrainCityRequestBuilder_FullParams(t *testing.T) {
	req := NewListTrainCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp(1583484681).
		CompanyId("test_company").
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
}

func TestListTrainCityRequestBuilder_PartialParams(t *testing.T) {
	req := NewListTrainCityRequestBuilder().
		ClientId("test_client").
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

func TestListTrainCityRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewListTrainCityRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestListTrainCityRequestBuilder_OnlyCommonParams(t *testing.T) {
	// ListTrainCity 无业务参数，仅通用参数
	req := NewListTrainCityRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
}

func TestListTrainCityApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{
			"city_id": "2001",
				"city_name_cn": "北京",
				"city_name_en": "Beijing",
				"province_id": "1",
				"province_name_cn": "北京市",
				"province_name_en": "Beijing",
				"country_id": "1",
				"canonical_country_code": "CN",
				"country_code": "CHN",
				"country_name_cn": "中国",
				"train_station": [
					{"station_name_cn": "北京南站", "station_name_en": "Beijing South", "station_name": "北京南", "station_id": 2001}
				]
			}
		],
		"request_id": "req_001"
	}`

	var reply ListTrainCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if len(reply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(reply.Data))
	}
	record := reply.Data[0]
	if record.CityId == nil || *record.CityId != "2001" {
		t.Errorf("CityId = %v, want 2001", record.CityId)
	}
	if record.CityNameCn == nil || *record.CityNameCn != "北京" {
		t.Errorf("CityNameCn = %v, want 北京", record.CityNameCn)
	}
	if record.ProvinceId == nil || *record.ProvinceId != "1" {
		t.Errorf("ProvinceId = %v, want 1", record.ProvinceId)
	}
	if record.CountryId == nil || *record.CountryId != "1" {
		t.Errorf("CountryId = %v, want 1", record.CountryId)
	}
	if record.CanonicalCountryCode == nil || *record.CanonicalCountryCode != "CN" {
		t.Errorf("CanonicalCountryCode = %v, want CN", record.CanonicalCountryCode)
	}
	if len(record.TrainStation) != 1 {
		t.Fatalf("TrainStation len = %d, want 1", len(record.TrainStation))
	}
	if record.TrainStation[0].StationId == nil || *record.TrainStation[0].StationId != 2001 {
		t.Errorf("TrainStation[0].StationId = %v, want 2001", record.TrainStation[0].StationId)
	}
	if record.TrainStation[0].StationNameCn == nil || *record.TrainStation[0].StationNameCn != "北京南站" {
		t.Errorf("TrainStation[0].StationNameCn = %v, want 北京南站", record.TrainStation[0].StationNameCn)
	}
}

func TestListTrainCityApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`

	var reply ListTrainCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
}

func TestListTrainCityApiReply_MultipleItems(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": [
			{"city_id": "2001", "city_name_cn": "北京"},
			{"city_id": "2002", "city_name_cn": "上海"}
		],
		"request_id": "req_multi"
	}`

	var reply ListTrainCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 2 {
		t.Fatalf("Data len = %d, want 2", len(reply.Data))
	}
	// 第2条部分字段缺失
	if reply.Data[1].TrainStation != nil {
		t.Errorf("Data[1].TrainStation = %v, want nil", reply.Data[1].TrainStation)
	}
}

func TestListTrainCityApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply ListTrainCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

func TestListTrainCityApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListTrainCityApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if len(reply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(reply.Data))
	}
}

// =====================================================================
// ListTrainCity 资源方法测试
// =====================================================================

func TestListTrainCity_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/DemeterAres/TrainCity" {
			t.Errorf("expected path /river/DemeterAres/TrainCity, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": [
			{"city_id": "2001", "city_name_cn": "北京", "train_station": [{"station_id": 2001, "station_name_cn": "北京南站"}]}
			],
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if resp.ListTrainCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListTrainCityApiReply.Errno)
	}
	if len(resp.ListTrainCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListTrainCityApiReply.Data))
	}
	if resp.ListTrainCityApiReply.Data[0].CityId == nil || *resp.ListTrainCityApiReply.Data[0].CityId != "2001" {
		t.Errorf("Data[0].CityId = %v, want 2001", resp.ListTrainCityApiReply.Data[0].CityId)
	}
}

func TestListTrainCity_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if len(resp.ListTrainCityApiReply.Data) != 0 {
		t.Errorf("Data len = %d, want 0", len(resp.ListTrainCityApiReply.Data))
	}
}

func TestListTrainCity_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if resp.ListTrainCityApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.ListTrainCityApiReply.Errno)
	}
}

func TestListTrainCity_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListTrainCityApiReply != nil {
		t.Errorf("ListTrainCityApiReply should be nil for non-200 response")
	}
}

func TestListTrainCity_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"2001","city_name_cn":"北京"}],"request_id":"req_enc"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if resp.ListTrainCityApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.ListTrainCityApiReply.Errno)
	}
	if len(resp.ListTrainCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListTrainCityApiReply.Data))
	}
}

func TestListTrainCity_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"2001","city_name_cn":"北京"}],"request_id":"req_enc256"}`
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

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if len(resp.ListTrainCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListTrainCityApiReply.Data))
	}
}

func TestListTrainCity_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[{"city_id":"2001"}],"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	c := &city{option: option}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
	if len(resp.ListTrainCityApiReply.Data) != 1 {
		t.Fatalf("Data len = %d, want 1", len(resp.ListTrainCityApiReply.Data))
	}
}

func TestListTrainCity_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newCityTestOption(testServer.URL)
	c := &city{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}

	req := NewListTrainCityApiReqBuilder().
		ListTrainCityRequest(NewListTrainCityRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := c.ListTrainCity(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListTrainCity() error = %v", err)
	}
	if resp.ListTrainCityApiReply == nil {
		t.Fatal("ListTrainCityApiReply is nil")
	}
}
