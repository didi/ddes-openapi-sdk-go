package v1

// ExtendFieldList 扩展信息list
type ExtendFieldList struct {
	ExtendField01 *string `json:"extend_field_01,omitempty"` // 扩展信息一
	ExtendField02 *string `json:"extend_field_02,omitempty"` // 扩展信息二
	ExtendField03 *string `json:"extend_field_03,omitempty"` // 扩展信息三
}

type ExtendFieldListBuilder struct {
	extendField01    string // 扩展信息一
	extendField01Set bool
	extendField02    string // 扩展信息二
	extendField02Set bool
	extendField03    string // 扩展信息三
	extendField03Set bool
}

func NewExtendFieldListBuilder() *ExtendFieldListBuilder {
	return &ExtendFieldListBuilder{}
}
func (builder *ExtendFieldListBuilder) ExtendField01(extendField01 string) *ExtendFieldListBuilder {
	builder.extendField01 = extendField01
	builder.extendField01Set = true
	return builder
}
func (builder *ExtendFieldListBuilder) ExtendField02(extendField02 string) *ExtendFieldListBuilder {
	builder.extendField02 = extendField02
	builder.extendField02Set = true
	return builder
}
func (builder *ExtendFieldListBuilder) ExtendField03(extendField03 string) *ExtendFieldListBuilder {
	builder.extendField03 = extendField03
	builder.extendField03Set = true
	return builder
}

func (builder *ExtendFieldListBuilder) Build() *ExtendFieldList {
	data := &ExtendFieldList{}
	if builder.extendField01Set {
		data.ExtendField01 = &builder.extendField01
	}
	if builder.extendField02Set {
		data.ExtendField02 = &builder.extendField02
	}
	if builder.extendField03Set {
		data.ExtendField03 = &builder.extendField03
	}
	return data
}
