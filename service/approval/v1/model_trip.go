package v1

// Trip 行程信息(TravelDetail中trips列表Obj信息)
type Trip struct {
	DepartureCity     *string      `json:"departure_city,omitempty"`      // 出发城市名称，出发城市名称 城市名称和城市ID建议传一种，不要一起传。
	DepartureCityId   *int32       `json:"departure_city_id,omitempty"`   // 滴滴出发城市ID，出发城市 ID，国际需要使用ID
	DestinationCity   *string      `json:"destination_city,omitempty"`    // 到达城市名称，到达城市名称 ；非轻度管控必传
	DestinationCityId *int32       `json:"destination_city_id,omitempty"` // 滴滴到达城市ID，到达城市 ID；非轻度管控必传 国际需要使用ID
	StartDate         *string      `json:"start_date,omitempty"`          // 行程段开始时间，行程段开始时间 轻度管控模式下，必须与 travel_detail 中的 start_date 相同 举例：2023-01-02
	EndDate           *string      `json:"end_date,omitempty"`            // 行程段结束时间，行程段结束时间 轻度管控模式下，必须与 travel_detail 中的 end_date 相同 举例：2023-01-14
	TripType          *string      `json:"trip_type,omitempty"`           // 出行方式，出行方式：0-其他；1-火车 ;2-飞机；支持多个，多个用英文逗号分开（用车酒店默认根据制度配置生效）举例：0,1,2，中度管控按照第一段生效
	ToCitys           []TravelCity `json:"to_citys,omitempty"`            // 目的城市集合，轻度管控必传，其他模式传了不生效 举例：[{\"id\":1,\"name\":\"北京\"}]
	IsReturn          *int32       `json:"is_return,omitempty"`           // 单程/往返，控制严格管控模式下，滴滴侧是否识别这个trip为往返行程生成返程行程权限 非必传字段，0-单程， 1-往返；默认为单程；传1后的拆分逻辑为：去程使用trip的传参生成，返程生成规则为：目的地-出发地，end_date-end_date，使用去程的trip_type；若对去程和返程的交通方式有分别管控诉求，不推荐使用此字段，建议拆成两个trip传输
}

type TripBuilder struct {
	departureCity        string // 出发城市名称，出发城市名称 城市名称和城市ID建议传一种，不要一起传。
	departureCitySet     bool
	departureCityId      int32 // 滴滴出发城市ID，出发城市 ID，国际需要使用ID
	departureCityIdSet   bool
	destinationCity      string // 到达城市名称，到达城市名称 ；非轻度管控必传
	destinationCitySet   bool
	destinationCityId    int32 // 滴滴到达城市ID，到达城市 ID；非轻度管控必传 国际需要使用ID
	destinationCityIdSet bool
	startDate            string // 行程段开始时间，行程段开始时间 轻度管控模式下，必须与 travel_detail 中的 start_date 相同 举例：2023-01-02
	startDateSet         bool
	endDate              string // 行程段结束时间，行程段结束时间 轻度管控模式下，必须与 travel_detail 中的 end_date 相同 举例：2023-01-14
	endDateSet           bool
	tripType             string // 出行方式，出行方式：0-其他；1-火车 ;2-飞机；支持多个，多个用英文逗号分开（用车酒店默认根据制度配置生效）举例：0,1,2，中度管控按照第一段生效
	tripTypeSet          bool
	toCitys              []TravelCity // 目的城市集合，轻度管控必传，其他模式传了不生效 举例：[{\"id\":1,\"name\":\"北京\"}]
	toCitysSet           bool
	isReturn             int32 // 单程/往返，控制严格管控模式下，滴滴侧是否识别这个trip为往返行程生成返程行程权限 非必传字段，0-单程， 1-往返；默认为单程；传1后的拆分逻辑为：去程使用trip的传参生成，返程生成规则为：目的地-出发地，end_date-end_date，使用去程的trip_type；若对去程和返程的交通方式有分别管控诉求，不推荐使用此字段，建议拆成两个trip传输
	isReturnSet          bool
}

func NewTripBuilder() *TripBuilder {
	return &TripBuilder{}
}
func (builder *TripBuilder) DepartureCity(departureCity string) *TripBuilder {
	builder.departureCity = departureCity
	builder.departureCitySet = true
	return builder
}
func (builder *TripBuilder) DepartureCityId(departureCityId int32) *TripBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *TripBuilder) DestinationCity(destinationCity string) *TripBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *TripBuilder) DestinationCityId(destinationCityId int32) *TripBuilder {
	builder.destinationCityId = destinationCityId
	builder.destinationCityIdSet = true
	return builder
}
func (builder *TripBuilder) StartDate(startDate string) *TripBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *TripBuilder) EndDate(endDate string) *TripBuilder {
	builder.endDate = endDate
	builder.endDateSet = true
	return builder
}
func (builder *TripBuilder) TripType(tripType string) *TripBuilder {
	builder.tripType = tripType
	builder.tripTypeSet = true
	return builder
}
func (builder *TripBuilder) ToCitys(toCitys []TravelCity) *TripBuilder {
	builder.toCitys = toCitys
	builder.toCitysSet = true
	return builder
}
func (builder *TripBuilder) IsReturn(isReturn int32) *TripBuilder {
	builder.isReturn = isReturn
	builder.isReturnSet = true
	return builder
}

func (builder *TripBuilder) Build() *Trip {
	data := &Trip{}
	if builder.departureCitySet {
		data.DepartureCity = &builder.departureCity
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.destinationCitySet {
		data.DestinationCity = &builder.destinationCity
	}
	if builder.destinationCityIdSet {
		data.DestinationCityId = &builder.destinationCityId
	}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.endDateSet {
		data.EndDate = &builder.endDate
	}
	if builder.tripTypeSet {
		data.TripType = &builder.tripType
	}
	if builder.toCitysSet {
		data.ToCitys = builder.toCitys
	}
	if builder.isReturnSet {
		data.IsReturn = &builder.isReturn
	}
	return data
}
