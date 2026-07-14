package v1

// UpdateMemberRequest 绑定项目与人员关系请求体
type UpdateMemberRequest struct {
	ClientId             *string `json:"client_id,omitempty"`              // 申请应用时分配的AppKey
	AccessToken          *string `json:"access_token,omitempty"`           // 授权后的access token
	CompanyId            *string `json:"company_id,omitempty"`             // 企业ID
	Timestamp            *int64  `json:"timestamp,omitempty"`              // 当前时间戳，精确到秒级
	ProjectId            *string `json:"project_id,omitempty"`             // 项目ID
	ProjectCode          *string `json:"project_code,omitempty"`           // 项目编码
	ProjectName          *string `json:"project_name,omitempty"`           // 项目名称
	MemberIds            *string `json:"member_ids,omitempty"`             // 员工ID，多个用逗号分隔，最多支持100个
	MemberType           *int32  `json:"member_type,omitempty"`            // 员工信息类型：0=手机号，1=员工编号，2=邮箱
	MemberValues         *string `json:"member_values,omitempty"`          // 员工信息数据，对应member_type，多个用逗号分隔，最多支持100个
	BelongEnterpriseName *string `json:"belong_enterprise_name,omitempty"` // 归属企业名称（集团账户参数）
	TaxpayerNo           *string `json:"taxpayer_no,omitempty"`            // 纳税人识别号（集团账户参数）
	Sign                 *string `json:"sign,omitempty"`                   // 签名
}

type UpdateMemberRequestBuilder struct {
	clientId                string
	clientIdSet             bool
	accessToken             string
	accessTokenSet          bool
	companyId               string
	companyIdSet            bool
	timestamp               int64
	timestampSet            bool
	projectId               string
	projectIdSet            bool
	projectCode             string
	projectCodeSet          bool
	projectName             string
	projectNameSet          bool
	memberIds               string
	memberIdsSet            bool
	memberType              int32
	memberTypeSet           bool
	memberValues            string
	memberValuesSet         bool
	belongEnterpriseName    string
	belongEnterpriseNameSet bool
	taxpayerNo              string
	taxpayerNoSet           bool
	sign                    string
	signSet                 bool
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
func (builder *UpdateMemberRequestBuilder) Timestamp(timestamp int64) *UpdateMemberRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) ProjectId(projectId string) *UpdateMemberRequestBuilder {
	builder.projectId = projectId
	builder.projectIdSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) ProjectCode(projectCode string) *UpdateMemberRequestBuilder {
	builder.projectCode = projectCode
	builder.projectCodeSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) ProjectName(projectName string) *UpdateMemberRequestBuilder {
	builder.projectName = projectName
	builder.projectNameSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) MemberIds(memberIds string) *UpdateMemberRequestBuilder {
	builder.memberIds = memberIds
	builder.memberIdsSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) MemberType(memberType int32) *UpdateMemberRequestBuilder {
	builder.memberType = memberType
	builder.memberTypeSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) MemberValues(memberValues string) *UpdateMemberRequestBuilder {
	builder.memberValues = memberValues
	builder.memberValuesSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) BelongEnterpriseName(belongEnterpriseName string) *UpdateMemberRequestBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) TaxpayerNo(taxpayerNo string) *UpdateMemberRequestBuilder {
	builder.taxpayerNo = taxpayerNo
	builder.taxpayerNoSet = true
	return builder
}
func (builder *UpdateMemberRequestBuilder) Sign(sign string) *UpdateMemberRequestBuilder {
	builder.sign = sign
	builder.signSet = true
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
	if builder.projectIdSet {
		data.ProjectId = &builder.projectId
	}
	if builder.projectCodeSet {
		data.ProjectCode = &builder.projectCode
	}
	if builder.projectNameSet {
		data.ProjectName = &builder.projectName
	}
	if builder.memberIdsSet {
		data.MemberIds = &builder.memberIds
	}
	if builder.memberTypeSet {
		data.MemberType = &builder.memberType
	}
	if builder.memberValuesSet {
		data.MemberValues = &builder.memberValues
	}
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	if builder.taxpayerNoSet {
		data.TaxpayerNo = &builder.taxpayerNo
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	return data
}
