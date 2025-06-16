package v1

// ListOrderReply struct for ListOrderReply
type ListOrderReply struct {
	DomesticflightData      *DomesticFlightIdData      `json:"domesticflight_data,omitempty"`
	InternationalflightData *InternationalFlightIdData `json:"internationalflight_data,omitempty"`
	DomestichotelData       *DomesticHotelIdData       `json:"domestichotel_data,omitempty"`
	InternationalhotelData  *InternationalHotelIdData  `json:"internationalhotel_data,omitempty"`
	TrainData               *TrainData                 `json:"train_data,omitempty"`
	CarData                 *CarData                   `json:"car_data,omitempty"`
}
