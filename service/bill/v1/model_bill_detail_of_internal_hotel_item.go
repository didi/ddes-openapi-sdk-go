package v1

// BillDetailOfInternalHotelItem 已出账单 ~ 海外酒店返回参数,参考内部文档进行定义；
type BillDetailOfInternalHotelItem struct {
	BillId                     *int64   `json:"bill_id,omitempty"`                        // 账单id 根结点批次id
	SettleTime                 *string  `json:"settle_time,omitempty"`                    // 结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\" 入账时间（交易维度，支付或者退款）
	MemberId                   *string  `json:"member_id,omitempty"`                      // 入住人员工id
	PassengerMemberNumber      *string  `json:"passenger_member_number,omitempty"`        // 入住人工号
	PassengerName              *string  `json:"passenger_name,omitempty"`                 // 入住人员工姓名
	Department                 *string  `json:"department,omitempty"`                     // 入住人部门名称
	BudgetCenterName           *string  `json:"budget_center_name,omitempty"`             // 成本中心名称
	BudgetCenterId             *string  `json:"budget_center_id,omitempty"`               // 成本中心ID
	BudgetCenterParentPathName *string  `json:"budget_center_parent_path_name,omitempty"` // 所在成本中心 路径 示例：\"白全岭测试公司01>项目二测试\"
	OrderId                    *string  `json:"order_id,omitempty"`                       // 订单号
	OrderType                  *string  `json:"order_type,omitempty"`                     // 订单类型 枚举【预订、退款】
	OrderStatus                *string  `json:"order_status,omitempty"`                   // 订单状态 枚举【\"已提交\"、\"已提交预订\"、\"商家审核中\"、\"预订成功\"、\"已入住\"、\"已完成\"、\"已取消订单\"、\"已取消预订\"、\"预订未到\"】 部分退款对应的订单状态：已完成（提前离店）、已取消预订（取消有罚金）、预订未到（人没到收一天罚金）
	OrderSource                *string  `json:"order_source,omitempty"`                   // 订单来源 枚举【企业app、H5等】
	PayType                    *string  `json:"pay_type,omitempty"`                       // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	BookingMemberName          *string  `json:"booking_member_name,omitempty"`            // 预订人姓名
	BookingMemberNumber        *string  `json:"booking_member_number,omitempty"`          // 预定人员工编号
	BookingDepName             *string  `json:"booking_dep_name,omitempty"`               // 预订人部门
	BookingParentPathDepName   *string  `json:"booking_parent_path_dep_name,omitempty"`   // 预订人所在部门 示例：\"白全岭测试公司01\"
	BookingDepId               *string  `json:"booking_dep_id,omitempty"`                 // 预订人部门ID
	BookingDepCode             *string  `json:"booking_dep_code,omitempty"`               // 预订人部门编号
	BookingDate                *string  `json:"booking_date,omitempty"`                   // 预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	CheckInDate                *string  `json:"check_in_date,omitempty"`                  // 预订入住日期 格式：\"yyyy-MM-dd\" 示例：\"2023-06-15\"
	OriginalCheckOutDate       *string  `json:"original_check_out_date,omitempty"`        // 预订离店日期 格式：\"yyyy-MM-dd\" 示例：\"2023-06-18\"
	CheckOutDate               *string  `json:"check_out_date,omitempty"`                 // 实际退房时间 格式：\"yyyy-MM-dd\" 示例：\"2023-06-18\"
	CheckInNames               *string  `json:"check_in_names,omitempty"`                 // 入住人姓名 (冗余字段)
	CheckinNamesCn             *string  `json:"checkin_names_cn,omitempty"`               // 入住此酒店订单的人员中文姓名（若多人入住，则使用逗号分隔）
	CheckInMemberNumbers       *string  `json:"check_in_member_numbers,omitempty"`        // 入住人员工编号
	CheckInMemberDepartments   *string  `json:"check_in_member_departments,omitempty"`    // 入住人部门名称 (冗余字段)
	CheckInDepName             *string  `json:"check_in_dep_name,omitempty"`              // 入住人部门 (冗余字段)
	CheckInPathDepName         *string  `json:"check_in_path_dep_name,omitempty"`         // 入住人所在部门 示例：\"白全岭测试公司01\"
	CheckInDepId               *string  `json:"check_in_dep_id,omitempty"`                // 入住人部门id
	CheckInDepCode             *string  `json:"check_in_dep_code,omitempty"`              // 入住人部门编号
	CityName                   *string  `json:"city_name,omitempty"`                      // 城市名称
	HotelType                  *string  `json:"hotel_type,omitempty"`                     // 酒店类型 枚举【\"酒店协议价\"、\"会员酒店\"】
	HotelName                  *string  `json:"hotel_name,omitempty"`                     // 酒店名称
	Rooms                      *string  `json:"rooms,omitempty"`                          // 房型 示例：高级大床房、豪华双床房
	RoomNum                    *float32 `json:"room_num,omitempty"`                       // 预订房间数 1个人住:1 2个人住:0.5 3个人住:0.33、0.33、0.34
	BookDayNumber              *int32   `json:"book_day_number,omitempty"`                // 预订天数
	TotalNights                *float32 `json:"total_nights,omitempty"`                   // 合计间夜数（=预订天数*预订房间数）
	CompanyRealAmount          *float32 `json:"company_real_amount,omitempty"`            // 企业实付金额（=房费+平台使用费）
	CompanyPayType             *string  `json:"company_pay_type,omitempty"`               // 交易类型 枚举【”消费“、”退款“】
	PersonalRealPay            *float32 `json:"personal_real_pay,omitempty"`              // 个人实付金额（入住人包含预订人，个人费用挂到预订人；入住人不包含预订人，个人费用挂到订单的第一个人）
	PersonalRealRefund         *float32 `json:"personal_real_refund,omitempty"`           // 个人实际退款金额
	ServiceFee                 *float32 `json:"service_fee,omitempty"`                    // 平台使用费
	CostDiscountBefore         *float32 `json:"cost_discount_before,omitempty"`           // 房费总价
	Amount                     *float32 `json:"amount,omitempty"`                         // 订单总金额（平台使用费同步收，等于企业实付金额；平台使用费异步收，等于房费总价）
	RcLowPrice                 *string  `json:"rc_low_price,omitempty"`                   // 酒店房费RC 房间费用超标原因，示例：为客户预订/招待客户
	RcAgreedPrice              *string  `json:"rc_agreed_price,omitempty"`                // 酒店协议价RC 预订协议酒店超标原因
	RcBook                     *string  `json:"rc_book,omitempty"`                        // 酒店提前预订天数RC 预订天数超标原因
	RcLevel                    *string  `json:"rc_level,omitempty"`                       // 酒店星级RC 预订酒店星级超标原因
	TravelPurpose              *string  `json:"travel_purpose,omitempty"`                 // 出行人目的
	TravelDescription          *string  `json:"travel_description,omitempty"`             // 出行人描述
	PersonalCouponCost         *float32 `json:"personal_coupon_cost,omitempty"`           // 个人用券抵扣金额
	CompanyCouponCost          *float32 `json:"company_coupon_cost,omitempty"`            // 企业用券抵扣金额
	ApprovalFirst              *string  `json:"approval_first,omitempty"`                 // 一级审批人
	ApprovalSecond             *string  `json:"approval_second,omitempty"`                // 二级审批人
	ApprovalThird              *string  `json:"approval_third,omitempty"`                 // 三级审批人
	ApprovalHistory            *string  `json:"approval_history,omitempty"`               // 预定/超标情况下的审批历史
	RequisitionId              *string  `json:"requisition_id,omitempty"`                 // 出差申请单号
	OutRequisitionId           *string  `json:"out_requisition_id,omitempty"`             // 外部审批单id
	IsVat                      *string  `json:"is_vat,omitempty"`                         // 是否增值税 枚举【“是”，“否”】
	TaxRatio                   *string  `json:"tax_ratio,omitempty"`                      // 税率
	ExtraInfo                  *string  `json:"extra_info,omitempty"`                     // 审批单自定义
	Remark                     *string  `json:"remark,omitempty"`                         // 备注
	TripReason                 *string  `json:"trip_reason,omitempty"`                    // 出差事由
	UniqueKey                  *string  `json:"unique_key,omitempty"`                     // 唯一键 此条明细的唯一键
	LegalEntityId              *int64   `json:"legal_entity_id,omitempty"`                // 预定人开票子公司id
	LegalEntityName            *string  `json:"legal_entity_name,omitempty"`              // 预定人开票子公司名称
	PassengerLegalEntityId     *int64   `json:"passenger_legal_entity_id,omitempty"`      // 入住人开票子公司id
	PassengerLegalEntityName   *string  `json:"passenger_legal_entity_name,omitempty"`    // 入住人开票子公司名称
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
	PassengerMemberId          *int64   `json:"passenger_member_id,omitempty"`            // 入住人ID 如：123432
	Reason                     *string  `json:"reason,omitempty"`                         // 取消原因 如：临时取消
	ProjectExtInfo             *string  `json:"project_ext_info,omitempty"`               // 项目自定义
	ApprovalNormalHistory      *string  `json:"approval_normal_history,omitempty"`        // 正常差旅申请的审批历史
	RankName                   *string  `json:"rank_name,omitempty"`                      // 职级 如：\"员工\"
	CityId                     *string  `json:"city_id,omitempty"`                        // 入住城市id
	EsDiffTaxFee               *float64 `json:"es_diff_tax_fee,omitempty"`                // 税点服务费 如：1.00
	InvoiceTaxRatio            *string  `json:"invoice_tax_ratio,omitempty"`              // 发票税率 如：\"1%\"
	InvoiceType                *string  `json:"invoice_type,omitempty"`                   // 发票类型 如：\"普票\"
	OutLegalEntityId           *string  `json:"out_legal_entity_id,omitempty"`            // 外部公司主体编号
	CountryNameCn              *string  `json:"country_name_cn,omitempty"`                // 国家
	Currency                   *string  `json:"currency,omitempty"`                       // 币种
	TaxSupplementServiceFee    *float64 `json:"tax_supplement_service_fee,omitempty"`     // 补税服务费
	InstitutionId              *int64   `json:"institution_id,omitempty"`                 // 制度ID
	BookingMemberId            *int64   `json:"booking_member_id,omitempty"`              // 预定人ID
	ExcludingServiceFee        *float64 `json:"excluding_service_fee,omitempty"`          // 企业实付（不包含平台使用费）
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

type BillDetailOfInternalHotelItemBuilder struct {
	billId                        int64 // 账单id 根结点批次id
	billIdSet                     bool
	settleTime                    string // 结算时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-05-27 22:00:47\" 入账时间（交易维度，支付或者退款）
	settleTimeSet                 bool
	memberId                      string // 入住人员工id
	memberIdSet                   bool
	passengerMemberNumber         string // 入住人工号
	passengerMemberNumberSet      bool
	passengerName                 string // 入住人员工姓名
	passengerNameSet              bool
	department                    string // 入住人部门名称
	departmentSet                 bool
	budgetCenterName              string // 成本中心名称
	budgetCenterNameSet           bool
	budgetCenterId                string // 成本中心ID
	budgetCenterIdSet             bool
	budgetCenterParentPathName    string // 所在成本中心 路径 示例：\"白全岭测试公司01>项目二测试\"
	budgetCenterParentPathNameSet bool
	orderId                       string // 订单号
	orderIdSet                    bool
	orderType                     string // 订单类型 枚举【预订、退款】
	orderTypeSet                  bool
	orderStatus                   string // 订单状态 枚举【\"已提交\"、\"已提交预订\"、\"商家审核中\"、\"预订成功\"、\"已入住\"、\"已完成\"、\"已取消订单\"、\"已取消预订\"、\"预订未到\"】 部分退款对应的订单状态：已完成（提前离店）、已取消预订（取消有罚金）、预订未到（人没到收一天罚金）
	orderStatusSet                bool
	orderSource                   string // 订单来源 枚举【企业app、H5等】
	orderSourceSet                bool
	payType                       string // 支付方式 枚举【\"企业支付\"、\"个人支付\"、\"混合支付\"、\"企业钱包支付\"】
	payTypeSet                    bool
	bookingMemberName             string // 预订人姓名
	bookingMemberNameSet          bool
	bookingMemberNumber           string // 预定人员工编号
	bookingMemberNumberSet        bool
	bookingDepName                string // 预订人部门
	bookingDepNameSet             bool
	bookingParentPathDepName      string // 预订人所在部门 示例：\"白全岭测试公司01\"
	bookingParentPathDepNameSet   bool
	bookingDepId                  string // 预订人部门ID
	bookingDepIdSet               bool
	bookingDepCode                string // 预订人部门编号
	bookingDepCodeSet             bool
	bookingDate                   string // 预订时间 格式：\"yyyy-MM-dd HH:mm:ss\" 示例： \"2023-06-13 15:58:49\"
	bookingDateSet                bool
	checkInDate                   string // 预订入住日期 格式：\"yyyy-MM-dd\" 示例：\"2023-06-15\"
	checkInDateSet                bool
	originalCheckOutDate          string // 预订离店日期 格式：\"yyyy-MM-dd\" 示例：\"2023-06-18\"
	originalCheckOutDateSet       bool
	checkOutDate                  string // 实际退房时间 格式：\"yyyy-MM-dd\" 示例：\"2023-06-18\"
	checkOutDateSet               bool
	checkInNames                  string // 入住人姓名 (冗余字段)
	checkInNamesSet               bool
	checkinNamesCn                string // 入住此酒店订单的人员中文姓名（若多人入住，则使用逗号分隔）
	checkinNamesCnSet             bool
	checkInMemberNumbers          string // 入住人员工编号
	checkInMemberNumbersSet       bool
	checkInMemberDepartments      string // 入住人部门名称 (冗余字段)
	checkInMemberDepartmentsSet   bool
	checkInDepName                string // 入住人部门 (冗余字段)
	checkInDepNameSet             bool
	checkInPathDepName            string // 入住人所在部门 示例：\"白全岭测试公司01\"
	checkInPathDepNameSet         bool
	checkInDepId                  string // 入住人部门id
	checkInDepIdSet               bool
	checkInDepCode                string // 入住人部门编号
	checkInDepCodeSet             bool
	cityName                      string // 城市名称
	cityNameSet                   bool
	hotelType                     string // 酒店类型 枚举【\"酒店协议价\"、\"会员酒店\"】
	hotelTypeSet                  bool
	hotelName                     string // 酒店名称
	hotelNameSet                  bool
	rooms                         string // 房型 示例：高级大床房、豪华双床房
	roomsSet                      bool
	roomNum                       float32 // 预订房间数 1个人住:1 2个人住:0.5 3个人住:0.33、0.33、0.34
	roomNumSet                    bool
	bookDayNumber                 int32 // 预订天数
	bookDayNumberSet              bool
	totalNights                   float32 // 合计间夜数（=预订天数*预订房间数）
	totalNightsSet                bool
	companyRealAmount             float32 // 企业实付金额（=房费+平台使用费）
	companyRealAmountSet          bool
	companyPayType                string // 交易类型 枚举【”消费“、”退款“】
	companyPayTypeSet             bool
	personalRealPay               float32 // 个人实付金额（入住人包含预订人，个人费用挂到预订人；入住人不包含预订人，个人费用挂到订单的第一个人）
	personalRealPaySet            bool
	personalRealRefund            float32 // 个人实际退款金额
	personalRealRefundSet         bool
	serviceFee                    float32 // 平台使用费
	serviceFeeSet                 bool
	costDiscountBefore            float32 // 房费总价
	costDiscountBeforeSet         bool
	amount                        float32 // 订单总金额（平台使用费同步收，等于企业实付金额；平台使用费异步收，等于房费总价）
	amountSet                     bool
	rcLowPrice                    string // 酒店房费RC 房间费用超标原因，示例：为客户预订/招待客户
	rcLowPriceSet                 bool
	rcAgreedPrice                 string // 酒店协议价RC 预订协议酒店超标原因
	rcAgreedPriceSet              bool
	rcBook                        string // 酒店提前预订天数RC 预订天数超标原因
	rcBookSet                     bool
	rcLevel                       string // 酒店星级RC 预订酒店星级超标原因
	rcLevelSet                    bool
	travelPurpose                 string // 出行人目的
	travelPurposeSet              bool
	travelDescription             string // 出行人描述
	travelDescriptionSet          bool
	personalCouponCost            float32 // 个人用券抵扣金额
	personalCouponCostSet         bool
	companyCouponCost             float32 // 企业用券抵扣金额
	companyCouponCostSet          bool
	approvalFirst                 string // 一级审批人
	approvalFirstSet              bool
	approvalSecond                string // 二级审批人
	approvalSecondSet             bool
	approvalThird                 string // 三级审批人
	approvalThirdSet              bool
	approvalHistory               string // 预定/超标情况下的审批历史
	approvalHistorySet            bool
	requisitionId                 string // 出差申请单号
	requisitionIdSet              bool
	outRequisitionId              string // 外部审批单id
	outRequisitionIdSet           bool
	isVat                         string // 是否增值税 枚举【“是”，“否”】
	isVatSet                      bool
	taxRatio                      string // 税率
	taxRatioSet                   bool
	extraInfo                     string // 审批单自定义
	extraInfoSet                  bool
	remark                        string // 备注
	remarkSet                     bool
	tripReason                    string // 出差事由
	tripReasonSet                 bool
	uniqueKey                     string // 唯一键 此条明细的唯一键
	uniqueKeySet                  bool
	legalEntityId                 int64 // 预定人开票子公司id
	legalEntityIdSet              bool
	legalEntityName               string // 预定人开票子公司名称
	legalEntityNameSet            bool
	passengerLegalEntityId        int64 // 入住人开票子公司id
	passengerLegalEntityIdSet     bool
	passengerLegalEntityName      string // 入住人开票子公司名称
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
	passengerMemberId             int64 // 入住人ID 如：123432
	passengerMemberIdSet          bool
	reason                        string // 取消原因 如：临时取消
	reasonSet                     bool
	projectExtInfo                string // 项目自定义
	projectExtInfoSet             bool
	approvalNormalHistory         string // 正常差旅申请的审批历史
	approvalNormalHistorySet      bool
	rankName                      string // 职级 如：\"员工\"
	rankNameSet                   bool
	cityId                        string // 入住城市id
	cityIdSet                     bool
	esDiffTaxFee                  float64 // 税点服务费 如：1.00
	esDiffTaxFeeSet               bool
	invoiceTaxRatio               string // 发票税率 如：\"1%\"
	invoiceTaxRatioSet            bool
	invoiceType                   string // 发票类型 如：\"普票\"
	invoiceTypeSet                bool
	outLegalEntityId              string // 外部公司主体编号
	outLegalEntityIdSet           bool
	countryNameCn                 string // 国家
	countryNameCnSet              bool
	currency                      string // 币种
	currencySet                   bool
	taxSupplementServiceFee       float64 // 补税服务费
	taxSupplementServiceFeeSet    bool
	institutionId                 int64 // 制度ID
	institutionIdSet              bool
	bookingMemberId               int64 // 预定人ID
	bookingMemberIdSet            bool
	excludingServiceFee           float64 // 企业实付（不包含平台使用费）
	excludingServiceFeeSet        bool
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

func NewBillDetailOfInternalHotelItemBuilder() *BillDetailOfInternalHotelItemBuilder {
	return &BillDetailOfInternalHotelItemBuilder{}
}
func (builder *BillDetailOfInternalHotelItemBuilder) BillId(billId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) SettleTime(settleTime string) *BillDetailOfInternalHotelItemBuilder {
	builder.settleTime = settleTime
	builder.settleTimeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) MemberId(memberId string) *BillDetailOfInternalHotelItemBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillDetailOfInternalHotelItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PassengerName(passengerName string) *BillDetailOfInternalHotelItemBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Department(department string) *BillDetailOfInternalHotelItemBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BudgetCenterName(budgetCenterName string) *BillDetailOfInternalHotelItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BudgetCenterId(budgetCenterId string) *BillDetailOfInternalHotelItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BudgetCenterParentPathName(budgetCenterParentPathName string) *BillDetailOfInternalHotelItemBuilder {
	builder.budgetCenterParentPathName = budgetCenterParentPathName
	builder.budgetCenterParentPathNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OrderId(orderId string) *BillDetailOfInternalHotelItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OrderType(orderType string) *BillDetailOfInternalHotelItemBuilder {
	builder.orderType = orderType
	builder.orderTypeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OrderStatus(orderStatus string) *BillDetailOfInternalHotelItemBuilder {
	builder.orderStatus = orderStatus
	builder.orderStatusSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OrderSource(orderSource string) *BillDetailOfInternalHotelItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PayType(payType string) *BillDetailOfInternalHotelItemBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingMemberName(bookingMemberName string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingMemberNumber(bookingMemberNumber string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingDepName(bookingDepName string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingParentPathDepName(bookingParentPathDepName string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingParentPathDepName = bookingParentPathDepName
	builder.bookingParentPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingDepId(bookingDepId string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingDepId = bookingDepId
	builder.bookingDepIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingDepCode(bookingDepCode string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingDate(bookingDate string) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingDate = bookingDate
	builder.bookingDateSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInDate(checkInDate string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInDate = checkInDate
	builder.checkInDateSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OriginalCheckOutDate(originalCheckOutDate string) *BillDetailOfInternalHotelItemBuilder {
	builder.originalCheckOutDate = originalCheckOutDate
	builder.originalCheckOutDateSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckOutDate(checkOutDate string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkOutDate = checkOutDate
	builder.checkOutDateSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInNames(checkInNames string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInNames = checkInNames
	builder.checkInNamesSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckinNamesCn(checkinNamesCn string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkinNamesCn = checkinNamesCn
	builder.checkinNamesCnSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInMemberNumbers(checkInMemberNumbers string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInMemberNumbers = checkInMemberNumbers
	builder.checkInMemberNumbersSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInMemberDepartments(checkInMemberDepartments string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInMemberDepartments = checkInMemberDepartments
	builder.checkInMemberDepartmentsSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInDepName(checkInDepName string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInDepName = checkInDepName
	builder.checkInDepNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInPathDepName(checkInPathDepName string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInPathDepName = checkInPathDepName
	builder.checkInPathDepNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInDepId(checkInDepId string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInDepId = checkInDepId
	builder.checkInDepIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CheckInDepCode(checkInDepCode string) *BillDetailOfInternalHotelItemBuilder {
	builder.checkInDepCode = checkInDepCode
	builder.checkInDepCodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CityName(cityName string) *BillDetailOfInternalHotelItemBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) HotelType(hotelType string) *BillDetailOfInternalHotelItemBuilder {
	builder.hotelType = hotelType
	builder.hotelTypeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) HotelName(hotelName string) *BillDetailOfInternalHotelItemBuilder {
	builder.hotelName = hotelName
	builder.hotelNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Rooms(rooms string) *BillDetailOfInternalHotelItemBuilder {
	builder.rooms = rooms
	builder.roomsSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RoomNum(roomNum float32) *BillDetailOfInternalHotelItemBuilder {
	builder.roomNum = roomNum
	builder.roomNumSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookDayNumber(bookDayNumber int32) *BillDetailOfInternalHotelItemBuilder {
	builder.bookDayNumber = bookDayNumber
	builder.bookDayNumberSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TotalNights(totalNights float32) *BillDetailOfInternalHotelItemBuilder {
	builder.totalNights = totalNights
	builder.totalNightsSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CompanyRealAmount(companyRealAmount float32) *BillDetailOfInternalHotelItemBuilder {
	builder.companyRealAmount = companyRealAmount
	builder.companyRealAmountSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CompanyPayType(companyPayType string) *BillDetailOfInternalHotelItemBuilder {
	builder.companyPayType = companyPayType
	builder.companyPayTypeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PersonalRealPay(personalRealPay float32) *BillDetailOfInternalHotelItemBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PersonalRealRefund(personalRealRefund float32) *BillDetailOfInternalHotelItemBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ServiceFee(serviceFee float32) *BillDetailOfInternalHotelItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CostDiscountBefore(costDiscountBefore float32) *BillDetailOfInternalHotelItemBuilder {
	builder.costDiscountBefore = costDiscountBefore
	builder.costDiscountBeforeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Amount(amount float32) *BillDetailOfInternalHotelItemBuilder {
	builder.amount = amount
	builder.amountSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RcLowPrice(rcLowPrice string) *BillDetailOfInternalHotelItemBuilder {
	builder.rcLowPrice = rcLowPrice
	builder.rcLowPriceSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RcAgreedPrice(rcAgreedPrice string) *BillDetailOfInternalHotelItemBuilder {
	builder.rcAgreedPrice = rcAgreedPrice
	builder.rcAgreedPriceSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RcBook(rcBook string) *BillDetailOfInternalHotelItemBuilder {
	builder.rcBook = rcBook
	builder.rcBookSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RcLevel(rcLevel string) *BillDetailOfInternalHotelItemBuilder {
	builder.rcLevel = rcLevel
	builder.rcLevelSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TravelPurpose(travelPurpose string) *BillDetailOfInternalHotelItemBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TravelDescription(travelDescription string) *BillDetailOfInternalHotelItemBuilder {
	builder.travelDescription = travelDescription
	builder.travelDescriptionSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PersonalCouponCost(personalCouponCost float32) *BillDetailOfInternalHotelItemBuilder {
	builder.personalCouponCost = personalCouponCost
	builder.personalCouponCostSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CompanyCouponCost(companyCouponCost float32) *BillDetailOfInternalHotelItemBuilder {
	builder.companyCouponCost = companyCouponCost
	builder.companyCouponCostSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ApprovalFirst(approvalFirst string) *BillDetailOfInternalHotelItemBuilder {
	builder.approvalFirst = approvalFirst
	builder.approvalFirstSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ApprovalSecond(approvalSecond string) *BillDetailOfInternalHotelItemBuilder {
	builder.approvalSecond = approvalSecond
	builder.approvalSecondSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ApprovalThird(approvalThird string) *BillDetailOfInternalHotelItemBuilder {
	builder.approvalThird = approvalThird
	builder.approvalThirdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ApprovalHistory(approvalHistory string) *BillDetailOfInternalHotelItemBuilder {
	builder.approvalHistory = approvalHistory
	builder.approvalHistorySet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RequisitionId(requisitionId string) *BillDetailOfInternalHotelItemBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OutRequisitionId(outRequisitionId string) *BillDetailOfInternalHotelItemBuilder {
	builder.outRequisitionId = outRequisitionId
	builder.outRequisitionIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) IsVat(isVat string) *BillDetailOfInternalHotelItemBuilder {
	builder.isVat = isVat
	builder.isVatSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TaxRatio(taxRatio string) *BillDetailOfInternalHotelItemBuilder {
	builder.taxRatio = taxRatio
	builder.taxRatioSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExtraInfo(extraInfo string) *BillDetailOfInternalHotelItemBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Remark(remark string) *BillDetailOfInternalHotelItemBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TripReason(tripReason string) *BillDetailOfInternalHotelItemBuilder {
	builder.tripReason = tripReason
	builder.tripReasonSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) UniqueKey(uniqueKey string) *BillDetailOfInternalHotelItemBuilder {
	builder.uniqueKey = uniqueKey
	builder.uniqueKeySet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) LegalEntityId(legalEntityId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) LegalEntityName(legalEntityName string) *BillDetailOfInternalHotelItemBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PassengerLegalEntityId(passengerLegalEntityId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *BillDetailOfInternalHotelItemBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo01(exInfo01 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo02(exInfo02 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo03(exInfo03 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo04(exInfo04 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo05(exInfo05 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo06(exInfo06 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo07(exInfo07 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo08(exInfo08 string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) InstitutionName(institutionName string) *BillDetailOfInternalHotelItemBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BudgetCenterCode(budgetCenterCode string) *BillDetailOfInternalHotelItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) PassengerMemberId(passengerMemberId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Reason(reason string) *BillDetailOfInternalHotelItemBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ProjectExtInfo(projectExtInfo string) *BillDetailOfInternalHotelItemBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ApprovalNormalHistory(approvalNormalHistory string) *BillDetailOfInternalHotelItemBuilder {
	builder.approvalNormalHistory = approvalNormalHistory
	builder.approvalNormalHistorySet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) RankName(rankName string) *BillDetailOfInternalHotelItemBuilder {
	builder.rankName = rankName
	builder.rankNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CityId(cityId string) *BillDetailOfInternalHotelItemBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) EsDiffTaxFee(esDiffTaxFee float64) *BillDetailOfInternalHotelItemBuilder {
	builder.esDiffTaxFee = esDiffTaxFee
	builder.esDiffTaxFeeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) InvoiceTaxRatio(invoiceTaxRatio string) *BillDetailOfInternalHotelItemBuilder {
	builder.invoiceTaxRatio = invoiceTaxRatio
	builder.invoiceTaxRatioSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) InvoiceType(invoiceType string) *BillDetailOfInternalHotelItemBuilder {
	builder.invoiceType = invoiceType
	builder.invoiceTypeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) OutLegalEntityId(outLegalEntityId string) *BillDetailOfInternalHotelItemBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) CountryNameCn(countryNameCn string) *BillDetailOfInternalHotelItemBuilder {
	builder.countryNameCn = countryNameCn
	builder.countryNameCnSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) Currency(currency string) *BillDetailOfInternalHotelItemBuilder {
	builder.currency = currency
	builder.currencySet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) TaxSupplementServiceFee(taxSupplementServiceFee float64) *BillDetailOfInternalHotelItemBuilder {
	builder.taxSupplementServiceFee = taxSupplementServiceFee
	builder.taxSupplementServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) InstitutionId(institutionId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) BookingMemberId(bookingMemberId int64) *BillDetailOfInternalHotelItemBuilder {
	builder.bookingMemberId = bookingMemberId
	builder.bookingMemberIdSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExcludingServiceFee(excludingServiceFee float64) *BillDetailOfInternalHotelItemBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExtendInfo(extendInfo string) *BillDetailOfInternalHotelItemBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) SubAccountCompanyName(subAccountCompanyName string) *BillDetailOfInternalHotelItemBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo01Code(exInfo01Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo02Code(exInfo02Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo03Code(exInfo03Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo04Code(exInfo04Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo05Code(exInfo05Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo06Code(exInfo06Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo07Code(exInfo07Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *BillDetailOfInternalHotelItemBuilder) ExInfo08Code(exInfo08Code string) *BillDetailOfInternalHotelItemBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}

func (builder *BillDetailOfInternalHotelItemBuilder) Build() *BillDetailOfInternalHotelItem {
	data := &BillDetailOfInternalHotelItem{}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.settleTimeSet {
		data.SettleTime = &builder.settleTime
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
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
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.budgetCenterParentPathNameSet {
		data.BudgetCenterParentPathName = &builder.budgetCenterParentPathName
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.orderTypeSet {
		data.OrderType = &builder.orderType
	}
	if builder.orderStatusSet {
		data.OrderStatus = &builder.orderStatus
	}
	if builder.orderSourceSet {
		data.OrderSource = &builder.orderSource
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
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
	if builder.bookingDepIdSet {
		data.BookingDepId = &builder.bookingDepId
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.bookingDateSet {
		data.BookingDate = &builder.bookingDate
	}
	if builder.checkInDateSet {
		data.CheckInDate = &builder.checkInDate
	}
	if builder.originalCheckOutDateSet {
		data.OriginalCheckOutDate = &builder.originalCheckOutDate
	}
	if builder.checkOutDateSet {
		data.CheckOutDate = &builder.checkOutDate
	}
	if builder.checkInNamesSet {
		data.CheckInNames = &builder.checkInNames
	}
	if builder.checkinNamesCnSet {
		data.CheckinNamesCn = &builder.checkinNamesCn
	}
	if builder.checkInMemberNumbersSet {
		data.CheckInMemberNumbers = &builder.checkInMemberNumbers
	}
	if builder.checkInMemberDepartmentsSet {
		data.CheckInMemberDepartments = &builder.checkInMemberDepartments
	}
	if builder.checkInDepNameSet {
		data.CheckInDepName = &builder.checkInDepName
	}
	if builder.checkInPathDepNameSet {
		data.CheckInPathDepName = &builder.checkInPathDepName
	}
	if builder.checkInDepIdSet {
		data.CheckInDepId = &builder.checkInDepId
	}
	if builder.checkInDepCodeSet {
		data.CheckInDepCode = &builder.checkInDepCode
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.hotelTypeSet {
		data.HotelType = &builder.hotelType
	}
	if builder.hotelNameSet {
		data.HotelName = &builder.hotelName
	}
	if builder.roomsSet {
		data.Rooms = &builder.rooms
	}
	if builder.roomNumSet {
		data.RoomNum = &builder.roomNum
	}
	if builder.bookDayNumberSet {
		data.BookDayNumber = &builder.bookDayNumber
	}
	if builder.totalNightsSet {
		data.TotalNights = &builder.totalNights
	}
	if builder.companyRealAmountSet {
		data.CompanyRealAmount = &builder.companyRealAmount
	}
	if builder.companyPayTypeSet {
		data.CompanyPayType = &builder.companyPayType
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.costDiscountBeforeSet {
		data.CostDiscountBefore = &builder.costDiscountBefore
	}
	if builder.amountSet {
		data.Amount = &builder.amount
	}
	if builder.rcLowPriceSet {
		data.RcLowPrice = &builder.rcLowPrice
	}
	if builder.rcAgreedPriceSet {
		data.RcAgreedPrice = &builder.rcAgreedPrice
	}
	if builder.rcBookSet {
		data.RcBook = &builder.rcBook
	}
	if builder.rcLevelSet {
		data.RcLevel = &builder.rcLevel
	}
	if builder.travelPurposeSet {
		data.TravelPurpose = &builder.travelPurpose
	}
	if builder.travelDescriptionSet {
		data.TravelDescription = &builder.travelDescription
	}
	if builder.personalCouponCostSet {
		data.PersonalCouponCost = &builder.personalCouponCost
	}
	if builder.companyCouponCostSet {
		data.CompanyCouponCost = &builder.companyCouponCost
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
	if builder.approvalHistorySet {
		data.ApprovalHistory = &builder.approvalHistory
	}
	if builder.requisitionIdSet {
		data.RequisitionId = &builder.requisitionId
	}
	if builder.outRequisitionIdSet {
		data.OutRequisitionId = &builder.outRequisitionId
	}
	if builder.isVatSet {
		data.IsVat = &builder.isVat
	}
	if builder.taxRatioSet {
		data.TaxRatio = &builder.taxRatio
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.tripReasonSet {
		data.TripReason = &builder.tripReason
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
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.approvalNormalHistorySet {
		data.ApprovalNormalHistory = &builder.approvalNormalHistory
	}
	if builder.rankNameSet {
		data.RankName = &builder.rankName
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.esDiffTaxFeeSet {
		data.EsDiffTaxFee = &builder.esDiffTaxFee
	}
	if builder.invoiceTaxRatioSet {
		data.InvoiceTaxRatio = &builder.invoiceTaxRatio
	}
	if builder.invoiceTypeSet {
		data.InvoiceType = &builder.invoiceType
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.countryNameCnSet {
		data.CountryNameCn = &builder.countryNameCn
	}
	if builder.currencySet {
		data.Currency = &builder.currency
	}
	if builder.taxSupplementServiceFeeSet {
		data.TaxSupplementServiceFee = &builder.taxSupplementServiceFee
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.bookingMemberIdSet {
		data.BookingMemberId = &builder.bookingMemberId
	}
	if builder.excludingServiceFeeSet {
		data.ExcludingServiceFee = &builder.excludingServiceFee
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
