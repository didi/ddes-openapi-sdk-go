package v1

// BudgetCenterRecord struct for BudgetCenterRecord
type BudgetCenterRecord struct {
	Id              *string      `json:"id,omitempty"`                // 滴滴内部项目或部门的ID
	Name            *string      `json:"name,omitempty"`              // 名称
	Type            *string      `json:"type,omitempty"`              // 类型，枚举值 1 部门 2 项目
	BudgetCycle     *int32       `json:"budget_cycle,omitempty"`      // 预算周期，枚举值数字 0 不限期 1 自然月 2 自然季度 3 自然年
	OutBudgetId     *string      `json:"out_budget_id,omitempty"`     // 外部成本中心id
	TotalQuota      *string      `json:"total_quota,omitempty"`       // 总金额，单位元 0 表示不限额度 精确到两位小数
	IsLimitQuota    *int32       `json:"is_limit_quota,omitempty"`    // 是否限额，枚举值数字 0 不限制 1 限制
	MemberNum       *int32       `json:"member_num,omitempty"`        // 在使用人数
	AvailableQuota  *string      `json:"available_quota,omitempty"`   // 可用金额，单位元 精确到两位小数
	FreezeQuota     *string      `json:"freeze_quota,omitempty"`      // 冻结金额，单位元 精确到两位小数
	LeaderId        *string      `json:"leader_id,omitempty"`         // 主管
	LeaderItemList  []LeaderItem `json:"leader_item_list,omitempty"`  // 主管列表
	ParentId        *string      `json:"parent_id,omitempty"`         // 上级部门或项目的id
	MemberUsed      *int32       `json:"member_used,omitempty"`       // 使用范围，枚举值数字 0 全员可见 1 项目成员可见 2 公司主体可见
	StartDate       *string      `json:"start_date,omitempty"`        // 项目开始日期，当类型为项目时，此参数有效 格式：yyyy-MM-dd
	ExpiryDate      *string      `json:"expiry_date,omitempty"`       // 项目结束日期，当类型为项目时，此参数有效 格式：yyyy-MM-dd，0为长期有效
	LegalEntityId   *string      `json:"legal_entity_id,omitempty"`   // 公司主体ID，多个英文逗号分开
	BudgetExtraInfo *string      `json:"budget_extra_info,omitempty"` // 项目拓展字段，仅项目有数据时返回
}

type BudgetCenterRecordBuilder struct {
	id                 string // 滴滴内部项目或部门的ID
	idSet              bool
	name               string // 名称
	nameSet            bool
	type_              string // 类型，枚举值 1 部门 2 项目
	type_Set           bool
	budgetCycle        int32 // 预算周期，枚举值数字 0 不限期 1 自然月 2 自然季度 3 自然年
	budgetCycleSet     bool
	outBudgetId        string // 外部成本中心id
	outBudgetIdSet     bool
	totalQuota         string // 总金额，单位元 0 表示不限额度 精确到两位小数
	totalQuotaSet      bool
	isLimitQuota       int32 // 是否限额，枚举值数字 0 不限制 1 限制
	isLimitQuotaSet    bool
	memberNum          int32 // 在使用人数
	memberNumSet       bool
	availableQuota     string // 可用金额，单位元 精确到两位小数
	availableQuotaSet  bool
	freezeQuota        string // 冻结金额，单位元 精确到两位小数
	freezeQuotaSet     bool
	leaderId           string // 主管
	leaderIdSet        bool
	leaderItemList     []LeaderItem // 主管列表
	leaderItemListSet  bool
	parentId           string // 上级部门或项目的id
	parentIdSet        bool
	memberUsed         int32 // 使用范围，枚举值数字 0 全员可见 1 项目成员可见 2 公司主体可见
	memberUsedSet      bool
	startDate          string // 项目开始日期，当类型为项目时，此参数有效 格式：yyyy-MM-dd
	startDateSet       bool
	expiryDate         string // 项目结束日期，当类型为项目时，此参数有效 格式：yyyy-MM-dd，0为长期有效
	expiryDateSet      bool
	legalEntityId      string // 公司主体ID，多个英文逗号分开
	legalEntityIdSet   bool
	budgetExtraInfo    string // 项目拓展字段，仅项目有数据时返回
	budgetExtraInfoSet bool
}

func NewBudgetCenterRecordBuilder() *BudgetCenterRecordBuilder {
	return &BudgetCenterRecordBuilder{}
}
func (builder *BudgetCenterRecordBuilder) Id(id string) *BudgetCenterRecordBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) Name(name string) *BudgetCenterRecordBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) Type(type_ string) *BudgetCenterRecordBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) BudgetCycle(budgetCycle int32) *BudgetCenterRecordBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) OutBudgetId(outBudgetId string) *BudgetCenterRecordBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) TotalQuota(totalQuota string) *BudgetCenterRecordBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) IsLimitQuota(isLimitQuota int32) *BudgetCenterRecordBuilder {
	builder.isLimitQuota = isLimitQuota
	builder.isLimitQuotaSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) MemberNum(memberNum int32) *BudgetCenterRecordBuilder {
	builder.memberNum = memberNum
	builder.memberNumSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) AvailableQuota(availableQuota string) *BudgetCenterRecordBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) FreezeQuota(freezeQuota string) *BudgetCenterRecordBuilder {
	builder.freezeQuota = freezeQuota
	builder.freezeQuotaSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) LeaderId(leaderId string) *BudgetCenterRecordBuilder {
	builder.leaderId = leaderId
	builder.leaderIdSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) LeaderItemList(leaderItemList []LeaderItem) *BudgetCenterRecordBuilder {
	builder.leaderItemList = leaderItemList
	builder.leaderItemListSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) ParentId(parentId string) *BudgetCenterRecordBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) MemberUsed(memberUsed int32) *BudgetCenterRecordBuilder {
	builder.memberUsed = memberUsed
	builder.memberUsedSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) StartDate(startDate string) *BudgetCenterRecordBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) ExpiryDate(expiryDate string) *BudgetCenterRecordBuilder {
	builder.expiryDate = expiryDate
	builder.expiryDateSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) LegalEntityId(legalEntityId string) *BudgetCenterRecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BudgetCenterRecordBuilder) BudgetExtraInfo(budgetExtraInfo string) *BudgetCenterRecordBuilder {
	builder.budgetExtraInfo = budgetExtraInfo
	builder.budgetExtraInfoSet = true
	return builder
}

func (builder *BudgetCenterRecordBuilder) Build() *BudgetCenterRecord {
	data := &BudgetCenterRecord{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
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
	if builder.leaderItemListSet {
		data.LeaderItemList = builder.leaderItemList
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
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
	return data
}
