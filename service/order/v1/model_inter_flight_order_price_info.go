package v1

// InterFlightOrderPriceInfo 价格信息（国际）
type InterFlightOrderPriceInfo struct {
	OrderAmount              *int32                    `json:"order_amount,omitempty"`             // 订单总价，预订时的价格，包含机票立减优惠，不包含退改，金额不会更改，不包含异步平台服务服务费。 定义预订出票时订单总花费=order_amount+asynchronous_service_fee
	CompanyPay               *int32                    `json:"company_pay,omitempty"`              // 公司支付金额，预订时订单总价内公司支付的金额 定义预订出票时公司支付金额=company_pay
	PersonalPay              *int32                    `json:"personal_pay,omitempty"`             // 个人支付金额，预订时订单总价内个人支付的金额
	ServiceFee               *int32                    `json:"service_fee,omitempty"`              // 平台使用费，不包含异步平台使用费
	AsynchronousServiceFee   *int32                    `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	CompanyRealPay           *int32                    `json:"company_real_pay,omitempty"`         // 订单公司支付金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司支付金额会变更。按照钱支出方向累计
	PersonalRealPay          *int32                    `json:"personal_real_pay,omitempty"`        // 订单个人支付金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人支付金额会变更。按照钱支出方向累计
	CompanyRealRefund        *int32                    `json:"company_real_refund,omitempty"`      // 订单公司收入金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司收入金额会变更。按照钱收入方向累计
	PersonalRealRefund       *int32                    `json:"personal_real_refund,omitempty"`     // 订单个人收入金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人收入付金额会变更。按照钱收入方向累计
	Currency                 *string                   `json:"currency,omitempty"`                 // 币种，默认 CNY
	PersonalPayDetail        *PersonalPayDetail        `json:"personal_pay_detail,omitempty"`      // 员工支付金额明细
	PersonalRealPayDetail    *PersonalRealPayDetail    `json:"personal_real_pay_detail,omitempty"`
	PersonalRealRefundDetail *PersonalRealRefundDetail `json:"personal_real_refund_detail,omitempty"`
	SrvPackCompanyRealPay    *int32                    `json:"srv_pack_company_real_pay,omitempty"` // 服务包企业实付金额 单位：分，退服务包时，金额会变
	CutFee                   *int32                    `json:"cut_fee,omitempty"`                   // 机票订立减优惠
}

type InterFlightOrderPriceInfoBuilder struct {
	orderAmount                 int32 // 订单总价，预订时的价格，包含机票立减优惠，不包含退改，金额不会更改，不包含异步平台服务服务费。 定义预订出票时订单总花费=order_amount+asynchronous_service_fee
	orderAmountSet              bool
	companyPay                  int32 // 公司支付金额，预订时订单总价内公司支付的金额 定义预订出票时公司支付金额=company_pay
	companyPaySet               bool
	personalPay                 int32 // 个人支付金额，预订时订单总价内个人支付的金额
	personalPaySet              bool
	serviceFee                  int32 // 平台使用费，不包含异步平台使用费
	serviceFeeSet               bool
	asynchronousServiceFee      int32 // 出票平台使用费（异步），单位：分
	asynchronousServiceFeeSet   bool
	companyRealPay              int32 // 订单公司支付金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司支付金额会变更。按照钱支出方向累计
	companyRealPaySet           bool
	personalRealPay             int32 // 订单个人支付金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人支付金额会变更。按照钱支出方向累计
	personalRealPaySet          bool
	companyRealRefund           int32 // 订单公司收入金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的公司收入金额会变更。按照钱收入方向累计
	companyRealRefundSet        bool
	personalRealRefund          int32 // 订单个人收入金额，单位：分 逻辑金额。跟随收费配置和退票改签，订单的个人收入付金额会变更。按照钱收入方向累计
	personalRealRefundSet       bool
	currency                    string // 币种，默认 CNY
	currencySet                 bool
	personalPayDetail           PersonalPayDetail // 员工支付金额明细
	personalPayDetailSet        bool
	personalRealPayDetail       PersonalRealPayDetail
	personalRealPayDetailSet    bool
	personalRealRefundDetail    PersonalRealRefundDetail
	personalRealRefundDetailSet bool
	srvPackCompanyRealPay       int32 // 服务包企业实付金额 单位：分，退服务包时，金额会变
	srvPackCompanyRealPaySet    bool
	cutFee                      int32 // 机票订立减优惠
	cutFeeSet                   bool
}

func NewInterFlightOrderPriceInfoBuilder() *InterFlightOrderPriceInfoBuilder {
	return &InterFlightOrderPriceInfoBuilder{}
}
func (builder *InterFlightOrderPriceInfoBuilder) OrderAmount(orderAmount int32) *InterFlightOrderPriceInfoBuilder {
	builder.orderAmount = orderAmount
	builder.orderAmountSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) CompanyPay(companyPay int32) *InterFlightOrderPriceInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalPay(personalPay int32) *InterFlightOrderPriceInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) ServiceFee(serviceFee int32) *InterFlightOrderPriceInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *InterFlightOrderPriceInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) CompanyRealPay(companyRealPay int32) *InterFlightOrderPriceInfoBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalRealPay(personalRealPay int32) *InterFlightOrderPriceInfoBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) CompanyRealRefund(companyRealRefund int32) *InterFlightOrderPriceInfoBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalRealRefund(personalRealRefund int32) *InterFlightOrderPriceInfoBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) Currency(currency string) *InterFlightOrderPriceInfoBuilder {
	builder.currency = currency
	builder.currencySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalPayDetail(personalPayDetail PersonalPayDetail) *InterFlightOrderPriceInfoBuilder {
	builder.personalPayDetail = personalPayDetail
	builder.personalPayDetailSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalRealPayDetail(personalRealPayDetail PersonalRealPayDetail) *InterFlightOrderPriceInfoBuilder {
	builder.personalRealPayDetail = personalRealPayDetail
	builder.personalRealPayDetailSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) PersonalRealRefundDetail(personalRealRefundDetail PersonalRealRefundDetail) *InterFlightOrderPriceInfoBuilder {
	builder.personalRealRefundDetail = personalRealRefundDetail
	builder.personalRealRefundDetailSet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) SrvPackCompanyRealPay(srvPackCompanyRealPay int32) *InterFlightOrderPriceInfoBuilder {
	builder.srvPackCompanyRealPay = srvPackCompanyRealPay
	builder.srvPackCompanyRealPaySet = true
	return builder
}
func (builder *InterFlightOrderPriceInfoBuilder) CutFee(cutFee int32) *InterFlightOrderPriceInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}

func (builder *InterFlightOrderPriceInfoBuilder) Build() *InterFlightOrderPriceInfo {
	data := &InterFlightOrderPriceInfo{}
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
	if builder.personalPayDetailSet {
		data.PersonalPayDetail = &builder.personalPayDetail
	}
	if builder.personalRealPayDetailSet {
		data.PersonalRealPayDetail = &builder.personalRealPayDetail
	}
	if builder.personalRealRefundDetailSet {
		data.PersonalRealRefundDetail = &builder.personalRealRefundDetail
	}
	if builder.srvPackCompanyRealPaySet {
		data.SrvPackCompanyRealPay = &builder.srvPackCompanyRealPay
	}
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	return data
}
