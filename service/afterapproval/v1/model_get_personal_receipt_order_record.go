package v1

import "github.com/didi/ddes-openapi-sdk-go/core"

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

// UnmarshalJSON 容错反序列化：OrderId 等字段真实流量 number/string 混存，
// *string 收 number 会失败。用 core.SmartDecode 经中间 map 转换，避免递归。
func (g *GetPersonalReceiptOrderRecord) UnmarshalJSON(data []byte) error {
	return core.SmartDecode(data, g)
}
