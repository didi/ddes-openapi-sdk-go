package v1

// InternationalFlightData 机票订单详情 ～ 国际机票数据
type InternationalFlightData struct {
	OrderList []InternationalFlightOrderListItem `json:"order_list,omitempty"` // 国际机票数据-订单列表
}

type InternationalFlightDataBuilder struct {
	orderList    []InternationalFlightOrderListItem // 国际机票数据-订单列表
	orderListSet bool
}

func NewInternationalFlightDataBuilder() *InternationalFlightDataBuilder {
	return &InternationalFlightDataBuilder{}
}
func (builder *InternationalFlightDataBuilder) OrderList(orderList []InternationalFlightOrderListItem) *InternationalFlightDataBuilder {
	builder.orderList = orderList
	builder.orderListSet = true
	return builder
}

func (builder *InternationalFlightDataBuilder) Build() *InternationalFlightData {
	data := &InternationalFlightData{}
	if builder.orderListSet {
		data.OrderList = builder.orderList
	}
	return data
}
