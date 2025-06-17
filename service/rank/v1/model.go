package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateRankApiReq struct {
	apiReq            *core.ApiReq
	createRankRequest *CreateRankRequest
}
type CreateRankApiResp struct {
	*core.ApiResp      `json:"-"`
	CreateRankApiReply *CreateRankApiReply `json:"createRankApiReply"`
}

type CreateRankApiReqBuilder struct {
	apiReq            *core.ApiReq
	createRankRequest *CreateRankRequest
}

func NewCreateRankApiReqBuilder() *CreateRankApiReqBuilder {
	builder := &CreateRankApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateRankApiReqBuilder) CreateRankRequest(createRankRequest *CreateRankRequest) *CreateRankApiReqBuilder {
	builder.createRankRequest = createRankRequest
	return builder
}

func (builder *CreateRankApiReqBuilder) Build() *CreateRankApiReq {
	req := &CreateRankApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createRankRequest
	return req
}

type DelRankApiReq struct {
	apiReq         *core.ApiReq
	delRankRequest *DelRankRequest
}
type DelRankApiResp struct {
	*core.ApiResp   `json:"-"`
	DelRankApiReply *DelRankApiReply `json:"delRankApiReply"`
}

type DelRankApiReqBuilder struct {
	apiReq         *core.ApiReq
	delRankRequest *DelRankRequest
}

func NewDelRankApiReqBuilder() *DelRankApiReqBuilder {
	builder := &DelRankApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DelRankApiReqBuilder) DelRankRequest(delRankRequest *DelRankRequest) *DelRankApiReqBuilder {
	builder.delRankRequest = delRankRequest
	return builder
}

func (builder *DelRankApiReqBuilder) Build() *DelRankApiReq {
	req := &DelRankApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.delRankRequest
	return req
}

type ListRankApiReq struct {
	apiReq *core.ApiReq
}
type ListRankApiResp struct {
	*core.ApiResp    `json:"-"`
	ListRankApiReply *ListRankApiReply `json:"listRankApiReply"`
}

type ListRankApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListRankApiReqBuilder() *ListRankApiReqBuilder {
	builder := &ListRankApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListRankApiReqBuilder) Build() *ListRankApiReq {
	req := &ListRankApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListRankApiReqBuilder) ClientId(clientid string) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListRankApiReqBuilder) AccessToken(accesstoken string) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListRankApiReqBuilder) CompanyId(companyid string) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListRankApiReqBuilder) Timestamp(timestamp string) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListRankApiReqBuilder) Offset(offset int32) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *ListRankApiReqBuilder) Length(length int32) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *ListRankApiReqBuilder) Sign(sign string) *ListRankApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}

type UpdateRankApiReq struct {
	apiReq            *core.ApiReq
	updateRankRequest *UpdateRankRequest
}
type UpdateRankApiResp struct {
	*core.ApiResp      `json:"-"`
	UpdateRankApiReply *UpdateRankApiReply `json:"updateRankApiReply"`
}

type UpdateRankApiReqBuilder struct {
	apiReq            *core.ApiReq
	updateRankRequest *UpdateRankRequest
}

func NewUpdateRankApiReqBuilder() *UpdateRankApiReqBuilder {
	builder := &UpdateRankApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateRankApiReqBuilder) UpdateRankRequest(updateRankRequest *UpdateRankRequest) *UpdateRankApiReqBuilder {
	builder.updateRankRequest = updateRankRequest
	return builder
}

func (builder *UpdateRankApiReqBuilder) Build() *UpdateRankApiReq {
	req := &UpdateRankApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateRankRequest
	return req
}
