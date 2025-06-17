package v1

// ListAirportCityRequest struct for ListAirportCityRequest
type ListAirportCityRequest struct {
	ClientId     *string `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	AccessToken  *string `json:"access_token,omitempty"`  // 授权后的access token
	Timestamp    *int64  `json:"timestamp,omitempty"`     // 当前时间戳
	CompanyId    *string `json:"company_id,omitempty"`    // 企业ID
	Sign         *string `json:"sign,omitempty"`          // 签名
	CountryLevel *int32  `json:"country_level,omitempty"` // 国家类型, 枚举值数字 1：国内城市（默认） 2：国外城市 3：国内和国外城市
}

type ListAirportCityRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	timestamp       int64 // 当前时间戳
	timestampSet    bool
	companyId       string // 企业ID
	companyIdSet    bool
	sign            string // 签名
	signSet         bool
	countryLevel    int32 // 国家类型, 枚举值数字 1：国内城市（默认） 2：国外城市 3：国内和国外城市
	countryLevelSet bool
}

func NewListAirportCityRequestBuilder() *ListAirportCityRequestBuilder {
	return &ListAirportCityRequestBuilder{}
}
func (builder *ListAirportCityRequestBuilder) ClientId(clientId string) *ListAirportCityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListAirportCityRequestBuilder) AccessToken(accessToken string) *ListAirportCityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListAirportCityRequestBuilder) Timestamp(timestamp int64) *ListAirportCityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListAirportCityRequestBuilder) CompanyId(companyId string) *ListAirportCityRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListAirportCityRequestBuilder) Sign(sign string) *ListAirportCityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListAirportCityRequestBuilder) CountryLevel(countryLevel int32) *ListAirportCityRequestBuilder {
	builder.countryLevel = countryLevel
	builder.countryLevelSet = true
	return builder
}

func (builder *ListAirportCityRequestBuilder) Build() *ListAirportCityRequest {
	data := &ListAirportCityRequest{}
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
	if builder.countryLevelSet {
		data.CountryLevel = &builder.countryLevel
	}
	return data
}
