package v1

// ProjectDetailMemberInfo 项目人员关联信息
type ProjectDetailMemberInfo struct {
	MemberId       *string `json:"member_id,omitempty"`       // 员工在滴滴的id
	Phone          *string `json:"phone,omitempty"`           // 员工手机号
	EmployeeNumber *string `json:"employee_number,omitempty"` // 员工工号
	Email          *string `json:"email,omitempty"`           // 员工邮箱
}

type ProjectDetailMemberInfoBuilder struct {
	memberId          string
	memberIdSet       bool
	phone             string
	phoneSet          bool
	employeeNumber    string
	employeeNumberSet bool
	email             string
	emailSet          bool
}

func NewProjectDetailMemberInfoBuilder() *ProjectDetailMemberInfoBuilder {
	return &ProjectDetailMemberInfoBuilder{}
}
func (builder *ProjectDetailMemberInfoBuilder) MemberId(memberId string) *ProjectDetailMemberInfoBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *ProjectDetailMemberInfoBuilder) Phone(phone string) *ProjectDetailMemberInfoBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *ProjectDetailMemberInfoBuilder) EmployeeNumber(employeeNumber string) *ProjectDetailMemberInfoBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *ProjectDetailMemberInfoBuilder) Email(email string) *ProjectDetailMemberInfoBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}

func (builder *ProjectDetailMemberInfoBuilder) Build() *ProjectDetailMemberInfo {
	data := &ProjectDetailMemberInfo{}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	return data
}
