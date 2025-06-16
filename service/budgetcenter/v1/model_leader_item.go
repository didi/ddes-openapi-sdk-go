package v1

// LeaderItem struct for LeaderItem
type LeaderItem struct {
	LeaderId   *string `json:"leader_id,omitempty"`   // 主管ID
	LeaderName *string `json:"leader_name,omitempty"` // 主管姓名
	LeaderType *string `json:"leader_type,omitempty"` // 主管类型，枚举英文：major 主要主管 other 其他主管
}

type LeaderItemBuilder struct {
	leaderId      string // 主管ID
	leaderIdSet   bool
	leaderName    string // 主管姓名
	leaderNameSet bool
	leaderType    string // 主管类型，枚举英文：major 主要主管 other 其他主管
	leaderTypeSet bool
}

func NewLeaderItemBuilder() *LeaderItemBuilder {
	return &LeaderItemBuilder{}
}
func (builder *LeaderItemBuilder) LeaderId(leaderId string) *LeaderItemBuilder {
	builder.leaderId = leaderId
	builder.leaderIdSet = true
	return builder
}
func (builder *LeaderItemBuilder) LeaderName(leaderName string) *LeaderItemBuilder {
	builder.leaderName = leaderName
	builder.leaderNameSet = true
	return builder
}
func (builder *LeaderItemBuilder) LeaderType(leaderType string) *LeaderItemBuilder {
	builder.leaderType = leaderType
	builder.leaderTypeSet = true
	return builder
}

func (builder *LeaderItemBuilder) Build() *LeaderItem {
	data := &LeaderItem{}
	if builder.leaderIdSet {
		data.LeaderId = &builder.leaderId
	}
	if builder.leaderNameSet {
		data.LeaderName = &builder.leaderName
	}
	if builder.leaderTypeSet {
		data.LeaderType = &builder.leaderType
	}
	return data
}
