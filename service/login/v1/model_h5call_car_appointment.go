package v1

// H5CallCarAppointment 预约用车
type H5CallCarAppointment struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 预约用车为callCarAppointment
	ApprovalId            *string `json:"approval_id,omitempty"`              // 差旅、申请用车，必传
	AppointmentTime       *string `json:"appointment_time,omitempty"`         // 预约时间，必传;如需制定预约时间，必传，格式”今天|14:30”，只支持[今天、明天、后天]
	LatFrom               *string `json:"lat_from,omitempty"`                 // 如不传以定位获取的地址为起点
	LngFrom               *string `json:"lng_from,omitempty"`                 // 如不传以定位获取的地址为起点
	PoiFromName           *string `json:"poi_from_name,omitempty"`            // 如不传以定位获取的地址为起点
	ToCityId              *string `json:"to_city_id,omitempty"`               // 如不传以手动选择地址为终点
	LatTo                 *string `json:"lat_to,omitempty"`                   // 如不传以手动选择地址为终点
	LngTo                 *string `json:"lng_to,omitempty"`                   // 如不传以手动选择地址为终点
	PoiToName             *string `json:"poi_to_name,omitempty"`              // 如不传以手动选择地址为终点
	CityId                *string `json:"city_id,omitempty"`                  // 如不传以手动选择地址为终点
	PassengerPhone        *string `json:"passenger_phone,omitempty"`          // 乘车人手机号，代叫必传
	PassengerName         *string `json:"passenger_name,omitempty"`           // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	RestrictPoiFlag       *string `json:"restrict_poi_flag,omitempty"`        // 如需锁定地址，必传
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5CallCarAppointmentBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 预约用车为callCarAppointment
	jumpPageSet              bool
	approvalId               string // 差旅、申请用车，必传
	approvalIdSet            bool
	appointmentTime          string // 预约时间，必传;如需制定预约时间，必传，格式”今天|14:30”，只支持[今天、明天、后天]
	appointmentTimeSet       bool
	latFrom                  string // 如不传以定位获取的地址为起点
	latFromSet               bool
	lngFrom                  string // 如不传以定位获取的地址为起点
	lngFromSet               bool
	poiFromName              string // 如不传以定位获取的地址为起点
	poiFromNameSet           bool
	toCityId                 string // 如不传以手动选择地址为终点
	toCityIdSet              bool
	latTo                    string // 如不传以手动选择地址为终点
	latToSet                 bool
	lngTo                    string // 如不传以手动选择地址为终点
	lngToSet                 bool
	poiToName                string // 如不传以手动选择地址为终点
	poiToNameSet             bool
	cityId                   string // 如不传以手动选择地址为终点
	cityIdSet                bool
	passengerPhone           string // 乘车人手机号，代叫必传
	passengerPhoneSet        bool
	passengerName            string // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	passengerNameSet         bool
	restrictPoiFlag          string // 如需锁定地址，必传
	restrictPoiFlagSet       bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5CallCarAppointmentBuilder() *H5CallCarAppointmentBuilder {
	return &H5CallCarAppointmentBuilder{}
}
func (builder *H5CallCarAppointmentBuilder) RuleId(ruleId string) *H5CallCarAppointmentBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) JumpPage(jumpPage string) *H5CallCarAppointmentBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) ApprovalId(approvalId string) *H5CallCarAppointmentBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) AppointmentTime(appointmentTime string) *H5CallCarAppointmentBuilder {
	builder.appointmentTime = appointmentTime
	builder.appointmentTimeSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) LatFrom(latFrom string) *H5CallCarAppointmentBuilder {
	builder.latFrom = latFrom
	builder.latFromSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) LngFrom(lngFrom string) *H5CallCarAppointmentBuilder {
	builder.lngFrom = lngFrom
	builder.lngFromSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) PoiFromName(poiFromName string) *H5CallCarAppointmentBuilder {
	builder.poiFromName = poiFromName
	builder.poiFromNameSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) ToCityId(toCityId string) *H5CallCarAppointmentBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) LatTo(latTo string) *H5CallCarAppointmentBuilder {
	builder.latTo = latTo
	builder.latToSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) LngTo(lngTo string) *H5CallCarAppointmentBuilder {
	builder.lngTo = lngTo
	builder.lngToSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) PoiToName(poiToName string) *H5CallCarAppointmentBuilder {
	builder.poiToName = poiToName
	builder.poiToNameSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) CityId(cityId string) *H5CallCarAppointmentBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) PassengerPhone(passengerPhone string) *H5CallCarAppointmentBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) PassengerName(passengerName string) *H5CallCarAppointmentBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5CallCarAppointmentBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5CallCarAppointmentBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5CallCarAppointmentBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5CallCarAppointmentBuilder) Build() *H5CallCarAppointment {
	data := &H5CallCarAppointment{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.appointmentTimeSet {
		data.AppointmentTime = &builder.appointmentTime
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
	if builder.cityIdSet {
		data.CityId = &builder.cityId
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
