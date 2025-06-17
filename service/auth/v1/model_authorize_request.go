package v1

// AuthorizeRequest struct for AuthorizeRequest
type AuthorizeRequest struct {
	ClientId     *string `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	ClientSecret *string `json:"client_secret,omitempty"` // 申请应用时分配的AppSecret
	GrantType    *string `json:"grant_type,omitempty"`    // 请求的类型，填写client_credentials
	Timestamp    *int64  `json:"timestamp,omitempty"`     // 当前时间戳
	Sign         *string `json:"sign,omitempty"`          // 签名
}

type AuthorizeRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	clientSecret    string // 申请应用时分配的AppSecret
	clientSecretSet bool
	grantType       string // 请求的类型，填写client_credentials
	grantTypeSet    bool
	timestamp       int64 // 当前时间戳
	timestampSet    bool
	sign            string // 签名
	signSet         bool
}

func NewAuthorizeRequestBuilder() *AuthorizeRequestBuilder {
	return &AuthorizeRequestBuilder{}
}
func (builder *AuthorizeRequestBuilder) ClientId(clientId string) *AuthorizeRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *AuthorizeRequestBuilder) ClientSecret(clientSecret string) *AuthorizeRequestBuilder {
	builder.clientSecret = clientSecret
	builder.clientSecretSet = true
	return builder
}
func (builder *AuthorizeRequestBuilder) GrantType(grantType string) *AuthorizeRequestBuilder {
	builder.grantType = grantType
	builder.grantTypeSet = true
	return builder
}
func (builder *AuthorizeRequestBuilder) Timestamp(timestamp int64) *AuthorizeRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *AuthorizeRequestBuilder) Sign(sign string) *AuthorizeRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}

func (builder *AuthorizeRequestBuilder) Build() *AuthorizeRequest {
	data := &AuthorizeRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.clientSecretSet {
		data.ClientSecret = &builder.clientSecret
	}
	if builder.grantTypeSet {
		data.GrantType = &builder.grantType
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	return data
}
