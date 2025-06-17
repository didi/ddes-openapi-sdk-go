package v1

// GetFlightOrderDetailRequest struct for GetFlightOrderDetailRequest
type GetFlightOrderDetailRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	CompanyId   *string `json:"company_id,omitempty"`   // 企业ID
	Timestamp   *int64  `json:"timestamp,omitempty"`    // 当前时间戳(精确到秒级)
	Sign        *string `json:"sign,omitempty"`         // 签名
	ProductType *int32  `json:"product_type,omitempty"` // 枚举数字：1  国内  2  国际
	OrderIds    *string `json:"order_ids,omitempty"`    // 订单号，多个订单用英文逗号连接，最多支持查询100个订单
}

type GetFlightOrderDetailRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	companyId      string // 企业ID
	companyIdSet   bool
	timestamp      int64 // 当前时间戳(精确到秒级)
	timestampSet   bool
	sign           string // 签名
	signSet        bool
	productType    int32 // 枚举数字：1  国内  2  国际
	productTypeSet bool
	orderIds       string // 订单号，多个订单用英文逗号连接，最多支持查询100个订单
	orderIdsSet    bool
}

func NewGetFlightOrderDetailRequestBuilder() *GetFlightOrderDetailRequestBuilder {
	return &GetFlightOrderDetailRequestBuilder{}
}
func (builder *GetFlightOrderDetailRequestBuilder) ClientId(clientId string) *GetFlightOrderDetailRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) AccessToken(accessToken string) *GetFlightOrderDetailRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) CompanyId(companyId string) *GetFlightOrderDetailRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) Timestamp(timestamp int64) *GetFlightOrderDetailRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) Sign(sign string) *GetFlightOrderDetailRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) ProductType(productType int32) *GetFlightOrderDetailRequestBuilder {
	builder.productType = productType
	builder.productTypeSet = true
	return builder
}
func (builder *GetFlightOrderDetailRequestBuilder) OrderIds(orderIds string) *GetFlightOrderDetailRequestBuilder {
	builder.orderIds = orderIds
	builder.orderIdsSet = true
	return builder
}

func (builder *GetFlightOrderDetailRequestBuilder) Build() *GetFlightOrderDetailRequest {
	data := &GetFlightOrderDetailRequest{}
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
	if builder.productTypeSet {
		data.ProductType = &builder.productType
	}
	if builder.orderIdsSet {
		data.OrderIds = &builder.orderIds
	}
	return data
}
