package v1

// GetMemberDetailApiReply struct for GetMemberDetailApiReply
type GetMemberDetailApiReply struct {
	Errno     *int32                `json:"errno,omitempty"`      // 错误码
	Errmsg    *string               `json:"errmsg,omitempty"`     // 错误文案
	RequestId *string               `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *GetMemberDetailReply `json:"data,omitempty"`
}
