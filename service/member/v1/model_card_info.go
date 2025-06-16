package v1

// CardInfo 证件信息
type CardInfo struct {
	CardType   *string `json:"card_type,omitempty"`   // 证件类型，枚举值数字 1. 身份证，2. 护照，3. 港澳台居民居住证，4. 台胞证，5. 军官证，6. 回乡证，7. 外国人永久居留身份证
	CardNo     *string `json:"card_no,omitempty"`     // 证件号码（已用AES算法加密）
	ExpireDate *string `json:"expire_date,omitempty"` // 证件过期日期，格式：2050-01-01（已用AES算法加密）
}

type CardInfoBuilder struct {
	cardType      string // 证件类型，枚举值数字 1. 身份证，2. 护照，3. 港澳台居民居住证，4. 台胞证，5. 军官证，6. 回乡证，7. 外国人永久居留身份证
	cardTypeSet   bool
	cardNo        string // 证件号码（已用AES算法加密）
	cardNoSet     bool
	expireDate    string // 证件过期日期，格式：2050-01-01（已用AES算法加密）
	expireDateSet bool
}

func NewCardInfoBuilder() *CardInfoBuilder {
	return &CardInfoBuilder{}
}
func (builder *CardInfoBuilder) CardType(cardType string) *CardInfoBuilder {
	builder.cardType = cardType
	builder.cardTypeSet = true
	return builder
}
func (builder *CardInfoBuilder) CardNo(cardNo string) *CardInfoBuilder {
	builder.cardNo = cardNo
	builder.cardNoSet = true
	return builder
}
func (builder *CardInfoBuilder) ExpireDate(expireDate string) *CardInfoBuilder {
	builder.expireDate = expireDate
	builder.expireDateSet = true
	return builder
}

func (builder *CardInfoBuilder) Build() *CardInfo {
	data := &CardInfo{}
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
