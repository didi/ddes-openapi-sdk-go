package v1

// TransactionOfWangYCItem 账单交易明细(网约车)
type TransactionOfWangYCItem struct {
	TransactionId            *int64   `json:"transaction_id,omitempty"`              // 资金流水id
	Type                     *string  `json:"type,omitempty"`                        // 资金操作类型
	OutMoney                 *float32 `json:"out_money,omitempty"`                   // 操作金额 - 支出
	InMoney                  *float32 `json:"in_money,omitempty"`                    // 操作金额 - 收入
	CreateTimeDate           *string  `json:"create_time_date,omitempty"`            // 创建时间
	OrderId                  *int64   `json:"order_id,omitempty"`                    // 订单号
	Remark                   *string  `json:"remark,omitempty"`                      // 备注
	IsCurrentTerm            *string  `json:"is_current_term,omitempty"`             // 所属账期
	CompanyOrderId           *int64   `json:"company_order_id,omitempty"`            // 企业订单号
	OutOrderId               *int64   `json:"out_order_id,omitempty"`                // 专快订单号
	OrderSource              *string  `json:"order_source,omitempty"`                // 订单来源
	Status                   *string  `json:"status,omitempty"`                      // 订单状态
	UseCarType               *string  `json:"use_car_type,omitempty"`                // 用车类型
	RequireLevel             *string  `json:"require_level,omitempty"`               // 用车类型(明细)
	MemberName               *string  `json:"member_name,omitempty"`                 // 下单人姓名
	EmployeeNumber           *string  `json:"employee_number,omitempty"`             // 下单人员工ID
	MemberEmail              *string  `json:"member_email,omitempty"`                // 企业邮箱
	CallPhone                *string  `json:"call_phone,omitempty"`                  // 下单人电话
	DepartmentName           *string  `json:"department_name,omitempty"`             // 所在部门(新)
	OutBudgetId              *string  `json:"out_budget_id,omitempty"`               // 所在部门(新)编号
	BranchName               *string  `json:"branch_name,omitempty"`                 // 下单人分公司(旧)
	Department               *string  `json:"department,omitempty"`                  // 下单人部门(旧)
	PassengerName            *string  `json:"passenger_name,omitempty"`              // 乘车人姓名
	PassengerPhone           *string  `json:"passenger_phone,omitempty"`             // 乘车人电话
	PayType                  *string  `json:"pay_type,omitempty"`                    // 支付方式
	TotalPrice               *float32 `json:"total_price,omitempty"`                 // 订单总金额
	CompanyRealPay           *float32 `json:"company_real_pay,omitempty"`            // 企业实付金额
	PersonalRealPay          *float32 `json:"personal_real_pay,omitempty"`           // 个人实付金额
	VoucherPay               *float32 `json:"voucher_pay,omitempty"`                 // 用券抵扣金额
	Cost                     *float32 `json:"cost,omitempty"`                        // 车费金额
	DiscountAfterPrice       *float32 `json:"discount_after_price,omitempty"`        // 平台使用费金额
	CompanyRealRefund        *float32 `json:"company_real_refund,omitempty"`         // 企业实退金额
	OrderType                *string  `json:"order_type,omitempty"`                  // 订单类型
	CreateTime               *string  `json:"create_time,omitempty"`                 // 下单时间
	DepartureTime            *string  `json:"departure_time,omitempty"`              // 计划用车时间
	MeetTime                 *string  `json:"meet_time,omitempty"`                   // 司机接单时间
	BeginChargeTime          *string  `json:"begin_charge_time,omitempty"`           // 开始计费时间
	FinishTime               *string  `json:"finish_time,omitempty"`                 // 结束计费时间
	PayTime                  *string  `json:"pay_time,omitempty"`                    // 支付时间
	RefundTime               *string  `json:"refund_time,omitempty"`                 // 退款时间
	City                     *string  `json:"city,omitempty"`                        // 用车城市
	DestinationCity          *string  `json:"destination_city,omitempty"`            // 到达城市
	NormalDistance           *float32 `json:"normal_distance,omitempty"`             // 用车行驶距离
	StartAddress             *string  `json:"start_address,omitempty"`               // 出发地地址
	StartName                *string  `json:"start_name,omitempty"`                  // 出发地名称
	EndAddress               *string  `json:"end_address,omitempty"`                 // 目的地地址
	EndName                  *string  `json:"end_name,omitempty"`                    // 目的地名称
	ActualStartName          *string  `json:"actual_start_name,omitempty"`           // 实际出发地
	ActualEndName            *string  `json:"actual_end_name,omitempty"`             // 实际目的地
	RuleName                 *string  `json:"rule_name,omitempty"`                   // 用车权限
	UseCarSrv                *string  `json:"use_car_srv,omitempty"`                 // 用车方式
	BudgetCenterName         *string  `json:"budget_center_name,omitempty"`          // 成本中心名称
	BudgetCenterId           *string  `json:"budget_center_id,omitempty"`            // 成本中心id
	UseCarRemark             *string  `json:"use_car_remark,omitempty"`              // 用车备注
	OrderAdditionalRemark    *string  `json:"order_additional_remark,omitempty"`     // 补充说明
	ApprovalReason           *string  `json:"approval_reason,omitempty"`             // 申请备注
	ApprovalCreateTime       *string  `json:"approval_create_time,omitempty"`        // 申请提交时间
	ApprovalLogs             *string  `json:"approval_logs,omitempty"`               // 审批人1|审批时间
	ApprovalId               *string  `json:"approval_id,omitempty"`                 // 外部审批单ID
	UpgradeType              *string  `json:"upgrade_type,omitempty"`                // 是否自费升舱
	IsFixedPrice             *string  `json:"is_fixed_price,omitempty"`              // 是否一口价
	IsSelfDrive              *string  `json:"is_self_drive,omitempty"`               // 是否自驾同行
	IsDiscountOrder          *string  `json:"is_discount_order,omitempty"`           // 是否折扣订单
	IsCarpool                *string  `json:"is_carpool,omitempty"`                  // 是否拼车
	IsReassignment           *string  `json:"is_reassignment,omitempty"`             // 是否改派
	CallbackInfo             *string  `json:"callback_info,omitempty"`               // API企业自定义
	AfterApprovalResult      *string  `json:"after_approval_result,omitempty"`       // 行后审批结果
	AfterApprovalContent     *string  `json:"after_approval_content,omitempty"`      // 审批历史
	ApprovalExtraInfo        *string  `json:"approval_extra_info,omitempty"`         // 审批单自定义
	UseCarTypeName           *string  `json:"use_car_type_name,omitempty"`           // 用车场景细分
	UseCarTypeV2             *string  `json:"use_car_type_v2,omitempty"`             // 用车场景
	UserPayType              *string  `json:"user_pay_type,omitempty"`               // 用户支付类型
	SubUseCarSrv             *string  `json:"sub_use_car_srv,omitempty"`             // 附加字段
	CompanyId                *int64   `json:"company_id,omitempty"`                  // 所属公司主体ID
	CompanyName              *string  `json:"company_name,omitempty"`                // 公司名称
	SpsPickUpServiceFee      *float32 `json:"sps_pick_up_service_fee,omitempty"`     // 司机举牌接机服务费
	IsUnusual                *string  `json:"is_unusual,omitempty"`                  // 是否异常
	UnusualType              *string  `json:"unusual_type,omitempty"`                // 异常类型
	UnusualContent           *string  `json:"unusual_content,omitempty"`             // 异常说明
	PositionName             *string  `json:"position_name,omitempty"`               // 职级
	InstitutionName          *string  `json:"institution_name,omitempty"`            // 用车制度
	UseCarService            *string  `json:"use_car_service,omitempty"`             // 用车服务类型
	ExInfo01                 *string  `json:"ex_info_01,omitempty"`                  // 拓展字段1
	ExInfo02                 *string  `json:"ex_info_02,omitempty"`                  // 拓展字段2
	ExInfo03                 *string  `json:"ex_info_03,omitempty"`                  // 拓展字段3
	LegalEntityId            *int64   `json:"legal_entity_id,omitempty"`             // 所属主体id
	LegalEntityName          *string  `json:"legal_entity_name,omitempty"`           // 所属主体名称
	BeforeApprovalContent    *string  `json:"before_approval_content,omitempty"`     // 待支付审批历史
	OriginalPrice            *float32 `json:"original_price,omitempty"`              // 原价
	OneTimeOfferSubsidy      *float32 `json:"one_time_offer_subsidy,omitempty"`      // 尊享折扣
	SubsidyAmount            *float32 `json:"subsidy_amount,omitempty"`              // 补贴金额
	VoucherDeductionTypeName *string  `json:"voucher_deduction_type_name,omitempty"` // 券抵扣类型
	ResidentCityNames        *string  `json:"resident_city_names,omitempty"`         // 常驻城市
	PassengerMemberNumber    *string  `json:"passenger_member_number,omitempty"`     // 乘车人工号
	PassengerMemberId        *string  `json:"passenger_member_id,omitempty"`         // 乘车人id
}

type TransactionOfWangYCItemBuilder struct {
	transactionId               int64 // 资金流水id
	transactionIdSet            bool
	type_                       string // 资金操作类型
	type_Set                    bool
	outMoney                    float32 // 操作金额 - 支出
	outMoneySet                 bool
	inMoney                     float32 // 操作金额 - 收入
	inMoneySet                  bool
	createTimeDate              string // 创建时间
	createTimeDateSet           bool
	orderId                     int64 // 订单号
	orderIdSet                  bool
	remark                      string // 备注
	remarkSet                   bool
	isCurrentTerm               string // 所属账期
	isCurrentTermSet            bool
	companyOrderId              int64 // 企业订单号
	companyOrderIdSet           bool
	outOrderId                  int64 // 专快订单号
	outOrderIdSet               bool
	orderSource                 string // 订单来源
	orderSourceSet              bool
	status                      string // 订单状态
	statusSet                   bool
	useCarType                  string // 用车类型
	useCarTypeSet               bool
	requireLevel                string // 用车类型(明细)
	requireLevelSet             bool
	memberName                  string // 下单人姓名
	memberNameSet               bool
	employeeNumber              string // 下单人员工ID
	employeeNumberSet           bool
	memberEmail                 string // 企业邮箱
	memberEmailSet              bool
	callPhone                   string // 下单人电话
	callPhoneSet                bool
	departmentName              string // 所在部门(新)
	departmentNameSet           bool
	outBudgetId                 string // 所在部门(新)编号
	outBudgetIdSet              bool
	branchName                  string // 下单人分公司(旧)
	branchNameSet               bool
	department                  string // 下单人部门(旧)
	departmentSet               bool
	passengerName               string // 乘车人姓名
	passengerNameSet            bool
	passengerPhone              string // 乘车人电话
	passengerPhoneSet           bool
	payType                     string // 支付方式
	payTypeSet                  bool
	totalPrice                  float32 // 订单总金额
	totalPriceSet               bool
	companyRealPay              float32 // 企业实付金额
	companyRealPaySet           bool
	personalRealPay             float32 // 个人实付金额
	personalRealPaySet          bool
	voucherPay                  float32 // 用券抵扣金额
	voucherPaySet               bool
	cost                        float32 // 车费金额
	costSet                     bool
	discountAfterPrice          float32 // 平台使用费金额
	discountAfterPriceSet       bool
	companyRealRefund           float32 // 企业实退金额
	companyRealRefundSet        bool
	orderType                   string // 订单类型
	orderTypeSet                bool
	createTime                  string // 下单时间
	createTimeSet               bool
	departureTime               string // 计划用车时间
	departureTimeSet            bool
	meetTime                    string // 司机接单时间
	meetTimeSet                 bool
	beginChargeTime             string // 开始计费时间
	beginChargeTimeSet          bool
	finishTime                  string // 结束计费时间
	finishTimeSet               bool
	payTime                     string // 支付时间
	payTimeSet                  bool
	refundTime                  string // 退款时间
	refundTimeSet               bool
	city                        string // 用车城市
	citySet                     bool
	destinationCity             string // 到达城市
	destinationCitySet          bool
	normalDistance              float32 // 用车行驶距离
	normalDistanceSet           bool
	startAddress                string // 出发地地址
	startAddressSet             bool
	startName                   string // 出发地名称
	startNameSet                bool
	endAddress                  string // 目的地地址
	endAddressSet               bool
	endName                     string // 目的地名称
	endNameSet                  bool
	actualStartName             string // 实际出发地
	actualStartNameSet          bool
	actualEndName               string // 实际目的地
	actualEndNameSet            bool
	ruleName                    string // 用车权限
	ruleNameSet                 bool
	useCarSrv                   string // 用车方式
	useCarSrvSet                bool
	budgetCenterName            string // 成本中心名称
	budgetCenterNameSet         bool
	budgetCenterId              string // 成本中心id
	budgetCenterIdSet           bool
	useCarRemark                string // 用车备注
	useCarRemarkSet             bool
	orderAdditionalRemark       string // 补充说明
	orderAdditionalRemarkSet    bool
	approvalReason              string // 申请备注
	approvalReasonSet           bool
	approvalCreateTime          string // 申请提交时间
	approvalCreateTimeSet       bool
	approvalLogs                string // 审批人1|审批时间
	approvalLogsSet             bool
	approvalId                  string // 外部审批单ID
	approvalIdSet               bool
	upgradeType                 string // 是否自费升舱
	upgradeTypeSet              bool
	isFixedPrice                string // 是否一口价
	isFixedPriceSet             bool
	isSelfDrive                 string // 是否自驾同行
	isSelfDriveSet              bool
	isDiscountOrder             string // 是否折扣订单
	isDiscountOrderSet          bool
	isCarpool                   string // 是否拼车
	isCarpoolSet                bool
	isReassignment              string // 是否改派
	isReassignmentSet           bool
	callbackInfo                string // API企业自定义
	callbackInfoSet             bool
	afterApprovalResult         string // 行后审批结果
	afterApprovalResultSet      bool
	afterApprovalContent        string // 审批历史
	afterApprovalContentSet     bool
	approvalExtraInfo           string // 审批单自定义
	approvalExtraInfoSet        bool
	useCarTypeName              string // 用车场景细分
	useCarTypeNameSet           bool
	useCarTypeV2                string // 用车场景
	useCarTypeV2Set             bool
	userPayType                 string // 用户支付类型
	userPayTypeSet              bool
	subUseCarSrv                string // 附加字段
	subUseCarSrvSet             bool
	companyId                   int64 // 所属公司主体ID
	companyIdSet                bool
	companyName                 string // 公司名称
	companyNameSet              bool
	spsPickUpServiceFee         float32 // 司机举牌接机服务费
	spsPickUpServiceFeeSet      bool
	isUnusual                   string // 是否异常
	isUnusualSet                bool
	unusualType                 string // 异常类型
	unusualTypeSet              bool
	unusualContent              string // 异常说明
	unusualContentSet           bool
	positionName                string // 职级
	positionNameSet             bool
	institutionName             string // 用车制度
	institutionNameSet          bool
	useCarService               string // 用车服务类型
	useCarServiceSet            bool
	exInfo01                    string // 拓展字段1
	exInfo01Set                 bool
	exInfo02                    string // 拓展字段2
	exInfo02Set                 bool
	exInfo03                    string // 拓展字段3
	exInfo03Set                 bool
	legalEntityId               int64 // 所属主体id
	legalEntityIdSet            bool
	legalEntityName             string // 所属主体名称
	legalEntityNameSet          bool
	beforeApprovalContent       string // 待支付审批历史
	beforeApprovalContentSet    bool
	originalPrice               float32 // 原价
	originalPriceSet            bool
	oneTimeOfferSubsidy         float32 // 尊享折扣
	oneTimeOfferSubsidySet      bool
	subsidyAmount               float32 // 补贴金额
	subsidyAmountSet            bool
	voucherDeductionTypeName    string // 券抵扣类型
	voucherDeductionTypeNameSet bool
	residentCityNames           string // 常驻城市
	residentCityNamesSet        bool
	passengerMemberNumber       string // 乘车人工号
	passengerMemberNumberSet    bool
	passengerMemberId           string // 乘车人id
	passengerMemberIdSet        bool
}

func NewTransactionOfWangYCItemBuilder() *TransactionOfWangYCItemBuilder {
	return &TransactionOfWangYCItemBuilder{}
}
func (builder *TransactionOfWangYCItemBuilder) TransactionId(transactionId int64) *TransactionOfWangYCItemBuilder {
	builder.transactionId = transactionId
	builder.transactionIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) Type(type_ string) *TransactionOfWangYCItemBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OutMoney(outMoney float32) *TransactionOfWangYCItemBuilder {
	builder.outMoney = outMoney
	builder.outMoneySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) InMoney(inMoney float32) *TransactionOfWangYCItemBuilder {
	builder.inMoney = inMoney
	builder.inMoneySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CreateTimeDate(createTimeDate string) *TransactionOfWangYCItemBuilder {
	builder.createTimeDate = createTimeDate
	builder.createTimeDateSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OrderId(orderId int64) *TransactionOfWangYCItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) Remark(remark string) *TransactionOfWangYCItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsCurrentTerm(isCurrentTerm string) *TransactionOfWangYCItemBuilder {
	builder.isCurrentTerm = isCurrentTerm
	builder.isCurrentTermSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CompanyOrderId(companyOrderId int64) *TransactionOfWangYCItemBuilder {
	builder.companyOrderId = companyOrderId
	builder.companyOrderIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OutOrderId(outOrderId int64) *TransactionOfWangYCItemBuilder {
	builder.outOrderId = outOrderId
	builder.outOrderIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OrderSource(orderSource string) *TransactionOfWangYCItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) Status(status string) *TransactionOfWangYCItemBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarType(useCarType string) *TransactionOfWangYCItemBuilder {
	builder.useCarType = useCarType
	builder.useCarTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) RequireLevel(requireLevel string) *TransactionOfWangYCItemBuilder {
	builder.requireLevel = requireLevel
	builder.requireLevelSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) MemberName(memberName string) *TransactionOfWangYCItemBuilder {
	builder.memberName = memberName
	builder.memberNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) EmployeeNumber(employeeNumber string) *TransactionOfWangYCItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) MemberEmail(memberEmail string) *TransactionOfWangYCItemBuilder {
	builder.memberEmail = memberEmail
	builder.memberEmailSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CallPhone(callPhone string) *TransactionOfWangYCItemBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) DepartmentName(departmentName string) *TransactionOfWangYCItemBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OutBudgetId(outBudgetId string) *TransactionOfWangYCItemBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) BranchName(branchName string) *TransactionOfWangYCItemBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) Department(department string) *TransactionOfWangYCItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PassengerName(passengerName string) *TransactionOfWangYCItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PassengerPhone(passengerPhone string) *TransactionOfWangYCItemBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PayType(payType string) *TransactionOfWangYCItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) TotalPrice(totalPrice float32) *TransactionOfWangYCItemBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CompanyRealPay(companyRealPay float32) *TransactionOfWangYCItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PersonalRealPay(personalRealPay float32) *TransactionOfWangYCItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) VoucherPay(voucherPay float32) *TransactionOfWangYCItemBuilder {
	builder.voucherPay = voucherPay
	builder.voucherPaySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) Cost(cost float32) *TransactionOfWangYCItemBuilder {
	builder.cost = cost
	builder.costSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) DiscountAfterPrice(discountAfterPrice float32) *TransactionOfWangYCItemBuilder {
	builder.discountAfterPrice = discountAfterPrice
	builder.discountAfterPriceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CompanyRealRefund(companyRealRefund float32) *TransactionOfWangYCItemBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OrderType(orderType string) *TransactionOfWangYCItemBuilder {
	builder.orderType = orderType
	builder.orderTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CreateTime(createTime string) *TransactionOfWangYCItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) DepartureTime(departureTime string) *TransactionOfWangYCItemBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) MeetTime(meetTime string) *TransactionOfWangYCItemBuilder {
	builder.meetTime = meetTime
	builder.meetTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) BeginChargeTime(beginChargeTime string) *TransactionOfWangYCItemBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) FinishTime(finishTime string) *TransactionOfWangYCItemBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PayTime(payTime string) *TransactionOfWangYCItemBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) RefundTime(refundTime string) *TransactionOfWangYCItemBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) City(city string) *TransactionOfWangYCItemBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) DestinationCity(destinationCity string) *TransactionOfWangYCItemBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) NormalDistance(normalDistance float32) *TransactionOfWangYCItemBuilder {
	builder.normalDistance = normalDistance
	builder.normalDistanceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) StartAddress(startAddress string) *TransactionOfWangYCItemBuilder {
	builder.startAddress = startAddress
	builder.startAddressSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) StartName(startName string) *TransactionOfWangYCItemBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) EndAddress(endAddress string) *TransactionOfWangYCItemBuilder {
	builder.endAddress = endAddress
	builder.endAddressSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) EndName(endName string) *TransactionOfWangYCItemBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ActualStartName(actualStartName string) *TransactionOfWangYCItemBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ActualEndName(actualEndName string) *TransactionOfWangYCItemBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) RuleName(ruleName string) *TransactionOfWangYCItemBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarSrv(useCarSrv string) *TransactionOfWangYCItemBuilder {
	builder.useCarSrv = useCarSrv
	builder.useCarSrvSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) BudgetCenterName(budgetCenterName string) *TransactionOfWangYCItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) BudgetCenterId(budgetCenterId string) *TransactionOfWangYCItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarRemark(useCarRemark string) *TransactionOfWangYCItemBuilder {
	builder.useCarRemark = useCarRemark
	builder.useCarRemarkSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OrderAdditionalRemark(orderAdditionalRemark string) *TransactionOfWangYCItemBuilder {
	builder.orderAdditionalRemark = orderAdditionalRemark
	builder.orderAdditionalRemarkSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ApprovalReason(approvalReason string) *TransactionOfWangYCItemBuilder {
	builder.approvalReason = approvalReason
	builder.approvalReasonSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ApprovalCreateTime(approvalCreateTime string) *TransactionOfWangYCItemBuilder {
	builder.approvalCreateTime = approvalCreateTime
	builder.approvalCreateTimeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ApprovalLogs(approvalLogs string) *TransactionOfWangYCItemBuilder {
	builder.approvalLogs = approvalLogs
	builder.approvalLogsSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ApprovalId(approvalId string) *TransactionOfWangYCItemBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UpgradeType(upgradeType string) *TransactionOfWangYCItemBuilder {
	builder.upgradeType = upgradeType
	builder.upgradeTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsFixedPrice(isFixedPrice string) *TransactionOfWangYCItemBuilder {
	builder.isFixedPrice = isFixedPrice
	builder.isFixedPriceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsSelfDrive(isSelfDrive string) *TransactionOfWangYCItemBuilder {
	builder.isSelfDrive = isSelfDrive
	builder.isSelfDriveSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsDiscountOrder(isDiscountOrder string) *TransactionOfWangYCItemBuilder {
	builder.isDiscountOrder = isDiscountOrder
	builder.isDiscountOrderSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsCarpool(isCarpool string) *TransactionOfWangYCItemBuilder {
	builder.isCarpool = isCarpool
	builder.isCarpoolSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsReassignment(isReassignment string) *TransactionOfWangYCItemBuilder {
	builder.isReassignment = isReassignment
	builder.isReassignmentSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CallbackInfo(callbackInfo string) *TransactionOfWangYCItemBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) AfterApprovalResult(afterApprovalResult string) *TransactionOfWangYCItemBuilder {
	builder.afterApprovalResult = afterApprovalResult
	builder.afterApprovalResultSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) AfterApprovalContent(afterApprovalContent string) *TransactionOfWangYCItemBuilder {
	builder.afterApprovalContent = afterApprovalContent
	builder.afterApprovalContentSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ApprovalExtraInfo(approvalExtraInfo string) *TransactionOfWangYCItemBuilder {
	builder.approvalExtraInfo = approvalExtraInfo
	builder.approvalExtraInfoSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarTypeName(useCarTypeName string) *TransactionOfWangYCItemBuilder {
	builder.useCarTypeName = useCarTypeName
	builder.useCarTypeNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarTypeV2(useCarTypeV2 string) *TransactionOfWangYCItemBuilder {
	builder.useCarTypeV2 = useCarTypeV2
	builder.useCarTypeV2Set = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UserPayType(userPayType string) *TransactionOfWangYCItemBuilder {
	builder.userPayType = userPayType
	builder.userPayTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) SubUseCarSrv(subUseCarSrv string) *TransactionOfWangYCItemBuilder {
	builder.subUseCarSrv = subUseCarSrv
	builder.subUseCarSrvSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CompanyId(companyId int64) *TransactionOfWangYCItemBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) CompanyName(companyName string) *TransactionOfWangYCItemBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) SpsPickUpServiceFee(spsPickUpServiceFee float32) *TransactionOfWangYCItemBuilder {
	builder.spsPickUpServiceFee = spsPickUpServiceFee
	builder.spsPickUpServiceFeeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) IsUnusual(isUnusual string) *TransactionOfWangYCItemBuilder {
	builder.isUnusual = isUnusual
	builder.isUnusualSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UnusualType(unusualType string) *TransactionOfWangYCItemBuilder {
	builder.unusualType = unusualType
	builder.unusualTypeSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UnusualContent(unusualContent string) *TransactionOfWangYCItemBuilder {
	builder.unusualContent = unusualContent
	builder.unusualContentSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PositionName(positionName string) *TransactionOfWangYCItemBuilder {
	builder.positionName = positionName
	builder.positionNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) InstitutionName(institutionName string) *TransactionOfWangYCItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) UseCarService(useCarService string) *TransactionOfWangYCItemBuilder {
	builder.useCarService = useCarService
	builder.useCarServiceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ExInfo01(exInfo01 string) *TransactionOfWangYCItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ExInfo02(exInfo02 string) *TransactionOfWangYCItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ExInfo03(exInfo03 string) *TransactionOfWangYCItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) LegalEntityId(legalEntityId int64) *TransactionOfWangYCItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) LegalEntityName(legalEntityName string) *TransactionOfWangYCItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) BeforeApprovalContent(beforeApprovalContent string) *TransactionOfWangYCItemBuilder {
	builder.beforeApprovalContent = beforeApprovalContent
	builder.beforeApprovalContentSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OriginalPrice(originalPrice float32) *TransactionOfWangYCItemBuilder {
	builder.originalPrice = originalPrice
	builder.originalPriceSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) OneTimeOfferSubsidy(oneTimeOfferSubsidy float32) *TransactionOfWangYCItemBuilder {
	builder.oneTimeOfferSubsidy = oneTimeOfferSubsidy
	builder.oneTimeOfferSubsidySet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) SubsidyAmount(subsidyAmount float32) *TransactionOfWangYCItemBuilder {
	builder.subsidyAmount = subsidyAmount
	builder.subsidyAmountSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) VoucherDeductionTypeName(voucherDeductionTypeName string) *TransactionOfWangYCItemBuilder {
	builder.voucherDeductionTypeName = voucherDeductionTypeName
	builder.voucherDeductionTypeNameSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) ResidentCityNames(residentCityNames string) *TransactionOfWangYCItemBuilder {
	builder.residentCityNames = residentCityNames
	builder.residentCityNamesSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *TransactionOfWangYCItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *TransactionOfWangYCItemBuilder) PassengerMemberId(passengerMemberId string) *TransactionOfWangYCItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}

func (builder *TransactionOfWangYCItemBuilder) Build() *TransactionOfWangYCItem {
	data := &TransactionOfWangYCItem{}
	if builder.transactionIdSet {
		data.TransactionId = &builder.transactionId
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.outMoneySet {
		data.OutMoney = &builder.outMoney
	}
	if builder.inMoneySet {
		data.InMoney = &builder.inMoney
	}
	if builder.createTimeDateSet {
		data.CreateTimeDate = &builder.createTimeDate
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.isCurrentTermSet {
		data.IsCurrentTerm = &builder.isCurrentTerm
	}
	if builder.companyOrderIdSet {
		data.CompanyOrderId = &builder.companyOrderId
	}
	if builder.outOrderIdSet {
		data.OutOrderId = &builder.outOrderId
	}
	if builder.orderSourceSet {
		data.OrderSource = &builder.orderSource
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.useCarTypeSet {
		data.UseCarType = &builder.useCarType
	}
	if builder.requireLevelSet {
		data.RequireLevel = &builder.requireLevel
	}
	if builder.memberNameSet {
		data.MemberName = &builder.memberName
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.memberEmailSet {
		data.MemberEmail = &builder.memberEmail
	}
	if builder.callPhoneSet {
		data.CallPhone = &builder.callPhone
	}
	if builder.departmentNameSet {
		data.DepartmentName = &builder.departmentName
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.branchNameSet {
		data.BranchName = &builder.branchName
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.passengerPhoneSet {
		data.PassengerPhone = &builder.passengerPhone
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.totalPriceSet {
		data.TotalPrice = &builder.totalPrice
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.voucherPaySet {
		data.VoucherPay = &builder.voucherPay
	}
	if builder.costSet {
		data.Cost = &builder.cost
	}
	if builder.discountAfterPriceSet {
		data.DiscountAfterPrice = &builder.discountAfterPrice
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.orderTypeSet {
		data.OrderType = &builder.orderType
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
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
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.destinationCitySet {
		data.DestinationCity = &builder.destinationCity
	}
	if builder.normalDistanceSet {
		data.NormalDistance = &builder.normalDistance
	}
	if builder.startAddressSet {
		data.StartAddress = &builder.startAddress
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.endAddressSet {
		data.EndAddress = &builder.endAddress
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
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.useCarSrvSet {
		data.UseCarSrv = &builder.useCarSrv
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.useCarRemarkSet {
		data.UseCarRemark = &builder.useCarRemark
	}
	if builder.orderAdditionalRemarkSet {
		data.OrderAdditionalRemark = &builder.orderAdditionalRemark
	}
	if builder.approvalReasonSet {
		data.ApprovalReason = &builder.approvalReason
	}
	if builder.approvalCreateTimeSet {
		data.ApprovalCreateTime = &builder.approvalCreateTime
	}
	if builder.approvalLogsSet {
		data.ApprovalLogs = &builder.approvalLogs
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.upgradeTypeSet {
		data.UpgradeType = &builder.upgradeType
	}
	if builder.isFixedPriceSet {
		data.IsFixedPrice = &builder.isFixedPrice
	}
	if builder.isSelfDriveSet {
		data.IsSelfDrive = &builder.isSelfDrive
	}
	if builder.isDiscountOrderSet {
		data.IsDiscountOrder = &builder.isDiscountOrder
	}
	if builder.isCarpoolSet {
		data.IsCarpool = &builder.isCarpool
	}
	if builder.isReassignmentSet {
		data.IsReassignment = &builder.isReassignment
	}
	if builder.callbackInfoSet {
		data.CallbackInfo = &builder.callbackInfo
	}
	if builder.afterApprovalResultSet {
		data.AfterApprovalResult = &builder.afterApprovalResult
	}
	if builder.afterApprovalContentSet {
		data.AfterApprovalContent = &builder.afterApprovalContent
	}
	if builder.approvalExtraInfoSet {
		data.ApprovalExtraInfo = &builder.approvalExtraInfo
	}
	if builder.useCarTypeNameSet {
		data.UseCarTypeName = &builder.useCarTypeName
	}
	if builder.useCarTypeV2Set {
		data.UseCarTypeV2 = &builder.useCarTypeV2
	}
	if builder.userPayTypeSet {
		data.UserPayType = &builder.userPayType
	}
	if builder.subUseCarSrvSet {
		data.SubUseCarSrv = &builder.subUseCarSrv
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.companyNameSet {
		data.CompanyName = &builder.companyName
	}
	if builder.spsPickUpServiceFeeSet {
		data.SpsPickUpServiceFee = &builder.spsPickUpServiceFee
	}
	if builder.isUnusualSet {
		data.IsUnusual = &builder.isUnusual
	}
	if builder.unusualTypeSet {
		data.UnusualType = &builder.unusualType
	}
	if builder.unusualContentSet {
		data.UnusualContent = &builder.unusualContent
	}
	if builder.positionNameSet {
		data.PositionName = &builder.positionName
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.useCarServiceSet {
		data.UseCarService = &builder.useCarService
	}
	if builder.exInfo01Set {
		data.ExInfo01 = &builder.exInfo01
	}
	if builder.exInfo02Set {
		data.ExInfo02 = &builder.exInfo02
	}
	if builder.exInfo03Set {
		data.ExInfo03 = &builder.exInfo03
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.beforeApprovalContentSet {
		data.BeforeApprovalContent = &builder.beforeApprovalContent
	}
	if builder.originalPriceSet {
		data.OriginalPrice = &builder.originalPrice
	}
	if builder.oneTimeOfferSubsidySet {
		data.OneTimeOfferSubsidy = &builder.oneTimeOfferSubsidy
	}
	if builder.subsidyAmountSet {
		data.SubsidyAmount = &builder.subsidyAmount
	}
	if builder.voucherDeductionTypeNameSet {
		data.VoucherDeductionTypeName = &builder.voucherDeductionTypeName
	}
	if builder.residentCityNamesSet {
		data.ResidentCityNames = &builder.residentCityNames
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
	}
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	return data
}
