package v1

// GetPersonalReceiptOrderApiReply struct for GetPersonalReceiptOrderApiReply
type GetPersonalReceiptOrderApiReply struct {
	Errno     int32                         `json:"errno"`                // 数字 0 表示成功，非0 表示失败
	Errmsg    string                        `json:"errmsg"`               // 错误文案
	RequestId *string                       `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *GetPersonalReceiptOrderReply `json:"data,omitempty"`
}
