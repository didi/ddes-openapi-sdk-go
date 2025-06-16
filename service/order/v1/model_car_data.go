package v1

// CarData 用车订单信息消息
type CarData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 用车订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type CarDataBuilder struct {
	orderIds    string // 用车订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewCarDataBuilder() *CarDataBuilder {
	return &CarDataBuilder{}
}
func (builder *CarDataBuilder) OrderIds(orderIds string) *CarDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *CarDataBuilder) Page(page PageInfo) *CarDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *CarDataBuilder) Build() *CarData {
	data := &CarData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
