package v1

// TrainOrderRefundInfo 退票信息
type TrainOrderRefundInfo struct {
	TicketUniqueKey  *string `json:"ticket_unique_key,omitempty"`  // 火车票ID，滴滴火车票内部唯一主键，退票成功的信息
	RefundHandleFee  *int32  `json:"refund_handle_fee,omitempty"`  // 退票手续费，单位：分
	UserRefund       *int32  `json:"user_refund,omitempty"`        // 退还用户金额，单位：分
	RefundServiceFee *int32  `json:"refund_service_fee,omitempty"` // 退票平台使用费，单位：分
	RefundedTime     *string `json:"refunded_time,omitempty"`      // 退票成功时间，格式：yyyy-MM-dd HH:mm:ss
	ReasonForRefund  *string `json:"reason_for_refund,omitempty"`  // 退票原因
	CompanyRefund    *int32  `json:"company_refund,omitempty"`     // 企业退款金额，单位：分
	PersonalRefund   *int32  `json:"personal_refund,omitempty"`    // 个人退款金额，单位：分
}

type TrainOrderRefundInfoBuilder struct {
	ticketUniqueKey     string // 火车票ID，滴滴火车票内部唯一主键，退票成功的信息
	ticketUniqueKeySet  bool
	refundHandleFee     int32 // 退票手续费，单位：分
	refundHandleFeeSet  bool
	userRefund          int32 // 退还用户金额，单位：分
	userRefundSet       bool
	refundServiceFee    int32 // 退票平台使用费，单位：分
	refundServiceFeeSet bool
	refundedTime        string // 退票成功时间，格式：yyyy-MM-dd HH:mm:ss
	refundedTimeSet     bool
	reasonForRefund     string // 退票原因
	reasonForRefundSet  bool
	companyRefund       int32 // 企业退款金额，单位：分
	companyRefundSet    bool
	personalRefund      int32 // 个人退款金额，单位：分
	personalRefundSet   bool
}

func NewTrainOrderRefundInfoBuilder() *TrainOrderRefundInfoBuilder {
	return &TrainOrderRefundInfoBuilder{}
}
func (builder *TrainOrderRefundInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *TrainOrderRefundInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) RefundHandleFee(refundHandleFee int32) *TrainOrderRefundInfoBuilder {
	builder.refundHandleFee = refundHandleFee
	builder.refundHandleFeeSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) UserRefund(userRefund int32) *TrainOrderRefundInfoBuilder {
	builder.userRefund = userRefund
	builder.userRefundSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) RefundServiceFee(refundServiceFee int32) *TrainOrderRefundInfoBuilder {
	builder.refundServiceFee = refundServiceFee
	builder.refundServiceFeeSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) RefundedTime(refundedTime string) *TrainOrderRefundInfoBuilder {
	builder.refundedTime = refundedTime
	builder.refundedTimeSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) ReasonForRefund(reasonForRefund string) *TrainOrderRefundInfoBuilder {
	builder.reasonForRefund = reasonForRefund
	builder.reasonForRefundSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) CompanyRefund(companyRefund int32) *TrainOrderRefundInfoBuilder {
	builder.companyRefund = companyRefund
	builder.companyRefundSet = true
	return builder
}
func (builder *TrainOrderRefundInfoBuilder) PersonalRefund(personalRefund int32) *TrainOrderRefundInfoBuilder {
	builder.personalRefund = personalRefund
	builder.personalRefundSet = true
	return builder
}

func (builder *TrainOrderRefundInfoBuilder) Build() *TrainOrderRefundInfo {
	data := &TrainOrderRefundInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.refundHandleFeeSet {
		data.RefundHandleFee = &builder.refundHandleFee
	}
	if builder.userRefundSet {
		data.UserRefund = &builder.userRefund
	}
	if builder.refundServiceFeeSet {
		data.RefundServiceFee = &builder.refundServiceFee
	}
	if builder.refundedTimeSet {
		data.RefundedTime = &builder.refundedTime
	}
	if builder.reasonForRefundSet {
		data.ReasonForRefund = &builder.reasonForRefund
	}
	if builder.companyRefundSet {
		data.CompanyRefund = &builder.companyRefund
	}
	if builder.personalRefundSet {
		data.PersonalRefund = &builder.personalRefund
	}
	return data
}
