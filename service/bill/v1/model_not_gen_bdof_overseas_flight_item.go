package v1

// NotGenBDOfOverseasFlightItem 未出账单 - 海外机票,参考内部文档进行定义；
type NotGenBDOfOverseasFlightItem struct {
	OriginTicketId             *string  `json:"origin_ticket_id,omitempty"`               // 原票id 滴滴企业版app生成的票id
	TicketIdText               *string  `json:"ticket_id_text,omitempty"`                 // 改签票id 滴滴企业版app生成的票id，改签票对应的票id
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`        // 当前机票的乘机人工号
	PassengerDepName           *string  `json:"passenger_dep_name,omitempty"`             // 当前机票的乘机人部门
	PassengerParentPathDepName *string  `json:"passenger_parent_path_dep_name,omitempty"` // 乘机人所在部门（部门路径） 示例：\"机票测试专用751>优秀部门\"
	PassengerDepCode           *string  `json:"passenger_dep_code,omitempty"`             // 当前机票的乘机人部门编号
	MemberId                   *string  `json:"member_id,omitempty"`                      // 当前机票的乘机人员工id
	OrderFixType               *int32   `json:"order_fix_type,omitempty"`                 // 订单补偿类型 枚举【0：非补偿单，1：补偿单】 该字段滴滴内部使用，外部可不用关注该字段
	OrderFixSettleTime         *string  `json:"order_fix_settle_time,omitempty"`          // 补偿单的结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	EmployeeNumber             *string  `json:"employee_number,omitempty"`                // 乘机人工号（冗余字段）
	PassengerName              *string  `json:"passenger_name,omitempty"`                 // 乘机人姓名（冗余字段）
	PassengerMemberName        *string  `json:"passenger_member_name,omitempty"`          // 乘机人中文姓名
	Department                 *string  `json:"department,omitempty"`                     // 乘机人部门（冗余字段）
	DepartmentId               *string  `json:"department_id,omitempty"`                  // 当前机票的乘机人部门id
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`             // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"` // 指当前订单员工所在成本中心（部门或项目），示例：公司>成本中心1>成本中心2
	BudgetCenterId             *string  `json:"budget_center_id,omitempty"`               // 成本中心ID 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`            // 预订机票订单的员工姓名
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`          // 预订机票订单的员工编号
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`               // 预订机票订单的员工部门
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`   // 预订机票订单的员工所在部门，示例：公司>部门1>部门2
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`               // 预订机票订单的员工部门编号（客户在ES后台输入的部门编号）
	RequisitionId              *string  `json:"requisition_id,omitempty"`                 // 出差申请单号
	OutRequisitionId           *string  `json:"out_requisition_id,omitempty"`             // 外部审批单ID 开放平台API客户接差旅场景的外部审批单号，差旅API客户通过此字段可以和滴滴订单关联，方便查询和对账
	OrderId                    *string  `json:"order_id,omitempty"`                       // 改签订单号 如：A改成B，则改签订单号为B
	OriginOrderId              *string  `json:"origin_order_id,omitempty"`                // 原始订单号 如：A改B改C，则原始订单号为A
	PreOrderId                 *string  `json:"pre_order_id,omitempty"`                   // 改签前订单号 如：A改成B，则改签前订单号为A
	TransactionType            *int32   `json:"transaction_type,omitempty"`               // 交易类型 枚举【0：出票，1：改签，2：退票，3：线下改签，4：线下退票】
	SupplierOrderId            *string  `json:"supplier_order_id,omitempty"`              // 供应商订单号 示例：若供应商票号为\"999-9744643597\"，则供应商订单号为\"9999744643597\"
	FlightSegmentTravel        *string  `json:"flight_segment_travel,omitempty"`          // 航段 当前机票的航段，包括出发地-目的地 示例：\"上海-北京\"
	FlightSegmentNumber        *int32   `json:"flight_segment_number,omitempty"`          // 当前机票的航段计数
	BookingDate                *string  `json:"booking_date,omitempty"`                   // 机票预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	SettlementTime             *string  `json:"settlement_time,omitempty"`                // 结算时间 国际机票结算入账的时间，以预订时间和退票完成时间（issue_time或refundTime，目前账单API无此字段）为准；若发生调整账单，则为账单调整时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	CreateTime                 *string  `json:"create_time,omitempty"`                    // 机票预订时间（冗余字段）格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	DepartureTime              *string  `json:"departure_time,omitempty"`                 // 起飞时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	FinalDepartureTime         *string  `json:"final_departure_time,omitempty"`           // 终票起飞时间或退款时间(滴滴专用) 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	CompanyRealPay             *string  `json:"company_real_pay,omitempty"`               // 企业实付(单位：元) 当前订单通过企业账户支付金额部分
	TotalFee                   *string  `json:"total_fee,omitempty"`                      // 订单总金额(单位：元) 当前订单支付的总金额，订单总金额=企业实付金额+个人实付金额（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	TicketFee                  *string  `json:"ticket_fee,omitempty"`                     // 票面价(单位：元) 机票中的票面价格
	ConstructionCost           *string  `json:"construction_cost,omitempty"`              // 机建费(单位：元) 机票中的机建费价格
	FuelCost                   *string  `json:"fuelCost,omitempty"`                       // 燃油费(单位：元) 机票中的燃油费价格
	UpgradeCost                *string  `json:"upgrade_cost,omitempty"`                   // 改签差价(单位：元) 机票改签产生的差价费用
	ChangeHandleCost           *string  `json:"change_handle_cost,omitempty"`             // 改签手续费(单位：元) 机票改签产生的手续费
	RefundHandleCost           *string  `json:"refund_handle_cost,omitempty"`             // 退款手续费(单位：元) 机票退票产生的退票手续费
	ServiceFee                 *string  `json:"service_fee,omitempty"`                    // 平台使用费(单位：元)
	Tax                        *string  `json:"tax,omitempty"`                            // 购买机票过程中所产生的税费
	Reason                     *string  `json:"reason,omitempty"`                         // 退改签原因
	PnrNumber                  *string  `json:"pnr_number,omitempty"`                     // PNR 机票发生的旅客订座记录
	AirlineSimpleName          *string  `json:"airline_simple_name,omitempty"`            // 航司简称
	FlightNumber               *string  `json:"flight_number,omitempty"`                  // 航班号
	DepartureCityName          *string  `json:"departure_city_name,omitempty"`            // 出发城市名称
	ArrivalCityName            *string  `json:"arrival_city_name,omitempty"`              // 达到城市名称
	DepartureAirportCode       *string  `json:"departure_airport_code,omitempty"`         // 出发机场三字码，如PEK
	ArrivalAirportCode         *string  `json:"arrival_airport_code,omitempty"`           // 达到城市三字码，如PEK
	DepartureAirportName       *string  `json:"departure_airport_name,omitempty"`         // 出发机场名称，如北京首都国际机场
	ArrivalAirportName         *string  `json:"arrival_airport_name,omitempty"`           // 到达机场名称，如北京首都国际机场
	CabinName                  *string  `json:"cabin_name,omitempty"`                     // 舱位名称，如经济舱
	CabinType                  *int32   `json:"cabin_type,omitempty"`                     // 舱位类型 枚举【0：未配置；1：头等舱；2：超值头等舱；3：商务舱；4：经济舱；5：经济舱精选；6：经济舱Y舱】
	CabinCode                  *string  `json:"cabin_code,omitempty"`                     // 子舱位代码
	ApprovalId                 *string  `json:"approval_id,omitempty"`                    // 出差申请单号（冗余字段）
	SubCompanyNo               *string  `json:"sub_company_no,omitempty"`                 // 公司主体代码（6位），仅滴滴账单展示
	FlightTravelType           *int32   `json:"flight_travel_type,omitempty"`             // 航程类别，枚举【0：国内大陆；1：国际+港澳台】
	ExtraInfo                  *string  `json:"extra_info,omitempty"`                     // 审批单自定义
	SupplierTicketNumber       *string  `json:"supplier_ticket_number,omitempty"`         // 供应商票号 示例：\"999-9744643597\"
	RcBook                     *string  `json:"rc_book,omitempty"`                        // 机票提前预定的超标原因
	RcLevel                    *string  `json:"rc_level,omitempty"`                       // 未按照规定仓位预定的超标原因
	RcLowPrice                 *string  `json:"rc_lowPrice,omitempty"`                    // 未按照机票最低价预定的超标原因
	RcTimePeriod               *string  `json:"rc_time_period,omitempty"`                 // 超出机票预订时间范围的超标原因
	ApprovalHistory            *string  `json:"approval_history,omitempty"`               // 审批历史记录，包括审批时间、审批人姓名、手机号、审批结果、不合规备注
	IsAgreement                *string  `json:"is_agreement,omitempty"`                   // 是否协议价 “是”或“否” 机票是否协议价预订
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`                 // 出行目的 出行人下单填写的出行目的
	TripDescription            *string  `json:"trip_description,omitempty"`               // 出行描述
	TripReason                 *string  `json:"trip_reason,omitempty"`                    // 出差事由
	AddedGoodsName             *string  `json:"added_goods_name,omitempty"`               // 手工单产品类型 线下预订
	UniqueKey                  *string  `json:"unique_key,omitempty"`                     // 唯一键 此条明细的唯一键
	IsAdded                    *int32   `json:"is_added,omitempty"`                       // 是否为手工单机票 枚举【0/null：正常机票 1：手工单机票】
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
	BeforeCutServiceFee        *float64 `json:"before_cut_service_fee,omitempty"`         // 手工单：立减前平台使用费 正常单：原平台使用费 （单位：元）
	InstitutionName            *string  `json:"institution_name,omitempty"`               // 制度名称
	BudgetCenterCode           *string  `json:"budget_center_code,omitempty"`             // 成本中心code 外部成本中心id
	ArrivalTime                *string  `json:"arrival_time,omitempty"`                   // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	PassengerMemberId          *string  `json:"passengerMemberId,omitempty"`              // 乘机人ID，如：123432
	BudgetCenterCode2          *string  `json:"budgetCenterCode,omitempty"`               // 成本中心编码，如：3899
	Rebook                     *string  `json:"rebook,omitempty"`                         // 退改类型 自愿/非自愿
	BookingMemberId            *string  `json:"bookingMemberId,omitempty"`                // 预订人员工ID 如：1111233
	ArrivalTime2               *string  `json:"arrivalTime,omitempty"`                    // 计划到达时间 如：2023-12-05 12:12:12
	ProjectExtInfo             *string  `json:"project_ext_info,omitempty"`               // 项目自定义
	ApprovalNormalHistory      *string  `json:"approval_normal_history,omitempty"`        // 正常差旅申请的审批历史
	ApprovalChangeHistory      *string  `json:"approval_change_history,omitempty"`        // 改签审批明细
	RankName                   *string  `json:"rank_name,omitempty"`                      // 职级 如：\"员工\"
	DepartureCityId            *string  `json:"departure_city_id,omitempty"`              // 出发城市id
	ArrivalCityId              *string  `json:"arrival_city_id,omitempty"`                // 到达城市id
	OutLegalEntityId           *string  `json:"out_legal_entity_id,omitempty"`            // 外部公司主体编号
	DepartureCountry           *string  `json:"departure_country,omitempty"`              // 出发国家名称
	ArrivalCountry             *string  `json:"arrival_country,omitempty"`                // 到达国家名称
	DepartureContinent         *string  `json:"departure_continent,omitempty"`            // 出发大洲名称
	ArrivalContinent           *string  `json:"arrival_continent,omitempty"`              // 到达大洲名称
	InstitutionId              *int64   `json:"institution_id,omitempty"`                 // 制度 ID
	ExcludingServiceFee        *float64 `json:"excluding_service_fee,omitempty"`          // 企业实付（不包含平台使用费）
	ExtendInfo                 *string  `json:"extend_info,omitempty"`                    // 自定义（开票主体等信息)
	ApplyRefundTime            *string  `json:"apply_refund_time,omitempty"`              // 退票申请时间
	ApplyChangeTime            *string  `json:"apply_change_time,omitempty"`              // 改签申请时间
	SubAccountCompanyName      *string  `json:"sub_account_company_name,omitempty"`       // 子账户公司名称
	ExInfo01Code               *string  `json:"ex_info_01_code,omitempty"`                // 用户自定义拓展字段 1 编码
	ExInfo02Code               *string  `json:"ex_info_02_code,omitempty"`                // 用户自定义拓展字段 2 编码
	ExInfo03Code               *string  `json:"ex_info_03_code,omitempty"`                // 用户自定义拓展字段 3 编码
	ExInfo04Code               *string  `json:"ex_info_04_code,omitempty"`                // 用户自定义拓展字段 4 编码
	ExInfo05Code               *string  `json:"ex_info_05_code,omitempty"`                // 用户自定义拓展字段 5 编码
	ExInfo06Code               *string  `json:"ex_info_06_code,omitempty"`                // 用户自定义拓展字段 6 编码
	ExInfo07Code               *string  `json:"ex_info_07_code,omitempty"`                // 用户自定义拓展字段 7 编码
	ExInfo08Code               *string  `json:"ex_info_08_code,omitempty"`                // 用户自定义拓展字段 8 编码
	ParentInstitutionId        *int64   `json:"parent_institution_id,omitempty"`          // 父制度 ID
}

type NotGenBDOfOverseasFlightItemBuilder struct {
	originTicketId                string // 原票id 滴滴企业版app生成的票id
	originTicketIdSet             bool
	ticketIdText                  string // 改签票id 滴滴企业版app生成的票id，改签票对应的票id
	ticketIdTextSet               bool
	passengerMemberNumber         string // 当前机票的乘机人工号
	passengerMemberNumberSet      bool
	passengerDepName              string // 当前机票的乘机人部门
	passengerDepNameSet           bool
	passengerParentPathDepName    string // 乘机人所在部门（部门路径） 示例：\"机票测试专用751>优秀部门\"
	passengerParentPathDepNameSet bool
	passengerDepCode              string // 当前机票的乘机人部门编号
	passengerDepCodeSet           bool
	memberId                      string // 当前机票的乘机人员工id
	memberIdSet                   bool
	orderFixType                  int32 // 订单补偿类型 枚举【0：非补偿单，1：补偿单】 该字段滴滴内部使用，外部可不用关注该字段
	orderFixTypeSet               bool
	orderFixSettleTime            string // 补偿单的结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\"
	orderFixSettleTimeSet         bool
	employeeNumber                string // 乘机人工号（冗余字段）
	employeeNumberSet             bool
	passengerName                 string // 乘机人姓名（冗余字段）
	passengerNameSet              bool
	passengerMemberName           string // 乘机人中文姓名
	passengerMemberNameSet        bool
	department                    string // 乘机人部门（冗余字段）
	departmentSet                 bool
	departmentId                  string // 当前机票的乘机人部门id
	departmentIdSet               bool
	budgetCenterName              string // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	budgetCenterNameSet           bool
	budgetCenterParentPathName    string // 指当前订单员工所在成本中心（部门或项目），示例：公司>成本中心1>成本中心2
	budgetCenterParentPathNameSet bool
	budgetCenterId                string // 成本中心ID 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	budgetCenterIdSet             bool
	bookingMemberName             string // 预订机票订单的员工姓名
	bookingMemberNameSet          bool
	bookingMemberNumber           string // 预订机票订单的员工编号
	bookingMemberNumberSet        bool
	bookingDepName                string // 预订机票订单的员工部门
	bookingDepNameSet             bool
	bookingParentPathDepName      string // 预订机票订单的员工所在部门，示例：公司>部门1>部门2
	bookingParentPathDepNameSet   bool
	bookingDepCode                string // 预订机票订单的员工部门编号（客户在ES后台输入的部门编号）
	bookingDepCodeSet             bool
	requisitionId                 string // 出差申请单号
	requisitionIdSet              bool
	outRequisitionId              string // 外部审批单ID 开放平台API客户接差旅场景的外部审批单号，差旅API客户通过此字段可以和滴滴订单关联，方便查询和对账
	outRequisitionIdSet           bool
	orderId                       string // 改签订单号 如：A改成B，则改签订单号为B
	orderIdSet                    bool
	originOrderId                 string // 原始订单号 如：A改B改C，则原始订单号为A
	originOrderIdSet              bool
	preOrderId                    string // 改签前订单号 如：A改成B，则改签前订单号为A
	preOrderIdSet                 bool
	transactionType               int32 // 交易类型 枚举【0：出票，1：改签，2：退票，3：线下改签，4：线下退票】
	transactionTypeSet            bool
	supplierOrderId               string // 供应商订单号 示例：若供应商票号为\"999-9744643597\"，则供应商订单号为\"9999744643597\"
	supplierOrderIdSet            bool
	flightSegmentTravel           string // 航段 当前机票的航段，包括出发地-目的地 示例：\"上海-北京\"
	flightSegmentTravelSet        bool
	flightSegmentNumber           int32 // 当前机票的航段计数
	flightSegmentNumberSet        bool
	bookingDate                   string // 机票预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	bookingDateSet                bool
	settlementTime                string // 结算时间 国际机票结算入账的时间，以预订时间和退票完成时间（issue_time或refundTime，目前账单API无此字段）为准；若发生调整账单，则为账单调整时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 15:58:49\"
	settlementTimeSet             bool
	createTime                    string // 机票预订时间（冗余字段）格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	createTimeSet                 bool
	departureTime                 string // 起飞时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	departureTimeSet              bool
	finalDepartureTime            string // 终票起飞时间或退款时间(滴滴专用) 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	finalDepartureTimeSet         bool
	companyRealPay                string // 企业实付(单位：元) 当前订单通过企业账户支付金额部分
	companyRealPaySet             bool
	totalFee                      string // 订单总金额(单位：元) 当前订单支付的总金额，订单总金额=企业实付金额+个人实付金额（异步收取平台使用费情况下，订单总金额不包含平台使用费，即订单总金额＜企业实付金额+个人实付金额）
	totalFeeSet                   bool
	ticketFee                     string // 票面价(单位：元) 机票中的票面价格
	ticketFeeSet                  bool
	constructionCost              string // 机建费(单位：元) 机票中的机建费价格
	constructionCostSet           bool
	fuelCost                      string // 燃油费(单位：元) 机票中的燃油费价格
	fuelCostSet                   bool
	upgradeCost                   string // 改签差价(单位：元) 机票改签产生的差价费用
	upgradeCostSet                bool
	changeHandleCost              string // 改签手续费(单位：元) 机票改签产生的手续费
	changeHandleCostSet           bool
	refundHandleCost              string // 退款手续费(单位：元) 机票退票产生的退票手续费
	refundHandleCostSet           bool
	serviceFee                    string // 平台使用费(单位：元)
	serviceFeeSet                 bool
	tax                           string // 购买机票过程中所产生的税费
	taxSet                        bool
	reason                        string // 退改签原因
	reasonSet                     bool
	pnrNumber                     string // PNR 机票发生的旅客订座记录
	pnrNumberSet                  bool
	airlineSimpleName             string // 航司简称
	airlineSimpleNameSet          bool
	flightNumber                  string // 航班号
	flightNumberSet               bool
	departureCityName             string // 出发城市名称
	departureCityNameSet          bool
	arrivalCityName               string // 达到城市名称
	arrivalCityNameSet            bool
	departureAirportCode          string // 出发机场三字码，如PEK
	departureAirportCodeSet       bool
	arrivalAirportCode            string // 达到城市三字码，如PEK
	arrivalAirportCodeSet         bool
	departureAirportName          string // 出发机场名称，如北京首都国际机场
	departureAirportNameSet       bool
	arrivalAirportName            string // 到达机场名称，如北京首都国际机场
	arrivalAirportNameSet         bool
	cabinName                     string // 舱位名称，如经济舱
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
	supplierTicketNumber          string // 供应商票号 示例：\"999-9744643597\"
	supplierTicketNumberSet       bool
	rcBook                        string // 机票提前预定的超标原因
	rcBookSet                     bool
	rcLevel                       string // 未按照规定仓位预定的超标原因
	rcLevelSet                    bool
	rcLowPrice                    string // 未按照机票最低价预定的超标原因
	rcLowPriceSet                 bool
	rcTimePeriod                  string // 超出机票预订时间范围的超标原因
	rcTimePeriodSet               bool
	approvalHistory               string // 审批历史记录，包括审批时间、审批人姓名、手机号、审批结果、不合规备注
	approvalHistorySet            bool
	isAgreement                   string // 是否协议价 “是”或“否” 机票是否协议价预订
	isAgreementSet                bool
	travelPurpose                 string // 出行目的 出行人下单填写的出行目的
	travelPurposeSet              bool
	tripDescription               string // 出行描述
	tripDescriptionSet            bool
	tripReason                    string // 出差事由
	tripReasonSet                 bool
	addedGoodsName                string // 手工单产品类型 线下预订
	addedGoodsNameSet             bool
	uniqueKey                     string // 唯一键 此条明细的唯一键
	uniqueKeySet                  bool
	isAdded                       int32 // 是否为手工单机票 枚举【0/null：正常机票 1：手工单机票】
	isAddedSet                    bool
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
	beforeCutServiceFee           float64 // 手工单：立减前平台使用费 正常单：原平台使用费 （单位：元）
	beforeCutServiceFeeSet        bool
	institutionName               string // 制度名称
	institutionNameSet            bool
	budgetCenterCode              string // 成本中心code 外部成本中心id
	budgetCenterCodeSet           bool
	arrivalTime                   string // 到达时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-16 10:58:49\"
	arrivalTimeSet                bool
	passengerMemberId             string // 乘机人ID，如：123432
	passengerMemberIdSet          bool
	budgetCenterCodeREPEAT2       string // 成本中心编码，如：3899
	budgetCenterCodeREPEAT2Set    bool
	rebook                        string // 退改类型 自愿/非自愿
	rebookSet                     bool
	bookingMemberId               string // 预订人员工ID 如：1111233
	bookingMemberIdSet            bool
	arrivalTimeREPEAT2            string // 计划到达时间 如：2023-12-05 12:12:12
	arrivalTimeREPEAT2Set         bool
	projectExtInfo                string // 项目自定义
	projectExtInfoSet             bool
	approvalNormalHistory         string // 正常差旅申请的审批历史
	approvalNormalHistorySet      bool
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
	departureCountry              string // 出发国家名称
	departureCountrySet           bool
	arrivalCountry                string // 到达国家名称
	arrivalCountrySet             bool
	departureContinent            string // 出发大洲名称
	departureContinentSet         bool
	arrivalContinent              string // 到达大洲名称
	arrivalContinentSet           bool
	institutionId                 int64 // 制度 ID
	institutionIdSet              bool
	excludingServiceFee           float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet        bool
	extendInfo                    string // 自定义（开票主体等信息)
	extendInfoSet                 bool
	applyRefundTime               string // 退票申请时间
	applyRefundTimeSet            bool
	applyChangeTime               string // 改签申请时间
	applyChangeTimeSet            bool
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
	parentInstitutionId           int64 // 父制度 ID
	parentInstitutionIdSet        bool
}

func NewNotGenBDOfOverseasFlightItemBuilder() *NotGenBDOfOverseasFlightItemBuilder {
	return &NotGenBDOfOverseasFlightItemBuilder{}
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OriginTicketId(originTicketId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.originTicketId = originTicketId
	builder.originTicketIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TicketIdText(ticketIdText string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.ticketIdText = ticketIdText
	builder.ticketIdTextSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerDepName(passengerDepName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerParentPathDepName(passengerParentPathDepName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerParentPathDepName = passengerParentPathDepName
	builder.passengerParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerDepCode(passengerDepCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) MemberId(memberId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OrderFixType(orderFixType int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.orderFixType = orderFixType
	builder.orderFixTypeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OrderFixSettleTime(orderFixSettleTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.orderFixSettleTime = orderFixSettleTime
	builder.orderFixSettleTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) EmployeeNumber(employeeNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerName(passengerName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerMemberName(passengerMemberName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) Department(department string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartmentId(departmentId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BudgetCenterName(budgetCenterName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BudgetCenterId(budgetCenterId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingMemberName(bookingMemberName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingMemberNumber(bookingMemberNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingDepName(bookingDepName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingDepCode(bookingDepCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RequisitionId(requisitionId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OutRequisitionId(outRequisitionId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OrderId(orderId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OriginOrderId(originOrderId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.originOrderId = originOrderId
	builder.originOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PreOrderId(preOrderId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.preOrderId = preOrderId
	builder.preOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TransactionType(transactionType int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.transactionType = transactionType
	builder.transactionTypeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) SupplierOrderId(supplierOrderId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.supplierOrderId = supplierOrderId
	builder.supplierOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FlightSegmentTravel(flightSegmentTravel string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.flightSegmentTravel = flightSegmentTravel
	builder.flightSegmentTravelSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FlightSegmentNumber(flightSegmentNumber int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.flightSegmentNumber = flightSegmentNumber
	builder.flightSegmentNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingDate(bookingDate string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingDate = bookingDate
	builder.bookingDateSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) SettlementTime(settlementTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.settlementTime = settlementTime
	builder.settlementTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) CreateTime(createTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureTime(departureTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FinalDepartureTime(finalDepartureTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.finalDepartureTime = finalDepartureTime
	builder.finalDepartureTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) CompanyRealPay(companyRealPay string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TotalFee(totalFee string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.totalFee = totalFee
	builder.totalFeeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TicketFee(ticketFee string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.ticketFee = ticketFee
	builder.ticketFeeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ConstructionCost(constructionCost string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.constructionCost = constructionCost
	builder.constructionCostSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FuelCost(fuelCost string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.fuelCost = fuelCost
	builder.fuelCostSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) UpgradeCost(upgradeCost string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.upgradeCost = upgradeCost
	builder.upgradeCostSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ChangeHandleCost(changeHandleCost string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.changeHandleCost = changeHandleCost
	builder.changeHandleCostSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RefundHandleCost(refundHandleCost string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.refundHandleCost = refundHandleCost
	builder.refundHandleCostSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ServiceFee(serviceFee string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) Tax(tax string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.tax = tax
	builder.taxSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) Reason(reason string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PnrNumber(pnrNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.pnrNumber = pnrNumber
	builder.pnrNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) AirlineSimpleName(airlineSimpleName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.airlineSimpleName = airlineSimpleName
	builder.airlineSimpleNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FlightNumber(flightNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureCityName(departureCityName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureCityName = departureCityName
	builder.departureCityNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalCityName(arrivalCityName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalCityName = arrivalCityName
	builder.arrivalCityNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureAirportCode(departureAirportCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureAirportCode = departureAirportCode
	builder.departureAirportCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalAirportCode(arrivalAirportCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalAirportCode = arrivalAirportCode
	builder.arrivalAirportCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureAirportName(departureAirportName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureAirportName = departureAirportName
	builder.departureAirportNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalAirportName(arrivalAirportName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalAirportName = arrivalAirportName
	builder.arrivalAirportNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) CabinName(cabinName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.cabinName = cabinName
	builder.cabinNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) CabinType(cabinType int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.cabinType = cabinType
	builder.cabinTypeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) CabinCode(cabinCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.cabinCode = cabinCode
	builder.cabinCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApprovalId(approvalId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) SubCompanyNo(subCompanyNo string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.subCompanyNo = subCompanyNo
	builder.subCompanyNoSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) FlightTravelType(flightTravelType int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.flightTravelType = flightTravelType
	builder.flightTravelTypeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExtraInfo(extraInfo string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) SupplierTicketNumber(supplierTicketNumber string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.supplierTicketNumber = supplierTicketNumber
	builder.supplierTicketNumberSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RcBook(rcBook string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rcBook = rcBook
	builder.rcBookSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RcLevel(rcLevel string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rcLevel = rcLevel
	builder.rcLevelSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RcLowPrice(rcLowPrice string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rcLowPrice = rcLowPrice
	builder.rcLowPriceSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RcTimePeriod(rcTimePeriod string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rcTimePeriod = rcTimePeriod
	builder.rcTimePeriodSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApprovalHistory(approvalHistory string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.approvalHistory = approvalHistory
	builder.approvalHistorySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) IsAgreement(isAgreement string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.isAgreement = isAgreement
	builder.isAgreementSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TravelPurpose(travelPurpose string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TripDescription(tripDescription string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.tripDescription = tripDescription
	builder.tripDescriptionSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) TripReason(tripReason string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) AddedGoodsName(addedGoodsName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.addedGoodsName = addedGoodsName
	builder.addedGoodsNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) UniqueKey(uniqueKey string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) IsAdded(isAdded int32) *NotGenBDOfOverseasFlightItemBuilder {
	builder.isAdded = isAdded
	builder.isAddedSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) LegalEntityId(legalEntityId int64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) LegalEntityName(legalEntityName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo01(exInfo01 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo02(exInfo02 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo03(exInfo03 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo04(exInfo04 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo05(exInfo05 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo06(exInfo06 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo07(exInfo07 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo08(exInfo08 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BeforeCutServiceFee(beforeCutServiceFee float64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.beforeCutServiceFee = beforeCutServiceFee
	builder.beforeCutServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) InstitutionName(institutionName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BudgetCenterCode(budgetCenterCode string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalTime(arrivalTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalTime = arrivalTime
	builder.arrivalTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) PassengerMemberId(passengerMemberId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BudgetCenterCode2(budgetCenterCodeREPEAT2 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.budgetCenterCodeREPEAT2 = budgetCenterCodeREPEAT2
	builder.budgetCenterCodeREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) Rebook(rebook string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rebook = rebook
	builder.rebookSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) BookingMemberId(bookingMemberId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalTime2(arrivalTimeREPEAT2 string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalTimeREPEAT2 = arrivalTimeREPEAT2
	builder.arrivalTimeREPEAT2Set = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ProjectExtInfo(projectExtInfo string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApprovalNormalHistory(approvalNormalHistory string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.approvalNormalHistory = approvalNormalHistory
	builder.approvalNormalHistorySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApprovalChangeHistory(approvalChangeHistory string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.approvalChangeHistory = approvalChangeHistory
	builder.approvalChangeHistorySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) RankName(rankName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureCityId(departureCityId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalCityId(arrivalCityId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) OutLegalEntityId(outLegalEntityId string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureCountry(departureCountry string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureCountry = departureCountry
	builder.departureCountrySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalCountry(arrivalCountry string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalCountry = arrivalCountry
	builder.arrivalCountrySet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) DepartureContinent(departureContinent string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.departureContinent = departureContinent
	builder.departureContinentSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ArrivalContinent(arrivalContinent string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.arrivalContinent = arrivalContinent
	builder.arrivalContinentSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) InstitutionId(institutionId int64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExtendInfo(extendInfo string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApplyRefundTime(applyRefundTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.applyRefundTime = applyRefundTime
	builder.applyRefundTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ApplyChangeTime(applyChangeTime string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.applyChangeTime = applyChangeTime
	builder.applyChangeTimeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo01Code(exInfo01Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo02Code(exInfo02Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo03Code(exInfo03Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo04Code(exInfo04Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo05Code(exInfo05Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo06Code(exInfo06Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo07Code(exInfo07Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ExInfo08Code(exInfo08Code string) *NotGenBDOfOverseasFlightItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}
func (builder *NotGenBDOfOverseasFlightItemBuilder) ParentInstitutionId(parentInstitutionId int64) *NotGenBDOfOverseasFlightItemBuilder {
	builder.parentInstitutionId = parentInstitutionId
	builder.parentInstitutionIdSet = true
	return builder
}

func (builder *NotGenBDOfOverseasFlightItemBuilder) Build() *NotGenBDOfOverseasFlightItem {
	data := &NotGenBDOfOverseasFlightItem{}
	if builder.originTicketIdSet {
		data.OriginTicketId = &builder.originTicketId
	}
	if builder.ticketIdTextSet {
		data.TicketIdText = &builder.ticketIdText
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
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.orderFixTypeSet {
		data.OrderFixType = &builder.orderFixType
	}
	if builder.orderFixSettleTimeSet {
		data.OrderFixSettleTime = &builder.orderFixSettleTime
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.passengerNameSet {
		data.PassengerName = &builder.passengerName
	}
	if builder.passengerMemberNameSet {
		data.PassengerMemberName = &builder.passengerMemberName
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.departmentIdSet {
		data.DepartmentId = &builder.departmentId
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
	if builder.originOrderIdSet {
		data.OriginOrderId = &builder.originOrderId
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
	if builder.finalDepartureTimeSet {
		data.FinalDepartureTime = &builder.finalDepartureTime
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
	if builder.taxSet {
		data.Tax = &builder.tax
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
	if builder.uniqueKeySet {
		data.UniqueKey = &builder.uniqueKey
	}
	if builder.isAddedSet {
		data.IsAdded = &builder.isAdded
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
	if builder.beforeCutServiceFeeSet {
		data.BeforeCutServiceFee = &builder.beforeCutServiceFee
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.arrivalTimeSet {
		data.ArrivalTime = &builder.arrivalTime
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
	if builder.arrivalTimeREPEAT2Set {
		data.ArrivalTime2 = &builder.arrivalTimeREPEAT2
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.approvalNormalHistorySet {
		data.ApprovalNormalHistory = &builder.approvalNormalHistory
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
	if builder.departureCountrySet {
		data.DepartureCountry = &builder.departureCountry
	}
	if builder.arrivalCountrySet {
		data.ArrivalCountry = &builder.arrivalCountry
	}
	if builder.departureContinentSet {
		data.DepartureContinent = &builder.departureContinent
	}
	if builder.arrivalContinentSet {
		data.ArrivalContinent = &builder.arrivalContinent
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
	if builder.parentInstitutionIdSet {
		data.ParentInstitutionId = &builder.parentInstitutionId
	}
	return data
}
