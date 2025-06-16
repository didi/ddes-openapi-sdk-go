package v1

// CarCityInfo struct for CarCityInfo
type CarCityInfo struct {
	CityId     *int32         `json:"city_id,omitempty"`     // 地级市ID（如：深圳ID为2）
	CityName   *string        `json:"city_name,omitempty"`   // 地级市名称
	CountyList []CountyRecord `json:"county_list,omitempty"` // 区县列表
}

type CarCityInfoBuilder struct {
	cityId        int32 // 地级市ID（如：深圳ID为2）
	cityIdSet     bool
	cityName      string // 地级市名称
	cityNameSet   bool
	countyList    []CountyRecord // 区县列表
	countyListSet bool
}

func NewCarCityInfoBuilder() *CarCityInfoBuilder {
	return &CarCityInfoBuilder{}
}
func (builder *CarCityInfoBuilder) CityId(cityId int32) *CarCityInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *CarCityInfoBuilder) CityName(cityName string) *CarCityInfoBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *CarCityInfoBuilder) CountyList(countyList []CountyRecord) *CarCityInfoBuilder {
	builder.countyList = countyList
	builder.countyListSet = true
	return builder
}

func (builder *CarCityInfoBuilder) Build() *CarCityInfo {
	data := &CarCityInfo{}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.countyListSet {
		data.CountyList = builder.countyList
	}
	return data
}
