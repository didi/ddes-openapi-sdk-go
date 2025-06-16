package v1

// GetPersonalReceiptOrderReply struct for GetPersonalReceiptOrderReply
type GetPersonalReceiptOrderReply struct {
	Total   int32                           `json:"total"`             // 申请应用时分配的AppKey
	Records []GetPersonalReceiptOrderRecord `json:"records,omitempty"` // 审批订单列表
}
