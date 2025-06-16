package v1

// InterFlightOrderPersonalPayDetail struct for InterFlightOrderPersonalPayDetail
type InterFlightOrderPersonalPayDetail struct {
	ReimburseFee   *int32 `json:"reimburse_fee,omitempty"`    // 可报销金额
	NoReimburseFee *int32 `json:"no_reimburse_fee,omitempty"` // 不可报销金额
}

type InterFlightOrderPersonalPayDetailBuilder struct {
	reimburseFee      int32 // 可报销金额
	reimburseFeeSet   bool
	noReimburseFee    int32 // 不可报销金额
	noReimburseFeeSet bool
}

func NewInterFlightOrderPersonalPayDetailBuilder() *InterFlightOrderPersonalPayDetailBuilder {
	return &InterFlightOrderPersonalPayDetailBuilder{}
}
func (builder *InterFlightOrderPersonalPayDetailBuilder) ReimburseFee(reimburseFee int32) *InterFlightOrderPersonalPayDetailBuilder {
	builder.reimburseFee = reimburseFee
	builder.reimburseFeeSet = true
	return builder
}
func (builder *InterFlightOrderPersonalPayDetailBuilder) NoReimburseFee(noReimburseFee int32) *InterFlightOrderPersonalPayDetailBuilder {
	builder.noReimburseFee = noReimburseFee
	builder.noReimburseFeeSet = true
	return builder
}

func (builder *InterFlightOrderPersonalPayDetailBuilder) Build() *InterFlightOrderPersonalPayDetail {
	data := &InterFlightOrderPersonalPayDetail{}
	if builder.reimburseFeeSet {
		data.ReimburseFee = &builder.reimburseFee
	}
	if builder.noReimburseFeeSet {
		data.NoReimburseFee = &builder.noReimburseFee
	}
	return data
}
