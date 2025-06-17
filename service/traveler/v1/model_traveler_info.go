package v1

// TravelerInfo struct for TravelerInfo
type TravelerInfo struct {
	TravelerId       *string          `json:"traveler_id,omitempty"`        // 滴滴侧出行人ID;更新的时候必填
	Phone            *string          `json:"phone,omitempty"`              // 外部出行人 手机号
	Name             *string          `json:"name,omitempty"`               // 外部出行人姓名
	EnglishSurname   *string          `json:"english_surname,omitempty"`    // 英文姓
	EnglishName      *string          `json:"english_name,omitempty"`       // 英文名
	Remark           *string          `json:"remark,omitempty"`             // 备注信息
	Sex              *int32           `json:"sex,omitempty"`                // 性别
	OutTravelerId    *string          `json:"out_traveler_id,omitempty"`    // 外部出行人唯一ID
	BirthDate        *string          `json:"birth_date,omitempty"`         // 出生日期
	CardList         []TravelCardInfo `json:"card_list,omitempty"`          // 身份证件信息
	ForceClearFields []string         `json:"force_clear_fields,omitempty"` // 强制清空字段;目前只支持传 card_list
}

type TravelerInfoBuilder struct {
	travelerId          string // 滴滴侧出行人ID;更新的时候必填
	travelerIdSet       bool
	phone               string // 外部出行人 手机号
	phoneSet            bool
	name                string // 外部出行人姓名
	nameSet             bool
	englishSurname      string // 英文姓
	englishSurnameSet   bool
	englishName         string // 英文名
	englishNameSet      bool
	remark              string // 备注信息
	remarkSet           bool
	sex                 int32 // 性别
	sexSet              bool
	outTravelerId       string // 外部出行人唯一ID
	outTravelerIdSet    bool
	birthDate           string // 出生日期
	birthDateSet        bool
	cardList            []TravelCardInfo // 身份证件信息
	cardListSet         bool
	forceClearFields    []string // 强制清空字段;目前只支持传 card_list
	forceClearFieldsSet bool
}

func NewTravelerInfoBuilder() *TravelerInfoBuilder {
	return &TravelerInfoBuilder{}
}
func (builder *TravelerInfoBuilder) TravelerId(travelerId string) *TravelerInfoBuilder {
	builder.travelerId = travelerId
	builder.travelerIdSet = true
	return builder
}
func (builder *TravelerInfoBuilder) Phone(phone string) *TravelerInfoBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *TravelerInfoBuilder) Name(name string) *TravelerInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *TravelerInfoBuilder) EnglishSurname(englishSurname string) *TravelerInfoBuilder {
	builder.englishSurname = englishSurname
	builder.englishSurnameSet = true
	return builder
}
func (builder *TravelerInfoBuilder) EnglishName(englishName string) *TravelerInfoBuilder {
	builder.englishName = englishName
	builder.englishNameSet = true
	return builder
}
func (builder *TravelerInfoBuilder) Remark(remark string) *TravelerInfoBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *TravelerInfoBuilder) Sex(sex int32) *TravelerInfoBuilder {
	builder.sex = sex
	builder.sexSet = true
	return builder
}
func (builder *TravelerInfoBuilder) OutTravelerId(outTravelerId string) *TravelerInfoBuilder {
	builder.outTravelerId = outTravelerId
	builder.outTravelerIdSet = true
	return builder
}
func (builder *TravelerInfoBuilder) BirthDate(birthDate string) *TravelerInfoBuilder {
	builder.birthDate = birthDate
	builder.birthDateSet = true
	return builder
}
func (builder *TravelerInfoBuilder) CardList(cardList []TravelCardInfo) *TravelerInfoBuilder {
	builder.cardList = cardList
	builder.cardListSet = true
	return builder
}
func (builder *TravelerInfoBuilder) ForceClearFields(forceClearFields []string) *TravelerInfoBuilder {
	builder.forceClearFields = forceClearFields
	builder.forceClearFieldsSet = true
	return builder
}

func (builder *TravelerInfoBuilder) Build() *TravelerInfo {
	data := &TravelerInfo{}
	if builder.travelerIdSet {
		data.TravelerId = &builder.travelerId
	}
	if builder.phoneSet {
		data.Phone = &builder.phone
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.englishSurnameSet {
		data.EnglishSurname = &builder.englishSurname
	}
	if builder.englishNameSet {
		data.EnglishName = &builder.englishName
	}
	if builder.remarkSet {
		data.Remark = &builder.remark
	}
	if builder.sexSet {
		data.Sex = &builder.sex
	}
	if builder.outTravelerIdSet {
		data.OutTravelerId = &builder.outTravelerId
	}
	if builder.birthDateSet {
		data.BirthDate = &builder.birthDate
	}
	if builder.cardListSet {
		data.CardList = builder.cardList
	}
	if builder.forceClearFieldsSet {
		data.ForceClearFields = builder.forceClearFields
	}
	return data
}
