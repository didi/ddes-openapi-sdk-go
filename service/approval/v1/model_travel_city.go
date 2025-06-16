package v1

// TravelCity struct for TravelCity
type TravelCity struct {
	Id   *int32  `json:"id,omitempty"`   // 城市ID
	Name *string `json:"name,omitempty"` // 城市名称
}

type TravelCityBuilder struct {
	id      int32 // 城市ID
	idSet   bool
	name    string // 城市名称
	nameSet bool
}

func NewTravelCityBuilder() *TravelCityBuilder {
	return &TravelCityBuilder{}
}
func (builder *TravelCityBuilder) Id(id int32) *TravelCityBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *TravelCityBuilder) Name(name string) *TravelCityBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}

func (builder *TravelCityBuilder) Build() *TravelCity {
	data := &TravelCity{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	return data
}
