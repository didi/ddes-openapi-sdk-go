package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type BillConfirmApiReq struct {
	apiReq             *core.ApiReq
	billConfirmRequest *BillConfirmRequest
}
type BillConfirmApiResp struct {
	*core.ApiResp       `json:"-"`
	BillConfirmApiReply *BillConfirmApiReply `json:"billConfirmApiReply"`
}

type BillConfirmApiReqBuilder struct {
	apiReq             *core.ApiReq
	billConfirmRequest *BillConfirmRequest
}

func NewBillConfirmApiReqBuilder() *BillConfirmApiReqBuilder {
	builder := &BillConfirmApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *BillConfirmApiReqBuilder) BillConfirmRequest(billConfirmRequest *BillConfirmRequest) *BillConfirmApiReqBuilder {
	builder.billConfirmRequest = billConfirmRequest
	return builder
}

func (builder *BillConfirmApiReqBuilder) Build() *BillConfirmApiReq {
	req := &BillConfirmApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.billConfirmRequest
	return req
}

type GetAdjustBillDataResultApiReq struct {
	apiReq                         *core.ApiReq
	getAdjustBillDataResultRequest *GetAdjustBillDataResultRequest
}
type GetAdjustBillDataResultApiResp struct {
	*core.ApiResp                   `json:"-"`
	GetAdjustBillDataResultApiReply *GetAdjustBillDataResultApiReply `json:"getAdjustBillDataResultApiReply"`
}

type GetAdjustBillDataResultApiReqBuilder struct {
	apiReq                         *core.ApiReq
	getAdjustBillDataResultRequest *GetAdjustBillDataResultRequest
}

func NewGetAdjustBillDataResultApiReqBuilder() *GetAdjustBillDataResultApiReqBuilder {
	builder := &GetAdjustBillDataResultApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *GetAdjustBillDataResultApiReqBuilder) GetAdjustBillDataResultRequest(getAdjustBillDataResultRequest *GetAdjustBillDataResultRequest) *GetAdjustBillDataResultApiReqBuilder {
	builder.getAdjustBillDataResultRequest = getAdjustBillDataResultRequest
	return builder
}

func (builder *GetAdjustBillDataResultApiReqBuilder) Build() *GetAdjustBillDataResultApiReq {
	req := &GetAdjustBillDataResultApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.getAdjustBillDataResultRequest
	return req
}

type GetBillDetailOfDaiJiaApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfDaiJiaApiResp struct {
	*core.ApiResp                 `json:"-"`
	GetBillDetailOfDaiJiaApiReply *GetBillDetailOfDaiJiaApiReply `json:"getBillDetailOfDaiJiaApiReply"`
}

type GetBillDetailOfDaiJiaApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfDaiJiaApiReqBuilder() *GetBillDetailOfDaiJiaApiReqBuilder {
	builder := &GetBillDetailOfDaiJiaApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfDaiJiaApiReqBuilder) Build() *GetBillDetailOfDaiJiaApiReq {
	req := &GetBillDetailOfDaiJiaApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) ClientId(clientid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) Sign(sign string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BillId(billid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) Type(type_ int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) LastId(lastid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) Length(length int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfDaiJiaApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfDomesticFlightApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfDomesticFlightApiResp struct {
	*core.ApiResp                         `json:"-"`
	GetBillDetailOfDomesticFlightApiReply *GetBillDetailOfDomesticFlightApiReply `json:"getBillDetailOfDomesticFlightApiReply"`
}

type GetBillDetailOfDomesticFlightApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfDomesticFlightApiReqBuilder() *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder := &GetBillDetailOfDomesticFlightApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) Build() *GetBillDetailOfDomesticFlightApiReq {
	req := &GetBillDetailOfDomesticFlightApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) ClientId(clientid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) Sign(sign string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BillId(billid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) Type(type_ int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) LastId(lastid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) Length(length int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfDomesticFlightApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfDomesticFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfDomesticHotelApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfDomesticHotelApiResp struct {
	*core.ApiResp                        `json:"-"`
	GetBillDetailOfDomesticHotelApiReply *GetBillDetailOfDomesticHotelApiReply `json:"getBillDetailOfDomesticHotelApiReply"`
}

type GetBillDetailOfDomesticHotelApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfDomesticHotelApiReqBuilder() *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder := &GetBillDetailOfDomesticHotelApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) Build() *GetBillDetailOfDomesticHotelApiReq {
	req := &GetBillDetailOfDomesticHotelApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) ClientId(clientid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) Sign(sign string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BillId(billid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) Type(type_ int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) LastId(lastid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) Length(length int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfDomesticHotelApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfDomesticHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfInterFlightApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfInterFlightApiResp struct {
	*core.ApiResp                      `json:"-"`
	GetBillDetailOfInterFlightApiReply *GetBillDetailOfInterFlightApiReply `json:"getBillDetailOfInterFlightApiReply"`
}

type GetBillDetailOfInterFlightApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfInterFlightApiReqBuilder() *GetBillDetailOfInterFlightApiReqBuilder {
	builder := &GetBillDetailOfInterFlightApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfInterFlightApiReqBuilder) Build() *GetBillDetailOfInterFlightApiReq {
	req := &GetBillDetailOfInterFlightApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) ClientId(clientid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) Sign(sign string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BillId(billid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) Type(type_ int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) LastId(lastid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) Length(length int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfInterFlightApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfInterHotelApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfInterHotelApiResp struct {
	*core.ApiResp                     `json:"-"`
	GetBillDetailOfInterHotelApiReply *GetBillDetailOfInterHotelApiReply `json:"getBillDetailOfInterHotelApiReply"`
}

type GetBillDetailOfInterHotelApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfInterHotelApiReqBuilder() *GetBillDetailOfInterHotelApiReqBuilder {
	builder := &GetBillDetailOfInterHotelApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfInterHotelApiReqBuilder) Build() *GetBillDetailOfInterHotelApiReq {
	req := &GetBillDetailOfInterHotelApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) ClientId(clientid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) Sign(sign string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BillId(billid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) Type(type_ int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) LastId(lastid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) Length(length int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfInterHotelApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfManualOrderApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfManualOrderApiResp struct {
	*core.ApiResp                      `json:"-"`
	GetBillDetailOfManualOrderApiReply *GetBillDetailOfManualOrderApiReply `json:"getBillDetailOfManualOrderApiReply"`
}

type GetBillDetailOfManualOrderApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfManualOrderApiReqBuilder() *GetBillDetailOfManualOrderApiReqBuilder {
	builder := &GetBillDetailOfManualOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfManualOrderApiReqBuilder) Build() *GetBillDetailOfManualOrderApiReq {
	req := &GetBillDetailOfManualOrderApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) ClientId(clientid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) Sign(sign string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BillId(billid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) Type(type_ int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) LastId(lastid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) Length(length int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfManualOrderApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfTaxiApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfTaxiApiResp struct {
	*core.ApiResp               `json:"-"`
	GetBillDetailOfTaxiApiReply *GetBillDetailOfTaxiApiReply `json:"getBillDetailOfTaxiApiReply"`
}

type GetBillDetailOfTaxiApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfTaxiApiReqBuilder() *GetBillDetailOfTaxiApiReqBuilder {
	builder := &GetBillDetailOfTaxiApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfTaxiApiReqBuilder) Build() *GetBillDetailOfTaxiApiReq {
	req := &GetBillDetailOfTaxiApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) ClientId(clientid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) Sign(sign string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BillId(billid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) Type(type_ int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) LastId(lastid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) Length(length int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfTaxiApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfTrainTicketApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfTrainTicketApiResp struct {
	*core.ApiResp                      `json:"-"`
	GetBillDetailOfTrainTicketApiReply *GetBillDetailOfTrainTicketApiReply `json:"getBillDetailOfTrainTicketApiReply"`
}

type GetBillDetailOfTrainTicketApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfTrainTicketApiReqBuilder() *GetBillDetailOfTrainTicketApiReqBuilder {
	builder := &GetBillDetailOfTrainTicketApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfTrainTicketApiReqBuilder) Build() *GetBillDetailOfTrainTicketApiReq {
	req := &GetBillDetailOfTrainTicketApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) ClientId(clientid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) Sign(sign string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BillId(billid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) Type(type_ int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) LastId(lastid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) Length(length int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfTrainTicketApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfTrainTicketApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillDetailOfWangYCApiReq struct {
	apiReq *core.ApiReq
}
type GetBillDetailOfWangYCApiResp struct {
	*core.ApiResp                 `json:"-"`
	GetBillDetailOfWangYCApiReply *GetBillDetailOfWangYCApiReply `json:"getBillDetailOfWangYCApiReply"`
}

type GetBillDetailOfWangYCApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillDetailOfWangYCApiReqBuilder() *GetBillDetailOfWangYCApiReqBuilder {
	builder := &GetBillDetailOfWangYCApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillDetailOfWangYCApiReqBuilder) Build() *GetBillDetailOfWangYCApiReq {
	req := &GetBillDetailOfWangYCApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) ClientId(clientid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) AccessToken(accesstoken string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) CompanyId(companyid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) Timestamp(timestamp int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) Sign(sign string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BillId(billid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) DepartmentId(departmentid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("department_id", departmentid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) Type(type_ int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("type", fmt.Sprint(type_))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) LastId(lastid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) Length(length int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BusinessType(businesstype int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BookingMemberId(bookingmemberid string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BillSplitType(billsplittype int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_type", fmt.Sprint(billsplittype))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BillSplitGroupType(billsplitgrouptype int32) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_type", fmt.Sprint(billsplitgrouptype))
	return builder
}
func (builder *GetBillDetailOfWangYCApiReqBuilder) BillSplitGroupKey(billsplitgroupkey string) *GetBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_split_group_key", billsplitgroupkey)
	return builder
}

type GetBillStructureApiReq struct {
	apiReq *core.ApiReq
}
type GetBillStructureApiResp struct {
	*core.ApiResp            `json:"-"`
	GetBillStructureApiReply *GetBillStructureApiReply `json:"getBillStructureApiReply"`
}

type GetBillStructureApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillStructureApiReqBuilder() *GetBillStructureApiReqBuilder {
	builder := &GetBillStructureApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillStructureApiReqBuilder) Build() *GetBillStructureApiReq {
	req := &GetBillStructureApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillStructureApiReqBuilder) ClientId(clientid string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) AccessToken(accesstoken string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) CompanyId(companyid string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) Timestamp(timestamp string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) Sign(sign string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) PaymentPeriod(paymentperiod string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("payment_period", paymentperiod)
	return builder
}
func (builder *GetBillStructureApiReqBuilder) BusinessType(businesstype string) *GetBillStructureApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", businesstype)
	return builder
}

type GetBillSummaryApiReq struct {
	apiReq *core.ApiReq
}
type GetBillSummaryApiResp struct {
	*core.ApiResp          `json:"-"`
	GetBillSummaryApiReply *GetBillSummaryApiReply `json:"getBillSummaryApiReply"`
}

type GetBillSummaryApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetBillSummaryApiReqBuilder() *GetBillSummaryApiReqBuilder {
	builder := &GetBillSummaryApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetBillSummaryApiReqBuilder) Build() *GetBillSummaryApiReq {
	req := &GetBillSummaryApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetBillSummaryApiReqBuilder) ClientId(clientid string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) AccessToken(accesstoken string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) CompanyId(companyid string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) Timestamp(timestamp string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) Sign(sign string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) BillId(billid string) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetBillSummaryApiReqBuilder) BusinessLine(businessline int32) *GetBillSummaryApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_line", fmt.Sprint(businessline))
	return builder
}

type GetNotGenBillDetailOfDaiJiaApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfDaiJiaApiResp struct {
	*core.ApiResp                       `json:"-"`
	GetNotGenBillDetailOfDaiJiaApiReply *GetNotGenBillDetailOfDaiJiaApiReply `json:"getNotGenBillDetailOfDaiJiaApiReply"`
}

type GetNotGenBillDetailOfDaiJiaApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfDaiJiaApiReqBuilder() *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder := &GetNotGenBillDetailOfDaiJiaApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) Build() *GetNotGenBillDetailOfDaiJiaApiReq {
	req := &GetNotGenBillDetailOfDaiJiaApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfDaiJiaApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfDaiJiaApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfFlightApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfFlightApiResp struct {
	*core.ApiResp                       `json:"-"`
	GetNotGenBillDetailOfFlightApiReply *GetNotGenBillDetailOfFlightApiReply `json:"getNotGenBillDetailOfFlightApiReply"`
}

type GetNotGenBillDetailOfFlightApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfFlightApiReqBuilder() *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder := &GetNotGenBillDetailOfFlightApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) Build() *GetNotGenBillDetailOfFlightApiReq {
	req := &GetNotGenBillDetailOfFlightApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfFlightApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfHotelApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfHotelApiResp struct {
	*core.ApiResp                      `json:"-"`
	GetNotGenBillDetailOfHotelApiReply *GetNotGenBillDetailOfHotelApiReply `json:"getNotGenBillDetailOfHotelApiReply"`
}

type GetNotGenBillDetailOfHotelApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfHotelApiReqBuilder() *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder := &GetNotGenBillDetailOfHotelApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) Build() *GetNotGenBillDetailOfHotelApiReq {
	req := &GetNotGenBillDetailOfHotelApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfHotelApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfInterFlightApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfInterFlightApiResp struct {
	*core.ApiResp                            `json:"-"`
	GetNotGenBillDetailOfInterFlightApiReply *GetNotGenBillDetailOfInterFlightApiReply `json:"getNotGenBillDetailOfInterFlightApiReply"`
}

type GetNotGenBillDetailOfInterFlightApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfInterFlightApiReqBuilder() *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder := &GetNotGenBillDetailOfInterFlightApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) Build() *GetNotGenBillDetailOfInterFlightApiReq {
	req := &GetNotGenBillDetailOfInterFlightApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterFlightApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfInterFlightApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfInterHotelApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfInterHotelApiResp struct {
	*core.ApiResp                           `json:"-"`
	GetNotGenBillDetailOfInterHotelApiReply *GetNotGenBillDetailOfInterHotelApiReply `json:"getNotGenBillDetailOfInterHotelApiReply"`
}

type GetNotGenBillDetailOfInterHotelApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfInterHotelApiReqBuilder() *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder := &GetNotGenBillDetailOfInterHotelApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) Build() *GetNotGenBillDetailOfInterHotelApiReq {
	req := &GetNotGenBillDetailOfInterHotelApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfInterHotelApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfInterHotelApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfManualOrderApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfManualOrderApiResp struct {
	*core.ApiResp                            `json:"-"`
	GetNotGenBillDetailOfManualOrderApiReply *GetNotGenBillDetailOfManualOrderApiReply `json:"getNotGenBillDetailOfManualOrderApiReply"`
}

type GetNotGenBillDetailOfManualOrderApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfManualOrderApiReqBuilder() *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder := &GetNotGenBillDetailOfManualOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) Build() *GetNotGenBillDetailOfManualOrderApiReq {
	req := &GetNotGenBillDetailOfManualOrderApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfManualOrderApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfManualOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfTaxiApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfTaxiApiResp struct {
	*core.ApiResp                     `json:"-"`
	GetNotGenBillDetailOfTaxiApiReply *GetNotGenBillDetailOfTaxiApiReply `json:"getNotGenBillDetailOfTaxiApiReply"`
}

type GetNotGenBillDetailOfTaxiApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfTaxiApiReqBuilder() *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder := &GetNotGenBillDetailOfTaxiApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) Build() *GetNotGenBillDetailOfTaxiApiReq {
	req := &GetNotGenBillDetailOfTaxiApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfTaxiApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfTrainApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfTrainApiResp struct {
	*core.ApiResp                      `json:"-"`
	GetNotGenBillDetailOfTrainApiReply *GetNotGenBillDetailOfTrainApiReply `json:"getNotGenBillDetailOfTrainApiReply"`
}

type GetNotGenBillDetailOfTrainApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfTrainApiReqBuilder() *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder := &GetNotGenBillDetailOfTrainApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) Build() *GetNotGenBillDetailOfTrainApiReq {
	req := &GetNotGenBillDetailOfTrainApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfTrainApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfTrainApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetNotGenBillDetailOfWangYCApiReq struct {
	apiReq *core.ApiReq
}
type GetNotGenBillDetailOfWangYCApiResp struct {
	*core.ApiResp                       `json:"-"`
	GetNotGenBillDetailOfWangYCApiReply *GetNotGenBillDetailOfWangYCApiReply `json:"getNotGenBillDetailOfWangYCApiReply"`
}

type GetNotGenBillDetailOfWangYCApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNotGenBillDetailOfWangYCApiReqBuilder() *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder := &GetNotGenBillDetailOfWangYCApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) Build() *GetNotGenBillDetailOfWangYCApiReq {
	req := &GetNotGenBillDetailOfWangYCApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) ClientId(clientid string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) AccessToken(accesstoken string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) CompanyId(companyid string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) Timestamp(timestamp int32) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) Sign(sign string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) BusinessType(businesstype int32) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) StartDate(startdate string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) EndDate(enddate string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) LastId(lastid string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", lastid)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) OutRequisitionId(outrequisitionid string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_requisition_id", outrequisitionid)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) BookingMemberId(bookingmemberid string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("booking_member_id", bookingmemberid)
	return builder
}
func (builder *GetNotGenBillDetailOfWangYCApiReqBuilder) DateQueryType(datequerytype string) *GetNotGenBillDetailOfWangYCApiReqBuilder {
	builder.apiReq.QueryParams.Set("date_query_type", datequerytype)
	return builder
}

type GetTransactionBillDetailApiReq struct {
	apiReq *core.ApiReq
}
type GetTransactionBillDetailApiResp struct {
	*core.ApiResp                    `json:"-"`
	GetTransactionBillDetailApiReply *GetTransactionBillDetailApiReply `json:"getTransactionBillDetailApiReply"`
}

type GetTransactionBillDetailApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetTransactionBillDetailApiReqBuilder() *GetTransactionBillDetailApiReqBuilder {
	builder := &GetTransactionBillDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetTransactionBillDetailApiReqBuilder) Build() *GetTransactionBillDetailApiReq {
	req := &GetTransactionBillDetailApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetTransactionBillDetailApiReqBuilder) ClientId(clientid string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) AccessToken(accesstoken string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) CompanyId(companyid string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) Timestamp(timestamp string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) Sign(sign string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) BillId(billid string) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) LastId(lastid int32) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", fmt.Sprint(lastid))
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) Length(length int32) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetTransactionBillDetailApiReqBuilder) BusinessType(businesstype int32) *GetTransactionBillDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}

type GetTransactionBillDetailOfTaxiApiReq struct {
	apiReq *core.ApiReq
}
type GetTransactionBillDetailOfTaxiApiResp struct {
	*core.ApiResp                          `json:"-"`
	GetTransactionBillDetailOfTaxiApiReply *GetTransactionBillDetailOfTaxiApiReply `json:"getTransactionBillDetailOfTaxiApiReply"`
}

type GetTransactionBillDetailOfTaxiApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetTransactionBillDetailOfTaxiApiReqBuilder() *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder := &GetTransactionBillDetailOfTaxiApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) Build() *GetTransactionBillDetailOfTaxiApiReq {
	req := &GetTransactionBillDetailOfTaxiApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) ClientId(clientid string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) AccessToken(accesstoken string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) CompanyId(companyid string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) Timestamp(timestamp string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) Sign(sign string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) BillId(billid string) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_id", billid)
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) LastId(lastid int32) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("last_id", fmt.Sprint(lastid))
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) Length(length int32) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetTransactionBillDetailOfTaxiApiReqBuilder) BusinessType(businesstype int32) *GetTransactionBillDetailOfTaxiApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_type", fmt.Sprint(businesstype))
	return builder
}

type ListBillApiReq struct {
	apiReq *core.ApiReq
}
type ListBillApiResp struct {
	*core.ApiResp    `json:"-"`
	ListBillApiReply *ListBillApiReply `json:"listBillApiReply"`
}

type ListBillApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListBillApiReqBuilder() *ListBillApiReqBuilder {
	builder := &ListBillApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListBillApiReqBuilder) Build() *ListBillApiReq {
	req := &ListBillApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListBillApiReqBuilder) ClientId(clientid string) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListBillApiReqBuilder) AccessToken(accesstoken string) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListBillApiReqBuilder) Timestamp(timestamp string) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListBillApiReqBuilder) CompanyId(companyid string) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListBillApiReqBuilder) BillStatus(billstatus int32) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("bill_status", fmt.Sprint(billstatus))
	return builder
}
func (builder *ListBillApiReqBuilder) Offset(offset int32) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *ListBillApiReqBuilder) Length(length int32) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *ListBillApiReqBuilder) BusinessLine(businessline int32) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("business_line", fmt.Sprint(businessline))
	return builder
}
func (builder *ListBillApiReqBuilder) Sign(sign string) *ListBillApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}

type UpdateAdjustBillDataApiReq struct {
	apiReq                      *core.ApiReq
	updateAdjustBillDataRequest *UpdateAdjustBillDataRequest
}
type UpdateAdjustBillDataApiResp struct {
	*core.ApiResp                `json:"-"`
	UpdateAdjustBillDataApiReply *UpdateAdjustBillDataApiReply `json:"updateAdjustBillDataApiReply"`
}

type UpdateAdjustBillDataApiReqBuilder struct {
	apiReq                      *core.ApiReq
	updateAdjustBillDataRequest *UpdateAdjustBillDataRequest
}

func NewUpdateAdjustBillDataApiReqBuilder() *UpdateAdjustBillDataApiReqBuilder {
	builder := &UpdateAdjustBillDataApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *UpdateAdjustBillDataApiReqBuilder) UpdateAdjustBillDataRequest(updateAdjustBillDataRequest *UpdateAdjustBillDataRequest) *UpdateAdjustBillDataApiReqBuilder {
	builder.updateAdjustBillDataRequest = updateAdjustBillDataRequest
	return builder
}

func (builder *UpdateAdjustBillDataApiReqBuilder) Build() *UpdateAdjustBillDataApiReq {
	req := &UpdateAdjustBillDataApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.updateAdjustBillDataRequest
	return req
}
