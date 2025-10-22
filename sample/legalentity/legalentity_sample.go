package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/legalentity/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	GetLegalEntity() // 公司主体查询
	//CreateLegalEntity() // 公司主体新增
	//UpdateLegalEntity() // 公司主体修改
	//DelLegalEntity()    // 公司主体停用
}

// GetLegalEntity 公司主体查询
func GetLegalEntity() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewGetLegalEntityApiReqBuilder().
		CompanyId(companyId).
		Build()
	apiResp, err := client.LegalentityService.V1.LegalEntity.GetLegalEntity(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.GetLegalEntityApiReply {
		bytes, err := json.Marshal(apiResp.GetLegalEntityApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// CreateLegalEntity 公司主体新增
func CreateLegalEntity() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewCreateLegalEntityApiReqBuilder().
		CreateLegalEntityRequest(v1.NewCreateLegalEntityRequestBuilder().
			CompanyId(companyId).
			Name("测试主体A").
			Build()).
		Build()
	apiResp, err := client.LegalentityService.V1.LegalEntity.CreateLegalEntity(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.CreateLegalEntityApiReply {
		bytes, err := json.Marshal(apiResp.CreateLegalEntityApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateLegalEntity 公司主体修改
func UpdateLegalEntity() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewUpdateLegalEntityApiReqBuilder().
		UpdateLegalEntityRequest(v1.NewUpdateLegalEntityRequestBuilder().
			CompanyId(companyId).
			Name("测试主体A-Update1").
			LegalEntityId("1125967001378661"). // legal_entity_id 和 out_legal_entity 二选一
			//OutLegalEntityId("").              // legal_entity_id 和 out_legal_entity 二选一
			Build()).
		Build()
	apiResp, err := client.LegalentityService.V1.LegalEntity.UpdateLegalEntity(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.UpdateLegalEntityApiReply {
		bytes, err := json.Marshal(apiResp.UpdateLegalEntityApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DelLegalEntity 公司主体停用
func DelLegalEntity() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true, // 启用TOKEN缓存
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	apiReq := v1.NewDelLegalEntityApiReqBuilder().
		DelLegalEntityRequest(v1.NewDelLegalEntityRequestBuilder().
			CompanyId(companyId).
			LegalEntityId("1125967001378661"). // legal_entity_id 和 out_legal_entity 二选一
			//OutLegalEntityId("").              // legal_entity_id 和 out_legal_entity 二选一
			Build()).
		Build()
	apiResp, err := client.LegalentityService.V1.LegalEntity.DelLegalEntity(context.Background(), apiReq, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if http.StatusOK != apiResp.StatusCode {
		log.Printf("HTTP Status Code: %d", apiResp.StatusCode)
		return
	}
	if nil != apiResp.DelLegalEntityApiReply {
		bytes, err := json.Marshal(apiResp.DelLegalEntityApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
