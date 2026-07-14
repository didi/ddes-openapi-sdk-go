package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

type GetProjectDetailApiReq struct {
	apiReq *core.ApiReq
}
type GetProjectDetailApiResp struct {
	*core.ApiResp            `json:"-"`
	GetProjectDetailApiReply *GetProjectDetailApiReply `json:"getProjectDetailApiReply"`
}

type GetProjectDetailApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetProjectDetailApiReqBuilder() *GetProjectDetailApiReqBuilder {
	builder := &GetProjectDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetProjectDetailApiReqBuilder) ClientId(clientid string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) AccessToken(accesstoken string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) CompanyId(companyid string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) Timestamp(timestamp string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) Sign(sign string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) ProjectId(projectid string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("project_id", projectid)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) ProjectName(projectname string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("project_name", projectname)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) ProjectCode(projectcode string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("project_code", projectcode)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) Offset(offset int32) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}

// Lenth 每页大小，最大为100，默认为20。参数名与文档保持一致
func (builder *GetProjectDetailApiReqBuilder) Lenth(lenth int32) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("lenth", fmt.Sprint(lenth))
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) BelongEnterpriseName(belongenterprisename string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("belong_enterprise_name", belongenterprisename)
	return builder
}
func (builder *GetProjectDetailApiReqBuilder) TaxpayerNo(taxpayerno string) *GetProjectDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("taxpayer_no", taxpayerno)
	return builder
}

func (builder *GetProjectDetailApiReqBuilder) Build() *GetProjectDetailApiReq {
	req := &GetProjectDetailApiReq{}
	req.apiReq = builder.apiReq
	return req
}

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

type UpdateMemberApiReq struct {
	apiReq              *core.ApiReq
	updateMemberRequest *UpdateMemberRequest
}
type UpdateMemberApiResp struct {
	*core.ApiResp        `json:"-"`
	UpdateMemberApiReply *UpdateMemberApiReply `json:"updateMemberApiReply"`
}

type UpdateMemberApiReqBuilder struct {
	apiReq              *core.ApiReq
	updateMemberRequest *UpdateMemberRequest
}

func NewUpdateMemberApiReqBuilder() *UpdateMemberApiReqBuilder {
	builder := &UpdateMemberApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateMemberApiReqBuilder) UpdateMemberRequest(updateMemberRequest *UpdateMemberRequest) *UpdateMemberApiReqBuilder {
	builder.updateMemberRequest = updateMemberRequest
	return builder
}

func (builder *UpdateMemberApiReqBuilder) Build() *UpdateMemberApiReq {
	req := &UpdateMemberApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateMemberRequest
	return req
}

type DelMemberApiReq struct {
	apiReq           *core.ApiReq
	delMemberRequest *DelMemberRequest
}
type DelMemberApiResp struct {
	*core.ApiResp     `json:"-"`
	DelMemberApiReply *DelMemberApiReply `json:"delMemberApiReply"`
}

type DelMemberApiReqBuilder struct {
	apiReq           *core.ApiReq
	delMemberRequest *DelMemberRequest
}

func NewDelMemberApiReqBuilder() *DelMemberApiReqBuilder {
	builder := &DelMemberApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DelMemberApiReqBuilder) DelMemberRequest(delMemberRequest *DelMemberRequest) *DelMemberApiReqBuilder {
	builder.delMemberRequest = delMemberRequest
	return builder
}

func (builder *DelMemberApiReqBuilder) Build() *DelMemberApiReq {
	req := &DelMemberApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.delMemberRequest
	return req
}
