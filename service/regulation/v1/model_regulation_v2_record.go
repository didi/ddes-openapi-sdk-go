package v1

// RegulationV2Record 制度记录（/open-apis/v2/regulation/get 响应 records 元素，契约依据 idl/regulation.idl RegulationInfo）
type RegulationV2Record struct {
	RegulationId                  *string `json:"regulation_id,omitempty"`                   // 制度id
	RegulationName                *string `json:"regulation_name,omitempty"`                 // 制度名称
	RegulationEmployeeName        *string `json:"regulation_employee_name,omitempty"`        // 制度员工侧名称
	RegulationEmployeeDescription *string `json:"regulation_employee_description,omitempty"` // 制度员工侧描述
	RegulationStatus              *string `json:"regulation_status,omitempty"`               // 制度状态，（0-停用 1-正常 2-删除 3-过期）
	IsApprove                     *int32  `json:"is_approve,omitempty"`                      // 制度是否需要审批，(0-无需审批，1-行前审批，2-差旅)
	SceneType                     *string `json:"scene_type,omitempty"`                      // 因公出行场景
	IsUseQuota                    *int32  `json:"is_use_quota,omitempty"`                    // 是否使用个人限额，（0-不使用、1-使用）
	Source                        *int32  `json:"source,omitempty"`                          // 来源（1-通用规则 2-行前审批 3-差旅 4-无需审批）
	CityType                      *int32  `json:"city_type,omitempty"`                       // 管控方式
	ApprovalType                  *int32  `json:"approval_type,omitempty"`                   // 审批类型，(0-无需审批 1-差旅 2-行前审批按次数 3-行前审批按日期)
	BelongEnterpriseName          *string `json:"belong_enterprise_name,omitempty"`          // 归属企业名称（集团账户）
}

type RegulationV2RecordBuilder struct {
	regulationId                     string
	regulationIdSet                  bool
	regulationName                   string
	regulationNameSet                bool
	regulationEmployeeName           string
	regulationEmployeeNameSet        bool
	regulationEmployeeDescription    string
	regulationEmployeeDescriptionSet bool
	regulationStatus                 string
	regulationStatusSet              bool
	isApprove                        int32
	isApproveSet                     bool
	sceneType                        string
	sceneTypeSet                     bool
	isUseQuota                       int32
	isUseQuotaSet                    bool
	source                           int32
	sourceSet                        bool
	cityType                         int32
	cityTypeSet                      bool
	approvalType                     int32
	approvalTypeSet                  bool
	belongEnterpriseName             string
	belongEnterpriseNameSet          bool
}

func NewRegulationV2RecordBuilder() *RegulationV2RecordBuilder {
	return &RegulationV2RecordBuilder{}
}
func (builder *RegulationV2RecordBuilder) RegulationId(regulationId string) *RegulationV2RecordBuilder {
	builder.regulationId = regulationId
	builder.regulationIdSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) RegulationName(regulationName string) *RegulationV2RecordBuilder {
	builder.regulationName = regulationName
	builder.regulationNameSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) RegulationEmployeeName(regulationEmployeeName string) *RegulationV2RecordBuilder {
	builder.regulationEmployeeName = regulationEmployeeName
	builder.regulationEmployeeNameSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) RegulationEmployeeDescription(regulationEmployeeDescription string) *RegulationV2RecordBuilder {
	builder.regulationEmployeeDescription = regulationEmployeeDescription
	builder.regulationEmployeeDescriptionSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) RegulationStatus(regulationStatus string) *RegulationV2RecordBuilder {
	builder.regulationStatus = regulationStatus
	builder.regulationStatusSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) IsApprove(isApprove int32) *RegulationV2RecordBuilder {
	builder.isApprove = isApprove
	builder.isApproveSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) SceneType(sceneType string) *RegulationV2RecordBuilder {
	builder.sceneType = sceneType
	builder.sceneTypeSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) IsUseQuota(isUseQuota int32) *RegulationV2RecordBuilder {
	builder.isUseQuota = isUseQuota
	builder.isUseQuotaSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) Source(source int32) *RegulationV2RecordBuilder {
	builder.source = source
	builder.sourceSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) CityType(cityType int32) *RegulationV2RecordBuilder {
	builder.cityType = cityType
	builder.cityTypeSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) ApprovalType(approvalType int32) *RegulationV2RecordBuilder {
	builder.approvalType = approvalType
	builder.approvalTypeSet = true
	return builder
}
func (builder *RegulationV2RecordBuilder) BelongEnterpriseName(belongEnterpriseName string) *RegulationV2RecordBuilder {
	builder.belongEnterpriseName = belongEnterpriseName
	builder.belongEnterpriseNameSet = true
	return builder
}

func (builder *RegulationV2RecordBuilder) Build() *RegulationV2Record {
	data := &RegulationV2Record{}
	if builder.regulationIdSet {
		data.RegulationId = &builder.regulationId
	}
	if builder.regulationNameSet {
		data.RegulationName = &builder.regulationName
	}
	if builder.regulationEmployeeNameSet {
		data.RegulationEmployeeName = &builder.regulationEmployeeName
	}
	if builder.regulationEmployeeDescriptionSet {
		data.RegulationEmployeeDescription = &builder.regulationEmployeeDescription
	}
	if builder.regulationStatusSet {
		data.RegulationStatus = &builder.regulationStatus
	}
	if builder.isApproveSet {
		data.IsApprove = &builder.isApprove
	}
	if builder.sceneTypeSet {
		data.SceneType = &builder.sceneType
	}
	if builder.isUseQuotaSet {
		data.IsUseQuota = &builder.isUseQuota
	}
	if builder.sourceSet {
		data.Source = &builder.source
	}
	if builder.cityTypeSet {
		data.CityType = &builder.cityType
	}
	if builder.approvalTypeSet {
		data.ApprovalType = &builder.approvalType
	}
	if builder.belongEnterpriseNameSet {
		data.BelongEnterpriseName = &builder.belongEnterpriseName
	}
	return data
}
