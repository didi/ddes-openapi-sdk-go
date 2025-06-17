package v1

// Approval approval消息,申请单信息
type Approval struct {
	Type  *int32  `json:"type,omitempty"`  // 申请单号类型，枚举值数字，1 out_approval_id（外部申请单单号），2 approval_id（滴滴内部申请单号）
	Value *string `json:"value,omitempty"` // 单号信息
}

type ApprovalBuilder struct {
	type_    int32 // 申请单号类型，枚举值数字，1 out_approval_id（外部申请单单号），2 approval_id（滴滴内部申请单号）
	type_Set bool
	value    string // 单号信息
	valueSet bool
}

func NewApprovalBuilder() *ApprovalBuilder {
	return &ApprovalBuilder{}
}
func (builder *ApprovalBuilder) Type(type_ int32) *ApprovalBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *ApprovalBuilder) Value(value string) *ApprovalBuilder {
	builder.value = value
	builder.valueSet = true
	return builder
}

func (builder *ApprovalBuilder) Build() *Approval {
	data := &Approval{}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.valueSet {
		data.Value = &builder.value
	}
	return data
}
