package v1

// Trip 行程信息(TravelDetail中trips列表Obj信息)
type Trip struct {
	DepartureCity               *string      `json:"departure_city,omitempty"`                // 出发城市名称，出发城市名称 城市名称和城市ID建议传一种，不要一起传。
	DepartureCityId             *int32       `json:"departure_city_id,omitempty"`             // 滴滴出发城市ID，出发城市 ID，国际需要使用ID
	DestinationCity             *string      `json:"destination_city,omitempty"`              // 到达城市名称，到达城市名称 ；非轻度管控必传
	DestinationCityId           *int32       `json:"destination_city_id,omitempty"`           // 滴滴到达城市ID，到达城市 ID；非轻度管控必传 国际需要使用ID
	StartDate                   *string      `json:"start_date,omitempty"`                    // 行程段开始时间，行程段开始时间 轻度管控模式下，必须与 travel_detail 中的 start_date 相同 举例：2023-01-02
	EndDate                     *string      `json:"end_date,omitempty"`                      // 行程段结束时间，行程段结束时间 轻度管控模式下，必须与 travel_detail 中的 end_date 相同 举例：2023-01-14
	TripType                    *string      `json:"trip_type,omitempty"`                     // 出行方式，出行方式：0-其他；1-火车 ;2-飞机；支持多个，多个用英文逗号分开（用车酒店默认根据制度配置生效）举例：0,1,2，中度管控按照第一段生效
	ToCitys                     []TravelCity `json:"to_citys,omitempty"`                      // 目的城市集合，轻度管控必传，其他模式传了不生效 举例：[{\"id\":1,\"name\":\"北京\"}]
	IsReturn                    *int32       `json:"is_return,omitempty"`                     // 单程/往返，控制严格管控模式下，滴滴侧是否识别这个trip为往返行程生成返程行程权限 非必传字段，0-单程， 1-往返；默认为单程；传1后的拆分逻辑为：去程使用trip的传参生成，返程生成规则为：目的地-出发地，end_date-end_date，使用去程的trip_type；若对去程和返程的交通方式有分别管控诉求，不推荐使用此字段，建议拆成两个trip传输
	DepartureAddressDimension   *int32       `json:"departure_address_dimension,omitempty"`   // 出发地维度 默认0; 1：国家；2：省；0：城市(包括县级市)
	DepartureCountryId          *int32       `json:"departure_country_id,omitempty"`          // 出发国家ID
	DepartureCountryName        *string      `json:"departure_country_name,omitempty"`        // 出发国家名称
	DepartureProvinceId         *int32       `json:"departure_province_id,omitempty"`         // 出发省ID
	DepartureProvinceName       *string      `json:"departure_province_name,omitempty"`       // 出发省名称
	DestinationAddressDimension *int32       `json:"destination_address_dimension,omitempty"` // 目的地维度 默认0; 1：国家；2：省；0：城市(包括县级市) 地址维度字段
	DestinationCountryId        *int32       `json:"destination_country_id,omitempty"`        // 到达国家ID
	DestinationCountryName      *string      `json:"destination_country_name,omitempty"`      // 到达国家名称
	DestinationProvinceId       *int32       `json:"destination_province_id,omitempty"`       // 到达省ID
	DestinationProvinceName     *string      `json:"destination_province_name,omitempty"`     // 到达省名称
}

type TripBuilder struct {
	departureCity                  string // 出发城市名称，出发城市名称 城市名称和城市ID建议传一种，不要一起传。
	departureCitySet               bool
	departureCityId                int32 // 滴滴出发城市ID，出发城市 ID，国际需要使用ID
	departureCityIdSet             bool
	destinationCity                string // 到达城市名称，到达城市名称 ；非轻度管控必传
	destinationCitySet             bool
	destinationCityId              int32 // 滴滴到达城市ID，到达城市 ID；非轻度管控必传 国际需要使用ID
	destinationCityIdSet           bool
	startDate                      string // 行程段开始时间，行程段开始时间 轻度管控模式下，必须与 travel_detail 中的 start_date 相同 举例：2023-01-02
	startDateSet                   bool
	endDate                        string // 行程段结束时间，行程段结束时间 轻度管控模式下，必须与 travel_detail 中的 end_date 相同 举例：2023-01-14
	endDateSet                     bool
	tripType                       string // 出行方式，出行方式：0-其他；1-火车 ;2-飞机；支持多个，多个用英文逗号分开（用车酒店默认根据制度配置生效）举例：0,1,2，中度管控按照第一段生效
	tripTypeSet                    bool
	toCitys                        []TravelCity // 目的城市集合，轻度管控必传，其他模式传了不生效 举例：[{\"id\":1,\"name\":\"北京\"}]
	toCitysSet                     bool
	isReturn                       int32 // 单程/往返，控制严格管控模式下，滴滴侧是否识别这个trip为往返行程生成返程行程权限 非必传字段，0-单程， 1-往返；默认为单程；传1后的拆分逻辑为：去程使用trip的传参生成，返程生成规则为：目的地-出发地，end_date-end_date，使用去程的trip_type；若对去程和返程的交通方式有分别管控诉求，不推荐使用此字段，建议拆成两个trip传输
	isReturnSet                    bool
	departureAddressDimension      int32 // 出发地维度 默认0; 1：国家；2：省；0：城市(包括县级市)
	departureAddressDimensionSet   bool
	departureCountryId             int32 // 出发国家ID
	departureCountryIdSet          bool
	departureCountryName           string // 出发国家名称
	departureCountryNameSet        bool
	departureProvinceId            int32 // 出发省ID
	departureProvinceIdSet         bool
	departureProvinceName          string // 出发省名称
	departureProvinceNameSet       bool
	destinationAddressDimension    int32 // 目的地维度 默认0; 1：国家；2：省；0：城市(包括县级市) 地址维度字段
	destinationAddressDimensionSet bool
	destinationCountryId           int32 // 到达国家ID
	destinationCountryIdSet        bool
	destinationCountryName         string // 到达国家名称
	destinationCountryNameSet      bool
	destinationProvinceId          int32 // 到达省ID
	destinationProvinceIdSet       bool
	destinationProvinceName        string // 到达省名称
	destinationProvinceNameSet     bool
}

func NewTripBuilder() *TripBuilder {
	return &TripBuilder{}
}
func (builder *TripBuilder) DepartureCity(departureCity string) *TripBuilder {
	builder.departureCity = departureCity
	builder.departureCitySet = true
	return builder
}
func (builder *TripBuilder) DepartureCityId(departureCityId int32) *TripBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *TripBuilder) DestinationCity(destinationCity string) *TripBuilder {
	builder.destinationCity = destinationCity
	builder.destinationCitySet = true
	return builder
}
func (builder *TripBuilder) DestinationCityId(destinationCityId int32) *TripBuilder {
	builder.destinationCityId = destinationCityId
	builder.destinationCityIdSet = true
	return builder
}
func (builder *TripBuilder) StartDate(startDate string) *TripBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *TripBuilder) EndDate(endDate string) *TripBuilder {
	builder.endDate = endDate
	builder.endDateSet = true
	return builder
}
func (builder *TripBuilder) TripType(tripType string) *TripBuilder {
	builder.tripType = tripType
	builder.tripTypeSet = true
	return builder
}
func (builder *TripBuilder) ToCitys(toCitys []TravelCity) *TripBuilder {
	builder.toCitys = toCitys
	builder.toCitysSet = true
	return builder
}
func (builder *TripBuilder) IsReturn(isReturn int32) *TripBuilder {
	builder.isReturn = isReturn
	builder.isReturnSet = true
	return builder
}
func (builder *TripBuilder) DepartureAddressDimension(departureAddressDimension int32) *TripBuilder {
	builder.departureAddressDimension = departureAddressDimension
	builder.departureAddressDimensionSet = true
	return builder
}
func (builder *TripBuilder) DepartureCountryId(departureCountryId int32) *TripBuilder {
	builder.departureCountryId = departureCountryId
	builder.departureCountryIdSet = true
	return builder
}
func (builder *TripBuilder) DepartureCountryName(departureCountryName string) *TripBuilder {
	builder.departureCountryName = departureCountryName
	builder.departureCountryNameSet = true
	return builder
}
func (builder *TripBuilder) DepartureProvinceId(departureProvinceId int32) *TripBuilder {
	builder.departureProvinceId = departureProvinceId
	builder.departureProvinceIdSet = true
	return builder
}
func (builder *TripBuilder) DepartureProvinceName(departureProvinceName string) *TripBuilder {
	builder.departureProvinceName = departureProvinceName
	builder.departureProvinceNameSet = true
	return builder
}
func (builder *TripBuilder) DestinationAddressDimension(destinationAddressDimension int32) *TripBuilder {
	builder.destinationAddressDimension = destinationAddressDimension
	builder.destinationAddressDimensionSet = true
	return builder
}
func (builder *TripBuilder) DestinationCountryId(destinationCountryId int32) *TripBuilder {
	builder.destinationCountryId = destinationCountryId
	builder.destinationCountryIdSet = true
	return builder
}
func (builder *TripBuilder) DestinationCountryName(destinationCountryName string) *TripBuilder {
	builder.destinationCountryName = destinationCountryName
	builder.destinationCountryNameSet = true
	return builder
}
func (builder *TripBuilder) DestinationProvinceId(destinationProvinceId int32) *TripBuilder {
	builder.destinationProvinceId = destinationProvinceId
	builder.destinationProvinceIdSet = true
	return builder
}
func (builder *TripBuilder) DestinationProvinceName(destinationProvinceName string) *TripBuilder {
	builder.destinationProvinceName = destinationProvinceName
	builder.destinationProvinceNameSet = true
	return builder
}

func (builder *TripBuilder) Build() *Trip {
	data := &Trip{}
	if builder.departureCitySet {
		data.DepartureCity = &builder.departureCity
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.destinationCitySet {
		data.DestinationCity = &builder.destinationCity
	}
	if builder.destinationCityIdSet {
		data.DestinationCityId = &builder.destinationCityId
	}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.endDateSet {
		data.EndDate = &builder.endDate
	}
	if builder.tripTypeSet {
		data.TripType = &builder.tripType
	}
	if builder.toCitysSet {
		data.ToCitys = builder.toCitys
	}
	if builder.isReturnSet {
		data.IsReturn = &builder.isReturn
	}
	if builder.departureAddressDimensionSet {
		data.DepartureAddressDimension = &builder.departureAddressDimension
	}
	if builder.departureCountryIdSet {
		data.DepartureCountryId = &builder.departureCountryId
	}
	if builder.departureCountryNameSet {
		data.DepartureCountryName = &builder.departureCountryName
	}
	if builder.departureProvinceIdSet {
		data.DepartureProvinceId = &builder.departureProvinceId
	}
	if builder.departureProvinceNameSet {
		data.DepartureProvinceName = &builder.departureProvinceName
	}
	if builder.destinationAddressDimensionSet {
		data.DestinationAddressDimension = &builder.destinationAddressDimension
	}
	if builder.destinationCountryIdSet {
		data.DestinationCountryId = &builder.destinationCountryId
	}
	if builder.destinationCountryNameSet {
		data.DestinationCountryName = &builder.destinationCountryName
	}
	if builder.destinationProvinceIdSet {
		data.DestinationProvinceId = &builder.destinationProvinceId
	}
	if builder.destinationProvinceNameSet {
		data.DestinationProvinceName = &builder.destinationProvinceName
	}
	return data
}
