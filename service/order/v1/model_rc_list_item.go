package v1

// RcListItem 管控信息列表项
type RcListItem struct {
	RcInfo *RcInfo `json:"rc_info,omitempty"`
}

type RcListItemBuilder struct {
	rcInfo    RcInfo
	rcInfoSet bool
}

func NewRcListItemBuilder() *RcListItemBuilder {
	return &RcListItemBuilder{}
}
func (builder *RcListItemBuilder) RcInfo(rcInfo RcInfo) *RcListItemBuilder {
	builder.rcInfo = rcInfo
	builder.rcInfoSet = true
	return builder
}

func (builder *RcListItemBuilder) Build() *RcListItem {
	data := &RcListItem{}
	if builder.rcInfoSet {
		data.RcInfo = &builder.rcInfo
	}
	return data
}
