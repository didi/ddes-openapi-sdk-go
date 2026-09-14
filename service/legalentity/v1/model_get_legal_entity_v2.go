package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// GetLegalEntityV2ApiReq 集团账户（上级公司统一对接）查询公司主体请求，GET /open-apis/v2/legalEntity/get；普通企业用 GetLegalEntity
type GetLegalEntityV2ApiReq struct {
	apiReq *core.ApiReq
}
type GetLegalEntityV2ApiResp struct {
	*core.ApiResp            `json:"-"`
	GetLegalEntityV2ApiReply *GetLegalEntityV2ApiReply `json:"getLegalEntityV2ApiReply"`
}

type GetLegalEntityV2ApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetLegalEntityV2ApiReqBuilder() *GetLegalEntityV2ApiReqBuilder {
	builder := &GetLegalEntityV2ApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetLegalEntityV2ApiReqBuilder) Build() *GetLegalEntityV2ApiReq {
	req := &GetLegalEntityV2ApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetLegalEntityV2ApiReqBuilder) ClientId(clientid string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) AccessToken(accesstoken string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) CompanyId(companyid string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) Timestamp(timestamp int32) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) Sign(sign string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) Keyword(keyword string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("keyword", keyword)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) LegalEntityId(legalentityid string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("legal_entity_id", legalentityid)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) OutLegalEntityId(outlegalentityid string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_legal_entity_id", outlegalentityid)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) Length(length int32) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) NextToken(nextToken string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("next_token", nextToken)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) QuerySubCompany(querySubCompany string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("query_sub_company", querySubCompany)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) BelongEnterpriseName(belongEnterpriseName string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("belong_enterprise_name", belongEnterpriseName)
	return builder
}
func (builder *GetLegalEntityV2ApiReqBuilder) TaxpayerNo(taxpayerNo string) *GetLegalEntityV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("taxpayer_no", taxpayerNo)
	return builder
}
