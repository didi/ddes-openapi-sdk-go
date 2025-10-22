package main

import (
	"context"
	"encoding/json"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/member/v1"
	"log"
	"os"
)

func main() {
	ListMember() // 用户查询~员工列表（批量查询）
	//GetMemberDetail() // 用户查询~员工明细
	//GetMemberQuota()  // 员工限额查询
	//CreateMember()    // 用户新增
	//UpdateMember()    //  用户修改
	//DelMember()       // 用户删除
}

// ListMember 用户查询~员工列表（批量查询）
func ListMember() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewListMemberApiReqBuilder().
		CompanyId(companyId).
		Offset(0).
		Length(10).
		Build()
	apiResp, err := client.MemberService.V1.Member.ListMember(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.ListMemberApiReply {
		bytes, err := json.Marshal(apiResp.ListMemberApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetMemberDetail 用户查询~员工明细
func GetMemberDetail() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewGetMemberDetailApiReqBuilder().
		CompanyId(companyId).
		MemberId("1125967069970971"). // 参数member_id、phone、employee_number、third_user_id 必传其中一个；
		Build()
	apiResp, err := client.MemberService.V1.Member.GetMemberDetail(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.GetMemberDetailApiReply {
		bytes, err := json.Marshal(apiResp.GetMemberDetailApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// GetMemberQuota 员工限额查询
func GetMemberQuota() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewGetMemberQuotaApiReqBuilder().
		CompanyId(companyId).
		MemberIds("1125946987190896,1125946815403073"). // 参数member_ids 和 employee_numbers 不能同时为空
		Build()
	apiResp, err := client.MemberService.V1.Member.GetMemberQuota(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.GetMemberQuotaApiReply {
		bytes, err := json.Marshal(apiResp.GetMemberQuotaApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// CreateMember 用户新增
func CreateMember() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewCreateMemberApiReqBuilder().
		CreateMemberRequest(v1.NewCreateMemberRequestBuilder().
			CompanyId(companyId).
			DataObj(*v1.NewMemberInfoBuilder().
				Realname("SDK测试名称A").
				MemberType(0). // member_type不传或者member_type为0的时候phone必填
				Phone("13198987687").
				Build()).
			Build()).
		Build()
	apiResp, err := client.MemberService.V1.Member.CreateMember(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.CreateMemberApiReply {
		bytes, err := json.Marshal(apiResp.CreateMemberApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// UpdateMember  用户修改
func UpdateMember() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	build := v1.NewUpdateMemberApiReqBuilder().
		UpdateMemberRequest(v1.NewUpdateMemberRequestBuilder().
			CompanyId(companyId).
			MemberId(1125967069970971).
			DataObj(*v1.NewMemberInfoBuilder().
				Realname("SDK测试名称A-修改测试").
				Build()).
			Build()).
		Build()
	apiResp, err := client.MemberService.V1.Member.UpdateMember(context.Background(), build, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.UpdateMemberApiReply {
		bytes, err := json.Marshal(apiResp.UpdateMemberApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}

// DelMember 用户删除
func DelMember() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{ // 预发环境
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	apiReq := v1.NewDelMemberApiReqBuilder().
		DelMemberRequest(v1.NewDelMemberRequestBuilder().
			CompanyId(companyId).
			MemberId("1125967069970971").
			Build()).
		Build()
	apiResp, err := client.MemberService.V1.Member.DelMember(context.Background(), apiReq, nil)
	if err != nil {
		panic(err)
	}
	if nil != apiResp.DelMemberApiReply {
		bytes, err := json.Marshal(apiResp.DelMemberApiReply)
		if err != nil {
			return
		}
		log.Println(string(bytes))
	}
}
