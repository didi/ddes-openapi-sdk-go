package v1

// CurAppRover 当前审批人
type CurAppRover struct {
	Type  *string `json:"type,omitempty"`  // 人员类型主键；英文枚举：employee_number ，email  ，phone
	Value *string `json:"value,omitempty"` // 人员主键对应的值
}

type CurAppRoverBuilder struct {
	type_    string // 人员类型主键；英文枚举：employee_number ，email  ，phone
	type_Set bool
	value    string // 人员主键对应的值
	valueSet bool
}

func NewCurAppRoverBuilder() *CurAppRoverBuilder {
	return &CurAppRoverBuilder{}
}
func (builder *CurAppRoverBuilder) Type(type_ string) *CurAppRoverBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *CurAppRoverBuilder) Value(value string) *CurAppRoverBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}

func (builder *CurAppRoverBuilder) Build() *CurAppRover {
	data := &CurAppRover{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	return data
}
