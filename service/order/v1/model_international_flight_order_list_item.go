package v1

// InternationalFlightOrderListItem 国际机票信息
type InternationalFlightOrderListItem struct {
	OrderInfo     *InterFlightOrderInfo            `json:"order_info,omitempty"`
	TicketList    []InterFlightOrderTicketListItem `json:"ticket_list,omitempty"`    // 票张详情
	ChangeList    []InterFlightOrderChangeListItem `json:"change_list,omitempty"`    // 改签信息
	RefundList    []InterFlightOrderRefundListItem `json:"refund_list,omitempty"`    // 退票信息
	PassengerList []FlightOrderPassengerItem       `json:"passenger_list,omitempty"` // 出行人信息
	RcList        []FlightOrderRcListItem          `json:"rc_list,omitempty"`        // 管控信息
	PriceInfo     *InterFlightOrderPriceInfo       `json:"price_info,omitempty"`
	SrvPackInfo   *FlightOrderSrvPackInfo          `json:"srv_pack_info,omitempty"`
}

type InternationalFlightOrderListItemBuilder struct {
	orderInfo        InterFlightOrderInfo
	orderInfoSet     bool
	ticketList       []InterFlightOrderTicketListItem // 票张详情
	ticketListSet    bool
	changeList       []InterFlightOrderChangeListItem // 改签信息
	changeListSet    bool
	refundList       []InterFlightOrderRefundListItem // 退票信息
	refundListSet    bool
	passengerList    []FlightOrderPassengerItem // 出行人信息
	passengerListSet bool
	rcList           []FlightOrderRcListItem // 管控信息
	rcListSet        bool
	priceInfo        InterFlightOrderPriceInfo
	priceInfoSet     bool
	srvPackInfo      FlightOrderSrvPackInfo
	srvPackInfoSet   bool
}

func NewInternationalFlightOrderListItemBuilder() *InternationalFlightOrderListItemBuilder {
	return &InternationalFlightOrderListItemBuilder{}
}
func (builder *InternationalFlightOrderListItemBuilder) OrderInfo(orderInfo InterFlightOrderInfo) *InternationalFlightOrderListItemBuilder {
	builder.orderInfo = orderInfo
	builder.orderInfoSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) TicketList(ticketList []InterFlightOrderTicketListItem) *InternationalFlightOrderListItemBuilder {
	builder.ticketList = ticketList
	builder.ticketListSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) ChangeList(changeList []InterFlightOrderChangeListItem) *InternationalFlightOrderListItemBuilder {
	builder.changeList = changeList
	builder.changeListSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) RefundList(refundList []InterFlightOrderRefundListItem) *InternationalFlightOrderListItemBuilder {
	builder.refundList = refundList
	builder.refundListSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) PassengerList(passengerList []FlightOrderPassengerItem) *InternationalFlightOrderListItemBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) RcList(rcList []FlightOrderRcListItem) *InternationalFlightOrderListItemBuilder {
	builder.rcList = rcList
	builder.rcListSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) PriceInfo(priceInfo InterFlightOrderPriceInfo) *InternationalFlightOrderListItemBuilder {
	builder.priceInfo = priceInfo
	builder.priceInfoSet = true
	return builder
}
func (builder *InternationalFlightOrderListItemBuilder) SrvPackInfo(srvPackInfo FlightOrderSrvPackInfo) *InternationalFlightOrderListItemBuilder {
	builder.srvPackInfo = srvPackInfo
	builder.srvPackInfoSet = true
	return builder
}

func (builder *InternationalFlightOrderListItemBuilder) Build() *InternationalFlightOrderListItem {
	data := &InternationalFlightOrderListItem{}
	if builder.orderInfoSet {
		data.OrderInfo = &builder.orderInfo
	}
	if builder.ticketListSet {
		data.TicketList = builder.ticketList
	}
	if builder.changeListSet {
		data.ChangeList = builder.changeList
	}
	if builder.refundListSet {
		data.RefundList = builder.refundList
	}
	if builder.passengerListSet {
		data.PassengerList = builder.passengerList
	}
	if builder.rcListSet {
		data.RcList = builder.rcList
	}
	if builder.priceInfoSet {
		data.PriceInfo = &builder.priceInfo
	}
	if builder.srvPackInfoSet {
		data.SrvPackInfo = &builder.srvPackInfo
	}
	return data
}
