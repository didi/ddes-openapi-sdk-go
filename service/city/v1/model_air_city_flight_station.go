package v1

// AirCityFlightStation 机票城市机场信息
type AirCityFlightStation struct {
	AirportNameCn *string `json:"airport_name_cn,omitempty"` // 机场中文名
	AirportNameEn *string `json:"airport_name_en,omitempty"` // 机场英文名
	AirportCode   *string `json:"airport_code,omitempty"`    // 机场三代
}

type AirCityFlightStationBuilder struct {
	airportNameCn    string // 机场中文名
	airportNameCnSet bool
	airportNameEn    string // 机场英文名
	airportNameEnSet bool
	airportCode      string // 机场三代
	airportCodeSet   bool
}

func NewAirCityFlightStationBuilder() *AirCityFlightStationBuilder {
	return &AirCityFlightStationBuilder{}
}
func (builder *AirCityFlightStationBuilder) AirportNameCn(airportNameCn string) *AirCityFlightStationBuilder {
	builder.airportNameCn = airportNameCn
	builder.airportNameCnSet = true
	return builder
}
func (builder *AirCityFlightStationBuilder) AirportNameEn(airportNameEn string) *AirCityFlightStationBuilder {
	builder.airportNameEn = airportNameEn
	builder.airportNameEnSet = true
	return builder
}
func (builder *AirCityFlightStationBuilder) AirportCode(airportCode string) *AirCityFlightStationBuilder {
	builder.airportCode = airportCode
	builder.airportCodeSet = true
	return builder
}

func (builder *AirCityFlightStationBuilder) Build() *AirCityFlightStation {
	data := &AirCityFlightStation{}
	if builder.airportNameCnSet {
		data.AirportNameCn = &builder.airportNameCn
	}
	if builder.airportNameEnSet {
		data.AirportNameEn = &builder.airportNameEn
	}
	if builder.airportCodeSet {
		data.AirportCode = &builder.airportCode
	}
	return data
}
