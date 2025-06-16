package v1

// ListTrainCityRequest struct for ListTrainCityRequest
type ListTrainCityRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	Timestamp   *int64  `json:"timestamp,omitempty"`    // 当前时间戳
	Sign        *string `json:"sign,omitempty"`         // 签名
}

type ListTrainCityRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	timestamp      int64 // 当前时间戳
	timestampSet   bool
	sign           string // 签名
	signSet        bool
}

func NewListTrainCityRequestBuilder() *ListTrainCityRequestBuilder {
	return &ListTrainCityRequestBuilder{}
}
func (builder *ListTrainCityRequestBuilder) ClientId(clientId string) *ListTrainCityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListTrainCityRequestBuilder) AccessToken(accessToken string) *ListTrainCityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListTrainCityRequestBuilder) Timestamp(timestamp int64) *ListTrainCityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListTrainCityRequestBuilder) Sign(sign string) *ListTrainCityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}

func (builder *ListTrainCityRequestBuilder) Build() *ListTrainCityRequest {
	data := &ListTrainCityRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	return data
}
