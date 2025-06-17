package v1

// H5ToAirport 送机
type H5ToAirport struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 送机为toAirport
	AirportId             *string `json:"airport_id,omitempty"`               // 机场ID，如不传需手动选择
	CityId                *string `json:"city_id,omitempty"`                  // 如不传以手动选择地址为起点
	LatFrom               *string `json:"lat_from,omitempty"`                 // 如不传以手动选择地址为起点
	LngFrom               *string `json:"lng_from,omitempty"`                 // 如不传以手动选择地址为起点
	PoiFromName           *string `json:"poi_from_name,omitempty"`            // 如不传以手动选择地址为起点
	PassengerPhone        *string `json:"passenger_phone,omitempty"`          // 乘车人手机号，代叫必传
	PassengerName         *string `json:"passenger_name,omitempty"`           // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	RestrictPoiFlag       *string `json:"restrict_poi_flag,omitempty"`        // 如需锁定地址，必传
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5ToAirportBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 送机为toAirport
	jumpPageSet              bool
	airportId                string // 机场ID，如不传需手动选择
	airportIdSet             bool
	cityId                   string // 如不传以手动选择地址为起点
	cityIdSet                bool
	latFrom                  string // 如不传以手动选择地址为起点
	latFromSet               bool
	lngFrom                  string // 如不传以手动选择地址为起点
	lngFromSet               bool
	poiFromName              string // 如不传以手动选择地址为起点
	poiFromNameSet           bool
	passengerPhone           string // 乘车人手机号，代叫必传
	passengerPhoneSet        bool
	passengerName            string // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	passengerNameSet         bool
	restrictPoiFlag          string // 如需锁定地址，必传
	restrictPoiFlagSet       bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5ToAirportBuilder() *H5ToAirportBuilder {
	return &H5ToAirportBuilder{}
}
func (builder *H5ToAirportBuilder) RuleId(ruleId string) *H5ToAirportBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5ToAirportBuilder) JumpPage(jumpPage string) *H5ToAirportBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5ToAirportBuilder) AirportId(airportId string) *H5ToAirportBuilder {
	builder.airportId = airportId
	builder.airportIdSet = true
	return builder
}
func (builder *H5ToAirportBuilder) CityId(cityId string) *H5ToAirportBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *H5ToAirportBuilder) LatFrom(latFrom string) *H5ToAirportBuilder {
	builder.latFrom = latFrom
	builder.latFromSet = true
	return builder
}
func (builder *H5ToAirportBuilder) LngFrom(lngFrom string) *H5ToAirportBuilder {
	builder.lngFrom = lngFrom
	builder.lngFromSet = true
	return builder
}
func (builder *H5ToAirportBuilder) PoiFromName(poiFromName string) *H5ToAirportBuilder {
	builder.poiFromName = poiFromName
	builder.poiFromNameSet = true
	return builder
}
func (builder *H5ToAirportBuilder) PassengerPhone(passengerPhone string) *H5ToAirportBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5ToAirportBuilder) PassengerName(passengerName string) *H5ToAirportBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *H5ToAirportBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5ToAirportBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5ToAirportBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5ToAirportBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5ToAirportBuilder) Build() *H5ToAirport {
	data := &H5ToAirport{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.airportIdSet {
		data.AirportId = &builder.airportId
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
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
