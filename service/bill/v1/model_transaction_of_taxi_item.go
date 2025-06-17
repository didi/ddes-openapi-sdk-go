package v1

// TransactionOfTaxiItem 账单交易明细(出租车)
type TransactionOfTaxiItem struct {
	TransactionId         *int64                  `json:"transaction_id,omitempty"`          // 资金流水id
	Type                  *string                 `json:"type,omitempty"`                    // 资金操作类型
	TransactionList       []TransactionOfTaxiItem `json:"transaction_list,omitempty"`        // 交易列表
	OutMoney              *float32                `json:"out_money,omitempty"`               // 操作金额 - 支出
	InMoney               *float32                `json:"in_money,omitempty"`                // 操作金额 - 收入
	CreateTimeDate        *string                 `json:"create_time_date,omitempty"`        // 创建时间
	OrderId               *int64                  `json:"order_id,omitempty"`                // 订单号
	Remark                *string                 `json:"remark,omitempty"`                  // 备注
	IsCurrentTerm         *string                 `json:"is_current_term,omitempty"`         // 所属账期
	CompanyOrderId        *int64                  `json:"company_order_id,omitempty"`        // 企业订单号
	OutOrderId            *int64                  `json:"out_order_id,omitempty"`            // 专快订单号
	OrderSource           *string                 `json:"order_source,omitempty"`            // 订单来源
	Status                *string                 `json:"status,omitempty"`                  // 订单状态
	UseCarType            *string                 `json:"use_car_type,omitempty"`            // 用车类型
	RequireLevel          *string                 `json:"require_level,omitempty"`           // 用车类型(明细)
	MemberName            *string                 `json:"member_name,omitempty"`             // 下单人姓名
	EmployeeNumber        *string                 `json:"employee_number,omitempty"`         // 下单人员工ID
	MemberEmail           *string                 `json:"member_email,omitempty"`            // 企业邮箱
	CallPhone             *string                 `json:"call_phone,omitempty"`              // 下单人电话
	DepartmentName        *string                 `json:"department_name,omitempty"`         // 所在部门(新)
	OutBudgetId           *string                 `json:"out_budget_id,omitempty"`           // 所在部门(新)编号
	BranchName            *string                 `json:"branch_name,omitempty"`             // 下单人分公司(旧)
	Department            *string                 `json:"department,omitempty"`              // 下单人部门(旧)
	PassengerName         *string                 `json:"passenger_name,omitempty"`          // 乘车人姓名
	PassengerPhone        *string                 `json:"passenger_phone,omitempty"`         // 乘车人电话
	PayType               *string                 `json:"pay_type,omitempty"`                // 支付方式
	TotalPrice            *float32                `json:"total_price,omitempty"`             // 订单总金额
	CompanyRealPay        *float32                `json:"company_real_pay,omitempty"`        // 企业实付金额
	PersonalRealPay       *float32                `json:"personal_real_pay,omitempty"`       // 个人实付金额
	VoucherPay            *float32                `json:"voucher_pay,omitempty"`             // 用券抵扣金额
	Cost                  *float32                `json:"cost,omitempty"`                    // 车费金额
	DiscountAfterPrice    *float32                `json:"discount_after_price,omitempty"`    // 平台使用费金额
	CompanyRealRefund     *float32                `json:"company_real_refund,omitempty"`     // 企业实退金额
	OrderType             *string                 `json:"order_type,omitempty"`              // 订单类型
	CreateTime            *string                 `json:"create_time,omitempty"`             // 下单时间
	DepartureTime         *string                 `json:"departure_time,omitempty"`          // 计划用车时间
	MeetTime              *string                 `json:"meet_time,omitempty"`               // 司机接单时间
	BeginChargeTime       *string                 `json:"begin_charge_time,omitempty"`       // 开始计费时间
	FinishTime            *string                 `json:"finish_time,omitempty"`             // 结束计费时间
	PayTime               *string                 `json:"pay_time,omitempty"`                // 支付时间
	RefundTime            *string                 `json:"refund_time,omitempty"`             // 退款时间
	City                  *string                 `json:"city,omitempty"`                    // 用车城市
	DestinationCity       *string                 `json:"destination_city,omitempty"`        // 到达城市
	NormalDistance        *float32                `json:"normal_distance,omitempty"`         // 用车行驶距离
	StartAddress          *string                 `json:"start_address,omitempty"`           // 出发地地址
	StartName             *string                 `json:"start_name,omitempty"`              // 出发地名称
	EndAddress            *string                 `json:"end_address,omitempty"`             // 目的地地址
	EndName               *string                 `json:"end_name,omitempty"`                // 目的地名称
	ActualStartName       *string                 `json:"actual_start_name,omitempty"`       // 实际出发地
	ActualEndName         *string                 `json:"actual_end_name,omitempty"`         // 实际目的地
	RuleName              *string                 `json:"rule_name,omitempty"`               // 用车权限
	UseCarSrv             *string                 `json:"use_car_srv,omitempty"`             // 用车方式
	BudgetCenterName      *string                 `json:"budget_center_name,omitempty"`      // 成本中心名称
	BudgetCenterId        *int64                  `json:"budget_center_id,omitempty"`        // 成本中心id
	UseCarRemark          *string                 `json:"use_car_remark,omitempty"`          // 用车备注
	OrderAdditionalRemark *string                 `json:"order_additional_remark,omitempty"` // 补充说明
	ApprovalReason        *string                 `json:"approval_reason,omitempty"`         // 申请备注
	ApprovalCreateTime    *string                 `json:"approval_create_time,omitempty"`    // 申请提交时间
	ApprovalLogs          *string                 `json:"approval_logs,omitempty"`           // 审批人1|审批时间
	ApprovalId            *int64                  `json:"approval_id,omitempty"`             // 外部审批单ID
	UpgradeType           *string                 `json:"upgrade_type,omitempty"`            // 是否自费升舱
	IsFixedPrice          *string                 `json:"is_fixed_price,omitempty"`          // 是否一口价
	IsSelfDrive           *string                 `json:"is_self_drive,omitempty"`           // 是否自驾同行
	IsDiscountOrder       *string                 `json:"is_discount_order,omitempty"`       // 是否折扣订单
	IsCarpool             *string                 `json:"is_carpool,omitempty"`              // 是否拼车
	IsReassignment        *string                 `json:"is_reassignment,omitempty"`         // 是否改派
	CallbackInfo          *string                 `json:"callback_info,omitempty"`           // API企业自定义
	AfterApprovalResult   *string                 `json:"after_approval_result,omitempty"`   // 行后审批结果
	AfterApprovalContent  *string                 `json:"after_approval_content,omitempty"`  // 审批历史
	ApprovalExtraInfo     *string                 `json:"approval_extra_info,omitempty"`     // 审批单自定义
	InstitutionId         *int64                  `json:"institution_id,omitempty"`          // 用车制度
}

type TransactionOfTaxiItemBuilder struct {
	transactionId            int64 // 资金流水id
	transactionIdSet         bool
	type_                    string // 资金操作类型
	type_Set                 bool
	transactionList          []TransactionOfTaxiItem // 交易列表
	transactionListSet       bool
	outMoney                 float32 // 操作金额 - 支出
	outMoneySet              bool
	inMoney                  float32 // 操作金额 - 收入
	inMoneySet               bool
	createTimeDate           string // 创建时间
	createTimeDateSet        bool
	orderId                  int64 // 订单号
	orderIdSet               bool
	remark                   string // 备注
	remarkSet                bool
	isCurrentTerm            string // 所属账期
	isCurrentTermSet         bool
	companyOrderId           int64 // 企业订单号
	companyOrderIdSet        bool
	outOrderId               int64 // 专快订单号
	outOrderIdSet            bool
	orderSource              string // 订单来源
	orderSourceSet           bool
	status                   string // 订单状态
	statusSet                bool
	useCarType               string // 用车类型
	useCarTypeSet            bool
	requireLevel             string // 用车类型(明细)
	requireLevelSet          bool
	memberName               string // 下单人姓名
	memberNameSet            bool
	employeeNumber           string // 下单人员工ID
	employeeNumberSet        bool
	memberEmail              string // 企业邮箱
	memberEmailSet           bool
	callPhone                string // 下单人电话
	callPhoneSet             bool
	departmentName           string // 所在部门(新)
	departmentNameSet        bool
	outBudgetId              string // 所在部门(新)编号
	outBudgetIdSet           bool
	branchName               string // 下单人分公司(旧)
	branchNameSet            bool
	department               string // 下单人部门(旧)
	departmentSet            bool
	passengerName            string // 乘车人姓名
	passengerNameSet         bool
	passengerPhone           string // 乘车人电话
	passengerPhoneSet        bool
	payType                  string // 支付方式
	payTypeSet               bool
	totalPrice               float32 // 订单总金额
	totalPriceSet            bool
	companyRealPay           float32 // 企业实付金额
	companyRealPaySet        bool
	personalRealPay          float32 // 个人实付金额
	personalRealPaySet       bool
	voucherPay               float32 // 用券抵扣金额
	voucherPaySet            bool
	cost                     float32 // 车费金额
	costSet                  bool
	discountAfterPrice       float32 // 平台使用费金额
	discountAfterPriceSet    bool
	companyRealRefund        float32 // 企业实退金额
	companyRealRefundSet     bool
	orderType                string // 订单类型
	orderTypeSet             bool
	createTime               string // 下单时间
	createTimeSet            bool
	departureTime            string // 计划用车时间
	departureTimeSet         bool
	meetTime                 string // 司机接单时间
	meetTimeSet              bool
	beginChargeTime          string // 开始计费时间
	beginChargeTimeSet       bool
	finishTime               string // 结束计费时间
	finishTimeSet            bool
	payTime                  string // 支付时间
	payTimeSet               bool
	refundTime               string // 退款时间
	refundTimeSet            bool
	city                     string // 用车城市
	citySet                  bool
	destinationCity          string // 到达城市
	destinationCitySet       bool
	normalDistance           float32 // 用车行驶距离
	normalDistanceSet        bool
	startAddress             string // 出发地地址
	startAddressSet          bool
	startName                string // 出发地名称
	startNameSet             bool
	endAddress               string // 目的地地址
	endAddressSet            bool
	endName                  string // 目的地名称
	endNameSet               bool
	actualStartName          string // 实际出发地
	actualStartNameSet       bool
	actualEndName            string // 实际目的地
	actualEndNameSet         bool
	ruleName                 string // 用车权限
	ruleNameSet              bool
	useCarSrv                string // 用车方式
	useCarSrvSet             bool
	budgetCenterName         string // 成本中心名称
	budgetCenterNameSet      bool
	budgetCenterId           int64 // 成本中心id
	budgetCenterIdSet        bool
	useCarRemark             string // 用车备注
	useCarRemarkSet          bool
	orderAdditionalRemark    string // 补充说明
	orderAdditionalRemarkSet bool
	approvalReason           string // 申请备注
	approvalReasonSet        bool
	approvalCreateTime       string // 申请提交时间
	approvalCreateTimeSet    bool
	approvalLogs             string // 审批人1|审批时间
	approvalLogsSet          bool
	approvalId               int64 // 外部审批单ID
	approvalIdSet            bool
	upgradeType              string // 是否自费升舱
	upgradeTypeSet           bool
	isFixedPrice             string // 是否一口价
	isFixedPriceSet          bool
	isSelfDrive              string // 是否自驾同行
	isSelfDriveSet           bool
	isDiscountOrder          string // 是否折扣订单
	isDiscountOrderSet       bool
	isCarpool                string // 是否拼车
	isCarpoolSet             bool
	isReassignment           string // 是否改派
	isReassignmentSet        bool
	callbackInfo             string // API企业自定义
	callbackInfoSet          bool
	afterApprovalResult      string // 行后审批结果
	afterApprovalResultSet   bool
	afterApprovalContent     string // 审批历史
	afterApprovalContentSet  bool
	approvalExtraInfo        string // 审批单自定义
	approvalExtraInfoSet     bool
	institutionId            int64 // 用车制度
	institutionIdSet         bool
}

func NewTransactionOfTaxiItemBuilder() *TransactionOfTaxiItemBuilder {
	return &TransactionOfTaxiItemBuilder{}
}
func (builder *TransactionOfTaxiItemBuilder) TransactionId(transactionId int64) *TransactionOfTaxiItemBuilder {
	builder.transactionId = transactionId
	builder.transactionIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) Type(type_ string) *TransactionOfTaxiItemBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) TransactionList(transactionList []TransactionOfTaxiItem) *TransactionOfTaxiItemBuilder {
	builder.transactionList = transactionList
	builder.transactionListSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OutMoney(outMoney float32) *TransactionOfTaxiItemBuilder {
	builder.outMoney = outMoney
	builder.outMoneySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) InMoney(inMoney float32) *TransactionOfTaxiItemBuilder {
	builder.inMoney = inMoney
	builder.inMoneySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CreateTimeDate(createTimeDate string) *TransactionOfTaxiItemBuilder {
	builder.createTimeDate = createTimeDate
	builder.createTimeDateSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OrderId(orderId int64) *TransactionOfTaxiItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) Remark(remark string) *TransactionOfTaxiItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsCurrentTerm(isCurrentTerm string) *TransactionOfTaxiItemBuilder {
	builder.isCurrentTerm = isCurrentTerm
	builder.isCurrentTermSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CompanyOrderId(companyOrderId int64) *TransactionOfTaxiItemBuilder {
	builder.companyOrderId = companyOrderId
	builder.companyOrderIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OutOrderId(outOrderId int64) *TransactionOfTaxiItemBuilder {
	builder.outOrderId = outOrderId
	builder.outOrderIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OrderSource(orderSource string) *TransactionOfTaxiItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) Status(status string) *TransactionOfTaxiItemBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) UseCarType(useCarType string) *TransactionOfTaxiItemBuilder {
	builder.useCarType = useCarType
	builder.useCarTypeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) RequireLevel(requireLevel string) *TransactionOfTaxiItemBuilder {
	builder.requireLevel = requireLevel
	builder.requireLevelSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) MemberName(memberName string) *TransactionOfTaxiItemBuilder {
	builder.memberName = memberName
	builder.memberNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) EmployeeNumber(employeeNumber string) *TransactionOfTaxiItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) MemberEmail(memberEmail string) *TransactionOfTaxiItemBuilder {
	builder.memberEmail = memberEmail
	builder.memberEmailSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CallPhone(callPhone string) *TransactionOfTaxiItemBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) DepartmentName(departmentName string) *TransactionOfTaxiItemBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OutBudgetId(outBudgetId string) *TransactionOfTaxiItemBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) BranchName(branchName string) *TransactionOfTaxiItemBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) Department(department string) *TransactionOfTaxiItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) PassengerName(passengerName string) *TransactionOfTaxiItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) PassengerPhone(passengerPhone string) *TransactionOfTaxiItemBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) PayType(payType string) *TransactionOfTaxiItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) TotalPrice(totalPrice float32) *TransactionOfTaxiItemBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CompanyRealPay(companyRealPay float32) *TransactionOfTaxiItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) PersonalRealPay(personalRealPay float32) *TransactionOfTaxiItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) VoucherPay(voucherPay float32) *TransactionOfTaxiItemBuilder {
	builder.voucherPay = voucherPay
	builder.voucherPaySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) Cost(cost float32) *TransactionOfTaxiItemBuilder {
	builder.cost = cost
	builder.costSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) DiscountAfterPrice(discountAfterPrice float32) *TransactionOfTaxiItemBuilder {
	builder.discountAfterPrice = discountAfterPrice
	builder.discountAfterPriceSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CompanyRealRefund(companyRealRefund float32) *TransactionOfTaxiItemBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OrderType(orderType string) *TransactionOfTaxiItemBuilder {
	builder.orderType = orderType
	builder.orderTypeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CreateTime(createTime string) *TransactionOfTaxiItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) DepartureTime(departureTime string) *TransactionOfTaxiItemBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) MeetTime(meetTime string) *TransactionOfTaxiItemBuilder {
	builder.meetTime = meetTime
	builder.meetTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) BeginChargeTime(beginChargeTime string) *TransactionOfTaxiItemBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) FinishTime(finishTime string) *TransactionOfTaxiItemBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) PayTime(payTime string) *TransactionOfTaxiItemBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) RefundTime(refundTime string) *TransactionOfTaxiItemBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) City(city string) *TransactionOfTaxiItemBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) DestinationCity(destinationCity string) *TransactionOfTaxiItemBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) NormalDistance(normalDistance float32) *TransactionOfTaxiItemBuilder {
	builder.normalDistance = normalDistance
	builder.normalDistanceSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) StartAddress(startAddress string) *TransactionOfTaxiItemBuilder {
	builder.startAddress = startAddress
	builder.startAddressSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) StartName(startName string) *TransactionOfTaxiItemBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) EndAddress(endAddress string) *TransactionOfTaxiItemBuilder {
	builder.endAddress = endAddress
	builder.endAddressSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) EndName(endName string) *TransactionOfTaxiItemBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ActualStartName(actualStartName string) *TransactionOfTaxiItemBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ActualEndName(actualEndName string) *TransactionOfTaxiItemBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) RuleName(ruleName string) *TransactionOfTaxiItemBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) UseCarSrv(useCarSrv string) *TransactionOfTaxiItemBuilder {
	builder.useCarSrv = useCarSrv
	builder.useCarSrvSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) BudgetCenterName(budgetCenterName string) *TransactionOfTaxiItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) BudgetCenterId(budgetCenterId int64) *TransactionOfTaxiItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) UseCarRemark(useCarRemark string) *TransactionOfTaxiItemBuilder {
	builder.useCarRemark = useCarRemark
	builder.useCarRemarkSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) OrderAdditionalRemark(orderAdditionalRemark string) *TransactionOfTaxiItemBuilder {
	builder.orderAdditionalRemark = orderAdditionalRemark
	builder.orderAdditionalRemarkSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ApprovalReason(approvalReason string) *TransactionOfTaxiItemBuilder {
	builder.approvalReason = approvalReason
	builder.approvalReasonSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ApprovalCreateTime(approvalCreateTime string) *TransactionOfTaxiItemBuilder {
	builder.approvalCreateTime = approvalCreateTime
	builder.approvalCreateTimeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ApprovalLogs(approvalLogs string) *TransactionOfTaxiItemBuilder {
	builder.approvalLogs = approvalLogs
	builder.approvalLogsSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ApprovalId(approvalId int64) *TransactionOfTaxiItemBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) UpgradeType(upgradeType string) *TransactionOfTaxiItemBuilder {
	builder.upgradeType = upgradeType
	builder.upgradeTypeSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsFixedPrice(isFixedPrice string) *TransactionOfTaxiItemBuilder {
	builder.isFixedPrice = isFixedPrice
	builder.isFixedPriceSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsSelfDrive(isSelfDrive string) *TransactionOfTaxiItemBuilder {
	builder.isSelfDrive = isSelfDrive
	builder.isSelfDriveSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsDiscountOrder(isDiscountOrder string) *TransactionOfTaxiItemBuilder {
	builder.isDiscountOrder = isDiscountOrder
	builder.isDiscountOrderSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsCarpool(isCarpool string) *TransactionOfTaxiItemBuilder {
	builder.isCarpool = isCarpool
	builder.isCarpoolSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) IsReassignment(isReassignment string) *TransactionOfTaxiItemBuilder {
	builder.isReassignment = isReassignment
	builder.isReassignmentSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) CallbackInfo(callbackInfo string) *TransactionOfTaxiItemBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) AfterApprovalResult(afterApprovalResult string) *TransactionOfTaxiItemBuilder {
	builder.afterApprovalResult = afterApprovalResult
	builder.afterApprovalResultSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) AfterApprovalContent(afterApprovalContent string) *TransactionOfTaxiItemBuilder {
	builder.afterApprovalContent = afterApprovalContent
	builder.afterApprovalContentSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) ApprovalExtraInfo(approvalExtraInfo string) *TransactionOfTaxiItemBuilder {
	builder.approvalExtraInfo = approvalExtraInfo
	builder.approvalExtraInfoSet = true
	return builder
}
func (builder *TransactionOfTaxiItemBuilder) InstitutionId(institutionId int64) *TransactionOfTaxiItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}

func (builder *TransactionOfTaxiItemBuilder) Build() *TransactionOfTaxiItem {
	data := &TransactionOfTaxiItem{}
	if builder.transactionIdSet {
		data.TransactionId = &builder.transactionId
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.transactionListSet {
		data.TransactionList = builder.transactionList
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
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	return data
}
