package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateLegalEntityApiReq struct {
	apiReq                   *core.ApiReq
	createLegalEntityRequest *CreateLegalEntityRequest
}
type CreateLegalEntityApiResp struct {
	*core.ApiResp             `json:"-"`
	CreateLegalEntityApiReply *CreateLegalEntityApiReply `json:"createLegalEntityApiReply"`
}

type CreateLegalEntityApiReqBuilder struct {
	apiReq                   *core.ApiReq
	createLegalEntityRequest *CreateLegalEntityRequest
}

func NewCreateLegalEntityApiReqBuilder() *CreateLegalEntityApiReqBuilder {
	builder := &CreateLegalEntityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateLegalEntityApiReqBuilder) CreateLegalEntityRequest(createLegalEntityRequest *CreateLegalEntityRequest) *CreateLegalEntityApiReqBuilder {
	builder.createLegalEntityRequest = createLegalEntityRequest
	return builder
}

func (builder *CreateLegalEntityApiReqBuilder) Build() *CreateLegalEntityApiReq {
	req := &CreateLegalEntityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createLegalEntityRequest
	return req
}

type DelLegalEntityApiReq struct {
	apiReq                *core.ApiReq
	delLegalEntityRequest *DelLegalEntityRequest
}
type DelLegalEntityApiResp struct {
	*core.ApiResp          `json:"-"`
	DelLegalEntityApiReply *DelLegalEntityApiReply `json:"delLegalEntityApiReply"`
}

type DelLegalEntityApiReqBuilder struct {
	apiReq                *core.ApiReq
	delLegalEntityRequest *DelLegalEntityRequest
}

func NewDelLegalEntityApiReqBuilder() *DelLegalEntityApiReqBuilder {
	builder := &DelLegalEntityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DelLegalEntityApiReqBuilder) DelLegalEntityRequest(delLegalEntityRequest *DelLegalEntityRequest) *DelLegalEntityApiReqBuilder {
	builder.delLegalEntityRequest = delLegalEntityRequest
	return builder
}

func (builder *DelLegalEntityApiReqBuilder) Build() *DelLegalEntityApiReq {
	req := &DelLegalEntityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.delLegalEntityRequest
	return req
}

type GetLegalEntityApiReq struct {
	apiReq *core.ApiReq
}
type GetLegalEntityApiResp struct {
	*core.ApiResp          `json:"-"`
	GetLegalEntityApiReply *GetLegalEntityApiReply `json:"getLegalEntityApiReply"`
}

type GetLegalEntityApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetLegalEntityApiReqBuilder() *GetLegalEntityApiReqBuilder {
	builder := &GetLegalEntityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetLegalEntityApiReqBuilder) Build() *GetLegalEntityApiReq {
	req := &GetLegalEntityApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetLegalEntityApiReqBuilder) ClientId(clientid string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) AccessToken(accesstoken string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) CompanyId(companyid string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) Timestamp(timestamp string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) Sign(sign string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) Keyword(keyword string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("keyword", keyword)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) LegalEntityId(legalentityid string) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("legal_entity_id", legalentityid)
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) Offset(offset int32) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *GetLegalEntityApiReqBuilder) Length(length int32) *GetLegalEntityApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}

type UpdateLegalEntityApiReq struct {
	apiReq                   *core.ApiReq
	updateLegalEntityRequest *UpdateLegalEntityRequest
}
type UpdateLegalEntityApiResp struct {
	*core.ApiResp             `json:"-"`
	UpdateLegalEntityApiReply *UpdateLegalEntityApiReply `json:"updateLegalEntityApiReply"`
}

type UpdateLegalEntityApiReqBuilder struct {
	apiReq                   *core.ApiReq
	updateLegalEntityRequest *UpdateLegalEntityRequest
}

func NewUpdateLegalEntityApiReqBuilder() *UpdateLegalEntityApiReqBuilder {
	builder := &UpdateLegalEntityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateLegalEntityApiReqBuilder) UpdateLegalEntityRequest(updateLegalEntityRequest *UpdateLegalEntityRequest) *UpdateLegalEntityApiReqBuilder {
	builder.updateLegalEntityRequest = updateLegalEntityRequest
	return builder
}

func (builder *UpdateLegalEntityApiReqBuilder) Build() *UpdateLegalEntityApiReq {
	req := &UpdateLegalEntityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateLegalEntityRequest
	return req
}
