package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/afterapproval/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	CreatePersonalReceipt() // 行后审批结果同步
	//GetPersonalReceiptOrder() // 行后审批完订单查询个人付款单
}

// CreatePersonalReceipt 行后审批结果同步
func CreatePersonalReceipt() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCreatePersonalReceiptApiReqBuilder().
		CreatePersonalReceiptRequest(v1.NewCreatePersonalReceiptRequestBuilder().
			CompanyId(companyId).
			OrderId("your order id").
			IsPass(1).
			Build()).
		Build()
	apiResp, err := client.AfterapprovalService.V1.AfterApproval.CreatePersonalReceipt(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreatePersonalReceiptApiReply {
		bytes, err := json.Marshal(apiResp.CreatePersonalReceiptApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetPersonalReceiptOrder 行后审批完订单～查询个人付款单
func GetPersonalReceiptOrder() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetPersonalReceiptOrderApiReqBuilder().
		CompanyId(companyId).
		StartDate("2024-10-18").
		EndDate("2024-10-25").
		Offset("0").
		Length("10").
		Build()
	apiResp, err := client.AfterapprovalService.V1.AfterApproval.GetPersonalReceiptOrder(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}

	if nil != apiResp.GetPersonalReceiptOrderApiReply {
		bytes, err := json.Marshal(apiResp.GetPersonalReceiptOrderApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
