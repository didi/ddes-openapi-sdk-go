package v1

// GetLegalEntityReply 定义包含公司主体记录列表的消息
type GetLegalEntityReply struct {
	Total   *int32              `json:"total,omitempty"`   // 总数据条数
	Records []LegalEntityRecord `json:"records,omitempty"` // 公司列表
}
