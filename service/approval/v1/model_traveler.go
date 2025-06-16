package v1

// Traveler 出行人信息
type Traveler struct {
	Name             *string                `json:"name,omitempty"`
	Phone            *string                `json:"phone,omitempty"`
	PassengerType    *string                `json:"passenger_type,omitempty"`
	PassengerNumber  *string                `json:"passenger_number,omitempty"`
	Email            *string                `json:"email,omitempty"`
	BudgetCenterList []BudgetCenterListItem `json:"budget_center_list,omitempty"`
}

type TravelerBuilder struct {
	name                string
	nameSet             bool
	phone               string
	phoneSet            bool
	passengerType       string
	passengerTypeSet    bool
	passengerNumber     string
	passengerNumberSet  bool
	email               string
	emailSet            bool
	budgetCenterList    []BudgetCenterListItem
	budgetCenterListSet bool
}

func NewTravelerBuilder() *TravelerBuilder {
	return &TravelerBuilder{}
}
func (builder *TravelerBuilder) Name(name string) *TravelerBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *TravelerBuilder) Phone(phone string) *TravelerBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *TravelerBuilder) PassengerType(passengerType string) *TravelerBuilder {
	builder.passengerType = passengerType
	builder.passengerTypeSet = true
	return builder
}
func (builder *TravelerBuilder) PassengerNumber(passengerNumber string) *TravelerBuilder {
	builder.passengerNumber = passengerNumber
	builder.passengerNumberSet = true
	return builder
}
func (builder *TravelerBuilder) Email(email string) *TravelerBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *TravelerBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *TravelerBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *TravelerBuilder) Build() *Traveler {
	data := &Traveler{}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.passengerTypeSet {
		data.PassengerType = &builder.passengerType
	}
	if builder.passengerNumberSet {
		data.PassengerNumber = &builder.passengerNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	return data
}
