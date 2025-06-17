package v1

// TrainOrderGrabInfo 抢票信息消息类型
type TrainOrderGrabInfo struct {
	PassengerTravelerId  []string `json:"passenger_traveler_id,omitempty"`  // 滴滴ID或travelerId用于关联乘客信息，多个英文逗号隔开
	SelectedTime         []string `json:"selected_time,omitempty"`          // 格式：yyyy-MM-dd  ，多个英文逗号隔开，包含主选，主选是第一个。
	SelectedTrainCode    []string `json:"selected_train_code,omitempty"`    // 多个英文逗号隔开
	SelectedSeatType     []string `json:"selected_seat_type,omitempty"`     // 多个英文逗号隔开，字母
	SeatTypeName         []string `json:"seat_type_name,omitempty"`         // 多个英文逗号隔开，名称
	DepartureCityName    *string  `json:"departure_city_name,omitempty"`    // 城市中文名，主选
	DepartureCityId      *string  `json:"departure_city_id,omitempty"`      // 滴滴城市ID，主选
	ArrivalCityName      *string  `json:"arrival_city_name,omitempty"`      // 城市中文名，主选
	ArrivalCityId        *string  `json:"arrival_city_id,omitempty"`        // 滴滴城市ID，主选
	DepartureStationName *string  `json:"departure_station_name,omitempty"` // 主选
	ArrivalStationName   *string  `json:"arrival_station_name,omitempty"`   // 主选
	TotalGrabPrice       *int32   `json:"total_grab_price,omitempty"`       // 单位：分 按照每人每种抢票情况最高价汇总
	TotalServiceFee      *int32   `json:"total_service_fee,omitempty"`      // 单位：分 实时收取的服务费 抢票按照百分比收时，该字段会变更
	TotalGrabServiceFee  *int32   `json:"total_grab_service_fee,omitempty"` // 单位：分 抢票成功时收取的服务费
	CompanyPay           *int32   `json:"company_pay,omitempty"`            // 单位：分 total_grab_price+total_service_fee+total_grab_service_fee总和的公司支付部分
	PersonalPay          *int32   `json:"personal_pay,omitempty"`           // 单位：分 total_grab_price+total_service_fee+total_grab_service_fee总和的个人支付部分
}

type TrainOrderGrabInfoBuilder struct {
	passengerTravelerId     []string // 滴滴ID或travelerId用于关联乘客信息，多个英文逗号隔开
	passengerTravelerIdSet  bool
	selectedTime            []string // 格式：yyyy-MM-dd  ，多个英文逗号隔开，包含主选，主选是第一个。
	selectedTimeSet         bool
	selectedTrainCode       []string // 多个英文逗号隔开
	selectedTrainCodeSet    bool
	selectedSeatType        []string // 多个英文逗号隔开，字母
	selectedSeatTypeSet     bool
	seatTypeName            []string // 多个英文逗号隔开，名称
	seatTypeNameSet         bool
	departureCityName       string // 城市中文名，主选
	departureCityNameSet    bool
	departureCityId         string // 滴滴城市ID，主选
	departureCityIdSet      bool
	arrivalCityName         string // 城市中文名，主选
	arrivalCityNameSet      bool
	arrivalCityId           string // 滴滴城市ID，主选
	arrivalCityIdSet        bool
	departureStationName    string // 主选
	departureStationNameSet bool
	arrivalStationName      string // 主选
	arrivalStationNameSet   bool
	totalGrabPrice          int32 // 单位：分 按照每人每种抢票情况最高价汇总
	totalGrabPriceSet       bool
	totalServiceFee         int32 // 单位：分 实时收取的服务费 抢票按照百分比收时，该字段会变更
	totalServiceFeeSet      bool
	totalGrabServiceFee     int32 // 单位：分 抢票成功时收取的服务费
	totalGrabServiceFeeSet  bool
	companyPay              int32 // 单位：分 total_grab_price+total_service_fee+total_grab_service_fee总和的公司支付部分
	companyPaySet           bool
	personalPay             int32 // 单位：分 total_grab_price+total_service_fee+total_grab_service_fee总和的个人支付部分
	personalPaySet          bool
}

func NewTrainOrderGrabInfoBuilder() *TrainOrderGrabInfoBuilder {
	return &TrainOrderGrabInfoBuilder{}
}
func (builder *TrainOrderGrabInfoBuilder) PassengerTravelerId(passengerTravelerId []string) *TrainOrderGrabInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) SelectedTime(selectedTime []string) *TrainOrderGrabInfoBuilder {
	builder.selectedTime = selectedTime
	builder.selectedTimeSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) SelectedTrainCode(selectedTrainCode []string) *TrainOrderGrabInfoBuilder {
	builder.selectedTrainCode = selectedTrainCode
	builder.selectedTrainCodeSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) SelectedSeatType(selectedSeatType []string) *TrainOrderGrabInfoBuilder {
	builder.selectedSeatType = selectedSeatType
	builder.selectedSeatTypeSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) SeatTypeName(seatTypeName []string) *TrainOrderGrabInfoBuilder {
	builder.seatTypeName = seatTypeName
	builder.seatTypeNameSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) DepartureCityName(departureCityName string) *TrainOrderGrabInfoBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) DepartureCityId(departureCityId string) *TrainOrderGrabInfoBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) ArrivalCityName(arrivalCityName string) *TrainOrderGrabInfoBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) ArrivalCityId(arrivalCityId string) *TrainOrderGrabInfoBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) DepartureStationName(departureStationName string) *TrainOrderGrabInfoBuilder {
	builder.departureStationName = departureStationName
	builder.departureStationNameSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) ArrivalStationName(arrivalStationName string) *TrainOrderGrabInfoBuilder {
	builder.arrivalStationName = arrivalStationName
	builder.arrivalStationNameSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) TotalGrabPrice(totalGrabPrice int32) *TrainOrderGrabInfoBuilder {
	builder.totalGrabPrice = totalGrabPrice
	builder.totalGrabPriceSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) TotalServiceFee(totalServiceFee int32) *TrainOrderGrabInfoBuilder {
	builder.totalServiceFee = totalServiceFee
	builder.totalServiceFeeSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) TotalGrabServiceFee(totalGrabServiceFee int32) *TrainOrderGrabInfoBuilder {
	builder.totalGrabServiceFee = totalGrabServiceFee
	builder.totalGrabServiceFeeSet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) CompanyPay(companyPay int32) *TrainOrderGrabInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *TrainOrderGrabInfoBuilder) PersonalPay(personalPay int32) *TrainOrderGrabInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}

func (builder *TrainOrderGrabInfoBuilder) Build() *TrainOrderGrabInfo {
	data := &TrainOrderGrabInfo{}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = builder.passengerTravelerId
	}
	if builder.selectedTimeSet {
		data.SelectedTime = builder.selectedTime
	}
	if builder.selectedTrainCodeSet {
		data.SelectedTrainCode = builder.selectedTrainCode
	}
	if builder.selectedSeatTypeSet {
		data.SelectedSeatType = builder.selectedSeatType
	}
	if builder.seatTypeNameSet {
		data.SeatTypeName = builder.seatTypeName
	}
	if builder.departureCityNameSet {
		data.DepartureCityName = &builder.departureCityName
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.arrivalCityNameSet {
		data.ArrivalCityName = &builder.arrivalCityName
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	if builder.departureStationNameSet {
		data.DepartureStationName = &builder.departureStationName
	}
	if builder.arrivalStationNameSet {
		data.ArrivalStationName = &builder.arrivalStationName
	}
	if builder.totalGrabPriceSet {
		data.TotalGrabPrice = &builder.totalGrabPrice
	}
	if builder.totalServiceFeeSet {
		data.TotalServiceFee = &builder.totalServiceFee
	}
	if builder.totalGrabServiceFeeSet {
		data.TotalGrabServiceFee = &builder.totalGrabServiceFee
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	return data
}
