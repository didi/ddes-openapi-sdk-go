package v1

// TrainOrderRcInfo struct for TrainOrderRcInfo
type TrainOrderRcInfo struct {
	RcId            *string `json:"rc_id,omitempty"`             // RC审批单ID
	RcType          *string `json:"rc_type,omitempty"`           // RC类型，枚举值文本：超标说明
	RcCode          *string `json:"rc_code,omitempty"`           // RCcode
	RcReason        *string `json:"rc_reason,omitempty"`         // RC原因
	TicketUniqueKey *string `json:"ticket_unique_key,omitempty"` // 适用对应RC的车票号，空是全部适用
}

type TrainOrderRcInfoBuilder struct {
	rcId               string // RC审批单ID
	rcIdSet            bool
	rcType             string // RC类型，枚举值文本：超标说明
	rcTypeSet          bool
	rcCode             string // RCcode
	rcCodeSet          bool
	rcReason           string // RC原因
	rcReasonSet        bool
	ticketUniqueKey    string // 适用对应RC的车票号，空是全部适用
	ticketUniqueKeySet bool
}

func NewTrainOrderRcInfoBuilder() *TrainOrderRcInfoBuilder {
	return &TrainOrderRcInfoBuilder{}
}
func (builder *TrainOrderRcInfoBuilder) RcId(rcId string) *TrainOrderRcInfoBuilder {
	builder.rcId = rcId
	builder.rcIdSet = true
	return builder
}
func (builder *TrainOrderRcInfoBuilder) RcType(rcType string) *TrainOrderRcInfoBuilder {
	builder.rcType = rcType
	builder.rcTypeSet = true
	return builder
}
func (builder *TrainOrderRcInfoBuilder) RcCode(rcCode string) *TrainOrderRcInfoBuilder {
	builder.rcCode = rcCode
	builder.rcCodeSet = true
	return builder
}
func (builder *TrainOrderRcInfoBuilder) RcReason(rcReason string) *TrainOrderRcInfoBuilder {
	builder.rcReason = rcReason
	builder.rcReasonSet = true
	return builder
}
func (builder *TrainOrderRcInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *TrainOrderRcInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}

func (builder *TrainOrderRcInfoBuilder) Build() *TrainOrderRcInfo {
	data := &TrainOrderRcInfo{}
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
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	return data
}
