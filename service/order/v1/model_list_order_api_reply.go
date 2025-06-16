package v1

// ListOrderApiReply struct for ListOrderApiReply
type ListOrderApiReply struct {
	Errno     int32           `json:"errno"`      // 错误，0表示成功，非0表示失败
	Errmsg    string          `json:"errmsg"`     // 成功返回\"SUCCESS\",失败返回对应的错误信息
	RequestId string          `json:"request_id"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *ListOrderReply `json:"data,omitempty"`
}
