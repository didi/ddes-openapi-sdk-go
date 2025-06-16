package v1

// CreateMemberRequest struct for CreateMemberRequest
type CreateMemberRequest struct {
	ClientId    *string     `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	AccessToken *string     `json:"access_token,omitempty"`  // 授权后的access token
	CompanyId   *string     `json:"company_id,omitempty"`    // 企业ID
	Timestamp   *int32      `json:"timestamp,omitempty"`     // 当前时间戳(精确到秒级)
	Sign        *string     `json:"sign,omitempty"`          // 签名
	SendMessage *int32      `json:"send_message,omitempty"`  // 是否发短信配置，枚举值数字 0 不发送，1 发送， 员工添加成功后发短信（包含双向确认短信） 默认值为0
	HasCardInfo *int32      `json:"has_card_info,omitempty"` // 是否含有证件信息，当传证件信息时，此字符传1，其他情况不传或传0
	Data        *string     `json:"data,omitempty"`          // 员工信息，详见 data
	DataObj     *MemberInfo `json:"data__obj__,omitempty"`
}

type CreateMemberRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	companyId      string // 企业ID
	companyIdSet   bool
	timestamp      int32 // 当前时间戳(精确到秒级)
	timestampSet   bool
	sign           string // 签名
	signSet        bool
	sendMessage    int32 // 是否发短信配置，枚举值数字 0 不发送，1 发送， 员工添加成功后发短信（包含双向确认短信） 默认值为0
	sendMessageSet bool
	hasCardInfo    int32 // 是否含有证件信息，当传证件信息时，此字符传1，其他情况不传或传0
	hasCardInfoSet bool
	data           string // 员工信息，详见 data
	dataSet        bool
	dataObj        MemberInfo
	dataObjSet     bool
}

func NewCreateMemberRequestBuilder() *CreateMemberRequestBuilder {
	return &CreateMemberRequestBuilder{}
}
func (builder *CreateMemberRequestBuilder) ClientId(clientId string) *CreateMemberRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) AccessToken(accessToken string) *CreateMemberRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) CompanyId(companyId string) *CreateMemberRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) Timestamp(timestamp int32) *CreateMemberRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) Sign(sign string) *CreateMemberRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) SendMessage(sendMessage int32) *CreateMemberRequestBuilder {
	builder.sendMessage = sendMessage
	builder.sendMessageSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) HasCardInfo(hasCardInfo int32) *CreateMemberRequestBuilder {
	builder.hasCardInfo = hasCardInfo
	builder.hasCardInfoSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) Data(data string) *CreateMemberRequestBuilder {
	builder.data = data
	builder.dataSet = true
	return builder
}
func (builder *CreateMemberRequestBuilder) DataObj(dataObj MemberInfo) *CreateMemberRequestBuilder {
	builder.dataObj = dataObj
	builder.dataObjSet = true
	return builder
}

func (builder *CreateMemberRequestBuilder) Build() *CreateMemberRequest {
	data := &CreateMemberRequest{}
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
	if builder.sendMessageSet {
		data.SendMessage = &builder.sendMessage
	}
	if builder.hasCardInfoSet {
		data.HasCardInfo = &builder.hasCardInfo
	}
	if builder.dataSet {
		data.Data = &builder.data
	}
	if builder.dataObjSet {
		data.DataObj = &builder.dataObj
	}
	return data
}
