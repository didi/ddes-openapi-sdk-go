package v1

// DetailsAmount // 员工支付金额明细 message PersonalPayDetail {   option (openapi.v3.schema) = {     required: [\"reimburse_fee\", \"no_reimburse_fee\"]   };   int32 reimburse_fee = 1 [(openapi.v3.property) = {description:\"可报销金额\", default:{number:0}}];   int32 no_reimburse_fee = 2 [(openapi.v3.property) = {description:\"不可报销金额\", default:{number:0}}]; }  // 订单员工实际支付金额明细 message PersonalRealPayDetail {   option (openapi.v3.schema) = {     required: [\"reimburse_fee\", \"no_reimburse_fee\"]   };   int32 reimburse_fee = 1 [(openapi.v3.property) = {description:\"可报销金额\", default:{number:0}}];   int32 no_reimburse_fee = 2 [(openapi.v3.property) = {description:\"不可报销金额\", default:{number:0}}]; }  // 员工退款金额明细 message PersonalRealRefundDetail {   option (openapi.v3.schema) = {     required: [\"reimburse_fee\", \"no_reimburse_fee\"]   };   int32 reimburse_fee = 1 [(openapi.v3.property) = {description:\"可报销金额\", default:{number:0}}];   int32 no_reimburse_fee = 2 [(openapi.v3.property) = {description:\"不可报销金额\", default:{number:0}}]; }  每日价格
type DetailsAmount struct {
	Date  *string `json:"date,omitempty"`  // 日期，格式：yyyy-MM-dd
	Price *int32  `json:"price,omitempty"` // 价格，单位：分
}

type DetailsAmountBuilder struct {
	date     string // 日期，格式：yyyy-MM-dd
	dateSet  bool
	price    int32 // 价格，单位：分
	priceSet bool
}

func NewDetailsAmountBuilder() *DetailsAmountBuilder {
	return &DetailsAmountBuilder{}
}
func (builder *DetailsAmountBuilder) Date(date string) *DetailsAmountBuilder {
	builder.date = date
	builder.dateSet = true
	return builder
}
func (builder *DetailsAmountBuilder) Price(price int32) *DetailsAmountBuilder {
	builder.price = price
	builder.priceSet = true
	return builder
}

func (builder *DetailsAmountBuilder) Build() *DetailsAmount {
	data := &DetailsAmount{}
	if builder.dateSet {
		data.Date = &builder.date
	}
	if builder.priceSet {
		data.Price = &builder.price
	}
	return data
}
