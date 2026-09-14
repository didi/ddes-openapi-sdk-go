package v1

// MemberV2CardInfo 集团账户 v2 员工证件信息（v2 契约独立类型，禁止与 /river 版 CardInfo 复用）
type MemberV2CardInfo struct {
	CardType   *string `json:"card_type,omitempty"`   // 证件类型，枚举值数字 1. 身份证，2. 护照，3. 港澳台居民居住证，4. 台胞证，5. 军官证，6. 回乡证，7. 外国人永久居留身份证
	CardNo     *string `json:"card_no,omitempty"`     // 证件号码（已用AES算法加密）
	ExpireDate *string `json:"expire_date,omitempty"` // 证件过期日期（已用AES算法加密）
}

type MemberV2CardInfoBuilder struct {
	cardType      string // 证件类型，枚举值数字 1. 身份证，2. 护照，3. 港澳台居民居住证，4. 台胞证，5. 军官证，6. 回乡证，7. 外国人永久居留身份证
	cardTypeSet   bool
	cardNo        string // 证件号码（已用AES算法加密）
	cardNoSet     bool
	expireDate    string // 证件过期日期（已用AES算法加密）
	expireDateSet bool
}

func NewMemberV2CardInfoBuilder() *MemberV2CardInfoBuilder {
	return &MemberV2CardInfoBuilder{}
}
func (builder *MemberV2CardInfoBuilder) CardType(cardType string) *MemberV2CardInfoBuilder {
	builder.cardType = cardType
	builder.cardTypeSet = true
	return builder
}
func (builder *MemberV2CardInfoBuilder) CardNo(cardNo string) *MemberV2CardInfoBuilder {
	builder.cardNo = cardNo
	builder.cardNoSet = true
	return builder
}
func (builder *MemberV2CardInfoBuilder) ExpireDate(expireDate string) *MemberV2CardInfoBuilder {
	builder.expireDate = expireDate
	builder.expireDateSet = true
	return builder
}

func (builder *MemberV2CardInfoBuilder) Build() *MemberV2CardInfo {
	data := &MemberV2CardInfo{}
	if builder.cardTypeSet {
		data.CardType = &builder.cardType
	}
	if builder.cardNoSet {
		data.CardNo = &builder.cardNo
	}
	if builder.expireDateSet {
		data.ExpireDate = &builder.expireDate
	}
	return data
}
