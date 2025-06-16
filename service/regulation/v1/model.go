package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type GetRegulationApiReq struct {
	apiReq *core.ApiReq
}
type GetRegulationApiResp struct {
	*core.ApiResp         `json:"-"`
	GetRegulationApiReply *GetRegulationApiReply `json:"getRegulationApiReply"`
}

type GetRegulationApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetRegulationApiReqBuilder() *GetRegulationApiReqBuilder {
	builder := &GetRegulationApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetRegulationApiReqBuilder) Build() *GetRegulationApiReq {
	req := &GetRegulationApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetRegulationApiReqBuilder) ClientId(clientid string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetRegulationApiReqBuilder) AccessToken(accesstoken string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetRegulationApiReqBuilder) CompanyId(companyid string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetRegulationApiReqBuilder) Timestamp(timestamp string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetRegulationApiReqBuilder) Sign(sign string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetRegulationApiReqBuilder) RegulationId(regulationid string) *GetRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("regulation_id", regulationid)
	return builder
}

type ListRegulationApiReq struct {
	apiReq *core.ApiReq
}
type ListRegulationApiResp struct {
	*core.ApiResp          `json:"-"`
	ListRegulationApiReply *ListRegulationApiReply `json:"listRegulationApiReply"`
}

type ListRegulationApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListRegulationApiReqBuilder() *ListRegulationApiReqBuilder {
	builder := &ListRegulationApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListRegulationApiReqBuilder) Build() *ListRegulationApiReq {
	req := &ListRegulationApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListRegulationApiReqBuilder) ClientId(clientid string) *ListRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListRegulationApiReqBuilder) AccessToken(accesstoken string) *ListRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListRegulationApiReqBuilder) CompanyId(companyid string) *ListRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListRegulationApiReqBuilder) Timestamp(timestamp string) *ListRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListRegulationApiReqBuilder) Sign(sign string) *ListRegulationApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
