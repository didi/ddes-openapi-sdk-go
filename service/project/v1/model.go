package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

type OutTravelerListApiReq struct {
	apiReq *core.ApiReq
}
type OutTravelerListApiResp struct {
	*core.ApiResp           `json:"-"`
	OutTravelerListApiReply *OutTravelerListApiReply `json:"outTravelerListApiReply"`
}

type OutTravelerListApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewOutTravelerListApiReqBuilder() *OutTravelerListApiReqBuilder {
	builder := &OutTravelerListApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *OutTravelerListApiReqBuilder) ClientId(clientid string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) AccessToken(accesstoken string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) CompanyId(companyid string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) Timestamp(timestamp string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) Sign(sign string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) ProjectId(projectid string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("project_id", projectid)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) OutBudgetId(outbudgetid string) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_budget_id", outbudgetid)
	return builder
}
func (builder *OutTravelerListApiReqBuilder) Page(page int32) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("page", fmt.Sprint(page))
	return builder
}
func (builder *OutTravelerListApiReqBuilder) PageSize(pagesize int32) *OutTravelerListApiReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pagesize))
	return builder
}

func (builder *OutTravelerListApiReqBuilder) Build() *OutTravelerListApiReq {
	req := &OutTravelerListApiReq{}
	req.apiReq = builder.apiReq
	return req
}
