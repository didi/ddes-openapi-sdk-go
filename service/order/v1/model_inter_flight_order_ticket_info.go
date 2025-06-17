package v1

// InterFlightOrderTicketInfo 原始票信息
type InterFlightOrderTicketInfo struct {
	TicketUniqueKey        *string                       `json:"ticket_unique_key,omitempty"`        // 机票唯一标记，非票号, 不包含改签票信息，预订时的票张信息
	TicketNumber           *string                       `json:"ticket_number,omitempty"`            // 票号，廉价航空直接给下游的票号，可以为空
	PassengerTravelerId    *string                       `json:"passenger_traveler_id,omitempty"`    // 乘客ID，memberid或travelerid用于关联乘客信息
	TicketStatus           *string                       `json:"ticket_status,omitempty"`            // 票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	IsAccountPrice         *int32                        `json:"is_account_price,omitempty"`         // 是否协议价，枚举数字 ：0 否 1 是
	SalePrice              *int32                        `json:"sale_price,omitempty"`               // 机票销售价，单位：分 包含立减金
	CutFee                 *int32                        `json:"cut_fee,omitempty"`                  // 机票订立减优惠，单位：分
	TicketPrice            *int32                        `json:"ticket_price,omitempty"`             // 机票实际出票价格，单位：分
	ConstructionFee        *int32                        `json:"construction_fee,omitempty"`         // 机建费，单位：分
	OilFee                 *int32                        `json:"oil_fee,omitempty"`                  // 燃油费，单位：分
	ServiceFee             *int32                        `json:"service_fee,omitempty"`              // 出票平台使用费，单位：分
	AsynchronousServiceFee *int32                        `json:"asynchronous_service_fee,omitempty"` // 出票平台使用费（异步），单位：分
	SequenceList           []FlightOrderSequenceListItem `json:"sequence_list,omitempty"`            // 航段信息
	Tax                    *int32                        `json:"tax,omitempty"`                      // 税，单位：分
}

type InterFlightOrderTicketInfoBuilder struct {
	ticketUniqueKey           string // 机票唯一标记，非票号, 不包含改签票信息，预订时的票张信息
	ticketUniqueKeySet        bool
	ticketNumber              string // 票号，廉价航空直接给下游的票号，可以为空
	ticketNumberSet           bool
	passengerTravelerId       string // 乘客ID，memberid或travelerid用于关联乘客信息
	passengerTravelerIdSet    bool
	ticketStatus              string // 票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	ticketStatusSet           bool
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
	tax                       int32 // 税，单位：分
	taxSet                    bool
}

func NewInterFlightOrderTicketInfoBuilder() *InterFlightOrderTicketInfoBuilder {
	return &InterFlightOrderTicketInfoBuilder{}
}
func (builder *InterFlightOrderTicketInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *InterFlightOrderTicketInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) TicketNumber(ticketNumber string) *InterFlightOrderTicketInfoBuilder {
	builder.ticketNumber = ticketNumber
	builder.ticketNumberSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) PassengerTravelerId(passengerTravelerId string) *InterFlightOrderTicketInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) TicketStatus(ticketStatus string) *InterFlightOrderTicketInfoBuilder {
	builder.ticketStatus = ticketStatus
	builder.ticketStatusSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) IsAccountPrice(isAccountPrice int32) *InterFlightOrderTicketInfoBuilder {
	builder.isAccountPrice = isAccountPrice
	builder.isAccountPriceSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) SalePrice(salePrice int32) *InterFlightOrderTicketInfoBuilder {
	builder.salePrice = salePrice
	builder.salePriceSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) CutFee(cutFee int32) *InterFlightOrderTicketInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) TicketPrice(ticketPrice int32) *InterFlightOrderTicketInfoBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) ConstructionFee(constructionFee int32) *InterFlightOrderTicketInfoBuilder {
	builder.constructionFee = constructionFee
	builder.constructionFeeSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) OilFee(oilFee int32) *InterFlightOrderTicketInfoBuilder {
	builder.oilFee = oilFee
	builder.oilFeeSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) ServiceFee(serviceFee int32) *InterFlightOrderTicketInfoBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) AsynchronousServiceFee(asynchronousServiceFee int32) *InterFlightOrderTicketInfoBuilder {
	builder.asynchronousServiceFee = asynchronousServiceFee
	builder.asynchronousServiceFeeSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) SequenceList(sequenceList []FlightOrderSequenceListItem) *InterFlightOrderTicketInfoBuilder {
	builder.sequenceList = sequenceList
	builder.sequenceListSet = true
	return builder
}
func (builder *InterFlightOrderTicketInfoBuilder) Tax(tax int32) *InterFlightOrderTicketInfoBuilder {
	builder.tax = tax
	builder.taxSet = true
	return builder
}

func (builder *InterFlightOrderTicketInfoBuilder) Build() *InterFlightOrderTicketInfo {
	data := &InterFlightOrderTicketInfo{}
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
	if builder.taxSet {
		data.Tax = &builder.tax
	}
	return data
}
