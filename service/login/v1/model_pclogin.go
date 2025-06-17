package v1

// PCLogin PC单点说明
type PCLogin struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转页面-默认bill; bill 单据页index 首页;trip 商旅pc首页;tripHotel 商旅pc首页 - 选中酒店tab;tripFlight 商旅pc首页 - 选中机票tab;tripTrain 商旅pc首页 - 选中火车票tab。注意：当没有跳转没有权限时管理后台会跳转404商旅会跳到敬请期待的兜底页面
}

type PCLoginBuilder struct {
	jumpPage    string // 跳转页面-默认bill; bill 单据页index 首页;trip 商旅pc首页;tripHotel 商旅pc首页 - 选中酒店tab;tripFlight 商旅pc首页 - 选中机票tab;tripTrain 商旅pc首页 - 选中火车票tab。注意：当没有跳转没有权限时管理后台会跳转404商旅会跳到敬请期待的兜底页面
	jumpPageSet bool
}

func NewPCLoginBuilder() *PCLoginBuilder {
	return &PCLoginBuilder{}
}
func (builder *PCLoginBuilder) JumpPage(jumpPage string) *PCLoginBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *PCLoginBuilder) Build() *PCLogin {
	data := &PCLogin{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
