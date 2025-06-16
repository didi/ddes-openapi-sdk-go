package v1

// FlightInfo 航班信息的结构
type FlightInfo struct {
	Routes []Routes `json:"routes,omitempty"` // 航班信息，可为直飞或中转。该数组长度为1，表示直飞；长度为2表示中转
}

type FlightInfoBuilder struct {
	routes    []Routes // 航班信息，可为直飞或中转。该数组长度为1，表示直飞；长度为2表示中转
	routesSet bool
}

func NewFlightInfoBuilder() *FlightInfoBuilder {
	return &FlightInfoBuilder{}
}
func (builder *FlightInfoBuilder) Routes(routes []Routes) *FlightInfoBuilder {
	builder.routes = routes
	builder.routesSet = true
	return builder
}

func (builder *FlightInfoBuilder) Build() *FlightInfo {
	data := &FlightInfo{}
	if builder.routesSet {
		data.Routes = builder.routes
	}
	return data
}
