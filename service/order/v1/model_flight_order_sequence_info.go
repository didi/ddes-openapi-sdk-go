package v1

// FlightOrderSequenceInfo 航段信息
type FlightOrderSequenceInfo struct {
	Sequence             *int32  `json:"sequence,omitempty"`               // 航程序号，票下的航段计数，单程时为1
	DepartureTime        *string `json:"departure_time,omitempty"`         // 预计起飞时间，格式：yyyy-MM-dd HH:mm:ss
	ArrivalTime          *string `json:"arrival_time,omitempty"`           // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	FlightNumber         *string `json:"flight_number,omitempty"`          // 航班号
	ShareFlightNumber    *string `json:"share_flight_number,omitempty"`    // 共享航班号
	Class                *string `json:"class,omitempty"`                  // 舱等，枚举值字母： 物理舱位   F  头等舱     C 商务      Y经济
	CabinType            *int32  `json:"cabin_type,omitempty"`             // 舱位类型，枚举值数字： 0 未配置,1 头等舱,2 超值头等舱,3 商务舱,4 经济舱,5 经济舱精选,6 经济舱y舱
	DepartureCity        *string `json:"departure_city,omitempty"`         // 起飞城市，城市中文名
	DepartureCityId      *string `json:"departure_city_id,omitempty"`      // 起飞城市ID，滴滴城市ID
	ArrivalCity          *string `json:"arrival_city,omitempty"`           // 到达城市，城市中文名
	ArrivalCityId        *string `json:"arrival_city_id,omitempty"`        // 到达城市ID，滴滴城市ID
	StopCity             *string `json:"stop_city,omitempty"`              // 经停城市，城市中文名
	StopCityId           *string `json:"stop_city_id,omitempty"`           // 经停城市ID，滴滴城市ID
	DepartureAirport     *string `json:"departure_airport,omitempty"`      // 起飞机场，机场中文名
	DepartureAirportCode *string `json:"departure_airport_code,omitempty"` // 起飞机场代码，机场IATACODE
	ArrivalAirport       *string `json:"arrival_airport,omitempty"`        // 到达机场，机场中文名
	ArrivalAirportCode   *string `json:"arrival_airport_code,omitempty"`   // 到达机场代码，机场IATACODE
	StopAirport          *string `json:"stop_airport,omitempty"`           // 经停机场，机场中文名
	StopAirportCode      *string `json:"stop_airport_code,omitempty"`      // 经停机场代码，机场IATACODE
}

type FlightOrderSequenceInfoBuilder struct {
	sequence                int32 // 航程序号，票下的航段计数，单程时为1
	sequenceSet             bool
	departureTime           string // 预计起飞时间，格式：yyyy-MM-dd HH:mm:ss
	departureTimeSet        bool
	arrivalTime             string // 预计到达时间，格式：yyyy-MM-dd HH:mm:ss
	arrivalTimeSet          bool
	flightNumber            string // 航班号
	flightNumberSet         bool
	shareFlightNumber       string // 共享航班号
	shareFlightNumberSet    bool
	class                   string // 舱等，枚举值字母： 物理舱位   F  头等舱     C 商务      Y经济
	classSet                bool
	cabinType               int32 // 舱位类型，枚举值数字： 0 未配置,1 头等舱,2 超值头等舱,3 商务舱,4 经济舱,5 经济舱精选,6 经济舱y舱
	cabinTypeSet            bool
	departureCity           string // 起飞城市，城市中文名
	departureCitySet        bool
	departureCityId         string // 起飞城市ID，滴滴城市ID
	departureCityIdSet      bool
	arrivalCity             string // 到达城市，城市中文名
	arrivalCitySet          bool
	arrivalCityId           string // 到达城市ID，滴滴城市ID
	arrivalCityIdSet        bool
	stopCity                string // 经停城市，城市中文名
	stopCitySet             bool
	stopCityId              string // 经停城市ID，滴滴城市ID
	stopCityIdSet           bool
	departureAirport        string // 起飞机场，机场中文名
	departureAirportSet     bool
	departureAirportCode    string // 起飞机场代码，机场IATACODE
	departureAirportCodeSet bool
	arrivalAirport          string // 到达机场，机场中文名
	arrivalAirportSet       bool
	arrivalAirportCode      string // 到达机场代码，机场IATACODE
	arrivalAirportCodeSet   bool
	stopAirport             string // 经停机场，机场中文名
	stopAirportSet          bool
	stopAirportCode         string // 经停机场代码，机场IATACODE
	stopAirportCodeSet      bool
}

func NewFlightOrderSequenceInfoBuilder() *FlightOrderSequenceInfoBuilder {
	return &FlightOrderSequenceInfoBuilder{}
}
func (builder *FlightOrderSequenceInfoBuilder) Sequence(sequence int32) *FlightOrderSequenceInfoBuilder {
	builder.sequence = sequence
	builder.sequenceSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) DepartureTime(departureTime string) *FlightOrderSequenceInfoBuilder {
	builder.departureTime = departureTime
	builder.departureTimeSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ArrivalTime(arrivalTime string) *FlightOrderSequenceInfoBuilder {
	builder.arrivalTime = arrivalTime
	builder.arrivalTimeSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) FlightNumber(flightNumber string) *FlightOrderSequenceInfoBuilder {
	builder.flightNumber = flightNumber
	builder.flightNumberSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ShareFlightNumber(shareFlightNumber string) *FlightOrderSequenceInfoBuilder {
	builder.shareFlightNumber = shareFlightNumber
	builder.shareFlightNumberSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) Class(class string) *FlightOrderSequenceInfoBuilder {
	builder.class = class
	builder.classSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) CabinType(cabinType int32) *FlightOrderSequenceInfoBuilder {
	builder.cabinType = cabinType
	builder.cabinTypeSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) DepartureCity(departureCity string) *FlightOrderSequenceInfoBuilder {
	builder.departureCity = departureCity
	builder.departureCitySet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) DepartureCityId(departureCityId string) *FlightOrderSequenceInfoBuilder {
	builder.departureCityId = departureCityId
	builder.departureCityIdSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ArrivalCity(arrivalCity string) *FlightOrderSequenceInfoBuilder {
	builder.arrivalCity = arrivalCity
	builder.arrivalCitySet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ArrivalCityId(arrivalCityId string) *FlightOrderSequenceInfoBuilder {
	builder.arrivalCityId = arrivalCityId
	builder.arrivalCityIdSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) StopCity(stopCity string) *FlightOrderSequenceInfoBuilder {
	builder.stopCity = stopCity
	builder.stopCitySet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) StopCityId(stopCityId string) *FlightOrderSequenceInfoBuilder {
	builder.stopCityId = stopCityId
	builder.stopCityIdSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) DepartureAirport(departureAirport string) *FlightOrderSequenceInfoBuilder {
	builder.departureAirport = departureAirport
	builder.departureAirportSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) DepartureAirportCode(departureAirportCode string) *FlightOrderSequenceInfoBuilder {
	builder.departureAirportCode = departureAirportCode
	builder.departureAirportCodeSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ArrivalAirport(arrivalAirport string) *FlightOrderSequenceInfoBuilder {
	builder.arrivalAirport = arrivalAirport
	builder.arrivalAirportSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) ArrivalAirportCode(arrivalAirportCode string) *FlightOrderSequenceInfoBuilder {
	builder.arrivalAirportCode = arrivalAirportCode
	builder.arrivalAirportCodeSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) StopAirport(stopAirport string) *FlightOrderSequenceInfoBuilder {
	builder.stopAirport = stopAirport
	builder.stopAirportSet = true
	return builder
}
func (builder *FlightOrderSequenceInfoBuilder) StopAirportCode(stopAirportCode string) *FlightOrderSequenceInfoBuilder {
	builder.stopAirportCode = stopAirportCode
	builder.stopAirportCodeSet = true
	return builder
}

func (builder *FlightOrderSequenceInfoBuilder) Build() *FlightOrderSequenceInfo {
	data := &FlightOrderSequenceInfo{}
	if builder.sequenceSet {
		data.Sequence = &builder.sequence
	}
	if builder.departureTimeSet {
		data.DepartureTime = &builder.departureTime
	}
	if builder.arrivalTimeSet {
		data.ArrivalTime = &builder.arrivalTime
	}
	if builder.flightNumberSet {
		data.FlightNumber = &builder.flightNumber
	}
	if builder.shareFlightNumberSet {
		data.ShareFlightNumber = &builder.shareFlightNumber
	}
	if builder.classSet {
		data.Class = &builder.class
	}
	if builder.cabinTypeSet {
		data.CabinType = &builder.cabinType
	}
	if builder.departureCitySet {
		data.DepartureCity = &builder.departureCity
	}
	if builder.departureCityIdSet {
		data.DepartureCityId = &builder.departureCityId
	}
	if builder.arrivalCitySet {
		data.ArrivalCity = &builder.arrivalCity
	}
	if builder.arrivalCityIdSet {
		data.ArrivalCityId = &builder.arrivalCityId
	}
	if builder.stopCitySet {
		data.StopCity = &builder.stopCity
	}
	if builder.stopCityIdSet {
		data.StopCityId = &builder.stopCityId
	}
	if builder.departureAirportSet {
		data.DepartureAirport = &builder.departureAirport
	}
	if builder.departureAirportCodeSet {
		data.DepartureAirportCode = &builder.departureAirportCode
	}
	if builder.arrivalAirportSet {
		data.ArrivalAirport = &builder.arrivalAirport
	}
	if builder.arrivalAirportCodeSet {
		data.ArrivalAirportCode = &builder.arrivalAirportCode
	}
	if builder.stopAirportSet {
		data.StopAirport = &builder.stopAirport
	}
	if builder.stopAirportCodeSet {
		data.StopAirportCode = &builder.stopAirportCode
	}
	return data
}
