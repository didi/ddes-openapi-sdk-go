package v1

// PersonalRefundDetail 退票信息（国内）
type PersonalRefundDetail struct {
	ReimburseFee   *int32 `json:"reimburse_fee,omitempty"`    // 可报销金额
	NoReimburseFee *int32 `json:"no_reimburse_fee,omitempty"` // 不可报销金额
}

type PersonalRefundDetailBuilder struct {
	reimburseFee      int32 // 可报销金额
	reimburseFeeSet   bool
	noReimburseFee    int32 // 不可报销金额
	noReimburseFeeSet bool
}

func NewPersonalRefundDetailBuilder() *PersonalRefundDetailBuilder {
	return &PersonalRefundDetailBuilder{}
}
func (builder *PersonalRefundDetailBuilder) ReimburseFee(reimburseFee int32) *PersonalRefundDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *PersonalRefundDetailBuilder) NoReimburseFee(noReimburseFee int32) *PersonalRefundDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}

func (builder *PersonalRefundDetailBuilder) Build() *PersonalRefundDetail {
	data := &PersonalRefundDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	return data
}
