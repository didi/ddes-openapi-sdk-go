package v1

// H5RuleDesc 制度详情
type H5RuleDesc struct {
	InstitutionId *string `json:"institution_id,omitempty"` // 制度id
	JumpPage      *string `json:"jumpPage,omitempty"`       // 跳转地址为ruleDesc
}

type H5RuleDescBuilder struct {
	institutionId    string // 制度id
	institutionIdSet bool
	jumpPage         string // 跳转地址为ruleDesc
	jumpPageSet      bool
}

func NewH5RuleDescBuilder() *H5RuleDescBuilder {
	return &H5RuleDescBuilder{}
}
func (builder *H5RuleDescBuilder) InstitutionId(institutionId string) *H5RuleDescBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *H5RuleDescBuilder) JumpPage(jumpPage string) *H5RuleDescBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5RuleDescBuilder) Build() *H5RuleDesc {
	data := &H5RuleDesc{}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
