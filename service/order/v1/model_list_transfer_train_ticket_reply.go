package v1

// ListTransferTrainTicketReply struct for ListTransferTrainTicketReply
type ListTransferTrainTicketReply struct {
	List            []TransferTrainTicketListItem `json:"list,omitempty"`     // 车次详细信息
	HasMore         *int32                        `json:"has_more,omitempty"` // 是否有下一页，非0则有下一页
	SelectCondition *SelectCondition              `json:"select_condition,omitempty"`
	Trace           *string                       `json:"trace,omitempty"`     // 数据请求唯一标识，系统生成
	CurPage         *int32                        `json:"cur_page,omitempty"`  // 当前页码，默认：1，必须大于等于1
	PageSize        *int32                        `json:"page_size,omitempty"` // 每页车次条数，默认：10
}
