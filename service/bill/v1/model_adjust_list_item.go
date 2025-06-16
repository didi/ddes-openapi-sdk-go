package v1

// AdjustListItem 账单调整 ～ 列表项目
type AdjustListItem struct {
	SubOrderId      *string       `json:"sub_order_id,omitempty"`      // 子订单号, 国内/国际机票/火车票：票id, 国内/海外酒店/增值服务：子订单号, 网约车/代驾/出租车：订单号;
	SubBusinessType *int32        `json:"sub_business_type,omitempty"` // 子品类,网约车：1 出租车：100 代驾：40 国内酒店：201 国内机票：202 火车票：203 国际酒店：204 国际机票：205 增值服务：531(还不支持,未来扩展)
	AdjustFields    *AdjustFields `json:"adjust_fields,omitempty"`
}

type AdjustListItemBuilder struct {
	subOrderId         string // 子订单号, 国内/国际机票/火车票：票id, 国内/海外酒店/增值服务：子订单号, 网约车/代驾/出租车：订单号;
	subOrderIdSet      bool
	subBusinessType    int32 // 子品类,网约车：1 出租车：100 代驾：40 国内酒店：201 国内机票：202 火车票：203 国际酒店：204 国际机票：205 增值服务：531(还不支持,未来扩展)
	subBusinessTypeSet bool
	adjustFields       AdjustFields
	adjustFieldsSet    bool
}

func NewAdjustListItemBuilder() *AdjustListItemBuilder {
	return &AdjustListItemBuilder{}
}
func (builder *AdjustListItemBuilder) SubOrderId(subOrderId string) *AdjustListItemBuilder {
	builder.subOrderId = subOrderId
	builder.subOrderIdSet = true
	return builder
}
func (builder *AdjustListItemBuilder) SubBusinessType(subBusinessType int32) *AdjustListItemBuilder {
	builder.subBusinessType = subBusinessType
	builder.subBusinessTypeSet = true
	return builder
}
func (builder *AdjustListItemBuilder) AdjustFields(adjustFields AdjustFields) *AdjustListItemBuilder {
	builder.adjustFields = adjustFields
	builder.adjustFieldsSet = true
	return builder
}

func (builder *AdjustListItemBuilder) Build() *AdjustListItem {
	data := &AdjustListItem{}
	if builder.subOrderIdSet {
		data.SubOrderId = &builder.subOrderId
	}
	if builder.subBusinessTypeSet {
		data.SubBusinessType = &builder.subBusinessType
	}
	if builder.adjustFieldsSet {
		data.AdjustFields = &builder.adjustFields
	}
	return data
}
