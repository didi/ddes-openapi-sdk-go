package v1

// LegalEntityRecord 公司主体
type LegalEntityRecord struct {
	Address          *string `json:"address,omitempty"`             // 公司主体注册地址
	BankCardNo       *string `json:"bank_card_no,omitempty"`        // 开户行账号
	EndTime          *string `json:"end_time,omitempty"`            // 公司主体有效期截止日期
	LegalEntityId    *int64  `json:"legal_entity_id,omitempty"`     // 公司主体 id
	Name             *string `json:"name,omitempty"`                // 公司主体名称
	OpenBank         *string `json:"open_bank,omitempty"`           // 开户行名称
	OutLegalEntityId *string `json:"out_legal_entity_id,omitempty"` // 外部主体编号
	ParentId         *int64  `json:"parent_id,omitempty"`           // 上级公司 id
	StartTime        *string `json:"start_time,omitempty"`          // 公司主体有效期开始日期
	TaxpayerNo       *string `json:"taxpayer_no,omitempty"`         // 纳税人识别号
	Telephone        *string `json:"telephone,omitempty"`           // 公司主体注册电话
	Status           *int32  `json:"status,omitempty"`              // 公司状态，0 停用，1 启用
}

type LegalEntityRecordBuilder struct {
	address             string // 公司主体注册地址
	addressSet          bool
	bankCardNo          string // 开户行账号
	bankCardNoSet       bool
	endTime             string // 公司主体有效期截止日期
	endTimeSet          bool
	legalEntityId       int64 // 公司主体 id
	legalEntityIdSet    bool
	name                string // 公司主体名称
	nameSet             bool
	openBank            string // 开户行名称
	openBankSet         bool
	outLegalEntityId    string // 外部主体编号
	outLegalEntityIdSet bool
	parentId            int64 // 上级公司 id
	parentIdSet         bool
	startTime           string // 公司主体有效期开始日期
	startTimeSet        bool
	taxpayerNo          string // 纳税人识别号
	taxpayerNoSet       bool
	telephone           string // 公司主体注册电话
	telephoneSet        bool
	status              int32 // 公司状态，0 停用，1 启用
	statusSet           bool
}

func NewLegalEntityRecordBuilder() *LegalEntityRecordBuilder {
	return &LegalEntityRecordBuilder{}
}
func (builder *LegalEntityRecordBuilder) Address(address string) *LegalEntityRecordBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) BankCardNo(bankCardNo string) *LegalEntityRecordBuilder {
	builder.bankCardNo = bankCardNo
	builder.bankCardNoSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) EndTime(endTime string) *LegalEntityRecordBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) LegalEntityId(legalEntityId int64) *LegalEntityRecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) Name(name string) *LegalEntityRecordBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) OpenBank(openBank string) *LegalEntityRecordBuilder {
	builder.openBank = openBank
	builder.openBankSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) OutLegalEntityId(outLegalEntityId string) *LegalEntityRecordBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) ParentId(parentId int64) *LegalEntityRecordBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) StartTime(startTime string) *LegalEntityRecordBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) TaxpayerNo(taxpayerNo string) *LegalEntityRecordBuilder {
	builder.taxpayerNo = taxpayerNo
	builder.taxpayerNoSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) Telephone(telephone string) *LegalEntityRecordBuilder {
	builder.telephone = telephone
	builder.telephoneSet = true
	return builder
}
func (builder *LegalEntityRecordBuilder) Status(status int32) *LegalEntityRecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}

func (builder *LegalEntityRecordBuilder) Build() *LegalEntityRecord {
	data := &LegalEntityRecord{}
	if builder.addressSet {
		data.Address = &builder.address
	}
	if builder.bankCardNoSet {
		data.BankCardNo = &builder.bankCardNo
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.openBankSet {
		data.OpenBank = &builder.openBank
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.taxpayerNoSet {
		data.TaxpayerNo = &builder.taxpayerNo
	}
	if builder.telephoneSet {
		data.Telephone = &builder.telephone
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	return data
}
