package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/workspace/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	CreateWorkplace() // 地点新增
	//DeleteWorkplace() // 地点删除
	//UpdateWorkplace() // 地点修改
}

// CreateWorkplace 地点新增
func CreateWorkplace() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCreateWorkplaceApiReqBuilder().
		CreateWorkplaceRequest(v1.NewCreateWorkplaceRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewWorkplaceInfoBuilder().
				CityId(17).
				Address("华润大厦-ByGo").
				Lng("104.1136").
				Lat("30.6525").
				PointRange(1000).
				Build()).
			Build()).
		Build()
	apiResp, err := client.WorkspaceService.V1.Workspace.CreateWorkplace(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateWorkplaceApiReply {
		bytes, err := json.Marshal(apiResp.CreateWorkplaceApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DeleteWorkplace 地点删除
func DeleteWorkplace() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewDeleteWorkplaceApiReqBuilder().
		DeleteWorkplaceRequest(v1.NewDeleteWorkplaceRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewWorkplaceInfoBuilder().
				Id("1125967181614189").
				Build()).
			Build()).
		Build()
	apiResp, err := client.WorkspaceService.V1.Workspace.DeleteWorkplace(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.DeleteWorkplaceApiReply {
		bytes, err := json.Marshal(apiResp.DeleteWorkplaceApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateWorkplace 地点修改
func UpdateWorkplace() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewUpdateWorkplaceApiReqBuilder().
		UpdateWorkplaceRequest(v1.NewUpdateWorkplaceRequestBuilder().
			CompanyId(companyId).
			ParamJsonObj(*v1.NewWorkplaceInfoBuilder().
				CityId(17).
				Address("华润大厦-ByGo1").
				Lng("104.1136").
				Lat("30.6525").
				PointRange(1000).
				Build()).
			Build()).
		Build()
	apiResp, err := client.WorkspaceService.V1.Workspace.UpdateWorkplace(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateWorkplaceApiReply {
		bytes, err := json.Marshal(apiResp.UpdateWorkplaceApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
