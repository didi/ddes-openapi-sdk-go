package v1

// PoiItem struct for PoiItem
type PoiItem struct {
	City       *string  `json:"city,omitempty"`        // 城市名称
	CityId     *int32   `json:"city_id,omitempty"`     // 城市滴滴侧ID
	CityAdcode *string  `json:"city_adcode,omitempty"` // 城市行政区划编码
	Flat       *float64 `json:"flat,omitempty"`        // 纬度
	Flng       *float64 `json:"flng,omitempty"`        // 经度
	PoiRange   *int32   `json:"poi_range,omitempty"`   // POI范围(米)
	Label      *string  `json:"label,omitempty"`       // POI标签名称
}

type PoiItemBuilder struct {
	city          string // 城市名称
	citySet       bool
	cityId        int32 // 城市滴滴侧ID
	cityIdSet     bool
	cityAdcode    string // 城市行政区划编码
	cityAdcodeSet bool
	flat          float64 // 纬度
	flatSet       bool
	flng          float64 // 经度
	flngSet       bool
	poiRange      int32 // POI范围(米)
	poiRangeSet   bool
	label         string // POI标签名称
	labelSet      bool
}

func NewPoiItemBuilder() *PoiItemBuilder {
	return &PoiItemBuilder{}
}
func (builder *PoiItemBuilder) City(city string) *PoiItemBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *PoiItemBuilder) CityId(cityId int32) *PoiItemBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *PoiItemBuilder) CityAdcode(cityAdcode string) *PoiItemBuilder {
	builder.cityAdcode = cityAdcode
	builder.cityAdcodeSet = true
	return builder
}
func (builder *PoiItemBuilder) Flat(flat float64) *PoiItemBuilder {
	builder.flat = flat
	builder.flatSet = true
	return builder
}
func (builder *PoiItemBuilder) Flng(flng float64) *PoiItemBuilder {
	builder.flng = flng
	builder.flngSet = true
	return builder
}
func (builder *PoiItemBuilder) PoiRange(poiRange int32) *PoiItemBuilder {
	builder.poiRange = poiRange
	builder.poiRangeSet = true
	return builder
}
func (builder *PoiItemBuilder) Label(label string) *PoiItemBuilder {
	builder.label = label
	builder.labelSet = true
	return builder
}

func (builder *PoiItemBuilder) Build() *PoiItem {
	data := &PoiItem{}
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.cityAdcodeSet {
		data.CityAdcode = &builder.cityAdcode
	}
	if builder.flatSet {
		data.Flat = &builder.flat
	}
	if builder.flngSet {
		data.Flng = &builder.flng
	}
	if builder.poiRangeSet {
		data.PoiRange = &builder.poiRange
	}
	if builder.labelSet {
		data.Label = &builder.label
	}
	return data
}
