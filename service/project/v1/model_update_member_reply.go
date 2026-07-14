package v1

// UpdateMemberReply 绑定项目与人员关系响应数据
type UpdateMemberReply struct {
	SuccessData         []string                `json:"success_data,omitempty"`          // 成功绑定的员工ID数组
	ErrorData           []UpdateMemberErrorInfo `json:"error_data,omitempty"`            // 失败的员工信息数组
	SuccessMemberValues []string                `json:"success_member_values,omitempty"` // 成功绑定的员工值数组
}

// UpdateMemberErrorInfo 绑定失败的员工信息
type UpdateMemberErrorInfo struct {
	ErrorMsg          *string  `json:"error_msg,omitempty"`           // 错误原因
	ErrorMemberIds    []string `json:"error_member_ids,omitempty"`    // 失败的员工ID数组
	ErrorMemberValues []string `json:"error_member_values,omitempty"` // 失败的员工值数组
}

type UpdateMemberErrorInfoBuilder struct {
	errorMsg             string
	errorMsgSet          bool
	errorMemberIds       []string
	errorMemberIdsSet    bool
	errorMemberValues    []string
	errorMemberValuesSet bool
}

func NewUpdateMemberErrorInfoBuilder() *UpdateMemberErrorInfoBuilder {
	return &UpdateMemberErrorInfoBuilder{}
}
func (builder *UpdateMemberErrorInfoBuilder) ErrorMsg(errorMsg string) *UpdateMemberErrorInfoBuilder {
	builder.errorMsg = errorMsg
	builder.errorMsgSet = true
	return builder
}
func (builder *UpdateMemberErrorInfoBuilder) ErrorMemberIds(errorMemberIds []string) *UpdateMemberErrorInfoBuilder {
	builder.errorMemberIds = errorMemberIds
	builder.errorMemberIdsSet = true
	return builder
}
func (builder *UpdateMemberErrorInfoBuilder) ErrorMemberValues(errorMemberValues []string) *UpdateMemberErrorInfoBuilder {
	builder.errorMemberValues = errorMemberValues
	builder.errorMemberValuesSet = true
	return builder
}

func (builder *UpdateMemberErrorInfoBuilder) Build() *UpdateMemberErrorInfo {
	data := &UpdateMemberErrorInfo{}
	if builder.errorMsgSet {
		data.ErrorMsg = &builder.errorMsg
	}
	if builder.errorMemberIdsSet {
		data.ErrorMemberIds = builder.errorMemberIds
	}
	if builder.errorMemberValuesSet {
		data.ErrorMemberValues = builder.errorMemberValues
	}
	return data
}
