package v1

// SubListItem 账单汇总查询-商旅、网约车、出租车 ～ 网约车 ～ 子账户汇总【4.0账单返回】
type SubListItem struct {
	PayChannel           *string  `json:"pay_channel,omitempty"`            // 结算方式
	ConsumeAmountOnline  *float32 `json:"consume_amount_online,omitempty"`  // 线上企业实付金额
	ConsumeAmountOffline *float32 `json:"consume_amount_offline,omitempty"` // 线下扣企业账户余额
	RefundAmount         *float32 `json:"refund_amount,omitempty"`          // 企业实退金额
	AdjustAmount         *float32 `json:"adjust_amount,omitempty"`          // 调整金额
	BillAmount           *float32 `json:"bill_amount,omitempty"`            // 账单金额
	CurrentBalance       *float32 `json:"current_balance,omitempty"`        // 期末余额
}

type SubListItemBuilder struct {
	payChannel              string // 结算方式
	payChannelSet           bool
	consumeAmountOnline     float32 // 线上企业实付金额
	consumeAmountOnlineSet  bool
	consumeAmountOffline    float32 // 线下扣企业账户余额
	consumeAmountOfflineSet bool
	refundAmount            float32 // 企业实退金额
	refundAmountSet         bool
	adjustAmount            float32 // 调整金额
	adjustAmountSet         bool
	billAmount              float32 // 账单金额
	billAmountSet           bool
	currentBalance          float32 // 期末余额
	currentBalanceSet       bool
}

func NewSubListItemBuilder() *SubListItemBuilder {
	return &SubListItemBuilder{}
}
func (builder *SubListItemBuilder) PayChannel(payChannel string) *SubListItemBuilder {
	builder.payChannel = payChannel
	builder.payChannelSet = true
	return builder
}
func (builder *SubListItemBuilder) ConsumeAmountOnline(consumeAmountOnline float32) *SubListItemBuilder {
	builder.consumeAmountOnline = consumeAmountOnline
	builder.consumeAmountOnlineSet = true
	return builder
}
func (builder *SubListItemBuilder) ConsumeAmountOffline(consumeAmountOffline float32) *SubListItemBuilder {
	builder.consumeAmountOffline = consumeAmountOffline
	builder.consumeAmountOfflineSet = true
	return builder
}
func (builder *SubListItemBuilder) RefundAmount(refundAmount float32) *SubListItemBuilder {
	builder.refundAmount = refundAmount
	builder.refundAmountSet = true
	return builder
}
func (builder *SubListItemBuilder) AdjustAmount(adjustAmount float32) *SubListItemBuilder {
	builder.adjustAmount = adjustAmount
	builder.adjustAmountSet = true
	return builder
}
func (builder *SubListItemBuilder) BillAmount(billAmount float32) *SubListItemBuilder {
	builder.billAmount = billAmount
	builder.billAmountSet = true
	return builder
}
func (builder *SubListItemBuilder) CurrentBalance(currentBalance float32) *SubListItemBuilder {
	builder.currentBalance = currentBalance
	builder.currentBalanceSet = true
	return builder
}

func (builder *SubListItemBuilder) Build() *SubListItem {
	data := &SubListItem{}
	if builder.payChannelSet {
		data.PayChannel = &builder.payChannel
	}
	if builder.consumeAmountOnlineSet {
		data.ConsumeAmountOnline = &builder.consumeAmountOnline
	}
	if builder.consumeAmountOfflineSet {
		data.ConsumeAmountOffline = &builder.consumeAmountOffline
	}
	if builder.refundAmountSet {
		data.RefundAmount = &builder.refundAmount
	}
	if builder.adjustAmountSet {
		data.AdjustAmount = &builder.adjustAmount
	}
	if builder.billAmountSet {
		data.BillAmount = &builder.billAmount
	}
	if builder.currentBalanceSet {
		data.CurrentBalance = &builder.currentBalance
	}
	return data
}
