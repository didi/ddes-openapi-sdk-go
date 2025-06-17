package v1

// TrainCityInfo struct for TrainCityInfo
type TrainCityInfo struct {
	CityId               *string        `json:"city_id,omitempty"`                // 滴滴城市ID
	CityNameCn           *string        `json:"city_name_cn,omitempty"`           // 城市中文名
	CityNameEn           *string        `json:"city_name_en,omitempty"`           // 城市英文名
	ProvinceId           *string        `json:"province_id,omitempty"`            // 省ID
	ProvinceNameCn       *string        `json:"province_name_cn,omitempty"`       // 省中文名
	ProvinceNameEn       *string        `json:"province_name_en,omitempty"`       // 省英文文名
	CountryId            *string        `json:"country_id,omitempty"`             // 国家ID
	CanonicalCountryCode *string        `json:"canonical_country_code,omitempty"` // 国家二字代码
	CountryCode          *string        `json:"country_code,omitempty"`           // 国家三字代码
	CountryNameCn        *string        `json:"country_name_cn,omitempty"`        // 国家中文名
	TrainStation         []TrainStation `json:"train_station,omitempty"`          // 火车站信息
}

type TrainCityInfoBuilder struct {
	cityId                  string // 滴滴城市ID
	cityIdSet               bool
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
	trainStation            []TrainStation // 火车站信息
	trainStationSet         bool
}

func NewTrainCityInfoBuilder() *TrainCityInfoBuilder {
	return &TrainCityInfoBuilder{}
}
func (builder *TrainCityInfoBuilder) CityId(cityId string) *TrainCityInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CityNameCn(cityNameCn string) *TrainCityInfoBuilder {
	builder.cityNameCn = cityNameCn
	builder.cityNameCnSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CityNameEn(cityNameEn string) *TrainCityInfoBuilder {
	builder.cityNameEn = cityNameEn
	builder.cityNameEnSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) ProvinceId(provinceId string) *TrainCityInfoBuilder {
	builder.provinceId = provinceId
	builder.provinceIdSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) ProvinceNameCn(provinceNameCn string) *TrainCityInfoBuilder {
	builder.provinceNameCn = provinceNameCn
	builder.provinceNameCnSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) ProvinceNameEn(provinceNameEn string) *TrainCityInfoBuilder {
	builder.provinceNameEn = provinceNameEn
	builder.provinceNameEnSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CountryId(countryId string) *TrainCityInfoBuilder {
	builder.countryId = countryId
	builder.countryIdSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CanonicalCountryCode(canonicalCountryCode string) *TrainCityInfoBuilder {
	builder.canonicalCountryCode = canonicalCountryCode
	builder.canonicalCountryCodeSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CountryCode(countryCode string) *TrainCityInfoBuilder {
	builder.countryCode = countryCode
	builder.countryCodeSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) CountryNameCn(countryNameCn string) *TrainCityInfoBuilder {
	builder.countryNameCn = countryNameCn
	builder.countryNameCnSet = true
	return builder
}
func (builder *TrainCityInfoBuilder) TrainStation(trainStation []TrainStation) *TrainCityInfoBuilder {
	builder.trainStation = trainStation
	builder.trainStationSet = true
	return builder
}

func (builder *TrainCityInfoBuilder) Build() *TrainCityInfo {
	data := &TrainCityInfo{}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
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
	if builder.trainStationSet {
		data.TrainStation = builder.trainStation
	}
	return data
}
