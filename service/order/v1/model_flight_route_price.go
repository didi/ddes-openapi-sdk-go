package v1

// FlightRoutePrice 航班价格信息的结构
type FlightRoutePrice struct {
	FlightNumber  *string `json:"flight_number,omitempty"`  // 航班号，直飞航班可以不关注。中转航班时，与routes数组里的route对象内的airline_info.flight_number 关联，确定每一个中转段对应的价格。
	EstimateFirst *int32  `json:"estimate_first,omitempty"` // 头等舱预估价，单位分，0表示该航班此舱售罄
	EstimateBiz   *int32  `json:"estimate_biz,omitempty"`   // 公务舱预估价，单位分，0表示该航班此舱售罄
	EstimateEco   *int32  `json:"estimate_eco,omitempty"`   // 经济舱预估价，单位分，0表示该航班此舱售罄
}

type FlightRoutePriceBuilder struct {
	flightNumber     string // 航班号，直飞航班可以不关注。中转航班时，与routes数组里的route对象内的airline_info.flight_number 关联，确定每一个中转段对应的价格。
	flightNumberSet  bool
	estimateFirst    int32 // 头等舱预估价，单位分，0表示该航班此舱售罄
	estimateFirstSet bool
	estimateBiz      int32 // 公务舱预估价，单位分，0表示该航班此舱售罄
	estimateBizSet   bool
	estimateEco      int32 // 经济舱预估价，单位分，0表示该航班此舱售罄
	estimateEcoSet   bool
}

func NewFlightRoutePriceBuilder() *FlightRoutePriceBuilder {
	return &FlightRoutePriceBuilder{}
}
func (builder *FlightRoutePriceBuilder) FlightNumber(flightNumber string) *FlightRoutePriceBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}
func (builder *FlightRoutePriceBuilder) EstimateFirst(estimateFirst int32) *FlightRoutePriceBuilder {
	builder.estimateFirst = estimateFirst
	builder.estimateFirstSet = true
	return builder
}
func (builder *FlightRoutePriceBuilder) EstimateBiz(estimateBiz int32) *FlightRoutePriceBuilder {
	builder.estimateBiz = estimateBiz
	builder.estimateBizSet = true
	return builder
}
func (builder *FlightRoutePriceBuilder) EstimateEco(estimateEco int32) *FlightRoutePriceBuilder {
	builder.estimateEco = estimateEco
	builder.estimateEcoSet = true
	return builder
}

func (builder *FlightRoutePriceBuilder) Build() *FlightRoutePrice {
	data := &FlightRoutePrice{}
	if builder.flightNumberSet {
		data.FlightNumber = &builder.flightNumber
	}
	if builder.estimateFirstSet {
		data.EstimateFirst = &builder.estimateFirst
	}
	if builder.estimateBizSet {
		data.EstimateBiz = &builder.estimateBiz
	}
	if builder.estimateEcoSet {
		data.EstimateEco = &builder.estimateEco
	}
	return data
}
