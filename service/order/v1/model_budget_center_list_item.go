package v1

import (
	"encoding/json"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// BudgetCenterListItem 多成本中心(flight、hotel订单共用)
type BudgetCenterListItem struct {
	AppName  *string `json:"app_name,omitempty"` // 字段员工侧展示名称
	Sequence *int32  `json:"sequence,omitempty"` // 字段序号 支持从1到9
	Id       *string `json:"id,omitempty"`       // 滴滴主键，sequence为1时，生效
	Value    *string `json:"value,omitempty"`    // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	Code     *string `json:"code,omitempty"`     // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效
}

type BudgetCenterListItemBuilder struct {
	appName     string // 字段员工侧展示名称
	appNameSet  bool
	sequence    int32 // 字段序号 支持从1到9
	sequenceSet bool
	id          string // 滴滴主键，sequence为1时，生效
	idSet       bool
	value       string // 对应成本中心的值。sequence为1时与name字段一致，sequence为2时，对应extend_field_01，sequence为3时，对应extend_field_02，sequence为4时，对应extend_field_03
	valueSet    bool
	code        string // 对应成本中心的编码。sequence为1时与out_budget_id字段一致 sequence为2到9时，CODE无效
	codeSet     bool
}

func NewBudgetCenterListItemBuilder() *BudgetCenterListItemBuilder {
	return &BudgetCenterListItemBuilder{}
}
func (builder *BudgetCenterListItemBuilder) AppName(appName string) *BudgetCenterListItemBuilder {
	builder.appName = appName
	builder.appNameSet = true
	return builder
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

// UnmarshalJSON 容错反序列化：Id 字段真实流量 int/str/空串混合，
// 标准反序列化对 *string 收 number 会失败。这里用 alias 避免递归，
// Id 单独以 RawMessage 接收后统一转 string。
func (b *BudgetCenterListItem) UnmarshalJSON(data []byte) error {
	type alias BudgetCenterListItem
	aux := struct {
		alias
		Id json.RawMessage `json:"id"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*b = BudgetCenterListItem(aux.alias)
	if len(aux.Id) > 0 && string(aux.Id) != "null" {
		s, err := core.RawMessageToString(aux.Id)
		if err != nil {
			return err
		}
		b.Id = &s
	}
	return nil
}
