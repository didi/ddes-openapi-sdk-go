package v1

// H5RuleList 制度中心
type H5RuleList struct {
	JumpPage *string `json:"jumpPage,omitempty"` // 跳转地址为ruleList
}

type H5RuleListBuilder struct {
	jumpPage    string // 跳转地址为ruleList
	jumpPageSet bool
}

func NewH5RuleListBuilder() *H5RuleListBuilder {
	return &H5RuleListBuilder{}
}
func (builder *H5RuleListBuilder) JumpPage(jumpPage string) *H5RuleListBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5RuleListBuilder) Build() *H5RuleList {
	data := &H5RuleList{}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
