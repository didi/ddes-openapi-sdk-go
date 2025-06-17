package v1

// TrainOrderPriceInfo 价格信息
type TrainOrderPriceInfo struct {
	OrderAmount            *int32  `json:"order_amount,omitempty"`             // 订单总价，单位：分  金额字段下同 ，预订时的价格，不包含异步服务费，包含实时服务费，会变更到实际预订支付价格。 出票完成时订单总花费=order_amount+asynchronous_service_fee
	CompanyPay             *int32  `json:"company_pay,omitempty"`              // 公司支付金额，单位：分预订时的价格，不包含异步服务费，会变更到实际预订支付价格。 出票完成时公司支付金额=company_pay
	PersonalPay            *int32  `json:"personal_pay,omitempty"`             // 个人支付金额，单位：分预订时的价格，不包含异步服务费，会变更到实际预订支付价格。
	ServiceFee             *int32  `json:"service_fee,omitempty"`              // 平台使用费，单位：实时收取的服务费
	AsynchronousServiceFee *int32  `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	GrabServiceFee         *int32  `json:"grab_service_fee,omitempty"`         // 抢票服务费，单位：抢票成功后收取的抢票服务费
	CompanyRealPay         *int32  `json:"company_real_pay,omitempty"`         // 订单公司支付金额，单位：分 ，订单的公司支付金额会变更。
	PersonalRealPay        *int32  `json:"personal_real_pay,omitempty"`        // 订单个人支付金额，单位：分 ，订单的个人支付金额会变更。
	CompanyRealRefund      *int32  `json:"company_real_refund,omitempty"`      // 企业实退
	PersonalRealRefund     *int32  `json:"personal_real_refund,omitempty"`     // 个人实退
	Currency               *string `json:"currency,omitempty"`                 // 币种，默认 CNY
}

type TrainOrderPriceInfoBuilder struct {
	orderAmount               int32 // 订单总价，单位：分  金额字段下同 ，预订时的价格，不包含异步服务费，包含实时服务费，会变更到实际预订支付价格。 出票完成时订单总花费=order_amount+asynchronous_service_fee
	orderAmountSet            bool
	companyPay                int32 // 公司支付金额，单位：分预订时的价格，不包含异步服务费，会变更到实际预订支付价格。 出票完成时公司支付金额=company_pay
	companyPaySet             bool
	personalPay               int32 // 个人支付金额，单位：分预订时的价格，不包含异步服务费，会变更到实际预订支付价格。
	personalPaySet            bool
	serviceFee                int32 // 平台使用费，单位：实时收取的服务费
	serviceFeeSet             bool
	asynchronousServiceFee    int32 // 出票平台使用费（异步），单位：分
	asynchronousServiceFeeSet bool
	grabServiceFee            int32 // 抢票服务费，单位：抢票成功后收取的抢票服务费
	grabServiceFeeSet         bool
	companyRealPay            int32 // 订单公司支付金额，单位：分 ，订单的公司支付金额会变更。
	companyRealPaySet         bool
	personalRealPay           int32 // 订单个人支付金额，单位：分 ，订单的个人支付金额会变更。
	personalRealPaySet        bool
	companyRealRefund         int32 // 企业实退
	companyRealRefundSet      bool
	personalRealRefund        int32 // 个人实退
	personalRealRefundSet     bool
	currency                  string // 币种，默认 CNY
	currencySet               bool
}

func NewTrainOrderPriceInfoBuilder() *TrainOrderPriceInfoBuilder {
	return &TrainOrderPriceInfoBuilder{}
}
func (builder *TrainOrderPriceInfoBuilder) OrderAmount(orderAmount int32) *TrainOrderPriceInfoBuilder {
	builder.orderAmount = orderAmount
	builder.orderAmountSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) CompanyPay(companyPay int32) *TrainOrderPriceInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) PersonalPay(personalPay int32) *TrainOrderPriceInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) ServiceFee(serviceFee int32) *TrainOrderPriceInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *TrainOrderPriceInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) GrabServiceFee(grabServiceFee int32) *TrainOrderPriceInfoBuilder {
	builder.grabServiceFee = grabServiceFee
	builder.grabServiceFeeSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) CompanyRealPay(companyRealPay int32) *TrainOrderPriceInfoBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) PersonalRealPay(personalRealPay int32) *TrainOrderPriceInfoBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) CompanyRealRefund(companyRealRefund int32) *TrainOrderPriceInfoBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) PersonalRealRefund(personalRealRefund int32) *TrainOrderPriceInfoBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *TrainOrderPriceInfoBuilder) Currency(currency string) *TrainOrderPriceInfoBuilder {
	builder.currency = currency
	builder.currencySet = true
	return builder
}

func (builder *TrainOrderPriceInfoBuilder) Build() *TrainOrderPriceInfo {
	data := &TrainOrderPriceInfo{}
	if builder.orderAmountSet {
		data.OrderAmount = &builder.orderAmount
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.asynchronousServiceFeeSet {
		data.AsynchronousServiceFee = &builder.asynchronousServiceFee
	}
	if builder.grabServiceFeeSet {
		data.GrabServiceFee = &builder.grabServiceFee
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	if builder.currencySet {
		data.Currency = &builder.currency
	}
	return data
}
