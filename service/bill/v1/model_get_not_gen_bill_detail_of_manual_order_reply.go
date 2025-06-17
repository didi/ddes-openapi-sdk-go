package v1

// GetNotGenBillDetailOfManualOrderReply struct for GetNotGenBillDetailOfManualOrderReply
type GetNotGenBillDetailOfManualOrderReply struct {
	Total      int64                       `json:"total"`                 // 符合查询条件的总条数
	LastId     *string                     `json:"last_id,omitempty"`     // 本次查询结果的最后一条记录的order_id
	DetailList []NotGenBDOfManualOrderItem `json:"detail_list,omitempty"` // 未出账单明细列表
}
