package v1

// DomesticFlightOrderPersonalPayDetail struct for DomesticFlightOrderPersonalPayDetail
type DomesticFlightOrderPersonalPayDetail struct {
	ReimburseFee   *int32 `json:"reimburse_fee,omitempty"`    // 可报销金额
	NoReimburseFee *int32 `json:"no_reimburse_fee,omitempty"` // 不可报销金额
}

type DomesticFlightOrderPersonalPayDetailBuilder struct {
	reimburseFee      int32 // 可报销金额
	reimburseFeeSet   bool
	noReimburseFee    int32 // 不可报销金额
	noReimburseFeeSet bool
}

func NewDomesticFlightOrderPersonalPayDetailBuilder() *DomesticFlightOrderPersonalPayDetailBuilder {
	return &DomesticFlightOrderPersonalPayDetailBuilder{}
}
func (builder *DomesticFlightOrderPersonalPayDetailBuilder) ReimburseFee(reimburseFee int32) *DomesticFlightOrderPersonalPayDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderPersonalPayDetailBuilder) NoReimburseFee(noReimburseFee int32) *DomesticFlightOrderPersonalPayDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}

func (builder *DomesticFlightOrderPersonalPayDetailBuilder) Build() *DomesticFlightOrderPersonalPayDetail {
	data := &DomesticFlightOrderPersonalPayDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	return data
}
