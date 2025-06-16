package v1

// ListMemberReply struct for ListMemberReply
type ListMemberReply struct {
	Total   *int32         `json:"total,omitempty"`   // 此次查询符合条件的员工总数
	Records []MemberRecord `json:"records,omitempty"` // 员工集合
}
