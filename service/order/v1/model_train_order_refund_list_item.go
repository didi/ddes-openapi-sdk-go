package v1

// TrainOrderRefundListItem struct for TrainOrderRefundListItem
type TrainOrderRefundListItem struct {
	RefundInfo *TrainOrderRefundInfo `json:"refund_info,omitempty"`
}

type TrainOrderRefundListItemBuilder struct {
	refundInfo    TrainOrderRefundInfo
	refundInfoSet bool
}

func NewTrainOrderRefundListItemBuilder() *TrainOrderRefundListItemBuilder {
	return &TrainOrderRefundListItemBuilder{}
}
func (builder *TrainOrderRefundListItemBuilder) RefundInfo(refundInfo TrainOrderRefundInfo) *TrainOrderRefundListItemBuilder {
	builder.refundInfo = refundInfo
	builder.refundInfoSet = true
	return builder
}

func (builder *TrainOrderRefundListItemBuilder) Build() *TrainOrderRefundListItem {
	data := &TrainOrderRefundListItem{}
	if builder.refundInfoSet {
		data.RefundInfo = &builder.refundInfo
	}
	return data
}
