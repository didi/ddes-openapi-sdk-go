package v1

// HotelRule 酒店规则信息
type HotelRule struct {
	RuleId         *string      `json:"rule_id,omitempty"`
	RuleName       *string      `json:"rule_name,omitempty"`
	RuleStatus     *string      `json:"rule_status,omitempty"`
	CityList       []TravelCity `json:"city_list,omitempty"`
	TotalCount     *int32       `json:"total_count,omitempty"`
	AvailableCount *int32       `json:"available_count,omitempty"`
	StartTime      *string      `json:"start_time,omitempty"`
	EndTime        *string      `json:"end_time,omitempty"`
}

type HotelRuleBuilder struct {
	ruleId            string
	ruleIdSet         bool
	ruleName          string
	ruleNameSet       bool
	ruleStatus        string
	ruleStatusSet     bool
	cityList          []TravelCity
	cityListSet       bool
	totalCount        int32
	totalCountSet     bool
	availableCount    int32
	availableCountSet bool
	startTime         string
	startTimeSet      bool
	endTime           string
	endTimeSet        bool
}

func NewHotelRuleBuilder() *HotelRuleBuilder {
	return &HotelRuleBuilder{}
}
func (builder *HotelRuleBuilder) RuleId(ruleId string) *HotelRuleBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *HotelRuleBuilder) RuleName(ruleName string) *HotelRuleBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *HotelRuleBuilder) RuleStatus(ruleStatus string) *HotelRuleBuilder {
	builder.ruleStatus = ruleStatus
	builder.ruleStatusSet = true
	return builder
}
func (builder *HotelRuleBuilder) CityList(cityList []TravelCity) *HotelRuleBuilder {
	builder.cityList = cityList
	builder.cityListSet = true
	return builder
}
func (builder *HotelRuleBuilder) TotalCount(totalCount int32) *HotelRuleBuilder {
	builder.totalCount = totalCount
	builder.totalCountSet = true
	return builder
}
func (builder *HotelRuleBuilder) AvailableCount(availableCount int32) *HotelRuleBuilder {
	builder.availableCount = availableCount
	builder.availableCountSet = true
	return builder
}
func (builder *HotelRuleBuilder) StartTime(startTime string) *HotelRuleBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *HotelRuleBuilder) EndTime(endTime string) *HotelRuleBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}

func (builder *HotelRuleBuilder) Build() *HotelRule {
	data := &HotelRule{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.ruleStatusSet {
		data.RuleStatus = &builder.ruleStatus
	}
	if builder.cityListSet {
		data.CityList = builder.cityList
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
	return data
}
