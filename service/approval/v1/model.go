package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type ApprovalPassApiReq struct {
	apiReq              *core.ApiReq
	approvalPassRequest *ApprovalPassRequest
}
type ApprovalPassApiResp struct {
	*core.ApiResp        `json:"-"`
	ApprovalPassApiReply *ApprovalPassApiReply `json:"approvalPassApiReply"`
}

type ApprovalPassApiReqBuilder struct {
	apiReq              *core.ApiReq
	approvalPassRequest *ApprovalPassRequest
}

func NewApprovalPassApiReqBuilder() *ApprovalPassApiReqBuilder {
	builder := &ApprovalPassApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ApprovalPassApiReqBuilder) ApprovalPassRequest(approvalPassRequest *ApprovalPassRequest) *ApprovalPassApiReqBuilder {
	builder.approvalPassRequest = approvalPassRequest
	return builder
}

func (builder *ApprovalPassApiReqBuilder) Build() *ApprovalPassApiReq {
	req := &ApprovalPassApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.approvalPassRequest
	return req
}

type CancelApprovalApiReq struct {
	apiReq                *core.ApiReq
	cancelApprovalRequest *CancelApprovalRequest
}
type CancelApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	CancelApprovalApiReply *CancelApprovalApiReply `json:"cancelApprovalApiReply"`
}

type CancelApprovalApiReqBuilder struct {
	apiReq                *core.ApiReq
	cancelApprovalRequest *CancelApprovalRequest
}

func NewCancelApprovalApiReqBuilder() *CancelApprovalApiReqBuilder {
	builder := &CancelApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CancelApprovalApiReqBuilder) CancelApprovalRequest(cancelApprovalRequest *CancelApprovalRequest) *CancelApprovalApiReqBuilder {
	builder.cancelApprovalRequest = cancelApprovalRequest
	return builder
}

func (builder *CancelApprovalApiReqBuilder) Build() *CancelApprovalApiReq {
	req := &CancelApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.cancelApprovalRequest
	return req
}

type CreateTravelApprovalApiReq struct {
	apiReq                *core.ApiReq
	createApprovalRequest *CreateApprovalRequest
}
type CreateTravelApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	CreateApprovalApiReply *CreateApprovalApiReply `json:"createApprovalApiReply"`
}

type CreateTravelApprovalApiReqBuilder struct {
	apiReq                *core.ApiReq
	createApprovalRequest *CreateApprovalRequest
}

func NewCreateTravelApprovalApiReqBuilder() *CreateTravelApprovalApiReqBuilder {
	builder := &CreateTravelApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateTravelApprovalApiReqBuilder) CreateTravelApprovalRequest(createApprovalRequest *CreateApprovalRequest) *CreateTravelApprovalApiReqBuilder {
	builder.createApprovalRequest = createApprovalRequest
	return builder
}

func (builder *CreateTravelApprovalApiReqBuilder) Build() *CreateTravelApprovalApiReq {
	req := &CreateTravelApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createApprovalRequest
	return req
}

type CreateBusinessByTimesApprovalApiReq struct {
	apiReq                               *core.ApiReq
	createApprovalBusinessByTimesRequest *CreateApprovalBusinessByTimesRequest
}
type CreateBusinessByTimesApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	CreateApprovalApiReply *CreateApprovalApiReply `json:"createApprovalApiReply"`
}

type CreateBusinessByTimesApprovalApiReqBuilder struct {
	apiReq                               *core.ApiReq
	createApprovalBusinessByTimesRequest *CreateApprovalBusinessByTimesRequest
}

func NewCreateBusinessByTimesApprovalApiReqBuilder() *CreateBusinessByTimesApprovalApiReqBuilder {
	builder := &CreateBusinessByTimesApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateBusinessByTimesApprovalApiReqBuilder) CreateBusinessByTimesApprovalRequest(createApprovalBusinessByTimesRequest *CreateApprovalBusinessByTimesRequest) *CreateBusinessByTimesApprovalApiReqBuilder {
	builder.createApprovalBusinessByTimesRequest = createApprovalBusinessByTimesRequest
	return builder
}

func (builder *CreateBusinessByTimesApprovalApiReqBuilder) Build() *CreateBusinessByTimesApprovalApiReq {
	req := &CreateBusinessByTimesApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createApprovalBusinessByTimesRequest
	return req
}

type CreateBusinessByDateApprovalApiReq struct {
	apiReq                              *core.ApiReq
	createApprovalBusinessByDateRequest *CreateApprovalBusinessByDateRequest
}
type CreateBusinessByDateApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	CreateApprovalApiReply *CreateApprovalApiReply `json:"createApprovalApiReply"`
}

type CreateBusinessByDateApprovalApiReqBuilder struct {
	apiReq                              *core.ApiReq
	createApprovalBusinessByDateRequest *CreateApprovalBusinessByDateRequest
}

func NewCreateBusinessByDateApprovalApiReqBuilder() *CreateBusinessByDateApprovalApiReqBuilder {
	builder := &CreateBusinessByDateApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateBusinessByDateApprovalApiReqBuilder) CreateBusinessByDateApprovalRequest(createApprovalBusinessByDateRequest *CreateApprovalBusinessByDateRequest) *CreateBusinessByDateApprovalApiReqBuilder {
	builder.createApprovalBusinessByDateRequest = createApprovalBusinessByDateRequest
	return builder
}

func (builder *CreateBusinessByDateApprovalApiReqBuilder) Build() *CreateBusinessByDateApprovalApiReq {
	req := &CreateBusinessByDateApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createApprovalBusinessByDateRequest
	return req
}

type GetApprovalDetailApiReq struct {
	apiReq *core.ApiReq
}
type GetApprovalDetailApiResp struct {
	*core.ApiResp             `json:"-"`
	GetApprovalDetailApiReply *GetApprovalDetailApiReply `json:"getApprovalDetailApiReply"`
}

type GetApprovalDetailApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetApprovalDetailApiReqBuilder() *GetApprovalDetailApiReqBuilder {
	builder := &GetApprovalDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetApprovalDetailApiReqBuilder) Build() *GetApprovalDetailApiReq {
	req := &GetApprovalDetailApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetApprovalDetailApiReqBuilder) ClientId(clientid string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) AccessToken(accesstoken string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) Timestamp(timestamp string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) CompanyId(companyid string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) Sign(sign string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) ApprovalId(approvalid string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("approval_id", approvalid)
	return builder
}
func (builder *GetApprovalDetailApiReqBuilder) OutApprovalId(outapprovalid string) *GetApprovalDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_approval_id", outapprovalid)
	return builder
}

type ListApprovalOrderApiReq struct {
	apiReq *core.ApiReq
}
type ListApprovalOrderApiResp struct {
	*core.ApiResp             `json:"-"`
	ListApprovalOrderApiReply *ListApprovalOrderApiReply `json:"listApprovalOrderApiReply"`
}

type ListApprovalOrderApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListApprovalOrderApiReqBuilder() *ListApprovalOrderApiReqBuilder {
	builder := &ListApprovalOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListApprovalOrderApiReqBuilder) Build() *ListApprovalOrderApiReq {
	req := &ListApprovalOrderApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListApprovalOrderApiReqBuilder) ClientId(clientid string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) AccessToken(accesstoken string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) Timestamp(timestamp string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) CompanyId(companyid string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) Sign(sign string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) ApprovalId(approvalid string) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("approval_id", approvalid)
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) Offset(offset int32) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *ListApprovalOrderApiReqBuilder) Length(length int32) *ListApprovalOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}

type UpdateApprovalApiReq struct {
	apiReq                *core.ApiReq
	updateApprovalRequest *UpdateApprovalRequest
}
type UpdateApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	UpdateApprovalApiReply *UpdateApprovalApiReply `json:"updateApprovalApiReply"`
}

type UpdateApprovalApiReqBuilder struct {
	apiReq                *core.ApiReq
	updateApprovalRequest *UpdateApprovalRequest
}

func NewUpdateApprovalApiReqBuilder() *UpdateApprovalApiReqBuilder {
	builder := &UpdateApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateApprovalApiReqBuilder) UpdateApprovalRequest(updateApprovalRequest *UpdateApprovalRequest) *UpdateApprovalApiReqBuilder {
	builder.updateApprovalRequest = updateApprovalRequest
	return builder
}

func (builder *UpdateApprovalApiReqBuilder) Build() *UpdateApprovalApiReq {
	req := &UpdateApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateApprovalRequest
	return req
}

type UpdateBusinessByTimesApprovalApiReq struct {
	apiReq                               *core.ApiReq
	updateApprovalBusinessByTimesRequest *UpdateApprovalBusinessByTimesRequest
}
type UpdateBusinessByTimesApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	UpdateApprovalApiReply *UpdateApprovalApiReply `json:"updateApprovalApiReply"`
}

type UpdateBusinessByTimesApprovalApiReqBuilder struct {
	apiReq                               *core.ApiReq
	updateApprovalBusinessByTimesRequest *UpdateApprovalBusinessByTimesRequest
}

func NewUpdateBusinessByTimesApprovalApiReqBuilder() *UpdateBusinessByTimesApprovalApiReqBuilder {
	builder := &UpdateBusinessByTimesApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateBusinessByTimesApprovalApiReqBuilder) UpdateBusinessByTimesApprovalRequest(updateApprovalBusinessByTimesRequest *UpdateApprovalBusinessByTimesRequest) *UpdateBusinessByTimesApprovalApiReqBuilder {
	builder.updateApprovalBusinessByTimesRequest = updateApprovalBusinessByTimesRequest
	return builder
}

func (builder *UpdateBusinessByTimesApprovalApiReqBuilder) Build() *UpdateBusinessByTimesApprovalApiReq {
	req := &UpdateBusinessByTimesApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateApprovalBusinessByTimesRequest
	return req
}

type UpdateBusinessByDateApprovalApiReq struct {
	apiReq                              *core.ApiReq
	updateApprovalBusinessByDateRequest *UpdateApprovalBusinessByDateRequest
}
type UpdateBusinessByDateApprovalApiResp struct {
	*core.ApiResp          `json:"-"`
	UpdateApprovalApiReply *UpdateApprovalApiReply `json:"updateApprovalApiReply"`
}

type UpdateBusinessByDateApprovalApiReqBuilder struct {
	apiReq                              *core.ApiReq
	updateApprovalBusinessByDateRequest *UpdateApprovalBusinessByDateRequest
}

func NewUpdateBusinessByDateApprovalApiReqBuilder() *UpdateBusinessByDateApprovalApiReqBuilder {
	builder := &UpdateBusinessByDateApprovalApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateBusinessByDateApprovalApiReqBuilder) UpdateBusinessByDateApprovalRequest(updateApprovalBusinessByDateRequest *UpdateApprovalBusinessByDateRequest) *UpdateBusinessByDateApprovalApiReqBuilder {
	builder.updateApprovalBusinessByDateRequest = updateApprovalBusinessByDateRequest
	return builder
}

func (builder *UpdateBusinessByDateApprovalApiReqBuilder) Build() *UpdateBusinessByDateApprovalApiReq {
	req := &UpdateBusinessByDateApprovalApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateApprovalBusinessByDateRequest
	return req
}
