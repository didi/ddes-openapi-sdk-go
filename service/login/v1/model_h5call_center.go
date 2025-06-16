package v1

// H5CallCenter 客服中心
type H5CallCenter struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为: callCenter
}

type H5CallCenterBuilder struct {
	jumpPage    string // 跳转地址为: callCenter
	jumpPageSet bool
}

func NewH5CallCenterBuilder() *H5CallCenterBuilder {
	return &H5CallCenterBuilder{}
}
func (builder *H5CallCenterBuilder) JumpPage(jumpPage string) *H5CallCenterBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5CallCenterBuilder) Build() *H5CallCenter {
	data := &H5CallCenter{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
