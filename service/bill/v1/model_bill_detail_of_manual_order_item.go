package v1

// BillDetailOfManualOrderItem 已出账单 ~ 增值手工单返回参数,参考内部文档进行定义；
type BillDetailOfManualOrderItem struct {
	BillId                *int64   `json:"bill_id,omitempty"`                 // 账单id 根结点批次id
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

type BillDetailOfManualOrderItemBuilder struct {
	billId                   int64 // 账单id 根结点批次id
	billIdSet                bool
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

func NewBillDetailOfManualOrderItemBuilder() *BillDetailOfManualOrderItemBuilder {
	return &BillDetailOfManualOrderItemBuilder{}
}
func (builder *BillDetailOfManualOrderItemBuilder) BillId(billId int64) *BillDetailOfManualOrderItemBuilder {
	builder.billId = billId
	builder.billIdSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) OrderId(orderId string) *BillDetailOfManualOrderItemBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) MainOrderId(mainOrderId string) *BillDetailOfManualOrderItemBuilder {
	builder.mainOrderId = mainOrderId
	builder.mainOrderIdSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) EsOrderId(esOrderId string) *BillDetailOfManualOrderItemBuilder {
	builder.esOrderId = esOrderId
	builder.esOrderIdSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) SupplierTicketNumber(supplierTicketNumber string) *BillDetailOfManualOrderItemBuilder {
	builder.supplierTicketNumber = supplierTicketNumber
	builder.supplierTicketNumberSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) BookingMemberName(bookingMemberName string) *BillDetailOfManualOrderItemBuilder {
	builder.bookingMemberName = bookingMemberName
	builder.bookingMemberNameSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) BookingMemberNumber(bookingMemberNumber string) *BillDetailOfManualOrderItemBuilder {
	builder.bookingMemberNumber = bookingMemberNumber
	builder.bookingMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) PassengerMemberNumber(passengerMemberNumber string) *BillDetailOfManualOrderItemBuilder {
	builder.passengerMemberNumber = passengerMemberNumber
	builder.passengerMemberNumberSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) PassengerMemberName(passengerMemberName string) *BillDetailOfManualOrderItemBuilder {
	builder.passengerMemberName = passengerMemberName
	builder.passengerMemberNameSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) BudgetCenterName(budgetCenterName string) *BillDetailOfManualOrderItemBuilder {
	builder.budgetCenterName = budgetCenterName
	builder.budgetCenterNameSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) BudgetCenterCode(budgetCenterCode string) *BillDetailOfManualOrderItemBuilder {
	builder.budgetCenterCode = budgetCenterCode
	builder.budgetCenterCodeSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) GoodName(goodName string) *BillDetailOfManualOrderItemBuilder {
	builder.goodName = goodName
	builder.goodNameSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) CompanyRealPay(companyRealPay float64) *BillDetailOfManualOrderItemBuilder {
	builder.companyRealPay = companyRealPay
	builder.companyRealPaySet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) Desc(desc string) *BillDetailOfManualOrderItemBuilder {
	builder.desc = desc
	builder.descSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) OrderSource(orderSource string) *BillDetailOfManualOrderItemBuilder {
	builder.orderSource = orderSource
	builder.orderSourceSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) ServiceFee(serviceFee float32) *BillDetailOfManualOrderItemBuilder {
	builder.serviceFee = serviceFee
	builder.serviceFeeSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) BeforeCutServiceFee(beforeCutServiceFee float32) *BillDetailOfManualOrderItemBuilder {
	builder.beforeCutServiceFee = beforeCutServiceFee
	builder.beforeCutServiceFeeSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) AddedEsCutFee(addedEsCutFee float32) *BillDetailOfManualOrderItemBuilder {
	builder.addedEsCutFee = addedEsCutFee
	builder.addedEsCutFeeSet = true
	return builder
}
func (builder *BillDetailOfManualOrderItemBuilder) AddedCutReason(addedCutReason string) *BillDetailOfManualOrderItemBuilder {
	builder.addedCutReason = addedCutReason
	builder.addedCutReasonSet = true
	return builder
}

func (builder *BillDetailOfManualOrderItemBuilder) Build() *BillDetailOfManualOrderItem {
	data := &BillDetailOfManualOrderItem{}
	if builder.billIdSet {
		data.BillId = &builder.billId
	}
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
