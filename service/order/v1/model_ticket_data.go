package v1

// TicketData 票信息 ～ 直达火车票
type TicketData struct {
	SeatType             *string `json:"seat_type,omitempty"`               // 席别码，1:硬座,2:软座,3:硬卧,4:软卧,5:包厢硬卧,6:高级软卧,7:一等软座,8:二等软座,9:商务座,A:高级动卧,B:混编硬座,C:混编硬卧,D:包厢软座,E:特等软座,F:动卧,G:二人软包,H:一人软包,I:一等卧,J:二等卧,K:混编软座,L:混编软卧,M:一等座,O:二等座,P:特等座,Q:多功能座,S:二等包座
	SeatTypeName         *string `json:"seat_type_name,omitempty"`          // 席别码名称
	TicketNum            *string `json:"ticket_num,omitempty"`              // 余票张数，大于 30 张返回有，大于 0 小于等于 30 显示实际张数
	TicketPrice          *int32  `json:"ticket_price,omitempty"`            // 票面价格，单位分
	IsSupportChooseSeats *int32  `json:"is_support_choose_seats,omitempty"` // 该席别是否支 持选座，0 - 否，1 - 是
	IsCanHB              *int32  `json:"is_can_HB,omitempty"`               // 该席别候补状态，0 表示该席别不支持候补;1 表示该席别可候补;2 表示当前该席别候补人数已达上限，无法候补。
	Discount             *int32  `json:"discount,omitempty"`                // 该席别是否是 折扣车票，0 - 否，1 - 是
	DiscountNum          *string `json:"discount_num,omitempty"`            // 具体折扣，9.5表示9.5折，-表示无折扣
}

type TicketDataBuilder struct {
	seatType                string // 席别码，1:硬座,2:软座,3:硬卧,4:软卧,5:包厢硬卧,6:高级软卧,7:一等软座,8:二等软座,9:商务座,A:高级动卧,B:混编硬座,C:混编硬卧,D:包厢软座,E:特等软座,F:动卧,G:二人软包,H:一人软包,I:一等卧,J:二等卧,K:混编软座,L:混编软卧,M:一等座,O:二等座,P:特等座,Q:多功能座,S:二等包座
	seatTypeSet             bool
	seatTypeName            string // 席别码名称
	seatTypeNameSet         bool
	ticketNum               string // 余票张数，大于 30 张返回有，大于 0 小于等于 30 显示实际张数
	ticketNumSet            bool
	ticketPrice             int32 // 票面价格，单位分
	ticketPriceSet          bool
	isSupportChooseSeats    int32 // 该席别是否支 持选座，0 - 否，1 - 是
	isSupportChooseSeatsSet bool
	isCanHB                 int32 // 该席别候补状态，0 表示该席别不支持候补;1 表示该席别可候补;2 表示当前该席别候补人数已达上限，无法候补。
	isCanHBSet              bool
	discount                int32 // 该席别是否是 折扣车票，0 - 否，1 - 是
	discountSet             bool
	discountNum             string // 具体折扣，9.5表示9.5折，-表示无折扣
	discountNumSet          bool
}

func NewTicketDataBuilder() *TicketDataBuilder {
	return &TicketDataBuilder{}
}
func (builder *TicketDataBuilder) SeatType(seatType string) *TicketDataBuilder {
	builder.seatType = seatType
	builder.seatTypeSet = true
	return builder
}
func (builder *TicketDataBuilder) SeatTypeName(seatTypeName string) *TicketDataBuilder {
	builder.seatTypeName = seatTypeName
	builder.seatTypeNameSet = true
	return builder
}
func (builder *TicketDataBuilder) TicketNum(ticketNum string) *TicketDataBuilder {
	builder.ticketNum = ticketNum
	builder.ticketNumSet = true
	return builder
}
func (builder *TicketDataBuilder) TicketPrice(ticketPrice int32) *TicketDataBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *TicketDataBuilder) IsSupportChooseSeats(isSupportChooseSeats int32) *TicketDataBuilder {
	builder.isSupportChooseSeats = isSupportChooseSeats
	builder.isSupportChooseSeatsSet = true
	return builder
}
func (builder *TicketDataBuilder) IsCanHB(isCanHB int32) *TicketDataBuilder {
	builder.isCanHB = isCanHB
	builder.isCanHBSet = true
	return builder
}
func (builder *TicketDataBuilder) Discount(discount int32) *TicketDataBuilder {
	builder.discount = discount
	builder.discountSet = true
	return builder
}
func (builder *TicketDataBuilder) DiscountNum(discountNum string) *TicketDataBuilder {
	builder.discountNum = discountNum
	builder.discountNumSet = true
	return builder
}

func (builder *TicketDataBuilder) Build() *TicketData {
	data := &TicketData{}
	if builder.seatTypeSet {
		data.SeatType = &builder.seatType
	}
	if builder.seatTypeNameSet {
		data.SeatTypeName = &builder.seatTypeName
	}
	if builder.ticketNumSet {
		data.TicketNum = &builder.ticketNum
	}
	if builder.ticketPriceSet {
		data.TicketPrice = &builder.ticketPrice
	}
	if builder.isSupportChooseSeatsSet {
		data.IsSupportChooseSeats = &builder.isSupportChooseSeats
	}
	if builder.isCanHBSet {
		data.IsCanHB = &builder.isCanHB
	}
	if builder.discountSet {
		data.Discount = &builder.discount
	}
	if builder.discountNumSet {
		data.DiscountNum = &builder.discountNum
	}
	return data
}
