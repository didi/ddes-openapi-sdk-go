package v1

// TrainOrderListItem struct for TrainOrderListItem
type TrainOrderListItem struct {
	OrderInfo     *TrainOrderInfo               `json:"order_info,omitempty"`
	PriceInfo     *TrainOrderPriceInfo          `json:"price_info,omitempty"`
	TicketList    []TrainOrderTicketListItem    `json:"ticket_list,omitempty"`    // 火车票基本信息
	GrabInfo      []TrainOrderGrabInfo          `json:"grab_info,omitempty"`      // 抢票基本信息
	ChangeList    []TrainOrderChangeListItem    `json:"change_list,omitempty"`    // 改签信息
	RefundList    []TrainOrderRefundListItem    `json:"refund_list,omitempty"`    // 退票信息
	PassengerList []TrainOrderPassengerListItem `json:"passenger_list,omitempty"` // 出行人信息
	RcList        []TrainOrderRcListItem        `json:"rc_list,omitempty"`        // 管控信息
}

type TrainOrderListItemBuilder struct {
	orderInfo        TrainOrderInfo
	orderInfoSet     bool
	priceInfo        TrainOrderPriceInfo
	priceInfoSet     bool
	ticketList       []TrainOrderTicketListItem // 火车票基本信息
	ticketListSet    bool
	grabInfo         []TrainOrderGrabInfo // 抢票基本信息
	grabInfoSet      bool
	changeList       []TrainOrderChangeListItem // 改签信息
	changeListSet    bool
	refundList       []TrainOrderRefundListItem // 退票信息
	refundListSet    bool
	passengerList    []TrainOrderPassengerListItem // 出行人信息
	passengerListSet bool
	rcList           []TrainOrderRcListItem // 管控信息
	rcListSet        bool
}

func NewTrainOrderListItemBuilder() *TrainOrderListItemBuilder {
	return &TrainOrderListItemBuilder{}
}
func (builder *TrainOrderListItemBuilder) OrderInfo(orderInfo TrainOrderInfo) *TrainOrderListItemBuilder {
	builder.orderInfo = orderInfo
	builder.orderInfoSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) PriceInfo(priceInfo TrainOrderPriceInfo) *TrainOrderListItemBuilder {
	builder.priceInfo = priceInfo
	builder.priceInfoSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) TicketList(ticketList []TrainOrderTicketListItem) *TrainOrderListItemBuilder {
	builder.ticketList = ticketList
	builder.ticketListSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) GrabInfo(grabInfo []TrainOrderGrabInfo) *TrainOrderListItemBuilder {
	builder.grabInfo = grabInfo
	builder.grabInfoSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) ChangeList(changeList []TrainOrderChangeListItem) *TrainOrderListItemBuilder {
	builder.changeList = changeList
	builder.changeListSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) RefundList(refundList []TrainOrderRefundListItem) *TrainOrderListItemBuilder {
	builder.refundList = refundList
	builder.refundListSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) PassengerList(passengerList []TrainOrderPassengerListItem) *TrainOrderListItemBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *TrainOrderListItemBuilder) RcList(rcList []TrainOrderRcListItem) *TrainOrderListItemBuilder {
	builder.rcList = rcList
	builder.rcListSet = true
	return builder
}

func (builder *TrainOrderListItemBuilder) Build() *TrainOrderListItem {
	data := &TrainOrderListItem{}
	if builder.orderInfoSet {
		data.OrderInfo = &builder.orderInfo
	}
	if builder.priceInfoSet {
		data.PriceInfo = &builder.priceInfo
	}
	if builder.ticketListSet {
		data.TicketList = builder.ticketList
	}
	if builder.grabInfoSet {
		data.GrabInfo = builder.grabInfo
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
	return data
}
