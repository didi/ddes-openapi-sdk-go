package v1

// UpdateAdjustBillDataRequest struct for UpdateAdjustBillDataRequest
type UpdateAdjustBillDataRequest struct {
	ClientId     *string          `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	AccessToken  *string          `json:"access_token,omitempty"`  // 授权后的access token
	Timestamp    *int64           `json:"timestamp,omitempty"`     // 当前时间戳，精确到秒级
	CompanyId    *string          `json:"company_id,omitempty"`    // 企业ID
	AdjustReqId  *string          `json:"adjust_req_id,omitempty"` // 自行传递的调整id,后面用于查询(成功的提交id不能重复)
	BusinessType *int32           `json:"business_type,omitempty"` // 业务线（1：网约车；2：商旅, 40：代驾 100：出租车）
	AdjustType   *int32           `json:"adjust_type,omitempty"`   // 调整类型1:订单计入下期2:订单信息调整
	BillId       *int64           `json:"bill_id,omitempty"`       // 账单ID
	AdjustList   []AdjustListItem `json:"adjust_list,omitempty"`   // adjust_list，要求长度大于0
	Remark       *string          `json:"remark,omitempty"`        // 备注
}

type UpdateAdjustBillDataRequestBuilder struct {
	clientId        string // 申请应用时分配的AppKey
	clientIdSet     bool
	accessToken     string // 授权后的access token
	accessTokenSet  bool
	timestamp       int64 // 当前时间戳，精确到秒级
	timestampSet    bool
	companyId       string // 企业ID
	companyIdSet    bool
	adjustReqId     string // 自行传递的调整id,后面用于查询(成功的提交id不能重复)
	adjustReqIdSet  bool
	businessType    int32 // 业务线（1：网约车；2：商旅, 40：代驾 100：出租车）
	businessTypeSet bool
	adjustType      int32 // 调整类型1:订单计入下期2:订单信息调整
	adjustTypeSet   bool
	billId          int64 // 账单ID
	billIdSet       bool
	adjustList      []AdjustListItem // adjust_list，要求长度大于0
	adjustListSet   bool
	remark          string // 备注
	remarkSet       bool
}

func NewUpdateAdjustBillDataRequestBuilder() *UpdateAdjustBillDataRequestBuilder {
	return &UpdateAdjustBillDataRequestBuilder{}
}
func (builder *UpdateAdjustBillDataRequestBuilder) ClientId(clientId string) *UpdateAdjustBillDataRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) AccessToken(accessToken string) *UpdateAdjustBillDataRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) Timestamp(timestamp int64) *UpdateAdjustBillDataRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) CompanyId(companyId string) *UpdateAdjustBillDataRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) AdjustReqId(adjustReqId string) *UpdateAdjustBillDataRequestBuilder {
	builder.adjustReqId = adjustReqId
	builder.adjustReqIdSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) BusinessType(businessType int32) *UpdateAdjustBillDataRequestBuilder {
	builder.businessType = businessType
	builder.businessTypeSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) AdjustType(adjustType int32) *UpdateAdjustBillDataRequestBuilder {
	builder.adjustType = adjustType
	builder.adjustTypeSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) BillId(billId int64) *UpdateAdjustBillDataRequestBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) AdjustList(adjustList []AdjustListItem) *UpdateAdjustBillDataRequestBuilder {
	builder.adjustList = adjustList
	builder.adjustListSet = true
	return builder
}
func (builder *UpdateAdjustBillDataRequestBuilder) Remark(remark string) *UpdateAdjustBillDataRequestBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}

func (builder *UpdateAdjustBillDataRequestBuilder) Build() *UpdateAdjustBillDataRequest {
	data := &UpdateAdjustBillDataRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.adjustReqIdSet {
		data.AdjustReqId = &builder.adjustReqId
	}
	if builder.businessTypeSet {
		data.BusinessType = &builder.businessType
	}
	if builder.adjustTypeSet {
		data.AdjustType = &builder.adjustType
	}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.adjustListSet {
		data.AdjustList = builder.adjustList
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	return data
}
