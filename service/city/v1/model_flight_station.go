package v1

// FlightStation 机场信息
type FlightStation struct {
	AirportId     *int64  `json:"airport_id,omitempty"`      // 机场ID，滴滴主键
	AirportNameCn *string `json:"airport_name_cn,omitempty"` // 机场中文名
	AirportNameEn *string `json:"airport_name_en,omitempty"` // 机场英文名
	AirportCode   *string `json:"airport_code,omitempty"`    // 机场三代
}

type FlightStationBuilder struct {
	airportId        int64 // 机场ID，滴滴主键
	airportIdSet     bool
	airportNameCn    string // 机场中文名
	airportNameCnSet bool
	airportNameEn    string // 机场英文名
	airportNameEnSet bool
	airportCode      string // 机场三代
	airportCodeSet   bool
}

func NewFlightStationBuilder() *FlightStationBuilder {
	return &FlightStationBuilder{}
}
func (builder *FlightStationBuilder) AirportId(airportId int64) *FlightStationBuilder {
	builder.airportId = airportId
	builder.airportIdSet = true
	return builder
}
func (builder *FlightStationBuilder) AirportNameCn(airportNameCn string) *FlightStationBuilder {
	builder.airportNameCn = airportNameCn
	builder.airportNameCnSet = true
	return builder
}
func (builder *FlightStationBuilder) AirportNameEn(airportNameEn string) *FlightStationBuilder {
	builder.airportNameEn = airportNameEn
	builder.airportNameEnSet = true
	return builder
}
func (builder *FlightStationBuilder) AirportCode(airportCode string) *FlightStationBuilder {
	builder.airportCode = airportCode
	builder.airportCodeSet = true
	return builder
}

func (builder *FlightStationBuilder) Build() *FlightStation {
	data := &FlightStation{}
	if builder.airportIdSet {
		data.AirportId = &builder.airportId
	}
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
