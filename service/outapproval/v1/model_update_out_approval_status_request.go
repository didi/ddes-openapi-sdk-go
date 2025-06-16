package v1

// UpdateOutApprovalStatusRequest struct for UpdateOutApprovalStatusRequest
type UpdateOutApprovalStatusRequest struct {
	ClientId     *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string `json:"company_id,omitempty"`   // 企业ID，滴滴企业ID
	Timestamp    *int64  `json:"timestamp,omitempty"`    // 当前时间戳(精确到秒级)
	Sign         *string `json:"sign,omitempty"`         // 签名，签名（注：本接口的传参都要参与签名）
	OutId        *string `json:"out_id,omitempty"`       // 外部审批单id，使用滴滴消息通知中的out_id
	Status       *int32  `json:"status,omitempty"`       // 审批单状态，0 驳回 1 通过
	ApprovalType *int32  `json:"approval_type,omitempty"`
}

type UpdateOutApprovalStatusRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	companyId       string // 企业ID，滴滴企业ID
	companyIdSet    bool
	timestamp       int64 // 当前时间戳(精确到秒级)
	timestampSet    bool
	sign            string // 签名，签名（注：本接口的传参都要参与签名）
	signSet         bool
	outId           string // 外部审批单id，使用滴滴消息通知中的out_id
	outIdSet        bool
	status          int32 // 审批单状态，0 驳回 1 通过
	statusSet       bool
	approvalType    int32
	approvalTypeSet bool
}

func NewUpdateOutApprovalStatusRequestBuilder() *UpdateOutApprovalStatusRequestBuilder {
	return &UpdateOutApprovalStatusRequestBuilder{}
}
func (builder *UpdateOutApprovalStatusRequestBuilder) ClientId(clientId string) *UpdateOutApprovalStatusRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) AccessToken(accessToken string) *UpdateOutApprovalStatusRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) CompanyId(companyId string) *UpdateOutApprovalStatusRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) Timestamp(timestamp int64) *UpdateOutApprovalStatusRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) Sign(sign string) *UpdateOutApprovalStatusRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) OutId(outId string) *UpdateOutApprovalStatusRequestBuilder {
	builder.outId = outId
	builder.outIdSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) Status(status int32) *UpdateOutApprovalStatusRequestBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *UpdateOutApprovalStatusRequestBuilder) ApprovalType(approvalType int32) *UpdateOutApprovalStatusRequestBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}

func (builder *UpdateOutApprovalStatusRequestBuilder) Build() *UpdateOutApprovalStatusRequest {
	data := &UpdateOutApprovalStatusRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	if builder.outIdSet {
		data.OutId = &builder.outId
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	return data
}
