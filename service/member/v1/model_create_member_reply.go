package v1

// CreateMemberReply struct for CreateMemberReply
type CreateMemberReply struct {
	Id    int64   `json:"id"`              // 员工在滴滴侧的 ID
	Phone *string `json:"phone,omitempty"` // 员工手机号
}
