package v1

// HotelOrderRcInfo struct for HotelOrderRcInfo
type HotelOrderRcInfo struct {
	RcId      *string  `json:"rc_id,omitempty"`     // RC审批单ID
	RcType    *string  `json:"rc_type,omitempty"`   // RC类型，枚举值文本：房费RC，星级RC，提前预定RC,协议价RC
	RcCode    *string  `json:"rc_code,omitempty"`   // RCcode
	RcReason  *string  `json:"rc_reason,omitempty"` // RC原因
	Sequences []string `json:"sequences,omitempty"` // 房间序号，分房在订单下标记房间的唯一性
}

type HotelOrderRcInfoBuilder struct {
	rcId         string // RC审批单ID
	rcIdSet      bool
	rcType       string // RC类型，枚举值文本：房费RC，星级RC，提前预定RC,协议价RC
	rcTypeSet    bool
	rcCode       string // RCcode
	rcCodeSet    bool
	rcReason     string // RC原因
	rcReasonSet  bool
	sequences    []string // 房间序号，分房在订单下标记房间的唯一性
	sequencesSet bool
}

func NewHotelOrderRcInfoBuilder() *HotelOrderRcInfoBuilder {
	return &HotelOrderRcInfoBuilder{}
}
func (builder *HotelOrderRcInfoBuilder) RcId(rcId string) *HotelOrderRcInfoBuilder {
	builder.rcId = rcId
	builder.rcIdSet = true
	return builder
}
func (builder *HotelOrderRcInfoBuilder) RcType(rcType string) *HotelOrderRcInfoBuilder {
	builder.rcType = rcType
	builder.rcTypeSet = true
	return builder
}
func (builder *HotelOrderRcInfoBuilder) RcCode(rcCode string) *HotelOrderRcInfoBuilder {
	builder.rcCode = rcCode
	builder.rcCodeSet = true
	return builder
}
func (builder *HotelOrderRcInfoBuilder) RcReason(rcReason string) *HotelOrderRcInfoBuilder {
	builder.rcReason = rcReason
	builder.rcReasonSet = true
	return builder
}
func (builder *HotelOrderRcInfoBuilder) Sequences(sequences []string) *HotelOrderRcInfoBuilder {
	builder.sequences = sequences
	builder.sequencesSet = true
	return builder
}

func (builder *HotelOrderRcInfoBuilder) Build() *HotelOrderRcInfo {
	data := &HotelOrderRcInfo{}
	if builder.rcIdSet {
		data.RcId = &builder.rcId
	}
	if builder.rcTypeSet {
		data.RcType = &builder.rcType
	}
	if builder.rcCodeSet {
		data.RcCode = &builder.rcCode
	}
	if builder.rcReasonSet {
		data.RcReason = &builder.rcReason
	}
	if builder.sequencesSet {
		data.Sequences = builder.sequences
	}
	return data
}
