package v1

// Booker booker消息,预订人信息
type Booker struct {
	Type  *string `json:"type,omitempty"`  // 预订人主键类型，枚举值英文，employee_number ，email  ，phone
	Value *string `json:"value,omitempty"` // 查询主键信息
}

type BookerBuilder struct {
	type_    string // 预订人主键类型，枚举值英文，employee_number ，email  ，phone
	type_Set bool
	value    string // 查询主键信息
	valueSet bool
}

func NewBookerBuilder() *BookerBuilder {
	return &BookerBuilder{}
}
func (builder *BookerBuilder) Type(type_ string) *BookerBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BookerBuilder) Value(value string) *BookerBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}

func (builder *BookerBuilder) Build() *Booker {
	data := &Booker{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	return data
}
