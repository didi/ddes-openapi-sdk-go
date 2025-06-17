package v1

// InvoiceInfo invoice_info消息,发票信息
type InvoiceInfo struct {
	IsInvoice *int32 `json:"is_invoice,omitempty"` // 开票状态，枚举值数字：1：开过 0：未开 仅用车生效
}

type InvoiceInfoBuilder struct {
	isInvoice    int32 // 开票状态，枚举值数字：1：开过 0：未开 仅用车生效
	isInvoiceSet bool
}

func NewInvoiceInfoBuilder() *InvoiceInfoBuilder {
	return &InvoiceInfoBuilder{}
}
func (builder *InvoiceInfoBuilder) IsInvoice(isInvoice int32) *InvoiceInfoBuilder {
	builder.isInvoice = isInvoice
	builder.isInvoiceSet = true
	return builder
}

func (builder *InvoiceInfoBuilder) Build() *InvoiceInfo {
	data := &InvoiceInfo{}
	if builder.isInvoiceSet {
		data.IsInvoice = &builder.isInvoice
	}
	return data
}
