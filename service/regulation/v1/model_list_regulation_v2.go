package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// ListRegulationV2ApiReq 集团账户（上级公司统一对接）查询制度列表请求，GET /open-apis/v2/regulation/get；普通企业用 ListRegulation。
// v2 契约为 next_token 跨公司分页（/river 版无分页）。
type ListRegulationV2ApiReq struct {
	apiReq *core.ApiReq
}
type ListRegulationV2ApiResp struct {
	*core.ApiResp            `json:"-"`
	ListRegulationV2ApiReply *ListRegulationV2ApiReply `json:"listRegulationV2ApiReply"`
}

type ListRegulationV2ApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListRegulationV2ApiReqBuilder() *ListRegulationV2ApiReqBuilder {
	builder := &ListRegulationV2ApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListRegulationV2ApiReqBuilder) Build() *ListRegulationV2ApiReq {
	req := &ListRegulationV2ApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListRegulationV2ApiReqBuilder) ClientId(clientid string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) AccessToken(accesstoken string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) CompanyId(companyid string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) Timestamp(timestamp string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) Sign(sign string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) Length(length int32) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) NextToken(nextToken string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("next_token", nextToken)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) QuerySubCompany(querySubCompany string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("query_sub_company", querySubCompany)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) BelongEnterpriseName(belongEnterpriseName string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("belong_enterprise_name", belongEnterpriseName)
	return builder
}
func (builder *ListRegulationV2ApiReqBuilder) TaxpayerNo(taxpayerNo string) *ListRegulationV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("taxpayer_no", taxpayerNo)
	return builder
}
