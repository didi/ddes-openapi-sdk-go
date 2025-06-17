package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type GetLoginEncryptStrApiReq struct {
	apiReq *core.ApiReq
}
type GetLoginEncryptStrApiResp struct {
	*core.ApiResp              `json:"-"`
	GetLoginEncryptStrApiReply *GetLoginEncryptStrApiReply `json:"getLoginEncryptStrApiReply"`
}

type GetLoginEncryptStrApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetLoginEncryptStrApiReqBuilder() *GetLoginEncryptStrApiReqBuilder {
	builder := &GetLoginEncryptStrApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetLoginEncryptStrApiReqBuilder) Build() *GetLoginEncryptStrApiReq {
	req := &GetLoginEncryptStrApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetLoginEncryptStrApiReqBuilder) ClientId(clientid string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) AccessToken(accesstoken string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) Timestamp(timestamp string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) Sign(sign string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) AppType(apptype string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("app_type", apptype)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) Phone(phone string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) Email(email string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("email", email)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) EmployeeNumber(employeenumber string) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("employee_number", employeenumber)
	return builder
}
func (builder *GetLoginEncryptStrApiReqBuilder) ProductType(producttype int32) *GetLoginEncryptStrApiReqBuilder {
	builder.apiReq.QueryParams.Set("product_type", fmt.Sprint(producttype))
	return builder
}
