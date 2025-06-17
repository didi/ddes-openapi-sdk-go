package v1

// H5OrderDetail 订单详情
type H5OrderDetail struct {
	OrderId *string `json:"order_id,omitempty"` // 订单id
}

type H5OrderDetailBuilder struct {
	orderId    string // 订单id
	orderIdSet bool
}

func NewH5OrderDetailBuilder() *H5OrderDetailBuilder {
	return &H5OrderDetailBuilder{}
}
func (builder *H5OrderDetailBuilder) OrderId(orderId string) *H5OrderDetailBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}

func (builder *H5OrderDetailBuilder) Build() *H5OrderDetail {
	data := &H5OrderDetail{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	return data
}
