package v1

// TrainOrderInfo 订单基础信息
type TrainOrderInfo struct {
	OrderId          *string            `json:"order_id,omitempty"`           // 订单号，主订单号
	IsGrabOrder      *string            `json:"is_grab_order,omitempty"`      // 是否抢票订单，枚举值数字 0：否 1：是
	CompanyName      *string            `json:"company_name,omitempty"`       // 企业名称，在滴滴注册的企业名称
	CompanyId        *string            `json:"company_id,omitempty"`         // 企业ID，滴滴企业ID
	BookerName       *string            `json:"booker_name,omitempty"`        // 预订人姓名，下单人
	BookerEmployeeId *string            `json:"booker_employee_id,omitempty"` // 预订人员工编号，客户维护或者外部对接的数据
	BookerMemberId   *string            `json:"booker_member_id,omitempty"`   // 预订人滴滴ID，内部员工
	BookerPhone      *string            `json:"booker_phone,omitempty"`       // 预订人手机号，脱敏展示
	DepartmentName   *string            `json:"department_name,omitempty"`    // 预订人部门名称，部门链路
	DepartmentId     *string            `json:"department_id,omitempty"`      // 预订人部门ID，末级部门ID
	InstitutionName  *string            `json:"institution_name,omitempty"`   // 制度名称，滴滴制度名称
	InstitutionId    *string            `json:"institution_id,omitempty"`     // 制度ID
	ApprovalId       *string            `json:"approval_id,omitempty"`        // 差旅申请单ID，滴滴申请单ID
	OutApprovalId    *string            `json:"out_approval_id,omitempty"`    // 外部申请单ID，系统对接外部的申请单号
	OrderStatus      *string            `json:"order_status,omitempty"`       // 订单状态，枚举英文code：Holding 占座中 Tobepaid 待支付 Waitforapproval 待审批 Cancelled 已取消 Closed 订单关闭 Ticketing 待出票 Ticketed 已出票 refunding 退票中 Partialrefund 部分退票 Refunded 已退票 Change Holding 改签占座中 Change to be paid 改签待支付 Changing 改签中 Completed 已完成 Unknow 未知 Grabbing 抢票中 Grab error 抢票异常 Canceling 取消中 Cancel error 取消异常 Grab failed 抢票失败
	CreateTime       *string            `json:"create_time,omitempty"`        // 订单创建时间，格式：yyyy-MM-dd HH:mm:ss
	IsPay            *string            `json:"is_pay,omitempty"`             // 支付状态，枚举数字：0 未支付 1 已支付
	PayType          *string            `json:"pay_type,omitempty"`           // 订单支付类型，枚举数字：0 企业支付 1 个人支付 2 混合支付
	ChangeTime       *string            `json:"change_time,omitempty"`        // 订单更新时间，格式：yyyy-MM-dd HH:mm:ss 查询条件
	CancelTime       *string            `json:"cancel_time,omitempty"`        // 取消时间，格式：yyyy-MM-dd HH:mm:ss
	BudgetCenterList []BudgetCenterList `json:"budget_center_list,omitempty"` // 多成本中心，参考budget_center_list对象
	DepartmentCode   *string            `json:"department_code,omitempty"`    // * 接口返回，文档未标注字段
}

type TrainOrderInfoBuilder struct {
	orderId             string // 订单号，主订单号
	orderIdSet          bool
	isGrabOrder         string // 是否抢票订单，枚举值数字 0：否 1：是
	isGrabOrderSet      bool
	companyName         string // 企业名称，在滴滴注册的企业名称
	companyNameSet      bool
	companyId           string // 企业ID，滴滴企业ID
	companyIdSet        bool
	bookerName          string // 预订人姓名，下单人
	bookerNameSet       bool
	bookerEmployeeId    string // 预订人员工编号，客户维护或者外部对接的数据
	bookerEmployeeIdSet bool
	bookerMemberId      string // 预订人滴滴ID，内部员工
	bookerMemberIdSet   bool
	bookerPhone         string // 预订人手机号，脱敏展示
	bookerPhoneSet      bool
	departmentName      string // 预订人部门名称，部门链路
	departmentNameSet   bool
	departmentId        string // 预订人部门ID，末级部门ID
	departmentIdSet     bool
	institutionName     string // 制度名称，滴滴制度名称
	institutionNameSet  bool
	institutionId       string // 制度ID
	institutionIdSet    bool
	approvalId          string // 差旅申请单ID，滴滴申请单ID
	approvalIdSet       bool
	outApprovalId       string // 外部申请单ID，系统对接外部的申请单号
	outApprovalIdSet    bool
	orderStatus         string // 订单状态，枚举英文code：Holding 占座中 Tobepaid 待支付 Waitforapproval 待审批 Cancelled 已取消 Closed 订单关闭 Ticketing 待出票 Ticketed 已出票 refunding 退票中 Partialrefund 部分退票 Refunded 已退票 Change Holding 改签占座中 Change to be paid 改签待支付 Changing 改签中 Completed 已完成 Unknow 未知 Grabbing 抢票中 Grab error 抢票异常 Canceling 取消中 Cancel error 取消异常 Grab failed 抢票失败
	orderStatusSet      bool
	createTime          string // 订单创建时间，格式：yyyy-MM-dd HH:mm:ss
	createTimeSet       bool
	isPay               string // 支付状态，枚举数字：0 未支付 1 已支付
	isPaySet            bool
	payType             string // 订单支付类型，枚举数字：0 企业支付 1 个人支付 2 混合支付
	payTypeSet          bool
	changeTime          string // 订单更新时间，格式：yyyy-MM-dd HH:mm:ss 查询条件
	changeTimeSet       bool
	cancelTime          string // 取消时间，格式：yyyy-MM-dd HH:mm:ss
	cancelTimeSet       bool
	budgetCenterList    []BudgetCenterList // 多成本中心，参考budget_center_list对象
	budgetCenterListSet bool
	departmentCode      string // * 接口返回，文档未标注字段
	departmentCodeSet   bool
}

func NewTrainOrderInfoBuilder() *TrainOrderInfoBuilder {
	return &TrainOrderInfoBuilder{}
}
func (builder *TrainOrderInfoBuilder) OrderId(orderId string) *TrainOrderInfoBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) IsGrabOrder(isGrabOrder string) *TrainOrderInfoBuilder {
	builder.isGrabOrder = isGrabOrder
	builder.isGrabOrderSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) CompanyName(companyName string) *TrainOrderInfoBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) CompanyId(companyId string) *TrainOrderInfoBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) BookerName(bookerName string) *TrainOrderInfoBuilder {
	builder.bookerName = bookerName
	builder.bookerNameSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) BookerEmployeeId(bookerEmployeeId string) *TrainOrderInfoBuilder {
	builder.bookerEmployeeId = bookerEmployeeId
	builder.bookerEmployeeIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) BookerMemberId(bookerMemberId string) *TrainOrderInfoBuilder {
	builder.bookerMemberId = bookerMemberId
	builder.bookerMemberIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) BookerPhone(bookerPhone string) *TrainOrderInfoBuilder {
	builder.bookerPhone = bookerPhone
	builder.bookerPhoneSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) DepartmentName(departmentName string) *TrainOrderInfoBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) DepartmentId(departmentId string) *TrainOrderInfoBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) InstitutionName(institutionName string) *TrainOrderInfoBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) InstitutionId(institutionId string) *TrainOrderInfoBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) ApprovalId(approvalId string) *TrainOrderInfoBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) OutApprovalId(outApprovalId string) *TrainOrderInfoBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) OrderStatus(orderStatus string) *TrainOrderInfoBuilder {
	builder.orderStatus = orderStatus
	builder.orderStatusSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) CreateTime(createTime string) *TrainOrderInfoBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) IsPay(isPay string) *TrainOrderInfoBuilder {
	builder.isPay = isPay
	builder.isPaySet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) PayType(payType string) *TrainOrderInfoBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) ChangeTime(changeTime string) *TrainOrderInfoBuilder {
	builder.changeTime = changeTime
	builder.changeTimeSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) CancelTime(cancelTime string) *TrainOrderInfoBuilder {
	builder.cancelTime = cancelTime
	builder.cancelTimeSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) BudgetCenterList(budgetCenterList []BudgetCenterList) *TrainOrderInfoBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *TrainOrderInfoBuilder) DepartmentCode(departmentCode string) *TrainOrderInfoBuilder {
	builder.departmentCode = departmentCode
	builder.departmentCodeSet = true
	return builder
}

func (builder *TrainOrderInfoBuilder) Build() *TrainOrderInfo {
	data := &TrainOrderInfo{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.isGrabOrderSet {
		data.IsGrabOrder = &builder.isGrabOrder
	}
	if builder.companyNameSet {
		data.CompanyName = &builder.companyName
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.bookerNameSet {
		data.BookerName = &builder.bookerName
	}
	if builder.bookerEmployeeIdSet {
		data.BookerEmployeeId = &builder.bookerEmployeeId
	}
	if builder.bookerMemberIdSet {
		data.BookerMemberId = &builder.bookerMemberId
	}
	if builder.bookerPhoneSet {
		data.BookerPhone = &builder.bookerPhone
	}
	if builder.departmentNameSet {
		data.DepartmentName = &builder.departmentName
	}
	if builder.departmentIdSet {
		data.DepartmentId = &builder.departmentId
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.orderStatusSet {
		data.OrderStatus = &builder.orderStatus
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.isPaySet {
		data.IsPay = &builder.isPay
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.changeTimeSet {
		data.ChangeTime = &builder.changeTime
	}
	if builder.cancelTimeSet {
		data.CancelTime = &builder.cancelTime
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	if builder.departmentCodeSet {
		data.DepartmentCode = &builder.departmentCode
	}
	return data
}
