package v1

// DomesticFlightIdData 国内机票订单信息消息
type DomesticFlightIdData struct {
	OrderIds *string   `json:"order_ids,omitempty"` // 国内机票订单号，以英文逗号连接
	Page     *PageInfo `json:"page,omitempty"`
}

type DomesticFlightIdDataBuilder struct {
	orderIds    string // 国内机票订单号，以英文逗号连接
	orderIdsSet bool
	page        PageInfo
	pageSet     bool
}

func NewDomesticFlightIdDataBuilder() *DomesticFlightIdDataBuilder {
	return &DomesticFlightIdDataBuilder{}
}
func (builder *DomesticFlightIdDataBuilder) OrderIds(orderIds string) *DomesticFlightIdDataBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}
func (builder *DomesticFlightIdDataBuilder) Page(page PageInfo) *DomesticFlightIdDataBuilder {
	builder.page = page
	builder.pageSet = true
	return builder
}

func (builder *DomesticFlightIdDataBuilder) Build() *DomesticFlightIdData {
	data := &DomesticFlightIdData{}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	if builder.pageSet {
		data.Page = &builder.page
	}
	return data
}
