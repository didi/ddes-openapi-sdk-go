package v1

// H5OrderList 订单列表
type H5OrderList struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为orderList
}

type H5OrderListBuilder struct {
	jumpPage    string // 跳转地址为orderList
	jumpPageSet bool
}

func NewH5OrderListBuilder() *H5OrderListBuilder {
	return &H5OrderListBuilder{}
}
func (builder *H5OrderListBuilder) JumpPage(jumpPage string) *H5OrderListBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5OrderListBuilder) Build() *H5OrderList {
	data := &H5OrderList{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
