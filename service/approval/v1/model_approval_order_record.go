package v1

// ApprovalOrderRecord 定义订单记录消息
type ApprovalOrderRecord struct {
	OrderId          *string                `json:"order_id,omitempty"`           // 订单ID
	ApprovalId       *string                `json:"approval_id,omitempty"`        // 审批单ID
	OutApprovalId    *string                `json:"out_approval_id,omitempty"`    // 企业主内部审批单ID
	RuleId           *string                `json:"rule_id,omitempty"`            // 规则ID
	RegulationId     *string                `json:"regulation_id,omitempty"`      // 制度ID
	SceneType        *string                `json:"scene_type,omitempty"`         // 因公出行场景（0-个人用车、1-商务出行、2-差旅、3-加班、4-办公地点通勤、91-代叫车、92-接送机）
	OrderCreateTime  *string                `json:"order_create_time,omitempty"`  // 订单创建时间，如：1675602713
	BeginChargeTime  *string                `json:"begin_charge_time,omitempty"`  // 订单开始计费时间，如：1675582413
	FinishTime       *string                `json:"finish_time,omitempty"`        // 订单结束计费时间，如：1675583857
	DepartureTime    *string                `json:"departure_time,omitempty"`     // 预约上车时间，如：1675582103
	UseCarType       *int32                 `json:"use_car_type,omitempty"`       // 用车方式（2:专车，3:快车）
	CarLevel         *int32                 `json:"car_level,omitempty"`          // 车型，如（100舒适型，400六座商务, 200行政级,600普通快车,900优享快车,1000豪华车,1100企业出租车,2000优选出租车）
	CityName         *string                `json:"city_name,omitempty"`          // 出发城市名称
	StartName        *string                `json:"start_name,omitempty"`         // 出发地名称
	EndName          *string                `json:"end_name,omitempty"`           // 目的地名称
	ActualStartName  *string                `json:"actual_start_name,omitempty"`  // 实际出发地（司机点击开始的位置）
	ActualEndName    *string                `json:"actual_end_name,omitempty"`    // 实际目的地（司机点击结束的位置）
	ActualFlat       *string                `json:"actual_flat,omitempty"`        // 实际出发地纬度（司机点击开始的位置）
	ActualFlng       *string                `json:"actual_flng,omitempty"`        // 实际出发地经度（司机点击开始的位置）
	ActualTlat       *string                `json:"actual_tlat,omitempty"`        // 实际目的地纬度（司机点击结束的位置）
	ActualTlng       *string                `json:"actual_tlng,omitempty"`        // 实际目的地经度（司机点击结束的位置）
	PayTime          *string                `json:"pay_time,omitempty"`           // 支付时间
	OrderStatus      *int32                 `json:"order_status,omitempty"`       // 订单状态（1-发单失败、2-已支付、3-已退款、4-已取消、5-待支付、6-部分支付、7-部分退款）
	PayType          *int32                 `json:"pay_type,omitempty"`           // 支付方式（0-企业支付、1-个人支付需报销、2-混合支付（企业和个人各支付部分））
	IsInvoice        *int32                 `json:"is_invoice,omitempty"`         // 开票状态（0-未开、1-开过 ）
	CallPhone        *string                `json:"call_phone,omitempty"`         // 叫车人手机号
	PassengerPhone   *string                `json:"passenger_phone,omitempty"`    // 乘车人手机号
	TotalPrice       *string                `json:"total_price,omitempty"`        // 订单总金额
	ActualPrice      *float32               `json:"actual_price,omitempty"`       // 订单实付金额（总金额-券折扣金额）
	RefundPrice      *string                `json:"refund_price,omitempty"`       // 订单总退款金额（企业支付退款+个人支付退款）
	CompanyPay       *string                `json:"company_pay,omitempty"`        // 企业支付应付金额
	PersonalPay      *string                `json:"personal_pay,omitempty"`       // 个人支付应付金额
	CompanyRealPay   *string                `json:"company_real_pay,omitempty"`   // 企业支付实付金额
	PersonalRealPay  *string                `json:"personal_real_pay,omitempty"`  // 个人支付实付金额
	CompanyRefund    *string                `json:"company_refund,omitempty"`     // 企业支付退款金额
	PersonalRefund   *string                `json:"personal_refund,omitempty"`    // 个人支付退款金额
	BudgetCenterId   *string                `json:"budget_center_id,omitempty"`   // 成本中心ID
	ExtraInfo        *string                `json:"extra_info,omitempty"`         // 扩展信息，自定义字段
	SceneId          *string                `json:"scene_id,omitempty"`           // 场景ID，101 申请单模式差旅；102 申请用车；103 自驾同行
	PreTotalFee      *string                `json:"pre_total_fee,omitempty"`      // 特惠快车原价金额 单位元
	FixedDiscountFee *string                `json:"fixed_discount_fee,omitempty"` // 特惠快车尊享折扣金额 单位元
	DiscountFee      *string                `json:"discount_fee,omitempty"`       // 补贴金额 单位元 在使用企业券时，补贴金额 = 尊享折扣 + 券金额，其他情况直接等于尊享折扣
	BudgetCenterList []BudgetCenterListItem `json:"budget_center_list,omitempty"` // 多成本中心
}

type ApprovalOrderRecordBuilder struct {
	orderId             string // 订单ID
	orderIdSet          bool
	approvalId          string // 审批单ID
	approvalIdSet       bool
	outApprovalId       string // 企业主内部审批单ID
	outApprovalIdSet    bool
	ruleId              string // 规则ID
	ruleIdSet           bool
	regulationId        string // 制度ID
	regulationIdSet     bool
	sceneType           string // 因公出行场景（0-个人用车、1-商务出行、2-差旅、3-加班、4-办公地点通勤、91-代叫车、92-接送机）
	sceneTypeSet        bool
	orderCreateTime     string // 订单创建时间，如：1675602713
	orderCreateTimeSet  bool
	beginChargeTime     string // 订单开始计费时间，如：1675582413
	beginChargeTimeSet  bool
	finishTime          string // 订单结束计费时间，如：1675583857
	finishTimeSet       bool
	departureTime       string // 预约上车时间，如：1675582103
	departureTimeSet    bool
	useCarType          int32 // 用车方式（2:专车，3:快车）
	useCarTypeSet       bool
	carLevel            int32 // 车型，如（100舒适型，400六座商务, 200行政级,600普通快车,900优享快车,1000豪华车,1100企业出租车,2000优选出租车）
	carLevelSet         bool
	cityName            string // 出发城市名称
	cityNameSet         bool
	startName           string // 出发地名称
	startNameSet        bool
	endName             string // 目的地名称
	endNameSet          bool
	actualStartName     string // 实际出发地（司机点击开始的位置）
	actualStartNameSet  bool
	actualEndName       string // 实际目的地（司机点击结束的位置）
	actualEndNameSet    bool
	actualFlat          string // 实际出发地纬度（司机点击开始的位置）
	actualFlatSet       bool
	actualFlng          string // 实际出发地经度（司机点击开始的位置）
	actualFlngSet       bool
	actualTlat          string // 实际目的地纬度（司机点击结束的位置）
	actualTlatSet       bool
	actualTlng          string // 实际目的地经度（司机点击结束的位置）
	actualTlngSet       bool
	payTime             string // 支付时间
	payTimeSet          bool
	orderStatus         int32 // 订单状态（1-发单失败、2-已支付、3-已退款、4-已取消、5-待支付、6-部分支付、7-部分退款）
	orderStatusSet      bool
	payType             int32 // 支付方式（0-企业支付、1-个人支付需报销、2-混合支付（企业和个人各支付部分））
	payTypeSet          bool
	isInvoice           int32 // 开票状态（0-未开、1-开过 ）
	isInvoiceSet        bool
	callPhone           string // 叫车人手机号
	callPhoneSet        bool
	passengerPhone      string // 乘车人手机号
	passengerPhoneSet   bool
	totalPrice          string // 订单总金额
	totalPriceSet       bool
	actualPrice         float32 // 订单实付金额（总金额-券折扣金额）
	actualPriceSet      bool
	refundPrice         string // 订单总退款金额（企业支付退款+个人支付退款）
	refundPriceSet      bool
	companyPay          string // 企业支付应付金额
	companyPaySet       bool
	personalPay         string // 个人支付应付金额
	personalPaySet      bool
	companyRealPay      string // 企业支付实付金额
	companyRealPaySet   bool
	personalRealPay     string // 个人支付实付金额
	personalRealPaySet  bool
	companyRefund       string // 企业支付退款金额
	companyRefundSet    bool
	personalRefund      string // 个人支付退款金额
	personalRefundSet   bool
	budgetCenterId      string // 成本中心ID
	budgetCenterIdSet   bool
	extraInfo           string // 扩展信息，自定义字段
	extraInfoSet        bool
	sceneId             string // 场景ID，101 申请单模式差旅；102 申请用车；103 自驾同行
	sceneIdSet          bool
	preTotalFee         string // 特惠快车原价金额 单位元
	preTotalFeeSet      bool
	fixedDiscountFee    string // 特惠快车尊享折扣金额 单位元
	fixedDiscountFeeSet bool
	discountFee         string // 补贴金额 单位元 在使用企业券时，补贴金额 = 尊享折扣 + 券金额，其他情况直接等于尊享折扣
	discountFeeSet      bool
	budgetCenterList    []BudgetCenterListItem // 多成本中心
	budgetCenterListSet bool
}

func NewApprovalOrderRecordBuilder() *ApprovalOrderRecordBuilder {
	return &ApprovalOrderRecordBuilder{}
}
func (builder *ApprovalOrderRecordBuilder) OrderId(orderId string) *ApprovalOrderRecordBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ApprovalId(approvalId string) *ApprovalOrderRecordBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) OutApprovalId(outApprovalId string) *ApprovalOrderRecordBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) RuleId(ruleId string) *ApprovalOrderRecordBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) RegulationId(regulationId string) *ApprovalOrderRecordBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) SceneType(sceneType string) *ApprovalOrderRecordBuilder {
	builder.sceneType = sceneType
	builder.sceneTypeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) OrderCreateTime(orderCreateTime string) *ApprovalOrderRecordBuilder {
	builder.orderCreateTime = orderCreateTime
	builder.orderCreateTimeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) BeginChargeTime(beginChargeTime string) *ApprovalOrderRecordBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) FinishTime(finishTime string) *ApprovalOrderRecordBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) DepartureTime(departureTime string) *ApprovalOrderRecordBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) UseCarType(useCarType int32) *ApprovalOrderRecordBuilder {
	builder.useCarType = useCarType
	builder.useCarTypeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CarLevel(carLevel int32) *ApprovalOrderRecordBuilder {
	builder.carLevel = carLevel
	builder.carLevelSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CityName(cityName string) *ApprovalOrderRecordBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) StartName(startName string) *ApprovalOrderRecordBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) EndName(endName string) *ApprovalOrderRecordBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualStartName(actualStartName string) *ApprovalOrderRecordBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualEndName(actualEndName string) *ApprovalOrderRecordBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualFlat(actualFlat string) *ApprovalOrderRecordBuilder {
	builder.actualFlat = actualFlat
	builder.actualFlatSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualFlng(actualFlng string) *ApprovalOrderRecordBuilder {
	builder.actualFlng = actualFlng
	builder.actualFlngSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualTlat(actualTlat string) *ApprovalOrderRecordBuilder {
	builder.actualTlat = actualTlat
	builder.actualTlatSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualTlng(actualTlng string) *ApprovalOrderRecordBuilder {
	builder.actualTlng = actualTlng
	builder.actualTlngSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PayTime(payTime string) *ApprovalOrderRecordBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) OrderStatus(orderStatus int32) *ApprovalOrderRecordBuilder {
	builder.orderStatus = orderStatus
	builder.orderStatusSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PayType(payType int32) *ApprovalOrderRecordBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) IsInvoice(isInvoice int32) *ApprovalOrderRecordBuilder {
	builder.isInvoice = isInvoice
	builder.isInvoiceSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CallPhone(callPhone string) *ApprovalOrderRecordBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PassengerPhone(passengerPhone string) *ApprovalOrderRecordBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) TotalPrice(totalPrice string) *ApprovalOrderRecordBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ActualPrice(actualPrice float32) *ApprovalOrderRecordBuilder {
	builder.actualPrice = actualPrice
	builder.actualPriceSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) RefundPrice(refundPrice string) *ApprovalOrderRecordBuilder {
	builder.refundPrice = refundPrice
	builder.refundPriceSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CompanyPay(companyPay string) *ApprovalOrderRecordBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PersonalPay(personalPay string) *ApprovalOrderRecordBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CompanyRealPay(companyRealPay string) *ApprovalOrderRecordBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PersonalRealPay(personalRealPay string) *ApprovalOrderRecordBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) CompanyRefund(companyRefund string) *ApprovalOrderRecordBuilder {
	builder.companyRefund = companyRefund
	builder.companyRefundSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PersonalRefund(personalRefund string) *ApprovalOrderRecordBuilder {
	builder.personalRefund = personalRefund
	builder.personalRefundSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) BudgetCenterId(budgetCenterId string) *ApprovalOrderRecordBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) ExtraInfo(extraInfo string) *ApprovalOrderRecordBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) SceneId(sceneId string) *ApprovalOrderRecordBuilder {
	builder.sceneId = sceneId
	builder.sceneIdSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) PreTotalFee(preTotalFee string) *ApprovalOrderRecordBuilder {
	builder.preTotalFee = preTotalFee
	builder.preTotalFeeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) FixedDiscountFee(fixedDiscountFee string) *ApprovalOrderRecordBuilder {
	builder.fixedDiscountFee = fixedDiscountFee
	builder.fixedDiscountFeeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) DiscountFee(discountFee string) *ApprovalOrderRecordBuilder {
	builder.discountFee = discountFee
	builder.discountFeeSet = true
	return builder
}
func (builder *ApprovalOrderRecordBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *ApprovalOrderRecordBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}

func (builder *ApprovalOrderRecordBuilder) Build() *ApprovalOrderRecord {
	data := &ApprovalOrderRecord{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.sceneTypeSet {
		data.SceneType = &builder.sceneType
	}
	if builder.orderCreateTimeSet {
		data.OrderCreateTime = &builder.orderCreateTime
	}
	if builder.beginChargeTimeSet {
		data.BeginChargeTime = &builder.beginChargeTime
	}
	if builder.finishTimeSet {
		data.FinishTime = &builder.finishTime
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
	}
	if builder.useCarTypeSet {
		data.UseCarType = &builder.useCarType
	}
	if builder.carLevelSet {
		data.CarLevel = &builder.carLevel
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.endNameSet {
		data.EndName = &builder.endName
	}
	if builder.actualStartNameSet {
		data.ActualStartName = &builder.actualStartName
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
	if builder.payTimeSet {
		data.PayTime = &builder.payTime
	}
	if builder.orderStatusSet {
		data.OrderStatus = &builder.orderStatus
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.isInvoiceSet {
		data.IsInvoice = &builder.isInvoice
	}
	if builder.callPhoneSet {
		data.CallPhone = &builder.callPhone
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
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
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.companyRefundSet {
		data.CompanyRefund = &builder.companyRefund
	}
	if builder.personalRefundSet {
		data.PersonalRefund = &builder.personalRefund
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.sceneIdSet {
		data.SceneId = &builder.sceneId
	}
	if builder.preTotalFeeSet {
		data.PreTotalFee = &builder.preTotalFee
	}
	if builder.fixedDiscountFeeSet {
		data.FixedDiscountFee = &builder.fixedDiscountFee
	}
	if builder.discountFeeSet {
		data.DiscountFee = &builder.discountFee
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	return data
}
