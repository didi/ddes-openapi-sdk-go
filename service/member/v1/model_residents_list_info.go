package v1

// ResidentsListInfo 常驻地信息
type ResidentsListInfo struct {
	Id     *string `json:"id,omitempty"`     // 常驻地ID
	Name   *string `json:"name,omitempty"`   // 常驻地名称
	Adcode *string `json:"adcode,omitempty"` // 常驻地行政区划代码
}

type ResidentsListInfoBuilder struct {
	id        string // 常驻地ID
	idSet     bool
	name      string // 常驻地名称
	nameSet   bool
	adcode    string // 常驻地行政区划代码
	adcodeSet bool
}

func NewResidentsListInfoBuilder() *ResidentsListInfoBuilder {
	return &ResidentsListInfoBuilder{}
}
func (builder *ResidentsListInfoBuilder) Id(id string) *ResidentsListInfoBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *ResidentsListInfoBuilder) Name(name string) *ResidentsListInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *ResidentsListInfoBuilder) Adcode(adcode string) *ResidentsListInfoBuilder {
	builder.adcode = adcode
	builder.adcodeSet = true
	return builder
}

func (builder *ResidentsListInfoBuilder) Build() *ResidentsListInfo {
	data := &ResidentsListInfo{}
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
