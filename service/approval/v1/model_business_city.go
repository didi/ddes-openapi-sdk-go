package v1

// BusinessCity struct for BusinessCity
type BusinessCity struct {
	CityId *int32  `json:"city_id,omitempty"` // 城市ID
	City   *string `json:"city,omitempty"`    // 城市名称
}

type BusinessCityBuilder struct {
	cityId    int32 // 城市ID
	cityIdSet bool
	city      string // 城市名称
	citySet   bool
}

func NewBusinessCityBuilder() *BusinessCityBuilder {
	return &BusinessCityBuilder{}
}
func (builder *BusinessCityBuilder) CityId(cityId int32) *BusinessCityBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *BusinessCityBuilder) City(city string) *BusinessCityBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}

func (builder *BusinessCityBuilder) Build() *BusinessCity {
	data := &BusinessCity{}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.citySet {
		data.City = &builder.city
	}
	return data
}
