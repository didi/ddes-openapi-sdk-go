package v1

// TrainStation 火车站信息
type TrainStation struct {
	StationNameCn *string `json:"station_name_cn,omitempty"` // 火车站中文名
	StationNameEn *string `json:"station_name_en,omitempty"` // 火车站英文名
	StationName   *string `json:"station_name,omitempty"`    // 火车站名
	StationId     *int32  `json:"station_id,omitempty"`      // 火车站ID，滴滴主键
}

type TrainStationBuilder struct {
	stationNameCn    string // 火车站中文名
	stationNameCnSet bool
	stationNameEn    string // 火车站英文名
	stationNameEnSet bool
	stationName      string // 火车站名
	stationNameSet   bool
	stationId        int32 // 火车站ID，滴滴主键
	stationIdSet     bool
}

func NewTrainStationBuilder() *TrainStationBuilder {
	return &TrainStationBuilder{}
}
func (builder *TrainStationBuilder) StationNameCn(stationNameCn string) *TrainStationBuilder {
	builder.stationNameCn = stationNameCn
	builder.stationNameCnSet = true
	return builder
}
func (builder *TrainStationBuilder) StationNameEn(stationNameEn string) *TrainStationBuilder {
	builder.stationNameEn = stationNameEn
	builder.stationNameEnSet = true
	return builder
}
func (builder *TrainStationBuilder) StationName(stationName string) *TrainStationBuilder {
	builder.stationName = stationName
	builder.stationNameSet = true
	return builder
}
func (builder *TrainStationBuilder) StationId(stationId int32) *TrainStationBuilder {
	builder.stationId = stationId
	builder.stationIdSet = true
	return builder
}

func (builder *TrainStationBuilder) Build() *TrainStation {
	data := &TrainStation{}
	if builder.stationNameCnSet {
		data.StationNameCn = &builder.stationNameCn
	}
	if builder.stationNameEnSet {
		data.StationNameEn = &builder.stationNameEn
	}
	if builder.stationNameSet {
		data.StationName = &builder.stationName
	}
	if builder.stationIdSet {
		data.StationId = &builder.stationId
	}
	return data
}
