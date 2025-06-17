package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/bill/v1"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ListBill() // 账单列表

	////已出账单接口(/river/Bill/detail)，已将此接口按业务类型拆分为：国内机票、国内酒店、火车票、国际机票、国际酒店、增值人工单、网约车、代驾、出租车，需按业务类型分别调用；
	//GetBillDetailOfDomesticFlight() // 已出账单~国内机票（1125967084045554）
	//GetBillDetailOfDomesticHotel()  // 已出账单~国内酒店（1125966905240344）
	//GetBillDetailOfTrainTicket()    // 已出账单~火车票（1125967160953370）
	//GetBillDetailOfInterFlight()    // 已出账单~国际机票
	//GetBillDetailOfInterHotel()     // 已出账单~国际酒店（1125966074000462）
	//GetBillDetailOfManualOrder()    // 已出账单~增值人工单
	//GetBillDetailOfWangYC()         // 已出账单～网约车【已验证】（1125966721993062）
	//GetBillDetailOfDaiJia()         // 已出账单～代驾(1125957079305278)
	//GetBillDetailOfTaxi()           // 已出账单～出租车(1125962263860239)

	////未出账单接口(/river/Bill/getNotGeneratedBillDetail)，已将此接口按业务类型拆分为：国内机票、国内酒店、火车票、国际机票、国际酒店、增值人工单、网约车、代驾、出租车，需按业务类型分别调用；
	//GetNotGenBillDetailOfFlight()      // 未出账单~国内机票
	//GetNotGenBillDetailOfHotel()       // 未出账单~国内酒店
	//GetNotGenBillDetailOfTrain()       // 未出账单~火车票
	//GetNotGenBillDetailOfInterFlight() // 未出账单～国际机票
	//GetNotGenBillDetailOfInterHotel()  // 未出账单～国际酒店
	//GetNotGenBillDetailOfManualOrder() // 未出账单～增值手工单
	//GetNotGenBillDetailOfWangYC()      // 未出账单～网约车
	//GetNotGenBillDetailOfDaiJia()      // 未出账单～代驾
	//GetNotGenBillDetailOfTaxi()        // 未出账单～出租车

	//GetTransactionBillDetail() // 网约车，出租车交易明细
	//GetBillSummary()           // 账单汇总查询-商旅、网约车、出租车
	//GetBillStructure()         // 网约车、商旅账单树查询
	//BillConfirm()              // 商旅、网约车账单确认
	//UpdateAdjustBillData()     // 调账提交
	//GetAdjustBillDataResult()  // 调账结果查询
}

// ListBill 账单列表接口(/river/Bill/get)
func ListBill() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListBillApiReqBuilder().
		CompanyId(companyId).
		BillStatus(3).   // 1：待确认；3：已确认 商旅暂不支持
		BusinessLine(1). // 枚举值数字：0：网约车；1：商旅；100：出租车）
		Offset(0).
		Length(10).
		Build()
	apiResp, err := client.BillService.V1.Bill.ListBill(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListBillApiReply {
		bytes, err := json.Marshal(apiResp.ListBillApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}

}

// GetBillDetailOfDomesticFlight 已出账单~国内机票
func GetBillDetailOfDomesticFlight() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfDomesticFlightApiReqBuilder().
		CompanyId(companyId).
		BillId("1125967084045554").
		BusinessType(11). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfDomesticFlight(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfDomesticFlightApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfDomesticFlightApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfDomesticHotel 已出账单~国内酒店
func GetBillDetailOfDomesticHotel() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfDomesticHotelApiReqBuilder().
		CompanyId(companyId).
		BillId("1125966905240344").
		BusinessType(201). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfDomesticHotel(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfDomesticHotelApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfDomesticHotelApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfTrainTicket 已出账单~火车票
func GetBillDetailOfTrainTicket() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfTrainTicketApiReqBuilder().
		CompanyId(companyId).
		BillId("1125967160953370").
		BusinessType(203). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfTrainTicket(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfTrainTicketApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfTrainTicketApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfInterFlight 已出账单 ～ 国际机票
func GetBillDetailOfInterFlight() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfInterFlightApiReqBuilder().
		CompanyId(companyId).
		BillId("").
		BusinessType(205). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfInterFlight(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfInterFlightApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfInterFlightApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfInterHotel 已出账单～国际酒店
func GetBillDetailOfInterHotel() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfInterHotelApiReqBuilder().
		CompanyId(companyId).
		BillId("1125966074000462").
		BusinessType(204). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfInterHotel(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfInterHotelApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfInterHotelApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfManualOrder 已出账单 ～ 增值人工单
func GetBillDetailOfManualOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfManualOrderApiReqBuilder().
		CompanyId(companyId).
		BillId("1125966074000462").
		BusinessType(531). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfManualOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfManualOrderApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfManualOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfWangYC 已出账单～网约车
func GetBillDetailOfWangYC() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfWangYCApiReqBuilder().
		CompanyId(companyId).
		//BillId("1125934784288183"). // 网约车账单号～已确认账单
		BillId("1125966721993062"). // 网约车账单号～已确认账单 1125966721993062
		BusinessType(0).            // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfWangYC(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfWangYCApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfWangYCApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
func GetBillDetailOfDaiJia() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfDaiJiaApiReqBuilder().
		CompanyId(companyId).
		BillId("1125957079305278").
		BusinessType(40). // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfDaiJia(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfDaiJiaApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfDaiJiaApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillDetailOfTaxi 已出账单～出租车
func GetBillDetailOfTaxi() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillDetailOfTaxiApiReqBuilder().
		CompanyId(companyId).
		BillId("1125962263860239"). // 出租车账单号
		BusinessType(100).          // 网约车：0，代驾：40，国内机票：11，出租车：100，国内酒店：201，火车票：203，国际酒店：204，国际机票：205，增值服务：531；默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillDetailOfTaxi(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillDetailOfTaxiApiReply {
		bytes, err := json.Marshal(apiResp.GetBillDetailOfTaxiApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfFlight 未出账单~国内机票
func GetNotGenBillDetailOfFlight() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfFlightApiReqBuilder().
		CompanyId(companyId).
		BusinessType(202). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfFlight(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfFlightApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfFlightApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfHotel 未出账单～国内酒店
func GetNotGenBillDetailOfHotel() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfHotelApiReqBuilder().
		CompanyId(companyId).
		BusinessType(201). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-28").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfHotel(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfHotelApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfHotelApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfTrain 未出账单～火车票
func GetNotGenBillDetailOfTrain() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfTrainApiReqBuilder().
		CompanyId(companyId).
		BusinessType(203). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfTrain(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfTrainApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfTrainApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfInterFlight 未出账单～国际机票
func GetNotGenBillDetailOfInterFlight() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfInterFlightApiReqBuilder().
		CompanyId(companyId).
		BusinessType(205). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfInterFlight(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfInterFlightApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfInterFlightApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfInterHotel 未出账单～国际酒店
func GetNotGenBillDetailOfInterHotel() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfInterHotelApiReqBuilder().
		CompanyId(companyId).
		BusinessType(204). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfInterHotel(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfInterHotelApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfInterHotelApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfManualOrder 未出账单～增值手工单
func GetNotGenBillDetailOfManualOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfManualOrderApiReqBuilder().
		CompanyId(companyId).
		BusinessType(531). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfManualOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfManualOrderApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfManualOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfWangYC 未出账单～网约车
func GetNotGenBillDetailOfWangYC() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfWangYCApiReqBuilder().
		CompanyId(companyId).
		BusinessType(0). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-20").
		EndDate("2025-05-01").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfWangYC(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfWangYCApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfWangYCApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfDaiJia 未出账单～代驾
func GetNotGenBillDetailOfDaiJia() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfDaiJiaApiReqBuilder().
		CompanyId(companyId).
		BusinessType(40). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-01-01").
		EndDate("2025-01-30").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfDaiJia(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfDaiJiaApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfDaiJiaApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetNotGenBillDetailOfTaxi 未出账单～出租车
func GetNotGenBillDetailOfTaxi() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetNotGenBillDetailOfTaxiApiReqBuilder().
		CompanyId(companyId).
		BusinessType(100). // 业务类型【国内酒店:201,国内机票:202,火车票:203,国际酒店:204,国际机票:205,增值服务:531 网约车:0（且支持4.0网约车） 代驾:40 （且支持4.0代驾） 出租车 100】
		StartDate("2025-04-01").
		EndDate("2025-04-30").
		Build()
	apiResp, err := client.BillService.V1.Bill.GetNotGenBillDetailOfTaxi(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetNotGenBillDetailOfTaxiApiReply {
		bytes, err := json.Marshal(apiResp.GetNotGenBillDetailOfTaxiApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetTransactionBillDetail 账单交易明细
func GetTransactionBillDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetTransactionBillDetailApiReqBuilder().
		CompanyId(companyId).
		BillId("1125966721993062").
		BusinessType(0). // 品类，0：网约车，100：出租车，默认是0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetTransactionBillDetail(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}

	if nil != apiResp.GetTransactionBillDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetTransactionBillDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillSummary 账单汇总查询-商旅、网约车、出租车
func GetBillSummary() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillSummaryApiReqBuilder().
		CompanyId(companyId).
		BillId("1125966721993062").
		BusinessLine(0). // 0：网约车，1：商旅，100：出租车，默认为0
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillSummary(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillSummaryApiReply {
		bytes, err := json.Marshal(apiResp.GetBillSummaryApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetBillStructure 网约车、商旅账单树查询
func GetBillStructure() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetBillStructureApiReqBuilder().
		CompanyId(companyId).
		PaymentPeriod("2025-04-01~2025-04-30"). // 账单周期
		BusinessType("2").                      // 1：网约车；40 代驾；100 出租车；2：商旅；默认不传为网约车
		Build()
	apiResp, err := client.BillService.V1.Bill.GetBillStructure(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetBillStructureApiReply {
		bytes, err := json.Marshal(apiResp.GetBillStructureApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// BillConfirm 账单确认-网约车、商旅
func BillConfirm() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewBillConfirmApiReqBuilder().
		BillConfirmRequest(
			v1.NewBillConfirmRequestBuilder().
				CompanyId(companyId).
				BillId("1125966721993062").
				BusinessType(0). // 业务线（0：网约车；1：商旅）默认是0，不传默认网约车
				Build()).
		Build()
	apiResp, err := client.BillService.V1.Bill.BillConfirm(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.BillConfirmApiReply {
		bytes, err := json.Marshal(apiResp.BillConfirmApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateAdjustBillData 账单调整提交
func UpdateAdjustBillData() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	adjustList := make([]v1.AdjustListItem, 0)
	adjustList = append(adjustList, *v1.NewAdjustListItemBuilder().
		SubBusinessType(1).
		SubOrderId("1125963892688390").
		Build())

	apiReq := v1.NewUpdateAdjustBillDataApiReqBuilder().
		UpdateAdjustBillDataRequest(
			v1.NewUpdateAdjustBillDataRequestBuilder().
				CompanyId(companyId).
				BillId(1125966721993062).
				AdjustReqId(fmt.Sprintf("%d", time.Now().Unix())). // 自行传递的调整id,后面用于查询(成功的提交id不能重复)
				AdjustType(1).                                     // 调整类型1:订单计入下期2:订单信息调整
				BusinessType(1).                                   // 业务线（1：网约车；2：商旅, 40代驾 100 出租车）
				AdjustList(adjustList).                            // 调整集合
				Build()).
		Build()
	apiResp, err := client.BillService.V1.Bill.UpdateAdjustBillData(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 使用JSON验证响应结构
	resultJson := `{"errno":0,"errmsg":"SUCCESS","data":{"adjust_req_id":"a8655091-f453-4cfd-ba00-9f623d60752c","adjust_status":-1,"adjust_type":2,"bill_id":1125962365727024,"business_type":2,"check_pass":false,"company_id":1125947747934237,"err_msg":"\u4ee5\u4e0b\u53d1\u9001\u9519\u8bef,\u4e0d\u80fd\u901a\u8fc7:1125961277887128-203:subBusinessType:203 subOrderId:1125961277887128budgetCenterName\/budgetCenterCode\u540d\u79f0\u6216\u7f16\u53f7\u67e5\u8be2\u4e0d\u5230","err_msg_list":[{"adjust_err":"subBusinessType:203 subOrderId:1125961277887128budgetCenterName\/budgetCenterCode\u540d\u79f0\u6216\u7f16\u53f7\u67e5\u8be2\u4e0d\u5230","sub_business_type":203,"sub_order_id":"1125961277887128"}]},"request_id":"zP0OSTssYlY2UyF3j0lkOd7hSrIB2XZz0JaJ06JuKwVs5obTgjF4EtAnflxuGsYu"}`
	log.Println("resultJson = ", resultJson)
	dataApiReply := v1.UpdateAdjustBillDataApiReply{}
	if err := json.Unmarshal([]byte(resultJson), &dataApiReply); err != nil {
		log.Fatal(err)
		return
	}

	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateAdjustBillDataApiReply {
		bytes, err := json.Marshal(apiResp.UpdateAdjustBillDataApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetAdjustBillDataResult 调账结果查询
func GetAdjustBillDataResult() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetAdjustBillDataResultApiReqBuilder().
		GetAdjustBillDataResultRequest(
			v1.NewGetAdjustBillDataResultRequestBuilder().
				CompanyId(companyId).
				BillId(1125966721993062).
				AdjustReqId("abc123").
				Build()).
		Build()
	apiResp, err := client.BillService.V1.Bill.GetAdjustBillDataResult(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetAdjustBillDataResultApiReply {
		bytes, err := json.Marshal(apiResp.GetAdjustBillDataResultApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
