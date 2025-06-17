package v1

// ExtendInfo struct for ExtendInfo
type ExtendInfo struct {
	Id       *string      `json:"id,omitempty"`       // 子档案唯一标识
	Status   *int32       `json:"status,omitempty"`   // 子档案状态，1正常，停用的还有返回么？
	Code     *string      `json:"code,omitempty"`     // 子档案code
	Name     *string      `json:"name,omitempty"`     // 子档案名称
	Children []ExtendInfo `json:"children,omitempty"` // 下一级子档案列表
}

type ExtendInfoBuilder struct {
	id          string // 子档案唯一标识
	idSet       bool
	status      int32 // 子档案状态，1正常，停用的还有返回么？
	statusSet   bool
	code        string // 子档案code
	codeSet     bool
	name        string // 子档案名称
	nameSet     bool
	children    []ExtendInfo // 下一级子档案列表
	childrenSet bool
}

func NewExtendInfoBuilder() *ExtendInfoBuilder {
	return &ExtendInfoBuilder{}
}
func (builder *ExtendInfoBuilder) Id(id string) *ExtendInfoBuilder {
	builder.id = id
	builder.idSet = true
	return builder
}
func (builder *ExtendInfoBuilder) Status(status int32) *ExtendInfoBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}
func (builder *ExtendInfoBuilder) Code(code string) *ExtendInfoBuilder {
	builder.code = code
	builder.codeSet = true
	return builder
}
func (builder *ExtendInfoBuilder) Name(name string) *ExtendInfoBuilder {
	builder.name = name
	builder.nameSet = true
	return builder
}
func (builder *ExtendInfoBuilder) Children(children []ExtendInfo) *ExtendInfoBuilder {
	builder.children = children
	builder.childrenSet = true
	return builder
}

func (builder *ExtendInfoBuilder) Build() *ExtendInfo {
	data := &ExtendInfo{}
	if builder.idSet {
		data.Id = &builder.id
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	if builder.codeSet {
		data.Code = &builder.code
	}
	if builder.nameSet {
		data.Name = &builder.name
	}
	if builder.childrenSet {
		data.Children = builder.children
	}
	return data
}
