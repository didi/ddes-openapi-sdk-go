package v1

// DelMemberApiReply 删除项目与人员关系响应
type DelMemberApiReply struct {
	Errno     int32          `json:"errno"`      // 错误码
	Errmsg    string         `json:"errmsg"`     // 错误文案
	RequestId string         `json:"request_id"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      DelMemberReply `json:"data"`       // 响应数据：type=1时为null，type=2+member_values时可能包含success_data/error_data等
}
