package v1

import (
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type GetCarOrderDetailApiReq struct {
	apiReq *core.ApiReq
}
type GetCarOrderDetailApiResp struct {
	*core.ApiResp             `json:"-"`
	GetCarOrderDetailApiReply *GetCarOrderDetailApiReply `json:"getCarOrderDetailApiReply"`
}

type GetCarOrderDetailApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetCarOrderDetailApiReqBuilder() *GetCarOrderDetailApiReqBuilder {
	builder := &GetCarOrderDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetCarOrderDetailApiReqBuilder) Build() *GetCarOrderDetailApiReq {
	req := &GetCarOrderDetailApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetCarOrderDetailApiReqBuilder) ClientId(clientid string) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) AccessToken(accesstoken string) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) CompanyId(companyid string) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) Timestamp(timestamp int32) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", fmt.Sprint(timestamp))
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) OrderId(orderid string) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("order_id", orderid)
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) NeedAbnormalMsg(needabnormalmsg int32) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_abnormal_msg", fmt.Sprint(needabnormalmsg))
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) NeedRuleInfo(needruleinfo int32) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_rule_info", fmt.Sprint(needruleinfo))
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) NeedCallEmployeeNumber(needcallemployeenumber int32) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_call_employee_number", fmt.Sprint(needcallemployeenumber))
	return builder
}
func (builder *GetCarOrderDetailApiReqBuilder) Sign(sign string) *GetCarOrderDetailApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}

type GetFlightEstimatePriceApiReq struct {
	apiReq                        *core.ApiReq
	getFlightEstimatePriceRequest *GetFlightEstimatePriceRequest
}
type GetFlightEstimatePriceApiResp struct {
	*core.ApiResp                  `json:"-"`
	GetFlightEstimatePriceApiReply *GetFlightEstimatePriceApiReply `json:"getFlightEstimatePriceApiReply"`
}

type GetFlightEstimatePriceApiReqBuilder struct {
	apiReq                        *core.ApiReq
	getFlightEstimatePriceRequest *GetFlightEstimatePriceRequest
}

func NewGetFlightEstimatePriceApiReqBuilder() *GetFlightEstimatePriceApiReqBuilder {
	builder := &GetFlightEstimatePriceApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *GetFlightEstimatePriceApiReqBuilder) GetFlightEstimatePriceRequest(getFlightEstimatePriceRequest *GetFlightEstimatePriceRequest) *GetFlightEstimatePriceApiReqBuilder {
	builder.getFlightEstimatePriceRequest = getFlightEstimatePriceRequest
	return builder
}

func (builder *GetFlightEstimatePriceApiReqBuilder) Build() *GetFlightEstimatePriceApiReq {
	req := &GetFlightEstimatePriceApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.getFlightEstimatePriceRequest
	return req
}

type GetFlightOrderDetailApiReq struct {
	apiReq                      *core.ApiReq
	getFlightOrderDetailRequest *GetFlightOrderDetailRequest
}
type GetFlightOrderDetailApiResp struct {
	*core.ApiResp                `json:"-"`
	GetFlightOrderDetailApiReply *GetFlightOrderDetailApiReply `json:"getFlightOrderDetailApiReply"`
}

type GetFlightOrderDetailApiReqBuilder struct {
	apiReq                      *core.ApiReq
	getFlightOrderDetailRequest *GetFlightOrderDetailRequest
}

func NewGetFlightOrderDetailApiReqBuilder() *GetFlightOrderDetailApiReqBuilder {
	builder := &GetFlightOrderDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *GetFlightOrderDetailApiReqBuilder) GetFlightOrderDetailRequest(getFlightOrderDetailRequest *GetFlightOrderDetailRequest) *GetFlightOrderDetailApiReqBuilder {
	builder.getFlightOrderDetailRequest = getFlightOrderDetailRequest
	return builder
}

func (builder *GetFlightOrderDetailApiReqBuilder) Build() *GetFlightOrderDetailApiReq {
	req := &GetFlightOrderDetailApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.getFlightOrderDetailRequest
	return req
}

type GetHotelOrderDetailApiReq struct {
	apiReq                     *core.ApiReq
	getHotelOrderDetailRequest *GetHotelOrderDetailRequest
}
type GetHotelOrderDetailApiResp struct {
	*core.ApiResp               `json:"-"`
	GetHotelOrderDetailApiReply *GetHotelOrderDetailApiReply `json:"getHotelOrderDetailApiReply"`
}

type GetHotelOrderDetailApiReqBuilder struct {
	apiReq                     *core.ApiReq
	getHotelOrderDetailRequest *GetHotelOrderDetailRequest
}

func NewGetHotelOrderDetailApiReqBuilder() *GetHotelOrderDetailApiReqBuilder {
	builder := &GetHotelOrderDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *GetHotelOrderDetailApiReqBuilder) GetHotelOrderDetailRequest(getHotelOrderDetailRequest *GetHotelOrderDetailRequest) *GetHotelOrderDetailApiReqBuilder {
	builder.getHotelOrderDetailRequest = getHotelOrderDetailRequest
	return builder
}

func (builder *GetHotelOrderDetailApiReqBuilder) Build() *GetHotelOrderDetailApiReq {
	req := &GetHotelOrderDetailApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.getHotelOrderDetailRequest
	return req
}

type GetOrderApiReq struct {
	apiReq *core.ApiReq
}
type GetOrderApiResp struct {
	*core.ApiResp    `json:"-"`
	GetOrderApiReply *GetOrderApiReply `json:"getOrderApiReply"`
}

type GetOrderApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetOrderApiReqBuilder() *GetOrderApiReqBuilder {
	builder := &GetOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *GetOrderApiReqBuilder) Build() *GetOrderApiReq {
	req := &GetOrderApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *GetOrderApiReqBuilder) ClientId(clientid string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *GetOrderApiReqBuilder) AccessToken(accesstoken string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *GetOrderApiReqBuilder) Timestamp(timestamp string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *GetOrderApiReqBuilder) CompanyId(companyid string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *GetOrderApiReqBuilder) CallPhone(callphone string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("call_phone", callphone)
	return builder
}
func (builder *GetOrderApiReqBuilder) Phone(phone string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("phone", phone)
	return builder
}
func (builder *GetOrderApiReqBuilder) StartDate(startdate string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_date", startdate)
	return builder
}
func (builder *GetOrderApiReqBuilder) EndDate(enddate string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_date", enddate)
	return builder
}
func (builder *GetOrderApiReqBuilder) StartTime(starttime string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("start_time", starttime)
	return builder
}
func (builder *GetOrderApiReqBuilder) EndTime(endtime string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("end_time", endtime)
	return builder
}
func (builder *GetOrderApiReqBuilder) UseCarType(usecartype int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("use_car_type", fmt.Sprint(usecartype))
	return builder
}
func (builder *GetOrderApiReqBuilder) PayType(paytype int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("pay_type", fmt.Sprint(paytype))
	return builder
}
func (builder *GetOrderApiReqBuilder) IsInvoice(isinvoice int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("is_invoice", fmt.Sprint(isinvoice))
	return builder
}
func (builder *GetOrderApiReqBuilder) BudgetCenterId(budgetcenterid string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("budget_center_id", budgetcenterid)
	return builder
}
func (builder *GetOrderApiReqBuilder) OutBudgetId(outbudgetid string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("out_budget_id", outbudgetid)
	return builder
}
func (builder *GetOrderApiReqBuilder) Name(name string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("name", name)
	return builder
}
func (builder *GetOrderApiReqBuilder) Offset(offset int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("offset", fmt.Sprint(offset))
	return builder
}
func (builder *GetOrderApiReqBuilder) Length(length int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("length", fmt.Sprint(length))
	return builder
}
func (builder *GetOrderApiReqBuilder) NeedApprovalId(needapprovalid int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_approval_id", fmt.Sprint(needapprovalid))
	return builder
}
func (builder *GetOrderApiReqBuilder) NeedRuleInfo(needruleinfo int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_rule_info", fmt.Sprint(needruleinfo))
	return builder
}
func (builder *GetOrderApiReqBuilder) NeedAbnormalMsg(needabnormalmsg int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_abnormal_msg", fmt.Sprint(needabnormalmsg))
	return builder
}
func (builder *GetOrderApiReqBuilder) NeedProjectInfo(needprojectinfo int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_project_info", fmt.Sprint(needprojectinfo))
	return builder
}
func (builder *GetOrderApiReqBuilder) Sign(sign string) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}
func (builder *GetOrderApiReqBuilder) NeedCallEmployeeNumber(needcallemployeenumber int32) *GetOrderApiReqBuilder {
	builder.apiReq.QueryParams.Set("need_call_employee_number", fmt.Sprint(needcallemployeenumber))
	return builder
}

type GetTrainOrderDetailApiReq struct {
	apiReq                     *core.ApiReq
	getTrainOrderDetailRequest *GetTrainOrderDetailRequest
}
type GetTrainOrderDetailApiResp struct {
	*core.ApiResp               `json:"-"`
	GetTrainOrderDetailApiReply *GetTrainOrderDetailApiReply `json:"getTrainOrderDetailApiReply"`
}

type GetTrainOrderDetailApiReqBuilder struct {
	apiReq                     *core.ApiReq
	getTrainOrderDetailRequest *GetTrainOrderDetailRequest
}

func NewGetTrainOrderDetailApiReqBuilder() *GetTrainOrderDetailApiReqBuilder {
	builder := &GetTrainOrderDetailApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *GetTrainOrderDetailApiReqBuilder) GetTrainOrderDetailRequest(getTrainOrderDetailRequest *GetTrainOrderDetailRequest) *GetTrainOrderDetailApiReqBuilder {
	builder.getTrainOrderDetailRequest = getTrainOrderDetailRequest
	return builder
}

func (builder *GetTrainOrderDetailApiReqBuilder) Build() *GetTrainOrderDetailApiReq {
	req := &GetTrainOrderDetailApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.getTrainOrderDetailRequest
	return req
}

type ListOrderApiReq struct {
	apiReq           *core.ApiReq
	listOrderRequest *ListOrderRequest
}
type ListOrderApiResp struct {
	*core.ApiResp     `json:"-"`
	ListOrderApiReply *ListOrderApiReply `json:"listOrderApiReply"`
}

type ListOrderApiReqBuilder struct {
	apiReq           *core.ApiReq
	listOrderRequest *ListOrderRequest
}

func NewListOrderApiReqBuilder() *ListOrderApiReqBuilder {
	builder := &ListOrderApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListOrderApiReqBuilder) ListOrderRequest(listOrderRequest *ListOrderRequest) *ListOrderApiReqBuilder {
	builder.listOrderRequest = listOrderRequest
	return builder
}

func (builder *ListOrderApiReqBuilder) Build() *ListOrderApiReq {
	req := &ListOrderApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listOrderRequest
	return req
}

type ListTrainLeftTicketApiReq struct {
	apiReq                     *core.ApiReq
	listTrainLeftTicketRequest *ListTrainLeftTicketRequest
}
type ListTrainLeftTicketApiResp struct {
	*core.ApiResp               `json:"-"`
	ListTrainLeftTicketApiReply *ListTrainLeftTicketApiReply `json:"listTrainLeftTicketApiReply"`
}

type ListTrainLeftTicketApiReqBuilder struct {
	apiReq                     *core.ApiReq
	listTrainLeftTicketRequest *ListTrainLeftTicketRequest
}

func NewListTrainLeftTicketApiReqBuilder() *ListTrainLeftTicketApiReqBuilder {
	builder := &ListTrainLeftTicketApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListTrainLeftTicketApiReqBuilder) ListTrainLeftTicketRequest(listTrainLeftTicketRequest *ListTrainLeftTicketRequest) *ListTrainLeftTicketApiReqBuilder {
	builder.listTrainLeftTicketRequest = listTrainLeftTicketRequest
	return builder
}

func (builder *ListTrainLeftTicketApiReqBuilder) Build() *ListTrainLeftTicketApiReq {
	req := &ListTrainLeftTicketApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listTrainLeftTicketRequest
	return req
}

type ListTransferTrainTicketApiReq struct {
	apiReq                         *core.ApiReq
	listTransferTrainTicketRequest *ListTransferTrainTicketRequest
}
type ListTransferTrainTicketApiResp struct {
	*core.ApiResp                   `json:"-"`
	ListTransferTrainTicketApiReply *ListTransferTrainTicketApiReply `json:"listTransferTrainTicketApiReply"`
}

type ListTransferTrainTicketApiReqBuilder struct {
	apiReq                         *core.ApiReq
	listTransferTrainTicketRequest *ListTransferTrainTicketRequest
}

func NewListTransferTrainTicketApiReqBuilder() *ListTransferTrainTicketApiReqBuilder {
	builder := &ListTransferTrainTicketApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListTransferTrainTicketApiReqBuilder) ListTransferTrainTicketRequest(listTransferTrainTicketRequest *ListTransferTrainTicketRequest) *ListTransferTrainTicketApiReqBuilder {
	builder.listTransferTrainTicketRequest = listTransferTrainTicketRequest
	return builder
}

func (builder *ListTransferTrainTicketApiReqBuilder) Build() *ListTransferTrainTicketApiReq {
	req := &ListTransferTrainTicketApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listTransferTrainTicketRequest
	return req
}
