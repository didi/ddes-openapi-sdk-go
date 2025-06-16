package v1

// GetLegalEntityApiReply 查询公司主体响应
type GetLegalEntityApiReply struct {
	Errno     *int32               `json:"errno,omitempty"`      // 错误码
	Errmsg    *string              `json:"errmsg,omitempty"`     // 错误文案
	RequestId *string              `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *GetLegalEntityReply `json:"data,omitempty"`
}
