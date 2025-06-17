package v1

// GetApprovalDetailApiReply 响应消息
type GetApprovalDetailApiReply struct {
	Errno     int32          `json:"errno"`      // 错误码
	Errmsg    string         `json:"errmsg"`     // 错误文案
	RequestId string         `json:"request_id"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      ApprovalDetail `json:"data"`
}
