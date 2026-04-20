package v1

// PassengerInfo 出行人对象
type PassengerInfo struct {
	PassengerName       *string    `json:"passenger_name,omitempty"`        // 乘客姓名
	PassengerPhone      *string    `json:"passenger_phone,omitempty"`       // 乘客手机号
	PassengerTravelerId *string    `json:"passenger_traveler_id,omitempty"` // 乘客唯一值
	PassengerEmployeeId *string    `json:"passenger_employee_id,omitempty"` // 乘客员工编号
	OutTravelerId       *string    `json:"out_traveler_id,omitempty"`       // 外部出行人唯一ID
	IsTraveler          *int32     `json:"is_traveler,omitempty"`           // 外部员工标记，枚举值数字：0 内部员工 1 外部员工
	CostShare           *CostShare `json:"cost_share,omitempty"`
}

type PassengerInfoBuilder struct {
	passengerName          string // 乘客姓名
	passengerNameSet       bool
	passengerPhone         string // 乘客手机号
	passengerPhoneSet      bool
	passengerTravelerId    string // 乘客唯一值
	passengerTravelerIdSet bool
	passengerEmployeeId    string // 乘客员工编号
	passengerEmployeeIdSet bool
	outTravelerId          string // 外部出行人唯一ID
	outTravelerIdSet       bool
	isTraveler             int32 // 外部员工标记，枚举值数字：0 内部员工 1 外部员工
	isTravelerSet          bool
	costShare              CostShare
	costShareSet           bool
}

func NewPassengerInfoBuilder() *PassengerInfoBuilder {
	return &PassengerInfoBuilder{}
}
func (builder *PassengerInfoBuilder) PassengerName(passengerName string) *PassengerInfoBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *PassengerInfoBuilder) PassengerPhone(passengerPhone string) *PassengerInfoBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *PassengerInfoBuilder) PassengerTravelerId(passengerTravelerId string) *PassengerInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *PassengerInfoBuilder) PassengerEmployeeId(passengerEmployeeId string) *PassengerInfoBuilder {
	builder.passengerEmployeeId = passengerEmployeeId
	builder.passengerEmployeeIdSet = true
	return builder
}
func (builder *PassengerInfoBuilder) OutTravelerId(outTravelerId string) *PassengerInfoBuilder {
	builder.outTravelerId = outTravelerId
	builder.outTravelerIdSet = true
	return builder
}
func (builder *PassengerInfoBuilder) IsTraveler(isTraveler int32) *PassengerInfoBuilder {
	builder.isTraveler = isTraveler
	builder.isTravelerSet = true
	return builder
}
func (builder *PassengerInfoBuilder) CostShare(costShare CostShare) *PassengerInfoBuilder {
	builder.costShare = costShare
	builder.costShareSet = true
	return builder
}

func (builder *PassengerInfoBuilder) Build() *PassengerInfo {
	data := &PassengerInfo{}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
	}
	if builder.passengerTravelerIdSet {
		data.PassengerTravelerId = &builder.passengerTravelerId
	}
	if builder.passengerEmployeeIdSet {
		data.PassengerEmployeeId = &builder.passengerEmployeeId
	}
	if builder.outTravelerIdSet {
		data.OutTravelerId = &builder.outTravelerId
	}
	if builder.isTravelerSet {
		data.IsTraveler = &builder.isTraveler
	}
	if builder.costShareSet {
		data.CostShare = &builder.costShare
	}
	return data
}
