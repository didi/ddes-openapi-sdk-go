package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateWorkplaceApiReq struct {
	apiReq                 *core.ApiReq
	createWorkplaceRequest *CreateWorkplaceRequest
}
type CreateWorkplaceApiResp struct {
	*core.ApiResp           `json:"-"`
	CreateWorkplaceApiReply *CreateWorkplaceApiReply `json:"createWorkplaceApiReply"`
}

type CreateWorkplaceApiReqBuilder struct {
	apiReq                 *core.ApiReq
	createWorkplaceRequest *CreateWorkplaceRequest
}

func NewCreateWorkplaceApiReqBuilder() *CreateWorkplaceApiReqBuilder {
	builder := &CreateWorkplaceApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateWorkplaceApiReqBuilder) CreateWorkplaceRequest(createWorkplaceRequest *CreateWorkplaceRequest) *CreateWorkplaceApiReqBuilder {
	builder.createWorkplaceRequest = createWorkplaceRequest
	return builder
}

func (builder *CreateWorkplaceApiReqBuilder) Build() *CreateWorkplaceApiReq {
	req := &CreateWorkplaceApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createWorkplaceRequest
	return req
}

type DeleteWorkplaceApiReq struct {
	apiReq                 *core.ApiReq
	deleteWorkplaceRequest *DeleteWorkplaceRequest
}
type DeleteWorkplaceApiResp struct {
	*core.ApiResp           `json:"-"`
	DeleteWorkplaceApiReply *DeleteWorkplaceApiReply `json:"deleteWorkplaceApiReply"`
}

type DeleteWorkplaceApiReqBuilder struct {
	apiReq                 *core.ApiReq
	deleteWorkplaceRequest *DeleteWorkplaceRequest
}

func NewDeleteWorkplaceApiReqBuilder() *DeleteWorkplaceApiReqBuilder {
	builder := &DeleteWorkplaceApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *DeleteWorkplaceApiReqBuilder) DeleteWorkplaceRequest(deleteWorkplaceRequest *DeleteWorkplaceRequest) *DeleteWorkplaceApiReqBuilder {
	builder.deleteWorkplaceRequest = deleteWorkplaceRequest
	return builder
}

func (builder *DeleteWorkplaceApiReqBuilder) Build() *DeleteWorkplaceApiReq {
	req := &DeleteWorkplaceApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.deleteWorkplaceRequest
	return req
}

type UpdateWorkplaceApiReq struct {
	apiReq                 *core.ApiReq
	updateWorkplaceRequest *UpdateWorkplaceRequest
}
type UpdateWorkplaceApiResp struct {
	*core.ApiResp           `json:"-"`
	UpdateWorkplaceApiReply *UpdateWorkplaceApiReply `json:"updateWorkplaceApiReply"`
}

type UpdateWorkplaceApiReqBuilder struct {
	apiReq                 *core.ApiReq
	updateWorkplaceRequest *UpdateWorkplaceRequest
}

func NewUpdateWorkplaceApiReqBuilder() *UpdateWorkplaceApiReqBuilder {
	builder := &UpdateWorkplaceApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateWorkplaceApiReqBuilder) UpdateWorkplaceRequest(updateWorkplaceRequest *UpdateWorkplaceRequest) *UpdateWorkplaceApiReqBuilder {
	builder.updateWorkplaceRequest = updateWorkplaceRequest
	return builder
}

func (builder *UpdateWorkplaceApiReqBuilder) Build() *UpdateWorkplaceApiReq {
	req := &UpdateWorkplaceApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateWorkplaceRequest
	return req
}
