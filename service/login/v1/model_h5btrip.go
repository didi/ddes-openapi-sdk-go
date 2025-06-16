package v1

// H5BTrip 商旅H5页面
type H5BTrip struct {
	ApprovalId *string `json:"approval_id,omitempty"` // 差旅、申请用车，必传
	OrderId    *string `json:"order_id,omitempty"`    // 订单id
	JumpPage   *string `json:"jump_page,omitempty"`   // 跳转页orderDetail:订单详情页 注：此处的jump_page是小写的， 和网约车页面跳转的jumpPage 的区别；
}

type H5BTripBuilder struct {
	approvalId    string // 差旅、申请用车，必传
	approvalIdSet bool
	orderId       string // 订单id
	orderIdSet    bool
	jumpPage      string // 跳转页orderDetail:订单详情页 注：此处的jump_page是小写的， 和网约车页面跳转的jumpPage 的区别；
	jumpPageSet   bool
}

func NewH5BTripBuilder() *H5BTripBuilder {
	return &H5BTripBuilder{}
}
func (builder *H5BTripBuilder) ApprovalId(approvalId string) *H5BTripBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5BTripBuilder) OrderId(orderId string) *H5BTripBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *H5BTripBuilder) JumpPage(jumpPage string) *H5BTripBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5BTripBuilder) Build() *H5BTrip {
	data := &H5BTrip{}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
