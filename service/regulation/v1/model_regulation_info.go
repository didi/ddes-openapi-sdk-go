package v1

// RegulationInfo struct for RegulationInfo
type RegulationInfo struct {
	RegulationId     *string `json:"regulation_id,omitempty"`     // 制度id
	RegulationName   *string `json:"regulation_name,omitempty"`   // 制度名称
	RegulationStatus *string `json:"regulation_status,omitempty"` // 制度状态，（0-停用 1-正常 2-删除 3-过期）
	Source           *int32  `json:"source,omitempty"`            // 来源，来源（1-通用规则 2-行前审批 3-差旅 4-无需审批）
	SceneType        *string `json:"scene_type,omitempty"`        // 因公出行场景，（0-个人用车、1-商务出行、2-需审批差旅制度、3-加班、4-办公地点通勤、5-疫期通勤、11-会议用车、91-代叫车、92-接送机、93-包车、94-代驾、95-私车同行、96-行前审批、104-无需审批差旅制度）
	IsUseQuota       *int32  `json:"is_use_quota,omitempty"`      // 是否使用个人限额，（0-不使用、1-使用），差旅制度默认返回不使用
	IsApprove        *int32  `json:"is_approve,omitempty"`        // 制度是否需要审批，(0-无需审批，1-行前审批，2-差旅)
	InstitutionId    *string `json:"institution_id,omitempty"`    // 制度模板ID
	CityType         *int32  `json:"city_type,omitempty"`         // 管控方式（0-不管控、1-中度管控、2-严格管控、3-交叉预定（轻度管控）、4-会议管控）。默认值-1无意义
	ApprovalType     *int32  `json:"approval_type,omitempty"`     // 审批类型，(0-无需审批 1-差旅 2-行前审批按次数 3-行前审批按日期)
}

type RegulationInfoBuilder struct {
	regulationId        string // 制度id
	regulationIdSet     bool
	regulationName      string // 制度名称
	regulationNameSet   bool
	regulationStatus    string // 制度状态，（0-停用 1-正常 2-删除 3-过期）
	regulationStatusSet bool
	source              int32 // 来源，来源（1-通用规则 2-行前审批 3-差旅 4-无需审批）
	sourceSet           bool
	sceneType           string // 因公出行场景，（0-个人用车、1-商务出行、2-需审批差旅制度、3-加班、4-办公地点通勤、5-疫期通勤、11-会议用车、91-代叫车、92-接送机、93-包车、94-代驾、95-私车同行、96-行前审批、104-无需审批差旅制度）
	sceneTypeSet        bool
	isUseQuota          int32 // 是否使用个人限额，（0-不使用、1-使用），差旅制度默认返回不使用
	isUseQuotaSet       bool
	isApprove           int32 // 制度是否需要审批，(0-无需审批，1-行前审批，2-差旅)
	isApproveSet        bool
	institutionId       string // 制度模板ID
	institutionIdSet    bool
	cityType            int32 // 管控方式（0-不管控、1-中度管控、2-严格管控、3-交叉预定（轻度管控）、4-会议管控）。默认值-1无意义
	cityTypeSet         bool
	approvalType        int32 // 审批类型，(0-无需审批 1-差旅 2-行前审批按次数 3-行前审批按日期)
	approvalTypeSet     bool
}

func NewRegulationInfoBuilder() *RegulationInfoBuilder {
	return &RegulationInfoBuilder{}
}
func (builder *RegulationInfoBuilder) RegulationId(regulationId string) *RegulationInfoBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *RegulationInfoBuilder) RegulationName(regulationName string) *RegulationInfoBuilder {
	builder.regulationName = regulationName
	builder.regulationNameSet = true
	return builder
}
func (builder *RegulationInfoBuilder) RegulationStatus(regulationStatus string) *RegulationInfoBuilder {
	builder.regulationStatus = regulationStatus
	builder.regulationStatusSet = true
	return builder
}
func (builder *RegulationInfoBuilder) Source(source int32) *RegulationInfoBuilder {
	builder.source = source
	builder.sourceSet = true
	return builder
}
func (builder *RegulationInfoBuilder) SceneType(sceneType string) *RegulationInfoBuilder {
	builder.sceneType = sceneType
	builder.sceneTypeSet = true
	return builder
}
func (builder *RegulationInfoBuilder) IsUseQuota(isUseQuota int32) *RegulationInfoBuilder {
	builder.isUseQuota = isUseQuota
	builder.isUseQuotaSet = true
	return builder
}
func (builder *RegulationInfoBuilder) IsApprove(isApprove int32) *RegulationInfoBuilder {
	builder.isApprove = isApprove
	builder.isApproveSet = true
	return builder
}
func (builder *RegulationInfoBuilder) InstitutionId(institutionId string) *RegulationInfoBuilder {
	builder.institutionId = institutionId
	builder.institutionIdSet = true
	return builder
}
func (builder *RegulationInfoBuilder) CityType(cityType int32) *RegulationInfoBuilder {
	builder.cityType = cityType
	builder.cityTypeSet = true
	return builder
}
func (builder *RegulationInfoBuilder) ApprovalType(approvalType int32) *RegulationInfoBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}

func (builder *RegulationInfoBuilder) Build() *RegulationInfo {
	data := &RegulationInfo{}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.regulationNameSet {
		data.RegulationName = &builder.regulationName
	}
	if builder.regulationStatusSet {
		data.RegulationStatus = &builder.regulationStatus
	}
	if builder.sourceSet {
		data.Source = &builder.source
	}
	if builder.sceneTypeSet {
		data.SceneType = &builder.sceneType
	}
	if builder.isUseQuotaSet {
		data.IsUseQuota = &builder.isUseQuota
	}
	if builder.isApproveSet {
		data.IsApprove = &builder.isApprove
	}
	if builder.institutionIdSet {
		data.InstitutionId = &builder.institutionId
	}
	if builder.cityTypeSet {
		data.CityType = &builder.cityType
	}
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	return data
}
