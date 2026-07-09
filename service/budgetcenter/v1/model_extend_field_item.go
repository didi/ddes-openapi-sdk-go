package v1

// ExtendFieldItem struct for ExtendFieldItem
type ExtendFieldItem struct {
	Id    *int32  `json:"id,omitempty"`    // 扩展字段ID
	Code  *string `json:"code,omitempty"`  // 扩展字段编码
	Value *string `json:"value,omitempty"` // 扩展字段值
}

type ExtendFieldItemBuilder struct {
	id       int32 // 扩展字段ID
	idSet    bool
	code     string // 扩展字段编码
	codeSet  bool
	value    string // 扩展字段值
	valueSet bool
}

func NewExtendFieldItemBuilder() *ExtendFieldItemBuilder {
	return &ExtendFieldItemBuilder{}
}
func (builder *ExtendFieldItemBuilder) Id(id int32) *ExtendFieldItemBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *ExtendFieldItemBuilder) Code(code string) *ExtendFieldItemBuilder {
	builder.code = code
	builder.codeSet = true
	return builder
}
func (builder *ExtendFieldItemBuilder) Value(value string) *ExtendFieldItemBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}

func (builder *ExtendFieldItemBuilder) Build() *ExtendFieldItem {
	data := &ExtendFieldItem{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.codeSet {
		data.Code = &builder.code
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	return data
}
