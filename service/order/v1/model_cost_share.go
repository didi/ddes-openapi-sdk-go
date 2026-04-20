package v1

// CostShare 个人分摊对象
type CostShare struct {
	ShareRatio         *float64 `json:"share_ratio,omitempty"`          // 分配比例
	ShareCompanyRatio  *float64 `json:"share_company_ratio,omitempty"`  // 公司承担比例
	SharePersonalRatio *float64 `json:"share_personal_ratio,omitempty"` // 个人承担比例
	CompanyPay         *string  `json:"company_pay,omitempty"`          // 企业应付总额
	PersonalPay        *string  `json:"personal_pay,omitempty"`         // 个人应付总额
	CompanyRealPay     *string  `json:"company_real_pay,omitempty"`     // 公司实付金额
	PersonalRealPay    *string  `json:"personal_real_pay,omitempty"`    // 个人实付金额
	CompanyRealRefund  *string  `json:"company_real_refund,omitempty"`  // 公司实际退款金额
	PersonalRealRefund *string  `json:"personal_real_refund,omitempty"` // 个人实际退款金额
}

type CostShareBuilder struct {
	shareRatio            float64 // 分配比例
	shareRatioSet         bool
	shareCompanyRatio     float64 // 公司承担比例
	shareCompanyRatioSet  bool
	sharePersonalRatio    float64 // 个人承担比例
	sharePersonalRatioSet bool
	companyPay            string // 企业应付总额
	companyPaySet         bool
	personalPay           string // 个人应付总额
	personalPaySet        bool
	companyRealPay        string // 公司实付金额
	companyRealPaySet     bool
	personalRealPay       string // 个人实付金额
	personalRealPaySet    bool
	companyRealRefund     string // 公司实际退款金额
	companyRealRefundSet  bool
	personalRealRefund    string // 个人实际退款金额
	personalRealRefundSet bool
}

func NewCostShareBuilder() *CostShareBuilder {
	return &CostShareBuilder{}
}
func (builder *CostShareBuilder) ShareRatio(shareRatio float64) *CostShareBuilder {
	builder.shareRatio = shareRatio
	builder.shareRatioSet = true
	return builder
}
func (builder *CostShareBuilder) ShareCompanyRatio(shareCompanyRatio float64) *CostShareBuilder {
	builder.shareCompanyRatio = shareCompanyRatio
	builder.shareCompanyRatioSet = true
	return builder
}
func (builder *CostShareBuilder) SharePersonalRatio(sharePersonalRatio float64) *CostShareBuilder {
	builder.sharePersonalRatio = sharePersonalRatio
	builder.sharePersonalRatioSet = true
	return builder
}
func (builder *CostShareBuilder) CompanyPay(companyPay string) *CostShareBuilder {
	builder.companyPay = companyPay
	builder.companyPaySet = true
	return builder
}
func (builder *CostShareBuilder) PersonalPay(personalPay string) *CostShareBuilder {
	builder.personalPay = personalPay
	builder.personalPaySet = true
	return builder
}
func (builder *CostShareBuilder) CompanyRealPay(companyRealPay string) *CostShareBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *CostShareBuilder) PersonalRealPay(personalRealPay string) *CostShareBuilder {
	builder.personalRealPay = personalRealPay
	builder.personalRealPaySet = true
	return builder
}
func (builder *CostShareBuilder) CompanyRealRefund(companyRealRefund string) *CostShareBuilder {
	builder.companyRealRefund = companyRealRefund
	builder.companyRealRefundSet = true
	return builder
}
func (builder *CostShareBuilder) PersonalRealRefund(personalRealRefund string) *CostShareBuilder {
	builder.personalRealRefund = personalRealRefund
	builder.personalRealRefundSet = true
	return builder
}

func (builder *CostShareBuilder) Build() *CostShare {
	data := &CostShare{}
	if builder.shareRatioSet {
		data.ShareRatio = &builder.shareRatio
	}
	if builder.shareCompanyRatioSet {
		data.ShareCompanyRatio = &builder.shareCompanyRatio
	}
	if builder.sharePersonalRatioSet {
		data.SharePersonalRatio = &builder.sharePersonalRatio
	}
	if builder.companyPaySet {
		data.CompanyPay = &builder.companyPay
	}
	if builder.personalPaySet {
		data.PersonalPay = &builder.personalPay
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.personalRealPaySet {
		data.PersonalRealPay = &builder.personalRealPay
	}
	if builder.companyRealRefundSet {
		data.CompanyRealRefund = &builder.companyRealRefund
	}
	if builder.personalRealRefundSet {
		data.PersonalRealRefund = &builder.personalRealRefund
	}
	return data
}
