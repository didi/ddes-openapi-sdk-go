package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// ListMemberV2ApiReq 集团账户（上级公司统一对接）查询员工列表请求，GET /open-apis/v2/member/get；普通企业用 ListMember
type ListMemberV2ApiReq struct {
	apiReq *core.ApiReq
}
type ListMemberV2ApiResp struct {
	*core.ApiResp        `json:"-"`
	ListMemberV2ApiReply *ListMemberV2ApiReply `json:"listMemberV2ApiReply"`
}

type ListMemberV2ApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListMemberV2ApiReqBuilder() *ListMemberV2ApiReqBuilder {
	builder := &ListMemberV2ApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListMemberV2ApiReqBuilder) Build() *ListMemberV2ApiReq {
	req := &ListMemberV2ApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListMemberV2ApiReqBuilder) ClientId(clientid string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) AccessToken(accesstoken string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) CompanyId(companyid string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Timestamp(timestamp int32) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Sign(sign string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Email(email string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("email", email)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) EmployeeNumber(employeenumber string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("employee_number", employeenumber)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Phone(phone string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Status(status string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("status", status)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) Length(length int32) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) NextToken(nextToken string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("next_token", nextToken)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) QuerySubCompany(querySubCompany string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("query_sub_company", querySubCompany)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) BelongEnterpriseName(belongEnterpriseName string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("belong_enterprise_name", belongEnterpriseName)
	return builder
}
func (builder *ListMemberV2ApiReqBuilder) TaxpayerNo(taxpayerNo string) *ListMemberV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("taxpayer_no", taxpayerNo)
	return builder
}
