package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/rank/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	CreateRank() // 职级新增
	//UpdateRank() // 职级更新
	//DelRank()    // 职级删除
	//ListRank()   // 职级查询
}

// CreateRank 职级新增
func CreateRank() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCreateRankApiReqBuilder().
		CreateRankRequest(v1.NewCreateRankRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewRankInfoBuilder().
				Name("SDK测试职级B").
				Build()).
			Build()).
		Build()
	apiResp, err := client.RankService.V1.Rank.CreateRank(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateRankApiReply {
		bytes, err := json.Marshal(apiResp.CreateRankApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateRank 职级更新
func UpdateRank() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewUpdateRankApiReqBuilder().
		UpdateRankRequest(v1.NewUpdateRankRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewRankInfoBuilder().
				RankId("1125967162701387").
				Name("SDK测试职级B-Update").
				Build()).
			Build()).
		Build()
	apiResp, err := client.RankService.V1.Rank.UpdateRank(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateRankApiReply {
		bytes, err := json.Marshal(apiResp.UpdateRankApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DelRank 职级删除
func DelRank() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewDelRankApiReqBuilder().
		DelRankRequest(v1.NewDelRankRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewRankInfoBuilder().
				RankId("1125967162701387").
				Build()).
			Build()).
		Build()
	apiResp, err := client.RankService.V1.Rank.DelRank(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.DelRankApiReply {
		bytes, err := json.Marshal(apiResp.DelRankApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListRank 职级查询
func ListRank() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListRankApiReqBuilder().
		CompanyId(companyId).
		Build()
	apiResp, err := client.RankService.V1.Rank.ListRank(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListRankApiReply {
		bytes, err := json.Marshal(apiResp.ListRankApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
