package v1

// DomesticFlightOrderListItem 国内机票信息
type DomesticFlightOrderListItem struct {
	OrderInfo     *DomesticFlightOrderInfo            `json:"order_info,omitempty"`
	TicketList    []DomesticFlightOrderTicketListItem `json:"ticket_list,omitempty"`    // 票张详情
	ChangeList    []DomesticFlightOrderChangeListItem `json:"change_list,omitempty"`    // 改签信息
	RefundList    []DomesticFlightOrderRefundListItem `json:"refund_list,omitempty"`    // 退票信息
	PassengerList []FlightOrderPassengerItem          `json:"passenger_list,omitempty"` // 出行人信息
	RcList        []FlightOrderRcListItem             `json:"rc_list,omitempty"`        // 管控信息
	PriceInfo     *DomesticFlightOrderPriceInfo       `json:"price_info,omitempty"`
	SrvPackInfo   *FlightOrderSrvPackInfo             `json:"srv_pack_info,omitempty"`
}

type DomesticFlightOrderListItemBuilder struct {
	orderInfo        DomesticFlightOrderInfo
	orderInfoSet     bool
	ticketList       []DomesticFlightOrderTicketListItem // 票张详情
	ticketListSet    bool
	changeList       []DomesticFlightOrderChangeListItem // 改签信息
	changeListSet    bool
	refundList       []DomesticFlightOrderRefundListItem // 退票信息
	refundListSet    bool
	passengerList    []FlightOrderPassengerItem // 出行人信息
	passengerListSet bool
	rcList           []FlightOrderRcListItem // 管控信息
	rcListSet        bool
	priceInfo        DomesticFlightOrderPriceInfo
	priceInfoSet     bool
	srvPackInfo      FlightOrderSrvPackInfo
	srvPackInfoSet   bool
}

func NewDomesticFlightOrderListItemBuilder() *DomesticFlightOrderListItemBuilder {
	return &DomesticFlightOrderListItemBuilder{}
}
func (builder *DomesticFlightOrderListItemBuilder) OrderInfo(orderInfo DomesticFlightOrderInfo) *DomesticFlightOrderListItemBuilder {
	builder.orderInfo = orderInfo
	builder.orderInfoSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) TicketList(ticketList []DomesticFlightOrderTicketListItem) *DomesticFlightOrderListItemBuilder {
	builder.ticketList = ticketList
	builder.ticketListSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) ChangeList(changeList []DomesticFlightOrderChangeListItem) *DomesticFlightOrderListItemBuilder {
	builder.changeList = changeList
	builder.changeListSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) RefundList(refundList []DomesticFlightOrderRefundListItem) *DomesticFlightOrderListItemBuilder {
	builder.refundList = refundList
	builder.refundListSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) PassengerList(passengerList []FlightOrderPassengerItem) *DomesticFlightOrderListItemBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) RcList(rcList []FlightOrderRcListItem) *DomesticFlightOrderListItemBuilder {
	builder.rcList = rcList
	builder.rcListSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) PriceInfo(priceInfo DomesticFlightOrderPriceInfo) *DomesticFlightOrderListItemBuilder {
	builder.priceInfo = priceInfo
	builder.priceInfoSet = true
	return builder
}
func (builder *DomesticFlightOrderListItemBuilder) SrvPackInfo(srvPackInfo FlightOrderSrvPackInfo) *DomesticFlightOrderListItemBuilder {
	builder.srvPackInfo = srvPackInfo
	builder.srvPackInfoSet = true
	return builder
}

func (builder *DomesticFlightOrderListItemBuilder) Build() *DomesticFlightOrderListItem {
	data := &DomesticFlightOrderListItem{}
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
