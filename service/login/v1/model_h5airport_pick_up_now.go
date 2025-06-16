package v1

// H5AirportPickUpNow 实时接机
type H5AirportPickUpNow struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 实时接机为airportPickUpNow
	AirportId             *string `json:"airport_id,omitempty"`               // 机场ID，如不传需手动选择
	ApprovalId            *string `json:"approval_id,omitempty"`              // 差旅、申请用车，必传
	ToCityId              *string `json:"to_city_id,omitempty"`               // 如不传以手动选择地址为终点
	LatTo                 *string `json:"lat_to,omitempty"`                   // 如不传以手动选择地址为终点
	LngTo                 *string `json:"lng_to,omitempty"`                   // 如不传以手动选择地址为终点
	PoiToName             *string `json:"poi_to_name,omitempty"`              // 如不传以手动选择地址为终点
	CityId                *string `json:"city_id,omitempty"`                  // 如不传以手动选择地址为终点
	PassengerPhone        *string `json:"passenger_phone,omitempty"`          // 乘车人手机号，代叫必传
	RestrictPoiFlag       *string `json:"restrict_poi_flag,omitempty"`        // 如需锁定地址，必传
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5AirportPickUpNowBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 实时接机为airportPickUpNow
	jumpPageSet              bool
	airportId                string // 机场ID，如不传需手动选择
	airportIdSet             bool
	approvalId               string // 差旅、申请用车，必传
	approvalIdSet            bool
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
	restrictPoiFlag          string // 如需锁定地址，必传
	restrictPoiFlagSet       bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5AirportPickUpNowBuilder() *H5AirportPickUpNowBuilder {
	return &H5AirportPickUpNowBuilder{}
}
func (builder *H5AirportPickUpNowBuilder) RuleId(ruleId string) *H5AirportPickUpNowBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) JumpPage(jumpPage string) *H5AirportPickUpNowBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) AirportId(airportId string) *H5AirportPickUpNowBuilder {
	builder.airportId = airportId
	builder.airportIdSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) ApprovalId(approvalId string) *H5AirportPickUpNowBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) ToCityId(toCityId string) *H5AirportPickUpNowBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) LatTo(latTo string) *H5AirportPickUpNowBuilder {
	builder.latTo = latTo
	builder.latToSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) LngTo(lngTo string) *H5AirportPickUpNowBuilder {
	builder.lngTo = lngTo
	builder.lngToSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) PoiToName(poiToName string) *H5AirportPickUpNowBuilder {
	builder.poiToName = poiToName
	builder.poiToNameSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) CityId(cityId string) *H5AirportPickUpNowBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) PassengerPhone(passengerPhone string) *H5AirportPickUpNowBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5AirportPickUpNowBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5AirportPickUpNowBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5AirportPickUpNowBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5AirportPickUpNowBuilder) Build() *H5AirportPickUpNow {
	data := &H5AirportPickUpNow{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.airportIdSet {
		data.AirportId = &builder.airportId
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
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
	if builder.restrictPoiFlagSet {
		data.RestrictPoiFlag = &builder.restrictPoiFlag
	}
	if builder.hideEstimatePriceFlagSet {
		data.HideEstimatePriceFlag = &builder.hideEstimatePriceFlag
	}
	return data
}
