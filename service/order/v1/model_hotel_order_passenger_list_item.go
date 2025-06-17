package v1

// HotelOrderPassengerListItem 出行人信息
type HotelOrderPassengerListItem struct {
	PassengerName             *string                `json:"passenger_name,omitempty"`        // 入住人姓名
	PassengerPhone            *string                `json:"passenger_phone,omitempty"`       // 入住人手机号
	PassengerTravelerId       *string                `json:"passenger_traveler_id,omitempty"` // 入住人唯一值，内部等于memberId
	PassengerEmployeeId       *string                `json:"passenger_employee_id,omitempty"` // 入住人员工编号
	IsTraveler                *int32                 `json:"is_traveler,omitempty"`           // 外部员工标记，枚举值数字：1 外部员工 0 内部员工
	BudgetCenterName          *string                `json:"budget_center_name,omitempty"`    // 成本中心名称，建议使用budget_center_list
	BudgetCenterCode          *string                `json:"budget_center_code,omitempty"`    // 成本中心code，外部项目或部门code，建议使用budget_center_list
	BudgetCenterId            *string                `json:"budget_center_id,omitempty"`      // 成本中心id，滴滴内部主键，建议使用budget_center_list
	BudgetCenterList          []BudgetCenterListItem `json:"budget_center_list,omitempty"`    // 多成本中心，参考budget_center_list对象
	PassengerPhoneCountryCode *string                `json:"passenger_phone_country_code,omitempty"`
}

type HotelOrderPassengerListItemBuilder struct {
	passengerName                string // 入住人姓名
	passengerNameSet             bool
	passengerPhone               string // 入住人手机号
	passengerPhoneSet            bool
	passengerTravelerId          string // 入住人唯一值，内部等于memberId
	passengerTravelerIdSet       bool
	passengerEmployeeId          string // 入住人员工编号
	passengerEmployeeIdSet       bool
	isTraveler                   int32 // 外部员工标记，枚举值数字：1 外部员工 0 内部员工
	isTravelerSet                bool
	budgetCenterName             string // 成本中心名称，建议使用budget_center_list
	budgetCenterNameSet          bool
	budgetCenterCode             string // 成本中心code，外部项目或部门code，建议使用budget_center_list
	budgetCenterCodeSet          bool
	budgetCenterId               string // 成本中心id，滴滴内部主键，建议使用budget_center_list
	budgetCenterIdSet            bool
	budgetCenterList             []BudgetCenterListItem // 多成本中心，参考budget_center_list对象
	budgetCenterListSet          bool
	passengerPhoneCountryCode    string
	passengerPhoneCountryCodeSet bool
}

func NewHotelOrderPassengerListItemBuilder() *HotelOrderPassengerListItemBuilder {
	return &HotelOrderPassengerListItemBuilder{}
}
func (builder *HotelOrderPassengerListItemBuilder) PassengerName(passengerName string) *HotelOrderPassengerListItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) PassengerPhone(passengerPhone string) *HotelOrderPassengerListItemBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) PassengerTravelerId(passengerTravelerId string) *HotelOrderPassengerListItemBuilder {
	builder.passengerTravelerId = passengerTravelerId
	builder.passengerTravelerIdSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) PassengerEmployeeId(passengerEmployeeId string) *HotelOrderPassengerListItemBuilder {
	builder.passengerEmployeeId = passengerEmployeeId
	builder.passengerEmployeeIdSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) IsTraveler(isTraveler int32) *HotelOrderPassengerListItemBuilder {
	builder.isTraveler = isTraveler
	builder.isTravelerSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) BudgetCenterName(budgetCenterName string) *HotelOrderPassengerListItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) BudgetCenterCode(budgetCenterCode string) *HotelOrderPassengerListItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) BudgetCenterId(budgetCenterId string) *HotelOrderPassengerListItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *HotelOrderPassengerListItemBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *HotelOrderPassengerListItemBuilder) PassengerPhoneCountryCode(passengerPhoneCountryCode string) *HotelOrderPassengerListItemBuilder {
	builder.passengerPhoneCountryCode = passengerPhoneCountryCode
	builder.passengerPhoneCountryCodeSet = true
	return builder
}

func (builder *HotelOrderPassengerListItemBuilder) Build() *HotelOrderPassengerListItem {
	data := &HotelOrderPassengerListItem{}
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
	if builder.passengerPhoneCountryCodeSet {
		data.PassengerPhoneCountryCode = &builder.passengerPhoneCountryCode
	}
	return data
}
