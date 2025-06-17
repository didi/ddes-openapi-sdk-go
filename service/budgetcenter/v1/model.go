package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateBudgetCenterApiReq struct {
	apiReq                    *core.ApiReq
	createBudgetCenterRequest *CreateBudgetCenterRequest
}
type CreateBudgetCenterApiResp struct {
	*core.ApiResp              `json:"-"`
	CreateBudgetCenterApiReply *CreateBudgetCenterApiReply `json:"createBudgetCenterApiReply"`
}

type CreateBudgetCenterApiReqBuilder struct {
	apiReq                    *core.ApiReq
	createBudgetCenterRequest *CreateBudgetCenterRequest
}

func NewCreateBudgetCenterApiReqBuilder() *CreateBudgetCenterApiReqBuilder {
	builder := &CreateBudgetCenterApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateBudgetCenterApiReqBuilder) CreateBudgetCenterRequest(createBudgetCenterRequest *CreateBudgetCenterRequest) *CreateBudgetCenterApiReqBuilder {
	builder.createBudgetCenterRequest = createBudgetCenterRequest
	return builder
}

func (builder *CreateBudgetCenterApiReqBuilder) Build() *CreateBudgetCenterApiReq {
	req := &CreateBudgetCenterApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createBudgetCenterRequest
	return req
}

type DelBudgetCenterApiReq struct {
	apiReq                 *core.ApiReq
	delBudgetCenterRequest *DelBudgetCenterRequest
}
type DelBudgetCenterApiResp struct {
	*core.ApiResp           `json:"-"`
	DelBudgetCenterApiReply *DelBudgetCenterApiReply `json:"delBudgetCenterApiReply"`
}

type DelBudgetCenterApiReqBuilder struct {
	apiReq                 *core.ApiReq
	delBudgetCenterRequest *DelBudgetCenterRequest
}

func NewDelBudgetCenterApiReqBuilder() *DelBudgetCenterApiReqBuilder {
	builder := &DelBudgetCenterApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DelBudgetCenterApiReqBuilder) DelBudgetCenterRequest(delBudgetCenterRequest *DelBudgetCenterRequest) *DelBudgetCenterApiReqBuilder {
	builder.delBudgetCenterRequest = delBudgetCenterRequest
	return builder
}

func (builder *DelBudgetCenterApiReqBuilder) Build() *DelBudgetCenterApiReq {
	req := &DelBudgetCenterApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.delBudgetCenterRequest
	return req
}

type GetBudgetCenterApiReq struct {
	apiReq *core.ApiReq
}
type GetBudgetCenterApiResp struct {
	*core.ApiResp           `json:"-"`
	GetBudgetCenterApiReply *GetBudgetCenterApiReply `json:"getBudgetCenterApiReply"`
}

type GetBudgetCenterApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBudgetCenterApiReqBuilder() *GetBudgetCenterApiReqBuilder {
	builder := &GetBudgetCenterApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBudgetCenterApiReqBuilder) Build() *GetBudgetCenterApiReq {
	req := &GetBudgetCenterApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBudgetCenterApiReqBuilder) ClientId(clientid string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) AccessToken(accesstoken string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) CompanyId(companyid string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Timestamp(timestamp int32) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Sign(sign string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Id(id string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("id", id)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) OutBudgetId(outbudgetid string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_budget_id", outbudgetid)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Type(type_ int32) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) IsExactName(isexactname int32) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("is_exact_name", fmt.Sprint(isexactname))
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Name(name string) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("name", name)
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Offset(offset int32) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *GetBudgetCenterApiReqBuilder) Length(length int32) *GetBudgetCenterApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}

type UpdateBudgetCenterApiReq struct {
	apiReq                    *core.ApiReq
	updateBudgetCenterRequest *UpdateBudgetCenterRequest
}
type UpdateBudgetCenterApiResp struct {
	*core.ApiResp              `json:"-"`
	UpdateBudgetCenterApiReply *UpdateBudgetCenterApiReply `json:"updateBudgetCenterApiReply"`
}

type UpdateBudgetCenterApiReqBuilder struct {
	apiReq                    *core.ApiReq
	updateBudgetCenterRequest *UpdateBudgetCenterRequest
}

func NewUpdateBudgetCenterApiReqBuilder() *UpdateBudgetCenterApiReqBuilder {
	builder := &UpdateBudgetCenterApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateBudgetCenterApiReqBuilder) UpdateBudgetCenterRequest(updateBudgetCenterRequest *UpdateBudgetCenterRequest) *UpdateBudgetCenterApiReqBuilder {
	builder.updateBudgetCenterRequest = updateBudgetCenterRequest
	return builder
}

func (builder *UpdateBudgetCenterApiReqBuilder) Build() *UpdateBudgetCenterApiReq {
	req := &UpdateBudgetCenterApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateBudgetCenterRequest
	return req
}
