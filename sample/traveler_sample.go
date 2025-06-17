package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/traveler/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	CreateTraveler() // 外部出行人新增
	//UpdateTraveler() // 外部出行人修改
	//DelTraveler()    // 外部出行人删除
}

// CreateTraveler 外部出行人新增
func CreateTraveler() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCreateTravelerApiReqBuilder().
		CreateTravelerRequest(v1.NewCreateTravelerRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewTravelerInfoBuilder().
				Name("SDK测试姓名B").
				Phone("18977778888").
				Build()).
			Build()).
		Build()
	apiResp, err := client.TravelerService.V1.Traveler.CreateTraveler(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateTravelerApiReply {
		bytes, err := json.Marshal(apiResp.CreateTravelerApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateTraveler 外部出行人修改
func UpdateTraveler() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewUpdateTravelerApiReqBuilder().
		UpdateTravelerRequest(v1.NewUpdateTravelerRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewTravelerInfoBuilder().
				TravelerId("1125967098438637").
				Name("SDK测试姓名B1").
				Phone("18977778888").
				Build()).
			Build()).
		Build()
	apiResp, err := client.TravelerService.V1.Traveler.UpdateTraveler(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateTravelerApiReply {
		bytes, err := json.Marshal(apiResp.UpdateTravelerApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DelTraveler 外部出行人删除
func DelTraveler() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewDelTravelerApiReqBuilder().
		DelTravelerRequest(v1.NewDelTravelerRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewTravelerInfoBuilder().
				TravelerId("1125967098438637").
				Build()).
			Build()).
		Build()
	apiResp, err := client.TravelerService.V1.Traveler.DelTraveler(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.DelTravelerApiReply {
		bytes, err := json.Marshal(apiResp.DelTravelerApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
