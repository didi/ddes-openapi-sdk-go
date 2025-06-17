package v1

// TravelDetail 差旅单行程信息 -> travel_detail
type TravelDetail struct {
	StartDate        *string      `json:"start_date,omitempty"` // 行程开始时间，格式： YYYY-MM-DD
	EndDate          *string      `json:"end_date,omitempty"`   // 行程结束时间，格式： YYYY-MM-DD
	Trips            []Trip       `json:"trips,omitempty"`      // 行程信息，行程信息；行程使用的制度(regulation_id)中，出差城市配置为无需填写时，trips可为空。否则必传，需按照行程顺序填写 详见 trip
	MeetingTrip      *MeetingTrip `json:"meeting_trip,omitempty"`
	StartCityRule    *int32       `json:"start_city_rule,omitempty"`    // 行程起点城市是否包含市内用车权限和酒店，（0-不包含，1-包含），trips第一行视为起点城市；默认为无市内用车和酒店，不管控和轻度管控不生效 不影响接送机服务
	EndCityRule      *int32       `json:"end_city_rule,omitempty"`      // 行程终点城市是否包含市内用车权限和酒店，（0-不包含，1-包含），trips最后一行视为终点城市，不管控和轻度管控不生效 不影响接送机服务,中度管控只对市内用车生效。严格管控下，对市内用车和酒店生效
	TrainTotalCount  *int32       `json:"train_total_count,omitempty"`  // 火车票次数，火车票次数；火车票次数 （不传走默认计算逻辑，传0代表火车票不可用）
	FlightTotalCount *int32       `json:"flight_total_count,omitempty"` // 机票次数，机票次数（不传走默认计算逻辑，传0代表机票不可用，对国内机票、国际机票都生效）
	HotelTotalCount  *int32       `json:"hotel_total_count,omitempty"`  // 酒店总间夜，酒店总间夜（不传走默认计算逻辑，传0代表酒店不可用，对国内酒店、国际酒店都生效）
	PickupTotalCount *int32       `json:"pickup_total_count,omitempty"` // 接送机总次数，接送机总次数（不传走默认计算逻辑，其他参数代表接送机总次数）
	CategoryControl  []int32      `json:"category_control,omitempty"`   // 品类控制开关，品类控制开关；制度有权限时,控制申请单是否创建品类对应服务 3-市内用车；4-火车票服务；6-机票服务；9-接送服务；10-酒店服务；11-国际机票服务，12 国际酒店服务 如果传了机火两个品类。制度上只开了机票品牌。则创建申请单成功，只能使用机票服务 category_control 传空数组会报错 举例：[4,10]，同时传递国内国际机票类别时，需要至少有一段国内城市对，否则报错。
}

type TravelDetailBuilder struct {
	startDate           string // 行程开始时间，格式： YYYY-MM-DD
	startDateSet        bool
	endDate             string // 行程结束时间，格式： YYYY-MM-DD
	endDateSet          bool
	trips               []Trip // 行程信息，行程信息；行程使用的制度(regulation_id)中，出差城市配置为无需填写时，trips可为空。否则必传，需按照行程顺序填写 详见 trip
	tripsSet            bool
	meetingTrip         MeetingTrip
	meetingTripSet      bool
	startCityRule       int32 // 行程起点城市是否包含市内用车权限和酒店，（0-不包含，1-包含），trips第一行视为起点城市；默认为无市内用车和酒店，不管控和轻度管控不生效 不影响接送机服务
	startCityRuleSet    bool
	endCityRule         int32 // 行程终点城市是否包含市内用车权限和酒店，（0-不包含，1-包含），trips最后一行视为终点城市，不管控和轻度管控不生效 不影响接送机服务,中度管控只对市内用车生效。严格管控下，对市内用车和酒店生效
	endCityRuleSet      bool
	trainTotalCount     int32 // 火车票次数，火车票次数；火车票次数 （不传走默认计算逻辑，传0代表火车票不可用）
	trainTotalCountSet  bool
	flightTotalCount    int32 // 机票次数，机票次数（不传走默认计算逻辑，传0代表机票不可用，对国内机票、国际机票都生效）
	flightTotalCountSet bool
	hotelTotalCount     int32 // 酒店总间夜，酒店总间夜（不传走默认计算逻辑，传0代表酒店不可用，对国内酒店、国际酒店都生效）
	hotelTotalCountSet  bool
	pickupTotalCount    int32 // 接送机总次数，接送机总次数（不传走默认计算逻辑，其他参数代表接送机总次数）
	pickupTotalCountSet bool
	categoryControl     []int32 // 品类控制开关，品类控制开关；制度有权限时,控制申请单是否创建品类对应服务 3-市内用车；4-火车票服务；6-机票服务；9-接送服务；10-酒店服务；11-国际机票服务，12 国际酒店服务 如果传了机火两个品类。制度上只开了机票品牌。则创建申请单成功，只能使用机票服务 category_control 传空数组会报错 举例：[4,10]，同时传递国内国际机票类别时，需要至少有一段国内城市对，否则报错。
	categoryControlSet  bool
}

func NewTravelDetailBuilder() *TravelDetailBuilder {
	return &TravelDetailBuilder{}
}
func (builder *TravelDetailBuilder) StartDate(startDate string) *TravelDetailBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *TravelDetailBuilder) EndDate(endDate string) *TravelDetailBuilder {
	builder.endDate = endDate
	builder.endDateSet = true
	return builder
}
func (builder *TravelDetailBuilder) Trips(trips []Trip) *TravelDetailBuilder {
	builder.trips = trips
	builder.tripsSet = true
	return builder
}
func (builder *TravelDetailBuilder) MeetingTrip(meetingTrip MeetingTrip) *TravelDetailBuilder {
	builder.meetingTrip = meetingTrip
	builder.meetingTripSet = true
	return builder
}
func (builder *TravelDetailBuilder) StartCityRule(startCityRule int32) *TravelDetailBuilder {
	builder.startCityRule = startCityRule
	builder.startCityRuleSet = true
	return builder
}
func (builder *TravelDetailBuilder) EndCityRule(endCityRule int32) *TravelDetailBuilder {
	builder.endCityRule = endCityRule
	builder.endCityRuleSet = true
	return builder
}
func (builder *TravelDetailBuilder) TrainTotalCount(trainTotalCount int32) *TravelDetailBuilder {
	builder.trainTotalCount = trainTotalCount
	builder.trainTotalCountSet = true
	return builder
}
func (builder *TravelDetailBuilder) FlightTotalCount(flightTotalCount int32) *TravelDetailBuilder {
	builder.flightTotalCount = flightTotalCount
	builder.flightTotalCountSet = true
	return builder
}
func (builder *TravelDetailBuilder) HotelTotalCount(hotelTotalCount int32) *TravelDetailBuilder {
	builder.hotelTotalCount = hotelTotalCount
	builder.hotelTotalCountSet = true
	return builder
}
func (builder *TravelDetailBuilder) PickupTotalCount(pickupTotalCount int32) *TravelDetailBuilder {
	builder.pickupTotalCount = pickupTotalCount
	builder.pickupTotalCountSet = true
	return builder
}
func (builder *TravelDetailBuilder) CategoryControl(categoryControl []int32) *TravelDetailBuilder {
	builder.categoryControl = categoryControl
	builder.categoryControlSet = true
	return builder
}

func (builder *TravelDetailBuilder) Build() *TravelDetail {
	data := &TravelDetail{}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.endDateSet {
		data.EndDate = &builder.endDate
	}
	if builder.tripsSet {
		data.Trips = builder.trips
	}
	if builder.meetingTripSet {
		data.MeetingTrip = &builder.meetingTrip
	}
	if builder.startCityRuleSet {
		data.StartCityRule = &builder.startCityRule
	}
	if builder.endCityRuleSet {
		data.EndCityRule = &builder.endCityRule
	}
	if builder.trainTotalCountSet {
		data.TrainTotalCount = &builder.trainTotalCount
	}
	if builder.flightTotalCountSet {
		data.FlightTotalCount = &builder.flightTotalCount
	}
	if builder.hotelTotalCountSet {
		data.HotelTotalCount = &builder.hotelTotalCount
	}
	if builder.pickupTotalCountSet {
		data.PickupTotalCount = &builder.pickupTotalCount
	}
	if builder.categoryControlSet {
		data.CategoryControl = builder.categoryControl
	}
	return data
}
