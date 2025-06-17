package v1

// InterHotelOrderBaseInfo 酒店基本信息(国际)
type InterHotelOrderBaseInfo struct {
	DetailsAmount    []DetailsAmount `json:"details_amount,omitempty"`     // 每日价格
	PayTime          *string         `json:"pay_time,omitempty"`           // 用户支付时间，格式：yyyy-MM-dd HH:mm:ss
	CheckinTime      *string         `json:"checkin_time,omitempty"`       // 预订入住时间，格式：yyyy-MM-dd
	CheckoutTime     *string         `json:"checkout_time,omitempty"`      // 预订离店时间，格式：yyyy-MM-dd
	ConfirmTime      *string         `json:"confirm_time,omitempty"`       // 预订成功时间，格式：yyyy-MM-dd HH:mm:ss
	RealCheckinTime  *string         `json:"real_checkin_time,omitempty"`  // 实际入住时间，最晚按照预订离店时间+1天更新 格式：yyyy-MM-dd
	RealCheckoutTime *string         `json:"real_checkout_time,omitempty"` // 实际离店时间，最晚按照预订离店时间+1天更新 格式：yyyy-MM-dd
	CityName         *string         `json:"city_name,omitempty"`          // 城市，城市中文名
	CityId           *string         `json:"city_id,omitempty"`            // 城市ID，滴滴城市ID
	HotelName        *string         `json:"hotel_name,omitempty"`         // 酒店中文名称
	HotelEnglishName *string         `json:"hotel_english_name,omitempty"` // 酒店英文名称
	RoomName         *string         `json:"room_name,omitempty"`          // 房型名称
	RoomNum          *int32          `json:"room_num,omitempty"`           // 预订房间数 ，不包含提前离店
	RoomNight        *int32          `json:"room_night,omitempty"`         // 预订间夜数 ，不包含提前离店
	RealRoomNumber   *int32          `json:"real_room_number,omitempty"`   // 实际房间数 ，最晚按照预订离店时间+1天更新
	RealRoomNight    *int32          `json:"real_room_night,omitempty"`    // 实际间夜数 ，最晚按照预订离店时间+1天更新
	InvoiceType      *string         `json:"invoice_type,omitempty"`       // 酒店发票类型，枚举值：增专，增普
	Country          *string         `json:"country,omitempty"`            // * 接口实际返回，文档未标注
	CountyId         *string         `json:"county_id,omitempty"`
	CountyName       *string         `json:"county_name,omitempty"`
}

type InterHotelOrderBaseInfoBuilder struct {
	detailsAmount       []DetailsAmount // 每日价格
	detailsAmountSet    bool
	payTime             string // 用户支付时间，格式：yyyy-MM-dd HH:mm:ss
	payTimeSet          bool
	checkinTime         string // 预订入住时间，格式：yyyy-MM-dd
	checkinTimeSet      bool
	checkoutTime        string // 预订离店时间，格式：yyyy-MM-dd
	checkoutTimeSet     bool
	confirmTime         string // 预订成功时间，格式：yyyy-MM-dd HH:mm:ss
	confirmTimeSet      bool
	realCheckinTime     string // 实际入住时间，最晚按照预订离店时间+1天更新 格式：yyyy-MM-dd
	realCheckinTimeSet  bool
	realCheckoutTime    string // 实际离店时间，最晚按照预订离店时间+1天更新 格式：yyyy-MM-dd
	realCheckoutTimeSet bool
	cityName            string // 城市，城市中文名
	cityNameSet         bool
	cityId              string // 城市ID，滴滴城市ID
	cityIdSet           bool
	hotelName           string // 酒店中文名称
	hotelNameSet        bool
	hotelEnglishName    string // 酒店英文名称
	hotelEnglishNameSet bool
	roomName            string // 房型名称
	roomNameSet         bool
	roomNum             int32 // 预订房间数 ，不包含提前离店
	roomNumSet          bool
	roomNight           int32 // 预订间夜数 ，不包含提前离店
	roomNightSet        bool
	realRoomNumber      int32 // 实际房间数 ，最晚按照预订离店时间+1天更新
	realRoomNumberSet   bool
	realRoomNight       int32 // 实际间夜数 ，最晚按照预订离店时间+1天更新
	realRoomNightSet    bool
	invoiceType         string // 酒店发票类型，枚举值：增专，增普
	invoiceTypeSet      bool
	country             string // * 接口实际返回，文档未标注
	countrySet          bool
	countyId            string
	countyIdSet         bool
	countyName          string
	countyNameSet       bool
}

func NewInterHotelOrderBaseInfoBuilder() *InterHotelOrderBaseInfoBuilder {
	return &InterHotelOrderBaseInfoBuilder{}
}
func (builder *InterHotelOrderBaseInfoBuilder) DetailsAmount(detailsAmount []DetailsAmount) *InterHotelOrderBaseInfoBuilder {
	builder.detailsAmount = detailsAmount
	builder.detailsAmountSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) PayTime(payTime string) *InterHotelOrderBaseInfoBuilder {
	builder.payTime = payTime
	builder.payTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CheckinTime(checkinTime string) *InterHotelOrderBaseInfoBuilder {
	builder.checkinTime = checkinTime
	builder.checkinTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CheckoutTime(checkoutTime string) *InterHotelOrderBaseInfoBuilder {
	builder.checkoutTime = checkoutTime
	builder.checkoutTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) ConfirmTime(confirmTime string) *InterHotelOrderBaseInfoBuilder {
	builder.confirmTime = confirmTime
	builder.confirmTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RealCheckinTime(realCheckinTime string) *InterHotelOrderBaseInfoBuilder {
	builder.realCheckinTime = realCheckinTime
	builder.realCheckinTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RealCheckoutTime(realCheckoutTime string) *InterHotelOrderBaseInfoBuilder {
	builder.realCheckoutTime = realCheckoutTime
	builder.realCheckoutTimeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CityName(cityName string) *InterHotelOrderBaseInfoBuilder {
	builder.cityName = cityName
	builder.cityNameSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CityId(cityId string) *InterHotelOrderBaseInfoBuilder {
	builder.cityId = cityId
	builder.cityIdSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) HotelName(hotelName string) *InterHotelOrderBaseInfoBuilder {
	builder.hotelName = hotelName
	builder.hotelNameSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) HotelEnglishName(hotelEnglishName string) *InterHotelOrderBaseInfoBuilder {
	builder.hotelEnglishName = hotelEnglishName
	builder.hotelEnglishNameSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RoomName(roomName string) *InterHotelOrderBaseInfoBuilder {
	builder.roomName = roomName
	builder.roomNameSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RoomNum(roomNum int32) *InterHotelOrderBaseInfoBuilder {
	builder.roomNum = roomNum
	builder.roomNumSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RoomNight(roomNight int32) *InterHotelOrderBaseInfoBuilder {
	builder.roomNight = roomNight
	builder.roomNightSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RealRoomNumber(realRoomNumber int32) *InterHotelOrderBaseInfoBuilder {
	builder.realRoomNumber = realRoomNumber
	builder.realRoomNumberSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) RealRoomNight(realRoomNight int32) *InterHotelOrderBaseInfoBuilder {
	builder.realRoomNight = realRoomNight
	builder.realRoomNightSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) InvoiceType(invoiceType string) *InterHotelOrderBaseInfoBuilder {
	builder.invoiceType = invoiceType
	builder.invoiceTypeSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) Country(country string) *InterHotelOrderBaseInfoBuilder {
	builder.country = country
	builder.countrySet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CountyId(countyId string) *InterHotelOrderBaseInfoBuilder {
	builder.countyId = countyId
	builder.countyIdSet = true
	return builder
}
func (builder *InterHotelOrderBaseInfoBuilder) CountyName(countyName string) *InterHotelOrderBaseInfoBuilder {
	builder.countyName = countyName
	builder.countyNameSet = true
	return builder
}

func (builder *InterHotelOrderBaseInfoBuilder) Build() *InterHotelOrderBaseInfo {
	data := &InterHotelOrderBaseInfo{}
	if builder.detailsAmountSet {
		data.DetailsAmount = builder.detailsAmount
	}
	if builder.payTimeSet {
		data.PayTime = &builder.payTime
	}
	if builder.checkinTimeSet {
		data.CheckinTime = &builder.checkinTime
	}
	if builder.checkoutTimeSet {
		data.CheckoutTime = &builder.checkoutTime
	}
	if builder.confirmTimeSet {
		data.ConfirmTime = &builder.confirmTime
	}
	if builder.realCheckinTimeSet {
		data.RealCheckinTime = &builder.realCheckinTime
	}
	if builder.realCheckoutTimeSet {
		data.RealCheckoutTime = &builder.realCheckoutTime
	}
	if builder.cityNameSet {
		data.CityName = &builder.cityName
	}
	if builder.cityIdSet {
		data.CityId = &builder.cityId
	}
	if builder.hotelNameSet {
		data.HotelName = &builder.hotelName
	}
	if builder.hotelEnglishNameSet {
		data.HotelEnglishName = &builder.hotelEnglishName
	}
	if builder.roomNameSet {
		data.RoomName = &builder.roomName
	}
	if builder.roomNumSet {
		data.RoomNum = &builder.roomNum
	}
	if builder.roomNightSet {
		data.RoomNight = &builder.roomNight
	}
	if builder.realRoomNumberSet {
		data.RealRoomNumber = &builder.realRoomNumber
	}
	if builder.realRoomNightSet {
		data.RealRoomNight = &builder.realRoomNight
	}
	if builder.invoiceTypeSet {
		data.InvoiceType = &builder.invoiceType
	}
	if builder.countrySet {
		data.Country = &builder.country
	}
	if builder.countyIdSet {
		data.CountyId = &builder.countyId
	}
	if builder.countyNameSet {
		data.CountyName = &builder.countyName
	}
	return data
}
