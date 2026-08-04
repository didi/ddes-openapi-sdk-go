package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// newBillTestOption 构造 bill 域测试用的 core.Option
func newBillTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

// billCommonParams 构造通用 query params 断言用例（GET 型）
func billCommonParams() []struct{ key, want string } {
	return []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
	}
}

// =====================================================================
// 模型层测试 - 数据模型 Builder
// =====================================================================

func TestErrMsgListItemBuilder(t *testing.T) {
	info := NewErrMsgListItemBuilder().
		SubBusinessType(202).
		SubOrderId("1125922289295589").
		AdjustErr("金额不匹配").
		Build()
	if info.SubBusinessType == nil || *info.SubBusinessType != 202 {
		t.Errorf("SubBusinessType = %v, want 202", info.SubBusinessType)
	}
	if info.SubOrderId == nil || *info.SubOrderId != "1125922289295589" {
		t.Errorf("SubOrderId = %v, want 1125922289295589", info.SubOrderId)
	}
	if info.AdjustErr == nil || *info.AdjustErr != "金额不匹配" {
		t.Errorf("AdjustErr = %v, want 金额不匹配", info.AdjustErr)
	}
	// 部分设置
	info2 := NewErrMsgListItemBuilder().SubOrderId("1001").Build()
	if info2.SubOrderId == nil || *info2.SubOrderId != "1001" {
		t.Errorf("SubOrderId = %v, want 1001", info2.SubOrderId)
	}
	if info2.SubBusinessType != nil {
		t.Errorf("SubBusinessType = %v, want nil", info2.SubBusinessType)
	}
}

func TestAdjustListItemBuilder(t *testing.T) {
	fields := NewAdjustFieldsBuilder().BookingDepName("技术部").Build()
	info := NewAdjustListItemBuilder().
		SubOrderId("1125922289295589").
		SubBusinessType(1).
		AdjustFields(*fields).
		Build()
	if info.SubOrderId == nil || *info.SubOrderId != "1125922289295589" {
		t.Errorf("SubOrderId = %v, want 1125922289295589", info.SubOrderId)
	}
	if info.SubBusinessType == nil || *info.SubBusinessType != 1 {
		t.Errorf("SubBusinessType = %v, want 1", info.SubBusinessType)
	}
	if info.AdjustFields == nil || info.AdjustFields.BookingDepName == nil || *info.AdjustFields.BookingDepName != "技术部" {
		t.Errorf("AdjustFields.BookingDepName mismatch")
	}
	// 部分设置
	info2 := NewAdjustListItemBuilder().SubOrderId("1001").Build()
	if info2.SubBusinessType != nil {
		t.Errorf("SubBusinessType = %v, want nil", info2.SubBusinessType)
	}
	if info2.AdjustFields != nil {
		t.Errorf("AdjustFields = %v, want nil", info2.AdjustFields)
	}
}

func TestAdjustFieldsBuilder(t *testing.T) {
	info := NewAdjustFieldsBuilder().
		BookingDepName("技术部").
		BookingDepCode("DEP001").
		BudgetCenterName("成本中心A").
		BudgetCenterCode("BC001").
		LegalEntityName("主体公司").
		LegalEntityCode("LE001").
		ProjectExtInfo("项目信息").
		ExInfo01("ext01").
		ExInfo02("ext02").
		Build()
	if info.BookingDepName == nil || *info.BookingDepName != "技术部" {
		t.Errorf("BookingDepName = %v, want 技术部", info.BookingDepName)
	}
	if info.BudgetCenterCode == nil || *info.BudgetCenterCode != "BC001" {
		t.Errorf("BudgetCenterCode = %v, want BC001", info.BudgetCenterCode)
	}
	if info.ExInfo02 == nil || *info.ExInfo02 != "ext02" {
		t.Errorf("ExInfo02 = %v, want ext02", info.ExInfo02)
	}
	// 部分设置
	info2 := NewAdjustFieldsBuilder().BookingDepName("行政部").Build()
	if info2.BookingDepName == nil || *info2.BookingDepName != "行政部" {
		t.Errorf("BookingDepName = %v, want 行政部", info2.BookingDepName)
	}
	if info2.BudgetCenterName != nil {
		t.Errorf("BudgetCenterName = %v, want nil", info2.BudgetCenterName)
	}
}

func TestBillRecordBuilder(t *testing.T) {
	record := NewBillRecordBuilder().
		BillId(1125922289295589).
		BillPeriod("2020-01-01~2020-01-31").
		BillStatus(3).
		BillAmount(5000.50).
		CompanyName("测试公司").
		Status(1).
		Build()
	if record.BillId == nil || *record.BillId != 1125922289295589 {
		t.Errorf("BillId = %v, want 1125922289295589", record.BillId)
	}
	if record.BillStatus == nil || *record.BillStatus != 3 {
		t.Errorf("BillStatus = %v, want 3", record.BillStatus)
	}
	if record.CompanyName == nil || *record.CompanyName != "测试公司" {
		t.Errorf("CompanyName = %v, want 测试公司", record.CompanyName)
	}
	// 部分设置
	record2 := NewBillRecordBuilder().BillId(1001).Build()
	if record2.BillId == nil || *record2.BillId != 1001 {
		t.Errorf("BillId = %v, want 1001", record2.BillId)
	}
	if record2.BillPeriod != nil {
		t.Errorf("BillPeriod = %v, want nil", record2.BillPeriod)
	}
}

func TestBillListItemBuilder(t *testing.T) {
	sub := NewSubBillSummaryItemBuilder().BusinessType(202).AmountMoney("100.00").Build()
	info := NewBillListItemBuilder().
		BillId("1125922289295589").
		Pid("0").
		Cid([]string{"1001", "1002"}).
		BillEntity("测试公司").
		BillStatus(3).
		BillAmount(5000.50).
		BillSplitType(1).
		SubBillSummary([]SubBillSummaryItem{*sub}).
		Build()
	if info.BillId == nil || *info.BillId != "1125922289295589" {
		t.Errorf("BillId = %v, want 1125922289295589", info.BillId)
	}
	if len(info.Cid) != 2 {
		t.Errorf("Cid len = %d, want 2", len(info.Cid))
	}
	if info.BillStatus == nil || *info.BillStatus != 3 {
		t.Errorf("BillStatus = %v, want 3", info.BillStatus)
	}
	if len(info.SubBillSummary) != 1 {
		t.Errorf("SubBillSummary len = %d, want 1", len(info.SubBillSummary))
	}
	// 部分设置
	info2 := NewBillListItemBuilder().BillId("1001").Build()
	if info2.Cid != nil {
		t.Errorf("Cid = %v, want nil", info2.Cid)
	}
}

func TestSubBillSummaryItemBuilder(t *testing.T) {
	info := NewSubBillSummaryItemBuilder().
		BusinessType(202).
		AmountMoney("1000.00").
		ConsumeAmount("950.00").
		RefundAmount("50.00").
		PreviousRefundAmount("10.00").
		Build()
	if info.BusinessType == nil || *info.BusinessType != 202 {
		t.Errorf("BusinessType = %v, want 202", info.BusinessType)
	}
	if info.AmountMoney == nil || *info.AmountMoney != "1000.00" {
		t.Errorf("AmountMoney = %v, want 1000.00", info.AmountMoney)
	}
	// 部分设置
	info2 := NewSubBillSummaryItemBuilder().BusinessType(203).Build()
	if info2.AmountMoney != nil {
		t.Errorf("AmountMoney = %v, want nil", info2.AmountMoney)
	}
}

func TestSubBillSummaryDTOBuilder(t *testing.T) {
	info := NewSubBillSummaryDTOBuilder().
		BusinessType(201).
		ConsumeAmount("500.00").
		AmountMoney("500.00").
		Build()
	if info.BusinessType == nil || *info.BusinessType != 201 {
		t.Errorf("BusinessType = %v, want 201", info.BusinessType)
	}
	if info.ConsumeAmount == nil || *info.ConsumeAmount != "500.00" {
		t.Errorf("ConsumeAmount = %v, want 500.00", info.ConsumeAmount)
	}
	info2 := NewSubBillSummaryDTOBuilder().BusinessType(204).Build()
	if info2.ConsumeAmount != nil {
		t.Errorf("ConsumeAmount = %v, want nil", info2.ConsumeAmount)
	}
}

func TestSubListItemBuilder(t *testing.T) {
	info := NewSubListItemBuilder().
		PayChannel("企业账户").
		ConsumeAmountOnline(1000.50).
		BillAmount(1000.50).
		Build()
	if info.PayChannel == nil || *info.PayChannel != "企业账户" {
		t.Errorf("PayChannel = %v, want 企业账户", info.PayChannel)
	}
	if info.ConsumeAmountOnline == nil || *info.ConsumeAmountOnline != 1000.50 {
		t.Errorf("ConsumeAmountOnline = %v, want 1000.50", info.ConsumeAmountOnline)
	}
	info2 := NewSubListItemBuilder().PayChannel("支付宝").Build()
	if info2.BillAmount != nil {
		t.Errorf("BillAmount = %v, want nil", info2.BillAmount)
	}
}

func TestBillDetailOfManualOrderItemBuilder(t *testing.T) {
	info := NewBillDetailOfManualOrderItemBuilder().
		BillId(1125922289295589).
		OrderId("ORDER001").
		BookingMemberName("张三").
		CompanyRealPay(100.50).
		ServiceFee(5.0).
		Build()
	if info.BillId == nil || *info.BillId != 1125922289295589 {
		t.Errorf("BillId = %v, want 1125922289295589", info.BillId)
	}
	if info.OrderId == nil || *info.OrderId != "ORDER001" {
		t.Errorf("OrderId = %v, want ORDER001", info.OrderId)
	}
	if info.BookingMemberName == nil || *info.BookingMemberName != "张三" {
		t.Errorf("BookingMemberName = %v, want 张三", info.BookingMemberName)
	}
	// 部分设置
	info2 := NewBillDetailOfManualOrderItemBuilder().OrderId("O2").Build()
	if info2.BillId != nil {
		t.Errorf("BillId = %v, want nil", info2.BillId)
	}
}

// =====================================================================
// 模型层测试 - POST 型 Request Builder (BillConfirm / GetAdjustBillDataResult / UpdateAdjustBillData)
// =====================================================================

func TestBillConfirmRequestBuilder(t *testing.T) {
	req := NewBillConfirmRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		BillId("1125922289295589").
		BusinessType(1).
		PaymentPeriod("2020-01-01~2020-01-31").
		DepartmentId("DEP001").
		BudgetCenterId("BC001").
		BillSplitType(1).
		BillSplitGroupType(1).
		BillSplitGroupKey("KEY001").
		Build()
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.BillId == nil || *req.BillId != "1125922289295589" {
		t.Errorf("BillId = %v, want 1125922289295589", req.BillId)
	}
	if req.BusinessType == nil || *req.BusinessType != 1 {
		t.Errorf("BusinessType = %v, want 1", req.BusinessType)
	}
	if req.BillSplitGroupKey == nil || *req.BillSplitGroupKey != "KEY001" {
		t.Errorf("BillSplitGroupKey = %v, want KEY001", req.BillSplitGroupKey)
	}
	// 部分设置
	req2 := NewBillConfirmRequestBuilder().ClientId("test_client").BillId("B1").Build()
	if req2.PaymentPeriod != nil {
		t.Errorf("PaymentPeriod = %v, want nil", req2.PaymentPeriod)
	}
	if req2.BusinessType != nil {
		t.Errorf("BusinessType = %v, want nil", req2.BusinessType)
	}
}

func TestBillConfirmRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewBillConfirmRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		BusinessType(0).
		BillSplitType(0).
		BillSplitGroupType(0).
		Build()
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.BusinessType == nil || *req.BusinessType != 0 {
		t.Errorf("BusinessType = %v, want 0", req.BusinessType)
	}
	if req.BillSplitType == nil || *req.BillSplitType != 0 {
		t.Errorf("BillSplitType = %v, want 0", req.BillSplitType)
	}
}

func TestGetAdjustBillDataResultRequestBuilder(t *testing.T) {
	req := NewGetAdjustBillDataResultRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		AdjustReqId("ADJ001").
		BillId(1125922289295589).
		Build()
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AdjustReqId == nil || *req.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId = %v, want ADJ001", req.AdjustReqId)
	}
	if req.BillId == nil || *req.BillId != 1125922289295589 {
		t.Errorf("BillId = %v, want 1125922289295589", req.BillId)
	}
	// 部分设置
	req2 := NewGetAdjustBillDataResultRequestBuilder().ClientId("c1").Build()
	if req2.AdjustReqId != nil {
		t.Errorf("AdjustReqId = %v, want nil", req2.AdjustReqId)
	}
}

func TestGetAdjustBillDataResultRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetAdjustBillDataResultRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		BillId(0).
		Build()
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.BillId == nil || *req.BillId != 0 {
		t.Errorf("BillId = %v, want 0", req.BillId)
	}
}

func TestUpdateAdjustBillDataRequestBuilder(t *testing.T) {
	adjustItem := NewAdjustListItemBuilder().
		SubOrderId("SUB001").
		SubBusinessType(1).
		AdjustFields(*NewAdjustFieldsBuilder().BookingDepName("技术部").Build()).
		Build()
	req := NewUpdateAdjustBillDataRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		Timestamp(1583484681).
		CompanyId("test_company").
		AdjustReqId("ADJ001").
		BusinessType(1).
		AdjustType(2).
		BillId(1125922289295589).
		AdjustList([]AdjustListItem{*adjustItem}).
		Remark("测试调账").
		Build()
	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AdjustReqId == nil || *req.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId = %v, want ADJ001", req.AdjustReqId)
	}
	if req.BillId == nil || *req.BillId != 1125922289295589 {
		t.Errorf("BillId = %v, want 1125922289295589", req.BillId)
	}
	if len(req.AdjustList) != 1 {
		t.Fatalf("AdjustList len = %d, want 1", len(req.AdjustList))
	}
	if req.AdjustList[0].SubOrderId == nil || *req.AdjustList[0].SubOrderId != "SUB001" {
		t.Errorf("AdjustList[0].SubOrderId mismatch")
	}
	if req.Remark == nil || *req.Remark != "测试调账" {
		t.Errorf("Remark = %v, want 测试调账", req.Remark)
	}
	// 部分设置
	req2 := NewUpdateAdjustBillDataRequestBuilder().ClientId("c1").Build()
	if req2.AdjustList != nil {
		t.Errorf("AdjustList = %v, want nil", req2.AdjustList)
	}
}

func TestUpdateAdjustBillDataRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewUpdateAdjustBillDataRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		BusinessType(0).
		AdjustType(0).
		BillId(0).
		Build()
	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.BusinessType == nil || *req.BusinessType != 0 {
		t.Errorf("BusinessType = %v, want 0", req.BusinessType)
	}
	if req.BillId == nil || *req.BillId != 0 {
		t.Errorf("BillId = %v, want 0", req.BillId)
	}
}

// =====================================================================
// 模型层测试 - GET 型 ApiReqBuilder QueryParams
// =====================================================================

func TestGetBillDetailOfDaiJiaApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		BillId("1125922289295589").
		PaymentPeriod("2020-01-01~2020-01-31").
		DepartmentId("DEP001").
		BudgetCenterId("BC001").
		Type(1).
		LastId("0").
		Length(10).
		BusinessType(1).
		OutRequisitionId("REQ001").
		BookingMemberId("MBR001").
		BillSplitType(1).
		BillSplitGroupType(1).
		BillSplitGroupKey("KEY001").
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"bill_id", "1125922289295589"},
		{"payment_period", "2020-01-01~2020-01-31"},
		{"department_id", "DEP001"},
		{"budget_center_id", "BC001"},
		{"type", "1"},
		{"last_id", "0"},
		{"length", "10"},
		{"business_type", "1"},
		{"out_requisition_id", "REQ001"},
		{"booking_member_id", "MBR001"},
		{"bill_split_type", "1"},
		{"bill_split_group_type", "1"},
		{"bill_split_group_key", "KEY001"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetBillDetailOfDaiJiaApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		BillId("1125922289295589").
		Build()
	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("bill_id") != "1125922289295589" {
		t.Errorf("bill_id mismatch")
	}
	if req.apiReq.QueryParams.Get("payment_period") != "" {
		t.Errorf("payment_period should be empty")
	}
	if req.apiReq.QueryParams.Get("business_type") != "" {
		t.Errorf("business_type should be empty")
	}
}

func TestGetBillDetailOfDaiJiaApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		Type(0).
		Length(0).
		BusinessType(0).
		BillSplitType(0).
		BillSplitGroupType(0).
		Build()
	if req.apiReq.QueryParams.Get("type") != "0" {
		t.Errorf("type = %q, want \"0\"", req.apiReq.QueryParams.Get("type"))
	}
	if req.apiReq.QueryParams.Get("length") != "0" {
		t.Errorf("length = %q, want \"0\"", req.apiReq.QueryParams.Get("length"))
	}
	if req.apiReq.QueryParams.Get("business_type") != "0" {
		t.Errorf("business_type = %q, want \"0\"", req.apiReq.QueryParams.Get("business_type"))
	}
}

func TestGetBillDetailOfDaiJiaApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()
	for _, tt := range billCommonParams() {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
	for _, key := range []string{"bill_id", "payment_period", "department_id", "budget_center_id", "type", "last_id", "length", "business_type", "bill_split_type", "bill_split_group_key"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

func TestGetBillStructureApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetBillStructureApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		PaymentPeriod("2020-01-01~2020-01-31").
		BusinessType("1").
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"payment_period", "2020-01-01~2020-01-31"},
		{"business_type", "1"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetBillStructureApiReqBuilder_PartialParams(t *testing.T) {
	req := NewGetBillStructureApiReqBuilder().
		ClientId("test_client").
		PaymentPeriod("2020-01-01~2020-01-31").
		Build()
	if req.apiReq.QueryParams.Get("payment_period") != "2020-01-01~2020-01-31" {
		t.Errorf("payment_period mismatch")
	}
	if req.apiReq.QueryParams.Get("business_type") != "" {
		t.Errorf("business_type should be empty")
	}
}

func TestGetBillStructureApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetBillStructureApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()
	for _, key := range []string{"payment_period", "business_type"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

func TestGetBillSummaryApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetBillSummaryApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		BillId("1125922289295589").
		BusinessLine(2).
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"bill_id", "1125922289295589"},
		{"business_line", "2"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetBillSummaryApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetBillSummaryApiReqBuilder().
		ClientId("test_client").
		BusinessLine(0).
		Build()
	if req.apiReq.QueryParams.Get("business_line") != "0" {
		t.Errorf("business_line = %q, want \"0\"", req.apiReq.QueryParams.Get("business_line"))
	}
}

func TestGetBillSummaryApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewGetBillSummaryApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()
	for _, key := range []string{"bill_id", "business_line"} {
		if v := req.apiReq.QueryParams.Get(key); v != "" {
			t.Errorf("QueryParams[%s] = %q, want empty", key, v)
		}
	}
}

func TestListBillApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListBillApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		BillStatus(1).
		Offset(0).
		Length(10).
		BusinessLine(2).
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"bill_status", "1"},
		{"offset", "0"},
		{"length", "10"},
		{"business_line", "2"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListBillApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewListBillApiReqBuilder().
		ClientId("test_client").
		BillStatus(0).
		Offset(0).
		Length(0).
		BusinessLine(0).
		Build()
	if req.apiReq.QueryParams.Get("bill_status") != "0" {
		t.Errorf("bill_status = %q, want \"0\"", req.apiReq.QueryParams.Get("bill_status"))
	}
	if req.apiReq.QueryParams.Get("offset") != "0" {
		t.Errorf("offset = %q, want \"0\"", req.apiReq.QueryParams.Get("offset"))
	}
}

func TestGetTransactionBillDetailApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetTransactionBillDetailApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		BillId("1125922289295589").
		LastId(0).
		Length(10).
		BusinessType(1).
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"bill_id", "1125922289295589"},
		{"last_id", "0"},
		{"length", "10"},
		{"business_type", "1"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetNotGenBillDetailOfDaiJiaApiReqBuilder_QueryParams(t *testing.T) {
	req := NewGetNotGenBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		BusinessType(1).
		StartDate("2020-01-01").
		EndDate("2020-01-31").
		LastId("0").
		OutRequisitionId("REQ001").
		BookingMemberId("MBR001").
		DateQueryType("1").
		Build()
	tests := []struct{ key, want string }{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"business_type", "1"},
		{"start_date", "2020-01-01"},
		{"end_date", "2020-01-31"},
		{"last_id", "0"},
		{"out_requisition_id", "REQ001"},
		{"booking_member_id", "MBR001"},
		{"date_query_type", "1"},
	}
	for _, tt := range tests {
		if got := req.apiReq.QueryParams.Get(tt.key); got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestGetNotGenBillDetailOfDaiJiaApiReqBuilder_ZeroIntValues(t *testing.T) {
	req := NewGetNotGenBillDetailOfDaiJiaApiReqBuilder().
		ClientId("test_client").
		BusinessType(0).
		Build()
	if req.apiReq.QueryParams.Get("business_type") != "0" {
		t.Errorf("business_type = %q, want \"0\"", req.apiReq.QueryParams.Get("business_type"))
	}
}

// =====================================================================
// 模型层测试 - 反序列化
// =====================================================================

func TestBillConfirmApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"CONFIRM001"},"request_id":"req_001"}`
	var reply BillConfirmApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %s, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.ConfirmBillId == nil || *reply.Data.ConfirmBillId != "CONFIRM001" {
		t.Errorf("ConfirmBillId = %v, want CONFIRM001", reply.Data.ConfirmBillId)
	}
}

func TestBillConfirmApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
	var reply BillConfirmApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestBillConfirmApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","request_id":"req_nodata"}`
	var reply BillConfirmApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestGetAdjustBillDataResultApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ001","check_pass":true,"adjust_status":3,"company_id":1125922289295589,"business_type":1,"adjust_type":2,"bill_id":1125922289295589,"err_msg_list":[{"sub_business_type":1,"sub_order_id":"SUB001","adjust_err":"金额不匹配"}]},"request_id":"req_001"}`
	var reply GetAdjustBillDataResultApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.AdjustReqId == nil || *reply.Data.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId = %v, want ADJ001", reply.Data.AdjustReqId)
	}
	if reply.Data.CheckPass == nil || !*reply.Data.CheckPass {
		t.Errorf("CheckPass = %v, want true", reply.Data.CheckPass)
	}
	if reply.Data.AdjustStatus == nil || *reply.Data.AdjustStatus != 3 {
		t.Errorf("AdjustStatus = %v, want 3", reply.Data.AdjustStatus)
	}
	if len(reply.Data.ErrMsgList) != 1 {
		t.Fatalf("ErrMsgList len = %d, want 1", len(reply.Data.ErrMsgList))
	}
	if reply.Data.ErrMsgList[0].SubOrderId == nil || *reply.Data.ErrMsgList[0].SubOrderId != "SUB001" {
		t.Errorf("ErrMsgList[0].SubOrderId mismatch")
	}
}

func TestGetAdjustBillDataResultApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_err"}`
	var reply GetAdjustBillDataResultApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestUpdateAdjustBillDataApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ001","check_pass":true,"adjust_status":1,"company_id":1125922289295589,"bill_id":1125922289295589,"business_type":1,"adjust_type":2,"err_msg":null,"err_msg_list":[]},"request_id":"req_001"}`
	var reply UpdateAdjustBillDataApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.AdjustReqId == nil || *reply.Data.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId = %v, want ADJ001", reply.Data.AdjustReqId)
	}
	if reply.Data.AdjustStatus == nil || *reply.Data.AdjustStatus != 1 {
		t.Errorf("AdjustStatus = %v, want 1", reply.Data.AdjustStatus)
	}
	if reply.Data.BillId == nil || *reply.Data.BillId != 1125922289295589 {
		t.Errorf("BillId = %v, want 1125922289295589", reply.Data.BillId)
	}
}

func TestListBillApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":78,"records":[{"bill_id":1125922289295589,"bill_period":"2020-01-01~2020-01-31","bill_status":3,"companyName":"测试公司","bill_amount":5000.50}],"lastId":100},"request_id":"req_001"}`
	var reply ListBillApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total != 78 {
		t.Errorf("Total = %d, want 78", reply.Data.Total)
	}
	if len(reply.Data.Records) != 1 {
		t.Fatalf("Records len = %d, want 1", len(reply.Data.Records))
	}
	if reply.Data.Records[0].BillId == nil || *reply.Data.Records[0].BillId != 1125922289295589 {
		t.Errorf("Records[0].BillId mismatch")
	}
	if reply.Data.Records[0].BillStatus == nil || *reply.Data.Records[0].BillStatus != 3 {
		t.Errorf("Records[0].BillStatus = %v, want 3", reply.Data.Records[0].BillStatus)
	}
}

func TestListBillApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[],"lastId":0},"request_id":"req_empty"}`
	var reply ListBillApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Records) != 0 {
		t.Errorf("Records len = %d, want 0", len(reply.Data.Records))
	}
}

func TestListBillApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`
	var reply ListBillApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestGetBillDetailOfManualOrderApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"is_last":true,"last_id":"1125922289295589","orders":[{"bill_id":1125922289295589,"order_id":"1125922289295589","booking_member_name":"张三","company_real_pay":100.50,"service_fee":5.0}]},"request_id":"req_001"}`
	var reply GetBillDetailOfManualOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total == nil || *reply.Data.Total != 1 {
		t.Errorf("Total = %v, want 1", reply.Data.Total)
	}
	if len(reply.Data.Orders) != 1 {
		t.Fatalf("Orders len = %d, want 1", len(reply.Data.Orders))
	}
	if reply.Data.Orders[0].BookingMemberName == nil || *reply.Data.Orders[0].BookingMemberName != "张三" {
		t.Errorf("Orders[0].BookingMemberName mismatch")
	}
}

func TestGetBillDetailOfManualOrderApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"is_last":true,"orders":[]},"request_id":"req_empty"}`
	var reply GetBillDetailOfManualOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Orders) != 0 {
		t.Errorf("Orders len = %d, want 0", len(reply.Data.Orders))
	}
}

func TestGetBillDetailOfManualOrderApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`
	var reply GetBillDetailOfManualOrderApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestGetBillStructureApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[{"bill_id":"1125922289295589","pid":"0","bill_entity":"测试公司","bill_status":3,"bill_amount":5000.50,"sub_bill_summary":[{"business_type":202,"amount_money":"1000.00","consume_amount":"950.00"}]}]},"request_id":"req_001"}`
	var reply GetBillStructureApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.List) != 1 {
		t.Fatalf("List len = %d, want 1", len(reply.Data.List))
	}
	if reply.Data.List[0].BillId == nil || *reply.Data.List[0].BillId != "1125922289295589" {
		t.Errorf("List[0].BillId mismatch")
	}
	if len(reply.Data.List[0].SubBillSummary) != 1 {
		t.Errorf("SubBillSummary len = %d, want 1", len(reply.Data.List[0].SubBillSummary))
	}
}

func TestGetBillStructureApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_empty"}`
	var reply GetBillStructureApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.List) != 0 {
		t.Errorf("List len = %d, want 0", len(reply.Data.List))
	}
}

func TestGetBillSummaryApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"bill_id":"1125922289295589","company_name":"测试公司","bill_period":"2020-01","consume_amount_online":1000.50,"amount_money":950.00,"business_type":2,"consume_total_amount":"1000.50","sub_list":[{"pay_channel":"企业账户","bill_amount":1000.50}],"sub_bill_summary":[{"business_type":202,"consume_amount":"950.00","amount_money":"950.00"}]},"request_id":"req_001"}`
	var reply GetBillSummaryApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.BillId == nil || *reply.Data.BillId != "1125922289295589" {
		t.Errorf("BillId mismatch")
	}
	if reply.Data.ConsumeTotalAmount == nil || *reply.Data.ConsumeTotalAmount != "1000.50" {
		t.Errorf("ConsumeTotalAmount mismatch")
	}
	if len(reply.Data.SubList) != 1 {
		t.Errorf("SubList len = %d, want 1", len(reply.Data.SubList))
	}
	if len(reply.Data.SubBillSummary) != 1 {
		t.Errorf("SubBillSummary len = %d, want 1", len(reply.Data.SubBillSummary))
	}
}

func TestGetTransactionBillDetailApiReply_Deserialization(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"last_id":100,"transaction_list":[{"order_id":1125922289295589}]},"request_id":"req_001"}`
	var reply GetTransactionBillDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.Total == nil || *reply.Data.Total != 1 {
		t.Errorf("Total = %v, want 1", reply.Data.Total)
	}
	if len(reply.Data.TransactionList) != 1 {
		t.Errorf("TransactionList len = %d, want 1", len(reply.Data.TransactionList))
	}
}

func TestGetTransactionBillDetailApiReply_EmptyDataArray(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"transaction_list":[]},"request_id":"req_empty"}`
	var reply GetTransactionBillDetailApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.TransactionList) != 0 {
		t.Errorf("TransactionList len = %d, want 0", len(reply.Data.TransactionList))
	}
}

// =====================================================================
// 资源方法测试 - POST 型：BillConfirm / GetAdjustBillDataResult / UpdateAdjustBillData
// =====================================================================

func TestBillConfirm_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/Bill/confirm" {
			t.Errorf("expected path /river/Bill/confirm, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"CONFIRM001"},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").BillId("1125922289295589").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.BillConfirmApiReply == nil {
		t.Fatal("BillConfirmApiReply is nil")
	}
	if resp.BillConfirmApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.BillConfirmApiReply.Errno)
	}
	if resp.BillConfirmApiReply.Data == nil || resp.BillConfirmApiReply.Data.ConfirmBillId == nil || *resp.BillConfirmApiReply.Data.ConfirmBillId != "CONFIRM001" {
		t.Errorf("ConfirmBillId mismatch")
	}
}

func TestBillConfirm_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply == nil {
		t.Fatal("BillConfirmApiReply is nil")
	}
	if resp.BillConfirmApiReply.Data != nil {
		t.Errorf("Data should be nil")
	}
}

func TestBillConfirm_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.BillConfirmApiReply.Errno)
	}
}

func TestBillConfirm_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.BillConfirmApiReply != nil {
		t.Errorf("BillConfirmApiReply should be nil for non-200 response")
	}
}

func TestBillConfirm_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"CONFIRM_ENC"},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply == nil || resp.BillConfirmApiReply.Errno != 0 {
		t.Fatalf("decryption failed: %v", resp.BillConfirmApiReply)
	}
	if resp.BillConfirmApiReply.Data.ConfirmBillId == nil || *resp.BillConfirmApiReply.Data.ConfirmBillId != "CONFIRM_ENC" {
		t.Errorf("ConfirmBillId mismatch")
	}
}

func TestBillConfirm_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"CONFIRM_256"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply == nil || resp.BillConfirmApiReply.Data == nil {
		t.Fatal("decryption failed")
	}
	if resp.BillConfirmApiReply.Data.ConfirmBillId == nil || *resp.BillConfirmApiReply.Data.ConfirmBillId != "CONFIRM_256" {
		t.Errorf("ConfirmBillId mismatch")
	}
}

func TestBillConfirm_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"CONFIRM_NOENC"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	b := &bill{option: option}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply == nil || resp.BillConfirmApiReply.Data == nil {
		t.Fatal("reply/data is nil")
	}
}

func TestBillConfirm_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", h)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"confirm_bill_id":"X"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}
	req := NewBillConfirmApiReqBuilder().
		BillConfirmRequest(NewBillConfirmRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.BillConfirm(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("BillConfirm() error = %v", err)
	}
	if resp.BillConfirmApiReply == nil {
		t.Fatal("BillConfirmApiReply is nil")
	}
}

func TestGetAdjustBillDataResult_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/Bill/queryAdjustBillDataResult" {
			t.Errorf("expected path /river/Bill/queryAdjustBillDataResult, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ001","check_pass":true,"adjust_status":3},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").AdjustReqId("ADJ001").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("GetAdjustBillDataResult() error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply == nil {
		t.Fatal("GetAdjustBillDataResultApiReply is nil")
	}
	if resp.GetAdjustBillDataResultApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.GetAdjustBillDataResultApiReply.Errno)
	}
	if resp.GetAdjustBillDataResultApiReply.Data.AdjustReqId == nil || *resp.GetAdjustBillDataResultApiReply.Data.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId mismatch")
	}
}

func TestGetAdjustBillDataResult_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply.Data != nil {
		t.Errorf("Data should be nil")
	}
}

func TestGetAdjustBillDataResult_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.GetAdjustBillDataResultApiReply.Errno)
	}
}

func TestGetAdjustBillDataResult_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.GetAdjustBillDataResultApiReply != nil {
		t.Errorf("GetAdjustBillDataResultApiReply should be nil")
	}
}

func TestGetAdjustBillDataResult_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_ENC","adjust_status":3},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply == nil || resp.GetAdjustBillDataResultApiReply.Data == nil {
		t.Fatal("decryption failed")
	}
	if resp.GetAdjustBillDataResultApiReply.Data.AdjustReqId == nil || *resp.GetAdjustBillDataResultApiReply.Data.AdjustReqId != "ADJ_ENC" {
		t.Errorf("AdjustReqId mismatch")
	}
}

func TestGetAdjustBillDataResult_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_256"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply == nil || resp.GetAdjustBillDataResultApiReply.Data == nil {
		t.Fatal("decryption failed")
	}
}

func TestGetAdjustBillDataResult_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_NOENC"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	b := &bill{option: option}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply == nil {
		t.Fatal("reply is nil")
	}
}

func TestGetAdjustBillDataResult_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
			t.Errorf("expected X-Custom-Header, got %s", h)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"X"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}
	req := NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(NewGetAdjustBillDataResultRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.GetAdjustBillDataResult(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.GetAdjustBillDataResultApiReply == nil {
		t.Fatal("reply is nil")
	}
}

func TestUpdateAdjustBillData_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/Bill/adjustBillData" {
			t.Errorf("expected path /river/Bill/adjustBillData, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ001","check_pass":true,"adjust_status":1,"bill_id":1125922289295589},"request_id":"req_001"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").AdjustReqId("ADJ001").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply == nil {
		t.Fatal("UpdateAdjustBillDataApiReply is nil")
	}
	if resp.UpdateAdjustBillDataApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateAdjustBillDataApiReply.Errno)
	}
	if resp.UpdateAdjustBillDataApiReply.Data.AdjustReqId == nil || *resp.UpdateAdjustBillDataApiReply.Data.AdjustReqId != "ADJ001" {
		t.Errorf("AdjustReqId mismatch")
	}
}

func TestUpdateAdjustBillData_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply.Data != nil {
		t.Errorf("Data should be nil")
	}
}

func TestUpdateAdjustBillData_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.UpdateAdjustBillDataApiReply.Errno)
	}
}

func TestUpdateAdjustBillData_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateAdjustBillDataApiReply != nil {
		t.Errorf("UpdateAdjustBillDataApiReply should be nil")
	}
}

func TestUpdateAdjustBillData_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_ENC","adjust_status":1},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply == nil || resp.UpdateAdjustBillDataApiReply.Data == nil {
		t.Fatal("decryption failed")
	}
}

func TestUpdateAdjustBillData_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_256"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply == nil {
		t.Fatal("decryption failed")
	}
}

func TestUpdateAdjustBillData_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"ADJ_NOENC"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
	b := &bill{option: option}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply == nil {
		t.Fatal("reply is nil")
	}
}

func TestUpdateAdjustBillData_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
			t.Errorf("expected X-Custom-Header, got %s", h)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"X"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newBillTestOption(testServer.URL)
	b := &bill{option: option}
	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{Header: customHeader}
	req := NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(NewUpdateAdjustBillDataRequestBuilder().ClientId("test_client").Build()).
		Build()

	resp, err := b.UpdateAdjustBillData(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if resp.UpdateAdjustBillDataApiReply == nil {
		t.Fatal("reply is nil")
	}
}

// =====================================================================
// 资源方法测试 - GET 型 Bill Detail（9 个业务类型，共用 /river/Bill/detail 路径）
// 用表驱动方式覆盖 8 个场景，减少重复代码
// =====================================================================

// billDetailMethod 描述一个 GET 型 bill detail 资源方法的元信息
type billDetailMethod struct {
	name        string // 资源方法名，如 "GetBillDetailOfDaiJia"
	path        string // API 路径
	buildReq    func() interface{}
	invoke      func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error)
	hasReply    func(resp interface{}) bool
	successBody string // 成功响应体（data 含一条记录）
	emptyBody   string // 空 data 响应体
}

// runBillDetailScenarios 对给定的 bill detail 方法运行 8 个标准场景
func runBillDetailScenarios(t *testing.T, m billDetailMethod) {
	t.Helper()

	t.Run("Success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != m.path {
				t.Errorf("expected path %s, got %s", m.path, r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		b := &bill{option: newBillTestOption(ts.URL)}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("EmptyData", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.emptyBody))
		}))
		defer ts.Close()
		b := &bill{option: newBillTestOption(ts.URL)}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("ApiError", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_err"}`))
		}))
		defer ts.Close()
		b := &bill{option: newBillTestOption(ts.URL)}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		// reply 非 nil，errno 非 0
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("HttpError", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()
		b := &bill{option: newBillTestOption(ts.URL)}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if m.hasReply(resp) {
			t.Errorf("reply should be nil for non-200 response")
		}
	})

	t.Run("WithEncryption_AES128", func(t *testing.T) {
		plaintext := m.successBody
		key := []byte("16byte-key-12345")
		encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		encryptData := base64.StdEncoding.EncodeToString(encrypted)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
		}))
		defer ts.Close()
		opt := newBillTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: string(key)}
		b := &bill{option: opt}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("decryption failed: reply is nil")
		}
	})

	t.Run("WithEncryption_AES256", func(t *testing.T) {
		plaintext := m.successBody
		key := []byte("32byte-key-1234567890abcdefghijk")
		encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		encryptData := base64.URLEncoding.EncodeToString(encrypted)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
		}))
		defer ts.Close()
		opt := newBillTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 2, Key: string(key)}
		b := &bill{option: opt}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("decryption failed: reply is nil")
		}
	})

	t.Run("EncryptionNoEncryptData", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		opt := newBillTestOption(ts.URL)
		opt.EnableEncryption = true
		opt.EncryptionOption = &core.EncryptionOption{Ent: 1, Key: "16byte-key-12345"}
		b := &bill{option: opt}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, nil)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})

	t.Run("WithReqOption", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if h := r.Header.Get("X-Custom-Header"); h != "custom-value" {
				t.Errorf("expected X-Custom-Header, got %s", h)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(m.successBody))
		}))
		defer ts.Close()
		b := &bill{option: newBillTestOption(ts.URL)}
		customHeader := http.Header{}
		customHeader.Set("X-Custom-Header", "custom-value")
		reqOption := &core.ReqOption{Header: customHeader}
		req := m.buildReq()
		resp, err := m.invoke(b, context.Background(), req, reqOption)
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if !m.hasReply(resp) {
			t.Fatal("reply is nil")
		}
	})
}

// 已出账单 detail 通用成功响应（含一条 order）
// 由于各业务类型 LastId/OrderId 字段类型不统一（SDK 历史技术债），按类型分组提供 mock：
//   - IntInt:  LastId/OrderId 均为 *int64（dai_jia、wang_yc），无 bill_id
//   - StrInt:  LastId 为 *string、OrderId 为 *int64（taxi），无 bill_id
//   - StrStr:  LastId/OrderId 均为 *string，bill_id 为 *int64（flight/hotel/inter_*/manual_order/train）
const billDetailSuccessBody_IntInt = `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"is_last":true,"last_id":1125922289295589,"orders":[{"order_id":1125922289295589,"booking_member_name":"张三"}]},"request_id":"req_001"}`
const billDetailSuccessBody_StrInt = `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"is_last":true,"last_id":"1125922289295589","orders":[{"order_id":1125922289295589,"booking_member_name":"张三"}]},"request_id":"req_001"}`
const billDetailSuccessBody_StrStr = `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"is_last":true,"last_id":"1125922289295589","orders":[{"order_id":"1125922289295589","bill_id":1125922289295589,"booking_member_name":"张三"}]},"request_id":"req_001"}`
const billDetailEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"is_last":true,"orders":[]},"request_id":"req_empty"}`

func TestGetBillDetailOfDaiJia_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfDaiJiaApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfDaiJia(ctx, req.(*GetBillDetailOfDaiJiaApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfDaiJiaApiResp).GetBillDetailOfDaiJiaApiReply != nil
		},
		successBody: billDetailSuccessBody_IntInt,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfDomesticFlight_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfDomesticFlightApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfDomesticFlight(ctx, req.(*GetBillDetailOfDomesticFlightApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfDomesticFlightApiResp).GetBillDetailOfDomesticFlightApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfDomesticHotel_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfDomesticHotelApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfDomesticHotel(ctx, req.(*GetBillDetailOfDomesticHotelApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfDomesticHotelApiResp).GetBillDetailOfDomesticHotelApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfInterFlight_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfInterFlightApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfInterFlight(ctx, req.(*GetBillDetailOfInterFlightApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfInterFlightApiResp).GetBillDetailOfInterFlightApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfInterHotel_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfInterHotelApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfInterHotel(ctx, req.(*GetBillDetailOfInterHotelApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfInterHotelApiResp).GetBillDetailOfInterHotelApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfManualOrder_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfManualOrderApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfManualOrder(ctx, req.(*GetBillDetailOfManualOrderApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfManualOrderApiResp).GetBillDetailOfManualOrderApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfTaxi_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfTaxiApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfTaxi(ctx, req.(*GetBillDetailOfTaxiApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfTaxiApiResp).GetBillDetailOfTaxiApiReply != nil
		},
		successBody: billDetailSuccessBody_StrInt,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfTrainTicket_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfTrainTicketApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfTrainTicket(ctx, req.(*GetBillDetailOfTrainTicketApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfTrainTicketApiResp).GetBillDetailOfTrainTicketApiReply != nil
		},
		successBody: billDetailSuccessBody_StrStr,
		emptyBody:   billDetailEmptyBody,
	})
}

func TestGetBillDetailOfWangYC_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/detail",
		buildReq: func() interface{} {
			return NewGetBillDetailOfWangYCApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillDetailOfWangYC(ctx, req.(*GetBillDetailOfWangYCApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetBillDetailOfWangYCApiResp).GetBillDetailOfWangYCApiReply != nil
		},
		successBody: billDetailSuccessBody_IntInt,
		emptyBody:   billDetailEmptyBody,
	})
}

// =====================================================================
// 资源方法测试 - GET 型 未出账单 detail（8 个业务类型，/river/Bill/getNotGeneratedBillDetail）
// 复用 runBillDetailScenarios，只是 path 和 builder/invoke 不同
// =====================================================================

// 未出账单 detail 通用成功响应（含一条 detail）
// LastId 在所有 not_gen Reply 中均为 *string；OrderId 类型不统一：
//   - StrInt: OrderId 为 *int64（dai_jia、taxi、wang_yc）
//   - StrStr: OrderId 为 *string（flight/hotel/inter_*/manual_order/train）
const notGenBillDetailSuccessBody_StrInt = `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"last_id":"1125922289295589","detail_list":[{"order_id":1125922289295589}]},"request_id":"req_001"}`
const notGenBillDetailSuccessBody_StrStr = `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"last_id":"1125922289295589","detail_list":[{"order_id":"1125922289295589"}]},"request_id":"req_001"}`
const notGenBillDetailEmptyBody = `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"detail_list":[]},"request_id":"req_empty"}`

func TestGetNotGenBillDetailOfDaiJia_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfDaiJiaApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfDaiJia(ctx, req.(*GetNotGenBillDetailOfDaiJiaApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfDaiJiaApiResp).GetNotGenBillDetailOfDaiJiaApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrInt,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfFlight_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfFlightApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfFlight(ctx, req.(*GetNotGenBillDetailOfFlightApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfFlightApiResp).GetNotGenBillDetailOfFlightApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfHotel_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfHotelApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfHotel(ctx, req.(*GetNotGenBillDetailOfHotelApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfHotelApiResp).GetNotGenBillDetailOfHotelApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfInterFlight_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfInterFlightApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfInterFlight(ctx, req.(*GetNotGenBillDetailOfInterFlightApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfInterFlightApiResp).GetNotGenBillDetailOfInterFlightApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfInterHotel_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfInterHotelApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfInterHotel(ctx, req.(*GetNotGenBillDetailOfInterHotelApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfInterHotelApiResp).GetNotGenBillDetailOfInterHotelApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfManualOrder_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfManualOrderApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfManualOrder(ctx, req.(*GetNotGenBillDetailOfManualOrderApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfManualOrderApiResp).GetNotGenBillDetailOfManualOrderApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfTaxi_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfTaxiApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfTaxi(ctx, req.(*GetNotGenBillDetailOfTaxiApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfTaxiApiResp).GetNotGenBillDetailOfTaxiApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrInt,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfTrain_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfTrainApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfTrain(ctx, req.(*GetNotGenBillDetailOfTrainApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfTrainApiResp).GetNotGenBillDetailOfTrainApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrStr,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

func TestGetNotGenBillDetailOfWangYC_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getNotGeneratedBillDetail",
		buildReq: func() interface{} {
			return NewGetNotGenBillDetailOfWangYCApiReqBuilder().ClientId("test_client").BusinessType(1).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetNotGenBillDetailOfWangYC(ctx, req.(*GetNotGenBillDetailOfWangYCApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetNotGenBillDetailOfWangYCApiResp).GetNotGenBillDetailOfWangYCApiReply != nil
		},
		successBody: notGenBillDetailSuccessBody_StrInt,
		emptyBody:   notGenBillDetailEmptyBody,
	})
}

// =====================================================================
// 资源方法测试 - GET 型 其它接口（GetBillStructure / GetBillSummary / ListBill /
// GetTransactionBillDetail / GetTransactionBillDetailOfTaxi）
// 复用 runBillDetailScenarios
// =====================================================================

func TestGetBillStructure_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/getBillStructure",
		buildReq: func() interface{} {
			return NewGetBillStructureApiReqBuilder().ClientId("test_client").PaymentPeriod("2020-01-01~2020-01-31").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillStructure(ctx, req.(*GetBillStructureApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*GetBillStructureApiResp).GetBillStructureApiReply != nil },
		successBody: `{"errno":0,"errmsg":"SUCCESS","data":{"list":[{"bill_id":"1125922289295589","bill_entity":"测试公司","bill_status":3}]},"request_id":"req_001"}`,
		emptyBody:   `{"errno":0,"errmsg":"SUCCESS","data":{"list":[]},"request_id":"req_empty"}`,
	})
}

func TestGetBillSummary_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/summary",
		buildReq: func() interface{} {
			return NewGetBillSummaryApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetBillSummary(ctx, req.(*GetBillSummaryApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*GetBillSummaryApiResp).GetBillSummaryApiReply != nil },
		successBody: `{"errno":0,"errmsg":"SUCCESS","data":{"bill_id":"1125922289295589","company_name":"测试公司","bill_period":"2020-01"},"request_id":"req_001"}`,
		emptyBody:   `{"errno":0,"errmsg":"SUCCESS","data":null,"request_id":"req_empty"}`,
	})
}

func TestListBill_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/get",
		buildReq: func() interface{} {
			return NewListBillApiReqBuilder().ClientId("test_client").BillStatus(1).Offset(0).Length(10).Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.ListBill(ctx, req.(*ListBillApiReq), opt)
		},
		hasReply:    func(resp interface{}) bool { return resp.(*ListBillApiResp).ListBillApiReply != nil },
		successBody: `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"records":[{"bill_id":1125922289295589,"bill_status":3}],"lastId":1},"request_id":"req_001"}`,
		emptyBody:   `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"records":[],"lastId":0},"request_id":"req_empty"}`,
	})
}

func TestGetTransactionBillDetail_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/transactionDetail",
		buildReq: func() interface{} {
			return NewGetTransactionBillDetailApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetTransactionBillDetail(ctx, req.(*GetTransactionBillDetailApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetTransactionBillDetailApiResp).GetTransactionBillDetailApiReply != nil
		},
		successBody: `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"last_id":100,"transaction_list":[{"order_id":1125922289295589}]},"request_id":"req_001"}`,
		emptyBody:   `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"transaction_list":[]},"request_id":"req_empty"}`,
	})
}

func TestGetTransactionBillDetailOfTaxi_AllScenarios(t *testing.T) {
	runBillDetailScenarios(t, billDetailMethod{
		path: "/river/Bill/transactionDetail",
		buildReq: func() interface{} {
			return NewGetTransactionBillDetailOfTaxiApiReqBuilder().ClientId("test_client").BillId("1125922289295589").Build()
		},
		invoke: func(b *bill, ctx context.Context, req interface{}, opt *core.ReqOption) (interface{}, error) {
			return b.GetTransactionBillDetailOfTaxi(ctx, req.(*GetTransactionBillDetailOfTaxiApiReq), opt)
		},
		hasReply: func(resp interface{}) bool {
			return resp.(*GetTransactionBillDetailOfTaxiApiResp).GetTransactionBillDetailOfTaxiApiReply != nil
		},
		successBody: `{"errno":0,"errmsg":"SUCCESS","data":{"total":1,"last_id":100,"transaction_list":[{"order_id":1125922289295589}]},"request_id":"req_001"}`,
		emptyBody:   `{"errno":0,"errmsg":"SUCCESS","data":{"total":0,"transaction_list":[]},"request_id":"req_empty"}`,
	})
}

func TestNotGenBDOfWangYCItemCompanyRealPayCompatibility(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		want      float64
		wantNil   bool
		wantError bool
	}{
		{name: "number", payload: `{"company_real_pay":12.34}`, want: 12.34},
		{name: "numeric string", payload: `{"company_real_pay":"12.34"}`, want: 12.34},
		{name: "empty string", payload: `{"company_real_pay":""}`, wantNil: true},
		{name: "null", payload: `{"company_real_pay":null}`, wantNil: true},
		{name: "non numeric string", payload: `{"company_real_pay":"invalid"}`, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item NotGenBDOfWangYCItem
			err := json.Unmarshal([]byte(tt.payload), &item)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected unmarshaling error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if tt.wantNil {
				if item.CompanyRealPay != nil {
					t.Fatalf("CompanyRealPay = %v, want nil", item.CompanyRealPay)
				}
				return
			}
			if item.CompanyRealPay == nil || *item.CompanyRealPay != tt.want {
				t.Fatalf("CompanyRealPay = %v, want %v", item.CompanyRealPay, tt.want)
			}
		})
	}

	item := NewNotGenBDOfWangYCItemBuilder().CompanyRealPay(12.34).Build()
	if item.CompanyRealPay == nil || *item.CompanyRealPay != 12.34 {
		t.Fatalf("Builder CompanyRealPay = %v, want 12.34", item.CompanyRealPay)
	}
}
