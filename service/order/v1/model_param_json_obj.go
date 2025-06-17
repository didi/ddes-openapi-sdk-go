package v1

// ParamJsonObj param_json消息
type ParamJsonObj struct {
	OrderType    *string       `json:"order_type,omitempty"` // 订单类型，枚举值英文code：all 所有类别，domesticflight 国内机票，internationalflight 国际机票，domestichotel 国内酒店，internationalhotel 国际酒店，train 火车票，car 用车
	SearchTime   *SearchTime   `json:"search_time,omitempty"`
	Approval     *Approval     `json:"approval,omitempty"`
	Booker       *Booker       `json:"booker,omitempty"`
	Passenger    *Passenger    `json:"passenger,omitempty"`
	BudgetCenter *BudgetCenter `json:"budget_center,omitempty"`
	PayInfo      *PayInfo      `json:"pay_info,omitempty"`
	InvoiceInfo  *InvoiceInfo  `json:"invoice_info,omitempty"`
	CurPage      *int32        `json:"cur_page,omitempty"` // 当前页，当前页，从1开始
	Limit        *int32        `json:"limit,omitempty"`    // 每页条数，每页查询多少条(限制100)，可支持的分页最大返回10000条，超过部分不返回。单次超过查询范围，请减少查询时间范围。最多单次单品类返回10000条
}

type ParamJsonObjBuilder struct {
	orderType       string // 订单类型，枚举值英文code：all 所有类别，domesticflight 国内机票，internationalflight 国际机票，domestichotel 国内酒店，internationalhotel 国际酒店，train 火车票，car 用车
	orderTypeSet    bool
	searchTime      SearchTime
	searchTimeSet   bool
	approval        Approval
	approvalSet     bool
	booker          Booker
	bookerSet       bool
	passenger       Passenger
	passengerSet    bool
	budgetCenter    BudgetCenter
	budgetCenterSet bool
	payInfo         PayInfo
	payInfoSet      bool
	invoiceInfo     InvoiceInfo
	invoiceInfoSet  bool
	curPage         int32 // 当前页，当前页，从1开始
	curPageSet      bool
	limit           int32 // 每页条数，每页查询多少条(限制100)，可支持的分页最大返回10000条，超过部分不返回。单次超过查询范围，请减少查询时间范围。最多单次单品类返回10000条
	limitSet        bool
}

func NewParamJsonObjBuilder() *ParamJsonObjBuilder {
	return &ParamJsonObjBuilder{}
}
func (builder *ParamJsonObjBuilder) OrderType(orderType string) *ParamJsonObjBuilder {
	builder.orderType = orderType
	builder.orderTypeSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) SearchTime(searchTime SearchTime) *ParamJsonObjBuilder {
	builder.searchTime = searchTime
	builder.searchTimeSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) Approval(approval Approval) *ParamJsonObjBuilder {
	builder.approval = approval
	builder.approvalSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) Booker(booker Booker) *ParamJsonObjBuilder {
	builder.booker = booker
	builder.bookerSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) Passenger(passenger Passenger) *ParamJsonObjBuilder {
	builder.passenger = passenger
	builder.passengerSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) BudgetCenter(budgetCenter BudgetCenter) *ParamJsonObjBuilder {
	builder.budgetCenter = budgetCenter
	builder.budgetCenterSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) PayInfo(payInfo PayInfo) *ParamJsonObjBuilder {
	builder.payInfo = payInfo
	builder.payInfoSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) InvoiceInfo(invoiceInfo InvoiceInfo) *ParamJsonObjBuilder {
	builder.invoiceInfo = invoiceInfo
	builder.invoiceInfoSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) CurPage(curPage int32) *ParamJsonObjBuilder {
	builder.curPage = curPage
	builder.curPageSet = true
	return builder
}
func (builder *ParamJsonObjBuilder) Limit(limit int32) *ParamJsonObjBuilder {
	builder.limit = limit
	builder.limitSet = true
	return builder
}

func (builder *ParamJsonObjBuilder) Build() *ParamJsonObj {
	data := &ParamJsonObj{}
	if builder.orderTypeSet {
		data.OrderType = &builder.orderType
	}
	if builder.searchTimeSet {
		data.SearchTime = &builder.searchTime
	}
	if builder.approvalSet {
		data.Approval = &builder.approval
	}
	if builder.bookerSet {
		data.Booker = &builder.booker
	}
	if builder.passengerSet {
		data.Passenger = &builder.passenger
	}
	if builder.budgetCenterSet {
		data.BudgetCenter = &builder.budgetCenter
	}
	if builder.payInfoSet {
		data.PayInfo = &builder.payInfo
	}
	if builder.invoiceInfoSet {
		data.InvoiceInfo = &builder.invoiceInfo
	}
	if builder.curPageSet {
		data.CurPage = &builder.curPage
	}
	if builder.limitSet {
		data.Limit = &builder.limit
	}
	return data
}
