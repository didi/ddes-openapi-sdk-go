package v1

import (
	"encoding/json"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// UpdateMemberReply struct for UpdateMemberReply
type UpdateMemberReply struct {
	MemberId *string `json:"member_id,omitempty"` // 员工滴滴侧ID（员工信息更新失败-员工已离职或被删除场景返回）
	Phone    *string `json:"phone,omitempty"`     // 员工手机号（员工信息更新失败场景返回）
	Status   *int32  `json:"status,omitempty"`    // 员工状态（仅员工信息更新失败场景返回），枚举值数字：1-正常，4-离职，6-未绑定手机号
}

type UpdateMemberReplyBuilder struct {
	memberId    string // 员工滴滴侧ID
	memberIdSet bool
	phone       string // 员工手机号
	phoneSet    bool
	status      int32 // 员工状态，枚举值数字：1-正常，4-离职，6-未绑定手机号
	statusSet   bool
}

func NewUpdateMemberReplyBuilder() *UpdateMemberReplyBuilder {
	return &UpdateMemberReplyBuilder{}
}
func (builder *UpdateMemberReplyBuilder) MemberId(memberId string) *UpdateMemberReplyBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *UpdateMemberReplyBuilder) Phone(phone string) *UpdateMemberReplyBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *UpdateMemberReplyBuilder) Status(status int32) *UpdateMemberReplyBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}

func (builder *UpdateMemberReplyBuilder) Build() *UpdateMemberReply {
	data := &UpdateMemberReply{}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	return data
}

// UnmarshalJSON 容错反序列化：MemberId 真实流量可能为 number（文档 int64）或 string，
// 标准反序列化对 *string 收 number 会失败。alias 避免递归，MemberId 以 RawMessage 接收后统一转 string。
func (u *UpdateMemberReply) UnmarshalJSON(data []byte) error {
	type alias UpdateMemberReply
	aux := struct {
		alias
		MemberId json.RawMessage `json:"member_id"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*u = UpdateMemberReply(aux.alias)
	if len(aux.MemberId) > 0 && string(aux.MemberId) != "null" {
		s, err := core.RawMessageToString(aux.MemberId)
		if err != nil {
			return err
		}
		u.MemberId = &s
	}
	return nil
}
