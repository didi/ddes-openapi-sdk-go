package v1

// BudgetCenter budget_center消息,成本中心信息
type BudgetCenter struct {
	Type *int32  `json:"type,omitempty"` // 成本中心类型，枚举值数字 ，1：部门；2：项目 type为1时，code有值。为2时code和name成对有值
	Code *string `json:"code,omitempty"` // 成本中心code
	Name *string `json:"name,omitempty"` // 成本中心名称
}

type BudgetCenterBuilder struct {
	type_    int32 // 成本中心类型，枚举值数字 ，1：部门；2：项目 type为1时，code有值。为2时code和name成对有值
	type_Set bool
	code     string // 成本中心code
	codeSet  bool
	name     string // 成本中心名称
	nameSet  bool
}

func NewBudgetCenterBuilder() *BudgetCenterBuilder {
	return &BudgetCenterBuilder{}
}
func (builder *BudgetCenterBuilder) Type(type_ int32) *BudgetCenterBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BudgetCenterBuilder) Code(code string) *BudgetCenterBuilder {
	builder.code = code
	builder.codeSet = true
	return builder
}
func (builder *BudgetCenterBuilder) Name(name string) *BudgetCenterBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}

func (builder *BudgetCenterBuilder) Build() *BudgetCenter {
	data := &BudgetCenter{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.codeSet {
		data.Code = &builder.code
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	return data
}
