package v1

// TrainOrderRcListItem struct for TrainOrderRcListItem
type TrainOrderRcListItem struct {
	RcInfo *TrainOrderRcInfo `json:"rc_info,omitempty"`
}

type TrainOrderRcListItemBuilder struct {
	rcInfo    TrainOrderRcInfo
	rcInfoSet bool
}

func NewTrainOrderRcListItemBuilder() *TrainOrderRcListItemBuilder {
	return &TrainOrderRcListItemBuilder{}
}
func (builder *TrainOrderRcListItemBuilder) RcInfo(rcInfo TrainOrderRcInfo) *TrainOrderRcListItemBuilder {
	builder.rcInfo = rcInfo
	builder.rcInfoSet = true
	return builder
}

func (builder *TrainOrderRcListItemBuilder) Build() *TrainOrderRcListItem {
	data := &TrainOrderRcListItem{}
	if builder.rcInfoSet {
		data.RcInfo = &builder.rcInfo
	}
	return data
}
