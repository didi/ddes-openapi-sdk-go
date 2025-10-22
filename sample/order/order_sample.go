package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/order/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	ListOrder() // 订单号列表查询

	// 订单明细
	//GetFlightOrderDetail() // 机票订单详情查询
	//GetHotelOrderDetail()  // 酒店订单详情查询
	//GetTrainOrderDetail()  // 火车票订单详情查询
	//GetCarOrderDetail()    // 用车订单详情查询

	//GetOrder()                // 用车列表
	//GetFlightEstimatePrice()  // 机票预估价获取
	//ListTrainLeftTicket()     // 火车票直达列表
	//ListTransferTrainTicket() // 火车票中转车次列表
}

// ListOrder 订单号列表查询
func ListOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListOrderApiReqBuilder().
		ListOrderRequest(
			v1.NewListOrderRequestBuilder().
				CompanyId(companyId).
				ParamJsonObj(
					*v1.NewParamJsonObjBuilder().
						OrderType("all"). // all 所有类别，domesticflight 国内机票，internationalflight 国际机票，domestichotel 国内酒店，internationalhotel 国际酒店，train 火车票，car 用车
						SearchTime(*v1.NewSearchTimeBuilder().
							Type("all").
							StartTime("2025-04-05 00:00:00").
							EndTime("2025-04-09 16:52:00").
							Build()).
						Limit(10).
						CurPage(1).
						Build()).
				Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.ListOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListOrderApiReply {
		bytes, err := json.Marshal(apiResp.ListOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetOrder 用车列表
func GetOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetOrderApiReqBuilder().
		CompanyId(companyId).
		Offset(0).
		Length(10).
		Build()
	apiResp, err := client.OrderService.V1.Order.GetOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetOrderApiReply {
		bytes, err := json.Marshal(apiResp.GetOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetFlightEstimatePrice 机票预估价获取
func GetFlightEstimatePrice() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetFlightEstimatePriceApiReqBuilder().
		GetFlightEstimatePriceRequest(
			v1.NewGetFlightEstimatePriceRequestBuilder().
				CompanyId(companyId).
				DepartureCityId("1"). // 出发城市id
				ArrivalCityId("4").   // 到达城市id
				Date("2025-05-23").   // 出发日期 （查询当天的有数据）
				//SearchType(1).
				Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.GetFlightEstimatePrice(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetFlightEstimatePriceApiReply {
		bytes, err := json.Marshal(apiResp.GetFlightEstimatePriceApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListTrainLeftTicket 火车票直达列表
func ListTrainLeftTicket() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListTrainLeftTicketApiReqBuilder().
		ListTrainLeftTicketRequest(
			v1.NewListTrainLeftTicketRequestBuilder().
				CompanyId(companyId).
				TrainDate("2025-05-24"). // 必须是今天之后的日期
				FromStationName("成都").
				ToStationName("重庆").
				Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.ListTrainLeftTicket(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListTrainLeftTicketApiReply {
		bytes, err := json.Marshal(apiResp.ListTrainLeftTicketApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListTransferTrainTicket 火车票中转车次列表
func ListTransferTrainTicket() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListTransferTrainTicketApiReqBuilder().
		ListTransferTrainTicketRequest(
			v1.NewListTransferTrainTicketRequestBuilder().
				CompanyId(companyId).
				TrainDate("2025-05-24").
				FromStationName("成都").
				ToStationName("重庆").
				CurPage(1).
				Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.ListTransferTrainTicket(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListTransferTrainTicketApiReply {
		bytes, err := json.Marshal(apiResp.ListTransferTrainTicketApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetFlightOrderDetail 机票订单详情查询
func GetFlightOrderDetail() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetFlightOrderDetailApiReqBuilder().
		GetFlightOrderDetailRequest(
			v1.NewGetFlightOrderDetailRequestBuilder().
				//CompanyId("1125923905976852").
				CompanyId("1125938266433647").
				ProductType(1).                                                 // 1：国内  2：国际
				OrderIds("1125962778746701,1125962778527503,1125962777697800"). // 【国内机票测试】订单号，多个订单用英文逗号连接，最多支持查询100个订单
				//ProductType(2).               // 1：国内  2：国际
				//OrderIds("1125966560704576"). // 【国际机票测试】订单号，多个订单用英文逗号连接，最多支持查询100个订单
				Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.GetFlightOrderDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetFlightOrderDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetFlightOrderDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetCarOrderDetail 用车订单详情查询
func GetCarOrderDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetCarOrderDetailApiReqBuilder().
		CompanyId(companyId).
		OrderId("1125953467690078"). // 用车订单ID
		Build()
	apiResp, err := client.OrderService.V1.Order.GetCarOrderDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetCarOrderDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetCarOrderDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetHotelOrderDetail 酒店订单详情查询
func GetHotelOrderDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetHotelOrderDetailApiReqBuilder().
		GetHotelOrderDetailRequest(v1.NewGetHotelOrderDetailRequestBuilder().
			CompanyId(companyId).
			ProductType(1).                                // 1：国内  2：国际
			OrderIds("1125954342995054,1125954342579658"). // 【测试国内酒店】多个订单用逗号分隔
			//ProductType(2).               // 1：国内  2：国际
			//OrderIds("1125957255257651"). // 【测试国际酒店】多个订单用逗号分隔
			Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.GetHotelOrderDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetHotelOrderDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetHotelOrderDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetTrainOrderDetail 火车票订单详情查询
func GetTrainOrderDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetTrainOrderDetailApiReqBuilder().
		GetTrainOrderDetailRequest(v1.NewGetTrainOrderDetailRequestBuilder().
			CompanyId(companyId).
			OrderIds("1125963001422088"). // [火车票订单-国内]
			Build()).
		Build()
	apiResp, err := client.OrderService.V1.Order.GetTrainOrderDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetTrainOrderDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetTrainOrderDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
