package v1

import "github.com/didi/ddes-openapi-sdk-go/core"

// BudgetCenterV2LimitRuleItem 集团账户 v2 限额规则信息（v2 契约独立类型，禁止与 /river 版 LimitRuleItem 复用）
type BudgetCenterV2LimitRuleItem struct {
	RuleName             *string  `json:"rule_name,omitempty"`              // 限额规则名称(员工侧展示名称)
	BudgetCycle          *int32   `json:"budget_cycle,omitempty"`           // 预算周期：0=不限额，1=自然月，2=自然季度，3=自然年，4=一次性，5=自然日，6=自定义，7=使用原部门周期
	IsAccumulative       *int32   `json:"is_accumulative,omitempty"`        // 是否累计：0=不可累计，1=可累计
	TotalQuota           *float64 `json:"total_quota,omitempty"`            // 限额(元)，0表示不限额度，精确到两位小数
	LimitManagementScope *int32   `json:"limit_management_scope,omitempty"` // 限额管理范围：0=对当前部门/项目生效，1=对当前和下级生效
	AvailableQuota       *string  `json:"available_quota,omitempty"`        // 剩余额度(元)，精确到两位小数
	FreezeQuota          *string  `json:"freeze_quota,omitempty"`           // 冻结金额(元)，精确到两位小数
}

type BudgetCenterV2LimitRuleItemBuilder struct {
	ruleName                string // 限额规则名称(员工侧展示名称)
	ruleNameSet             bool
	budgetCycle             int32 // 预算周期：0=不限额，1=自然月，2=自然季度，3=自然年，4=一次性，5=自然日，6=自定义，7=使用原部门周期
	budgetCycleSet          bool
	isAccumulative          int32 // 是否累计：0=不可累计，1=可累计
	isAccumulativeSet       bool
	totalQuota              float64 // 限额(元)，0表示不限额度，精确到两位小数
	totalQuotaSet           bool
	limitManagementScope    int32 // 限额管理范围：0=对当前部门/项目生效，1=对当前和下级生效
	limitManagementScopeSet bool
	availableQuota          string // 剩余额度(元)，精确到两位小数
	availableQuotaSet       bool
	freezeQuota             string // 冻结金额(元)，精确到两位小数
	freezeQuotaSet          bool
}

func NewBudgetCenterV2LimitRuleItemBuilder() *BudgetCenterV2LimitRuleItemBuilder {
	return &BudgetCenterV2LimitRuleItemBuilder{}
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) RuleName(ruleName string) *BudgetCenterV2LimitRuleItemBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) BudgetCycle(budgetCycle int32) *BudgetCenterV2LimitRuleItemBuilder {
	builder.budgetCycle = budgetCycle
	builder.budgetCycleSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) IsAccumulative(isAccumulative int32) *BudgetCenterV2LimitRuleItemBuilder {
	builder.isAccumulative = isAccumulative
	builder.isAccumulativeSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) TotalQuota(totalQuota float64) *BudgetCenterV2LimitRuleItemBuilder {
	builder.totalQuota = totalQuota
	builder.totalQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) LimitManagementScope(limitManagementScope int32) *BudgetCenterV2LimitRuleItemBuilder {
	builder.limitManagementScope = limitManagementScope
	builder.limitManagementScopeSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) AvailableQuota(availableQuota string) *BudgetCenterV2LimitRuleItemBuilder {
	builder.availableQuota = availableQuota
	builder.availableQuotaSet = true
	return builder
}
func (builder *BudgetCenterV2LimitRuleItemBuilder) FreezeQuota(freezeQuota string) *BudgetCenterV2LimitRuleItemBuilder {
	builder.freezeQuota = freezeQuota
	builder.freezeQuotaSet = true
	return builder
}

func (builder *BudgetCenterV2LimitRuleItemBuilder) Build() *BudgetCenterV2LimitRuleItem {
	data := &BudgetCenterV2LimitRuleItem{}
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

// UnmarshalJSON 容错反序列化：AvailableQuota 等字段真实流量 number/string 混存，
// *string 收 number 会失败。用 core.SmartDecode 经中间 map 转换，避免递归。
func (l *BudgetCenterV2LimitRuleItem) UnmarshalJSON(data []byte) error {
	return core.SmartDecode(data, l)
}
