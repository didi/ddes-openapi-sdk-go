package v1

// UpdateTravelerRequest struct for UpdateTravelerRequest
type UpdateTravelerRequest struct {
	ClientId     *string       `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string       `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string       `json:"company_id,omitempty"`   // 公司 ID
	Timestamp    *int64        `json:"timestamp,omitempty"`    // 当前时间戳
	Sign         *string       `json:"sign,omitempty"`         // 签名
	ParamJson    *string       `json:"param_json,omitempty"`   // 请求参数，(将 param_json 转为 json 字符串)
	ParamJsonObj *TravelerInfo `json:"param_json__obj__,omitempty"`
}

type UpdateTravelerRequestBuilder struct {
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
	paramJson       string // 请求参数，(将 param_json 转为 json 字符串)
	paramJsonSet    bool
	paramJsonObj    TravelerInfo
	paramJsonObjSet bool
}

func NewUpdateTravelerRequestBuilder() *UpdateTravelerRequestBuilder {
	return &UpdateTravelerRequestBuilder{}
}
func (builder *UpdateTravelerRequestBuilder) ClientId(clientId string) *UpdateTravelerRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) AccessToken(accessToken string) *UpdateTravelerRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) CompanyId(companyId string) *UpdateTravelerRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) Timestamp(timestamp int64) *UpdateTravelerRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) Sign(sign string) *UpdateTravelerRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) ParamJson(paramJson string) *UpdateTravelerRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *UpdateTravelerRequestBuilder) ParamJsonObj(paramJsonObj TravelerInfo) *UpdateTravelerRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *UpdateTravelerRequestBuilder) Build() *UpdateTravelerRequest {
	data := &UpdateTravelerRequest{}
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
