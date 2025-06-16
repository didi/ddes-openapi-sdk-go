package v1

// BudgetCenterListItem 多成本中心Item项
type BudgetCenterListItem struct {
	Sequence *int32  `json:"sequence,omitempty"` // 字段序号 支持从1到9
	Id       *string `json:"id,omitempty"`       // sequence为1时，生效
	Value    *string `json:"value,omitempty"`    // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	Code     *string `json:"code,omitempty"`     // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效。
}

type BudgetCenterListItemBuilder struct {
	sequence    int32 // 字段序号 支持从1到9
	sequenceSet bool
	id          string // sequence为1时，生效
	idSet       bool
	value       string // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	valueSet    bool
	code        string // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效。
	codeSet     bool
}

func NewBudgetCenterListItemBuilder() *BudgetCenterListItemBuilder {
	return &BudgetCenterListItemBuilder{}
}
func (builder *BudgetCenterListItemBuilder) Sequence(sequence int32) *BudgetCenterListItemBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}
func (builder *BudgetCenterListItemBuilder) Id(id string) *BudgetCenterListItemBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *BudgetCenterListItemBuilder) Value(value string) *BudgetCenterListItemBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}
func (builder *BudgetCenterListItemBuilder) Code(code string) *BudgetCenterListItemBuilder {
	builder.code = code
	builder.codeSet = true
	return builder
}

func (builder *BudgetCenterListItemBuilder) Build() *BudgetCenterListItem {
	data := &BudgetCenterListItem{}
	if builder.sequenceSet {
		data.Sequence = &builder.sequence
	}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	if builder.codeSet {
		data.Code = &builder.code
	}
	return data
}
