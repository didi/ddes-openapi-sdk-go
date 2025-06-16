package v1

// StopoverPoint 途经点信息
type StopoverPoint struct {
	Cityid  *string `json:"cityid,omitempty"`  // 途经点城市id
	Address *string `json:"address,omitempty"` // 详细地址，途经点的详细地址
	StopId  *int32  `json:"stop_id,omitempty"` // 途径点id
	City    *string `json:"city,omitempty"`    // 途经点城市名称
	Lng     *string `json:"lng,omitempty"`     // 途经点经度
	Lat     *string `json:"lat,omitempty"`     // 途经点维度
	Name    *string `json:"name,omitempty"`    // 途经点名称
	Status  *int32  `json:"status,omitempty"`  // 途径点状态，状态分为0和1 , 0代表还未经过该途经点,1 代表已经经过该途经点
}

type StopoverPointBuilder struct {
	cityid     string // 途经点城市id
	cityidSet  bool
	address    string // 详细地址，途经点的详细地址
	addressSet bool
	stopId     int32 // 途径点id
	stopIdSet  bool
	city       string // 途经点城市名称
	citySet    bool
	lng        string // 途经点经度
	lngSet     bool
	lat        string // 途经点维度
	latSet     bool
	name       string // 途经点名称
	nameSet    bool
	status     int32 // 途径点状态，状态分为0和1 , 0代表还未经过该途经点,1 代表已经经过该途经点
	statusSet  bool
}

func NewStopoverPointBuilder() *StopoverPointBuilder {
	return &StopoverPointBuilder{}
}
func (builder *StopoverPointBuilder) Cityid(cityid string) *StopoverPointBuilder {
	builder.cityid = cityid
	builder.cityidSet = true
	return builder
}
func (builder *StopoverPointBuilder) Address(address string) *StopoverPointBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *StopoverPointBuilder) StopId(stopId int32) *StopoverPointBuilder {
	builder.stopId = stopId
	builder.stopIdSet = true
	return builder
}
func (builder *StopoverPointBuilder) City(city string) *StopoverPointBuilder {
	builder.city = city
	builder.citySet = true
	return builder
}
func (builder *StopoverPointBuilder) Lng(lng string) *StopoverPointBuilder {
	builder.lng = lng
	builder.lngSet = true
	return builder
}
func (builder *StopoverPointBuilder) Lat(lat string) *StopoverPointBuilder {
	builder.lat = lat
	builder.latSet = true
	return builder
}
func (builder *StopoverPointBuilder) Name(name string) *StopoverPointBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *StopoverPointBuilder) Status(status int32) *StopoverPointBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}

func (builder *StopoverPointBuilder) Build() *StopoverPoint {
	data := &StopoverPoint{}
	if builder.cityidSet {
		data.Cityid = &builder.cityid
	}
	if builder.addressSet {
		data.Address = &builder.address
	}
	if builder.stopIdSet {
		data.StopId = &builder.stopId
	}
	if builder.citySet {
		data.City = &builder.city
	}
	if builder.lngSet {
		data.Lng = &builder.lng
	}
	if builder.latSet {
		data.Lat = &builder.lat
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	return data
}
