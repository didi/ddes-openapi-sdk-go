package v1

// ListBillReply struct for ListBillReply
type ListBillReply struct {
	Total   int32        `json:"total"`   // 记录总数
	Records []BillRecord `json:"records"` // 账单记录列表
	LastId  int32        `json:"lastId"`  // 本次查询结果的最后一条记录ID
}
