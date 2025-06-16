package v1

// PersonalRealPayDetail 实付支付金额明细 ~ 机票订单、酒店订单共用
type PersonalRealPayDetail struct {
	ReimburseFee       *int32 `json:"reimburse_fee,omitempty"`         // 可报销金额
	NoReimburseFee     *int32 `json:"no_reimburse_fee,omitempty"`      // 不可报销金额
	CompanyCardRealPay *int32 `json:"company_card_real_pay,omitempty"` // * 酒店订单所含字段
}

type PersonalRealPayDetailBuilder struct {
	reimburseFee          int32 // 可报销金额
	reimburseFeeSet       bool
	noReimburseFee        int32 // 不可报销金额
	noReimburseFeeSet     bool
	companyCardRealPay    int32 // * 酒店订单所含字段
	companyCardRealPaySet bool
}

func NewPersonalRealPayDetailBuilder() *PersonalRealPayDetailBuilder {
	return &PersonalRealPayDetailBuilder{}
}
func (builder *PersonalRealPayDetailBuilder) ReimburseFee(reimburseFee int32) *PersonalRealPayDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *PersonalRealPayDetailBuilder) NoReimburseFee(noReimburseFee int32) *PersonalRealPayDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}
func (builder *PersonalRealPayDetailBuilder) CompanyCardRealPay(companyCardRealPay int32) *PersonalRealPayDetailBuilder {
	builder.companyCardRealPay = companyCardRealPay
	builder.companyCardRealPaySet = true
	return builder
}

func (builder *PersonalRealPayDetailBuilder) Build() *PersonalRealPayDetail {
	data := &PersonalRealPayDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	if builder.companyCardRealPaySet {
		data.CompanyCardRealPay = &builder.companyCardRealPay
	}
	return data
}
