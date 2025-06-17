package v1

// AirlineInfo 航司信息的结构
type AirlineInfo struct {
	AirlineName           *string `json:"airline_name,omitempty"`             // 航司名称，如 中国东方航空股份有限公司
	AirlineSimpleName     *string `json:"airline_simple_name,omitempty"`      // 航司简称，如 东方航空
	AirlineVerySimpleName *string `json:"airline_very_simple_name,omitempty"` // 航司极简称，如 东航
	FlightNumber          *string `json:"flight_number,omitempty"`            // 航班号，如 MU5100
}

type AirlineInfoBuilder struct {
	airlineName              string // 航司名称，如 中国东方航空股份有限公司
	airlineNameSet           bool
	airlineSimpleName        string // 航司简称，如 东方航空
	airlineSimpleNameSet     bool
	airlineVerySimpleName    string // 航司极简称，如 东航
	airlineVerySimpleNameSet bool
	flightNumber             string // 航班号，如 MU5100
	flightNumberSet          bool
}

func NewAirlineInfoBuilder() *AirlineInfoBuilder {
	return &AirlineInfoBuilder{}
}
func (builder *AirlineInfoBuilder) AirlineName(airlineName string) *AirlineInfoBuilder {
	builder.airlineName = airlineName
	builder.airlineNameSet = true
	return builder
}
func (builder *AirlineInfoBuilder) AirlineSimpleName(airlineSimpleName string) *AirlineInfoBuilder {
	builder.airlineSimpleName = airlineSimpleName
	builder.airlineSimpleNameSet = true
	return builder
}
func (builder *AirlineInfoBuilder) AirlineVerySimpleName(airlineVerySimpleName string) *AirlineInfoBuilder {
	builder.airlineVerySimpleName = airlineVerySimpleName
	builder.airlineVerySimpleNameSet = true
	return builder
}
func (builder *AirlineInfoBuilder) FlightNumber(flightNumber string) *AirlineInfoBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}

func (builder *AirlineInfoBuilder) Build() *AirlineInfo {
	data := &AirlineInfo{}
	if builder.airlineNameSet {
		data.AirlineName = &builder.airlineName
	}
	if builder.airlineSimpleNameSet {
		data.AirlineSimpleName = &builder.airlineSimpleName
	}
	if builder.airlineVerySimpleNameSet {
		data.AirlineVerySimpleName = &builder.airlineVerySimpleName
	}
	if builder.flightNumberSet {
		data.FlightNumber = &builder.flightNumber
	}
	return data
}
