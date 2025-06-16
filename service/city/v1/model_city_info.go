package v1

// CityInfo 城市信息
type CityInfo struct {
	CityId        *int32          `json:"city_id,omitempty"`         // 城市ID
	ProductType   []int32         `json:"product_type,omitempty"`    // 可预订品类枚举，[10,20,30]表示可预定：用车/机票/酒店
	CityNameCn    *string         `json:"city_name_cn,omitempty"`    // 城市中文名
	CityShortName *string         `json:"city_short_name,omitempty"` // 城市短名
	CityPathId    *string         `json:"city_path_id,omitempty"`    // 城市路径id，23-302
	CityPathCn    *string         `json:"city_path_cn,omitempty"`    // 城市路径中文，库尔佩珀县 - 波士顿
	CityPathEn    *string         `json:"city_path_en,omitempty"`    // 城市路径英文，Culpeper County - Boston
	CityNameEn    *string         `json:"city_name_en,omitempty"`    // 城市英文名
	FlightStation []FlightStation `json:"flight_station,omitempty"`  // 机场信息
	TrainStation  []TrainStation  `json:"train_station,omitempty"`   // 火车站信息
}

type CityInfoBuilder struct {
	cityId           int32 // 城市ID
	cityIdSet        bool
	productType      []int32 // 可预订品类枚举，[10,20,30]表示可预定：用车/机票/酒店
	productTypeSet   bool
	cityNameCn       string // 城市中文名
	cityNameCnSet    bool
	cityShortName    string // 城市短名
	cityShortNameSet bool
	cityPathId       string // 城市路径id，23-302
	cityPathIdSet    bool
	cityPathCn       string // 城市路径中文，库尔佩珀县 - 波士顿
	cityPathCnSet    bool
	cityPathEn       string // 城市路径英文，Culpeper County - Boston
	cityPathEnSet    bool
	cityNameEn       string // 城市英文名
	cityNameEnSet    bool
	flightStation    []FlightStation // 机场信息
	flightStationSet bool
	trainStation     []TrainStation // 火车站信息
	trainStationSet  bool
}

func NewCityInfoBuilder() *CityInfoBuilder {
	return &CityInfoBuilder{}
}
func (builder *CityInfoBuilder) CityId(cityId int32) *CityInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *CityInfoBuilder) ProductType(productType []int32) *CityInfoBuilder {
	builder.productType = productType
	builder.productTypeSet = true
	return builder
}
func (builder *CityInfoBuilder) CityNameCn(cityNameCn string) *CityInfoBuilder {
	builder.cityNameCn = cityNameCn
	builder.cityNameCnSet = true
	return builder
}
func (builder *CityInfoBuilder) CityShortName(cityShortName string) *CityInfoBuilder {
	builder.cityShortName = cityShortName
	builder.cityShortNameSet = true
	return builder
}
func (builder *CityInfoBuilder) CityPathId(cityPathId string) *CityInfoBuilder {
	builder.cityPathId = cityPathId
	builder.cityPathIdSet = true
	return builder
}
func (builder *CityInfoBuilder) CityPathCn(cityPathCn string) *CityInfoBuilder {
	builder.cityPathCn = cityPathCn
	builder.cityPathCnSet = true
	return builder
}
func (builder *CityInfoBuilder) CityPathEn(cityPathEn string) *CityInfoBuilder {
	builder.cityPathEn = cityPathEn
	builder.cityPathEnSet = true
	return builder
}
func (builder *CityInfoBuilder) CityNameEn(cityNameEn string) *CityInfoBuilder {
	builder.cityNameEn = cityNameEn
	builder.cityNameEnSet = true
	return builder
}
func (builder *CityInfoBuilder) FlightStation(flightStation []FlightStation) *CityInfoBuilder {
	builder.flightStation = flightStation
	builder.flightStationSet = true
	return builder
}
func (builder *CityInfoBuilder) TrainStation(trainStation []TrainStation) *CityInfoBuilder {
	builder.trainStation = trainStation
	builder.trainStationSet = true
	return builder
}

func (builder *CityInfoBuilder) Build() *CityInfo {
	data := &CityInfo{}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.productTypeSet {
		data.ProductType = builder.productType
	}
	if builder.cityNameCnSet {
		data.CityNameCn = &builder.cityNameCn
	}
	if builder.cityShortNameSet {
		data.CityShortName = &builder.cityShortName
	}
	if builder.cityPathIdSet {
		data.CityPathId = &builder.cityPathId
	}
	if builder.cityPathCnSet {
		data.CityPathCn = &builder.cityPathCn
	}
	if builder.cityPathEnSet {
		data.CityPathEn = &builder.cityPathEn
	}
	if builder.cityNameEnSet {
		data.CityNameEn = &builder.cityNameEn
	}
	if builder.flightStationSet {
		data.FlightStation = builder.flightStation
	}
	if builder.trainStationSet {
		data.TrainStation = builder.trainStation
	}
	return data
}
