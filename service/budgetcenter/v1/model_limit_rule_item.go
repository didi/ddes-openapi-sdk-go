package v1

// LimitRuleItem struct for LimitRuleItem
type LimitRuleItem struct {
	RuleName             *string `json:"rule_name,omitempty"`              // 限额规则名称(员工侧展示名称)
	BudgetCycle          *int32  `json:"budget_cycle,omitempty"`           // 预算周期：0=不限额，1=自然月，2=自然季度，3=自然年，4=一次性，5=自然日，6=自定义，7=使用原部门周期
	IsAccumulative       *int32  `json:"is_accumulative,omitempty"`        // 是否累计：0=不可累计，1=可累计
	TotalQuota           *string `json:"total_quota,omitempty"`            // 限额(元)，0表示不限额度，精确到两位小数
	LimitManagementScope *int32  `json:"limit_management_scope,omitempty"` // 限额管理范围：0=对当前部门/项目生效，1=对当前和下级生效
	AvailableQuota       *string `json:"available_quota,omitempty"`        // 剩余额度(元)，精确到两位小数
	FreezeQuota          *string `json:"freeze_quota,omitempty"`           // 冻结金额(元)，精确到两位小数
}

type LimitRuleItemBuilder struct {
	ruleName                string // 限额规则名称(员工侧展示名称)
	ruleNameSet             bool
	budgetCycle             int32 // 预算周期：0=不限额，1=自然月，2=自然季度，3=自然年，4=一次性，5=自然日，6=自定义，7=使用原部门周期
	budgetCycleSet          bool
	isAccumulative          int32 // 是否累计：0=不可累计，1=可累计
	isAccumulativeSet       bool
	totalQuota              string // 限额(元)，0表示不限额度，精确到两位小数
	totalQuotaSet           bool
	limitManagementScope    int32 // 限额管理范围：0=对当前部门/项目生效，1=对当前和下级生效
	limitManagementScopeSet bool
	availableQuota          string // 剩余额度(元)，精确到两位小数
	availableQuotaSet       bool
	freezeQuota             string // 冻结金额(元)，精确到两位小数
	freezeQuotaSet          bool
}

func NewLimitRuleItemBuilder() *LimitRuleItemBuilder {
	return &LimitRuleItemBuilder{}
}
func (builder *LimitRuleItemBuilder) RuleName(ruleName string) *LimitRuleItemBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) BudgetCycle(budgetCycle int32) *LimitRuleItemBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) IsAccumulative(isAccumulative int32) *LimitRuleItemBuilder {
	builder.isAccumulative = isAccumulative
	builder.isAccumulativeSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) TotalQuota(totalQuota string) *LimitRuleItemBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) LimitManagementScope(limitManagementScope int32) *LimitRuleItemBuilder {
	builder.limitManagementScope = limitManagementScope
	builder.limitManagementScopeSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) AvailableQuota(availableQuota string) *LimitRuleItemBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *LimitRuleItemBuilder) FreezeQuota(freezeQuota string) *LimitRuleItemBuilder {
	builder.freezeQuota = freezeQuota
	builder.freezeQuotaSet = true
	return builder
}

func (builder *LimitRuleItemBuilder) Build() *LimitRuleItem {
	data := &LimitRuleItem{}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.budgetCycleSet {
		data.BudgetCycle = &builder.budgetCycle
	}
	if builder.isAccumulativeSet {
		data.IsAccumulative = &builder.isAccumulative
	}
	if builder.totalQuotaSet {
		data.TotalQuota = &builder.totalQuota
	}
	if builder.limitManagementScopeSet {
		data.LimitManagementScope = &builder.limitManagementScope
	}
	if builder.availableQuotaSet {
		data.AvailableQuota = &builder.availableQuota
	}
	if builder.freezeQuotaSet {
		data.FreezeQuota = &builder.freezeQuota
	}
	return data
}
