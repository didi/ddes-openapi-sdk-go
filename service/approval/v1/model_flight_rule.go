package v1

// FlightRule 航班规则信息
type FlightRule struct {
	RuleId         *string    `json:"rule_id,omitempty"`
	RuleName       *string    `json:"rule_name,omitempty"`
	RuleStatus     *string    `json:"rule_status,omitempty"`
	TotalCount     *int32     `json:"total_count,omitempty"`
	AvailableCount *int32     `json:"available_count,omitempty"`
	StartTime      *string    `json:"start_time,omitempty"`
	EndTime        *string    `json:"end_time,omitempty"`
	CityLine       []CityLine `json:"city_line,omitempty"`
}

type FlightRuleBuilder struct {
	ruleId            string
	ruleIdSet         bool
	ruleName          string
	ruleNameSet       bool
	ruleStatus        string
	ruleStatusSet     bool
	totalCount        int32
	totalCountSet     bool
	availableCount    int32
	availableCountSet bool
	startTime         string
	startTimeSet      bool
	endTime           string
	endTimeSet        bool
	cityLine          []CityLine
	cityLineSet       bool
}

func NewFlightRuleBuilder() *FlightRuleBuilder {
	return &FlightRuleBuilder{}
}
func (builder *FlightRuleBuilder) RuleId(ruleId string) *FlightRuleBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *FlightRuleBuilder) RuleName(ruleName string) *FlightRuleBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *FlightRuleBuilder) RuleStatus(ruleStatus string) *FlightRuleBuilder {
	builder.ruleStatus = ruleStatus
	builder.ruleStatusSet = true
	return builder
}
func (builder *FlightRuleBuilder) TotalCount(totalCount int32) *FlightRuleBuilder {
	builder.totalCount = totalCount
	builder.totalCountSet = true
	return builder
}
func (builder *FlightRuleBuilder) AvailableCount(availableCount int32) *FlightRuleBuilder {
	builder.availableCount = availableCount
	builder.availableCountSet = true
	return builder
}
func (builder *FlightRuleBuilder) StartTime(startTime string) *FlightRuleBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *FlightRuleBuilder) EndTime(endTime string) *FlightRuleBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *FlightRuleBuilder) CityLine(cityLine []CityLine) *FlightRuleBuilder {
	builder.cityLine = cityLine
	builder.cityLineSet = true
	return builder
}

func (builder *FlightRuleBuilder) Build() *FlightRule {
	data := &FlightRule{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.ruleStatusSet {
		data.RuleStatus = &builder.ruleStatus
	}
	if builder.totalCountSet {
		data.TotalCount = &builder.totalCount
	}
	if builder.availableCountSet {
		data.AvailableCount = &builder.availableCount
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.cityLineSet {
		data.CityLine = builder.cityLine
	}
	return data
}
