package v1

// RelatedEmployeeItem struct for RelatedEmployeeItem
type RelatedEmployeeItem struct {
	RelatedEmployeeId *string `json:"related_employee_id,omitempty"` // 关联员工滴滴侧ID
	EmployeeNumber    *string `json:"employee_number,omitempty"`     // 员工工号
	Phone             *string `json:"phone,omitempty"`               // 员工手机号
	Email             *string `json:"email,omitempty"`               // 员工邮箱
}

type RelatedEmployeeItemBuilder struct {
	relatedEmployeeId    string // 关联员工滴滴侧ID
	relatedEmployeeIdSet bool
	employeeNumber       string // 员工工号
	employeeNumberSet    bool
	phone                string // 员工手机号
	phoneSet             bool
	email                string // 员工邮箱
	emailSet             bool
}

func NewRelatedEmployeeItemBuilder() *RelatedEmployeeItemBuilder {
	return &RelatedEmployeeItemBuilder{}
}
func (builder *RelatedEmployeeItemBuilder) RelatedEmployeeId(relatedEmployeeId string) *RelatedEmployeeItemBuilder {
	builder.relatedEmployeeId = relatedEmployeeId
	builder.relatedEmployeeIdSet = true
	return builder
}
func (builder *RelatedEmployeeItemBuilder) EmployeeNumber(employeeNumber string) *RelatedEmployeeItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *RelatedEmployeeItemBuilder) Phone(phone string) *RelatedEmployeeItemBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *RelatedEmployeeItemBuilder) Email(email string) *RelatedEmployeeItemBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}

func (builder *RelatedEmployeeItemBuilder) Build() *RelatedEmployeeItem {
	data := &RelatedEmployeeItem{}
	if builder.relatedEmployeeIdSet {
		data.RelatedEmployeeId = &builder.relatedEmployeeId
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	return data
}
