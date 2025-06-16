package v1

// ListTrainLeftTicketRequest struct for ListTrainLeftTicketRequest
type ListTrainLeftTicketRequest struct {
	ClientId        *string `json:"client_id,omitempty"`         // 申请应用时分配的AppKey
	AccessToken     *string `json:"access_token,omitempty"`      // 授权后的access token
	CompanyId       *string `json:"company_id,omitempty"`        // 企业ID
	Timestamp       *int32  `json:"timestamp,omitempty"`         // 当前时间戳(精确到秒级)
	Sign            *string `json:"sign,omitempty"`              // 签名
	TrainDate       *string `json:"train_date,omitempty"`        // 上车日期，格式2021-01-01
	FromStationName *string `json:"from_station_name,omitempty"` // 上车站站名
	ToStationName   *string `json:"to_station_name,omitempty"`   // 下车站站名
	StartTime       *string `json:"start_time,omitempty"`        // 查询开始时间(列车出发时间)，格式12:00:00，不传为00:00:00
	EndTime         *string `json:"end_time,omitempty"`          // 查询结束时间(列车出发时间)，格式12:00:00，不传为23:59:59，包含最后一秒
}

type ListTrainLeftTicketRequestBuilder struct {
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
	startTime          string // 查询开始时间(列车出发时间)，格式12:00:00，不传为00:00:00
	startTimeSet       bool
	endTime            string // 查询结束时间(列车出发时间)，格式12:00:00，不传为23:59:59，包含最后一秒
	endTimeSet         bool
}

func NewListTrainLeftTicketRequestBuilder() *ListTrainLeftTicketRequestBuilder {
	return &ListTrainLeftTicketRequestBuilder{}
}
func (builder *ListTrainLeftTicketRequestBuilder) ClientId(clientId string) *ListTrainLeftTicketRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) AccessToken(accessToken string) *ListTrainLeftTicketRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) CompanyId(companyId string) *ListTrainLeftTicketRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) Timestamp(timestamp int32) *ListTrainLeftTicketRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) Sign(sign string) *ListTrainLeftTicketRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) TrainDate(trainDate string) *ListTrainLeftTicketRequestBuilder {
	builder.trainDate = trainDate
	builder.trainDateSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) FromStationName(fromStationName string) *ListTrainLeftTicketRequestBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) ToStationName(toStationName string) *ListTrainLeftTicketRequestBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) StartTime(startTime string) *ListTrainLeftTicketRequestBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *ListTrainLeftTicketRequestBuilder) EndTime(endTime string) *ListTrainLeftTicketRequestBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}

func (builder *ListTrainLeftTicketRequestBuilder) Build() *ListTrainLeftTicketRequest {
	data := &ListTrainLeftTicketRequest{}
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
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	return data
}
