package v1

// FlightOrderPassengerInfo 出行人信息
type FlightOrderPassengerInfo struct {
	PassengerName       *string                `json:"passenger_name,omitempty"`        // 乘客姓名
	PassengerPhone      *string                `json:"passenger_phone,omitempty"`       // 乘客手机号
	PassengerTravelerId *string                `json:"passenger_traveler_id,omitempty"` // 乘客唯一值，内部员工同memberId
	PassengerEmployeeId *string                `json:"passenger_employee_id,omitempty"` // 乘客员工编号
	IsTraveler          *int32                 `json:"is_traveler,omitempty"`           // 外部员工标记，枚举值数字：1 外部员工 0 内部员工
	BudgetCenterName    *string                `json:"budget_center_name,omitempty"`    // 成本中心名称，建议使用budget_center_list
	BudgetCenterCode    *string                `json:"budget_center_code,omitempty"`    // 成本中心code，外部项目或部门code，建议使用budget_center_list
	BudgetCenterId      *string                `json:"budget_center_id,omitempty"`      // 成本中心id，滴滴内部主键，建议使用budget_center_list
	BudgetCenterList    []BudgetCenterListItem `json:"budget_center_list,omitempty"`    // 多成本中心，参考budget_center_list对象
}

type FlightOrderPassengerInfoBuilder struct {
	passengerName          string // 乘客姓名
	passengerNameSet       bool
	passengerPhone         string // 乘客手机号
	passengerPhoneSet      bool
	passengerTravelerId    string // 乘客唯一值，内部员工同memberId
	passengerTravelerIdSet bool
	passengerEmployeeId    string // 乘客员工编号
	passengerEmployeeIdSet bool
	isTraveler             int32 // 外部员工标记，枚举值数字：1 外部员工 0 内部员工
	isTravelerSet          bool
	budgetCenterName       string // 成本中心名称，建议使用budget_center_list
	budgetCenterNameSet    bool
	budgetCenterCode       string // 成本中心code，外部项目或部门code，建议使用budget_center_list
	budgetCenterCodeSet    bool
	budgetCenterId         string // 成本中心id，滴滴内部主键，建议使用budget_center_list
	budgetCenterIdSet      bool
	budgetCenterList       []BudgetCenterListItem // 多成本中心，参考budget_center_list对象
	budgetCenterListSet    bool
}

func NewFlightOrderPassengerInfoBuilder() *FlightOrderPassengerInfoBuilder {
	return &FlightOrderPassengerInfoBuilder{}
}
func (builder *FlightOrderPassengerInfoBuilder) PassengerName(passengerName string) *FlightOrderPassengerInfoBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) PassengerPhone(passengerPhone string) *FlightOrderPassengerInfoBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) PassengerTravelerId(passengerTravelerId string) *FlightOrderPassengerInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) PassengerEmployeeId(passengerEmployeeId string) *FlightOrderPassengerInfoBuilder {
	builder.passengerEmployeeId = passengerEmployeeId
	builder.passengerEmployeeIdSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) IsTraveler(isTraveler int32) *FlightOrderPassengerInfoBuilder {
	builder.isTraveler = isTraveler
	builder.isTravelerSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) BudgetCenterName(budgetCenterName string) *FlightOrderPassengerInfoBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) BudgetCenterCode(budgetCenterCode string) *FlightOrderPassengerInfoBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) BudgetCenterId(budgetCenterId string) *FlightOrderPassengerInfoBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *FlightOrderPassengerInfoBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *FlightOrderPassengerInfoBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *FlightOrderPassengerInfoBuilder) Build() *FlightOrderPassengerInfo {
	data := &FlightOrderPassengerInfo{}
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
	if builder.isTravelerSet {
		data.IsTraveler = &builder.isTraveler
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	return data
}
