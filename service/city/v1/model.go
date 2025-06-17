package v1

import (
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/url"
)

type ListAirportCityApiReq struct {
	apiReq                 *core.ApiReq
	listAirportCityRequest *ListAirportCityRequest
}
type ListAirportCityApiResp struct {
	*core.ApiResp           `json:"-"`
	ListAirportCityApiReply *ListAirportCityApiReply `json:"listAirportCityApiReply"`
}

type ListAirportCityApiReqBuilder struct {
	apiReq                 *core.ApiReq
	listAirportCityRequest *ListAirportCityRequest
}

func NewListAirportCityApiReqBuilder() *ListAirportCityApiReqBuilder {
	builder := &ListAirportCityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListAirportCityApiReqBuilder) ListAirportCityRequest(listAirportCityRequest *ListAirportCityRequest) *ListAirportCityApiReqBuilder {
	builder.listAirportCityRequest = listAirportCityRequest
	return builder
}

func (builder *ListAirportCityApiReqBuilder) Build() *ListAirportCityApiReq {
	req := &ListAirportCityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listAirportCityRequest
	return req
}

type ListCarCityApiReq struct {
	apiReq *core.ApiReq
}
type ListCarCityApiResp struct {
	*core.ApiResp       `json:"-"`
	ListCarCityApiReply *ListCarCityApiReply `json:"listCarCityApiReply"`
}

type ListCarCityApiReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListCarCityApiReqBuilder() *ListCarCityApiReqBuilder {
	builder := &ListCarCityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}

func (builder *ListCarCityApiReqBuilder) Build() *ListCarCityApiReq {
	req := &ListCarCityApiReq{}
	req.apiReq = builder.apiReq
	return req
}
func (builder *ListCarCityApiReqBuilder) ClientId(clientid string) *ListCarCityApiReqBuilder {
	builder.apiReq.QueryParams.Set("client_id", clientid)
	return builder
}
func (builder *ListCarCityApiReqBuilder) AccessToken(accesstoken string) *ListCarCityApiReqBuilder {
	builder.apiReq.QueryParams.Set("access_token", accesstoken)
	return builder
}
func (builder *ListCarCityApiReqBuilder) Timestamp(timestamp string) *ListCarCityApiReqBuilder {
	builder.apiReq.QueryParams.Set("timestamp", timestamp)
	return builder
}
func (builder *ListCarCityApiReqBuilder) CompanyId(companyid string) *ListCarCityApiReqBuilder {
	builder.apiReq.QueryParams.Set("company_id", companyid)
	return builder
}
func (builder *ListCarCityApiReqBuilder) Sign(sign string) *ListCarCityApiReqBuilder {
	builder.apiReq.QueryParams.Set("sign", sign)
	return builder
}

type ListCityApiReq struct {
	apiReq          *core.ApiReq
	listCityRequest *ListCityRequest
}
type ListCityApiResp struct {
	*core.ApiResp    `json:"-"`
	ListCityApiReply *ListCityApiReply `json:"listCityApiReply"`
}

type ListCityApiReqBuilder struct {
	apiReq          *core.ApiReq
	listCityRequest *ListCityRequest
}

func NewListCityApiReqBuilder() *ListCityApiReqBuilder {
	builder := &ListCityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListCityApiReqBuilder) ListCityRequest(listCityRequest *ListCityRequest) *ListCityApiReqBuilder {
	builder.listCityRequest = listCityRequest
	return builder
}

func (builder *ListCityApiReqBuilder) Build() *ListCityApiReq {
	req := &ListCityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listCityRequest
	return req
}

type ListCountryApiReq struct {
	apiReq             *core.ApiReq
	listCountryRequest *ListCountryRequest
}
type ListCountryApiResp struct {
	*core.ApiResp       `json:"-"`
	ListCountryApiReply *ListCountryApiReply `json:"listCountryApiReply"`
}

type ListCountryApiReqBuilder struct {
	apiReq             *core.ApiReq
	listCountryRequest *ListCountryRequest
}

func NewListCountryApiReqBuilder() *ListCountryApiReqBuilder {
	builder := &ListCountryApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListCountryApiReqBuilder) ListCountryRequest(listCountryRequest *ListCountryRequest) *ListCountryApiReqBuilder {
	builder.listCountryRequest = listCountryRequest
	return builder
}

func (builder *ListCountryApiReqBuilder) Build() *ListCountryApiReq {
	req := &ListCountryApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listCountryRequest
	return req
}

type ListHotelCityApiReq struct {
	apiReq               *core.ApiReq
	listHotelCityRequest *ListHotelCityRequest
}
type ListHotelCityApiResp struct {
	*core.ApiResp         `json:"-"`
	ListHotelCityApiReply *ListHotelCityApiReply `json:"listHotelCityApiReply"`
}

type ListHotelCityApiReqBuilder struct {
	apiReq               *core.ApiReq
	listHotelCityRequest *ListHotelCityRequest
}

func NewListHotelCityApiReqBuilder() *ListHotelCityApiReqBuilder {
	builder := &ListHotelCityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListHotelCityApiReqBuilder) ListHotelCityRequest(listHotelCityRequest *ListHotelCityRequest) *ListHotelCityApiReqBuilder {
	builder.listHotelCityRequest = listHotelCityRequest
	return builder
}

func (builder *ListHotelCityApiReqBuilder) Build() *ListHotelCityApiReq {
	req := &ListHotelCityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listHotelCityRequest
	return req
}

type ListTrainCityApiReq struct {
	apiReq               *core.ApiReq
	listTrainCityRequest *ListTrainCityRequest
}
type ListTrainCityApiResp struct {
	*core.ApiResp         `json:"-"`
	ListTrainCityApiReply *ListTrainCityApiReply `json:"listTrainCityApiReply"`
}

type ListTrainCityApiReqBuilder struct {
	apiReq               *core.ApiReq
	listTrainCityRequest *ListTrainCityRequest
}

func NewListTrainCityApiReqBuilder() *ListTrainCityApiReqBuilder {
	builder := &ListTrainCityApiReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  make(map[string]string),
		QueryParams: url.Values{},
	}
	return builder
}
func (builder *ListTrainCityApiReqBuilder) ListTrainCityRequest(listTrainCityRequest *ListTrainCityRequest) *ListTrainCityApiReqBuilder {
	builder.listTrainCityRequest = listTrainCityRequest
	return builder
}

func (builder *ListTrainCityApiReqBuilder) Build() *ListTrainCityApiReq {
	req := &ListTrainCityApiReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.listTrainCityRequest
	return req
}
