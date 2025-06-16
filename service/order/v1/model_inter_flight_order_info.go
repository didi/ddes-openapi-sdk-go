package v1

// InterFlightOrderInfo 国际机票订单信息
type InterFlightOrderInfo struct {
	OrderId          *string                `json:"order_id,omitempty"`           // 订单号，主单号
	CompanyName      *string                `json:"company_name,omitempty"`       // 企业名称，在滴滴注册的企业名称
	CompanyId        *string                `json:"company_id,omitempty"`         // 企业ID，滴滴企业ID
	BookerName       *string                `json:"booker_name,omitempty"`        // 预订人姓名，下单人
	BookerEmployeeId *string                `json:"booker_employee_id,omitempty"` // 预订人员工工号，客户维护或者外部对接的数据
	BookerMemberId   *string                `json:"booker_member_id,omitempty"`   // 预订人滴滴ID，内部员工
	BookerPhone      *string                `json:"booker_phone,omitempty"`       // 预订人手机号，脱敏展示
	DepartmentName   *string                `json:"department_name,omitempty"`    // 预订人部门名称，部门链路
	DepartmentId     *string                `json:"department_id,omitempty"`      // 预订人部门ID，滴滴末级部门ID
	InstitutionName  *string                `json:"institution_name,omitempty"`   // 制度名称，滴滴制度名称
	InstitutionId    *string                `json:"institution_id,omitempty"`     // 制度ID
	ApprovalId       *string                `json:"approval_id,omitempty"`        // 差旅申请单ID，滴滴申请单ID
	OutApprovalId    *string                `json:"out_approval_id,omitempty"`    // 外部申请单ID，系统对接外部的申请单号
	FlightWay        *int32                 `json:"flight_way,omitempty"`         // 航程类别，枚举数字： 1 单程, 2 往返
	OrderStatus      *string                `json:"order_status,omitempty"`       // 订单状态，枚举英文code：Tobepaid 待支付；Waitforapproval 待审批；Ticketing 待出票；Ticketed 已出票；Accountingforchangefee 改签待核算；Changing 改签中；refunding 退票中；Completed 已完成；Refunded 已退票；Partialrefund 部分退票；Closed 订单关闭；Cancelled 已取消；Unknow 未知
	CreateTime       *string                `json:"create_time,omitempty"`        // 订单创建时间，格式：yyyy-MM-dd HH:mm:ss
	IsPay            *int32                 `json:"is_pay,omitempty"`             // 支付状态，枚举数字：0 未支付 1 已支付
	PayType          *int32                 `json:"pay_type,omitempty"`           // 订单支付类型，枚举数字：0 企业支付 1 个人支付 2 混合支付
	CancelTime       *string                `json:"cancel_time,omitempty"`        // 取消订单时间，格式：yyyy-MM-dd HH:mm:ss
	ChangeTime       *string                `json:"change_time,omitempty"`        // 订单更新时间，格式：yyyy-MM-dd HH:mm:ss 查询条件
	BudgetCenterList []BudgetCenterListItem `json:"budget_center_list,omitempty"` // 多成本中心，参考budget_center_list对象
	DepartmentCode   *string                `json:"department_code,omitempty"`    // * 接口实际返回，文档未标注
}

type InterFlightOrderInfoBuilder struct {
	orderId             string // 订单号，主单号
	orderIdSet          bool
	companyName         string // 企业名称，在滴滴注册的企业名称
	companyNameSet      bool
	companyId           string // 企业ID，滴滴企业ID
	companyIdSet        bool
	bookerName          string // 预订人姓名，下单人
	bookerNameSet       bool
	bookerEmployeeId    string // 预订人员工工号，客户维护或者外部对接的数据
	bookerEmployeeIdSet bool
	bookerMemberId      string // 预订人滴滴ID，内部员工
	bookerMemberIdSet   bool
	bookerPhone         string // 预订人手机号，脱敏展示
	bookerPhoneSet      bool
	departmentName      string // 预订人部门名称，部门链路
	departmentNameSet   bool
	departmentId        string // 预订人部门ID，滴滴末级部门ID
	departmentIdSet     bool
	institutionName     string // 制度名称，滴滴制度名称
	institutionNameSet  bool
	institutionId       string // 制度ID
	institutionIdSet    bool
	approvalId          string // 差旅申请单ID，滴滴申请单ID
	approvalIdSet       bool
	outApprovalId       string // 外部申请单ID，系统对接外部的申请单号
	outApprovalIdSet    bool
	flightWay           int32 // 航程类别，枚举数字： 1 单程, 2 往返
	flightWaySet        bool
	orderStatus         string // 订单状态，枚举英文code：Tobepaid 待支付；Waitforapproval 待审批；Ticketing 待出票；Ticketed 已出票；Accountingforchangefee 改签待核算；Changing 改签中；refunding 退票中；Completed 已完成；Refunded 已退票；Partialrefund 部分退票；Closed 订单关闭；Cancelled 已取消；Unknow 未知
	orderStatusSet      bool
	createTime          string // 订单创建时间，格式：yyyy-MM-dd HH:mm:ss
	createTimeSet       bool
	isPay               int32 // 支付状态，枚举数字：0 未支付 1 已支付
	isPaySet            bool
	payType             int32 // 订单支付类型，枚举数字：0 企业支付 1 个人支付 2 混合支付
	payTypeSet          bool
	cancelTime          string // 取消订单时间，格式：yyyy-MM-dd HH:mm:ss
	cancelTimeSet       bool
	changeTime          string // 订单更新时间，格式：yyyy-MM-dd HH:mm:ss 查询条件
	changeTimeSet       bool
	budgetCenterList    []BudgetCenterListItem // 多成本中心，参考budget_center_list对象
	budgetCenterListSet bool
	departmentCode      string // * 接口实际返回，文档未标注
	departmentCodeSet   bool
}

func NewInterFlightOrderInfoBuilder() *InterFlightOrderInfoBuilder {
	return &InterFlightOrderInfoBuilder{}
}
func (builder *InterFlightOrderInfoBuilder) OrderId(orderId string) *InterFlightOrderInfoBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) CompanyName(companyName string) *InterFlightOrderInfoBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) CompanyId(companyId string) *InterFlightOrderInfoBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) BookerName(bookerName string) *InterFlightOrderInfoBuilder {
	builder.bookerName = bookerName
	builder.bookerNameSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) BookerEmployeeId(bookerEmployeeId string) *InterFlightOrderInfoBuilder {
	builder.bookerEmployeeId = bookerEmployeeId
	builder.bookerEmployeeIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) BookerMemberId(bookerMemberId string) *InterFlightOrderInfoBuilder {
	builder.bookerMemberId = bookerMemberId
	builder.bookerMemberIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) BookerPhone(bookerPhone string) *InterFlightOrderInfoBuilder {
	builder.bookerPhone = bookerPhone
	builder.bookerPhoneSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) DepartmentName(departmentName string) *InterFlightOrderInfoBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) DepartmentId(departmentId string) *InterFlightOrderInfoBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) InstitutionName(institutionName string) *InterFlightOrderInfoBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) InstitutionId(institutionId string) *InterFlightOrderInfoBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) ApprovalId(approvalId string) *InterFlightOrderInfoBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) OutApprovalId(outApprovalId string) *InterFlightOrderInfoBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) FlightWay(flightWay int32) *InterFlightOrderInfoBuilder {
	builder.flightWay = flightWay
	builder.flightWaySet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) OrderStatus(orderStatus string) *InterFlightOrderInfoBuilder {
	builder.orderStatus = orderStatus
	builder.orderStatusSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) CreateTime(createTime string) *InterFlightOrderInfoBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) IsPay(isPay int32) *InterFlightOrderInfoBuilder {
	builder.isPay = isPay
	builder.isPaySet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) PayType(payType int32) *InterFlightOrderInfoBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) CancelTime(cancelTime string) *InterFlightOrderInfoBuilder {
	builder.cancelTime = cancelTime
	builder.cancelTimeSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) ChangeTime(changeTime string) *InterFlightOrderInfoBuilder {
	builder.changeTime = changeTime
	builder.changeTimeSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *InterFlightOrderInfoBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *InterFlightOrderInfoBuilder) DepartmentCode(departmentCode string) *InterFlightOrderInfoBuilder {
	builder.departmentCode = departmentCode
	builder.departmentCodeSet = true
	return builder
}

func (builder *InterFlightOrderInfoBuilder) Build() *InterFlightOrderInfo {
	data := &InterFlightOrderInfo{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
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
	if builder.flightWaySet {
		data.FlightWay = &builder.flightWay
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
	if builder.cancelTimeSet {
		data.CancelTime = &builder.cancelTime
	}
	if builder.changeTimeSet {
		data.ChangeTime = &builder.changeTime
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	if builder.departmentCodeSet {
		data.DepartmentCode = &builder.departmentCode
	}
	return data
}
