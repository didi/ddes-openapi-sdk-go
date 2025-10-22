package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/approval/v1"
	"log"
	"net/http"
	"os"
)

// 审批对接
func main() {
	// 创建申请单，已按业务拆分
	CreateTravelApproval() // 创建申请单(差旅)
	//CreateBusinessByDateApproval()  // 创建用车申请单(按日期)
	//CreateBusinessByTimesApproval() // 创建用车申请单(按次数)

	// 修改申请单，已按业务拆分
	//UpdateTravelApproval()          // 修改申请单(差旅)
	//UpdateBusinessByDateApproval()  // 修改用车申请单(按日期)
	//UpdateBusinessByTimesApproval() // 修改用车申请单(按次数)

	//ApprovalCancel()    // 取消审批单
	//ListApprovalOrder() // 查询审批单
	//GetApprovalDetail() // 查询审批单详情
	//ApprovalPass()      // 外部审批处理
}

// ApprovalPass 外部审批处理
func ApprovalPass() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewApprovalPassApiReqBuilder().
		ApprovalPassRequest(v1.NewApprovalPassRequestBuilder().
			CompanyId(companyId).
			ObjectType(1).
			ObjectId(1125965135320697).
			ObjectApprovalType(4).
			IsPass(1).
			Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.ApprovalPass(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ApprovalPassApiReply {
		bytes, err := json.Marshal(apiResp.ApprovalPassApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// CreateTravelApproval 创建差旅申请单
func CreateTravelApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	trips := make([]v1.Trip, 0)
	cities := make([]v1.TravelCity, 0)
	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam03").
		MemberType(0).PassengerPhone("00016218408").Build())
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam04").
		MemberType(0).PassengerPhone("00016213494").Build())
	cities = append(cities, *v1.NewTravelCityBuilder().Id(17).Name("成都市").Build())
	trips = append(trips, *v1.NewTripBuilder().DepartureCity("北京市").DepartureCityId(1).
		StartDate("2025-05-22").EndDate("2025-05-22").TripType("1,2").ToCitys(cities).Build())
	apiReq := v1.NewCreateTravelApprovalApiReqBuilder().
		CreateTravelApprovalRequest(
			v1.NewCreateApprovalRequestBuilder().CompanyId(companyId).ApprovalType(1).RegulationId("1125937551844010").
				OutApprovalId("202504230002-go-sdk-test-2").Reason("sdk-test").
				TravelDetailObj(*v1.NewTravelDetailBuilder().StartDate("2025-05-22").EndDate("2025-05-22").Trips(trips).
					StartCityRule(1).EndCityRule(1).Build()).
				MemberType(0).Phone("00016213494").PassengerListObj(passengers).ExecutiveRegulationType(0).
				ExecutiveRegulationId("1125937551844010").ExecutiveRegulationMemberType(0).ExecutiveRegulationMember("00016213494").
				Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.CreateTravelApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.CreateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// CreateBusinessByDateApproval 创建用车按日期申请单
func CreateBusinessByDateApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())

	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam02").
		MemberType(0).PassengerPhone("00016244610").Build())

	apiReq := v1.NewCreateBusinessByDateApprovalApiReqBuilder().
		CreateBusinessByDateApprovalRequest(
			v1.NewCreateApprovalBusinessByDateRequestBuilder().CompanyId(companyId).
				OutApprovalId("go-sdk-test-by-date").Reason("go-sdk-test-by-date").
				ApprovalType(3). // 审批类型，1-差旅，2-用车按次数，3-用车按日期
				RegulationId("1125927807992439").
				BudgetCenterId("1125950150116339").
				Phone("11100022245").MemberType(0).
				ExtendFieldListObj(*v1.NewExtendFieldListBuilder().ExtendField01("906188916").Build()).
				BusinessTripDetailObj(*v1.NewBusinessTripDetailByDateBuilder().StartTime("2025-05-22").EndTime("2025-05-22").
					TripAmount(70882).PerorderMoneyQuota(70882).Build()).
				PassengerListObj(passengers).Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.CreateBusinessByDateApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.CreateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}

}

// CreateBusinessByTimesApproval 创建用车按次数申请单
func CreateBusinessByTimesApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())

	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam02").
		MemberType(0).PassengerPhone("00016244610").Build())

	apiReq := v1.NewCreateBusinessByTimesApprovalApiReqBuilder().
		CreateBusinessByTimesApprovalRequest(
			v1.NewCreateApprovalBusinessByTimesRequestBuilder().CompanyId(companyId).
				OutApprovalId("go-sdk-test-by-times").Reason("go-sdk-test-by-date").
				ApprovalType(2). // 审批类型，1-差旅，2-用车按次数，3-用车按日期
				RegulationId("1125928451544592").
				BudgetCenterId("1125950150116339").
				Phone("11100022245").MemberType(0).
				ExtendFieldListObj(*v1.NewExtendFieldListBuilder().ExtendField01("按次数测试").Build()).
				BusinessTripDetailObj(*v1.NewBusinessTripDetailByTimesBuilder().StartTime("2025-05-22").EndTime("2025-05-22").TripTimes(2).Build()).
				PassengerListObj(passengers).Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.CreateBusinessByTimesApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.CreateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}

}

// UpdateTravelApproval 更新差旅申请单
func UpdateTravelApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam03").
		MemberType(0).PassengerPhone("00016218408").Build())
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam04").
		MemberType(0).PassengerPhone("00016213494").Build())

	cities := make([]v1.TravelCity, 0)
	cities = append(cities, *v1.NewTravelCityBuilder().Id(17).Name("成都市").Build())

	trips := make([]v1.Trip, 0)
	trips = append(trips, *v1.NewTripBuilder().DepartureCity("北京市").DepartureCityId(1).
		StartDate("2025-05-22").EndDate("2025-05-23").TripType("1,2").ToCitys(cities).Build())

	apiReq := v1.NewUpdateApprovalApiReqBuilder().
		UpdateApprovalRequest(
			v1.NewUpdateApprovalRequestBuilder().CompanyId(companyId).
				ApprovalType(1).ApprovalId("1125967188986629").Reason("go-sdk-test-update").
				TravelDetailObj(*v1.NewTravelDetailBuilder().StartDate("2025-05-22").EndDate("2025-05-23").
					Trips(trips).StartCityRule(1).EndCityRule(1).
					Build()).
				PassengerListObj(passengers).
				Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.UpdateApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.UpdateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateBusinessByDateApproval 更新用车按日期申请单
func UpdateBusinessByDateApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam02").
		MemberType(0).PassengerPhone("00016244610").Build())

	cities := make([]v1.TravelCity, 0)
	cities = append(cities, *v1.NewTravelCityBuilder().Id(17).Name("成都市").Build())

	trips := make([]v1.Trip, 0)
	trips = append(trips, *v1.NewTripBuilder().DepartureCity("北京市").DepartureCityId(1).
		StartDate("2025-05-22").EndDate("2025-05-23").TripType("1,2").ToCitys(cities).Build())

	apiReq := v1.NewUpdateBusinessByDateApprovalApiReqBuilder().UpdateBusinessByDateApprovalRequest(
		v1.NewUpdateApprovalBusinessByDateRequestBuilder().
			CompanyId(companyId).
			ApprovalType(3).ApprovalId("1125967184604018").Reason("go-sdk-test-update-by-date").
			BusinessTripDetailObj(*v1.NewBusinessTripDetailByDateBuilder().StartTime("2025-05-22").EndTime("2025-05-22").
				TripAmount(80000).PerorderMoneyQuota(80000).
				Build()).
			PassengerListObj(passengers).
			BudgetCenterId("1125950150116339").
			ExtendFieldListObj(*v1.NewExtendFieldListBuilder().ExtendField01("906188916").Build()).
			Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.UpdateBusinessByDateApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.UpdateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateBusinessByTimesApproval 更新用车按次数申请单
func UpdateBusinessByTimesApproval() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	passengers := make([]v1.TripPassenger, 0)
	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam01").
		MemberType(0).PassengerPhone("11100022245").Build())

	passengers = append(passengers, *v1.NewTripPassengerBuilder().PassengerType(0).PassengerName("tNam02").
		MemberType(0).PassengerPhone("00016244610").Build())

	apiReq := v1.NewUpdateBusinessByTimesApprovalApiReqBuilder().
		UpdateBusinessByTimesApprovalRequest(
			v1.NewUpdateApprovalBusinessByTimesRequestBuilder().CompanyId(companyId).
				ApprovalId("1125967184585606").OutApprovalId("go-sdk-test-by-times").Reason("go-sdk-test-by-times-update").
				ApprovalType(2). // 审批类型，1-差旅，2-用车按次数，3-用车按日期
				ExtendFieldListObj(*v1.NewExtendFieldListBuilder().ExtendField01("按次数测试").Build()).
				BusinessTripDetailObj(*v1.NewBusinessTripDetailByTimesBuilder().StartTime("2025-05-22 00:00:00").EndTime("2025-05-23 12:00:00").TripTimes(2).Build()).
				PassengerListObj(passengers).Build()).
		Build()
	apiResp, err := client.ApprovalService.V1.Approval.UpdateBusinessByTimesApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateApprovalApiReply {
		bytes, err := json.Marshal(apiResp.UpdateApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ApprovalCancel 取消申请单
func ApprovalCancel() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCancelApprovalApiReqBuilder().CancelApprovalRequest(
		v1.NewCancelApprovalRequestBuilder().
			CompanyId(companyId).
			ApprovalId("1125967185832931").
			Build()).
		Build()

	apiResp, err := client.ApprovalService.V1.Approval.CancelApproval(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CancelApprovalApiReply {
		bytes, err := json.Marshal(apiResp.CancelApprovalApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListApprovalOrder 审批单查询
func ListApprovalOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListApprovalOrderApiReqBuilder().
		CompanyId(companyId).
		ApprovalId("1125960004198548").
		Offset(0).
		Length(10).Build()
	apiResp, err := client.ApprovalService.V1.Approval.ListApprovalOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListApprovalOrderApiReply {
		bytes, err := json.Marshal(apiResp.ListApprovalOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetApprovalDetail 申请单详情接口
func GetApprovalDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetApprovalDetailApiReqBuilder().
		CompanyId(companyId).
		ApprovalId("1125967184585606").Build()
	apiResp, err := client.ApprovalService.V1.Approval.GetApprovalDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetApprovalDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetApprovalDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
