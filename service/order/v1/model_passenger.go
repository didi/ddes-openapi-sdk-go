package v1

// Passenger passenger消息,出行人信息
type Passenger struct {
	Type  *string `json:"type,omitempty"`  // 出行人主键类型，枚举值英文，employee_number ，email  ，phone，外部出行人只能使用手机号
	Value *string `json:"value,omitempty"` // 查询主键信息
}

type PassengerBuilder struct {
	type_    string // 出行人主键类型，枚举值英文，employee_number ，email  ，phone，外部出行人只能使用手机号
	type_Set bool
	value    string // 查询主键信息
	valueSet bool
}

func NewPassengerBuilder() *PassengerBuilder {
	return &PassengerBuilder{}
}
func (builder *PassengerBuilder) Type(type_ string) *PassengerBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *PassengerBuilder) Value(value string) *PassengerBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}

func (builder *PassengerBuilder) Build() *Passenger {
	data := &Passenger{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	return data
}
