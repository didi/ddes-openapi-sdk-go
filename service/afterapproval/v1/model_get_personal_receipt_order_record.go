package v1

// GetPersonalReceiptOrderRecord struct for GetPersonalReceiptOrderRecord
type GetPersonalReceiptOrderRecord struct {
	OrderId    *string `json:"order_id,omitempty"`    // 订单号
	ApprovalId *string `json:"approval_id,omitempty"` // 审批单号
}

type GetPersonalReceiptOrderRecordBuilder struct {
	orderId       string // 订单号
	orderIdSet    bool
	approvalId    string // 审批单号
	approvalIdSet bool
}

func NewGetPersonalReceiptOrderRecordBuilder() *GetPersonalReceiptOrderRecordBuilder {
	return &GetPersonalReceiptOrderRecordBuilder{}
}
func (builder *GetPersonalReceiptOrderRecordBuilder) OrderId(orderId string) *GetPersonalReceiptOrderRecordBuilder {
	builder.orderId = orderId
	builder.orderIdSet = true
	return builder
}
func (builder *GetPersonalReceiptOrderRecordBuilder) ApprovalId(approvalId string) *GetPersonalReceiptOrderRecordBuilder {
	builder.approvalId = approvalId
	builder.approvalIdSet = true
	return builder
}

func (builder *GetPersonalReceiptOrderRecordBuilder) Build() *GetPersonalReceiptOrderRecord {
	data := &GetPersonalReceiptOrderRecord{}
	if builder.orderIdSet {
		data.OrderId = &builder.orderId
	}
	if builder.approvalIdSet {
		data.ApprovalId = &builder.approvalId
	}
	return data
}
