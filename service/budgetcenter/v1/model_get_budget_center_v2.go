package v1

import (
	"fmt"
	"net/url"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// GetBudgetCenterV2ApiReq 集团账户（上级公司统一对接）查询部门/项目请求，POST /open-apis/v2/budgetCenter/get；普通企业用 GetBudgetCenter。
// v2 契约无 offset/is_get_poi/is_get_extend_fields（分页用 next_token，无 POI/扩展字段返回）。
type GetBudgetCenterV2ApiReq struct {
	apiReq *core.ApiReq
}
type GetBudgetCenterV2ApiResp struct {
	*core.ApiResp             `json:"-"`
	GetBudgetCenterV2ApiReply *GetBudgetCenterV2ApiReply `json:"getBudgetCenterV2ApiReply"`
}

type GetBudgetCenterV2ApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBudgetCenterV2ApiReqBuilder() *GetBudgetCenterV2ApiReqBuilder {
	builder := &GetBudgetCenterV2ApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBudgetCenterV2ApiReqBuilder) Build() *GetBudgetCenterV2ApiReq {
	req := &GetBudgetCenterV2ApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBudgetCenterV2ApiReqBuilder) ClientId(clientid string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) AccessToken(accesstoken string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) CompanyId(companyid string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Timestamp(timestamp int32) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Sign(sign string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Id(id string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("id", id)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) OutBudgetId(outbudgetid string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_budget_id", outbudgetid)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Type(type_ int32) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) IsExactName(isexactname int32) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("is_exact_name", fmt.Sprint(isexactname))
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Name(name string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("name", name)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) ThirdOutId(thirdOutId string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("third_out_id", thirdOutId)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) IsNeedLimitRule(isNeedLimitRule int32) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("is_need_limit_rule", fmt.Sprint(isNeedLimitRule))
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) Length(length int32) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) NextToken(nextToken string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("next_token", nextToken)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) QuerySubCompany(querySubCompany string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("query_sub_company", querySubCompany)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) BelongEnterpriseName(belongEnterpriseName string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("belong_enterprise_name", belongEnterpriseName)
	return builder
}
func (builder *GetBudgetCenterV2ApiReqBuilder) TaxpayerNo(taxpayerNo string) *GetBudgetCenterV2ApiReqBuilder {
	builder.apiReq.QueryParams.Set("taxpayer_no", taxpayerNo)
	return builder
}
