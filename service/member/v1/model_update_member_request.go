package v1

// UpdateMemberRequest struct for UpdateMemberRequest
type UpdateMemberRequest struct {
	ClientId       *string     `json:"client_id,omitempty"`       // 申请应用时分配的AppKey
	AccessToken    *string     `json:"access_token,omitempty"`    // 授权后的access token
	CompanyId      *string     `json:"company_id,omitempty"`      // 企业ID
	Timestamp      *int32      `json:"timestamp,omitempty"`       // 当前时间戳(精确到秒级)
	Sign           *string     `json:"sign,omitempty"`            // 签名
	MemberId       *int64      `json:"member_id,omitempty"`       // 员工在滴滴企业的ID，member_id或employee_number不能同时为空。优先级member_id高于employee_number
	EmployeeNumber *string     `json:"employee_number,omitempty"` // 员工工号
	HasCardInfo    *int32      `json:"has_card_info,omitempty"`   // 是否含有证件信息，是否含有证件信息，当传证件信息时，此字符传1，其他情况不传或传0
	Data           *string     `json:"data,omitempty"`            // 员工信息，详见 data
	DataObj        *MemberInfo `json:"data__obj__,omitempty"`
}

type UpdateMemberRequestBuilder struct {
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
	memberId          int64 // 员工在滴滴企业的ID，member_id或employee_number不能同时为空。优先级member_id高于employee_number
	memberIdSet       bool
	employeeNumber    string // 员工工号
	employeeNumberSet bool
	hasCardInfo       int32 // 是否含有证件信息，是否含有证件信息，当传证件信息时，此字符传1，其他情况不传或传0
	hasCardInfoSet    bool
	data              string // 员工信息，详见 data
	dataSet           bool
	dataObj           MemberInfo
	dataObjSet        bool
}

func NewUpdateMemberRequestBuilder() *UpdateMemberRequestBuilder {
	return &UpdateMemberRequestBuilder{}
}
func (builder *UpdateMemberRequestBuilder) ClientId(clientId string) *UpdateMemberRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) AccessToken(accessToken string) *UpdateMemberRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) CompanyId(companyId string) *UpdateMemberRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) Timestamp(timestamp int32) *UpdateMemberRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) Sign(sign string) *UpdateMemberRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) MemberId(memberId int64) *UpdateMemberRequestBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) EmployeeNumber(employeeNumber string) *UpdateMemberRequestBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) HasCardInfo(hasCardInfo int32) *UpdateMemberRequestBuilder {
	builder.hasCardInfo = hasCardInfo
	builder.hasCardInfoSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) Data(data string) *UpdateMemberRequestBuilder {
	builder.data = data
	builder.dataSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) DataObj(dataObj MemberInfo) *UpdateMemberRequestBuilder {
	builder.dataObj = dataObj
	builder.dataObjSet = true
	return builder
}

func (builder *UpdateMemberRequestBuilder) Build() *UpdateMemberRequest {
	data := &UpdateMemberRequest{}
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
