package v1

// DomesticFlightOrderTicketInfo 原始票信息
type DomesticFlightOrderTicketInfo struct {
	TicketUniqueKey        *string                       `json:"ticket_unique_key,omitempty"`        // 机票唯一标记，非票号, 不包含改签票信息，预订时的票张信息
	TicketNumber           *string                       `json:"ticket_number,omitempty"`            // 票号，廉价航空直接给下游的票号，可以为空
	PassengerTravelerId    *string                       `json:"passenger_traveler_id,omitempty"`    // 乘客ID，memberid或travelerid用于关联乘客信息
	TicketStatus           *string                       `json:"ticket_status,omitempty"`            // 票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	Discount               *int32                        `json:"discount,omitempty"`                 // 机票折扣，举例：85  八五折   全价显示为 100
	IsAccountPrice         *int32                        `json:"is_account_price,omitempty"`         // 是否协议价，枚举数字 ：0 否 1 是
	SalePrice              *int32                        `json:"sale_price,omitempty"`               // 机票销售价，单位：分 包含立减金
	CutFee                 *int32                        `json:"cut_fee,omitempty"`                  // 机票订立减优惠，单位：分
	TicketPrice            *int32                        `json:"ticket_price,omitempty"`             // 机票实际出票价格，单位：分
	ConstructionFee        *int32                        `json:"construction_fee,omitempty"`         // 机建费，单位：分
	OilFee                 *int32                        `json:"oil_fee,omitempty"`                  // 燃油费，单位：分
	ServiceFee             *int32                        `json:"service_fee,omitempty"`              // 出票平台使用费，单位：分
	AsynchronousServiceFee *int32                        `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	SequenceList           []FlightOrderSequenceListItem `json:"sequence_list,omitempty"`            // 航段信息
}

type DomesticFlightOrderTicketInfoBuilder struct {
	ticketUniqueKey           string // 机票唯一标记，非票号, 不包含改签票信息，预订时的票张信息
	ticketUniqueKeySet        bool
	ticketNumber              string // 票号，廉价航空直接给下游的票号，可以为空
	ticketNumberSet           bool
	passengerTravelerId       string // 乘客ID，memberid或travelerid用于关联乘客信息
	passengerTravelerIdSet    bool
	ticketStatus              string // 票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	ticketStatusSet           bool
	discount                  int32 // 机票折扣，举例：85  八五折   全价显示为 100
	discountSet               bool
	isAccountPrice            int32 // 是否协议价，枚举数字 ：0 否 1 是
	isAccountPriceSet         bool
	salePrice                 int32 // 机票销售价，单位：分 包含立减金
	salePriceSet              bool
	cutFee                    int32 // 机票订立减优惠，单位：分
	cutFeeSet                 bool
	ticketPrice               int32 // 机票实际出票价格，单位：分
	ticketPriceSet            bool
	constructionFee           int32 // 机建费，单位：分
	constructionFeeSet        bool
	oilFee                    int32 // 燃油费，单位：分
	oilFeeSet                 bool
	serviceFee                int32 // 出票平台使用费，单位：分
	serviceFeeSet             bool
	asynchronousServiceFee    int32 // 出票平台使用费（异步），单位：分
	asynchronousServiceFeeSet bool
	sequenceList              []FlightOrderSequenceListItem // 航段信息
	sequenceListSet           bool
}

func NewDomesticFlightOrderTicketInfoBuilder() *DomesticFlightOrderTicketInfoBuilder {
	return &DomesticFlightOrderTicketInfoBuilder{}
}
func (builder *DomesticFlightOrderTicketInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *DomesticFlightOrderTicketInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) TicketNumber(ticketNumber string) *DomesticFlightOrderTicketInfoBuilder {
	builder.ticketNumber = ticketNumber
	builder.ticketNumberSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) PassengerTravelerId(passengerTravelerId string) *DomesticFlightOrderTicketInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) TicketStatus(ticketStatus string) *DomesticFlightOrderTicketInfoBuilder {
	builder.ticketStatus = ticketStatus
	builder.ticketStatusSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) Discount(discount int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.discount = discount
	builder.discountSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) IsAccountPrice(isAccountPrice int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.isAccountPrice = isAccountPrice
	builder.isAccountPriceSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) SalePrice(salePrice int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.salePrice = salePrice
	builder.salePriceSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) CutFee(cutFee int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) TicketPrice(ticketPrice int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) ConstructionFee(constructionFee int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.constructionFee = constructionFee
	builder.constructionFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) OilFee(oilFee int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.oilFee = oilFee
	builder.oilFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) ServiceFee(serviceFee int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *DomesticFlightOrderTicketInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderTicketInfoBuilder) SequenceList(sequenceList []FlightOrderSequenceListItem) *DomesticFlightOrderTicketInfoBuilder {
	builder.sequenceList = sequenceList
	builder.sequenceListSet = true
	return builder
}

func (builder *DomesticFlightOrderTicketInfoBuilder) Build() *DomesticFlightOrderTicketInfo {
	data := &DomesticFlightOrderTicketInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.ticketNumberSet {
		data.TicketNumber = &builder.ticketNumber
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
	}
	if builder.ticketStatusSet {
		data.TicketStatus = &builder.ticketStatus
	}
	if builder.discountSet {
		data.Discount = &builder.discount
	}
	if builder.isAccountPriceSet {
		data.IsAccountPrice = &builder.isAccountPrice
	}
	if builder.salePriceSet {
		data.SalePrice = &builder.salePrice
	}
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	if builder.ticketPriceSet {
		data.TicketPrice = &builder.ticketPrice
	}
	if builder.constructionFeeSet {
		data.ConstructionFee = &builder.constructionFee
	}
	if builder.oilFeeSet {
		data.OilFee = &builder.oilFee
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.asynchronousServiceFeeSet {
		data.AsynchronousServiceFee = &builder.asynchronousServiceFee
	}
	if builder.sequenceListSet {
		data.SequenceList = builder.sequenceList
	}
	return data
}
