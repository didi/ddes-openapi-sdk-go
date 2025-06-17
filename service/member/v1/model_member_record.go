package v1

// MemberRecord struct for MemberRecord
type MemberRecord struct {
	Id                     *string    `json:"id,omitempty"`                       // 员工在滴滴企业平台的ID
	Phone                  *string    `json:"phone,omitempty"`                    // 员工手机号
	Realname               *string    `json:"realname,omitempty"`                 // 员工姓名
	EmployeeNumber         *string    `json:"employee_number,omitempty"`          // 员工工号，员工在公司的员工号
	Email                  *string    `json:"email,omitempty"`                    // 邮箱
	SystemRole             *int32     `json:"system_role,omitempty"`              // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	RoleIds                *string    `json:"role_ids,omitempty"`                 // 角色，可以通过角色API获取对应的ID
	ImmediateSuperiorPhone *string    `json:"immediate_superior_phone,omitempty"` // 员工直属上级的手机号码，直属上级可在审批流中担任审批人
	ImmediateSuperiorEid   *string    `json:"immediate_superior_eid,omitempty"`   // 员工直属上级的员工编号
	Residentsname          *string    `json:"residentsname,omitempty"`            // 常驻地中文
	UseCompanyMoney        *int32     `json:"use_company_money,omitempty"`        // 是否企业支付余额，枚举值数字 0 否，1 是
	TotalQuota             *string    `json:"total_quota,omitempty"`              // 每月配额，单位元
	IsRemark               *string    `json:"is_remark,omitempty"`                // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	BudgetCenterId         *string    `json:"budget_center_id,omitempty"`         // 所在部门ID
	ConDepartmentIds       []string   `json:"con_department_ids,omitempty"`       // 所在兼岗部门ID，员工返回兼岗的部门ID数组
	ProjectIds             *string    `json:"project_ids,omitempty"`              // 所在项目ID
	UseCarConfig           []string   `json:"use_car_config,omitempty"`           // 用车规则ID数组
	AvailableQuota         *string    `json:"available_quota,omitempty"`          // 员工可用限额，单位元 例如\"937.70\"
	BranchName             *string    `json:"branch_name,omitempty"`              // 所在分公司名称（老），后续此参数会去掉
	Department             *string    `json:"department,omitempty"`               // 部门名称（老），后续此参数会去掉
	RegulationId           []string   `json:"regulation_id,omitempty"`            // 用车制度ID数组，跟新增、修改、详情等API统一
	SetDismissTime         *string    `json:"set_dismiss_time,omitempty"`         // 设置的员工离职日期，为空时表示未设置离职日期，格式为 yyyy-MM-dd
	DismissTime            *string    `json:"dismiss_time,omitempty"`             // 员工实际离职日期，为空时表示未离职,格式为 格式：yyyy-MM-dd HH:mm:ss
	InvoiceInfo            *string    `json:"invoice_info,omitempty"`             // 开票主体信息
	LegalEntityId          *string    `json:"legal_entity_id,omitempty"`          // 所在公司主体id
	OutLegalEntityId       *string    `json:"out_legal_entity_id,omitempty"`      // 外部公司主体编号
	RankId                 *string    `json:"rank_id,omitempty"`                  // 职级id
	OutRankId              *string    `json:"out_rank_id,omitempty"`              // 外部职级 ID
	EnglishSurname         *string    `json:"english_surname,omitempty"`          // 英文姓
	EnglishName            *string    `json:"english_name,omitempty"`             // 英文名
	Nickname               *string    `json:"nickname,omitempty"`                 // 昵称
	Sex                    *int32     `json:"sex,omitempty"`                      // 性别，枚举值数字 0 未知 1 男 2 女
	BirthDate              *string    `json:"birth_date,omitempty"`               // 出生日期，格式2000-01-01（已用AES算法加密）
	CardList               []CardInfo `json:"card_list,omitempty"`                // 证件信息
	Source                 *string    `json:"source,omitempty"`                   // 员工加入滴滴企业平台的渠道，枚举值数字：0;未知 1;PC逐一添加 2;PC批量导入 3;邮件邀请
}

type MemberRecordBuilder struct {
	id                        string // 员工在滴滴企业平台的ID
	idSet                     bool
	phone                     string // 员工手机号
	phoneSet                  bool
	realname                  string // 员工姓名
	realnameSet               bool
	employeeNumber            string // 员工工号，员工在公司的员工号
	employeeNumberSet         bool
	email                     string // 邮箱
	emailSet                  bool
	systemRole                int32 // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	systemRoleSet             bool
	roleIds                   string // 角色，可以通过角色API获取对应的ID
	roleIdsSet                bool
	immediateSuperiorPhone    string // 员工直属上级的手机号码，直属上级可在审批流中担任审批人
	immediateSuperiorPhoneSet bool
	immediateSuperiorEid      string // 员工直属上级的员工编号
	immediateSuperiorEidSet   bool
	residentsname             string // 常驻地中文
	residentsnameSet          bool
	useCompanyMoney           int32 // 是否企业支付余额，枚举值数字 0 否，1 是
	useCompanyMoneySet        bool
	totalQuota                string // 每月配额，单位元
	totalQuotaSet             bool
	isRemark                  string // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	isRemarkSet               bool
	budgetCenterId            string // 所在部门ID
	budgetCenterIdSet         bool
	conDepartmentIds          []string // 所在兼岗部门ID，员工返回兼岗的部门ID数组
	conDepartmentIdsSet       bool
	projectIds                string // 所在项目ID
	projectIdsSet             bool
	useCarConfig              []string // 用车规则ID数组
	useCarConfigSet           bool
	availableQuota            string // 员工可用限额，单位元 例如\"937.70\"
	availableQuotaSet         bool
	branchName                string // 所在分公司名称（老），后续此参数会去掉
	branchNameSet             bool
	department                string // 部门名称（老），后续此参数会去掉
	departmentSet             bool
	regulationId              []string // 用车制度ID数组，跟新增、修改、详情等API统一
	regulationIdSet           bool
	setDismissTime            string // 设置的员工离职日期，为空时表示未设置离职日期，格式为 yyyy-MM-dd
	setDismissTimeSet         bool
	dismissTime               string // 员工实际离职日期，为空时表示未离职,格式为 格式：yyyy-MM-dd HH:mm:ss
	dismissTimeSet            bool
	invoiceInfo               string // 开票主体信息
	invoiceInfoSet            bool
	legalEntityId             string // 所在公司主体id
	legalEntityIdSet          bool
	outLegalEntityId          string // 外部公司主体编号
	outLegalEntityIdSet       bool
	rankId                    string // 职级id
	rankIdSet                 bool
	outRankId                 string // 外部职级 ID
	outRankIdSet              bool
	englishSurname            string // 英文姓
	englishSurnameSet         bool
	englishName               string // 英文名
	englishNameSet            bool
	nickname                  string // 昵称
	nicknameSet               bool
	sex                       int32 // 性别，枚举值数字 0 未知 1 男 2 女
	sexSet                    bool
	birthDate                 string // 出生日期，格式2000-01-01（已用AES算法加密）
	birthDateSet              bool
	cardList                  []CardInfo // 证件信息
	cardListSet               bool
	source                    string // 员工加入滴滴企业平台的渠道，枚举值数字：0;未知 1;PC逐一添加 2;PC批量导入 3;邮件邀请
	sourceSet                 bool
}

func NewMemberRecordBuilder() *MemberRecordBuilder {
	return &MemberRecordBuilder{}
}
func (builder *MemberRecordBuilder) Id(id string) *MemberRecordBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *MemberRecordBuilder) Phone(phone string) *MemberRecordBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *MemberRecordBuilder) Realname(realname string) *MemberRecordBuilder {
	builder.realname = realname
	builder.realnameSet = true
	return builder
}
func (builder *MemberRecordBuilder) EmployeeNumber(employeeNumber string) *MemberRecordBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *MemberRecordBuilder) Email(email string) *MemberRecordBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *MemberRecordBuilder) SystemRole(systemRole int32) *MemberRecordBuilder {
	builder.systemRole = systemRole
	builder.systemRoleSet = true
	return builder
}
func (builder *MemberRecordBuilder) RoleIds(roleIds string) *MemberRecordBuilder {
	builder.roleIds = roleIds
	builder.roleIdsSet = true
	return builder
}
func (builder *MemberRecordBuilder) ImmediateSuperiorPhone(immediateSuperiorPhone string) *MemberRecordBuilder {
	builder.immediateSuperiorPhone = immediateSuperiorPhone
	builder.immediateSuperiorPhoneSet = true
	return builder
}
func (builder *MemberRecordBuilder) ImmediateSuperiorEid(immediateSuperiorEid string) *MemberRecordBuilder {
	builder.immediateSuperiorEid = immediateSuperiorEid
	builder.immediateSuperiorEidSet = true
	return builder
}
func (builder *MemberRecordBuilder) Residentsname(residentsname string) *MemberRecordBuilder {
	builder.residentsname = residentsname
	builder.residentsnameSet = true
	return builder
}
func (builder *MemberRecordBuilder) UseCompanyMoney(useCompanyMoney int32) *MemberRecordBuilder {
	builder.useCompanyMoney = useCompanyMoney
	builder.useCompanyMoneySet = true
	return builder
}
func (builder *MemberRecordBuilder) TotalQuota(totalQuota string) *MemberRecordBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *MemberRecordBuilder) IsRemark(isRemark string) *MemberRecordBuilder {
	builder.isRemark = isRemark
	builder.isRemarkSet = true
	return builder
}
func (builder *MemberRecordBuilder) BudgetCenterId(budgetCenterId string) *MemberRecordBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) ConDepartmentIds(conDepartmentIds []string) *MemberRecordBuilder {
	builder.conDepartmentIds = conDepartmentIds
	builder.conDepartmentIdsSet = true
	return builder
}
func (builder *MemberRecordBuilder) ProjectIds(projectIds string) *MemberRecordBuilder {
	builder.projectIds = projectIds
	builder.projectIdsSet = true
	return builder
}
func (builder *MemberRecordBuilder) UseCarConfig(useCarConfig []string) *MemberRecordBuilder {
	builder.useCarConfig = useCarConfig
	builder.useCarConfigSet = true
	return builder
}
func (builder *MemberRecordBuilder) AvailableQuota(availableQuota string) *MemberRecordBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *MemberRecordBuilder) BranchName(branchName string) *MemberRecordBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *MemberRecordBuilder) Department(department string) *MemberRecordBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *MemberRecordBuilder) RegulationId(regulationId []string) *MemberRecordBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) SetDismissTime(setDismissTime string) *MemberRecordBuilder {
	builder.setDismissTime = setDismissTime
	builder.setDismissTimeSet = true
	return builder
}
func (builder *MemberRecordBuilder) DismissTime(dismissTime string) *MemberRecordBuilder {
	builder.dismissTime = dismissTime
	builder.dismissTimeSet = true
	return builder
}
func (builder *MemberRecordBuilder) InvoiceInfo(invoiceInfo string) *MemberRecordBuilder {
	builder.invoiceInfo = invoiceInfo
	builder.invoiceInfoSet = true
	return builder
}
func (builder *MemberRecordBuilder) LegalEntityId(legalEntityId string) *MemberRecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) OutLegalEntityId(outLegalEntityId string) *MemberRecordBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) RankId(rankId string) *MemberRecordBuilder {
	builder.rankId = rankId
	builder.rankIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) OutRankId(outRankId string) *MemberRecordBuilder {
	builder.outRankId = outRankId
	builder.outRankIdSet = true
	return builder
}
func (builder *MemberRecordBuilder) EnglishSurname(englishSurname string) *MemberRecordBuilder {
	builder.englishSurname = englishSurname
	builder.englishSurnameSet = true
	return builder
}
func (builder *MemberRecordBuilder) EnglishName(englishName string) *MemberRecordBuilder {
	builder.englishName = englishName
	builder.englishNameSet = true
	return builder
}
func (builder *MemberRecordBuilder) Nickname(nickname string) *MemberRecordBuilder {
	builder.nickname = nickname
	builder.nicknameSet = true
	return builder
}
func (builder *MemberRecordBuilder) Sex(sex int32) *MemberRecordBuilder {
	builder.sex = sex
	builder.sexSet = true
	return builder
}
func (builder *MemberRecordBuilder) BirthDate(birthDate string) *MemberRecordBuilder {
	builder.birthDate = birthDate
	builder.birthDateSet = true
	return builder
}
func (builder *MemberRecordBuilder) CardList(cardList []CardInfo) *MemberRecordBuilder {
	builder.cardList = cardList
	builder.cardListSet = true
	return builder
}
func (builder *MemberRecordBuilder) Source(source string) *MemberRecordBuilder {
	builder.source = source
	builder.sourceSet = true
	return builder
}

func (builder *MemberRecordBuilder) Build() *MemberRecord {
	data := &MemberRecord{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.realnameSet {
		data.Realname = &builder.realname
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.systemRoleSet {
		data.SystemRole = &builder.systemRole
	}
	if builder.roleIdsSet {
		data.RoleIds = &builder.roleIds
	}
	if builder.immediateSuperiorPhoneSet {
		data.ImmediateSuperiorPhone = &builder.immediateSuperiorPhone
	}
	if builder.immediateSuperiorEidSet {
		data.ImmediateSuperiorEid = &builder.immediateSuperiorEid
	}
	if builder.residentsnameSet {
		data.Residentsname = &builder.residentsname
	}
	if builder.useCompanyMoneySet {
		data.UseCompanyMoney = &builder.useCompanyMoney
	}
	if builder.totalQuotaSet {
		data.TotalQuota = &builder.totalQuota
	}
	if builder.isRemarkSet {
		data.IsRemark = &builder.isRemark
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.conDepartmentIdsSet {
		data.ConDepartmentIds = builder.conDepartmentIds
	}
	if builder.projectIdsSet {
		data.ProjectIds = &builder.projectIds
	}
	if builder.useCarConfigSet {
		data.UseCarConfig = builder.useCarConfig
	}
	if builder.availableQuotaSet {
		data.AvailableQuota = &builder.availableQuota
	}
	if builder.branchNameSet {
		data.BranchName = &builder.branchName
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.regulationIdSet {
		data.RegulationId = builder.regulationId
	}
	if builder.setDismissTimeSet {
		data.SetDismissTime = &builder.setDismissTime
	}
	if builder.dismissTimeSet {
		data.DismissTime = &builder.dismissTime
	}
	if builder.invoiceInfoSet {
		data.InvoiceInfo = &builder.invoiceInfo
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.rankIdSet {
		data.RankId = &builder.rankId
	}
	if builder.outRankIdSet {
		data.OutRankId = &builder.outRankId
	}
	if builder.englishSurnameSet {
		data.EnglishSurname = &builder.englishSurname
	}
	if builder.englishNameSet {
		data.EnglishName = &builder.englishName
	}
	if builder.nicknameSet {
		data.Nickname = &builder.nickname
	}
	if builder.sexSet {
		data.Sex = &builder.sex
	}
	if builder.birthDateSet {
		data.BirthDate = &builder.birthDate
	}
	if builder.cardListSet {
		data.CardList = builder.cardList
	}
	if builder.sourceSet {
		data.Source = &builder.source
	}
	return data
}
