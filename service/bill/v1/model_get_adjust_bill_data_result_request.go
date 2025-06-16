package v1

// GetAdjustBillDataResultRequest struct for GetAdjustBillDataResultRequest
type GetAdjustBillDataResultRequest struct {
	ClientId    *string `json:"client_id,omitempty"`     // 申请应用时分配的AppKey
	AccessToken *string `json:"access_token,omitempty"`  // 授权后的access token
	Timestamp   *int64  `json:"timestamp,omitempty"`     // 当前时间戳，精确到秒级
	CompanyId   *string `json:"company_id,omitempty"`    // 企业ID
	AdjustReqId *string `json:"adjust_req_id,omitempty"` // 幂等ID 调整提交id
	BillId      *int64  `json:"bill_id,omitempty"`       // 账单ID
}

type GetAdjustBillDataResultRequestBuilder struct {
	clientId       string // 申请应用时分配的AppKey
	clientIdSet    bool
	accessToken    string // 授权后的access token
	accessTokenSet bool
	timestamp      int64 // 当前时间戳，精确到秒级
	timestampSet   bool
	companyId      string // 企业ID
	companyIdSet   bool
	adjustReqId    string // 幂等ID 调整提交id
	adjustReqIdSet bool
	billId         int64 // 账单ID
	billIdSet      bool
}

func NewGetAdjustBillDataResultRequestBuilder() *GetAdjustBillDataResultRequestBuilder {
	return &GetAdjustBillDataResultRequestBuilder{}
}
func (builder *GetAdjustBillDataResultRequestBuilder) ClientId(clientId string) *GetAdjustBillDataResultRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *GetAdjustBillDataResultRequestBuilder) AccessToken(accessToken string) *GetAdjustBillDataResultRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *GetAdjustBillDataResultRequestBuilder) Timestamp(timestamp int64) *GetAdjustBillDataResultRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *GetAdjustBillDataResultRequestBuilder) CompanyId(companyId string) *GetAdjustBillDataResultRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *GetAdjustBillDataResultRequestBuilder) AdjustReqId(adjustReqId string) *GetAdjustBillDataResultRequestBuilder {
	builder.adjustReqId = adjustReqId
	builder.adjustReqIdSet = true
	return builder
}
func (builder *GetAdjustBillDataResultRequestBuilder) BillId(billId int64) *GetAdjustBillDataResultRequestBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}

func (builder *GetAdjustBillDataResultRequestBuilder) Build() *GetAdjustBillDataResultRequest {
	data := &GetAdjustBillDataResultRequest{}
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
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	return data
}
