package v1

// AdjustFields struct for AdjustFields
type AdjustFields struct {
	BookingDepName           *string `json:"bookingDepName,omitempty"`           // 部门名称相关（对于用车是下单人部门名称，对于商旅是预订人部门名称）
	BookingDepCode           *string `json:"bookingDepCode,omitempty"`           // 部门代码相关（对于用车是下单人部门代码，对于商旅是预订人部门代码）
	BudgetCenterName         *string `json:"budgetCenterName,omitempty"`         // 成本中心名称
	BudgetCenterCode         *string `json:"budgetCenterCode,omitempty"`         // 成本中心代码
	LegalEntityName          *string `json:"legalEntityName,omitempty"`          // 所属公司名称相关（对于用车是下单人所属公司名称，对于商旅是预订人所属公司名称）
	LegalEntityCode          *string `json:"legalEntityCode,omitempty"`          // 所属公司代码相关（对于用车是下单人所属公司代码，对于商旅是预订人所属公司代码）
	ProjectExtInfo           *string `json:"projectExtInfo,omitempty"`           // 项目自定义信息
	ExInfo01                 *string `json:"exInfo01,omitempty"`                 // 自定义扩展信息 01
	ExInfo02                 *string `json:"exInfo02,omitempty"`                 // 自定义扩展信息 02
	ExInfo03                 *string `json:"exInfo03,omitempty"`                 // 自定义扩展信息 03
	ExInfo04                 *string `json:"exInfo04,omitempty"`                 // 自定义扩展信息 04
	ExInfo05                 *string `json:"exInfo05,omitempty"`                 // 自定义扩展信息 05
	ExInfo06                 *string `json:"exInfo06,omitempty"`                 // 自定义扩展信息 06
	ExInfo07                 *string `json:"exInfo07,omitempty"`                 // 自定义扩展信息 07
	ExInfo08                 *string `json:"exInfo08,omitempty"`                 // 自定义扩展信息 08
	BookingDepId             *string `json:"bookingDepId,omitempty"`             // 下单人部门 不直接传递,请bookingDepName+bookingDepCode匹配;name支持传递或>进行路径连接,code请直接传递;会尝试 或者+并且+有效 进行逻辑得到唯一匹配项,若不是则报错
	BudgetCenterId           *string `json:"budgetCenterId,omitempty"`           // 成本中心名称 不直接传递,用名称+code匹配,逻辑同bookingDepId
	LegalEntityId            *string `json:"legalEntityId,omitempty"`            // 下单人所属公司 不直接传递,用名称+code匹配,逻辑同bookingDepId
	ExtendInfo               *string `json:"extendInfo,omitempty"`               // 自定义(开票主体等信息)
	ApprovalInfo             *string `json:"approvalInfo,omitempty"`             // 审批自定义
	InstitutionName          *string `json:"institutionName,omitempty"`          // 用车制度 会check是否存在
	RuleName                 *string `json:"ruleName,omitempty"`                 // 用车权限 会check是否存在
	PassengerDepId           *string `json:"passengerDepId,omitempty"`           // 入住人部门 不直接传递,用名称+code匹配,逻辑同网约车bookingDepId
	PassengerLegalEntityId   *string `json:"passengerLegalEntityId,omitempty"`   // 入住人所属公司名称 不直接传递,用名称+code匹配,逻辑同网约车bookingDepId
	ExtraInfo                *string `json:"extraInfo,omitempty"`                // 审批自定义
	PassengerLegalEntityCode *string `json:"passengerLegalEntityCode,omitempty"` // 入住人所属公司代码 [商旅字段]
	PassengerDepName         *string `json:"passengerDepName,omitempty"`         // 入住人部门名称 [商旅字段]
	PassengerDepCode         *string `json:"passengerDepCode,omitempty"`         // 入住人部门代码 [商旅字段]
	PassengerLegalEntityName *string `json:"passengerLegalEntityName,omitempty"` // 入住人所属公司名称 [商旅字段]
}

type AdjustFieldsBuilder struct {
	bookingDepName              string // 部门名称相关（对于用车是下单人部门名称，对于商旅是预订人部门名称）
	bookingDepNameSet           bool
	bookingDepCode              string // 部门代码相关（对于用车是下单人部门代码，对于商旅是预订人部门代码）
	bookingDepCodeSet           bool
	budgetCenterName            string // 成本中心名称
	budgetCenterNameSet         bool
	budgetCenterCode            string // 成本中心代码
	budgetCenterCodeSet         bool
	legalEntityName             string // 所属公司名称相关（对于用车是下单人所属公司名称，对于商旅是预订人所属公司名称）
	legalEntityNameSet          bool
	legalEntityCode             string // 所属公司代码相关（对于用车是下单人所属公司代码，对于商旅是预订人所属公司代码）
	legalEntityCodeSet          bool
	projectExtInfo              string // 项目自定义信息
	projectExtInfoSet           bool
	exInfo01                    string // 自定义扩展信息 01
	exInfo01Set                 bool
	exInfo02                    string // 自定义扩展信息 02
	exInfo02Set                 bool
	exInfo03                    string // 自定义扩展信息 03
	exInfo03Set                 bool
	exInfo04                    string // 自定义扩展信息 04
	exInfo04Set                 bool
	exInfo05                    string // 自定义扩展信息 05
	exInfo05Set                 bool
	exInfo06                    string // 自定义扩展信息 06
	exInfo06Set                 bool
	exInfo07                    string // 自定义扩展信息 07
	exInfo07Set                 bool
	exInfo08                    string // 自定义扩展信息 08
	exInfo08Set                 bool
	bookingDepId                string // 下单人部门 不直接传递,请bookingDepName+bookingDepCode匹配;name支持传递或>进行路径连接,code请直接传递;会尝试 或者+并且+有效 进行逻辑得到唯一匹配项,若不是则报错
	bookingDepIdSet             bool
	budgetCenterId              string // 成本中心名称 不直接传递,用名称+code匹配,逻辑同bookingDepId
	budgetCenterIdSet           bool
	legalEntityId               string // 下单人所属公司 不直接传递,用名称+code匹配,逻辑同bookingDepId
	legalEntityIdSet            bool
	extendInfo                  string // 自定义(开票主体等信息)
	extendInfoSet               bool
	approvalInfo                string // 审批自定义
	approvalInfoSet             bool
	institutionName             string // 用车制度 会check是否存在
	institutionNameSet          bool
	ruleName                    string // 用车权限 会check是否存在
	ruleNameSet                 bool
	passengerDepId              string // 入住人部门 不直接传递,用名称+code匹配,逻辑同网约车bookingDepId
	passengerDepIdSet           bool
	passengerLegalEntityId      string // 入住人所属公司名称 不直接传递,用名称+code匹配,逻辑同网约车bookingDepId
	passengerLegalEntityIdSet   bool
	extraInfo                   string // 审批自定义
	extraInfoSet                bool
	passengerLegalEntityCode    string // 入住人所属公司代码 [商旅字段]
	passengerLegalEntityCodeSet bool
	passengerDepName            string // 入住人部门名称 [商旅字段]
	passengerDepNameSet         bool
	passengerDepCode            string // 入住人部门代码 [商旅字段]
	passengerDepCodeSet         bool
	passengerLegalEntityName    string // 入住人所属公司名称 [商旅字段]
	passengerLegalEntityNameSet bool
}

func NewAdjustFieldsBuilder() *AdjustFieldsBuilder {
	return &AdjustFieldsBuilder{}
}
func (builder *AdjustFieldsBuilder) BookingDepName(bookingDepName string) *AdjustFieldsBuilder {
	builder.bookingDepName = bookingDepName
	builder.bookingDepNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) BookingDepCode(bookingDepCode string) *AdjustFieldsBuilder {
	builder.bookingDepCode = bookingDepCode
	builder.bookingDepCodeSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) BudgetCenterName(budgetCenterName string) *AdjustFieldsBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) BudgetCenterCode(budgetCenterCode string) *AdjustFieldsBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) LegalEntityName(legalEntityName string) *AdjustFieldsBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) LegalEntityCode(legalEntityCode string) *AdjustFieldsBuilder {
	builder.legalEntityCode = legalEntityCode
	builder.legalEntityCodeSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) ProjectExtInfo(projectExtInfo string) *AdjustFieldsBuilder {
	builder.projectExtInfo = projectExtInfo
	builder.projectExtInfoSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo01(exInfo01 string) *AdjustFieldsBuilder {
	builder.exInfo01 = exInfo01
	builder.exInfo01Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo02(exInfo02 string) *AdjustFieldsBuilder {
	builder.exInfo02 = exInfo02
	builder.exInfo02Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo03(exInfo03 string) *AdjustFieldsBuilder {
	builder.exInfo03 = exInfo03
	builder.exInfo03Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo04(exInfo04 string) *AdjustFieldsBuilder {
	builder.exInfo04 = exInfo04
	builder.exInfo04Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo05(exInfo05 string) *AdjustFieldsBuilder {
	builder.exInfo05 = exInfo05
	builder.exInfo05Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo06(exInfo06 string) *AdjustFieldsBuilder {
	builder.exInfo06 = exInfo06
	builder.exInfo06Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo07(exInfo07 string) *AdjustFieldsBuilder {
	builder.exInfo07 = exInfo07
	builder.exInfo07Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExInfo08(exInfo08 string) *AdjustFieldsBuilder {
	builder.exInfo08 = exInfo08
	builder.exInfo08Set = true
	return builder
}
func (builder *AdjustFieldsBuilder) BookingDepId(bookingDepId string) *AdjustFieldsBuilder {
	builder.bookingDepId = bookingDepId
	builder.bookingDepIdSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) BudgetCenterId(budgetCenterId string) *AdjustFieldsBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) LegalEntityId(legalEntityId string) *AdjustFieldsBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExtendInfo(extendInfo string) *AdjustFieldsBuilder {
	builder.extendInfo = extendInfo
	builder.extendInfoSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) ApprovalInfo(approvalInfo string) *AdjustFieldsBuilder {
	builder.approvalInfo = approvalInfo
	builder.approvalInfoSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) InstitutionName(institutionName string) *AdjustFieldsBuilder {
	builder.institutionName = institutionName
	builder.institutionNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) RuleName(ruleName string) *AdjustFieldsBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerDepId(passengerDepId string) *AdjustFieldsBuilder {
	builder.passengerDepId = passengerDepId
	builder.passengerDepIdSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerLegalEntityId(passengerLegalEntityId string) *AdjustFieldsBuilder {
	builder.passengerLegalEntityId = passengerLegalEntityId
	builder.passengerLegalEntityIdSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) ExtraInfo(extraInfo string) *AdjustFieldsBuilder {
	builder.extraInfo = extraInfo
	builder.extraInfoSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerLegalEntityCode(passengerLegalEntityCode string) *AdjustFieldsBuilder {
	builder.passengerLegalEntityCode = passengerLegalEntityCode
	builder.passengerLegalEntityCodeSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerDepName(passengerDepName string) *AdjustFieldsBuilder {
	builder.passengerDepName = passengerDepName
	builder.passengerDepNameSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerDepCode(passengerDepCode string) *AdjustFieldsBuilder {
	builder.passengerDepCode = passengerDepCode
	builder.passengerDepCodeSet = true
	return builder
}
func (builder *AdjustFieldsBuilder) PassengerLegalEntityName(passengerLegalEntityName string) *AdjustFieldsBuilder {
	builder.passengerLegalEntityName = passengerLegalEntityName
	builder.passengerLegalEntityNameSet = true
	return builder
}

func (builder *AdjustFieldsBuilder) Build() *AdjustFields {
	data := &AdjustFields{}
	if builder.bookingDepNameSet {
		data.BookingDepName = &builder.bookingDepName
	}
	if builder.bookingDepCodeSet {
		data.BookingDepCode = &builder.bookingDepCode
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	if builder.legalEntityCodeSet {
		data.LegalEntityCode = &builder.legalEntityCode
	}
	if builder.projectExtInfoSet {
		data.ProjectExtInfo = &builder.projectExtInfo
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
	if builder.bookingDepIdSet {
		data.BookingDepId = &builder.bookingDepId
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.extendInfoSet {
		data.ExtendInfo = &builder.extendInfo
	}
	if builder.approvalInfoSet {
		data.ApprovalInfo = &builder.approvalInfo
	}
	if builder.institutionNameSet {
		data.InstitutionName = &builder.institutionName
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.passengerDepIdSet {
		data.PassengerDepId = &builder.passengerDepId
	}
	if builder.passengerLegalEntityIdSet {
		data.PassengerLegalEntityId = &builder.passengerLegalEntityId
	}
	if builder.extraInfoSet {
		data.ExtraInfo = &builder.extraInfo
	}
	if builder.passengerLegalEntityCodeSet {
		data.PassengerLegalEntityCode = &builder.passengerLegalEntityCode
	}
	if builder.passengerDepNameSet {
		data.PassengerDepName = &builder.passengerDepName
	}
	if builder.passengerDepCodeSet {
		data.PassengerDepCode = &builder.passengerDepCode
	}
	if builder.passengerLegalEntityNameSet {
		data.PassengerLegalEntityName = &builder.passengerLegalEntityName
	}
	return data
}
