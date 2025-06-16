package v1

// GetLoginEncryptStrReply struct for GetLoginEncryptStrReply
type GetLoginEncryptStrReply struct {
	EncryptStr         *string               `json:"encrypt_str,omitempty"` // 跳转链接
	CallCarNow         *H5CallCarNow         `json:"callCarNow,omitempty"`  // 实时用车需要用到的参数
	CallCarAppointment *H5CallCarAppointment `json:"callCarAppointment,omitempty"`
	AirportPickUpNow   *H5AirportPickUpNow   `json:"airportPickUpNow,omitempty"`
	AirportPickUp      *H5AirportPickUp      `json:"airportPickUp,omitempty"`
	ToAirport          *H5ToAirport          `json:"toAirport,omitempty"`
	StationPickUp      *H5StationPickUp      `json:"stationPickUp,omitempty"`
	ToStation          *H5ToStation          `json:"toStation,omitempty"`
	OrderDetail        *H5OrderDetail        `json:"orderDetail,omitempty"`
	TravelList         *H5TravelList         `json:"travelList,omitempty"`
	H5RuleList         *H5RuleList           `json:"h5RuleList,omitempty"`
	H5RuleDesc         *H5RuleDesc           `json:"h5RuleDesc,omitempty"`
	H5OrderList        *H5OrderList          `json:"H5OrderList,omitempty"`
	H5Invoice          *H5Invoice            `json:"h5Invoice,omitempty"`
	H5MyWallet         *H5MyWallet           `json:"h5MyWallet,omitempty"`
	H5CallCenter       *H5CallCenter         `json:"h5CallCenter,omitempty"`
	H5HomeAddress      *H5HomeAddress        `json:"h5HomeAddress,omitempty"`
	H5CharteredCar     *H5CharteredCar       `json:"h5CharteredCar,omitempty"`
	H5Trip             *H5BTrip              `json:"h5Trip,omitempty"`
	H5BTripCategory    *H5BTripCategory      `json:"h5BTripCategory,omitempty"`
	H5BTripOrderDetail *H5BTripOrderDetail   `json:"h5BTripOrderDetail,omitempty"`
	PcLogin            *PCLogin              `json:"pcLogin,omitempty"`
}
