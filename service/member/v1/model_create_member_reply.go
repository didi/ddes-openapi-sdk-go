package v1

import (
	"encoding/json"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// CreateMemberReply struct for CreateMemberReply
type CreateMemberReply struct {
	Id       int64   `json:"id"`                  // 员工在滴滴侧的 ID（新增员工成功场景返回）
	Phone    *string `json:"phone,omitempty"`     // 员工手机号（新增成功、添加失败（员工已存在）场景返回）
	MemberId *string `json:"member_id,omitempty"` // 员工滴滴侧ID（添加员工失败-员工已存在场景返回）
	Status   *int32  `json:"status,omitempty"`    // 员工状态（仅添加失败-员工已存在场景返回），枚举值数字：1-正常，4-离职，6-未绑定手机号
}

type CreateMemberReplyBuilder struct {
	id          int64 // 员工在滴滴侧的 ID
	idSet       bool
	phone       string // 员工手机号
	phoneSet    bool
	memberId    string // 员工滴滴侧ID
	memberIdSet bool
	status      int32 // 员工状态，枚举值数字：1-正常，4-离职，6-未绑定手机号
	statusSet   bool
}

func NewCreateMemberReplyBuilder() *CreateMemberReplyBuilder {
	return &CreateMemberReplyBuilder{}
}
func (builder *CreateMemberReplyBuilder) Id(id int64) *CreateMemberReplyBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *CreateMemberReplyBuilder) Phone(phone string) *CreateMemberReplyBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *CreateMemberReplyBuilder) MemberId(memberId string) *CreateMemberReplyBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *CreateMemberReplyBuilder) Status(status int32) *CreateMemberReplyBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}

func (builder *CreateMemberReplyBuilder) Build() *CreateMemberReply {
	data := &CreateMemberReply{}
	if builder.idSet {
		data.Id = builder.id
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	return data
}

// UnmarshalJSON 容错反序列化：MemberId 真实流量可能为 number（文档 int64）或 string，
// 标准反序列化对 *string 收 number 会失败。alias 避免递归，MemberId 以 RawMessage 接收后统一转 string。
func (c *CreateMemberReply) UnmarshalJSON(data []byte) error {
	type alias CreateMemberReply
	aux := struct {
		alias
		MemberId json.RawMessage `json:"member_id"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*c = CreateMemberReply(aux.alias)
	if len(aux.MemberId) > 0 && string(aux.MemberId) != "null" {
		s, err := core.RawMessageToString(aux.MemberId)
		if err != nil {
			return err
		}
		c.MemberId = &s
	}
	return nil
}
