package v1

// GetMemberDetailReply struct for GetMemberDetailReply
type GetMemberDetailReply struct {
	MemberId               *string             `json:"member_id,omitempty"`                // 员工在滴滴企业平台的ID
	Phone                  *string             `json:"phone,omitempty"`                    // 员工手机号
	Realname               *string             `json:"realname,omitempty"`                 // 员工姓名
	EmployeeNumber         *string             `json:"employee_number,omitempty"`          // 员工工号，员工在公司的员工号
	Email                  *string             `json:"email,omitempty"`                    // 邮箱
	SystemRole             *int32              `json:"system_role,omitempty"`              // 系统角色，枚举值数字 0 车辆预定人员，1 普通管理员，2 超级管理员
	RoleIds                *string             `json:"role_ids,omitempty"`                 // 角色，可以通过角色API获取对应的ID
	ImmediateSuperiorPhone *string             `json:"immediate_superior_phone,omitempty"` // 员工直属上级的手机号码，直属上级可在审批流中担任审批人
	ImmediateSuperiorEid   *string             `json:"immediate_superior_eid,omitempty"`   // 员工直属上级的员工编号
	Residentsname          *string             `json:"residentsname,omitempty"`            // 常驻地中文
	UseCompanyMoney        *int32              `json:"use_company_money,omitempty"`        // 是否企业支付余额，枚举值数字 0 否，1 是
	TotalQuota             *string             `json:"total_quota,omitempty"`              // 每月配额，单位元
	IsRemark               *int32              `json:"is_remark,omitempty"`                // 叫车时备注信息是否必填，枚举值数字 0 选填，1 必填，2 按制度填写
	BudgetCenterId         *string             `json:"budget_center_id,omitempty"`         // 所在部门ID
	ConDepartmentIds       []string            `json:"con_department_ids,omitempty"`       // 所在兼岗部门ID，兼岗部门ID数组
	UseCarConfig           []string            `json:"use_car_config,omitempty"`           // 用车规则ID数组
	ProjectIds             *string             `json:"project_ids,omitempty"`              // 所在项目ID
	AvailableQuota         *string             `json:"available_quota,omitempty"`          // 员工可用限额，单位元 例如\"937.70\"
	BranchName             *string             `json:"branch_name,omitempty"`              // 所在分公司名称（老），后续此参数会去掉
	Department             *string             `json:"department,omitempty"`               // 部门名称（老），后续此参数会去掉
	RegulationId           []string            `json:"regulation_id,omitempty"`            // 用车制度ID数组，跟新增、修改、详情等API统一
	SetDismissTime         *string             `json:"set_dismiss_time,omitempty"`         // 设置的员工离职日期，为空时表示未设置离职日期，格式为 yyyy-MM-dd
	DismissTime            *string             `json:"dismiss_time,omitempty"`             // 员工实际离职日期，为空时表示未离职,格式为 格式：yyyy-MM-dd HH:mm:ss
	InvoiceInfo            *string             `json:"invoice_info,omitempty"`             // 开票主体信息
	LegalEntityId          *string             `json:"legal_entity_id,omitempty"`          // 所在公司主体id
	OutLegalEntityId       *string             `json:"out_legal_entity_id,omitempty"`      // 外部公司主体编号
	RankId                 *string             `json:"rank_id,omitempty"`                  // 职级id
	OutRankId              *string             `json:"out_rank_id,omitempty"`              // 外部职级 ID
	EnglishSurname         *string             `json:"english_surname,omitempty"`          // 英文姓，同lastname
	EnglishName            *string             `json:"english_name,omitempty"`             // 英文名，同firstname 有middlename时 english_name=firstname middlename
	Nickname               *string             `json:"nickname,omitempty"`                 // 昵称
	Sex                    *int32              `json:"sex,omitempty"`                      // 性别，枚举值数字 0 未知 1 男 2 女
	BirthDate              *string             `json:"birth_date,omitempty"`               // 出生日期，格式2000-01-01（已用AES算法加密）
	CardList               []CardInfo          `json:"card_list,omitempty"`                // 证件信息
	Source                 *string             `json:"source,omitempty"`                   // 员工加入滴滴企业平台的渠道，枚举值数字：0;未知 1;PC逐一添加 2;PC批量导入 3;邮件邀请 4;API添加
	Status                 *int32              `json:"status,omitempty"`                   // 员工状态，枚举值数字：1 正常 4 离职 6 未绑定手机号
	ResidentsList          []ResidentsListInfo `json:"residents_list,omitempty"`           // 常驻地列表
	LimitRuleList          []LimitRuleInfo     `json:"limit_rule_list,omitempty"`          // 限额规则列表
	CertRealname           *string             `json:"cert_realname,omitempty"`            // 证件中文姓名，证件上的真实中文姓名
	CertEnglishSurname     *string             `json:"cert_english_surname,omitempty"`     // 证件英文姓，证件上的真实英文姓
	CertEnglishName        *string             `json:"cert_english_name,omitempty"`        // 证件英文名，证件上的真实英文名
	HomeAddress            []HomeAddressInfo   `json:"home_address,omitempty"`             // 家庭住址信息，用于用车管控
}
