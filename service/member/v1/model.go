package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreateMemberApiReq struct {
	apiReq              *core.ApiReq
	createMemberRequest *CreateMemberRequest
}
type CreateMemberApiResp struct {
	*core.ApiResp        `json:"-"`
	CreateMemberApiReply *CreateMemberApiReply `json:"createMemberApiReply"`
}

type CreateMemberApiReqBuilder struct {
	apiReq              *core.ApiReq
	createMemberRequest *CreateMemberRequest
}

func NewCreateMemberApiReqBuilder() *CreateMemberApiReqBuilder {
	builder := &CreateMemberApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreateMemberApiReqBuilder) CreateMemberRequest(createMemberRequest *CreateMemberRequest) *CreateMemberApiReqBuilder {
	builder.createMemberRequest = createMemberRequest
	return builder
}

func (builder *CreateMemberApiReqBuilder) Build() *CreateMemberApiReq {
	req := &CreateMemberApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createMemberRequest
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

type GetMemberDetailApiReq struct {
	apiReq *core.ApiReq
}
type GetMemberDetailApiResp struct {
	*core.ApiResp           `json:"-"`
	GetMemberDetailApiReply *GetMemberDetailApiReply `json:"getMemberDetailApiReply"`
}

type GetMemberDetailApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetMemberDetailApiReqBuilder() *GetMemberDetailApiReqBuilder {
	builder := &GetMemberDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetMemberDetailApiReqBuilder) Build() *GetMemberDetailApiReq {
	req := &GetMemberDetailApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetMemberDetailApiReqBuilder) ClientId(clientid string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) AccessToken(accesstoken string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) CompanyId(companyid string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) Timestamp(timestamp int32) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) Sign(sign string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) MemberId(memberid string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("member_id", memberid)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) EmployeeNumber(employeenumber string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("employee_number", employeenumber)
	return builder
}
func (builder *GetMemberDetailApiReqBuilder) Phone(phone string) *GetMemberDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}

type GetMemberQuotaApiReq struct {
	apiReq *core.ApiReq
}
type GetMemberQuotaApiResp struct {
	*core.ApiResp          `json:"-"`
	GetMemberQuotaApiReply *GetMemberQuotaApiReply `json:"getMemberQuotaApiReply"`
}

type GetMemberQuotaApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetMemberQuotaApiReqBuilder() *GetMemberQuotaApiReqBuilder {
	builder := &GetMemberQuotaApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetMemberQuotaApiReqBuilder) Build() *GetMemberQuotaApiReq {
	req := &GetMemberQuotaApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetMemberQuotaApiReqBuilder) ClientId(clientid string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) AccessToken(accesstoken string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) CompanyId(companyid string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) Timestamp(timestamp int32) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) Sign(sign string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) MemberIds(memberids string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("member_ids", memberids)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) EmployeeNumbers(employeenumbers string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("employee_numbers", employeenumbers)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) StartDate(startdate string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetMemberQuotaApiReqBuilder) EndDate(enddate string) *GetMemberQuotaApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}

type ListMemberApiReq struct {
	apiReq *core.ApiReq
}
type ListMemberApiResp struct {
	*core.ApiResp      `json:"-"`
	ListMemberApiReply *ListMemberApiReply `json:"listMemberApiReply"`
}

type ListMemberApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListMemberApiReqBuilder() *ListMemberApiReqBuilder {
	builder := &ListMemberApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListMemberApiReqBuilder) Build() *ListMemberApiReq {
	req := &ListMemberApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListMemberApiReqBuilder) ClientId(clientid string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListMemberApiReqBuilder) AccessToken(accesstoken string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListMemberApiReqBuilder) CompanyId(companyid string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListMemberApiReqBuilder) Timestamp(timestamp int32) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *ListMemberApiReqBuilder) Sign(sign string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *ListMemberApiReqBuilder) Email(email string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("email", email)
	return builder
}
func (builder *ListMemberApiReqBuilder) EmployeeNumber(employeenumber string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("employee_number", employeenumber)
	return builder
}
func (builder *ListMemberApiReqBuilder) Phone(phone string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}
func (builder *ListMemberApiReqBuilder) Realname(realname string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("realname", realname)
	return builder
}
func (builder *ListMemberApiReqBuilder) Status(status string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("status", status)
	return builder
}
func (builder *ListMemberApiReqBuilder) Offset(offset int32) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *ListMemberApiReqBuilder) Length(length int32) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *ListMemberApiReqBuilder) LastId(lastid string) *ListMemberApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
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
