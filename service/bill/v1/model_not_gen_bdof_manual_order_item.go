package v1

// NotGenBDOfManualOrderItem 未出账单 - 增值手工订单,参考内部文档进行定义；
type NotGenBDOfManualOrderItem struct {
	OrderId               *string  `json:"order_id,omitempty"`                // 订单号
	MainOrderId           *string  `json:"main_order_id,omitempty"`           // 主订单号
	EsOrderId             *string  `json:"es_order_id,omitempty"`             // 关联订单号
	SupplierTicketNumber  *string  `json:"supplier_ticket_number,omitempty"`  // 票号
	BookingMemberName     *string  `json:"booking_member_name,omitempty"`     // 预订人姓名
	BookingMemberNumber   *string  `json:"booking_member_number,omitempty"`   // 预订人工号
	PassengerMemberNumber *string  `json:"passenger_member_number,omitempty"` // 使用人工号
	PassengerMemberName   *string  `json:"passenger_member_name,omitempty"`   // 使用人姓名
	BudgetCenterName      *string  `json:"budget_center_name,omitempty"`      // 成本中心
	BudgetCenterCode      *string  `json:"budget_center_code,omitempty"`      // 成本中心id
	GoodName              *string  `json:"good_name,omitempty"`               // 产品名称
	CompanyRealPay        *float64 `json:"company_real_pay,omitempty"`        // 单价
	Desc                  *string  `json:"desc,omitempty"`                    // 描述
	OrderSource           *string  `json:"order_source,omitempty"`            // 订单来源
	ServiceFee            *float32 `json:"service_fee,omitempty"`             // 平台使用费
	BeforeCutServiceFee   *float32 `json:"before_cut_service_fee,omitempty"`  // 平台使用费（立减前）
	AddedEsCutFee         *float32 `json:"added_es_cut_fee,omitempty"`        // 立减金额
	AddedCutReason        *string  `json:"added_cut_reason,omitempty"`        // 立减原因
}

type NotGenBDOfManualOrderItemBuilder struct {
	orderId                  string // 订单号
	orderIdSet               bool
	mainOrderId              string // 主订单号
	mainOrderIdSet           bool
	esOrderId                string // 关联订单号
	esOrderIdSet             bool
	supplierTicketNumber     string // 票号
	supplierTicketNumberSet  bool
	bookingMemberName        string // 预订人姓名
	bookingMemberNameSet     bool
	bookingMemberNumber      string // 预订人工号
	bookingMemberNumberSet   bool
	passengerMemberNumber    string // 使用人工号
	passengerMemberNumberSet bool
	passengerMemberName      string // 使用人姓名
	passengerMemberNameSet   bool
	budgetCenterName         string // 成本中心
	budgetCenterNameSet      bool
	budgetCenterCode         string // 成本中心id
	budgetCenterCodeSet      bool
	goodName                 string // 产品名称
	goodNameSet              bool
	companyRealPay           float64 // 单价
	companyRealPaySet        bool
	desc                     string // 描述
	descSet                  bool
	orderSource              string // 订单来源
	orderSourceSet           bool
	serviceFee               float32 // 平台使用费
	serviceFeeSet            bool
	beforeCutServiceFee      float32 // 平台使用费（立减前）
	beforeCutServiceFeeSet   bool
	addedEsCutFee            float32 // 立减金额
	addedEsCutFeeSet         bool
	addedCutReason           string // 立减原因
	addedCutReasonSet        bool
}

func NewNotGenBDOfManualOrderItemBuilder() *NotGenBDOfManualOrderItemBuilder {
	return &NotGenBDOfManualOrderItemBuilder{}
}
func (builder *NotGenBDOfManualOrderItemBuilder) OrderId(orderId string) *NotGenBDOfManualOrderItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) MainOrderId(mainOrderId string) *NotGenBDOfManualOrderItemBuilder {
	builder.mainOrderId = mainOrderId
	builder.mainOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) EsOrderId(esOrderId string) *NotGenBDOfManualOrderItemBuilder {
	builder.esOrderId = esOrderId
	builder.esOrderIdSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) SupplierTicketNumber(supplierTicketNumber string) *NotGenBDOfManualOrderItemBuilder {
	builder.supplierTicketNumber = supplierTicketNumber
	builder.supplierTicketNumberSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) BookingMemberName(bookingMemberName string) *NotGenBDOfManualOrderItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) BookingMemberNumber(bookingMemberNumber string) *NotGenBDOfManualOrderItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *NotGenBDOfManualOrderItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) PassengerMemberName(passengerMemberName string) *NotGenBDOfManualOrderItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) BudgetCenterName(budgetCenterName string) *NotGenBDOfManualOrderItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) BudgetCenterCode(budgetCenterCode string) *NotGenBDOfManualOrderItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) GoodName(goodName string) *NotGenBDOfManualOrderItemBuilder {
	builder.goodName = goodName
	builder.goodNameSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) CompanyRealPay(companyRealPay float64) *NotGenBDOfManualOrderItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) Desc(desc string) *NotGenBDOfManualOrderItemBuilder {
	builder.desc = desc
	builder.descSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) OrderSource(orderSource string) *NotGenBDOfManualOrderItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) ServiceFee(serviceFee float32) *NotGenBDOfManualOrderItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) BeforeCutServiceFee(beforeCutServiceFee float32) *NotGenBDOfManualOrderItemBuilder {
	builder.beforeCutServiceFee = beforeCutServiceFee
	builder.beforeCutServiceFeeSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) AddedEsCutFee(addedEsCutFee float32) *NotGenBDOfManualOrderItemBuilder {
	builder.addedEsCutFee = addedEsCutFee
	builder.addedEsCutFeeSet = true
	return builder
}
func (builder *NotGenBDOfManualOrderItemBuilder) AddedCutReason(addedCutReason string) *NotGenBDOfManualOrderItemBuilder {
	builder.addedCutReason = addedCutReason
	builder.addedCutReasonSet = true
	return builder
}

func (builder *NotGenBDOfManualOrderItemBuilder) Build() *NotGenBDOfManualOrderItem {
	data := &NotGenBDOfManualOrderItem{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.mainOrderIdSet {
		data.MainOrderId = &builder.mainOrderId
	}
	if builder.esOrderIdSet {
		data.EsOrderId = &builder.esOrderId
	}
	if builder.supplierTicketNumberSet {
		data.SupplierTicketNumber = &builder.supplierTicketNumber
	}
	if builder.bookingMemberNameSet {
		data.BookingMemberName = &builder.bookingMemberName
	}
	if builder.bookingMemberNumberSet {
		data.BookingMemberNumber = &builder.bookingMemberNumber
	}
	if builder.passengerMemberNumberSet {
		data.PassengerMemberNumber = &builder.passengerMemberNumber
	}
	if builder.passengerMemberNameSet {
		data.PassengerMemberName = &builder.passengerMemberName
	}
	if builder.budgetCenterNameSet {
		data.BudgetCenterName = &builder.budgetCenterName
	}
	if builder.budgetCenterCodeSet {
		data.BudgetCenterCode = &builder.budgetCenterCode
	}
	if builder.goodNameSet {
		data.GoodName = &builder.goodName
	}
	if builder.companyRealPaySet {
		data.CompanyRealPay = &builder.companyRealPay
	}
	if builder.descSet {
		data.Desc = &builder.desc
	}
	if builder.orderSourceSet {
		data.OrderSource = &builder.orderSource
	}
	if builder.serviceFeeSet {
		data.ServiceFee = &builder.serviceFee
	}
	if builder.beforeCutServiceFeeSet {
		data.BeforeCutServiceFee = &builder.beforeCutServiceFee
	}
	if builder.addedEsCutFeeSet {
		data.AddedEsCutFee = &builder.addedEsCutFee
	}
	if builder.addedCutReasonSet {
		data.AddedCutReason = &builder.addedCutReason
	}
	return data
}
