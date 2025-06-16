package v1

// HotelCityInfo struct for HotelCityInfo
type HotelCityInfo struct {
	CityId               *string `json:"city_id,omitempty"`                // 滴滴城市ID
	DescriptionCn        *string `json:"description_cn,omitempty"`         // 城市中文描述
	DescriptionEn        *string `json:"description_en,omitempty"`         // 城市英文描述
	CityNameCn           *string `json:"city_name_cn,omitempty"`           // 城市中文名
	CityNameEn           *string `json:"city_name_en,omitempty"`           // 城市英文名
	ProvinceId           *string `json:"province_id,omitempty"`            // 省ID
	ProvinceNameCn       *string `json:"province_name_cn,omitempty"`       // 省中文名
	ProvinceNameEn       *string `json:"province_name_en,omitempty"`       // 省英文文名
	CountryId            *string `json:"country_id,omitempty"`             // 国家ID
	CanonicalCountryCode *string `json:"canonical_country_code,omitempty"` // 国家二字代码
	CountryCode          *string `json:"country_code,omitempty"`           // 国家三字代码
	CountryNameCn        *string `json:"country_name_cn,omitempty"`        // 国家中文名
	CountryNameEn        *string `json:"country_name_en,omitempty"`        // 国家英文名
}

type HotelCityInfoBuilder struct {
	cityId                  string // 滴滴城市ID
	cityIdSet               bool
	descriptionCn           string // 城市中文描述
	descriptionCnSet        bool
	descriptionEn           string // 城市英文描述
	descriptionEnSet        bool
	cityNameCn              string // 城市中文名
	cityNameCnSet           bool
	cityNameEn              string // 城市英文名
	cityNameEnSet           bool
	provinceId              string // 省ID
	provinceIdSet           bool
	provinceNameCn          string // 省中文名
	provinceNameCnSet       bool
	provinceNameEn          string // 省英文文名
	provinceNameEnSet       bool
	countryId               string // 国家ID
	countryIdSet            bool
	canonicalCountryCode    string // 国家二字代码
	canonicalCountryCodeSet bool
	countryCode             string // 国家三字代码
	countryCodeSet          bool
	countryNameCn           string // 国家中文名
	countryNameCnSet        bool
	countryNameEn           string // 国家英文名
	countryNameEnSet        bool
}

func NewHotelCityInfoBuilder() *HotelCityInfoBuilder {
	return &HotelCityInfoBuilder{}
}
func (builder *HotelCityInfoBuilder) CityId(cityId string) *HotelCityInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) DescriptionCn(descriptionCn string) *HotelCityInfoBuilder {
	builder.descriptionCn = descriptionCn
	builder.descriptionCnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) DescriptionEn(descriptionEn string) *HotelCityInfoBuilder {
	builder.descriptionEn = descriptionEn
	builder.descriptionEnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CityNameCn(cityNameCn string) *HotelCityInfoBuilder {
	builder.cityNameCn = cityNameCn
	builder.cityNameCnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CityNameEn(cityNameEn string) *HotelCityInfoBuilder {
	builder.cityNameEn = cityNameEn
	builder.cityNameEnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) ProvinceId(provinceId string) *HotelCityInfoBuilder {
	builder.provinceId = provinceId
	builder.provinceIdSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) ProvinceNameCn(provinceNameCn string) *HotelCityInfoBuilder {
	builder.provinceNameCn = provinceNameCn
	builder.provinceNameCnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) ProvinceNameEn(provinceNameEn string) *HotelCityInfoBuilder {
	builder.provinceNameEn = provinceNameEn
	builder.provinceNameEnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CountryId(countryId string) *HotelCityInfoBuilder {
	builder.countryId = countryId
	builder.countryIdSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CanonicalCountryCode(canonicalCountryCode string) *HotelCityInfoBuilder {
	builder.canonicalCountryCode = canonicalCountryCode
	builder.canonicalCountryCodeSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CountryCode(countryCode string) *HotelCityInfoBuilder {
	builder.countryCode = countryCode
	builder.countryCodeSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CountryNameCn(countryNameCn string) *HotelCityInfoBuilder {
	builder.countryNameCn = countryNameCn
	builder.countryNameCnSet = true
	return builder
}
func (builder *HotelCityInfoBuilder) CountryNameEn(countryNameEn string) *HotelCityInfoBuilder {
	builder.countryNameEn = countryNameEn
	builder.countryNameEnSet = true
	return builder
}

func (builder *HotelCityInfoBuilder) Build() *HotelCityInfo {
	data := &HotelCityInfo{}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.descriptionCnSet {
		data.DescriptionCn = &builder.descriptionCn
	}
	if builder.descriptionEnSet {
		data.DescriptionEn = &builder.descriptionEn
	}
	if builder.cityNameCnSet {
		data.CityNameCn = &builder.cityNameCn
	}
	if builder.cityNameEnSet {
		data.CityNameEn = &builder.cityNameEn
	}
	if builder.provinceIdSet {
		data.ProvinceId = &builder.provinceId
	}
	if builder.provinceNameCnSet {
		data.ProvinceNameCn = &builder.provinceNameCn
	}
	if builder.provinceNameEnSet {
		data.ProvinceNameEn = &builder.provinceNameEn
	}
	if builder.countryIdSet {
		data.CountryId = &builder.countryId
	}
	if builder.canonicalCountryCodeSet {
		data.CanonicalCountryCode = &builder.canonicalCountryCode
	}
	if builder.countryCodeSet {
		data.CountryCode = &builder.countryCode
	}
	if builder.countryNameCnSet {
		data.CountryNameCn = &builder.countryNameCn
	}
	if builder.countryNameEnSet {
		data.CountryNameEn = &builder.countryNameEn
	}
	return data
}
