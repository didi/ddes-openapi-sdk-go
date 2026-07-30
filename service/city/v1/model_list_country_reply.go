package v1

import "github.com/didi/ddes-openapi-sdk-go/core"

// ListCountryReply struct for ListCountryReply
type ListCountryReply struct {
	CountryId            *string `json:"country_id,omitempty"`             // 国家id（大ID，服务端返回值超 int32，用 *string）
	CanonicalCountryCode *string `json:"canonical_country_code,omitempty"` // 国家二次码
	CountryCode          *string `json:"country_code,omitempty"`           // 国家三字码
	CountryNameCn        *string `json:"country_name_cn,omitempty"`        // 国家中文名
	CountryNameEn        *string `json:"country_name_en,omitempty"`        // 国家英文名
	ContinentId          *int32  `json:"continent_id,omitempty"`           // 大洲id
	ContinentNameCn      *string `json:"continent_name_cn,omitempty"`      // 大洲中文名
	ContinentNameEn      *string `json:"continent_name_en,omitempty"`      // 大洲英文名
}

// UnmarshalJSON 容错反序列化：CountryId 真实流量返回 number，*string 收 number 会失败。
// 用 core.SmartDecode 经中间 map 转换，不经 json.Unmarshal 到自身以避免递归。
func (l *ListCountryReply) UnmarshalJSON(data []byte) error {
	return core.SmartDecode(data, l)
}
