package v1

// H5HomeAddress 常用地址设置
type H5HomeAddress struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为: homeAddress
}

type H5HomeAddressBuilder struct {
	jumpPage    string // 跳转地址为: homeAddress
	jumpPageSet bool
}

func NewH5HomeAddressBuilder() *H5HomeAddressBuilder {
	return &H5HomeAddressBuilder{}
}
func (builder *H5HomeAddressBuilder) JumpPage(jumpPage string) *H5HomeAddressBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5HomeAddressBuilder) Build() *H5HomeAddress {
	data := &H5HomeAddress{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
