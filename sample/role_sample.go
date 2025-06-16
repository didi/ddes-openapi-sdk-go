package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/role/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	ListRole() // 角色查询
}

// ListRole 角色查询
func ListRole() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewListRoleApiReqBuilder().
		CompanyId(companyId).
		Build()
	apiResp, err := client.RoleService.V1.Role.ListRole(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.ListRoleApiReply {
		bytes, err := json.Marshal(apiResp.ListRoleApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
