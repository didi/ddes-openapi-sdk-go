

<h1 align="center">DiDi OpenApi SDK for Go</h1>


## 开发前准备
使用GO-SDK接入服务前，请确保你已在[滴滴企业开放平台](https://open.es.xiaojukeji.com/)创建了应用并安装了Go环境。


## 安装

使用 `go get` 下载安装 SDK

```sh
go get -u github.com/didi/ddes-openapi-sdk-go@latest
```

## 快速使用

开始使用前请确保已注册账号并获取到client_id、client_secret、sign_key、company_id、接口加密传输的Key等基础信息；

### 创建客户端

```go
package main

import (
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
)

func main() {
	clientId, clientSecret, signKey := "your client_id", "your client_secret", "your sign_key"
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true,               // 启用token缓存后将由SDK进行access token的获取、缓存和更新；
		LogLevel:         core.LogLevelDebug, // 查看调试信息
	})
	if err != nil {
		// handle error
		panic(err)
	}
	// TODO handle business
}

```

**core.Option参数说明**

```go
option := core.Option{
	BaseUrl:          "", // 请求地址，不设置则访问生产环境，需访问测试环境时可传入测试环境地址；
	RequestTimeOut:   0, // 超时设置
	EnableTokenCache: false, // 是否启用token缓存
	TokenCache:       nil, // 自定义缓存实现，不传则使用默认实现
	EnableEncryption: false, // 是否启用接口传输加密
	EncryptionOption: &core.EncryptionOption{
		Ent: 1, // 1: Md5 2: sha256
		Key: "your key(md5Key/sha256Key)", 
	}, // 加密传输配置参数
	LogLevel:         core.LogLevelDebug, // 日志级别
	Logger:           nil, // 自定义日志实现，不传则使用默认实现
}
```



### 启用加密传输

```go
package main

import (
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	"os"
)

func main() {
	clientId, clientSecret, signKey := "your client_id", "your client_secret", "your sign_key"
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		EnableTokenCache: true,
		EnableEncryption: true, // 是否启用加密传输
		// 加密传输配置参数
		EncryptionOption: &core.EncryptionOption{
			Ent: 1, // 1: Md5 2: sha256
			Key: "your key(md5Key/sha256Key)", 
		},
		LogLevel: core.LogLevelDebug,
	})
	if err != nil {
		// handle error
		panic(err)
	}
	// TODO handle business
}

```



### 自定义序列化

参考示例代码路径：sample/serializer_sample.go



**1、实现Serializer序列化接口**

```go
// 待实现的序列化接口
type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

// 自定义实现
type MySerializer struct{}

func (serializer *MySerializer) Serialize(v interface{}) ([]byte, error) {
	// TODO 自定义实现
	// 示例：
	jsonData, err := json.Marshal(v)
	return jsonData, err
}

func (serializer *MySerializer) Deserialize(bytes []byte, v interface{}) error {
	// TODO 自定义实现
	// 示例：
	return json.Unmarshal(bytes, v)
}

```

**2、全局自定义序列化**

```go
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
	// TODO handle business
}
```



**3、只针对单个请求的自定义序列化**

```go
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
	// TODO handle business
}
```









## 示例

### 接口认证(获取access_token)

```go
package main

import (
	"context"
	"github.com/didi/ddes-openapi-sdk-go"
	"github.com/didi/ddes-openapi-sdk-go/core"
	v1 "github.com/didi/ddes-openapi-sdk-go/service/auth/v1"
	"log"
)

func main() {
	clientId, clientSecret, signKey := "your client_id", "your client_secret", "your sign_key"
	client, err := didi.NewClientWithOption(clientId, clientSecret, signKey, &core.Option{
		LogLevel:         core.LogLevelDebug, // 查看调试信息
	})
	if err != nil {
		// handle errors
		panic(err)
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
	// TODO handle business
}

```



### 更多示例

| 示例模块     | 示例代码路径                   | 包含接口                          | 接口地址                                      |
| ------------ | ------------------------------ | --------------------------------- | --------------------------------------------- |
| 接口认证     | sample/auth_sample.go          | 接口认证(获取access_token)        | /river/Auth/authorize                         |
| 行后授权     | sample/afterapproval_sample.go | 行后审批结果同步                  | /river/AfterApproval/createPersonalReceipt    |
|              |                                | 行后审批完订单查询个人付款单      | /river/AfterApproval/getPersonalReceiptOrder  |
| 审批对接     | sample/approval_sample.go      | 创建申请单(已按业务拆分)          | /river/Approval/create                        |
|              |                                | 修改申请单(已按业务拆分)          | /river/Approval/update                        |
|              |                                | 取消申请单                        | /river/Approval/cancel                        |
|              |                                | 查询审批单列表                    | /river/Approval/getOrder                      |
|              |                                | 查询申请单详情                    | /open-apis/v1/approval/detail                 |
|              |                                | 外部审批处理                      | /river/Approval/pass                          |
| 外部审批处理 | sample/outapproval_sample.go   | 外部通知审批单状态变更            | /river/OutApproval/Status                     |
| 财税信息     | sample/bill_sample.go          | 账单列表                          | /river/Bill/get                               |
|              |                                | 已出账单(已按业务类型拆分)        | /river/Bill/detail                            |
|              |                                | 未出账单(已按业务类型拆分)        | /river/Bill/getNotGeneratedBillDetail         |
|              |                                | 网约车，出租车交易明细            | /river/Bill/transactionDetail                 |
|              |                                | 账单汇总查询-商旅、网约车、出租车 | /river/Bill/summary                           |
|              |                                | 网约车、商旅账单树查询            | /river/Bill/getBillStructure                  |
|              |                                | 商旅、网约车账单确认              | /river/Bill/confirm                           |
|              |                                | 调账提交                          | /river/Bill/adjustBillData                    |
|              |                                | 调账结果查询                      | /river/Bill/queryAdjustBillDataResult         |
| 部门或项目   | sample/budgetcenter_sample.go  | 部门或项目查询                    | /river/BudgetCenter/get                       |
|              |                                | 部门或项目新增                    | /river/BudgetCenter/add                       |
|              |                                | 部门或项目修改                    | /river/BudgetCenter/edit                      |
|              |                                | 部门或项目停用                    | /river/BudgetCenter/del                       |
| 城市         | sample/city_sample.go          | 国家查询(全量获取国际ID)          | /river/DemeterAres/Country/index              |
|              |                                | 全量开城城市列表查询              | /open-apis/v1/city/list                       |
| 地点         | sample/workspace_sample.go     | 地点新增                          | /open-apis/v1/workplace/create                |
|              |                                | 地点删除                          | /open-apis/v1/workplace/del                   |
|              |                                | 地点修改                          | /open-apis/v1/workplace/update                |
| 外部出行人   | sample/traveler_sample.go      | 外部出行人新增                    | /open-apis/v1/traveler/create                 |
|              |                                | 外部出行人修改                    | /open-apis/v1/traveler/update                 |
|              |                                | 外部出行人删除                    | /open-apis/v1/traveler/del                    |
| 角色         | sample/role_sample.go          | 角色查询                          | /river/Role/get                               |
| 制度         | sample/regulation_sample.go    | 制度列表                          | /river/Regulation/get                         |
|              |                                | 制度详情                          | /river/Regulation/detail                      |
| 职级         | sample/rank_sample.go          | 职级新增                          | /open-apis/v1/rank/create                     |
|              |                                | 职级更新                          | /open-apis/v1/rank/update                     |
|              |                                | 职级删除                          | /open-apis/v1/rank/del                        |
|              |                                | 职级查询                          | /river/Rank/getRanks                          |
| 订单         | sample/order_sample.go         | 订单号列表查询                    | /open-apis/v1/order/list                      |
|              |                                | 机票订单详情查询                  | /api-gateway/g/flight/orderDetail             |
|              |                                | 酒店订单详情查询                  | /api-gateway/g/hotel/orderDetail              |
|              |                                | 火车票订单详情查询                | /api-gateway/g/train/orderDetail              |
|              |                                | 用车订单详情查询                  | /river/Order/detail                           |
|              |                                | 用车列表                          | /river/Order/get                              |
|              |                                | 机票预估价获取                    | /api-gateway/g/flight/info/estimatePrice      |
|              |                                | 火车票直达列表                    | /api-gateway/train/queryLeftTicket            |
|              |                                | 火车票中转车次列表                | /api-gateway/g/train/transfer/queryLeftTicket |
| 用户         | sample/member_sample.go        | 用户查询~员工列表（批量查询）     | /river/Member/get                             |
|              |                                | 用户查询~员工明细                 | /river/Member/detail                          |
|              |                                | 员工限额查询                      | /river/Member/getQuota                        |
|              |                                | 用户新增                          | /river/Member/single                          |
|              |                                | 用户修改                          | /river/Member/edit                            |
|              |                                | 用户删除                          | /river/Member/del                             |
| 单点页面     | sample/login_sample.go         | 单点请求                          | /river/Login/getLoginEncryptStr               |
| 公司主体     | sample/legalentity_sample.go   | 公司主体查询                      | /river/LegalEntity/get                        |
|              |                                | 公司主体新增                      | /river/LegalEntity/add                        |
|              |                                | 公司主体修改                      | /river/LegalEntity/edit                       |
|              |                                | 公司主体停用                      | /river/LegalEntity/del                        |
| 拓展         | sample/extend_sample.go        | 拓展信息批量创建                  | /river/ExtendInfo/BatchSync                   |
|              |                                | 档案状态处理                      | /river/ExtendInfo/Status                      |
|              |                                | 拓展信息查询                      | /river/ExtendInfo/Get                         |

