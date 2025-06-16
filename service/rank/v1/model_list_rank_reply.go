package v1

// ListRankReply struct for ListRankReply
type ListRankReply struct {
	Total   *int32     `json:"total,omitempty"`   // 职级集合中的条数
	Records []RankInfo `json:"records,omitempty"` // 职级数据集合
}
