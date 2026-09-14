package v1

// MemberV2Record 员工记录（/open-apis/v2/member/get 响应 records 元素，契约依据 idl/member.idl MemberInfo）
type MemberV2Record struct {
	Id                     *string                     `json:"id,omitempty"`                       // 员工在滴滴企业平台的ID
	Phone                  *string                     `json:"phone,omitempty"`                    // 员工手机号
	Realname               *string                     `json:"realname,omitempty"`                 // 员工姓名
	EmployeeNumber         *string                     `json:"employee_number,omitempty"`          // 员工工号
	ThirdUserId            *string                     `json:"third_user_id,omitempty"`            // 第三方用户ID
	ConDepartmentIds       []string                    `json:"con_department_ids,omitempty"`       // 所在兼岗部门ID数组
	Email                  *string                     `json:"email,omitempty"`                    // 邮箱
	Department             *string                     `json:"department,omitempty"`               // 部门名称（老），后续此参数会去掉
	BranchName             *string                     `json:"branch_name,omitempty"`              // 所在分公司名称（老），后续此参数会去掉
	SystemRole             *int32                      `json:"system_role,omitempty"`              // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	RoleIds                *string                     `json:"role_ids,omitempty"`                 // 角色，多个用_连接
	Residentsname          *string                     `json:"residentsname,omitempty"`            // 常驻地中文
	ResidentsList          []MemberV2ResidentsListInfo `json:"residents_list,omitempty"`           // 常驻地列表
	UseCompanyMoney        *int32                      `json:"use_company_money,omitempty"`        // 是否企业支付余额，枚举值数字 0 否，1 是
	TotalQuota             *string                     `json:"total_quota,omitempty"`              // 每月配额，单位元
	AvailableQuota         *string                     `json:"available_quota,omitempty"`          // 员工可用限额，单位元
	LimitRuleList          []MemberV2LimitRuleInfo     `json:"limit_rule_list,omitempty"`          // 限额规则列表
	IsRemark               *string                     `json:"is_remark,omitempty"`                // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	BudgetCenterId         *string                     `json:"budget_center_id,omitempty"`         // 所在部门ID
	UseCarConfig           []string                    `json:"use_car_config,omitempty"`           // 用车规则ID数组
	RegulationId           []string                    `json:"regulation_id,omitempty"`            // 用车制度ID数组
	GuestCarRight          *int32                      `json:"guest_car_right,omitempty"`          // 客人用车权限
	SetDismissTime         *string                     `json:"set_dismiss_time,omitempty"`         // 设置的员工离职日期，为空时表示未设置离职日期，格式为 yyyy-MM-dd
	DismissTime            *string                     `json:"dismiss_time,omitempty"`             // 员工实际离职日期，为空时表示未离职，格式为 yyyy-MM-dd HH:mm:ss
	ImmediateSuperiorPhone *string                     `json:"immediate_superior_phone,omitempty"` // 员工直属上级的手机号码
	InvoiceInfo            *string                     `json:"invoice_info,omitempty"`             // 开票主体信息
	LegalEntityId          *string                     `json:"legal_entity_id,omitempty"`          // 所在公司主体id
	RankId                 *string                     `json:"rank_id,omitempty"`                  // 职级id
	OutRankId              *string                     `json:"out_rank_id,omitempty"`              // 外部职级 ID
	EnglishSurname         *string                     `json:"english_surname,omitempty"`          // 英文姓
	EnglishName            *string                     `json:"english_name,omitempty"`             // 英文名
	Nickname               *string                     `json:"nickname,omitempty"`                 // 昵称
	Sex                    *int32                      `json:"sex,omitempty"`                      // 性别，枚举值数字 0 未知 1 男 2 女
	BirthDate              *string                     `json:"birth_date,omitempty"`               // 出生日期，格式2000-01-01
	CardList               []MemberV2CardInfo          `json:"card_list,omitempty"`                // 证件信息
	Source                 *string                     `json:"source,omitempty"`                   // 员工加入滴滴企业平台的渠道
	ImmediateSuperiorEid   *string                     `json:"immediate_superior_eid,omitempty"`   // 员工直属上级的员工编号
	OutLegalEntityId       *string                     `json:"out_legal_entity_id,omitempty"`      // 外部公司主体编号
	CertRealname           *string                     `json:"cert_realname,omitempty"`            // 证件中文姓名
	CertEnglishSurname     *string                     `json:"cert_english_surname,omitempty"`     // 证件英文姓
	CertEnglishName        *string                     `json:"cert_english_name,omitempty"`        // 证件英文名
	Status                 *int32                      `json:"status,omitempty"`                   // 员工状态，枚举值数字：1 正常 4 离职 6 未绑定手机号
	MonthQuota             *string                     `json:"month_quota,omitempty"`              // 每月配额，同 total_quota
	BelongEnterpriseName   *string                     `json:"belong_enterprise_name,omitempty"`   // 归属企业名称（集团账户）
}

type MemberV2RecordBuilder struct {
	id                        string
	idSet                     bool
	phone                     string
	phoneSet                  bool
	realname                  string
	realnameSet               bool
	employeeNumber            string
	employeeNumberSet         bool
	thirdUserId               string
	thirdUserIdSet            bool
	email                     string
	emailSet                  bool
	department                string
	departmentSet             bool
	branchName                string
	branchNameSet             bool
	systemRole                int32
	systemRoleSet             bool
	roleIds                   string
	roleIdsSet                bool
	residentsname             string
	residentsnameSet          bool
	useCompanyMoney           int32
	useCompanyMoneySet        bool
	totalQuota                string
	totalQuotaSet             bool
	availableQuota            string
	availableQuotaSet         bool
	isRemark                  string
	isRemarkSet               bool
	budgetCenterId            string
	budgetCenterIdSet         bool
	setDismissTime            string
	setDismissTimeSet         bool
	dismissTime               string
	dismissTimeSet            bool
	immediateSuperiorPhone    string
	immediateSuperiorPhoneSet bool
	invoiceInfo               string
	invoiceInfoSet            bool
	legalEntityId             string
	legalEntityIdSet          bool
	rankId                    string
	rankIdSet                 bool
	outRankId                 string
	outRankIdSet              bool
	englishSurname            string
	englishSurnameSet         bool
	englishName               string
	englishNameSet            bool
	nickname                  string
	nicknameSet               bool
	sex                       int32
	sexSet                    bool
	birthDate                 string
	birthDateSet              bool
	source                    string
	sourceSet                 bool
	immediateSuperiorEid      string
	immediateSuperiorEidSet   bool
	outLegalEntityId          string
	outLegalEntityIdSet       bool
	certRealname              string
	certRealnameSet           bool
	certEnglishSurname        string
	certEnglishSurnameSet     bool
	certEnglishName           string
	certEnglishNameSet        bool
	status                    int32
	statusSet                 bool
	monthQuota                string
	monthQuotaSet             bool
	belongEnterpriseName      string
	belongEnterpriseNameSet   bool
	conDepartmentIds          []string
	conDepartmentIdsSet       bool
	residentsList             []MemberV2ResidentsListInfo
	residentsListSet          bool
	limitRuleList             []MemberV2LimitRuleInfo
	limitRuleListSet          bool
	useCarConfig              []string
	useCarConfigSet           bool
	regulationId              []string
	regulationIdSet           bool
	guestCarRight             int32
	guestCarRightSet          bool
	cardList                  []MemberV2CardInfo
	cardListSet               bool
}

func NewMemberV2RecordBuilder() *MemberV2RecordBuilder {
	return &MemberV2RecordBuilder{}
}
func (builder *MemberV2RecordBuilder) Id(id string) *MemberV2RecordBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Phone(phone string) *MemberV2RecordBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Realname(realname string) *MemberV2RecordBuilder {
	builder.realname = realname
	builder.realnameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) EmployeeNumber(employeeNumber string) *MemberV2RecordBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) ThirdUserId(thirdUserId string) *MemberV2RecordBuilder {
	builder.thirdUserId = thirdUserId
	builder.thirdUserIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Email(email string) *MemberV2RecordBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Department(department string) *MemberV2RecordBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) BranchName(branchName string) *MemberV2RecordBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) SystemRole(systemRole int32) *MemberV2RecordBuilder {
	builder.systemRole = systemRole
	builder.systemRoleSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) RoleIds(roleIds string) *MemberV2RecordBuilder {
	builder.roleIds = roleIds
	builder.roleIdsSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Residentsname(residentsname string) *MemberV2RecordBuilder {
	builder.residentsname = residentsname
	builder.residentsnameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) UseCompanyMoney(useCompanyMoney int32) *MemberV2RecordBuilder {
	builder.useCompanyMoney = useCompanyMoney
	builder.useCompanyMoneySet = true
	return builder
}
func (builder *MemberV2RecordBuilder) TotalQuota(totalQuota string) *MemberV2RecordBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) AvailableQuota(availableQuota string) *MemberV2RecordBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) IsRemark(isRemark string) *MemberV2RecordBuilder {
	builder.isRemark = isRemark
	builder.isRemarkSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) BudgetCenterId(budgetCenterId string) *MemberV2RecordBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) SetDismissTime(setDismissTime string) *MemberV2RecordBuilder {
	builder.setDismissTime = setDismissTime
	builder.setDismissTimeSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) DismissTime(dismissTime string) *MemberV2RecordBuilder {
	builder.dismissTime = dismissTime
	builder.dismissTimeSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) ImmediateSuperiorPhone(immediateSuperiorPhone string) *MemberV2RecordBuilder {
	builder.immediateSuperiorPhone = immediateSuperiorPhone
	builder.immediateSuperiorPhoneSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) InvoiceInfo(invoiceInfo string) *MemberV2RecordBuilder {
	builder.invoiceInfo = invoiceInfo
	builder.invoiceInfoSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) LegalEntityId(legalEntityId string) *MemberV2RecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) RankId(rankId string) *MemberV2RecordBuilder {
	builder.rankId = rankId
	builder.rankIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) OutRankId(outRankId string) *MemberV2RecordBuilder {
	builder.outRankId = outRankId
	builder.outRankIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) EnglishSurname(englishSurname string) *MemberV2RecordBuilder {
	builder.englishSurname = englishSurname
	builder.englishSurnameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) EnglishName(englishName string) *MemberV2RecordBuilder {
	builder.englishName = englishName
	builder.englishNameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Nickname(nickname string) *MemberV2RecordBuilder {
	builder.nickname = nickname
	builder.nicknameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Sex(sex int32) *MemberV2RecordBuilder {
	builder.sex = sex
	builder.sexSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) BirthDate(birthDate string) *MemberV2RecordBuilder {
	builder.birthDate = birthDate
	builder.birthDateSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Source(source string) *MemberV2RecordBuilder {
	builder.source = source
	builder.sourceSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) ImmediateSuperiorEid(immediateSuperiorEid string) *MemberV2RecordBuilder {
	builder.immediateSuperiorEid = immediateSuperiorEid
	builder.immediateSuperiorEidSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) OutLegalEntityId(outLegalEntityId string) *MemberV2RecordBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) CertRealname(certRealname string) *MemberV2RecordBuilder {
	builder.certRealname = certRealname
	builder.certRealnameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) CertEnglishSurname(certEnglishSurname string) *MemberV2RecordBuilder {
	builder.certEnglishSurname = certEnglishSurname
	builder.certEnglishSurnameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) CertEnglishName(certEnglishName string) *MemberV2RecordBuilder {
	builder.certEnglishName = certEnglishName
	builder.certEnglishNameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) Status(status int32) *MemberV2RecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) MonthQuota(monthQuota string) *MemberV2RecordBuilder {
	builder.monthQuota = monthQuota
	builder.monthQuotaSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) BelongEnterpriseName(belongEnterpriseName string) *MemberV2RecordBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) ConDepartmentIds(conDepartmentIds []string) *MemberV2RecordBuilder {
	builder.conDepartmentIds = conDepartmentIds
	builder.conDepartmentIdsSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) ResidentsList(residentsList []MemberV2ResidentsListInfo) *MemberV2RecordBuilder {
	builder.residentsList = residentsList
	builder.residentsListSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) LimitRuleList(limitRuleList []MemberV2LimitRuleInfo) *MemberV2RecordBuilder {
	builder.limitRuleList = limitRuleList
	builder.limitRuleListSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) UseCarConfig(useCarConfig []string) *MemberV2RecordBuilder {
	builder.useCarConfig = useCarConfig
	builder.useCarConfigSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) RegulationId(regulationId []string) *MemberV2RecordBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) GuestCarRight(guestCarRight int32) *MemberV2RecordBuilder {
	builder.guestCarRight = guestCarRight
	builder.guestCarRightSet = true
	return builder
}
func (builder *MemberV2RecordBuilder) CardList(cardList []MemberV2CardInfo) *MemberV2RecordBuilder {
	builder.cardList = cardList
	builder.cardListSet = true
	return builder
}

func (builder *MemberV2RecordBuilder) Build() *MemberV2Record {
	data := &MemberV2Record{}
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
	if builder.thirdUserIdSet {
		data.ThirdUserId = &builder.thirdUserId
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.branchNameSet {
		data.BranchName = &builder.branchName
	}
	if builder.systemRoleSet {
		data.SystemRole = &builder.systemRole
	}
	if builder.roleIdsSet {
		data.RoleIds = &builder.roleIds
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
	if builder.availableQuotaSet {
		data.AvailableQuota = &builder.availableQuota
	}
	if builder.isRemarkSet {
		data.IsRemark = &builder.isRemark
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.setDismissTimeSet {
		data.SetDismissTime = &builder.setDismissTime
	}
	if builder.dismissTimeSet {
		data.DismissTime = &builder.dismissTime
	}
	if builder.immediateSuperiorPhoneSet {
		data.ImmediateSuperiorPhone = &builder.immediateSuperiorPhone
	}
	if builder.invoiceInfoSet {
		data.InvoiceInfo = &builder.invoiceInfo
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
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
	if builder.sourceSet {
		data.Source = &builder.source
	}
	if builder.immediateSuperiorEidSet {
		data.ImmediateSuperiorEid = &builder.immediateSuperiorEid
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.certRealnameSet {
		data.CertRealname = &builder.certRealname
	}
	if builder.certEnglishSurnameSet {
		data.CertEnglishSurname = &builder.certEnglishSurname
	}
	if builder.certEnglishNameSet {
		data.CertEnglishName = &builder.certEnglishName
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.monthQuotaSet {
		data.MonthQuota = &builder.monthQuota
	}
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	if builder.conDepartmentIdsSet {
		data.ConDepartmentIds = builder.conDepartmentIds
	}
	if builder.residentsListSet {
		data.ResidentsList = builder.residentsList
	}
	if builder.limitRuleListSet {
		data.LimitRuleList = builder.limitRuleList
	}
	if builder.useCarConfigSet {
		data.UseCarConfig = builder.useCarConfig
	}
	if builder.regulationIdSet {
		data.RegulationId = builder.regulationId
	}
	if builder.guestCarRightSet {
		data.GuestCarRight = &builder.guestCarRight
	}
	if builder.cardListSet {
		data.CardList = builder.cardList
	}
	return data
}
