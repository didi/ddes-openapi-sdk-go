package v1

// InterHotelOrderPriceInfo 酒店订单价格信息（国际）
type InterHotelOrderPriceInfo struct {
	OrderAmount              *int32                    `json:"order_amount,omitempty"` // 订单总价，酒店预订时的费用，不包含离店后的结算信息  =公司支付金额+个人支付金额=房费总额+平台使用费（实时） ，不包含立减优惠，税点服务费
	CompanyPay               *int32                    `json:"company_pay,omitempty"`  // 公司应付支付金额，单位：分
	PersonalPay              *int32                    `json:"personal_pay,omitempty"` // 员工应付支付金额，单位：分
	PersonalPayDetail        *PersonalPayDetail        `json:"personal_pay_detail,omitempty"`
	CutFee                   *int32                    `json:"cut_fee,omitempty"`                  // 立减金，单位：分
	CompanyCoupon            *int32                    `json:"company_coupon,omitempty"`           // 公司优惠券，单位：分 退款时，该字段不更新
	PersonalCoupon           *int32                    `json:"personal_coupon,omitempty"`          // 个人优惠券，单位：分 退款时，该字段不更新
	ServiceFee               *int32                    `json:"service_fee,omitempty"`              // 平台使用费，单位：分 实时随单收取平台使用费
	AsynchronousServiceFee   *int32                    `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	DifftaxFee               *int32                    `json:"difftax_fee,omitempty"`              // 税点服务费，单位：分 发生退款时，会变更为最终值金额
	RoomAmount               *int32                    `json:"room_amount,omitempty"`              // 房费总额，单位：分
	CompanyRealPay           *int32                    `json:"company_real_pay,omitempty"`         // 订单公司实际支付金额，单位：分
	PersonalRealPay          *int32                    `json:"personal_real_pay,omitempty"`        // 订单个人实际支付金额，单位：分
	PersonalRealPayDetail    *PersonalRealPayDetail    `json:"personal_real_pay_detail,omitempty"`
	CompanyRealRefund        *int32                    `json:"company_real_refund,omitempty"`  // 公司实际退款金额，单位：分 发生退款款时有值
	PersonalRealRefund       *int32                    `json:"personal_real_refund,omitempty"` // 员工实际退款金额，单位：分 发生退款款时有值
	PersonalRealRefundDetail *PersonalRealRefundDetail `json:"personal_real_refund_detail,omitempty"`
	Currency                 *string                   `json:"currency,omitempty"` // 币种，默认 CNY
}

type InterHotelOrderPriceInfoBuilder struct {
	orderAmount                 int32 // 订单总价，酒店预订时的费用，不包含离店后的结算信息  =公司支付金额+个人支付金额=房费总额+平台使用费（实时） ，不包含立减优惠，税点服务费
	orderAmountSet              bool
	companyPay                  int32 // 公司应付支付金额，单位：分
	companyPaySet               bool
	personalPay                 int32 // 员工应付支付金额，单位：分
	personalPaySet              bool
	personalPayDetail           PersonalPayDetail
	personalPayDetailSet        bool
	cutFee                      int32 // 立减金，单位：分
	cutFeeSet                   bool
	companyCoupon               int32 // 公司优惠券，单位：分 退款时，该字段不更新
	companyCouponSet            bool
	personalCoupon              int32 // 个人优惠券，单位：分 退款时，该字段不更新
	personalCouponSet           bool
	serviceFee                  int32 // 平台使用费，单位：分 实时随单收取平台使用费
	serviceFeeSet               bool
	asynchronousServiceFee      int32 // 出票平台使用费（异步），单位：分
	asynchronousServiceFeeSet   bool
	difftaxFee                  int32 // 税点服务费，单位：分 发生退款时，会变更为最终值金额
	difftaxFeeSet               bool
	roomAmount                  int32 // 房费总额，单位：分
	roomAmountSet               bool
	companyRealPay              int32 // 订单公司实际支付金额，单位：分
	companyRealPaySet           bool
	personalRealPay             int32 // 订单个人实际支付金额，单位：分
	personalRealPaySet          bool
	personalRealPayDetail       PersonalRealPayDetail
	personalRealPayDetailSet    bool
	companyRealRefund           int32 // 公司实际退款金额，单位：分 发生退款款时有值
	companyRealRefundSet        bool
	personalRealRefund          int32 // 员工实际退款金额，单位：分 发生退款款时有值
	personalRealRefundSet       bool
	personalRealRefundDetail    PersonalRealRefundDetail
	personalRealRefundDetailSet bool
	currency                    string // 币种，默认 CNY
	currencySet                 bool
}

func NewInterHotelOrderPriceInfoBuilder() *InterHotelOrderPriceInfoBuilder {
	return &InterHotelOrderPriceInfoBuilder{}
}
func (builder *InterHotelOrderPriceInfoBuilder) OrderAmount(orderAmount int32) *InterHotelOrderPriceInfoBuilder {
	builder.orderAmount = orderAmount
	builder.orderAmountSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) CompanyPay(companyPay int32) *InterHotelOrderPriceInfoBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalPay(personalPay int32) *InterHotelOrderPriceInfoBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalPayDetail(personalPayDetail PersonalPayDetail) *InterHotelOrderPriceInfoBuilder {
	builder.personalPayDetail = personalPayDetail
	builder.personalPayDetailSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) CutFee(cutFee int32) *InterHotelOrderPriceInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) CompanyCoupon(companyCoupon int32) *InterHotelOrderPriceInfoBuilder {
	builder.companyCoupon = companyCoupon
	builder.companyCouponSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalCoupon(personalCoupon int32) *InterHotelOrderPriceInfoBuilder {
	builder.personalCoupon = personalCoupon
	builder.personalCouponSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) ServiceFee(serviceFee int32) *InterHotelOrderPriceInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *InterHotelOrderPriceInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) DifftaxFee(difftaxFee int32) *InterHotelOrderPriceInfoBuilder {
	builder.difftaxFee = difftaxFee
	builder.difftaxFeeSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) RoomAmount(roomAmount int32) *InterHotelOrderPriceInfoBuilder {
	builder.roomAmount = roomAmount
	builder.roomAmountSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) CompanyRealPay(companyRealPay int32) *InterHotelOrderPriceInfoBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalRealPay(personalRealPay int32) *InterHotelOrderPriceInfoBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalRealPayDetail(personalRealPayDetail PersonalRealPayDetail) *InterHotelOrderPriceInfoBuilder {
	builder.personalRealPayDetail = personalRealPayDetail
	builder.personalRealPayDetailSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) CompanyRealRefund(companyRealRefund int32) *InterHotelOrderPriceInfoBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalRealRefund(personalRealRefund int32) *InterHotelOrderPriceInfoBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) PersonalRealRefundDetail(personalRealRefundDetail PersonalRealRefundDetail) *InterHotelOrderPriceInfoBuilder {
	builder.personalRealRefundDetail = personalRealRefundDetail
	builder.personalRealRefundDetailSet = true
	return builder
}
func (builder *InterHotelOrderPriceInfoBuilder) Currency(currency string) *InterHotelOrderPriceInfoBuilder {
	builder.currency = currency
	builder.currencySet = true
	return builder
}

func (builder *InterHotelOrderPriceInfoBuilder) Build() *InterHotelOrderPriceInfo {
	data := &InterHotelOrderPriceInfo{}
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
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	if builder.companyCouponSet {
		data.CompanyCoupon = &builder.companyCoupon
	}
	if builder.personalCouponSet {
		data.PersonalCoupon = &builder.personalCoupon
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.asynchronousServiceFeeSet {
		data.AsynchronousServiceFee = &builder.asynchronousServiceFee
	}
	if builder.difftaxFeeSet {
		data.DifftaxFee = &builder.difftaxFee
	}
	if builder.roomAmountSet {
		data.RoomAmount = &builder.roomAmount
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
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	if builder.personalRealRefundDetailSet {
		data.PersonalRealRefundDetail = &builder.personalRealRefundDetail
	}
	if builder.currencySet {
		data.Currency = &builder.currency
	}
	return data
}
