package v1

// ListCityParamObj struct for ListCityParamObj
type ListCityParamObj struct {
	CountryId   *int32  `json:"country_id,omitempty"`   // 国家ID
	ProvinceId  *int32  `json:"province_id,omitempty"`  // 省ID，country_id，province_id，city_id 请求的时候省和城市ID加入时作为条件查询返回。country_id+province_id 返回该省的数据country_id+city_id 返回该城市的数据city_id 优先级大于province_id
	CityId      *int32  `json:"city_id,omitempty"`      // 城市ID，country_id，province_id，city_id 请求的时候省和城市ID加入时作为条件查询返回。country_id+province_id 返回该省的数据country_id+city_id 返回该城市的数据city_id 优先级大于province_id
	ProductType *string `json:"product_type,omitempty"` // 可预订品类枚举，枚举值：10: 用车20: 机票30: 酒店40: 火车票 查询多个品类英文逗号隔开。举例：10,20,30
}

type ListCityParamObjBuilder struct {
	countryId      int32 // 国家ID
	countryIdSet   bool
	provinceId     int32 // 省ID，country_id，province_id，city_id 请求的时候省和城市ID加入时作为条件查询返回。country_id+province_id 返回该省的数据country_id+city_id 返回该城市的数据city_id 优先级大于province_id
	provinceIdSet  bool
	cityId         int32 // 城市ID，country_id，province_id，city_id 请求的时候省和城市ID加入时作为条件查询返回。country_id+province_id 返回该省的数据country_id+city_id 返回该城市的数据city_id 优先级大于province_id
	cityIdSet      bool
	productType    string // 可预订品类枚举，枚举值：10: 用车20: 机票30: 酒店40: 火车票 查询多个品类英文逗号隔开。举例：10,20,30
	productTypeSet bool
}

func NewListCityParamObjBuilder() *ListCityParamObjBuilder {
	return &ListCityParamObjBuilder{}
}
func (builder *ListCityParamObjBuilder) CountryId(countryId int32) *ListCityParamObjBuilder {
	builder.countryId = countryId
	builder.countryIdSet = true
	return builder
}
func (builder *ListCityParamObjBuilder) ProvinceId(provinceId int32) *ListCityParamObjBuilder {
	builder.provinceId = provinceId
	builder.provinceIdSet = true
	return builder
}
func (builder *ListCityParamObjBuilder) CityId(cityId int32) *ListCityParamObjBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *ListCityParamObjBuilder) ProductType(productType string) *ListCityParamObjBuilder {
	builder.productType = productType
	builder.productTypeSet = true
	return builder
}

func (builder *ListCityParamObjBuilder) Build() *ListCityParamObj {
	data := &ListCityParamObj{}
	if builder.countryIdSet {
		data.CountryId = &builder.countryId
	}
	if builder.provinceIdSet {
		data.ProvinceId = &builder.provinceId
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.productTypeSet {
		data.ProductType = &builder.productType
	}
	return data
}
