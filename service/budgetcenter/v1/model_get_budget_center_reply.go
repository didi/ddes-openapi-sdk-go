package v1

// GetBudgetCenterReply struct for GetBudgetCenterReply
type GetBudgetCenterReply struct {
	Total   string               `json:"total"`   // 总数据条数
	Records []BudgetCenterRecord `json:"records"` // 部门或者项目集合
}
