package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/member/v1"
)

// ListMemberV2 集团账户（上级公司统一对接）~员工列表查询，走 /open-apis/v2/member/get。
// 普通企业用 ListMember；集团客户用本方法，可通过 belong_enterprise_name/taxpayer_no 定位子公司，
// query_sub_company=1 聚合下级公司数据，next_token 跨公司翻页（为空表示没有下一页）。
func ListMemberV2() {
	clientId, clientSecret, signKey, companyId := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SIGN_KEY"), os.Getenv("COMPANY_ID")
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true,
		LogLevel:         core.LogLevelDebug,
	})
	if err != nil {
		panic(err)
	}

	nextToken := "" // 首次查询不传
	for page := 1; ; page++ {
		builder := v1.NewListMemberV2ApiReqBuilder().
			CompanyId(companyId).
			Length(10).
			QuerySubCompany("1").        // 聚合下级公司数据；仅查当前公司可不传或传 0
			BelongEnterpriseName("某某集团") // 与 taxpayer_no 任选其一；传子公司 company_id 时可不传
		if nextToken != "" {
			builder = builder.NextToken(nextToken)
		}
		apiResp, err := client.MemberService.V1.Member.ListMemberV2(context.Background(), builder.Build(), nil)
		if err != nil {
			log.Fatalf("ListMemberV2 error: %v", err)
		}
		if apiResp.ListMemberV2ApiReply == nil || apiResp.ListMemberV2ApiReply.Data == nil {
			log.Fatalf("ListMemberV2 empty reply, http status: %d", apiResp.StatusCode)
		}
		data := apiResp.ListMemberV2ApiReply.Data
		out, _ := json.MarshalIndent(data.Records, "", "  ")
		log.Printf("page=%d company_total=%d progress=%v records=%s", page, derefInt32(data.CompanyTotal), derefFloat64(data.Progress), out)

		// 翻页结束：next_token 为空字符串或空指针
		if data.NextToken == nil || *data.NextToken == "" {
			break
		}
		nextToken = *data.NextToken
	}
}

func derefInt32(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}

func derefFloat64(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
