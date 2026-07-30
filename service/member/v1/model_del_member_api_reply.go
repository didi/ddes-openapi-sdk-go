package v1

import (
	"encoding/json"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// DelMemberApiReply struct for DelMemberApiReply
type DelMemberApiReply struct {
	Errno     *int32   `json:"errno,omitempty"`      // 错误码
	Errmsg    *string  `json:"errmsg,omitempty"`     // 错误文案
	RequestId *string  `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      []string `json:"data,omitempty"`       // 数据对象（真实流量元素 int/str 混合，UnmarshalJSON 容错）
}

// UnmarshalJSON 容错反序列化：Data 真实流量是数组，元素 int/str 混存（如 1126...、Z072967），
// 标准反序列化对 []string 收 number 元素会失败。用 alias 避免递归，Data 单独按元素转 string。
func (r *DelMemberApiReply) UnmarshalJSON(data []byte) error {
	type alias DelMemberApiReply
	aux := struct {
		alias
		Data []json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = DelMemberApiReply(aux.alias)
	if len(aux.Data) > 0 {
		r.Data = make([]string, 0, len(aux.Data))
		for _, raw := range aux.Data {
			s, _ := core.RawMessageToString(raw)
			r.Data = append(r.Data, s)
		}
	}
	return nil
}
