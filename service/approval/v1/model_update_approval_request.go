package v1

// UpdateApprovalRequest struct for UpdateApprovalRequest
type UpdateApprovalRequest struct {
	ClientId            *string                `json:"client_id,omitempty"`            // 申请应用时分配的AppKey
	AccessToken         *string                `json:"access_token,omitempty"`         // 授权后的access token
	Timestamp           *int64                 `json:"timestamp,omitempty"`            // 当前时间戳(精确到秒级)
	CompanyId           *string                `json:"company_id,omitempty"`           // 滴滴公司 ID
	Sign                *string                `json:"sign,omitempty"`                 // 签名
	ApprovalType        *int32                 `json:"approval_type,omitempty"`        // 申批单类型，枚举值数字：1:差旅单；2: 行前审批-按次数；3: 行前审批-按日期
	ApprovalId          *string                `json:"approval_id,omitempty"`          // 滴滴申请单ID。approval_id与out_approval_id至少需要传一个
	OutApprovalId       *string                `json:"out_approval_id,omitempty"`      // 外部申请单ID，外部申请单ID approval_id优先级大于out_approval_id
	Reason              *string                `json:"reason,omitempty"`               // 申请原因，申请原因，限制 200 字符；默认为空字符 eg: 北京出差，行程维度事由。使用卡片事由时，可以为空
	Name                *string                `json:"name,omitempty"`                 // 部门/项目名称，不大于 200 字符
	OutBudgetId         *string                `json:"out_budget_id,omitempty"`        // 编号，长度限制：≤ 64 字符
	TravelBudget        *string                `json:"travel_budget,omitempty"`        // 差旅预算总额，差旅预算控制，一个申请单对应一个总预算纬度，一次只能全部或者选择部分品类。字段不传输，修改不处理，字段传输，空对象时处理为清空。给值为修改
	TravelManagement    *string                `json:"travel_management,omitempty"`    // 差旅管控，差旅管控，目前仅支持市内用车每日限额设置（为json字符串类型，具体看请求示例）
	ExtraInfo           *string                `json:"extra_info,omitempty"`           // 扩展信息，扩展信息，自定义字段；最长不大于 500 字符；(必须为json字符串)；默认为空字符
	TravelDetail        *string                `json:"travel_detail,omitempty"`        // 差旅单行程信息，差旅单行程信息，将travel_detail的值转为 json 字符串，approval_type = 1 时必传
	PassengerList       *string                `json:"passenger_list,omitempty"`       // 出行人信息，出行人信息，不传时默认出行人为申请人，将passenger_list 转为 json 数组字符串。详见passenger_list
	BudgetCenterList    *string                `json:"budget_center_list,omitempty"`   // 多成本中心
	ExtendFieldList     *string                `json:"extend_field_list,omitempty"`    // 扩展信息list，自定义字段，最长不大于 500 字符，将extend_field_list转为 json 字符串，三个字段仅作为备注性字段。详见extend_field_list
	BudgetCenterId      *string                `json:"budget_center_id,omitempty"`     // 滴滴侧成本中心ID，滴滴侧成本中心ID；获取方式接口返回的 ID（可为项目/部门ID) eg : 1125920020961744 关联部门 时， id和out_budget_id 优先处理id关联项目时 id和 out_budget_id与 name 组合唯一值时 优先处理id ，out_budget_id与 name 同时有值时按照项目处理，out_budget_id有值且name为空时，按照部门处理
	TravelBudgetObj     *TravelBudget          `json:"travel_budget__obj__,omitempty"` // 差旅预算总额,可使用脚本将其转换为json后赋值给 travel_budget 字段
	TravelManagementObj *TravelManagement      `json:"travel_management__obj__,omitempty"`
	ExtraInfoObj        *map[string]string     `json:"extra_info__obj__,omitempty"` // 扩展信息,可使用脚本将其转换为json后赋值给 extra_info 字段
	TravelDetailObj     *TravelDetail          `json:"travel_detail__obj__,omitempty"`
	PassengerListObj    []TripPassenger        `json:"passenger_list__obj__,omitempty"`     // 出行人信息,可使用脚本将其转换为json后赋值给 passenger_list 字段
	BudgetCenterListObj []BudgetCenterListItem `json:"budget_center_list__obj__,omitempty"` // 差旅预算总额,可使用脚本将其转换为json后赋值给 budget_center_list 字段
	ExtendFieldListObj  *ExtendFieldList       `json:"extend_field_list__obj__,omitempty"`
}

type UpdateApprovalRequestBuilder struct {
	clientId               string // 申请应用时分配的AppKey
	clientIdSet            bool
	accessToken            string // 授权后的access token
	accessTokenSet         bool
	timestamp              int64 // 当前时间戳(精确到秒级)
	timestampSet           bool
	companyId              string // 滴滴公司 ID
	companyIdSet           bool
	sign                   string // 签名
	signSet                bool
	approvalType           int32 // 申批单类型，枚举值数字：1:差旅单；2: 行前审批-按次数；3: 行前审批-按日期
	approvalTypeSet        bool
	approvalId             string // 滴滴申请单ID。approval_id与out_approval_id至少需要传一个
	approvalIdSet          bool
	outApprovalId          string // 外部申请单ID，外部申请单ID approval_id优先级大于out_approval_id
	outApprovalIdSet       bool
	reason                 string // 申请原因，申请原因，限制 200 字符；默认为空字符 eg: 北京出差，行程维度事由。使用卡片事由时，可以为空
	reasonSet              bool
	name                   string // 部门/项目名称，不大于 200 字符
	nameSet                bool
	outBudgetId            string // 编号，长度限制：≤ 64 字符
	outBudgetIdSet         bool
	travelBudget           string // 差旅预算总额，差旅预算控制，一个申请单对应一个总预算纬度，一次只能全部或者选择部分品类。字段不传输，修改不处理，字段传输，空对象时处理为清空。给值为修改
	travelBudgetSet        bool
	travelManagement       string // 差旅管控，差旅管控，目前仅支持市内用车每日限额设置（为json字符串类型，具体看请求示例）
	travelManagementSet    bool
	extraInfo              string // 扩展信息，扩展信息，自定义字段；最长不大于 500 字符；(必须为json字符串)；默认为空字符
	extraInfoSet           bool
	travelDetail           string // 差旅单行程信息，差旅单行程信息，将travel_detail的值转为 json 字符串，approval_type = 1 时必传
	travelDetailSet        bool
	passengerList          string // 出行人信息，出行人信息，不传时默认出行人为申请人，将passenger_list 转为 json 数组字符串。详见passenger_list
	passengerListSet       bool
	budgetCenterList       string // 多成本中心
	budgetCenterListSet    bool
	extendFieldList        string // 扩展信息list，自定义字段，最长不大于 500 字符，将extend_field_list转为 json 字符串，三个字段仅作为备注性字段。详见extend_field_list
	extendFieldListSet     bool
	budgetCenterId         string // 滴滴侧成本中心ID，滴滴侧成本中心ID；获取方式接口返回的 ID（可为项目/部门ID) eg : 1125920020961744 关联部门 时， id和out_budget_id 优先处理id关联项目时 id和 out_budget_id与 name 组合唯一值时 优先处理id ，out_budget_id与 name 同时有值时按照项目处理，out_budget_id有值且name为空时，按照部门处理
	budgetCenterIdSet      bool
	travelBudgetObj        TravelBudget // 差旅预算总额,可使用脚本将其转换为json后赋值给 travel_budget 字段
	travelBudgetObjSet     bool
	travelManagementObj    TravelManagement
	travelManagementObjSet bool
	extraInfoObj           map[string]string // 扩展信息,可使用脚本将其转换为json后赋值给 extra_info 字段
	extraInfoObjSet        bool
	travelDetailObj        TravelDetail
	travelDetailObjSet     bool
	passengerListObj       []TripPassenger // 出行人信息,可使用脚本将其转换为json后赋值给 passenger_list 字段
	passengerListObjSet    bool
	budgetCenterListObj    []BudgetCenterListItem // 差旅预算总额,可使用脚本将其转换为json后赋值给 budget_center_list 字段
	budgetCenterListObjSet bool
	extendFieldListObj     ExtendFieldList
	extendFieldListObjSet  bool
}

func NewUpdateApprovalRequestBuilder() *UpdateApprovalRequestBuilder {
	return &UpdateApprovalRequestBuilder{}
}
func (builder *UpdateApprovalRequestBuilder) ClientId(clientId string) *UpdateApprovalRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) AccessToken(accessToken string) *UpdateApprovalRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) Timestamp(timestamp int64) *UpdateApprovalRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) CompanyId(companyId string) *UpdateApprovalRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) Sign(sign string) *UpdateApprovalRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ApprovalType(approvalType int32) *UpdateApprovalRequestBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ApprovalId(approvalId string) *UpdateApprovalRequestBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) OutApprovalId(outApprovalId string) *UpdateApprovalRequestBuilder {
	builder.outApprovalId = outApprovalId
	builder.outApprovalIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) Reason(reason string) *UpdateApprovalRequestBuilder {
	builder.reason = reason
	builder.reasonSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) Name(name string) *UpdateApprovalRequestBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) OutBudgetId(outBudgetId string) *UpdateApprovalRequestBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelBudget(travelBudget string) *UpdateApprovalRequestBuilder {
	builder.travelBudget = travelBudget
	builder.travelBudgetSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelManagement(travelManagement string) *UpdateApprovalRequestBuilder {
	builder.travelManagement = travelManagement
	builder.travelManagementSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ExtraInfo(extraInfo string) *UpdateApprovalRequestBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelDetail(travelDetail string) *UpdateApprovalRequestBuilder {
	builder.travelDetail = travelDetail
	builder.travelDetailSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) PassengerList(passengerList string) *UpdateApprovalRequestBuilder {
	builder.passengerList = passengerList
	builder.passengerListSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) BudgetCenterList(budgetCenterList string) *UpdateApprovalRequestBuilder {
	builder.budgetCenterList = budgetCenterList
	builder.budgetCenterListSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ExtendFieldList(extendFieldList string) *UpdateApprovalRequestBuilder {
	builder.extendFieldList = extendFieldList
	builder.extendFieldListSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) BudgetCenterId(budgetCenterId string) *UpdateApprovalRequestBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelBudgetObj(travelBudgetObj TravelBudget) *UpdateApprovalRequestBuilder {
	builder.travelBudgetObj = travelBudgetObj
	builder.travelBudgetObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelManagementObj(travelManagementObj TravelManagement) *UpdateApprovalRequestBuilder {
	builder.travelManagementObj = travelManagementObj
	builder.travelManagementObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ExtraInfoObj(extraInfoObj map[string]string) *UpdateApprovalRequestBuilder {
	builder.extraInfoObj = extraInfoObj
	builder.extraInfoObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) TravelDetailObj(travelDetailObj TravelDetail) *UpdateApprovalRequestBuilder {
	builder.travelDetailObj = travelDetailObj
	builder.travelDetailObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) PassengerListObj(passengerListObj []TripPassenger) *UpdateApprovalRequestBuilder {
	builder.passengerListObj = passengerListObj
	builder.passengerListObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) BudgetCenterListObj(budgetCenterListObj []BudgetCenterListItem) *UpdateApprovalRequestBuilder {
	builder.budgetCenterListObj = budgetCenterListObj
	builder.budgetCenterListObjSet = true
	return builder
}
func (builder *UpdateApprovalRequestBuilder) ExtendFieldListObj(extendFieldListObj ExtendFieldList) *UpdateApprovalRequestBuilder {
	builder.extendFieldListObj = extendFieldListObj
	builder.extendFieldListObjSet = true
	return builder
}

func (builder *UpdateApprovalRequestBuilder) Build() *UpdateApprovalRequest {
	data := &UpdateApprovalRequest{}
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
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.outApprovalIdSet {
		data.OutApprovalId = &builder.outApprovalId
	}
	if builder.reasonSet {
		data.Reason = &builder.reason
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	if builder.travelBudgetSet {
		data.TravelBudget = &builder.travelBudget
	}
	if builder.travelManagementSet {
		data.TravelManagement = &builder.travelManagement
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.travelDetailSet {
		data.TravelDetail = &builder.travelDetail
	}
	if builder.passengerListSet {
		data.PassengerList = &builder.passengerList
	}
	if builder.budgetCenterListSet {
		data.BudgetCenterList = &builder.budgetCenterList
	}
	if builder.extendFieldListSet {
		data.ExtendFieldList = &builder.extendFieldList
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.travelBudgetObjSet {
		data.TravelBudgetObj = &builder.travelBudgetObj
	}
	if builder.travelManagementObjSet {
		data.TravelManagementObj = &builder.travelManagementObj
	}
	if builder.extraInfoObjSet {
		data.ExtraInfoObj = &builder.extraInfoObj
	}
	if builder.travelDetailObjSet {
		data.TravelDetailObj = &builder.travelDetailObj
	}
	if builder.passengerListObjSet {
		data.PassengerListObj = builder.passengerListObj
	}
	if builder.budgetCenterListObjSet {
		data.BudgetCenterListObj = builder.budgetCenterListObj
	}
	if builder.extendFieldListObjSet {
		data.ExtendFieldListObj = &builder.extendFieldListObj
	}
	return data
}
