package didi

import (
	"context"
	"errors"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"github.com/didi/ddes-openapi-sdk-go/service/afterapproval"
	"github.com/didi/ddes-openapi-sdk-go/service/approval"
	"github.com/didi/ddes-openapi-sdk-go/service/auth"
	"github.com/didi/ddes-openapi-sdk-go/service/bill"
	"github.com/didi/ddes-openapi-sdk-go/service/budgetcenter"
	"github.com/didi/ddes-openapi-sdk-go/service/city"
	"github.com/didi/ddes-openapi-sdk-go/service/extend"
	"github.com/didi/ddes-openapi-sdk-go/service/legalentity"
	"github.com/didi/ddes-openapi-sdk-go/service/login"
	"github.com/didi/ddes-openapi-sdk-go/service/member"
	"github.com/didi/ddes-openapi-sdk-go/service/order"
	"github.com/didi/ddes-openapi-sdk-go/service/outapproval"
	"github.com/didi/ddes-openapi-sdk-go/service/rank"
	"github.com/didi/ddes-openapi-sdk-go/service/regulation"
	"github.com/didi/ddes-openapi-sdk-go/service/role"
	"github.com/didi/ddes-openapi-sdk-go/service/traveler"
	"github.com/didi/ddes-openapi-sdk-go/service/workspace"

	"net/http"
)

type Client struct {
	option *core.Option // 配置项

	WorkspaceService     *workspace.Service
	RoleService          *role.Service
	TravelerService      *traveler.Service
	AuthService          *auth.Service
	CityService          *city.Service
	ApprovalService      *approval.Service
	BudgetcenterService  *budgetcenter.Service
	BillService          *bill.Service
	LoginService         *login.Service
	ExtendService        *extend.Service
	OutapprovalService   *outapproval.Service
	LegalentityService   *legalentity.Service
	RegulationService    *regulation.Service
	MemberService        *member.Service
	AfterapprovalService *afterapproval.Service
	RankService          *rank.Service
	OrderService         *order.Service
}

func NewClient(appId, appSecret, signKey string) (*Client, error) {
	return NewClientWithOption(appId, appSecret, signKey, nil)
}
func NewClientWithOption(clientId, clientSecret, signKey string, option *core.Option) (*Client, error) {
	if "" == clientId || "" == clientSecret || "" == signKey {
		return nil, errors.New("clientId or clientSecret or signKey is empty")
	}
	client := &Client{}
	defaultOption := core.NewDefaultOption(clientId, clientSecret, signKey)
	if nil != option {
		if "" != option.BaseUrl {
			defaultOption.BaseUrl = option.BaseUrl
		}
		if option.RequestTimeOut > 0 {
			defaultOption.RequestTimeOut = option.RequestTimeOut
		}
		defaultOption.EnableTokenCache = option.EnableTokenCache
		if nil != option.TokenCache {
			defaultOption.TokenCache = option.TokenCache
			core.NewTokenCache(defaultOption)
		}
		defaultOption.EnableEncryption = option.EnableEncryption
		if nil != option.EncryptionOption {
			defaultOption.EncryptionOption = option.EncryptionOption
		}
		if option.SignMethod == 1 || option.SignMethod == 2 {
			defaultOption.SignMethod = option.SignMethod
		}
		if option.LogLevel >= core.LogLevelDebug && option.LogLevel <= core.LogLevelError {
			defaultOption.LogLevel = option.LogLevel
		}
		if nil != option.Logger {
			defaultOption.Logger = option.Logger
		} else {
			defaultOption.Logger = core.NewDefaultLogger(defaultOption.LogLevel)
		}
	}
	client.option = defaultOption

	// 初始化服务
	initService(client, client.option)
	return client, nil
}

func initService(client *Client, option *core.Option) {
	client.WorkspaceService = workspace.NewService(option)
	client.RoleService = role.NewService(option)
	client.TravelerService = traveler.NewService(option)
	client.AuthService = auth.NewService(option)
	client.CityService = city.NewService(option)
	client.ApprovalService = approval.NewService(option)
	client.BudgetcenterService = budgetcenter.NewService(option)
	client.BillService = bill.NewService(option)
	client.LoginService = login.NewService(option)
	client.ExtendService = extend.NewService(option)
	client.OutapprovalService = outapproval.NewService(option)
	client.LegalentityService = legalentity.NewService(option)
	client.RegulationService = regulation.NewService(option)
	client.MemberService = member.NewService(option)
	client.AfterapprovalService = afterapproval.NewService(option)
	client.RankService = rank.NewService(option)
	client.OrderService = order.NewService(option)
}

func (client *Client) Post(ctx context.Context, apiPath string, body interface{}, reqOption *core.ReqOption) (*core.ApiResp, error) {
	apiReq := core.ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    apiPath,
		Body:       body,
	}
	if nil != reqOption {
		apiReq.QueryParams = reqOption.QueryParams
		apiReq.PathParams = reqOption.PathParams
	}
	return client.Do(ctx, &apiReq, reqOption)
}
func (client *Client) Get(ctx context.Context, apiPath string, body interface{}, reqOption *core.ReqOption) (*core.ApiResp, error) {
	apiReq := core.ApiReq{
		HttpMethod: http.MethodGet,
		ApiPath:    apiPath,
		Body:       body,
	}
	if nil != reqOption {
		apiReq.QueryParams = reqOption.QueryParams
		apiReq.PathParams = reqOption.PathParams
	}
	return client.Do(ctx, &apiReq, reqOption)
}
func (client *Client) Delete(ctx context.Context, apiPath string, body interface{}, reqOption *core.ReqOption) (*core.ApiResp, error) {
	return client.Do(ctx, &core.ApiReq{
		HttpMethod:  http.MethodDelete,
		ApiPath:     apiPath,
		Body:        body,
		QueryParams: reqOption.QueryParams,
		PathParams:  reqOption.PathParams,
	}, reqOption)
}
func (client *Client) Put(ctx context.Context, apiPath string, body interface{}, reqOption *core.ReqOption) (*core.ApiResp, error) {
	return client.Do(ctx, &core.ApiReq{
		HttpMethod:  http.MethodPut,
		ApiPath:     apiPath,
		Body:        body,
		QueryParams: reqOption.QueryParams,
		PathParams:  reqOption.PathParams,
	}, reqOption)
}
func (client *Client) Patch(ctx context.Context, apiPath string, body interface{}, reqOption *core.ReqOption) (*core.ApiResp, error) {
	return client.Do(ctx, &core.ApiReq{
		HttpMethod:  http.MethodPatch,
		ApiPath:     apiPath,
		Body:        body,
		QueryParams: reqOption.QueryParams,
		PathParams:  reqOption.PathParams,
	}, reqOption)
}
func (client *Client) Do(ctx context.Context, apiReq *core.ApiReq, reqOption *core.ReqOption) (*core.ApiResp, error) {
	return core.Request(ctx, apiReq, client.option, reqOption)
}
