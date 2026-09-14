package v1

// BudgetCenterV2Record 部门/项目记录（/open-apis/v2/budgetCenter/get 响应 records 元素，契约依据 idl/budget-center.idl BudgetCenterInfo）。
// v2 契约与 /river 版分叉：新增 company_id/belong_enterprise_name/create_time/update_time；无 poi_list/extend_field/department_id/out_department_id/scope。
type BudgetCenterV2Record struct {
	Id                   *string                       `json:"id,omitempty"`                     // 滴滴内部项目或部门的ID
	Name                 *string                       `json:"name,omitempty"`                   // 名称
	CompanyId            *string                       `json:"company_id,omitempty"`             // 公司ID
	BelongEnterpriseName *string                       `json:"belong_enterprise_name,omitempty"` // 归属企业名称（集团账户）
	Status               *string                       `json:"status,omitempty"`                 // 状态
	ThirdOutId           *string                       `json:"third_out_id,omitempty"`           // 第三方外部ID
	Type                 *string                       `json:"type,omitempty"`                   // 类型，枚举值 1 部门 2 项目
	BudgetCycle          *int32                        `json:"budget_cycle,omitempty"`           // 预算周期，枚举值数字 0 不限期 1 自然月 2 自然季度 3 自然年
	OutBudgetId          *string                       `json:"out_budget_id,omitempty"`          // 外部成本中心id
	TotalQuota           *string                       `json:"total_quota,omitempty"`            // 总金额，单位元
	IsLimitQuota         *int32                        `json:"is_limit_quota,omitempty"`         // 是否限额，枚举值数字 0 不限制 1 限制
	MemberNum            *int64                        `json:"member_num,omitempty"`             // 在使用人数
	AvailableQuota       *string                       `json:"available_quota,omitempty"`        // 可用金额，单位元
	FreezeQuota          *string                       `json:"freeze_quota,omitempty"`           // 冻结金额，单位元
	LeaderId             *string                       `json:"leader_id,omitempty"`              // 主管
	LeaderItemList       []BudgetCenterV2LeaderItem    `json:"leader_item_list,omitempty"`       // 主管列表
	ParentId             *string                       `json:"parent_id,omitempty"`              // 上级部门或项目的id
	OutParentId          *string                       `json:"out_parent_id,omitempty"`          // 上级部门/项目外部编码
	MemberUsed           *int32                        `json:"member_used,omitempty"`            // 使用范围，枚举值数字 0 全员可见 1 项目成员可见 2 公司主体可见
	StartDate            *string                       `json:"start_date,omitempty"`             // 项目开始日期，当类型为项目时有效 格式：yyyy-MM-dd
	ExpiryDate           *string                       `json:"expiry_date,omitempty"`            // 项目结束日期，当类型为项目时有效 格式：yyyy-MM-dd，0为长期有效
	CreateTime           *string                       `json:"create_time,omitempty"`            // 创建时间
	UpdateTime           *string                       `json:"update_time,omitempty"`            // 更新时间
	LegalEntityId        *string                       `json:"legal_entity_id,omitempty"`        // 公司主体ID，多个英文逗号分开
	OutLegalEntityId     *string                       `json:"out_legal_entity_id,omitempty"`    // 项目所属公司主体外部编码
	BudgetExtraInfo      *string                       `json:"budget_extra_info,omitempty"`      // 项目拓展字段，仅项目有数据时返回
	LimitRuleList        []BudgetCenterV2LimitRuleItem `json:"limit_rule_list,omitempty"`        // 限额规则列表
}

type BudgetCenterV2RecordBuilder struct {
	id                      string
	idSet                   bool
	name                    string
	nameSet                 bool
	companyId               string
	companyIdSet            bool
	belongEnterpriseName    string
	belongEnterpriseNameSet bool
	status                  string
	statusSet               bool
	thirdOutId              string
	thirdOutIdSet           bool
	type_                   string
	type_Set                bool
	budgetCycle             int32
	budgetCycleSet          bool
	outBudgetId             string
	outBudgetIdSet          bool
	totalQuota              string
	totalQuotaSet           bool
	isLimitQuota            int32
	isLimitQuotaSet         bool
	memberNum               int64
	memberNumSet            bool
	availableQuota          string
	availableQuotaSet       bool
	freezeQuota             string
	freezeQuotaSet          bool
	leaderId                string
	leaderIdSet             bool
	parentId                string
	parentIdSet             bool
	outParentId             string
	outParentIdSet          bool
	memberUsed              int32
	memberUsedSet           bool
	startDate               string
	startDateSet            bool
	expiryDate              string
	expiryDateSet           bool
	createTime              string
	createTimeSet           bool
	updateTime              string
	updateTimeSet           bool
	legalEntityId           string
	legalEntityIdSet        bool
	outLegalEntityId        string
	outLegalEntityIdSet     bool
	budgetExtraInfo         string
	budgetExtraInfoSet      bool
	leaderItemList          []BudgetCenterV2LeaderItem
	leaderItemListSet       bool
	limitRuleList           []BudgetCenterV2LimitRuleItem
	limitRuleListSet        bool
}

func NewBudgetCenterV2RecordBuilder() *BudgetCenterV2RecordBuilder {
	return &BudgetCenterV2RecordBuilder{}
}
func (builder *BudgetCenterV2RecordBuilder) Id(id string) *BudgetCenterV2RecordBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) Name(name string) *BudgetCenterV2RecordBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) CompanyId(companyId string) *BudgetCenterV2RecordBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) BelongEnterpriseName(belongEnterpriseName string) *BudgetCenterV2RecordBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) Status(status string) *BudgetCenterV2RecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) ThirdOutId(thirdOutId string) *BudgetCenterV2RecordBuilder {
	builder.thirdOutId = thirdOutId
	builder.thirdOutIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) Type(type_ string) *BudgetCenterV2RecordBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) BudgetCycle(budgetCycle int32) *BudgetCenterV2RecordBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) OutBudgetId(outBudgetId string) *BudgetCenterV2RecordBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) TotalQuota(totalQuota string) *BudgetCenterV2RecordBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) IsLimitQuota(isLimitQuota int32) *BudgetCenterV2RecordBuilder {
	builder.isLimitQuota = isLimitQuota
	builder.isLimitQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) MemberNum(memberNum int64) *BudgetCenterV2RecordBuilder {
	builder.memberNum = memberNum
	builder.memberNumSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) AvailableQuota(availableQuota string) *BudgetCenterV2RecordBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) FreezeQuota(freezeQuota string) *BudgetCenterV2RecordBuilder {
	builder.freezeQuota = freezeQuota
	builder.freezeQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) LeaderId(leaderId string) *BudgetCenterV2RecordBuilder {
	builder.leaderId = leaderId
	builder.leaderIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) ParentId(parentId string) *BudgetCenterV2RecordBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) OutParentId(outParentId string) *BudgetCenterV2RecordBuilder {
	builder.outParentId = outParentId
	builder.outParentIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) MemberUsed(memberUsed int32) *BudgetCenterV2RecordBuilder {
	builder.memberUsed = memberUsed
	builder.memberUsedSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) StartDate(startDate string) *BudgetCenterV2RecordBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) ExpiryDate(expiryDate string) *BudgetCenterV2RecordBuilder {
	builder.expiryDate = expiryDate
	builder.expiryDateSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) CreateTime(createTime string) *BudgetCenterV2RecordBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) UpdateTime(updateTime string) *BudgetCenterV2RecordBuilder {
	builder.updateTime = updateTime
	builder.updateTimeSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) LegalEntityId(legalEntityId string) *BudgetCenterV2RecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) OutLegalEntityId(outLegalEntityId string) *BudgetCenterV2RecordBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) BudgetExtraInfo(budgetExtraInfo string) *BudgetCenterV2RecordBuilder {
	builder.budgetExtraInfo = budgetExtraInfo
	builder.budgetExtraInfoSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) LeaderItemList(leaderItemList []BudgetCenterV2LeaderItem) *BudgetCenterV2RecordBuilder {
	builder.leaderItemList = leaderItemList
	builder.leaderItemListSet = true
	return builder
}
func (builder *BudgetCenterV2RecordBuilder) LimitRuleList(limitRuleList []BudgetCenterV2LimitRuleItem) *BudgetCenterV2RecordBuilder {
	builder.limitRuleList = limitRuleList
	builder.limitRuleListSet = true
	return builder
}

func (builder *BudgetCenterV2RecordBuilder) Build() *BudgetCenterV2Record {
	data := &BudgetCenterV2Record{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.thirdOutIdSet {
		data.ThirdOutId = &builder.thirdOutId
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.budgetCycleSet {
		data.BudgetCycle = &builder.budgetCycle
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.totalQuotaSet {
		data.TotalQuota = &builder.totalQuota
	}
	if builder.isLimitQuotaSet {
		data.IsLimitQuota = &builder.isLimitQuota
	}
	if builder.memberNumSet {
		data.MemberNum = &builder.memberNum
	}
	if builder.availableQuotaSet {
		data.AvailableQuota = &builder.availableQuota
	}
	if builder.freezeQuotaSet {
		data.FreezeQuota = &builder.freezeQuota
	}
	if builder.leaderIdSet {
		data.LeaderId = &builder.leaderId
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
	}
	if builder.outParentIdSet {
		data.OutParentId = &builder.outParentId
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
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.updateTimeSet {
		data.UpdateTime = &builder.updateTime
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.budgetExtraInfoSet {
		data.BudgetExtraInfo = &builder.budgetExtraInfo
	}
	if builder.leaderItemListSet {
		data.LeaderItemList = builder.leaderItemList
	}
	if builder.limitRuleListSet {
		data.LimitRuleList = builder.limitRuleList
	}
	return data
}
