package v1

// FlightList 航班列表项的结构
type FlightList struct {
	FlightInfo       *FlightInfo        `json:"flight_info,omitempty"`
	FlightRoutePrice []FlightRoutePrice `json:"flight_route_price,omitempty"` // 航班的价格信息，设计为数组是应对中转航班时，按航班号返回每个航班的预估价格。该数组长度为1，表示直飞；长度为2表示中转
}

type FlightListBuilder struct {
	flightInfo          FlightInfo
	flightInfoSet       bool
	flightRoutePrice    []FlightRoutePrice // 航班的价格信息，设计为数组是应对中转航班时，按航班号返回每个航班的预估价格。该数组长度为1，表示直飞；长度为2表示中转
	flightRoutePriceSet bool
}

func NewFlightListBuilder() *FlightListBuilder {
	return &FlightListBuilder{}
}
func (builder *FlightListBuilder) FlightInfo(flightInfo FlightInfo) *FlightListBuilder {
	builder.flightInfo = flightInfo
	builder.flightInfoSet = true
	return builder
}
func (builder *FlightListBuilder) FlightRoutePrice(flightRoutePrice []FlightRoutePrice) *FlightListBuilder {
	builder.flightRoutePrice = flightRoutePrice
	builder.flightRoutePriceSet = true
	return builder
}

func (builder *FlightListBuilder) Build() *FlightList {
	data := &FlightList{}
	if builder.flightInfoSet {
		data.FlightInfo = &builder.flightInfo
	}
	if builder.flightRoutePriceSet {
		data.FlightRoutePrice = builder.flightRoutePrice
	}
	return data
}
