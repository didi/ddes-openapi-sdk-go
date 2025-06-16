package v1

// TrainData 火车票订单信息消息
type TrainData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 火车票订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type TrainDataBuilder struct {
	orderIds    string // 火车票订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewTrainDataBuilder() *TrainDataBuilder {
	return &TrainDataBuilder{}
}
func (builder *TrainDataBuilder) OrderIds(orderIds string) *TrainDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *TrainDataBuilder) Page(page PageInfo) *TrainDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *TrainDataBuilder) Build() *TrainData {
	data := &TrainData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
