package v1

// DepartureInfo 出发信息的结构
type DepartureInfo struct {
	DepartureAirportCode       *string `json:"departure_airport_code,omitempty"`        // 出发机场三字码，如PEK
	DepartureAirportName       *string `json:"departure_airport_name,omitempty"`        // 出发机场名称，如 首都机场
	DepartureAirportFullName   *string `json:"departure_airport_full_name,omitempty"`   // 出发机场全称，如 北京首都国际机场
	DepartureAirportSimpleName *string `json:"departure_airport_simple_name,omitempty"` // 出发机场简称，如 首都
	DepartureCityName          *string `json:"departure_city_name,omitempty"`           // 出发城市名，如 北京
	DepartureTerminal          *string `json:"departure_terminal,omitempty"`            // 出发机场航站楼，可为空。如 T2
	DepartureDatetime          *string `json:"departure_datetime,omitempty"`            // 出发日期时间，如 2024-11-25 07:00
	DepartureCityId            *string `json:"departure_city_id,omitempty"`             // 出发城市id，如 1
}

type DepartureInfoBuilder struct {
	departureAirportCode          string // 出发机场三字码，如PEK
	departureAirportCodeSet       bool
	departureAirportName          string // 出发机场名称，如 首都机场
	departureAirportNameSet       bool
	departureAirportFullName      string // 出发机场全称，如 北京首都国际机场
	departureAirportFullNameSet   bool
	departureAirportSimpleName    string // 出发机场简称，如 首都
	departureAirportSimpleNameSet bool
	departureCityName             string // 出发城市名，如 北京
	departureCityNameSet          bool
	departureTerminal             string // 出发机场航站楼，可为空。如 T2
	departureTerminalSet          bool
	departureDatetime             string // 出发日期时间，如 2024-11-25 07:00
	departureDatetimeSet          bool
	departureCityId               string // 出发城市id，如 1
	departureCityIdSet            bool
}

func NewDepartureInfoBuilder() *DepartureInfoBuilder {
	return &DepartureInfoBuilder{}
}
func (builder *DepartureInfoBuilder) DepartureAirportCode(departureAirportCode string) *DepartureInfoBuilder {
	builder.departureAirportCode = departureAirportCode
	builder.departureAirportCodeSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureAirportName(departureAirportName string) *DepartureInfoBuilder {
	builder.departureAirportName = departureAirportName
	builder.departureAirportNameSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureAirportFullName(departureAirportFullName string) *DepartureInfoBuilder {
	builder.departureAirportFullName = departureAirportFullName
	builder.departureAirportFullNameSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureAirportSimpleName(departureAirportSimpleName string) *DepartureInfoBuilder {
	builder.departureAirportSimpleName = departureAirportSimpleName
	builder.departureAirportSimpleNameSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureCityName(departureCityName string) *DepartureInfoBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureTerminal(departureTerminal string) *DepartureInfoBuilder {
	builder.departureTerminal = departureTerminal
	builder.departureTerminalSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureDatetime(departureDatetime string) *DepartureInfoBuilder {
	builder.departureDatetime = departureDatetime
	builder.departureDatetimeSet = true
	return builder
}
func (builder *DepartureInfoBuilder) DepartureCityId(departureCityId string) *DepartureInfoBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}

func (builder *DepartureInfoBuilder) Build() *DepartureInfo {
	data := &DepartureInfo{}
	if builder.departureAirportCodeSet {
		data.DepartureAirportCode = &builder.departureAirportCode
	}
	if builder.departureAirportNameSet {
		data.DepartureAirportName = &builder.departureAirportName
	}
	if builder.departureAirportFullNameSet {
		data.DepartureAirportFullName = &builder.departureAirportFullName
	}
	if builder.departureAirportSimpleNameSet {
		data.DepartureAirportSimpleName = &builder.departureAirportSimpleName
	}
	if builder.departureCityNameSet {
		data.DepartureCityName = &builder.departureCityName
	}
	if builder.departureTerminalSet {
		data.DepartureTerminal = &builder.departureTerminal
	}
	if builder.departureDatetimeSet {
		data.DepartureDatetime = &builder.departureDatetime
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	return data
}
