package v1

// FlightOrderSequenceListItem struct for FlightOrderSequenceListItem
type FlightOrderSequenceListItem struct {
	Sequenceinfo *FlightOrderSequenceInfo `json:"sequenceinfo,omitempty"`
}

type FlightOrderSequenceListItemBuilder struct {
	sequenceinfo    FlightOrderSequenceInfo
	sequenceinfoSet bool
}

func NewFlightOrderSequenceListItemBuilder() *FlightOrderSequenceListItemBuilder {
	return &FlightOrderSequenceListItemBuilder{}
}
func (builder *FlightOrderSequenceListItemBuilder) Sequenceinfo(sequenceinfo FlightOrderSequenceInfo) *FlightOrderSequenceListItemBuilder {
	builder.sequenceinfo = sequenceinfo
	builder.sequenceinfoSet = true
	return builder
}

func (builder *FlightOrderSequenceListItemBuilder) Build() *FlightOrderSequenceListItem {
	data := &FlightOrderSequenceListItem{}
	if builder.sequenceinfoSet {
		data.Sequenceinfo = &builder.sequenceinfo
	}
	return data
}
