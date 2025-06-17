package v1

// CreateExtendBatchRequest struct for CreateExtendBatchRequest
type CreateExtendBatchRequest struct {
	ClientId    *string      `json:"client_id,omitempty"`        // 申请应用时分配的AppKey
	AccessToken *string      `json:"access_token,omitempty"`     // 授权后的access token
	CompanyId   *string      `json:"company_id,omitempty"`       // 企业ID
	Timestamp   *int64       `json:"timestamp,omitempty"`        // 当前时间戳，精确到秒级
	Sign        *string      `json:"sign,omitempty"`             // 签名
	RootCode    *string      `json:"root_code,omitempty"`        // 档案code，150字符以内
	RootName    *string      `json:"root_name,omitempty"`        // 档案名称，150字符以内
	ItemList    *string      `json:"item_list,omitempty"`        // 子档案列表，json字符串
	ItemListObj []ExtendInfo `json:"item_list__obj__,omitempty"` // 子档案列表Obj(脚本转为json后赋值给item_list字段)
}

type CreateExtendBatchRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	companyId      string // 企业ID
	companyIdSet   bool
	timestamp      int64 // 当前时间戳，精确到秒级
	timestampSet   bool
	sign           string // 签名
	signSet        bool
	rootCode       string // 档案code，150字符以内
	rootCodeSet    bool
	rootName       string // 档案名称，150字符以内
	rootNameSet    bool
	itemList       string // 子档案列表，json字符串
	itemListSet    bool
	itemListObj    []ExtendInfo // 子档案列表Obj(脚本转为json后赋值给item_list字段)
	itemListObjSet bool
}

func NewCreateExtendBatchRequestBuilder() *CreateExtendBatchRequestBuilder {
	return &CreateExtendBatchRequestBuilder{}
}
func (builder *CreateExtendBatchRequestBuilder) ClientId(clientId string) *CreateExtendBatchRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) AccessToken(accessToken string) *CreateExtendBatchRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) CompanyId(companyId string) *CreateExtendBatchRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) Timestamp(timestamp int64) *CreateExtendBatchRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) Sign(sign string) *CreateExtendBatchRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) RootCode(rootCode string) *CreateExtendBatchRequestBuilder {
	builder.rootCode = rootCode
	builder.rootCodeSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) RootName(rootName string) *CreateExtendBatchRequestBuilder {
	builder.rootName = rootName
	builder.rootNameSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) ItemList(itemList string) *CreateExtendBatchRequestBuilder {
	builder.itemList = itemList
	builder.itemListSet = true
	return builder
}
func (builder *CreateExtendBatchRequestBuilder) ItemListObj(itemListObj []ExtendInfo) *CreateExtendBatchRequestBuilder {
	builder.itemListObj = itemListObj
	builder.itemListObjSet = true
	return builder
}

func (builder *CreateExtendBatchRequestBuilder) Build() *CreateExtendBatchRequest {
	data := &CreateExtendBatchRequest{}
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
	if builder.rootCodeSet {
		data.RootCode = &builder.rootCode
	}
	if builder.rootNameSet {
		data.RootName = &builder.rootName
	}
	if builder.itemListSet {
		data.ItemList = &builder.itemList
	}
	if builder.itemListObjSet {
		data.ItemListObj = builder.itemListObj
	}
	return data
}
