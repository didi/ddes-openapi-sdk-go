package v1

// FlightOrderRcInfo 管控信息
type FlightOrderRcInfo struct {
	TicketUniqueKey *string `json:"ticket_unique_key,omitempty"` // 票号唯一值
	RcId            *string `json:"rc_id,omitempty"`             // RC审批单ID
	RcType          *string `json:"rc_type,omitempty"`           // RC类型，枚举值文本：舱等RC，低价RC，提前预定RC，同行程火车耗时超标，同行程火车价格超标
	RcCode          *string `json:"rc_code,omitempty"`           // RCcode
	RcReason        *string `json:"rc_reason,omitempty"`         // RC原因
}

type FlightOrderRcInfoBuilder struct {
	ticketUniqueKey    string // 票号唯一值
	ticketUniqueKeySet bool
	rcId               string // RC审批单ID
	rcIdSet            bool
	rcType             string // RC类型，枚举值文本：舱等RC，低价RC，提前预定RC，同行程火车耗时超标，同行程火车价格超标
	rcTypeSet          bool
	rcCode             string // RCcode
	rcCodeSet          bool
	rcReason           string // RC原因
	rcReasonSet        bool
}

func NewFlightOrderRcInfoBuilder() *FlightOrderRcInfoBuilder {
	return &FlightOrderRcInfoBuilder{}
}
func (builder *FlightOrderRcInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *FlightOrderRcInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *FlightOrderRcInfoBuilder) RcId(rcId string) *FlightOrderRcInfoBuilder {
	builder.rcId = rcId
	builder.rcIdSet = true
	return builder
}
func (builder *FlightOrderRcInfoBuilder) RcType(rcType string) *FlightOrderRcInfoBuilder {
	builder.rcType = rcType
	builder.rcTypeSet = true
	return builder
}
func (builder *FlightOrderRcInfoBuilder) RcCode(rcCode string) *FlightOrderRcInfoBuilder {
	builder.rcCode = rcCode
	builder.rcCodeSet = true
	return builder
}
func (builder *FlightOrderRcInfoBuilder) RcReason(rcReason string) *FlightOrderRcInfoBuilder {
	builder.rcReason = rcReason
	builder.rcReasonSet = true
	return builder
}

func (builder *FlightOrderRcInfoBuilder) Build() *FlightOrderRcInfo {
	data := &FlightOrderRcInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
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
	return data
}
