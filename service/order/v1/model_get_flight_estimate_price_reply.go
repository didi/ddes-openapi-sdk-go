package v1

// GetFlightEstimatePriceReply struct for GetFlightEstimatePriceReply
type GetFlightEstimatePriceReply struct {
	FlightList []FlightList `json:"flight_list"` // 航班列表
}
