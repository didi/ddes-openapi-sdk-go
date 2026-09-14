package v1

import "github.com/didi/ddes-openapi-sdk-go/core"

// MemberV2ResidentsListInfo 集团账户 v2 员工常驻地信息（v2 契约独立类型，禁止与 /river 版 ResidentsListInfo 复用）
type MemberV2ResidentsListInfo struct {
	Id     *string `json:"id,omitempty"`     // 常驻地ID
	Name   *string `json:"name,omitempty"`   // 常驻地名称
	Adcode *string `json:"adcode,omitempty"` // 常驻地行政区划代码
}

type MemberV2ResidentsListInfoBuilder struct {
	id        string // 常驻地ID
	idSet     bool
	name      string // 常驻地名称
	nameSet   bool
	adcode    string // 常驻地行政区划代码
	adcodeSet bool
}

func NewMemberV2ResidentsListInfoBuilder() *MemberV2ResidentsListInfoBuilder {
	return &MemberV2ResidentsListInfoBuilder{}
}
func (builder *MemberV2ResidentsListInfoBuilder) Id(id string) *MemberV2ResidentsListInfoBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *MemberV2ResidentsListInfoBuilder) Name(name string) *MemberV2ResidentsListInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *MemberV2ResidentsListInfoBuilder) Adcode(adcode string) *MemberV2ResidentsListInfoBuilder {
	builder.adcode = adcode
	builder.adcodeSet = true
	return builder
}

func (builder *MemberV2ResidentsListInfoBuilder) Build() *MemberV2ResidentsListInfo {
	data := &MemberV2ResidentsListInfo{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.adcodeSet {
		data.Adcode = &builder.adcode
	}
	return data
}

// UnmarshalJSON 容错反序列化：Id 真实流量返回 number，*string 收 number 会失败。
// 用 core.SmartDecode 经中间 map 转换（number→string 容错），不经 json.Unmarshal
// 到自身以避免递归。
func (r *MemberV2ResidentsListInfo) UnmarshalJSON(data []byte) error {
	return core.SmartDecode(data, r)
}
