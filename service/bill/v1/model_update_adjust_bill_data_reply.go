package v1

// UpdateAdjustBillDataReply struct for UpdateAdjustBillDataReply
type UpdateAdjustBillDataReply struct {
	AdjustReqId  *string          `json:"adjust_req_id,omitempty"` // 调账id(提交过来的),可用于后续查询
	CheckPass    *bool            `json:"check_pass,omitempty"`    // 是否接口请求提交有效, true为有效,false请参考错误信息重新提交
	AdjustStatus *int32           `json:"adjust_status,omitempty"` // 状态1.提交成功2.执行中3.执行成功-1.提交失败-2.执行失败
	ErrMsg       *string          `json:"err_msg,omitempty"`       // 没通过时,errmsg 简述
	ErrMsgList   []ErrMsgListItem `json:"err_msg_list,omitempty"`  // 没通过时,errMsg 详细
	CompanyId    *int64           `json:"company_id,omitempty"`    // 企业ID
	BusinessType *int32           `json:"business_type,omitempty"` // 业务类型
	AdjustType   *int32           `json:"adjust_type,omitempty"`   // 调账类型
	BillId       *int64           `json:"bill_id,omitempty"`       // 账单ID
}
