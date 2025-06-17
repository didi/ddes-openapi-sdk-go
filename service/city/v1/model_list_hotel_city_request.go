package v1

// ListHotelCityRequest struct for ListHotelCityRequest
type ListHotelCityRequest struct {
	ClientId    *string `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"` // 授权后的access token
	Timestamp   *int64  `json:"timestamp,omitempty"`    // 当前时间戳
	Sign        *string `json:"sign,omitempty"`         // 签名
	CountryId   *int32  `json:"country_id,omitempty"`   // 国家ID
}

type ListHotelCityRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	timestamp      int64 // 当前时间戳
	timestampSet   bool
	sign           string // 签名
	signSet        bool
	countryId      int32 // 国家ID
	countryIdSet   bool
}

func NewListHotelCityRequestBuilder() *ListHotelCityRequestBuilder {
	return &ListHotelCityRequestBuilder{}
}
func (builder *ListHotelCityRequestBuilder) ClientId(clientId string) *ListHotelCityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListHotelCityRequestBuilder) AccessToken(accessToken string) *ListHotelCityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListHotelCityRequestBuilder) Timestamp(timestamp int64) *ListHotelCityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListHotelCityRequestBuilder) Sign(sign string) *ListHotelCityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListHotelCityRequestBuilder) CountryId(countryId int32) *ListHotelCityRequestBuilder {
	builder.countryId = countryId
	builder.countryIdSet = true
	return builder
}

func (builder *ListHotelCityRequestBuilder) Build() *ListHotelCityRequest {
	data := &ListHotelCityRequest{}
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
	if builder.countryIdSet {
		data.CountryId = &builder.countryId
	}
	return data
}
