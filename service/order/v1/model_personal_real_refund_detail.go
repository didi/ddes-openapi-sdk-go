package v1

// PersonalRealRefundDetail 实付退款金额明细 ~ 机票订单、酒店订单共用
type PersonalRealRefundDetail struct {
	ReimburseFee          *int32 `json:"reimburse_fee,omitempty"`            // 可报销金额
	NoReimburseFee        *int32 `json:"no_reimburse_fee,omitempty"`         // 不可报销金额
	CompanyCardRealRefund *int32 `json:"company_card_real_refund,omitempty"` // 不可报销金额
}

type PersonalRealRefundDetailBuilder struct {
	reimburseFee             int32 // 可报销金额
	reimburseFeeSet          bool
	noReimburseFee           int32 // 不可报销金额
	noReimburseFeeSet        bool
	companyCardRealRefund    int32 // 不可报销金额
	companyCardRealRefundSet bool
}

func NewPersonalRealRefundDetailBuilder() *PersonalRealRefundDetailBuilder {
	return &PersonalRealRefundDetailBuilder{}
}
func (builder *PersonalRealRefundDetailBuilder) ReimburseFee(reimburseFee int32) *PersonalRealRefundDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *PersonalRealRefundDetailBuilder) NoReimburseFee(noReimburseFee int32) *PersonalRealRefundDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}
func (builder *PersonalRealRefundDetailBuilder) CompanyCardRealRefund(companyCardRealRefund int32) *PersonalRealRefundDetailBuilder {
	builder.companyCardRealRefund = companyCardRealRefund
	builder.companyCardRealRefundSet = true
	return builder
}

func (builder *PersonalRealRefundDetailBuilder) Build() *PersonalRealRefundDetail {
	data := &PersonalRealRefundDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	if builder.companyCardRealRefundSet {
		data.CompanyCardRealRefund = &builder.companyCardRealRefund
	}
	return data
}
