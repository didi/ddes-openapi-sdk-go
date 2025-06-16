package v1

// H5Invoice 发票与报销
type H5Invoice struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为invoice
}

type H5InvoiceBuilder struct {
	jumpPage    string // 跳转地址为invoice
	jumpPageSet bool
}

func NewH5InvoiceBuilder() *H5InvoiceBuilder {
	return &H5InvoiceBuilder{}
}
func (builder *H5InvoiceBuilder) JumpPage(jumpPage string) *H5InvoiceBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5InvoiceBuilder) Build() *H5Invoice {
	data := &H5Invoice{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
