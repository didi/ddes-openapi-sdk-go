package v1

// CreateApprovalBusinessByTimesRequest 创建用车按次数审批单入参
type CreateApprovalBusinessByTimesRequest struct {
	ClientId                      *string                    `json:"client_id,omitempty"`                        // 申请应用时分配的AppKey
	AccessToken                   *string                    `json:"access_token,omitempty"`                     // 授权后的access token
	Timestamp                     *int64                     `json:"timestamp,omitempty"`                        // 当前时间戳(精确到秒级)
	CompanyId                     *string                    `json:"company_id,omitempty"`                       // 滴滴公司 ID
	Sign                          *string                    `json:"sign,omitempty"`                             // 签名
	OutTripId                     *string                    `json:"out_trip_id,omitempty"`                      // 外部tripid，用于滴滴首页合并卡片展示
	ApprovalType                  *int32                     `json:"approval_type,omitempty"`                    // 申批单类型，枚举值数字：1:差旅单；2: 行前审批-按次数；3: 行前审批-按日期
	RegulationId                  *string                    `json:"regulation_id,omitempty"`                    // (差旅/行前)申请单对应的需审批制度ID regulation_id与policy_type、policy_type_value至少有一组有值
	PolicyType                    *int32                     `json:"policy_type,omitempty"`                      // 执行政策指定项类型，1：out_rank_id 外部职级编号；2：employee_number 员工工号；3：employee_phone 员工手机号；4：employee_email 员工邮箱；5：regulation_id 制度I，指定生效policy_type_value内对应的值
	PolicyTypeValue               *string                    `json:"policy_type_value,omitempty"`                // 执行政策指定项
	OutApprovalId                 *string                    `json:"out_approval_id,omitempty"`                  // 外部申请单ID，客户侧申请单ID，不大于 120 字符。eg: TA_100002，与返回中的approvalid一一对应
	BudgetCenterId                *string                    `json:"budget_center_id,omitempty"`                 // 滴滴侧成本中心ID，获取方式接口返回的 ID（可为项目/部门ID)，关联部门 时， id和out_budget_id 优先处理id关联项目时 id和 out_budget_id与 name 组合唯一值时 优先处理id ，out_budget_id与 name 同时有值时按照项目处理，out_budget_id有值且name为空时，按照部门处理。
	Name                          *string                    `json:"name,omitempty"`                             // 部门/项目名称，不大于 200 字符
	OutBudgetId                   *string                    `json:"out_budget_id,omitempty"`                    // 编号，长度限制：≤ 64 字符
	Reason                        *string                    `json:"reason,omitempty"`                           // 申请原因，申请原因，限制 200 字符；默认为空字符，eg: 北京出差，行程维度事由。使用卡片事由时，可以为空。
	MemberType                    *int32                     `json:"member_type,omitempty"`                      // 申请人唯一键类型，枚举值数字：0：手机号，phone 必填；1：工号，employee_number 必填；2：邮箱，email 必填；默认为0
	Phone                         *string                    `json:"phone,omitempty"`                            // 申请人手机号，申请人手机号；默认为空字符；member_type不传 或者member_type 为0时phone必传
	EmployeeNumber                *string                    `json:"employee_number,omitempty"`                  // 申请人工号，申请人工号，默认为空字符
	Email                         *string                    `json:"email,omitempty"`                            // 申请人邮箱，申请人邮箱；默认为空字符
	ExecutiveRegulationType       *int32                     `json:"executive_regulation_type,omitempty"`        // 代订执行制度类型，枚举值数字：0: 执行申请人制度；1: 执行外部出行人不限差标:（需要passenger_list都是外部出行人）；2: 执行出行人制度，作用于预订时生效的差标。；3: 执行多人多差标制度
	ExecutiveRegulationId         *string                    `json:"executive_regulation_id,omitempty"`          // 代订执行人的制度，代订执行人的制度，executive_regulation_type=2 时选填，传了就生效，不传：选择 executive_regulation_member 最新绑定的一条制度
	ExecutiveRegulationMemberType *int32                     `json:"executive_regulation_member_type,omitempty"` // 代订执行人标识类型，枚举值数字：0：手机号，executive_regulation_member 为执行人手机号；1：工号，executive_regulation_member 为执行人工号；2：邮箱：executive_regulation_member 为执行人邮箱；默认为0
	ExecutiveRegulationMember     *string                    `json:"executive_regulation_member,omitempty"`      // 代订执行人(手机号/工号/邮箱)，代订执行人(手机号/工号/邮箱)；默认手机号，注意：执行出行人制度的时候，需要是出行人passenger_list里的员工，executive_regulation_type=2时必传， 只能是内部员工
	OutTripInfo                   *string                    `json:"out_trip_info,omitempty"`                    // 外部trip信息，用于滴滴首页合并卡片展示
	ExtraInfo                     *string                    `json:"extra_info,omitempty"`                       // 扩展信息，扩展信息，自定义字段；最长不大于 500 字符；(必须为json字符串)；默认为空字符
	ExtendFieldList               *string                    `json:"extend_field_list,omitempty"`                // 扩展信息list，自定义字段，最长不大于 500 字符，将extend_field_list转为 json 字符串，三个字段仅作为备注性字段。详见extend_field_list
	BusinessTripDetail            *string                    `json:"business_trip_detail,omitempty"`             // 行前行程信息, 将 business_trip_detail转为 json 字符串approval_type = 2 或 approval_type = 3 必填。详见 [business_trip_detail](#business_trip_detail （按次数）
	TravelBudget                  *string                    `json:"travel_budget,omitempty"`                    // 差旅预算总额，差旅预算控制，一个申请单对应一个总预算纬度，一次只能全部或者选择部分品类。
	TravelManagement              *string                    `json:"travel_management,omitempty"`                // 差旅管控，差旅管控，目前仅支持市内用车每日限额设置（为json字符串类型，具体看请求示例）
	PassengerList                 *string                    `json:"passenger_list,omitempty"`                   // 出行人信息，出行人信息，不传时默认出行人为申请人，将passenger_list 转为 json 数组字符串。详见passenger_list
	BudgetCenterList              *string                    `json:"budget_center_list,omitempty"`               // 多成本中心(array)，序号1对应远成本中心字段，依然支持滴滴内部主键ID。使用部门CODE主键，和项目时名称和code作为主键。
	OutTripInfoObj                *OutTripInfo               `json:"out_trip_info_obj,omitempty"`
	ExtraInfoObj                  *map[string]string         `json:"extra_info__obj__,omitempty"` // 扩展信息。转成json赋值给extra_info字段
	ExtendFieldListObj            *ExtendFieldList           `json:"extend_field_list__obj__,omitempty"`
	BusinessTripDetailObj         *BusinessTripDetailByTimes `json:"business_trip_detail__obj__,omitempty"`
	TravelBudgetObj               *TravelBudget              `json:"travel_budget__obj__,omitempty"`
	TravelManagementObj           *TravelManagement          `json:"travel_management__obj__,omitempty"`
	PassengerListObj              []TripPassenger            `json:"passenger_list__obj__,omitempty"`     // 出行人信息,可使用脚本将其转换为json后赋值给 passenger_list 字段
	BudgetCenterListObj           []BudgetCenterListItem     `json:"budget_center_list__obj__,omitempty"` // 差旅预算总额,可使用脚本将其转换为json后赋值给 budget_center_list 字段
}

type CreateApprovalBusinessByTimesRequestBuilder struct {
	clientId                         string // 申请应用时分配的AppKey
	clientIdSet                      bool
	accessToken                      string // 授权后的access token
	accessTokenSet                   bool
	timestamp                        int64 // 当前时间戳(精确到秒级)
	timestampSet                     bool
	companyId                        string // 滴滴公司 ID
	companyIdSet                     bool
	sign                             string // 签名
	signSet                          bool
	outTripId                        string // 外部tripid，用于滴滴首页合并卡片展示
	outTripIdSet                     bool
	approvalType                     int32 // 申批单类型，枚举值数字：1:差旅单；2: 行前审批-按次数；3: 行前审批-按日期
	approvalTypeSet                  bool
	regulationId                     string // (差旅/行前)申请单对应的需审批制度ID regulation_id与policy_type、policy_type_value至少有一组有值
	regulationIdSet                  bool
	policyType                       int32 // 执行政策指定项类型，1：out_rank_id 外部职级编号；2：employee_number 员工工号；3：employee_phone 员工手机号；4：employee_email 员工邮箱；5：regulation_id 制度I，指定生效policy_type_value内对应的值
	policyTypeSet                    bool
	policyTypeValue                  string // 执行政策指定项
	policyTypeValueSet               bool
	outApprovalId                    string // 外部申请单ID，客户侧申请单ID，不大于 120 字符。eg: TA_100002，与返回中的approvalid一一对应
	outApprovalIdSet                 bool
	budgetCenterId                   string // 滴滴侧成本中心ID，获取方式接口返回的 ID（可为项目/部门ID)，关联部门 时， id和out_budget_id 优先处理id关联项目时 id和 out_budget_id与 name 组合唯一值时 优先处理id ，out_budget_id与 name 同时有值时按照项目处理，out_budget_id有值且name为空时，按照部门处理。
	budgetCenterIdSet                bool
	name                             string // 部门/项目名称，不大于 200 字符
	nameSet                          bool
	outBudgetId                      string // 编号，长度限制：≤ 64 字符
	outBudgetIdSet                   bool
	reason                           string // 申请原因，申请原因，限制 200 字符；默认为空字符，eg: 北京出差，行程维度事由。使用卡片事由时，可以为空。
	reasonSet                        bool
	memberType                       int32 // 申请人唯一键类型，枚举值数字：0：手机号，phone 必填；1：工号，employee_number 必填；2：邮箱，email 必填；默认为0
	memberTypeSet                    bool
	phone                            string // 申请人手机号，申请人手机号；默认为空字符；member_type不传 或者member_type 为0时phone必传
	phoneSet                         bool
	employeeNumber                   string // 申请人工号，申请人工号，默认为空字符
	employeeNumberSet                bool
	email                            string // 申请人邮箱，申请人邮箱；默认为空字符
	emailSet                         bool
	executiveRegulationType          int32 // 代订执行制度类型，枚举值数字：0: 执行申请人制度；1: 执行外部出行人不限差标:（需要passenger_list都是外部出行人）；2: 执行出行人制度，作用于预订时生效的差标。；3: 执行多人多差标制度
	executiveRegulationTypeSet       bool
	executiveRegulationId            string // 代订执行人的制度，代订执行人的制度，executive_regulation_type=2 时选填，传了就生效，不传：选择 executive_regulation_member 最新绑定的一条制度
	executiveRegulationIdSet         bool
	executiveRegulationMemberType    int32 // 代订执行人标识类型，枚举值数字：0：手机号，executive_regulation_member 为执行人手机号；1：工号，executive_regulation_member 为执行人工号；2：邮箱：executive_regulation_member 为执行人邮箱；默认为0
	executiveRegulationMemberTypeSet bool
	executiveRegulationMember        string // 代订执行人(手机号/工号/邮箱)，代订执行人(手机号/工号/邮箱)；默认手机号，注意：执行出行人制度的时候，需要是出行人passenger_list里的员工，executive_regulation_type=2时必传， 只能是内部员工
	executiveRegulationMemberSet     bool
	outTripInfo                      string // 外部trip信息，用于滴滴首页合并卡片展示
	outTripInfoSet                   bool
	extraInfo                        string // 扩展信息，扩展信息，自定义字段；最长不大于 500 字符；(必须为json字符串)；默认为空字符
	extraInfoSet                     bool
	extendFieldList                  string // 扩展信息list，自定义字段，最长不大于 500 字符，将extend_field_list转为 json 字符串，三个字段仅作为备注性字段。详见extend_field_list
	extendFieldListSet               bool
	businessTripDetail               string // 行前行程信息, 将 business_trip_detail转为 json 字符串approval_type = 2 或 approval_type = 3 必填。详见 [business_trip_detail](#business_trip_detail （按次数）
	businessTripDetailSet            bool
	travelBudget                     string // 差旅预算总额，差旅预算控制，一个申请单对应一个总预算纬度，一次只能全部或者选择部分品类。
	travelBudgetSet                  bool
	travelManagement                 string // 差旅管控，差旅管控，目前仅支持市内用车每日限额设置（为json字符串类型，具体看请求示例）
	travelManagementSet              bool
	passengerList                    string // 出行人信息，出行人信息，不传时默认出行人为申请人，将passenger_list 转为 json 数组字符串。详见passenger_list
	passengerListSet                 bool
	budgetCenterList                 string // 多成本中心(array)，序号1对应远成本中心字段，依然支持滴滴内部主键ID。使用部门CODE主键，和项目时名称和code作为主键。
	budgetCenterListSet              bool
	outTripInfoObj                   OutTripInfo
	outTripInfoObjSet                bool
	extraInfoObj                     map[string]string // 扩展信息。转成json赋值给extra_info字段
	extraInfoObjSet                  bool
	extendFieldListObj               ExtendFieldList
	extendFieldListObjSet            bool
	businessTripDetailObj            BusinessTripDetailByTimes
	businessTripDetailObjSet         bool
	travelBudgetObj                  TravelBudget
	travelBudgetObjSet               bool
	travelManagementObj              TravelManagement
	travelManagementObjSet           bool
	passengerListObj                 []TripPassenger // 出行人信息,可使用脚本将其转换为json后赋值给 passenger_list 字段
	passengerListObjSet              bool
	budgetCenterListObj              []BudgetCenterListItem // 差旅预算总额,可使用脚本将其转换为json后赋值给 budget_center_list 字段
	budgetCenterListObjSet           bool
}

func NewCreateApprovalBusinessByTimesRequestBuilder() *CreateApprovalBusinessByTimesRequestBuilder {
	return &CreateApprovalBusinessByTimesRequestBuilder{}
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ClientId(clientId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) AccessToken(accessToken string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Timestamp(timestamp int64) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) CompanyId(companyId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Sign(sign string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) OutTripId(outTripId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.outTripId = outTripId
	builder.outTripIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ApprovalType(approvalType int32) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) RegulationId(regulationId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) PolicyType(policyType int32) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.policyType = policyType
	builder.policyTypeSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) PolicyTypeValue(policyTypeValue string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.policyTypeValue = policyTypeValue
	builder.policyTypeValueSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) OutApprovalId(outApprovalId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) BudgetCenterId(budgetCenterId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Name(name string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) OutBudgetId(outBudgetId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Reason(reason string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) MemberType(memberType int32) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.memberType = memberType
	builder.memberTypeSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Phone(phone string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) EmployeeNumber(employeeNumber string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.employeeNumber = employeeNumber
	builder.employeeNumberSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) Email(email string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.email = email
	builder.emailSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExecutiveRegulationType(executiveRegulationType int32) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.executiveRegulationType = executiveRegulationType
	builder.executiveRegulationTypeSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExecutiveRegulationId(executiveRegulationId string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.executiveRegulationId = executiveRegulationId
	builder.executiveRegulationIdSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExecutiveRegulationMemberType(executiveRegulationMemberType int32) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.executiveRegulationMemberType = executiveRegulationMemberType
	builder.executiveRegulationMemberTypeSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExecutiveRegulationMember(executiveRegulationMember string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.executiveRegulationMember = executiveRegulationMember
	builder.executiveRegulationMemberSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) OutTripInfo(outTripInfo string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.outTripInfo = outTripInfo
	builder.outTripInfoSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExtraInfo(extraInfo string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExtendFieldList(extendFieldList string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.extendFieldList = extendFieldList
	builder.extendFieldListSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) BusinessTripDetail(businessTripDetail string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.businessTripDetail = businessTripDetail
	builder.businessTripDetailSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) TravelBudget(travelBudget string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.travelBudget = travelBudget
	builder.travelBudgetSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) TravelManagement(travelManagement string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.travelManagement = travelManagement
	builder.travelManagementSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) PassengerList(passengerList string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) BudgetCenterList(budgetCenterList string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) OutTripInfoObj(outTripInfoObj OutTripInfo) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.outTripInfoObj = outTripInfoObj
	builder.outTripInfoObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExtraInfoObj(extraInfoObj map[string]string) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.extraInfoObj = extraInfoObj
	builder.extraInfoObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) ExtendFieldListObj(extendFieldListObj ExtendFieldList) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.extendFieldListObj = extendFieldListObj
	builder.extendFieldListObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) BusinessTripDetailObj(businessTripDetailObj BusinessTripDetailByTimes) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.businessTripDetailObj = businessTripDetailObj
	builder.businessTripDetailObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) TravelBudgetObj(travelBudgetObj TravelBudget) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.travelBudgetObj = travelBudgetObj
	builder.travelBudgetObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) TravelManagementObj(travelManagementObj TravelManagement) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.travelManagementObj = travelManagementObj
	builder.travelManagementObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) PassengerListObj(passengerListObj []TripPassenger) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.passengerListObj = passengerListObj
	builder.passengerListObjSet = true
	return builder
}
func (builder *CreateApprovalBusinessByTimesRequestBuilder) BudgetCenterListObj(budgetCenterListObj []BudgetCenterListItem) *CreateApprovalBusinessByTimesRequestBuilder {
	builder.budgetCenterListObj = budgetCenterListObj
	builder.budgetCenterListObjSet = true
	return builder
}

func (builder *CreateApprovalBusinessByTimesRequestBuilder) Build() *CreateApprovalBusinessByTimesRequest {
	data := &CreateApprovalBusinessByTimesRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	if builder.outTripIdSet {
		data.OutTripId = &builder.outTripId
	}
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.policyTypeSet {
		data.PolicyType = &builder.policyType
	}
	if builder.policyTypeValueSet {
		data.PolicyTypeValue = &builder.policyTypeValue
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.memberTypeSet {
		data.MemberType = &builder.memberType
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.employeeNumberSet {
		data.EmployeeNumber = &builder.employeeNumber
	}
	if builder.emailSet {
		data.Email = &builder.email
	}
	if builder.executiveRegulationTypeSet {
		data.ExecutiveRegulationType = &builder.executiveRegulationType
	}
	if builder.executiveRegulationIdSet {
		data.ExecutiveRegulationId = &builder.executiveRegulationId
	}
	if builder.executiveRegulationMemberTypeSet {
		data.ExecutiveRegulationMemberType = &builder.executiveRegulationMemberType
	}
	if builder.executiveRegulationMemberSet {
		data.ExecutiveRegulationMember = &builder.executiveRegulationMember
	}
	if builder.outTripInfoSet {
		data.OutTripInfo = &builder.outTripInfo
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.extendFieldListSet {
		data.ExtendFieldList = &builder.extendFieldList
	}
	if builder.businessTripDetailSet {
		data.BusinessTripDetail = &builder.businessTripDetail
	}
	if builder.travelBudgetSet {
		data.TravelBudget = &builder.travelBudget
	}
	if builder.travelManagementSet {
		data.TravelManagement = &builder.travelManagement
	}
	if builder.passengerListSet {
		data.PassengerList = &builder.passengerList
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = &builder.budgetCenterList
	}
	if builder.outTripInfoObjSet {
		data.OutTripInfoObj = &builder.outTripInfoObj
	}
	if builder.extraInfoObjSet {
		data.ExtraInfoObj = &builder.extraInfoObj
	}
	if builder.extendFieldListObjSet {
		data.ExtendFieldListObj = &builder.extendFieldListObj
	}
	if builder.businessTripDetailObjSet {
		data.BusinessTripDetailObj = &builder.businessTripDetailObj
	}
	if builder.travelBudgetObjSet {
		data.TravelBudgetObj = &builder.travelBudgetObj
	}
	if builder.travelManagementObjSet {
		data.TravelManagementObj = &builder.travelManagementObj
	}
	if builder.passengerListObjSet {
		data.PassengerListObj = builder.passengerListObj
	}
	if builder.budgetCenterListObjSet {
		data.BudgetCenterListObj = builder.budgetCenterListObj
	}
	return data
}
