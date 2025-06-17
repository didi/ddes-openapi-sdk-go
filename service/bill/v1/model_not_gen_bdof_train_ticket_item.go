package v1

// NotGenBDOfTrainTicketItem 未出账单 - 火车票,参考内部文档进行定义；
type NotGenBDOfTrainTicketItem struct {
	OriginTicketIdText         *string  `json:"origin_ticket_id_text,omitempty"`          // 原票id
	TicketIdText               *string  `json:"ticket_id_text,omitempty"`                 // 改签票id
	OrderId                    *string  `json:"order_id,omitempty"`                       // 当前订单号
	OrderIdText                *string  `json:"order_id_text,omitempty"`                  // 改签订单号
	TransactionType            *string  `json:"transaction_type,omitempty"`               // 订单类型 枚举【\"出票\"、\"改签\"、\"退票\"】
	Abnormal                   *string  `json:"abnormal,omitempty"`                       // 异常状态 枚举【\"\"、\"线下已改签\"、”线下已退票\"、“线下已改签后线下退票”】
	PayType                    *string  `json:"pay_type,omitempty"`                       // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	BookingMemberId            *int64   `json:"booking_member_id,omitempty"`              // 预定人id
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`            // 预订人姓名
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`          // 预订人工号
	BookingDepId               *int64   `json:"booking_dep_id,omitempty"`                 // 预定人部门id
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`               // 预订人部门名称
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`               // 预定人部门code
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`   // 预定人部门路径 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康昊公司>营销事业部>华南事业部>南宁省公司>南宁省公司(本部)\"
	PassengerMemberId          *int64   `json:"passenger_member_id,omitempty"`            // 乘车人id
	PassengerMemberName        *string  `json:"passenger_member_name,omitempty"`          // 乘车人姓名
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`        // 乘车人工号
	PassengerDepId             *int64   `json:"passenger_dep_id,omitempty"`               // 乘车人部门id
	PassengerDepName           *string  `json:"passenger_dep_name,omitempty"`             // 乘车人部门名称
	PassengerDepCode           *string  `json:"passenger_dep_code,omitempty"`             // 乘车人部门code
	PassengerParentPathDepName *string  `json:"passenger_parent_path_dep_name,omitempty"` // 乘车人部门路径 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康缘股份>康昊公司>营销事业部>江苏事业部>无锡省公司>无锡省公司(本部)\"
	BudgetCenterId             *int64   `json:"budget_center_id,omitempty"`               // 成本中心id
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"` // 所在成本中心
	BudgetCenterCode           *string  `json:"budget_center_code,omitempty"`             // 成本中心code 外部成本中心id
	PassengerMemberId2         *string  `json:"passengerMemberId,omitempty"`              // 乘车人ID 如：123432
	BudgetCenterCode2          *string  `json:"budgetCenterCode,omitempty"`               // 成本中心编码 如：3899
	BookingMemberId2           *string  `json:"bookingMemberId,omitempty"`                // 预订人员工ID 如：1111233
	BookingMemberNumber2       *string  `json:"bookingMemberNumber,omitempty"`            // 客户侧预订人工号 如：D234
	PassengerMemberNumber2     *string  `json:"passengerMemberNumber,omitempty"`          // 客户侧乘车人工号 如：D234
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`             // 成本中心名称
	OrderTime                  *string  `json:"order_time,omitempty"`                     // 预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	StartTime                  *string  `json:"start_time,omitempty"`                     // 出发时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 14:41:00\"
	EndTime                    *string  `json:"end_time,omitempty"`                       // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	FromStationName            *string  `json:"from_station_name,omitempty"`              // 出发站点
	ToStationName              *string  `json:"to_station_name,omitempty"`                // 到达站点
	TrainCode                  *string  `json:"train_code,omitempty"`                     // 车次
	SeatLevel                  *string  `json:"seat_level,omitempty"`                     // 座位
	SeatName                   *string  `json:"seat_name,omitempty"`                      // 席别
	TotalTicketPrice           *float64 `json:"total_ticket_price,omitempty"`             // 票价（单位：元）
	TicketPrice                *float64 `json:"ticket_price,omitempty"`                   // 企业票价（单位：元）
	PersonTicketPrice          *float64 `json:"person_ticket_price,omitempty"`            // 个人票价（单位：元）
	TotalHandingExtraFee       *float64 `json:"total_handing_extra_fee,omitempty"`        // 退改手续费（单位：元）
	HandingExtraFee            *float64 `json:"handing_extra_fee,omitempty"`              // 企业退改手续费（单位：元）
	PersonHandingExtraFee      *float64 `json:"person_handing_extra_fee,omitempty"`       // 个人退改手续费（单位：元）
	RefundFee                  *float64 `json:"refund_fee,omitempty"`                     // 退票费（单位：元）
	ChangeExtraFee             *float64 `json:"change_extra_fee,omitempty"`               // 改签手续费（单位：元）
	RefundExtraFee             *float64 `json:"refund_extra_fee,omitempty"`               // 退票手续费（单位：元）
	ServiceFee                 *float64 `json:"service_fee,omitempty"`                    // 平台使用费（单位：元）
	CompanyRealPay             *float64 `json:"company_real_pay,omitempty"`               // 企业实付金额（单位：元）
	PersonRealPay              *float64 `json:"person_real_pay,omitempty"`                // 个人实付金额（单位：元）
	OrderPay                   *float64 `json:"order_pay,omitempty"`                      // 订单金额（单位：元） 当前订单支付的总金额，订单总金额=票价+退改手续费+平台使用费+抢票服务费（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	StubPay                    *float64 `json:"stub_pay,omitempty"`                       // 票根金额（单位：元）
	IsChange                   *string  `json:"is_change,omitempty"`                      // 是否改签
	ChangeTime                 *string  `json:"change_time,omitempty"`                    // 改签时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	IsRefund                   *string  `json:"is_refund,omitempty"`                      // 是否退票
	RefundTime                 *string  `json:"refund_time,omitempty"`                    // 退票时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	RcNormalLevel              *string  `json:"rc_normal_level,omitempty"`                // 预订座席RC
	RcChangeLevel              *string  `json:"rc_change_level,omitempty"`                // 改签座席RC
	RequisitionId              *string  `json:"requisition_id,omitempty"`                 // 出差申请单号
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`                 // 出行目的
	TripDescription            *string  `json:"trip_description,omitempty"`               // 出行描述
	TripReason                 *string  `json:"trip_reason,omitempty"`                    // 出差事由
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
	Reason                     *string  `json:"reason,omitempty"`                         // 改签原因 如：临时取消
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
	ExtendInfo                 *string  `json:"extend_info,omitempty"`                    // 自定义（开票主体等信息)
	ApplyRefundTime            *string  `json:"apply_refund_time,omitempty"`              // 退票申请时间
	ApplyChangeTime            *string  `json:"apply_change_time,omitempty"`              // 改签申请时间
	SubAccountCompanyName      *string  `json:"sub_account_company_name,omitempty"`       // 子账户公司名称
	ExInfo01Code               *string  `json:"ex_info_01_code,omitempty"`                // 用户自定义拓展字段1编码
	ExInfo02Code               *string  `json:"ex_info_02_code,omitempty"`                // 用户自定义拓展字段2编码
	ExInfo03Code               *string  `json:"ex_info_03_code,omitempty"`                // 用户自定义拓展字段3编码
	ExInfo04Code               *string  `json:"ex_info_04_code,omitempty"`                // 用户自定义拓展字段4编码
	ExInfo05Code               *string  `json:"ex_info_05_code,omitempty"`                // 用户自定义拓展字段5编码
	ExInfo06Code               *string  `json:"ex_info_06_code,omitempty"`                // 用户自定义拓展字段6编码
	ExInfo07Code               *string  `json:"ex_info_07_code,omitempty"`                // 用户自定义拓展字段7编码
	ExInfo08Code               *string  `json:"ex_info_08_code,omitempty"`                // 用户自定义拓展字段8编码
	ParentInstiutionId         *int64   `json:"parent_instiution_id,omitempty"`           // 父制度ID
}

type NotGenBDOfTrainTicketItemBuilder struct {
	originTicketIdText              string // 原票id
	originTicketIdTextSet           bool
	ticketIdText                    string // 改签票id
	ticketIdTextSet                 bool
	orderId                         string // 当前订单号
	orderIdSet                      bool
	orderIdText                     string // 改签订单号
	orderIdTextSet                  bool
	transactionType                 string // 订单类型 枚举【\"出票\"、\"改签\"、\"退票\"】
	transactionTypeSet              bool
	abnormal                        string // 异常状态 枚举【\"\"、\"线下已改签\"、”线下已退票\"、“线下已改签后线下退票”】
	abnormalSet                     bool
	payType                         string // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	payTypeSet                      bool
	bookingMemberId                 int64 // 预定人id
	bookingMemberIdSet              bool
	bookingMemberName               string // 预订人姓名
	bookingMemberNameSet            bool
	bookingMemberNumber             string // 预订人工号
	bookingMemberNumberSet          bool
	bookingDepId                    int64 // 预定人部门id
	bookingDepIdSet                 bool
	bookingDepName                  string // 预订人部门名称
	bookingDepNameSet               bool
	bookingDepCode                  string // 预定人部门code
	bookingDepCodeSet               bool
	bookingParentPathDepName        string // 预定人部门路径 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康昊公司>营销事业部>华南事业部>南宁省公司>南宁省公司(本部)\"
	bookingParentPathDepNameSet     bool
	passengerMemberId               int64 // 乘车人id
	passengerMemberIdSet            bool
	passengerMemberName             string // 乘车人姓名
	passengerMemberNameSet          bool
	passengerMemberNumber           string // 乘车人工号
	passengerMemberNumberSet        bool
	passengerDepId                  int64 // 乘车人部门id
	passengerDepIdSet               bool
	passengerDepName                string // 乘车人部门名称
	passengerDepNameSet             bool
	passengerDepCode                string // 乘车人部门code
	passengerDepCodeSet             bool
	passengerParentPathDepName      string // 乘车人部门路径 示例：\"江苏康缘药业股份有限公司>江苏康缘药业股份有限公司>康缘集团>康缘股份>康昊公司>营销事业部>江苏事业部>无锡省公司>无锡省公司(本部)\"
	passengerParentPathDepNameSet   bool
	budgetCenterId                  int64 // 成本中心id
	budgetCenterIdSet               bool
	budgetCenterParentPathName      string // 所在成本中心
	budgetCenterParentPathNameSet   bool
	budgetCenterCode                string // 成本中心code 外部成本中心id
	budgetCenterCodeSet             bool
	passengerMemberIdREPEAT2        string // 乘车人ID 如：123432
	passengerMemberIdREPEAT2Set     bool
	budgetCenterCodeREPEAT2         string // 成本中心编码 如：3899
	budgetCenterCodeREPEAT2Set      bool
	bookingMemberIdREPEAT2          string // 预订人员工ID 如：1111233
	bookingMemberIdREPEAT2Set       bool
	bookingMemberNumberREPEAT2      string // 客户侧预订人工号 如：D234
	bookingMemberNumberREPEAT2Set   bool
	passengerMemberNumberREPEAT2    string // 客户侧乘车人工号 如：D234
	passengerMemberNumberREPEAT2Set bool
	budgetCenterName                string // 成本中心名称
	budgetCenterNameSet             bool
	orderTime                       string // 预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	orderTimeSet                    bool
	startTime                       string // 出发时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 14:41:00\"
	startTimeSet                    bool
	endTime                         string // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	endTimeSet                      bool
	fromStationName                 string // 出发站点
	fromStationNameSet              bool
	toStationName                   string // 到达站点
	toStationNameSet                bool
	trainCode                       string // 车次
	trainCodeSet                    bool
	seatLevel                       string // 座位
	seatLevelSet                    bool
	seatName                        string // 席别
	seatNameSet                     bool
	totalTicketPrice                float64 // 票价（单位：元）
	totalTicketPriceSet             bool
	ticketPrice                     float64 // 企业票价（单位：元）
	ticketPriceSet                  bool
	personTicketPrice               float64 // 个人票价（单位：元）
	personTicketPriceSet            bool
	totalHandingExtraFee            float64 // 退改手续费（单位：元）
	totalHandingExtraFeeSet         bool
	handingExtraFee                 float64 // 企业退改手续费（单位：元）
	handingExtraFeeSet              bool
	personHandingExtraFee           float64 // 个人退改手续费（单位：元）
	personHandingExtraFeeSet        bool
	refundFee                       float64 // 退票费（单位：元）
	refundFeeSet                    bool
	changeExtraFee                  float64 // 改签手续费（单位：元）
	changeExtraFeeSet               bool
	refundExtraFee                  float64 // 退票手续费（单位：元）
	refundExtraFeeSet               bool
	serviceFee                      float64 // 平台使用费（单位：元）
	serviceFeeSet                   bool
	companyRealPay                  float64 // 企业实付金额（单位：元）
	companyRealPaySet               bool
	personRealPay                   float64 // 个人实付金额（单位：元）
	personRealPaySet                bool
	orderPay                        float64 // 订单金额（单位：元） 当前订单支付的总金额，订单总金额=票价+退改手续费+平台使用费+抢票服务费（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	orderPaySet                     bool
	stubPay                         float64 // 票根金额（单位：元）
	stubPaySet                      bool
	isChange                        string // 是否改签
	isChangeSet                     bool
	changeTime                      string // 改签时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	changeTimeSet                   bool
	isRefund                        string // 是否退票
	isRefundSet                     bool
	refundTime                      string // 退票时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-08 15:11:00\"
	refundTimeSet                   bool
	rcNormalLevel                   string // 预订座席RC
	rcNormalLevelSet                bool
	rcChangeLevel                   string // 改签座席RC
	rcChangeLevelSet                bool
	requisitionId                   string // 出差申请单号
	requisitionIdSet                bool
	travelPurpose                   string // 出行目的
	travelPurposeSet                bool
	tripDescription                 string // 出行描述
	tripDescriptionSet              bool
	tripReason                      string // 出差事由
	tripReasonSet                   bool
	remark                          string // 备注
	remarkSet                       bool
	uniqueKey                       string // 唯一键 此条明细的唯一键
	uniqueKeySet                    bool
	legalEntityId                   int64 // 预定人开票子公司id
	legalEntityIdSet                bool
	legalEntityName                 string // 预定人开票子公司名称
	legalEntityNameSet              bool
	passengerLegalEntityId          int64 // 乘车人开票子公司id
	passengerLegalEntityIdSet       bool
	passengerLegalEntityName        string // 乘车人开票子公司名称
	passengerLegalEntityNameSet     bool
	outRequisitionId                string // 外部审批单ID
	outRequisitionIdSet             bool
	extraInfo                       string // 审批自定义
	extraInfoSet                    bool
	exInfo01                        string // 用户自定义拓展字段1
	exInfo01Set                     bool
	exInfo02                        string // 用户自定义拓展字段2
	exInfo02Set                     bool
	exInfo03                        string // 用户自定义拓展字段3
	exInfo03Set                     bool
	exInfo04                        string // 用户自定义拓展字段4
	exInfo04Set                     bool
	exInfo05                        string // 用户自定义拓展字段5
	exInfo05Set                     bool
	exInfo06                        string // 用户自定义拓展字段6
	exInfo06Set                     bool
	exInfo07                        string // 用户自定义拓展字段7
	exInfo07Set                     bool
	exInfo08                        string // 用户自定义拓展字段8
	exInfo08Set                     bool
	institutionName                 string // 制度名称
	institutionNameSet              bool
	fromCityName                    string // 出发城市名称
	fromCityNameSet                 bool
	fromCityId                      int64 // 出发城市id
	fromCityIdSet                   bool
	toCityName                      string // 到达城市名称
	toCityNameSet                   bool
	toCityId                        int64 // 到达城市id
	toCityIdSet                     bool
	settleTime                      string // 结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	settleTimeSet                   bool
	trainTicketRefundHandleFee      float64 // 退改手续费票根金额
	trainTicketRefundHandleFeeSet   bool
	reason                          string // 改签原因 如：临时取消
	reasonSet                       bool
	purchaseType                    string // 出票类型 枚举：正常出票、抢票出票
	purchaseTypeSet                 bool
	grabServiceFee                  float64 // 抢票服务费（单位：元）
	grabServiceFeeSet               bool
	rankName                        string // 职级 如：\"员工\"
	rankNameSet                     bool
	outLegalEntityId                string // 外部公司主体编号
	outLegalEntityIdSet             bool
	totalDifferenceFee              float64 // 差价退票费(总)
	totalDifferenceFeeSet           bool
	companyDifferenceFee            float64 // 差价退票费(企业)
	companyDifferenceFeeSet         bool
	personDifferenceFee             float64 // 差价退票费(个人)
	personDifferenceFeeSet          bool
	differenceFeeStub               float64 // 差价退票费(票根)
	differenceFeeStubSet            bool
	totalChangeFee                  float64 // 改签手续费(总)
	totalChangeFeeSet               bool
	companyChangeFee                float64 // 改签手续费(企业)
	companyChangeFeeSet             bool
	personChangeFee                 float64 // 改签手续费(个人)
	personChangeFeeSet              bool
	changeFeeStub                   float64 // 改签手续费(票根)
	changeFeeStubSet                bool
	institutionId                   int64 // 制度ID
	institutionIdSet                bool
	excludingServiceFee             float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet          bool
	extendInfo                      string // 自定义（开票主体等信息)
	extendInfoSet                   bool
	applyRefundTime                 string // 退票申请时间
	applyRefundTimeSet              bool
	applyChangeTime                 string // 改签申请时间
	applyChangeTimeSet              bool
	subAccountCompanyName           string // 子账户公司名称
	subAccountCompanyNameSet        bool
	exInfo01Code                    string // 用户自定义拓展字段1编码
	exInfo01CodeSet                 bool
	exInfo02Code                    string // 用户自定义拓展字段2编码
	exInfo02CodeSet                 bool
	exInfo03Code                    string // 用户自定义拓展字段3编码
	exInfo03CodeSet                 bool
	exInfo04Code                    string // 用户自定义拓展字段4编码
	exInfo04CodeSet                 bool
	exInfo05Code                    string // 用户自定义拓展字段5编码
	exInfo05CodeSet                 bool
	exInfo06Code                    string // 用户自定义拓展字段6编码
	exInfo06CodeSet                 bool
	exInfo07Code                    string // 用户自定义拓展字段7编码
	exInfo07CodeSet                 bool
	exInfo08Code                    string // 用户自定义拓展字段8编码
	exInfo08CodeSet                 bool
	parentInstiutionId              int64 // 父制度ID
	parentInstiutionIdSet           bool
}

func NewNotGenBDOfTrainTicketItemBuilder() *NotGenBDOfTrainTicketItemBuilder {
	return &NotGenBDOfTrainTicketItemBuilder{}
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OriginTicketIdText(originTicketIdText string) *NotGenBDOfTrainTicketItemBuilder {
	builder.originTicketIdText = originTicketIdText
	builder.originTicketIdTextSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TicketIdText(ticketIdText string) *NotGenBDOfTrainTicketItemBuilder {
	builder.ticketIdText = ticketIdText
	builder.ticketIdTextSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OrderId(orderId string) *NotGenBDOfTrainTicketItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OrderIdText(orderIdText string) *NotGenBDOfTrainTicketItemBuilder {
	builder.orderIdText = orderIdText
	builder.orderIdTextSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TransactionType(transactionType string) *NotGenBDOfTrainTicketItemBuilder {
	builder.transactionType = transactionType
	builder.transactionTypeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) Abnormal(abnormal string) *NotGenBDOfTrainTicketItemBuilder {
	builder.abnormal = abnormal
	builder.abnormalSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PayType(payType string) *NotGenBDOfTrainTicketItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingMemberId(bookingMemberId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingMemberName(bookingMemberName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingMemberNumber(bookingMemberNumber string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingDepId(bookingDepId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingDepId = bookingDepId
	builder.bookingDepIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingDepName(bookingDepName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingDepCode(bookingDepCode string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerMemberId(passengerMemberId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerMemberName(passengerMemberName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerDepId(passengerDepId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerDepId = passengerDepId
	builder.passengerDepIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerDepName(passengerDepName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerDepCode(passengerDepCode string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerParentPathDepName(passengerParentPathDepName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerParentPathDepName = passengerParentPathDepName
	builder.passengerParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BudgetCenterId(budgetCenterId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BudgetCenterCode(budgetCenterCode string) *NotGenBDOfTrainTicketItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerMemberId2(passengerMemberIdREPEAT2 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerMemberIdREPEAT2 = passengerMemberIdREPEAT2
	builder.passengerMemberIdREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BudgetCenterCode2(budgetCenterCodeREPEAT2 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.budgetCenterCodeREPEAT2 = budgetCenterCodeREPEAT2
	builder.budgetCenterCodeREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingMemberId2(bookingMemberIdREPEAT2 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingMemberIdREPEAT2 = bookingMemberIdREPEAT2
	builder.bookingMemberIdREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BookingMemberNumber2(bookingMemberNumberREPEAT2 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.bookingMemberNumberREPEAT2 = bookingMemberNumberREPEAT2
	builder.bookingMemberNumberREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerMemberNumber2(passengerMemberNumberREPEAT2 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerMemberNumberREPEAT2 = passengerMemberNumberREPEAT2
	builder.passengerMemberNumberREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) BudgetCenterName(budgetCenterName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OrderTime(orderTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.orderTime = orderTime
	builder.orderTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) StartTime(startTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) EndTime(endTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) FromStationName(fromStationName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ToStationName(toStationName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TrainCode(trainCode string) *NotGenBDOfTrainTicketItemBuilder {
	builder.trainCode = trainCode
	builder.trainCodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) SeatLevel(seatLevel string) *NotGenBDOfTrainTicketItemBuilder {
	builder.seatLevel = seatLevel
	builder.seatLevelSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) SeatName(seatName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.seatName = seatName
	builder.seatNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TotalTicketPrice(totalTicketPrice float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.totalTicketPrice = totalTicketPrice
	builder.totalTicketPriceSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TicketPrice(ticketPrice float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.ticketPrice = ticketPrice
	builder.ticketPriceSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PersonTicketPrice(personTicketPrice float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.personTicketPrice = personTicketPrice
	builder.personTicketPriceSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TotalHandingExtraFee(totalHandingExtraFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.totalHandingExtraFee = totalHandingExtraFee
	builder.totalHandingExtraFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) HandingExtraFee(handingExtraFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.handingExtraFee = handingExtraFee
	builder.handingExtraFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PersonHandingExtraFee(personHandingExtraFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.personHandingExtraFee = personHandingExtraFee
	builder.personHandingExtraFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RefundFee(refundFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.refundFee = refundFee
	builder.refundFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ChangeExtraFee(changeExtraFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.changeExtraFee = changeExtraFee
	builder.changeExtraFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RefundExtraFee(refundExtraFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.refundExtraFee = refundExtraFee
	builder.refundExtraFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ServiceFee(serviceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) CompanyRealPay(companyRealPay float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PersonRealPay(personRealPay float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.personRealPay = personRealPay
	builder.personRealPaySet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OrderPay(orderPay float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.orderPay = orderPay
	builder.orderPaySet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) StubPay(stubPay float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.stubPay = stubPay
	builder.stubPaySet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) IsChange(isChange string) *NotGenBDOfTrainTicketItemBuilder {
	builder.isChange = isChange
	builder.isChangeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ChangeTime(changeTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.changeTime = changeTime
	builder.changeTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) IsRefund(isRefund string) *NotGenBDOfTrainTicketItemBuilder {
	builder.isRefund = isRefund
	builder.isRefundSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RefundTime(refundTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RcNormalLevel(rcNormalLevel string) *NotGenBDOfTrainTicketItemBuilder {
	builder.rcNormalLevel = rcNormalLevel
	builder.rcNormalLevelSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RcChangeLevel(rcChangeLevel string) *NotGenBDOfTrainTicketItemBuilder {
	builder.rcChangeLevel = rcChangeLevel
	builder.rcChangeLevelSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RequisitionId(requisitionId string) *NotGenBDOfTrainTicketItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TravelPurpose(travelPurpose string) *NotGenBDOfTrainTicketItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TripDescription(tripDescription string) *NotGenBDOfTrainTicketItemBuilder {
	builder.tripDescription = tripDescription
	builder.tripDescriptionSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TripReason(tripReason string) *NotGenBDOfTrainTicketItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) Remark(remark string) *NotGenBDOfTrainTicketItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) UniqueKey(uniqueKey string) *NotGenBDOfTrainTicketItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) LegalEntityId(legalEntityId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) LegalEntityName(legalEntityName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OutRequisitionId(outRequisitionId string) *NotGenBDOfTrainTicketItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExtraInfo(extraInfo string) *NotGenBDOfTrainTicketItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo01(exInfo01 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo02(exInfo02 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo03(exInfo03 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo04(exInfo04 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo05(exInfo05 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo06(exInfo06 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo07(exInfo07 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo08(exInfo08 string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) InstitutionName(institutionName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) FromCityName(fromCityName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.fromCityName = fromCityName
	builder.fromCityNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) FromCityId(fromCityId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.fromCityId = fromCityId
	builder.fromCityIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ToCityName(toCityName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.toCityName = toCityName
	builder.toCityNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ToCityId(toCityId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.toCityId = toCityId
	builder.toCityIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) SettleTime(settleTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.settleTime = settleTime
	builder.settleTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TrainTicketRefundHandleFee(trainTicketRefundHandleFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.trainTicketRefundHandleFee = trainTicketRefundHandleFee
	builder.trainTicketRefundHandleFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) Reason(reason string) *NotGenBDOfTrainTicketItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PurchaseType(purchaseType string) *NotGenBDOfTrainTicketItemBuilder {
	builder.purchaseType = purchaseType
	builder.purchaseTypeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) GrabServiceFee(grabServiceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.grabServiceFee = grabServiceFee
	builder.grabServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) RankName(rankName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) OutLegalEntityId(outLegalEntityId string) *NotGenBDOfTrainTicketItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TotalDifferenceFee(totalDifferenceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.totalDifferenceFee = totalDifferenceFee
	builder.totalDifferenceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) CompanyDifferenceFee(companyDifferenceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.companyDifferenceFee = companyDifferenceFee
	builder.companyDifferenceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PersonDifferenceFee(personDifferenceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.personDifferenceFee = personDifferenceFee
	builder.personDifferenceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) DifferenceFeeStub(differenceFeeStub float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.differenceFeeStub = differenceFeeStub
	builder.differenceFeeStubSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) TotalChangeFee(totalChangeFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.totalChangeFee = totalChangeFee
	builder.totalChangeFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) CompanyChangeFee(companyChangeFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.companyChangeFee = companyChangeFee
	builder.companyChangeFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) PersonChangeFee(personChangeFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.personChangeFee = personChangeFee
	builder.personChangeFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ChangeFeeStub(changeFeeStub float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.changeFeeStub = changeFeeStub
	builder.changeFeeStubSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) InstitutionId(institutionId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *NotGenBDOfTrainTicketItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExtendInfo(extendInfo string) *NotGenBDOfTrainTicketItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ApplyRefundTime(applyRefundTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.applyRefundTime = applyRefundTime
	builder.applyRefundTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ApplyChangeTime(applyChangeTime string) *NotGenBDOfTrainTicketItemBuilder {
	builder.applyChangeTime = applyChangeTime
	builder.applyChangeTimeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *NotGenBDOfTrainTicketItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo01Code(exInfo01Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo02Code(exInfo02Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo03Code(exInfo03Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo04Code(exInfo04Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo05Code(exInfo05Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo06Code(exInfo06Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo07Code(exInfo07Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ExInfo08Code(exInfo08Code string) *NotGenBDOfTrainTicketItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}
func (builder *NotGenBDOfTrainTicketItemBuilder) ParentInstiutionId(parentInstiutionId int64) *NotGenBDOfTrainTicketItemBuilder {
	builder.parentInstiutionId = parentInstiutionId
	builder.parentInstiutionIdSet = true
	return builder
}

func (builder *NotGenBDOfTrainTicketItemBuilder) Build() *NotGenBDOfTrainTicketItem {
	data := &NotGenBDOfTrainTicketItem{}
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
	if builder.passengerMemberIdREPEAT2Set {
		data.PassengerMemberId2 = &builder.passengerMemberIdREPEAT2
	}
	if builder.budgetCenterCodeREPEAT2Set {
		data.BudgetCenterCode2 = &builder.budgetCenterCodeREPEAT2
	}
	if builder.bookingMemberIdREPEAT2Set {
		data.BookingMemberId2 = &builder.bookingMemberIdREPEAT2
	}
	if builder.bookingMemberNumberREPEAT2Set {
		data.BookingMemberNumber2 = &builder.bookingMemberNumberREPEAT2
	}
	if builder.passengerMemberNumberREPEAT2Set {
		data.PassengerMemberNumber2 = &builder.passengerMemberNumberREPEAT2
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
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.applyRefundTimeSet {
		data.ApplyRefundTime = &builder.applyRefundTime
	}
	if builder.applyChangeTimeSet {
		data.ApplyChangeTime = &builder.applyChangeTime
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
	if builder.parentInstiutionIdSet {
		data.ParentInstiutionId = &builder.parentInstiutionId
	}
	return data
}
