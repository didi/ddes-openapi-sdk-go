package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/budgetcenter/v1"
	"log"
	"os"
)

func main() {
	GetBudgetCenter() // 部门或项目查询
	//CreateBudgetCenter() // 部门或项目新增
	//UpdateBudgetCenter() // 部门或项目修改
	//DelBudgetCenter()    // 部门或项目停用
}

// GetBudgetCenter 部门或项目查询
func GetBudgetCenter() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		EnableEncryption: false,
		EncryptionOption: &core.EncryptionOption{
			Ent: 1,
			Key: core.MD5([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
			//Ent: 2,
			//Key: core.Sha256([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
		},
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewGetBudgetCenterApiReqBuilder().
		CompanyId(companyId).
		Offset(0).
		Length(10).
		Build()
	apiResp, err := client.BudgetcenterService.V1.BudgetCenter.GetBudgetCenter(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.GetBudgetCenterApiReply {
		bytes, err := json.Marshal(apiResp.GetBudgetCenterApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// CreateBudgetCenter 部门或项目新增
func CreateBudgetCenter() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		EnableEncryption: false,
		EncryptionOption: &core.EncryptionOption{
			Ent: 1,
			Key: core.MD5([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
			//Ent: 2,
			//Key: core.Sha256([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
		},
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewCreateBudgetCenterApiReqBuilder().
		CreateBudgetCenterRequest(v1.NewCreateBudgetCenterRequestBuilder().
			CompanyId(companyId).
			Name("SDK测试部门A").
			Type(1).        // 1:部门 2:项目
			BudgetCycle(0). // 0：不限额；1：自然月 2：自然季度 3：自然年(其中2 3只对部门生效，需要设置白名单，须联系客户经理)
			TotalQuota("100").
			Build()).
		Build()
	apiResp, err := client.BudgetcenterService.V1.BudgetCenter.CreateBudgetCenter(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.CreateBudgetCenterApiReply {
		bytes, err := json.Marshal(apiResp.CreateBudgetCenterApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateBudgetCenter 部门或项目修改
func UpdateBudgetCenter() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		EnableEncryption: false,
		EncryptionOption: &core.EncryptionOption{
			Ent: 1,
			Key: core.MD5([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
			//Ent: 2,
			//Key: core.Sha256([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
		},
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewUpdateBudgetCenterApiReqBuilder().
		UpdateBudgetCenterRequest(v1.NewUpdateBudgetCenterRequestBuilder().
			CompanyId(companyId).
			Name("SDK测试部门A-up1").
			Type(1).        // 1:部门 2:项目
			BudgetCycle(0). // 0：不限额；1：自然月 2：自然季度 3：自然年(其中2 3只对部门生效，需要设置白名单，须联系客户经理)
			TotalQuota("101").
			Id("1125967024871738"). // 部门或项目ID
			Build()).
		Build()
	apiResp, err := client.BudgetcenterService.V1.BudgetCenter.UpdateBudgetCenter(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.UpdateBudgetCenterApiReply {
		bytes, err := json.Marshal(apiResp.UpdateBudgetCenterApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DelBudgetCenter 部门或项目停用
func DelBudgetCenter() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
		EnableEncryption: false,
		EncryptionOption: &core.EncryptionOption{
			Ent: 1,
			Key: core.MD5([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
			//Ent: 2,
			//Key: core.Sha256([]byte(fmt.Sprintf("es_traveler_%s", companyId))),
		},
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewDelBudgetCenterApiReqBuilder().
		DelBudgetCenterRequest(v1.NewDelBudgetCenterRequestBuilder().
			CompanyId(companyId).
			//Name("SDK测试部门A-up1").
			Type(1).                // 1:部门 2:项目
			Id("1125967024871738"). // 需要删除的部门或者项目的id,删除部门 时， id和out_budget_id 优先处理id t删除项目时 id和 out_budget_id与 name 组合唯一值时 优先处理id
			Build()).
		Build()
	apiResp, err := client.BudgetcenterService.V1.BudgetCenter.DelBudgetCenter(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.DelBudgetCenterApiReply {
		bytes, err := json.Marshal(apiResp.DelBudgetCenterApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
