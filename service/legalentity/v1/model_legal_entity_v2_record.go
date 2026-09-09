package v1

// LegalEntityV2Record 公司主体记录（/open-apis/v2/legalEntity/get 响应 records 元素，契约依据 idl/legal-entity.idl LegalEntityInfo）
type LegalEntityV2Record struct {
	LegalEntityId        *string `json:"legal_entity_id,omitempty"`        // 公司主体滴滴侧id（v2 契约为 string，超大数值不丢精度）
	Name                 *string `json:"name,omitempty"`                   // 公司主体名称
	OutLegalEntityId     *string `json:"out_legal_entity_id,omitempty"`    // 外部公司主体编号
	Status               *int32  `json:"status,omitempty"`                 // 公司状态，0 停用，1 启用，2 过期
	ParentId             *int64  `json:"parent_id,omitempty"`              // 上级公司滴滴侧公司主体id
	StartTime            *string `json:"start_time,omitempty"`             // 公司主体有效期开始日期
	EndTime              *string `json:"end_time,omitempty"`               // 公司主体有效期截止日期
	TaxpayerNo           *string `json:"taxpayer_no,omitempty"`            // 纳税人识别号
	Address              *string `json:"address,omitempty"`                // 公司主体注册地址
	Telephone            *string `json:"telephone,omitempty"`              // 公司主体注册电话
	OpenBank             *string `json:"open_bank,omitempty"`              // 开户行名称
	BankCardNo           *string `json:"bank_card_no,omitempty"`           // 开户行账号
	OutParentId          *string `json:"out_parent_id,omitempty"`          // 上级公司外部编号
	BelongEnterpriseName *string `json:"belong_enterprise_name,omitempty"` // 归属企业名称（集团账户）
}

type LegalEntityV2RecordBuilder struct {
	legalEntityId           string
	legalEntityIdSet        bool
	name                    string
	nameSet                 bool
	outLegalEntityId        string
	outLegalEntityIdSet     bool
	status                  int32
	statusSet               bool
	parentId                int64
	parentIdSet             bool
	startTime               string
	startTimeSet            bool
	endTime                 string
	endTimeSet              bool
	taxpayerNo              string
	taxpayerNoSet           bool
	address                 string
	addressSet              bool
	telephone               string
	telephoneSet            bool
	openBank                string
	openBankSet             bool
	bankCardNo              string
	bankCardNoSet           bool
	outParentId             string
	outParentIdSet          bool
	belongEnterpriseName    string
	belongEnterpriseNameSet bool
}

func NewLegalEntityV2RecordBuilder() *LegalEntityV2RecordBuilder {
	return &LegalEntityV2RecordBuilder{}
}
func (builder *LegalEntityV2RecordBuilder) LegalEntityId(legalEntityId string) *LegalEntityV2RecordBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) Name(name string) *LegalEntityV2RecordBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) OutLegalEntityId(outLegalEntityId string) *LegalEntityV2RecordBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) Status(status int32) *LegalEntityV2RecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) ParentId(parentId int64) *LegalEntityV2RecordBuilder {
	builder.parentId = parentId
	builder.parentIdSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) StartTime(startTime string) *LegalEntityV2RecordBuilder {
	builder.startTime = startTime
	builder.startTimeSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) EndTime(endTime string) *LegalEntityV2RecordBuilder {
	builder.endTime = endTime
	builder.endTimeSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) TaxpayerNo(taxpayerNo string) *LegalEntityV2RecordBuilder {
	builder.taxpayerNo = taxpayerNo
	builder.taxpayerNoSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) Address(address string) *LegalEntityV2RecordBuilder {
	builder.address = address
	builder.addressSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) Telephone(telephone string) *LegalEntityV2RecordBuilder {
	builder.telephone = telephone
	builder.telephoneSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) OpenBank(openBank string) *LegalEntityV2RecordBuilder {
	builder.openBank = openBank
	builder.openBankSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) BankCardNo(bankCardNo string) *LegalEntityV2RecordBuilder {
	builder.bankCardNo = bankCardNo
	builder.bankCardNoSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) OutParentId(outParentId string) *LegalEntityV2RecordBuilder {
	builder.outParentId = outParentId
	builder.outParentIdSet = true
	return builder
}
func (builder *LegalEntityV2RecordBuilder) BelongEnterpriseName(belongEnterpriseName string) *LegalEntityV2RecordBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}

func (builder *LegalEntityV2RecordBuilder) Build() *LegalEntityV2Record {
	data := &LegalEntityV2Record{}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.parentIdSet {
		data.ParentId = &builder.parentId
	}
	if builder.startTimeSet {
		data.StartTime = &builder.startTime
	}
	if builder.endTimeSet {
		data.EndTime = &builder.endTime
	}
	if builder.taxpayerNoSet {
		data.TaxpayerNo = &builder.taxpayerNo
	}
	if builder.addressSet {
		data.Address = &builder.address
	}
	if builder.telephoneSet {
		data.Telephone = &builder.telephone
	}
	if builder.openBankSet {
		data.OpenBank = &builder.openBank
	}
	if builder.bankCardNoSet {
		data.BankCardNo = &builder.bankCardNo
	}
	if builder.outParentIdSet {
		data.OutParentId = &builder.outParentId
	}
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	return data
}
