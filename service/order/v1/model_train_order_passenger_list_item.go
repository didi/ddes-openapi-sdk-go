package v1

// TrainOrderPassengerListItem struct for TrainOrderPassengerListItem
type TrainOrderPassengerListItem struct {
	PassengerInfo *TrainOrderPassengerInfo `json:"passenger_info,omitempty"`
}

type TrainOrderPassengerListItemBuilder struct {
	passengerInfo    TrainOrderPassengerInfo
	passengerInfoSet bool
}

func NewTrainOrderPassengerListItemBuilder() *TrainOrderPassengerListItemBuilder {
	return &TrainOrderPassengerListItemBuilder{}
}
func (builder *TrainOrderPassengerListItemBuilder) PassengerInfo(passengerInfo TrainOrderPassengerInfo) *TrainOrderPassengerListItemBuilder {
	builder.passengerInfo = passengerInfo
	builder.passengerInfoSet = true
	return builder
}

func (builder *TrainOrderPassengerListItemBuilder) Build() *TrainOrderPassengerListItem {
	data := &TrainOrderPassengerListItem{}
	if builder.passengerInfoSet {
		data.PassengerInfo = &builder.passengerInfo
	}
	return data
}
