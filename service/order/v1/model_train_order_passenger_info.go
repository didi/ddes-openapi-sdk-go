package v1

// TrainOrderPassengerInfo 定义乘客信息的消息
type TrainOrderPassengerInfo struct {
	PassengerName       *string                `json:"passenger_name,omitempty"`        // 乘客姓名
	PassengerPhone      *string                `json:"passenger_phone,omitempty"`       // 乘客手机号，掩码形式呈现
	PassengerTravelerId *string                `json:"passenger_traveler_id,omitempty"` // 乘客唯一值，若为内部员工则同memberId
	PassengerEmployeeId *string                `json:"passenger_employee_id,omitempty"` // 乘客员工编号
	IsTraveler          *int32                 `json:"is_traveler,omitempty"`           // 外部员工标记，1表示外部员工，0表示内部员工
	BudgetCenterName    *string                `json:"budget_center_name,omitempty"`    // 成本中心名称，建议使用budget_center_list字段替代
	BudgetCenterCode    *string                `json:"budget_center_code,omitempty"`    // 成本中心code，外部项目或部门code，建议使用budget_center_list字段替代
	BudgetCenterId      *string                `json:"budget_center_id,omitempty"`      // 成本中心id，滴滴内部主键，建议使用budget_center_list字段替代
	BudgetCenterList    []BudgetCenterListItem `json:"budget_center_list,omitempty"`    // 多成本中心信息，参考BudgetCenter对象
}

type TrainOrderPassengerInfoBuilder struct {
	passengerName          string // 乘客姓名
	passengerNameSet       bool
	passengerPhone         string // 乘客手机号，掩码形式呈现
	passengerPhoneSet      bool
	passengerTravelerId    string // 乘客唯一值，若为内部员工则同memberId
	passengerTravelerIdSet bool
	passengerEmployeeId    string // 乘客员工编号
	passengerEmployeeIdSet bool
	isTraveler             int32 // 外部员工标记，1表示外部员工，0表示内部员工
	isTravelerSet          bool
	budgetCenterName       string // 成本中心名称，建议使用budget_center_list字段替代
	budgetCenterNameSet    bool
	budgetCenterCode       string // 成本中心code，外部项目或部门code，建议使用budget_center_list字段替代
	budgetCenterCodeSet    bool
	budgetCenterId         string // 成本中心id，滴滴内部主键，建议使用budget_center_list字段替代
	budgetCenterIdSet      bool
	budgetCenterList       []BudgetCenterListItem // 多成本中心信息，参考BudgetCenter对象
	budgetCenterListSet    bool
}

func NewTrainOrderPassengerInfoBuilder() *TrainOrderPassengerInfoBuilder {
	return &TrainOrderPassengerInfoBuilder{}
}
func (builder *TrainOrderPassengerInfoBuilder) PassengerName(passengerName string) *TrainOrderPassengerInfoBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) PassengerPhone(passengerPhone string) *TrainOrderPassengerInfoBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) PassengerTravelerId(passengerTravelerId string) *TrainOrderPassengerInfoBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) PassengerEmployeeId(passengerEmployeeId string) *TrainOrderPassengerInfoBuilder {
	builder.passengerEmployeeId = passengerEmployeeId
	builder.passengerEmployeeIdSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) IsTraveler(isTraveler int32) *TrainOrderPassengerInfoBuilder {
	builder.isTraveler = isTraveler
	builder.isTravelerSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) BudgetCenterName(budgetCenterName string) *TrainOrderPassengerInfoBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) BudgetCenterCode(budgetCenterCode string) *TrainOrderPassengerInfoBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) BudgetCenterId(budgetCenterId string) *TrainOrderPassengerInfoBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *TrainOrderPassengerInfoBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *TrainOrderPassengerInfoBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *TrainOrderPassengerInfoBuilder) Build() *TrainOrderPassengerInfo {
	data := &TrainOrderPassengerInfo{}
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
