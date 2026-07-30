package v1

import (
	"encoding/json"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// GetBillDetailOfWangYCReply struct for GetBillDetailOfWangYCReply
type GetBillDetailOfWangYCReply struct {
	Total  *int64                 `json:"total,omitempty"`   // 总条数
	IsLast *bool                  `json:"is_last,omitempty"` // 最后一页标记字段，是否最后一页 true：是 false：否
	LastId *string                `json:"last_id,omitempty"` // 本次查询结果的最后一条记录的order_id
	Orders []BillListItemOfWangYC `json:"orders,omitempty"`  // 账单明细列表
}

// UnmarshalJSON 容错反序列化：LastId 真实流量 int/str/非数字混存，
// 标准反序列化对 *string 收 number 会失败。用 alias 避免递归，LastId 单独转 string。
func (r *GetBillDetailOfWangYCReply) UnmarshalJSON(data []byte) error {
	type alias GetBillDetailOfWangYCReply
	aux := struct {
		alias
		LastId json.RawMessage `json:"last_id"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = GetBillDetailOfWangYCReply(aux.alias)
	if len(aux.LastId) > 0 && string(aux.LastId) != "null" {
		s, err := core.RawMessageToString(aux.LastId)
		if err != nil {
			return err
		}
		r.LastId = &s
	}
	return nil
}
