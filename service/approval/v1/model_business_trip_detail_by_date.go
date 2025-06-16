package v1

// BusinessTripDetailByDate 因公用车审批行程数据(按日期) -> business_trip_detail
type BusinessTripDetailByDate struct {
	StartTime          *string        `json:"start_time,omitempty"`           // 开始时间，需要大于等于接口调用当前时间。时间格式为：2015-06-16 12:00:09，必填
	EndTime            *string        `json:"end_time,omitempty"`             // 结束时间，需大于等于开始时间。时间格式为：2015-06-16 12:00:09，必填
	Trips              []BusinessCity `json:"trips,omitempty"`                // 用车城市列表，创建规则时如果选择了“由员工填写”则必填，制度配置为无需填写城市时，传了城市列表也不生效，示例：[{ \"city_id\": 1, \"city\": \"北京\" },{ \"city_id\": 2, \"city\": \"上海\" }]
	TripAmount         *int32         `json:"trip_amount,omitempty"`          // 用车金额，单位：【分】（100表示100分，也就是1块钱。大于100生效，小于100为不限，选填
	TripTimes          *int32         `json:"trip_times,omitempty"`           // 用车次数，非必传只在制度配置为【由员工填写时】生效，不传、传空报错，建议传大于0的整数，若想不限次数可以传负数比如 -1，选填
	PerorderMoneyQuota *int32         `json:"perorder_money_quota,omitempty"` // 每单限额。单位：【分】（100表示100分，也就是1块钱）。1、es 后台制度配置为【由员工填写】时：该参数必须要传，且参数值要大于0。参数值大于100才生效，参数值小于100时不限制用车金额 （建议研发人员针对员工填写的参数做一下处理，比如说员工填写100元，该参数传10000）；2. es 后台制度配置为【不限】或者【限制 每单企业支付不超过xx元】时，即使接口传了该参数，该参数也不会生效，以后台配置为主，选填
}

type BusinessTripDetailByDateBuilder struct {
	startTime             string // 开始时间，需要大于等于接口调用当前时间。时间格式为：2015-06-16 12:00:09，必填
	startTimeSet          bool
	endTime               string // 结束时间，需大于等于开始时间。时间格式为：2015-06-16 12:00:09，必填
	endTimeSet            bool
	trips                 []BusinessCity // 用车城市列表，创建规则时如果选择了“由员工填写”则必填，制度配置为无需填写城市时，传了城市列表也不生效，示例：[{ \"city_id\": 1, \"city\": \"北京\" },{ \"city_id\": 2, \"city\": \"上海\" }]
	tripsSet              bool
	tripAmount            int32 // 用车金额，单位：【分】（100表示100分，也就是1块钱。大于100生效，小于100为不限，选填
	tripAmountSet         bool
	tripTimes             int32 // 用车次数，非必传只在制度配置为【由员工填写时】生效，不传、传空报错，建议传大于0的整数，若想不限次数可以传负数比如 -1，选填
	tripTimesSet          bool
	perorderMoneyQuota    int32 // 每单限额。单位：【分】（100表示100分，也就是1块钱）。1、es 后台制度配置为【由员工填写】时：该参数必须要传，且参数值要大于0。参数值大于100才生效，参数值小于100时不限制用车金额 （建议研发人员针对员工填写的参数做一下处理，比如说员工填写100元，该参数传10000）；2. es 后台制度配置为【不限】或者【限制 每单企业支付不超过xx元】时，即使接口传了该参数，该参数也不会生效，以后台配置为主，选填
	perorderMoneyQuotaSet bool
}

func NewBusinessTripDetailByDateBuilder() *BusinessTripDetailByDateBuilder {
	return &BusinessTripDetailByDateBuilder{}
}
func (builder *BusinessTripDetailByDateBuilder) StartTime(startTime string) *BusinessTripDetailByDateBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *BusinessTripDetailByDateBuilder) EndTime(endTime string) *BusinessTripDetailByDateBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *BusinessTripDetailByDateBuilder) Trips(trips []BusinessCity) *BusinessTripDetailByDateBuilder {
	builder.trips = trips
	builder.tripsSet = true
	return builder
}
func (builder *BusinessTripDetailByDateBuilder) TripAmount(tripAmount int32) *BusinessTripDetailByDateBuilder {
	builder.tripAmount = tripAmount
	builder.tripAmountSet = true
	return builder
}
func (builder *BusinessTripDetailByDateBuilder) TripTimes(tripTimes int32) *BusinessTripDetailByDateBuilder {
	builder.tripTimes = tripTimes
	builder.tripTimesSet = true
	return builder
}
func (builder *BusinessTripDetailByDateBuilder) PerorderMoneyQuota(perorderMoneyQuota int32) *BusinessTripDetailByDateBuilder {
	builder.perorderMoneyQuota = perorderMoneyQuota
	builder.perorderMoneyQuotaSet = true
	return builder
}

func (builder *BusinessTripDetailByDateBuilder) Build() *BusinessTripDetailByDate {
	data := &BusinessTripDetailByDate{}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.tripsSet {
		data.Trips = builder.trips
	}
	if builder.tripAmountSet {
		data.TripAmount = &builder.tripAmount
	}
	if builder.tripTimesSet {
		data.TripTimes = &builder.tripTimes
	}
	if builder.perorderMoneyQuotaSet {
		data.PerorderMoneyQuota = &builder.perorderMoneyQuota
	}
	return data
}
