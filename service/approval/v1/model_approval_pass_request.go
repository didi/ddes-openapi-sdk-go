package v1

// ApprovalPassRequest struct for ApprovalPassRequest
type ApprovalPassRequest struct {
	ClientId           *string      `json:"client_id,omitempty"`            // 申请应用时分配的AppKey
	AccessToken        *string      `json:"access_token,omitempty"`         // 授权后的access token
	CompanyId          *string      `json:"company_id,omitempty"`           // 企业ID
	Timestamp          *int64       `json:"timestamp,omitempty"`            // 当前时间戳(精确到秒级)
	Sign               *string      `json:"sign,omitempty"`                 // 签名
	ObjectType         *int32       `json:"object_type,omitempty"`          // 授权对象，授权对象枚举：1：订单 2：审批单
	ObjectId           *int64       `json:"object_id,omitempty"`            // 授权id，object_type对应对象的值 滴滴订单号或滴滴审批单号
	ObjectApprovalType *int32       `json:"object_approval_type,omitempty"` // 审批类型，审批类型枚举：4：个转企类型 21：企业级事前审批单用车
	IsPass             *int32       `json:"is_pass,omitempty"`              // 审批类型，审批类型枚举：1：通过 2 拒绝
	CurApprover        *string      `json:"cur_approver,omitempty"`         // 当前审批人，举例：\"{\\\"type\\\":\\\"phone\\\",\\\"value\\\":\\\"11100077098\\\"}\"
	CurApproverObj     *CurAppRover `json:"cur_approver__obj__,omitempty"`
}

type ApprovalPassRequestBuilder struct {
	clientId              string // 申请应用时分配的AppKey
	clientIdSet           bool
	accessToken           string // 授权后的access token
	accessTokenSet        bool
	companyId             string // 企业ID
	companyIdSet          bool
	timestamp             int64 // 当前时间戳(精确到秒级)
	timestampSet          bool
	sign                  string // 签名
	signSet               bool
	objectType            int32 // 授权对象，授权对象枚举：1：订单 2：审批单
	objectTypeSet         bool
	objectId              int64 // 授权id，object_type对应对象的值 滴滴订单号或滴滴审批单号
	objectIdSet           bool
	objectApprovalType    int32 // 审批类型，审批类型枚举：4：个转企类型 21：企业级事前审批单用车
	objectApprovalTypeSet bool
	isPass                int32 // 审批类型，审批类型枚举：1：通过 2 拒绝
	isPassSet             bool
	curApprover           string // 当前审批人，举例：\"{\\\"type\\\":\\\"phone\\\",\\\"value\\\":\\\"11100077098\\\"}\"
	curApproverSet        bool
	curApproverObj        CurAppRover
	curApproverObjSet     bool
}

func NewApprovalPassRequestBuilder() *ApprovalPassRequestBuilder {
	return &ApprovalPassRequestBuilder{}
}
func (builder *ApprovalPassRequestBuilder) ClientId(clientId string) *ApprovalPassRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) AccessToken(accessToken string) *ApprovalPassRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) CompanyId(companyId string) *ApprovalPassRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) Timestamp(timestamp int64) *ApprovalPassRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) Sign(sign string) *ApprovalPassRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) ObjectType(objectType int32) *ApprovalPassRequestBuilder {
	builder.objectType = objectType
	builder.objectTypeSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) ObjectId(objectId int64) *ApprovalPassRequestBuilder {
	builder.objectId = objectId
	builder.objectIdSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) ObjectApprovalType(objectApprovalType int32) *ApprovalPassRequestBuilder {
	builder.objectApprovalType = objectApprovalType
	builder.objectApprovalTypeSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) IsPass(isPass int32) *ApprovalPassRequestBuilder {
	builder.isPass = isPass
	builder.isPassSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) CurApprover(curApprover string) *ApprovalPassRequestBuilder {
	builder.curApprover = curApprover
	builder.curApproverSet = true
	return builder
}
func (builder *ApprovalPassRequestBuilder) CurApproverObj(curApproverObj CurAppRover) *ApprovalPassRequestBuilder {
	builder.curApproverObj = curApproverObj
	builder.curApproverObjSet = true
	return builder
}

func (builder *ApprovalPassRequestBuilder) Build() *ApprovalPassRequest {
	data := &ApprovalPassRequest{}
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
	if builder.objectTypeSet {
		data.ObjectType = &builder.objectType
	}
	if builder.objectIdSet {
		data.ObjectId = &builder.objectId
	}
	if builder.objectApprovalTypeSet {
		data.ObjectApprovalType = &builder.objectApprovalType
	}
	if builder.isPassSet {
		data.IsPass = &builder.isPass
	}
	if builder.curApproverSet {
		data.CurApprover = &builder.curApprover
	}
	if builder.curApproverObjSet {
		data.CurApproverObj = &builder.curApproverObj
	}
	return data
}
