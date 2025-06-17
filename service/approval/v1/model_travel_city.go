package v1

// TravelCity struct for TravelCity
type TravelCity struct {
	Id               *int32  `json:"id,omitempty"`                // 城市ID
	Name             *string `json:"name,omitempty"`              // 城市名称
	AddressDimension *int32  `json:"address_dimension,omitempty"` // 目的地维度 默认0; 1：国家；2：省；0：城市(包括县级市)
	CountryId        *int32  `json:"country_id,omitempty"`        // 出发国家ID
	CountryName      *string `json:"country_name,omitempty"`      // 出发国家名称
	ProvinceId       *int32  `json:"province_id,omitempty"`       // 出发省ID
	ProvinceName     *string `json:"province_name,omitempty"`     // 出发省名称
}

type TravelCityBuilder struct {
	id                  int32 // 城市ID
	idSet               bool
	name                string // 城市名称
	nameSet             bool
	addressDimension    int32 // 目的地维度 默认0; 1：国家；2：省；0：城市(包括县级市)
	addressDimensionSet bool
	countryId           int32 // 出发国家ID
	countryIdSet        bool
	countryName         string // 出发国家名称
	countryNameSet      bool
	provinceId          int32 // 出发省ID
	provinceIdSet       bool
	provinceName        string // 出发省名称
	provinceNameSet     bool
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
func (builder *TravelCityBuilder) AddressDimension(addressDimension int32) *TravelCityBuilder {
	builder.addressDimension = addressDimension
	builder.addressDimensionSet = true
	return builder
}
func (builder *TravelCityBuilder) CountryId(countryId int32) *TravelCityBuilder {
	builder.countryId = countryId
	builder.countryIdSet = true
	return builder
}
func (builder *TravelCityBuilder) CountryName(countryName string) *TravelCityBuilder {
	builder.countryName = countryName
	builder.countryNameSet = true
	return builder
}
func (builder *TravelCityBuilder) ProvinceId(provinceId int32) *TravelCityBuilder {
	builder.provinceId = provinceId
	builder.provinceIdSet = true
	return builder
}
func (builder *TravelCityBuilder) ProvinceName(provinceName string) *TravelCityBuilder {
	builder.provinceName = provinceName
	builder.provinceNameSet = true
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
	if builder.addressDimensionSet {
		data.AddressDimension = &builder.addressDimension
	}
	if builder.countryIdSet {
		data.CountryId = &builder.countryId
	}
	if builder.countryNameSet {
		data.CountryName = &builder.countryName
	}
	if builder.provinceIdSet {
		data.ProvinceId = &builder.provinceId
	}
	if builder.provinceNameSet {
		data.ProvinceName = &builder.provinceName
	}
	return data
}
