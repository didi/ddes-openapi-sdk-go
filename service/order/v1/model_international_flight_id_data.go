package v1

// InternationalFlightIdData 国际机票订单信息消息
type InternationalFlightIdData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 国际机票订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type InternationalFlightIdDataBuilder struct {
	orderIds    string // 国际机票订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewInternationalFlightIdDataBuilder() *InternationalFlightIdDataBuilder {
	return &InternationalFlightIdDataBuilder{}
}
func (builder *InternationalFlightIdDataBuilder) OrderIds(orderIds string) *InternationalFlightIdDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *InternationalFlightIdDataBuilder) Page(page PageInfo) *InternationalFlightIdDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *InternationalFlightIdDataBuilder) Build() *InternationalFlightIdData {
	data := &InternationalFlightIdData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
