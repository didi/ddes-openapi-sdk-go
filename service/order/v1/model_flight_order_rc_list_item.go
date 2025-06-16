package v1

// FlightOrderRcListItem struct for FlightOrderRcListItem
type FlightOrderRcListItem struct {
	RcInfo *FlightOrderRcInfo `json:"rc_info,omitempty"`
}

type FlightOrderRcListItemBuilder struct {
	rcInfo    FlightOrderRcInfo
	rcInfoSet bool
}

func NewFlightOrderRcListItemBuilder() *FlightOrderRcListItemBuilder {
	return &FlightOrderRcListItemBuilder{}
}
func (builder *FlightOrderRcListItemBuilder) RcInfo(rcInfo FlightOrderRcInfo) *FlightOrderRcListItemBuilder {
	builder.rcInfo = rcInfo
	builder.rcInfoSet = true
	return builder
}

func (builder *FlightOrderRcListItemBuilder) Build() *FlightOrderRcListItem {
	data := &FlightOrderRcListItem{}
	if builder.rcInfoSet {
		data.RcInfo = &builder.rcInfo
	}
	return data
}
