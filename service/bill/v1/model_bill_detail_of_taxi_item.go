package v1

// BillDetailOfTaxiItem 已出账单 ~ 出租车返回参数
type BillDetailOfTaxiItem struct {
	IsCurrentTerm             *string  `json:"is_current_term,omitempty"`              // 账期 枚举【本期、往期】
	OrderId                   *int64   `json:"order_id,omitempty"`                     // 订单id -企业级内部
	OutOrderId                *int64   `json:"out_order_id,omitempty"`                 // 订单id -网约车
	OrderSource               *string  `json:"order_source,omitempty"`                 // 订单来源 枚举【企业app、H5等】
	Status                    *string  `json:"status,omitempty"`                       // 订单状态 枚举【全部支付、部分支付、超时取消扣费、部分退款、全部退款】
	Rule                      *string  `json:"rule,omitempty"`                         // 用车类型 枚举【出租车、专车、快车、代驾、豪华车、私车公用、同时呼叫】
	RequireLevel              *string  `json:"require_level,omitempty"`                // 车型等级 枚举【专车-舒适型、专车-豪华型、快车-普通型、豪华车-奔驰S级 等】
	MemberName                *string  `json:"member_name,omitempty"`                  // 预定人姓名
	EmployeeNumber            *string  `json:"employee_number,omitempty"`              // 预定人工号
	MemberMail                *string  `json:"member_mail,omitempty"`                  // 预定人邮箱
	CallPhone                 *string  `json:"call_phone,omitempty"`                   // 预定人手机号
	DepartmentName            *string  `json:"department_name,omitempty"`              // 预定人部门路径 如：\"jrm测试公司2023>同一级\"
	OutBudgetId               *string  `json:"out_budget_id,omitempty"`                // 外部成本中心code
	Department                *string  `json:"department,omitempty"`                   // 预定人员工部门 取member department 字段（旧）
	BranchName                *string  `json:"branch_name,omitempty"`                  // 预定人所在分公司名称
	PassengerMemberNumber     *string  `json:"passenger_member_number,omitempty"`      // 乘车人工号
	PassengerName             *string  `json:"passenger_name,omitempty"`               // 乘车人姓名
	PassengerPhone            *string  `json:"passenger_phone,omitempty"`              // 乘车人电话
	PayType                   *string  `json:"pay_type,omitempty"`                     // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】 当前订单支付金额的方式，如：企业支付、混合支付（指企业支付和个人支付两种支付方式同时支付一笔订单）
	TotalPrice                *float64 `json:"total_price,omitempty"`                  // 订单总金额
	CompanyRealPay            *float64 `json:"company_real_pay,omitempty"`             // 企业实付金额
	PersonalRealPay           *float64 `json:"personal_real_pay,omitempty"`            // 个人实付金额
	RealVoucherPay            *float64 `json:"real_voucher_pay,omitempty"`             // 代金券(实际使用金额)
	Cost                      *float64 `json:"cost,omitempty"`                         // 专车订单金额
	DiscountAfterPrice        *float64 `json:"discount_after_price,omitempty"`         // 服务费金额
	CompanyRealRefund         *float64 `json:"company_real_refund,omitempty"`          // 企业实退金额
	Type                      *string  `json:"type,omitempty"`                         // 订单类型 枚举【实时、预约、调账：订单调增、调账：订单调减、调账：主体调增、调账：主体调减、订单计入下一期】
	CreateTime                *string  `json:"create_time,omitempty"`                  // 订单创建时间
	DepartureDime             *string  `json:"departure_dime,omitempty"`               // 出发时间
	MeetTime                  *string  `json:"meetTime,omitempty"`                     // 接驾时间
	BeginChargeTime           *string  `json:"begin_charge_time,omitempty"`            // 开始计费时间
	FinishTime                *string  `json:"finish_time,omitempty"`                  // 完单时间
	PayTime                   *string  `json:"pay_time,omitempty"`                     // 支付时间
	RefundTime                *string  `json:"refund_time,omitempty"`                  // 退款时间
	City                      *string  `json:"city,omitempty"`                         // 城市
	DestinationCity           *string  `json:"destination_city,omitempty"`             // 订单到达城市
	NormalDistance            *string  `json:"normal_distance,omitempty"`              // 实际用车里程
	StartAddress              *string  `json:"start_address,omitempty"`                // 出发地地址
	EndAddress                *string  `json:"end_address,omitempty"`                  // 目的地地址
	StartName                 *string  `json:"start_name,omitempty"`                   // 出发地名称
	EndName                   *string  `json:"end_name,omitempty"`                     // 目的地名称
	ActualStartName           *string  `json:"actual_start_name,omitempty"`            // 实际出发地名称
	ActualEndName             *string  `json:"actual_end_name,omitempty"`              // 实际目的地名称
	RuleName                  *string  `json:"rule_name,omitempty"`                    // 用车规则
	UseCarSrv                 *string  `json:"use_car_srv,omitempty"`                  // 用车类型 枚举【快车用车、专车用车、专车包车、豪华车送机 等】
	BudgetCenterName          *string  `json:"budget_center_name,omitempty"`           // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	BudgetId                  *string  `json:"budget_id,omitempty"`                    // 成本中心ID 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	Remark                    *string  `json:"remark,omitempty"`                       // 备注
	OrderAdditionalRemark     *string  `json:"order_additional_remark,omitempty"`      // 补充说明
	ApprovalReason            *string  `json:"approval_reason,omitempty"`              // 审批备注
	ApprovalCreateTime        *string  `json:"approval_create_time,omitempty"`         // 申请提交时间
	ApprovalLogs              *string  `json:"approval_logs,omitempty"`                // 审批历史-审批时间
	OutApprovalId             *string  `json:"out_approval_id,omitempty"`              // 外部申请单id
	UpgradeType               *string  `json:"upgrade_type,omitempty"`                 // 升舱类型 枚举【非升舱单、升舱、专车拼车、降舱】
	IsFixedPrice              *string  `json:"is_fixed_price,omitempty"`               // 是否一口价 枚举【是、否】
	IsSelfDrive               *string  `json:"is_self_drive,omitempty"`                // 是否是私车同行 枚举【是、否】
	IsDiscountOrder           *string  `json:"is_discount_order,omitempty"`            // 是否有折扣 枚举【是、否】
	IsCarpool                 *string  `json:"is_carpool,omitempty"`                   // 是否拼车 枚举【是、否】
	IsReassignment            *string  `json:"is_reassignment,omitempty"`              // 是否改派 枚举【是、否】
	CallbackInfo              *string  `json:"callback_info,omitempty"`                // 用户自定义字段
	ApprovalExtraInfo         *string  `json:"approval_extra_info,omitempty"`          // 审批单自定义
	AfterApprovalResult       *string  `json:"after_approval_result,omitempty"`        // 行后审批结果
	AfterApprovalContent      *string  `json:"after_approval_content,omitempty"`       // 审批历史
	CompanyId                 *int64   `json:"company_id,omitempty"`                   // 公司id
	CompanyName               *string  `json:"company_name,omitempty"`                 // 公司名称
	LegalEntityId             *int64   `json:"legal_entity_id,omitempty"`              // 预定人开票子公司id
	LegalEntityName           *string  `json:"legal_entity_name,omitempty"`            // 预定人开票子公司名称
	CompanyCardRealPay        *float64 `json:"company_card_real_pay,omitempty"`        // 企业钱包实付金额
	UseCarTypeV2              *string  `json:"use_car_type_v2,omitempty"`              // 用车场景 枚举【加班、日常出勤、接送机 等】
	BudgetItemName            *string  `json:"budget_item_name,omitempty"`             // 科目费用
	UseType                   *string  `json:"use_type,omitempty"`                     // 是否代叫车 枚举【代客户叫车、员工自己用车】
	StartPrice                *float64 `json:"start_price,omitempty"`                  // 起步价
	NormalFee                 *float64 `json:"normal_fee,omitempty"`                   // 里程费
	LowSpeedFee               *float64 `json:"low_speed_fee,omitempty"`                // 低速费
	EmptyFee                  *float64 `json:"empty_fee,omitempty"`                    // 远途费
	NightFee                  *float64 `json:"night_fee,omitempty"`                    // 夜间费
	ResponsibleCancelFee      *float64 `json:"responsible_cancel_fee,omitempty"`       // 有责取消费
	LimitFee                  *float64 `json:"limit_fee,omitempty"`                    // 基础费
	LimitPay                  *float64 `json:"limit_pay,omitempty"`                    // 基础费补足金额
	NormalTimeFee             *float64 `json:"normal_time_fee,omitempty"`              // 时长费
	DynamicPrice              *float64 `json:"dynamic_price,omitempty"`                // 动态调价
	TipFee                    *float64 `json:"tip_fee,omitempty"`                      // 小费
	BridgeFee                 *float64 `json:"bridge_fee,omitempty"`                   // 路桥费
	HighwayFee                *float64 `json:"highway_fee,omitempty"`                  // 高速费
	ParkFee                   *float64 `json:"park_fee,omitempty"`                     // 停车费
	WaitFee                   *float64 `json:"wait_fee,omitempty"`                     // 等待费
	IncentiveFee              *float64 `json:"incentive_fee,omitempty"`                // 调度费
	RedPacket                 *float64 `json:"red_packet,omitempty"`                   // 春节加价费
	OtherFee                  *float64 `json:"other_fee,omitempty"`                    // 其他费用
	MemberId                  *string  `json:"member_id,omitempty"`                    // 预定人id
	RefundPrice               *float64 `json:"refund_price,omitempty"`                 // 退款金额
	GroupName                 *string  `json:"group_name,omitempty"`                   // 预定人部门名称（新）
	GroupId                   *int64   `json:"group_id,omitempty"`                     // 预定人部门id
	CompanyRealPreTaxPay      *float64 `json:"company_real_pre_tax_pay,omitempty"`     // 不含税实付金额 =企业实付金额*（1-税率）
	CompanyRealTaxPay         *float64 `json:"company_real_tax_pay,omitempty"`         // 税额 =企业实付金额*税率
	UserPayTime               *string  `json:"user_pay_time,omitempty"`                // 用户首次确认支付时间
	UserPayType               *string  `json:"user_pay_type,omitempty"`                // 用户支付类型 枚举【用户主动支付、系统自动支付、系统自动支付后用户主动支付】
	SettleType                *string  `json:"settle_type,omitempty"`                  // 业务线名称 网约车
	SubAccountName            *string  `json:"sub_account_name,omitempty"`             // 所属子账户
	BelongBudgetCenterName    *string  `json:"belong_budget_center_name,omitempty"`    // 恒为空 可忽略
	BelongOutBudgetId         *string  `json:"belong_out_budget_id,omitempty"`         // 恒为空 可忽略
	IsCarpoolSuccess          *string  `json:"is_carpool_success,omitempty"`           // 是否拼成 1:是 0:否
	CompanyPayAble            *float64 `json:"company_pay_able,omitempty"`             // 企业应付金额
	ServiceType               *string  `json:"service_type,omitempty"`                 // 包车套餐
	BeforeApprovalResult      *string  `json:"before_approval_result,omitempty"`       // 行前审批
	TimeSlotDiscount          *float64 `json:"time_slot_discount,omitempty"`           // 时段单单惠
	ExtendInfo                *string  `json:"extend_info,omitempty"`                  // 附加信息（开票主体信息） 对应用户invoice_info信息
	Period                    *string  `json:"period,omitempty"`                       // 支付账期
	ExcludingTaxPayPrice      *float64 `json:"excluding_tax_pay_price,omitempty"`      // 不含税实付金额 =企业实付金额*（1-税率） 同company_real_pre_tax_pay 冗余字段
	TaxPayPrice               *float64 `json:"tax_pay_price,omitempty"`                // 实付税额 =企业实付金额*税率 同company_real_tax_pay 冗余字段
	ExcludingTaxRefundPrice   *float64 `json:"excluding_tax_refund_price,omitempty"`   // 不含税实退金额 =企业实退金额*（1-税率）
	TaxRefundPrice            *float64 `json:"tax_refund_price,omitempty"`             // 实退税额 =企业实退金额*税率
	IsSensitive               *string  `json:"is_sensitive,omitempty"`                 // 是否敏感订单 1:是 0:否
	SensitiveReason           *string  `json:"sensitive_reason,omitempty"`             // 敏感订单原因
	CompanyCardRealRefund     *float64 `json:"company_card_real_refund,omitempty"`     // 企业出行卡实退金额
	CancelTime                *string  `json:"cancel_time,omitempty"`                  // 取消时间
	CrossCityFee              *float64 `json:"cross_city_fee,omitempty"`               // 跨城费
	RemoteAreaFee             *float64 `json:"remote_area_fee,omitempty"`              // 偏远地区接驾费
	UseCarTypeName            *string  `json:"use_car_type_name,omitempty"`            // 用车场景细分 枚举【日常用车、加班用车、办公地通勤 等】
	SubUseCarSrv              *string  `json:"sub_use_car_srv,omitempty"`              // 使用场景 枚举【市内用车、接送机、接送火车站、接送汽车站、接送渡口】
	EnergyConsumeFee          *float64 `json:"energy_consume_fee,omitempty"`           // 综合能耗费
	IsUnusual                 *string  `json:"is_unusual,omitempty"`                   // 是否异常 枚举【是、否】
	UnusualType               *string  `json:"unusual_type,omitempty"`                 // 异常类型
	UnusualContent            *string  `json:"unusual_content,omitempty"`              // 异常说明
	PositionName              *string  `json:"position_name,omitempty"`                // 职级
	InstitutionName           *string  `json:"institution_name,omitempty"`             // 用车制度
	UseCarService             *string  `json:"use_car_service,omitempty"`              // 用车服务 枚举【接送服务、出差市内用车】
	ExInfo01                  *string  `json:"ex_info_01,omitempty"`                   // 用户自定义拓展字段1
	ExInfo02                  *string  `json:"ex_info_02,omitempty"`                   // 用户自定义拓展字段2
	ExInfo03                  *string  `json:"ex_info_03,omitempty"`                   // 用户自定义拓展字段3
	ExInfo04                  *string  `json:"ex_info_04,omitempty"`                   // 用户自定义拓展字段4
	ExInfo05                  *string  `json:"ex_info_05,omitempty"`                   // 用户自定义拓展字段5
	ExInfo06                  *string  `json:"ex_info_06,omitempty"`                   // 用户自定义拓展字段6
	ExInfo07                  *string  `json:"ex_info_07,omitempty"`                   // 用户自定义拓展字段7
	ExInfo08                  *string  `json:"ex_info_08,omitempty"`                   // 用户自定义拓展字段8
	SpsPickUpServiceFee       *float64 `json:"sps_pick_up_service_fee,omitempty"`      // 司机举牌接机服务费
	DeductionTypeName         *string  `json:"deduction_type_name,omitempty"`          // 交易类型 枚举：订单企业支付扣款、订单企业支付后退款、企业支付转个人支付退款、行后审批驳回后退款、个人支付转企业支付扣款、线下企业支付扣款、企业支付转个人支付充值、行后审批驳回后充值
	PayChannel                *string  `json:"pay_channel,omitempty"`                  // 结算方式 枚举：银行卡、德付通、预存、返现、授信、试用金、电子账户
	ResidentCityNames         *string  `json:"resident_city_names,omitempty"`          // 常驻城市
	OriginalPrice             *float64 `json:"original_price,omitempty"`               // 原价
	OneTimeOfferSubsidy       *float64 `json:"one_time_offer_subsidy,omitempty"`       // 尊享折扣
	SubsidyAmount             *float64 `json:"subsidy_amount,omitempty"`               // 补贴金额
	VoucherDeductionTypeName  *string  `json:"voucher_deduction_type_name,omitempty"`  // 券抵扣类型
	PassengerMemberId         *int64   `json:"passenger_member_id,omitempty"`          // 乘车人员工id
	ProjectExtInfo            *string  `json:"project_ext_info,omitempty"`             // 项目自定义
	OutLegalEntityId          *string  `json:"out_legal_entity_id,omitempty"`          // 外部公司主体编号
	CarTravelInfo             *string  `json:"car_travel_Info,omitempty"`              // 途经点地址 如：丰台区-大红门街道1号；多个途经点用｜连接，丰台区-大红门街道1号｜朝阳区-国贸5号楼
	InstitutionId             *int64   `json:"institution_id,omitempty"`               // 制度ID 制度-规则ID
	ParentInstitutionId       *int64   `json:"parent_institution_id,omitempty"`        // 制度ID（新） 管控制度，同申请单regulation_id
	ExcludingServiceFee       *float64 `json:"excluding_service_fee,omitempty"`        // 企业实付（不包含平台使用费）
	EstimatePrice             *float64 `json:"estimate_price,omitempty"`               // 预估金额
	SubAccountCompanyName     *string  `json:"sub_account_company_name,omitempty"`     // 子账户公司名称
	ChargeTime                *string  `json:"charge_time,omitempty"`                  // 服务费收取方式
	StartCountyId             *string  `json:"start_county_id,omitempty"`              // 开始区县id
	EndCountyId               *string  `json:"end_county_id,omitempty"`                // 结束区县id
	StartCountyName           *string  `json:"start_county_name,omitempty"`            // 开始区县名称
	EndCountyName             *string  `json:"end_county_name,omitempty"`              // 结束区县名称
	FlagExt                   *string  `json:"flag_ext,omitempty"`                     // 是否追加车型
	CWeight                   *string  `json:"c_weight,omitempty"`                     // 碳排放
	Co2Emissions              *string  `json:"co2_emissions,omitempty"`                // 碳减排
	AfterApproveMemberOne     *string  `json:"after_approve_member_one,omitempty"`     // 行后审批人工号1
	AfterApproveMemberTwo     *string  `json:"after_approve_member_two,omitempty"`     // 行后审批人工号2
	AfterApproveMemberThree   *string  `json:"after_approve_member_three,omitempty"`   // 行后审批人工号3
	CallPhoneCountryCode      *string  `json:"call_phone_country_code,omitempty"`      // 预订人手机号区号
	PassengerPhoneCountryCode *string  `json:"passenger_phone_country_code,omitempty"` // 乘车人手机号区号
	EnterpriseCouponBatch     *string  `json:"enterprise_coupon_batch,omitempty"`      // 企业券批次
	EnterpriseCouponName      *string  `json:"enterprise_coupon_name,omitempty"`       // 企业券名称
	LastApproves              *string  `json:"last_approves,omitempty"`                // 行后审批最后审批人
	LastApprovalTime          *string  `json:"last_approval_time,omitempty"`           // 行后审批最后审批时间
	ExInfo01Code              *string  `json:"ex_info_01_code,omitempty"`              // 用户自定义拓展字段1编码
	ExInfo02Code              *string  `json:"ex_info_02_code,omitempty"`              // 用户自定义拓展字段2编码
	ExInfo03Code              *string  `json:"ex_info_03_code,omitempty"`              // 用户自定义拓展字段3编码
	ExInfo04Code              *string  `json:"ex_info_04_code,omitempty"`              // 用户自定义拓展字段4编码
	ExInfo05Code              *string  `json:"ex_info_05_code,omitempty"`              // 用户自定义拓展字段5编码
	ExInfo06Code              *string  `json:"ex_info_06_code,omitempty"`              // 用户自定义拓展字段6编码
	ExInfo07Code              *string  `json:"ex_info_07_code,omitempty"`              // 用户自定义拓展字段7编码
	ExInfo08Code              *string  `json:"ex_info_08_code,omitempty"`              // 用户自定义拓展字段8编码
}

type BillDetailOfTaxiItemBuilder struct {
	isCurrentTerm                string // 账期 枚举【本期、往期】
	isCurrentTermSet             bool
	orderId                      int64 // 订单id -企业级内部
	orderIdSet                   bool
	outOrderId                   int64 // 订单id -网约车
	outOrderIdSet                bool
	orderSource                  string // 订单来源 枚举【企业app、H5等】
	orderSourceSet               bool
	status                       string // 订单状态 枚举【全部支付、部分支付、超时取消扣费、部分退款、全部退款】
	statusSet                    bool
	rule                         string // 用车类型 枚举【出租车、专车、快车、代驾、豪华车、私车公用、同时呼叫】
	ruleSet                      bool
	requireLevel                 string // 车型等级 枚举【专车-舒适型、专车-豪华型、快车-普通型、豪华车-奔驰S级 等】
	requireLevelSet              bool
	memberName                   string // 预定人姓名
	memberNameSet                bool
	employeeNumber               string // 预定人工号
	employeeNumberSet            bool
	memberMail                   string // 预定人邮箱
	memberMailSet                bool
	callPhone                    string // 预定人手机号
	callPhoneSet                 bool
	departmentName               string // 预定人部门路径 如：\"jrm测试公司2023>同一级\"
	departmentNameSet            bool
	outBudgetId                  string // 外部成本中心code
	outBudgetIdSet               bool
	department                   string // 预定人员工部门 取member department 字段（旧）
	departmentSet                bool
	branchName                   string // 预定人所在分公司名称
	branchNameSet                bool
	passengerMemberNumber        string // 乘车人工号
	passengerMemberNumberSet     bool
	passengerName                string // 乘车人姓名
	passengerNameSet             bool
	passengerPhone               string // 乘车人电话
	passengerPhoneSet            bool
	payType                      string // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】 当前订单支付金额的方式，如：企业支付、混合支付（指企业支付和个人支付两种支付方式同时支付一笔订单）
	payTypeSet                   bool
	totalPrice                   float64 // 订单总金额
	totalPriceSet                bool
	companyRealPay               float64 // 企业实付金额
	companyRealPaySet            bool
	personalRealPay              float64 // 个人实付金额
	personalRealPaySet           bool
	realVoucherPay               float64 // 代金券(实际使用金额)
	realVoucherPaySet            bool
	cost                         float64 // 专车订单金额
	costSet                      bool
	discountAfterPrice           float64 // 服务费金额
	discountAfterPriceSet        bool
	companyRealRefund            float64 // 企业实退金额
	companyRealRefundSet         bool
	type_                        string // 订单类型 枚举【实时、预约、调账：订单调增、调账：订单调减、调账：主体调增、调账：主体调减、订单计入下一期】
	type_Set                     bool
	createTime                   string // 订单创建时间
	createTimeSet                bool
	departureDime                string // 出发时间
	departureDimeSet             bool
	meetTime                     string // 接驾时间
	meetTimeSet                  bool
	beginChargeTime              string // 开始计费时间
	beginChargeTimeSet           bool
	finishTime                   string // 完单时间
	finishTimeSet                bool
	payTime                      string // 支付时间
	payTimeSet                   bool
	refundTime                   string // 退款时间
	refundTimeSet                bool
	city                         string // 城市
	citySet                      bool
	destinationCity              string // 订单到达城市
	destinationCitySet           bool
	normalDistance               string // 实际用车里程
	normalDistanceSet            bool
	startAddress                 string // 出发地地址
	startAddressSet              bool
	endAddress                   string // 目的地地址
	endAddressSet                bool
	startName                    string // 出发地名称
	startNameSet                 bool
	endName                      string // 目的地名称
	endNameSet                   bool
	actualStartName              string // 实际出发地名称
	actualStartNameSet           bool
	actualEndName                string // 实际目的地名称
	actualEndNameSet             bool
	ruleName                     string // 用车规则
	ruleNameSet                  bool
	useCarSrv                    string // 用车类型 枚举【快车用车、专车用车、专车包车、豪华车送机 等】
	useCarSrvSet                 bool
	budgetCenterName             string // 成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段
	budgetCenterNameSet          bool
	budgetId                     string // 成本中心ID 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段
	budgetIdSet                  bool
	remark                       string // 备注
	remarkSet                    bool
	orderAdditionalRemark        string // 补充说明
	orderAdditionalRemarkSet     bool
	approvalReason               string // 审批备注
	approvalReasonSet            bool
	approvalCreateTime           string // 申请提交时间
	approvalCreateTimeSet        bool
	approvalLogs                 string // 审批历史-审批时间
	approvalLogsSet              bool
	outApprovalId                string // 外部申请单id
	outApprovalIdSet             bool
	upgradeType                  string // 升舱类型 枚举【非升舱单、升舱、专车拼车、降舱】
	upgradeTypeSet               bool
	isFixedPrice                 string // 是否一口价 枚举【是、否】
	isFixedPriceSet              bool
	isSelfDrive                  string // 是否是私车同行 枚举【是、否】
	isSelfDriveSet               bool
	isDiscountOrder              string // 是否有折扣 枚举【是、否】
	isDiscountOrderSet           bool
	isCarpool                    string // 是否拼车 枚举【是、否】
	isCarpoolSet                 bool
	isReassignment               string // 是否改派 枚举【是、否】
	isReassignmentSet            bool
	callbackInfo                 string // 用户自定义字段
	callbackInfoSet              bool
	approvalExtraInfo            string // 审批单自定义
	approvalExtraInfoSet         bool
	afterApprovalResult          string // 行后审批结果
	afterApprovalResultSet       bool
	afterApprovalContent         string // 审批历史
	afterApprovalContentSet      bool
	companyId                    int64 // 公司id
	companyIdSet                 bool
	companyName                  string // 公司名称
	companyNameSet               bool
	legalEntityId                int64 // 预定人开票子公司id
	legalEntityIdSet             bool
	legalEntityName              string // 预定人开票子公司名称
	legalEntityNameSet           bool
	companyCardRealPay           float64 // 企业钱包实付金额
	companyCardRealPaySet        bool
	useCarTypeV2                 string // 用车场景 枚举【加班、日常出勤、接送机 等】
	useCarTypeV2Set              bool
	budgetItemName               string // 科目费用
	budgetItemNameSet            bool
	useType                      string // 是否代叫车 枚举【代客户叫车、员工自己用车】
	useTypeSet                   bool
	startPrice                   float64 // 起步价
	startPriceSet                bool
	normalFee                    float64 // 里程费
	normalFeeSet                 bool
	lowSpeedFee                  float64 // 低速费
	lowSpeedFeeSet               bool
	emptyFee                     float64 // 远途费
	emptyFeeSet                  bool
	nightFee                     float64 // 夜间费
	nightFeeSet                  bool
	responsibleCancelFee         float64 // 有责取消费
	responsibleCancelFeeSet      bool
	limitFee                     float64 // 基础费
	limitFeeSet                  bool
	limitPay                     float64 // 基础费补足金额
	limitPaySet                  bool
	normalTimeFee                float64 // 时长费
	normalTimeFeeSet             bool
	dynamicPrice                 float64 // 动态调价
	dynamicPriceSet              bool
	tipFee                       float64 // 小费
	tipFeeSet                    bool
	bridgeFee                    float64 // 路桥费
	bridgeFeeSet                 bool
	highwayFee                   float64 // 高速费
	highwayFeeSet                bool
	parkFee                      float64 // 停车费
	parkFeeSet                   bool
	waitFee                      float64 // 等待费
	waitFeeSet                   bool
	incentiveFee                 float64 // 调度费
	incentiveFeeSet              bool
	redPacket                    float64 // 春节加价费
	redPacketSet                 bool
	otherFee                     float64 // 其他费用
	otherFeeSet                  bool
	memberId                     string // 预定人id
	memberIdSet                  bool
	refundPrice                  float64 // 退款金额
	refundPriceSet               bool
	groupName                    string // 预定人部门名称（新）
	groupNameSet                 bool
	groupId                      int64 // 预定人部门id
	groupIdSet                   bool
	companyRealPreTaxPay         float64 // 不含税实付金额 =企业实付金额*（1-税率）
	companyRealPreTaxPaySet      bool
	companyRealTaxPay            float64 // 税额 =企业实付金额*税率
	companyRealTaxPaySet         bool
	userPayTime                  string // 用户首次确认支付时间
	userPayTimeSet               bool
	userPayType                  string // 用户支付类型 枚举【用户主动支付、系统自动支付、系统自动支付后用户主动支付】
	userPayTypeSet               bool
	settleType                   string // 业务线名称 网约车
	settleTypeSet                bool
	subAccountName               string // 所属子账户
	subAccountNameSet            bool
	belongBudgetCenterName       string // 恒为空 可忽略
	belongBudgetCenterNameSet    bool
	belongOutBudgetId            string // 恒为空 可忽略
	belongOutBudgetIdSet         bool
	isCarpoolSuccess             string // 是否拼成 1:是 0:否
	isCarpoolSuccessSet          bool
	companyPayAble               float64 // 企业应付金额
	companyPayAbleSet            bool
	serviceType                  string // 包车套餐
	serviceTypeSet               bool
	beforeApprovalResult         string // 行前审批
	beforeApprovalResultSet      bool
	timeSlotDiscount             float64 // 时段单单惠
	timeSlotDiscountSet          bool
	extendInfo                   string // 附加信息（开票主体信息） 对应用户invoice_info信息
	extendInfoSet                bool
	period                       string // 支付账期
	periodSet                    bool
	excludingTaxPayPrice         float64 // 不含税实付金额 =企业实付金额*（1-税率） 同company_real_pre_tax_pay 冗余字段
	excludingTaxPayPriceSet      bool
	taxPayPrice                  float64 // 实付税额 =企业实付金额*税率 同company_real_tax_pay 冗余字段
	taxPayPriceSet               bool
	excludingTaxRefundPrice      float64 // 不含税实退金额 =企业实退金额*（1-税率）
	excludingTaxRefundPriceSet   bool
	taxRefundPrice               float64 // 实退税额 =企业实退金额*税率
	taxRefundPriceSet            bool
	isSensitive                  string // 是否敏感订单 1:是 0:否
	isSensitiveSet               bool
	sensitiveReason              string // 敏感订单原因
	sensitiveReasonSet           bool
	companyCardRealRefund        float64 // 企业出行卡实退金额
	companyCardRealRefundSet     bool
	cancelTime                   string // 取消时间
	cancelTimeSet                bool
	crossCityFee                 float64 // 跨城费
	crossCityFeeSet              bool
	remoteAreaFee                float64 // 偏远地区接驾费
	remoteAreaFeeSet             bool
	useCarTypeName               string // 用车场景细分 枚举【日常用车、加班用车、办公地通勤 等】
	useCarTypeNameSet            bool
	subUseCarSrv                 string // 使用场景 枚举【市内用车、接送机、接送火车站、接送汽车站、接送渡口】
	subUseCarSrvSet              bool
	energyConsumeFee             float64 // 综合能耗费
	energyConsumeFeeSet          bool
	isUnusual                    string // 是否异常 枚举【是、否】
	isUnusualSet                 bool
	unusualType                  string // 异常类型
	unusualTypeSet               bool
	unusualContent               string // 异常说明
	unusualContentSet            bool
	positionName                 string // 职级
	positionNameSet              bool
	institutionName              string // 用车制度
	institutionNameSet           bool
	useCarService                string // 用车服务 枚举【接送服务、出差市内用车】
	useCarServiceSet             bool
	exInfo01                     string // 用户自定义拓展字段1
	exInfo01Set                  bool
	exInfo02                     string // 用户自定义拓展字段2
	exInfo02Set                  bool
	exInfo03                     string // 用户自定义拓展字段3
	exInfo03Set                  bool
	exInfo04                     string // 用户自定义拓展字段4
	exInfo04Set                  bool
	exInfo05                     string // 用户自定义拓展字段5
	exInfo05Set                  bool
	exInfo06                     string // 用户自定义拓展字段6
	exInfo06Set                  bool
	exInfo07                     string // 用户自定义拓展字段7
	exInfo07Set                  bool
	exInfo08                     string // 用户自定义拓展字段8
	exInfo08Set                  bool
	spsPickUpServiceFee          float64 // 司机举牌接机服务费
	spsPickUpServiceFeeSet       bool
	deductionTypeName            string // 交易类型 枚举：订单企业支付扣款、订单企业支付后退款、企业支付转个人支付退款、行后审批驳回后退款、个人支付转企业支付扣款、线下企业支付扣款、企业支付转个人支付充值、行后审批驳回后充值
	deductionTypeNameSet         bool
	payChannel                   string // 结算方式 枚举：银行卡、德付通、预存、返现、授信、试用金、电子账户
	payChannelSet                bool
	residentCityNames            string // 常驻城市
	residentCityNamesSet         bool
	originalPrice                float64 // 原价
	originalPriceSet             bool
	oneTimeOfferSubsidy          float64 // 尊享折扣
	oneTimeOfferSubsidySet       bool
	subsidyAmount                float64 // 补贴金额
	subsidyAmountSet             bool
	voucherDeductionTypeName     string // 券抵扣类型
	voucherDeductionTypeNameSet  bool
	passengerMemberId            int64 // 乘车人员工id
	passengerMemberIdSet         bool
	projectExtInfo               string // 项目自定义
	projectExtInfoSet            bool
	outLegalEntityId             string // 外部公司主体编号
	outLegalEntityIdSet          bool
	carTravelInfo                string // 途经点地址 如：丰台区-大红门街道1号；多个途经点用｜连接，丰台区-大红门街道1号｜朝阳区-国贸5号楼
	carTravelInfoSet             bool
	institutionId                int64 // 制度ID 制度-规则ID
	institutionIdSet             bool
	parentInstitutionId          int64 // 制度ID（新） 管控制度，同申请单regulation_id
	parentInstitutionIdSet       bool
	excludingServiceFee          float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet       bool
	estimatePrice                float64 // 预估金额
	estimatePriceSet             bool
	subAccountCompanyName        string // 子账户公司名称
	subAccountCompanyNameSet     bool
	chargeTime                   string // 服务费收取方式
	chargeTimeSet                bool
	startCountyId                string // 开始区县id
	startCountyIdSet             bool
	endCountyId                  string // 结束区县id
	endCountyIdSet               bool
	startCountyName              string // 开始区县名称
	startCountyNameSet           bool
	endCountyName                string // 结束区县名称
	endCountyNameSet             bool
	flagExt                      string // 是否追加车型
	flagExtSet                   bool
	cWeight                      string // 碳排放
	cWeightSet                   bool
	co2Emissions                 string // 碳减排
	co2EmissionsSet              bool
	afterApproveMemberOne        string // 行后审批人工号1
	afterApproveMemberOneSet     bool
	afterApproveMemberTwo        string // 行后审批人工号2
	afterApproveMemberTwoSet     bool
	afterApproveMemberThree      string // 行后审批人工号3
	afterApproveMemberThreeSet   bool
	callPhoneCountryCode         string // 预订人手机号区号
	callPhoneCountryCodeSet      bool
	passengerPhoneCountryCode    string // 乘车人手机号区号
	passengerPhoneCountryCodeSet bool
	enterpriseCouponBatch        string // 企业券批次
	enterpriseCouponBatchSet     bool
	enterpriseCouponName         string // 企业券名称
	enterpriseCouponNameSet      bool
	lastApproves                 string // 行后审批最后审批人
	lastApprovesSet              bool
	lastApprovalTime             string // 行后审批最后审批时间
	lastApprovalTimeSet          bool
	exInfo01Code                 string // 用户自定义拓展字段1编码
	exInfo01CodeSet              bool
	exInfo02Code                 string // 用户自定义拓展字段2编码
	exInfo02CodeSet              bool
	exInfo03Code                 string // 用户自定义拓展字段3编码
	exInfo03CodeSet              bool
	exInfo04Code                 string // 用户自定义拓展字段4编码
	exInfo04CodeSet              bool
	exInfo05Code                 string // 用户自定义拓展字段5编码
	exInfo05CodeSet              bool
	exInfo06Code                 string // 用户自定义拓展字段6编码
	exInfo06CodeSet              bool
	exInfo07Code                 string // 用户自定义拓展字段7编码
	exInfo07CodeSet              bool
	exInfo08Code                 string // 用户自定义拓展字段8编码
	exInfo08CodeSet              bool
}

func NewBillDetailOfTaxiItemBuilder() *BillDetailOfTaxiItemBuilder {
	return &BillDetailOfTaxiItemBuilder{}
}
func (builder *BillDetailOfTaxiItemBuilder) IsCurrentTerm(isCurrentTerm string) *BillDetailOfTaxiItemBuilder {
	builder.isCurrentTerm = isCurrentTerm
	builder.isCurrentTermSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OrderId(orderId int64) *BillDetailOfTaxiItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OutOrderId(outOrderId int64) *BillDetailOfTaxiItemBuilder {
	builder.outOrderId = outOrderId
	builder.outOrderIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OrderSource(orderSource string) *BillDetailOfTaxiItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Status(status string) *BillDetailOfTaxiItemBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Rule(rule string) *BillDetailOfTaxiItemBuilder {
	builder.rule = rule
	builder.ruleSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RequireLevel(requireLevel string) *BillDetailOfTaxiItemBuilder {
	builder.requireLevel = requireLevel
	builder.requireLevelSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) MemberName(memberName string) *BillDetailOfTaxiItemBuilder {
	builder.memberName = memberName
	builder.memberNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EmployeeNumber(employeeNumber string) *BillDetailOfTaxiItemBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) MemberMail(memberMail string) *BillDetailOfTaxiItemBuilder {
	builder.memberMail = memberMail
	builder.memberMailSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CallPhone(callPhone string) *BillDetailOfTaxiItemBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DepartmentName(departmentName string) *BillDetailOfTaxiItemBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OutBudgetId(outBudgetId string) *BillDetailOfTaxiItemBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Department(department string) *BillDetailOfTaxiItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BranchName(branchName string) *BillDetailOfTaxiItemBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillDetailOfTaxiItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PassengerName(passengerName string) *BillDetailOfTaxiItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PassengerPhone(passengerPhone string) *BillDetailOfTaxiItemBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PayType(payType string) *BillDetailOfTaxiItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) TotalPrice(totalPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyRealPay(companyRealPay float64) *BillDetailOfTaxiItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PersonalRealPay(personalRealPay float64) *BillDetailOfTaxiItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RealVoucherPay(realVoucherPay float64) *BillDetailOfTaxiItemBuilder {
	builder.realVoucherPay = realVoucherPay
	builder.realVoucherPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Cost(cost float64) *BillDetailOfTaxiItemBuilder {
	builder.cost = cost
	builder.costSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DiscountAfterPrice(discountAfterPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.discountAfterPrice = discountAfterPrice
	builder.discountAfterPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyRealRefund(companyRealRefund float64) *BillDetailOfTaxiItemBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Type(type_ string) *BillDetailOfTaxiItemBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CreateTime(createTime string) *BillDetailOfTaxiItemBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DepartureDime(departureDime string) *BillDetailOfTaxiItemBuilder {
	builder.departureDime = departureDime
	builder.departureDimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) MeetTime(meetTime string) *BillDetailOfTaxiItemBuilder {
	builder.meetTime = meetTime
	builder.meetTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BeginChargeTime(beginChargeTime string) *BillDetailOfTaxiItemBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) FinishTime(finishTime string) *BillDetailOfTaxiItemBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PayTime(payTime string) *BillDetailOfTaxiItemBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RefundTime(refundTime string) *BillDetailOfTaxiItemBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) City(city string) *BillDetailOfTaxiItemBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DestinationCity(destinationCity string) *BillDetailOfTaxiItemBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) NormalDistance(normalDistance string) *BillDetailOfTaxiItemBuilder {
	builder.normalDistance = normalDistance
	builder.normalDistanceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) StartAddress(startAddress string) *BillDetailOfTaxiItemBuilder {
	builder.startAddress = startAddress
	builder.startAddressSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EndAddress(endAddress string) *BillDetailOfTaxiItemBuilder {
	builder.endAddress = endAddress
	builder.endAddressSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) StartName(startName string) *BillDetailOfTaxiItemBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EndName(endName string) *BillDetailOfTaxiItemBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ActualStartName(actualStartName string) *BillDetailOfTaxiItemBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ActualEndName(actualEndName string) *BillDetailOfTaxiItemBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RuleName(ruleName string) *BillDetailOfTaxiItemBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UseCarSrv(useCarSrv string) *BillDetailOfTaxiItemBuilder {
	builder.useCarSrv = useCarSrv
	builder.useCarSrvSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BudgetCenterName(budgetCenterName string) *BillDetailOfTaxiItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BudgetId(budgetId string) *BillDetailOfTaxiItemBuilder {
	builder.budgetId = budgetId
	builder.budgetIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Remark(remark string) *BillDetailOfTaxiItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OrderAdditionalRemark(orderAdditionalRemark string) *BillDetailOfTaxiItemBuilder {
	builder.orderAdditionalRemark = orderAdditionalRemark
	builder.orderAdditionalRemarkSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ApprovalReason(approvalReason string) *BillDetailOfTaxiItemBuilder {
	builder.approvalReason = approvalReason
	builder.approvalReasonSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ApprovalCreateTime(approvalCreateTime string) *BillDetailOfTaxiItemBuilder {
	builder.approvalCreateTime = approvalCreateTime
	builder.approvalCreateTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ApprovalLogs(approvalLogs string) *BillDetailOfTaxiItemBuilder {
	builder.approvalLogs = approvalLogs
	builder.approvalLogsSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OutApprovalId(outApprovalId string) *BillDetailOfTaxiItemBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UpgradeType(upgradeType string) *BillDetailOfTaxiItemBuilder {
	builder.upgradeType = upgradeType
	builder.upgradeTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsFixedPrice(isFixedPrice string) *BillDetailOfTaxiItemBuilder {
	builder.isFixedPrice = isFixedPrice
	builder.isFixedPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsSelfDrive(isSelfDrive string) *BillDetailOfTaxiItemBuilder {
	builder.isSelfDrive = isSelfDrive
	builder.isSelfDriveSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsDiscountOrder(isDiscountOrder string) *BillDetailOfTaxiItemBuilder {
	builder.isDiscountOrder = isDiscountOrder
	builder.isDiscountOrderSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsCarpool(isCarpool string) *BillDetailOfTaxiItemBuilder {
	builder.isCarpool = isCarpool
	builder.isCarpoolSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsReassignment(isReassignment string) *BillDetailOfTaxiItemBuilder {
	builder.isReassignment = isReassignment
	builder.isReassignmentSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CallbackInfo(callbackInfo string) *BillDetailOfTaxiItemBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ApprovalExtraInfo(approvalExtraInfo string) *BillDetailOfTaxiItemBuilder {
	builder.approvalExtraInfo = approvalExtraInfo
	builder.approvalExtraInfoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) AfterApprovalResult(afterApprovalResult string) *BillDetailOfTaxiItemBuilder {
	builder.afterApprovalResult = afterApprovalResult
	builder.afterApprovalResultSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) AfterApprovalContent(afterApprovalContent string) *BillDetailOfTaxiItemBuilder {
	builder.afterApprovalContent = afterApprovalContent
	builder.afterApprovalContentSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyId(companyId int64) *BillDetailOfTaxiItemBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyName(companyName string) *BillDetailOfTaxiItemBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LegalEntityId(legalEntityId int64) *BillDetailOfTaxiItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LegalEntityName(legalEntityName string) *BillDetailOfTaxiItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyCardRealPay(companyCardRealPay float64) *BillDetailOfTaxiItemBuilder {
	builder.companyCardRealPay = companyCardRealPay
	builder.companyCardRealPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UseCarTypeV2(useCarTypeV2 string) *BillDetailOfTaxiItemBuilder {
	builder.useCarTypeV2 = useCarTypeV2
	builder.useCarTypeV2Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BudgetItemName(budgetItemName string) *BillDetailOfTaxiItemBuilder {
	builder.budgetItemName = budgetItemName
	builder.budgetItemNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UseType(useType string) *BillDetailOfTaxiItemBuilder {
	builder.useType = useType
	builder.useTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) StartPrice(startPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.startPrice = startPrice
	builder.startPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) NormalFee(normalFee float64) *BillDetailOfTaxiItemBuilder {
	builder.normalFee = normalFee
	builder.normalFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LowSpeedFee(lowSpeedFee float64) *BillDetailOfTaxiItemBuilder {
	builder.lowSpeedFee = lowSpeedFee
	builder.lowSpeedFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EmptyFee(emptyFee float64) *BillDetailOfTaxiItemBuilder {
	builder.emptyFee = emptyFee
	builder.emptyFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) NightFee(nightFee float64) *BillDetailOfTaxiItemBuilder {
	builder.nightFee = nightFee
	builder.nightFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ResponsibleCancelFee(responsibleCancelFee float64) *BillDetailOfTaxiItemBuilder {
	builder.responsibleCancelFee = responsibleCancelFee
	builder.responsibleCancelFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LimitFee(limitFee float64) *BillDetailOfTaxiItemBuilder {
	builder.limitFee = limitFee
	builder.limitFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LimitPay(limitPay float64) *BillDetailOfTaxiItemBuilder {
	builder.limitPay = limitPay
	builder.limitPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) NormalTimeFee(normalTimeFee float64) *BillDetailOfTaxiItemBuilder {
	builder.normalTimeFee = normalTimeFee
	builder.normalTimeFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DynamicPrice(dynamicPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.dynamicPrice = dynamicPrice
	builder.dynamicPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) TipFee(tipFee float64) *BillDetailOfTaxiItemBuilder {
	builder.tipFee = tipFee
	builder.tipFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BridgeFee(bridgeFee float64) *BillDetailOfTaxiItemBuilder {
	builder.bridgeFee = bridgeFee
	builder.bridgeFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) HighwayFee(highwayFee float64) *BillDetailOfTaxiItemBuilder {
	builder.highwayFee = highwayFee
	builder.highwayFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ParkFee(parkFee float64) *BillDetailOfTaxiItemBuilder {
	builder.parkFee = parkFee
	builder.parkFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) WaitFee(waitFee float64) *BillDetailOfTaxiItemBuilder {
	builder.waitFee = waitFee
	builder.waitFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IncentiveFee(incentiveFee float64) *BillDetailOfTaxiItemBuilder {
	builder.incentiveFee = incentiveFee
	builder.incentiveFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RedPacket(redPacket float64) *BillDetailOfTaxiItemBuilder {
	builder.redPacket = redPacket
	builder.redPacketSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OtherFee(otherFee float64) *BillDetailOfTaxiItemBuilder {
	builder.otherFee = otherFee
	builder.otherFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) MemberId(memberId string) *BillDetailOfTaxiItemBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RefundPrice(refundPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.refundPrice = refundPrice
	builder.refundPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) GroupName(groupName string) *BillDetailOfTaxiItemBuilder {
	builder.groupName = groupName
	builder.groupNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) GroupId(groupId int64) *BillDetailOfTaxiItemBuilder {
	builder.groupId = groupId
	builder.groupIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyRealPreTaxPay(companyRealPreTaxPay float64) *BillDetailOfTaxiItemBuilder {
	builder.companyRealPreTaxPay = companyRealPreTaxPay
	builder.companyRealPreTaxPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyRealTaxPay(companyRealTaxPay float64) *BillDetailOfTaxiItemBuilder {
	builder.companyRealTaxPay = companyRealTaxPay
	builder.companyRealTaxPaySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UserPayTime(userPayTime string) *BillDetailOfTaxiItemBuilder {
	builder.userPayTime = userPayTime
	builder.userPayTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UserPayType(userPayType string) *BillDetailOfTaxiItemBuilder {
	builder.userPayType = userPayType
	builder.userPayTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SettleType(settleType string) *BillDetailOfTaxiItemBuilder {
	builder.settleType = settleType
	builder.settleTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SubAccountName(subAccountName string) *BillDetailOfTaxiItemBuilder {
	builder.subAccountName = subAccountName
	builder.subAccountNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BelongBudgetCenterName(belongBudgetCenterName string) *BillDetailOfTaxiItemBuilder {
	builder.belongBudgetCenterName = belongBudgetCenterName
	builder.belongBudgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BelongOutBudgetId(belongOutBudgetId string) *BillDetailOfTaxiItemBuilder {
	builder.belongOutBudgetId = belongOutBudgetId
	builder.belongOutBudgetIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsCarpoolSuccess(isCarpoolSuccess string) *BillDetailOfTaxiItemBuilder {
	builder.isCarpoolSuccess = isCarpoolSuccess
	builder.isCarpoolSuccessSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyPayAble(companyPayAble float64) *BillDetailOfTaxiItemBuilder {
	builder.companyPayAble = companyPayAble
	builder.companyPayAbleSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ServiceType(serviceType string) *BillDetailOfTaxiItemBuilder {
	builder.serviceType = serviceType
	builder.serviceTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) BeforeApprovalResult(beforeApprovalResult string) *BillDetailOfTaxiItemBuilder {
	builder.beforeApprovalResult = beforeApprovalResult
	builder.beforeApprovalResultSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) TimeSlotDiscount(timeSlotDiscount float64) *BillDetailOfTaxiItemBuilder {
	builder.timeSlotDiscount = timeSlotDiscount
	builder.timeSlotDiscountSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExtendInfo(extendInfo string) *BillDetailOfTaxiItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Period(period string) *BillDetailOfTaxiItemBuilder {
	builder.period = period
	builder.periodSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExcludingTaxPayPrice(excludingTaxPayPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.excludingTaxPayPrice = excludingTaxPayPrice
	builder.excludingTaxPayPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) TaxPayPrice(taxPayPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.taxPayPrice = taxPayPrice
	builder.taxPayPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExcludingTaxRefundPrice(excludingTaxRefundPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.excludingTaxRefundPrice = excludingTaxRefundPrice
	builder.excludingTaxRefundPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) TaxRefundPrice(taxRefundPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.taxRefundPrice = taxRefundPrice
	builder.taxRefundPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsSensitive(isSensitive string) *BillDetailOfTaxiItemBuilder {
	builder.isSensitive = isSensitive
	builder.isSensitiveSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SensitiveReason(sensitiveReason string) *BillDetailOfTaxiItemBuilder {
	builder.sensitiveReason = sensitiveReason
	builder.sensitiveReasonSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CompanyCardRealRefund(companyCardRealRefund float64) *BillDetailOfTaxiItemBuilder {
	builder.companyCardRealRefund = companyCardRealRefund
	builder.companyCardRealRefundSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CancelTime(cancelTime string) *BillDetailOfTaxiItemBuilder {
	builder.cancelTime = cancelTime
	builder.cancelTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CrossCityFee(crossCityFee float64) *BillDetailOfTaxiItemBuilder {
	builder.crossCityFee = crossCityFee
	builder.crossCityFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) RemoteAreaFee(remoteAreaFee float64) *BillDetailOfTaxiItemBuilder {
	builder.remoteAreaFee = remoteAreaFee
	builder.remoteAreaFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UseCarTypeName(useCarTypeName string) *BillDetailOfTaxiItemBuilder {
	builder.useCarTypeName = useCarTypeName
	builder.useCarTypeNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SubUseCarSrv(subUseCarSrv string) *BillDetailOfTaxiItemBuilder {
	builder.subUseCarSrv = subUseCarSrv
	builder.subUseCarSrvSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EnergyConsumeFee(energyConsumeFee float64) *BillDetailOfTaxiItemBuilder {
	builder.energyConsumeFee = energyConsumeFee
	builder.energyConsumeFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) IsUnusual(isUnusual string) *BillDetailOfTaxiItemBuilder {
	builder.isUnusual = isUnusual
	builder.isUnusualSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UnusualType(unusualType string) *BillDetailOfTaxiItemBuilder {
	builder.unusualType = unusualType
	builder.unusualTypeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UnusualContent(unusualContent string) *BillDetailOfTaxiItemBuilder {
	builder.unusualContent = unusualContent
	builder.unusualContentSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PositionName(positionName string) *BillDetailOfTaxiItemBuilder {
	builder.positionName = positionName
	builder.positionNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) InstitutionName(institutionName string) *BillDetailOfTaxiItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) UseCarService(useCarService string) *BillDetailOfTaxiItemBuilder {
	builder.useCarService = useCarService
	builder.useCarServiceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo01(exInfo01 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo02(exInfo02 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo03(exInfo03 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo04(exInfo04 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo05(exInfo05 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo06(exInfo06 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo07(exInfo07 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo08(exInfo08 string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SpsPickUpServiceFee(spsPickUpServiceFee float64) *BillDetailOfTaxiItemBuilder {
	builder.spsPickUpServiceFee = spsPickUpServiceFee
	builder.spsPickUpServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) DeductionTypeName(deductionTypeName string) *BillDetailOfTaxiItemBuilder {
	builder.deductionTypeName = deductionTypeName
	builder.deductionTypeNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PayChannel(payChannel string) *BillDetailOfTaxiItemBuilder {
	builder.payChannel = payChannel
	builder.payChannelSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ResidentCityNames(residentCityNames string) *BillDetailOfTaxiItemBuilder {
	builder.residentCityNames = residentCityNames
	builder.residentCityNamesSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OriginalPrice(originalPrice float64) *BillDetailOfTaxiItemBuilder {
	builder.originalPrice = originalPrice
	builder.originalPriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OneTimeOfferSubsidy(oneTimeOfferSubsidy float64) *BillDetailOfTaxiItemBuilder {
	builder.oneTimeOfferSubsidy = oneTimeOfferSubsidy
	builder.oneTimeOfferSubsidySet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SubsidyAmount(subsidyAmount float64) *BillDetailOfTaxiItemBuilder {
	builder.subsidyAmount = subsidyAmount
	builder.subsidyAmountSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) VoucherDeductionTypeName(voucherDeductionTypeName string) *BillDetailOfTaxiItemBuilder {
	builder.voucherDeductionTypeName = voucherDeductionTypeName
	builder.voucherDeductionTypeNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PassengerMemberId(passengerMemberId int64) *BillDetailOfTaxiItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ProjectExtInfo(projectExtInfo string) *BillDetailOfTaxiItemBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) OutLegalEntityId(outLegalEntityId string) *BillDetailOfTaxiItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CarTravelInfo(carTravelInfo string) *BillDetailOfTaxiItemBuilder {
	builder.carTravelInfo = carTravelInfo
	builder.carTravelInfoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) InstitutionId(institutionId int64) *BillDetailOfTaxiItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ParentInstitutionId(parentInstitutionId int64) *BillDetailOfTaxiItemBuilder {
	builder.parentInstitutionId = parentInstitutionId
	builder.parentInstitutionIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *BillDetailOfTaxiItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EstimatePrice(estimatePrice float64) *BillDetailOfTaxiItemBuilder {
	builder.estimatePrice = estimatePrice
	builder.estimatePriceSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *BillDetailOfTaxiItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ChargeTime(chargeTime string) *BillDetailOfTaxiItemBuilder {
	builder.chargeTime = chargeTime
	builder.chargeTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) StartCountyId(startCountyId string) *BillDetailOfTaxiItemBuilder {
	builder.startCountyId = startCountyId
	builder.startCountyIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EndCountyId(endCountyId string) *BillDetailOfTaxiItemBuilder {
	builder.endCountyId = endCountyId
	builder.endCountyIdSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) StartCountyName(startCountyName string) *BillDetailOfTaxiItemBuilder {
	builder.startCountyName = startCountyName
	builder.startCountyNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EndCountyName(endCountyName string) *BillDetailOfTaxiItemBuilder {
	builder.endCountyName = endCountyName
	builder.endCountyNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) FlagExt(flagExt string) *BillDetailOfTaxiItemBuilder {
	builder.flagExt = flagExt
	builder.flagExtSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CWeight(cWeight string) *BillDetailOfTaxiItemBuilder {
	builder.cWeight = cWeight
	builder.cWeightSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) Co2Emissions(co2Emissions string) *BillDetailOfTaxiItemBuilder {
	builder.co2Emissions = co2Emissions
	builder.co2EmissionsSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) AfterApproveMemberOne(afterApproveMemberOne string) *BillDetailOfTaxiItemBuilder {
	builder.afterApproveMemberOne = afterApproveMemberOne
	builder.afterApproveMemberOneSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) AfterApproveMemberTwo(afterApproveMemberTwo string) *BillDetailOfTaxiItemBuilder {
	builder.afterApproveMemberTwo = afterApproveMemberTwo
	builder.afterApproveMemberTwoSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) AfterApproveMemberThree(afterApproveMemberThree string) *BillDetailOfTaxiItemBuilder {
	builder.afterApproveMemberThree = afterApproveMemberThree
	builder.afterApproveMemberThreeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) CallPhoneCountryCode(callPhoneCountryCode string) *BillDetailOfTaxiItemBuilder {
	builder.callPhoneCountryCode = callPhoneCountryCode
	builder.callPhoneCountryCodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) PassengerPhoneCountryCode(passengerPhoneCountryCode string) *BillDetailOfTaxiItemBuilder {
	builder.passengerPhoneCountryCode = passengerPhoneCountryCode
	builder.passengerPhoneCountryCodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EnterpriseCouponBatch(enterpriseCouponBatch string) *BillDetailOfTaxiItemBuilder {
	builder.enterpriseCouponBatch = enterpriseCouponBatch
	builder.enterpriseCouponBatchSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) EnterpriseCouponName(enterpriseCouponName string) *BillDetailOfTaxiItemBuilder {
	builder.enterpriseCouponName = enterpriseCouponName
	builder.enterpriseCouponNameSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LastApproves(lastApproves string) *BillDetailOfTaxiItemBuilder {
	builder.lastApproves = lastApproves
	builder.lastApprovesSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) LastApprovalTime(lastApprovalTime string) *BillDetailOfTaxiItemBuilder {
	builder.lastApprovalTime = lastApprovalTime
	builder.lastApprovalTimeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo01Code(exInfo01Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo02Code(exInfo02Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo03Code(exInfo03Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo04Code(exInfo04Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo05Code(exInfo05Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo06Code(exInfo06Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo07Code(exInfo07Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *BillDetailOfTaxiItemBuilder) ExInfo08Code(exInfo08Code string) *BillDetailOfTaxiItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}

func (builder *BillDetailOfTaxiItemBuilder) Build() *BillDetailOfTaxiItem {
	data := &BillDetailOfTaxiItem{}
	if builder.isCurrentTermSet {
		data.IsCurrentTerm = &builder.isCurrentTerm
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
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
	if builder.ruleSet {
		data.Rule = &builder.rule
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
	if builder.memberMailSet {
		data.MemberMail = &builder.memberMail
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
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.branchNameSet {
		data.BranchName = &builder.branchName
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
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
	if builder.realVoucherPaySet {
		data.RealVoucherPay = &builder.realVoucherPay
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
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.departureDimeSet {
		data.DepartureDime = &builder.departureDime
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
	if builder.endAddressSet {
		data.EndAddress = &builder.endAddress
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
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.useCarSrvSet {
		data.UseCarSrv = &builder.useCarSrv
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetIdSet {
		data.BudgetId = &builder.budgetId
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
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
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
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
	if builder.approvalExtraInfoSet {
		data.ApprovalExtraInfo = &builder.approvalExtraInfo
	}
	if builder.afterApprovalResultSet {
		data.AfterApprovalResult = &builder.afterApprovalResult
	}
	if builder.afterApprovalContentSet {
		data.AfterApprovalContent = &builder.afterApprovalContent
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.companyNameSet {
		data.CompanyName = &builder.companyName
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.companyCardRealPaySet {
		data.CompanyCardRealPay = &builder.companyCardRealPay
	}
	if builder.useCarTypeV2Set {
		data.UseCarTypeV2 = &builder.useCarTypeV2
	}
	if builder.budgetItemNameSet {
		data.BudgetItemName = &builder.budgetItemName
	}
	if builder.useTypeSet {
		data.UseType = &builder.useType
	}
	if builder.startPriceSet {
		data.StartPrice = &builder.startPrice
	}
	if builder.normalFeeSet {
		data.NormalFee = &builder.normalFee
	}
	if builder.lowSpeedFeeSet {
		data.LowSpeedFee = &builder.lowSpeedFee
	}
	if builder.emptyFeeSet {
		data.EmptyFee = &builder.emptyFee
	}
	if builder.nightFeeSet {
		data.NightFee = &builder.nightFee
	}
	if builder.responsibleCancelFeeSet {
		data.ResponsibleCancelFee = &builder.responsibleCancelFee
	}
	if builder.limitFeeSet {
		data.LimitFee = &builder.limitFee
	}
	if builder.limitPaySet {
		data.LimitPay = &builder.limitPay
	}
	if builder.normalTimeFeeSet {
		data.NormalTimeFee = &builder.normalTimeFee
	}
	if builder.dynamicPriceSet {
		data.DynamicPrice = &builder.dynamicPrice
	}
	if builder.tipFeeSet {
		data.TipFee = &builder.tipFee
	}
	if builder.bridgeFeeSet {
		data.BridgeFee = &builder.bridgeFee
	}
	if builder.highwayFeeSet {
		data.HighwayFee = &builder.highwayFee
	}
	if builder.parkFeeSet {
		data.ParkFee = &builder.parkFee
	}
	if builder.waitFeeSet {
		data.WaitFee = &builder.waitFee
	}
	if builder.incentiveFeeSet {
		data.IncentiveFee = &builder.incentiveFee
	}
	if builder.redPacketSet {
		data.RedPacket = &builder.redPacket
	}
	if builder.otherFeeSet {
		data.OtherFee = &builder.otherFee
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.refundPriceSet {
		data.RefundPrice = &builder.refundPrice
	}
	if builder.groupNameSet {
		data.GroupName = &builder.groupName
	}
	if builder.groupIdSet {
		data.GroupId = &builder.groupId
	}
	if builder.companyRealPreTaxPaySet {
		data.CompanyRealPreTaxPay = &builder.companyRealPreTaxPay
	}
	if builder.companyRealTaxPaySet {
		data.CompanyRealTaxPay = &builder.companyRealTaxPay
	}
	if builder.userPayTimeSet {
		data.UserPayTime = &builder.userPayTime
	}
	if builder.userPayTypeSet {
		data.UserPayType = &builder.userPayType
	}
	if builder.settleTypeSet {
		data.SettleType = &builder.settleType
	}
	if builder.subAccountNameSet {
		data.SubAccountName = &builder.subAccountName
	}
	if builder.belongBudgetCenterNameSet {
		data.BelongBudgetCenterName = &builder.belongBudgetCenterName
	}
	if builder.belongOutBudgetIdSet {
		data.BelongOutBudgetId = &builder.belongOutBudgetId
	}
	if builder.isCarpoolSuccessSet {
		data.IsCarpoolSuccess = &builder.isCarpoolSuccess
	}
	if builder.companyPayAbleSet {
		data.CompanyPayAble = &builder.companyPayAble
	}
	if builder.serviceTypeSet {
		data.ServiceType = &builder.serviceType
	}
	if builder.beforeApprovalResultSet {
		data.BeforeApprovalResult = &builder.beforeApprovalResult
	}
	if builder.timeSlotDiscountSet {
		data.TimeSlotDiscount = &builder.timeSlotDiscount
	}
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.periodSet {
		data.Period = &builder.period
	}
	if builder.excludingTaxPayPriceSet {
		data.ExcludingTaxPayPrice = &builder.excludingTaxPayPrice
	}
	if builder.taxPayPriceSet {
		data.TaxPayPrice = &builder.taxPayPrice
	}
	if builder.excludingTaxRefundPriceSet {
		data.ExcludingTaxRefundPrice = &builder.excludingTaxRefundPrice
	}
	if builder.taxRefundPriceSet {
		data.TaxRefundPrice = &builder.taxRefundPrice
	}
	if builder.isSensitiveSet {
		data.IsSensitive = &builder.isSensitive
	}
	if builder.sensitiveReasonSet {
		data.SensitiveReason = &builder.sensitiveReason
	}
	if builder.companyCardRealRefundSet {
		data.CompanyCardRealRefund = &builder.companyCardRealRefund
	}
	if builder.cancelTimeSet {
		data.CancelTime = &builder.cancelTime
	}
	if builder.crossCityFeeSet {
		data.CrossCityFee = &builder.crossCityFee
	}
	if builder.remoteAreaFeeSet {
		data.RemoteAreaFee = &builder.remoteAreaFee
	}
	if builder.useCarTypeNameSet {
		data.UseCarTypeName = &builder.useCarTypeName
	}
	if builder.subUseCarSrvSet {
		data.SubUseCarSrv = &builder.subUseCarSrv
	}
	if builder.energyConsumeFeeSet {
		data.EnergyConsumeFee = &builder.energyConsumeFee
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
	if builder.spsPickUpServiceFeeSet {
		data.SpsPickUpServiceFee = &builder.spsPickUpServiceFee
	}
	if builder.deductionTypeNameSet {
		data.DeductionTypeName = &builder.deductionTypeName
	}
	if builder.payChannelSet {
		data.PayChannel = &builder.payChannel
	}
	if builder.residentCityNamesSet {
		data.ResidentCityNames = &builder.residentCityNames
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
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.carTravelInfoSet {
		data.CarTravelInfo = &builder.carTravelInfo
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.parentInstitutionIdSet {
		data.ParentInstitutionId = &builder.parentInstitutionId
	}
	if builder.excludingServiceFeeSet {
		data.ExcludingServiceFee = &builder.excludingServiceFee
	}
	if builder.estimatePriceSet {
		data.EstimatePrice = &builder.estimatePrice
	}
	if builder.subAccountCompanyNameSet {
		data.SubAccountCompanyName = &builder.subAccountCompanyName
	}
	if builder.chargeTimeSet {
		data.ChargeTime = &builder.chargeTime
	}
	if builder.startCountyIdSet {
		data.StartCountyId = &builder.startCountyId
	}
	if builder.endCountyIdSet {
		data.EndCountyId = &builder.endCountyId
	}
	if builder.startCountyNameSet {
		data.StartCountyName = &builder.startCountyName
	}
	if builder.endCountyNameSet {
		data.EndCountyName = &builder.endCountyName
	}
	if builder.flagExtSet {
		data.FlagExt = &builder.flagExt
	}
	if builder.cWeightSet {
		data.CWeight = &builder.cWeight
	}
	if builder.co2EmissionsSet {
		data.Co2Emissions = &builder.co2Emissions
	}
	if builder.afterApproveMemberOneSet {
		data.AfterApproveMemberOne = &builder.afterApproveMemberOne
	}
	if builder.afterApproveMemberTwoSet {
		data.AfterApproveMemberTwo = &builder.afterApproveMemberTwo
	}
	if builder.afterApproveMemberThreeSet {
		data.AfterApproveMemberThree = &builder.afterApproveMemberThree
	}
	if builder.callPhoneCountryCodeSet {
		data.CallPhoneCountryCode = &builder.callPhoneCountryCode
	}
	if builder.passengerPhoneCountryCodeSet {
		data.PassengerPhoneCountryCode = &builder.passengerPhoneCountryCode
	}
	if builder.enterpriseCouponBatchSet {
		data.EnterpriseCouponBatch = &builder.enterpriseCouponBatch
	}
	if builder.enterpriseCouponNameSet {
		data.EnterpriseCouponName = &builder.enterpriseCouponName
	}
	if builder.lastApprovesSet {
		data.LastApproves = &builder.lastApproves
	}
	if builder.lastApprovalTimeSet {
		data.LastApprovalTime = &builder.lastApprovalTime
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
