package v1

// OutTripInfo 外部trip信息,用于滴滴首页合并卡片展示
type OutTripInfo struct {
	BeginTime     *int64  `json:"begin_time,omitempty"`     // 外部申请单开始时间，时间戳
	EndTime       *int64  `json:"end_time,omitempty"`       // 外部申请单结束时间，时间戳
	TravelPurpose *string `json:"travel_purpose,omitempty"` // 出差事由，滴滴首页卡片合并展示的事由
}

type OutTripInfoBuilder struct {
	beginTime        int64 // 外部申请单开始时间，时间戳
	beginTimeSet     bool
	endTime          int64 // 外部申请单结束时间，时间戳
	endTimeSet       bool
	travelPurpose    string // 出差事由，滴滴首页卡片合并展示的事由
	travelPurposeSet bool
}

func NewOutTripInfoBuilder() *OutTripInfoBuilder {
	return &OutTripInfoBuilder{}
}
func (builder *OutTripInfoBuilder) BeginTime(beginTime int64) *OutTripInfoBuilder {
	builder.beginTime = beginTime
	builder.beginTimeSet = true
	return builder
}
func (builder *OutTripInfoBuilder) EndTime(endTime int64) *OutTripInfoBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *OutTripInfoBuilder) TravelPurpose(travelPurpose string) *OutTripInfoBuilder {
	builder.travelPurpose = travelPurpose
	builder.travelPurposeSet = true
	return builder
}

func (builder *OutTripInfoBuilder) Build() *OutTripInfo {
	data := &OutTripInfo{}
	if builder.beginTimeSet {
		data.BeginTime = &builder.beginTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.travelPurposeSet {
		data.TravelPurpose = &builder.travelPurpose
	}
	return data
}
