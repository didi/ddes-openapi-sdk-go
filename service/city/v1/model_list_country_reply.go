package v1

// ListCountryReply struct for ListCountryReply
type ListCountryReply struct {
	CountryId            *int32  `json:"country_id,omitempty"`             // 国家id
	CanonicalCountryCode *string `json:"canonical_country_code,omitempty"` // 国家二次码
	CountryCode          *string `json:"country_code,omitempty"`           // 国家三字码
	CountryNameCn        *string `json:"country_name_cn,omitempty"`        // 国家中文名
	CountryNameEn        *string `json:"country_name_en,omitempty"`        // 国家英文名
	ContinentId          *int32  `json:"continent_id,omitempty"`           // 大洲id
	ContinentNameCn      *string `json:"continent_name_cn,omitempty"`      // 大洲中文名
	ContinentNameEn      *string `json:"continent_name_en,omitempty"`      // 大洲英文名
}
