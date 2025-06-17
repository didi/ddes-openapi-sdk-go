package v1

// DelBudgetCenterRequest struct for DelBudgetCenterRequest
type DelBudgetCenterRequest struct {
	ClientId    *string `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"`  // 授权后的access token
	CompanyId   *string `json:"company_id,omitempty"`    // 企业ID
	Timestamp   *int32  `json:"timestamp,omitempty"`     // 当前时间戳(精确到秒级)
	Sign        *string `json:"sign,omitempty"`          // 签名
	Type        *int32  `json:"type,omitempty"`          // 类型，推测为 1 代表部门，2 代表项目，未传时可根据业务规则处理
	Id          *string `json:"id,omitempty"`            // 需要删除的部门或者项目的id，删除部门时，id和out_budget_id优先处理id；删除项目时，id和out_budget_id与name组合唯一值时，优先处理id
	Name        *string `json:"name,omitempty"`          // 部门/项目名称，不大于 200 字符，使用out_budget_id删除项目时，项目名称需要一起提供
	OutBudgetId *string `json:"out_budget_id,omitempty"` // 编号，type = 1 时必填；type = 2 非必填；长度限制：≤ 64 字符
}

type DelBudgetCenterRequestBuilder struct {
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
	type_          int32 // 类型，推测为 1 代表部门，2 代表项目，未传时可根据业务规则处理
	type_Set       bool
	id             string // 需要删除的部门或者项目的id，删除部门时，id和out_budget_id优先处理id；删除项目时，id和out_budget_id与name组合唯一值时，优先处理id
	idSet          bool
	name           string // 部门/项目名称，不大于 200 字符，使用out_budget_id删除项目时，项目名称需要一起提供
	nameSet        bool
	outBudgetId    string // 编号，type = 1 时必填；type = 2 非必填；长度限制：≤ 64 字符
	outBudgetIdSet bool
}

func NewDelBudgetCenterRequestBuilder() *DelBudgetCenterRequestBuilder {
	return &DelBudgetCenterRequestBuilder{}
}
func (builder *DelBudgetCenterRequestBuilder) ClientId(clientId string) *DelBudgetCenterRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) AccessToken(accessToken string) *DelBudgetCenterRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) CompanyId(companyId string) *DelBudgetCenterRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) Timestamp(timestamp int32) *DelBudgetCenterRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) Sign(sign string) *DelBudgetCenterRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) Type(type_ int32) *DelBudgetCenterRequestBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) Id(id string) *DelBudgetCenterRequestBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) Name(name string) *DelBudgetCenterRequestBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *DelBudgetCenterRequestBuilder) OutBudgetId(outBudgetId string) *DelBudgetCenterRequestBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}

func (builder *DelBudgetCenterRequestBuilder) Build() *DelBudgetCenterRequest {
	data := &DelBudgetCenterRequest{}
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
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	return data
}
