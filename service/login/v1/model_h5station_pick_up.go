package v1

// H5StationPickUp 接站
type H5StationPickUp struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 接站为stationPickUp
	AppointmentTime       *string `json:"appointment_time,omitempty"`         // 预约时间，必传;如需制定预约时间，必传，格式”今天|14:30”，只支持[今天、明天、后天]
	ApprovalId            *string `json:"approval_id,omitempty"`              // 差旅、申请用车，必传
	LatFrom               *string `json:"lat_from,omitempty"`                 // 如不传以定位获取的地址为起点
	LngFrom               *string `json:"lng_from,omitempty"`                 // 如不传以定位获取的地址为起点
	PoiFromName           *string `json:"poi_from_name,omitempty"`            // 如不传以定位获取的地址为起点
	CityId                *string `json:"city_id,omitempty"`                  // 如不传以手动选择地址为起点
	ToCityId              *string `json:"to_city_id,omitempty"`               // 如不传以手动选择地址为终点
	LatTo                 *string `json:"lat_to,omitempty"`                   // 如不传以手动选择地址为终点
	LngTo                 *string `json:"lng_to,omitempty"`                   // 如不传以手动选择地址为终点
	PoiToName             *string `json:"poi_to_name,omitempty"`              // 如不传以手动选择地址为终点
	PassengerPhone        *string `json:"passenger_phone,omitempty"`          // 乘车人手机号，代叫必传
	PassengerName         *string `json:"passenger_name,omitempty"`           // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	RestrictPoiFlag       *string `json:"restrict_poi_flag,omitempty"`        // 如需锁定地址，必传
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5StationPickUpBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 接站为stationPickUp
	jumpPageSet              bool
	appointmentTime          string // 预约时间，必传;如需制定预约时间，必传，格式”今天|14:30”，只支持[今天、明天、后天]
	appointmentTimeSet       bool
	approvalId               string // 差旅、申请用车，必传
	approvalIdSet            bool
	latFrom                  string // 如不传以定位获取的地址为起点
	latFromSet               bool
	lngFrom                  string // 如不传以定位获取的地址为起点
	lngFromSet               bool
	poiFromName              string // 如不传以定位获取的地址为起点
	poiFromNameSet           bool
	cityId                   string // 如不传以手动选择地址为起点
	cityIdSet                bool
	toCityId                 string // 如不传以手动选择地址为终点
	toCityIdSet              bool
	latTo                    string // 如不传以手动选择地址为终点
	latToSet                 bool
	lngTo                    string // 如不传以手动选择地址为终点
	lngToSet                 bool
	poiToName                string // 如不传以手动选择地址为终点
	poiToNameSet             bool
	passengerPhone           string // 乘车人手机号，代叫必传
	passengerPhoneSet        bool
	passengerName            string // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	passengerNameSet         bool
	restrictPoiFlag          string // 如需锁定地址，必传
	restrictPoiFlagSet       bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5StationPickUpBuilder() *H5StationPickUpBuilder {
	return &H5StationPickUpBuilder{}
}
func (builder *H5StationPickUpBuilder) RuleId(ruleId string) *H5StationPickUpBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) JumpPage(jumpPage string) *H5StationPickUpBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) AppointmentTime(appointmentTime string) *H5StationPickUpBuilder {
	builder.appointmentTime = appointmentTime
	builder.appointmentTimeSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) ApprovalId(approvalId string) *H5StationPickUpBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) LatFrom(latFrom string) *H5StationPickUpBuilder {
	builder.latFrom = latFrom
	builder.latFromSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) LngFrom(lngFrom string) *H5StationPickUpBuilder {
	builder.lngFrom = lngFrom
	builder.lngFromSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) PoiFromName(poiFromName string) *H5StationPickUpBuilder {
	builder.poiFromName = poiFromName
	builder.poiFromNameSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) CityId(cityId string) *H5StationPickUpBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) ToCityId(toCityId string) *H5StationPickUpBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) LatTo(latTo string) *H5StationPickUpBuilder {
	builder.latTo = latTo
	builder.latToSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) LngTo(lngTo string) *H5StationPickUpBuilder {
	builder.lngTo = lngTo
	builder.lngToSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) PoiToName(poiToName string) *H5StationPickUpBuilder {
	builder.poiToName = poiToName
	builder.poiToNameSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) PassengerPhone(passengerPhone string) *H5StationPickUpBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) PassengerName(passengerName string) *H5StationPickUpBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5StationPickUpBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5StationPickUpBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5StationPickUpBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5StationPickUpBuilder) Build() *H5StationPickUp {
	data := &H5StationPickUp{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.appointmentTimeSet {
		data.AppointmentTime = &builder.appointmentTime
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.latFromSet {
		data.LatFrom = &builder.latFrom
	}
	if builder.lngFromSet {
		data.LngFrom = &builder.lngFrom
	}
	if builder.poiFromNameSet {
		data.PoiFromName = &builder.poiFromName
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.toCityIdSet {
		data.ToCityId = &builder.toCityId
	}
	if builder.latToSet {
		data.LatTo = &builder.latTo
	}
	if builder.lngToSet {
		data.LngTo = &builder.lngTo
	}
	if builder.poiToNameSet {
		data.PoiToName = &builder.poiToName
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.restrictPoiFlagSet {
		data.RestrictPoiFlag = &builder.restrictPoiFlag
	}
	if builder.hideEstimatePriceFlagSet {
		data.HideEstimatePriceFlag = &builder.hideEstimatePriceFlag
	}
	return data
}
