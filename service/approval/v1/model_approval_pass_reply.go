package v1

// ApprovalPassReply struct for ApprovalPassReply
type ApprovalPassReply struct {
	ApprovalId    *string             `json:"approval_id,omitempty"` // 审批单号
	ErrorInfoList []ErrorInfoListItem `json:"error_info_list,omitempty"`
}
