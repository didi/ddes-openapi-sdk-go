package v1

// PersonalPayDetail 支付金额明细 ~ 机票订单、酒店订单共用
type PersonalPayDetail struct {
	ReimburseFee   *int32 `json:"reimburse_fee,omitempty"`    // 可报销金额
	NoReimburseFee *int32 `json:"no_reimburse_fee,omitempty"` // 不可报销金额
}

type PersonalPayDetailBuilder struct {
	reimburseFee      int32 // 可报销金额
	reimburseFeeSet   bool
	noReimburseFee    int32 // 不可报销金额
	noReimburseFeeSet bool
}

func NewPersonalPayDetailBuilder() *PersonalPayDetailBuilder {
	return &PersonalPayDetailBuilder{}
}
func (builder *PersonalPayDetailBuilder) ReimburseFee(reimburseFee int32) *PersonalPayDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *PersonalPayDetailBuilder) NoReimburseFee(noReimburseFee int32) *PersonalPayDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}

func (builder *PersonalPayDetailBuilder) Build() *PersonalPayDetail {
	data := &PersonalPayDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	return data
}
