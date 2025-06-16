package v1

// SubBillSummaryDTO 账单汇总查询-商旅、网约车、出租车 ～ 商旅 ～ 子业务线汇总金额
type SubBillSummaryDTO struct {
	BusinessType         *int32  `json:"business_type,omitempty"`          // 业务类型 201:国内酒店 202:国内机票 203:火车票 204:海外酒店 205:国际机票 531:增值服务
	ConsumeAmount        *string `json:"consume_amount,omitempty"`         // 消耗金额（单位：元）
	RefundAmount         *string `json:"refund_amount,omitempty"`          // 退款金额（单位：元）
	PreviousRefundAmount *string `json:"previous_refund_amount,omitempty"` // 往期退款金额（单位：元）
	AmountMoney          *string `json:"amount_money,omitempty"`           // 结算金额（单位：元）
}

type SubBillSummaryDTOBuilder struct {
	businessType            int32 // 业务类型 201:国内酒店 202:国内机票 203:火车票 204:海外酒店 205:国际机票 531:增值服务
	businessTypeSet         bool
	consumeAmount           string // 消耗金额（单位：元）
	consumeAmountSet        bool
	refundAmount            string // 退款金额（单位：元）
	refundAmountSet         bool
	previousRefundAmount    string // 往期退款金额（单位：元）
	previousRefundAmountSet bool
	amountMoney             string // 结算金额（单位：元）
	amountMoneySet          bool
}

func NewSubBillSummaryDTOBuilder() *SubBillSummaryDTOBuilder {
	return &SubBillSummaryDTOBuilder{}
}
func (builder *SubBillSummaryDTOBuilder) BusinessType(businessType int32) *SubBillSummaryDTOBuilder {
	builder.businessType = businessType
	builder.businessTypeSet = true
	return builder
}
func (builder *SubBillSummaryDTOBuilder) ConsumeAmount(consumeAmount string) *SubBillSummaryDTOBuilder {
	builder.consumeAmount = consumeAmount
	builder.consumeAmountSet = true
	return builder
}
func (builder *SubBillSummaryDTOBuilder) RefundAmount(refundAmount string) *SubBillSummaryDTOBuilder {
	builder.refundAmount = refundAmount
	builder.refundAmountSet = true
	return builder
}
func (builder *SubBillSummaryDTOBuilder) PreviousRefundAmount(previousRefundAmount string) *SubBillSummaryDTOBuilder {
	builder.previousRefundAmount = previousRefundAmount
	builder.previousRefundAmountSet = true
	return builder
}
func (builder *SubBillSummaryDTOBuilder) AmountMoney(amountMoney string) *SubBillSummaryDTOBuilder {
	builder.amountMoney = amountMoney
	builder.amountMoneySet = true
	return builder
}

func (builder *SubBillSummaryDTOBuilder) Build() *SubBillSummaryDTO {
	data := &SubBillSummaryDTO{}
	if builder.businessTypeSet {
		data.BusinessType = &builder.businessType
	}
	if builder.consumeAmountSet {
		data.ConsumeAmount = &builder.consumeAmount
	}
	if builder.refundAmountSet {
		data.RefundAmount = &builder.refundAmount
	}
	if builder.previousRefundAmountSet {
		data.PreviousRefundAmount = &builder.previousRefundAmount
	}
	if builder.amountMoneySet {
		data.AmountMoney = &builder.amountMoney
	}
	return data
}
