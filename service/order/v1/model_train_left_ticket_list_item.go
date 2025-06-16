package v1

// TrainLeftTicketListItem 车次详细信息
type TrainLeftTicketListItem struct {
	TrainNo             *string      `json:"train_no,omitempty"`              // 全车次
	TrainCode           *string      `json:"train_code,omitempty"`            // 站车次
	FromStationTelecode *string      `json:"from_station_telecode,omitempty"` // 上车站电报码
	FromStationName     *string      `json:"from_station_name,omitempty"`     // 上车站站名
	ToStationTelecode   *string      `json:"to_station_telecode,omitempty"`   // 下车站电报码
	ToStationName       *string      `json:"to_station_name,omitempty"`       // 下车站站名
	StartStationName    *string      `json:"start_station_name,omitempty"`    // 始发站站名
	EndStationName      *string      `json:"end_station_name,omitempty"`      // 终到站站名
	StartTime           *string      `json:"start_time,omitempty"`            // 上车时间，格式2021-01-01 12:00:00
	ArriveTime          *string      `json:"arrive_time,omitempty"`           // 下车时间，格式2021-01-01 14:23:43
	TravelTime          *int32       `json:"travel_time,omitempty"`           // 历时，单位分钟
	TravelTimeName      *string      `json:"travel_time_name,omitempty"`      // 历时文案，如：历时0小时33分
	SaleFlag            *int32       `json:"sale_flag,omitempty"`             // 是否放票0 表示正常售票; 1 表示不能发售(未到放票时间或暂停发售，不可购买); 2 表示进藏等特殊车辆，下单前需将 saleFlagMsg 信息明确提示给用户。
	SaleFlagMsg         *string      `json:"sale_flag_msg,omitempty"`         // 放票信息sale_flag =1 时表示不能发售原因信息(前端控制车次的历时、票价等信息不显示)。 sale_flag =2 时表示特殊提示信息，下单前需提示用户，用户确认后继续下单。
	IsSupportCard       *int32       `json:"is_support_card,omitempty"`       // 是否支持刷二代身份证乘车: 0 是不支持，1 是支持
	DayDifference       *int32       `json:"day_difference,omitempty"`        // 车次跨t天数，0 代表当天到达，1 代表第二天到达
	TicketData          []TicketData `json:"ticket_data,omitempty"`           // 车次余票信息
	EticketTrainFlag    *int32       `json:"eticket_train_flag,omitempty"`    // 电子票乘车标识，0 是不支持，1 是支持
	SpecialFlag         *int32       `json:"special_flag,omitempty"`          // 是否有静音车厢，0 - 无，1 - 有
	IsStartStation      *int32       `json:"is_start_station,omitempty"`      // 是否为始发站，0 - 经过站，1 - 始发站
	IsEndStation        *int32       `json:"is_end_station,omitempty"`        // 是否为终点站，0 - 经过站，1 - 站点站
}

type TrainLeftTicketListItemBuilder struct {
	trainNo                string // 全车次
	trainNoSet             bool
	trainCode              string // 站车次
	trainCodeSet           bool
	fromStationTelecode    string // 上车站电报码
	fromStationTelecodeSet bool
	fromStationName        string // 上车站站名
	fromStationNameSet     bool
	toStationTelecode      string // 下车站电报码
	toStationTelecodeSet   bool
	toStationName          string // 下车站站名
	toStationNameSet       bool
	startStationName       string // 始发站站名
	startStationNameSet    bool
	endStationName         string // 终到站站名
	endStationNameSet      bool
	startTime              string // 上车时间，格式2021-01-01 12:00:00
	startTimeSet           bool
	arriveTime             string // 下车时间，格式2021-01-01 14:23:43
	arriveTimeSet          bool
	travelTime             int32 // 历时，单位分钟
	travelTimeSet          bool
	travelTimeName         string // 历时文案，如：历时0小时33分
	travelTimeNameSet      bool
	saleFlag               int32 // 是否放票0 表示正常售票; 1 表示不能发售(未到放票时间或暂停发售，不可购买); 2 表示进藏等特殊车辆，下单前需将 saleFlagMsg 信息明确提示给用户。
	saleFlagSet            bool
	saleFlagMsg            string // 放票信息sale_flag =1 时表示不能发售原因信息(前端控制车次的历时、票价等信息不显示)。 sale_flag =2 时表示特殊提示信息，下单前需提示用户，用户确认后继续下单。
	saleFlagMsgSet         bool
	isSupportCard          int32 // 是否支持刷二代身份证乘车: 0 是不支持，1 是支持
	isSupportCardSet       bool
	dayDifference          int32 // 车次跨t天数，0 代表当天到达，1 代表第二天到达
	dayDifferenceSet       bool
	ticketData             []TicketData // 车次余票信息
	ticketDataSet          bool
	eticketTrainFlag       int32 // 电子票乘车标识，0 是不支持，1 是支持
	eticketTrainFlagSet    bool
	specialFlag            int32 // 是否有静音车厢，0 - 无，1 - 有
	specialFlagSet         bool
	isStartStation         int32 // 是否为始发站，0 - 经过站，1 - 始发站
	isStartStationSet      bool
	isEndStation           int32 // 是否为终点站，0 - 经过站，1 - 站点站
	isEndStationSet        bool
}

func NewTrainLeftTicketListItemBuilder() *TrainLeftTicketListItemBuilder {
	return &TrainLeftTicketListItemBuilder{}
}
func (builder *TrainLeftTicketListItemBuilder) TrainNo(trainNo string) *TrainLeftTicketListItemBuilder {
	builder.trainNo = trainNo
	builder.trainNoSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) TrainCode(trainCode string) *TrainLeftTicketListItemBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) FromStationTelecode(fromStationTelecode string) *TrainLeftTicketListItemBuilder {
	builder.fromStationTelecode = fromStationTelecode
	builder.fromStationTelecodeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) FromStationName(fromStationName string) *TrainLeftTicketListItemBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) ToStationTelecode(toStationTelecode string) *TrainLeftTicketListItemBuilder {
	builder.toStationTelecode = toStationTelecode
	builder.toStationTelecodeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) ToStationName(toStationName string) *TrainLeftTicketListItemBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) StartStationName(startStationName string) *TrainLeftTicketListItemBuilder {
	builder.startStationName = startStationName
	builder.startStationNameSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) EndStationName(endStationName string) *TrainLeftTicketListItemBuilder {
	builder.endStationName = endStationName
	builder.endStationNameSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) StartTime(startTime string) *TrainLeftTicketListItemBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) ArriveTime(arriveTime string) *TrainLeftTicketListItemBuilder {
	builder.arriveTime = arriveTime
	builder.arriveTimeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) TravelTime(travelTime int32) *TrainLeftTicketListItemBuilder {
	builder.travelTime = travelTime
	builder.travelTimeSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) TravelTimeName(travelTimeName string) *TrainLeftTicketListItemBuilder {
	builder.travelTimeName = travelTimeName
	builder.travelTimeNameSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) SaleFlag(saleFlag int32) *TrainLeftTicketListItemBuilder {
	builder.saleFlag = saleFlag
	builder.saleFlagSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) SaleFlagMsg(saleFlagMsg string) *TrainLeftTicketListItemBuilder {
	builder.saleFlagMsg = saleFlagMsg
	builder.saleFlagMsgSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) IsSupportCard(isSupportCard int32) *TrainLeftTicketListItemBuilder {
	builder.isSupportCard = isSupportCard
	builder.isSupportCardSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) DayDifference(dayDifference int32) *TrainLeftTicketListItemBuilder {
	builder.dayDifference = dayDifference
	builder.dayDifferenceSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) TicketData(ticketData []TicketData) *TrainLeftTicketListItemBuilder {
	builder.ticketData = ticketData
	builder.ticketDataSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) EticketTrainFlag(eticketTrainFlag int32) *TrainLeftTicketListItemBuilder {
	builder.eticketTrainFlag = eticketTrainFlag
	builder.eticketTrainFlagSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) SpecialFlag(specialFlag int32) *TrainLeftTicketListItemBuilder {
	builder.specialFlag = specialFlag
	builder.specialFlagSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) IsStartStation(isStartStation int32) *TrainLeftTicketListItemBuilder {
	builder.isStartStation = isStartStation
	builder.isStartStationSet = true
	return builder
}
func (builder *TrainLeftTicketListItemBuilder) IsEndStation(isEndStation int32) *TrainLeftTicketListItemBuilder {
	builder.isEndStation = isEndStation
	builder.isEndStationSet = true
	return builder
}

func (builder *TrainLeftTicketListItemBuilder) Build() *TrainLeftTicketListItem {
	data := &TrainLeftTicketListItem{}
	if builder.trainNoSet {
		data.TrainNo = &builder.trainNo
	}
	if builder.trainCodeSet {
		data.TrainCode = &builder.trainCode
	}
	if builder.fromStationTelecodeSet {
		data.FromStationTelecode = &builder.fromStationTelecode
	}
	if builder.fromStationNameSet {
		data.FromStationName = &builder.fromStationName
	}
	if builder.toStationTelecodeSet {
		data.ToStationTelecode = &builder.toStationTelecode
	}
	if builder.toStationNameSet {
		data.ToStationName = &builder.toStationName
	}
	if builder.startStationNameSet {
		data.StartStationName = &builder.startStationName
	}
	if builder.endStationNameSet {
		data.EndStationName = &builder.endStationName
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.arriveTimeSet {
		data.ArriveTime = &builder.arriveTime
	}
	if builder.travelTimeSet {
		data.TravelTime = &builder.travelTime
	}
	if builder.travelTimeNameSet {
		data.TravelTimeName = &builder.travelTimeName
	}
	if builder.saleFlagSet {
		data.SaleFlag = &builder.saleFlag
	}
	if builder.saleFlagMsgSet {
		data.SaleFlagMsg = &builder.saleFlagMsg
	}
	if builder.isSupportCardSet {
		data.IsSupportCard = &builder.isSupportCard
	}
	if builder.dayDifferenceSet {
		data.DayDifference = &builder.dayDifference
	}
	if builder.ticketDataSet {
		data.TicketData = builder.ticketData
	}
	if builder.eticketTrainFlagSet {
		data.EticketTrainFlag = &builder.eticketTrainFlag
	}
	if builder.specialFlagSet {
		data.SpecialFlag = &builder.specialFlag
	}
	if builder.isStartStationSet {
		data.IsStartStation = &builder.isStartStation
	}
	if builder.isEndStationSet {
		data.IsEndStation = &builder.isEndStation
	}
	return data
}
