package v1

// OutTravelerInfo 外部出行人信息
type OutTravelerInfo struct {
	TravelerId       *string                `json:"traveler_id,omitempty"`       // 外部出行人滴滴侧ID
	OutTravelerId    *string                `json:"out_traveler_id,omitempty"`   // 外部出行人编码
	Phone            *string                `json:"phone,omitempty"`             // 手机号
	Name             *string                `json:"name,omitempty"`              // 姓名
	EnglishSurname   *string                `json:"english_surname,omitempty"`   // 英文姓
	EnglishName      *string                `json:"english_name,omitempty"`      // 英文名
	Remark           *string                `json:"remark,omitempty"`            // 备注信息
	Sex              *int32                 `json:"sex,omitempty"`               // 性别 0:未知 1:男 2:女
	BirthDate        *string                `json:"birth_date,omitempty"`        // 出生日期 yyyy-MM-dd
	CardList         []*CardInfo            `json:"card_list,omitempty"`         // 证件信息列表
	RelatedEmployees []*RelatedEmployeeInfo `json:"related_employees,omitempty"` // 关联内部员工列表
	ProjectId        *string                `json:"project_id,omitempty"`        // 所属项目滴滴侧ID
	ProjectName      *string                `json:"project_name,omitempty"`      // 所属项目名称
	OutBudgetId      *string                `json:"out_budget_id,omitempty"`     // 所属项目编码
}

type OutTravelerInfoBuilder struct {
	travelerId          string
	travelerIdSet       bool
	outTravelerId       string
	outTravelerIdSet    bool
	phone               string
	phoneSet            bool
	name                string
	nameSet             bool
	englishSurname      string
	englishSurnameSet   bool
	englishName         string
	englishNameSet      bool
	remark              string
	remarkSet           bool
	sex                 int32
	sexSet              bool
	birthDate           string
	birthDateSet        bool
	cardList            []*CardInfo
	cardListSet         bool
	relatedEmployees    []*RelatedEmployeeInfo
	relatedEmployeesSet bool
	projectId           string
	projectIdSet        bool
	projectName         string
	projectNameSet      bool
	outBudgetId         string
	outBudgetIdSet      bool
}

func NewOutTravelerInfoBuilder() *OutTravelerInfoBuilder {
	return &OutTravelerInfoBuilder{}
}
func (builder *OutTravelerInfoBuilder) TravelerId(travelerId string) *OutTravelerInfoBuilder {
	builder.travelerId = travelerId
	builder.travelerIdSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) OutTravelerId(outTravelerId string) *OutTravelerInfoBuilder {
	builder.outTravelerId = outTravelerId
	builder.outTravelerIdSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) Phone(phone string) *OutTravelerInfoBuilder {
	builder.phone = phone
	builder.phoneSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) Name(name string) *OutTravelerInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) EnglishSurname(englishSurname string) *OutTravelerInfoBuilder {
	builder.englishSurname = englishSurname
	builder.englishSurnameSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) EnglishName(englishName string) *OutTravelerInfoBuilder {
	builder.englishName = englishName
	builder.englishNameSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) Remark(remark string) *OutTravelerInfoBuilder {
	builder.remark = remark
	builder.remarkSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) Sex(sex int32) *OutTravelerInfoBuilder {
	builder.sex = sex
	builder.sexSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) BirthDate(birthDate string) *OutTravelerInfoBuilder {
	builder.birthDate = birthDate
	builder.birthDateSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) CardList(cardList []*CardInfo) *OutTravelerInfoBuilder {
	builder.cardList = cardList
	builder.cardListSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) RelatedEmployees(relatedEmployees []*RelatedEmployeeInfo) *OutTravelerInfoBuilder {
	builder.relatedEmployees = relatedEmployees
	builder.relatedEmployeesSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) ProjectId(projectId string) *OutTravelerInfoBuilder {
	builder.projectId = projectId
	builder.projectIdSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) ProjectName(projectName string) *OutTravelerInfoBuilder {
	builder.projectName = projectName
	builder.projectNameSet = true
	return builder
}
func (builder *OutTravelerInfoBuilder) OutBudgetId(outBudgetId string) *OutTravelerInfoBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}

func (builder *OutTravelerInfoBuilder) Build() *OutTravelerInfo {
	data := &OutTravelerInfo{}
	if builder.travelerIdSet {
		data.TravelerId = &builder.travelerId
	}
	if builder.outTravelerIdSet {
		data.OutTravelerId = &builder.outTravelerId
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
	if builder.birthDateSet {
		data.BirthDate = &builder.birthDate
	}
	if builder.cardListSet {
		data.CardList = builder.cardList
	}
	if builder.relatedEmployeesSet {
		data.RelatedEmployees = builder.relatedEmployees
	}
	if builder.projectIdSet {
		data.ProjectId = &builder.projectId
	}
	if builder.projectNameSet {
		data.ProjectName = &builder.projectName
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
	}
	return data
}

// CardInfo 证件信息
type CardInfo struct {
	CardNo     *string `json:"card_no,omitempty"`     // 证件号码
	CardType   *int32  `json:"card_type,omitempty"`   // 证件类型编码
	ExpireDate *string `json:"expire_date,omitempty"` // 证件有效期
}

type CardInfoBuilder struct {
	cardNo        string
	cardNoSet     bool
	cardType      int32
	cardTypeSet   bool
	expireDate    string
	expireDateSet bool
}

func NewCardInfoBuilder() *CardInfoBuilder {
	return &CardInfoBuilder{}
}
func (builder *CardInfoBuilder) CardNo(cardNo string) *CardInfoBuilder {
	builder.cardNo = cardNo
	builder.cardNoSet = true
	return builder
}
func (builder *CardInfoBuilder) CardType(cardType int32) *CardInfoBuilder {
	builder.cardType = cardType
	builder.cardTypeSet = true
	return builder
}
func (builder *CardInfoBuilder) ExpireDate(expireDate string) *CardInfoBuilder {
	builder.expireDate = expireDate
	builder.expireDateSet = true
	return builder
}

func (builder *CardInfoBuilder) Build() *CardInfo {
	data := &CardInfo{}
	if builder.cardNoSet {
		data.CardNo = &builder.cardNo
	}
	if builder.cardTypeSet {
		data.CardType = &builder.cardType
	}
	if builder.expireDateSet {
		data.ExpireDate = &builder.expireDate
	}
	return data
}

// RelatedEmployeeInfo 关联员工信息
type RelatedEmployeeInfo struct {
	RelatedEmployeeId             *string `json:"related_employee_id,omitempty"`              // 关联员工滴滴ID
	RelatedEmployeePhone          *string `json:"related_employee_phone,omitempty"`           // 关联员工手机号
	RelatedEmployeeEmployeeNumber *string `json:"related_employee_employee_number,omitempty"` // 关联员工工号
	RelatedEmployeeEmail          *string `json:"related_employee_email,omitempty"`           // 关联员工邮箱
}

type RelatedEmployeeInfoBuilder struct {
	relatedEmployeeId                string
	relatedEmployeeIdSet             bool
	relatedEmployeePhone             string
	relatedEmployeePhoneSet          bool
	relatedEmployeeEmployeeNumber    string
	relatedEmployeeEmployeeNumberSet bool
	relatedEmployeeEmail             string
	relatedEmployeeEmailSet          bool
}

func NewRelatedEmployeeInfoBuilder() *RelatedEmployeeInfoBuilder {
	return &RelatedEmployeeInfoBuilder{}
}
func (builder *RelatedEmployeeInfoBuilder) RelatedEmployeeId(relatedEmployeeId string) *RelatedEmployeeInfoBuilder {
	builder.relatedEmployeeId = relatedEmployeeId
	builder.relatedEmployeeIdSet = true
	return builder
}
func (builder *RelatedEmployeeInfoBuilder) RelatedEmployeePhone(relatedEmployeePhone string) *RelatedEmployeeInfoBuilder {
	builder.relatedEmployeePhone = relatedEmployeePhone
	builder.relatedEmployeePhoneSet = true
	return builder
}
func (builder *RelatedEmployeeInfoBuilder) RelatedEmployeeEmployeeNumber(relatedEmployeeEmployeeNumber string) *RelatedEmployeeInfoBuilder {
	builder.relatedEmployeeEmployeeNumber = relatedEmployeeEmployeeNumber
	builder.relatedEmployeeEmployeeNumberSet = true
	return builder
}
func (builder *RelatedEmployeeInfoBuilder) RelatedEmployeeEmail(relatedEmployeeEmail string) *RelatedEmployeeInfoBuilder {
	builder.relatedEmployeeEmail = relatedEmployeeEmail
	builder.relatedEmployeeEmailSet = true
	return builder
}

func (builder *RelatedEmployeeInfoBuilder) Build() *RelatedEmployeeInfo {
	data := &RelatedEmployeeInfo{}
	if builder.relatedEmployeeIdSet {
		data.RelatedEmployeeId = &builder.relatedEmployeeId
	}
	if builder.relatedEmployeePhoneSet {
		data.RelatedEmployeePhone = &builder.relatedEmployeePhone
	}
	if builder.relatedEmployeeEmployeeNumberSet {
		data.RelatedEmployeeEmployeeNumber = &builder.relatedEmployeeEmployeeNumber
	}
	if builder.relatedEmployeeEmailSet {
		data.RelatedEmployeeEmail = &builder.relatedEmployeeEmail
	}
	return data
}
