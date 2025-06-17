package v1

// BillRecord 网约车、商旅、出租车三个业务线字段已合并到同一个结构中
type BillRecord struct {
	BillId          *int64   `json:"bill_id,omitempty"`          // 结算单号
	BillPeriod      *string  `json:"bill_period,omitempty"`      // 结算周期
	BillStatus      *int32   `json:"bill_status,omitempty"`      // 当前状态，枚举值数字：1：待确认；3：已确认
	BillAmount      *float32 `json:"bill_amount,omitempty"`      // 应结算金额
	BillConsumption *float32 `json:"bill_consumption,omitempty"` // 消耗金额
	BillRefund      *float32 `json:"bill_refund,omitempty"`      // 退款金额
	BillDifference  *float32 `json:"bill_difference,omitempty"`  // 差异金额（出租车为调整金额）
	CompanyId       *string  `json:"companyId,omitempty"`        // 公司 ID
	CompanyName     *string  `json:"companyName,omitempty"`      // 公司名称
	ConfirmDate     *string  `json:"confirm_date,omitempty"`     // 确认日期
	StartDate       *string  `json:"start_date,omitempty"`       // 开始时间（商旅适用）
	EndDate         *string  `json:"end_date,omitempty"`         // 结束时间（商旅适用）
	GroupType       *int32   `json:"groupType,omitempty"`        // 分组类型
	RootBatchId     *int64   `json:"root_batch_id,omitempty"`    // 根结点批次（商旅适用）
	Status          *int32   `json:"status,omitempty"`           // 账单状态（商旅适用），0. 生成中 1. 生成成功（待确认） 2.生成失败 3.已确认
}

type BillRecordBuilder struct {
	billId             int64 // 结算单号
	billIdSet          bool
	billPeriod         string // 结算周期
	billPeriodSet      bool
	billStatus         int32 // 当前状态，枚举值数字：1：待确认；3：已确认
	billStatusSet      bool
	billAmount         float32 // 应结算金额
	billAmountSet      bool
	billConsumption    float32 // 消耗金额
	billConsumptionSet bool
	billRefund         float32 // 退款金额
	billRefundSet      bool
	billDifference     float32 // 差异金额（出租车为调整金额）
	billDifferenceSet  bool
	companyId          string // 公司 ID
	companyIdSet       bool
	companyName        string // 公司名称
	companyNameSet     bool
	confirmDate        string // 确认日期
	confirmDateSet     bool
	startDate          string // 开始时间（商旅适用）
	startDateSet       bool
	endDate            string // 结束时间（商旅适用）
	endDateSet         bool
	groupType          int32 // 分组类型
	groupTypeSet       bool
	rootBatchId        int64 // 根结点批次（商旅适用）
	rootBatchIdSet     bool
	status             int32 // 账单状态（商旅适用），0. 生成中 1. 生成成功（待确认） 2.生成失败 3.已确认
	statusSet          bool
}

func NewBillRecordBuilder() *BillRecordBuilder {
	return &BillRecordBuilder{}
}
func (builder *BillRecordBuilder) BillId(billId int64) *BillRecordBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillRecordBuilder) BillPeriod(billPeriod string) *BillRecordBuilder {
	builder.billPeriod = billPeriod
	builder.billPeriodSet = true
	return builder
}
func (builder *BillRecordBuilder) BillStatus(billStatus int32) *BillRecordBuilder {
	builder.billStatus = billStatus
	builder.billStatusSet = true
	return builder
}
func (builder *BillRecordBuilder) BillAmount(billAmount float32) *BillRecordBuilder {
	builder.billAmount = billAmount
	builder.billAmountSet = true
	return builder
}
func (builder *BillRecordBuilder) BillConsumption(billConsumption float32) *BillRecordBuilder {
	builder.billConsumption = billConsumption
	builder.billConsumptionSet = true
	return builder
}
func (builder *BillRecordBuilder) BillRefund(billRefund float32) *BillRecordBuilder {
	builder.billRefund = billRefund
	builder.billRefundSet = true
	return builder
}
func (builder *BillRecordBuilder) BillDifference(billDifference float32) *BillRecordBuilder {
	builder.billDifference = billDifference
	builder.billDifferenceSet = true
	return builder
}
func (builder *BillRecordBuilder) CompanyId(companyId string) *BillRecordBuilder {
	builder.companyId = companyId
	builder.companyIdSet = true
	return builder
}
func (builder *BillRecordBuilder) CompanyName(companyName string) *BillRecordBuilder {
	builder.companyName = companyName
	builder.companyNameSet = true
	return builder
}
func (builder *BillRecordBuilder) ConfirmDate(confirmDate string) *BillRecordBuilder {
	builder.confirmDate = confirmDate
	builder.confirmDateSet = true
	return builder
}
func (builder *BillRecordBuilder) StartDate(startDate string) *BillRecordBuilder {
	builder.startDate = startDate
	builder.startDateSet = true
	return builder
}
func (builder *BillRecordBuilder) EndDate(endDate string) *BillRecordBuilder {
	builder.endDate = endDate
	builder.endDateSet = true
	return builder
}
func (builder *BillRecordBuilder) GroupType(groupType int32) *BillRecordBuilder {
	builder.groupType = groupType
	builder.groupTypeSet = true
	return builder
}
func (builder *BillRecordBuilder) RootBatchId(rootBatchId int64) *BillRecordBuilder {
	builder.rootBatchId = rootBatchId
	builder.rootBatchIdSet = true
	return builder
}
func (builder *BillRecordBuilder) Status(status int32) *BillRecordBuilder {
	builder.status = status
	builder.statusSet = true
	return builder
}

func (builder *BillRecordBuilder) Build() *BillRecord {
	data := &BillRecord{}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.billPeriodSet {
		data.BillPeriod = &builder.billPeriod
	}
	if builder.billStatusSet {
		data.BillStatus = &builder.billStatus
	}
	if builder.billAmountSet {
		data.BillAmount = &builder.billAmount
	}
	if builder.billConsumptionSet {
		data.BillConsumption = &builder.billConsumption
	}
	if builder.billRefundSet {
		data.BillRefund = &builder.billRefund
	}
	if builder.billDifferenceSet {
		data.BillDifference = &builder.billDifference
	}
	if builder.companyIdSet {
		data.CompanyId = &builder.companyId
	}
	if builder.companyNameSet {
		data.CompanyName = &builder.companyName
	}
	if builder.confirmDateSet {
		data.ConfirmDate = &builder.confirmDate
	}
	if builder.startDateSet {
		data.StartDate = &builder.startDate
	}
	if builder.endDateSet {
		data.EndDate = &builder.endDate
	}
	if builder.groupTypeSet {
		data.GroupType = &builder.groupType
	}
	if builder.rootBatchIdSet {
		data.RootBatchId = &builder.rootBatchId
	}
	if builder.statusSet {
		data.Status = &builder.status
	}
	return data
}
