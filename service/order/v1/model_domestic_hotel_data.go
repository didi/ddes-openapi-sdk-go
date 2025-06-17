package v1

// DomesticHotelData struct for DomesticHotelData
type DomesticHotelData struct {
	OrderList []DomesticHotelOrderListItem `json:"order_list,omitempty"` // 国内酒店订单列表
}

type DomesticHotelDataBuilder struct {
	orderList    []DomesticHotelOrderListItem // 国内酒店订单列表
	orderListSet bool
}

func NewDomesticHotelDataBuilder() *DomesticHotelDataBuilder {
	return &DomesticHotelDataBuilder{}
}
func (builder *DomesticHotelDataBuilder) OrderList(orderList []DomesticHotelOrderListItem) *DomesticHotelDataBuilder {
	builder.orderList = orderList
	builder.orderListSet = true
	return builder
}

func (builder *DomesticHotelDataBuilder) Build() *DomesticHotelData {
	data := &DomesticHotelData{}
	if builder.orderListSet {
		data.OrderList = builder.orderList
	}
	return data
}
