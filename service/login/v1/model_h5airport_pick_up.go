package v1

// H5AirportPickUp 预约接机
type H5AirportPickUp struct {
	RuleId                *string `json:"rule_id,omitempty"`                  // 制度规则 id
	JumpPage              *string `json:"jumpPage,omitempty"`                 // 预约接机为airportPickUp
	FlightNo              *string `json:"flight_no,omitempty"`                // 航班号，如不传需手动选择
	FlightDepartDate      *string `json:"flight_depart_date,omitempty"`       // 航班出发日期，如 flight_no有参数，此字段必传
	ApprovalId            *string `json:"approval_id,omitempty"`              // 差旅、申请用车，必传
	ToCityId              *string `json:"to_city_id,omitempty"`               // 如不传以手动选择地址为终点
	LatTo                 *string `json:"lat_to,omitempty"`                   // 如不传以手动选择地址为终点
	LngTo                 *string `json:"lng_to,omitempty"`                   // 如不传以手动选择地址为终点
	PoiToName             *string `json:"poi_to_name,omitempty"`              // 如不传以手动选择地址为终点
	PassengerPhone        *string `json:"passenger_phone,omitempty"`          // 乘车人手机号，代叫必传
	PassengerName         *string `json:"passenger_name,omitempty"`           // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	RestrictPoiFlag       *string `json:"restrict_poi_flag,omitempty"`        // 如需锁定地址，必传
	HideEstimatePriceFlag *string `json:"hide_estimate_price_flag,omitempty"` // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
}

type H5AirportPickUpBuilder struct {
	ruleId                   string // 制度规则 id
	ruleIdSet                bool
	jumpPage                 string // 预约接机为airportPickUp
	jumpPageSet              bool
	flightNo                 string // 航班号，如不传需手动选择
	flightNoSet              bool
	flightDepartDate         string // 航班出发日期，如 flight_no有参数，此字段必传
	flightDepartDateSet      bool
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
	passengerPhone           string // 乘车人手机号，代叫必传
	passengerPhoneSet        bool
	passengerName            string // 设置passenger_phone后可传入 如不传入默认使用在 滴滴管理后台填入的名称或乘客+手机号后四位
	passengerNameSet         bool
	restrictPoiFlag          string // 如需锁定地址，必传
	restrictPoiFlagSet       bool
	hideEstimatePriceFlag    string // 是否在预估页面隐藏价格相关信息 1 是 0 否，不传等于0
	hideEstimatePriceFlagSet bool
}

func NewH5AirportPickUpBuilder() *H5AirportPickUpBuilder {
	return &H5AirportPickUpBuilder{}
}
func (builder *H5AirportPickUpBuilder) RuleId(ruleId string) *H5AirportPickUpBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) JumpPage(jumpPage string) *H5AirportPickUpBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) FlightNo(flightNo string) *H5AirportPickUpBuilder {
	builder.flightNo = flightNo
	builder.flightNoSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) FlightDepartDate(flightDepartDate string) *H5AirportPickUpBuilder {
	builder.flightDepartDate = flightDepartDate
	builder.flightDepartDateSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) ApprovalId(approvalId string) *H5AirportPickUpBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) ToCityId(toCityId string) *H5AirportPickUpBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) LatTo(latTo string) *H5AirportPickUpBuilder {
	builder.latTo = latTo
	builder.latToSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) LngTo(lngTo string) *H5AirportPickUpBuilder {
	builder.lngTo = lngTo
	builder.lngToSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) PoiToName(poiToName string) *H5AirportPickUpBuilder {
	builder.poiToName = poiToName
	builder.poiToNameSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) PassengerPhone(passengerPhone string) *H5AirportPickUpBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) PassengerName(passengerName string) *H5AirportPickUpBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) RestrictPoiFlag(restrictPoiFlag string) *H5AirportPickUpBuilder {
	builder.restrictPoiFlag = restrictPoiFlag
	builder.restrictPoiFlagSet = true
	return builder
}
func (builder *H5AirportPickUpBuilder) HideEstimatePriceFlag(hideEstimatePriceFlag string) *H5AirportPickUpBuilder {
	builder.hideEstimatePriceFlag = hideEstimatePriceFlag
	builder.hideEstimatePriceFlagSet = true
	return builder
}

func (builder *H5AirportPickUpBuilder) Build() *H5AirportPickUp {
	data := &H5AirportPickUp{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.flightNoSet {
		data.FlightNo = &builder.flightNo
	}
	if builder.flightDepartDateSet {
		data.FlightDepartDate = &builder.flightDepartDate
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
