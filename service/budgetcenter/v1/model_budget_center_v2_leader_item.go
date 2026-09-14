package v1

// BudgetCenterV2LeaderItem 集团账户 v2 部门/项目主管信息（v2 契约独立类型，禁止与 /river 版 LeaderItem 复用）
type BudgetCenterV2LeaderItem struct {
	LeaderId   *string `json:"leader_id,omitempty"`   // 主管ID
	LeaderName *string `json:"leader_name,omitempty"` // 主管姓名
	LeaderType *string `json:"leader_type,omitempty"` // 主管类型，枚举英文：major 主要主管 other 其他主管
}

type BudgetCenterV2LeaderItemBuilder struct {
	leaderId      string // 主管ID
	leaderIdSet   bool
	leaderName    string // 主管姓名
	leaderNameSet bool
	leaderType    string // 主管类型，枚举英文：major 主要主管 other 其他主管
	leaderTypeSet bool
}

func NewBudgetCenterV2LeaderItemBuilder() *BudgetCenterV2LeaderItemBuilder {
	return &BudgetCenterV2LeaderItemBuilder{}
}
func (builder *BudgetCenterV2LeaderItemBuilder) LeaderId(leaderId string) *BudgetCenterV2LeaderItemBuilder {
	builder.leaderId = leaderId
	builder.leaderIdSet = true
	return builder
}
func (builder *BudgetCenterV2LeaderItemBuilder) LeaderName(leaderName string) *BudgetCenterV2LeaderItemBuilder {
	builder.leaderName = leaderName
	builder.leaderNameSet = true
	return builder
}
func (builder *BudgetCenterV2LeaderItemBuilder) LeaderType(leaderType string) *BudgetCenterV2LeaderItemBuilder {
	builder.leaderType = leaderType
	builder.leaderTypeSet = true
	return builder
}

func (builder *BudgetCenterV2LeaderItemBuilder) Build() *BudgetCenterV2LeaderItem {
	data := &BudgetCenterV2LeaderItem{}
	if builder.leaderIdSet {
		data.LeaderId = &builder.leaderId
	}
	if builder.leaderNameSet {
		data.LeaderName = &builder.leaderName
	}
	if builder.leaderTypeSet {
		data.LeaderType = &builder.leaderType
	}
	return data
}
