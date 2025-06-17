package v1

// Routes 航段的结构，文档标注字段为数组，实际返回为obj;
type Routes struct {
	DepartureInfo *DepartureInfo `json:"departure_info,omitempty"`
	ArrivalInfo   *ArrivalInfo   `json:"arrival_info,omitempty"`
	AirlineInfo   *AirlineInfo   `json:"airline_info,omitempty"`
}

type RoutesBuilder struct {
	departureInfo    DepartureInfo
	departureInfoSet bool
	arrivalInfo      ArrivalInfo
	arrivalInfoSet   bool
	airlineInfo      AirlineInfo
	airlineInfoSet   bool
}

func NewRoutesBuilder() *RoutesBuilder {
	return &RoutesBuilder{}
}
func (builder *RoutesBuilder) DepartureInfo(departureInfo DepartureInfo) *RoutesBuilder {
	builder.departureInfo = departureInfo
	builder.departureInfoSet = true
	return builder
}
func (builder *RoutesBuilder) ArrivalInfo(arrivalInfo ArrivalInfo) *RoutesBuilder {
	builder.arrivalInfo = arrivalInfo
	builder.arrivalInfoSet = true
	return builder
}
func (builder *RoutesBuilder) AirlineInfo(airlineInfo AirlineInfo) *RoutesBuilder {
	builder.airlineInfo = airlineInfo
	builder.airlineInfoSet = true
	return builder
}

func (builder *RoutesBuilder) Build() *Routes {
	data := &Routes{}
	if builder.departureInfoSet {
		data.DepartureInfo = &builder.departureInfo
	}
	if builder.arrivalInfoSet {
		data.ArrivalInfo = &builder.arrivalInfo
	}
	if builder.airlineInfoSet {
		data.AirlineInfo = &builder.airlineInfo
	}
	return data
}
