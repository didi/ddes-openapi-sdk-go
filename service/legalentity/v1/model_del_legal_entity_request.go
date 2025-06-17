package v1

// DelLegalEntityRequest struct for DelLegalEntityRequest
type DelLegalEntityRequest struct {
	ClientId         *string `json:"client_id,omitempty"`           // 申请应用时分配的AppKey
	AccessToken      *string `json:"access_token,omitempty"`        // 授权后的access token
	CompanyId        *string `json:"company_id,omitempty"`          // 企业ID
	Timestamp        *int32  `json:"timestamp,omitempty"`           // 当前时间戳(精确到秒级)
	Sign             *string `json:"sign,omitempty"`                // 签名
	LegalEntityId    *string `json:"legal_entity_id,omitempty"`     // 需要删除的公司主体滴滴id，公司主体滴滴id和外部公司主体编号二者之一必填
	OutLegalEntityId *string `json:"out_legal_entity_id,omitempty"` // 需要删除的外部公司主体编号，不可重复，滴滴后台有效的公司ID不重复
}

type DelLegalEntityRequestBuilder struct {
	clientId            string // 申请应用时分配的AppKey
	clientIdSet         bool
	accessToken         string // 授权后的access token
	accessTokenSet      bool
	companyId           string // 企业ID
	companyIdSet        bool
	timestamp           int32 // 当前时间戳(精确到秒级)
	timestampSet        bool
	sign                string // 签名
	signSet             bool
	legalEntityId       string // 需要删除的公司主体滴滴id，公司主体滴滴id和外部公司主体编号二者之一必填
	legalEntityIdSet    bool
	outLegalEntityId    string // 需要删除的外部公司主体编号，不可重复，滴滴后台有效的公司ID不重复
	outLegalEntityIdSet bool
}

func NewDelLegalEntityRequestBuilder() *DelLegalEntityRequestBuilder {
	return &DelLegalEntityRequestBuilder{}
}
func (builder *DelLegalEntityRequestBuilder) ClientId(clientId string) *DelLegalEntityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) AccessToken(accessToken string) *DelLegalEntityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) CompanyId(companyId string) *DelLegalEntityRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) Timestamp(timestamp int32) *DelLegalEntityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) Sign(sign string) *DelLegalEntityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) LegalEntityId(legalEntityId string) *DelLegalEntityRequestBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *DelLegalEntityRequestBuilder) OutLegalEntityId(outLegalEntityId string) *DelLegalEntityRequestBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}

func (builder *DelLegalEntityRequestBuilder) Build() *DelLegalEntityRequest {
	data := &DelLegalEntityRequest{}
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
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	return data
}
