package v1

// ListCityRequest struct for ListCityRequest
type ListCityRequest struct {
	ClientId     *string           `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string           `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string           `json:"company_id,omitempty"`   // 公司 ID
	Timestamp    *int64            `json:"timestamp,omitempty"`    // 当前时间戳
	Sign         *string           `json:"sign,omitempty"`         // 签名
	ParamJson    *string           `json:"param_json,omitempty"`   // 查询信息，(将 param_json 转为 json 字符串)
	ParamJsonObj *ListCityParamObj `json:"param_json__obj__,omitempty"`
}

type ListCityRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	companyId       string // 公司 ID
	companyIdSet    bool
	timestamp       int64 // 当前时间戳
	timestampSet    bool
	sign            string // 签名
	signSet         bool
	paramJson       string // 查询信息，(将 param_json 转为 json 字符串)
	paramJsonSet    bool
	paramJsonObj    ListCityParamObj
	paramJsonObjSet bool
}

func NewListCityRequestBuilder() *ListCityRequestBuilder {
	return &ListCityRequestBuilder{}
}
func (builder *ListCityRequestBuilder) ClientId(clientId string) *ListCityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListCityRequestBuilder) AccessToken(accessToken string) *ListCityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListCityRequestBuilder) CompanyId(companyId string) *ListCityRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListCityRequestBuilder) Timestamp(timestamp int64) *ListCityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListCityRequestBuilder) Sign(sign string) *ListCityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListCityRequestBuilder) ParamJson(paramJson string) *ListCityRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *ListCityRequestBuilder) ParamJsonObj(paramJsonObj ListCityParamObj) *ListCityRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *ListCityRequestBuilder) Build() *ListCityRequest {
	data := &ListCityRequest{}
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
	if builder.paramJsonSet {
		data.ParamJson = &builder.paramJson
	}
	if builder.paramJsonObjSet {
		data.ParamJsonObj = &builder.paramJsonObj
	}
	return data
}
