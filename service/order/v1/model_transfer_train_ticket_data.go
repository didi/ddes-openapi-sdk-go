package v1

// TransferTrainTicketData 票信息 ～ 中转火车票
type TransferTrainTicketData struct {
	TicketNum            *string `json:"ticket_num,omitempty"`              // 票张数
	TicketPrice          *int32  `json:"ticket_price,omitempty"`            // 票价（卧铺场景，使用下铺金额）
	ParPrice             *int32  `json:"par_price,omitempty"`               // 票价
	IsSupportChooseSeats *string `json:"is_support_choose_seats,omitempty"` // 是否支持选座，1：是 0：否
	IsCanHB              *string `json:"is_can_HB,omitempty"`               // 该席别候补状态，0 表示该席别不支持候补;1 表示该席别可候补;2 表示当前该席别候补人数已达上限，无法候补。
	SeatTypeName         *string `json:"seat_type_name,omitempty"`          // 坐席名称，与席别码对应
	SeatType             *string `json:"seat_type,omitempty"`               // 席别码，1:硬座,2:软座,3:硬卧,4:软卧,5:包厢硬卧,6:高级软卧,7:一等软座,8:二等软座,9:商务座,A:高级动卧,B:混编硬座,C:混编硬卧,D:包厢软座,E:特等软座,F:动卧,G:二人软包,H:一人软包,I:一等卧,J:二等卧,K:混编软座,L:混编软卧,M:一等座,O:二等座,P:特等座,Q:多功能座,S:二等包座
	Discount             *string `json:"discount,omitempty"`                // 该席别是否是 折扣车票，0 - 否，1 - 是
	DiscountNum          *string `json:"discount_num,omitempty"`            // 具体折扣，9.5表示9.5折，-表示无折扣
}

type TransferTrainTicketDataBuilder struct {
	ticketNum               string // 票张数
	ticketNumSet            bool
	ticketPrice             int32 // 票价（卧铺场景，使用下铺金额）
	ticketPriceSet          bool
	parPrice                int32 // 票价
	parPriceSet             bool
	isSupportChooseSeats    string // 是否支持选座，1：是 0：否
	isSupportChooseSeatsSet bool
	isCanHB                 string // 该席别候补状态，0 表示该席别不支持候补;1 表示该席别可候补;2 表示当前该席别候补人数已达上限，无法候补。
	isCanHBSet              bool
	seatTypeName            string // 坐席名称，与席别码对应
	seatTypeNameSet         bool
	seatType                string // 席别码，1:硬座,2:软座,3:硬卧,4:软卧,5:包厢硬卧,6:高级软卧,7:一等软座,8:二等软座,9:商务座,A:高级动卧,B:混编硬座,C:混编硬卧,D:包厢软座,E:特等软座,F:动卧,G:二人软包,H:一人软包,I:一等卧,J:二等卧,K:混编软座,L:混编软卧,M:一等座,O:二等座,P:特等座,Q:多功能座,S:二等包座
	seatTypeSet             bool
	discount                string // 该席别是否是 折扣车票，0 - 否，1 - 是
	discountSet             bool
	discountNum             string // 具体折扣，9.5表示9.5折，-表示无折扣
	discountNumSet          bool
}

func NewTransferTrainTicketDataBuilder() *TransferTrainTicketDataBuilder {
	return &TransferTrainTicketDataBuilder{}
}
func (builder *TransferTrainTicketDataBuilder) TicketNum(ticketNum string) *TransferTrainTicketDataBuilder {
	builder.ticketNum = ticketNum
	builder.ticketNumSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) TicketPrice(ticketPrice int32) *TransferTrainTicketDataBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) ParPrice(parPrice int32) *TransferTrainTicketDataBuilder {
	builder.parPrice = parPrice
	builder.parPriceSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) IsSupportChooseSeats(isSupportChooseSeats string) *TransferTrainTicketDataBuilder {
	builder.isSupportChooseSeats = isSupportChooseSeats
	builder.isSupportChooseSeatsSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) IsCanHB(isCanHB string) *TransferTrainTicketDataBuilder {
	builder.isCanHB = isCanHB
	builder.isCanHBSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) SeatTypeName(seatTypeName string) *TransferTrainTicketDataBuilder {
	builder.seatTypeName = seatTypeName
	builder.seatTypeNameSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) SeatType(seatType string) *TransferTrainTicketDataBuilder {
	builder.seatType = seatType
	builder.seatTypeSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) Discount(discount string) *TransferTrainTicketDataBuilder {
	builder.discount = discount
	builder.discountSet = true
	return builder
}
func (builder *TransferTrainTicketDataBuilder) DiscountNum(discountNum string) *TransferTrainTicketDataBuilder {
	builder.discountNum = discountNum
	builder.discountNumSet = true
	return builder
}

func (builder *TransferTrainTicketDataBuilder) Build() *TransferTrainTicketData {
	data := &TransferTrainTicketData{}
	if builder.ticketNumSet {
		data.TicketNum = &builder.ticketNum
	}
	if builder.ticketPriceSet {
		data.TicketPrice = &builder.ticketPrice
	}
	if builder.parPriceSet {
		data.ParPrice = &builder.parPrice
	}
	if builder.isSupportChooseSeatsSet {
		data.IsSupportChooseSeats = &builder.isSupportChooseSeats
	}
	if builder.isCanHBSet {
		data.IsCanHB = &builder.isCanHB
	}
	if builder.seatTypeNameSet {
		data.SeatTypeName = &builder.seatTypeName
	}
	if builder.seatTypeSet {
		data.SeatType = &builder.seatType
	}
	if builder.discountSet {
		data.Discount = &builder.discount
	}
	if builder.discountNumSet {
		data.DiscountNum = &builder.discountNum
	}
	return data
}
