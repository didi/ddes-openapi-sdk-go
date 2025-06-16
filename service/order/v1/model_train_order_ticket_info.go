package v1

// TrainOrderTicketInfo 火车票基本信息
type TrainOrderTicketInfo struct {
	TicketUniqueKey        *string `json:"ticket_unique_key,omitempty"`        // 火车票ID，非票号, 不包含改签票信息，预订时的票张信息
	PassengerTravelerId    *string `json:"passenger_traveler_id,omitempty"`    // 乘客ID，滴滴ID或travelerId用于关联乘客信息
	StartTime              *string `json:"start_time,omitempty"`               // 预计发车时间，格式：yyyy-MM-dd HH:mm:ss
	ArriveTime             *string `json:"arrive_time,omitempty"`              // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	TrainCode              *string `json:"train_code,omitempty"`               // 车次编号
	TicketStatus           *string `json:"ticket_status,omitempty"`            // 车票状态，枚举英文code：Holding 占座中 Held 占座成功 Cancelled 已取消 Ticketing 出票中 Openforuse 待乘车 Outbound 已出站 refunding 退票中 Refunded 已退票 Changing 改签中 Changed 已改签 Waitforapproval 待审批 Refoundedother 其它平台退票 Chagnedother 其它平台改签 Refoundedafterchagnedother 其它平台退改 Unknow 未知
	SeatType               *string `json:"seat_type,omitempty"`                // 座位席别，展示中文名，无返回时，显示“其他”，参考坐席列表
	CoachNumber            *string `json:"coach_number,omitempty"`             // 车厢号，车厢号，示例：09车厢
	SeatNo                 *string `json:"seat_no,omitempty"`                  // 车票座位号，示例：03D号
	DepartureCityName      *string `json:"departure_city_name,omitempty"`      // 出发城市名称，城市中文名
	DepartureCityId        *string `json:"departure_city_id,omitempty"`        // 出发城市ID，滴滴城市ID
	ArrivalCityName        *string `json:"arrival_city_name,omitempty"`        // 到达城市名称，城市中文名
	ArrivalCityId          *string `json:"arrival_city_id,omitempty"`          // 到达城市ID，滴滴城市ID
	DepartureStationName   *string `json:"departure_station_name,omitempty"`   // 出发车站名称
	ArrivalStationName     *string `json:"arrival_station_name,omitempty"`     // 到达车站名称
	SalePrice              *int32  `json:"sale_price,omitempty"`               // 车票票面价，单位：分 票面销售费用 ，实际出票价格
	ServiceFee             *int32  `json:"service_fee,omitempty"`              // 出票平台使用费，单位：分 实时收取的服务费
	AsynchronousServiceFee *int32  `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	GrabServiceFee         *int32  `json:"grab_service_fee,omitempty"`         // 抢票服务费，单位：分 抢票成功时收取的服务费
	CompanyPay             *int32  `json:"company_pay,omitempty"`              // 车票企业支付金额，单位：分 车票公司支付部分
	PersonalPay            *int32  `json:"personal_pay,omitempty"`             // 车票个人支付金额，单位：分 车票个人支付部分
	PrintTicketTime        *string `json:"print_ticket_time,omitempty"`        // 出票时间，格式：yyyy-MM-dd HH:mm:ss
}

type TrainOrderTicketInfoBuilder struct {
	ticketUniqueKey           string // 火车票ID，非票号, 不包含改签票信息，预订时的票张信息
	ticketUniqueKeySet        bool
	passengerTravelerId       string // 乘客ID，滴滴ID或travelerId用于关联乘客信息
	passengerTravelerIdSet    bool
	startTime                 string // 预计发车时间，格式：yyyy-MM-dd HH:mm:ss
	startTimeSet              bool
	arriveTime                string // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	arriveTimeSet             bool
	trainCode                 string // 车次编号
	trainCodeSet              bool
	ticketStatus              string // 车票状态，枚举英文code：Holding 占座中 Held 占座成功 Cancelled 已取消 Ticketing 出票中 Openforuse 待乘车 Outbound 已出站 refunding 退票中 Refunded 已退票 Changing 改签中 Changed 已改签 Waitforapproval 待审批 Refoundedother 其它平台退票 Chagnedother 其它平台改签 Refoundedafterchagnedother 其它平台退改 Unknow 未知
	ticketStatusSet           bool
	seatType                  string // 座位席别，展示中文名，无返回时，显示“其他”，参考坐席列表
	seatTypeSet               bool
	coachNumber               string // 车厢号，车厢号，示例：09车厢
	coachNumberSet            bool
	seatNo                    string // 车票座位号，示例：03D号
	seatNoSet                 bool
	departureCityName         string // 出发城市名称，城市中文名
	departureCityNameSet      bool
	departureCityId           string // 出发城市ID，滴滴城市ID
	departureCityIdSet        bool
	arrivalCityName           string // 到达城市名称，城市中文名
	arrivalCityNameSet        bool
	arrivalCityId             string // 到达城市ID，滴滴城市ID
	arrivalCityIdSet          bool
	departureStationName      string // 出发车站名称
	departureStationNameSet   bool
	arrivalStationName        string // 到达车站名称
	arrivalStationNameSet     bool
	salePrice                 int32 // 车票票面价，单位：分 票面销售费用 ，实际出票价格
	salePriceSet              bool
	serviceFee                int32 // 出票平台使用费，单位：分 实时收取的服务费
	serviceFeeSet             bool
	asynchronousServiceFee    int32 // 出票平台使用费（异步），单位：分
	asynchronousServiceFeeSet bool
	grabServiceFee            int32 // 抢票服务费，单位：分 抢票成功时收取的服务费
	grabServiceFeeSet         bool
	companyPay                int32 // 车票企业支付金额，单位：分 车票公司支付部分
	companyPaySet             bool
	personalPay               int32 // 车票个人支付金额，单位：分 车票个人支付部分
	personalPaySet            bool
	printTicketTime           string // 出票时间，格式：yyyy-MM-dd HH:mm:ss
	printTicketTimeSet        bool
}

func NewTrainOrderTicketInfoBuilder() *TrainOrderTicketInfoBuilder {
	return &TrainOrderTicketInfoBuilder{}
}
func (builder *TrainOrderTicketInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *TrainOrderTicketInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) PassengerTravelerId(passengerTravelerId string) *TrainOrderTicketInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) StartTime(startTime string) *TrainOrderTicketInfoBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) ArriveTime(arriveTime string) *TrainOrderTicketInfoBuilder {
	builder.arriveTime = arriveTime
	builder.arriveTimeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) TrainCode(trainCode string) *TrainOrderTicketInfoBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) TicketStatus(ticketStatus string) *TrainOrderTicketInfoBuilder {
	builder.ticketStatus = ticketStatus
	builder.ticketStatusSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) SeatType(seatType string) *TrainOrderTicketInfoBuilder {
	builder.seatType = seatType
	builder.seatTypeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) CoachNumber(coachNumber string) *TrainOrderTicketInfoBuilder {
	builder.coachNumber = coachNumber
	builder.coachNumberSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) SeatNo(seatNo string) *TrainOrderTicketInfoBuilder {
	builder.seatNo = seatNo
	builder.seatNoSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) DepartureCityName(departureCityName string) *TrainOrderTicketInfoBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) DepartureCityId(departureCityId string) *TrainOrderTicketInfoBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) ArrivalCityName(arrivalCityName string) *TrainOrderTicketInfoBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) ArrivalCityId(arrivalCityId string) *TrainOrderTicketInfoBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) DepartureStationName(departureStationName string) *TrainOrderTicketInfoBuilder {
	builder.departureStationName = departureStationName
	builder.departureStationNameSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) ArrivalStationName(arrivalStationName string) *TrainOrderTicketInfoBuilder {
	builder.arrivalStationName = arrivalStationName
	builder.arrivalStationNameSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) SalePrice(salePrice int32) *TrainOrderTicketInfoBuilder {
	builder.salePrice = salePrice
	builder.salePriceSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) ServiceFee(serviceFee int32) *TrainOrderTicketInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *TrainOrderTicketInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) GrabServiceFee(grabServiceFee int32) *TrainOrderTicketInfoBuilder {
	builder.grabServiceFee = grabServiceFee
	builder.grabServiceFeeSet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) CompanyPay(companyPay int32) *TrainOrderTicketInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) PersonalPay(personalPay int32) *TrainOrderTicketInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *TrainOrderTicketInfoBuilder) PrintTicketTime(printTicketTime string) *TrainOrderTicketInfoBuilder {
	builder.printTicketTime = printTicketTime
	builder.printTicketTimeSet = true
	return builder
}

func (builder *TrainOrderTicketInfoBuilder) Build() *TrainOrderTicketInfo {
	data := &TrainOrderTicketInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
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
	if builder.salePriceSet {
		data.SalePrice = &builder.salePrice
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.asynchronousServiceFeeSet {
		data.AsynchronousServiceFee = &builder.asynchronousServiceFee
	}
	if builder.grabServiceFeeSet {
		data.GrabServiceFee = &builder.grabServiceFee
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.printTicketTimeSet {
		data.PrintTicketTime = &builder.printTicketTime
	}
	return data
}
