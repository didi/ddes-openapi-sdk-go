package v1

// HomeAddressInfo 家庭住址信息
type HomeAddressInfo struct {
	City        *string `json:"city,omitempty"`         // 城市名称
	CityId      *int32  `json:"city_id,omitempty"`      // 滴滴城市ID
	CityAdcode  *string `json:"city_adcode,omitempty"`  // 国内城市行政区划代码
	AddressName *string `json:"address_name,omitempty"` // 家庭住址详细名称
}

type HomeAddressInfoBuilder struct {
	city           string // 城市名称
	citySet        bool
	cityId         int32 // 滴滴城市ID
	cityIdSet      bool
	cityAdcode     string // 国内城市行政区划代码
	cityAdcodeSet  bool
	addressName    string // 家庭住址详细名称
	addressNameSet bool
}

func NewHomeAddressInfoBuilder() *HomeAddressInfoBuilder {
	return &HomeAddressInfoBuilder{}
}
func (builder *HomeAddressInfoBuilder) City(city string) *HomeAddressInfoBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *HomeAddressInfoBuilder) CityId(cityId int32) *HomeAddressInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *HomeAddressInfoBuilder) CityAdcode(cityAdcode string) *HomeAddressInfoBuilder {
	builder.cityAdcode = cityAdcode
	builder.cityAdcodeSet = true
	return builder
}
func (builder *HomeAddressInfoBuilder) AddressName(addressName string) *HomeAddressInfoBuilder {
	builder.addressName = addressName
	builder.addressNameSet = true
	return builder
}

func (builder *HomeAddressInfoBuilder) Build() *HomeAddressInfo {
	data := &HomeAddressInfo{}
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.cityAdcodeSet {
		data.CityAdcode = &builder.cityAdcode
	}
	if builder.addressNameSet {
		data.AddressName = &builder.addressName
	}
	return data
}
