package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateTravelerApiReq struct {
	apiReq                *core.ApiReq
	createTravelerRequest *CreateTravelerRequest
}
type CreateTravelerApiResp struct {
	*core.ApiResp          `json:"-"`
	CreateTravelerApiReply *CreateTravelerApiReply `json:"createTravelerApiReply"`
}

type CreateTravelerApiReqBuilder struct {
	apiReq                *core.ApiReq
	createTravelerRequest *CreateTravelerRequest
}

func NewCreateTravelerApiReqBuilder() *CreateTravelerApiReqBuilder {
	builder := &CreateTravelerApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateTravelerApiReqBuilder) CreateTravelerRequest(createTravelerRequest *CreateTravelerRequest) *CreateTravelerApiReqBuilder {
	builder.createTravelerRequest = createTravelerRequest
	return builder
}

func (builder *CreateTravelerApiReqBuilder) Build() *CreateTravelerApiReq {
	req := &CreateTravelerApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createTravelerRequest
	return req
}

type DelTravelerApiReq struct {
	apiReq             *core.ApiReq
	delTravelerRequest *DelTravelerRequest
}
type DelTravelerApiResp struct {
	*core.ApiResp       `json:"-"`
	DelTravelerApiReply *DelTravelerApiReply `json:"delTravelerApiReply"`
}

type DelTravelerApiReqBuilder struct {
	apiReq             *core.ApiReq
	delTravelerRequest *DelTravelerRequest
}

func NewDelTravelerApiReqBuilder() *DelTravelerApiReqBuilder {
	builder := &DelTravelerApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DelTravelerApiReqBuilder) DelTravelerRequest(delTravelerRequest *DelTravelerRequest) *DelTravelerApiReqBuilder {
	builder.delTravelerRequest = delTravelerRequest
	return builder
}

func (builder *DelTravelerApiReqBuilder) Build() *DelTravelerApiReq {
	req := &DelTravelerApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.delTravelerRequest
	return req
}

type UpdateTravelerApiReq struct {
	apiReq                *core.ApiReq
	updateTravelerRequest *UpdateTravelerRequest
}
type UpdateTravelerApiResp struct {
	*core.ApiResp          `json:"-"`
	UpdateTravelerApiReply *UpdateTravelerApiReply `json:"updateTravelerApiReply"`
}

type UpdateTravelerApiReqBuilder struct {
	apiReq                *core.ApiReq
	updateTravelerRequest *UpdateTravelerRequest
}

func NewUpdateTravelerApiReqBuilder() *UpdateTravelerApiReqBuilder {
	builder := &UpdateTravelerApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateTravelerApiReqBuilder) UpdateTravelerRequest(updateTravelerRequest *UpdateTravelerRequest) *UpdateTravelerApiReqBuilder {
	builder.updateTravelerRequest = updateTravelerRequest
	return builder
}

func (builder *UpdateTravelerApiReqBuilder) Build() *UpdateTravelerApiReq {
	req := &UpdateTravelerApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateTravelerRequest
	return req
}
