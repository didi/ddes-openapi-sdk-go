package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/regulation/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	ListRegulation() // 制度列表
	//GetRegulation()  // 制度详情
}

// ListRegulation 制度列表
func ListRegulation() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListRegulationApiReqBuilder().
		CompanyId(companyId).
		Build()
	apiResp, err := client.RegulationService.V1.Regulation.ListRegulation(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListRegulationApiReply {
		bytes, err := json.Marshal(apiResp.ListRegulationApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetRegulation 制度详情
func GetRegulation() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetRegulationApiReqBuilder().
		CompanyId(companyId).
		RegulationId("1125950821265772").
		Build()
	apiResp, err := client.RegulationService.V1.Regulation.GetRegulation(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetRegulationApiReply {
		bytes, err := json.Marshal(apiResp.GetRegulationApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
