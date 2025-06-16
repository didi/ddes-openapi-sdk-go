package v1

// SubBillSummaryItem 账单树查询(网约车、商旅) ~ 账单集合 ~ 账单集合
type SubBillSummaryItem struct {
	AmountMoney          *string `json:"amount_money,omitempty"`           // 结算金额（单位：元）
	BusinessType         *int32  `json:"business_type,omitempty"`          // 业务类型 201:国内酒店 202:国内机票 203:火车票 204:海外酒店 205:国际机票 531:增值服务
	ConsumeAmount        *string `json:"consume_amount,omitempty"`         // 消耗金额（单位：元）
	PreviousRefundAmount *string `json:"previous_refund_amount,omitempty"` // 往期退款金额（单位：元）
	RefundAmount         *string `json:"refund_amount,omitempty"`          // 退款金额（单位：元）
}

type SubBillSummaryItemBuilder struct {
	amountMoney             string // 结算金额（单位：元）
	amountMoneySet          bool
	businessType            int32 // 业务类型 201:国内酒店 202:国内机票 203:火车票 204:海外酒店 205:国际机票 531:增值服务
	businessTypeSet         bool
	consumeAmount           string // 消耗金额（单位：元）
	consumeAmountSet        bool
	previousRefundAmount    string // 往期退款金额（单位：元）
	previousRefundAmountSet bool
	refundAmount            string // 退款金额（单位：元）
	refundAmountSet         bool
}

func NewSubBillSummaryItemBuilder() *SubBillSummaryItemBuilder {
	return &SubBillSummaryItemBuilder{}
}
func (builder *SubBillSummaryItemBuilder) AmountMoney(amountMoney string) *SubBillSummaryItemBuilder {
	builder.amountMoney = amountMoney
	builder.amountMoneySet = true
	return builder
}
func (builder *SubBillSummaryItemBuilder) BusinessType(businessType int32) *SubBillSummaryItemBuilder {
	builder.businessType = businessType
	builder.businessTypeSet = true
	return builder
}
func (builder *SubBillSummaryItemBuilder) ConsumeAmount(consumeAmount string) *SubBillSummaryItemBuilder {
	builder.consumeAmount = consumeAmount
	builder.consumeAmountSet = true
	return builder
}
func (builder *SubBillSummaryItemBuilder) PreviousRefundAmount(previousRefundAmount string) *SubBillSummaryItemBuilder {
	builder.previousRefundAmount = previousRefundAmount
	builder.previousRefundAmountSet = true
	return builder
}
func (builder *SubBillSummaryItemBuilder) RefundAmount(refundAmount string) *SubBillSummaryItemBuilder {
	builder.refundAmount = refundAmount
	builder.refundAmountSet = true
	return builder
}

func (builder *SubBillSummaryItemBuilder) Build() *SubBillSummaryItem {
	data := &SubBillSummaryItem{}
	if builder.amountMoneySet {
		data.AmountMoney = &builder.amountMoney
	}
	if builder.businessTypeSet {
		data.BusinessType = &builder.businessType
	}
	if builder.consumeAmountSet {
		data.ConsumeAmount = &builder.consumeAmount
	}
	if builder.previousRefundAmountSet {
		data.PreviousRefundAmount = &builder.previousRefundAmount
	}
	if builder.refundAmountSet {
		data.RefundAmount = &builder.refundAmount
	}
	return data
}
