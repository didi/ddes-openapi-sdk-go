package v1

// DomesticFlightData 机票订单详情 ～ 国内机票数据
type DomesticFlightData struct {
	OrderList []DomesticFlightOrderListItem `json:"order_list,omitempty"` // 国内机票订单列表
}

type DomesticFlightDataBuilder struct {
	orderList    []DomesticFlightOrderListItem // 国内机票订单列表
	orderListSet bool
}

func NewDomesticFlightDataBuilder() *DomesticFlightDataBuilder {
	return &DomesticFlightDataBuilder{}
}
func (builder *DomesticFlightDataBuilder) OrderList(orderList []DomesticFlightOrderListItem) *DomesticFlightDataBuilder {
	builder.orderList = orderList
	builder.orderListSet = true
	return builder
}

func (builder *DomesticFlightDataBuilder) Build() *DomesticFlightData {
	data := &DomesticFlightData{}
	if builder.orderListSet {
		data.OrderList = builder.orderList
	}
	return data
}
