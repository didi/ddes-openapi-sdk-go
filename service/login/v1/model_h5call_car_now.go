package v1

// H5CallCarNow 实时用车
type H5CallCarNow struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 跳转场景页面(参考参数使用说明) 实时用车为 callCarNow
	ApprovalId            *string `json:"approval_id,omitempty"`              // 差旅、申请用车，必传
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
	RestrictPassenger     *string `json:"restrict_passenger,omitempty"`       // 如需锁定乘车人，必传
	RequireLevelList      *string `json:"require_level_list,omitempty"`       // 如需控制车型选择及默认勾选，必传
	AppendCar             *string `json:"append_car,omitempty"`               // 等待追加车型
	CallbackInfo          *string `json:"callback_info,omitempty"`            // 客户自定义自段，可关联订单信息
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5CallCarNowBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 跳转场景页面(参考参数使用说明) 实时用车为 callCarNow
	jumpPageSet              bool
	approvalId               string // 差旅、申请用车，必传
	approvalIdSet            bool
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
	restrictPassenger        string // 如需锁定乘车人，必传
	restrictPassengerSet     bool
	requireLevelList         string // 如需控制车型选择及默认勾选，必传
	requireLevelListSet      bool
	appendCar                string // 等待追加车型
	appendCarSet             bool
	callbackInfo             string // 客户自定义自段，可关联订单信息
	callbackInfoSet          bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5CallCarNowBuilder() *H5CallCarNowBuilder {
	return &H5CallCarNowBuilder{}
}
func (builder *H5CallCarNowBuilder) RuleId(ruleId string) *H5CallCarNowBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) JumpPage(jumpPage string) *H5CallCarNowBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) ApprovalId(approvalId string) *H5CallCarNowBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) LatFrom(latFrom string) *H5CallCarNowBuilder {
	builder.latFrom = latFrom
	builder.latFromSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) LngFrom(lngFrom string) *H5CallCarNowBuilder {
	builder.lngFrom = lngFrom
	builder.lngFromSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) PoiFromName(poiFromName string) *H5CallCarNowBuilder {
	builder.poiFromName = poiFromName
	builder.poiFromNameSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) ToCityId(toCityId string) *H5CallCarNowBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) LatTo(latTo string) *H5CallCarNowBuilder {
	builder.latTo = latTo
	builder.latToSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) LngTo(lngTo string) *H5CallCarNowBuilder {
	builder.lngTo = lngTo
	builder.lngToSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) PoiToName(poiToName string) *H5CallCarNowBuilder {
	builder.poiToName = poiToName
	builder.poiToNameSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) CityId(cityId string) *H5CallCarNowBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) PassengerPhone(passengerPhone string) *H5CallCarNowBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) PassengerName(passengerName string) *H5CallCarNowBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5CallCarNowBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) RestrictPassenger(restrictPassenger string) *H5CallCarNowBuilder {
	builder.restrictPassenger = restrictPassenger
	builder.restrictPassengerSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) RequireLevelList(requireLevelList string) *H5CallCarNowBuilder {
	builder.requireLevelList = requireLevelList
	builder.requireLevelListSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) AppendCar(appendCar string) *H5CallCarNowBuilder {
	builder.appendCar = appendCar
	builder.appendCarSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) CallbackInfo(callbackInfo string) *H5CallCarNowBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *H5CallCarNowBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5CallCarNowBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5CallCarNowBuilder) Build() *H5CallCarNow {
	data := &H5CallCarNow{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
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
	if builder.restrictPassengerSet {
		data.RestrictPassenger = &builder.restrictPassenger
	}
	if builder.requireLevelListSet {
		data.RequireLevelList = &builder.requireLevelList
	}
	if builder.appendCarSet {
		data.AppendCar = &builder.appendCar
	}
	if builder.callbackInfoSet {
		data.CallbackInfo = &builder.callbackInfo
	}
	if builder.hideEstimatePriceFlagSet {
		data.HideEstimatePriceFlag = &builder.hideEstimatePriceFlag
	}
	return data
}
