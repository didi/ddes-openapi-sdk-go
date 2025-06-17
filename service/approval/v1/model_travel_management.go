package v1

// TravelManagement 差旅管控 -> travel_management
type TravelManagement struct {
	DailyAmountControl []DailyAmountControl `json:"daily_amount_control,omitempty"` // 每日金额管控内容
}

type TravelManagementBuilder struct {
	dailyAmountControl    []DailyAmountControl // 每日金额管控内容
	dailyAmountControlSet bool
}

func NewTravelManagementBuilder() *TravelManagementBuilder {
	return &TravelManagementBuilder{}
}
func (builder *TravelManagementBuilder) DailyAmountControl(dailyAmountControl []DailyAmountControl) *TravelManagementBuilder {
	builder.dailyAmountControl = dailyAmountControl
	builder.dailyAmountControlSet = true
	return builder
}

func (builder *TravelManagementBuilder) Build() *TravelManagement {
	data := &TravelManagement{}
	if builder.dailyAmountControlSet {
		data.DailyAmountControl = builder.dailyAmountControl
	}
	return data
}
