package v1

// HistoryDestination 历史目的地信息
type HistoryDestination struct {
	Address  *string `json:"address,omitempty"`  // 历史目的地详细地址
	Lng      *string `json:"lng,omitempty"`      // 历史目的地经度
	Lat      *string `json:"lat,omitempty"`      // 历史目的地纬度
	Name     *string `json:"name,omitempty"`     // 历史目的地地址名称
	Sequence *int32  `json:"sequence,omitempty"` // 历史目的地序号，序号从1开始递增，存在多次修改目的地时，1表示最早修改的那次
}

type HistoryDestinationBuilder struct {
	address     string // 历史目的地详细地址
	addressSet  bool
	lng         string // 历史目的地经度
	lngSet      bool
	lat         string // 历史目的地纬度
	latSet      bool
	name        string // 历史目的地地址名称
	nameSet     bool
	sequence    int32 // 历史目的地序号，序号从1开始递增，存在多次修改目的地时，1表示最早修改的那次
	sequenceSet bool
}

func NewHistoryDestinationBuilder() *HistoryDestinationBuilder {
	return &HistoryDestinationBuilder{}
}
func (builder *HistoryDestinationBuilder) Address(address string) *HistoryDestinationBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *HistoryDestinationBuilder) Lng(lng string) *HistoryDestinationBuilder {
	builder.lng = lng
	builder.lngSet = true
	return builder
}
func (builder *HistoryDestinationBuilder) Lat(lat string) *HistoryDestinationBuilder {
	builder.lat = lat
	builder.latSet = true
	return builder
}
func (builder *HistoryDestinationBuilder) Name(name string) *HistoryDestinationBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *HistoryDestinationBuilder) Sequence(sequence int32) *HistoryDestinationBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}

func (builder *HistoryDestinationBuilder) Build() *HistoryDestination {
	data := &HistoryDestination{}
	if builder.addressSet {
		data.Address = &builder.address
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
	if builder.sequenceSet {
		data.Sequence = &builder.sequence
	}
	return data
}
