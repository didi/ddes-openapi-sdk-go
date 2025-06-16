package main

import (
	"context"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"net/http"
	"os"
)

type MySerializer struct{}

func (serializer *MySerializer) Serialize(v interface{}) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (serializer *MySerializer) Deserialize(bytes []byte, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

// 替换全局序列化实现
func globalSerializer() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		Serializer:       &MySerializer{}, // 全局使用自定义序列化
	})
	if err != nil {
		// handle error
		panic(err)
	}
	// handle business
	_, _ = client.AuthService.V1.Auth.Authorize(context.Background(), nil, nil)
}

// 替换单个请求的序列化
func singleSerializer() {
	clientId, clientSecret, signKey := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		// handle error
		panic(err)
	}
	authorizeApiResp, err := client.AuthService.V1.Auth.Authorize(context.Background(), nil, &core.ReqOption{
		Serializer: &MySerializer{}, // 仅针对本次请求的自定义序列化
	})
	if err != nil {
		return
	}
	// handle business
	if http.StatusOK == authorizeApiResp.StatusCode {

	}
}
