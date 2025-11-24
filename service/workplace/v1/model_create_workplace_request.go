package v1

// CreateWorkplaceRequest struct for CreateWorkplaceRequest
type CreateWorkplaceRequest struct {
	ClientId     *string        `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string        `json:"access_token,omitempty"` // 授权后的access token
	Timestamp    *int64         `json:"timestamp,omitempty"`    // 当前时间戳,精确到秒级
	CompanyId    *string        `json:"company_id,omitempty"`   // 公司 ID
	Sign         *string        `json:"sign,omitempty"`         // 签名
	ParamJson    *string        `json:"param_json,omitempty"`   // 查询信息，(将 param_json 转为 json 字符串)
	ParamJsonObj *WorkplaceInfo `json:"param_json__obj__,omitempty"`
}

type CreateWorkplaceRequestBuilder struct {
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

func NewCreateWorkplaceRequestBuilder() *CreateWorkplaceRequestBuilder {
	return &CreateWorkplaceRequestBuilder{}
}
func (builder *CreateWorkplaceRequestBuilder) ClientId(clientId string) *CreateWorkplaceRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) AccessToken(accessToken string) *CreateWorkplaceRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) Timestamp(timestamp int64) *CreateWorkplaceRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) CompanyId(companyId string) *CreateWorkplaceRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) Sign(sign string) *CreateWorkplaceRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) ParamJson(paramJson string) *CreateWorkplaceRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *CreateWorkplaceRequestBuilder) ParamJsonObj(paramJsonObj WorkplaceInfo) *CreateWorkplaceRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *CreateWorkplaceRequestBuilder) Build() *CreateWorkplaceRequest {
	data := &CreateWorkplaceRequest{}
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
