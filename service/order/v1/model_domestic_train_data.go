package v1

// DomesticTrainData 国内火车票数据
type DomesticTrainData struct {
	OrderList []TrainOrderListItem `json:"order_list,omitempty"` // 国内火车票数据-订单列表，国内订单返回有值
}

type DomesticTrainDataBuilder struct {
	orderList    []TrainOrderListItem // 国内火车票数据-订单列表，国内订单返回有值
	orderListSet bool
}

func NewDomesticTrainDataBuilder() *DomesticTrainDataBuilder {
	return &DomesticTrainDataBuilder{}
}
func (builder *DomesticTrainDataBuilder) OrderList(orderList []TrainOrderListItem) *DomesticTrainDataBuilder {
	builder.orderList = orderList
	builder.orderListSet = true
	return builder
}

func (builder *DomesticTrainDataBuilder) Build() *DomesticTrainData {
	data := &DomesticTrainData{}
	if builder.orderListSet {
		data.OrderList = builder.orderList
	}
	return data
}
