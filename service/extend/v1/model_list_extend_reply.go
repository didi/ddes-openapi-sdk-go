package v1

// ListExtendReply struct for ListExtendReply
type ListExtendReply struct {
	RootId     *string      `json:"root_id,omitempty"`     // 档案唯一标识
	RootCode   *string      `json:"root_code,omitempty"`   // 档案code
	RootName   *string      `json:"root_name,omitempty"`   // 档案名称
	RootStatus *int32       `json:"root_status,omitempty"` // 档案状态，1正常，2停用
	Children   []ExtendInfo `json:"children,omitempty"`    // 子档案列表
}
