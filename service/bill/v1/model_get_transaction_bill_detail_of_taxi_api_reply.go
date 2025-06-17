package v1

// GetTransactionBillDetailOfTaxiApiReply struct for GetTransactionBillDetailOfTaxiApiReply
type GetTransactionBillDetailOfTaxiApiReply struct {
	Errno     int32                                `json:"errno"`      // 错误码
	Errmsg    string                               `json:"errmsg"`     // 错误文案
	RequestId string                               `json:"request_id"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *GetTransactionBillDetailOfTaxiReply `json:"data,omitempty"`
}
