package v1

// MemberInfo struct for MemberInfo
type MemberInfo struct {
	MemberType                      *int32     `json:"member_type,omitempty"`                        // 员工信息类型，枚举值数字 0：手机号，1：工号。2：邮箱；默认为0
	Phone                           *string    `json:"phone,omitempty"`                              // 员工手机号，member_type 为0时必传
	Realname                        *string    `json:"realname,omitempty"`                           // 员工姓名
	EmployeeNumber                  *string    `json:"employee_number,omitempty"`                    // 员工工号，member_type 为1时必传 员工在公司的员工号
	Email                           *string    `json:"email,omitempty"`                              // 邮箱，member_type 为2时必传
	SystemRole                      *int32     `json:"system_role,omitempty"`                        // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	RoleIds                         *string    `json:"role_ids,omitempty"`                           // 角色，可以通过角色API获取对应的ID
	ImmediateSuperiorPhone          *string    `json:"immediate_superior_phone,omitempty"`           // 员工直属上级的手机号码，直属上级可在审批流中担任审批人 immediate_superior_phone与immediate_superior_eid以手机号优先
	ImmediateSuperiorEmail          *string    `json:"immediate_superior_email,omitempty"`           // 直属上级邮箱
	ImmediateSuperiorEmployeeNumber *string    `json:"immediate_superior_employee_number,omitempty"` // 员工直属上级的员工编号，直属上级可在审批流中担任审批人
	ImmediateSuperiorMemberID       *int32     `json:"immediate_superior_memberID,omitempty"`        // 直属上级 ID
	ClearImmediateSuperior          *int32     `json:"clear_immediate_superior,omitempty"`           // 清除上级；数字1 清除；【更新参数】
	Residentsname                   *string    `json:"residentsname,omitempty"`                      // 常驻地中文
	RedisentsIds                    *string    `json:"redisents_ids,omitempty"`                      // 常驻地ID，不传不更新，传“”清空，多个使用“_”连接；
	UseCompanyMoney                 *int32     `json:"use_company_money,omitempty"`                  // 是否企业支付余额，枚举值数字 0 否，1 是
	TotalQuota                      *string    `json:"total_quota,omitempty"`                        // 每月配额，单位元 0 不限
	IsRemark                        *int32     `json:"is_remark,omitempty"`                          // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	BudgetCenterId                  *int64     `json:"budget_center_id,omitempty"`                   // 所在部门ID，budget_center_id与out_budget_id同时存在时，以budget_center_id为准
	OutBudgetId                     *string    `json:"out_budget_id,omitempty"`                      // 客户部门CODE
	ConDepartmentIds                *string    `json:"con_department_ids,omitempty"`                 // 所在兼岗部门ID，con_department_ids与con_department_codes都存在时，以con_department_ids为准 多个使用“_”连接
	ConDepartmentCodes              *string    `json:"con_department_codes,omitempty"`               // 所在兼岗部门CODE（同部门新增修改的out_budget_id），con_department_ids与con_department_codes都存在时，以con_department_ids为准 多个使用“_”连接
	RegulationId                    *string    `json:"regulation_id,omitempty"`                      // 用车制度ID数组，制度ID；通过制度列表接口查询；多个用 _ 连接；默认为空，若该员工的所有制度都是在es后台通过部门/职级/全员方式分配，则员工身上的制度字段不用传；同时注意检查use_company_money字段是否传输，制度和企业支付权限都有才能企业支付
	SetDismissTime                  *string    `json:"set_dismiss_time,omitempty"`                   // 设置的员工离职日期，设置员工离职日期，到期后自动加入已离职名单，不传或为空时认为不设置离职时间，不传不更新；传空更新为空 格式为 yyyy-MM-dd，时间需要传入大于今天，隔天凌晨处理离职时间。
	ProjectIds                      *string    `json:"project_ids,omitempty"`                        // 所在项目ID，可以填多个，以_分隔。通过成本中心查询api获取id（类型为2）
	ProjectCodesDetail              *string    `json:"project_codes_detail,omitempty"`               // 项目信息，人员上绑定的项目信息，将project_codes_detail的值转为 json 字符 \"project_codes_detail\":\"[{\"project_name\":\"出差巡视\",\"project_code\":\"travelcode\" },{\"project_name\":\"出差巡视\",\"project_code\":\"travelcode\" }] \"，project_ids与project_code_detail同时有值时，以project_ids为准。元素上限100个
	LegalEntityId                   *string    `json:"legal_entity_id,omitempty"`                    // 所在公司主体id
	OutLegalEntityId                *string    `json:"out_legal_entity_id,omitempty"`                // 外部所在公司主体id
	RankId                          *string    `json:"rank_id,omitempty"`                            // 职级id
	OutRankId                       *string    `json:"out_rank_id,omitempty"`                        // 外部职级 ID
	EnglishSurname                  *string    `json:"english_surname,omitempty"`                    // 英文姓，同lastname
	EnglishName                     *string    `json:"english_name,omitempty"`                       // 英文名，同firstname 有middlename时 english_name=firstname middlename
	Nickname                        *string    `json:"nickname,omitempty"`                           // 昵称
	Sex                             *int32     `json:"sex,omitempty"`                                // 性别，枚举值数字 0 未知 1 男 2 女
	BirthDate                       *string    `json:"birth_date,omitempty"`                         // 出生日期，格式2000-01-01，注：1、若采用AES256整体加密，此字段需明文传输，无需单独再加密 2、若不整体加密传输时，此字段只可采用AES128加密传输 3、若采用AES128整体加密，此字段仍需采用AES128单独加密（存在历史客户原因）
	CardList                        []CardInfo `json:"card_list,omitempty"`                          // 证件信息
}

type MemberInfoBuilder struct {
	memberType                         int32 // 员工信息类型，枚举值数字 0：手机号，1：工号。2：邮箱；默认为0
	memberTypeSet                      bool
	phone                              string // 员工手机号，member_type 为0时必传
	phoneSet                           bool
	realname                           string // 员工姓名
	realnameSet                        bool
	employeeNumber                     string // 员工工号，member_type 为1时必传 员工在公司的员工号
	employeeNumberSet                  bool
	email                              string // 邮箱，member_type 为2时必传
	emailSet                           bool
	systemRole                         int32 // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	systemRoleSet                      bool
	roleIds                            string // 角色，可以通过角色API获取对应的ID
	roleIdsSet                         bool
	immediateSuperiorPhone             string // 员工直属上级的手机号码，直属上级可在审批流中担任审批人 immediate_superior_phone与immediate_superior_eid以手机号优先
	immediateSuperiorPhoneSet          bool
	immediateSuperiorEmail             string // 直属上级邮箱
	immediateSuperiorEmailSet          bool
	immediateSuperiorEmployeeNumber    string // 员工直属上级的员工编号，直属上级可在审批流中担任审批人
	immediateSuperiorEmployeeNumberSet bool
	immediateSuperiorMemberID          int32 // 直属上级 ID
	immediateSuperiorMemberIDSet       bool
	clearImmediateSuperior             int32 // 清除上级；数字1 清除；【更新参数】
	clearImmediateSuperiorSet          bool
	residentsname                      string // 常驻地中文
	residentsnameSet                   bool
	redisentsIds                       string // 常驻地ID，不传不更新，传“”清空，多个使用“_”连接；
	redisentsIdsSet                    bool
	useCompanyMoney                    int32 // 是否企业支付余额，枚举值数字 0 否，1 是
	useCompanyMoneySet                 bool
	totalQuota                         string // 每月配额，单位元 0 不限
	totalQuotaSet                      bool
	isRemark                           int32 // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	isRemarkSet                        bool
	budgetCenterId                     int64 // 所在部门ID，budget_center_id与out_budget_id同时存在时，以budget_center_id为准
	budgetCenterIdSet                  bool
	outBudgetId                        string // 客户部门CODE
	outBudgetIdSet                     bool
	conDepartmentIds                   string // 所在兼岗部门ID，con_department_ids与con_department_codes都存在时，以con_department_ids为准 多个使用“_”连接
	conDepartmentIdsSet                bool
	conDepartmentCodes                 string // 所在兼岗部门CODE（同部门新增修改的out_budget_id），con_department_ids与con_department_codes都存在时，以con_department_ids为准 多个使用“_”连接
	conDepartmentCodesSet              bool
	regulationId                       string // 用车制度ID数组，制度ID；通过制度列表接口查询；多个用 _ 连接；默认为空，若该员工的所有制度都是在es后台通过部门/职级/全员方式分配，则员工身上的制度字段不用传；同时注意检查use_company_money字段是否传输，制度和企业支付权限都有才能企业支付
	regulationIdSet                    bool
	setDismissTime                     string // 设置的员工离职日期，设置员工离职日期，到期后自动加入已离职名单，不传或为空时认为不设置离职时间，不传不更新；传空更新为空 格式为 yyyy-MM-dd，时间需要传入大于今天，隔天凌晨处理离职时间。
	setDismissTimeSet                  bool
	projectIds                         string // 所在项目ID，可以填多个，以_分隔。通过成本中心查询api获取id（类型为2）
	projectIdsSet                      bool
	projectCodesDetail                 string // 项目信息，人员上绑定的项目信息，将project_codes_detail的值转为 json 字符 \"project_codes_detail\":\"[{\"project_name\":\"出差巡视\",\"project_code\":\"travelcode\" },{\"project_name\":\"出差巡视\",\"project_code\":\"travelcode\" }] \"，project_ids与project_code_detail同时有值时，以project_ids为准。元素上限100个
	projectCodesDetailSet              bool
	legalEntityId                      string // 所在公司主体id
	legalEntityIdSet                   bool
	outLegalEntityId                   string // 外部所在公司主体id
	outLegalEntityIdSet                bool
	rankId                             string // 职级id
	rankIdSet                          bool
	outRankId                          string // 外部职级 ID
	outRankIdSet                       bool
	englishSurname                     string // 英文姓，同lastname
	englishSurnameSet                  bool
	englishName                        string // 英文名，同firstname 有middlename时 english_name=firstname middlename
	englishNameSet                     bool
	nickname                           string // 昵称
	nicknameSet                        bool
	sex                                int32 // 性别，枚举值数字 0 未知 1 男 2 女
	sexSet                             bool
	birthDate                          string // 出生日期，格式2000-01-01，注：1、若采用AES256整体加密，此字段需明文传输，无需单独再加密 2、若不整体加密传输时，此字段只可采用AES128加密传输 3、若采用AES128整体加密，此字段仍需采用AES128单独加密（存在历史客户原因）
	birthDateSet                       bool
	cardList                           []CardInfo // 证件信息
	cardListSet                        bool
}

func NewMemberInfoBuilder() *MemberInfoBuilder {
	return &MemberInfoBuilder{}
}
func (builder *MemberInfoBuilder) MemberType(memberType int32) *MemberInfoBuilder {
	builder.memberType = memberType
	builder.memberTypeSet = true
	return builder
}
func (builder *MemberInfoBuilder) Phone(phone string) *MemberInfoBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *MemberInfoBuilder) Realname(realname string) *MemberInfoBuilder {
	builder.realname = realname
	builder.realnameSet = true
	return builder
}
func (builder *MemberInfoBuilder) EmployeeNumber(employeeNumber string) *MemberInfoBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *MemberInfoBuilder) Email(email string) *MemberInfoBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *MemberInfoBuilder) SystemRole(systemRole int32) *MemberInfoBuilder {
	builder.systemRole = systemRole
	builder.systemRoleSet = true
	return builder
}
func (builder *MemberInfoBuilder) RoleIds(roleIds string) *MemberInfoBuilder {
	builder.roleIds = roleIds
	builder.roleIdsSet = true
	return builder
}
func (builder *MemberInfoBuilder) ImmediateSuperiorPhone(immediateSuperiorPhone string) *MemberInfoBuilder {
	builder.immediateSuperiorPhone = immediateSuperiorPhone
	builder.immediateSuperiorPhoneSet = true
	return builder
}
func (builder *MemberInfoBuilder) ImmediateSuperiorEmail(immediateSuperiorEmail string) *MemberInfoBuilder {
	builder.immediateSuperiorEmail = immediateSuperiorEmail
	builder.immediateSuperiorEmailSet = true
	return builder
}
func (builder *MemberInfoBuilder) ImmediateSuperiorEmployeeNumber(immediateSuperiorEmployeeNumber string) *MemberInfoBuilder {
	builder.immediateSuperiorEmployeeNumber = immediateSuperiorEmployeeNumber
	builder.immediateSuperiorEmployeeNumberSet = true
	return builder
}
func (builder *MemberInfoBuilder) ImmediateSuperiorMemberID(immediateSuperiorMemberID int32) *MemberInfoBuilder {
	builder.immediateSuperiorMemberID = immediateSuperiorMemberID
	builder.immediateSuperiorMemberIDSet = true
	return builder
}
func (builder *MemberInfoBuilder) ClearImmediateSuperior(clearImmediateSuperior int32) *MemberInfoBuilder {
	builder.clearImmediateSuperior = clearImmediateSuperior
	builder.clearImmediateSuperiorSet = true
	return builder
}
func (builder *MemberInfoBuilder) Residentsname(residentsname string) *MemberInfoBuilder {
	builder.residentsname = residentsname
	builder.residentsnameSet = true
	return builder
}
func (builder *MemberInfoBuilder) RedisentsIds(redisentsIds string) *MemberInfoBuilder {
	builder.redisentsIds = redisentsIds
	builder.redisentsIdsSet = true
	return builder
}
func (builder *MemberInfoBuilder) UseCompanyMoney(useCompanyMoney int32) *MemberInfoBuilder {
	builder.useCompanyMoney = useCompanyMoney
	builder.useCompanyMoneySet = true
	return builder
}
func (builder *MemberInfoBuilder) TotalQuota(totalQuota string) *MemberInfoBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *MemberInfoBuilder) IsRemark(isRemark int32) *MemberInfoBuilder {
	builder.isRemark = isRemark
	builder.isRemarkSet = true
	return builder
}
func (builder *MemberInfoBuilder) BudgetCenterId(budgetCenterId int64) *MemberInfoBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) OutBudgetId(outBudgetId string) *MemberInfoBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) ConDepartmentIds(conDepartmentIds string) *MemberInfoBuilder {
	builder.conDepartmentIds = conDepartmentIds
	builder.conDepartmentIdsSet = true
	return builder
}
func (builder *MemberInfoBuilder) ConDepartmentCodes(conDepartmentCodes string) *MemberInfoBuilder {
	builder.conDepartmentCodes = conDepartmentCodes
	builder.conDepartmentCodesSet = true
	return builder
}
func (builder *MemberInfoBuilder) RegulationId(regulationId string) *MemberInfoBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) SetDismissTime(setDismissTime string) *MemberInfoBuilder {
	builder.setDismissTime = setDismissTime
	builder.setDismissTimeSet = true
	return builder
}
func (builder *MemberInfoBuilder) ProjectIds(projectIds string) *MemberInfoBuilder {
	builder.projectIds = projectIds
	builder.projectIdsSet = true
	return builder
}
func (builder *MemberInfoBuilder) ProjectCodesDetail(projectCodesDetail string) *MemberInfoBuilder {
	builder.projectCodesDetail = projectCodesDetail
	builder.projectCodesDetailSet = true
	return builder
}
func (builder *MemberInfoBuilder) LegalEntityId(legalEntityId string) *MemberInfoBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) OutLegalEntityId(outLegalEntityId string) *MemberInfoBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) RankId(rankId string) *MemberInfoBuilder {
	builder.rankId = rankId
	builder.rankIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) OutRankId(outRankId string) *MemberInfoBuilder {
	builder.outRankId = outRankId
	builder.outRankIdSet = true
	return builder
}
func (builder *MemberInfoBuilder) EnglishSurname(englishSurname string) *MemberInfoBuilder {
	builder.englishSurname = englishSurname
	builder.englishSurnameSet = true
	return builder
}
func (builder *MemberInfoBuilder) EnglishName(englishName string) *MemberInfoBuilder {
	builder.englishName = englishName
	builder.englishNameSet = true
	return builder
}
func (builder *MemberInfoBuilder) Nickname(nickname string) *MemberInfoBuilder {
	builder.nickname = nickname
	builder.nicknameSet = true
	return builder
}
func (builder *MemberInfoBuilder) Sex(sex int32) *MemberInfoBuilder {
	builder.sex = sex
	builder.sexSet = true
	return builder
}
func (builder *MemberInfoBuilder) BirthDate(birthDate string) *MemberInfoBuilder {
	builder.birthDate = birthDate
	builder.birthDateSet = true
	return builder
}
func (builder *MemberInfoBuilder) CardList(cardList []CardInfo) *MemberInfoBuilder {
	builder.cardList = cardList
	builder.cardListSet = true
	return builder
}

func (builder *MemberInfoBuilder) Build() *MemberInfo {
	data := &MemberInfo{}
	if builder.memberTypeSet {
		data.MemberType = &builder.memberType
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
	if builder.immediateSuperiorEmailSet {
		data.ImmediateSuperiorEmail = &builder.immediateSuperiorEmail
	}
	if builder.immediateSuperiorEmployeeNumberSet {
		data.ImmediateSuperiorEmployeeNumber = &builder.immediateSuperiorEmployeeNumber
	}
	if builder.immediateSuperiorMemberIDSet {
		data.ImmediateSuperiorMemberID = &builder.immediateSuperiorMemberID
	}
	if builder.clearImmediateSuperiorSet {
		data.ClearImmediateSuperior = &builder.clearImmediateSuperior
	}
	if builder.residentsnameSet {
		data.Residentsname = &builder.residentsname
	}
	if builder.redisentsIdsSet {
		data.RedisentsIds = &builder.redisentsIds
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
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.conDepartmentIdsSet {
		data.ConDepartmentIds = &builder.conDepartmentIds
	}
	if builder.conDepartmentCodesSet {
		data.ConDepartmentCodes = &builder.conDepartmentCodes
	}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.setDismissTimeSet {
		data.SetDismissTime = &builder.setDismissTime
	}
	if builder.projectIdsSet {
		data.ProjectIds = &builder.projectIds
	}
	if builder.projectCodesDetailSet {
		data.ProjectCodesDetail = &builder.projectCodesDetail
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
	return data
}
