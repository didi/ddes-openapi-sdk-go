package v1

// GetHotelOrderDetailReply struct for GetHotelOrderDetailReply
type GetHotelOrderDetailReply struct {
	Total                  int32                   `json:"total"` // 符合查询条件的数据总数
	DomestichotelData      *DomesticHotelData      `json:"domestichotel_data,omitempty"`
	InternationalhotelData *InternationalHotelData `json:"internationalhotel_data,omitempty"`
}
