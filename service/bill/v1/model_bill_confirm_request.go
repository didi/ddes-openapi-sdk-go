package v1

// BillConfirmRequest struct for BillConfirmRequest
type BillConfirmRequest struct {
	ClientId           *string `json:"client_id,omitempty"`             // 申请应用时分配的AppKey
	AccessToken        *string `json:"access_token,omitempty"`          // 授权后的access token
	CompanyId          *string `json:"company_id,omitempty"`            // 企业ID
	Timestamp          *int64  `json:"timestamp,omitempty"`             // 当前时间戳，精确到秒级
	BillId             *string `json:"bill_id,omitempty"`               // 账单id
	BusinessType       *int32  `json:"business_type,omitempty"`         // 业务线（0：网约车；1：商旅）默认是0，不传默认网约车
	PaymentPeriod      *string `json:"payment_period,omitempty"`        // 账单周期(当不传bill_id时，payment_period和department_id必须填写 或 payment_period和budget_center_id必须填写 或 payment_period和bill_split_type和bill_split_group_key必须填写) 示例：\"2020-01-01~2020-01-31\"
	DepartmentId       *string `json:"department_id,omitempty"`         // 部门id(当按账单周期 + 部门查询时，department_id必须填写) 仅支持网约车
	BudgetCenterId     *string `json:"budget_center_id,omitempty"`      // 成本中心id(当按账单周期 + 成本中心查询时，budget_center_id必须填写) 仅支持网约车
	BillSplitType      *int32  `json:"bill_split_type,omitempty"`       // 拆分账单维度（仅支持 国内机票、国内酒店、火车票、国际机票、国际酒店、增值手工单）1=按出行人部门拆分，2=按预订人部门拆分，3=按成本中心拆分，4=按出行人所属公司拆分，5=按预订人所属公司拆分，6=按自定义字段拆分 注：bill_split_type和bill_split_group_key必须结合使用，且当传bill_id时按照拆分维度查不生效7=多维拆帐 注：bill_split_type、bill_split_group_type和bill_split_group_key这3个字段必须结合使用，且当传bill_id时按照拆分维度查不生效
	BillSplitGroupType *int32  `json:"bill_split_group_type,omitempty"` // 多维拆帐节点类型（仅支持 国内机票、国内酒店、火车票、国际机票、国际酒店、增值手工单），1=按部门拆分，2=按项目拆分，4=按公司拆分，6=按自定义字段、用车制度、用车权限拆分，8= 未知主体
	BillSplitGroupKey  *string `json:"bill_split_group_key,omitempty"`  // 拆分账单主体id或名称（仅支持 机票、酒店、火车票、国际机票、国际酒店、增值手工单） 1.单维拆账时，bill_split_type为1~5时，拆分账单主体id或名称请填写id；bill_split_type为6时，请填写自定义字段名称 注：bill_split_type和bill_split_group_key必须结合使用，且当传bill_id时按照拆分维度查不生效； 2.多维拆帐时，bill_split_group_type为1、2、4时，填写对应的id；bill_split_group_type为6时，值为滴滴生成的帐单主体名称；bill_split_group_type为8时，本字段为空。
	Sign               *string `json:"sign,omitempty"`                  // 签名
}

type BillConfirmRequestBuilder struct {
	clientId              string // 申请应用时分配的AppKey
	clientIdSet           bool
	accessToken           string // 授权后的access token
	accessTokenSet        bool
	companyId             string // 企业ID
	companyIdSet          bool
	timestamp             int64 // 当前时间戳，精确到秒级
	timestampSet          bool
	billId                string // 账单id
	billIdSet             bool
	businessType          int32 // 业务线（0：网约车；1：商旅）默认是0，不传默认网约车
	businessTypeSet       bool
	paymentPeriod         string // 账单周期(当不传bill_id时，payment_period和department_id必须填写 或 payment_period和budget_center_id必须填写 或 payment_period和bill_split_type和bill_split_group_key必须填写) 示例：\"2020-01-01~2020-01-31\"
	paymentPeriodSet      bool
	departmentId          string // 部门id(当按账单周期 + 部门查询时，department_id必须填写) 仅支持网约车
	departmentIdSet       bool
	budgetCenterId        string // 成本中心id(当按账单周期 + 成本中心查询时，budget_center_id必须填写) 仅支持网约车
	budgetCenterIdSet     bool
	billSplitType         int32 // 拆分账单维度（仅支持 国内机票、国内酒店、火车票、国际机票、国际酒店、增值手工单）1=按出行人部门拆分，2=按预订人部门拆分，3=按成本中心拆分，4=按出行人所属公司拆分，5=按预订人所属公司拆分，6=按自定义字段拆分 注：bill_split_type和bill_split_group_key必须结合使用，且当传bill_id时按照拆分维度查不生效7=多维拆帐 注：bill_split_type、bill_split_group_type和bill_split_group_key这3个字段必须结合使用，且当传bill_id时按照拆分维度查不生效
	billSplitTypeSet      bool
	billSplitGroupType    int32 // 多维拆帐节点类型（仅支持 国内机票、国内酒店、火车票、国际机票、国际酒店、增值手工单），1=按部门拆分，2=按项目拆分，4=按公司拆分，6=按自定义字段、用车制度、用车权限拆分，8= 未知主体
	billSplitGroupTypeSet bool
	billSplitGroupKey     string // 拆分账单主体id或名称（仅支持 机票、酒店、火车票、国际机票、国际酒店、增值手工单） 1.单维拆账时，bill_split_type为1~5时，拆分账单主体id或名称请填写id；bill_split_type为6时，请填写自定义字段名称 注：bill_split_type和bill_split_group_key必须结合使用，且当传bill_id时按照拆分维度查不生效； 2.多维拆帐时，bill_split_group_type为1、2、4时，填写对应的id；bill_split_group_type为6时，值为滴滴生成的帐单主体名称；bill_split_group_type为8时，本字段为空。
	billSplitGroupKeySet  bool
	sign                  string // 签名
	signSet               bool
}

func NewBillConfirmRequestBuilder() *BillConfirmRequestBuilder {
	return &BillConfirmRequestBuilder{}
}
func (builder *BillConfirmRequestBuilder) ClientId(clientId string) *BillConfirmRequestBuilder {
	builder.clientId = clientId
	builder.clientIdSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) AccessToken(accessToken string) *BillConfirmRequestBuilder {
	builder.accessToken = accessToken
	builder.accessTokenSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) CompanyId(companyId string) *BillConfirmRequestBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) Timestamp(timestamp int64) *BillConfirmRequestBuilder {
	builder.timestamp = timestamp
	builder.timestampSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BillId(billId string) *BillConfirmRequestBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BusinessType(businessType int32) *BillConfirmRequestBuilder {
	builder.businessType = businessType
	builder.businessTypeSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) PaymentPeriod(paymentPeriod string) *BillConfirmRequestBuilder {
	builder.paymentPeriod = paymentPeriod
	builder.paymentPeriodSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) DepartmentId(departmentId string) *BillConfirmRequestBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BudgetCenterId(budgetCenterId string) *BillConfirmRequestBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BillSplitType(billSplitType int32) *BillConfirmRequestBuilder {
	builder.billSplitType = billSplitType
	builder.billSplitTypeSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BillSplitGroupType(billSplitGroupType int32) *BillConfirmRequestBuilder {
	builder.billSplitGroupType = billSplitGroupType
	builder.billSplitGroupTypeSet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) BillSplitGroupKey(billSplitGroupKey string) *BillConfirmRequestBuilder {
	builder.billSplitGroupKey = billSplitGroupKey
	builder.billSplitGroupKeySet = true
	return builder
}
func (builder *BillConfirmRequestBuilder) Sign(sign string) *BillConfirmRequestBuilder {
	builder.sign = sign
	builder.signSet = true
	return builder
}

func (builder *BillConfirmRequestBuilder) Build() *BillConfirmRequest {
	data := &BillConfirmRequest{}
	if builder.clientIdSet {
		data.ClientId = &builder.clientId
	}
	if builder.accessTokenSet {
		data.AccessToken = &builder.accessToken
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.timestampSet {
		data.Timestamp = &builder.timestamp
	}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.businessTypeSet {
		data.BusinessType = &builder.businessType
	}
	if builder.paymentPeriodSet {
		data.PaymentPeriod = &builder.paymentPeriod
	}
	if builder.departmentIdSet {
		data.DepartmentId = &builder.departmentId
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.billSplitTypeSet {
		data.BillSplitType = &builder.billSplitType
	}
	if builder.billSplitGroupTypeSet {
		data.BillSplitGroupType = &builder.billSplitGroupType
	}
	if builder.billSplitGroupKeySet {
		data.BillSplitGroupKey = &builder.billSplitGroupKey
	}
	if builder.signSet {
		data.Sign = &builder.sign
	}
	return data
}
