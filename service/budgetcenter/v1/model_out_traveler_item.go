package v1

// OutTravelerItem struct for OutTravelerItem
type OutTravelerItem struct {
	OutTravelerId    *string               `json:"out_traveler_id,omitempty"`   // 外部出行人编码
	Id               *int64                `json:"id,omitempty"`                // 外部出行人滴滴侧ID
	RelatedEmployees []RelatedEmployeeItem `json:"related_employees,omitempty"` // 关联员工列表，最多20个
}

type OutTravelerItemBuilder struct {
	outTravelerId       string // 外部出行人编码
	outTravelerIdSet    bool
	id                  int64 // 外部出行人滴滴侧ID
	idSet               bool
	relatedEmployees    []RelatedEmployeeItem // 关联员工列表
	relatedEmployeesSet bool
}

func NewOutTravelerItemBuilder() *OutTravelerItemBuilder {
	return &OutTravelerItemBuilder{}
}
func (builder *OutTravelerItemBuilder) OutTravelerId(outTravelerId string) *OutTravelerItemBuilder {
	builder.outTravelerId = outTravelerId
	builder.outTravelerIdSet = true
	return builder
}
func (builder *OutTravelerItemBuilder) Id(id int64) *OutTravelerItemBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *OutTravelerItemBuilder) RelatedEmployees(relatedEmployees []RelatedEmployeeItem) *OutTravelerItemBuilder {
	builder.relatedEmployees = relatedEmployees
	builder.relatedEmployeesSet = true
	return builder
}

func (builder *OutTravelerItemBuilder) Build() *OutTravelerItem {
	data := &OutTravelerItem{}
	if builder.outTravelerIdSet {
		data.OutTravelerId = &builder.outTravelerId
	}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.relatedEmployeesSet {
		data.RelatedEmployees = builder.relatedEmployees
	}
	return data
}
