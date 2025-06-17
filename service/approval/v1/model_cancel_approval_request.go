package v1

// CancelApprovalRequest 取消申请单
type CancelApprovalRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	Timestamp   *int64  `json:"timestamp,omitempty"`    // 当前时间戳(精确到秒级)
	CompanyId   *string `json:"company_id,omitempty"`   // 滴滴公司 ID
	Sign        *string `json:"sign,omitempty"`         // 签名
	ApprovalId  *string `json:"approval_id,omitempty"`  // 滴滴申请单ID，approval_id与out_approval_id同时存在时，以approval_id为准
	IsForce     *int32  `json:"is_force,omitempty"`     // 强制取消标志，1，是（需要开白名单）0，否
	// Deprecated
	OutApprovalId *string `json:"out_approval_id,omitempty"` // 外部申请单ID，客户侧申请单ID，不大于 120 字符。eg: TA_100002，与返回中的approvalid一一对应
}

type CancelApprovalRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	timestamp      int64 // 当前时间戳(精确到秒级)
	timestampSet   bool
	companyId      string // 滴滴公司 ID
	companyIdSet   bool
	sign           string // 签名
	signSet        bool
	approvalId     string // 滴滴申请单ID，approval_id与out_approval_id同时存在时，以approval_id为准
	approvalIdSet  bool
	isForce        int32 // 强制取消标志，1，是（需要开白名单）0，否
	isForceSet     bool
	// Deprecated
	outApprovalId    string // 外部申请单ID，客户侧申请单ID，不大于 120 字符。eg: TA_100002，与返回中的approvalid一一对应
	outApprovalIdSet bool
}

func NewCancelApprovalRequestBuilder() *CancelApprovalRequestBuilder {
	return &CancelApprovalRequestBuilder{}
}
func (builder *CancelApprovalRequestBuilder) ClientId(clientId string) *CancelApprovalRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) AccessToken(accessToken string) *CancelApprovalRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) Timestamp(timestamp int64) *CancelApprovalRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) CompanyId(companyId string) *CancelApprovalRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) Sign(sign string) *CancelApprovalRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) ApprovalId(approvalId string) *CancelApprovalRequestBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *CancelApprovalRequestBuilder) IsForce(isForce int32) *CancelApprovalRequestBuilder {
	builder.isForce = isForce
	builder.isForceSet = true
	return builder
}

// Deprecated
func (builder *CancelApprovalRequestBuilder) OutApprovalId(outApprovalId string) *CancelApprovalRequestBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}

func (builder *CancelApprovalRequestBuilder) Build() *CancelApprovalRequest {
	data := &CancelApprovalRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.isForceSet {
		data.IsForce = &builder.isForce
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	return data
}
