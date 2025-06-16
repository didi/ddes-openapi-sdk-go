package v1

// DailyAmountControl 每日金额管控内容
type DailyAmountControl struct {
	DailyAmount    *int32  `json:"daily_amount,omitempty"`    // 每人日金额管控
	ControlType    *int32  `json:"control_type,omitempty"`    // 管控方式1 日清
	ControlProduct []int32 `json:"control_product,omitempty"` // 1 市内用车
}

type DailyAmountControlBuilder struct {
	dailyAmount       int32 // 每人日金额管控
	dailyAmountSet    bool
	controlType       int32 // 管控方式1 日清
	controlTypeSet    bool
	controlProduct    []int32 // 1 市内用车
	controlProductSet bool
}

func NewDailyAmountControlBuilder() *DailyAmountControlBuilder {
	return &DailyAmountControlBuilder{}
}
func (builder *DailyAmountControlBuilder) DailyAmount(dailyAmount int32) *DailyAmountControlBuilder {
	builder.dailyAmount = dailyAmount
	builder.dailyAmountSet = true
	return builder
}
func (builder *DailyAmountControlBuilder) ControlType(controlType int32) *DailyAmountControlBuilder {
	builder.controlType = controlType
	builder.controlTypeSet = true
	return builder
}
func (builder *DailyAmountControlBuilder) ControlProduct(controlProduct []int32) *DailyAmountControlBuilder {
	builder.controlProduct = controlProduct
	builder.controlProductSet = true
	return builder
}

func (builder *DailyAmountControlBuilder) Build() *DailyAmountControl {
	data := &DailyAmountControl{}
	if builder.dailyAmountSet {
		data.DailyAmount = &builder.dailyAmount
	}
	if builder.controlTypeSet {
		data.ControlType = &builder.controlType
	}
	if builder.controlProductSet {
		data.ControlProduct = builder.controlProduct
	}
	return data
}
