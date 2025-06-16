package v1

// BillListItem 账单树查询(网约车、商旅) ~ 账单集合
type BillListItem struct {
	BillId               *string              `json:"bill_id,omitempty"`                  // 账单id
	Pid                  *string              `json:"pid,omitempty"`                      // 父级账单id
	Cid                  []string             `json:"cid,omitempty"`                      // 子集账单id集合
	BillEntityId         *string              `json:"bill_entity_id,omitempty"`           // 主体id 未拆分账单为公司id，拆分账单为组id（groupId）
	BillEntity           *string              `json:"bill_entity,omitempty"`              // 主体名称 未拆分账单为公司名称，拆分账单为组名称（groupName）
	DepartmentId         *string              `json:"department_id,omitempty"`            // 部门id（不支持多维拆账）
	BudgetCenterId       *string              `json:"budget_center_id,omitempty"`         // 成本中心id（不支持多维拆账）
	OutBudgetId          *string              `json:"out_budget_id,omitempty"`            // 外部成本中心id，根据本接口budget_center_id字段的值关联查询（不支持多维拆账）
	BillStatus           *int32               `json:"bill_status,omitempty"`              // 账单状态（1：待确认；3：已确认）
	BillAmount           *float32             `json:"bill_amount,omitempty"`              // 应结算金额
	BillConsumption      *float32             `json:"bill_consumption,omitempty"`         // 消耗金额
	BillRefund           *float32             `json:"bill_refund,omitempty"`              // 退款金额
	BillDifference       *float32             `json:"bill_difference,omitempty"`          // 调整金额
	BillSplitType        *int32               `json:"bill_split_type,omitempty"`          // 拆帐类型 1=按出行人部门拆分，2=按预订人部门拆分，3=按成本中心拆分，4=按出行人所属公司拆分，5=按预订人所属公司拆分，6=按自定义字段拆分，7=多维拆帐
	BillSplitGroupType   *int32               `json:"bill_split_group_type,omitempty"`    // 多维拆帐节点类型，支持网约车、代驾、出租车、商旅；1=按部门拆分，2=按项目拆分，4=按公司拆分，6=按自定义字段、用车制度、用车权限拆分，8= 未知主体注：如果选择支持账单合并，合并的子账单的拆帐类型只取优先级第一个的拆帐类型
	OutBillSplitGroupKey *string              `json:"out_bill_split_group_key,omitempty"` // 账单主体外部id或名称，支持网约车、代驾、出租车、商旅； 单维拆账时，bill_split_type为1~3时，根据部门id或成本中心id的值，来取外部成本中心（out_budget_id）；bill_split_type为4、5时，根据公司id的值，来取外部公司主体编号（out_legal_entity_id），bill_split_type为6时，填写自定义字段的值； 多维拆账单时，bill_split_group_type为1、2时，来取外部成本中心（out_budget_id）；bill_split_group_type为4时，来取外部公司主体编号（out_legal_entity_id）；bill_split_group_type为6时，值为滴滴生成的帐单主体名称，例如 按【自定义01】拆分的帐单主体为AAA则填写【AAA】，按 【自定义01】、【自定义02】拆分的帐单主体为AAA（自定义01）、BBB（自定义02）则填写【AAA&BBB】；bill_split_group_type为8时，值为滴滴生成的帐单主体名称，例如 1级拆分的未知主体名称为“未知主体”，2级拆分的未知主体名称为“xxx部门-未知主体”
	SubBillSummary       []SubBillSummaryItem `json:"sub_bill_summary,omitempty"`         // 账单集合
}

type BillListItemBuilder struct {
	billId                  string // 账单id
	billIdSet               bool
	pid                     string // 父级账单id
	pidSet                  bool
	cid                     []string // 子集账单id集合
	cidSet                  bool
	billEntityId            string // 主体id 未拆分账单为公司id，拆分账单为组id（groupId）
	billEntityIdSet         bool
	billEntity              string // 主体名称 未拆分账单为公司名称，拆分账单为组名称（groupName）
	billEntitySet           bool
	departmentId            string // 部门id（不支持多维拆账）
	departmentIdSet         bool
	budgetCenterId          string // 成本中心id（不支持多维拆账）
	budgetCenterIdSet       bool
	outBudgetId             string // 外部成本中心id，根据本接口budget_center_id字段的值关联查询（不支持多维拆账）
	outBudgetIdSet          bool
	billStatus              int32 // 账单状态（1：待确认；3：已确认）
	billStatusSet           bool
	billAmount              float32 // 应结算金额
	billAmountSet           bool
	billConsumption         float32 // 消耗金额
	billConsumptionSet      bool
	billRefund              float32 // 退款金额
	billRefundSet           bool
	billDifference          float32 // 调整金额
	billDifferenceSet       bool
	billSplitType           int32 // 拆帐类型 1=按出行人部门拆分，2=按预订人部门拆分，3=按成本中心拆分，4=按出行人所属公司拆分，5=按预订人所属公司拆分，6=按自定义字段拆分，7=多维拆帐
	billSplitTypeSet        bool
	billSplitGroupType      int32 // 多维拆帐节点类型，支持网约车、代驾、出租车、商旅；1=按部门拆分，2=按项目拆分，4=按公司拆分，6=按自定义字段、用车制度、用车权限拆分，8= 未知主体注：如果选择支持账单合并，合并的子账单的拆帐类型只取优先级第一个的拆帐类型
	billSplitGroupTypeSet   bool
	outBillSplitGroupKey    string // 账单主体外部id或名称，支持网约车、代驾、出租车、商旅； 单维拆账时，bill_split_type为1~3时，根据部门id或成本中心id的值，来取外部成本中心（out_budget_id）；bill_split_type为4、5时，根据公司id的值，来取外部公司主体编号（out_legal_entity_id），bill_split_type为6时，填写自定义字段的值； 多维拆账单时，bill_split_group_type为1、2时，来取外部成本中心（out_budget_id）；bill_split_group_type为4时，来取外部公司主体编号（out_legal_entity_id）；bill_split_group_type为6时，值为滴滴生成的帐单主体名称，例如 按【自定义01】拆分的帐单主体为AAA则填写【AAA】，按 【自定义01】、【自定义02】拆分的帐单主体为AAA（自定义01）、BBB（自定义02）则填写【AAA&BBB】；bill_split_group_type为8时，值为滴滴生成的帐单主体名称，例如 1级拆分的未知主体名称为“未知主体”，2级拆分的未知主体名称为“xxx部门-未知主体”
	outBillSplitGroupKeySet bool
	subBillSummary          []SubBillSummaryItem // 账单集合
	subBillSummarySet       bool
}

func NewBillListItemBuilder() *BillListItemBuilder {
	return &BillListItemBuilder{}
}
func (builder *BillListItemBuilder) BillId(billId string) *BillListItemBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillListItemBuilder) Pid(pid string) *BillListItemBuilder {
	builder.pid = pid
	builder.pidSet = true
	return builder
}
func (builder *BillListItemBuilder) Cid(cid []string) *BillListItemBuilder {
	builder.cid = cid
	builder.cidSet = true
	return builder
}
func (builder *BillListItemBuilder) BillEntityId(billEntityId string) *BillListItemBuilder {
	builder.billEntityId = billEntityId
	builder.billEntityIdSet = true
	return builder
}
func (builder *BillListItemBuilder) BillEntity(billEntity string) *BillListItemBuilder {
	builder.billEntity = billEntity
	builder.billEntitySet = true
	return builder
}
func (builder *BillListItemBuilder) DepartmentId(departmentId string) *BillListItemBuilder {
	builder.departmentId = departmentId
	builder.departmentIdSet = true
	return builder
}
func (builder *BillListItemBuilder) BudgetCenterId(budgetCenterId string) *BillListItemBuilder {
	builder.budgetCenterId = budgetCenterId
	builder.budgetCenterIdSet = true
	return builder
}
func (builder *BillListItemBuilder) OutBudgetId(outBudgetId string) *BillListItemBuilder {
	builder.outBudgetId = outBudgetId
	builder.outBudgetIdSet = true
	return builder
}
func (builder *BillListItemBuilder) BillStatus(billStatus int32) *BillListItemBuilder {
	builder.billStatus = billStatus
	builder.billStatusSet = true
	return builder
}
func (builder *BillListItemBuilder) BillAmount(billAmount float32) *BillListItemBuilder {
	builder.billAmount = billAmount
	builder.billAmountSet = true
	return builder
}
func (builder *BillListItemBuilder) BillConsumption(billConsumption float32) *BillListItemBuilder {
	builder.billConsumption = billConsumption
	builder.billConsumptionSet = true
	return builder
}
func (builder *BillListItemBuilder) BillRefund(billRefund float32) *BillListItemBuilder {
	builder.billRefund = billRefund
	builder.billRefundSet = true
	return builder
}
func (builder *BillListItemBuilder) BillDifference(billDifference float32) *BillListItemBuilder {
	builder.billDifference = billDifference
	builder.billDifferenceSet = true
	return builder
}
func (builder *BillListItemBuilder) BillSplitType(billSplitType int32) *BillListItemBuilder {
	builder.billSplitType = billSplitType
	builder.billSplitTypeSet = true
	return builder
}
func (builder *BillListItemBuilder) BillSplitGroupType(billSplitGroupType int32) *BillListItemBuilder {
	builder.billSplitGroupType = billSplitGroupType
	builder.billSplitGroupTypeSet = true
	return builder
}
func (builder *BillListItemBuilder) OutBillSplitGroupKey(outBillSplitGroupKey string) *BillListItemBuilder {
	builder.outBillSplitGroupKey = outBillSplitGroupKey
	builder.outBillSplitGroupKeySet = true
	return builder
}
func (builder *BillListItemBuilder) SubBillSummary(subBillSummary []SubBillSummaryItem) *BillListItemBuilder {
	builder.subBillSummary = subBillSummary
	builder.subBillSummarySet = true
	return builder
}

func (builder *BillListItemBuilder) Build() *BillListItem {
	data := &BillListItem{}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
	if builder.pidSet {
		data.Pid = &builder.pid
	}
	if builder.cidSet {
		data.Cid = builder.cid
	}
	if builder.billEntityIdSet {
		data.BillEntityId = &builder.billEntityId
	}
	if builder.billEntitySet {
		data.BillEntity = &builder.billEntity
	}
	if builder.departmentIdSet {
		data.DepartmentId = &builder.departmentId
	}
	if builder.budgetCenterIdSet {
		data.BudgetCenterId = &builder.budgetCenterId
	}
	if builder.outBudgetIdSet {
		data.OutBudgetId = &builder.outBudgetId
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
	if builder.billSplitTypeSet {
		data.BillSplitType = &builder.billSplitType
	}
	if builder.billSplitGroupTypeSet {
		data.BillSplitGroupType = &builder.billSplitGroupType
	}
	if builder.outBillSplitGroupKeySet {
		data.OutBillSplitGroupKey = &builder.outBillSplitGroupKey
	}
	if builder.subBillSummarySet {
		data.SubBillSummary = builder.subBillSummary
	}
	return data
}
