package v1

// PageInfo 分页信息消息
type PageInfo struct {
	CurPage *int32 `json:"cur_page,omitempty"` // 查询页
	Limit   *int32 `json:"limit,omitempty"`    // 页大小
	Total   *int32 `json:"total,omitempty"`    // 符合条件的数据总条数
}

type PageInfoBuilder struct {
	curPage    int32 // 查询页
	curPageSet bool
	limit      int32 // 页大小
	limitSet   bool
	total      int32 // 符合条件的数据总条数
	totalSet   bool
}

func NewPageInfoBuilder() *PageInfoBuilder {
	return &PageInfoBuilder{}
}
func (builder *PageInfoBuilder) CurPage(curPage int32) *PageInfoBuilder {
	builder.curPage = curPage
	builder.curPageSet = true
	return builder
}
func (builder *PageInfoBuilder) Limit(limit int32) *PageInfoBuilder {
	builder.limit = limit
	builder.limitSet = true
	return builder
}
func (builder *PageInfoBuilder) Total(total int32) *PageInfoBuilder {
	builder.total = total
	builder.totalSet = true
	return builder
}

func (builder *PageInfoBuilder) Build() *PageInfo {
	data := &PageInfo{}
	if builder.curPageSet {
		data.CurPage = &builder.curPage
	}
	if builder.limitSet {
		data.Limit = &builder.limit
	}
	if builder.totalSet {
		data.Total = &builder.total
	}
	return data
}
