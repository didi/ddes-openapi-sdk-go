package v1

// DomesticFlightOrderPriceInfo 价格信息(国内)
type DomesticFlightOrderPriceInfo struct {
	OrderAmount              *int32                    `json:"order_amount,omitempty"` // 订单总价 预订时的价格，包含机票立减优惠，不包含退改，金额不会更改，不包含异步平台服务服务费。 定义预订出票时订单总花费=order_amount+asynchronous_service_fee
	CompanyPay               *int32                    `json:"company_pay,omitempty"`  // 公司支付金额 预订时订单总价内公司支付的金额 定义预订出票时公司支付金额=company_pay
	PersonalPay              *int32                    `json:"personal_pay,omitempty"` // 员工支付金额 预订时订单总价内员工支付的金额
	PersonalPayDetail        *PersonalPayDetail        `json:"personal_pay_detail,omitempty"`
	ServiceFee               *int32                    `json:"service_fee,omitempty"`              // 平台使用费 不包含异步平台使用费
	AsynchronousServiceFee   *int32                    `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步） 单位：分
	CompanyRealPay           *int32                    `json:"company_real_pay,omitempty"`         // 订单公司支付金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司支付金额会变更。按照钱支出方向累计
	PersonalRealPay          *int32                    `json:"personal_real_pay,omitempty"`        // 订单个人支付金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人支付金额会变更。按照钱支出方向累计
	PersonalRealPayDetail    *PersonalRealPayDetail    `json:"personal_real_pay_detail,omitempty"`
	CutFee                   *int32                    `json:"cut_fee,omitempty"`              // 机票订立减优惠
	CompanyRealRefund        *int32                    `json:"company_real_refund,omitempty"`  // 订单公司收入金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司收入金额会变更。按照钱收入方向累计
	PersonalRealRefund       *int32                    `json:"personal_real_refund,omitempty"` // 订单个人收入金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人收入付金额会变更。按照钱收入方向累计
	PersonalRealRefundDetail *PersonalRealRefundDetail `json:"personal_real_refund_detail,omitempty"`
	SrvPackCompanyRealPay    *int32                    `json:"srv_pack_company_real_pay,omitempty"` // 服务包企业实付金额 单位：分，退服务包时，金额会变
	Currency                 *string                   `json:"currency,omitempty"`                  // 币种 默认 CNY
}

type DomesticFlightOrderPriceInfoBuilder struct {
	orderAmount                 int32 // 订单总价 预订时的价格，包含机票立减优惠，不包含退改，金额不会更改，不包含异步平台服务服务费。 定义预订出票时订单总花费=order_amount+asynchronous_service_fee
	orderAmountSet              bool
	companyPay                  int32 // 公司支付金额 预订时订单总价内公司支付的金额 定义预订出票时公司支付金额=company_pay
	companyPaySet               bool
	personalPay                 int32 // 员工支付金额 预订时订单总价内员工支付的金额
	personalPaySet              bool
	personalPayDetail           PersonalPayDetail
	personalPayDetailSet        bool
	serviceFee                  int32 // 平台使用费 不包含异步平台使用费
	serviceFeeSet               bool
	asynchronousServiceFee      int32 // 出票平台使用费（异步） 单位：分
	asynchronousServiceFeeSet   bool
	companyRealPay              int32 // 订单公司支付金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司支付金额会变更。按照钱支出方向累计
	companyRealPaySet           bool
	personalRealPay             int32 // 订单个人支付金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人支付金额会变更。按照钱支出方向累计
	personalRealPaySet          bool
	personalRealPayDetail       PersonalRealPayDetail
	personalRealPayDetailSet    bool
	cutFee                      int32 // 机票订立减优惠
	cutFeeSet                   bool
	companyRealRefund           int32 // 订单公司收入金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司收入金额会变更。按照钱收入方向累计
	companyRealRefundSet        bool
	personalRealRefund          int32 // 订单个人收入金额 单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人收入付金额会变更。按照钱收入方向累计
	personalRealRefundSet       bool
	personalRealRefundDetail    PersonalRealRefundDetail
	personalRealRefundDetailSet bool
	srvPackCompanyRealPay       int32 // 服务包企业实付金额 单位：分，退服务包时，金额会变
	srvPackCompanyRealPaySet    bool
	currency                    string // 币种 默认 CNY
	currencySet                 bool
}

func NewDomesticFlightOrderPriceInfoBuilder() *DomesticFlightOrderPriceInfoBuilder {
	return &DomesticFlightOrderPriceInfoBuilder{}
}
func (builder *DomesticFlightOrderPriceInfoBuilder) OrderAmount(orderAmount int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.orderAmount = orderAmount
	builder.orderAmountSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) CompanyPay(companyPay int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalPay(personalPay int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalPayDetail(personalPayDetail PersonalPayDetail) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalPayDetail = personalPayDetail
	builder.personalPayDetailSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) ServiceFee(serviceFee int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) CompanyRealPay(companyRealPay int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalRealPay(personalRealPay int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalRealPayDetail(personalRealPayDetail PersonalRealPayDetail) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalRealPayDetail = personalRealPayDetail
	builder.personalRealPayDetailSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) CutFee(cutFee int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) CompanyRealRefund(companyRealRefund int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalRealRefund(personalRealRefund int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) PersonalRealRefundDetail(personalRealRefundDetail PersonalRealRefundDetail) *DomesticFlightOrderPriceInfoBuilder {
	builder.personalRealRefundDetail = personalRealRefundDetail
	builder.personalRealRefundDetailSet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) SrvPackCompanyRealPay(srvPackCompanyRealPay int32) *DomesticFlightOrderPriceInfoBuilder {
	builder.srvPackCompanyRealPay = srvPackCompanyRealPay
	builder.srvPackCompanyRealPaySet = true
	return builder
}
func (builder *DomesticFlightOrderPriceInfoBuilder) Currency(currency string) *DomesticFlightOrderPriceInfoBuilder {
	builder.currency = currency
	builder.currencySet = true
	return builder
}

func (builder *DomesticFlightOrderPriceInfoBuilder) Build() *DomesticFlightOrderPriceInfo {
	data := &DomesticFlightOrderPriceInfo{}
	if builder.orderAmountSet {
		data.OrderAmount = &builder.orderAmount
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.personalPayDetailSet {
		data.PersonalPayDetail = &builder.personalPayDetail
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.asynchronousServiceFeeSet {
		data.AsynchronousServiceFee = &builder.asynchronousServiceFee
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.personalRealPayDetailSet {
		data.PersonalRealPayDetail = &builder.personalRealPayDetail
	}
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	if builder.personalRealRefundDetailSet {
		data.PersonalRealRefundDetail = &builder.personalRealRefundDetail
	}
	if builder.srvPackCompanyRealPaySet {
		data.SrvPackCompanyRealPay = &builder.srvPackCompanyRealPay
	}
	if builder.currencySet {
		data.Currency = &builder.currency
	}
	return data
}
