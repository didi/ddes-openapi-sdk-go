package v1

// SelectCondition 筛选条件
type SelectCondition struct {
	TransferCityList []string `json:"transfer_city_list,omitempty"` // 中台城市列表
	ToStationList    []string `json:"to_station_list,omitempty"`    // 到达站点列表
	FromStationList  []string `json:"from_station_list,omitempty"`  // 到达站点列表
	IsShowCross      *int32   `json:"is_show_cross,omitempty"`
	IsShowLocal      *int32   `json:"is_show_local,omitempty"`
}

type SelectConditionBuilder struct {
	transferCityList    []string // 中台城市列表
	transferCityListSet bool
	toStationList       []string // 到达站点列表
	toStationListSet    bool
	fromStationList     []string // 到达站点列表
	fromStationListSet  bool
	isShowCross         int32
	isShowCrossSet      bool
	isShowLocal         int32
	isShowLocalSet      bool
}

func NewSelectConditionBuilder() *SelectConditionBuilder {
	return &SelectConditionBuilder{}
}
func (builder *SelectConditionBuilder) TransferCityList(transferCityList []string) *SelectConditionBuilder {
	builder.transferCityList = transferCityList
	builder.transferCityListSet = true
	return builder
}
func (builder *SelectConditionBuilder) ToStationList(toStationList []string) *SelectConditionBuilder {
	builder.toStationList = toStationList
	builder.toStationListSet = true
	return builder
}
func (builder *SelectConditionBuilder) FromStationList(fromStationList []string) *SelectConditionBuilder {
	builder.fromStationList = fromStationList
	builder.fromStationListSet = true
	return builder
}
func (builder *SelectConditionBuilder) IsShowCross(isShowCross int32) *SelectConditionBuilder {
	builder.isShowCross = isShowCross
	builder.isShowCrossSet = true
	return builder
}
func (builder *SelectConditionBuilder) IsShowLocal(isShowLocal int32) *SelectConditionBuilder {
	builder.isShowLocal = isShowLocal
	builder.isShowLocalSet = true
	return builder
}

func (builder *SelectConditionBuilder) Build() *SelectCondition {
	data := &SelectCondition{}
	if builder.transferCityListSet {
		data.TransferCityList = builder.transferCityList
	}
	if builder.toStationListSet {
		data.ToStationList = builder.toStationList
	}
	if builder.fromStationListSet {
		data.FromStationList = builder.fromStationList
	}
	if builder.isShowCrossSet {
		data.IsShowCross = &builder.isShowCross
	}
	if builder.isShowLocalSet {
		data.IsShowLocal = &builder.isShowLocal
	}
	return data
}
