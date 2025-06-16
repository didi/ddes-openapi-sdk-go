package v1

// HotelOrder message HotelOrderTicketListItem {   TicketInfo ticket_info = 2 [(openapi.v3.property) = {description: \"原始票信息\"}]; }  国内酒店和国际酒店的order_info字段相同，可以共用；
type HotelOrder struct {
	OrderId                *string                `json:"order_id,omitempty"`           // 订单号 主单号
	CompanyName            *string                `json:"company_name,omitempty"`       // 企业名称 在滴滴注册的企业名称
	CompanyId              *string                `json:"company_id,omitempty"`         // 企业ID 滴滴企业ID
	BookerName             *string                `json:"booker_name,omitempty"`        // 预订人姓名 下单人
	BookerEmployeeId       *string                `json:"booker_employee_id,omitempty"` // 预订人员工编号 客户维护或者外部对接的数据
	BookerMemberId         *string                `json:"booker_member_id,omitempty"`   // 预订人滴滴ID 内部员工
	BookerPhone            *string                `json:"booker_phone,omitempty"`       // 预订人手机号 脱敏展示
	DepartmentName         *string                `json:"department_name,omitempty"`    // 预订人部门名称 部门链路
	DepartmentId           *string                `json:"department_id,omitempty"`      // 预订人部门ID 末级部门ID
	InstitutionName        *string                `json:"institution_name,omitempty"`   // 制度名称 滴滴制度名称
	InstitutionId          *string                `json:"institution_id,omitempty"`     // 制度ID
	ApprovalId             *string                `json:"approval_id,omitempty"`        // 差旅申请单ID 滴滴申请单ID
	OutApprovalId          *string                `json:"out_approval_id,omitempty"`    // 外部申请单ID 系统对接外部的申请单号
	OrderStatus            *string                `json:"order_status,omitempty"`       // 订单状态 枚举英文code: Submitted 已提交 Booking 已提交预订 Confirming 商家确认中 Booked 预订成功 Checkedin 已入住 Completed 已完成 Cancelledbeforepay 支付前取消 Cancelledafterpay 支付后取消 Noshow 预订未到 Unknow 未知
	IsPay                  *int32                 `json:"is_pay,omitempty"`             // 支付状态 枚举数字：0 未支付 1 已支付
	IsEarlyCheckout        *int32                 `json:"is_early_checkout,omitempty"`  // 提前离店 枚举数字： 0 非提前离店 1 提前离店
	PayType                *int32                 `json:"pay_type,omitempty"`           // 订单支付类型 枚举数字： 0 企业支付 1 个人支付 2 混合支付
	CreateTime             *string                `json:"create_time,omitempty"`        // 订单创建时间 格式：yyyy-MM-dd HH:mm:ss
	CompleteTime           *string                `json:"complete_time,omitempty"`      // 完单时间 格式：yyyy-MM-dd HH:mm:ss
	CancelTime             *string                `json:"cancel_time,omitempty"`        // 取消时间 格式：yyyy-MM-dd HH:mm:ss
	RefundTime             *string                `json:"refund_time,omitempty"`        // 退款时间 格式：yyyy-MM-dd HH:mm:ss
	ChangeTime             *string                `json:"change_time,omitempty"`        // 订单更新时间 格式：yyyy-MM-dd HH:mm:ss 查询条件
	CancelReason           *string                `json:"cancel_reason,omitempty"`      // 取消原因 用户取消原因
	BudgetCenterList       []BudgetCenterListItem `json:"budget_center_list,omitempty"` // 多成本中心 参考budget_center_list对象
	DepartmentCode         *string                `json:"department_code,omitempty"`    // * 以下为接口实际返回，单文档为标注的字段
	ExtraInfo              *string                `json:"extra_info,omitempty"`
	BookerPhoneCountryCode *string                `json:"booker_phone_country_code,omitempty"`
	ExtendFieldList        *string                `json:"extend_field_list,omitempty"`
}

type HotelOrderBuilder struct {
	orderId                   string // 订单号 主单号
	orderIdSet                bool
	companyName               string // 企业名称 在滴滴注册的企业名称
	companyNameSet            bool
	companyId                 string // 企业ID 滴滴企业ID
	companyIdSet              bool
	bookerName                string // 预订人姓名 下单人
	bookerNameSet             bool
	bookerEmployeeId          string // 预订人员工编号 客户维护或者外部对接的数据
	bookerEmployeeIdSet       bool
	bookerMemberId            string // 预订人滴滴ID 内部员工
	bookerMemberIdSet         bool
	bookerPhone               string // 预订人手机号 脱敏展示
	bookerPhoneSet            bool
	departmentName            string // 预订人部门名称 部门链路
	departmentNameSet         bool
	departmentId              string // 预订人部门ID 末级部门ID
	departmentIdSet           bool
	institutionName           string // 制度名称 滴滴制度名称
	institutionNameSet        bool
	institutionId             string // 制度ID
	institutionIdSet          bool
	approvalId                string // 差旅申请单ID 滴滴申请单ID
	approvalIdSet             bool
	outApprovalId             string // 外部申请单ID 系统对接外部的申请单号
	outApprovalIdSet          bool
	orderStatus               string // 订单状态 枚举英文code: Submitted 已提交 Booking 已提交预订 Confirming 商家确认中 Booked 预订成功 Checkedin 已入住 Completed 已完成 Cancelledbeforepay 支付前取消 Cancelledafterpay 支付后取消 Noshow 预订未到 Unknow 未知
	orderStatusSet            bool
	isPay                     int32 // 支付状态 枚举数字：0 未支付 1 已支付
	isPaySet                  bool
	isEarlyCheckout           int32 // 提前离店 枚举数字： 0 非提前离店 1 提前离店
	isEarlyCheckoutSet        bool
	payType                   int32 // 订单支付类型 枚举数字： 0 企业支付 1 个人支付 2 混合支付
	payTypeSet                bool
	createTime                string // 订单创建时间 格式：yyyy-MM-dd HH:mm:ss
	createTimeSet             bool
	completeTime              string // 完单时间 格式：yyyy-MM-dd HH:mm:ss
	completeTimeSet           bool
	cancelTime                string // 取消时间 格式：yyyy-MM-dd HH:mm:ss
	cancelTimeSet             bool
	refundTime                string // 退款时间 格式：yyyy-MM-dd HH:mm:ss
	refundTimeSet             bool
	changeTime                string // 订单更新时间 格式：yyyy-MM-dd HH:mm:ss 查询条件
	changeTimeSet             bool
	cancelReason              string // 取消原因 用户取消原因
	cancelReasonSet           bool
	budgetCenterList          []BudgetCenterListItem // 多成本中心 参考budget_center_list对象
	budgetCenterListSet       bool
	departmentCode            string // * 以下为接口实际返回，单文档为标注的字段
	departmentCodeSet         bool
	extraInfo                 string
	extraInfoSet              bool
	bookerPhoneCountryCode    string
	bookerPhoneCountryCodeSet bool
	extendFieldList           string
	extendFieldListSet        bool
}

func NewHotelOrderBuilder() *HotelOrderBuilder {
	return &HotelOrderBuilder{}
}
func (builder *HotelOrderBuilder) OrderId(orderId string) *HotelOrderBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) CompanyName(companyName string) *HotelOrderBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *HotelOrderBuilder) CompanyId(companyId string) *HotelOrderBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) BookerName(bookerName string) *HotelOrderBuilder {
	builder.bookerName = bookerName
	builder.bookerNameSet = true
	return builder
}
func (builder *HotelOrderBuilder) BookerEmployeeId(bookerEmployeeId string) *HotelOrderBuilder {
	builder.bookerEmployeeId = bookerEmployeeId
	builder.bookerEmployeeIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) BookerMemberId(bookerMemberId string) *HotelOrderBuilder {
	builder.bookerMemberId = bookerMemberId
	builder.bookerMemberIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) BookerPhone(bookerPhone string) *HotelOrderBuilder {
	builder.bookerPhone = bookerPhone
	builder.bookerPhoneSet = true
	return builder
}
func (builder *HotelOrderBuilder) DepartmentName(departmentName string) *HotelOrderBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *HotelOrderBuilder) DepartmentId(departmentId string) *HotelOrderBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) InstitutionName(institutionName string) *HotelOrderBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *HotelOrderBuilder) InstitutionId(institutionId string) *HotelOrderBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) ApprovalId(approvalId string) *HotelOrderBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) OutApprovalId(outApprovalId string) *HotelOrderBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *HotelOrderBuilder) OrderStatus(orderStatus string) *HotelOrderBuilder {
	builder.orderStatus = orderStatus
	builder.orderStatusSet = true
	return builder
}
func (builder *HotelOrderBuilder) IsPay(isPay int32) *HotelOrderBuilder {
	builder.isPay = isPay
	builder.isPaySet = true
	return builder
}
func (builder *HotelOrderBuilder) IsEarlyCheckout(isEarlyCheckout int32) *HotelOrderBuilder {
	builder.isEarlyCheckout = isEarlyCheckout
	builder.isEarlyCheckoutSet = true
	return builder
}
func (builder *HotelOrderBuilder) PayType(payType int32) *HotelOrderBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *HotelOrderBuilder) CreateTime(createTime string) *HotelOrderBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *HotelOrderBuilder) CompleteTime(completeTime string) *HotelOrderBuilder {
	builder.completeTime = completeTime
	builder.completeTimeSet = true
	return builder
}
func (builder *HotelOrderBuilder) CancelTime(cancelTime string) *HotelOrderBuilder {
	builder.cancelTime = cancelTime
	builder.cancelTimeSet = true
	return builder
}
func (builder *HotelOrderBuilder) RefundTime(refundTime string) *HotelOrderBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *HotelOrderBuilder) ChangeTime(changeTime string) *HotelOrderBuilder {
	builder.changeTime = changeTime
	builder.changeTimeSet = true
	return builder
}
func (builder *HotelOrderBuilder) CancelReason(cancelReason string) *HotelOrderBuilder {
	builder.cancelReason = cancelReason
	builder.cancelReasonSet = true
	return builder
}
func (builder *HotelOrderBuilder) BudgetCenterList(budgetCenterList []BudgetCenterListItem) *HotelOrderBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *HotelOrderBuilder) DepartmentCode(departmentCode string) *HotelOrderBuilder {
	builder.departmentCode = departmentCode
	builder.departmentCodeSet = true
	return builder
}
func (builder *HotelOrderBuilder) ExtraInfo(extraInfo string) *HotelOrderBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *HotelOrderBuilder) BookerPhoneCountryCode(bookerPhoneCountryCode string) *HotelOrderBuilder {
	builder.bookerPhoneCountryCode = bookerPhoneCountryCode
	builder.bookerPhoneCountryCodeSet = true
	return builder
}
func (builder *HotelOrderBuilder) ExtendFieldList(extendFieldList string) *HotelOrderBuilder {
	builder.extendFieldList = extendFieldList
	builder.extendFieldListSet = true
	return builder
}

func (builder *HotelOrderBuilder) Build() *HotelOrder {
	data := &HotelOrder{}
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
	if builder.orderStatusSet {
		data.OrderStatus = &builder.orderStatus
	}
	if builder.isPaySet {
		data.IsPay = &builder.isPay
	}
	if builder.isEarlyCheckoutSet {
		data.IsEarlyCheckout = &builder.isEarlyCheckout
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.completeTimeSet {
		data.CompleteTime = &builder.completeTime
	}
	if builder.cancelTimeSet {
		data.CancelTime = &builder.cancelTime
	}
	if builder.refundTimeSet {
		data.RefundTime = &builder.refundTime
	}
	if builder.changeTimeSet {
		data.ChangeTime = &builder.changeTime
	}
	if builder.cancelReasonSet {
		data.CancelReason = &builder.cancelReason
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = builder.budgetCenterList
	}
	if builder.departmentCodeSet {
		data.DepartmentCode = &builder.departmentCode
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.bookerPhoneCountryCodeSet {
		data.BookerPhoneCountryCode = &builder.bookerPhoneCountryCode
	}
	if builder.extendFieldListSet {
		data.ExtendFieldList = &builder.extendFieldList
	}
	return data
}
