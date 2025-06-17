package v1

// DelMemberRequest struct for DelMemberRequest
type DelMemberRequest struct {
	ClientId       *string `json:"client_id,omitempty"`       // 申请应用时分配的AppKey
	AccessToken    *string `json:"access_token,omitempty"`    // 授权后的access token
	CompanyId      *string `json:"company_id,omitempty"`      // 企业ID
	Timestamp      *int32  `json:"timestamp,omitempty"`       // 当前时间戳(精确到秒级)
	Sign           *string `json:"sign,omitempty"`            // 签名
	MemberId       *string `json:"member_id,omitempty"`       // 员工在滴滴企业的ID，员工在滴滴企业的ID（同员工新增接口中返回的id；删除多个员工时，员工号需要以\"_\"隔开，一次最多删除100条）member_id与employee_number同时存在以member_id为准
	EmployeeNumber *string `json:"employee_number,omitempty"` // 员工在滴滴企业的工号，json 字符串（一次最多删除100条），member_id和employee_number仅有一个生效，不能同时为空，member_id 优先级高于employee_number, 示例： [\"D1001\", \"D1002\", \"D1003\"]
}

type DelMemberRequestBuilder struct {
	clientId          string // 申请应用时分配的AppKey
	clientIdSet       bool
	accessToken       string // 授权后的access token
	accessTokenSet    bool
	companyId         string // 企业ID
	companyIdSet      bool
	timestamp         int32 // 当前时间戳(精确到秒级)
	timestampSet      bool
	sign              string // 签名
	signSet           bool
	memberId          string // 员工在滴滴企业的ID，员工在滴滴企业的ID（同员工新增接口中返回的id；删除多个员工时，员工号需要以\"_\"隔开，一次最多删除100条）member_id与employee_number同时存在以member_id为准
	memberIdSet       bool
	employeeNumber    string // 员工在滴滴企业的工号，json 字符串（一次最多删除100条），member_id和employee_number仅有一个生效，不能同时为空，member_id 优先级高于employee_number, 示例： [\"D1001\", \"D1002\", \"D1003\"]
	employeeNumberSet bool
}

func NewDelMemberRequestBuilder() *DelMemberRequestBuilder {
	return &DelMemberRequestBuilder{}
}
func (builder *DelMemberRequestBuilder) ClientId(clientId string) *DelMemberRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) AccessToken(accessToken string) *DelMemberRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) CompanyId(companyId string) *DelMemberRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) Timestamp(timestamp int32) *DelMemberRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) Sign(sign string) *DelMemberRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) MemberId(memberId string) *DelMemberRequestBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *DelMemberRequestBuilder) EmployeeNumber(employeeNumber string) *DelMemberRequestBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}

func (builder *DelMemberRequestBuilder) Build() *DelMemberRequest {
	data := &DelMemberRequest{}
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
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	return data
}
