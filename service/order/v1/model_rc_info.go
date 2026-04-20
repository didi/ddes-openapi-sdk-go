package v1

// RcInfo 管控信息
type RcInfo struct {
	RcType   *string `json:"rc_type,omitempty"`   // 超标类型
	RcCode   *string `json:"rc_code,omitempty"`   // 超标编码
	RcReason *string `json:"rc_reason,omitempty"` // 超标原因
	RcRemark *string `json:"rc_remark,omitempty"` // 超标备注
}

type RcInfoBuilder struct {
	rcType      string // 超标类型
	rcTypeSet   bool
	rcCode      string // 超标编码
	rcCodeSet   bool
	rcReason    string // 超标原因
	rcReasonSet bool
	rcRemark    string // 超标备注
	rcRemarkSet bool
}

func NewRcInfoBuilder() *RcInfoBuilder {
	return &RcInfoBuilder{}
}
func (builder *RcInfoBuilder) RcType(rcType string) *RcInfoBuilder {
	builder.rcType = rcType
	builder.rcTypeSet = true
	return builder
}
func (builder *RcInfoBuilder) RcCode(rcCode string) *RcInfoBuilder {
	builder.rcCode = rcCode
	builder.rcCodeSet = true
	return builder
}
func (builder *RcInfoBuilder) RcReason(rcReason string) *RcInfoBuilder {
	builder.rcReason = rcReason
	builder.rcReasonSet = true
	return builder
}
func (builder *RcInfoBuilder) RcRemark(rcRemark string) *RcInfoBuilder {
	builder.rcRemark = rcRemark
	builder.rcRemarkSet = true
	return builder
}

func (builder *RcInfoBuilder) Build() *RcInfo {
	data := &RcInfo{}
	if builder.rcTypeSet {
		data.RcType = &builder.rcType
	}
	if builder.rcCodeSet {
		data.RcCode = &builder.rcCode
	}
	if builder.rcReasonSet {
		data.RcReason = &builder.rcReason
	}
	if builder.rcRemarkSet {
		data.RcRemark = &builder.rcRemark
	}
	return data
}
