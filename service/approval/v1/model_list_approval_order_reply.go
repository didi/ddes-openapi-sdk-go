package v1

// ListApprovalOrderReply struct for ListApprovalOrderReply
type ListApprovalOrderReply struct {
	Records []ApprovalOrderRecord `json:"records,omitempty"` // 订单记录数组
}
