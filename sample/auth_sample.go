package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/auth/v1"
	"log"
	"os"
)

func main() {
	Authorize()
}

// Authorize 接口认证
func Authorize() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: false, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	apiReq := v1.NewAuthorizeApiReqBuilder().
		AuthorizeRequest(
			v1.NewAuthorizeRequestBuilder().
				GrantType("client_credentials").
				Build()).
		Build()
	authorizeApiResp, err := client.AuthService.V1.Auth.Authorize(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	if nil == authorizeApiResp.AuthorizeApiReply {
		log.Fatal("authorize api response is nil")
		return
	}
	bytes, err := json.Marshal(authorizeApiResp.AuthorizeApiReply)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(bytes))
}
