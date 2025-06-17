package v1

// GetOrderReply struct for GetOrderReply
type GetOrderReply struct {
	Total   *int32        `json:"total,omitempty"`   // 此次查询符合条件的订单总数
	Records []OrderRecord `json:"records,omitempty"` // 订单集合
}
