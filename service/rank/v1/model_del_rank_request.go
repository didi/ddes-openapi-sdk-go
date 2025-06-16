package v1

// DelRankRequest struct for DelRankRequest
type DelRankRequest struct {
	ClientId     *string   `json:"client_id,omitempty"`    // 申请应用时分配的AppKey
	AccessToken  *string   `json:"access_token,omitempty"` // 授权后的access token
	CompanyId    *string   `json:"company_id,omitempty"`   // 公司ID(滴滴侧公司ID,托管模式下为操作的公司ID，同时需要开启托管的员工管理权限,非托管模式下为自己公司ID)
	Timestamp    *int64    `json:"timestamp,omitempty"`    // 当前时间戳
	ParamJson    *string   `json:"param_json,omitempty"`   // 请求参数，(将 param_json 转为 json 字符串)
	Sign         *string   `json:"sign,omitempty"`         // 签名
	ParamJsonObj *RankInfo `json:"param_json__obj__,omitempty"`
}

type DelRankRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	companyId       string // 公司ID(滴滴侧公司ID,托管模式下为操作的公司ID，同时需要开启托管的员工管理权限,非托管模式下为自己公司ID)
	companyIdSet    bool
	timestamp       int64 // 当前时间戳
	timestampSet    bool
	paramJson       string // 请求参数，(将 param_json 转为 json 字符串)
	paramJsonSet    bool
	sign            string // 签名
	signSet         bool
	paramJsonObj    RankInfo
	paramJsonObjSet bool
}

func NewDelRankRequestBuilder() *DelRankRequestBuilder {
	return &DelRankRequestBuilder{}
}
func (builder *DelRankRequestBuilder) ClientId(clientId string) *DelRankRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *DelRankRequestBuilder) AccessToken(accessToken string) *DelRankRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *DelRankRequestBuilder) CompanyId(companyId string) *DelRankRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *DelRankRequestBuilder) Timestamp(timestamp int64) *DelRankRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *DelRankRequestBuilder) ParamJson(paramJson string) *DelRankRequestBuilder {
	builder.paramJson = paramJson
	builder.paramJsonSet = true
	return builder
}
func (builder *DelRankRequestBuilder) Sign(sign string) *DelRankRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *DelRankRequestBuilder) ParamJsonObj(paramJsonObj RankInfo) *DelRankRequestBuilder {
	builder.paramJsonObj = paramJsonObj
	builder.paramJsonObjSet = true
	return builder
}

func (builder *DelRankRequestBuilder) Build() *DelRankRequest {
	data := &DelRankRequest{}
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
	if builder.paramJsonSet {
		data.ParamJson = &builder.paramJson
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	if builder.paramJsonObjSet {
		data.ParamJsonObj = &builder.paramJsonObj
	}
	return data
}
