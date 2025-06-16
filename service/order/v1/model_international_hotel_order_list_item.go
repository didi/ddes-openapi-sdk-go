package v1

// InternationalHotelOrderListItem 酒店信息（国际）
type InternationalHotelOrderListItem struct {
	OrderInfo     *HotelOrder                   `json:"order_info,omitempty"`
	PriceInfo     *InterHotelOrderPriceInfo     `json:"price_info,omitempty"`
	BaseInfo      *InterHotelOrderBaseInfo      `json:"base_info,omitempty"`
	PassengerList []HotelOrderPassengerListItem `json:"passenger_list,omitempty"` // 出行人信息
	SequenceList  []HotelOrderSequenceListItem  `json:"sequence_list,omitempty"`  // 入住房间列表
	RcList        []HotelOrderRcInfo            `json:"rc_list,omitempty"`        // 管控信息
}

type InternationalHotelOrderListItemBuilder struct {
	orderInfo        HotelOrder
	orderInfoSet     bool
	priceInfo        InterHotelOrderPriceInfo
	priceInfoSet     bool
	baseInfo         InterHotelOrderBaseInfo
	baseInfoSet      bool
	passengerList    []HotelOrderPassengerListItem // 出行人信息
	passengerListSet bool
	sequenceList     []HotelOrderSequenceListItem // 入住房间列表
	sequenceListSet  bool
	rcList           []HotelOrderRcInfo // 管控信息
	rcListSet        bool
}

func NewInternationalHotelOrderListItemBuilder() *InternationalHotelOrderListItemBuilder {
	return &InternationalHotelOrderListItemBuilder{}
}
func (builder *InternationalHotelOrderListItemBuilder) OrderInfo(orderInfo HotelOrder) *InternationalHotelOrderListItemBuilder {
	builder.orderInfo = orderInfo
	builder.orderInfoSet = true
	return builder
}
func (builder *InternationalHotelOrderListItemBuilder) PriceInfo(priceInfo InterHotelOrderPriceInfo) *InternationalHotelOrderListItemBuilder {
	builder.priceInfo = priceInfo
	builder.priceInfoSet = true
	return builder
}
func (builder *InternationalHotelOrderListItemBuilder) BaseInfo(baseInfo InterHotelOrderBaseInfo) *InternationalHotelOrderListItemBuilder {
	builder.baseInfo = baseInfo
	builder.baseInfoSet = true
	return builder
}
func (builder *InternationalHotelOrderListItemBuilder) PassengerList(passengerList []HotelOrderPassengerListItem) *InternationalHotelOrderListItemBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *InternationalHotelOrderListItemBuilder) SequenceList(sequenceList []HotelOrderSequenceListItem) *InternationalHotelOrderListItemBuilder {
	builder.sequenceList = sequenceList
	builder.sequenceListSet = true
	return builder
}
func (builder *InternationalHotelOrderListItemBuilder) RcList(rcList []HotelOrderRcInfo) *InternationalHotelOrderListItemBuilder {
	builder.rcList = rcList
	builder.rcListSet = true
	return builder
}

func (builder *InternationalHotelOrderListItemBuilder) Build() *InternationalHotelOrderListItem {
	data := &InternationalHotelOrderListItem{}
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
