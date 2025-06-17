package v1

// RankInfo 外部
type RankInfo struct {
	RankId    *string `json:"rank_id,omitempty"`     // 滴滴侧职级 ID；更新时和out_rank_id二选一
	OutRankId *string `json:"out_rank_id,omitempty"` // 外部职级编号
	Name      *string `json:"name,omitempty"`        // 职级名称
}

type RankInfoBuilder struct {
	rankId       string // 滴滴侧职级 ID；更新时和out_rank_id二选一
	rankIdSet    bool
	outRankId    string // 外部职级编号
	outRankIdSet bool
	name         string // 职级名称
	nameSet      bool
}

func NewRankInfoBuilder() *RankInfoBuilder {
	return &RankInfoBuilder{}
}
func (builder *RankInfoBuilder) RankId(rankId string) *RankInfoBuilder {
	builder.rankId = rankId
	builder.rankIdSet = true
	return builder
}
func (builder *RankInfoBuilder) OutRankId(outRankId string) *RankInfoBuilder {
	builder.outRankId = outRankId
	builder.outRankIdSet = true
	return builder
}
func (builder *RankInfoBuilder) Name(name string) *RankInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}

func (builder *RankInfoBuilder) Build() *RankInfo {
	data := &RankInfo{}
	if builder.rankIdSet {
		data.RankId = &builder.rankId
	}
	if builder.outRankIdSet {
		data.OutRankId = &builder.outRankId
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	return data
}
