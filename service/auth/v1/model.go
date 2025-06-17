package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type AuthorizeApiReq struct {
	apiReq           *core.ApiReq
	authorizeRequest *AuthorizeRequest
}
type AuthorizeApiResp struct {
	*core.ApiResp     `json:"-"`
	AuthorizeApiReply *AuthorizeApiReply `json:"authorizeApiReply"`
}

type AuthorizeApiReqBuilder struct {
	apiReq           *core.ApiReq
	authorizeRequest *AuthorizeRequest
}

func NewAuthorizeApiReqBuilder() *AuthorizeApiReqBuilder {
	builder := &AuthorizeApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *AuthorizeApiReqBuilder) AuthorizeRequest(authorizeRequest *AuthorizeRequest) *AuthorizeApiReqBuilder {
	builder.authorizeRequest = authorizeRequest
	return builder
}

func (builder *AuthorizeApiReqBuilder) Build() *AuthorizeApiReq {
	req := &AuthorizeApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.authorizeRequest
	return req
}
