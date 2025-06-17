package v1

// NotGenBDOfDomesticFlightItem 未出账单 - 国内机票,参考内部文档进行定义；
type NotGenBDOfDomesticFlightItem struct {
	OriginTicketId             *string  `json:"origin_ticket_id,omitempty"`               // 原票id
	TicketIdText               *string  `json:"ticket_id_text,omitempty"`                 // 改签票id
	PassengerMemberName        *string  `json:"passenger_member_name,omitempty"`          // 乘机人姓名
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`        // 乘机人工号
	PassengerDepName           *string  `json:"passenger_dep_name,omitempty"`             // 乘机人部门
	PassengerParentPathDepName *string  `json:"passenger_parent_path_dep_name,omitempty"` // 乘机人所在部门 示例：\"机票测试专用751>优秀部门\"
	PassengerDepCode           *string  `json:"passenger_dep_code,omitempty"`             // 乘机人部门编号
	SettleTime                 *string  `json:"settle_time,omitempty"`                    // 完单时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	MemberId                   *string  `json:"member_id,omitempty"`                      // 乘机人员工ID
	OrderFixType               *int32   `json:"order_fix_type,omitempty"`                 // 订单补偿类型 枚举【0：非补偿单，1：补偿单】 该字段滴滴内部使用，外部可不用关注该字段
	OrderFixSettleTime         *string  `json:"order_fix_settle_time,omitempty"`          // 补偿单的结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	PassengerName              *string  `json:"passenger_name,omitempty"`                 // 乘机人姓名（冗余字段）
	Department                 *string  `json:"department,omitempty"`                     // 乘机人部门（冗余字段）
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`             // 成本中心名称
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"` // 所在成本中心 示例：\"机票测试专用751>高优项目\"
	BudgetCenterId             *string  `json:"budget_center_id,omitempty"`               // 成本中心ID
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`            // 预订人姓名
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`          // 预订人工号
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`               // 预订人部门
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`   // 预订人所在部门 示例：\"机票测试专用751>优秀部门\"
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`               // 预订人部门编号
	RequisitionId              *string  `json:"requisition_id,omitempty"`                 // 出差申请单号
	OutRequisitionId           *string  `json:"out_requisition_id,omitempty"`             // 外部审批单ID
	OrderId                    *string  `json:"order_id,omitempty"`                       // 改签订单号 如：A改成B，则改签订单号为B
	PreOrderId                 *string  `json:"pre_order_id,omitempty"`                   // 改签前订单号 如：A改成B，则改签前订单号为A
	TransactionType            *int32   `json:"transaction_type,omitempty"`               // 交易类型 枚举【0：出票，1：改签，2：退票，3：线下改签，4：线下退票】
	SupplierOrderId            *string  `json:"supplier_order_id,omitempty"`              // 供应商订单号 示例：若供应商票号为\"999-9744643597\"，则供应商订单号为\"9999744643597\"
	FlightSegmentTravel        *string  `json:"flight_segment_travel,omitempty"`          // 航段
	FlightSegmentNumber        *int32   `json:"flight_segment_number,omitempty"`          // 航段数量
	BookingDate                *string  `json:"booking_date,omitempty"`                   // 预定日期 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	SettlementTime             *string  `json:"settlement_time,omitempty"`                // 完单时间 (冗余字段) 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	CreateTime                 *string  `json:"create_time,omitempty"`                    // 下单时间（冗余字段）格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	DepartureTime              *string  `json:"departure_time,omitempty"`                 // 起飞时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	PayType                    *string  `json:"pay_type,omitempty"`                       // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	CompanyRealPay             *string  `json:"company_real_pay,omitempty"`               // 企业实付(单位：元)
	TotalFee                   *string  `json:"total_fee,omitempty"`                      // 订单总金额(单位：元) 当前订单支付的总金额，订单总金额=企业实付金额+个人实付金额（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	TicketFee                  *string  `json:"ticket_fee,omitempty"`                     // 票面费用(单位：元)
	ConstructionCost           *string  `json:"construction_cost,omitempty"`              // 机建费(单位：元)
	FuelCost                   *string  `json:"fuelCost,omitempty"`                       // 燃油费(单位：元)
	UpgradeCost                *string  `json:"upgrade_cost,omitempty"`                   // 改签差价(单位：元)
	ChangeHandleCost           *string  `json:"change_handle_cost,omitempty"`             // 改签手续费(单位：元)
	RefundHandleCost           *string  `json:"refund_handle_cost,omitempty"`             // 退款手续费(单位：元)
	ServiceFee                 *string  `json:"service_fee,omitempty"`                    // 平台使用费(单位：元)
	TravelItineraryFee         *string  `json:"travel_itinerary_fee,omitempty"`           // 行程单金额(单位：元)
	InvoiceFee                 *string  `json:"invoice_fee,omitempty"`                    // 发票金额(单位：元)
	Reason                     *string  `json:"reason,omitempty"`                         // 退改签原因
	PnrNumber                  *string  `json:"pnr_number,omitempty"`                     // PNR
	AirlineSimpleName          *string  `json:"airline_simple_name,omitempty"`            // 航司简称
	FlightNumber               *string  `json:"flight_number,omitempty"`                  // 航班号
	DepartureCityName          *string  `json:"departure_city_name,omitempty"`            // 出发城市名称
	ArrivalCityName            *string  `json:"arrival_city_name,omitempty"`              // 达到城市名称
	DepartureAirportCode       *string  `json:"departure_airport_code,omitempty"`         // 出发机场三字码
	ArrivalAirportCode         *string  `json:"arrival_airport_code,omitempty"`           // 达到城市三字码
	DepartureAirportName       *string  `json:"departure_airport_name,omitempty"`         // 出发机场名称
	ArrivalAirportName         *string  `json:"arrival_airport_name,omitempty"`           // 达到机场名称
	CabinName                  *string  `json:"cabin_name,omitempty"`                     // 舱位名称
	CabinType                  *int32   `json:"cabin_type,omitempty"`                     // 舱位类型 枚举【0：未配置；1：头等舱；2：超值头等舱；3：商务舱；4：经济舱；5：经济舱精选；6：经济舱Y舱】
	CabinCode                  *string  `json:"cabin_code,omitempty"`                     // 子舱位代码
	ApprovalId                 *string  `json:"approval_id,omitempty"`                    // 出差申请单号（冗余字段）
	SubCompanyNo               *string  `json:"sub_company_no,omitempty"`                 // 公司主体代码（6位），仅滴滴账单展示
	FlightTravelType           *int32   `json:"flight_travel_type,omitempty"`             // 航程类别，枚举【0：国内大陆；1：国际+港澳台】
	ExtraInfo                  *string  `json:"extra_info,omitempty"`                     // 审批单自定义
	FinalOfflineStatus         *string  `json:"final_offline_status,omitempty"`           // 终票异常状态
	SupplierTicketNumber       *string  `json:"supplier_ticket_number,omitempty"`         // 供应商票号
	RcBook                     *string  `json:"rc_book,omitempty"`                        // 机票提前预定天数RC
	RcLevel                    *string  `json:"rc_level,omitempty"`                       // 机票舱等RC
	RcLowPrice                 *string  `json:"rc_lowPrice,omitempty"`                    // 机票低价RC
	RcTimePeriod               *string  `json:"rc_time_period,omitempty"`                 // 机票时间范围预订RC
	ApprovalHistory            *string  `json:"approval_history,omitempty"`               // 审批历史
	IsAgreement                *string  `json:"is_agreement,omitempty"`                   // 是否协议价 “是”或“否”
	Discount                   *string  `json:"discount,omitempty"`                       // 折扣 示例：”全价”、“8.8折”
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`                 // 出行目的
	TripDescription            *string  `json:"trip_description,omitempty"`               // 出行描述
	TripReason                 *string  `json:"trip_reason,omitempty"`                    // 出差事由
	AddedGoodsName             *string  `json:"added_goods_name,omitempty"`               // 手工单产品类型
	Remark                     *string  `json:"remark,omitempty"`                         // 备注
	UniqueKey                  *string  `json:"unique_key,omitempty"`                     // 唯一键 此条明细的唯一键
	LegalEntityId              *int64   `json:"legal_entity_id,omitempty"`                // 预定人开票子公司id
	LegalEntityName            *string  `json:"legal_entity_name,omitempty"`              // 预定人开票子公司名称
	PassengerLegalEntityId     *int64   `json:"passenger_legal_entity_id,omitempty"`      // 乘车人开票子公司id
	PassengerLegalEntityName   *string  `json:"passenger_legal_entity_name,omitempty"`    // 乘车人开票子公司名称
	ExInfo01                   *string  `json:"ex_info_01,omitempty"`                     // 用户自定义拓展字段1
	ExInfo02                   *string  `json:"ex_info_02,omitempty"`                     // 用户自定义拓展字段2
	ExInfo03                   *string  `json:"ex_info_03,omitempty"`                     // 用户自定义拓展字段3
	ExInfo04                   *string  `json:"ex_info_04,omitempty"`                     // 用户自定义拓展字段4
	ExInfo05                   *string  `json:"ex_info_05,omitempty"`                     // 用户自定义拓展字段5
	ExInfo06                   *string  `json:"ex_info_06,omitempty"`                     // 用户自定义拓展字段6
	ExInfo07                   *string  `json:"ex_info_07,omitempty"`                     // 用户自定义拓展字段7
	ExInfo08                   *string  `json:"ex_info_08,omitempty"`                     // 用户自定义拓展字段8
	InstitutionName            *string  `json:"institution_name,omitempty"`               // 制度名称
	BudgetCenterCode           *string  `json:"budget_center_code,omitempty"`             // 成本中心code 外部成本中心id
	OriginOrderId              *string  `json:"origin_order_id,omitempty"`                // 原订单号 如：\"9999744643597\"
	ArrivalTime                *string  `json:"arrival_time,omitempty"`                   // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	EconomyOriginalCost        *string  `json:"economy_original_cost,omitempty"`          // Y舱全价 如：\"168.80\" 保留两位小数
	UpgradeFinalStatus         *int32   `json:"upgrade_final_status,omitempty"`           // 终票状态 枚举【-1, \"起飞未到终态/预定口径\" 0, \"特殊-私退私改\" 1, \"终态-已使用\" 2, \"终态-已退票\"】
	ExceptionInfo              *string  `json:"exception_info,omitempty"`                 // 机票订单异常信息 如：\"张三-单程-北京至上海：行程冲突；李四-返程-上海至北京：行程冲突、行程重复\"
	PassengerMemberId          *string  `json:"passengerMemberId,omitempty"`              // 乘车人ID，如：123432
	BudgetCenterCode2          *string  `json:"budgetCenterCode,omitempty"`               // 成本中心编码，如：3899
	Rebook                     *string  `json:"rebook,omitempty"`                         // 退改类型 自愿/非自愿
	BookingMemberId            *string  `json:"bookingMemberId,omitempty"`                // 预订人员工ID 如：1111233
	MainFlightNumber           *string  `json:"mainFlightNumber,omitempty"`               // 共享航班号 如：CU8282
	BaseFee                    *string  `json:"baseFee,omitempty"`                        // 机票全价 如：2000
	ApprovalChangeHistory      *string  `json:"approval_change_history,omitempty"`        // 改签审批明细
	RankName                   *string  `json:"rank_name,omitempty"`                      // 职级 如：\"员工\"
	DepartureCityId            *string  `json:"departure_city_id,omitempty"`              // 出发城市id
	ArrivalCityId              *string  `json:"arrival_city_id,omitempty"`                // 到达城市id
	OutLegalEntityId           *string  `json:"out_legal_entity_id,omitempty"`            // 外部公司主体编号
	ReductionFee               *float64 `json:"reduction_fee,omitempty"`                  // 滴滴订立减 如：10.00
	PersonalRealPay            *string  `json:"personal_real_pay,omitempty"`              // 个人实付
	CompanyChangeServiceCost   *string  `json:"company_change_service_cost,omitempty"`    // 企业改签手续费
	PersonalChangeServiceCost  *string  `json:"personal_change_service_cost,omitempty"`   // 个人改签手续费
	CompanyUpgradeCost         *string  `json:"company_upgrade_cost,omitempty"`           // 企业改签差价
	PersonalUpgradeCost        *string  `json:"personal_upgrade_cost,omitempty"`          // 个人改签差价
	CompanyTicketCost          *string  `json:"company_ticket_cost,omitempty"`            // 企业票面价
	PersonalTicketCost         *string  `json:"personal_ticket_cost,omitempty"`           // 个人票面价
	InstitutionId              *int64   `json:"institution_id,omitempty"`                 // 制度ID
	LastupdateTime             *string  `json:"lastupdate_time,omitempty"`                // 最后更新时间
	RcDistancePeriod           *string  `json:"rc_distance_period,omitempty"`             // 超出机票里程范围的超标原因
	ExcludingServiceFee        *float64 `json:"excluding_service_fee,omitempty"`          // 企业实付（不包含平台使用费）
	ExtendInfo                 *string  `json:"extend_info,omitempty"`                    // 自定义（开票主体等信息)
	ApplyRefundTime            *string  `json:"apply_refund_time,omitempty"`              // 退票申请时间
	ApplyChangeTime            *string  `json:"apply_change_time,omitempty"`              // 改签申请时间
	IsDiscountTicket           *string  `json:"is_discount_ticket,omitempty"`             // 是否折扣机票
	SubAccountCompanyName      *string  `json:"sub_account_company_name,omitempty"`       // 子账户公司名称
	ExInfo01Code               *string  `json:"ex_info_01_code,omitempty"`                // 用户自定义拓展字段 1 编码
	ExInfo02Code               *string  `json:"ex_info_02_code,omitempty"`                // 用户自定义拓展字段 2 编码
	ExInfo03Code               *string  `json:"ex_info_03_code,omitempty"`                // 用户自定义拓展字段 3 编码
	ExInfo04Code               *string  `json:"ex_info_04_code,omitempty"`                // 用户自定义拓展字段 4 编码
	ExInfo05Code               *string  `json:"ex_info_05_code,omitempty"`                // 用户自定义拓展字段 5 编码
	ExInfo06Code               *string  `json:"ex_info_06_code,omitempty"`                // 用户自定义拓展字段 6 编码
	ExInfo07Code               *string  `json:"ex_info_07_code,omitempty"`                // 用户自定义拓展字段 7 编码
	ExInfo08Code               *string  `json:"ex_info_08_code,omitempty"`                // 用户自定义拓展字段 8 编码
	ParentInstiutionId         *int64   `json:"parent_instiution_id,omitempty"`           // 父制度 ID
	TicketFeeSp                *string  `json:"ticket_fee_sp,omitempty"`                  // 票面价 (单位：元) 出票时为出票的票面价；改签时为 0；当（出 - 退）退票时为原票面价票的负数，当（出 - 改 - 退）退票时为（原票面价 + 改签差价）的负数 (白名单客户字段，其他客户不提供)
	ConstructionCostSp         *string  `json:"construction_cost_sp,omitempty"`           // 机建费 (单位：元) 出票时为出票的机建费；改签时为 0；退票时为原机建费票的负数 (白名单客户字段，其他客户不提供)
	FuelCostSp                 *string  `json:"fuelCost_sp,omitempty"`                    // 燃油费 (单位：元) 出票时为出票的燃油费；改签为 0；退票为原燃油费的负数 (白名单客户字段，其他客户不提供)
	ExemptType                 *string  `json:"exempt_type,omitempty"`                    // 豁免类型
	ExemptReason               *string  `json:"exempt_reason,omitempty"`                  // 豁免原因
}

type NotGenBDOfDomesticFlightItemBuilder struct {
	originTicketId                string // 原票id
	originTicketIdSet             bool
	ticketIdText                  string // 改签票id
	ticketIdTextSet               bool
	passengerMemberName           string // 乘机人姓名
	passengerMemberNameSet        bool
	passengerMemberNumber         string // 乘机人工号
	passengerMemberNumberSet      bool
	passengerDepName              string // 乘机人部门
	passengerDepNameSet           bool
	passengerParentPathDepName    string // 乘机人所在部门 示例：\"机票测试专用751>优秀部门\"
	passengerParentPathDepNameSet bool
	passengerDepCode              string // 乘机人部门编号
	passengerDepCodeSet           bool
	settleTime                    string // 完单时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	settleTimeSet                 bool
	memberId                      string // 乘机人员工ID
	memberIdSet                   bool
	orderFixType                  int32 // 订单补偿类型 枚举【0：非补偿单，1：补偿单】 该字段滴滴内部使用，外部可不用关注该字段
	orderFixTypeSet               bool
	orderFixSettleTime            string // 补偿单的结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	orderFixSettleTimeSet         bool
	passengerName                 string // 乘机人姓名（冗余字段）
	passengerNameSet              bool
	department                    string // 乘机人部门（冗余字段）
	departmentSet                 bool
	budgetCenterName              string // 成本中心名称
	budgetCenterNameSet           bool
	budgetCenterParentPathName    string // 所在成本中心 示例：\"机票测试专用751>高优项目\"
	budgetCenterParentPathNameSet bool
	budgetCenterId                string // 成本中心ID
	budgetCenterIdSet             bool
	bookingMemberName             string // 预订人姓名
	bookingMemberNameSet          bool
	bookingMemberNumber           string // 预订人工号
	bookingMemberNumberSet        bool
	bookingDepName                string // 预订人部门
	bookingDepNameSet             bool
	bookingParentPathDepName      string // 预订人所在部门 示例：\"机票测试专用751>优秀部门\"
	bookingParentPathDepNameSet   bool
	bookingDepCode                string // 预订人部门编号
	bookingDepCodeSet             bool
	requisitionId                 string // 出差申请单号
	requisitionIdSet              bool
	outRequisitionId              string // 外部审批单ID
	outRequisitionIdSet           bool
	orderId                       string // 改签订单号 如：A改成B，则改签订单号为B
	orderIdSet                    bool
	preOrderId                    string // 改签前订单号 如：A改成B，则改签前订单号为A
	preOrderIdSet                 bool
	transactionType               int32 // 交易类型 枚举【0：出票，1：改签，2：退票，3：线下改签，4：线下退票】
	transactionTypeSet            bool
	supplierOrderId               string // 供应商订单号 示例：若供应商票号为\"999-9744643597\"，则供应商订单号为\"9999744643597\"
	supplierOrderIdSet            bool
	flightSegmentTravel           string // 航段
	flightSegmentTravelSet        bool
	flightSegmentNumber           int32 // 航段数量
	flightSegmentNumberSet        bool
	bookingDate                   string // 预定日期 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	bookingDateSet                bool
	settlementTime                string // 完单时间 (冗余字段) 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	settlementTimeSet             bool
	createTime                    string // 下单时间（冗余字段）格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	createTimeSet                 bool
	departureTime                 string // 起飞时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	departureTimeSet              bool
	payType                       string // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	payTypeSet                    bool
	companyRealPay                string // 企业实付(单位：元)
	companyRealPaySet             bool
	totalFee                      string // 订单总金额(单位：元) 当前订单支付的总金额，订单总金额=企业实付金额+个人实付金额（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	totalFeeSet                   bool
	ticketFee                     string // 票面费用(单位：元)
	ticketFeeSet                  bool
	constructionCost              string // 机建费(单位：元)
	constructionCostSet           bool
	fuelCost                      string // 燃油费(单位：元)
	fuelCostSet                   bool
	upgradeCost                   string // 改签差价(单位：元)
	upgradeCostSet                bool
	changeHandleCost              string // 改签手续费(单位：元)
	changeHandleCostSet           bool
	refundHandleCost              string // 退款手续费(单位：元)
	refundHandleCostSet           bool
	serviceFee                    string // 平台使用费(单位：元)
	serviceFeeSet                 bool
	travelItineraryFee            string // 行程单金额(单位：元)
	travelItineraryFeeSet         bool
	invoiceFee                    string // 发票金额(单位：元)
	invoiceFeeSet                 bool
	reason                        string // 退改签原因
	reasonSet                     bool
	pnrNumber                     string // PNR
	pnrNumberSet                  bool
	airlineSimpleName             string // 航司简称
	airlineSimpleNameSet          bool
	flightNumber                  string // 航班号
	flightNumberSet               bool
	departureCityName             string // 出发城市名称
	departureCityNameSet          bool
	arrivalCityName               string // 达到城市名称
	arrivalCityNameSet            bool
	departureAirportCode          string // 出发机场三字码
	departureAirportCodeSet       bool
	arrivalAirportCode            string // 达到城市三字码
	arrivalAirportCodeSet         bool
	departureAirportName          string // 出发机场名称
	departureAirportNameSet       bool
	arrivalAirportName            string // 达到机场名称
	arrivalAirportNameSet         bool
	cabinName                     string // 舱位名称
	cabinNameSet                  bool
	cabinType                     int32 // 舱位类型 枚举【0：未配置；1：头等舱；2：超值头等舱；3：商务舱；4：经济舱；5：经济舱精选；6：经济舱Y舱】
	cabinTypeSet                  bool
	cabinCode                     string // 子舱位代码
	cabinCodeSet                  bool
	approvalId                    string // 出差申请单号（冗余字段）
	approvalIdSet                 bool
	subCompanyNo                  string // 公司主体代码（6位），仅滴滴账单展示
	subCompanyNoSet               bool
	flightTravelType              int32 // 航程类别，枚举【0：国内大陆；1：国际+港澳台】
	flightTravelTypeSet           bool
	extraInfo                     string // 审批单自定义
	extraInfoSet                  bool
	finalOfflineStatus            string // 终票异常状态
	finalOfflineStatusSet         bool
	supplierTicketNumber          string // 供应商票号
	supplierTicketNumberSet       bool
	rcBook                        string // 机票提前预定天数RC
	rcBookSet                     bool
	rcLevel                       string // 机票舱等RC
	rcLevelSet                    bool
	rcLowPrice                    string // 机票低价RC
	rcLowPriceSet                 bool
	rcTimePeriod                  string // 机票时间范围预订RC
	rcTimePeriodSet               bool
	approvalHistory               string // 审批历史
	approvalHistorySet            bool
	isAgreement                   string // 是否协议价 “是”或“否”
	isAgreementSet                bool
	discount                      string // 折扣 示例：”全价”、“8.8折”
	discountSet                   bool
	travelPurpose                 string // 出行目的
	travelPurposeSet              bool
	tripDescription               string // 出行描述
	tripDescriptionSet            bool
	tripReason                    string // 出差事由
	tripReasonSet                 bool
	addedGoodsName                string // 手工单产品类型
	addedGoodsNameSet             bool
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
	budgetCenterCode              string // 成本中心code 外部成本中心id
	budgetCenterCodeSet           bool
	originOrderId                 string // 原订单号 如：\"9999744643597\"
	originOrderIdSet              bool
	arrivalTime                   string // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	arrivalTimeSet                bool
	economyOriginalCost           string // Y舱全价 如：\"168.80\" 保留两位小数
	economyOriginalCostSet        bool
	upgradeFinalStatus            int32 // 终票状态 枚举【-1, \"起飞未到终态/预定口径\" 0, \"特殊-私退私改\" 1, \"终态-已使用\" 2, \"终态-已退票\"】
	upgradeFinalStatusSet         bool
	exceptionInfo                 string // 机票订单异常信息 如：\"张三-单程-北京至上海：行程冲突；李四-返程-上海至北京：行程冲突、行程重复\"
	exceptionInfoSet              bool
	passengerMemberId             string // 乘车人ID，如：123432
	passengerMemberIdSet          bool
	budgetCenterCodeREPEAT2       string // 成本中心编码，如：3899
	budgetCenterCodeREPEAT2Set    bool
	rebook                        string // 退改类型 自愿/非自愿
	rebookSet                     bool
	bookingMemberId               string // 预订人员工ID 如：1111233
	bookingMemberIdSet            bool
	mainFlightNumber              string // 共享航班号 如：CU8282
	mainFlightNumberSet           bool
	baseFee                       string // 机票全价 如：2000
	baseFeeSet                    bool
	approvalChangeHistory         string // 改签审批明细
	approvalChangeHistorySet      bool
	rankName                      string // 职级 如：\"员工\"
	rankNameSet                   bool
	departureCityId               string // 出发城市id
	departureCityIdSet            bool
	arrivalCityId                 string // 到达城市id
	arrivalCityIdSet              bool
	outLegalEntityId              string // 外部公司主体编号
	outLegalEntityIdSet           bool
	reductionFee                  float64 // 滴滴订立减 如：10.00
	reductionFeeSet               bool
	personalRealPay               string // 个人实付
	personalRealPaySet            bool
	companyChangeServiceCost      string // 企业改签手续费
	companyChangeServiceCostSet   bool
	personalChangeServiceCost     string // 个人改签手续费
	personalChangeServiceCostSet  bool
	companyUpgradeCost            string // 企业改签差价
	companyUpgradeCostSet         bool
	personalUpgradeCost           string // 个人改签差价
	personalUpgradeCostSet        bool
	companyTicketCost             string // 企业票面价
	companyTicketCostSet          bool
	personalTicketCost            string // 个人票面价
	personalTicketCostSet         bool
	institutionId                 int64 // 制度ID
	institutionIdSet              bool
	lastupdateTime                string // 最后更新时间
	lastupdateTimeSet             bool
	rcDistancePeriod              string // 超出机票里程范围的超标原因
	rcDistancePeriodSet           bool
	excludingServiceFee           float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet        bool
	extendInfo                    string // 自定义（开票主体等信息)
	extendInfoSet                 bool
	applyRefundTime               string // 退票申请时间
	applyRefundTimeSet            bool
	applyChangeTime               string // 改签申请时间
	applyChangeTimeSet            bool
	isDiscountTicket              string // 是否折扣机票
	isDiscountTicketSet           bool
	subAccountCompanyName         string // 子账户公司名称
	subAccountCompanyNameSet      bool
	exInfo01Code                  string // 用户自定义拓展字段 1 编码
	exInfo01CodeSet               bool
	exInfo02Code                  string // 用户自定义拓展字段 2 编码
	exInfo02CodeSet               bool
	exInfo03Code                  string // 用户自定义拓展字段 3 编码
	exInfo03CodeSet               bool
	exInfo04Code                  string // 用户自定义拓展字段 4 编码
	exInfo04CodeSet               bool
	exInfo05Code                  string // 用户自定义拓展字段 5 编码
	exInfo05CodeSet               bool
	exInfo06Code                  string // 用户自定义拓展字段 6 编码
	exInfo06CodeSet               bool
	exInfo07Code                  string // 用户自定义拓展字段 7 编码
	exInfo07CodeSet               bool
	exInfo08Code                  string // 用户自定义拓展字段 8 编码
	exInfo08CodeSet               bool
	parentInstiutionId            int64 // 父制度 ID
	parentInstiutionIdSet         bool
	ticketFeeSp                   string // 票面价 (单位：元) 出票时为出票的票面价；改签时为 0；当（出 - 退）退票时为原票面价票的负数，当（出 - 改 - 退）退票时为（原票面价 + 改签差价）的负数 (白名单客户字段，其他客户不提供)
	ticketFeeSpSet                bool
	constructionCostSp            string // 机建费 (单位：元) 出票时为出票的机建费；改签时为 0；退票时为原机建费票的负数 (白名单客户字段，其他客户不提供)
	constructionCostSpSet         bool
	fuelCostSp                    string // 燃油费 (单位：元) 出票时为出票的燃油费；改签为 0；退票为原燃油费的负数 (白名单客户字段，其他客户不提供)
	fuelCostSpSet                 bool
	exemptType                    string // 豁免类型
	exemptTypeSet                 bool
	exemptReason                  string // 豁免原因
	exemptReasonSet               bool
}

func NewNotGenBDOfDomesticFlightItemBuilder() *NotGenBDOfDomesticFlightItemBuilder {
	return &NotGenBDOfDomesticFlightItemBuilder{}
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OriginTicketId(originTicketId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.originTicketId = originTicketId
	builder.originTicketIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TicketIdText(ticketIdText string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.ticketIdText = ticketIdText
	builder.ticketIdTextSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerMemberName(passengerMemberName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerDepName(passengerDepName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerParentPathDepName(passengerParentPathDepName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerParentPathDepName = passengerParentPathDepName
	builder.passengerParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerDepCode(passengerDepCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SettleTime(settleTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.settleTime = settleTime
	builder.settleTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) MemberId(memberId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OrderFixType(orderFixType int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.orderFixType = orderFixType
	builder.orderFixTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OrderFixSettleTime(orderFixSettleTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.orderFixSettleTime = orderFixSettleTime
	builder.orderFixSettleTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerName(passengerName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) Department(department string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BudgetCenterName(budgetCenterName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BudgetCenterId(budgetCenterId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingMemberName(bookingMemberName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingMemberNumber(bookingMemberNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingDepName(bookingDepName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingDepCode(bookingDepCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RequisitionId(requisitionId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OutRequisitionId(outRequisitionId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OrderId(orderId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PreOrderId(preOrderId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.preOrderId = preOrderId
	builder.preOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TransactionType(transactionType int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.transactionType = transactionType
	builder.transactionTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SupplierOrderId(supplierOrderId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.supplierOrderId = supplierOrderId
	builder.supplierOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FlightSegmentTravel(flightSegmentTravel string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.flightSegmentTravel = flightSegmentTravel
	builder.flightSegmentTravelSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FlightSegmentNumber(flightSegmentNumber int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.flightSegmentNumber = flightSegmentNumber
	builder.flightSegmentNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingDate(bookingDate string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingDate = bookingDate
	builder.bookingDateSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SettlementTime(settlementTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.settlementTime = settlementTime
	builder.settlementTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CreateTime(createTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) DepartureTime(departureTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PayType(payType string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CompanyRealPay(companyRealPay string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TotalFee(totalFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.totalFee = totalFee
	builder.totalFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TicketFee(ticketFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.ticketFee = ticketFee
	builder.ticketFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ConstructionCost(constructionCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.constructionCost = constructionCost
	builder.constructionCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FuelCost(fuelCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.fuelCost = fuelCost
	builder.fuelCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) UpgradeCost(upgradeCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.upgradeCost = upgradeCost
	builder.upgradeCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ChangeHandleCost(changeHandleCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.changeHandleCost = changeHandleCost
	builder.changeHandleCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RefundHandleCost(refundHandleCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.refundHandleCost = refundHandleCost
	builder.refundHandleCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ServiceFee(serviceFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TravelItineraryFee(travelItineraryFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.travelItineraryFee = travelItineraryFee
	builder.travelItineraryFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) InvoiceFee(invoiceFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.invoiceFee = invoiceFee
	builder.invoiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) Reason(reason string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PnrNumber(pnrNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.pnrNumber = pnrNumber
	builder.pnrNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) AirlineSimpleName(airlineSimpleName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.airlineSimpleName = airlineSimpleName
	builder.airlineSimpleNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FlightNumber(flightNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) DepartureCityName(departureCityName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ArrivalCityName(arrivalCityName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) DepartureAirportCode(departureAirportCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.departureAirportCode = departureAirportCode
	builder.departureAirportCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ArrivalAirportCode(arrivalAirportCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.arrivalAirportCode = arrivalAirportCode
	builder.arrivalAirportCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) DepartureAirportName(departureAirportName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.departureAirportName = departureAirportName
	builder.departureAirportNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ArrivalAirportName(arrivalAirportName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.arrivalAirportName = arrivalAirportName
	builder.arrivalAirportNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CabinName(cabinName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.cabinName = cabinName
	builder.cabinNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CabinType(cabinType int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.cabinType = cabinType
	builder.cabinTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CabinCode(cabinCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.cabinCode = cabinCode
	builder.cabinCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ApprovalId(approvalId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SubCompanyNo(subCompanyNo string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.subCompanyNo = subCompanyNo
	builder.subCompanyNoSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FlightTravelType(flightTravelType int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.flightTravelType = flightTravelType
	builder.flightTravelTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExtraInfo(extraInfo string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FinalOfflineStatus(finalOfflineStatus string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.finalOfflineStatus = finalOfflineStatus
	builder.finalOfflineStatusSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SupplierTicketNumber(supplierTicketNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.supplierTicketNumber = supplierTicketNumber
	builder.supplierTicketNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RcBook(rcBook string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rcBook = rcBook
	builder.rcBookSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RcLevel(rcLevel string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rcLevel = rcLevel
	builder.rcLevelSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RcLowPrice(rcLowPrice string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rcLowPrice = rcLowPrice
	builder.rcLowPriceSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RcTimePeriod(rcTimePeriod string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rcTimePeriod = rcTimePeriod
	builder.rcTimePeriodSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ApprovalHistory(approvalHistory string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.approvalHistory = approvalHistory
	builder.approvalHistorySet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) IsAgreement(isAgreement string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.isAgreement = isAgreement
	builder.isAgreementSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) Discount(discount string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.discount = discount
	builder.discountSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TravelPurpose(travelPurpose string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TripDescription(tripDescription string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.tripDescription = tripDescription
	builder.tripDescriptionSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TripReason(tripReason string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) AddedGoodsName(addedGoodsName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.addedGoodsName = addedGoodsName
	builder.addedGoodsNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) Remark(remark string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) UniqueKey(uniqueKey string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) LegalEntityId(legalEntityId int64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) LegalEntityName(legalEntityName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo01(exInfo01 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo02(exInfo02 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo03(exInfo03 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo04(exInfo04 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo05(exInfo05 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo06(exInfo06 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo07(exInfo07 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo08(exInfo08 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) InstitutionName(institutionName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BudgetCenterCode(budgetCenterCode string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OriginOrderId(originOrderId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.originOrderId = originOrderId
	builder.originOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ArrivalTime(arrivalTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.arrivalTime = arrivalTime
	builder.arrivalTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) EconomyOriginalCost(economyOriginalCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.economyOriginalCost = economyOriginalCost
	builder.economyOriginalCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) UpgradeFinalStatus(upgradeFinalStatus int32) *NotGenBDOfDomesticFlightItemBuilder {
	builder.upgradeFinalStatus = upgradeFinalStatus
	builder.upgradeFinalStatusSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExceptionInfo(exceptionInfo string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exceptionInfo = exceptionInfo
	builder.exceptionInfoSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PassengerMemberId(passengerMemberId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BudgetCenterCode2(budgetCenterCodeREPEAT2 string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.budgetCenterCodeREPEAT2 = budgetCenterCodeREPEAT2
	builder.budgetCenterCodeREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) Rebook(rebook string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rebook = rebook
	builder.rebookSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BookingMemberId(bookingMemberId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) MainFlightNumber(mainFlightNumber string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.mainFlightNumber = mainFlightNumber
	builder.mainFlightNumberSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) BaseFee(baseFee string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.baseFee = baseFee
	builder.baseFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ApprovalChangeHistory(approvalChangeHistory string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.approvalChangeHistory = approvalChangeHistory
	builder.approvalChangeHistorySet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RankName(rankName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) DepartureCityId(departureCityId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ArrivalCityId(arrivalCityId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) OutLegalEntityId(outLegalEntityId string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ReductionFee(reductionFee float64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.reductionFee = reductionFee
	builder.reductionFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PersonalRealPay(personalRealPay string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CompanyChangeServiceCost(companyChangeServiceCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.companyChangeServiceCost = companyChangeServiceCost
	builder.companyChangeServiceCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PersonalChangeServiceCost(personalChangeServiceCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.personalChangeServiceCost = personalChangeServiceCost
	builder.personalChangeServiceCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CompanyUpgradeCost(companyUpgradeCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.companyUpgradeCost = companyUpgradeCost
	builder.companyUpgradeCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PersonalUpgradeCost(personalUpgradeCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.personalUpgradeCost = personalUpgradeCost
	builder.personalUpgradeCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) CompanyTicketCost(companyTicketCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.companyTicketCost = companyTicketCost
	builder.companyTicketCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) PersonalTicketCost(personalTicketCost string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.personalTicketCost = personalTicketCost
	builder.personalTicketCostSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) InstitutionId(institutionId int64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) LastupdateTime(lastupdateTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.lastupdateTime = lastupdateTime
	builder.lastupdateTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) RcDistancePeriod(rcDistancePeriod string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.rcDistancePeriod = rcDistancePeriod
	builder.rcDistancePeriodSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExtendInfo(extendInfo string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ApplyRefundTime(applyRefundTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.applyRefundTime = applyRefundTime
	builder.applyRefundTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ApplyChangeTime(applyChangeTime string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.applyChangeTime = applyChangeTime
	builder.applyChangeTimeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) IsDiscountTicket(isDiscountTicket string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.isDiscountTicket = isDiscountTicket
	builder.isDiscountTicketSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo01Code(exInfo01Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo02Code(exInfo02Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo03Code(exInfo03Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo04Code(exInfo04Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo05Code(exInfo05Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo06Code(exInfo06Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo07Code(exInfo07Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExInfo08Code(exInfo08Code string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ParentInstiutionId(parentInstiutionId int64) *NotGenBDOfDomesticFlightItemBuilder {
	builder.parentInstiutionId = parentInstiutionId
	builder.parentInstiutionIdSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) TicketFeeSp(ticketFeeSp string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.ticketFeeSp = ticketFeeSp
	builder.ticketFeeSpSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ConstructionCostSp(constructionCostSp string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.constructionCostSp = constructionCostSp
	builder.constructionCostSpSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) FuelCostSp(fuelCostSp string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.fuelCostSp = fuelCostSp
	builder.fuelCostSpSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExemptType(exemptType string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exemptType = exemptType
	builder.exemptTypeSet = true
	return builder
}
func (builder *NotGenBDOfDomesticFlightItemBuilder) ExemptReason(exemptReason string) *NotGenBDOfDomesticFlightItemBuilder {
	builder.exemptReason = exemptReason
	builder.exemptReasonSet = true
	return builder
}

func (builder *NotGenBDOfDomesticFlightItemBuilder) Build() *NotGenBDOfDomesticFlightItem {
	data := &NotGenBDOfDomesticFlightItem{}
	if builder.originTicketIdSet {
		data.OriginTicketId = &builder.originTicketId
	}
	if builder.ticketIdTextSet {
		data.TicketIdText = &builder.ticketIdText
	}
	if builder.passengerMemberNameSet {
		data.PassengerMemberName = &builder.passengerMemberName
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
	}
	if builder.passengerDepNameSet {
		data.PassengerDepName = &builder.passengerDepName
	}
	if builder.passengerParentPathDepNameSet {
		data.PassengerParentPathDepName = &builder.passengerParentPathDepName
	}
	if builder.passengerDepCodeSet {
		data.PassengerDepCode = &builder.passengerDepCode
	}
	if builder.settleTimeSet {
		data.SettleTime = &builder.settleTime
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.orderFixTypeSet {
		data.OrderFixType = &builder.orderFixType
	}
	if builder.orderFixSettleTimeSet {
		data.OrderFixSettleTime = &builder.orderFixSettleTime
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterParentPathNameSet {
		data.BudgetCenterParentPathName = &builder.budgetCenterParentPathName
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.bookingMemberNameSet {
		data.BookingMemberName = &builder.bookingMemberName
	}
	if builder.bookingMemberNumberSet {
		data.BookingMemberNumber = &builder.bookingMemberNumber
	}
	if builder.bookingDepNameSet {
		data.BookingDepName = &builder.bookingDepName
	}
	if builder.bookingParentPathDepNameSet {
		data.BookingParentPathDepName = &builder.bookingParentPathDepName
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.requisitionIdSet {
		data.RequisitionId = &builder.requisitionId
	}
	if builder.outRequisitionIdSet {
		data.OutRequisitionId = &builder.outRequisitionId
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.preOrderIdSet {
		data.PreOrderId = &builder.preOrderId
	}
	if builder.transactionTypeSet {
		data.TransactionType = &builder.transactionType
	}
	if builder.supplierOrderIdSet {
		data.SupplierOrderId = &builder.supplierOrderId
	}
	if builder.flightSegmentTravelSet {
		data.FlightSegmentTravel = &builder.flightSegmentTravel
	}
	if builder.flightSegmentNumberSet {
		data.FlightSegmentNumber = &builder.flightSegmentNumber
	}
	if builder.bookingDateSet {
		data.BookingDate = &builder.bookingDate
	}
	if builder.settlementTimeSet {
		data.SettlementTime = &builder.settlementTime
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.totalFeeSet {
		data.TotalFee = &builder.totalFee
	}
	if builder.ticketFeeSet {
		data.TicketFee = &builder.ticketFee
	}
	if builder.constructionCostSet {
		data.ConstructionCost = &builder.constructionCost
	}
	if builder.fuelCostSet {
		data.FuelCost = &builder.fuelCost
	}
	if builder.upgradeCostSet {
		data.UpgradeCost = &builder.upgradeCost
	}
	if builder.changeHandleCostSet {
		data.ChangeHandleCost = &builder.changeHandleCost
	}
	if builder.refundHandleCostSet {
		data.RefundHandleCost = &builder.refundHandleCost
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.travelItineraryFeeSet {
		data.TravelItineraryFee = &builder.travelItineraryFee
	}
	if builder.invoiceFeeSet {
		data.InvoiceFee = &builder.invoiceFee
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.pnrNumberSet {
		data.PnrNumber = &builder.pnrNumber
	}
	if builder.airlineSimpleNameSet {
		data.AirlineSimpleName = &builder.airlineSimpleName
	}
	if builder.flightNumberSet {
		data.FlightNumber = &builder.flightNumber
	}
	if builder.departureCityNameSet {
		data.DepartureCityName = &builder.departureCityName
	}
	if builder.arrivalCityNameSet {
		data.ArrivalCityName = &builder.arrivalCityName
	}
	if builder.departureAirportCodeSet {
		data.DepartureAirportCode = &builder.departureAirportCode
	}
	if builder.arrivalAirportCodeSet {
		data.ArrivalAirportCode = &builder.arrivalAirportCode
	}
	if builder.departureAirportNameSet {
		data.DepartureAirportName = &builder.departureAirportName
	}
	if builder.arrivalAirportNameSet {
		data.ArrivalAirportName = &builder.arrivalAirportName
	}
	if builder.cabinNameSet {
		data.CabinName = &builder.cabinName
	}
	if builder.cabinTypeSet {
		data.CabinType = &builder.cabinType
	}
	if builder.cabinCodeSet {
		data.CabinCode = &builder.cabinCode
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.subCompanyNoSet {
		data.SubCompanyNo = &builder.subCompanyNo
	}
	if builder.flightTravelTypeSet {
		data.FlightTravelType = &builder.flightTravelType
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.finalOfflineStatusSet {
		data.FinalOfflineStatus = &builder.finalOfflineStatus
	}
	if builder.supplierTicketNumberSet {
		data.SupplierTicketNumber = &builder.supplierTicketNumber
	}
	if builder.rcBookSet {
		data.RcBook = &builder.rcBook
	}
	if builder.rcLevelSet {
		data.RcLevel = &builder.rcLevel
	}
	if builder.rcLowPriceSet {
		data.RcLowPrice = &builder.rcLowPrice
	}
	if builder.rcTimePeriodSet {
		data.RcTimePeriod = &builder.rcTimePeriod
	}
	if builder.approvalHistorySet {
		data.ApprovalHistory = &builder.approvalHistory
	}
	if builder.isAgreementSet {
		data.IsAgreement = &builder.isAgreement
	}
	if builder.discountSet {
		data.Discount = &builder.discount
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
	if builder.addedGoodsNameSet {
		data.AddedGoodsName = &builder.addedGoodsName
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
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.originOrderIdSet {
		data.OriginOrderId = &builder.originOrderId
	}
	if builder.arrivalTimeSet {
		data.ArrivalTime = &builder.arrivalTime
	}
	if builder.economyOriginalCostSet {
		data.EconomyOriginalCost = &builder.economyOriginalCost
	}
	if builder.upgradeFinalStatusSet {
		data.UpgradeFinalStatus = &builder.upgradeFinalStatus
	}
	if builder.exceptionInfoSet {
		data.ExceptionInfo = &builder.exceptionInfo
	}
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	if builder.budgetCenterCodeREPEAT2Set {
		data.BudgetCenterCode2 = &builder.budgetCenterCodeREPEAT2
	}
	if builder.rebookSet {
		data.Rebook = &builder.rebook
	}
	if builder.bookingMemberIdSet {
		data.BookingMemberId = &builder.bookingMemberId
	}
	if builder.mainFlightNumberSet {
		data.MainFlightNumber = &builder.mainFlightNumber
	}
	if builder.baseFeeSet {
		data.BaseFee = &builder.baseFee
	}
	if builder.approvalChangeHistorySet {
		data.ApprovalChangeHistory = &builder.approvalChangeHistory
	}
	if builder.rankNameSet {
		data.RankName = &builder.rankName
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.reductionFeeSet {
		data.ReductionFee = &builder.reductionFee
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.companyChangeServiceCostSet {
		data.CompanyChangeServiceCost = &builder.companyChangeServiceCost
	}
	if builder.personalChangeServiceCostSet {
		data.PersonalChangeServiceCost = &builder.personalChangeServiceCost
	}
	if builder.companyUpgradeCostSet {
		data.CompanyUpgradeCost = &builder.companyUpgradeCost
	}
	if builder.personalUpgradeCostSet {
		data.PersonalUpgradeCost = &builder.personalUpgradeCost
	}
	if builder.companyTicketCostSet {
		data.CompanyTicketCost = &builder.companyTicketCost
	}
	if builder.personalTicketCostSet {
		data.PersonalTicketCost = &builder.personalTicketCost
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.lastupdateTimeSet {
		data.LastupdateTime = &builder.lastupdateTime
	}
	if builder.rcDistancePeriodSet {
		data.RcDistancePeriod = &builder.rcDistancePeriod
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
	if builder.isDiscountTicketSet {
		data.IsDiscountTicket = &builder.isDiscountTicket
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
	if builder.ticketFeeSpSet {
		data.TicketFeeSp = &builder.ticketFeeSp
	}
	if builder.constructionCostSpSet {
		data.ConstructionCostSp = &builder.constructionCostSp
	}
	if builder.fuelCostSpSet {
		data.FuelCostSp = &builder.fuelCostSp
	}
	if builder.exemptTypeSet {
		data.ExemptType = &builder.exemptType
	}
	if builder.exemptReasonSet {
		data.ExemptReason = &builder.exemptReason
	}
	return data
}
