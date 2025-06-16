package v1

// ErrorListItem struct for ErrorListItem
type ErrorListItem struct {
	Data   *ExtendInfo `json:"data,omitempty"`
	ErrMsg *string     `json:"err_msg,omitempty"` // 错误信息
}

type ErrorListItemBuilder struct {
	data      ExtendInfo
	dataSet   bool
	errMsg    string // 错误信息
	errMsgSet bool
}

func NewErrorListItemBuilder() *ErrorListItemBuilder {
	return &ErrorListItemBuilder{}
}
func (builder *ErrorListItemBuilder) Data(data ExtendInfo) *ErrorListItemBuilder {
	builder.data = data
	builder.dataSet = true
	return builder
}
func (builder *ErrorListItemBuilder) ErrMsg(errMsg string) *ErrorListItemBuilder {
	builder.errMsg = errMsg
	builder.errMsgSet = true
	return builder
}

func (builder *ErrorListItemBuilder) Build() *ErrorListItem {
	data := &ErrorListItem{}
	if builder.dataSet {
		data.Data = &builder.data
	}
	if builder.errMsgSet {
		data.ErrMsg = &builder.errMsg
	}
	return data
}
