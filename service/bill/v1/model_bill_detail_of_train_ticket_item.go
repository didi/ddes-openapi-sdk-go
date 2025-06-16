package v1

// BillDetailOfTrainTicketItem 已出账单 ~ 火车票返回参数,参考内部文档进行定义；
type BillDetailOfTrainTicketItem struct {
	BillId                     *int64   `json:"bill_id,omitempty"`                        // 账单id 根结点批次id
	OriginTicketIdText         *string  `json:"origin_ticket_id_text,omitempty"`          // 原票id 滴滴企业版app生成的原始票id
	TicketIdText               *string  `json:"ticket_id_text,omitempty"`                 // 改签票id 滴滴企业版app生成的改签票id
	OrderId                    *string  `json:"order_id,omitempty"`                       // 当前订单号
	OrderIdText                *string  `json:"order_id_text,omitempty"`                  // 改签订单号 滴滴企业版app生成的改签订单号
	TransactionType            *string  `json:"transaction_type,omitempty"`               // 该笔订单的交易类型 枚举【\"出票\"、\"改签\"、\"退票\"】
	Abnormal                   *string  `json:"abnormal,omitempty"`                       // 火车票发生线下退改的异常状态 枚举【\"\"、\"线下已改签\"、”线下已退票\"、“线下已改签后线下退票”】
	PayType                    *string  `json:"pay_type,omitempty"`                       // 当前订单支付金额的方式，如：企业支付、混合支付（指企业支付和个人支付两种支付方式同时支付一笔订单）
	BookingMemberId            *int64   `json:"booking_member_id,omitempty"`              // 预定人id
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`            // 预订人姓名
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`          // 预订人工号
	BookingDepId               *int64   `json:"booking_dep_id,omitempty"`                 // 预定人部门id
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`               // 预订人部门名称
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`               // 预定人部门编号（客户在ES后台输入的部门编号）
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`   // 预订火车票订单的员工所在部门 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康昊公司>营销事业部>华南事业部>南宁省公司>南宁省公司(本部)\"
	PassengerMemberId          *int64   `json:"passenger_member_id,omitempty"`            // 乘车人id
	PassengerMemberName        *string  `json:"passenger_member_name,omitempty"`          // 乘车人姓名
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`        // 乘车人工号
	PassengerDepId             *int64   `json:"passenger_dep_id,omitempty"`               // 乘车人部门id
	PassengerDepName           *string  `json:"passenger_dep_name,omitempty"`             // 乘车人部门名称
	PassengerDepCode           *string  `json:"passenger_dep_code,omitempty"`             // 乘车人部门编号
	PassengerParentPathDepName *string  `json:"passenger_parent_path_dep_name,omitempty"` // 当前火车票的乘车人所在部门 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康缘股份>康昊公司>营销事业部>江苏事业部>无锡省公司>无锡省公司(本部)\"
	BudgetCenterId             *int64   `json:"budget_center_id,omitempty"`               // 成本中心id 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"` // 所在成本中心 管理员在ES后台差旅管控模块维护的所在成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	BudgetCenterCode           *string  `json:"budget_center_code,omitempty"`             // 成本中心code 外部成本中心id
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`             // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	OrderTime                  *string  `json:"order_time,omitempty"`                     // 预订时间 火车票订单的下单时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	StartTime                  *string  `json:"start_time,omitempty"`                     // 出发时间 火车票展示的发车时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 14:41:00\"
	EndTime                    *string  `json:"end_time,omitempty"`                       // 到达时间 火车票展示的到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	FromStationName            *string  `json:"from_station_name,omitempty"`              // 出发站点 火车票展示的出发站点
	ToStationName              *string  `json:"to_station_name,omitempty"`                // 到达站点 火车票展示的到达站点
	TrainCode                  *string  `json:"train_code,omitempty"`                     // 车次 火车票所展示的列车车次
	SeatLevel                  *string  `json:"seat_level,omitempty"`                     // 座位 当前火车票的座位号
	SeatName                   *string  `json:"seat_name,omitempty"`                      // 席别 火车席位的类别，如硬座、硬卧、一等座、二等座
	TotalTicketPrice           *float64 `json:"total_ticket_price,omitempty"`             // 票价（单位：元） =企业票价+个人票价
	TicketPrice                *float64 `json:"ticket_price,omitempty"`                   // 企业票价（单位：元） 当前票价通过企业账户支付金额部分
	PersonTicketPrice          *float64 `json:"person_ticket_price,omitempty"`            // 个人票价（单位：元） 当前票价通过个人账户支付金额部分
	TotalHandingExtraFee       *float64 `json:"total_handing_extra_fee,omitempty"`        // 退改手续费（单位：元） 退票和改签时所产生的手续费 =企业退改手续费+个人退改手续费
	HandingExtraFee            *float64 `json:"handing_extra_fee,omitempty"`              // 企业退改手续费（单位：元） 退票和改签时企业所支付的手续费
	PersonHandingExtraFee      *float64 `json:"person_handing_extra_fee,omitempty"`       // 个人退改手续费（单位：元） 退票和改签时个人所支付的手续费
	RefundFee                  *float64 `json:"refund_fee,omitempty"`                     // 企业退票费（单位：元） 【火车票改版后 此字段已在excel账单下掉】
	ChangeExtraFee             *float64 `json:"change_extra_fee,omitempty"`               // 企业改签手续费（单位：元）【火车票改版后 此字段已在excel账单下掉】
	RefundExtraFee             *float64 `json:"refund_extra_fee,omitempty"`               // 企业退票手续费（单位：元） 【火车票改版后 此字段已在excel账单下掉】
	ServiceFee                 *float64 `json:"service_fee,omitempty"`                    // 平台使用费（单位：元）
	CompanyRealPay             *float64 `json:"company_real_pay,omitempty"`               // 企业实付金额（单位：元）
	PersonRealPay              *float64 `json:"person_real_pay,omitempty"`                // 个人实付金额（单位：元）
	OrderPay                   *float64 `json:"order_pay,omitempty"`                      // 订单金额（单位：元） 当前订单支付的总金额，订单总金额=票价+退改手续费+平台使用费+抢票服务费。（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	StubPay                    *float64 `json:"stub_pay,omitempty"`                       // 票根金额（单位：元） 票据的存根金额
	IsChange                   *string  `json:"is_change,omitempty"`                      // 是否改签 枚举：”是”、“否”
	ChangeTime                 *string  `json:"change_time,omitempty"`                    // 改签时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	IsRefund                   *string  `json:"is_refund,omitempty"`                      // 是否退票 枚举：”是”、“否”
	RefundTime                 *string  `json:"refund_time,omitempty"`                    // 退票时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	RcNormalLevel              *string  `json:"rc_normal_level,omitempty"`                // 预订座席RC 超出预订规定坐席的超标原因
	RcChangeLevel              *string  `json:"rc_change_level,omitempty"`                // 改签座席RC 超出改签规定坐席的超标原因
	RequisitionId              *string  `json:"requisition_id,omitempty"`                 // 出差申请单号 出差时所生成的单号
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`                 // 出行目的 出行人下单填写的出行目的
	TripDescription            *string  `json:"trip_description,omitempty"`               // 出行描述 出行人下单填写的出行描述
	TripReason                 *string  `json:"trip_reason,omitempty"`                    // 出差事由 出行人下单填写的出行事由
	Remark                     *string  `json:"remark,omitempty"`                         // 备注
	UniqueKey                  *string  `json:"unique_key,omitempty"`                     // 唯一键 此条明细的唯一键
	LegalEntityId              *int64   `json:"legal_entity_id,omitempty"`                // 预定人开票子公司id
	LegalEntityName            *string  `json:"legal_entity_name,omitempty"`              // 预定人开票子公司名称
	PassengerLegalEntityId     *int64   `json:"passenger_legal_entity_id,omitempty"`      // 乘车人开票子公司id
	PassengerLegalEntityName   *string  `json:"passenger_legal_entity_name,omitempty"`    // 乘车人开票子公司名称
	OutRequisitionId           *string  `json:"out_requisition_id,omitempty"`             // 外部审批单ID
	ExtraInfo                  *string  `json:"extra_info,omitempty"`                     // 审批自定义
	ExInfo01                   *string  `json:"ex_info_01,omitempty"`                     // 用户自定义拓展字段1
	ExInfo02                   *string  `json:"ex_info_02,omitempty"`                     // 用户自定义拓展字段2
	ExInfo03                   *string  `json:"ex_info_03,omitempty"`                     // 用户自定义拓展字段3
	ExInfo04                   *string  `json:"ex_info_04,omitempty"`                     // 用户自定义拓展字段4
	ExInfo05                   *string  `json:"ex_info_05,omitempty"`                     // 用户自定义拓展字段5
	ExInfo06                   *string  `json:"ex_info_06,omitempty"`                     // 用户自定义拓展字段6
	ExInfo07                   *string  `json:"ex_info_07,omitempty"`                     // 用户自定义拓展字段7
	ExInfo08                   *string  `json:"ex_info_08,omitempty"`                     // 用户自定义拓展字段8
	InstitutionName            *string  `json:"institution_name,omitempty"`               // 制度名称
	FromCityName               *string  `json:"from_city_name,omitempty"`                 // 出发城市名称
	FromCityId                 *int64   `json:"from_city_id,omitempty"`                   // 出发城市id
	ToCityName                 *string  `json:"to_city_name,omitempty"`                   // 到达城市名称
	ToCityId                   *int64   `json:"to_city_id,omitempty"`                     // 到达城市id
	SettleTime                 *string  `json:"settle_time,omitempty"`                    // 结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	TrainTicketRefundHandleFee *float64 `json:"train_ticket_refund_handle_fee,omitempty"` // 退改手续费票根金额
	Reason                     *string  `json:"reason,omitempty"`                         // 退改签原因 如：临时取消
	AddedGoodsName             *string  `json:"added_goods_name,omitempty"`               // 手工单产品类型 线下预订
	AddedCutReason             *string  `json:"added_cut_reason,omitempty"`               // 手工单立减原因 中文
	AddedCutServiceFee         *float64 `json:"added_cut_service_fee,omitempty"`          // 手工单 立减平台使用费
	BeforeCutServiceFee        *float64 `json:"before_cut_service_fee,omitempty"`         // 手工单 立减前平台使用费
	ApprovalFirst              *string  `json:"approval_first,omitempty"`                 // 一级审批人
	ApprovalSecond             *string  `json:"approval_second,omitempty"`                // 二级审批人
	ApprovalThird              *string  `json:"approval_third,omitempty"`                 // 三级审批人
	ProjectExtInfo             *string  `json:"project_ext_info,omitempty"`               // 项目自定义
	OriginOrderId              *string  `json:"origin_order_id,omitempty"`                // 原始订单号 如：A改B改C，则原始订单号为A
	ApprovalHistory            *string  `json:"approval_history,omitempty"`               // 预定/超标情况下的审批历史
	ApprovalNormalHistory      *string  `json:"approval_normal_history,omitempty"`        // 正常差旅申请的审批历史
	PurchaseType               *string  `json:"purchase_type,omitempty"`                  // 出票类型 枚举：正常出票、抢票出票
	GrabServiceFee             *float64 `json:"grab_service_fee,omitempty"`               // 抢票服务费（单位：元）
	RankName                   *string  `json:"rank_name,omitempty"`                      // 职级 如：\"员工\"
	OutLegalEntityId           *string  `json:"out_legal_entity_id,omitempty"`            // 外部公司主体编号
	TotalDifferenceFee         *float64 `json:"total_difference_fee,omitempty"`           // 差价退票费(总)
	CompanyDifferenceFee       *float64 `json:"company_difference_fee,omitempty"`         // 差价退票费(企业)
	PersonDifferenceFee        *float64 `json:"person_difference_fee,omitempty"`          // 差价退票费(个人)
	DifferenceFeeStub          *float64 `json:"difference_fee_stub,omitempty"`            // 差价退票费(票根)
	TotalChangeFee             *float64 `json:"total_change_fee,omitempty"`               // 改签手续费(总)
	CompanyChangeFee           *float64 `json:"company_change_fee,omitempty"`             // 改签手续费(企业)
	PersonChangeFee            *float64 `json:"person_change_fee,omitempty"`              // 改签手续费(个人)
	ChangeFeeStub              *float64 `json:"change_fee_stub,omitempty"`                // 改签手续费(票根)
	InstitutionId              *int64   `json:"institution_id,omitempty"`                 // 制度ID
	ExcludingServiceFee        *float64 `json:"excluding_service_fee,omitempty"`          // 企业实付（不包含平台使用费）
	ApplyRefundTime            *string  `json:"apply_refund_time,omitempty"`              // 退票申请时间
	ApplyChangeTime            *string  `json:"apply_change_time,omitempty"`              // 改签申请时间
	ExtendInfo                 *string  `json:"extend_info,omitempty"`                    // 自定义（开票主体等信息)
	SubAccountCompanyName      *string  `json:"sub_account_company_name,omitempty"`       // 子账户公司名称
	ExInfo01Code               *string  `json:"ex_info_01_code,omitempty"`                // 用户自定义拓展字段1编码
	ExInfo02Code               *string  `json:"ex_info_02_code,omitempty"`                // 用户自定义拓展字段2编码
	ExInfo03Code               *string  `json:"ex_info_03_code,omitempty"`                // 用户自定义拓展字段3编码
	ExInfo04Code               *string  `json:"ex_info_04_code,omitempty"`                // 用户自定义拓展字段4编码
	ExInfo05Code               *string  `json:"ex_info_05_code,omitempty"`                // 用户自定义拓展字段5编码
	ExInfo06Code               *string  `json:"ex_info_06_code,omitempty"`                // 用户自定义拓展字段6编码
	ExInfo07Code               *string  `json:"ex_info_07_code,omitempty"`                // 用户自定义拓展字段7编码
	ExInfo08Code               *string  `json:"ex_info_08_code,omitempty"`                // 用户自定义拓展字段8编码
}

type BillDetailOfTrainTicketItemBuilder struct {
	billId                        int64 // 账单id 根结点批次id
	billIdSet                     bool
	originTicketIdText            string // 原票id 滴滴企业版app生成的原始票id
	originTicketIdTextSet         bool
	ticketIdText                  string // 改签票id 滴滴企业版app生成的改签票id
	ticketIdTextSet               bool
	orderId                       string // 当前订单号
	orderIdSet                    bool
	orderIdText                   string // 改签订单号 滴滴企业版app生成的改签订单号
	orderIdTextSet                bool
	transactionType               string // 该笔订单的交易类型 枚举【\"出票\"、\"改签\"、\"退票\"】
	transactionTypeSet            bool
	abnormal                      string // 火车票发生线下退改的异常状态 枚举【\"\"、\"线下已改签\"、”线下已退票\"、“线下已改签后线下退票”】
	abnormalSet                   bool
	payType                       string // 当前订单支付金额的方式，如：企业支付、混合支付（指企业支付和个人支付两种支付方式同时支付一笔订单）
	payTypeSet                    bool
	bookingMemberId               int64 // 预定人id
	bookingMemberIdSet            bool
	bookingMemberName             string // 预订人姓名
	bookingMemberNameSet          bool
	bookingMemberNumber           string // 预订人工号
	bookingMemberNumberSet        bool
	bookingDepId                  int64 // 预定人部门id
	bookingDepIdSet               bool
	bookingDepName                string // 预订人部门名称
	bookingDepNameSet             bool
	bookingDepCode                string // 预定人部门编号（客户在ES后台输入的部门编号）
	bookingDepCodeSet             bool
	bookingParentPathDepName      string // 预订火车票订单的员工所在部门 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康昊公司>营销事业部>华南事业部>南宁省公司>南宁省公司(本部)\"
	bookingParentPathDepNameSet   bool
	passengerMemberId             int64 // 乘车人id
	passengerMemberIdSet          bool
	passengerMemberName           string // 乘车人姓名
	passengerMemberNameSet        bool
	passengerMemberNumber         string // 乘车人工号
	passengerMemberNumberSet      bool
	passengerDepId                int64 // 乘车人部门id
	passengerDepIdSet             bool
	passengerDepName              string // 乘车人部门名称
	passengerDepNameSet           bool
	passengerDepCode              string // 乘车人部门编号
	passengerDepCodeSet           bool
	passengerParentPathDepName    string // 当前火车票的乘车人所在部门 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康缘股份>康昊公司>营销事业部>江苏事业部>无锡省公司>无锡省公司(本部)\"
	passengerParentPathDepNameSet bool
	budgetCenterId                int64 // 成本中心id 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	budgetCenterIdSet             bool
	budgetCenterParentPathName    string // 所在成本中心 管理员在ES后台差旅管控模块维护的所在成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	budgetCenterParentPathNameSet bool
	budgetCenterCode              string // 成本中心code 外部成本中心id
	budgetCenterCodeSet           bool
	budgetCenterName              string // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	budgetCenterNameSet           bool
	orderTime                     string // 预订时间 火车票订单的下单时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	orderTimeSet                  bool
	startTime                     string // 出发时间 火车票展示的发车时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 14:41:00\"
	startTimeSet                  bool
	endTime                       string // 到达时间 火车票展示的到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	endTimeSet                    bool
	fromStationName               string // 出发站点 火车票展示的出发站点
	fromStationNameSet            bool
	toStationName                 string // 到达站点 火车票展示的到达站点
	toStationNameSet              bool
	trainCode                     string // 车次 火车票所展示的列车车次
	trainCodeSet                  bool
	seatLevel                     string // 座位 当前火车票的座位号
	seatLevelSet                  bool
	seatName                      string // 席别 火车席位的类别，如硬座、硬卧、一等座、二等座
	seatNameSet                   bool
	totalTicketPrice              float64 // 票价（单位：元） =企业票价+个人票价
	totalTicketPriceSet           bool
	ticketPrice                   float64 // 企业票价（单位：元） 当前票价通过企业账户支付金额部分
	ticketPriceSet                bool
	personTicketPrice             float64 // 个人票价（单位：元） 当前票价通过个人账户支付金额部分
	personTicketPriceSet          bool
	totalHandingExtraFee          float64 // 退改手续费（单位：元） 退票和改签时所产生的手续费 =企业退改手续费+个人退改手续费
	totalHandingExtraFeeSet       bool
	handingExtraFee               float64 // 企业退改手续费（单位：元） 退票和改签时企业所支付的手续费
	handingExtraFeeSet            bool
	personHandingExtraFee         float64 // 个人退改手续费（单位：元） 退票和改签时个人所支付的手续费
	personHandingExtraFeeSet      bool
	refundFee                     float64 // 企业退票费（单位：元） 【火车票改版后 此字段已在excel账单下掉】
	refundFeeSet                  bool
	changeExtraFee                float64 // 企业改签手续费（单位：元）【火车票改版后 此字段已在excel账单下掉】
	changeExtraFeeSet             bool
	refundExtraFee                float64 // 企业退票手续费（单位：元） 【火车票改版后 此字段已在excel账单下掉】
	refundExtraFeeSet             bool
	serviceFee                    float64 // 平台使用费（单位：元）
	serviceFeeSet                 bool
	companyRealPay                float64 // 企业实付金额（单位：元）
	companyRealPaySet             bool
	personRealPay                 float64 // 个人实付金额（单位：元）
	personRealPaySet              bool
	orderPay                      float64 // 订单金额（单位：元） 当前订单支付的总金额，订单总金额=票价+退改手续费+平台使用费+抢票服务费。（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	orderPaySet                   bool
	stubPay                       float64 // 票根金额（单位：元） 票据的存根金额
	stubPaySet                    bool
	isChange                      string // 是否改签 枚举：”是”、“否”
	isChangeSet                   bool
	changeTime                    string // 改签时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	changeTimeSet                 bool
	isRefund                      string // 是否退票 枚举：”是”、“否”
	isRefundSet                   bool
	refundTime                    string // 退票时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	refundTimeSet                 bool
	rcNormalLevel                 string // 预订座席RC 超出预订规定坐席的超标原因
	rcNormalLevelSet              bool
	rcChangeLevel                 string // 改签座席RC 超出改签规定坐席的超标原因
	rcChangeLevelSet              bool
	requisitionId                 string // 出差申请单号 出差时所生成的单号
	requisitionIdSet              bool
	travelPurpose                 string // 出行目的 出行人下单填写的出行目的
	travelPurposeSet              bool
	tripDescription               string // 出行描述 出行人下单填写的出行描述
	tripDescriptionSet            bool
	tripReason                    string // 出差事由 出行人下单填写的出行事由
	tripReasonSet                 bool
	remark                        string // 备注
	remarkSet                     bool
	uniqueKey                     string // 唯一键 此条明细的唯一键
	uniqueKeySet                  bool
	legalEntityId                 int64 // 预定人开票子公司id
	legalEntityIdSet              bool
	legalEntityName               string // 预定人开票子公司名称
	legalEntityNameSet            bool
	passengerLegalEntityId        int64 // 乘车人开票子公司id
	passengerLegalEntityIdSet     bool
	passengerLegalEntityName      string // 乘车人开票子公司名称
	passengerLegalEntityNameSet   bool
	outRequisitionId              string // 外部审批单ID
	outRequisitionIdSet           bool
	extraInfo                     string // 审批自定义
	extraInfoSet                  bool
	exInfo01                      string // 用户自定义拓展字段1
	exInfo01Set                   bool
	exInfo02                      string // 用户自定义拓展字段2
	exInfo02Set                   bool
	exInfo03                      string // 用户自定义拓展字段3
	exInfo03Set                   bool
	exInfo04                      string // 用户自定义拓展字段4
	exInfo04Set                   bool
	exInfo05                      string // 用户自定义拓展字段5
	exInfo05Set                   bool
	exInfo06                      string // 用户自定义拓展字段6
	exInfo06Set                   bool
	exInfo07                      string // 用户自定义拓展字段7
	exInfo07Set                   bool
	exInfo08                      string // 用户自定义拓展字段8
	exInfo08Set                   bool
	institutionName               string // 制度名称
	institutionNameSet            bool
	fromCityName                  string // 出发城市名称
	fromCityNameSet               bool
	fromCityId                    int64 // 出发城市id
	fromCityIdSet                 bool
	toCityName                    string // 到达城市名称
	toCityNameSet                 bool
	toCityId                      int64 // 到达城市id
	toCityIdSet                   bool
	settleTime                    string // 结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	settleTimeSet                 bool
	trainTicketRefundHandleFee    float64 // 退改手续费票根金额
	trainTicketRefundHandleFeeSet bool
	reason                        string // 退改签原因 如：临时取消
	reasonSet                     bool
	addedGoodsName                string // 手工单产品类型 线下预订
	addedGoodsNameSet             bool
	addedCutReason                string // 手工单立减原因 中文
	addedCutReasonSet             bool
	addedCutServiceFee            float64 // 手工单 立减平台使用费
	addedCutServiceFeeSet         bool
	beforeCutServiceFee           float64 // 手工单 立减前平台使用费
	beforeCutServiceFeeSet        bool
	approvalFirst                 string // 一级审批人
	approvalFirstSet              bool
	approvalSecond                string // 二级审批人
	approvalSecondSet             bool
	approvalThird                 string // 三级审批人
	approvalThirdSet              bool
	projectExtInfo                string // 项目自定义
	projectExtInfoSet             bool
	originOrderId                 string // 原始订单号 如：A改B改C，则原始订单号为A
	originOrderIdSet              bool
	approvalHistory               string // 预定/超标情况下的审批历史
	approvalHistorySet            bool
	approvalNormalHistory         string // 正常差旅申请的审批历史
	approvalNormalHistorySet      bool
	purchaseType                  string // 出票类型 枚举：正常出票、抢票出票
	purchaseTypeSet               bool
	grabServiceFee                float64 // 抢票服务费（单位：元）
	grabServiceFeeSet             bool
	rankName                      string // 职级 如：\"员工\"
	rankNameSet                   bool
	outLegalEntityId              string // 外部公司主体编号
	outLegalEntityIdSet           bool
	totalDifferenceFee            float64 // 差价退票费(总)
	totalDifferenceFeeSet         bool
	companyDifferenceFee          float64 // 差价退票费(企业)
	companyDifferenceFeeSet       bool
	personDifferenceFee           float64 // 差价退票费(个人)
	personDifferenceFeeSet        bool
	differenceFeeStub             float64 // 差价退票费(票根)
	differenceFeeStubSet          bool
	totalChangeFee                float64 // 改签手续费(总)
	totalChangeFeeSet             bool
	companyChangeFee              float64 // 改签手续费(企业)
	companyChangeFeeSet           bool
	personChangeFee               float64 // 改签手续费(个人)
	personChangeFeeSet            bool
	changeFeeStub                 float64 // 改签手续费(票根)
	changeFeeStubSet              bool
	institutionId                 int64 // 制度ID
	institutionIdSet              bool
	excludingServiceFee           float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet        bool
	applyRefundTime               string // 退票申请时间
	applyRefundTimeSet            bool
	applyChangeTime               string // 改签申请时间
	applyChangeTimeSet            bool
	extendInfo                    string // 自定义（开票主体等信息)
	extendInfoSet                 bool
	subAccountCompanyName         string // 子账户公司名称
	subAccountCompanyNameSet      bool
	exInfo01Code                  string // 用户自定义拓展字段1编码
	exInfo01CodeSet               bool
	exInfo02Code                  string // 用户自定义拓展字段2编码
	exInfo02CodeSet               bool
	exInfo03Code                  string // 用户自定义拓展字段3编码
	exInfo03CodeSet               bool
	exInfo04Code                  string // 用户自定义拓展字段4编码
	exInfo04CodeSet               bool
	exInfo05Code                  string // 用户自定义拓展字段5编码
	exInfo05CodeSet               bool
	exInfo06Code                  string // 用户自定义拓展字段6编码
	exInfo06CodeSet               bool
	exInfo07Code                  string // 用户自定义拓展字段7编码
	exInfo07CodeSet               bool
	exInfo08Code                  string // 用户自定义拓展字段8编码
	exInfo08CodeSet               bool
}

func NewBillDetailOfTrainTicketItemBuilder() *BillDetailOfTrainTicketItemBuilder {
	return &BillDetailOfTrainTicketItemBuilder{}
}
func (builder *BillDetailOfTrainTicketItemBuilder) BillId(billId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OriginTicketIdText(originTicketIdText string) *BillDetailOfTrainTicketItemBuilder {
	builder.originTicketIdText = originTicketIdText
	builder.originTicketIdTextSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TicketIdText(ticketIdText string) *BillDetailOfTrainTicketItemBuilder {
	builder.ticketIdText = ticketIdText
	builder.ticketIdTextSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OrderId(orderId string) *BillDetailOfTrainTicketItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OrderIdText(orderIdText string) *BillDetailOfTrainTicketItemBuilder {
	builder.orderIdText = orderIdText
	builder.orderIdTextSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TransactionType(transactionType string) *BillDetailOfTrainTicketItemBuilder {
	builder.transactionType = transactionType
	builder.transactionTypeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) Abnormal(abnormal string) *BillDetailOfTrainTicketItemBuilder {
	builder.abnormal = abnormal
	builder.abnormalSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PayType(payType string) *BillDetailOfTrainTicketItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingMemberId(bookingMemberId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingMemberName(bookingMemberName string) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingMemberNumber(bookingMemberNumber string) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingDepId(bookingDepId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingDepId = bookingDepId
	builder.bookingDepIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingDepName(bookingDepName string) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingDepCode(bookingDepCode string) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *BillDetailOfTrainTicketItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerMemberId(passengerMemberId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerMemberName(passengerMemberName string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerDepId(passengerDepId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerDepId = passengerDepId
	builder.passengerDepIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerDepName(passengerDepName string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerDepCode(passengerDepCode string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerParentPathDepName(passengerParentPathDepName string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerParentPathDepName = passengerParentPathDepName
	builder.passengerParentPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BudgetCenterId(budgetCenterId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *BillDetailOfTrainTicketItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BudgetCenterCode(budgetCenterCode string) *BillDetailOfTrainTicketItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BudgetCenterName(budgetCenterName string) *BillDetailOfTrainTicketItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OrderTime(orderTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.orderTime = orderTime
	builder.orderTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) StartTime(startTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) EndTime(endTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) FromStationName(fromStationName string) *BillDetailOfTrainTicketItemBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ToStationName(toStationName string) *BillDetailOfTrainTicketItemBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TrainCode(trainCode string) *BillDetailOfTrainTicketItemBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) SeatLevel(seatLevel string) *BillDetailOfTrainTicketItemBuilder {
	builder.seatLevel = seatLevel
	builder.seatLevelSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) SeatName(seatName string) *BillDetailOfTrainTicketItemBuilder {
	builder.seatName = seatName
	builder.seatNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TotalTicketPrice(totalTicketPrice float64) *BillDetailOfTrainTicketItemBuilder {
	builder.totalTicketPrice = totalTicketPrice
	builder.totalTicketPriceSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TicketPrice(ticketPrice float64) *BillDetailOfTrainTicketItemBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PersonTicketPrice(personTicketPrice float64) *BillDetailOfTrainTicketItemBuilder {
	builder.personTicketPrice = personTicketPrice
	builder.personTicketPriceSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TotalHandingExtraFee(totalHandingExtraFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.totalHandingExtraFee = totalHandingExtraFee
	builder.totalHandingExtraFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) HandingExtraFee(handingExtraFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.handingExtraFee = handingExtraFee
	builder.handingExtraFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PersonHandingExtraFee(personHandingExtraFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.personHandingExtraFee = personHandingExtraFee
	builder.personHandingExtraFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RefundFee(refundFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.refundFee = refundFee
	builder.refundFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ChangeExtraFee(changeExtraFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.changeExtraFee = changeExtraFee
	builder.changeExtraFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RefundExtraFee(refundExtraFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.refundExtraFee = refundExtraFee
	builder.refundExtraFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ServiceFee(serviceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) CompanyRealPay(companyRealPay float64) *BillDetailOfTrainTicketItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PersonRealPay(personRealPay float64) *BillDetailOfTrainTicketItemBuilder {
	builder.personRealPay = personRealPay
	builder.personRealPaySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OrderPay(orderPay float64) *BillDetailOfTrainTicketItemBuilder {
	builder.orderPay = orderPay
	builder.orderPaySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) StubPay(stubPay float64) *BillDetailOfTrainTicketItemBuilder {
	builder.stubPay = stubPay
	builder.stubPaySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) IsChange(isChange string) *BillDetailOfTrainTicketItemBuilder {
	builder.isChange = isChange
	builder.isChangeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ChangeTime(changeTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.changeTime = changeTime
	builder.changeTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) IsRefund(isRefund string) *BillDetailOfTrainTicketItemBuilder {
	builder.isRefund = isRefund
	builder.isRefundSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RefundTime(refundTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RcNormalLevel(rcNormalLevel string) *BillDetailOfTrainTicketItemBuilder {
	builder.rcNormalLevel = rcNormalLevel
	builder.rcNormalLevelSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RcChangeLevel(rcChangeLevel string) *BillDetailOfTrainTicketItemBuilder {
	builder.rcChangeLevel = rcChangeLevel
	builder.rcChangeLevelSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RequisitionId(requisitionId string) *BillDetailOfTrainTicketItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TravelPurpose(travelPurpose string) *BillDetailOfTrainTicketItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TripDescription(tripDescription string) *BillDetailOfTrainTicketItemBuilder {
	builder.tripDescription = tripDescription
	builder.tripDescriptionSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TripReason(tripReason string) *BillDetailOfTrainTicketItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) Remark(remark string) *BillDetailOfTrainTicketItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) UniqueKey(uniqueKey string) *BillDetailOfTrainTicketItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) LegalEntityId(legalEntityId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) LegalEntityName(legalEntityName string) *BillDetailOfTrainTicketItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *BillDetailOfTrainTicketItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OutRequisitionId(outRequisitionId string) *BillDetailOfTrainTicketItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExtraInfo(extraInfo string) *BillDetailOfTrainTicketItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo01(exInfo01 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo02(exInfo02 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo03(exInfo03 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo04(exInfo04 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo05(exInfo05 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo06(exInfo06 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo07(exInfo07 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo08(exInfo08 string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) InstitutionName(institutionName string) *BillDetailOfTrainTicketItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) FromCityName(fromCityName string) *BillDetailOfTrainTicketItemBuilder {
	builder.fromCityName = fromCityName
	builder.fromCityNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) FromCityId(fromCityId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.fromCityId = fromCityId
	builder.fromCityIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ToCityName(toCityName string) *BillDetailOfTrainTicketItemBuilder {
	builder.toCityName = toCityName
	builder.toCityNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ToCityId(toCityId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) SettleTime(settleTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.settleTime = settleTime
	builder.settleTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TrainTicketRefundHandleFee(trainTicketRefundHandleFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.trainTicketRefundHandleFee = trainTicketRefundHandleFee
	builder.trainTicketRefundHandleFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) Reason(reason string) *BillDetailOfTrainTicketItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) AddedGoodsName(addedGoodsName string) *BillDetailOfTrainTicketItemBuilder {
	builder.addedGoodsName = addedGoodsName
	builder.addedGoodsNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) AddedCutReason(addedCutReason string) *BillDetailOfTrainTicketItemBuilder {
	builder.addedCutReason = addedCutReason
	builder.addedCutReasonSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) AddedCutServiceFee(addedCutServiceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.addedCutServiceFee = addedCutServiceFee
	builder.addedCutServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) BeforeCutServiceFee(beforeCutServiceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.beforeCutServiceFee = beforeCutServiceFee
	builder.beforeCutServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApprovalFirst(approvalFirst string) *BillDetailOfTrainTicketItemBuilder {
	builder.approvalFirst = approvalFirst
	builder.approvalFirstSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApprovalSecond(approvalSecond string) *BillDetailOfTrainTicketItemBuilder {
	builder.approvalSecond = approvalSecond
	builder.approvalSecondSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApprovalThird(approvalThird string) *BillDetailOfTrainTicketItemBuilder {
	builder.approvalThird = approvalThird
	builder.approvalThirdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ProjectExtInfo(projectExtInfo string) *BillDetailOfTrainTicketItemBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OriginOrderId(originOrderId string) *BillDetailOfTrainTicketItemBuilder {
	builder.originOrderId = originOrderId
	builder.originOrderIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApprovalHistory(approvalHistory string) *BillDetailOfTrainTicketItemBuilder {
	builder.approvalHistory = approvalHistory
	builder.approvalHistorySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApprovalNormalHistory(approvalNormalHistory string) *BillDetailOfTrainTicketItemBuilder {
	builder.approvalNormalHistory = approvalNormalHistory
	builder.approvalNormalHistorySet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PurchaseType(purchaseType string) *BillDetailOfTrainTicketItemBuilder {
	builder.purchaseType = purchaseType
	builder.purchaseTypeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) GrabServiceFee(grabServiceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.grabServiceFee = grabServiceFee
	builder.grabServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) RankName(rankName string) *BillDetailOfTrainTicketItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) OutLegalEntityId(outLegalEntityId string) *BillDetailOfTrainTicketItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TotalDifferenceFee(totalDifferenceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.totalDifferenceFee = totalDifferenceFee
	builder.totalDifferenceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) CompanyDifferenceFee(companyDifferenceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.companyDifferenceFee = companyDifferenceFee
	builder.companyDifferenceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PersonDifferenceFee(personDifferenceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.personDifferenceFee = personDifferenceFee
	builder.personDifferenceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) DifferenceFeeStub(differenceFeeStub float64) *BillDetailOfTrainTicketItemBuilder {
	builder.differenceFeeStub = differenceFeeStub
	builder.differenceFeeStubSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) TotalChangeFee(totalChangeFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.totalChangeFee = totalChangeFee
	builder.totalChangeFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) CompanyChangeFee(companyChangeFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.companyChangeFee = companyChangeFee
	builder.companyChangeFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) PersonChangeFee(personChangeFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.personChangeFee = personChangeFee
	builder.personChangeFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ChangeFeeStub(changeFeeStub float64) *BillDetailOfTrainTicketItemBuilder {
	builder.changeFeeStub = changeFeeStub
	builder.changeFeeStubSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) InstitutionId(institutionId int64) *BillDetailOfTrainTicketItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *BillDetailOfTrainTicketItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApplyRefundTime(applyRefundTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.applyRefundTime = applyRefundTime
	builder.applyRefundTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ApplyChangeTime(applyChangeTime string) *BillDetailOfTrainTicketItemBuilder {
	builder.applyChangeTime = applyChangeTime
	builder.applyChangeTimeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExtendInfo(extendInfo string) *BillDetailOfTrainTicketItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *BillDetailOfTrainTicketItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo01Code(exInfo01Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo02Code(exInfo02Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo03Code(exInfo03Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo04Code(exInfo04Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo05Code(exInfo05Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo06Code(exInfo06Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo07Code(exInfo07Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *BillDetailOfTrainTicketItemBuilder) ExInfo08Code(exInfo08Code string) *BillDetailOfTrainTicketItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}

func (builder *BillDetailOfTrainTicketItemBuilder) Build() *BillDetailOfTrainTicketItem {
	data := &BillDetailOfTrainTicketItem{}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.originTicketIdTextSet {
		data.OriginTicketIdText = &builder.originTicketIdText
	}
	if builder.ticketIdTextSet {
		data.TicketIdText = &builder.ticketIdText
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.orderIdTextSet {
		data.OrderIdText = &builder.orderIdText
	}
	if builder.transactionTypeSet {
		data.TransactionType = &builder.transactionType
	}
	if builder.abnormalSet {
		data.Abnormal = &builder.abnormal
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.bookingMemberIdSet {
		data.BookingMemberId = &builder.bookingMemberId
	}
	if builder.bookingMemberNameSet {
		data.BookingMemberName = &builder.bookingMemberName
	}
	if builder.bookingMemberNumberSet {
		data.BookingMemberNumber = &builder.bookingMemberNumber
	}
	if builder.bookingDepIdSet {
		data.BookingDepId = &builder.bookingDepId
	}
	if builder.bookingDepNameSet {
		data.BookingDepName = &builder.bookingDepName
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.bookingParentPathDepNameSet {
		data.BookingParentPathDepName = &builder.bookingParentPathDepName
	}
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	if builder.passengerMemberNameSet {
		data.PassengerMemberName = &builder.passengerMemberName
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
	}
	if builder.passengerDepIdSet {
		data.PassengerDepId = &builder.passengerDepId
	}
	if builder.passengerDepNameSet {
		data.PassengerDepName = &builder.passengerDepName
	}
	if builder.passengerDepCodeSet {
		data.PassengerDepCode = &builder.passengerDepCode
	}
	if builder.passengerParentPathDepNameSet {
		data.PassengerParentPathDepName = &builder.passengerParentPathDepName
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.budgetCenterParentPathNameSet {
		data.BudgetCenterParentPathName = &builder.budgetCenterParentPathName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.orderTimeSet {
		data.OrderTime = &builder.orderTime
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.fromStationNameSet {
		data.FromStationName = &builder.fromStationName
	}
	if builder.toStationNameSet {
		data.ToStationName = &builder.toStationName
	}
	if builder.trainCodeSet {
		data.TrainCode = &builder.trainCode
	}
	if builder.seatLevelSet {
		data.SeatLevel = &builder.seatLevel
	}
	if builder.seatNameSet {
		data.SeatName = &builder.seatName
	}
	if builder.totalTicketPriceSet {
		data.TotalTicketPrice = &builder.totalTicketPrice
	}
	if builder.ticketPriceSet {
		data.TicketPrice = &builder.ticketPrice
	}
	if builder.personTicketPriceSet {
		data.PersonTicketPrice = &builder.personTicketPrice
	}
	if builder.totalHandingExtraFeeSet {
		data.TotalHandingExtraFee = &builder.totalHandingExtraFee
	}
	if builder.handingExtraFeeSet {
		data.HandingExtraFee = &builder.handingExtraFee
	}
	if builder.personHandingExtraFeeSet {
		data.PersonHandingExtraFee = &builder.personHandingExtraFee
	}
	if builder.refundFeeSet {
		data.RefundFee = &builder.refundFee
	}
	if builder.changeExtraFeeSet {
		data.ChangeExtraFee = &builder.changeExtraFee
	}
	if builder.refundExtraFeeSet {
		data.RefundExtraFee = &builder.refundExtraFee
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personRealPaySet {
		data.PersonRealPay = &builder.personRealPay
	}
	if builder.orderPaySet {
		data.OrderPay = &builder.orderPay
	}
	if builder.stubPaySet {
		data.StubPay = &builder.stubPay
	}
	if builder.isChangeSet {
		data.IsChange = &builder.isChange
	}
	if builder.changeTimeSet {
		data.ChangeTime = &builder.changeTime
	}
	if builder.isRefundSet {
		data.IsRefund = &builder.isRefund
	}
	if builder.refundTimeSet {
		data.RefundTime = &builder.refundTime
	}
	if builder.rcNormalLevelSet {
		data.RcNormalLevel = &builder.rcNormalLevel
	}
	if builder.rcChangeLevelSet {
		data.RcChangeLevel = &builder.rcChangeLevel
	}
	if builder.requisitionIdSet {
		data.RequisitionId = &builder.requisitionId
	}
	if builder.travelPurposeSet {
		data.TravelPurpose = &builder.travelPurpose
	}
	if builder.tripDescriptionSet {
		data.TripDescription = &builder.tripDescription
	}
	if builder.tripReasonSet {
		data.TripReason = &builder.tripReason
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.uniqueKeySet {
		data.UniqueKey = &builder.uniqueKey
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.passengerLegalEntityIdSet {
		data.PassengerLegalEntityId = &builder.passengerLegalEntityId
	}
	if builder.passengerLegalEntityNameSet {
		data.PassengerLegalEntityName = &builder.passengerLegalEntityName
	}
	if builder.outRequisitionIdSet {
		data.OutRequisitionId = &builder.outRequisitionId
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
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
	if builder.exInfo04Set {
		data.ExInfo04 = &builder.exInfo04
	}
	if builder.exInfo05Set {
		data.ExInfo05 = &builder.exInfo05
	}
	if builder.exInfo06Set {
		data.ExInfo06 = &builder.exInfo06
	}
	if builder.exInfo07Set {
		data.ExInfo07 = &builder.exInfo07
	}
	if builder.exInfo08Set {
		data.ExInfo08 = &builder.exInfo08
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.fromCityNameSet {
		data.FromCityName = &builder.fromCityName
	}
	if builder.fromCityIdSet {
		data.FromCityId = &builder.fromCityId
	}
	if builder.toCityNameSet {
		data.ToCityName = &builder.toCityName
	}
	if builder.toCityIdSet {
		data.ToCityId = &builder.toCityId
	}
	if builder.settleTimeSet {
		data.SettleTime = &builder.settleTime
	}
	if builder.trainTicketRefundHandleFeeSet {
		data.TrainTicketRefundHandleFee = &builder.trainTicketRefundHandleFee
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.addedGoodsNameSet {
		data.AddedGoodsName = &builder.addedGoodsName
	}
	if builder.addedCutReasonSet {
		data.AddedCutReason = &builder.addedCutReason
	}
	if builder.addedCutServiceFeeSet {
		data.AddedCutServiceFee = &builder.addedCutServiceFee
	}
	if builder.beforeCutServiceFeeSet {
		data.BeforeCutServiceFee = &builder.beforeCutServiceFee
	}
	if builder.approvalFirstSet {
		data.ApprovalFirst = &builder.approvalFirst
	}
	if builder.approvalSecondSet {
		data.ApprovalSecond = &builder.approvalSecond
	}
	if builder.approvalThirdSet {
		data.ApprovalThird = &builder.approvalThird
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.originOrderIdSet {
		data.OriginOrderId = &builder.originOrderId
	}
	if builder.approvalHistorySet {
		data.ApprovalHistory = &builder.approvalHistory
	}
	if builder.approvalNormalHistorySet {
		data.ApprovalNormalHistory = &builder.approvalNormalHistory
	}
	if builder.purchaseTypeSet {
		data.PurchaseType = &builder.purchaseType
	}
	if builder.grabServiceFeeSet {
		data.GrabServiceFee = &builder.grabServiceFee
	}
	if builder.rankNameSet {
		data.RankName = &builder.rankName
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.totalDifferenceFeeSet {
		data.TotalDifferenceFee = &builder.totalDifferenceFee
	}
	if builder.companyDifferenceFeeSet {
		data.CompanyDifferenceFee = &builder.companyDifferenceFee
	}
	if builder.personDifferenceFeeSet {
		data.PersonDifferenceFee = &builder.personDifferenceFee
	}
	if builder.differenceFeeStubSet {
		data.DifferenceFeeStub = &builder.differenceFeeStub
	}
	if builder.totalChangeFeeSet {
		data.TotalChangeFee = &builder.totalChangeFee
	}
	if builder.companyChangeFeeSet {
		data.CompanyChangeFee = &builder.companyChangeFee
	}
	if builder.personChangeFeeSet {
		data.PersonChangeFee = &builder.personChangeFee
	}
	if builder.changeFeeStubSet {
		data.ChangeFeeStub = &builder.changeFeeStub
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.excludingServiceFeeSet {
		data.ExcludingServiceFee = &builder.excludingServiceFee
	}
	if builder.applyRefundTimeSet {
		data.ApplyRefundTime = &builder.applyRefundTime
	}
	if builder.applyChangeTimeSet {
		data.ApplyChangeTime = &builder.applyChangeTime
	}
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.subAccountCompanyNameSet {
		data.SubAccountCompanyName = &builder.subAccountCompanyName
	}
	if builder.exInfo01CodeSet {
		data.ExInfo01Code = &builder.exInfo01Code
	}
	if builder.exInfo02CodeSet {
		data.ExInfo02Code = &builder.exInfo02Code
	}
	if builder.exInfo03CodeSet {
		data.ExInfo03Code = &builder.exInfo03Code
	}
	if builder.exInfo04CodeSet {
		data.ExInfo04Code = &builder.exInfo04Code
	}
	if builder.exInfo05CodeSet {
		data.ExInfo05Code = &builder.exInfo05Code
	}
	if builder.exInfo06CodeSet {
		data.ExInfo06Code = &builder.exInfo06Code
	}
	if builder.exInfo07CodeSet {
		data.ExInfo07Code = &builder.exInfo07Code
	}
	if builder.exInfo08CodeSet {
		data.ExInfo08Code = &builder.exInfo08Code
	}
	return data
}
