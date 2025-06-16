package v1

// PassengerInfo struct for PassengerInfo
type PassengerInfo struct {
	Name             *string                `json:"name,omitempty"`
	Phone            *string                `json:"phone,omitempty"`
	PassengerType    *string                `json:"passenger_type,omitempty"`
	EmployeeNumber   *string                `json:"employee_number,omitempty"`
	Email            *string                `json:"email,omitempty"`
	BudgetCenterList []BudgetCenterListItem `json:"budget_center_list,omitempty"`
}

type PassengerInfoBuilder struct {
	name                string
	nameSet             bool
	phone               string
	phoneSet            bool
	passengerType       string
	passengerTypeSet    bool
	employeeNumber      string
	employeeNumberSet   bool
	email               string
	emailSet            bool
	budgetCenterList    []BudgetCenterListItem
	budgetCenterListSet bool
}

func NewPassengerInfoBuilder() *PassengerInfoBuilder {
	return &PassengerInfoBuilder{}
}
func (builder *PassengerInfoBuilder) Name(name string) *PassengerInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *PassengerInfoBuilder) Phone(phone string) *PassengerInfoBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *PassengerInfoBuilder) PassengerType(passengerType string) *PassengerInfoBuilder {
	builder.passengerType = passengerType
	builder.passengerTypeSet = true
	return builder
}
func (builder *PassengerInfoBuilder) EmployeeNumber(employeeNumber string) *PassengerInfoBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *PassengerInfoBuilder) Email(email string) *PassengerInfoBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *PassengerInfoBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *PassengerInfoBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *PassengerInfoBuilder) Build() *PassengerInfo {
	data := &PassengerInfo{}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.passengerTypeSet {
		data.PassengerType = &builder.passengerType
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	return data
}
