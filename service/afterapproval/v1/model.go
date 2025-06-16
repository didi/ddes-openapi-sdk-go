package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type CreatePersonalReceiptApiReq struct {
	apiReq                       *core.ApiReq
	createPersonalReceiptRequest *CreatePersonalReceiptRequest
}
type CreatePersonalReceiptApiResp struct {
	*core.ApiResp                 `json:"-"`
	CreatePersonalReceiptApiReply *CreatePersonalReceiptApiReply `json:"createPersonalReceiptApiReply"`
}

type CreatePersonalReceiptApiReqBuilder struct {
	apiReq                       *core.ApiReq
	createPersonalReceiptRequest *CreatePersonalReceiptRequest
}

func NewCreatePersonalReceiptApiReqBuilder() *CreatePersonalReceiptApiReqBuilder {
	builder := &CreatePersonalReceiptApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *CreatePersonalReceiptApiReqBuilder) CreatePersonalReceiptRequest(createPersonalReceiptRequest *CreatePersonalReceiptRequest) *CreatePersonalReceiptApiReqBuilder {
	builder.createPersonalReceiptRequest = createPersonalReceiptRequest
	return builder
}

func (builder *CreatePersonalReceiptApiReqBuilder) Build() *CreatePersonalReceiptApiReq {
	req := &CreatePersonalReceiptApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.createPersonalReceiptRequest
	return req
}

type GetPersonalReceiptOrderApiReq struct {
	apiReq *core.ApiReq
}
type GetPersonalReceiptOrderApiResp struct {
	*core.ApiResp                   `json:"-"`
	GetPersonalReceiptOrderApiReply *GetPersonalReceiptOrderApiReply `json:"getPersonalReceiptOrderApiReply"`
}

type GetPersonalReceiptOrderApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetPersonalReceiptOrderApiReqBuilder() *GetPersonalReceiptOrderApiReqBuilder {
	builder := &GetPersonalReceiptOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetPersonalReceiptOrderApiReqBuilder) Build() *GetPersonalReceiptOrderApiReq {
	req := &GetPersonalReceiptOrderApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) ClientId(clientid string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) AccessToken(accesstoken string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) CompanyId(companyid string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Timestamp(timestamp string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Sign(sign string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) QueryTimeType(querytimetype int32) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("query_time_type", fmt.Sprint(querytimetype))
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) StartDate(startdate string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) EndDate(enddate string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Status(status int32) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("status", fmt.Sprint(status))
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Type(type_ int32) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Phone(phone string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Offset(offset string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", offset)
	return builder
}
func (builder *GetPersonalReceiptOrderApiReqBuilder) Length(length string) *GetPersonalReceiptOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", length)
	return builder
}
