package v1

// ListAirportCityReply struct for ListAirportCityReply
type ListAirportCityReply struct {
	CityId               *string                `json:"city_id,omitempty"`                // 滴滴城市ID（大ID，服务端返回值超 int32，用 *string）
	CityNameCn           *string                `json:"city_name_cn,omitempty"`           // 城市中文名
	CityNameEn           *string                `json:"city_name_en,omitempty"`           // 城市英文名
	ProvinceId           *int32                 `json:"province_id,omitempty"`            // 省ID
	ProvinceNameCn       *string                `json:"province_name_cn,omitempty"`       // 省中文名
	ProvinceNameEn       *string                `json:"province_name_en,omitempty"`       // 省英文文名
	CountryId            *string                `json:"country_id,omitempty"`             // 国家ID（大ID，用 *string）
	CanonicalCountryCode *string                `json:"canonical_country_code,omitempty"` // 国家二字代码
	CountryCode          *string                `json:"country_code,omitempty"`           // 国家三字代码
	CountryNameCn        *string                `json:"country_name_cn,omitempty"`        // 国家中文名
	CountryNameEn        *string                `json:"country_name_en,omitempty"`        // 国家英文名
	FlightStation        []AirCityFlightStation `json:"flight_station,omitempty"`         // 机场信息
}
