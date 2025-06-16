package v1

// ListOrderRequest struct for ListOrderRequest
type ListOrderRequest struct {
	ClientId     *string       `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string       `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string       `json:"company_id,omitempty"`   // 企业ID
	Timestamp    *int64        `json:"timestamp,omitempty"`    // 当前时间戳(精确到秒级)
	Sign         *string       `json:"sign,omitempty"`         // 签名
	ParamJson    *string       `json:"param_json,omitempty"`   // 查询信息
	ParamJsonObj *ParamJsonObj `json:"param_json__obj__,omitempty"`
}

type ListOrderRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	companyId       string // 企业ID
	companyIdSet    bool
	timestamp       int64 // 当前时间戳(精确到秒级)
	timestampSet    bool
	sign            string // 签名
	signSet         bool
	paramJson       string // 查询信息
	paramJsonSet    bool
	paramJsonObj    ParamJsonObj
	paramJsonObjSet bool
}

func NewListOrderRequestBuilder() *ListOrderRequestBuilder {
	return &ListOrderRequestBuilder{}
}
func (builder *ListOrderRequestBuilder) ClientId(clientId string) *ListOrderRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) AccessToken(accessToken string) *ListOrderRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) CompanyId(companyId string) *ListOrderRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) Timestamp(timestamp int64) *ListOrderRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) Sign(sign string) *ListOrderRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) ParamJson(paramJson string) *ListOrderRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *ListOrderRequestBuilder) ParamJsonObj(paramJsonObj ParamJsonObj) *ListOrderRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *ListOrderRequestBuilder) Build() *ListOrderRequest {
	data := &ListOrderRequest{}
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
