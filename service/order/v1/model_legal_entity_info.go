package v1

// LegalEntityInfo 所属公司信息
type LegalEntityInfo struct {
	LegalEntityId    *string `json:"legal_entity_id,omitempty"`     // 叫车人所属公司id
	OutLegalEntityId *string `json:"out_legal_entity_id,omitempty"` // 叫车人所属公司业务编码
	LegalEntityName  *string `json:"legal_entity_name,omitempty"`   // 叫车人所属公司名称
}

type LegalEntityInfoBuilder struct {
	legalEntityId       string // 叫车人所属公司id
	legalEntityIdSet    bool
	outLegalEntityId    string // 叫车人所属公司业务编码
	outLegalEntityIdSet bool
	legalEntityName     string // 叫车人所属公司名称
	legalEntityNameSet  bool
}

func NewLegalEntityInfoBuilder() *LegalEntityInfoBuilder {
	return &LegalEntityInfoBuilder{}
}
func (builder *LegalEntityInfoBuilder) LegalEntityId(legalEntityId string) *LegalEntityInfoBuilder {
	builder.legalEntityId = legalEntityId
	builder.legalEntityIdSet = true
	return builder
}
func (builder *LegalEntityInfoBuilder) OutLegalEntityId(outLegalEntityId string) *LegalEntityInfoBuilder {
	builder.outLegalEntityId = outLegalEntityId
	builder.outLegalEntityIdSet = true
	return builder
}
func (builder *LegalEntityInfoBuilder) LegalEntityName(legalEntityName string) *LegalEntityInfoBuilder {
	builder.legalEntityName = legalEntityName
	builder.legalEntityNameSet = true
	return builder
}

func (builder *LegalEntityInfoBuilder) Build() *LegalEntityInfo {
	data := &LegalEntityInfo{}
	if builder.legalEntityIdSet {
		data.LegalEntityId = &builder.legalEntityId
	}
	if builder.outLegalEntityIdSet {
		data.OutLegalEntityId = &builder.outLegalEntityId
	}
	if builder.legalEntityNameSet {
		data.LegalEntityName = &builder.legalEntityName
	}
	return data
}
