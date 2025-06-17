package v1

// TrainRule 列车规则信息
type TrainRule struct {
	RuleId         *string    `json:"rule_id,omitempty"`
	RuleName       *string    `json:"rule_name,omitempty"`
	RuleStatus     *string    `json:"rule_status,omitempty"`
	TotalCount     *int32     `json:"total_count,omitempty"`
	AvailableCount *int32     `json:"available_count,omitempty"`
	StartTime      *string    `json:"start_time,omitempty"`
	EndTime        *string    `json:"end_time,omitempty"`
	CityLine       []CityLine `json:"city_line,omitempty"`
}

type TrainRuleBuilder struct {
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

func NewTrainRuleBuilder() *TrainRuleBuilder {
	return &TrainRuleBuilder{}
}
func (builder *TrainRuleBuilder) RuleId(ruleId string) *TrainRuleBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *TrainRuleBuilder) RuleName(ruleName string) *TrainRuleBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *TrainRuleBuilder) RuleStatus(ruleStatus string) *TrainRuleBuilder {
	builder.ruleStatus = ruleStatus
	builder.ruleStatusSet = true
	return builder
}
func (builder *TrainRuleBuilder) TotalCount(totalCount int32) *TrainRuleBuilder {
	builder.totalCount = totalCount
	builder.totalCountSet = true
	return builder
}
func (builder *TrainRuleBuilder) AvailableCount(availableCount int32) *TrainRuleBuilder {
	builder.availableCount = availableCount
	builder.availableCountSet = true
	return builder
}
func (builder *TrainRuleBuilder) StartTime(startTime string) *TrainRuleBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *TrainRuleBuilder) EndTime(endTime string) *TrainRuleBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *TrainRuleBuilder) CityLine(cityLine []CityLine) *TrainRuleBuilder {
	builder.cityLine = cityLine
	builder.cityLineSet = true
	return builder
}

func (builder *TrainRuleBuilder) Build() *TrainRule {
	data := &TrainRule{}
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
