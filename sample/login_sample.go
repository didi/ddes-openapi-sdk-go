package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/login/v1"
	"log"
	"os"
)

func main() {
	GetLoginEncryptStr() // 单点请求
}

// GetLoginEncryptStr 单点请求
func GetLoginEncryptStr() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		EnableEncryption: false, // 是否启用加密传输
		EncryptionOption: &core.EncryptionOption{
			Ent: 1, // 1: md5加密，2:sha256加密
			Key: "your key(md5/sha256)",
		},
	})
	if err != nil {
		panic(err)
	}
	apiReq := v1.NewGetLoginEncryptStrApiReqBuilder().
		Phone("00016254068").
		AppType("2").
		Build()
	apiResp, err := client.LoginService.V1.Login.GetLoginEncryptStr(context.Background(), apiReq, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.GetLoginEncryptStrApiReply {
		bytes, err := json.Marshal(apiResp.GetLoginEncryptStrApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
