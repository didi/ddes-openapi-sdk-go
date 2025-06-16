package v1

// DomesticHotelOrderListItem 酒店信息（国内）
type DomesticHotelOrderListItem struct {
	OrderInfo     *HotelOrder                   `json:"order_info,omitempty"`
	PriceInfo     *DomesticHotelOrderPriceInfo  `json:"price_info,omitempty"`
	BaseInfo      *DomesticHotelOrderBaseInfo   `json:"base_info,omitempty"`
	PassengerList []HotelOrderPassengerListItem `json:"passenger_list,omitempty"` // 出行人信息
	SequenceList  []HotelOrderSequenceListItem  `json:"sequence_list,omitempty"`  // 入住房间列表
	RcList        []HotelOrderRcInfo            `json:"rc_list,omitempty"`        // 管控信息
}

type DomesticHotelOrderListItemBuilder struct {
	orderInfo        HotelOrder
	orderInfoSet     bool
	priceInfo        DomesticHotelOrderPriceInfo
	priceInfoSet     bool
	baseInfo         DomesticHotelOrderBaseInfo
	baseInfoSet      bool
	passengerList    []HotelOrderPassengerListItem // 出行人信息
	passengerListSet bool
	sequenceList     []HotelOrderSequenceListItem // 入住房间列表
	sequenceListSet  bool
	rcList           []HotelOrderRcInfo // 管控信息
	rcListSet        bool
}

func NewDomesticHotelOrderListItemBuilder() *DomesticHotelOrderListItemBuilder {
	return &DomesticHotelOrderListItemBuilder{}
}
func (builder *DomesticHotelOrderListItemBuilder) OrderInfo(orderInfo HotelOrder) *DomesticHotelOrderListItemBuilder {
	builder.orderInfo = orderInfo
	builder.orderInfoSet = true
	return builder
}
func (builder *DomesticHotelOrderListItemBuilder) PriceInfo(priceInfo DomesticHotelOrderPriceInfo) *DomesticHotelOrderListItemBuilder {
	builder.priceInfo = priceInfo
	builder.priceInfoSet = true
	return builder
}
func (builder *DomesticHotelOrderListItemBuilder) BaseInfo(baseInfo DomesticHotelOrderBaseInfo) *DomesticHotelOrderListItemBuilder {
	builder.baseInfo = baseInfo
	builder.baseInfoSet = true
	return builder
}
func (builder *DomesticHotelOrderListItemBuilder) PassengerList(passengerList []HotelOrderPassengerListItem) *DomesticHotelOrderListItemBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *DomesticHotelOrderListItemBuilder) SequenceList(sequenceList []HotelOrderSequenceListItem) *DomesticHotelOrderListItemBuilder {
	builder.sequenceList = sequenceList
	builder.sequenceListSet = true
	return builder
}
func (builder *DomesticHotelOrderListItemBuilder) RcList(rcList []HotelOrderRcInfo) *DomesticHotelOrderListItemBuilder {
	builder.rcList = rcList
	builder.rcListSet = true
	return builder
}

func (builder *DomesticHotelOrderListItemBuilder) Build() *DomesticHotelOrderListItem {
	data := &DomesticHotelOrderListItem{}
	if builder.orderInfoSet {
		data.OrderInfo = &builder.orderInfo
	}
	if builder.priceInfoSet {
		data.PriceInfo = &builder.priceInfo
	}
	if builder.baseInfoSet {
		data.BaseInfo = &builder.baseInfo
	}
	if builder.passengerListSet {
		data.PassengerList = builder.passengerList
	}
	if builder.sequenceListSet {
		data.SequenceList = builder.sequenceList
	}
	if builder.rcListSet {
		data.RcList = builder.rcList
	}
	return data
}
