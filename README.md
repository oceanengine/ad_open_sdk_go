# Marketing Go SDK

## 概述
巨量引擎开放平台 Marketing API(以下简称API) SDK 提供了Token获取、请求封装、响应解释等功能，以本地化方式轻松完成API的调用和结果的获取，旨在帮助开发者快速搭建投放管理系统。

## 使用条件
1. 使用SDK需要首先注册成为巨量引擎开发者，请参考[开发者快速入门文档](https://open.oceanengine.com/labels/7/docs/1696710498372623)
2. 使用SDK需要先拥有API的访问权限，所有SDK的使用与应用拥有的权限组相关联

## 安装

```shell
go get git@github.com:oceanengine/ad_open_sdk_go.git
```

import:

```golang
import ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go
```

## 使用

### sdk包结构

```golang
.
├── LICENSE
├── README.md
├── api
│   ├── api_project_create_v30.go
│   ├── api_xxx.go
│   ├── client.go
│   └── middleware.go
├── client.go
├── config
│   └── configuration.go
├── examples
│   ├── promotion_create_v30_example.go
│   └── xxx_example.go
├── git_push.sh
├── go.mod
├── go.sum
├── middleware
│   ├── auth.go
│   ├── header.go
│   └── log.go
├── models
│   ├── model_promotion_create_v3_0_request.go
│   ├── model_xxx_response.go
│   └── utils.go
└── test
```
- client.go 为 sdk 的入口文件，包含 sdk 的初始化以及 api 对象获取函数
- api 目录下包含了生成的 api 调用代码逻辑、api client 以及中间件逻辑
- config 目录包含sdk配置结构的声明
- example 目录中包含了每个接口的调用样例代码
- middleware 目录中包含默认中间件逻辑
- models 目录中包含每个接口的模型、枚举定义

### 配置

使用默认配置:
```golang
import "github.com/oceanengine/ad_open_sdk_go/config"

configuration := config.NewConfiguration()
```

配置定义如下:

```golang
type Configuration struct {
    Host          string            `json:"host,omitempty"`
    Scheme        string            `json:"scheme,omitempty"`
    DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
    UserAgent     string            `json:"userAgent,omitempty"`
    LogEnable     bool              `json:"log_enable,omitempty"`
    HTTPClient    *http.Client      `json:"-"`
}
```

### 初始化client
> 可参考下面的代码初始化 client，并可修改 client 的属性

```golang
func main() {
    configuration := config.NewConfiguration()
    configuration.LogEnable = true    // 启用日志
    configuration.HTTPClient = CustomeClient // 使用自定义 http client
    apiClient := ad_open_sdk_go.Init(configuration)
    apiClient.Use(xxxMiddleware)    // 增加中间件
    apiClient.SetLogEnable(true)    // 初始化后修改是否启用日志
    apiClient.SetHost("xxx.com")    // 修改域名
    apiClient.AddDefaultHeader("X-Debug-Header", "1")    // 为每个请求增加默认 header
    apiClient.ApiClient.Cfg.xxx = xxx         // 修改传入的config
}
```

### 自定义中间件

需要实现如下函数
```golang
// Endpoint represent one method for calling from remote.
type Endpoint func(ctx context.Context, req *http.Request) (resp *http.Response, err error)

// Middleware deal with input Endpoint and output Endpoint.
type Middleware func(Endpoint) Endpoint
```

可参考默认中间件的实现

```golang
func AuthMiddleware(next api.Endpoint) api.Endpoint {
    return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
        token := ctx.Value(config.ContextAccessToken)
        if token != nil {
            if tokenStr, ok := token.(string); ok {
                req.Header.Set("Access-Token", tokenStr)
            }
        }
        return next(ctx, req)
    }
}
```

### 获取Access Token
> 注：本示例适用于授权时通过Authorization Code获取Access Token和Refresh Token，如需更新Access Token请参考 ./examples/oauth2_access_token_example.go 示例
#### 调用接口
在请求参数内给app_id、secret、auth_code进行赋值，并调用接口即可获得access_token

### 接口调用
接口调用示例可参考 examples 目录下文件
#### 示例
```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ApiOpenApiV30ProjectListGetRequestExample 接口参数
type ApiOpenApiV30ProjectListGetRequestExample struct {
	AdvertiserId int64                   `json:"advertiser_id"`
	Fields       []string                `json:"fields,omitempty"`
	Filtering    ProjectListV30Filtering `json:"filtering,omitempty"`
	Page         int64                   `json:"page,omitempty"`
	PageSize     int64                   `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/project/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()
	// 初始化client，此client建议复用
	configuration := config.NewConfiguration()
	// 可修改默认 HTTP client 为自定义 client
	//configuration.HTTPClient = xxxClient
	apiClient := ad_open_sdk_go.Init(configuration)
	// client修改配置
	apiClient.SetLogEnable(true)
	apiClient.AddDefaultHeader("xxx", "xxx")

	var request ApiOpenApiV30ProjectListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	// 调用接口，首先使用 xxxApi() 方法获取 service
	resp, httpRes, err := apiClient.ProjectListV30Api().
		// 通过 Get/Post 方法获取 Request
		Get(ctx).
		// 传入接口参数，Get 请求通过链式调用传入参数，Post 请求通过结构体传入参数
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Fields(request.Fields).Filtering(request.Filtering).Page(request.Page).PageSize(request.PageSize).
		// 调用 Execute 方法执行请求
		Execute()
	// 接口会返回 resp 结构体、原生 http response、error
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)
}
```
#### 通用接口调用
* 如果想使用 sdk 调用未生成代码的接口，可使用 CommonApi 接口进行调用
* 通用接口支持 Get Post PostMultipart 三种方式调用
    * Get 方法使用 RequestQuery 方法传入参数
    * Post 方式使用 RequestBody 方法传入参数
    * PostMultipart 方法使用 RequestForm 方法传入参数, RequestFile 方法传入文件
* 更多示例可参考 examples/common_api_example.go

### API接口列表

接口域名 https://api.oceanengine.com/

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CdpBrandGetV30Api* | **OpenApiV30CdpBrandGetGet** | **Get** /open_api/v3.0/cdp/brand/get/ | 
*CreativeStrategyListV2Api* | **OpenApi2CreativeStrategyListGet** | **Get** /open_api/2/creative/strategy/list/ | 
*MaterialStatusUpdateV30Api* | **OpenApiV30MaterialStatusUpdatePost** | **Post** /open_api/v3.0/material/status/update/ | 
*Oauth2AccessTokenApi* | **OpenApiOauth2AccessTokenPost** | **Post** /open_api/oauth2/access_token/ | 
*Oauth2RefreshTokenApi* | **OpenApiOauth2RefreshTokenPost** | **Post** /open_api/oauth2/refresh_token/ | 
*ProjectBudgetUpdateV30Api* | **OpenApiV30ProjectBudgetUpdatePost** | **Post** /open_api/v3.0/project/budget/update/ | 
*ProjectCreateV30Api* | **OpenApiV30ProjectCreatePost** | **Post** /open_api/v3.0/project/create/ | 
*ProjectDeleteV30Api* | **OpenApiV30ProjectDeletePost** | **Post** /open_api/v3.0/project/delete/ | 
*ProjectListV30Api* | **OpenApiV30ProjectListGet** | **Get** /open_api/v3.0/project/list/ | 
*ProjectStatusUpdateV30Api* | **OpenApiV30ProjectStatusUpdatePost** | **Post** /open_api/v3.0/project/status/update/ | 
*ProjectUpdateV30Api* | **OpenApiV30ProjectUpdatePost** | **Post** /open_api/v3.0/project/update/ | 
*PromotionAutoGenerateConfigCreateV30Api* | **OpenApiV30PromotionAutoGenerateConfigCreatePost** | **Post** /open_api/v3.0/promotion/auto_generate_config/create/ | 
*PromotionAutoGenerateConfigGetV30Api* | **OpenApiV30PromotionAutoGenerateConfigGetGet** | **Get** /open_api/v3.0/promotion/auto_generate_config/get/ | 
*PromotionBidUpdateV30Api* | **OpenApiV30PromotionBidUpdatePost** | **Post** /open_api/v3.0/promotion/bid/update/ | 
*PromotionCostProtectStatusGetV30Api* | **OpenApiV30PromotionCostProtectStatusGetGet** | **Get** /open_api/v3.0/promotion/cost_protect_status/get/ | 
*PromotionCreateV30Api* | **OpenApiV30PromotionCreatePost** | **Post** /open_api/v3.0/promotion/create/ | 
*PromotionDeepbidUpdateV30Api* | **OpenApiV30PromotionDeepbidUpdatePost** | **Post** /open_api/v3.0/promotion/deepbid/update/ | 
*PromotionDeleteV30Api* | **OpenApiV30PromotionDeletePost** | **Post** /open_api/v3.0/promotion/delete/ | 
*PromotionListV30Api* | **OpenApiV30PromotionListGet** | **Get** /open_api/v3.0/promotion/list/ | 
*PromotionRejectReasonGetV30Api* | **OpenApiV30PromotionRejectReasonGetGet** | **Get** /open_api/v3.0/promotion/reject_reason/get/ | 
*PromotionStatusUpdateV30Api* | **OpenApiV30PromotionStatusUpdatePost** | **Post** /open_api/v3.0/promotion/status/update/ | 
*PromotionUpdateV30Api* | **OpenApiV30PromotionUpdatePost** | **Post** /open_api/v3.0/promotion/update/ | 
*ReportCustomConfigGetV30Api* | **OpenApiV30ReportCustomConfigGetGet** | **Get** /open_api/v3.0/report/custom/config/get/ | 
*ReportCustomGetV30Api* | **OpenApiV30ReportCustomGetGet** | **Get** /open_api/v3.0/report/custom/get/ | 
*SuggWordsV30Api* | **OpenApiV30SuggWordsPost** | **Post** /open_api/v3.0/sugg_words/ | 
*ToolsKeywordsBidRatioCreateV30Api* | **OpenApiV30ToolsKeywordsBidRatioCreatePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/create/ | 
*ToolsKeywordsBidRatioDeleteV30Api* | **OpenApiV30ToolsKeywordsBidRatioDeletePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/delete/ | 
*ToolsKeywordsBidRatioGetV30Api* | **OpenApiV30ToolsKeywordsBidRatioGetGet** | **Get** /open_api/v3.0/tools/keywords_bid_ratio/get/ | 
*ToolsKeywordsBidRatioUpdateV30Api* | **OpenApiV30ToolsKeywordsBidRatioUpdatePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/update/ | 
*ToolsKeywordsProjectInfoGetV30Api* | **OpenApiV30ToolsKeywordsProjectInfoGetGet** | **Get** /open_api/v3.0/tools/keywords_project_info/get/ | 
*ToolsPrivativeWordBatchGetV30Api* | **OpenApiV30ToolsPrivativeWordBatchGetPost** | **Post** /open_api/v3.0/tools/privative_word/batch_get/ | 
*ToolsPrivativeWordProjectAddV30Api* | **OpenApiV30ToolsPrivativeWordProjectAddPost** | **Post** /open_api/v3.0/tools/privative_word/project/add/ | 
*ToolsPrivativeWordProjectUpdateV30Api* | **OpenApiV30ToolsPrivativeWordProjectUpdatePost** | **Post** /open_api/v3.0/tools/privative_word/project/update/ | 
*ToolsPrivativeWordPromotionAddV30Api* | **OpenApiV30ToolsPrivativeWordPromotionAddPost** | **Post** /open_api/v3.0/tools/privative_word/promotion/add/ | 
*ToolsPrivativeWordPromotionUpdateV30Api* | **OpenApiV30ToolsPrivativeWordPromotionUpdatePost** | **Post** /open_api/v3.0/tools/privative_word/promotion/update/ | 


## 问题建议与反馈
如果您在使用SDK过程中有任何问题与建议，请随时登录[开发者官网](https://open.oceanengine.com/labels) ，点击右下角的"咨询"按钮，与我们的客服支持人员联系

## 后续计划
1. 丰富各类场景示例
2. 推出其他语言的SDK
