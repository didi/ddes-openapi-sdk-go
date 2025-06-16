package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateExtendBatchApiReq struct {
	apiReq                   *core.ApiReq
	createExtendBatchRequest *CreateExtendBatchRequest
}
type CreateExtendBatchApiResp struct {
	*core.ApiResp             `json:"-"`
	CreateExtendBatchApiReply *CreateExtendBatchApiReply `json:"createExtendBatchApiReply"`
}

type CreateExtendBatchApiReqBuilder struct {
	apiReq                   *core.ApiReq
	createExtendBatchRequest *CreateExtendBatchRequest
}

func NewCreateExtendBatchApiReqBuilder() *CreateExtendBatchApiReqBuilder {
	builder := &CreateExtendBatchApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateExtendBatchApiReqBuilder) CreateExtendBatchRequest(createExtendBatchRequest *CreateExtendBatchRequest) *CreateExtendBatchApiReqBuilder {
	builder.createExtendBatchRequest = createExtendBatchRequest
	return builder
}

func (builder *CreateExtendBatchApiReqBuilder) Build() *CreateExtendBatchApiReq {
	req := &CreateExtendBatchApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createExtendBatchRequest
	return req
}

type ListExtendApiReq struct {
	apiReq *core.ApiReq
}
type ListExtendApiResp struct {
	*core.ApiResp      `json:"-"`
	ListExtendApiReply *ListExtendApiReply `json:"listExtendApiReply"`
}

type ListExtendApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListExtendApiReqBuilder() *ListExtendApiReqBuilder {
	builder := &ListExtendApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListExtendApiReqBuilder) Build() *ListExtendApiReq {
	req := &ListExtendApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListExtendApiReqBuilder) ClientId(clientid string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListExtendApiReqBuilder) AccessToken(accesstoken string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListExtendApiReqBuilder) Timestamp(timestamp string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListExtendApiReqBuilder) CompanyId(companyid string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListExtendApiReqBuilder) Sign(sign string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *ListExtendApiReqBuilder) RootCode(rootcode string) *ListExtendApiReqBuilder {
	builder.apiReq.QueryParams.Set("root_code", rootcode)
	return builder
}

type UpdateExtendStatusApiReq struct {
	apiReq                    *core.ApiReq
	updateExtendStatusRequest *UpdateExtendStatusRequest
}
type UpdateExtendStatusApiResp struct {
	*core.ApiResp              `json:"-"`
	UpdateExtendStatusApiReply *UpdateExtendStatusApiReply `json:"updateExtendStatusApiReply"`
}

type UpdateExtendStatusApiReqBuilder struct {
	apiReq                    *core.ApiReq
	updateExtendStatusRequest *UpdateExtendStatusRequest
}

func NewUpdateExtendStatusApiReqBuilder() *UpdateExtendStatusApiReqBuilder {
	builder := &UpdateExtendStatusApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateExtendStatusApiReqBuilder) UpdateExtendStatusRequest(updateExtendStatusRequest *UpdateExtendStatusRequest) *UpdateExtendStatusApiReqBuilder {
	builder.updateExtendStatusRequest = updateExtendStatusRequest
	return builder
}

func (builder *UpdateExtendStatusApiReqBuilder) Build() *UpdateExtendStatusApiReq {
	req := &UpdateExtendStatusApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateExtendStatusRequest
	return req
}
