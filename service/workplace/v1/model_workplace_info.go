package v1

// WorkplaceInfo struct for WorkplaceInfo
type WorkplaceInfo struct {
	Id           *string `json:"id,omitempty"`             // 滴滴地址ID，客户ID或滴滴地址ID二选一，如果两个都有，则以滴滴ID为准
	CityId       *int32  `json:"city_id,omitempty"`        // 职场所在城市id，职场所在城市id
	OutAddressId *string `json:"out_address_id,omitempty"` // 地址所属城市id，客户地址ID，最大64字符
	Address      *string `json:"address,omitempty"`        // 地址详细名称，地址详细名称，最大100字符
	Lng          *string `json:"lng,omitempty"`            // 经度
	Lat          *string `json:"lat,omitempty"`            // 纬度
	IsWorkplace  *int32  `json:"is_workplace,omitempty"`   // 是否为职场，是否为职场，1是，0否，默认为否
	PointRange   *int32  `json:"point_range,omitempty"`    // 用车范围，用车范围 100 到1500 枚举为100的整数倍
	Remark       *string `json:"remark,omitempty"`         // 备注，备注，最大200字符
}

type WorkplaceInfoBuilder struct {
	id              string // 滴滴地址ID，客户ID或滴滴地址ID二选一，如果两个都有，则以滴滴ID为准
	idSet           bool
	cityId          int32 // 职场所在城市id，职场所在城市id
	cityIdSet       bool
	outAddressId    string // 地址所属城市id，客户地址ID，最大64字符
	outAddressIdSet bool
	address         string // 地址详细名称，地址详细名称，最大100字符
	addressSet      bool
	lng             string // 经度
	lngSet          bool
	lat             string // 纬度
	latSet          bool
	isWorkplace     int32 // 是否为职场，是否为职场，1是，0否，默认为否
	isWorkplaceSet  bool
	pointRange      int32 // 用车范围，用车范围 100 到1500 枚举为100的整数倍
	pointRangeSet   bool
	remark          string // 备注，备注，最大200字符
	remarkSet       bool
}

func NewWorkplaceInfoBuilder() *WorkplaceInfoBuilder {
	return &WorkplaceInfoBuilder{}
}
func (builder *WorkplaceInfoBuilder) Id(id string) *WorkplaceInfoBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) CityId(cityId int32) *WorkplaceInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) OutAddressId(outAddressId string) *WorkplaceInfoBuilder {
	builder.outAddressId = outAddressId
	builder.outAddressIdSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) Address(address string) *WorkplaceInfoBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) Lng(lng string) *WorkplaceInfoBuilder {
	builder.lng = lng
	builder.lngSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) Lat(lat string) *WorkplaceInfoBuilder {
	builder.lat = lat
	builder.latSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) IsWorkplace(isWorkplace int32) *WorkplaceInfoBuilder {
	builder.isWorkplace = isWorkplace
	builder.isWorkplaceSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) PointRange(pointRange int32) *WorkplaceInfoBuilder {
	builder.pointRange = pointRange
	builder.pointRangeSet = true
	return builder
}
func (builder *WorkplaceInfoBuilder) Remark(remark string) *WorkplaceInfoBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}

func (builder *WorkplaceInfoBuilder) Build() *WorkplaceInfo {
	data := &WorkplaceInfo{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.outAddressIdSet {
		data.OutAddressId = &builder.outAddressId
	}
	if builder.addressSet {
		data.Address = &builder.address
	}
	if builder.lngSet {
		data.Lng = &builder.lng
	}
	if builder.latSet {
		data.Lat = &builder.lat
	}
	if builder.isWorkplaceSet {
		data.IsWorkplace = &builder.isWorkplace
	}
	if builder.pointRangeSet {
		data.PointRange = &builder.pointRange
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	return data
}
