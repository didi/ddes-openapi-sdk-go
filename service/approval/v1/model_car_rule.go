package v1

// CarRule 用车规则信息
type CarRule struct {
	RuleId         *string  `json:"rule_id,omitempty"`
	RuleName       *string  `json:"rule_name,omitempty"`
	RuleStatus     *int32   `json:"rule_status,omitempty"`
	CityId         *string  `json:"city_id,omitempty"`
	CityName       *string  `json:"city_name,omitempty"`
	UseScene       *int32   `json:"use_scene,omitempty"`
	TotalCount     *int32   `json:"total_count,omitempty"`
	AvailableCount *int32   `json:"available_count,omitempty"`
	TotalMoney     *float32 `json:"total_money,omitempty"`
	AvailableMoney *float32 `json:"available_money,omitempty"`
	PerorderMoney  *float32 `json:"perorder_money,omitempty"`
	StartTime      *string  `json:"start_time,omitempty"`
	EndTime        *string  `json:"end_time,omitempty"`
}

type CarRuleBuilder struct {
	ruleId            string
	ruleIdSet         bool
	ruleName          string
	ruleNameSet       bool
	ruleStatus        int32
	ruleStatusSet     bool
	cityId            string
	cityIdSet         bool
	cityName          string
	cityNameSet       bool
	useScene          int32
	useSceneSet       bool
	totalCount        int32
	totalCountSet     bool
	availableCount    int32
	availableCountSet bool
	totalMoney        float32
	totalMoneySet     bool
	availableMoney    float32
	availableMoneySet bool
	perorderMoney     float32
	perorderMoneySet  bool
	startTime         string
	startTimeSet      bool
	endTime           string
	endTimeSet        bool
}

func NewCarRuleBuilder() *CarRuleBuilder {
	return &CarRuleBuilder{}
}
func (builder *CarRuleBuilder) RuleId(ruleId string) *CarRuleBuilder {
	builder.ruleId = ruleId
	builder.ruleIdSet = true
	return builder
}
func (builder *CarRuleBuilder) RuleName(ruleName string) *CarRuleBuilder {
	builder.ruleName = ruleName
	builder.ruleNameSet = true
	return builder
}
func (builder *CarRuleBuilder) RuleStatus(ruleStatus int32) *CarRuleBuilder {
	builder.ruleStatus = ruleStatus
	builder.ruleStatusSet = true
	return builder
}
func (builder *CarRuleBuilder) CityId(cityId string) *CarRuleBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *CarRuleBuilder) CityName(cityName string) *CarRuleBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *CarRuleBuilder) UseScene(useScene int32) *CarRuleBuilder {
	builder.useScene = useScene
	builder.useSceneSet = true
	return builder
}
func (builder *CarRuleBuilder) TotalCount(totalCount int32) *CarRuleBuilder {
	builder.totalCount = totalCount
	builder.totalCountSet = true
	return builder
}
func (builder *CarRuleBuilder) AvailableCount(availableCount int32) *CarRuleBuilder {
	builder.availableCount = availableCount
	builder.availableCountSet = true
	return builder
}
func (builder *CarRuleBuilder) TotalMoney(totalMoney float32) *CarRuleBuilder {
	builder.totalMoney = totalMoney
	builder.totalMoneySet = true
	return builder
}
func (builder *CarRuleBuilder) AvailableMoney(availableMoney float32) *CarRuleBuilder {
	builder.availableMoney = availableMoney
	builder.availableMoneySet = true
	return builder
}
func (builder *CarRuleBuilder) PerorderMoney(perorderMoney float32) *CarRuleBuilder {
	builder.perorderMoney = perorderMoney
	builder.perorderMoneySet = true
	return builder
}
func (builder *CarRuleBuilder) StartTime(startTime string) *CarRuleBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *CarRuleBuilder) EndTime(endTime string) *CarRuleBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}

func (builder *CarRuleBuilder) Build() *CarRule {
	data := &CarRule{}
	if builder.ruleIdSet {
		data.RuleId = &builder.ruleId
	}
	if builder.ruleNameSet {
		data.RuleName = &builder.ruleName
	}
	if builder.ruleStatusSet {
		data.RuleStatus = &builder.ruleStatus
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.useSceneSet {
		data.UseScene = &builder.useScene
	}
	if builder.totalCountSet {
		data.TotalCount = &builder.totalCount
	}
	if builder.availableCountSet {
		data.AvailableCount = &builder.availableCount
	}
	if builder.totalMoneySet {
		data.TotalMoney = &builder.totalMoney
	}
	if builder.availableMoneySet {
		data.AvailableMoney = &builder.availableMoney
	}
	if builder.perorderMoneySet {
		data.PerorderMoney = &builder.perorderMoney
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	return data
}
