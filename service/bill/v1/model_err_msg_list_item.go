package v1

// ErrMsgListItem struct for ErrMsgListItem
type ErrMsgListItem struct {
	SubBusinessType *int32  `json:"sub_business_type,omitempty"` // 提示信息品类
	SubOrderId      *string `json:"sub_order_id,omitempty"`      // 提示信息哪一单提交不通过
	AdjustErr       *string `json:"adjust_err,omitempty"`        // 该单不通过的原因
}

type ErrMsgListItemBuilder struct {
	subBusinessType    int32 // 提示信息品类
	subBusinessTypeSet bool
	subOrderId         string // 提示信息哪一单提交不通过
	subOrderIdSet      bool
	adjustErr          string // 该单不通过的原因
	adjustErrSet       bool
}

func NewErrMsgListItemBuilder() *ErrMsgListItemBuilder {
	return &ErrMsgListItemBuilder{}
}
func (builder *ErrMsgListItemBuilder) SubBusinessType(subBusinessType int32) *ErrMsgListItemBuilder {
	builder.subBusinessType = subBusinessType
	builder.subBusinessTypeSet = true
	return builder
}
func (builder *ErrMsgListItemBuilder) SubOrderId(subOrderId string) *ErrMsgListItemBuilder {
	builder.subOrderId = subOrderId
	builder.subOrderIdSet = true
	return builder
}
func (builder *ErrMsgListItemBuilder) AdjustErr(adjustErr string) *ErrMsgListItemBuilder {
	builder.adjustErr = adjustErr
	builder.adjustErrSet = true
	return builder
}

func (builder *ErrMsgListItemBuilder) Build() *ErrMsgListItem {
	data := &ErrMsgListItem{}
	if builder.subBusinessTypeSet {
		data.SubBusinessType = &builder.subBusinessType
	}
	if builder.subOrderIdSet {
		data.SubOrderId = &builder.subOrderId
	}
	if builder.adjustErrSet {
		data.AdjustErr = &builder.adjustErr
	}
	return data
}
