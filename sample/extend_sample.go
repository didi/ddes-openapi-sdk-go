package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/extend/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	CreateExtendBatch() // 拓展信息批量创建
	//UpdateExtendStatus() // 档案状态处理
	//ListExtend()         // 拓展信息查询
}

// CreateExtendBatch 拓展信息批量创建
func CreateExtendBatch() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	/**
	  测试用例
	  "root_code": "321",
	  "root_name": "测试档案_1",
	  "item_list": "[{\"code\":\"321321\",\"name\":\"测试子档案\"},{\"code\":\"123123\",\"name\":\"测试子档案01\"}]"
	*/
	extendInfos := make([]v1.ExtendInfo, 0)
	extendInfos = append(extendInfos, *v1.NewExtendInfoBuilder().Code("321321").Name("测试子档案").Build())
	extendInfos = append(extendInfos, *v1.NewExtendInfoBuilder().Code("123123").Name("测试子档案01").Build())

	apiReq := v1.NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(v1.NewCreateExtendBatchRequestBuilder().
			CompanyId(companyId).
			RootCode("321").
			RootName("GO测试档案_1").
			ItemListObj(extendInfos).
			Build()).
		Build()
	apiResp, err := client.ExtendService.V1.Extend.CreateExtendBatch(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateExtendBatchApiReply {
		bytes, err := json.Marshal(apiResp.CreateExtendBatchApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateExtendStatus 档案状态处理
func UpdateExtendStatus() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	extendInfos := make([]v1.ExtendInfo, 0)
	extendInfos = append(extendInfos, *v1.NewExtendInfoBuilder().Code("321321").Status(3).Build())
	//extendInfos = append(extendInfos, *v1.NewExtendInfoBuilder().Code("123123").Name("测试子档案01").Build())

	apiReq := v1.NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(v1.NewUpdateExtendStatusRequestBuilder().
			CompanyId(companyId).
			RootCode("3212").
			ItemListObj(extendInfos).
			Build()).
		Build()
	apiResp, err := client.ExtendService.V1.Extend.UpdateExtendStatus(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateExtendStatusApiReply {
		bytes, err := json.Marshal(apiResp.UpdateExtendStatusApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// ListExtend 拓展信息查询
func ListExtend() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListExtendApiReqBuilder().
		CompanyId(companyId).
		RootCode("321").
		Build()
	apiResp, err := client.ExtendService.V1.Extend.ListExtend(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListExtendApiReply {
		bytes, err := json.Marshal(apiResp.ListExtendApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
