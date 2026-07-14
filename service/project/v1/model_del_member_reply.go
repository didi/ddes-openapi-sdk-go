package v1

// DelMemberReply 删除项目与人员关系响应数据
type DelMemberReply struct {
	SuccessData         []string             `json:"success_data,omitempty"`          // 成功删除的员工ID数组
	ErrorData           []DelMemberErrorInfo `json:"error_data,omitempty"`            // 失败的员工信息数组
	SuccessMemberValues []string             `json:"success_member_values,omitempty"` // 成功删除的员工值数组
}

// DelMemberErrorInfo 删除失败的员工信息
type DelMemberErrorInfo struct {
	ErrorMsg          *string  `json:"error_msg,omitempty"`           // 错误原因
	ErrorMemberIds    []string `json:"error_member_ids,omitempty"`    // 失败的员工ID数组
	ErrorMemberValues []string `json:"error_member_values,omitempty"` // 失败的员工值数组
}

type DelMemberErrorInfoBuilder struct {
	errorMsg             string
	errorMsgSet          bool
	errorMemberIds       []string
	errorMemberIdsSet    bool
	errorMemberValues    []string
	errorMemberValuesSet bool
}

func NewDelMemberErrorInfoBuilder() *DelMemberErrorInfoBuilder {
	return &DelMemberErrorInfoBuilder{}
}
func (builder *DelMemberErrorInfoBuilder) ErrorMsg(errorMsg string) *DelMemberErrorInfoBuilder {
	builder.errorMsg = errorMsg
	builder.errorMsgSet = true
	return builder
}
func (builder *DelMemberErrorInfoBuilder) ErrorMemberIds(errorMemberIds []string) *DelMemberErrorInfoBuilder {
	builder.errorMemberIds = errorMemberIds
	builder.errorMemberIdsSet = true
	return builder
}
func (builder *DelMemberErrorInfoBuilder) ErrorMemberValues(errorMemberValues []string) *DelMemberErrorInfoBuilder {
	builder.errorMemberValues = errorMemberValues
	builder.errorMemberValuesSet = true
	return builder
}

func (builder *DelMemberErrorInfoBuilder) Build() *DelMemberErrorInfo {
	data := &DelMemberErrorInfo{}
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
