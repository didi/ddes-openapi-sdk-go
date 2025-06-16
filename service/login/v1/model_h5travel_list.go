package v1

// H5TravelList 差旅单/审批单详情
type H5TravelList struct {
	ApprovalId *string `json:"approval_id,omitempty"` // 差旅/审批单id
	JumpPage   *string `json:"jumpPage,omitempty"`    // 跳转地址为travelList
}

type H5TravelListBuilder struct {
	approvalId    string // 差旅/审批单id
	approvalIdSet bool
	jumpPage      string // 跳转地址为travelList
	jumpPageSet   bool
}

func NewH5TravelListBuilder() *H5TravelListBuilder {
	return &H5TravelListBuilder{}
}
func (builder *H5TravelListBuilder) ApprovalId(approvalId string) *H5TravelListBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}
func (builder *H5TravelListBuilder) JumpPage(jumpPage string) *H5TravelListBuilder {
	builder.jumpPage = jumpPage
	builder.jumpPageSet = true
	return builder
}

func (builder *H5TravelListBuilder) Build() *H5TravelList {
	data := &H5TravelList{}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	if builder.jumpPageSet {
		data.JumpPage = &builder.jumpPage
	}
	return data
}
