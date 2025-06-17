package v1

// FlightOrderSrvPackInfo struct for FlightOrderSrvPackInfo
type FlightOrderSrvPackInfo struct {
	List []SrvPackListItem `json:"list,omitempty"` // 服务包信息
}

type FlightOrderSrvPackInfoBuilder struct {
	list    []SrvPackListItem // 服务包信息
	listSet bool
}

func NewFlightOrderSrvPackInfoBuilder() *FlightOrderSrvPackInfoBuilder {
	return &FlightOrderSrvPackInfoBuilder{}
}
func (builder *FlightOrderSrvPackInfoBuilder) List(list []SrvPackListItem) *FlightOrderSrvPackInfoBuilder {
	builder.list = list
	builder.listSet = true
	return builder
}

func (builder *FlightOrderSrvPackInfoBuilder) Build() *FlightOrderSrvPackInfo {
	data := &FlightOrderSrvPackInfo{}
	if builder.listSet {
		data.List = builder.list
	}
	return data
}
