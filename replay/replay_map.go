package replay

import (
	"context"

	didi "github.com/didi/ddes-openapi-sdk-go"
	v1afterapproval "github.com/didi/ddes-openapi-sdk-go/service/afterapproval/v1"
	v1approval "github.com/didi/ddes-openapi-sdk-go/service/approval/v1"
	v1auth "github.com/didi/ddes-openapi-sdk-go/service/auth/v1"
	v1bill "github.com/didi/ddes-openapi-sdk-go/service/bill/v1"
	v1budgetcenter "github.com/didi/ddes-openapi-sdk-go/service/budgetcenter/v1"
	v1city "github.com/didi/ddes-openapi-sdk-go/service/city/v1"
	v1extend "github.com/didi/ddes-openapi-sdk-go/service/extend/v1"
	v1legalentity "github.com/didi/ddes-openapi-sdk-go/service/legalentity/v1"
	v1login "github.com/didi/ddes-openapi-sdk-go/service/login/v1"
	v1member "github.com/didi/ddes-openapi-sdk-go/service/member/v1"
	v1order "github.com/didi/ddes-openapi-sdk-go/service/order/v1"
	v1rank "github.com/didi/ddes-openapi-sdk-go/service/rank/v1"
	v1regulation "github.com/didi/ddes-openapi-sdk-go/service/regulation/v1"
	v1traveler "github.com/didi/ddes-openapi-sdk-go/service/traveler/v1"
)

// replayMap 是 uri → 接口元数据 的映射表。
// 新增接口只需在此追加一条 entry。
//
// river 家族：build 走 buildRiverApiReq（逐字段反射 + 类型适配）。
// open-apis 家族：build 走 buildOpenApisApiReq（param_json 注入；平铺字段接口传 nil requestBuilder）。
// call/reply 由通用 extractReply 反射处理，无需逐接口编写。
var replayMap = map[string]replayEntry{}

func init() {
	ctx := context.Background()
	_ = ctx

	// ===== /river/ 家族 =====
	replayMap["/river/Order/detail"] = mkRiver(
		func() interface{} { return v1order.NewGetCarOrderDetailApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.OrderService.V1.Order.GetCarOrderDetail(ctx, req.(*v1order.GetCarOrderDetailApiReq), nil)
		},
	)
	replayMap["/river/Order/get"] = mkRiver(
		func() interface{} { return v1order.NewGetOrderApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.OrderService.V1.Order.GetOrder(ctx, req.(*v1order.GetOrderApiReq), nil)
		},
	)
	replayMap["/river/Member/get"] = mkRiver(
		func() interface{} { return v1member.NewListMemberApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.ListMember(ctx, req.(*v1member.ListMemberApiReq), nil)
		},
	)
	replayMap["/river/Member/detail"] = mkRiver(
		func() interface{} { return v1member.NewGetMemberDetailApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.GetMemberDetail(ctx, req.(*v1member.GetMemberDetailApiReq), nil)
		},
	)
	replayMap["/river/Member/edit"] = mkRequest("river",
		func() interface{} { return v1member.NewUpdateMemberApiReqBuilder() },
		func() interface{} { return v1member.NewUpdateMemberRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.UpdateMember(ctx, req.(*v1member.UpdateMemberApiReq), nil)
		},
	)
	replayMap["/river/Member/del"] = mkRequest("river",
		func() interface{} { return v1member.NewDelMemberApiReqBuilder() },
		func() interface{} { return v1member.NewDelMemberRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.DelMember(ctx, req.(*v1member.DelMemberApiReq), nil)
		},
	)
	replayMap["/river/Member/getQuota"] = mkRiver(
		func() interface{} { return v1member.NewGetMemberQuotaApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.GetMemberQuota(ctx, req.(*v1member.GetMemberQuotaApiReq), nil)
		},
	)
	replayMap["/river/Member/single"] = mkRequest("river",
		func() interface{} { return v1member.NewCreateMemberApiReqBuilder() },
		func() interface{} { return v1member.NewCreateMemberRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.MemberService.V1.Member.CreateMember(ctx, req.(*v1member.CreateMemberApiReq), nil)
		},
	)
	replayMap["/river/BudgetCenter/get"] = mkRiver(
		func() interface{} { return v1budgetcenter.NewGetBudgetCenterApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BudgetcenterService.V1.BudgetCenter.GetBudgetCenter(ctx, req.(*v1budgetcenter.GetBudgetCenterApiReq), nil)
		},
	)
	replayMap["/river/BudgetCenter/add"] = mkRequest("river",
		func() interface{} { return v1budgetcenter.NewCreateBudgetCenterApiReqBuilder() },
		func() interface{} { return v1budgetcenter.NewCreateBudgetCenterRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BudgetcenterService.V1.BudgetCenter.CreateBudgetCenter(ctx, req.(*v1budgetcenter.CreateBudgetCenterApiReq), nil)
		},
	)
	replayMap["/river/BudgetCenter/edit"] = mkRequest("river",
		func() interface{} { return v1budgetcenter.NewUpdateBudgetCenterApiReqBuilder() },
		func() interface{} { return v1budgetcenter.NewUpdateBudgetCenterRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BudgetcenterService.V1.BudgetCenter.UpdateBudgetCenter(ctx, req.(*v1budgetcenter.UpdateBudgetCenterApiReq), nil)
		},
	)
	replayMap["/river/BudgetCenter/del"] = mkRequest("river",
		func() interface{} { return v1budgetcenter.NewDelBudgetCenterApiReqBuilder() },
		func() interface{} { return v1budgetcenter.NewDelBudgetCenterRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BudgetcenterService.V1.BudgetCenter.DelBudgetCenter(ctx, req.(*v1budgetcenter.DelBudgetCenterApiReq), nil)
		},
	)
	replayMap["/river/Approval/getOrder"] = mkRiver(
		func() interface{} { return v1approval.NewListApprovalOrderApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.ListApprovalOrder(ctx, req.(*v1approval.ListApprovalOrderApiReq), nil)
		},
	)
	replayMap["/river/Approval/cancel"] = mkRequest("river",
		func() interface{} { return v1approval.NewCancelApprovalApiReqBuilder() },
		func() interface{} { return v1approval.NewCancelApprovalRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.CancelApproval(ctx, req.(*v1approval.CancelApprovalApiReq), nil)
		},
	)
	replayMap["/river/Approval/pass"] = mkRequest("river",
		func() interface{} { return v1approval.NewApprovalPassApiReqBuilder() },
		func() interface{} { return v1approval.NewApprovalPassRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.ApprovalPass(ctx, req.(*v1approval.ApprovalPassApiReq), nil)
		},
	)
	// /river/Approval/create 日志 in 与 CreateBusinessByDateApproval 参数最接近
	replayMap["/river/Approval/create"] = mkRequest("river",
		func() interface{} { return v1approval.NewCreateBusinessByDateApprovalApiReqBuilder() },
		func() interface{} { return v1approval.NewCreateApprovalBusinessByDateRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.CreateBusinessByDateApproval(ctx, req.(*v1approval.CreateBusinessByDateApprovalApiReq), nil)
		},
	)
	replayMap["/river/Approval/update"] = mkRequest("river",
		func() interface{} { return v1approval.NewUpdateBusinessByDateApprovalApiReqBuilder() },
		func() interface{} { return v1approval.NewUpdateApprovalBusinessByDateRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.UpdateBusinessByDateApproval(ctx, req.(*v1approval.UpdateBusinessByDateApprovalApiReq), nil)
		},
	)
	replayMap["/river/Auth/authorize"] = mkRequest("river",
		func() interface{} { return v1auth.NewAuthorizeApiReqBuilder() },
		func() interface{} { return v1auth.NewAuthorizeRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.AuthService.V1.Auth.Authorize(ctx, req.(*v1auth.AuthorizeApiReq), nil)
		},
	)
	replayMap["/river/Bill/get"] = mkRiver(
		func() interface{} { return v1bill.NewListBillApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BillService.V1.Bill.ListBill(ctx, req.(*v1bill.ListBillApiReq), nil)
		},
	)
	replayMap["/river/Bill/detail"] = mkRiver(
		func() interface{} { return v1bill.NewGetBillDetailOfWangYCApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BillService.V1.Bill.GetBillDetailOfWangYC(ctx, req.(*v1bill.GetBillDetailOfWangYCApiReq), nil)
		},
	)
	replayMap["/river/Bill/getBillStructure"] = mkRiver(
		func() interface{} { return v1bill.NewGetBillStructureApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BillService.V1.Bill.GetBillStructure(ctx, req.(*v1bill.GetBillStructureApiReq), nil)
		},
	)
	replayMap["/river/Bill/getNotGeneratedBillDetail"] = mkRiver(
		func() interface{} { return v1bill.NewGetNotGenBillDetailOfWangYCApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.BillService.V1.Bill.GetNotGenBillDetailOfWangYC(ctx, req.(*v1bill.GetNotGenBillDetailOfWangYCApiReq), nil)
		},
	)
	replayMap["/river/City/get"] = mkRiver(
		func() interface{} { return v1city.NewListCarCityApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.CityService.V1.City.ListCarCity(ctx, req.(*v1city.ListCarCityApiReq), nil)
		},
	)
	replayMap["/river/DemeterAres/Country/index"] = mkRequest("river",
		func() interface{} { return v1city.NewListCountryApiReqBuilder() },
		func() interface{} { return v1city.NewListCountryRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.CityService.V1.City.ListCountry(ctx, req.(*v1city.ListCountryApiReq), nil)
		},
	)
	replayMap["/river/DemeterAres/TrainCity"] = mkRequest("river",
		func() interface{} { return v1city.NewListTrainCityApiReqBuilder() },
		func() interface{} { return v1city.NewListTrainCityRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.CityService.V1.City.ListTrainCity(ctx, req.(*v1city.ListTrainCityApiReq), nil)
		},
	)
	replayMap["/river/ExtendInfo/BatchSync"] = mkRequest("river",
		func() interface{} { return v1extend.NewCreateExtendBatchApiReqBuilder() },
		func() interface{} { return v1extend.NewCreateExtendBatchRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ExtendService.V1.Extend.CreateExtendBatch(ctx, req.(*v1extend.CreateExtendBatchApiReq), nil)
		},
	)
	replayMap["/river/LegalEntity/get"] = mkRiver(
		func() interface{} { return v1legalentity.NewGetLegalEntityApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.LegalentityService.V1.LegalEntity.GetLegalEntity(ctx, req.(*v1legalentity.GetLegalEntityApiReq), nil)
		},
	)
	replayMap["/river/LegalEntity/add"] = mkRequest("river",
		func() interface{} { return v1legalentity.NewCreateLegalEntityApiReqBuilder() },
		func() interface{} { return v1legalentity.NewCreateLegalEntityRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.LegalentityService.V1.LegalEntity.CreateLegalEntity(ctx, req.(*v1legalentity.CreateLegalEntityApiReq), nil)
		},
	)
	replayMap["/river/LegalEntity/edit"] = mkRequest("river",
		func() interface{} { return v1legalentity.NewUpdateLegalEntityApiReqBuilder() },
		func() interface{} { return v1legalentity.NewUpdateLegalEntityRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.LegalentityService.V1.LegalEntity.UpdateLegalEntity(ctx, req.(*v1legalentity.UpdateLegalEntityApiReq), nil)
		},
	)
	replayMap["/river/Login/getLoginEncryptStr"] = mkRiver(
		func() interface{} { return v1login.NewGetLoginEncryptStrApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.LoginService.V1.Login.GetLoginEncryptStr(ctx, req.(*v1login.GetLoginEncryptStrApiReq), nil)
		},
	)
	replayMap["/river/Rank/getRanks"] = mkRiver(
		func() interface{} { return v1rank.NewListRankApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.RankService.V1.Rank.ListRank(ctx, req.(*v1rank.ListRankApiReq), nil)
		},
	)
	replayMap["/river/Regulation/detail"] = mkRiver(
		func() interface{} { return v1regulation.NewGetRegulationApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.RegulationService.V1.Regulation.GetRegulation(ctx, req.(*v1regulation.GetRegulationApiReq), nil)
		},
	)
	replayMap["/river/Regulation/get"] = mkRiver(
		func() interface{} { return v1regulation.NewListRegulationApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.RegulationService.V1.Regulation.ListRegulation(ctx, req.(*v1regulation.ListRegulationApiReq), nil)
		},
	)
	replayMap["/river/AfterApproval/createPersonalReceipt"] = mkRequest("river",
		func() interface{} { return v1afterapproval.NewCreatePersonalReceiptApiReqBuilder() },
		func() interface{} { return v1afterapproval.NewCreatePersonalReceiptRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.AfterapprovalService.V1.AfterApproval.CreatePersonalReceipt(ctx, req.(*v1afterapproval.CreatePersonalReceiptApiReq), nil)
		},
	)
	replayMap["/river/AfterApproval/getPersonalReceiptOrder"] = mkRiver(
		func() interface{} { return v1afterapproval.NewGetPersonalReceiptOrderApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.AfterapprovalService.V1.AfterApproval.GetPersonalReceiptOrder(ctx, req.(*v1afterapproval.GetPersonalReceiptOrderApiReq), nil)
		},
	)

	// ===== /open-apis/ 家族 =====
	replayMap["/open-apis/v1/order/list"] = mkRequest("open-apis",
		func() interface{} { return v1order.NewListOrderApiReqBuilder() },
		func() interface{} { return v1order.NewListOrderRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.OrderService.V1.Order.ListOrder(ctx, req.(*v1order.ListOrderApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/approval/detail"] = mkFlat("open-apis",
		func() interface{} { return v1approval.NewGetApprovalDetailApiReqBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.ApprovalService.V1.Approval.GetApprovalDetail(ctx, req.(*v1approval.GetApprovalDetailApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/rank/create"] = mkRequest("open-apis",
		func() interface{} { return v1rank.NewCreateRankApiReqBuilder() },
		func() interface{} { return v1rank.NewCreateRankRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.RankService.V1.Rank.CreateRank(ctx, req.(*v1rank.CreateRankApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/rank/update"] = mkRequest("open-apis",
		func() interface{} { return v1rank.NewUpdateRankApiReqBuilder() },
		func() interface{} { return v1rank.NewUpdateRankRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.RankService.V1.Rank.UpdateRank(ctx, req.(*v1rank.UpdateRankApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/traveler/create"] = mkRequest("open-apis",
		func() interface{} { return v1traveler.NewCreateTravelerApiReqBuilder() },
		func() interface{} { return v1traveler.NewCreateTravelerRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.TravelerService.V1.Traveler.CreateTraveler(ctx, req.(*v1traveler.CreateTravelerApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/traveler/del"] = mkRequest("open-apis",
		func() interface{} { return v1traveler.NewDelTravelerApiReqBuilder() },
		func() interface{} { return v1traveler.NewDelTravelerRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.TravelerService.V1.Traveler.DelTraveler(ctx, req.(*v1traveler.DelTravelerApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/traveler/update"] = mkRequest("open-apis",
		func() interface{} { return v1traveler.NewUpdateTravelerApiReqBuilder() },
		func() interface{} { return v1traveler.NewUpdateTravelerRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.TravelerService.V1.Traveler.UpdateTraveler(ctx, req.(*v1traveler.UpdateTravelerApiReq), nil)
		},
	)
	replayMap["/open-apis/v1/city/list"] = mkRequest("open-apis",
		func() interface{} { return v1city.NewListCityApiReqBuilder() },
		func() interface{} { return v1city.NewListCityRequestBuilder() },
		func(c *didi.Client, req interface{}) (interface{}, error) {
			return c.CityService.V1.City.ListCity(ctx, req.(*v1city.ListCityApiReq), nil)
		},
	)
}

// mkRiver 构造 river 平铺字段 entry：build 走逐字段反射，call/reply 通用。
func mkRiver(newBuilder func() interface{}, call func(*didi.Client, interface{}) (interface{}, error)) replayEntry {
	return replayEntry{
		family: "river",
		build: func(in map[string]interface{}) (interface{}, error) {
			return buildRiverApiReq(newBuilder(), in)
		},
		call:  call,
		reply: func(resp interface{}) (int32, interface{}) { e, d, _ := extractReply(resp); return e, d },
	}
}

// mkRequest 构造 POST-body entry（river POST-body 与 open-apis param_json 共用）：
// build 走 buildRequestApiReq，自动判别 param_json vs 直接字段。
func mkRequest(family string, newApiReqBuilder func() interface{}, newRequestBuilder func() interface{}, call func(*didi.Client, interface{}) (interface{}, error)) replayEntry {
	return replayEntry{
		family: family,
		build: func(in map[string]interface{}) (interface{}, error) {
			return buildRequestApiReq(newApiReqBuilder(), newRequestBuilder(), in)
		},
		call:  call,
		reply: func(resp interface{}) (int32, interface{}) { e, d, _ := extractReply(resp); return e, d },
	}
}

// mkFlat 构造平铺字段 entry（open-apis approval/detail 等）：走 river 平铺反射。
func mkFlat(family string, newBuilder func() interface{}, call func(*didi.Client, interface{}) (interface{}, error)) replayEntry {
	return replayEntry{
		family: family,
		build: func(in map[string]interface{}) (interface{}, error) {
			return buildRiverApiReq(newBuilder(), in)
		},
		call:  call,
		reply: func(resp interface{}) (int32, interface{}) { e, d, _ := extractReply(resp); return e, d },
	}
}
