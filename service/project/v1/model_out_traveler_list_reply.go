package v1

// OutTravelerListReply struct for OutTravelerListReply
type OutTravelerListReply struct {
	OutTravelers []*OutTravelerInfo `json:"out_travelers,omitempty"` // 外部出行人列表
	Total        *int64             `json:"total,omitempty"`         // 此次查询符合条件的人员总数
}
