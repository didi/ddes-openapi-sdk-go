package v1

// PayInfo pay_info消息,支付信息
type PayInfo struct {
	PayType *int32 `json:"pay_type,omitempty"` // 支付方式，枚举值数字：0企业支付，1个人垫付，2混合支付 机票暂不支持混付
}

type PayInfoBuilder struct {
	payType    int32 // 支付方式，枚举值数字：0企业支付，1个人垫付，2混合支付 机票暂不支持混付
	payTypeSet bool
}

func NewPayInfoBuilder() *PayInfoBuilder {
	return &PayInfoBuilder{}
}
func (builder *PayInfoBuilder) PayType(payType int32) *PayInfoBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}

func (builder *PayInfoBuilder) Build() *PayInfo {
	data := &PayInfo{}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	return data
}
