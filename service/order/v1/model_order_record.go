package v1

// OrderRecord 订单信息(getOrder)
type OrderRecord struct {
	OrderId               *string                `json:"order_id,omitempty"`                 // 订单id
	CallPhone             *string                `json:"call_phone,omitempty"`               // 叫车人手机号
	PassengerPhone        *string                `json:"passenger_phone,omitempty"`          // 乘车人手机号
	PassengerName         *string                `json:"passenger_name,omitempty"`           // 乘车人名字
	MemberId              *string                `json:"member_id,omitempty"`                // 叫车人员工在滴滴的id
	City                  *string                `json:"city,omitempty"`                     // 城市id
	CityName              *string                `json:"city_name,omitempty"`                // 城市名称
	StartName             *string                `json:"start_name,omitempty"`               // 出发地名称
	ActualStartName       *string                `json:"actual_start_name,omitempty"`        // 实际出发地名称
	EndName               *string                `json:"end_name,omitempty"`                 // 目的的名称
	ActualEndName         *string                `json:"actual_end_name,omitempty"`          // 实际目的地名称
	ActualFlat            *string                `json:"actual_flat,omitempty"`              // 出发地纬度
	ActualFlng            *string                `json:"actual_flng,omitempty"`              // 出发地经度
	ActualTlat            *string                `json:"actual_tlat,omitempty"`              // 目的地纬度
	ActualTlng            *string                `json:"actual_tlng,omitempty"`              // 目的地经度
	NormalDistance        *string                `json:"normal_distance,omitempty"`          // 总里程
	CreateTime            *string                `json:"create_time,omitempty"`              // 下单时间(订单生成时间)
	DepartureTime         *string                `json:"departure_time,omitempty"`           // 出发时间(预约单为预约出发时间)
	StriveTime            *string                `json:"strive_time,omitempty"`              // 接单时间
	MeetTime              *string                `json:"meet_time,omitempty"`                // 接驾时间
	BeginChargeTime       *string                `json:"begin_charge_time,omitempty"`        // 开始计价时间
	FinishTime            *string                `json:"finish_time,omitempty"`              // 结束计价时间
	PayTime               *string                `json:"pay_time,omitempty"`                 // 支付时间
	RefundTime            *string                `json:"refund_time,omitempty"`              // 退款时间
	RequireLevel          *string                `json:"require_level,omitempty"`            // 用车车型（100舒适型，400六座商务, 200行政级,600普通快车,900优享快车,1000豪华车,1100企业出租车,2000优选出租车，40600企业特价快车） 以上为常用枚举。全量请参考附录展示的枚举。
	UseCarType            *string                `json:"use_car_type,omitempty"`             // 用车车型大类（2-专车、3-快车、5-豪华车）
	SubUseCarType         *int32                 `json:"sub_use_car_type,omitempty"`         // 用车场景， 0 - 市内用车，1 - 接送机，2 - 接送站，3 - 接送汽车站，4 - 接送渡口
	Status                *string                `json:"status,omitempty"`                   // 订单状态（2-已支付、3-已退款、4-已取消、7-部分退款）
	PayType               *string                `json:"pay_type,omitempty"`                 // 支付类型（0-企业支付、1-个人支付、2-混合支付）
	OrderSource           *string                `json:"order_source,omitempty"`             // 订单来源（0-Web、1-滴滴出行App、2-H5、3-OpenAPI、4-企业APP、5-邀约券、6-SDK、8-Webapp、10-企业app H5）
	SupplierType          *string                `json:"supplier_type,omitempty"`            // 枚举值 0自营 1API 2花小猪自营 3鸿鹄三方
	PricingMode           *int32                 `json:"pricing_mode,omitempty"`             // 计价模型（0-实时计价、1-一口价、2-两口价(一口价变成实时计价)）
	IsCarpool             *string                `json:"is_carpool,omitempty"`               // 是否为拼车（0-非拼车，1-拼车）
	IsInvoice             *int32                 `json:"is_invoice,omitempty"`               // 是否已开票（0-未开票、1-已开票）
	EstimatePrice         *string                `json:"estimate_price,omitempty"`           // 预估金额，单位：元
	TotalPrice            *string                `json:"total_price,omitempty"`              // 订单总金额，单位：元
	ActualPrice           *string                `json:"actual_price,omitempty"`             // 实付金额（总金额-券抵扣金额），单位：元
	RefundPrice           *string                `json:"refund_price,omitempty"`             // 退款金额，单位：元
	CompanyPay            *string                `json:"company_pay,omitempty"`              // 公司支付金额，单位：元
	CompanyCardPay        *string                `json:"company_card_pay,omitempty"`         // 公司出行卡支付金额，单位：元
	PersonalPay           *string                `json:"personal_pay,omitempty"`             // 个人支付金额，单位：元
	CompanyRealPay        *string                `json:"company_real_pay,omitempty"`         // 公司实付金额，单位：元
	CompanyCardRealPay    *string                `json:"company_card_real_pay,omitempty"`    // 公司出行卡实付金额，单位：元
	PersonalRealPay       *string                `json:"personal_real_pay,omitempty"`        // 个人实付金额，单位：元
	CompanyRealRefund     *string                `json:"company_real_refund,omitempty"`      // 公司实际退款金额，单位：元
	CompanyCardRealRefund *string                `json:"company_card_real_refund,omitempty"` // 公司出行卡实际退款金额，单位：元
	PersonalRealRefund    *string                `json:"personal_real_refund,omitempty"`     // 个人实际退款金额
	BudgetCenterId        *string                `json:"budget_center_id,omitempty"`         // 成本中心id
	UseCarConfigId        *string                `json:"use_car_config_id,omitempty"`        // 用车规则id
	IsSensitive           *int32                 `json:"is_sensitive,omitempty"`             // 是否为敏感订单（0-不是敏感订单、1-敏感订单）
	SensitiveRuleId       *int64                 `json:"sensitive_rule_id,omitempty"`        // 命中的敏感规则id
	ApprovalId            *string                `json:"approval_id,omitempty"`              // 滴滴内部审批单id
	OutApprovalId         *string                `json:"out_approval_id,omitempty"`          // 接入方审批单id
	CallbackInfo          *string                `json:"callback_info,omitempty"`            // 下单时用户携带的callback_info
	ExtraInfo             *string                `json:"extra_info,omitempty"`               // 扩展信息，自定义字段(创建审批单的扩展信息)
	ExtendFieldList       *string                `json:"extend_field_list,omitempty"`        // 申请单基础信息, 拓展信息list,json 串
	EncryptedInfo         *string                `json:"encrypted_info,omitempty"`           // 订单加密信息（可忽略）
	Remark                *string                `json:"remark,omitempty"`                   // 备注信息
	IsAbnormal            *int32                 `json:"is_abnormal,omitempty"`              // 是否为敏感订单(0-否，1-是) need_abnormal_msg=1时返回
	AbnormalExplanation   *string                `json:"abnormal_explanation,omitempty"`     // 敏感订单员工说明文案；need_abnormal_msg=1时返回
	AbnormalType          *string                `json:"abnormal_type,omitempty"`            // 敏感订单类型；need_abnormal_msg=1时返回
	ReasonType            *string                `json:"reason_type,omitempty"`              // 敏感订单解释原因类型；need_abnormal_msg=1时返回
	OperationType         *int32                 `json:"operation_type,omitempty"`           // 0 、无，限仅开通敏感订单1.0 或 未开通敏感订单（含1.0+2.0） 1 、标记，（已开通敏感订单2.0，配置标记 or 智能） 3 、解释说明，（已开通敏感订单2.0，配置解释说明 or 智能） 4 、个人支付，（已开通配置敏感订单2.0，配置个人支付 or 智能）need_abnormal_msg为 1时返回
	RuleName              *string                `json:"rule_name,omitempty"`                // 制度名称（need_rule_info = 1 时返回）
	RegulationId          *int64                 `json:"regulation_id,omitempty"`            // 制度ID （need_rule_info = 1 时返回
	Type                  *string                `json:"type,omitempty"`                     // 订单类型(0:实时、1:预约)
	DestCity              *string                `json:"dest_city,omitempty"`                // 目的城市ID
	DestCityName          *string                `json:"dest_city_name,omitempty"`           // 目的城市名称
	GroupId               *string                `json:"group_id,omitempty"`                 // 部门ID
	BudgetExtraInfo       *string                `json:"budget_extra_info,omitempty"`        // 项目扩展信息的自定义字段；最长不大于 500 字符；(必须为json字符串，json解析后不能为空)；
	CallEmployeeNumber    *string                `json:"call_employee_number,omitempty"`     // 叫车人工号
	BudgetCenterList      []BudgetCenterListItem `json:"budget_center_list,omitempty"`       // 参考budget_center_list对象
	ActualStartAddress    *string                `json:"actual_start_address,omitempty"`     // 实际出发地址
	ActualEndAddress      *string                `json:"actual_end_address,omitempty"`       // 实际到达地址
	DepartureCountyId     *int32                 `json:"departure_county_id,omitempty"`      // 出发区县 ID
	DepartureCountyName   *string                `json:"departure_county_name,omitempty"`    // 出发区县名称
	DestinationCountyId   *int32                 `json:"destination_county_id,omitempty"`    // 到达区县 ID
	DestinationCountyName *string                `json:"destination_county_name,omitempty"`  // 到达区县 名称
	StopoverPoints        []StopoverPoint        `json:"stopover_points,omitempty"`          // 参考stopover_points对象
}

type OrderRecordBuilder struct {
	orderId                  string // 订单id
	orderIdSet               bool
	callPhone                string // 叫车人手机号
	callPhoneSet             bool
	passengerPhone           string // 乘车人手机号
	passengerPhoneSet        bool
	passengerName            string // 乘车人名字
	passengerNameSet         bool
	memberId                 string // 叫车人员工在滴滴的id
	memberIdSet              bool
	city                     string // 城市id
	citySet                  bool
	cityName                 string // 城市名称
	cityNameSet              bool
	startName                string // 出发地名称
	startNameSet             bool
	actualStartName          string // 实际出发地名称
	actualStartNameSet       bool
	endName                  string // 目的的名称
	endNameSet               bool
	actualEndName            string // 实际目的地名称
	actualEndNameSet         bool
	actualFlat               string // 出发地纬度
	actualFlatSet            bool
	actualFlng               string // 出发地经度
	actualFlngSet            bool
	actualTlat               string // 目的地纬度
	actualTlatSet            bool
	actualTlng               string // 目的地经度
	actualTlngSet            bool
	normalDistance           string // 总里程
	normalDistanceSet        bool
	createTime               string // 下单时间(订单生成时间)
	createTimeSet            bool
	departureTime            string // 出发时间(预约单为预约出发时间)
	departureTimeSet         bool
	striveTime               string // 接单时间
	striveTimeSet            bool
	meetTime                 string // 接驾时间
	meetTimeSet              bool
	beginChargeTime          string // 开始计价时间
	beginChargeTimeSet       bool
	finishTime               string // 结束计价时间
	finishTimeSet            bool
	payTime                  string // 支付时间
	payTimeSet               bool
	refundTime               string // 退款时间
	refundTimeSet            bool
	requireLevel             string // 用车车型（100舒适型，400六座商务, 200行政级,600普通快车,900优享快车,1000豪华车,1100企业出租车,2000优选出租车，40600企业特价快车） 以上为常用枚举。全量请参考附录展示的枚举。
	requireLevelSet          bool
	useCarType               string // 用车车型大类（2-专车、3-快车、5-豪华车）
	useCarTypeSet            bool
	subUseCarType            int32 // 用车场景， 0 - 市内用车，1 - 接送机，2 - 接送站，3 - 接送汽车站，4 - 接送渡口
	subUseCarTypeSet         bool
	status                   string // 订单状态（2-已支付、3-已退款、4-已取消、7-部分退款）
	statusSet                bool
	payType                  string // 支付类型（0-企业支付、1-个人支付、2-混合支付）
	payTypeSet               bool
	orderSource              string // 订单来源（0-Web、1-滴滴出行App、2-H5、3-OpenAPI、4-企业APP、5-邀约券、6-SDK、8-Webapp、10-企业app H5）
	orderSourceSet           bool
	supplierType             string // 枚举值 0自营 1API 2花小猪自营 3鸿鹄三方
	supplierTypeSet          bool
	pricingMode              int32 // 计价模型（0-实时计价、1-一口价、2-两口价(一口价变成实时计价)）
	pricingModeSet           bool
	isCarpool                string // 是否为拼车（0-非拼车，1-拼车）
	isCarpoolSet             bool
	isInvoice                int32 // 是否已开票（0-未开票、1-已开票）
	isInvoiceSet             bool
	estimatePrice            string // 预估金额，单位：元
	estimatePriceSet         bool
	totalPrice               string // 订单总金额，单位：元
	totalPriceSet            bool
	actualPrice              string // 实付金额（总金额-券抵扣金额），单位：元
	actualPriceSet           bool
	refundPrice              string // 退款金额，单位：元
	refundPriceSet           bool
	companyPay               string // 公司支付金额，单位：元
	companyPaySet            bool
	companyCardPay           string // 公司出行卡支付金额，单位：元
	companyCardPaySet        bool
	personalPay              string // 个人支付金额，单位：元
	personalPaySet           bool
	companyRealPay           string // 公司实付金额，单位：元
	companyRealPaySet        bool
	companyCardRealPay       string // 公司出行卡实付金额，单位：元
	companyCardRealPaySet    bool
	personalRealPay          string // 个人实付金额，单位：元
	personalRealPaySet       bool
	companyRealRefund        string // 公司实际退款金额，单位：元
	companyRealRefundSet     bool
	companyCardRealRefund    string // 公司出行卡实际退款金额，单位：元
	companyCardRealRefundSet bool
	personalRealRefund       string // 个人实际退款金额
	personalRealRefundSet    bool
	budgetCenterId           string // 成本中心id
	budgetCenterIdSet        bool
	useCarConfigId           string // 用车规则id
	useCarConfigIdSet        bool
	isSensitive              int32 // 是否为敏感订单（0-不是敏感订单、1-敏感订单）
	isSensitiveSet           bool
	sensitiveRuleId          int64 // 命中的敏感规则id
	sensitiveRuleIdSet       bool
	approvalId               string // 滴滴内部审批单id
	approvalIdSet            bool
	outApprovalId            string // 接入方审批单id
	outApprovalIdSet         bool
	callbackInfo             string // 下单时用户携带的callback_info
	callbackInfoSet          bool
	extraInfo                string // 扩展信息，自定义字段(创建审批单的扩展信息)
	extraInfoSet             bool
	extendFieldList          string // 申请单基础信息, 拓展信息list,json 串
	extendFieldListSet       bool
	encryptedInfo            string // 订单加密信息（可忽略）
	encryptedInfoSet         bool
	remark                   string // 备注信息
	remarkSet                bool
	isAbnormal               int32 // 是否为敏感订单(0-否，1-是) need_abnormal_msg=1时返回
	isAbnormalSet            bool
	abnormalExplanation      string // 敏感订单员工说明文案；need_abnormal_msg=1时返回
	abnormalExplanationSet   bool
	abnormalType             string // 敏感订单类型；need_abnormal_msg=1时返回
	abnormalTypeSet          bool
	reasonType               string // 敏感订单解释原因类型；need_abnormal_msg=1时返回
	reasonTypeSet            bool
	operationType            int32 // 0 、无，限仅开通敏感订单1.0 或 未开通敏感订单（含1.0+2.0） 1 、标记，（已开通敏感订单2.0，配置标记 or 智能） 3 、解释说明，（已开通敏感订单2.0，配置解释说明 or 智能） 4 、个人支付，（已开通配置敏感订单2.0，配置个人支付 or 智能）need_abnormal_msg为 1时返回
	operationTypeSet         bool
	ruleName                 string // 制度名称（need_rule_info = 1 时返回）
	ruleNameSet              bool
	regulationId             int64 // 制度ID （need_rule_info = 1 时返回
	regulationIdSet          bool
	type_                    string // 订单类型(0:实时、1:预约)
	type_Set                 bool
	destCity                 string // 目的城市ID
	destCitySet              bool
	destCityName             string // 目的城市名称
	destCityNameSet          bool
	groupId                  string // 部门ID
	groupIdSet               bool
	budgetExtraInfo          string // 项目扩展信息的自定义字段；最长不大于 500 字符；(必须为json字符串，json解析后不能为空)；
	budgetExtraInfoSet       bool
	callEmployeeNumber       string // 叫车人工号
	callEmployeeNumberSet    bool
	budgetCenterList         []BudgetCenterListItem // 参考budget_center_list对象
	budgetCenterListSet      bool
	actualStartAddress       string // 实际出发地址
	actualStartAddressSet    bool
	actualEndAddress         string // 实际到达地址
	actualEndAddressSet      bool
	departureCountyId        int32 // 出发区县 ID
	departureCountyIdSet     bool
	departureCountyName      string // 出发区县名称
	departureCountyNameSet   bool
	destinationCountyId      int32 // 到达区县 ID
	destinationCountyIdSet   bool
	destinationCountyName    string // 到达区县 名称
	destinationCountyNameSet bool
	stopoverPoints           []StopoverPoint // 参考stopover_points对象
	stopoverPointsSet        bool
}

func NewOrderRecordBuilder() *OrderRecordBuilder {
	return &OrderRecordBuilder{}
}
func (builder *OrderRecordBuilder) OrderId(orderId string) *OrderRecordBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) CallPhone(callPhone string) *OrderRecordBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *OrderRecordBuilder) PassengerPhone(passengerPhone string) *OrderRecordBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *OrderRecordBuilder) PassengerName(passengerName string) *OrderRecordBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) MemberId(memberId string) *OrderRecordBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) City(city string) *OrderRecordBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *OrderRecordBuilder) CityName(cityName string) *OrderRecordBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) StartName(startName string) *OrderRecordBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualStartName(actualStartName string) *OrderRecordBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) EndName(endName string) *OrderRecordBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualEndName(actualEndName string) *OrderRecordBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualFlat(actualFlat string) *OrderRecordBuilder {
	builder.actualFlat = actualFlat
	builder.actualFlatSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualFlng(actualFlng string) *OrderRecordBuilder {
	builder.actualFlng = actualFlng
	builder.actualFlngSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualTlat(actualTlat string) *OrderRecordBuilder {
	builder.actualTlat = actualTlat
	builder.actualTlatSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualTlng(actualTlng string) *OrderRecordBuilder {
	builder.actualTlng = actualTlng
	builder.actualTlngSet = true
	return builder
}
func (builder *OrderRecordBuilder) NormalDistance(normalDistance string) *OrderRecordBuilder {
	builder.normalDistance = normalDistance
	builder.normalDistanceSet = true
	return builder
}
func (builder *OrderRecordBuilder) CreateTime(createTime string) *OrderRecordBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) DepartureTime(departureTime string) *OrderRecordBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) StriveTime(striveTime string) *OrderRecordBuilder {
	builder.striveTime = striveTime
	builder.striveTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) MeetTime(meetTime string) *OrderRecordBuilder {
	builder.meetTime = meetTime
	builder.meetTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) BeginChargeTime(beginChargeTime string) *OrderRecordBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) FinishTime(finishTime string) *OrderRecordBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) PayTime(payTime string) *OrderRecordBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) RefundTime(refundTime string) *OrderRecordBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *OrderRecordBuilder) RequireLevel(requireLevel string) *OrderRecordBuilder {
	builder.requireLevel = requireLevel
	builder.requireLevelSet = true
	return builder
}
func (builder *OrderRecordBuilder) UseCarType(useCarType string) *OrderRecordBuilder {
	builder.useCarType = useCarType
	builder.useCarTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) SubUseCarType(subUseCarType int32) *OrderRecordBuilder {
	builder.subUseCarType = subUseCarType
	builder.subUseCarTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) Status(status string) *OrderRecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *OrderRecordBuilder) PayType(payType string) *OrderRecordBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) OrderSource(orderSource string) *OrderRecordBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *OrderRecordBuilder) SupplierType(supplierType string) *OrderRecordBuilder {
	builder.supplierType = supplierType
	builder.supplierTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) PricingMode(pricingMode int32) *OrderRecordBuilder {
	builder.pricingMode = pricingMode
	builder.pricingModeSet = true
	return builder
}
func (builder *OrderRecordBuilder) IsCarpool(isCarpool string) *OrderRecordBuilder {
	builder.isCarpool = isCarpool
	builder.isCarpoolSet = true
	return builder
}
func (builder *OrderRecordBuilder) IsInvoice(isInvoice int32) *OrderRecordBuilder {
	builder.isInvoice = isInvoice
	builder.isInvoiceSet = true
	return builder
}
func (builder *OrderRecordBuilder) EstimatePrice(estimatePrice string) *OrderRecordBuilder {
	builder.estimatePrice = estimatePrice
	builder.estimatePriceSet = true
	return builder
}
func (builder *OrderRecordBuilder) TotalPrice(totalPrice string) *OrderRecordBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualPrice(actualPrice string) *OrderRecordBuilder {
	builder.actualPrice = actualPrice
	builder.actualPriceSet = true
	return builder
}
func (builder *OrderRecordBuilder) RefundPrice(refundPrice string) *OrderRecordBuilder {
	builder.refundPrice = refundPrice
	builder.refundPriceSet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyPay(companyPay string) *OrderRecordBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyCardPay(companyCardPay string) *OrderRecordBuilder {
	builder.companyCardPay = companyCardPay
	builder.companyCardPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) PersonalPay(personalPay string) *OrderRecordBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyRealPay(companyRealPay string) *OrderRecordBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyCardRealPay(companyCardRealPay string) *OrderRecordBuilder {
	builder.companyCardRealPay = companyCardRealPay
	builder.companyCardRealPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) PersonalRealPay(personalRealPay string) *OrderRecordBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyRealRefund(companyRealRefund string) *OrderRecordBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *OrderRecordBuilder) CompanyCardRealRefund(companyCardRealRefund string) *OrderRecordBuilder {
	builder.companyCardRealRefund = companyCardRealRefund
	builder.companyCardRealRefundSet = true
	return builder
}
func (builder *OrderRecordBuilder) PersonalRealRefund(personalRealRefund string) *OrderRecordBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *OrderRecordBuilder) BudgetCenterId(budgetCenterId string) *OrderRecordBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) UseCarConfigId(useCarConfigId string) *OrderRecordBuilder {
	builder.useCarConfigId = useCarConfigId
	builder.useCarConfigIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) IsSensitive(isSensitive int32) *OrderRecordBuilder {
	builder.isSensitive = isSensitive
	builder.isSensitiveSet = true
	return builder
}
func (builder *OrderRecordBuilder) SensitiveRuleId(sensitiveRuleId int64) *OrderRecordBuilder {
	builder.sensitiveRuleId = sensitiveRuleId
	builder.sensitiveRuleIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) ApprovalId(approvalId string) *OrderRecordBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) OutApprovalId(outApprovalId string) *OrderRecordBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) CallbackInfo(callbackInfo string) *OrderRecordBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *OrderRecordBuilder) ExtraInfo(extraInfo string) *OrderRecordBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *OrderRecordBuilder) ExtendFieldList(extendFieldList string) *OrderRecordBuilder {
	builder.extendFieldList = extendFieldList
	builder.extendFieldListSet = true
	return builder
}
func (builder *OrderRecordBuilder) EncryptedInfo(encryptedInfo string) *OrderRecordBuilder {
	builder.encryptedInfo = encryptedInfo
	builder.encryptedInfoSet = true
	return builder
}
func (builder *OrderRecordBuilder) Remark(remark string) *OrderRecordBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *OrderRecordBuilder) IsAbnormal(isAbnormal int32) *OrderRecordBuilder {
	builder.isAbnormal = isAbnormal
	builder.isAbnormalSet = true
	return builder
}
func (builder *OrderRecordBuilder) AbnormalExplanation(abnormalExplanation string) *OrderRecordBuilder {
	builder.abnormalExplanation = abnormalExplanation
	builder.abnormalExplanationSet = true
	return builder
}
func (builder *OrderRecordBuilder) AbnormalType(abnormalType string) *OrderRecordBuilder {
	builder.abnormalType = abnormalType
	builder.abnormalTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) ReasonType(reasonType string) *OrderRecordBuilder {
	builder.reasonType = reasonType
	builder.reasonTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) OperationType(operationType int32) *OrderRecordBuilder {
	builder.operationType = operationType
	builder.operationTypeSet = true
	return builder
}
func (builder *OrderRecordBuilder) RuleName(ruleName string) *OrderRecordBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) RegulationId(regulationId int64) *OrderRecordBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) Type(type_ string) *OrderRecordBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *OrderRecordBuilder) DestCity(destCity string) *OrderRecordBuilder {
	builder.destCity = destCity
	builder.destCitySet = true
	return builder
}
func (builder *OrderRecordBuilder) DestCityName(destCityName string) *OrderRecordBuilder {
	builder.destCityName = destCityName
	builder.destCityNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) GroupId(groupId string) *OrderRecordBuilder {
	builder.groupId = groupId
	builder.groupIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) BudgetExtraInfo(budgetExtraInfo string) *OrderRecordBuilder {
	builder.budgetExtraInfo = budgetExtraInfo
	builder.budgetExtraInfoSet = true
	return builder
}
func (builder *OrderRecordBuilder) CallEmployeeNumber(callEmployeeNumber string) *OrderRecordBuilder {
	builder.callEmployeeNumber = callEmployeeNumber
	builder.callEmployeeNumberSet = true
	return builder
}
func (builder *OrderRecordBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *OrderRecordBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualStartAddress(actualStartAddress string) *OrderRecordBuilder {
	builder.actualStartAddress = actualStartAddress
	builder.actualStartAddressSet = true
	return builder
}
func (builder *OrderRecordBuilder) ActualEndAddress(actualEndAddress string) *OrderRecordBuilder {
	builder.actualEndAddress = actualEndAddress
	builder.actualEndAddressSet = true
	return builder
}
func (builder *OrderRecordBuilder) DepartureCountyId(departureCountyId int32) *OrderRecordBuilder {
	builder.departureCountyId = departureCountyId
	builder.departureCountyIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) DepartureCountyName(departureCountyName string) *OrderRecordBuilder {
	builder.departureCountyName = departureCountyName
	builder.departureCountyNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) DestinationCountyId(destinationCountyId int32) *OrderRecordBuilder {
	builder.destinationCountyId = destinationCountyId
	builder.destinationCountyIdSet = true
	return builder
}
func (builder *OrderRecordBuilder) DestinationCountyName(destinationCountyName string) *OrderRecordBuilder {
	builder.destinationCountyName = destinationCountyName
	builder.destinationCountyNameSet = true
	return builder
}
func (builder *OrderRecordBuilder) StopoverPoints(stopoverPoints []StopoverPoint) *OrderRecordBuilder {
	builder.stopoverPoints = stopoverPoints
	builder.stopoverPointsSet = true
	return builder
}

func (builder *OrderRecordBuilder) Build() *OrderRecord {
	data := &OrderRecord{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.callPhoneSet {
		data.CallPhone = &builder.callPhone
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.actualStartNameSet {
		data.ActualStartName = &builder.actualStartName
	}
	if builder.endNameSet {
		data.EndName = &builder.endName
	}
	if builder.actualEndNameSet {
		data.ActualEndName = &builder.actualEndName
	}
	if builder.actualFlatSet {
		data.ActualFlat = &builder.actualFlat
	}
	if builder.actualFlngSet {
		data.ActualFlng = &builder.actualFlng
	}
	if builder.actualTlatSet {
		data.ActualTlat = &builder.actualTlat
	}
	if builder.actualTlngSet {
		data.ActualTlng = &builder.actualTlng
	}
	if builder.normalDistanceSet {
		data.NormalDistance = &builder.normalDistance
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
	}
	if builder.striveTimeSet {
		data.StriveTime = &builder.striveTime
	}
	if builder.meetTimeSet {
		data.MeetTime = &builder.meetTime
	}
	if builder.beginChargeTimeSet {
		data.BeginChargeTime = &builder.beginChargeTime
	}
	if builder.finishTimeSet {
		data.FinishTime = &builder.finishTime
	}
	if builder.payTimeSet {
		data.PayTime = &builder.payTime
	}
	if builder.refundTimeSet {
		data.RefundTime = &builder.refundTime
	}
	if builder.requireLevelSet {
		data.RequireLevel = &builder.requireLevel
	}
	if builder.useCarTypeSet {
		data.UseCarType = &builder.useCarType
	}
	if builder.subUseCarTypeSet {
		data.SubUseCarType = &builder.subUseCarType
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.orderSourceSet {
		data.OrderSource = &builder.orderSource
	}
	if builder.supplierTypeSet {
		data.SupplierType = &builder.supplierType
	}
	if builder.pricingModeSet {
		data.PricingMode = &builder.pricingMode
	}
	if builder.isCarpoolSet {
		data.IsCarpool = &builder.isCarpool
	}
	if builder.isInvoiceSet {
		data.IsInvoice = &builder.isInvoice
	}
	if builder.estimatePriceSet {
		data.EstimatePrice = &builder.estimatePrice
	}
	if builder.totalPriceSet {
		data.TotalPrice = &builder.totalPrice
	}
	if builder.actualPriceSet {
		data.ActualPrice = &builder.actualPrice
	}
	if builder.refundPriceSet {
		data.RefundPrice = &builder.refundPrice
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.companyCardPaySet {
		data.CompanyCardPay = &builder.companyCardPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.companyCardRealPaySet {
		data.CompanyCardRealPay = &builder.companyCardRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.companyCardRealRefundSet {
		data.CompanyCardRealRefund = &builder.companyCardRealRefund
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.useCarConfigIdSet {
		data.UseCarConfigId = &builder.useCarConfigId
	}
	if builder.isSensitiveSet {
		data.IsSensitive = &builder.isSensitive
	}
	if builder.sensitiveRuleIdSet {
		data.SensitiveRuleId = &builder.sensitiveRuleId
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.callbackInfoSet {
		data.CallbackInfo = &builder.callbackInfo
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.extendFieldListSet {
		data.ExtendFieldList = &builder.extendFieldList
	}
	if builder.encryptedInfoSet {
		data.EncryptedInfo = &builder.encryptedInfo
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.isAbnormalSet {
		data.IsAbnormal = &builder.isAbnormal
	}
	if builder.abnormalExplanationSet {
		data.AbnormalExplanation = &builder.abnormalExplanation
	}
	if builder.abnormalTypeSet {
		data.AbnormalType = &builder.abnormalType
	}
	if builder.reasonTypeSet {
		data.ReasonType = &builder.reasonType
	}
	if builder.operationTypeSet {
		data.OperationType = &builder.operationType
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.destCitySet {
		data.DestCity = &builder.destCity
	}
	if builder.destCityNameSet {
		data.DestCityName = &builder.destCityName
	}
	if builder.groupIdSet {
		data.GroupId = &builder.groupId
	}
	if builder.budgetExtraInfoSet {
		data.BudgetExtraInfo = &builder.budgetExtraInfo
	}
	if builder.callEmployeeNumberSet {
		data.CallEmployeeNumber = &builder.callEmployeeNumber
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	if builder.actualStartAddressSet {
		data.ActualStartAddress = &builder.actualStartAddress
	}
	if builder.actualEndAddressSet {
		data.ActualEndAddress = &builder.actualEndAddress
	}
	if builder.departureCountyIdSet {
		data.DepartureCountyId = &builder.departureCountyId
	}
	if builder.departureCountyNameSet {
		data.DepartureCountyName = &builder.departureCountyName
	}
	if builder.destinationCountyIdSet {
		data.DestinationCountyId = &builder.destinationCountyId
	}
	if builder.destinationCountyNameSet {
		data.DestinationCountyName = &builder.destinationCountyName
	}
	if builder.stopoverPointsSet {
		data.StopoverPoints = builder.stopoverPoints
	}
	return data
}
