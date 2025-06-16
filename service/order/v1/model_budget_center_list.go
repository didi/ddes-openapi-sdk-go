package v1

// BudgetCenterList 多成本中心
type BudgetCenterList struct {
	AppName  *string `json:"app_name,omitempty"` // 字段员工侧展示名称
	Sequence *int32  `json:"sequence,omitempty"` // 字段序号 支持从1到9
	Id       *string `json:"id,omitempty"`       // 滴滴主键，sequence为1时，生效
	Value    *string `json:"value,omitempty"`    // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	Code     *string `json:"code,omitempty"`     // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效。
}

type BudgetCenterListBuilder struct {
	appName     string // 字段员工侧展示名称
	appNameSet  bool
	sequence    int32 // 字段序号 支持从1到9
	sequenceSet bool
	id          string // 滴滴主键，sequence为1时，生效
	idSet       bool
	value       string // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	valueSet    bool
	code        string // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效。
	codeSet     bool
}

func NewBudgetCenterListBuilder() *BudgetCenterListBuilder {
	return &BudgetCenterListBuilder{}
}
func (builder *BudgetCenterListBuilder) AppName(appName string) *BudgetCenterListBuilder {
	builder.appName = appName
	builder.appNameSet = true
	return builder
}
func (builder *BudgetCenterListBuilder) Sequence(sequence int32) *BudgetCenterListBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}
func (builder *BudgetCenterListBuilder) Id(id string) *BudgetCenterListBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *BudgetCenterListBuilder) Value(value string) *BudgetCenterListBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}
func (builder *BudgetCenterListBuilder) Code(code string) *BudgetCenterListBuilder {
	builder.code = code
	builder.codeSet = true
	return builder
}

func (builder *BudgetCenterListBuilder) Build() *BudgetCenterList {
	data := &BudgetCenterList{}
	if builder.appNameSet {
		data.AppName = &builder.appName
	}
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
