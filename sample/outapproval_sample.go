package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/outapproval/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	UpdateOutApprovalStatus() // 外部通知审批单状态变更接口
}

// UpdateOutApprovalStatus 外部通知审批单状态变更接口
func UpdateOutApprovalStatus() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewUpdateOutApprovalStatusApiReqBuilder().
		UpdateOutApprovalStatusRequest(v1.NewUpdateOutApprovalStatusRequestBuilder().
			CompanyId(companyId).
			OutId("1125941460354413").
			ApprovalType(0).
			Status(0).
			Build()).
		Build()
	apiResp, err := client.OutapprovalService.V1.OutApproval.UpdateOutApprovalStatus(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateOutApprovalStatusApiReply {
		bytes, err := json.Marshal(apiResp.UpdateOutApprovalStatusApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
