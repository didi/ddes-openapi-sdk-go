package v1

// InterFlightOrderChangeListItem struct for InterFlightOrderChangeListItem
type InterFlightOrderChangeListItem struct {
	ChangeInfo *InterFlightOrderChangeInfo `json:"change_info,omitempty"`
}

type InterFlightOrderChangeListItemBuilder struct {
	changeInfo    InterFlightOrderChangeInfo
	changeInfoSet bool
}

func NewInterFlightOrderChangeListItemBuilder() *InterFlightOrderChangeListItemBuilder {
	return &InterFlightOrderChangeListItemBuilder{}
}
func (builder *InterFlightOrderChangeListItemBuilder) ChangeInfo(changeInfo InterFlightOrderChangeInfo) *InterFlightOrderChangeListItemBuilder {
	builder.changeInfo = changeInfo
	builder.changeInfoSet = true
	return builder
}

func (builder *InterFlightOrderChangeListItemBuilder) Build() *InterFlightOrderChangeListItem {
	data := &InterFlightOrderChangeListItem{}
	if builder.changeInfoSet {
		data.ChangeInfo = &builder.changeInfo
	}
	return data
}
