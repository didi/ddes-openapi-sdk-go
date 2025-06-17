package v1

// UpdateBudgetCenterRequest struct for UpdateBudgetCenterRequest
type UpdateBudgetCenterRequest struct {
	ClientId           *string            `json:"client_id,omitempty"`                // 申请应用时分配的AppKey
	AccessToken        *string            `json:"access_token,omitempty"`             // 授权后的access token
	CompanyId          *string            `json:"company_id,omitempty"`               // 企业ID
	Timestamp          *int32             `json:"timestamp,omitempty"`                // 当前时间戳(精确到秒级)
	Sign               *string            `json:"sign,omitempty"`                     // 签名
	Type               *int32             `json:"type,omitempty"`                     // 类型，枚举值数字 1 部门 2 项目 使用外部out_budget_id 和name 值更新 type字段必传
	BudgetCycle        *int32             `json:"budget_cycle,omitempty"`             // 预算周期，枚举值数字 0：不限额；1：自然月 2：自然季度 3：自然年（其中23只对部门生效，需要设置白名单，须联系客户经理）
	TotalQuota         *string            `json:"total_quota,omitempty"`              // 总金额，单位元 0表示不限额度 小数点后保留两位精度
	Id                 *string            `json:"id,omitempty"`                       // 部门/项目 ID，type = 1 时， id和out_budget_id 优先处理id type = 2时 id和 out_budget_id与 name 组合唯一值时 优先处理id
	Name               *string            `json:"name,omitempty"`                     // 部门/项目名称，不大于 200 字符
	OutBudgetId        *string            `json:"out_budget_id,omitempty"`            // 编号，type = 1 时必填；type = 2 非必填； 长度限制：≤ 64 字符 更新部门时。字段必填。
	LeaderId           *string            `json:"leader_id,omitempty"`                // 主管ID，人员同步时返回的memberid leader_id和leader_employee_id 时，优先处理leader_id，主管id，多个用英文逗号分开，最多30个
	LeaderEmployeeId   *string            `json:"leader_employee_id,omitempty"`       // 主管员工编号，主管工号，json字符串，第一个是主要主管，后续是其他主管，最多30个，leader_id存在时不生效，举例：[111,2222,44444]
	ParentId           *string            `json:"parent_id,omitempty"`                // 上级部门/项目 ID，新建部门时返回的部门ID，type = 1 非必填，不传默认为顶级部门ID，不可以传空字符，type=2 默认为，parent_id优先级大于 out_parent_id和 out_parent_name
	OutParentId        *string            `json:"out_parent_id,omitempty"`            // 上级部门/项目外部CODE，type = 1 非必填，不传默认为顶级部门code，不可以传空字符，type=2 默认为空，项目需要out_parent_id和out_parent_name一起传递，作为唯一值校验。
	OutParentName      *string            `json:"out_parent_name,omitempty"`          // 上级部门/项目外部名称，type = 1 非必填，不传默认为顶级部门code，type=2 默认为空 ，项目需要out_parent_id和out_parent_name一起传递，作为唯一值校验。
	MemberUsed         *int32             `json:"member_used,omitempty"`              // 使用范围，枚举值数字 type=2时生效，0 ：全员可见，1：项目成员内可见，2：公司主体内可见，不传默认为0 (枚举 2需要设置白名单，须联系客户经理。报错误码10001)
	StartDate          *string            `json:"start_date,omitempty"`               // 项目开始日期，默认为空 格式：yyyy-MM-dd type=2时生效
	ExpiryDate         *string            `json:"expiry_date,omitempty"`              // 项目结束日期，默认为空 格式：yyyy-MM-dd type=2时生效
	LegalEntityId      *string            `json:"legal_entity_id,omitempty"`          // 公司主体ID，多个用英文逗号分开，字段不传不生效，传空字符串清空，传非空字符串更新(更新已有值。后台无值时新建)，如果对应的公司主体id已经停用或者不存在 返回错误码10001 type=2时生效
	BudgetExtraInfo    *string            `json:"budget_extra_info,omitempty"`        // 项目扩展信息的自定义字段，项目扩展信息的自定义字段；最长不大于 500 字符；(必须为json字符串，json解析后不能为空)；仅对项目（type=2）生效，部门传了不生效；
	BudgetExtraInfoObj *map[string]string `json:"budget_extra_info__obj__,omitempty"` // 项目拓展字段对象，实际使用的时候会自动转换为json字符串赋值到budget_extra_info
}

type UpdateBudgetCenterRequestBuilder struct {
	clientId              string // 申请应用时分配的AppKey
	clientIdSet           bool
	accessToken           string // 授权后的access token
	accessTokenSet        bool
	companyId             string // 企业ID
	companyIdSet          bool
	timestamp             int32 // 当前时间戳(精确到秒级)
	timestampSet          bool
	sign                  string // 签名
	signSet               bool
	type_                 int32 // 类型，枚举值数字 1 部门 2 项目 使用外部out_budget_id 和name 值更新 type字段必传
	type_Set              bool
	budgetCycle           int32 // 预算周期，枚举值数字 0：不限额；1：自然月 2：自然季度 3：自然年（其中23只对部门生效，需要设置白名单，须联系客户经理）
	budgetCycleSet        bool
	totalQuota            string // 总金额，单位元 0表示不限额度 小数点后保留两位精度
	totalQuotaSet         bool
	id                    string // 部门/项目 ID，type = 1 时， id和out_budget_id 优先处理id type = 2时 id和 out_budget_id与 name 组合唯一值时 优先处理id
	idSet                 bool
	name                  string // 部门/项目名称，不大于 200 字符
	nameSet               bool
	outBudgetId           string // 编号，type = 1 时必填；type = 2 非必填； 长度限制：≤ 64 字符 更新部门时。字段必填。
	outBudgetIdSet        bool
	leaderId              string // 主管ID，人员同步时返回的memberid leader_id和leader_employee_id 时，优先处理leader_id，主管id，多个用英文逗号分开，最多30个
	leaderIdSet           bool
	leaderEmployeeId      string // 主管员工编号，主管工号，json字符串，第一个是主要主管，后续是其他主管，最多30个，leader_id存在时不生效，举例：[111,2222,44444]
	leaderEmployeeIdSet   bool
	parentId              string // 上级部门/项目 ID，新建部门时返回的部门ID，type = 1 非必填，不传默认为顶级部门ID，不可以传空字符，type=2 默认为，parent_id优先级大于 out_parent_id和 out_parent_name
	parentIdSet           bool
	outParentId           string // 上级部门/项目外部CODE，type = 1 非必填，不传默认为顶级部门code，不可以传空字符，type=2 默认为空，项目需要out_parent_id和out_parent_name一起传递，作为唯一值校验。
	outParentIdSet        bool
	outParentName         string // 上级部门/项目外部名称，type = 1 非必填，不传默认为顶级部门code，type=2 默认为空 ，项目需要out_parent_id和out_parent_name一起传递，作为唯一值校验。
	outParentNameSet      bool
	memberUsed            int32 // 使用范围，枚举值数字 type=2时生效，0 ：全员可见，1：项目成员内可见，2：公司主体内可见，不传默认为0 (枚举 2需要设置白名单，须联系客户经理。报错误码10001)
	memberUsedSet         bool
	startDate             string // 项目开始日期，默认为空 格式：yyyy-MM-dd type=2时生效
	startDateSet          bool
	expiryDate            string // 项目结束日期，默认为空 格式：yyyy-MM-dd type=2时生效
	expiryDateSet         bool
	legalEntityId         string // 公司主体ID，多个用英文逗号分开，字段不传不生效，传空字符串清空，传非空字符串更新(更新已有值。后台无值时新建)，如果对应的公司主体id已经停用或者不存在 返回错误码10001 type=2时生效
	legalEntityIdSet      bool
	budgetExtraInfo       string // 项目扩展信息的自定义字段，项目扩展信息的自定义字段；最长不大于 500 字符；(必须为json字符串，json解析后不能为空)；仅对项目（type=2）生效，部门传了不生效；
	budgetExtraInfoSet    bool
	budgetExtraInfoObj    map[string]string // 项目拓展字段对象，实际使用的时候会自动转换为json字符串赋值到budget_extra_info
	budgetExtraInfoObjSet bool
}

func NewUpdateBudgetCenterRequestBuilder() *UpdateBudgetCenterRequestBuilder {
	return &UpdateBudgetCenterRequestBuilder{}
}
func (builder *UpdateBudgetCenterRequestBuilder) ClientId(clientId string) *UpdateBudgetCenterRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) AccessToken(accessToken string) *UpdateBudgetCenterRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) CompanyId(companyId string) *UpdateBudgetCenterRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) Timestamp(timestamp int32) *UpdateBudgetCenterRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) Sign(sign string) *UpdateBudgetCenterRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) Type(type_ int32) *UpdateBudgetCenterRequestBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) BudgetCycle(budgetCycle int32) *UpdateBudgetCenterRequestBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) TotalQuota(totalQuota string) *UpdateBudgetCenterRequestBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) Id(id string) *UpdateBudgetCenterRequestBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) Name(name string) *UpdateBudgetCenterRequestBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) OutBudgetId(outBudgetId string) *UpdateBudgetCenterRequestBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) LeaderId(leaderId string) *UpdateBudgetCenterRequestBuilder {
	builder.leaderId = leaderId
	builder.leaderIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) LeaderEmployeeId(leaderEmployeeId string) *UpdateBudgetCenterRequestBuilder {
	builder.leaderEmployeeId = leaderEmployeeId
	builder.leaderEmployeeIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) ParentId(parentId string) *UpdateBudgetCenterRequestBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) OutParentId(outParentId string) *UpdateBudgetCenterRequestBuilder {
	builder.outParentId = outParentId
	builder.outParentIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) OutParentName(outParentName string) *UpdateBudgetCenterRequestBuilder {
	builder.outParentName = outParentName
	builder.outParentNameSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) MemberUsed(memberUsed int32) *UpdateBudgetCenterRequestBuilder {
	builder.memberUsed = memberUsed
	builder.memberUsedSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) StartDate(startDate string) *UpdateBudgetCenterRequestBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) ExpiryDate(expiryDate string) *UpdateBudgetCenterRequestBuilder {
	builder.expiryDate = expiryDate
	builder.expiryDateSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) LegalEntityId(legalEntityId string) *UpdateBudgetCenterRequestBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) BudgetExtraInfo(budgetExtraInfo string) *UpdateBudgetCenterRequestBuilder {
	builder.budgetExtraInfo = budgetExtraInfo
	builder.budgetExtraInfoSet = true
	return builder
}
func (builder *UpdateBudgetCenterRequestBuilder) BudgetExtraInfoObj(budgetExtraInfoObj map[string]string) *UpdateBudgetCenterRequestBuilder {
	builder.budgetExtraInfoObj = budgetExtraInfoObj
	builder.budgetExtraInfoObjSet = true
	return builder
}

func (builder *UpdateBudgetCenterRequestBuilder) Build() *UpdateBudgetCenterRequest {
	data := &UpdateBudgetCenterRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.budgetCycleSet {
		data.BudgetCycle = &builder.budgetCycle
	}
	if builder.totalQuotaSet {
		data.TotalQuota = &builder.totalQuota
	}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.leaderIdSet {
		data.LeaderId = &builder.leaderId
	}
	if builder.leaderEmployeeIdSet {
		data.LeaderEmployeeId = &builder.leaderEmployeeId
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
	}
	if builder.outParentIdSet {
		data.OutParentId = &builder.outParentId
	}
	if builder.outParentNameSet {
		data.OutParentName = &builder.outParentName
	}
	if builder.memberUsedSet {
		data.MemberUsed = &builder.memberUsed
	}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.expiryDateSet {
		data.ExpiryDate = &builder.expiryDate
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.budgetExtraInfoSet {
		data.BudgetExtraInfo = &builder.budgetExtraInfo
	}
	if builder.budgetExtraInfoObjSet {
		data.BudgetExtraInfoObj = &builder.budgetExtraInfoObj
	}
	return data
}
