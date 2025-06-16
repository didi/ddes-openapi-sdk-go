package v1

// ListTransferTrainTicketRequest struct for ListTransferTrainTicketRequest
type ListTransferTrainTicketRequest struct {
	ClientId        *string `json:"client_id,omitempty"`         // 申请应用时分配的AppKey
	AccessToken     *string `json:"access_token,omitempty"`      // 授权后的access token
	CompanyId       *string `json:"company_id,omitempty"`        // 企业ID
	Timestamp       *int32  `json:"timestamp,omitempty"`         // 当前时间戳(精确到秒级)
	Sign            *string `json:"sign,omitempty"`              // 签名
	TrainDate       *string `json:"train_date,omitempty"`        // 上车日期，格式2021-01-01
	FromStationName *string `json:"from_station_name,omitempty"` // 上车站站名
	ToStationName   *string `json:"to_station_name,omitempty"`   // 下车站站名
	CurPage         *int32  `json:"cur_page,omitempty"`          // 当前页码，大于等于1且为正整数
}

type ListTransferTrainTicketRequestBuilder struct {
	clientId           string // 申请应用时分配的AppKey
	clientIdSet        bool
	accessToken        string // 授权后的access token
	accessTokenSet     bool
	companyId          string // 企业ID
	companyIdSet       bool
	timestamp          int32 // 当前时间戳(精确到秒级)
	timestampSet       bool
	sign               string // 签名
	signSet            bool
	trainDate          string // 上车日期，格式2021-01-01
	trainDateSet       bool
	fromStationName    string // 上车站站名
	fromStationNameSet bool
	toStationName      string // 下车站站名
	toStationNameSet   bool
	curPage            int32 // 当前页码，大于等于1且为正整数
	curPageSet         bool
}

func NewListTransferTrainTicketRequestBuilder() *ListTransferTrainTicketRequestBuilder {
	return &ListTransferTrainTicketRequestBuilder{}
}
func (builder *ListTransferTrainTicketRequestBuilder) ClientId(clientId string) *ListTransferTrainTicketRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) AccessToken(accessToken string) *ListTransferTrainTicketRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) CompanyId(companyId string) *ListTransferTrainTicketRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) Timestamp(timestamp int32) *ListTransferTrainTicketRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) Sign(sign string) *ListTransferTrainTicketRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) TrainDate(trainDate string) *ListTransferTrainTicketRequestBuilder {
	builder.trainDate = trainDate
	builder.trainDateSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) FromStationName(fromStationName string) *ListTransferTrainTicketRequestBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) ToStationName(toStationName string) *ListTransferTrainTicketRequestBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *ListTransferTrainTicketRequestBuilder) CurPage(curPage int32) *ListTransferTrainTicketRequestBuilder {
	builder.curPage = curPage
	builder.curPageSet = true
	return builder
}

func (builder *ListTransferTrainTicketRequestBuilder) Build() *ListTransferTrainTicketRequest {
	data := &ListTransferTrainTicketRequest{}
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
	if builder.trainDateSet {
		data.TrainDate = &builder.trainDate
	}
	if builder.fromStationNameSet {
		data.FromStationName = &builder.fromStationName
	}
	if builder.toStationNameSet {
		data.ToStationName = &builder.toStationName
	}
	if builder.curPageSet {
		data.CurPage = &builder.curPage
	}
	return data
}
