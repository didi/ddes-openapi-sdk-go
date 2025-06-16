package v1

// CityRecord struct for CityRecord
type CityRecord struct {
	ProvinceId     *int32     `json:"province_id,omitempty"`      // 省ID，国家下一层级，滴滴企业级内部主键
	ProvinceNameCn *string    `json:"province_name_cn,omitempty"` // 省中文名
	ProvinceNameEn *string    `json:"province_name_en,omitempty"` // 省英文名
	CityList       []CityInfo `json:"city_list,omitempty"`        // 城市列表
}

type CityRecordBuilder struct {
	provinceId        int32 // 省ID，国家下一层级，滴滴企业级内部主键
	provinceIdSet     bool
	provinceNameCn    string // 省中文名
	provinceNameCnSet bool
	provinceNameEn    string // 省英文名
	provinceNameEnSet bool
	cityList          []CityInfo // 城市列表
	cityListSet       bool
}

func NewCityRecordBuilder() *CityRecordBuilder {
	return &CityRecordBuilder{}
}
func (builder *CityRecordBuilder) ProvinceId(provinceId int32) *CityRecordBuilder {
	builder.provinceId = provinceId
	builder.provinceIdSet = true
	return builder
}
func (builder *CityRecordBuilder) ProvinceNameCn(provinceNameCn string) *CityRecordBuilder {
	builder.provinceNameCn = provinceNameCn
	builder.provinceNameCnSet = true
	return builder
}
func (builder *CityRecordBuilder) ProvinceNameEn(provinceNameEn string) *CityRecordBuilder {
	builder.provinceNameEn = provinceNameEn
	builder.provinceNameEnSet = true
	return builder
}
func (builder *CityRecordBuilder) CityList(cityList []CityInfo) *CityRecordBuilder {
	builder.cityList = cityList
	builder.cityListSet = true
	return builder
}

func (builder *CityRecordBuilder) Build() *CityRecord {
	data := &CityRecord{}
	if builder.provinceIdSet {
		data.ProvinceId = &builder.provinceId
	}
	if builder.provinceNameCnSet {
		data.ProvinceNameCn = &builder.provinceNameCn
	}
	if builder.provinceNameEnSet {
		data.ProvinceNameEn = &builder.provinceNameEn
	}
	if builder.cityListSet {
		data.CityList = builder.cityList
	}
	return data
}
