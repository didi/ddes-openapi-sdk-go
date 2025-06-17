package v1

// DomesticFlightOrderRefundInfo struct for DomesticFlightOrderRefundInfo
type DomesticFlightOrderRefundInfo struct {
	TicketUniqueKey              *string               `json:"ticket_unique_key,omitempty"`               // 票号唯一值，退票成功的信息
	RefundHandleFee              *int32                `json:"refund_handle_fee,omitempty"`               // 退票手续费，单位：分 票价（包含改签手续费），不包含平台使用费
	CutFee                       *int32                `json:"cut_fee,omitempty"`                         // 机票订立减优惠，单位：分 退票时回收金额
	RefundServiceFee             *int32                `json:"refund_service_fee,omitempty"`              // 退票平台使用费，单位：分
	RefundAsynchronousServiceFee *int32                `json:"refund_asynchronous_service_fee,omitempty"` // 退票异步平台使用费，单位：分
	ReasonForRefund              *string               `json:"reason_for_refund,omitempty"`               // 退票原因
	RefundUserFee                *int32                `json:"refund_user_fee,omitempty"`                 // 退还用户金额，单位：分 refund_user_fee=company_refund+personal_refund
	CompanyRefund                *int32                `json:"company_refund,omitempty"`                  // 企业支付退款金额，单位：分 目前包含服务费
	PersonalRefund               *int32                `json:"personal_refund,omitempty"`                 // 员工支付退款金额，单位：分
	PersonalRefundDetail         *PersonalRefundDetail `json:"personal_refund_detail,omitempty"`
	RefundedTime                 *string               `json:"refunded_time,omitempty"` // 退票成功时间，格式：yyyy-MM-dd HH:mm:ss
}

type DomesticFlightOrderRefundInfoBuilder struct {
	ticketUniqueKey                 string // 票号唯一值，退票成功的信息
	ticketUniqueKeySet              bool
	refundHandleFee                 int32 // 退票手续费，单位：分 票价（包含改签手续费），不包含平台使用费
	refundHandleFeeSet              bool
	cutFee                          int32 // 机票订立减优惠，单位：分 退票时回收金额
	cutFeeSet                       bool
	refundServiceFee                int32 // 退票平台使用费，单位：分
	refundServiceFeeSet             bool
	refundAsynchronousServiceFee    int32 // 退票异步平台使用费，单位：分
	refundAsynchronousServiceFeeSet bool
	reasonForRefund                 string // 退票原因
	reasonForRefundSet              bool
	refundUserFee                   int32 // 退还用户金额，单位：分 refund_user_fee=company_refund+personal_refund
	refundUserFeeSet                bool
	companyRefund                   int32 // 企业支付退款金额，单位：分 目前包含服务费
	companyRefundSet                bool
	personalRefund                  int32 // 员工支付退款金额，单位：分
	personalRefundSet               bool
	personalRefundDetail            PersonalRefundDetail
	personalRefundDetailSet         bool
	refundedTime                    string // 退票成功时间，格式：yyyy-MM-dd HH:mm:ss
	refundedTimeSet                 bool
}

func NewDomesticFlightOrderRefundInfoBuilder() *DomesticFlightOrderRefundInfoBuilder {
	return &DomesticFlightOrderRefundInfoBuilder{}
}
func (builder *DomesticFlightOrderRefundInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *DomesticFlightOrderRefundInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) RefundHandleFee(refundHandleFee int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.refundHandleFee = refundHandleFee
	builder.refundHandleFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) CutFee(cutFee int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) RefundServiceFee(refundServiceFee int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.refundServiceFee = refundServiceFee
	builder.refundServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) RefundAsynchronousServiceFee(refundAsynchronousServiceFee int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.refundAsynchronousServiceFee = refundAsynchronousServiceFee
	builder.refundAsynchronousServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) ReasonForRefund(reasonForRefund string) *DomesticFlightOrderRefundInfoBuilder {
	builder.reasonForRefund = reasonForRefund
	builder.reasonForRefundSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) RefundUserFee(refundUserFee int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.refundUserFee = refundUserFee
	builder.refundUserFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) CompanyRefund(companyRefund int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.companyRefund = companyRefund
	builder.companyRefundSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) PersonalRefund(personalRefund int32) *DomesticFlightOrderRefundInfoBuilder {
	builder.personalRefund = personalRefund
	builder.personalRefundSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) PersonalRefundDetail(personalRefundDetail PersonalRefundDetail) *DomesticFlightOrderRefundInfoBuilder {
	builder.personalRefundDetail = personalRefundDetail
	builder.personalRefundDetailSet = true
	return builder
}
func (builder *DomesticFlightOrderRefundInfoBuilder) RefundedTime(refundedTime string) *DomesticFlightOrderRefundInfoBuilder {
	builder.refundedTime = refundedTime
	builder.refundedTimeSet = true
	return builder
}

func (builder *DomesticFlightOrderRefundInfoBuilder) Build() *DomesticFlightOrderRefundInfo {
	data := &DomesticFlightOrderRefundInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.refundHandleFeeSet {
		data.RefundHandleFee = &builder.refundHandleFee
	}
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	if builder.refundServiceFeeSet {
		data.RefundServiceFee = &builder.refundServiceFee
	}
	if builder.refundAsynchronousServiceFeeSet {
		data.RefundAsynchronousServiceFee = &builder.refundAsynchronousServiceFee
	}
	if builder.reasonForRefundSet {
		data.ReasonForRefund = &builder.reasonForRefund
	}
	if builder.refundUserFeeSet {
		data.RefundUserFee = &builder.refundUserFee
	}
	if builder.companyRefundSet {
		data.CompanyRefund = &builder.companyRefund
	}
	if builder.personalRefundSet {
		data.PersonalRefund = &builder.personalRefund
	}
	if builder.personalRefundDetailSet {
		data.PersonalRefundDetail = &builder.personalRefundDetail
	}
	if builder.refundedTimeSet {
		data.RefundedTime = &builder.refundedTime
	}
	return data
}
