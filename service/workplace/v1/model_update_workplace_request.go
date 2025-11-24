package v1

// UpdateWorkplaceRequest struct for UpdateWorkplaceRequest
type UpdateWorkplaceRequest struct {
	ClientId     *string        `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string        `json:"access_token,omitempty"` // 授权后的access token
	Timestamp    *int64         `json:"timestamp,omitempty"`    // 当前时间戳,精确到秒级
	CompanyId    *string        `json:"company_id,omitempty"`   // 公司 ID
	Sign         *string        `json:"sign,omitempty"`         // 签名
	ParamJson    *string        `json:"param_json,omitempty"`   // 查询信息，(将 param_json 转为 json 字符串)
	ParamJsonObj *WorkplaceInfo `json:"param_json__obj__,omitempty"`
}

type UpdateWorkplaceRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	timestamp       int64 // 当前时间戳,精确到秒级
	timestampSet    bool
	companyId       string // 公司 ID
	companyIdSet    bool
	sign            string // 签名
	signSet         bool
	paramJson       string // 查询信息，(将 param_json 转为 json 字符串)
	paramJsonSet    bool
	paramJsonObj    WorkplaceInfo
	paramJsonObjSet bool
}

func NewUpdateWorkplaceRequestBuilder() *UpdateWorkplaceRequestBuilder {
	return &UpdateWorkplaceRequestBuilder{}
}
func (builder *UpdateWorkplaceRequestBuilder) ClientId(clientId string) *UpdateWorkplaceRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) AccessToken(accessToken string) *UpdateWorkplaceRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) Timestamp(timestamp int64) *UpdateWorkplaceRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) CompanyId(companyId string) *UpdateWorkplaceRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) Sign(sign string) *UpdateWorkplaceRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) ParamJson(paramJson string) *UpdateWorkplaceRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *UpdateWorkplaceRequestBuilder) ParamJsonObj(paramJsonObj WorkplaceInfo) *UpdateWorkplaceRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *UpdateWorkplaceRequestBuilder) Build() *UpdateWorkplaceRequest {
	data := &UpdateWorkplaceRequest{}
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
	if builder.paramJsonSet {
		data.ParamJson = &builder.paramJson
	}
	if builder.paramJsonObjSet {
		data.ParamJsonObj = &builder.paramJsonObj
	}
	return data
}
