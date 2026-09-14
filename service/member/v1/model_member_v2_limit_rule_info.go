package v1

// MemberV2LimitRuleInfo 集团账户 v2 员工限额规则信息（v2 契约独立类型，禁止与 /river 版 LimitRuleInfo 复用）
type MemberV2LimitRuleInfo struct {
	RuleName             *string  `json:"rule_name,omitempty"`              // 限额规则名称
	BudgetCycle          *int32   `json:"budget_cycle,omitempty"`           // 预算周期，枚举值数字：0 不限额 1 自然月 2 自然季度 3 自然年 4 一次性 5 自然日 6 自定义 7 使用原部门的周期
	IsAccumulative       *int32   `json:"is_accumulative,omitempty"`        // 是否累计，枚举值数字：0 不可累计 1 可累计
	TotalQuota           *float64 `json:"total_quota,omitempty"`            // 限额，单位元，精确到两位小数，0表示不限额度
	AvailableQuota       *float64 `json:"available_quota,omitempty"`        // 剩余额度，单位元，精确到两位小数
	FreezeQuota          *float64 `json:"freeze_quota,omitempty"`           // 冻结金额，单位元，精确到两位小数
	LimitManagementScope *int32   `json:"limit_management_scope,omitempty"` // 限额管理范围
}

type MemberV2LimitRuleInfoBuilder struct {
	ruleName                string // 限额规则名称
	ruleNameSet             bool
	budgetCycle             int32 // 预算周期，枚举值数字：0 不限额 1 自然月 2 自然季度 3 自然年 4 一次性 5 自然日 6 自定义 7 使用原部门的周期
	budgetCycleSet          bool
	isAccumulative          int32 // 是否累计，枚举值数字：0 不可累计 1 可累计
	isAccumulativeSet       bool
	totalQuota              float64 // 限额，单位元，精确到两位小数，0表示不限额度
	totalQuotaSet           bool
	availableQuota          float64 // 剩余额度，单位元，精确到两位小数
	availableQuotaSet       bool
	freezeQuota             float64 // 冻结金额，单位元，精确到两位小数
	freezeQuotaSet          bool
	limitManagementScope    int32 // 限额管理范围
	limitManagementScopeSet bool
}

func NewMemberV2LimitRuleInfoBuilder() *MemberV2LimitRuleInfoBuilder {
	return &MemberV2LimitRuleInfoBuilder{}
}
func (builder *MemberV2LimitRuleInfoBuilder) RuleName(ruleName string) *MemberV2LimitRuleInfoBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) BudgetCycle(budgetCycle int32) *MemberV2LimitRuleInfoBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) IsAccumulative(isAccumulative int32) *MemberV2LimitRuleInfoBuilder {
	builder.isAccumulative = isAccumulative
	builder.isAccumulativeSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) TotalQuota(totalQuota float64) *MemberV2LimitRuleInfoBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) AvailableQuota(availableQuota float64) *MemberV2LimitRuleInfoBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) FreezeQuota(freezeQuota float64) *MemberV2LimitRuleInfoBuilder {
	builder.freezeQuota = freezeQuota
	builder.freezeQuotaSet = true
	return builder
}
func (builder *MemberV2LimitRuleInfoBuilder) LimitManagementScope(limitManagementScope int32) *MemberV2LimitRuleInfoBuilder {
	builder.limitManagementScope = limitManagementScope
	builder.limitManagementScopeSet = true
	return builder
}

func (builder *MemberV2LimitRuleInfoBuilder) Build() *MemberV2LimitRuleInfo {
	data := &MemberV2LimitRuleInfo{}
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
	if builder.availableQuotaSet {
		data.AvailableQuota = &builder.availableQuota
	}
	if builder.freezeQuotaSet {
		data.FreezeQuota = &builder.freezeQuota
	}
	if builder.limitManagementScopeSet {
		data.LimitManagementScope = &builder.limitManagementScope
	}
	return data
}
