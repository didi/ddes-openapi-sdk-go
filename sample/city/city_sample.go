package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/city/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	ListCountry() // 国家
	//ListCity()    // 全量开城城市列表查询
}

// ListCity 全量开城城市列表查询
func ListCity() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListCityApiReqBuilder().
		ListCityRequest(v1.NewListCityRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewListCityParamObjBuilder().
				CountryId(86).
				CityId(17).
				ProductType("10"). // 枚举值：10:用车 20:机票 30:酒店 40:火车票 查询多个品类英文逗号隔开。
				Build()).
			Build()).
		Build()
	apiResp, err := client.CityService.V1.City.ListCity(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListCityApiReply {
		bytes, err := json.Marshal(apiResp.ListCityApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListCountry 国家
func ListCountry() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListCountryApiReqBuilder().
		ListCountryRequest(v1.NewListCountryRequestBuilder().
			Build()).
		Build()
	apiResp, err := client.CityService.V1.City.ListCountry(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListCountryApiReply {
		bytes, err := json.Marshal(apiResp.ListCountryApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
