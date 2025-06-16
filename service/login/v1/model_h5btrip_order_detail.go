package v1

// H5BTripOrderDetail 订单详情页
type H5BTripOrderDetail struct {
	JumpPage *string `json:"jump_page,omitempty"` // 差旅、申请用车，必传
	OrderId  *string `json:"order_id,omitempty"`  // 订单id
}

type H5BTripOrderDetailBuilder struct {
	jumpPage    string // 差旅、申请用车，必传
	jumpPageSet bool
	orderId     string // 订单id
	orderIdSet  bool
}

func NewH5BTripOrderDetailBuilder() *H5BTripOrderDetailBuilder {
	return &H5BTripOrderDetailBuilder{}
}
func (builder *H5BTripOrderDetailBuilder) JumpPage(jumpPage string) *H5BTripOrderDetailBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5BTripOrderDetailBuilder) OrderId(orderId string) *H5BTripOrderDetailBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}

func (builder *H5BTripOrderDetailBuilder) Build() *H5BTripOrderDetail {
	data := &H5BTripOrderDetail{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	return data
}
