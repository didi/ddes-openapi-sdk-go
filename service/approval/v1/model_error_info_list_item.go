package v1

// ErrorInfoListItem struct for ErrorInfoListItem
type ErrorInfoListItem struct {
	ErrNo  *int32  `json:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty"`
}

type ErrorInfoListItemBuilder struct {
	errNo     int32
	errNoSet  bool
	errMsg    string
	errMsgSet bool
}

func NewErrorInfoListItemBuilder() *ErrorInfoListItemBuilder {
	return &ErrorInfoListItemBuilder{}
}
func (builder *ErrorInfoListItemBuilder) ErrNo(errNo int32) *ErrorInfoListItemBuilder {
	builder.errNo = errNo
	builder.errNoSet = true
	return builder
}
func (builder *ErrorInfoListItemBuilder) ErrMsg(errMsg string) *ErrorInfoListItemBuilder {
	builder.errMsg = errMsg
	builder.errMsgSet = true
	return builder
}

func (builder *ErrorInfoListItemBuilder) Build() *ErrorInfoListItem {
	data := &ErrorInfoListItem{}
	if builder.errNoSet {
		data.ErrNo = &builder.errNo
	}
	if builder.errMsgSet {
		data.ErrMsg = &builder.errMsg
	}
	return data
}
