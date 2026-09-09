package v1

// RankRecord 职级数据（getRanks 响应 records 元素；服务端返回字段名为 id，与请求侧 RankInfo 的 rank_id 不同）
type RankRecord struct {
	Id        *string `json:"id,omitempty"`          // 滴滴侧职级 ID
	Name      *string `json:"name,omitempty"`        // 职级名称
	OutRankId *string `json:"out_rank_id,omitempty"` // 外部职级 ID
}

type RankRecordBuilder struct {
	id           string // 滴滴侧职级 ID
	idSet        bool
	name         string // 职级名称
	nameSet      bool
	outRankId    string // 外部职级 ID
	outRankIdSet bool
}

func NewRankRecordBuilder() *RankRecordBuilder {
	return &RankRecordBuilder{}
}
func (builder *RankRecordBuilder) Id(id string) *RankRecordBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *RankRecordBuilder) Name(name string) *RankRecordBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *RankRecordBuilder) OutRankId(outRankId string) *RankRecordBuilder {
	builder.outRankId = outRankId
	builder.outRankIdSet = true
	return builder
}

func (builder *RankRecordBuilder) Build() *RankRecord {
	data := &RankRecord{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outRankIdSet {
		data.OutRankId = &builder.outRankId
	}
	return data
}
