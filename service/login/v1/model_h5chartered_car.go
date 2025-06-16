package v1

// H5CharteredCar 包车页面
type H5CharteredCar struct {
	RuleId     *string `json:"rule_id,omitempty"`     // 制度规则 id
	JumpPage   *string `json:"jumpPage,omitempty"`    // 跳转地址为: homeAddress
	ApprovalId *string `json:"approval_id,omitempty"` // 差旅、申请用车，必传
}

type H5CharteredCarBuilder struct {
	ruleId        string // 制度规则 id
	ruleIdSet     bool
	jumpPage      string // 跳转地址为: homeAddress
	jumpPageSet   bool
	approvalId    string // 差旅、申请用车，必传
	approvalIdSet bool
}

func NewH5CharteredCarBuilder() *H5CharteredCarBuilder {
	return &H5CharteredCarBuilder{}
}
func (builder *H5CharteredCarBuilder) RuleId(ruleId string) *H5CharteredCarBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *H5CharteredCarBuilder) JumpPage(jumpPage string) *H5CharteredCarBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}
func (builder *H5CharteredCarBuilder) ApprovalId(approvalId string) *H5CharteredCarBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}

func (builder *H5CharteredCarBuilder) Build() *H5CharteredCar {
	data := &H5CharteredCar{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	return data
}
