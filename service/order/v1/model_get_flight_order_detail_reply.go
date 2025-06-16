package v1

// GetFlightOrderDetailReply struct for GetFlightOrderDetailReply
type GetFlightOrderDetailReply struct {
	Total                   int32                    `json:"total"` // 符合查询条件的数据总数
	DomesticflightData      *DomesticFlightData      `json:"domesticflight_data,omitempty"`
	InternationalflightData *InternationalFlightData `json:"internationalflight_data,omitempty"`
}
