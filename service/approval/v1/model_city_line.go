package v1

// CityLine 城市线路信息
type CityLine struct {
	StartId   *string `json:"start_id,omitempty"`
	StartName *string `json:"start_name,omitempty"`
	EndId     *string `json:"end_id,omitempty"`
	EndName   *string `json:"end_name,omitempty"`
}

type CityLineBuilder struct {
	startId      string
	startIdSet   bool
	startName    string
	startNameSet bool
	endId        string
	endIdSet     bool
	endName      string
	endNameSet   bool
}

func NewCityLineBuilder() *CityLineBuilder {
	return &CityLineBuilder{}
}
func (builder *CityLineBuilder) StartId(startId string) *CityLineBuilder {
	builder.startId = startId
	builder.startIdSet = true
	return builder
}
func (builder *CityLineBuilder) StartName(startName string) *CityLineBuilder {
	builder.startName = startName
	builder.startNameSet = true
	return builder
}
func (builder *CityLineBuilder) EndId(endId string) *CityLineBuilder {
	builder.endId = endId
	builder.endIdSet = true
	return builder
}
func (builder *CityLineBuilder) EndName(endName string) *CityLineBuilder {
	builder.endName = endName
	builder.endNameSet = true
	return builder
}

func (builder *CityLineBuilder) Build() *CityLine {
	data := &CityLine{}
	if builder.startIdSet {
		data.StartId = &builder.startId
	}
	if builder.startNameSet {
		data.StartName = &builder.startName
	}
	if builder.endIdSet {
		data.EndId = &builder.endId
	}
	if builder.endNameSet {
		data.EndName = &builder.endName
	}
	return data
}
