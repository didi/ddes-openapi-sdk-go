package v1

// ListCountryRequest struct for ListCountryRequest
type ListCountryRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	Timestamp   *int64  `json:"timestamp,omitempty"`    // 当前时间戳
	CompanyId   *string `json:"company_id,omitempty"`   // 企业ID
	Sign        *string `json:"sign,omitempty"`         // 签名
}

type ListCountryRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	timestamp      int64 // 当前时间戳
	timestampSet   bool
	companyId      string // 企业ID
	companyIdSet   bool
	sign           string // 签名
	signSet        bool
}

func NewListCountryRequestBuilder() *ListCountryRequestBuilder {
	return &ListCountryRequestBuilder{}
}
func (builder *ListCountryRequestBuilder) ClientId(clientId string) *ListCountryRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListCountryRequestBuilder) AccessToken(accessToken string) *ListCountryRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListCountryRequestBuilder) Timestamp(timestamp int64) *ListCountryRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListCountryRequestBuilder) CompanyId(companyId string) *ListCountryRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListCountryRequestBuilder) Sign(sign string) *ListCountryRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}

func (builder *ListCountryRequestBuilder) Build() *ListCountryRequest {
	data := &ListCountryRequest{}
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
	return data
}
