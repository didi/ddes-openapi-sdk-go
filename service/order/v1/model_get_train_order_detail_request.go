package v1

// GetTrainOrderDetailRequest struct for GetTrainOrderDetailRequest
type GetTrainOrderDetailRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	CompanyId   *string `json:"company_id,omitempty"`   // 企业ID
	Timestamp   *int32  `json:"timestamp,omitempty"`    // 当前时间戳(精确到秒级)
	Sign        *string `json:"sign,omitempty"`         // 签名
	OrderIds    *string `json:"order_ids,omitempty"`    // 订单号，多个订单用英文逗号连接，最多支持查询100个订单
}

type GetTrainOrderDetailRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	companyId      string // 企业ID
	companyIdSet   bool
	timestamp      int32 // 当前时间戳(精确到秒级)
	timestampSet   bool
	sign           string // 签名
	signSet        bool
	orderIds       string // 订单号，多个订单用英文逗号连接，最多支持查询100个订单
	orderIdsSet    bool
}

func NewGetTrainOrderDetailRequestBuilder() *GetTrainOrderDetailRequestBuilder {
	return &GetTrainOrderDetailRequestBuilder{}
}
func (builder *GetTrainOrderDetailRequestBuilder) ClientId(clientId string) *GetTrainOrderDetailRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *GetTrainOrderDetailRequestBuilder) AccessToken(accessToken string) *GetTrainOrderDetailRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *GetTrainOrderDetailRequestBuilder) CompanyId(companyId string) *GetTrainOrderDetailRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *GetTrainOrderDetailRequestBuilder) Timestamp(timestamp int32) *GetTrainOrderDetailRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *GetTrainOrderDetailRequestBuilder) Sign(sign string) *GetTrainOrderDetailRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *GetTrainOrderDetailRequestBuilder) OrderIds(orderIds string) *GetTrainOrderDetailRequestBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}

func (builder *GetTrainOrderDetailRequestBuilder) Build() *GetTrainOrderDetailRequest {
	data := &GetTrainOrderDetailRequest{}
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
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	return data
}
