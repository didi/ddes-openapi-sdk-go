package v1

// PriceInfo 费用明细信息
type PriceInfo struct {
	Name   *string `json:"name,omitempty"`   // 费用的名称
	Amount *string `json:"amount,omitempty"` // 费用的金额，单位为元
	Type   *string `json:"type,omitempty"`   // 费用的类型
}

type PriceInfoBuilder struct {
	name      string // 费用的名称
	nameSet   bool
	amount    string // 费用的金额，单位为元
	amountSet bool
	type_     string // 费用的类型
	type_Set  bool
}

func NewPriceInfoBuilder() *PriceInfoBuilder {
	return &PriceInfoBuilder{}
}
func (builder *PriceInfoBuilder) Name(name string) *PriceInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *PriceInfoBuilder) Amount(amount string) *PriceInfoBuilder {
	builder.amount = amount
	builder.amountSet = true
	return builder
}
func (builder *PriceInfoBuilder) Type(type_ string) *PriceInfoBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}

func (builder *PriceInfoBuilder) Build() *PriceInfo {
	data := &PriceInfo{}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.amountSet {
		data.Amount = &builder.amount
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	return data
}
