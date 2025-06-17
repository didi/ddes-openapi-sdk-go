package v1

// TrainOrderChangeInfo 改签信息消息类型
type TrainOrderChangeInfo struct {
	TicketUniqueKey      *string `json:"ticket_unique_key,omitempty"`      // 火车票ID，滴滴火车票内部唯一主键，改签新票 改签成功的信息，包含0元改签
	PassengerTravelerId  *string `json:"passenger_traveler_id,omitempty"`  // 乘客ID，滴滴ID或travelerId用于关联乘客信息
	PreTicketUniqueKey   *string `json:"pre_ticket_unique_key,omitempty"`  // 改签前火车票ID，被改签的火车票ID
	ChangeOrderId        *string `json:"change_order_id,omitempty"`        // 改签处理订单号，滴滴内部的改签订单号
	ReasonForChange      *string `json:"reason_for_change,omitempty"`      // 改签原因
	ChangedTime          *string `json:"changed_time,omitempty"`           // 改签成功时间，格式：yyyy-MM-dd HH:mm:ss
	StartTime            *string `json:"start_time,omitempty"`             // 预计发车时间，格式：yyyy-MM-dd HH:mm:ss
	ArriveTime           *string `json:"arrive_time,omitempty"`            // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	TrainCode            *string `json:"train_code,omitempty"`             // 车次编号
	TicketStatus         *string `json:"ticket_status,omitempty"`          // 车票状态，枚举英文code：Holding 占座中；Held 占座成功；Cancelled 已取消；Ticketing 出票中；Openforuse 待乘车；Outbound 已出站；refunding 退票中；Refunded 已退票；Changing 改签中；Changed 已改签；Waitforapproval 待审批；Refoundedother 其它平台退票；Chagnedother 其它平台改签；Refoundedafterchagnedother 其它平台退改；Unknow 未知
	SeatType             *string `json:"seat_type,omitempty"`              // 座位席别，展示中文名，无返回时，显示“其他”
	CoachNumber          *string `json:"coach_number,omitempty"`           // 车厢号，示例：09车厢
	SeatNo               *string `json:"seat_no,omitempty"`                // 座位号，车票座位号，示例：03D号
	DepartureCityName    *string `json:"departure_city_name,omitempty"`    // 出发城市名称，城市中文名
	DepartureCityId      *string `json:"departure_city_id,omitempty"`      // 出发城市ID，滴滴城市ID
	ArrivalCityName      *string `json:"arrival_city_name,omitempty"`      // 到达城市名称，城市中文名
	ArrivalCityId        *string `json:"arrival_city_id,omitempty"`        // 到达城市ID，滴滴城市ID
	DepartureStationName *string `json:"departure_station_name,omitempty"` // 出发车站名称
	ArrivalStationName   *string `json:"arrival_station_name,omitempty"`   // 到达车站名称
	ChangeHandingFee     *int32  `json:"change_handing_fee,omitempty"`     // 改签手续费，单位：分 改签总手续费包含改签费和退票差价费
	ChangeTicketPrice    *int32  `json:"change_ticket_price,omitempty"`    // 改签新票票款，单位：分
	ChangeServiceFee     *int32  `json:"change_service_fee,omitempty"`     // 改签平台使用费，单位：分
	ChangeCompanyPay     *int32  `json:"change_company_pay,omitempty"`     // 改签企业支付金额，单位：分 改签车票公司支付部分
	ChangePersonalPay    *int32  `json:"change_personal_pay,omitempty"`    // 改签个人支付金额，单位：分 改签车票个人支付部分
}

type TrainOrderChangeInfoBuilder struct {
	ticketUniqueKey         string // 火车票ID，滴滴火车票内部唯一主键，改签新票 改签成功的信息，包含0元改签
	ticketUniqueKeySet      bool
	passengerTravelerId     string // 乘客ID，滴滴ID或travelerId用于关联乘客信息
	passengerTravelerIdSet  bool
	preTicketUniqueKey      string // 改签前火车票ID，被改签的火车票ID
	preTicketUniqueKeySet   bool
	changeOrderId           string // 改签处理订单号，滴滴内部的改签订单号
	changeOrderIdSet        bool
	reasonForChange         string // 改签原因
	reasonForChangeSet      bool
	changedTime             string // 改签成功时间，格式：yyyy-MM-dd HH:mm:ss
	changedTimeSet          bool
	startTime               string // 预计发车时间，格式：yyyy-MM-dd HH:mm:ss
	startTimeSet            bool
	arriveTime              string // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	arriveTimeSet           bool
	trainCode               string // 车次编号
	trainCodeSet            bool
	ticketStatus            string // 车票状态，枚举英文code：Holding 占座中；Held 占座成功；Cancelled 已取消；Ticketing 出票中；Openforuse 待乘车；Outbound 已出站；refunding 退票中；Refunded 已退票；Changing 改签中；Changed 已改签；Waitforapproval 待审批；Refoundedother 其它平台退票；Chagnedother 其它平台改签；Refoundedafterchagnedother 其它平台退改；Unknow 未知
	ticketStatusSet         bool
	seatType                string // 座位席别，展示中文名，无返回时，显示“其他”
	seatTypeSet             bool
	coachNumber             string // 车厢号，示例：09车厢
	coachNumberSet          bool
	seatNo                  string // 座位号，车票座位号，示例：03D号
	seatNoSet               bool
	departureCityName       string // 出发城市名称，城市中文名
	departureCityNameSet    bool
	departureCityId         string // 出发城市ID，滴滴城市ID
	departureCityIdSet      bool
	arrivalCityName         string // 到达城市名称，城市中文名
	arrivalCityNameSet      bool
	arrivalCityId           string // 到达城市ID，滴滴城市ID
	arrivalCityIdSet        bool
	departureStationName    string // 出发车站名称
	departureStationNameSet bool
	arrivalStationName      string // 到达车站名称
	arrivalStationNameSet   bool
	changeHandingFee        int32 // 改签手续费，单位：分 改签总手续费包含改签费和退票差价费
	changeHandingFeeSet     bool
	changeTicketPrice       int32 // 改签新票票款，单位：分
	changeTicketPriceSet    bool
	changeServiceFee        int32 // 改签平台使用费，单位：分
	changeServiceFeeSet     bool
	changeCompanyPay        int32 // 改签企业支付金额，单位：分 改签车票公司支付部分
	changeCompanyPaySet     bool
	changePersonalPay       int32 // 改签个人支付金额，单位：分 改签车票个人支付部分
	changePersonalPaySet    bool
}

func NewTrainOrderChangeInfoBuilder() *TrainOrderChangeInfoBuilder {
	return &TrainOrderChangeInfoBuilder{}
}
func (builder *TrainOrderChangeInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *TrainOrderChangeInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) PassengerTravelerId(passengerTravelerId string) *TrainOrderChangeInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) PreTicketUniqueKey(preTicketUniqueKey string) *TrainOrderChangeInfoBuilder {
	builder.preTicketUniqueKey = preTicketUniqueKey
	builder.preTicketUniqueKeySet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangeOrderId(changeOrderId string) *TrainOrderChangeInfoBuilder {
	builder.changeOrderId = changeOrderId
	builder.changeOrderIdSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ReasonForChange(reasonForChange string) *TrainOrderChangeInfoBuilder {
	builder.reasonForChange = reasonForChange
	builder.reasonForChangeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangedTime(changedTime string) *TrainOrderChangeInfoBuilder {
	builder.changedTime = changedTime
	builder.changedTimeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) StartTime(startTime string) *TrainOrderChangeInfoBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ArriveTime(arriveTime string) *TrainOrderChangeInfoBuilder {
	builder.arriveTime = arriveTime
	builder.arriveTimeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) TrainCode(trainCode string) *TrainOrderChangeInfoBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) TicketStatus(ticketStatus string) *TrainOrderChangeInfoBuilder {
	builder.ticketStatus = ticketStatus
	builder.ticketStatusSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) SeatType(seatType string) *TrainOrderChangeInfoBuilder {
	builder.seatType = seatType
	builder.seatTypeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) CoachNumber(coachNumber string) *TrainOrderChangeInfoBuilder {
	builder.coachNumber = coachNumber
	builder.coachNumberSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) SeatNo(seatNo string) *TrainOrderChangeInfoBuilder {
	builder.seatNo = seatNo
	builder.seatNoSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) DepartureCityName(departureCityName string) *TrainOrderChangeInfoBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) DepartureCityId(departureCityId string) *TrainOrderChangeInfoBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ArrivalCityName(arrivalCityName string) *TrainOrderChangeInfoBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ArrivalCityId(arrivalCityId string) *TrainOrderChangeInfoBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) DepartureStationName(departureStationName string) *TrainOrderChangeInfoBuilder {
	builder.departureStationName = departureStationName
	builder.departureStationNameSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ArrivalStationName(arrivalStationName string) *TrainOrderChangeInfoBuilder {
	builder.arrivalStationName = arrivalStationName
	builder.arrivalStationNameSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangeHandingFee(changeHandingFee int32) *TrainOrderChangeInfoBuilder {
	builder.changeHandingFee = changeHandingFee
	builder.changeHandingFeeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangeTicketPrice(changeTicketPrice int32) *TrainOrderChangeInfoBuilder {
	builder.changeTicketPrice = changeTicketPrice
	builder.changeTicketPriceSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangeServiceFee(changeServiceFee int32) *TrainOrderChangeInfoBuilder {
	builder.changeServiceFee = changeServiceFee
	builder.changeServiceFeeSet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangeCompanyPay(changeCompanyPay int32) *TrainOrderChangeInfoBuilder {
	builder.changeCompanyPay = changeCompanyPay
	builder.changeCompanyPaySet = true
	return builder
}
func (builder *TrainOrderChangeInfoBuilder) ChangePersonalPay(changePersonalPay int32) *TrainOrderChangeInfoBuilder {
	builder.changePersonalPay = changePersonalPay
	builder.changePersonalPaySet = true
	return builder
}

func (builder *TrainOrderChangeInfoBuilder) Build() *TrainOrderChangeInfo {
	data := &TrainOrderChangeInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
	}
	if builder.preTicketUniqueKeySet {
		data.PreTicketUniqueKey = &builder.preTicketUniqueKey
	}
	if builder.changeOrderIdSet {
		data.ChangeOrderId = &builder.changeOrderId
	}
	if builder.reasonForChangeSet {
		data.ReasonForChange = &builder.reasonForChange
	}
	if builder.changedTimeSet {
		data.ChangedTime = &builder.changedTime
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.arriveTimeSet {
		data.ArriveTime = &builder.arriveTime
	}
	if builder.trainCodeSet {
		data.TrainCode = &builder.trainCode
	}
	if builder.ticketStatusSet {
		data.TicketStatus = &builder.ticketStatus
	}
	if builder.seatTypeSet {
		data.SeatType = &builder.seatType
	}
	if builder.coachNumberSet {
		data.CoachNumber = &builder.coachNumber
	}
	if builder.seatNoSet {
		data.SeatNo = &builder.seatNo
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
	if builder.changeHandingFeeSet {
		data.ChangeHandingFee = &builder.changeHandingFee
	}
	if builder.changeTicketPriceSet {
		data.ChangeTicketPrice = &builder.changeTicketPrice
	}
	if builder.changeServiceFeeSet {
		data.ChangeServiceFee = &builder.changeServiceFee
	}
	if builder.changeCompanyPaySet {
		data.ChangeCompanyPay = &builder.changeCompanyPay
	}
	if builder.changePersonalPaySet {
		data.ChangePersonalPay = &builder.changePersonalPay
	}
	return data
}
