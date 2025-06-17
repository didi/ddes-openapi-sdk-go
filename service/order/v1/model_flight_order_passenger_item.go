package v1

// FlightOrderPassengerItem struct for FlightOrderPassengerItem
type FlightOrderPassengerItem struct {
	PassengerInfo *FlightOrderPassengerInfo `json:"passenger_info,omitempty"`
}

type FlightOrderPassengerItemBuilder struct {
	passengerInfo    FlightOrderPassengerInfo
	passengerInfoSet bool
}

func NewFlightOrderPassengerItemBuilder() *FlightOrderPassengerItemBuilder {
	return &FlightOrderPassengerItemBuilder{}
}
func (builder *FlightOrderPassengerItemBuilder) PassengerInfo(passengerInfo FlightOrderPassengerInfo) *FlightOrderPassengerItemBuilder {
	builder.passengerInfo = passengerInfo
	builder.passengerInfoSet = true
	return builder
}

func (builder *FlightOrderPassengerItemBuilder) Build() *FlightOrderPassengerItem {
	data := &FlightOrderPassengerItem{}
	if builder.passengerInfoSet {
		data.PassengerInfo = &builder.passengerInfo
	}
	return data
}
