package v1

// GetTransactionBillDetailOfTaxiReply struct for GetTransactionBillDetailOfTaxiReply
type GetTransactionBillDetailOfTaxiReply struct {
	Total           *int64                  `json:"total,omitempty"`            // 符合查询条件的总条数
	LastId          *int64                  `json:"last_id,omitempty"`          // 本次查询最后一条数据的偏移量
	TransactionList []TransactionOfTaxiItem `json:"transaction_list,omitempty"` // 交易列表
}
