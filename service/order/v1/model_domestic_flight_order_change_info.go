package v1

// DomesticFlightOrderChangeInfo 改签信息(国内)
type DomesticFlightOrderChangeInfo struct {
	TicketUniqueKey              *string                               `json:"ticket_unique_key,omitempty"`          // 票号唯一值，展示成功的改签后的票，改签所有的过程数据，包含0元改签
	OriginalTicketUniqueKey      *string                               `json:"original_ticket_unique_key,omitempty"` // 改签原票唯一值，原票，多次改签也是对应预订的原票
	PreTicketUniqueKey           *string                               `json:"pre_ticket_unique_key,omitempty"`      // 改签前一次唯一值
	ChangedticketNumber          *string                               `json:"changedticket_number,omitempty"`       // 改签后票号，廉价航空直接给下游的票号，可以为空
	ChangedticketStatus          *string                               `json:"changedticket_status,omitempty"`       // 改签后票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	ChangeOrderId                *string                               `json:"change_order_id,omitempty"`            // 改签处理订单号，滴滴内部的改签订单号,预算推送时关联
	PassengerTravelerId          *string                               `json:"passenger_traveler_id,omitempty"`      // 乘客唯一值
	ReasonForChange              *string                               `json:"reason_for_change,omitempty"`          // 改签原因
	ChangedTime                  *string                               `json:"changed_time,omitempty"`               // 改签成功时间，格式：yyyy-MM-dd HH:mm:ss
	ChangHandleFee               *string                               `json:"chang_handle_fee,omitempty"`           // 改签手续费，单位：分
	ChangeUpCabinFee             *string                               `json:"change_up_cabin_fee,omitempty"`        // 改签升舱费，单位：分
	ChangeFee                    *string                               `json:"change_fee,omitempty"`                 // 改签总费用，单位：分 change_fee=change_up_cabin_fee+chang_handle_fee
	CutFee                       *int32                                `json:"cut_fee,omitempty"`                    // 机票订立减优惠，单位：分 改签时回收
	ChangeCompanyPay             *int32                                `json:"change_company_pay,omitempty"`         // 改签企业支付金额，单位：分 改签时公司支付部分，包含平台使用费
	ChangePersonalPay            *int32                                `json:"change_personal_pay,omitempty"`        // 改签员工支付金额，单位：分 改签时员工支付部分 明细参考personal_pay_detail
	PersonalPayDetail            *DomesticFlightOrderPersonalPayDetail `json:"personal_pay_detail,omitempty"`
	ChangeServiceFee             *string                               `json:"change_service_fee,omitempty"`              // 改签出票平台使用费，单位：分 实时收取的服务费
	ChangeAsynchronousServiceFee *string                               `json:"change_asynchronous_service_fee,omitempty"` // 改签出票异步平台使用费，单位：分 异步收取的服务费
	SequenceList                 []FlightOrderSequenceListItem         `json:"sequence_list,omitempty"`                   // 航段信息
}

type DomesticFlightOrderChangeInfoBuilder struct {
	ticketUniqueKey                 string // 票号唯一值，展示成功的改签后的票，改签所有的过程数据，包含0元改签
	ticketUniqueKeySet              bool
	originalTicketUniqueKey         string // 改签原票唯一值，原票，多次改签也是对应预订的原票
	originalTicketUniqueKeySet      bool
	preTicketUniqueKey              string // 改签前一次唯一值
	preTicketUniqueKeySet           bool
	changedticketNumber             string // 改签后票号，廉价航空直接给下游的票号，可以为空
	changedticketNumberSet          bool
	changedticketStatus             string // 改签后票号状态，枚举英文code：OpenForUse 客票有效未使用,CheckedIn 已办理值机手续,Boarded     已登机,Used     已使用,Refunded 已退票,Exchanged 电子客票已换开为其他客票,Suspended 系统处理中或认为挂起禁止使用,Void      已作废,Airport CNTL 机场控制,CPN Note 机场控制,UnKnow      未知，起飞后票号才开始扫描获取最近的票号状态
	changedticketStatusSet          bool
	changeOrderId                   string // 改签处理订单号，滴滴内部的改签订单号,预算推送时关联
	changeOrderIdSet                bool
	passengerTravelerId             string // 乘客唯一值
	passengerTravelerIdSet          bool
	reasonForChange                 string // 改签原因
	reasonForChangeSet              bool
	changedTime                     string // 改签成功时间，格式：yyyy-MM-dd HH:mm:ss
	changedTimeSet                  bool
	changHandleFee                  string // 改签手续费，单位：分
	changHandleFeeSet               bool
	changeUpCabinFee                string // 改签升舱费，单位：分
	changeUpCabinFeeSet             bool
	changeFee                       string // 改签总费用，单位：分 change_fee=change_up_cabin_fee+chang_handle_fee
	changeFeeSet                    bool
	cutFee                          int32 // 机票订立减优惠，单位：分 改签时回收
	cutFeeSet                       bool
	changeCompanyPay                int32 // 改签企业支付金额，单位：分 改签时公司支付部分，包含平台使用费
	changeCompanyPaySet             bool
	changePersonalPay               int32 // 改签员工支付金额，单位：分 改签时员工支付部分 明细参考personal_pay_detail
	changePersonalPaySet            bool
	personalPayDetail               DomesticFlightOrderPersonalPayDetail
	personalPayDetailSet            bool
	changeServiceFee                string // 改签出票平台使用费，单位：分 实时收取的服务费
	changeServiceFeeSet             bool
	changeAsynchronousServiceFee    string // 改签出票异步平台使用费，单位：分 异步收取的服务费
	changeAsynchronousServiceFeeSet bool
	sequenceList                    []FlightOrderSequenceListItem // 航段信息
	sequenceListSet                 bool
}

func NewDomesticFlightOrderChangeInfoBuilder() *DomesticFlightOrderChangeInfoBuilder {
	return &DomesticFlightOrderChangeInfoBuilder{}
}
func (builder *DomesticFlightOrderChangeInfoBuilder) TicketUniqueKey(ticketUniqueKey string) *DomesticFlightOrderChangeInfoBuilder {
	builder.ticketUniqueKey = ticketUniqueKey
	builder.ticketUniqueKeySet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) OriginalTicketUniqueKey(originalTicketUniqueKey string) *DomesticFlightOrderChangeInfoBuilder {
	builder.originalTicketUniqueKey = originalTicketUniqueKey
	builder.originalTicketUniqueKeySet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) PreTicketUniqueKey(preTicketUniqueKey string) *DomesticFlightOrderChangeInfoBuilder {
	builder.preTicketUniqueKey = preTicketUniqueKey
	builder.preTicketUniqueKeySet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangedticketNumber(changedticketNumber string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changedticketNumber = changedticketNumber
	builder.changedticketNumberSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangedticketStatus(changedticketStatus string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changedticketStatus = changedticketStatus
	builder.changedticketStatusSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeOrderId(changeOrderId string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeOrderId = changeOrderId
	builder.changeOrderIdSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) PassengerTravelerId(passengerTravelerId string) *DomesticFlightOrderChangeInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ReasonForChange(reasonForChange string) *DomesticFlightOrderChangeInfoBuilder {
	builder.reasonForChange = reasonForChange
	builder.reasonForChangeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangedTime(changedTime string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changedTime = changedTime
	builder.changedTimeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangHandleFee(changHandleFee string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changHandleFee = changHandleFee
	builder.changHandleFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeUpCabinFee(changeUpCabinFee string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeUpCabinFee = changeUpCabinFee
	builder.changeUpCabinFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeFee(changeFee string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeFee = changeFee
	builder.changeFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) CutFee(cutFee int32) *DomesticFlightOrderChangeInfoBuilder {
	builder.cutFee = cutFee
	builder.cutFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeCompanyPay(changeCompanyPay int32) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeCompanyPay = changeCompanyPay
	builder.changeCompanyPaySet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangePersonalPay(changePersonalPay int32) *DomesticFlightOrderChangeInfoBuilder {
	builder.changePersonalPay = changePersonalPay
	builder.changePersonalPaySet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) PersonalPayDetail(personalPayDetail DomesticFlightOrderPersonalPayDetail) *DomesticFlightOrderChangeInfoBuilder {
	builder.personalPayDetail = personalPayDetail
	builder.personalPayDetailSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeServiceFee(changeServiceFee string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeServiceFee = changeServiceFee
	builder.changeServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) ChangeAsynchronousServiceFee(changeAsynchronousServiceFee string) *DomesticFlightOrderChangeInfoBuilder {
	builder.changeAsynchronousServiceFee = changeAsynchronousServiceFee
	builder.changeAsynchronousServiceFeeSet = true
	return builder
}
func (builder *DomesticFlightOrderChangeInfoBuilder) SequenceList(sequenceList []FlightOrderSequenceListItem) *DomesticFlightOrderChangeInfoBuilder {
	builder.sequenceList = sequenceList
	builder.sequenceListSet = true
	return builder
}

func (builder *DomesticFlightOrderChangeInfoBuilder) Build() *DomesticFlightOrderChangeInfo {
	data := &DomesticFlightOrderChangeInfo{}
	if builder.ticketUniqueKeySet {
		data.TicketUniqueKey = &builder.ticketUniqueKey
	}
	if builder.originalTicketUniqueKeySet {
		data.OriginalTicketUniqueKey = &builder.originalTicketUniqueKey
	}
	if builder.preTicketUniqueKeySet {
		data.PreTicketUniqueKey = &builder.preTicketUniqueKey
	}
	if builder.changedticketNumberSet {
		data.ChangedticketNumber = &builder.changedticketNumber
	}
	if builder.changedticketStatusSet {
		data.ChangedticketStatus = &builder.changedticketStatus
	}
	if builder.changeOrderIdSet {
		data.ChangeOrderId = &builder.changeOrderId
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
	}
	if builder.reasonForChangeSet {
		data.ReasonForChange = &builder.reasonForChange
	}
	if builder.changedTimeSet {
		data.ChangedTime = &builder.changedTime
	}
	if builder.changHandleFeeSet {
		data.ChangHandleFee = &builder.changHandleFee
	}
	if builder.changeUpCabinFeeSet {
		data.ChangeUpCabinFee = &builder.changeUpCabinFee
	}
	if builder.changeFeeSet {
		data.ChangeFee = &builder.changeFee
	}
	if builder.cutFeeSet {
		data.CutFee = &builder.cutFee
	}
	if builder.changeCompanyPaySet {
		data.ChangeCompanyPay = &builder.changeCompanyPay
	}
	if builder.changePersonalPaySet {
		data.ChangePersonalPay = &builder.changePersonalPay
	}
	if builder.personalPayDetailSet {
		data.PersonalPayDetail = &builder.personalPayDetail
	}
	if builder.changeServiceFeeSet {
		data.ChangeServiceFee = &builder.changeServiceFee
	}
	if builder.changeAsynchronousServiceFeeSet {
		data.ChangeAsynchronousServiceFee = &builder.changeAsynchronousServiceFee
	}
	if builder.sequenceListSet {
		data.SequenceList = builder.sequenceList
	}
	return data
}
