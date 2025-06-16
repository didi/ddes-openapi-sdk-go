package v1

// TransferTrainTicketListItem 中转车次详细信息
type TransferTrainTicketListItem struct {
	SupplierId          *int32        `json:"supplier_id,omitempty"`           // 渠道来源
	ToStationName       *string       `json:"to_station_name,omitempty"`       // 到达站点名称
	TransferStationName *string       `json:"transfer_station_name,omitempty"` // 中转站点名称
	TransferStationType *int32        `json:"transfer_station_type,omitempty"` // 中转站类型，1-同站换乘；2-跨站换乘
	TransferCityName    *string       `json:"transfer_city_name,omitempty"`    // 中转城市名称
	TotalRuntime        *string       `json:"total_runtime,omitempty"`         // 两程总运行时长
	DayDifference       *int32        `json:"day_difference,omitempty"`        // 全程跨天数 如:0 代表当天到达，1 代表第二天到达
	FromStationName     *string       `json:"from_station_name,omitempty"`     // 出发站名称
	TransferStopTime    *string       `json:"transfer_stop_time,omitempty"`    // 中转停留时长
	SegmentItems        []SegmentItem `json:"segment_items,omitempty"`         // 行程列表
}

type TransferTrainTicketListItemBuilder struct {
	supplierId             int32 // 渠道来源
	supplierIdSet          bool
	toStationName          string // 到达站点名称
	toStationNameSet       bool
	transferStationName    string // 中转站点名称
	transferStationNameSet bool
	transferStationType    int32 // 中转站类型，1-同站换乘；2-跨站换乘
	transferStationTypeSet bool
	transferCityName       string // 中转城市名称
	transferCityNameSet    bool
	totalRuntime           string // 两程总运行时长
	totalRuntimeSet        bool
	dayDifference          int32 // 全程跨天数 如:0 代表当天到达，1 代表第二天到达
	dayDifferenceSet       bool
	fromStationName        string // 出发站名称
	fromStationNameSet     bool
	transferStopTime       string // 中转停留时长
	transferStopTimeSet    bool
	segmentItems           []SegmentItem // 行程列表
	segmentItemsSet        bool
}

func NewTransferTrainTicketListItemBuilder() *TransferTrainTicketListItemBuilder {
	return &TransferTrainTicketListItemBuilder{}
}
func (builder *TransferTrainTicketListItemBuilder) SupplierId(supplierId int32) *TransferTrainTicketListItemBuilder {
	builder.supplierId = supplierId
	builder.supplierIdSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) ToStationName(toStationName string) *TransferTrainTicketListItemBuilder {
	builder.toStationName = toStationName
	builder.toStationNameSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) TransferStationName(transferStationName string) *TransferTrainTicketListItemBuilder {
	builder.transferStationName = transferStationName
	builder.transferStationNameSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) TransferStationType(transferStationType int32) *TransferTrainTicketListItemBuilder {
	builder.transferStationType = transferStationType
	builder.transferStationTypeSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) TransferCityName(transferCityName string) *TransferTrainTicketListItemBuilder {
	builder.transferCityName = transferCityName
	builder.transferCityNameSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) TotalRuntime(totalRuntime string) *TransferTrainTicketListItemBuilder {
	builder.totalRuntime = totalRuntime
	builder.totalRuntimeSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) DayDifference(dayDifference int32) *TransferTrainTicketListItemBuilder {
	builder.dayDifference = dayDifference
	builder.dayDifferenceSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) FromStationName(fromStationName string) *TransferTrainTicketListItemBuilder {
	builder.fromStationName = fromStationName
	builder.fromStationNameSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) TransferStopTime(transferStopTime string) *TransferTrainTicketListItemBuilder {
	builder.transferStopTime = transferStopTime
	builder.transferStopTimeSet = true
	return builder
}
func (builder *TransferTrainTicketListItemBuilder) SegmentItems(segmentItems []SegmentItem) *TransferTrainTicketListItemBuilder {
	builder.segmentItems = segmentItems
	builder.segmentItemsSet = true
	return builder
}

func (builder *TransferTrainTicketListItemBuilder) Build() *TransferTrainTicketListItem {
	data := &TransferTrainTicketListItem{}
	if builder.supplierIdSet {
		data.SupplierId = &builder.supplierId
	}
	if builder.toStationNameSet {
		data.ToStationName = &builder.toStationName
	}
	if builder.transferStationNameSet {
		data.TransferStationName = &builder.transferStationName
	}
	if builder.transferStationTypeSet {
		data.TransferStationType = &builder.transferStationType
	}
	if builder.transferCityNameSet {
		data.TransferCityName = &builder.transferCityName
	}
	if builder.totalRuntimeSet {
		data.TotalRuntime = &builder.totalRuntime
	}
	if builder.dayDifferenceSet {
		data.DayDifference = &builder.dayDifference
	}
	if builder.fromStationNameSet {
		data.FromStationName = &builder.fromStationName
	}
	if builder.transferStopTimeSet {
		data.TransferStopTime = &builder.transferStopTime
	}
	if builder.segmentItemsSet {
		data.SegmentItems = builder.segmentItems
	}
	return data
}
