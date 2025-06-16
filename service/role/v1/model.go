package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type ListRoleApiReq struct {
	apiReq *core.ApiReq
}
type ListRoleApiResp struct {
	*core.ApiResp    `json:"-"`
	ListRoleApiReply *ListRoleApiReply `json:"listRoleApiReply"`
}

type ListRoleApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListRoleApiReqBuilder() *ListRoleApiReqBuilder {
	builder := &ListRoleApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListRoleApiReqBuilder) Build() *ListRoleApiReq {
	req := &ListRoleApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListRoleApiReqBuilder) ClientId(clientid string) *ListRoleApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListRoleApiReqBuilder) AccessToken(accesstoken string) *ListRoleApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListRoleApiReqBuilder) CompanyId(companyid string) *ListRoleApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListRoleApiReqBuilder) Timestamp(timestamp string) *ListRoleApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListRoleApiReqBuilder) Sign(sign string) *ListRoleApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
