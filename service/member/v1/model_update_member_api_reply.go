package v1

// UpdateMemberApiReply struct for UpdateMemberApiReply
type UpdateMemberApiReply struct {
	Errno     *int32             `json:"errno,omitempty"`      // 错误码
	Errmsg    *string            `json:"errmsg,omitempty"`     // 错误文案
	RequestId *string            `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *UpdateMemberReply `json:"data,omitempty"`       // 员工信息更新失败（员工已离职或被删除）场景返回
}
