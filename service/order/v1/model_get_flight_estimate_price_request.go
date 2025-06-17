package v1

// GetFlightEstimatePriceRequest struct for GetFlightEstimatePriceRequest
type GetFlightEstimatePriceRequest struct {
	ClientId        *string `json:"client_id,omitempty"`         // 申请应用时分配的AppKey
	AccessToken     *string `json:"access_token,omitempty"`      // 授权后的access token
	CompanyId       *string `json:"company_id,omitempty"`        // 企业ID
	Timestamp       *int64  `json:"timestamp,omitempty"`         // 当前时间戳(精确到秒级)
	Sign            *string `json:"sign,omitempty"`              // 签名
	DepartureCityId *string `json:"departure_city_id,omitempty"` // 出发城市id
	ArrivalCityId   *string `json:"arrival_city_id,omitempty"`   // 到达城市id
	Date            *string `json:"date,omitempty"`              // 出发时间，格式 2024-01-31
	SearchType      *int32  `json:"search_type,omitempty"`       // 1表示只要直飞，2表示只要中转，3或0或者不传表示全都搜索
}

type GetFlightEstimatePriceRequestBuilder struct {
	clientId           string // 申请应用时分配的AppKey
	clientIdSet        bool
	accessToken        string // 授权后的access token
	accessTokenSet     bool
	companyId          string // 企业ID
	companyIdSet       bool
	timestamp          int64 // 当前时间戳(精确到秒级)
	timestampSet       bool
	sign               string // 签名
	signSet            bool
	departureCityId    string // 出发城市id
	departureCityIdSet bool
	arrivalCityId      string // 到达城市id
	arrivalCityIdSet   bool
	date               string // 出发时间，格式 2024-01-31
	dateSet            bool
	searchType         int32 // 1表示只要直飞，2表示只要中转，3或0或者不传表示全都搜索
	searchTypeSet      bool
}

func NewGetFlightEstimatePriceRequestBuilder() *GetFlightEstimatePriceRequestBuilder {
	return &GetFlightEstimatePriceRequestBuilder{}
}
func (builder *GetFlightEstimatePriceRequestBuilder) ClientId(clientId string) *GetFlightEstimatePriceRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) AccessToken(accessToken string) *GetFlightEstimatePriceRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) CompanyId(companyId string) *GetFlightEstimatePriceRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) Timestamp(timestamp int64) *GetFlightEstimatePriceRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) Sign(sign string) *GetFlightEstimatePriceRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) DepartureCityId(departureCityId string) *GetFlightEstimatePriceRequestBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) ArrivalCityId(arrivalCityId string) *GetFlightEstimatePriceRequestBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) Date(date string) *GetFlightEstimatePriceRequestBuilder {
	builder.date = date
	builder.dateSet = true
	return builder
}
func (builder *GetFlightEstimatePriceRequestBuilder) SearchType(searchType int32) *GetFlightEstimatePriceRequestBuilder {
	builder.searchType = searchType
	builder.searchTypeSet = true
	return builder
}

func (builder *GetFlightEstimatePriceRequestBuilder) Build() *GetFlightEstimatePriceRequest {
	data := &GetFlightEstimatePriceRequest{}
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
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	if builder.dateSet {
		data.Date = &builder.date
	}
	if builder.searchTypeSet {
		data.SearchType = &builder.searchType
	}
	return data
}
