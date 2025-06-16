package v1

// DomesticFlightOrderRefundListItem struct for DomesticFlightOrderRefundListItem
type DomesticFlightOrderRefundListItem struct {
	RefundInfo *DomesticFlightOrderRefundInfo `json:"refund_info,omitempty"`
}

type DomesticFlightOrderRefundListItemBuilder struct {
	refundInfo    DomesticFlightOrderRefundInfo
	refundInfoSet bool
}

func NewDomesticFlightOrderRefundListItemBuilder() *DomesticFlightOrderRefundListItemBuilder {
	return &DomesticFlightOrderRefundListItemBuilder{}
}
func (builder *DomesticFlightOrderRefundListItemBuilder) RefundInfo(refundInfo DomesticFlightOrderRefundInfo) *DomesticFlightOrderRefundListItemBuilder {
	builder.refundInfo = refundInfo
	builder.refundInfoSet = true
	return builder
}

func (builder *DomesticFlightOrderRefundListItemBuilder) Build() *DomesticFlightOrderRefundListItem {
	data := &DomesticFlightOrderRefundListItem{}
	if builder.refundInfoSet {
		data.RefundInfo = &builder.refundInfo
	}
	return data
}
