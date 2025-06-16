package v1

// ArrivalInfo 到达信息的结构
type ArrivalInfo struct {
	ArrivalAirportCode       *string `json:"arrival_airport_code,omitempty"`        // 到达机场三字码
	ArrivalAirportName       *string `json:"arrival_airport_name,omitempty"`        // 到达机场名称
	ArrivalAirportFullName   *string `json:"arrival_airport_full_name,omitempty"`   // 到达机场全称
	ArrivalAirportSimpleName *string `json:"arrival_airport_simple_name,omitempty"` // 到达机场简称
	ArrivalCityName          *string `json:"arrival_city_name,omitempty"`           // 到达城市名
	ArrivalTerminal          *string `json:"arrival_terminal,omitempty"`            // 到达机场航站楼
	ArrivalDatetime          *string `json:"arrival_datetime,omitempty"`            // 到达日期时间
	ArrivalCityId            *string `json:"arrival_city_id,omitempty"`             // 到达城市id
}

type ArrivalInfoBuilder struct {
	arrivalAirportCode          string // 到达机场三字码
	arrivalAirportCodeSet       bool
	arrivalAirportName          string // 到达机场名称
	arrivalAirportNameSet       bool
	arrivalAirportFullName      string // 到达机场全称
	arrivalAirportFullNameSet   bool
	arrivalAirportSimpleName    string // 到达机场简称
	arrivalAirportSimpleNameSet bool
	arrivalCityName             string // 到达城市名
	arrivalCityNameSet          bool
	arrivalTerminal             string // 到达机场航站楼
	arrivalTerminalSet          bool
	arrivalDatetime             string // 到达日期时间
	arrivalDatetimeSet          bool
	arrivalCityId               string // 到达城市id
	arrivalCityIdSet            bool
}

func NewArrivalInfoBuilder() *ArrivalInfoBuilder {
	return &ArrivalInfoBuilder{}
}
func (builder *ArrivalInfoBuilder) ArrivalAirportCode(arrivalAirportCode string) *ArrivalInfoBuilder {
	builder.arrivalAirportCode = arrivalAirportCode
	builder.arrivalAirportCodeSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalAirportName(arrivalAirportName string) *ArrivalInfoBuilder {
	builder.arrivalAirportName = arrivalAirportName
	builder.arrivalAirportNameSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalAirportFullName(arrivalAirportFullName string) *ArrivalInfoBuilder {
	builder.arrivalAirportFullName = arrivalAirportFullName
	builder.arrivalAirportFullNameSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalAirportSimpleName(arrivalAirportSimpleName string) *ArrivalInfoBuilder {
	builder.arrivalAirportSimpleName = arrivalAirportSimpleName
	builder.arrivalAirportSimpleNameSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalCityName(arrivalCityName string) *ArrivalInfoBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalTerminal(arrivalTerminal string) *ArrivalInfoBuilder {
	builder.arrivalTerminal = arrivalTerminal
	builder.arrivalTerminalSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalDatetime(arrivalDatetime string) *ArrivalInfoBuilder {
	builder.arrivalDatetime = arrivalDatetime
	builder.arrivalDatetimeSet = true
	return builder
}
func (builder *ArrivalInfoBuilder) ArrivalCityId(arrivalCityId string) *ArrivalInfoBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}

func (builder *ArrivalInfoBuilder) Build() *ArrivalInfo {
	data := &ArrivalInfo{}
	if builder.arrivalAirportCodeSet {
		data.ArrivalAirportCode = &builder.arrivalAirportCode
	}
	if builder.arrivalAirportNameSet {
		data.ArrivalAirportName = &builder.arrivalAirportName
	}
	if builder.arrivalAirportFullNameSet {
		data.ArrivalAirportFullName = &builder.arrivalAirportFullName
	}
	if builder.arrivalAirportSimpleNameSet {
		data.ArrivalAirportSimpleName = &builder.arrivalAirportSimpleName
	}
	if builder.arrivalCityNameSet {
		data.ArrivalCityName = &builder.arrivalCityName
	}
	if builder.arrivalTerminalSet {
		data.ArrivalTerminal = &builder.arrivalTerminal
	}
	if builder.arrivalDatetimeSet {
		data.ArrivalDatetime = &builder.arrivalDatetime
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	return data
}
