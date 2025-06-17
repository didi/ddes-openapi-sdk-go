package v1

// GetBillDetailOfDomesticHotelReply struct for GetBillDetailOfDomesticHotelReply
type GetBillDetailOfDomesticHotelReply struct {
	Total  *int64                          `json:"total,omitempty"`   // 总条数
	IsLast *bool                           `json:"is_last,omitempty"` // 最后一页标记字段，是否最后一页 true：是 false：否
	LastId *string                         `json:"last_id,omitempty"` // 本次查询结果的最后一条记录的order_id
	Orders []BillDetailOfDomesticHotelItem `json:"orders,omitempty"`  // 账单明细列表(国内酒店)
}
