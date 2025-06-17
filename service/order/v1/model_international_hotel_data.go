package v1

// InternationalHotelData struct for InternationalHotelData
type InternationalHotelData struct {
	OrderList []InternationalHotelOrderListItem `json:"order_list,omitempty"` // 国际机票数据-订单列表
}

type InternationalHotelDataBuilder struct {
	orderList    []InternationalHotelOrderListItem // 国际机票数据-订单列表
	orderListSet bool
}

func NewInternationalHotelDataBuilder() *InternationalHotelDataBuilder {
	return &InternationalHotelDataBuilder{}
}
func (builder *InternationalHotelDataBuilder) OrderList(orderList []InternationalHotelOrderListItem) *InternationalHotelDataBuilder {
	builder.orderList = orderList
	builder.orderListSet = true
	return builder
}

func (builder *InternationalHotelDataBuilder) Build() *InternationalHotelData {
	data := &InternationalHotelData{}
	if builder.orderListSet {
		data.OrderList = builder.orderList
	}
	return data
}
