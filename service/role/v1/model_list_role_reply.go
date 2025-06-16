package v1

// ListRoleReply struct for ListRoleReply
type ListRoleReply struct {
	Id    *string `json:"id,omitempty"`    // 角色ID
	Name  *string `json:"name,omitempty"`  // 角色中文名称
	Alias *string `json:"alias,omitempty"` // 角色英文别名
}
