package v1

// TrainOrderChangeListItem struct for TrainOrderChangeListItem
type TrainOrderChangeListItem struct {
	ChangeInfo *TrainOrderChangeInfo `json:"change_info,omitempty"`
}

type TrainOrderChangeListItemBuilder struct {
	changeInfo    TrainOrderChangeInfo
	changeInfoSet bool
}

func NewTrainOrderChangeListItemBuilder() *TrainOrderChangeListItemBuilder {
	return &TrainOrderChangeListItemBuilder{}
}
func (builder *TrainOrderChangeListItemBuilder) ChangeInfo(changeInfo TrainOrderChangeInfo) *TrainOrderChangeListItemBuilder {
	builder.changeInfo = changeInfo
	builder.changeInfoSet = true
	return builder
}

func (builder *TrainOrderChangeListItemBuilder) Build() *TrainOrderChangeListItem {
	data := &TrainOrderChangeListItem{}
	if builder.changeInfoSet {
		data.ChangeInfo = &builder.changeInfo
	}
	return data
}
