package v1

// H5BTripCategory 商旅品类预订
type H5BTripCategory struct {
	ApprovalId *string `json:"approval_id,omitempty"` // 差旅、申请用车，必传
}

type H5BTripCategoryBuilder struct {
	approvalId    string // 差旅、申请用车，必传
	approvalIdSet bool
}

func NewH5BTripCategoryBuilder() *H5BTripCategoryBuilder {
	return &H5BTripCategoryBuilder{}
}
func (builder *H5BTripCategoryBuilder) ApprovalId(approvalId string) *H5BTripCategoryBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}

func (builder *H5BTripCategoryBuilder) Build() *H5BTripCategory {
	data := &H5BTripCategory{}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	return data
}
