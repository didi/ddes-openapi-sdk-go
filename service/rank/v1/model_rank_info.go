package v1

// RankInfo 外部
type RankInfo struct {
	RankId               *string `json:"rank_id,omitempty"`                // 滴滴侧职级 ID；更新时和out_rank_id二选一
	OutRankId            *string `json:"out_rank_id,omitempty"`            // 外部职级编号
	Name                 *string `json:"name,omitempty"`                   // 职级名称
	BelongEnterpriseName *string `json:"belong_enterprise_name,omitempty"` // 归属企业名称（集团账户）
	TaxpayerNo           *string `json:"taxpayer_no,omitempty"`            // 纳税人识别号（集团账户）
}

type RankInfoBuilder struct {
	rankId                  string // 滴滴侧职级 ID；更新时和out_rank_id二选一
	rankIdSet               bool
	outRankId               string // 外部职级编号
	outRankIdSet            bool
	name                    string // 职级名称
	nameSet                 bool
	belongEnterpriseName    string // 归属企业名称（集团账户）
	belongEnterpriseNameSet bool
	taxpayerNo              string // 纳税人识别号（集团账户）
	taxpayerNoSet           bool
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
func (builder *RankInfoBuilder) BelongEnterpriseName(belongEnterpriseName string) *RankInfoBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}
func (builder *RankInfoBuilder) TaxpayerNo(taxpayerNo string) *RankInfoBuilder {
	builder.taxpayerNo = taxpayerNo
	builder.taxpayerNoSet = true
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
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	if builder.taxpayerNoSet {
		data.TaxpayerNo = &builder.taxpayerNo
	}
	return data
}
