package v1

// BusinessTripDetailByTimes 因公用车审批行程数据(按次数) -> business_trip_detail
type BusinessTripDetailByTimes struct {
	StartTime          *string  `json:"start_time,omitempty"`           // 开始时间，开始时间，需要大于等于接口调用当前时间。时间格式为：2015-06-16 12:00:09
	EndTime            *string  `json:"end_time,omitempty"`             // 结束时间，结束时间，需大于等于开始时间。时间格式为：2015-06-16 12:00:09
	DepartureCityId    *int32   `json:"departure_city_id,omitempty"`    // 滴滴出发城市 ID，出发城市 ID；制度配置需要填写时必填
	DepartureCity      *string  `json:"departure_city,omitempty"`       // 出发地城市名称，出发地城市名称（指定出发地时必填）
	DestinationCityId  *int32   `json:"destination_city_id,omitempty"`  // 滴滴目的地城市ID，目的地城市id（指定目的地时必填）
	DestinationCity    *string  `json:"destination_city,omitempty"`     // 目的地城市名称，目的地城市名称（指定目的地时必填）
	StartName          *string  `json:"start_name,omitempty"`           // 出发地名称，出发地名称（最多50个字）（指定出发地时必填）
	StartAddress       *string  `json:"start_address,omitempty"`        // 出发地详细地址，出发地详细地址（最多100个字）（指定出发地时必填）
	Flat               *float32 `json:"flat,omitempty"`                 // 出发地纬度，出发地纬度 （指定出发地时必填）
	Flng               *float32 `json:"flng,omitempty"`                 // 出发地经度，出发地经度 （指定出发地时必填）
	EndName            *string  `json:"end_name,omitempty"`             // 目的地名称，目的地名称（最多50个字）（指定目的地时必填）
	EndAddress         *string  `json:"end_address,omitempty"`          // 目的地详细地址，目的地详细地址（最多100个字）（指定目的地时必填）
	Tlat               *float32 `json:"tlat,omitempty"`                 // 目的地纬度，目的地纬度（指定目的地时必填）
	Tlng               *float32 `json:"tlng,omitempty"`                 // 目的地经度，目的地经度（指定目的地时必填）
	IsReturn           *int32   `json:"is_return,omitempty"`            // 是否往返，是否往返，0-不往返，1-往返，默认为0
	TripTimes          *int32   `json:"trip_times,omitempty"`           // 用车次数，用车次数，当is_return为1时，用车次数必须为偶数（去程一次+回程一次 制度配置为无需填写，传了不生效
	PerorderMoneyQuota *int32   `json:"perorder_money_quota,omitempty"` // 每单限额，每单限额，单位:分，需大于100生效, 需是 100 的整数倍 1. 制度配置了由员工填写，小于 100 为不限 2. 制度配置为无需填写，传了不生效 3. 制度配置金额，用户也传了。以配置的为主
	TotalMoneyQuota    *int32   `json:"total_money_quota,omitempty"`    // 总限额，总限额，单位：分，需大于100生效，需是 100 的整数倍 1. 制度配置了由员工填写，小于 100 为不限 2. 制度配置为无需填写，传了不生效
}

type BusinessTripDetailByTimesBuilder struct {
	startTime             string // 开始时间，开始时间，需要大于等于接口调用当前时间。时间格式为：2015-06-16 12:00:09
	startTimeSet          bool
	endTime               string // 结束时间，结束时间，需大于等于开始时间。时间格式为：2015-06-16 12:00:09
	endTimeSet            bool
	departureCityId       int32 // 滴滴出发城市 ID，出发城市 ID；制度配置需要填写时必填
	departureCityIdSet    bool
	departureCity         string // 出发地城市名称，出发地城市名称（指定出发地时必填）
	departureCitySet      bool
	destinationCityId     int32 // 滴滴目的地城市ID，目的地城市id（指定目的地时必填）
	destinationCityIdSet  bool
	destinationCity       string // 目的地城市名称，目的地城市名称（指定目的地时必填）
	destinationCitySet    bool
	startName             string // 出发地名称，出发地名称（最多50个字）（指定出发地时必填）
	startNameSet          bool
	startAddress          string // 出发地详细地址，出发地详细地址（最多100个字）（指定出发地时必填）
	startAddressSet       bool
	flat                  float32 // 出发地纬度，出发地纬度 （指定出发地时必填）
	flatSet               bool
	flng                  float32 // 出发地经度，出发地经度 （指定出发地时必填）
	flngSet               bool
	endName               string // 目的地名称，目的地名称（最多50个字）（指定目的地时必填）
	endNameSet            bool
	endAddress            string // 目的地详细地址，目的地详细地址（最多100个字）（指定目的地时必填）
	endAddressSet         bool
	tlat                  float32 // 目的地纬度，目的地纬度（指定目的地时必填）
	tlatSet               bool
	tlng                  float32 // 目的地经度，目的地经度（指定目的地时必填）
	tlngSet               bool
	isReturn              int32 // 是否往返，是否往返，0-不往返，1-往返，默认为0
	isReturnSet           bool
	tripTimes             int32 // 用车次数，用车次数，当is_return为1时，用车次数必须为偶数（去程一次+回程一次 制度配置为无需填写，传了不生效
	tripTimesSet          bool
	perorderMoneyQuota    int32 // 每单限额，每单限额，单位:分，需大于100生效, 需是 100 的整数倍 1. 制度配置了由员工填写，小于 100 为不限 2. 制度配置为无需填写，传了不生效 3. 制度配置金额，用户也传了。以配置的为主
	perorderMoneyQuotaSet bool
	totalMoneyQuota       int32 // 总限额，总限额，单位：分，需大于100生效，需是 100 的整数倍 1. 制度配置了由员工填写，小于 100 为不限 2. 制度配置为无需填写，传了不生效
	totalMoneyQuotaSet    bool
}

func NewBusinessTripDetailByTimesBuilder() *BusinessTripDetailByTimesBuilder {
	return &BusinessTripDetailByTimesBuilder{}
}
func (builder *BusinessTripDetailByTimesBuilder) StartTime(startTime string) *BusinessTripDetailByTimesBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) EndTime(endTime string) *BusinessTripDetailByTimesBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) DepartureCityId(departureCityId int32) *BusinessTripDetailByTimesBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) DepartureCity(departureCity string) *BusinessTripDetailByTimesBuilder {
	builder.departureCity = departureCity
	builder.departureCitySet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) DestinationCityId(destinationCityId int32) *BusinessTripDetailByTimesBuilder {
	builder.destinationCityId = destinationCityId
	builder.destinationCityIdSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) DestinationCity(destinationCity string) *BusinessTripDetailByTimesBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) StartName(startName string) *BusinessTripDetailByTimesBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) StartAddress(startAddress string) *BusinessTripDetailByTimesBuilder {
	builder.startAddress = startAddress
	builder.startAddressSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) Flat(flat float32) *BusinessTripDetailByTimesBuilder {
	builder.flat = flat
	builder.flatSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) Flng(flng float32) *BusinessTripDetailByTimesBuilder {
	builder.flng = flng
	builder.flngSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) EndName(endName string) *BusinessTripDetailByTimesBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) EndAddress(endAddress string) *BusinessTripDetailByTimesBuilder {
	builder.endAddress = endAddress
	builder.endAddressSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) Tlat(tlat float32) *BusinessTripDetailByTimesBuilder {
	builder.tlat = tlat
	builder.tlatSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) Tlng(tlng float32) *BusinessTripDetailByTimesBuilder {
	builder.tlng = tlng
	builder.tlngSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) IsReturn(isReturn int32) *BusinessTripDetailByTimesBuilder {
	builder.isReturn = isReturn
	builder.isReturnSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) TripTimes(tripTimes int32) *BusinessTripDetailByTimesBuilder {
	builder.tripTimes = tripTimes
	builder.tripTimesSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) PerorderMoneyQuota(perorderMoneyQuota int32) *BusinessTripDetailByTimesBuilder {
	builder.perorderMoneyQuota = perorderMoneyQuota
	builder.perorderMoneyQuotaSet = true
	return builder
}
func (builder *BusinessTripDetailByTimesBuilder) TotalMoneyQuota(totalMoneyQuota int32) *BusinessTripDetailByTimesBuilder {
	builder.totalMoneyQuota = totalMoneyQuota
	builder.totalMoneyQuotaSet = true
	return builder
}

func (builder *BusinessTripDetailByTimesBuilder) Build() *BusinessTripDetailByTimes {
	data := &BusinessTripDetailByTimes{}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.departureCitySet {
		data.DepartureCity = &builder.departureCity
	}
	if builder.destinationCityIdSet {
		data.DestinationCityId = &builder.destinationCityId
	}
	if builder.destinationCitySet {
		data.DestinationCity = &builder.destinationCity
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.startAddressSet {
		data.StartAddress = &builder.startAddress
	}
	if builder.flatSet {
		data.Flat = &builder.flat
	}
	if builder.flngSet {
		data.Flng = &builder.flng
	}
	if builder.endNameSet {
		data.EndName = &builder.endName
	}
	if builder.endAddressSet {
		data.EndAddress = &builder.endAddress
	}
	if builder.tlatSet {
		data.Tlat = &builder.tlat
	}
	if builder.tlngSet {
		data.Tlng = &builder.tlng
	}
	if builder.isReturnSet {
		data.IsReturn = &builder.isReturn
	}
	if builder.tripTimesSet {
		data.TripTimes = &builder.tripTimes
	}
	if builder.perorderMoneyQuotaSet {
		data.PerorderMoneyQuota = &builder.perorderMoneyQuota
	}
	if builder.totalMoneyQuotaSet {
		data.TotalMoneyQuota = &builder.totalMoneyQuota
	}
	return data
}
