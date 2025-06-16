package v1

// SegmentItem 行程列表项
type SegmentItem struct {
	StartStationName    *string                   `json:"start_station_name,omitempty"`    // 出发站点名称
	ArriveTime          *string                   `json:"arrive_time,omitempty"`           // 到达时间，eg: 2025-01-10 10:23
	DayDifference       *int32                    `json:"day_difference,omitempty"`        // 第一程和到二程跨天的天数
	IsEndStation        *int32                    `json:"is_end_station,omitempty"`        // 是否是终点站，1：是 0：否
	FromCityId          *string                   `json:"from_city_id,omitempty"`          // 出发城市ID
	FromStationTelecode *string                   `json:"from_station_telecode,omitempty"` // 车站点电报码
	ArriveTimestamp     *int32                    `json:"arrive_timestamp,omitempty"`      // 到达时间时间戳，eg： 1736475780
	SaleFlagMsg         *string                   `json:"sale_flag_msg,omitempty"`         // 是否开启售卖，1：是 0：否
	EticketTrainFlag    *string                   `json:"eticket_train_flag,omitempty"`    // 电子票乘车标识 0 是不支持，1 是支持
	TrainType           *string                   `json:"train_type,omitempty"`            // 车次类型
	TrainNo             *string                   `json:"train_no,omitempty"`              // 车次号
	TrainCode           *string                   `json:"train_code,omitempty"`            // 车次号长码
	StartTime           *string                   `json:"start_time,omitempty"`            // 发车时间
	FromStationNo       *string                   `json:"from_station_no,omitempty"`       // 出发站点
	ToCityId            *string                   `json:"to_city_id,omitempty"`            // 到达城市
	SaleFlag            *string                   `json:"sale_flag,omitempty"`             // 是否放票 0表示正常售票;1表示不能发售(未到放票时间或暂停发售，不可购买);2表示进藏等特殊车辆
	IsSupportCard       *string                   `json:"is_support_card,omitempty"`       // 是否支持刷卡，1：是 0：否
	TrainSecrets        *string                   `json:"train_secrets,omitempty"`         // 车票维度系统唯一标识（内部数据流转使用）
	TravelTimeName      *string                   `json:"travel_time_name,omitempty"`      // 格式化行驶时长，eg: 历史4小时15分
	StartTimestamp      *int32                    `json:"start_timestamp,omitempty"`       // 发车时间时间戳
	ToStationNo         *string                   `json:"to_station_no,omitempty"`         // 到达站点
	Sequence            *int32                    `json:"sequence,omitempty"`              // 第几程，第一程：1，第二程：2
	ToStationTelecode   *string                   `json:"to_station_telecode,omitempty"`   // 到达站点电报码
	ToStationName       *string                   `json:"to_station_name,omitempty"`       // 到达站点名称
	TrainTypeCode       *string                   `json:"train_type_code,omitempty"`       // 车次类型编码
	FromStationName     *string                   `json:"from_station_name,omitempty"`     // 出发站点名称
	SeatTypes           *string                   `json:"seat_types,omitempty"`            // 坐席类型
	TravelTime          *int32                    `json:"travel_time,omitempty"`           // 行驶时长，单位：分钟
	SpecialFlag         *int32                    `json:"special_flag,omitempty"`          // 是否存在静音，1：是 0：否
	IsStartStation      *int32                    `json:"is_start_station,omitempty"`      // 是否始发站，1：是 0：否
	EndStationName      *string                   `json:"end_station_name,omitempty"`      // 到达站点名称
	TicketData          []TransferTrainTicketData `json:"ticket_data,omitempty"`           // 票信息
}

type SegmentItemBuilder struct {
	startStationName       string // 出发站点名称
	startStationNameSet    bool
	arriveTime             string // 到达时间，eg: 2025-01-10 10:23
	arriveTimeSet          bool
	dayDifference          int32 // 第一程和到二程跨天的天数
	dayDifferenceSet       bool
	isEndStation           int32 // 是否是终点站，1：是 0：否
	isEndStationSet        bool
	fromCityId             string // 出发城市ID
	fromCityIdSet          bool
	fromStationTelecode    string // 车站点电报码
	fromStationTelecodeSet bool
	arriveTimestamp        int32 // 到达时间时间戳，eg： 1736475780
	arriveTimestampSet     bool
	saleFlagMsg            string // 是否开启售卖，1：是 0：否
	saleFlagMsgSet         bool
	eticketTrainFlag       string // 电子票乘车标识 0 是不支持，1 是支持
	eticketTrainFlagSet    bool
	trainType              string // 车次类型
	trainTypeSet           bool
	trainNo                string // 车次号
	trainNoSet             bool
	trainCode              string // 车次号长码
	trainCodeSet           bool
	startTime              string // 发车时间
	startTimeSet           bool
	fromStationNo          string // 出发站点
	fromStationNoSet       bool
	toCityId               string // 到达城市
	toCityIdSet            bool
	saleFlag               string // 是否放票 0表示正常售票;1表示不能发售(未到放票时间或暂停发售，不可购买);2表示进藏等特殊车辆
	saleFlagSet            bool
	isSupportCard          string // 是否支持刷卡，1：是 0：否
	isSupportCardSet       bool
	trainSecrets           string // 车票维度系统唯一标识（内部数据流转使用）
	trainSecretsSet        bool
	travelTimeName         string // 格式化行驶时长，eg: 历史4小时15分
	travelTimeNameSet      bool
	startTimestamp         int32 // 发车时间时间戳
	startTimestampSet      bool
	toStationNo            string // 到达站点
	toStationNoSet         bool
	sequence               int32 // 第几程，第一程：1，第二程：2
	sequenceSet            bool
	toStationTelecode      string // 到达站点电报码
	toStationTelecodeSet   bool
	toStationName          string // 到达站点名称
	toStationNameSet       bool
	trainTypeCode          string // 车次类型编码
	trainTypeCodeSet       bool
	fromStationName        string // 出发站点名称
	fromStationNameSet     bool
	seatTypes              string // 坐席类型
	seatTypesSet           bool
	travelTime             int32 // 行驶时长，单位：分钟
	travelTimeSet          bool
	specialFlag            int32 // 是否存在静音，1：是 0：否
	specialFlagSet         bool
	isStartStation         int32 // 是否始发站，1：是 0：否
	isStartStationSet      bool
	endStationName         string // 到达站点名称
	endStationNameSet      bool
	ticketData             []TransferTrainTicketData // 票信息
	ticketDataSet          bool
}

func NewSegmentItemBuilder() *SegmentItemBuilder {
	return &SegmentItemBuilder{}
}
func (builder *SegmentItemBuilder) StartStationName(startStationName string) *SegmentItemBuilder {
	builder.startStationName = startStationName
	builder.startStationNameSet = true
	return builder
}
func (builder *SegmentItemBuilder) ArriveTime(arriveTime string) *SegmentItemBuilder {
	builder.arriveTime = arriveTime
	builder.arriveTimeSet = true
	return builder
}
func (builder *SegmentItemBuilder) DayDifference(dayDifference int32) *SegmentItemBuilder {
	builder.dayDifference = dayDifference
	builder.dayDifferenceSet = true
	return builder
}
func (builder *SegmentItemBuilder) IsEndStation(isEndStation int32) *SegmentItemBuilder {
	builder.isEndStation = isEndStation
	builder.isEndStationSet = true
	return builder
}
func (builder *SegmentItemBuilder) FromCityId(fromCityId string) *SegmentItemBuilder {
	builder.fromCityId = fromCityId
	builder.fromCityIdSet = true
	return builder
}
func (builder *SegmentItemBuilder) FromStationTelecode(fromStationTelecode string) *SegmentItemBuilder {
	builder.fromStationTelecode = fromStationTelecode
	builder.fromStationTelecodeSet = true
	return builder
}
func (builder *SegmentItemBuilder) ArriveTimestamp(arriveTimestamp int32) *SegmentItemBuilder {
	builder.arriveTimestamp = arriveTimestamp
	builder.arriveTimestampSet = true
	return builder
}
func (builder *SegmentItemBuilder) SaleFlagMsg(saleFlagMsg string) *SegmentItemBuilder {
	builder.saleFlagMsg = saleFlagMsg
	builder.saleFlagMsgSet = true
	return builder
}
func (builder *SegmentItemBuilder) EticketTrainFlag(eticketTrainFlag string) *SegmentItemBuilder {
	builder.eticketTrainFlag = eticketTrainFlag
	builder.eticketTrainFlagSet = true
	return builder
}
func (builder *SegmentItemBuilder) TrainType(trainType string) *SegmentItemBuilder {
	builder.trainType = trainType
	builder.trainTypeSet = true
	return builder
}
func (builder *SegmentItemBuilder) TrainNo(trainNo string) *SegmentItemBuilder {
	builder.trainNo = trainNo
	builder.trainNoSet = true
	return builder
}
func (builder *SegmentItemBuilder) TrainCode(trainCode string) *SegmentItemBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *SegmentItemBuilder) StartTime(startTime string) *SegmentItemBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *SegmentItemBuilder) FromStationNo(fromStationNo string) *SegmentItemBuilder {
	builder.fromStationNo = fromStationNo
	builder.fromStationNoSet = true
	return builder
}
func (builder *SegmentItemBuilder) ToCityId(toCityId string) *SegmentItemBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *SegmentItemBuilder) SaleFlag(saleFlag string) *SegmentItemBuilder {
	builder.saleFlag = saleFlag
	builder.saleFlagSet = true
	return builder
}
func (builder *SegmentItemBuilder) IsSupportCard(isSupportCard string) *SegmentItemBuilder {
	builder.isSupportCard = isSupportCard
	builder.isSupportCardSet = true
	return builder
}
func (builder *SegmentItemBuilder) TrainSecrets(trainSecrets string) *SegmentItemBuilder {
	builder.trainSecrets = trainSecrets
	builder.trainSecretsSet = true
	return builder
}
func (builder *SegmentItemBuilder) TravelTimeName(travelTimeName string) *SegmentItemBuilder {
	builder.travelTimeName = travelTimeName
	builder.travelTimeNameSet = true
	return builder
}
func (builder *SegmentItemBuilder) StartTimestamp(startTimestamp int32) *SegmentItemBuilder {
	builder.startTimestamp = startTimestamp
	builder.startTimestampSet = true
	return builder
}
func (builder *SegmentItemBuilder) ToStationNo(toStationNo string) *SegmentItemBuilder {
	builder.toStationNo = toStationNo
	builder.toStationNoSet = true
	return builder
}
func (builder *SegmentItemBuilder) Sequence(sequence int32) *SegmentItemBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}
func (builder *SegmentItemBuilder) ToStationTelecode(toStationTelecode string) *SegmentItemBuilder {
	builder.toStationTelecode = toStationTelecode
	builder.toStationTelecodeSet = true
	return builder
}
func (builder *SegmentItemBuilder) ToStationName(toStationName string) *SegmentItemBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *SegmentItemBuilder) TrainTypeCode(trainTypeCode string) *SegmentItemBuilder {
	builder.trainTypeCode = trainTypeCode
	builder.trainTypeCodeSet = true
	return builder
}
func (builder *SegmentItemBuilder) FromStationName(fromStationName string) *SegmentItemBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *SegmentItemBuilder) SeatTypes(seatTypes string) *SegmentItemBuilder {
	builder.seatTypes = seatTypes
	builder.seatTypesSet = true
	return builder
}
func (builder *SegmentItemBuilder) TravelTime(travelTime int32) *SegmentItemBuilder {
	builder.travelTime = travelTime
	builder.travelTimeSet = true
	return builder
}
func (builder *SegmentItemBuilder) SpecialFlag(specialFlag int32) *SegmentItemBuilder {
	builder.specialFlag = specialFlag
	builder.specialFlagSet = true
	return builder
}
func (builder *SegmentItemBuilder) IsStartStation(isStartStation int32) *SegmentItemBuilder {
	builder.isStartStation = isStartStation
	builder.isStartStationSet = true
	return builder
}
func (builder *SegmentItemBuilder) EndStationName(endStationName string) *SegmentItemBuilder {
	builder.endStationName = endStationName
	builder.endStationNameSet = true
	return builder
}
func (builder *SegmentItemBuilder) TicketData(ticketData []TransferTrainTicketData) *SegmentItemBuilder {
	builder.ticketData = ticketData
	builder.ticketDataSet = true
	return builder
}

func (builder *SegmentItemBuilder) Build() *SegmentItem {
	data := &SegmentItem{}
	if builder.startStationNameSet {
		data.StartStationName = &builder.startStationName
	}
	if builder.arriveTimeSet {
		data.ArriveTime = &builder.arriveTime
	}
	if builder.dayDifferenceSet {
		data.DayDifference = &builder.dayDifference
	}
	if builder.isEndStationSet {
		data.IsEndStation = &builder.isEndStation
	}
	if builder.fromCityIdSet {
		data.FromCityId = &builder.fromCityId
	}
	if builder.fromStationTelecodeSet {
		data.FromStationTelecode = &builder.fromStationTelecode
	}
	if builder.arriveTimestampSet {
		data.ArriveTimestamp = &builder.arriveTimestamp
	}
	if builder.saleFlagMsgSet {
		data.SaleFlagMsg = &builder.saleFlagMsg
	}
	if builder.eticketTrainFlagSet {
		data.EticketTrainFlag = &builder.eticketTrainFlag
	}
	if builder.trainTypeSet {
		data.TrainType = &builder.trainType
	}
	if builder.trainNoSet {
		data.TrainNo = &builder.trainNo
	}
	if builder.trainCodeSet {
		data.TrainCode = &builder.trainCode
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.fromStationNoSet {
		data.FromStationNo = &builder.fromStationNo
	}
	if builder.toCityIdSet {
		data.ToCityId = &builder.toCityId
	}
	if builder.saleFlagSet {
		data.SaleFlag = &builder.saleFlag
	}
	if builder.isSupportCardSet {
		data.IsSupportCard = &builder.isSupportCard
	}
	if builder.trainSecretsSet {
		data.TrainSecrets = &builder.trainSecrets
	}
	if builder.travelTimeNameSet {
		data.TravelTimeName = &builder.travelTimeName
	}
	if builder.startTimestampSet {
		data.StartTimestamp = &builder.startTimestamp
	}
	if builder.toStationNoSet {
		data.ToStationNo = &builder.toStationNo
	}
	if builder.sequenceSet {
		data.Sequence = &builder.sequence
	}
	if builder.toStationTelecodeSet {
		data.ToStationTelecode = &builder.toStationTelecode
	}
	if builder.toStationNameSet {
		data.ToStationName = &builder.toStationName
	}
	if builder.trainTypeCodeSet {
		data.TrainTypeCode = &builder.trainTypeCode
	}
	if builder.fromStationNameSet {
		data.FromStationName = &builder.fromStationName
	}
	if builder.seatTypesSet {
		data.SeatTypes = &builder.seatTypes
	}
	if builder.travelTimeSet {
		data.TravelTime = &builder.travelTime
	}
	if builder.specialFlagSet {
		data.SpecialFlag = &builder.specialFlag
	}
	if builder.isStartStationSet {
		data.IsStartStation = &builder.isStartStation
	}
	if builder.endStationNameSet {
		data.EndStationName = &builder.endStationName
	}
	if builder.ticketDataSet {
		data.TicketData = builder.ticketData
	}
	return data
}
