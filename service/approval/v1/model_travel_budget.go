package v1

// TravelBudget 差旅预算总额，需转换为json后赋值给travel_budget字段；
type TravelBudget struct {
	BudgetAmount *int32  `json:"budget_amount,omitempty"` // 申请单总预算金额 单位 分
	BudgetType   *int32  `json:"budget_type,omitempty"`   // 预算类型，1 申请单维度
	BudgetShare  []int32 `json:"budget_share,omitempty"`  // 需要申请单使用的制度内开启对应的品类。且最终取两者交集。使用数字，预算可使用品类：1 市内用车 2 接送场景 301 国内酒店 401 火车票 601 国内机票 701 国际机票 801 国际酒店
}

type TravelBudgetBuilder struct {
	budgetAmount    int32 // 申请单总预算金额 单位 分
	budgetAmountSet bool
	budgetType      int32 // 预算类型，1 申请单维度
	budgetTypeSet   bool
	budgetShare     []int32 // 需要申请单使用的制度内开启对应的品类。且最终取两者交集。使用数字，预算可使用品类：1 市内用车 2 接送场景 301 国内酒店 401 火车票 601 国内机票 701 国际机票 801 国际酒店
	budgetShareSet  bool
}

func NewTravelBudgetBuilder() *TravelBudgetBuilder {
	return &TravelBudgetBuilder{}
}
func (builder *TravelBudgetBuilder) BudgetAmount(budgetAmount int32) *TravelBudgetBuilder {
	builder.budgetAmount = budgetAmount
	builder.budgetAmountSet = true
	return builder
}
func (builder *TravelBudgetBuilder) BudgetType(budgetType int32) *TravelBudgetBuilder {
	builder.budgetType = budgetType
	builder.budgetTypeSet = true
	return builder
}
func (builder *TravelBudgetBuilder) BudgetShare(budgetShare []int32) *TravelBudgetBuilder {
	builder.budgetShare = budgetShare
	builder.budgetShareSet = true
	return builder
}

func (builder *TravelBudgetBuilder) Build() *TravelBudget {
	data := &TravelBudget{}
	if builder.budgetAmountSet {
		data.BudgetAmount = &builder.budgetAmount
	}
	if builder.budgetTypeSet {
		data.BudgetType = &builder.budgetType
	}
	if builder.budgetShareSet {
		data.BudgetShare = builder.budgetShare
	}
	return data
}
