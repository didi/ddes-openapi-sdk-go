package v1

// BillListItemOfWangYC 已出账单 ~ 网约车返回参数,参考内部文档进行定义；[按文档整理，缺失字段太多，暂时 屏蔽] message BillListItemOfWangYC {   string is_current_term = 1 [(openapi.v3.property) = {description: \"账期 枚举【本期、往期】\"}];   int64 order_id = 2 [(openapi.v3.property) = {description: \"订单id -企业级内部\" type:\"integer\" format:\"int64\"}];   int64 out_order_id = 3 [(openapi.v3.property) = {description: \"订单id -网约车\" type:\"integer\" format:\"int64\"}];   string order_source = 4 [(openapi.v3.property) = {description: \"订单来源 枚举【企业app、H5等】\"}];   string status = 5 [(openapi.v3.property) = {description: \"订单状态 枚举【全部支付、部分支付、超时取消扣费、全部退款、部分退款】\"}];   string rule = 6 [(openapi.v3.property) = {description: \"用车类型 枚举【出租车、专车、快车、代驾、豪华车、私车公用、同时呼叫】\"}];   string require_level = 7 [(openapi.v3.property) = {description: \"车型等级 枚举【专车-舒适型、专车-豪华型、快车-普通型、豪华车-奔驰S级 等】\"}];   string member_name = 8 [(openapi.v3.property) = {description: \"预定人姓名\"}];   string employee_number = 9 [(openapi.v3.property) = {description: \"预定人工号\"}];   string member_mail = 10 [(openapi.v3.property) = {description: \"预定人邮箱\"}];   string call_phone = 11 [(openapi.v3.property) = {description: \"预定人手机号\"}];   string department_name = 12 [(openapi.v3.property) = {description: \"预定人部门路径 如：\\\"jrm测试公司2023>同一级\\\"\"}];   string out_budget_id = 13 [(openapi.v3.property) = {description: \"外部成本中心code\"}];   string department = 14 [(openapi.v3.property) = {description: \"预定人员工部门 取member department 字段（旧）\"}];   string branch_name = 15 [(openapi.v3.property) = {description: \"预定人所在分公司名称\"}];   string passenger_member_number = 16 [(openapi.v3.property) = {description: \"乘车人工号\"}];   string passenger_name = 17 [(openapi.v3.property) = {description: \"乘车人姓名\"}];   string passenger_phone = 18 [(openapi.v3.property) = {description: \"乘车人电话\"}];   double total_price = 19 [(openapi.v3.property) = {description: \"订单总金额\"}];   double company_real_pay = 20 [(openapi.v3.property) = {description: \"企业实付金额\"}];   double personal_real_pay = 21 [(openapi.v3.property) = {description: \"个人实付金额\"}];   double real_voucher_pay = 22 [(openapi.v3.property) = {description: \"代金券(实际使用金额)\"}];   double cost = 23 [(openapi.v3.property) = {description: \"专车订单金额\"}];   double discount_after_price = 24 [(openapi.v3.property) = {description: \"服务费金额\"}];   double company_real_refund = 25 [(openapi.v3.property) = {description: \"企业实退金额\"}];   string type = 26 [(openapi.v3.property) = {description: \"订单类型 枚举【实时、预约、调账：订单调增、调账：订单调减、调账：主体调增、调账：主体调减、订单计入下一期】\"}];   string create_time = 27 [(openapi.v3.property) = {description: \"订单创建时间\"}];   string departure_dime = 28 [(openapi.v3.property) = {description: \"出发时间\"}];   string meetTime = 29 [(openapi.v3.property) = {description: \"接驾时间\"}];   string begin_charge_time = 30 [(openapi.v3.property) = {description: \"开始计费时间\"}];   string finish_time = 31 [(openapi.v3.property) = {description: \"完单时间\"}];   string pay_time = 32 [(openapi.v3.property) = {description: \"支付时间\"}];   string refund_time = 33 [(openapi.v3.property) = {description: \"退款时间\"}];   string city = 34 [(openapi.v3.property) = {description: \"城市\"}];   string destination_city = 35 [(openapi.v3.property) = {description: \"订单到达城市\"}];   double normal_distance = 36 [(openapi.v3.property) = {description: \"实际用车里程\"}];   string start_address = 37 [(openapi.v3.property) = {description: \"出发地地址\"}];   string end_address = 38 [(openapi.v3.property) = {description: \"目的地地址\"}];   string start_name = 39 [(openapi.v3.property) = {description: \"出发地名称\"}];   string end_name = 40 [(openapi.v3.property) = {description: \"目的地名称\"}];   string actual_start_name = 41 [(openapi.v3.property) = {description: \"实际出发地名称\"}];   string actual_end_name = 42 [(openapi.v3.property) = {description: \"实际目的地名称\"}];   string rule_name = 43 [(openapi.v3.property) = {description: \"用车规则\"}];   string use_car_srv = 44 [(openapi.v3.property) = {description: \"用车类型 枚举【快车用车、专车用车、专车包车、豪华车送机 等】\"}];   string budget_center_name = 45 [(openapi.v3.property) = {description: \"成本中心名称 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）名称，同时可以设置员工下单时是否必填此字段\"}];   string budget_id = 46 [(openapi.v3.property) = {description: \"成本中心ID 管理员在ES后台差旅管控模块维护的成本中心（部门或项目）ID，同时可以设置员工下单时是否必填此字段\"}];   string remark = 47 [(openapi.v3.property) = {description: \"备注\"}];   string order_additional_remark = 48 [(openapi.v3.property) = {description: \"补充说明\"}];   string approval_reason = 49 [(openapi.v3.property) = {description: \"审批备注\"}];   string approval_create_time = 50 [(openapi.v3.property) = {description: \"申请提交时间\"}];   string approval_logs = 51 [(openapi.v3.property) = {description: \"审批历史-审批时间\"}];   string out_approval_id = 52 [(openapi.v3.property) = {description: \"外部申请单id\"}];   string upgrade_type = 53 [(openapi.v3.property) = {description: \"升舱类型 枚举【非升舱单、升舱、专车拼车、降舱】\"}];   string is_fixed_price = 54 [(openapi.v3.property) = {description: \"是否一口价 枚举【是、否】\"}];   string is_self_drive = 55 [(openapi.v3.property) = {description: \"是否是私车同行 枚举【是、否】\"}];   string is_discount_order = 56 [(openapi.v3.property) = {description: \"是否有折扣 枚举【是、否】\"}];   string is_carpool = 57 [(openapi.v3.property) = {description: \"是否拼车 枚举【是、否】\"}];   string is_reassignment = 58 [(openapi.v3.property) = {description: \"是否改派 枚举【是、否】\"}];   string callback_info = 59 [(openapi.v3.property) = {description: \"用户自定义字段\"}];   string approval_extra_info = 60 [(openapi.v3.property) = {description: \"审批单自定义\"}];   string after_approval_result = 61 [(openapi.v3.property) = {description: \"行后审批结果\"}];   string after_approval_content = 62 [(openapi.v3.property) = {description: \"审批历史\"}];   int64 company_id = 63 [(openapi.v3.property) = {description: \"公司id\" type:\"integer\" format:\"int64\"}];   string company_name = 64 [(openapi.v3.property) = {description: \"公司名称\"}];   int64 legal_entity_id = 65 [(openapi.v3.property) = {description: \"预定人开票子公司id\" type:\"integer\" format:\"int64\"}];   string legal_entity_name = 66 [(openapi.v3.property) = {description: \"预定人开票子公司名称\"}];   double company_card_real_pay = 67 [(openapi.v3.property) = {description: \"企业钱包实付金额\"}];   string use_car_type_v2 = 68 [(openapi.v3.property) = {description: \"用车场景 枚举【加班、日常出勤、接送机 等】\"}];   string budget_item_name = 69 [(openapi.v3.property) = {description: \"科目费用\"}];   string use_type = 70 [(openapi.v3.property) = {description: \"是否代叫车 枚举【代客户叫车、员工自己用车】\"}];   double start_price = 71 [(openapi.v3.property) = {description: \"起步价\"}];   double normal_fee = 72 [(openapi.v3.property) = {description: \"里程费\"}];   double low_speed_fee = 73 [(openapi.v3.property) = {description: \"低速费\"}];   double empty_fee = 74 [(openapi.v3.property) = {description: \"远途费\"}];   double night_fee = 75 [(openapi.v3.property) = {description: \"夜间费\"}];   double responsible_cancel_fee = 76 [(openapi.v3.property) = {description: \"有责取消费\"}];   double limit_fee = 77 [(openapi.v3.property) = {description: \"基础费\"}];   double limit_pay = 78 [(openapi.v3.property) = {description: \"基础费补足金额\"}];   double normal_time_fee = 79 [(openapi.v3.property) = {description: \"时长费\"}];   double dynamic_price = 80 [(openapi.v3.property) = {description: \"动态调价\"}];   double tip_fee = 81 [(openapi.v3.property) = {description: \"小费\"}];   double bridge_fee = 82 [(openapi.v3.property) = {description: \"路桥费\"}];   double highway_fee = 83 [(openapi.v3.property) = {description: \"高速费\"}];   double park_fee = 84 [(openapi.v3.property) = {description: \"停车费\"}];   double wait_fee = 85 [(openapi.v3.property) = {description: \"等待费\"}];   double incentive_fee = 86 [(openapi.v3.property) = {description: \"调度费\"}];   double red_packet = 87 [(openapi.v3.property) = {description: \"春节加价费\"}];   double other_fee = 88 [(openapi.v3.property) = {description: \"其他费用\"}];   int64 member_id = 89 [(openapi.v3.property) = {description: \"预定人id\" type:\"integer\" format:\"int64\"}];   double refund_price = 90 [(openapi.v3.property) = {description: \"退款金额\"}];   string group_name = 91 [(openapi.v3.property) = {description: \"预定人部门名称（新）\"}];   string group_id = 92 [(openapi.v3.property) = {description: \"预定人部门id\"}];   double company_real_pre_tax_pay = 93 [(openapi.v3.property) = {description: \"不含税实付金额 =企业实付金额*（1-税率）\"}];   double company_real_tax_pay = 94 [(openapi.v3.property) = {description: \"税额 =企业实付金额*税率\"}];   string user_pay_time = 95 [(openapi.v3.property) = {description: \"用户首次确认支付时间\"}];   string user_pay_type = 96 [(openapi.v3.property) = {description: \"用户支付类型 枚举【用户主动支付、系统自动支付、系统自动支付后用户主动支付】\"}];   string settle_type = 97 [(openapi.v3.property) = {description: \"业务线名称 网约车\"}];   string sub_account_name = 98 [(openapi.v3.property) = {description: \"所属子账户\"}];   string belong_budget_center_name = 99 [(openapi.v3.property) = {description: \"恒为空 可忽略\"}];   string belong_out_budget_id = 100 [(openapi.v3.property) = {description: \"恒为空 可忽略\"}];   int32 is_carpool_success = 101 [(openapi.v3.property) = {description: \"是否拼成 1:是 0:否\"}];   double company_pay_able = 102 [(openapi.v3.property) = {description: \"企业应付金额\"}];   string service_type = 103 [(openapi.v3.property) = {description: \"包车套餐\"}];   string before_approval_result = 104 [(openapi.v3.property) = {description: \"行前审批\"}];   double time_slot_discount = 105 [(openapi.v3.property) = {description: \"时段单单惠\"}];   string extend_info = 106 [(openapi.v3.property) = {description: \"附加信息（开票主体信息） 对应用户invoice_info信息\"}];   string period = 107 [(openapi.v3.property) = {description: \"支付账期\"}];   double excluding_tax_pay_price = 108 [(openapi.v3.property) = {description: \"不含税实付金额 =企业实付金额*（1-税率） 同company_real_pre_tax_pay 冗余字段\"}];   double tax_pay_price = 109 [(openapi.v3.property) = {description: \"实付税额 =企业实付金额*税率 同company_real_tax_pay 冗余字段\"}];   double excluding_tax_refund_price = 110 [(openapi.v3.property) = {description: \"不含税实退金额 =企业实退金额*（1-税率）\"}];   double tax_refund_price = 111 [(openapi.v3.property) = {description: \"实退税额 =企业实退金额*税率\"}];   int32 is_sensitive = 112 [(openapi.v3.property) = {description: \"是否敏感订单 1:是 0:否\"}];   string sensitive_reason = 113 [(openapi.v3.property) = {description: \"敏感订单原因\"}];   double company_card_real_refund = 114 [(openapi.v3.property) = {description: \"企业出行卡实退金额\"}];   string cancel_time = 115 [(openapi.v3.property) = {description: \"取消时间\"}];   double cross_city_fee = 116 [(openapi.v3.property) = {description: \"跨城费\"}];   double remote_area_fee = 117 [(openapi.v3.property) = {description: \"偏远地区接驾费\"}];   string use_car_type_name = 118 [(openapi.v3.property) = {description: \"用车场景细分 枚举【日常用车、加班用车、办公地通勤 等】\"}];   string sub_use_car_srv = 119 [(openapi.v3.property) = {description: \"使用场景 枚举【市内用车、接送机、接送火车站、接送汽车站、接送渡口】\"}];   double energy_consume_fee = 120 [(openapi.v3.property) = {description: \"综合能耗费\"}];   string is_unusual = 121 [(openapi.v3.property) = {description: \"是否异常 枚举【是、否】\"}];   string unusual_type = 122 [(openapi.v3.property) = {description: \"异常类型\"}];   string unusual_content = 123 [(openapi.v3.property) = {description: \"异常说明\"}];   string position_name = 124 [(openapi.v3.property) = {description: \"职级\"}];   string institution_name = 125 [(openapi.v3.property) = {description: \"用车制度\"}];   string use_car_service = 126 [(openapi.v3.property) = {description: \"用车服务 枚举【接送服务、出差市内用车】\"}];   string ex_info_01 = 127 [(openapi.v3.property) = {description: \"用户自定义拓展字段1\"}];   string ex_info_02 = 128 [(openapi.v3.property) = {description: \"用户自定义拓展字段2\"}];   string ex_info_03 = 129 [(openapi.v3.property) = {description: \"用户自定义拓展字段3\"}];   string ex_info_04 = 130 [(openapi.v3.property) = {description: \"用户自定义拓展字段4\"}];   string ex_info_05 = 131 [(openapi.v3.property) = {description: \"用户自定义拓展字段5\"}];   string ex_info_06 = 132 [(openapi.v3.property) = {description: \"用户自定义拓展字段6\"}];   string ex_info_07 = 133 [(openapi.v3.property) = {description: \"用户自定义拓展字段7\"}];   string ex_info_08 = 134 [(openapi.v3.property) = {description: \"用户自定义拓展字段8\"}];   double sps_pick_up_service_fee = 135 [(openapi.v3.property) = {description: \"司机举牌接机服务费\"}];   string deduction_type_name = 136 [(openapi.v3.property) = {description: \"交易类型 枚举：订单企业支付扣款、订单企业支付后退款、企业支付转个人支付退款、行后审批驳回后退款、个人支付转企业支付扣款、线下企业支付扣款、企业支付转个人支付充值、行后审批驳回后充值\"}];   string pay_channel = 137 [(openapi.v3.property) = {description: \"结算方式 枚举：银行卡、德付通、预存、返现、授信、试用金、电子账户\"}];   string resident_city_names = 138 [(openapi.v3.property) = {description: \"常驻城市\"}];   double original_price = 139 [(openapi.v3.property) = {description: \"原价\"}];   double one_time_offer_subsidy = 140 [(openapi.v3.property) = {description: \"尊享折扣\"}];   double subsidy_amount = 141 [(openapi.v3.property) = {description: \"补贴金额\"}];   string voucher_deduction_type_name = 142 [(openapi.v3.property) = {description: \"券抵扣类型\"}]; //  string passenger_member_number = 143 [(openapi.v3.property) = {description: \"乘车人员工编号\"}]; // 文档中出现的重复字段   int64 passenger_member_id = 144 [(openapi.v3.property) = {description: \"乘车人员工id\" type:\"integer\" format:\"int64\"}];   string project_ext_info = 145 [(openapi.v3.property) = {description: \"项目自定义\"}];   string out_legal_entity_id = 146 [(openapi.v3.property) = {description: \"外部公司主体编号\"}];   string car_travel_Info = 147 [(openapi.v3.property) = {description: \"途经点地址 如：丰台区-大红门街道1号；多个途经点用｜连接，丰台区-大红门街道1号｜朝阳区-国贸5号楼\"}];   string supplier_type = 148 [(openapi.v3.property) = {description: \"运力来源 枚举值：滴滴自营、滴滴旗下品牌、第三方服务\"}];   string supplier_name = 149 [(openapi.v3.property) = {description: \"三方名称\"}];   string supplier_car_name = 150 [(openapi.v3.property) = {description: \"三方车型\"}];   int64 institution_id = 151 [(openapi.v3.property) = {description: \"制度ID\" type:\"integer\" format:\"int64\"}];   string excluding_service_fee = 152 [(openapi.v3.property) = {description: \"企业实付（不包含平台使用费）\"}];   double estimate_price = 153 [(openapi.v3.property) = {description: \"预估金额\"}];   string sub_account_company_name = 154 [(openapi.v3.property) = {description: \"子账户公司名称\"}];   string start_county_id = 155 [(openapi.v3.property) = {description: \"开始区县id\"}];   string end_county_id = 156 [(openapi.v3.property) = {description: \"结束区县id\"}];   string start_county_name = 157 [(openapi.v3.property) = {description: \"开始区县名称\"}];   string end_county_name = 158 [(openapi.v3.property) = {description: \"结束区县名称\"}];   string flag_ext = 159 [(openapi.v3.property) = {description: \"是否追加车型\"}];   string c_weight = 160 [(openapi.v3.property) = {description: \"碳排放\"}];   string co2_emissions = 161 [(openapi.v3.property) = {description: \"碳减排\"}];   string after_approve_member_one = 162 [(openapi.v3.property) = {description: \"行后审批人工号1\"}];   string after_approve_member_two = 163 [(openapi.v3.property) = {description: \"行后审批人工号2\"}];   string after_approve_member_three = 164 [(openapi.v3.property) = {description: \"行后审批人工号3\"}];   string call_phone_country_code = 165 [(openapi.v3.property) = {description: \"预订人手机号区号\"}];   string passenger_phone_country_code = 166 [(openapi.v3.property) = {description: \"乘车人手机号区号\"}];   string enterprise_coupon_batch = 167 [(openapi.v3.property) = {description: \"企业券批次\"}];   string enterprise_coupon_name = 168 [(openapi.v3.property) = {description: \"企业券名称\"}];   string upgrade_base_require_level = 169 [(openapi.v3.property) = {description: \"企业支付车型\"}];   string last_approves = 170 [(openapi.v3.property) = {description: \"行后审批最后审批人\"}];   string last_approval_time = 171 [(openapi.v3.property) = {description: \"行后审批最后审批时间\"}];   string ex_info_01_code = 172 [(openapi.v3.property) = {description: \"用户自定义拓展字段1编码\"}];   string ex_info_02_code = 173 [(openapi.v3.property) = {description: \"用户自定义拓展字段2编码\"}];   string ex_info_03_code = 174 [(openapi.v3.property) = {description: \"用户自定义拓展字段3编码\"}];   string ex_info_04_code = 175 [(openapi.v3.property) = {description: \"用户自定义拓展字段4编码\"}];   string ex_info_05_code = 176 [(openapi.v3.property) = {description: \"用户自定义拓展字段5编码\"}];   string ex_info_06_code = 177 [(openapi.v3.property) = {description: \"用户自定义拓展字段6编码\"}];   string ex_info_07_code = 178 [(openapi.v3.property) = {description: \"用户自定义拓展字段7编码\"}];   string ex_info_08_code = 179 [(openapi.v3.property) = {description: \"用户自定义拓展字段8编码\"}];   string originalEndName = 180 [(openapi.v3.property) = {description: \"原始目的地字段\"}];   string originalEndAddress = 181 [(openapi.v3.property) = {description: \"原始目的地地址\"}];   string actualEndNameParsed = 182 [(openapi.v3.property) = {description: \"实际目的地未脱敏（中金用）\"}];   string pay_type = 183 [(openapi.v3.property) = {description: \"支付方式 枚举【\\\"企业支付\\\"、\\\"个人支付\\\"、\\\"混合支付\\\"、\\\"企业钱包支付\\\"】 当前订单支付金额的方式，如：企业支付、混合支付（指企业支付和个人支付两种支付方式同时支付一笔订单）\"}]; }  已出账单 ~ 网约车返回参数
type BillListItemOfWangYC struct {
	ActualEndName             *string  `json:"actual_end_name,omitempty"`
	ActualFromIsWorkPlace     *string  `json:"actual_from_is_work_place,omitempty"`
	ActualStartName           *string  `json:"actual_start_name,omitempty"`
	ActualToIsWorkPlace       *string  `json:"actual_to_is_work_place,omitempty"`
	AfterApprovalContent      *string  `json:"after_approval_content,omitempty"`
	AfterApprovalResult       *string  `json:"after_approval_result,omitempty"`
	AfterApproveMemberOne     *string  `json:"after_approve_member_one,omitempty"`
	AfterApproveMemberThree   *string  `json:"after_approve_member_three,omitempty"`
	AfterApproveMemberTwo     *string  `json:"after_approve_member_two,omitempty"`
	ApprovalCreateTime        *string  `json:"approval_create_time,omitempty"`
	ApprovalExtraInfo         *string  `json:"approval_extra_info,omitempty"`
	ApprovalLogs              *string  `json:"approval_logs,omitempty"`
	ApprovalReason            *string  `json:"approval_reason,omitempty"`
	BeforeApprovalResult      *string  `json:"before_approval_result,omitempty"`
	BeginChargeTime           *string  `json:"begin_charge_time,omitempty"`
	BelongBudgetCenterName    *string  `json:"belong_budget_center_name,omitempty"`
	BelongOutBudgetId         *string  `json:"belong_out_budget_id,omitempty"`
	BookingDepCode            *string  `json:"booking_dep_code,omitempty"`
	BookingServiceFee         *float32 `json:"booking_service_fee,omitempty"`
	BranchName                *string  `json:"branch_name,omitempty"`
	BridgeFee                 *float32 `json:"bridge_fee,omitempty"`
	BudgetCenterName          *string  `json:"budget_center_name,omitempty"`
	BudgetId                  *string  `json:"budget_id,omitempty"`
	BudgetItemName            *string  `json:"budget_item_name,omitempty"`
	CWeight                   *string  `json:"c_weight,omitempty"`
	CallPhone                 *string  `json:"call_phone,omitempty"`
	CallPhoneCountryCode      *string  `json:"call_phone_country_code,omitempty"`
	CallbackInfo              *string  `json:"callback_info,omitempty"`
	CancelTime                *string  `json:"cancel_time,omitempty"`
	CarArriveProvince         *string  `json:"car_arrive_province,omitempty"`
	CarDepartProvince         *string  `json:"car_depart_province,omitempty"`
	CarTravelInfo             *string  `json:"car_travel_info,omitempty"`
	ChargeTime                *string  `json:"charge_time,omitempty"`
	City                      *string  `json:"city,omitempty"`
	CleanCarFee               *float32 `json:"clean_car_fee,omitempty"`
	Co2Emissions              *string  `json:"co2_emissions,omitempty"`
	CompanyCardRealPay        *float32 `json:"company_card_real_pay,omitempty"`
	CompanyCardRealRefund     *float32 `json:"company_card_real_refund,omitempty"`
	CompanyId                 *int64   `json:"company_id,omitempty"`
	CompanyName               *string  `json:"company_name,omitempty"`
	CompanyPayAble            *float32 `json:"company_pay_able,omitempty"`
	CompanyRealPay            *float32 `json:"company_real_pay,omitempty"`
	CompanyRealPreTaxPay      *float32 `json:"company_real_pre_tax_pay,omitempty"`
	CompanyRealRefund         *float32 `json:"company_real_refund,omitempty"`
	CompanyRealTaxPay         *float32 `json:"company_real_tax_pay,omitempty"`
	Cost                      *float32 `json:"cost,omitempty"`
	CreateTime                *string  `json:"create_time,omitempty"`
	CrossCityFee              *float32 `json:"cross_city_fee,omitempty"`
	DeductionTypeName         *string  `json:"deduction_type_name,omitempty"`
	Department                *string  `json:"department,omitempty"`
	DepartmentName            *string  `json:"department_name,omitempty"`
	DepartureDime             *string  `json:"departure_dime,omitempty"`
	DestinationCity           *string  `json:"destination_city,omitempty"`
	DiscountAfterPrice        *float32 `json:"discount_after_price,omitempty"`
	DynamicPrice              *float32 `json:"dynamic_price,omitempty"`
	EmployeeNumber            *string  `json:"employee_number,omitempty"`
	EmptyFee                  *float32 `json:"empty_fee,omitempty"`
	EndAddress                *string  `json:"end_address,omitempty"`
	EndCountyId               *string  `json:"end_county_id,omitempty"`
	EndCountyName             *string  `json:"end_county_name,omitempty"`
	EndName                   *string  `json:"end_name,omitempty"`
	EnergyConsumeFee          *float32 `json:"energy_consume_fee,omitempty"`
	EnterpriseCouponBatch     *string  `json:"enterprise_coupon_batch,omitempty"`
	EnterpriseCouponName      *string  `json:"enterprise_coupon_name,omitempty"`
	EnterpriseInstantDiscount *float32 `json:"enterprise_instant_discount,omitempty"`
	EstimatePrice             *float32 `json:"estimate_price,omitempty"`
	EstimatedDistance         *string  `json:"estimated_distance,omitempty"`
	ExInfo01                  *string  `json:"ex_info_01,omitempty"`
	ExInfo01Code              *string  `json:"ex_info_01_code,omitempty"`
	ExInfo02                  *string  `json:"ex_info_02,omitempty"`
	ExInfo02Code              *string  `json:"ex_info_02_code,omitempty"`
	ExInfo03                  *string  `json:"ex_info_03,omitempty"`
	ExInfo03Code              *string  `json:"ex_info_03_code,omitempty"`
	ExInfo04                  *string  `json:"ex_info_04,omitempty"`
	ExInfo04Code              *string  `json:"ex_info_04_code,omitempty"`
	ExInfo05                  *string  `json:"ex_info_05,omitempty"`
	ExInfo05Code              *string  `json:"ex_info_05_code,omitempty"`
	ExInfo06                  *string  `json:"ex_info_06,omitempty"`
	ExInfo06Code              *string  `json:"ex_info_06_code,omitempty"`
	ExInfo07                  *string  `json:"ex_info_07,omitempty"`
	ExInfo07Code              *string  `json:"ex_info_07_code,omitempty"`
	ExInfo08                  *string  `json:"ex_info_08,omitempty"`
	ExInfo08Code              *string  `json:"ex_info_08_code,omitempty"`
	ExcludingServiceFee       *float32 `json:"excluding_service_fee,omitempty"`
	ExcludingTaxPayPrice      *float32 `json:"excluding_tax_pay_price,omitempty"`
	ExcludingTaxRefundPrice   *float32 `json:"excluding_tax_refund_price,omitempty"`
	ExtendInfo                *string  `json:"extend_info,omitempty"`
	FinishTime                *string  `json:"finish_time,omitempty"`
	FlagExt                   *string  `json:"flag_ext,omitempty"`
	FromIsWorkPlace           *string  `json:"from_is_work_place,omitempty"`
	Group1Code                *string  `json:"group_1_code,omitempty"`
	Group1Name                *string  `json:"group_1_name,omitempty"`
	Group2Code                *string  `json:"group_2_code,omitempty"`
	Group2Name                *string  `json:"group_2_name,omitempty"`
	Group3Code                *string  `json:"group_3_code,omitempty"`
	Group3Name                *string  `json:"group_3_name,omitempty"`
	Group4Code                *string  `json:"group_4_code,omitempty"`
	Group4Name                *string  `json:"group_4_name,omitempty"`
	Group5Code                *string  `json:"group_5_code,omitempty"`
	Group5Name                *string  `json:"group_5_name,omitempty"`
	Group6Code                *string  `json:"group_6_code,omitempty"`
	Group6Name                *string  `json:"group_6_name,omitempty"`
	Group7Code                *string  `json:"group_7_code,omitempty"`
	Group7Name                *string  `json:"group_7_name,omitempty"`
	Group8Code                *string  `json:"group_8_code,omitempty"`
	Group8Name                *string  `json:"group_8_name,omitempty"`
	Group9Code                *string  `json:"group_9_code,omitempty"`
	Group9Name                *string  `json:"group_9_name,omitempty"`
	GroupId                   *string  `json:"group_id,omitempty"`
	GroupName                 *string  `json:"group_name,omitempty"`
	HighwayFee                *float32 `json:"highway_fee,omitempty"`
	IncentiveFee              *float32 `json:"incentive_fee,omitempty"`
	InstitutionId             *float32 `json:"institution_id,omitempty"`
	InstitutionName           *string  `json:"institution_name,omitempty"`
	InvoiceEntityInfo         *string  `json:"invoice_entity_info,omitempty"`
	IsCarpool                 *string  `json:"is_carpool,omitempty"`
	IsCarpoolSuccess          *float32 `json:"is_carpool_success,omitempty"`
	IsCurrentTerm             *string  `json:"is_current_term,omitempty"`
	IsDiscountOrder           *string  `json:"is_discount_order,omitempty"`
	IsFixedPrice              *string  `json:"is_fixed_price,omitempty"`
	IsReassignment            *string  `json:"is_reassignment,omitempty"`
	IsSelfDrive               *string  `json:"is_self_drive,omitempty"`
	IsSensitive               *float32 `json:"is_sensitive,omitempty"`
	IsUnusual                 *string  `json:"is_unusual,omitempty"`
	LastApprovalTime          *string  `json:"last_approval_time,omitempty"`
	LastApproves              *string  `json:"last_approves,omitempty"`
	LegalEntityId             *int64   `json:"legal_entity_id,omitempty"`
	LegalEntityName           *string  `json:"legal_entity_name,omitempty"`
	LimitFee                  *float32 `json:"limit_fee,omitempty"`
	LimitPay                  *float32 `json:"limit_pay,omitempty"`
	LowSpeedFee               *float32 `json:"low_speed_fee,omitempty"`
	MeetTime                  *string  `json:"meetTime,omitempty"`
	MemberId                  *int64   `json:"member_id,omitempty"`
	MemberMail                *string  `json:"member_mail,omitempty"`
	MemberName                *string  `json:"member_name,omitempty"`
	NightFee                  *float32 `json:"night_fee,omitempty"`
	NormalDistance            *float32 `json:"normal_distance,omitempty"`
	NormalFee                 *float32 `json:"normal_fee,omitempty"`
	NormalTimeFee             *float32 `json:"normal_time_fee,omitempty"`
	OneTimeOfferSubsidy       *float32 `json:"one_time_offer_subsidy,omitempty"`
	OrderAdditionalRemark     *string  `json:"order_additional_remark,omitempty"`
	OrderId                   *int64   `json:"order_id,omitempty"`
	OrderSource               *string  `json:"order_source,omitempty"`
	OriginalEndAddress        *string  `json:"original_end_address,omitempty"`
	OriginalEndName           *string  `json:"original_end_name,omitempty"`
	OriginalPrice             *float32 `json:"original_price,omitempty"`
	OtherFee                  *float32 `json:"other_fee,omitempty"`
	OutApprovalId             *string  `json:"out_approval_id,omitempty"`
	OutBudgetId               *string  `json:"out_budget_id,omitempty"`
	OutLegalEntityId          *string  `json:"out_legal_entity_id,omitempty"`
	OutOrderId                *int64   `json:"out_order_id,omitempty"`
	ParkFee                   *float32 `json:"park_fee,omitempty"`
	PassengerMemberId         *int64   `json:"passenger_member_id,omitempty"`
	PassengerMemberNumber     *string  `json:"passenger_member_number,omitempty"`
	PassengerName             *string  `json:"passenger_name,omitempty"`
	PassengerPhone            *string  `json:"passenger_phone,omitempty"`
	PassengerPhoneCountryCode *string  `json:"passenger_phone_country_code,omitempty"`
	PayChannel                *string  `json:"pay_channel,omitempty"`
	PayTime                   *string  `json:"pay_time,omitempty"`
	PayType                   *string  `json:"pay_type,omitempty"`
	Period                    *string  `json:"period,omitempty"`
	PersonalInstantDiscount   *float32 `json:"personal_instant_discount,omitempty"`
	PersonalRealPay           *float32 `json:"personal_real_pay,omitempty"`
	PositionName              *string  `json:"position_name,omitempty"`
	ProjectExtInfo            *string  `json:"project_ext_info,omitempty"`
	RealVoucherPay            *float32 `json:"real_voucher_pay,omitempty"`
	RedPacket                 *float32 `json:"red_packet,omitempty"`
	RefundPrice               *float32 `json:"refund_price,omitempty"`
	RefundTime                *string  `json:"refund_time,omitempty"`
	Remark                    *string  `json:"remark,omitempty"`
	RemoteAreaFee             *float32 `json:"remote_area_fee,omitempty"`
	RequireLevel              *string  `json:"require_level,omitempty"`
	RequisitionId             *string  `json:"requisition_id,omitempty"`
	ResidentCityNames         *string  `json:"resident_city_names,omitempty"`
	ResponsibleCancelFee      *float32 `json:"responsible_cancel_fee,omitempty"`
	Rule                      *string  `json:"rule,omitempty"`
	RuleName                  *string  `json:"rule_name,omitempty"`
	SensitiveReason           *string  `json:"sensitive_reason,omitempty"`
	ServiceType               *string  `json:"service_type,omitempty"`
	SettleType                *string  `json:"settle_type,omitempty"`
	SpsPickUpServiceFee       *float32 `json:"sps_pick_up_service_fee,omitempty"`
	StartAddress              *string  `json:"start_address,omitempty"`
	StartCountyId             *string  `json:"start_county_id,omitempty"`
	StartCountyName           *string  `json:"start_county_name,omitempty"`
	StartName                 *string  `json:"start_name,omitempty"`
	StartPrice                *float32 `json:"start_price,omitempty"`
	Status                    *string  `json:"status,omitempty"`
	StriveTime                *string  `json:"strive_time,omitempty"`
	SubAccountCompanyName     *string  `json:"sub_account_company_name,omitempty"`
	SubAccountName            *string  `json:"sub_account_name,omitempty"`
	SubUseCarSrv              *string  `json:"sub_use_car_srv,omitempty"`
	SubsidyAmount             *float32 `json:"subsidy_amount,omitempty"`
	SupplierCarName           *string  `json:"supplier_car_name,omitempty"`
	SupplierName              *string  `json:"supplier_name,omitempty"`
	SupplierType              *string  `json:"supplier_type,omitempty"`
	TaxPayPrice               *float32 `json:"tax_pay_price,omitempty"`
	TaxRefundPrice            *float32 `json:"tax_refund_price,omitempty"`
	TimeSlotDiscount          *float32 `json:"time_slot_discount,omitempty"`
	TipFee                    *float32 `json:"tip_fee,omitempty"`
	ToIsWorkPlace             *string  `json:"to_is_work_place,omitempty"`
	TotalPrice                *float32 `json:"total_price,omitempty"`
	Type                      *string  `json:"type,omitempty"`
	UnusualContent            *string  `json:"unusual_content,omitempty"`
	UnusualType               *string  `json:"unusual_type,omitempty"`
	UpgradeBaseRequireLevel   *string  `json:"upgrade_base_require_level,omitempty"`
	UpgradeCompanyPayQuota    *float32 `json:"upgrade_company_pay_quota,omitempty"`
	UpgradeType               *string  `json:"upgrade_type,omitempty"`
	UseCarService             *string  `json:"use_car_service,omitempty"`
	UseCarSrv                 *string  `json:"use_car_srv,omitempty"`
	UseCarTypeName            *string  `json:"use_car_type_name,omitempty"`
	UseCarTypeV2              *string  `json:"use_car_type_v2,omitempty"`
	UseType                   *string  `json:"use_type,omitempty"`
	UserPayTime               *string  `json:"user_pay_time,omitempty"`
	UserPayType               *string  `json:"user_pay_type,omitempty"`
	VoucherDeductionTypeName  *string  `json:"voucher_deduction_type_name,omitempty"`
	WaitFee                   *float32 `json:"wait_fee,omitempty"`
}

type BillListItemOfWangYCBuilder struct {
	actualEndName                string
	actualEndNameSet             bool
	actualFromIsWorkPlace        string
	actualFromIsWorkPlaceSet     bool
	actualStartName              string
	actualStartNameSet           bool
	actualToIsWorkPlace          string
	actualToIsWorkPlaceSet       bool
	afterApprovalContent         string
	afterApprovalContentSet      bool
	afterApprovalResult          string
	afterApprovalResultSet       bool
	afterApproveMemberOne        string
	afterApproveMemberOneSet     bool
	afterApproveMemberThree      string
	afterApproveMemberThreeSet   bool
	afterApproveMemberTwo        string
	afterApproveMemberTwoSet     bool
	approvalCreateTime           string
	approvalCreateTimeSet        bool
	approvalExtraInfo            string
	approvalExtraInfoSet         bool
	approvalLogs                 string
	approvalLogsSet              bool
	approvalReason               string
	approvalReasonSet            bool
	beforeApprovalResult         string
	beforeApprovalResultSet      bool
	beginChargeTime              string
	beginChargeTimeSet           bool
	belongBudgetCenterName       string
	belongBudgetCenterNameSet    bool
	belongOutBudgetId            string
	belongOutBudgetIdSet         bool
	bookingDepCode               string
	bookingDepCodeSet            bool
	bookingServiceFee            float32
	bookingServiceFeeSet         bool
	branchName                   string
	branchNameSet                bool
	bridgeFee                    float32
	bridgeFeeSet                 bool
	budgetCenterName             string
	budgetCenterNameSet          bool
	budgetId                     string
	budgetIdSet                  bool
	budgetItemName               string
	budgetItemNameSet            bool
	cWeight                      string
	cWeightSet                   bool
	callPhone                    string
	callPhoneSet                 bool
	callPhoneCountryCode         string
	callPhoneCountryCodeSet      bool
	callbackInfo                 string
	callbackInfoSet              bool
	cancelTime                   string
	cancelTimeSet                bool
	carArriveProvince            string
	carArriveProvinceSet         bool
	carDepartProvince            string
	carDepartProvinceSet         bool
	carTravelInfo                string
	carTravelInfoSet             bool
	chargeTime                   string
	chargeTimeSet                bool
	city                         string
	citySet                      bool
	cleanCarFee                  float32
	cleanCarFeeSet               bool
	co2Emissions                 string
	co2EmissionsSet              bool
	companyCardRealPay           float32
	companyCardRealPaySet        bool
	companyCardRealRefund        float32
	companyCardRealRefundSet     bool
	companyId                    int64
	companyIdSet                 bool
	companyName                  string
	companyNameSet               bool
	companyPayAble               float32
	companyPayAbleSet            bool
	companyRealPay               float32
	companyRealPaySet            bool
	companyRealPreTaxPay         float32
	companyRealPreTaxPaySet      bool
	companyRealRefund            float32
	companyRealRefundSet         bool
	companyRealTaxPay            float32
	companyRealTaxPaySet         bool
	cost                         float32
	costSet                      bool
	createTime                   string
	createTimeSet                bool
	crossCityFee                 float32
	crossCityFeeSet              bool
	deductionTypeName            string
	deductionTypeNameSet         bool
	department                   string
	departmentSet                bool
	departmentName               string
	departmentNameSet            bool
	departureDime                string
	departureDimeSet             bool
	destinationCity              string
	destinationCitySet           bool
	discountAfterPrice           float32
	discountAfterPriceSet        bool
	dynamicPrice                 float32
	dynamicPriceSet              bool
	employeeNumber               string
	employeeNumberSet            bool
	emptyFee                     float32
	emptyFeeSet                  bool
	endAddress                   string
	endAddressSet                bool
	endCountyId                  string
	endCountyIdSet               bool
	endCountyName                string
	endCountyNameSet             bool
	endName                      string
	endNameSet                   bool
	energyConsumeFee             float32
	energyConsumeFeeSet          bool
	enterpriseCouponBatch        string
	enterpriseCouponBatchSet     bool
	enterpriseCouponName         string
	enterpriseCouponNameSet      bool
	enterpriseInstantDiscount    float32
	enterpriseInstantDiscountSet bool
	estimatePrice                float32
	estimatePriceSet             bool
	estimatedDistance            string
	estimatedDistanceSet         bool
	exInfo01                     string
	exInfo01Set                  bool
	exInfo01Code                 string
	exInfo01CodeSet              bool
	exInfo02                     string
	exInfo02Set                  bool
	exInfo02Code                 string
	exInfo02CodeSet              bool
	exInfo03                     string
	exInfo03Set                  bool
	exInfo03Code                 string
	exInfo03CodeSet              bool
	exInfo04                     string
	exInfo04Set                  bool
	exInfo04Code                 string
	exInfo04CodeSet              bool
	exInfo05                     string
	exInfo05Set                  bool
	exInfo05Code                 string
	exInfo05CodeSet              bool
	exInfo06                     string
	exInfo06Set                  bool
	exInfo06Code                 string
	exInfo06CodeSet              bool
	exInfo07                     string
	exInfo07Set                  bool
	exInfo07Code                 string
	exInfo07CodeSet              bool
	exInfo08                     string
	exInfo08Set                  bool
	exInfo08Code                 string
	exInfo08CodeSet              bool
	excludingServiceFee          float32
	excludingServiceFeeSet       bool
	excludingTaxPayPrice         float32
	excludingTaxPayPriceSet      bool
	excludingTaxRefundPrice      float32
	excludingTaxRefundPriceSet   bool
	extendInfo                   string
	extendInfoSet                bool
	finishTime                   string
	finishTimeSet                bool
	flagExt                      string
	flagExtSet                   bool
	fromIsWorkPlace              string
	fromIsWorkPlaceSet           bool
	group1Code                   string
	group1CodeSet                bool
	group1Name                   string
	group1NameSet                bool
	group2Code                   string
	group2CodeSet                bool
	group2Name                   string
	group2NameSet                bool
	group3Code                   string
	group3CodeSet                bool
	group3Name                   string
	group3NameSet                bool
	group4Code                   string
	group4CodeSet                bool
	group4Name                   string
	group4NameSet                bool
	group5Code                   string
	group5CodeSet                bool
	group5Name                   string
	group5NameSet                bool
	group6Code                   string
	group6CodeSet                bool
	group6Name                   string
	group6NameSet                bool
	group7Code                   string
	group7CodeSet                bool
	group7Name                   string
	group7NameSet                bool
	group8Code                   string
	group8CodeSet                bool
	group8Name                   string
	group8NameSet                bool
	group9Code                   string
	group9CodeSet                bool
	group9Name                   string
	group9NameSet                bool
	groupId                      string
	groupIdSet                   bool
	groupName                    string
	groupNameSet                 bool
	highwayFee                   float32
	highwayFeeSet                bool
	incentiveFee                 float32
	incentiveFeeSet              bool
	institutionId                float32
	institutionIdSet             bool
	institutionName              string
	institutionNameSet           bool
	invoiceEntityInfo            string
	invoiceEntityInfoSet         bool
	isCarpool                    string
	isCarpoolSet                 bool
	isCarpoolSuccess             float32
	isCarpoolSuccessSet          bool
	isCurrentTerm                string
	isCurrentTermSet             bool
	isDiscountOrder              string
	isDiscountOrderSet           bool
	isFixedPrice                 string
	isFixedPriceSet              bool
	isReassignment               string
	isReassignmentSet            bool
	isSelfDrive                  string
	isSelfDriveSet               bool
	isSensitive                  float32
	isSensitiveSet               bool
	isUnusual                    string
	isUnusualSet                 bool
	lastApprovalTime             string
	lastApprovalTimeSet          bool
	lastApproves                 string
	lastApprovesSet              bool
	legalEntityId                int64
	legalEntityIdSet             bool
	legalEntityName              string
	legalEntityNameSet           bool
	limitFee                     float32
	limitFeeSet                  bool
	limitPay                     float32
	limitPaySet                  bool
	lowSpeedFee                  float32
	lowSpeedFeeSet               bool
	meetTime                     string
	meetTimeSet                  bool
	memberId                     int64
	memberIdSet                  bool
	memberMail                   string
	memberMailSet                bool
	memberName                   string
	memberNameSet                bool
	nightFee                     float32
	nightFeeSet                  bool
	normalDistance               float32
	normalDistanceSet            bool
	normalFee                    float32
	normalFeeSet                 bool
	normalTimeFee                float32
	normalTimeFeeSet             bool
	oneTimeOfferSubsidy          float32
	oneTimeOfferSubsidySet       bool
	orderAdditionalRemark        string
	orderAdditionalRemarkSet     bool
	orderId                      int64
	orderIdSet                   bool
	orderSource                  string
	orderSourceSet               bool
	originalEndAddress           string
	originalEndAddressSet        bool
	originalEndName              string
	originalEndNameSet           bool
	originalPrice                float32
	originalPriceSet             bool
	otherFee                     float32
	otherFeeSet                  bool
	outApprovalId                string
	outApprovalIdSet             bool
	outBudgetId                  string
	outBudgetIdSet               bool
	outLegalEntityId             string
	outLegalEntityIdSet          bool
	outOrderId                   int64
	outOrderIdSet                bool
	parkFee                      float32
	parkFeeSet                   bool
	passengerMemberId            int64
	passengerMemberIdSet         bool
	passengerMemberNumber        string
	passengerMemberNumberSet     bool
	passengerName                string
	passengerNameSet             bool
	passengerPhone               string
	passengerPhoneSet            bool
	passengerPhoneCountryCode    string
	passengerPhoneCountryCodeSet bool
	payChannel                   string
	payChannelSet                bool
	payTime                      string
	payTimeSet                   bool
	payType                      string
	payTypeSet                   bool
	period                       string
	periodSet                    bool
	personalInstantDiscount      float32
	personalInstantDiscountSet   bool
	personalRealPay              float32
	personalRealPaySet           bool
	positionName                 string
	positionNameSet              bool
	projectExtInfo               string
	projectExtInfoSet            bool
	realVoucherPay               float32
	realVoucherPaySet            bool
	redPacket                    float32
	redPacketSet                 bool
	refundPrice                  float32
	refundPriceSet               bool
	refundTime                   string
	refundTimeSet                bool
	remark                       string
	remarkSet                    bool
	remoteAreaFee                float32
	remoteAreaFeeSet             bool
	requireLevel                 string
	requireLevelSet              bool
	requisitionId                string
	requisitionIdSet             bool
	residentCityNames            string
	residentCityNamesSet         bool
	responsibleCancelFee         float32
	responsibleCancelFeeSet      bool
	rule                         string
	ruleSet                      bool
	ruleName                     string
	ruleNameSet                  bool
	sensitiveReason              string
	sensitiveReasonSet           bool
	serviceType                  string
	serviceTypeSet               bool
	settleType                   string
	settleTypeSet                bool
	spsPickUpServiceFee          float32
	spsPickUpServiceFeeSet       bool
	startAddress                 string
	startAddressSet              bool
	startCountyId                string
	startCountyIdSet             bool
	startCountyName              string
	startCountyNameSet           bool
	startName                    string
	startNameSet                 bool
	startPrice                   float32
	startPriceSet                bool
	status                       string
	statusSet                    bool
	striveTime                   string
	striveTimeSet                bool
	subAccountCompanyName        string
	subAccountCompanyNameSet     bool
	subAccountName               string
	subAccountNameSet            bool
	subUseCarSrv                 string
	subUseCarSrvSet              bool
	subsidyAmount                float32
	subsidyAmountSet             bool
	supplierCarName              string
	supplierCarNameSet           bool
	supplierName                 string
	supplierNameSet              bool
	supplierType                 string
	supplierTypeSet              bool
	taxPayPrice                  float32
	taxPayPriceSet               bool
	taxRefundPrice               float32
	taxRefundPriceSet            bool
	timeSlotDiscount             float32
	timeSlotDiscountSet          bool
	tipFee                       float32
	tipFeeSet                    bool
	toIsWorkPlace                string
	toIsWorkPlaceSet             bool
	totalPrice                   float32
	totalPriceSet                bool
	type_                        string
	type_Set                     bool
	unusualContent               string
	unusualContentSet            bool
	unusualType                  string
	unusualTypeSet               bool
	upgradeBaseRequireLevel      string
	upgradeBaseRequireLevelSet   bool
	upgradeCompanyPayQuota       float32
	upgradeCompanyPayQuotaSet    bool
	upgradeType                  string
	upgradeTypeSet               bool
	useCarService                string
	useCarServiceSet             bool
	useCarSrv                    string
	useCarSrvSet                 bool
	useCarTypeName               string
	useCarTypeNameSet            bool
	useCarTypeV2                 string
	useCarTypeV2Set              bool
	useType                      string
	useTypeSet                   bool
	userPayTime                  string
	userPayTimeSet               bool
	userPayType                  string
	userPayTypeSet               bool
	voucherDeductionTypeName     string
	voucherDeductionTypeNameSet  bool
	waitFee                      float32
	waitFeeSet                   bool
}

func NewBillListItemOfWangYCBuilder() *BillListItemOfWangYCBuilder {
	return &BillListItemOfWangYCBuilder{}
}
func (builder *BillListItemOfWangYCBuilder) ActualEndName(actualEndName string) *BillListItemOfWangYCBuilder {
	builder.actualEndName = actualEndName
	builder.actualEndNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ActualFromIsWorkPlace(actualFromIsWorkPlace string) *BillListItemOfWangYCBuilder {
	builder.actualFromIsWorkPlace = actualFromIsWorkPlace
	builder.actualFromIsWorkPlaceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ActualStartName(actualStartName string) *BillListItemOfWangYCBuilder {
	builder.actualStartName = actualStartName
	builder.actualStartNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ActualToIsWorkPlace(actualToIsWorkPlace string) *BillListItemOfWangYCBuilder {
	builder.actualToIsWorkPlace = actualToIsWorkPlace
	builder.actualToIsWorkPlaceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) AfterApprovalContent(afterApprovalContent string) *BillListItemOfWangYCBuilder {
	builder.afterApprovalContent = afterApprovalContent
	builder.afterApprovalContentSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) AfterApprovalResult(afterApprovalResult string) *BillListItemOfWangYCBuilder {
	builder.afterApprovalResult = afterApprovalResult
	builder.afterApprovalResultSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) AfterApproveMemberOne(afterApproveMemberOne string) *BillListItemOfWangYCBuilder {
	builder.afterApproveMemberOne = afterApproveMemberOne
	builder.afterApproveMemberOneSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) AfterApproveMemberThree(afterApproveMemberThree string) *BillListItemOfWangYCBuilder {
	builder.afterApproveMemberThree = afterApproveMemberThree
	builder.afterApproveMemberThreeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) AfterApproveMemberTwo(afterApproveMemberTwo string) *BillListItemOfWangYCBuilder {
	builder.afterApproveMemberTwo = afterApproveMemberTwo
	builder.afterApproveMemberTwoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ApprovalCreateTime(approvalCreateTime string) *BillListItemOfWangYCBuilder {
	builder.approvalCreateTime = approvalCreateTime
	builder.approvalCreateTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ApprovalExtraInfo(approvalExtraInfo string) *BillListItemOfWangYCBuilder {
	builder.approvalExtraInfo = approvalExtraInfo
	builder.approvalExtraInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ApprovalLogs(approvalLogs string) *BillListItemOfWangYCBuilder {
	builder.approvalLogs = approvalLogs
	builder.approvalLogsSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ApprovalReason(approvalReason string) *BillListItemOfWangYCBuilder {
	builder.approvalReason = approvalReason
	builder.approvalReasonSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BeforeApprovalResult(beforeApprovalResult string) *BillListItemOfWangYCBuilder {
	builder.beforeApprovalResult = beforeApprovalResult
	builder.beforeApprovalResultSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BeginChargeTime(beginChargeTime string) *BillListItemOfWangYCBuilder {
	builder.beginChargeTime = beginChargeTime
	builder.beginChargeTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BelongBudgetCenterName(belongBudgetCenterName string) *BillListItemOfWangYCBuilder {
	builder.belongBudgetCenterName = belongBudgetCenterName
	builder.belongBudgetCenterNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BelongOutBudgetId(belongOutBudgetId string) *BillListItemOfWangYCBuilder {
	builder.belongOutBudgetId = belongOutBudgetId
	builder.belongOutBudgetIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BookingDepCode(bookingDepCode string) *BillListItemOfWangYCBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BookingServiceFee(bookingServiceFee float32) *BillListItemOfWangYCBuilder {
	builder.bookingServiceFee = bookingServiceFee
	builder.bookingServiceFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BranchName(branchName string) *BillListItemOfWangYCBuilder {
	builder.branchName = branchName
	builder.branchNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BridgeFee(bridgeFee float32) *BillListItemOfWangYCBuilder {
	builder.bridgeFee = bridgeFee
	builder.bridgeFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BudgetCenterName(budgetCenterName string) *BillListItemOfWangYCBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BudgetId(budgetId string) *BillListItemOfWangYCBuilder {
	builder.budgetId = budgetId
	builder.budgetIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) BudgetItemName(budgetItemName string) *BillListItemOfWangYCBuilder {
	builder.budgetItemName = budgetItemName
	builder.budgetItemNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CWeight(cWeight string) *BillListItemOfWangYCBuilder {
	builder.cWeight = cWeight
	builder.cWeightSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CallPhone(callPhone string) *BillListItemOfWangYCBuilder {
	builder.callPhone = callPhone
	builder.callPhoneSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CallPhoneCountryCode(callPhoneCountryCode string) *BillListItemOfWangYCBuilder {
	builder.callPhoneCountryCode = callPhoneCountryCode
	builder.callPhoneCountryCodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CallbackInfo(callbackInfo string) *BillListItemOfWangYCBuilder {
	builder.callbackInfo = callbackInfo
	builder.callbackInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CancelTime(cancelTime string) *BillListItemOfWangYCBuilder {
	builder.cancelTime = cancelTime
	builder.cancelTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CarArriveProvince(carArriveProvince string) *BillListItemOfWangYCBuilder {
	builder.carArriveProvince = carArriveProvince
	builder.carArriveProvinceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CarDepartProvince(carDepartProvince string) *BillListItemOfWangYCBuilder {
	builder.carDepartProvince = carDepartProvince
	builder.carDepartProvinceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CarTravelInfo(carTravelInfo string) *BillListItemOfWangYCBuilder {
	builder.carTravelInfo = carTravelInfo
	builder.carTravelInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ChargeTime(chargeTime string) *BillListItemOfWangYCBuilder {
	builder.chargeTime = chargeTime
	builder.chargeTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) City(city string) *BillListItemOfWangYCBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CleanCarFee(cleanCarFee float32) *BillListItemOfWangYCBuilder {
	builder.cleanCarFee = cleanCarFee
	builder.cleanCarFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Co2Emissions(co2Emissions string) *BillListItemOfWangYCBuilder {
	builder.co2Emissions = co2Emissions
	builder.co2EmissionsSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyCardRealPay(companyCardRealPay float32) *BillListItemOfWangYCBuilder {
	builder.companyCardRealPay = companyCardRealPay
	builder.companyCardRealPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyCardRealRefund(companyCardRealRefund float32) *BillListItemOfWangYCBuilder {
	builder.companyCardRealRefund = companyCardRealRefund
	builder.companyCardRealRefundSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyId(companyId int64) *BillListItemOfWangYCBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyName(companyName string) *BillListItemOfWangYCBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyPayAble(companyPayAble float32) *BillListItemOfWangYCBuilder {
	builder.companyPayAble = companyPayAble
	builder.companyPayAbleSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyRealPay(companyRealPay float32) *BillListItemOfWangYCBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyRealPreTaxPay(companyRealPreTaxPay float32) *BillListItemOfWangYCBuilder {
	builder.companyRealPreTaxPay = companyRealPreTaxPay
	builder.companyRealPreTaxPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyRealRefund(companyRealRefund float32) *BillListItemOfWangYCBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CompanyRealTaxPay(companyRealTaxPay float32) *BillListItemOfWangYCBuilder {
	builder.companyRealTaxPay = companyRealTaxPay
	builder.companyRealTaxPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Cost(cost float32) *BillListItemOfWangYCBuilder {
	builder.cost = cost
	builder.costSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CreateTime(createTime string) *BillListItemOfWangYCBuilder {
	builder.createTime = createTime
	builder.createTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) CrossCityFee(crossCityFee float32) *BillListItemOfWangYCBuilder {
	builder.crossCityFee = crossCityFee
	builder.crossCityFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DeductionTypeName(deductionTypeName string) *BillListItemOfWangYCBuilder {
	builder.deductionTypeName = deductionTypeName
	builder.deductionTypeNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Department(department string) *BillListItemOfWangYCBuilder {
	builder.department = department
	builder.departmentSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DepartmentName(departmentName string) *BillListItemOfWangYCBuilder {
	builder.departmentName = departmentName
	builder.departmentNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DepartureDime(departureDime string) *BillListItemOfWangYCBuilder {
	builder.departureDime = departureDime
	builder.departureDimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DestinationCity(destinationCity string) *BillListItemOfWangYCBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DiscountAfterPrice(discountAfterPrice float32) *BillListItemOfWangYCBuilder {
	builder.discountAfterPrice = discountAfterPrice
	builder.discountAfterPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) DynamicPrice(dynamicPrice float32) *BillListItemOfWangYCBuilder {
	builder.dynamicPrice = dynamicPrice
	builder.dynamicPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EmployeeNumber(employeeNumber string) *BillListItemOfWangYCBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EmptyFee(emptyFee float32) *BillListItemOfWangYCBuilder {
	builder.emptyFee = emptyFee
	builder.emptyFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EndAddress(endAddress string) *BillListItemOfWangYCBuilder {
	builder.endAddress = endAddress
	builder.endAddressSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EndCountyId(endCountyId string) *BillListItemOfWangYCBuilder {
	builder.endCountyId = endCountyId
	builder.endCountyIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EndCountyName(endCountyName string) *BillListItemOfWangYCBuilder {
	builder.endCountyName = endCountyName
	builder.endCountyNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EndName(endName string) *BillListItemOfWangYCBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EnergyConsumeFee(energyConsumeFee float32) *BillListItemOfWangYCBuilder {
	builder.energyConsumeFee = energyConsumeFee
	builder.energyConsumeFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EnterpriseCouponBatch(enterpriseCouponBatch string) *BillListItemOfWangYCBuilder {
	builder.enterpriseCouponBatch = enterpriseCouponBatch
	builder.enterpriseCouponBatchSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EnterpriseCouponName(enterpriseCouponName string) *BillListItemOfWangYCBuilder {
	builder.enterpriseCouponName = enterpriseCouponName
	builder.enterpriseCouponNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EnterpriseInstantDiscount(enterpriseInstantDiscount float32) *BillListItemOfWangYCBuilder {
	builder.enterpriseInstantDiscount = enterpriseInstantDiscount
	builder.enterpriseInstantDiscountSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EstimatePrice(estimatePrice float32) *BillListItemOfWangYCBuilder {
	builder.estimatePrice = estimatePrice
	builder.estimatePriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) EstimatedDistance(estimatedDistance string) *BillListItemOfWangYCBuilder {
	builder.estimatedDistance = estimatedDistance
	builder.estimatedDistanceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo01(exInfo01 string) *BillListItemOfWangYCBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo01Code(exInfo01Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo01Code = exInfo01Code
	builder.exInfo01CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo02(exInfo02 string) *BillListItemOfWangYCBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo02Code(exInfo02Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo02Code = exInfo02Code
	builder.exInfo02CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo03(exInfo03 string) *BillListItemOfWangYCBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo03Code(exInfo03Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo03Code = exInfo03Code
	builder.exInfo03CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo04(exInfo04 string) *BillListItemOfWangYCBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo04Code(exInfo04Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo04Code = exInfo04Code
	builder.exInfo04CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo05(exInfo05 string) *BillListItemOfWangYCBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo05Code(exInfo05Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo05Code = exInfo05Code
	builder.exInfo05CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo06(exInfo06 string) *BillListItemOfWangYCBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo06Code(exInfo06Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo06Code = exInfo06Code
	builder.exInfo06CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo07(exInfo07 string) *BillListItemOfWangYCBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo07Code(exInfo07Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo07Code = exInfo07Code
	builder.exInfo07CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo08(exInfo08 string) *BillListItemOfWangYCBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExInfo08Code(exInfo08Code string) *BillListItemOfWangYCBuilder {
	builder.exInfo08Code = exInfo08Code
	builder.exInfo08CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExcludingServiceFee(excludingServiceFee float32) *BillListItemOfWangYCBuilder {
	builder.excludingServiceFee = excludingServiceFee
	builder.excludingServiceFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExcludingTaxPayPrice(excludingTaxPayPrice float32) *BillListItemOfWangYCBuilder {
	builder.excludingTaxPayPrice = excludingTaxPayPrice
	builder.excludingTaxPayPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExcludingTaxRefundPrice(excludingTaxRefundPrice float32) *BillListItemOfWangYCBuilder {
	builder.excludingTaxRefundPrice = excludingTaxRefundPrice
	builder.excludingTaxRefundPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ExtendInfo(extendInfo string) *BillListItemOfWangYCBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) FinishTime(finishTime string) *BillListItemOfWangYCBuilder {
	builder.finishTime = finishTime
	builder.finishTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) FlagExt(flagExt string) *BillListItemOfWangYCBuilder {
	builder.flagExt = flagExt
	builder.flagExtSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) FromIsWorkPlace(fromIsWorkPlace string) *BillListItemOfWangYCBuilder {
	builder.fromIsWorkPlace = fromIsWorkPlace
	builder.fromIsWorkPlaceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group1Code(group1Code string) *BillListItemOfWangYCBuilder {
	builder.group1Code = group1Code
	builder.group1CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group1Name(group1Name string) *BillListItemOfWangYCBuilder {
	builder.group1Name = group1Name
	builder.group1NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group2Code(group2Code string) *BillListItemOfWangYCBuilder {
	builder.group2Code = group2Code
	builder.group2CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group2Name(group2Name string) *BillListItemOfWangYCBuilder {
	builder.group2Name = group2Name
	builder.group2NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group3Code(group3Code string) *BillListItemOfWangYCBuilder {
	builder.group3Code = group3Code
	builder.group3CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group3Name(group3Name string) *BillListItemOfWangYCBuilder {
	builder.group3Name = group3Name
	builder.group3NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group4Code(group4Code string) *BillListItemOfWangYCBuilder {
	builder.group4Code = group4Code
	builder.group4CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group4Name(group4Name string) *BillListItemOfWangYCBuilder {
	builder.group4Name = group4Name
	builder.group4NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group5Code(group5Code string) *BillListItemOfWangYCBuilder {
	builder.group5Code = group5Code
	builder.group5CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group5Name(group5Name string) *BillListItemOfWangYCBuilder {
	builder.group5Name = group5Name
	builder.group5NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group6Code(group6Code string) *BillListItemOfWangYCBuilder {
	builder.group6Code = group6Code
	builder.group6CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group6Name(group6Name string) *BillListItemOfWangYCBuilder {
	builder.group6Name = group6Name
	builder.group6NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group7Code(group7Code string) *BillListItemOfWangYCBuilder {
	builder.group7Code = group7Code
	builder.group7CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group7Name(group7Name string) *BillListItemOfWangYCBuilder {
	builder.group7Name = group7Name
	builder.group7NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group8Code(group8Code string) *BillListItemOfWangYCBuilder {
	builder.group8Code = group8Code
	builder.group8CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group8Name(group8Name string) *BillListItemOfWangYCBuilder {
	builder.group8Name = group8Name
	builder.group8NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group9Code(group9Code string) *BillListItemOfWangYCBuilder {
	builder.group9Code = group9Code
	builder.group9CodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Group9Name(group9Name string) *BillListItemOfWangYCBuilder {
	builder.group9Name = group9Name
	builder.group9NameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) GroupId(groupId string) *BillListItemOfWangYCBuilder {
	builder.groupId = groupId
	builder.groupIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) GroupName(groupName string) *BillListItemOfWangYCBuilder {
	builder.groupName = groupName
	builder.groupNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) HighwayFee(highwayFee float32) *BillListItemOfWangYCBuilder {
	builder.highwayFee = highwayFee
	builder.highwayFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IncentiveFee(incentiveFee float32) *BillListItemOfWangYCBuilder {
	builder.incentiveFee = incentiveFee
	builder.incentiveFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) InstitutionId(institutionId float32) *BillListItemOfWangYCBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) InstitutionName(institutionName string) *BillListItemOfWangYCBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) InvoiceEntityInfo(invoiceEntityInfo string) *BillListItemOfWangYCBuilder {
	builder.invoiceEntityInfo = invoiceEntityInfo
	builder.invoiceEntityInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsCarpool(isCarpool string) *BillListItemOfWangYCBuilder {
	builder.isCarpool = isCarpool
	builder.isCarpoolSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsCarpoolSuccess(isCarpoolSuccess float32) *BillListItemOfWangYCBuilder {
	builder.isCarpoolSuccess = isCarpoolSuccess
	builder.isCarpoolSuccessSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsCurrentTerm(isCurrentTerm string) *BillListItemOfWangYCBuilder {
	builder.isCurrentTerm = isCurrentTerm
	builder.isCurrentTermSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsDiscountOrder(isDiscountOrder string) *BillListItemOfWangYCBuilder {
	builder.isDiscountOrder = isDiscountOrder
	builder.isDiscountOrderSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsFixedPrice(isFixedPrice string) *BillListItemOfWangYCBuilder {
	builder.isFixedPrice = isFixedPrice
	builder.isFixedPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsReassignment(isReassignment string) *BillListItemOfWangYCBuilder {
	builder.isReassignment = isReassignment
	builder.isReassignmentSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsSelfDrive(isSelfDrive string) *BillListItemOfWangYCBuilder {
	builder.isSelfDrive = isSelfDrive
	builder.isSelfDriveSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsSensitive(isSensitive float32) *BillListItemOfWangYCBuilder {
	builder.isSensitive = isSensitive
	builder.isSensitiveSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) IsUnusual(isUnusual string) *BillListItemOfWangYCBuilder {
	builder.isUnusual = isUnusual
	builder.isUnusualSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LastApprovalTime(lastApprovalTime string) *BillListItemOfWangYCBuilder {
	builder.lastApprovalTime = lastApprovalTime
	builder.lastApprovalTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LastApproves(lastApproves string) *BillListItemOfWangYCBuilder {
	builder.lastApproves = lastApproves
	builder.lastApprovesSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LegalEntityId(legalEntityId int64) *BillListItemOfWangYCBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LegalEntityName(legalEntityName string) *BillListItemOfWangYCBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LimitFee(limitFee float32) *BillListItemOfWangYCBuilder {
	builder.limitFee = limitFee
	builder.limitFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LimitPay(limitPay float32) *BillListItemOfWangYCBuilder {
	builder.limitPay = limitPay
	builder.limitPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) LowSpeedFee(lowSpeedFee float32) *BillListItemOfWangYCBuilder {
	builder.lowSpeedFee = lowSpeedFee
	builder.lowSpeedFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) MeetTime(meetTime string) *BillListItemOfWangYCBuilder {
	builder.meetTime = meetTime
	builder.meetTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) MemberId(memberId int64) *BillListItemOfWangYCBuilder {
	builder.memberId = memberId
	builder.memberIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) MemberMail(memberMail string) *BillListItemOfWangYCBuilder {
	builder.memberMail = memberMail
	builder.memberMailSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) MemberName(memberName string) *BillListItemOfWangYCBuilder {
	builder.memberName = memberName
	builder.memberNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) NightFee(nightFee float32) *BillListItemOfWangYCBuilder {
	builder.nightFee = nightFee
	builder.nightFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) NormalDistance(normalDistance float32) *BillListItemOfWangYCBuilder {
	builder.normalDistance = normalDistance
	builder.normalDistanceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) NormalFee(normalFee float32) *BillListItemOfWangYCBuilder {
	builder.normalFee = normalFee
	builder.normalFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) NormalTimeFee(normalTimeFee float32) *BillListItemOfWangYCBuilder {
	builder.normalTimeFee = normalTimeFee
	builder.normalTimeFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OneTimeOfferSubsidy(oneTimeOfferSubsidy float32) *BillListItemOfWangYCBuilder {
	builder.oneTimeOfferSubsidy = oneTimeOfferSubsidy
	builder.oneTimeOfferSubsidySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OrderAdditionalRemark(orderAdditionalRemark string) *BillListItemOfWangYCBuilder {
	builder.orderAdditionalRemark = orderAdditionalRemark
	builder.orderAdditionalRemarkSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OrderId(orderId int64) *BillListItemOfWangYCBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OrderSource(orderSource string) *BillListItemOfWangYCBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OriginalEndAddress(originalEndAddress string) *BillListItemOfWangYCBuilder {
	builder.originalEndAddress = originalEndAddress
	builder.originalEndAddressSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OriginalEndName(originalEndName string) *BillListItemOfWangYCBuilder {
	builder.originalEndName = originalEndName
	builder.originalEndNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OriginalPrice(originalPrice float32) *BillListItemOfWangYCBuilder {
	builder.originalPrice = originalPrice
	builder.originalPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OtherFee(otherFee float32) *BillListItemOfWangYCBuilder {
	builder.otherFee = otherFee
	builder.otherFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OutApprovalId(outApprovalId string) *BillListItemOfWangYCBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OutBudgetId(outBudgetId string) *BillListItemOfWangYCBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OutLegalEntityId(outLegalEntityId string) *BillListItemOfWangYCBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) OutOrderId(outOrderId int64) *BillListItemOfWangYCBuilder {
	builder.outOrderId = outOrderId
	builder.outOrderIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ParkFee(parkFee float32) *BillListItemOfWangYCBuilder {
	builder.parkFee = parkFee
	builder.parkFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PassengerMemberId(passengerMemberId int64) *BillListItemOfWangYCBuilder {
	builder.passengerMemberId = passengerMemberId
	builder.passengerMemberIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillListItemOfWangYCBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PassengerName(passengerName string) *BillListItemOfWangYCBuilder {
	builder.passengerName = passengerName
	builder.passengerNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PassengerPhone(passengerPhone string) *BillListItemOfWangYCBuilder {
	builder.passengerPhone = passengerPhone
	builder.passengerPhoneSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PassengerPhoneCountryCode(passengerPhoneCountryCode string) *BillListItemOfWangYCBuilder {
	builder.passengerPhoneCountryCode = passengerPhoneCountryCode
	builder.passengerPhoneCountryCodeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PayChannel(payChannel string) *BillListItemOfWangYCBuilder {
	builder.payChannel = payChannel
	builder.payChannelSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PayTime(payTime string) *BillListItemOfWangYCBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PayType(payType string) *BillListItemOfWangYCBuilder {
	builder.payType = payType
	builder.payTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Period(period string) *BillListItemOfWangYCBuilder {
	builder.period = period
	builder.periodSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PersonalInstantDiscount(personalInstantDiscount float32) *BillListItemOfWangYCBuilder {
	builder.personalInstantDiscount = personalInstantDiscount
	builder.personalInstantDiscountSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PersonalRealPay(personalRealPay float32) *BillListItemOfWangYCBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) PositionName(positionName string) *BillListItemOfWangYCBuilder {
	builder.positionName = positionName
	builder.positionNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ProjectExtInfo(projectExtInfo string) *BillListItemOfWangYCBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RealVoucherPay(realVoucherPay float32) *BillListItemOfWangYCBuilder {
	builder.realVoucherPay = realVoucherPay
	builder.realVoucherPaySet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RedPacket(redPacket float32) *BillListItemOfWangYCBuilder {
	builder.redPacket = redPacket
	builder.redPacketSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RefundPrice(refundPrice float32) *BillListItemOfWangYCBuilder {
	builder.refundPrice = refundPrice
	builder.refundPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RefundTime(refundTime string) *BillListItemOfWangYCBuilder {
	builder.refundTime = refundTime
	builder.refundTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Remark(remark string) *BillListItemOfWangYCBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RemoteAreaFee(remoteAreaFee float32) *BillListItemOfWangYCBuilder {
	builder.remoteAreaFee = remoteAreaFee
	builder.remoteAreaFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RequireLevel(requireLevel string) *BillListItemOfWangYCBuilder {
	builder.requireLevel = requireLevel
	builder.requireLevelSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RequisitionId(requisitionId string) *BillListItemOfWangYCBuilder {
	builder.requisitionId = requisitionId
	builder.requisitionIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ResidentCityNames(residentCityNames string) *BillListItemOfWangYCBuilder {
	builder.residentCityNames = residentCityNames
	builder.residentCityNamesSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ResponsibleCancelFee(responsibleCancelFee float32) *BillListItemOfWangYCBuilder {
	builder.responsibleCancelFee = responsibleCancelFee
	builder.responsibleCancelFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Rule(rule string) *BillListItemOfWangYCBuilder {
	builder.rule = rule
	builder.ruleSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) RuleName(ruleName string) *BillListItemOfWangYCBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SensitiveReason(sensitiveReason string) *BillListItemOfWangYCBuilder {
	builder.sensitiveReason = sensitiveReason
	builder.sensitiveReasonSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ServiceType(serviceType string) *BillListItemOfWangYCBuilder {
	builder.serviceType = serviceType
	builder.serviceTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SettleType(settleType string) *BillListItemOfWangYCBuilder {
	builder.settleType = settleType
	builder.settleTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SpsPickUpServiceFee(spsPickUpServiceFee float32) *BillListItemOfWangYCBuilder {
	builder.spsPickUpServiceFee = spsPickUpServiceFee
	builder.spsPickUpServiceFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StartAddress(startAddress string) *BillListItemOfWangYCBuilder {
	builder.startAddress = startAddress
	builder.startAddressSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StartCountyId(startCountyId string) *BillListItemOfWangYCBuilder {
	builder.startCountyId = startCountyId
	builder.startCountyIdSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StartCountyName(startCountyName string) *BillListItemOfWangYCBuilder {
	builder.startCountyName = startCountyName
	builder.startCountyNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StartName(startName string) *BillListItemOfWangYCBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StartPrice(startPrice float32) *BillListItemOfWangYCBuilder {
	builder.startPrice = startPrice
	builder.startPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Status(status string) *BillListItemOfWangYCBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) StriveTime(striveTime string) *BillListItemOfWangYCBuilder {
	builder.striveTime = striveTime
	builder.striveTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SubAccountCompanyName(subAccountCompanyName string) *BillListItemOfWangYCBuilder {
	builder.subAccountCompanyName = subAccountCompanyName
	builder.subAccountCompanyNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SubAccountName(subAccountName string) *BillListItemOfWangYCBuilder {
	builder.subAccountName = subAccountName
	builder.subAccountNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SubUseCarSrv(subUseCarSrv string) *BillListItemOfWangYCBuilder {
	builder.subUseCarSrv = subUseCarSrv
	builder.subUseCarSrvSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SubsidyAmount(subsidyAmount float32) *BillListItemOfWangYCBuilder {
	builder.subsidyAmount = subsidyAmount
	builder.subsidyAmountSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SupplierCarName(supplierCarName string) *BillListItemOfWangYCBuilder {
	builder.supplierCarName = supplierCarName
	builder.supplierCarNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SupplierName(supplierName string) *BillListItemOfWangYCBuilder {
	builder.supplierName = supplierName
	builder.supplierNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) SupplierType(supplierType string) *BillListItemOfWangYCBuilder {
	builder.supplierType = supplierType
	builder.supplierTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) TaxPayPrice(taxPayPrice float32) *BillListItemOfWangYCBuilder {
	builder.taxPayPrice = taxPayPrice
	builder.taxPayPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) TaxRefundPrice(taxRefundPrice float32) *BillListItemOfWangYCBuilder {
	builder.taxRefundPrice = taxRefundPrice
	builder.taxRefundPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) TimeSlotDiscount(timeSlotDiscount float32) *BillListItemOfWangYCBuilder {
	builder.timeSlotDiscount = timeSlotDiscount
	builder.timeSlotDiscountSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) TipFee(tipFee float32) *BillListItemOfWangYCBuilder {
	builder.tipFee = tipFee
	builder.tipFeeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) ToIsWorkPlace(toIsWorkPlace string) *BillListItemOfWangYCBuilder {
	builder.toIsWorkPlace = toIsWorkPlace
	builder.toIsWorkPlaceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) TotalPrice(totalPrice float32) *BillListItemOfWangYCBuilder {
	builder.totalPrice = totalPrice
	builder.totalPriceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) Type(type_ string) *BillListItemOfWangYCBuilder {
	builder.type_ = type_
	builder.type_Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UnusualContent(unusualContent string) *BillListItemOfWangYCBuilder {
	builder.unusualContent = unusualContent
	builder.unusualContentSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UnusualType(unusualType string) *BillListItemOfWangYCBuilder {
	builder.unusualType = unusualType
	builder.unusualTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UpgradeBaseRequireLevel(upgradeBaseRequireLevel string) *BillListItemOfWangYCBuilder {
	builder.upgradeBaseRequireLevel = upgradeBaseRequireLevel
	builder.upgradeBaseRequireLevelSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UpgradeCompanyPayQuota(upgradeCompanyPayQuota float32) *BillListItemOfWangYCBuilder {
	builder.upgradeCompanyPayQuota = upgradeCompanyPayQuota
	builder.upgradeCompanyPayQuotaSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UpgradeType(upgradeType string) *BillListItemOfWangYCBuilder {
	builder.upgradeType = upgradeType
	builder.upgradeTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UseCarService(useCarService string) *BillListItemOfWangYCBuilder {
	builder.useCarService = useCarService
	builder.useCarServiceSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UseCarSrv(useCarSrv string) *BillListItemOfWangYCBuilder {
	builder.useCarSrv = useCarSrv
	builder.useCarSrvSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UseCarTypeName(useCarTypeName string) *BillListItemOfWangYCBuilder {
	builder.useCarTypeName = useCarTypeName
	builder.useCarTypeNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UseCarTypeV2(useCarTypeV2 string) *BillListItemOfWangYCBuilder {
	builder.useCarTypeV2 = useCarTypeV2
	builder.useCarTypeV2Set = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UseType(useType string) *BillListItemOfWangYCBuilder {
	builder.useType = useType
	builder.useTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UserPayTime(userPayTime string) *BillListItemOfWangYCBuilder {
	builder.userPayTime = userPayTime
	builder.userPayTimeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) UserPayType(userPayType string) *BillListItemOfWangYCBuilder {
	builder.userPayType = userPayType
	builder.userPayTypeSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) VoucherDeductionTypeName(voucherDeductionTypeName string) *BillListItemOfWangYCBuilder {
	builder.voucherDeductionTypeName = voucherDeductionTypeName
	builder.voucherDeductionTypeNameSet = true
	return builder
}
func (builder *BillListItemOfWangYCBuilder) WaitFee(waitFee float32) *BillListItemOfWangYCBuilder {
	builder.waitFee = waitFee
	builder.waitFeeSet = true
	return builder
}

func (builder *BillListItemOfWangYCBuilder) Build() *BillListItemOfWangYC {
	data := &BillListItemOfWangYC{}
	if builder.actualEndNameSet {
		data.ActualEndName = &builder.actualEndName
	}
	if builder.actualFromIsWorkPlaceSet {
		data.ActualFromIsWorkPlace = &builder.actualFromIsWorkPlace
	}
	if builder.actualStartNameSet {
		data.ActualStartName = &builder.actualStartName
	}
	if builder.actualToIsWorkPlaceSet {
		data.ActualToIsWorkPlace = &builder.actualToIsWorkPlace
	}
	if builder.afterApprovalContentSet {
		data.AfterApprovalContent = &builder.afterApprovalContent
	}
	if builder.afterApprovalResultSet {
		data.AfterApprovalResult = &builder.afterApprovalResult
	}
	if builder.afterApproveMemberOneSet {
		data.AfterApproveMemberOne = &builder.afterApproveMemberOne
	}
	if builder.afterApproveMemberThreeSet {
		data.AfterApproveMemberThree = &builder.afterApproveMemberThree
	}
	if builder.afterApproveMemberTwoSet {
		data.AfterApproveMemberTwo = &builder.afterApproveMemberTwo
	}
	if builder.approvalCreateTimeSet {
		data.ApprovalCreateTime = &builder.approvalCreateTime
	}
	if builder.approvalExtraInfoSet {
		data.ApprovalExtraInfo = &builder.approvalExtraInfo
	}
	if builder.approvalLogsSet {
		data.ApprovalLogs = &builder.approvalLogs
	}
	if builder.approvalReasonSet {
		data.ApprovalReason = &builder.approvalReason
	}
	if builder.beforeApprovalResultSet {
		data.BeforeApprovalResult = &builder.beforeApprovalResult
	}
	if builder.beginChargeTimeSet {
		data.BeginChargeTime = &builder.beginChargeTime
	}
	if builder.belongBudgetCenterNameSet {
		data.BelongBudgetCenterName = &builder.belongBudgetCenterName
	}
	if builder.belongOutBudgetIdSet {
		data.BelongOutBudgetId = &builder.belongOutBudgetId
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.bookingServiceFeeSet {
		data.BookingServiceFee = &builder.bookingServiceFee
	}
	if builder.branchNameSet {
		data.BranchName = &builder.branchName
	}
	if builder.bridgeFeeSet {
		data.BridgeFee = &builder.bridgeFee
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetIdSet {
		data.BudgetId = &builder.budgetId
	}
	if builder.budgetItemNameSet {
		data.BudgetItemName = &builder.budgetItemName
	}
	if builder.cWeightSet {
		data.CWeight = &builder.cWeight
	}
	if builder.callPhoneSet {
		data.CallPhone = &builder.callPhone
	}
	if builder.callPhoneCountryCodeSet {
		data.CallPhoneCountryCode = &builder.callPhoneCountryCode
	}
	if builder.callbackInfoSet {
		data.CallbackInfo = &builder.callbackInfo
	}
	if builder.cancelTimeSet {
		data.CancelTime = &builder.cancelTime
	}
	if builder.carArriveProvinceSet {
		data.CarArriveProvince = &builder.carArriveProvince
	}
	if builder.carDepartProvinceSet {
		data.CarDepartProvince = &builder.carDepartProvince
	}
	if builder.carTravelInfoSet {
		data.CarTravelInfo = &builder.carTravelInfo
	}
	if builder.chargeTimeSet {
		data.ChargeTime = &builder.chargeTime
	}
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.cleanCarFeeSet {
		data.CleanCarFee = &builder.cleanCarFee
	}
	if builder.co2EmissionsSet {
		data.Co2Emissions = &builder.co2Emissions
	}
	if builder.companyCardRealPaySet {
		data.CompanyCardRealPay = &builder.companyCardRealPay
	}
	if builder.companyCardRealRefundSet {
		data.CompanyCardRealRefund = &builder.companyCardRealRefund
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.companyNameSet {
		data.CompanyName = &builder.companyName
	}
	if builder.companyPayAbleSet {
		data.CompanyPayAble = &builder.companyPayAble
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.companyRealPreTaxPaySet {
		data.CompanyRealPreTaxPay = &builder.companyRealPreTaxPay
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.companyRealTaxPaySet {
		data.CompanyRealTaxPay = &builder.companyRealTaxPay
	}
	if builder.costSet {
		data.Cost = &builder.cost
	}
	if builder.createTimeSet {
		data.CreateTime = &builder.createTime
	}
	if builder.crossCityFeeSet {
		data.CrossCityFee = &builder.crossCityFee
	}
	if builder.deductionTypeNameSet {
		data.DeductionTypeName = &builder.deductionTypeName
	}
	if builder.departmentSet {
		data.Department = &builder.department
	}
	if builder.departmentNameSet {
		data.DepartmentName = &builder.departmentName
	}
	if builder.departureDimeSet {
		data.DepartureDime = &builder.departureDime
	}
	if builder.destinationCitySet {
		data.DestinationCity = &builder.destinationCity
	}
	if builder.discountAfterPriceSet {
		data.DiscountAfterPrice = &builder.discountAfterPrice
	}
	if builder.dynamicPriceSet {
		data.DynamicPrice = &builder.dynamicPrice
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emptyFeeSet {
		data.EmptyFee = &builder.emptyFee
	}
	if builder.endAddressSet {
		data.EndAddress = &builder.endAddress
	}
	if builder.endCountyIdSet {
		data.EndCountyId = &builder.endCountyId
	}
	if builder.endCountyNameSet {
		data.EndCountyName = &builder.endCountyName
	}
	if builder.endNameSet {
		data.EndName = &builder.endName
	}
	if builder.energyConsumeFeeSet {
		data.EnergyConsumeFee = &builder.energyConsumeFee
	}
	if builder.enterpriseCouponBatchSet {
		data.EnterpriseCouponBatch = &builder.enterpriseCouponBatch
	}
	if builder.enterpriseCouponNameSet {
		data.EnterpriseCouponName = &builder.enterpriseCouponName
	}
	if builder.enterpriseInstantDiscountSet {
		data.EnterpriseInstantDiscount = &builder.enterpriseInstantDiscount
	}
	if builder.estimatePriceSet {
		data.EstimatePrice = &builder.estimatePrice
	}
	if builder.estimatedDistanceSet {
		data.EstimatedDistance = &builder.estimatedDistance
	}
	if builder.exInfo01Set {
		data.ExInfo01 = &builder.exInfo01
	}
	if builder.exInfo01CodeSet {
		data.ExInfo01Code = &builder.exInfo01Code
	}
	if builder.exInfo02Set {
		data.ExInfo02 = &builder.exInfo02
	}
	if builder.exInfo02CodeSet {
		data.ExInfo02Code = &builder.exInfo02Code
	}
	if builder.exInfo03Set {
		data.ExInfo03 = &builder.exInfo03
	}
	if builder.exInfo03CodeSet {
		data.ExInfo03Code = &builder.exInfo03Code
	}
	if builder.exInfo04Set {
		data.ExInfo04 = &builder.exInfo04
	}
	if builder.exInfo04CodeSet {
		data.ExInfo04Code = &builder.exInfo04Code
	}
	if builder.exInfo05Set {
		data.ExInfo05 = &builder.exInfo05
	}
	if builder.exInfo05CodeSet {
		data.ExInfo05Code = &builder.exInfo05Code
	}
	if builder.exInfo06Set {
		data.ExInfo06 = &builder.exInfo06
	}
	if builder.exInfo06CodeSet {
		data.ExInfo06Code = &builder.exInfo06Code
	}
	if builder.exInfo07Set {
		data.ExInfo07 = &builder.exInfo07
	}
	if builder.exInfo07CodeSet {
		data.ExInfo07Code = &builder.exInfo07Code
	}
	if builder.exInfo08Set {
		data.ExInfo08 = &builder.exInfo08
	}
	if builder.exInfo08CodeSet {
		data.ExInfo08Code = &builder.exInfo08Code
	}
	if builder.excludingServiceFeeSet {
		data.ExcludingServiceFee = &builder.excludingServiceFee
	}
	if builder.excludingTaxPayPriceSet {
		data.ExcludingTaxPayPrice = &builder.excludingTaxPayPrice
	}
	if builder.excludingTaxRefundPriceSet {
		data.ExcludingTaxRefundPrice = &builder.excludingTaxRefundPrice
	}
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.finishTimeSet {
		data.FinishTime = &builder.finishTime
	}
	if builder.flagExtSet {
		data.FlagExt = &builder.flagExt
	}
	if builder.fromIsWorkPlaceSet {
		data.FromIsWorkPlace = &builder.fromIsWorkPlace
	}
	if builder.group1CodeSet {
		data.Group1Code = &builder.group1Code
	}
	if builder.group1NameSet {
		data.Group1Name = &builder.group1Name
	}
	if builder.group2CodeSet {
		data.Group2Code = &builder.group2Code
	}
	if builder.group2NameSet {
		data.Group2Name = &builder.group2Name
	}
	if builder.group3CodeSet {
		data.Group3Code = &builder.group3Code
	}
	if builder.group3NameSet {
		data.Group3Name = &builder.group3Name
	}
	if builder.group4CodeSet {
		data.Group4Code = &builder.group4Code
	}
	if builder.group4NameSet {
		data.Group4Name = &builder.group4Name
	}
	if builder.group5CodeSet {
		data.Group5Code = &builder.group5Code
	}
	if builder.group5NameSet {
		data.Group5Name = &builder.group5Name
	}
	if builder.group6CodeSet {
		data.Group6Code = &builder.group6Code
	}
	if builder.group6NameSet {
		data.Group6Name = &builder.group6Name
	}
	if builder.group7CodeSet {
		data.Group7Code = &builder.group7Code
	}
	if builder.group7NameSet {
		data.Group7Name = &builder.group7Name
	}
	if builder.group8CodeSet {
		data.Group8Code = &builder.group8Code
	}
	if builder.group8NameSet {
		data.Group8Name = &builder.group8Name
	}
	if builder.group9CodeSet {
		data.Group9Code = &builder.group9Code
	}
	if builder.group9NameSet {
		data.Group9Name = &builder.group9Name
	}
	if builder.groupIdSet {
		data.GroupId = &builder.groupId
	}
	if builder.groupNameSet {
		data.GroupName = &builder.groupName
	}
	if builder.highwayFeeSet {
		data.HighwayFee = &builder.highwayFee
	}
	if builder.incentiveFeeSet {
		data.IncentiveFee = &builder.incentiveFee
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.invoiceEntityInfoSet {
		data.InvoiceEntityInfo = &builder.invoiceEntityInfo
	}
	if builder.isCarpoolSet {
		data.IsCarpool = &builder.isCarpool
	}
	if builder.isCarpoolSuccessSet {
		data.IsCarpoolSuccess = &builder.isCarpoolSuccess
	}
	if builder.isCurrentTermSet {
		data.IsCurrentTerm = &builder.isCurrentTerm
	}
	if builder.isDiscountOrderSet {
		data.IsDiscountOrder = &builder.isDiscountOrder
	}
	if builder.isFixedPriceSet {
		data.IsFixedPrice = &builder.isFixedPrice
	}
	if builder.isReassignmentSet {
		data.IsReassignment = &builder.isReassignment
	}
	if builder.isSelfDriveSet {
		data.IsSelfDrive = &builder.isSelfDrive
	}
	if builder.isSensitiveSet {
		data.IsSensitive = &builder.isSensitive
	}
	if builder.isUnusualSet {
		data.IsUnusual = &builder.isUnusual
	}
	if builder.lastApprovalTimeSet {
		data.LastApprovalTime = &builder.lastApprovalTime
	}
	if builder.lastApprovesSet {
		data.LastApproves = &builder.lastApproves
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.limitFeeSet {
		data.LimitFee = &builder.limitFee
	}
	if builder.limitPaySet {
		data.LimitPay = &builder.limitPay
	}
	if builder.lowSpeedFeeSet {
		data.LowSpeedFee = &builder.lowSpeedFee
	}
	if builder.meetTimeSet {
		data.MeetTime = &builder.meetTime
	}
	if builder.memberIdSet {
		data.MemberId = &builder.memberId
	}
	if builder.memberMailSet {
		data.MemberMail = &builder.memberMail
	}
	if builder.memberNameSet {
		data.MemberName = &builder.memberName
	}
	if builder.nightFeeSet {
		data.NightFee = &builder.nightFee
	}
	if builder.normalDistanceSet {
		data.NormalDistance = &builder.normalDistance
	}
	if builder.normalFeeSet {
		data.NormalFee = &builder.normalFee
	}
	if builder.normalTimeFeeSet {
		data.NormalTimeFee = &builder.normalTimeFee
	}
	if builder.oneTimeOfferSubsidySet {
		data.OneTimeOfferSubsidy = &builder.oneTimeOfferSubsidy
	}
	if builder.orderAdditionalRemarkSet {
		data.OrderAdditionalRemark = &builder.orderAdditionalRemark
	}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.orderSourceSet {
		data.OrderSource = &builder.orderSource
	}
	if builder.originalEndAddressSet {
		data.OriginalEndAddress = &builder.originalEndAddress
	}
	if builder.originalEndNameSet {
		data.OriginalEndName = &builder.originalEndName
	}
	if builder.originalPriceSet {
		data.OriginalPrice = &builder.originalPrice
	}
	if builder.otherFeeSet {
		data.OtherFee = &builder.otherFee
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.outOrderIdSet {
		data.OutOrderId = &builder.outOrderId
	}
	if builder.parkFeeSet {
		data.ParkFee = &builder.parkFee
	}
	if builder.passengerMemberIdSet {
		data.PassengerMemberId = &builder.passengerMemberId
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
	if builder.passengerPhoneCountryCodeSet {
		data.PassengerPhoneCountryCode = &builder.passengerPhoneCountryCode
	}
	if builder.payChannelSet {
		data.PayChannel = &builder.payChannel
	}
	if builder.payTimeSet {
		data.PayTime = &builder.payTime
	}
	if builder.payTypeSet {
		data.PayType = &builder.payType
	}
	if builder.periodSet {
		data.Period = &builder.period
	}
	if builder.personalInstantDiscountSet {
		data.PersonalInstantDiscount = &builder.personalInstantDiscount
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.positionNameSet {
		data.PositionName = &builder.positionName
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
	}
	if builder.realVoucherPaySet {
		data.RealVoucherPay = &builder.realVoucherPay
	}
	if builder.redPacketSet {
		data.RedPacket = &builder.redPacket
	}
	if builder.refundPriceSet {
		data.RefundPrice = &builder.refundPrice
	}
	if builder.refundTimeSet {
		data.RefundTime = &builder.refundTime
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.remoteAreaFeeSet {
		data.RemoteAreaFee = &builder.remoteAreaFee
	}
	if builder.requireLevelSet {
		data.RequireLevel = &builder.requireLevel
	}
	if builder.requisitionIdSet {
		data.RequisitionId = &builder.requisitionId
	}
	if builder.residentCityNamesSet {
		data.ResidentCityNames = &builder.residentCityNames
	}
	if builder.responsibleCancelFeeSet {
		data.ResponsibleCancelFee = &builder.responsibleCancelFee
	}
	if builder.ruleSet {
		data.Rule = &builder.rule
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.sensitiveReasonSet {
		data.SensitiveReason = &builder.sensitiveReason
	}
	if builder.serviceTypeSet {
		data.ServiceType = &builder.serviceType
	}
	if builder.settleTypeSet {
		data.SettleType = &builder.settleType
	}
	if builder.spsPickUpServiceFeeSet {
		data.SpsPickUpServiceFee = &builder.spsPickUpServiceFee
	}
	if builder.startAddressSet {
		data.StartAddress = &builder.startAddress
	}
	if builder.startCountyIdSet {
		data.StartCountyId = &builder.startCountyId
	}
	if builder.startCountyNameSet {
		data.StartCountyName = &builder.startCountyName
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.startPriceSet {
		data.StartPrice = &builder.startPrice
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.striveTimeSet {
		data.StriveTime = &builder.striveTime
	}
	if builder.subAccountCompanyNameSet {
		data.SubAccountCompanyName = &builder.subAccountCompanyName
	}
	if builder.subAccountNameSet {
		data.SubAccountName = &builder.subAccountName
	}
	if builder.subUseCarSrvSet {
		data.SubUseCarSrv = &builder.subUseCarSrv
	}
	if builder.subsidyAmountSet {
		data.SubsidyAmount = &builder.subsidyAmount
	}
	if builder.supplierCarNameSet {
		data.SupplierCarName = &builder.supplierCarName
	}
	if builder.supplierNameSet {
		data.SupplierName = &builder.supplierName
	}
	if builder.supplierTypeSet {
		data.SupplierType = &builder.supplierType
	}
	if builder.taxPayPriceSet {
		data.TaxPayPrice = &builder.taxPayPrice
	}
	if builder.taxRefundPriceSet {
		data.TaxRefundPrice = &builder.taxRefundPrice
	}
	if builder.timeSlotDiscountSet {
		data.TimeSlotDiscount = &builder.timeSlotDiscount
	}
	if builder.tipFeeSet {
		data.TipFee = &builder.tipFee
	}
	if builder.toIsWorkPlaceSet {
		data.ToIsWorkPlace = &builder.toIsWorkPlace
	}
	if builder.totalPriceSet {
		data.TotalPrice = &builder.totalPrice
	}
	if builder.type_Set {
		data.Type = &builder.type_
	}
	if builder.unusualContentSet {
		data.UnusualContent = &builder.unusualContent
	}
	if builder.unusualTypeSet {
		data.UnusualType = &builder.unusualType
	}
	if builder.upgradeBaseRequireLevelSet {
		data.UpgradeBaseRequireLevel = &builder.upgradeBaseRequireLevel
	}
	if builder.upgradeCompanyPayQuotaSet {
		data.UpgradeCompanyPayQuota = &builder.upgradeCompanyPayQuota
	}
	if builder.upgradeTypeSet {
		data.UpgradeType = &builder.upgradeType
	}
	if builder.useCarServiceSet {
		data.UseCarService = &builder.useCarService
	}
	if builder.useCarSrvSet {
		data.UseCarSrv = &builder.useCarSrv
	}
	if builder.useCarTypeNameSet {
		data.UseCarTypeName = &builder.useCarTypeName
	}
	if builder.useCarTypeV2Set {
		data.UseCarTypeV2 = &builder.useCarTypeV2
	}
	if builder.useTypeSet {
		data.UseType = &builder.useType
	}
	if builder.userPayTimeSet {
		data.UserPayTime = &builder.userPayTime
	}
	if builder.userPayTypeSet {
		data.UserPayType = &builder.userPayType
	}
	if builder.voucherDeductionTypeNameSet {
		data.VoucherDeductionTypeName = &builder.voucherDeductionTypeName
	}
	if builder.waitFeeSet {
		data.WaitFee = &builder.waitFee
	}
	return data
}
