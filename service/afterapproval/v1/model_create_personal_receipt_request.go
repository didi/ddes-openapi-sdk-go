package v1

// CreatePersonalReceiptRequest struct for CreatePersonalReceiptRequest
type CreatePersonalReceiptRequest struct {
	ClientId      *string `json:"client_id,omitempty"`      // 申请应用时分配的AppKey
	AccessToken   *string `json:"access_token,omitempty"`   // 授权后的access token
	CompanyId     *string `json:"company_id,omitempty"`     // 企业ID
	Timestamp     *int64  `json:"timestamp,omitempty"`      // 当前时间戳(精确到秒级)
	Sign          *string `json:"sign,omitempty"`           // 签名
	OrderId       *string `json:"order_id,omitempty"`       // 订单号
	IsPass        *int32  `json:"is_pass,omitempty"`        // 审批类型 (0 - 驳回，1 - 通过，默认为0)
	ApproverPhone *string `json:"approver_phone,omitempty"` // 审批人手机号 (不填时默认为企业超管)
	Remark        *string `json:"remark,omitempty"`         // 备注，驳回原因或者通过备注，限制 120 个字符
}

type CreatePersonalReceiptRequestBuilder struct {
	clientId         string // 申请应用时分配的AppKey
	clientIdSet      bool
	accessToken      string // 授权后的access token
	accessTokenSet   bool
	companyId        string // 企业ID
	companyIdSet     bool
	timestamp        int64 // 当前时间戳(精确到秒级)
	timestampSet     bool
	sign             string // 签名
	signSet          bool
	orderId          string // 订单号
	orderIdSet       bool
	isPass           int32 // 审批类型 (0 - 驳回，1 - 通过，默认为0)
	isPassSet        bool
	approverPhone    string // 审批人手机号 (不填时默认为企业超管)
	approverPhoneSet bool
	remark           string // 备注，驳回原因或者通过备注，限制 120 个字符
	remarkSet        bool
}

func NewCreatePersonalReceiptRequestBuilder() *CreatePersonalReceiptRequestBuilder {
	return &CreatePersonalReceiptRequestBuilder{}
}
func (builder *CreatePersonalReceiptRequestBuilder) ClientId(clientId string) *CreatePersonalReceiptRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) AccessToken(accessToken string) *CreatePersonalReceiptRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) CompanyId(companyId string) *CreatePersonalReceiptRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) Timestamp(timestamp int64) *CreatePersonalReceiptRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) Sign(sign string) *CreatePersonalReceiptRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) OrderId(orderId string) *CreatePersonalReceiptRequestBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) IsPass(isPass int32) *CreatePersonalReceiptRequestBuilder {
	builder.isPass = isPass
	builder.isPassSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) ApproverPhone(approverPhone string) *CreatePersonalReceiptRequestBuilder {
	builder.approverPhone = approverPhone
	builder.approverPhoneSet = true
	return builder
}
func (builder *CreatePersonalReceiptRequestBuilder) Remark(remark string) *CreatePersonalReceiptRequestBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}

func (builder *CreatePersonalReceiptRequestBuilder) Build() *CreatePersonalReceiptRequest {
	data := &CreatePersonalReceiptRequest{}
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
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.isPassSet {
		data.IsPass = &builder.isPass
	}
	if builder.approverPhoneSet {
		data.ApproverPhone = &builder.approverPhone
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	return data
}
