package v1

// GetBillSummaryReply struct for GetBillSummaryReply
type GetBillSummaryReply struct {
	BillId                    *string             `json:"bill_id,omitempty"`                      // 结算单号
	CompanyName               *string             `json:"company_name,omitempty"`                 // 企业名称
	BillPeriod                *string             `json:"bill_period,omitempty"`                  // 账期
	SettleTypeName            *string             `json:"settle_type_name,omitempty"`             // 品类(网约车、代驾)
	ConsumeAmountOnline       *float32            `json:"consume_amount_online,omitempty"`        // 线上企业实付金额
	ConsumeAmountOffline      *float32            `json:"consume_amount_offline,omitempty"`       // 线下扣企业账户金额
	RefundAmount              *float32            `json:"refund_amount,omitempty"`                // 企业实退金额
	OutDateRefundAmount       *float32            `json:"out_date_refund_amount,omitempty"`       // 订单超期退款金额
	AmountMoney               *float32            `json:"amount_money,omitempty"`                 // 授信客户：本期应结算金额 预存客户：充值金额
	CurrentBalance            *float32            `json:"current_balance,omitempty"`              // 期末余额
	SubList                   []SubListItem       `json:"sub_list,omitempty"`                     // 子账户汇总 【4.0账单返回】
	BusinessType              *int32              `json:"business_type,omitempty"`                // 业务线 2:商旅
	ConsumeTotalAmount        *string             `json:"consume_total_amount,omitempty"`         // 消耗金额（单位：元）
	RefundTotalAmount         *string             `json:"refund_total_amount,omitempty"`          // 退款金额（单位：元）
	PreviousRefundTotalAmount *string             `json:"previous_refund_total_amount,omitempty"` // 往期退款金额（单位：元）
	AmountTotalMoney          *string             `json:"amount_total_money,omitempty"`           // 结算金额（单位：元）
	SubBillSummary            []SubBillSummaryDTO `json:"sub_bill_summary,omitempty"`             // 子业务线汇总金额 SubBillSummaryDTO
}
