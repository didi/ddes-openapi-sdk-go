package v1

// UpdateExtendStatusRequest struct for UpdateExtendStatusRequest
type UpdateExtendStatusRequest struct {
	ClientId    *string      `json:"client_id,omitempty"`        // 申请应用时分配的AppKey
	AccessToken *string      `json:"access_token,omitempty"`     // 授权后的access token
	CompanyId   *string      `json:"company_id,omitempty"`       // 企业ID
	Timestamp   *int64       `json:"timestamp,omitempty"`        // 当前时间戳，精确到秒级
	Sign        *string      `json:"sign,omitempty"`             // 签名
	RootCode    *string      `json:"root_code,omitempty"`        // 档案code
	RootStatus  *int32       `json:"root_status,omitempty"`      // 状态;1为正常，2为停用，3为删除
	ItemList    *string      `json:"item_list,omitempty"`        // 子档案列表，json字符串
	ItemListObj []ExtendInfo `json:"item_list__obj__,omitempty"` // 子档案列表Obj(脚本转为json后赋值给item_list字段)
}

type UpdateExtendStatusRequestBuilder struct {
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
	rootCode       string // 档案code
	rootCodeSet    bool
	rootStatus     int32 // 状态;1为正常，2为停用，3为删除
	rootStatusSet  bool
	itemList       string // 子档案列表，json字符串
	itemListSet    bool
	itemListObj    []ExtendInfo // 子档案列表Obj(脚本转为json后赋值给item_list字段)
	itemListObjSet bool
}

func NewUpdateExtendStatusRequestBuilder() *UpdateExtendStatusRequestBuilder {
	return &UpdateExtendStatusRequestBuilder{}
}
func (builder *UpdateExtendStatusRequestBuilder) ClientId(clientId string) *UpdateExtendStatusRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) AccessToken(accessToken string) *UpdateExtendStatusRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) CompanyId(companyId string) *UpdateExtendStatusRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) Timestamp(timestamp int64) *UpdateExtendStatusRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) Sign(sign string) *UpdateExtendStatusRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) RootCode(rootCode string) *UpdateExtendStatusRequestBuilder {
	builder.rootCode = rootCode
	builder.rootCodeSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) RootStatus(rootStatus int32) *UpdateExtendStatusRequestBuilder {
	builder.rootStatus = rootStatus
	builder.rootStatusSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) ItemList(itemList string) *UpdateExtendStatusRequestBuilder {
	builder.itemList = itemList
	builder.itemListSet = true
	return builder
}
func (builder *UpdateExtendStatusRequestBuilder) ItemListObj(itemListObj []ExtendInfo) *UpdateExtendStatusRequestBuilder {
	builder.itemListObj = itemListObj
	builder.itemListObjSet = true
	return builder
}

func (builder *UpdateExtendStatusRequestBuilder) Build() *UpdateExtendStatusRequest {
	data := &UpdateExtendStatusRequest{}
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
	if builder.rootStatusSet {
		data.RootStatus = &builder.rootStatus
	}
	if builder.itemListSet {
		data.ItemList = &builder.itemList
	}
	if builder.itemListObjSet {
		data.ItemListObj = builder.itemListObj
	}
	return data
}
