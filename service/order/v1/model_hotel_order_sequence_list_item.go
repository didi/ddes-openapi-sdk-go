package v1

// HotelOrderSequenceListItem 入住房间列表信息
type HotelOrderSequenceListItem struct {
	Sequence            *string  `json:"sequence,omitempty"`              // 房间序号，分房在订单下标记房间的唯一性
	PassengerTravelerId []string `json:"passenger_traveler_id,omitempty"` // 入住人唯一标记，入住人关联信息
	RealCheckinTime     *string  `json:"real_checkin_time,omitempty"`     // 实际入住时间，格式：yyyy-MM-dd HH:mm:ss
	RealCheckoutTime    *string  `json:"real_checkout_time,omitempty"`    // 实际离店时间，格式：yyyy-MM-dd HH:mm:ss
	SequenceId          *string  `json:"sequence_id,omitempty"`           // * 酒店、国际机票订单字段
}

type HotelOrderSequenceListItemBuilder struct {
	sequence               string // 房间序号，分房在订单下标记房间的唯一性
	sequenceSet            bool
	passengerTravelerId    []string // 入住人唯一标记，入住人关联信息
	passengerTravelerIdSet bool
	realCheckinTime        string // 实际入住时间，格式：yyyy-MM-dd HH:mm:ss
	realCheckinTimeSet     bool
	realCheckoutTime       string // 实际离店时间，格式：yyyy-MM-dd HH:mm:ss
	realCheckoutTimeSet    bool
	sequenceId             string // * 酒店、国际机票订单字段
	sequenceIdSet          bool
}

func NewHotelOrderSequenceListItemBuilder() *HotelOrderSequenceListItemBuilder {
	return &HotelOrderSequenceListItemBuilder{}
}
func (builder *HotelOrderSequenceListItemBuilder) Sequence(sequence string) *HotelOrderSequenceListItemBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}
func (builder *HotelOrderSequenceListItemBuilder) PassengerTravelerId(passengerTravelerId []string) *HotelOrderSequenceListItemBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *HotelOrderSequenceListItemBuilder) RealCheckinTime(realCheckinTime string) *HotelOrderSequenceListItemBuilder {
	builder.realCheckinTime = realCheckinTime
	builder.realCheckinTimeSet = true
	return builder
}
func (builder *HotelOrderSequenceListItemBuilder) RealCheckoutTime(realCheckoutTime string) *HotelOrderSequenceListItemBuilder {
	builder.realCheckoutTime = realCheckoutTime
	builder.realCheckoutTimeSet = true
	return builder
}
func (builder *HotelOrderSequenceListItemBuilder) SequenceId(sequenceId string) *HotelOrderSequenceListItemBuilder {
	builder.sequenceId = sequenceId
	builder.sequenceIdSet = true
	return builder
}

func (builder *HotelOrderSequenceListItemBuilder) Build() *HotelOrderSequenceListItem {
	data := &HotelOrderSequenceListItem{}
	if builder.sequenceSet {
		data.Sequence = &builder.sequence
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = builder.passengerTravelerId
	}
	if builder.realCheckinTimeSet {
		data.RealCheckinTime = &builder.realCheckinTime
	}
	if builder.realCheckoutTimeSet {
		data.RealCheckoutTime = &builder.realCheckoutTime
	}
	if builder.sequenceIdSet {
		data.SequenceId = &builder.sequenceId
	}
	return data
}
