package v1

// GetMemberQuotaReply 【提醒】此接口返回值涉及map中存放数组对象，已在rpc定义处重写响应值！！！
type GetMemberQuotaReply struct {
	MemberId       *int64   `json:"member_id,omitempty"`        // 员工在滴滴企业平台的ID
	QuotaStartDate *string  `json:"quota_start_date,omitempty"` // 限额开始日期，月限额时为自然月开始日 格式为 yyyy-MM-dd
	QuotaEndDate   *string  `json:"quota_end_date,omitempty"`   // 限额结束日期，月限额时为自然月结束日 格式为 yyyy-MM-dd
	QuotaCycle     *int32   `json:"quota_cycle,omitempty"`      // 限额类型，枚举值数字 0 不限期 1 自然月
	TotalQuota     *float32 `json:"total_quota,omitempty"`      // 总限额，单位：元
	AvailableQuota *float32 `json:"available_quota,omitempty"`  // 可用额度，单位：元
}
