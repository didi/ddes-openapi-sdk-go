package v1

// DomesticFlightOrderChangeListItem struct for DomesticFlightOrderChangeListItem
type DomesticFlightOrderChangeListItem struct {
	ChangeInfo *DomesticFlightOrderChangeInfo `json:"change_info,omitempty"`
}

type DomesticFlightOrderChangeListItemBuilder struct {
	changeInfo    DomesticFlightOrderChangeInfo
	changeInfoSet bool
}

func NewDomesticFlightOrderChangeListItemBuilder() *DomesticFlightOrderChangeListItemBuilder {
	return &DomesticFlightOrderChangeListItemBuilder{}
}
func (builder *DomesticFlightOrderChangeListItemBuilder) ChangeInfo(changeInfo DomesticFlightOrderChangeInfo) *DomesticFlightOrderChangeListItemBuilder {
	builder.changeInfo = changeInfo
	builder.changeInfoSet = true
	return builder
}

func (builder *DomesticFlightOrderChangeListItemBuilder) Build() *DomesticFlightOrderChangeListItem {
	data := &DomesticFlightOrderChangeListItem{}
	if builder.changeInfoSet {
		data.ChangeInfo = &builder.changeInfo
	}
	return data
}
