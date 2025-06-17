package v1

// H5MyWallet 我的钱包
type H5MyWallet struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为: mywallet
}

type H5MyWalletBuilder struct {
	jumpPage    string // 跳转地址为: mywallet
	jumpPageSet bool
}

func NewH5MyWalletBuilder() *H5MyWalletBuilder {
	return &H5MyWalletBuilder{}
}
func (builder *H5MyWalletBuilder) JumpPage(jumpPage string) *H5MyWalletBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5MyWalletBuilder) Build() *H5MyWallet {
	data := &H5MyWallet{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
