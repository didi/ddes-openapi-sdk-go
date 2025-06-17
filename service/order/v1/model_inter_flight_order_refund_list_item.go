package v1

// InterFlightOrderRefundListItem struct for InterFlightOrderRefundListItem
type InterFlightOrderRefundListItem struct {
	RefundInfo *InterFlightOrderRefundInfo `json:"refund_info,omitempty"`
}

type InterFlightOrderRefundListItemBuilder struct {
	refundInfo    InterFlightOrderRefundInfo
	refundInfoSet bool
}

func NewInterFlightOrderRefundListItemBuilder() *InterFlightOrderRefundListItemBuilder {
	return &InterFlightOrderRefundListItemBuilder{}
}
func (builder *InterFlightOrderRefundListItemBuilder) RefundInfo(refundInfo InterFlightOrderRefundInfo) *InterFlightOrderRefundListItemBuilder {
	builder.refundInfo = refundInfo
	builder.refundInfoSet = true
	return builder
}

func (builder *InterFlightOrderRefundListItemBuilder) Build() *InterFlightOrderRefundListItem {
	data := &InterFlightOrderRefundListItem{}
	if builder.refundInfoSet {
		data.RefundInfo = &builder.refundInfo
	}
	return data
}
