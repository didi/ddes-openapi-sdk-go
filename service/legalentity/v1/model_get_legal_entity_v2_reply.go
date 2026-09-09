package v1

// GetLegalEntityV2ApiReply /open-apis/v2/legalEntity/get 响应外层
type GetLegalEntityV2ApiReply struct {
	Errno     *int32                 `json:"errno,omitempty"`      // 错误码
	Errmsg    *string                `json:"errmsg,omitempty"`     // 错误文案
	RequestId *string                `json:"request_id,omitempty"` // 请求ID(该字段一定要保留，方便排查问题)
	Data      *GetLegalEntityV2Reply `json:"data,omitempty"`       // 数据对象
}

// GetLegalEntityV2Reply 跨公司分页数据（next_token 为空说明没有下一页）
type GetLegalEntityV2Reply struct {
	NextToken            *string               `json:"next_token,omitempty"`              // 下一页token，空说明没有下一页
	Records              []LegalEntityV2Record `json:"records,omitempty"`                 // 公司主体列表
	Progress             *float64              `json:"progress,omitempty"`                // 查询进度
	CompanyTotal         *int32                `json:"company_total,omitempty"`           // 涉及的公司总数
	CompanyListHasChange *bool                 `json:"company_list_has_change,omitempty"` // 翻页期间公司列表是否发生变化
}
