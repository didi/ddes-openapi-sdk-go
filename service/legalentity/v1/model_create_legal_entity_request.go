package v1

// CreateLegalEntityRequest struct for CreateLegalEntityRequest
type CreateLegalEntityRequest struct {
	ClientId         *string `json:"client_id,omitempty"`           // 申请应用时分配的AppKey
	AccessToken      *string `json:"access_token,omitempty"`        // 授权后的access token
	CompanyId        *string `json:"company_id,omitempty"`          // 企业ID
	Timestamp        *int32  `json:"timestamp,omitempty"`           // 当前时间戳(精确到秒级)
	Sign             *string `json:"sign,omitempty"`                // 签名
	Name             *string `json:"name,omitempty"`                // 公司主体名称
	OutLegalEntityId *string `json:"out_legal_entity_id,omitempty"` // 外部公司主体编号，不可重复
	ParentId         *int32  `json:"parent_id,omitempty"`           // 上级公司id，新建公司信息之后，滴滴返回的公司ID
	StartTime        *string `json:"start_time,omitempty"`          // 公司主体有效期开始日期，不传 默认为当前日期 格式：yyyy-MM-dd HH:mm:ss
	EndTime          *string `json:"end_time,omitempty"`            // 公司主体有效期结束日期，不传默认为长期有效 格式：yyyy-MM-dd HH:mm:ss 需大于当前时间
	TaxpayerNo       *string `json:"taxpayer_no,omitempty"`         // 纳税人识别号，长度为 15/18/20 位
	Address          *string `json:"address,omitempty"`             // 公司主体注册地址
	Telephone        *string `json:"telephone,omitempty"`           // 公司主体注册电话
	OpenBank         *string `json:"open_bank,omitempty"`           // 开户行名称
	BankCardNo       *string `json:"bank_card_no,omitempty"`        // 开户行账号，纯数字
}

type CreateLegalEntityRequestBuilder struct {
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
	name                string // 公司主体名称
	nameSet             bool
	outLegalEntityId    string // 外部公司主体编号，不可重复
	outLegalEntityIdSet bool
	parentId            int32 // 上级公司id，新建公司信息之后，滴滴返回的公司ID
	parentIdSet         bool
	startTime           string // 公司主体有效期开始日期，不传 默认为当前日期 格式：yyyy-MM-dd HH:mm:ss
	startTimeSet        bool
	endTime             string // 公司主体有效期结束日期，不传默认为长期有效 格式：yyyy-MM-dd HH:mm:ss 需大于当前时间
	endTimeSet          bool
	taxpayerNo          string // 纳税人识别号，长度为 15/18/20 位
	taxpayerNoSet       bool
	address             string // 公司主体注册地址
	addressSet          bool
	telephone           string // 公司主体注册电话
	telephoneSet        bool
	openBank            string // 开户行名称
	openBankSet         bool
	bankCardNo          string // 开户行账号，纯数字
	bankCardNoSet       bool
}

func NewCreateLegalEntityRequestBuilder() *CreateLegalEntityRequestBuilder {
	return &CreateLegalEntityRequestBuilder{}
}
func (builder *CreateLegalEntityRequestBuilder) ClientId(clientId string) *CreateLegalEntityRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) AccessToken(accessToken string) *CreateLegalEntityRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) CompanyId(companyId string) *CreateLegalEntityRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) Timestamp(timestamp int32) *CreateLegalEntityRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) Sign(sign string) *CreateLegalEntityRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) Name(name string) *CreateLegalEntityRequestBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) OutLegalEntityId(outLegalEntityId string) *CreateLegalEntityRequestBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) ParentId(parentId int32) *CreateLegalEntityRequestBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) StartTime(startTime string) *CreateLegalEntityRequestBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) EndTime(endTime string) *CreateLegalEntityRequestBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) TaxpayerNo(taxpayerNo string) *CreateLegalEntityRequestBuilder {
	builder.taxpayerNo = taxpayerNo
	builder.taxpayerNoSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) Address(address string) *CreateLegalEntityRequestBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) Telephone(telephone string) *CreateLegalEntityRequestBuilder {
	builder.telephone = telephone
	builder.telephoneSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) OpenBank(openBank string) *CreateLegalEntityRequestBuilder {
	builder.openBank = openBank
	builder.openBankSet = true
	return builder
}
func (builder *CreateLegalEntityRequestBuilder) BankCardNo(bankCardNo string) *CreateLegalEntityRequestBuilder {
	builder.bankCardNo = bankCardNo
	builder.bankCardNoSet = true
	return builder
}

func (builder *CreateLegalEntityRequestBuilder) Build() *CreateLegalEntityRequest {
	data := &CreateLegalEntityRequest{}
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
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.taxpayerNoSet {
		data.TaxpayerNo = &builder.taxpayerNo
	}
	if builder.addressSet {
		data.Address = &builder.address
	}
	if builder.telephoneSet {
		data.Telephone = &builder.telephone
	}
	if builder.openBankSet {
		data.OpenBank = &builder.openBank
	}
	if builder.bankCardNoSet {
		data.BankCardNo = &builder.bankCardNo
	}
	return data
}
