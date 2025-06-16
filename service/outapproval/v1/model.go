package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type UpdateOutApprovalStatusApiReq struct {
	apiReq                         *core.ApiReq
	updateOutApprovalStatusRequest *UpdateOutApprovalStatusRequest
}
type UpdateOutApprovalStatusApiResp struct {
	*core.ApiResp                   `json:"-"`
	UpdateOutApprovalStatusApiReply *UpdateOutApprovalStatusApiReply `json:"updateOutApprovalStatusApiReply"`
}

type UpdateOutApprovalStatusApiReqBuilder struct {
	apiReq                         *core.ApiReq
	updateOutApprovalStatusRequest *UpdateOutApprovalStatusRequest
}

func NewUpdateOutApprovalStatusApiReqBuilder() *UpdateOutApprovalStatusApiReqBuilder {
	builder := &UpdateOutApprovalStatusApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateOutApprovalStatusApiReqBuilder) UpdateOutApprovalStatusRequest(updateOutApprovalStatusRequest *UpdateOutApprovalStatusRequest) *UpdateOutApprovalStatusApiReqBuilder {
	builder.updateOutApprovalStatusRequest = updateOutApprovalStatusRequest
	return builder
}

func (builder *UpdateOutApprovalStatusApiReqBuilder) Build() *UpdateOutApprovalStatusApiReq {
	req := &UpdateOutApprovalStatusApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateOutApprovalStatusRequest
	return req
}
