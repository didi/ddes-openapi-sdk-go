package v1

// TripPassenger 出行人信息
type TripPassenger struct {
	PassengerType    *int32                 `json:"passenger_type,omitempty"`     // 出行人类型，枚举值数字：0：员工；1：外部出行人
	BudgetCenterList []BudgetCenterListItem `json:"budget_center_list,omitempty"` // 多成本中心
	PassengerName    *string                `json:"passenger_name,omitempty"`     // 出行人姓名/员工姓名
	MemberType       *int32                 `json:"member_type,omitempty"`        // 出行人类型，0：手机号，passenger_phone 1：工号；employee_number 必填 2：邮箱：email 必填 默认为0 以 member_type 对应的值为准，其他字段传了不生效
	PassengerPhone   *string                `json:"passenger_phone,omitempty"`    // 员工手机号，举例：15100000000
	EmployeeNumber   *string                `json:"employee_number,omitempty"`    // 员工工号
	Email            *string                `json:"email,omitempty"`              // 员工邮箱
	TravelerId       *string                `json:"traveler_id,omitempty"`        // 外部出行人 ID
	// Deprecated
	PolicyType *int32 `json:"policy_type,omitempty"` // 执行政策指定项类型，1：out_rank_id 外部职级编号 2：employee_number 员工工号 3：employee_phone 员工手机号 4：employee_email 员工邮箱 5：regulation_id 制度I 指定生效policy_type_value内对应的值
	// Deprecated
	PolicyTypeValue *string `json:"policy_type_value,omitempty"` // 执行政策指定项
}

type TripPassengerBuilder struct {
	passengerType       int32 // 出行人类型，枚举值数字：0：员工；1：外部出行人
	passengerTypeSet    bool
	budgetCenterList    []BudgetCenterListItem // 多成本中心
	budgetCenterListSet bool
	passengerName       string // 出行人姓名/员工姓名
	passengerNameSet    bool
	memberType          int32 // 出行人类型，0：手机号，passenger_phone 1：工号；employee_number 必填 2：邮箱：email 必填 默认为0 以 member_type 对应的值为准，其他字段传了不生效
	memberTypeSet       bool
	passengerPhone      string // 员工手机号，举例：15100000000
	passengerPhoneSet   bool
	employeeNumber      string // 员工工号
	employeeNumberSet   bool
	email               string // 员工邮箱
	emailSet            bool
	travelerId          string // 外部出行人 ID
	travelerIdSet       bool
	// Deprecated
	policyType    int32 // 执行政策指定项类型，1：out_rank_id 外部职级编号 2：employee_number 员工工号 3：employee_phone 员工手机号 4：employee_email 员工邮箱 5：regulation_id 制度I 指定生效policy_type_value内对应的值
	policyTypeSet bool
	// Deprecated
	policyTypeValue    string // 执行政策指定项
	policyTypeValueSet bool
}

func NewTripPassengerBuilder() *TripPassengerBuilder {
	return &TripPassengerBuilder{}
}
func (builder *TripPassengerBuilder) PassengerType(passengerType int32) *TripPassengerBuilder {
	builder.passengerType = passengerType
	builder.passengerTypeSet = true
	return builder
}
func (builder *TripPassengerBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *TripPassengerBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *TripPassengerBuilder) PassengerName(passengerName string) *TripPassengerBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *TripPassengerBuilder) MemberType(memberType int32) *TripPassengerBuilder {
	builder.memberType = memberType
	builder.memberTypeSet = true
	return builder
}
func (builder *TripPassengerBuilder) PassengerPhone(passengerPhone string) *TripPassengerBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *TripPassengerBuilder) EmployeeNumber(employeeNumber string) *TripPassengerBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *TripPassengerBuilder) Email(email string) *TripPassengerBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *TripPassengerBuilder) TravelerId(travelerId string) *TripPassengerBuilder {
	builder.travelerId = travelerId
	builder.travelerIdSet = true
	return builder
}

// Deprecated
func (builder *TripPassengerBuilder) PolicyType(policyType int32) *TripPassengerBuilder {
	builder.policyType = policyType
	builder.policyTypeSet = true
	return builder
}

// Deprecated
func (builder *TripPassengerBuilder) PolicyTypeValue(policyTypeValue string) *TripPassengerBuilder {
	builder.policyTypeValue = policyTypeValue
	builder.policyTypeValueSet = true
	return builder
}

func (builder *TripPassengerBuilder) Build() *TripPassenger {
	data := &TripPassenger{}
	if builder.passengerTypeSet {
		data.PassengerType = &builder.passengerType
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.memberTypeSet {
		data.MemberType = &builder.memberType
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.travelerIdSet {
		data.TravelerId = &builder.travelerId
	}
	if builder.policyTypeSet {
		data.PolicyType = &builder.policyType
	}
	if builder.policyTypeValueSet {
		data.PolicyTypeValue = &builder.policyTypeValue
	}
	return data
}
