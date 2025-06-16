package v1

// SrvPackListItem struct for SrvPackListItem
type SrvPackListItem struct {
	PassengerTravelerId *int32                  `json:"passenger_traveler_id,omitempty"` // 服务包使用人员
	SrvPackOrderId      *int32                  `json:"srv_pack_order_id,omitempty"`     // 服务包订单号
	PriceAmount         *int32                  `json:"price_amount,omitempty"`          // 服务包金额，单位分
	Data                []ServicePackageContent `json:"data,omitempty"`                  // 服务包内容
}

type SrvPackListItemBuilder struct {
	passengerTravelerId    int32 // 服务包使用人员
	passengerTravelerIdSet bool
	srvPackOrderId         int32 // 服务包订单号
	srvPackOrderIdSet      bool
	priceAmount            int32 // 服务包金额，单位分
	priceAmountSet         bool
	data                   []ServicePackageContent // 服务包内容
	dataSet                bool
}

func NewSrvPackListItemBuilder() *SrvPackListItemBuilder {
	return &SrvPackListItemBuilder{}
}
func (builder *SrvPackListItemBuilder) PassengerTravelerId(passengerTravelerId int32) *SrvPackListItemBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *SrvPackListItemBuilder) SrvPackOrderId(srvPackOrderId int32) *SrvPackListItemBuilder {
	builder.srvPackOrderId = srvPackOrderId
	builder.srvPackOrderIdSet = true
	return builder
}
func (builder *SrvPackListItemBuilder) PriceAmount(priceAmount int32) *SrvPackListItemBuilder {
	builder.priceAmount = priceAmount
	builder.priceAmountSet = true
	return builder
}
func (builder *SrvPackListItemBuilder) Data(data []ServicePackageContent) *SrvPackListItemBuilder {
	builder.data = data
	builder.dataSet = true
	return builder
}

func (builder *SrvPackListItemBuilder) Build() *SrvPackListItem {
	data := &SrvPackListItem{}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
	}
	if builder.srvPackOrderIdSet {
		data.SrvPackOrderId = &builder.srvPackOrderId
	}
	if builder.priceAmountSet {
		data.PriceAmount = &builder.priceAmount
	}
	if builder.dataSet {
		data.Data = builder.data
	}
	return data
}
