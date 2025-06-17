package v1

// GetNotGenBillDetailOfInterFlightReply struct for GetNotGenBillDetailOfInterFlightReply
type GetNotGenBillDetailOfInterFlightReply struct {
	Total      int64                          `json:"total"`                 // 符合查询条件的总条数
	LastId     *string                        `json:"last_id,omitempty"`     // 本次查询结果的最后一条记录的order_id
	DetailList []NotGenBDOfOverseasFlightItem `json:"detail_list,omitempty"` // 未出账单明细列表
}
