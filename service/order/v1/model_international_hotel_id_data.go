package v1

// InternationalHotelIdData 国际酒店订单信息消息
type InternationalHotelIdData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 国际酒店订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type InternationalHotelIdDataBuilder struct {
	orderIds    string // 国际酒店订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewInternationalHotelIdDataBuilder() *InternationalHotelIdDataBuilder {
	return &InternationalHotelIdDataBuilder{}
}
func (builder *InternationalHotelIdDataBuilder) OrderIds(orderIds string) *InternationalHotelIdDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *InternationalHotelIdDataBuilder) Page(page PageInfo) *InternationalHotelIdDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *InternationalHotelIdDataBuilder) Build() *InternationalHotelIdData {
	data := &InternationalHotelIdData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
