package v1

// CreateRankRequest struct for CreateRankRequest
type CreateRankRequest struct {
	ClientId     *string   `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string   `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string   `json:"company_id,omitempty"`   // 公司 ID
	Timestamp    *int64    `json:"timestamp,omitempty"`    // 当前时间戳
	Sign         *string   `json:"sign,omitempty"`         // 签名
	ParamJson    *string   `json:"param_json,omitempty"`   // 请求参数，(将 param_json 转为 json 字符串)
	ParamJsonObj *RankInfo `json:"param_json__obj__,omitempty"`
}

type CreateRankRequestBuilder struct {
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
	paramJsonObj    RankInfo
	paramJsonObjSet bool
}

func NewCreateRankRequestBuilder() *CreateRankRequestBuilder {
	return &CreateRankRequestBuilder{}
}
func (builder *CreateRankRequestBuilder) ClientId(clientId string) *CreateRankRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) AccessToken(accessToken string) *CreateRankRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) CompanyId(companyId string) *CreateRankRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) Timestamp(timestamp int64) *CreateRankRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) Sign(sign string) *CreateRankRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) ParamJson(paramJson string) *CreateRankRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *CreateRankRequestBuilder) ParamJsonObj(paramJsonObj RankInfo) *CreateRankRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *CreateRankRequestBuilder) Build() *CreateRankRequest {
	data := &CreateRankRequest{}
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
