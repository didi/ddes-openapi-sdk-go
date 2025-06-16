package v1

// DomesticHotelIdData 国内酒店订单信息消息
type DomesticHotelIdData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 国内酒店订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type DomesticHotelIdDataBuilder struct {
	orderIds    string // 国内酒店订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewDomesticHotelIdDataBuilder() *DomesticHotelIdDataBuilder {
	return &DomesticHotelIdDataBuilder{}
}
func (builder *DomesticHotelIdDataBuilder) OrderIds(orderIds string) *DomesticHotelIdDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *DomesticHotelIdDataBuilder) Page(page PageInfo) *DomesticHotelIdDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *DomesticHotelIdDataBuilder) Build() *DomesticHotelIdData {
	data := &DomesticHotelIdData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
