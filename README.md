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

接口列表 https://open.oceanengine.com/labels/7

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdCostProtectStatusGetV2Api* | **OpenApi2AdCostProtectStatusGetGet** | **Get** /open_api/2/ad/cost_protect_status/get/ | 
*AdGetV2Api* | **OpenApi2AdGetGet** | **Get** /open_api/2/ad/get/ | 
*AdRejectReasonV2Api* | **OpenApi2AdRejectReasonGet** | **Get** /open_api/2/ad/reject_reason/ | 
*AdUpdateStatusV2Api* | **OpenApi2AdUpdateStatusPost** | **Post** /open_api/2/ad/update/status/ | 
*AdlabGroupDetailV30Api* | **OpenApiV30AdlabGroupDetailGet** | **Get** /open_api/v3.0/adlab/group/detail/ | 
*AdlabGroupListV30Api* | **OpenApiV30AdlabGroupListGet** | **Get** /open_api/v3.0/adlab/group/list/ | 
*AdlabGroupUpdateStatusV30Api* | **OpenApiV30AdlabGroupUpdateStatusPost** | **Post** /open_api/v3.0/adlab/group/update_status/ | 
*AdvConvertOleConvertV2Api* | **OpenApi2AdvConvertOleConvertPost** | **Post** /open_api/2/adv_convert/ole/convert/ | 
*AdvertiserAvatarGetV2Api* | **OpenApi2AdvertiserAvatarGetGet** | **Get** /open_api/2/advertiser/avatar/get/ | 
*AdvertiserAvatarSubmitV2Api* | **OpenApi2AdvertiserAvatarSubmitPost** | **Post** /open_api/2/advertiser/avatar/submit/ | 
*AdvertiserBudgetGetV2Api* | **OpenApi2AdvertiserBudgetGetGet** | **Get** /open_api/2/advertiser/budget/get/ | 
*AdvertiserDeliveryQualificationListV30Api* | **OpenApiV30AdvertiserDeliveryQualificationListGet** | **Get** /open_api/v3.0/advertiser/delivery_qualification/list/ | 
*AdvertiserDeliveryQualificationSubmitV30Api* | **OpenApiV30AdvertiserDeliveryQualificationSubmitPost** | **Post** /open_api/v3.0/advertiser/delivery_qualification/submit/ | 
*AdvertiserFundDailyStatV2Api* | **OpenApi2AdvertiserFundDailyStatGet** | **Get** /open_api/2/advertiser/fund/daily_stat/ | 
*AdvertiserFundGetV2Api* | **OpenApi2AdvertiserFundGetGet** | **Get** /open_api/2/advertiser/fund/get/ | 
*AdvertiserFundTransactionGetV2Api* | **OpenApi2AdvertiserFundTransactionGetGet** | **Get** /open_api/2/advertiser/fund/transaction/get/ | 
*AdvertiserInfoV2Api* | **OpenApi2AdvertiserInfoGet** | **Get** /open_api/2/advertiser/info/ | 
*AdvertiserPublicInfoV2Api* | **OpenApi2AdvertiserPublicInfoGet** | **Get** /open_api/2/advertiser/public_info/ | 
*AdvertiserQualificationCreateV2V2Api* | **OpenApi2AdvertiserQualificationCreateV2Post** | **Post** /open_api/2/advertiser/qualification/create_v2/ | 
*AdvertiserQualificationGetV30Api* | **OpenApiV30AdvertiserQualificationGetGet** | **Get** /open_api/v3.0/advertiser/qualification/get/ | 
*AdvertiserQualificationSelectV2V2Api* | **OpenApi2AdvertiserQualificationSelectV2Get** | **Get** /open_api/2/advertiser/qualification/select_v2/ | 
*AdvertiserQualificationSubmitV30Api* | **OpenApiV30AdvertiserQualificationSubmitPost** | **Post** /open_api/v3.0/advertiser/qualification/submit/ | 
*AdvertiserTransferableFundGetV2Api* | **OpenApi2AdvertiserTransferableFundGetGet** | **Get** /open_api/2/advertiser/transferable_fund/get/ | 
*AdvertiserUpdateBudgetV2Api* | **OpenApi2AdvertiserUpdateBudgetPost** | **Post** /open_api/2/advertiser/update/budget/ | 
*AgentAdvertiserCopyV2Api* | **OpenApi2AgentAdvertiserCopyPost** | **Post** /open_api/2/agent/advertiser/copy/ | 
*AgentAdvertiserSelectV2Api* | **OpenApi2AgentAdvertiserSelectGet** | **Get** /open_api/2/agent/advertiser/select/ | 
*AgentAdvertiserUpdateV2Api* | **OpenApi2AgentAdvertiserUpdatePost** | **Post** /open_api/2/agent/advertiser/update/ | 
*AgentChildAgentSelectV2Api* | **OpenApi2AgentChildAgentSelectGet** | **Get** /open_api/2/agent/child_agent/select/ | 
*AgentFundTransferSeqCommitV2Api* | **OpenApi2AgentFundTransferSeqCommitPost** | **Post** /open_api/2/agent/fund/transfer_seq/commit/ | 
*AgentFundTransferSeqCreateV2Api* | **OpenApi2AgentFundTransferSeqCreatePost** | **Post** /open_api/2/agent/fund/transfer_seq/create/ | 
*AgentInfoV2Api* | **OpenApi2AgentInfoGet** | **Get** /open_api/2/agent/info/ | 
*AgentRefundTransferSeqCommitV2Api* | **OpenApi2AgentRefundTransferSeqCommitPost** | **Post** /open_api/2/agent/refund/transfer_seq/commit/ | 
*AgentRefundTransferSeqCreateV2Api* | **OpenApi2AgentRefundTransferSeqCreatePost** | **Post** /open_api/2/agent/refund/transfer_seq/create/ | 
*AnalyticsAttributionV30Api* | **OpenApiV30AnalyticsAttributionPost** | **Post** /open_api/v3.0/analytics/attribution/ | 
*AssetsCreativeComponentCreateV2Api* | **OpenApi2AssetsCreativeComponentCreatePost** | **Post** /open_api/2/assets/creative_component/create/ | 
*AssetsCreativeComponentGetV2Api* | **OpenApi2AssetsCreativeComponentGetGet** | **Get** /open_api/2/assets/creative_component/get/ | 
*AssetsCreativeComponentUpdateV2Api* | **OpenApi2AssetsCreativeComponentUpdatePost** | **Post** /open_api/2/assets/creative_component/update/ | 
*AsyncTaskCreateV2Api* | **OpenApi2AsyncTaskCreatePost** | **Post** /open_api/2/async_task/create/ | 
*AsyncTaskDownloadV2Api* | **OpenApi2AsyncTaskDownloadGet** | **Get** /open_api/2/async_task/download/ | 
*AsyncTaskGetV2Api* | **OpenApi2AsyncTaskGetGet** | **Get** /open_api/2/async_task/get/ | 
*AudiencePackageCreateV2Api* | **OpenApi2AudiencePackageCreatePost** | **Post** /open_api/2/audience_package/create/ | 
*AudiencePackageDeleteV2Api* | **OpenApi2AudiencePackageDeletePost** | **Post** /open_api/2/audience_package/delete/ | 
*AudiencePackageUpdateV2Api* | **OpenApi2AudiencePackageUpdatePost** | **Post** /open_api/2/audience_package/update/ | 
*BusinessPlatformCompanyAccountGetV30Api* | **OpenApiV30BusinessPlatformCompanyAccountGetGet** | **Get** /open_api/v3.0/business_platform/company_account/get/ | 
*BusinessPlatformCompanyInfoGetV30Api* | **OpenApiV30BusinessPlatformCompanyInfoGetGet** | **Get** /open_api/v3.0/business_platform/company_info/get/ | 
*BusinessPlatformPartnerOrganizationListV2Api* | **OpenApi2BusinessPlatformPartnerOrganizationListGet** | **Get** /open_api/2/business_platform/partner_organization/list/ | 
*CampaignGetV2Api* | **OpenApi2CampaignGetGet** | **Get** /open_api/2/campaign/get/ | 
*CampaignUpdateStatusV2Api* | **OpenApi2CampaignUpdateStatusPost** | **Post** /open_api/2/campaign/update/status/ | 
*CarouselAdGetV2Api* | **OpenApi2CarouselAdGetGet** | **Get** /open_api/2/carousel/ad/get/ | 
*CarouselCreateV2Api* | **OpenApi2CarouselCreatePost** | **Post** /open_api/2/carousel/create/ | 
*CarouselDeleteV2Api* | **OpenApi2CarouselDeletePost** | **Post** /open_api/2/carousel/delete/ | 
*CarouselListV2Api* | **OpenApi2CarouselListGet** | **Get** /open_api/2/carousel/list/ | 
*CarouselUpdateV2Api* | **OpenApi2CarouselUpdatePost** | **Post** /open_api/2/carousel/update/ | 
*CdpBrandGetV30Api* | **OpenApiV30CdpBrandGetGet** | **Get** /open_api/v3.0/cdp/brand/get/ | 
*ClueCouponCodeConsumeV2Api* | **OpenApi2ClueCouponCodeConsumePost** | **Post** /open_api/2/clue/coupon/code/consume/ | 
*ClueCouponCodeGetV2Api* | **OpenApi2ClueCouponCodeGetGet** | **Get** /open_api/2/clue/coupon/code/get/ | 
*ClueCouponCreateV2Api* | **OpenApi2ClueCouponCreatePost** | **Post** /open_api/2/clue/coupon/create/ | 
*ClueCouponDetailV2Api* | **OpenApi2ClueCouponDetailGet** | **Get** /open_api/2/clue/coupon/detail/ | 
*ClueCouponEmployeeCreateV2Api* | **OpenApi2ClueCouponEmployeeCreatePost** | **Post** /open_api/2/clue/coupon/employee/create/ | 
*ClueCouponEmployeeDeleteV2Api* | **OpenApi2ClueCouponEmployeeDeletePost** | **Post** /open_api/2/clue/coupon/employee/delete/ | 
*ClueCouponEmployeeGetV2Api* | **OpenApi2ClueCouponEmployeeGetGet** | **Get** /open_api/2/clue/coupon/employee/get/ | 
*ClueCouponGetV2Api* | **OpenApi2ClueCouponGetGet** | **Get** /open_api/2/clue/coupon/get/ | 
*ClueCouponUpdateV2Api* | **OpenApi2ClueCouponUpdatePost** | **Post** /open_api/2/clue/coupon/update/ | 
*ClueFormCreateV2Api* | **OpenApi2ClueFormCreatePost** | **Post** /open_api/2/clue/form/create/ | 
*ClueFormDeleteV2Api* | **OpenApi2ClueFormDeletePost** | **Post** /open_api/2/clue/form/delete/ | 
*ClueFormDetailV2Api* | **OpenApi2ClueFormDetailGet** | **Get** /open_api/2/clue/form/detail/ | 
*ClueFormListV2Api* | **OpenApi2ClueFormListGet** | **Get** /open_api/2/clue/form/list/ | 
*ClueFormUpdateV2Api* | **OpenApi2ClueFormUpdatePost** | **Post** /open_api/2/clue/form/update/ | 
*ClueSmartphoneCreateV2Api* | **OpenApi2ClueSmartphoneCreatePost** | **Post** /open_api/2/clue/smartphone/create/ | 
*ClueSmartphoneDeleteV2Api* | **OpenApi2ClueSmartphoneDeletePost** | **Post** /open_api/2/clue/smartphone/delete/ | 
*ClueSmartphoneGetV2Api* | **OpenApi2ClueSmartphoneGetGet** | **Get** /open_api/2/clue/smartphone/get/ | 
*ClueSmartphoneRecordV2Api* | **OpenApi2ClueSmartphoneRecordGet** | **Get** /open_api/2/clue/smartphone/record/ | 
*ClueWechatInstanceDetailV2Api* | **OpenApi2ClueWechatInstanceDetailGet** | **Get** /open_api/2/clue/wechat_instance/detail/ | 
*ClueWechatInstanceListV2Api* | **OpenApi2ClueWechatInstanceListGet** | **Get** /open_api/2/clue/wechat_instance/list/ | 
*ClueWechatInstanceUpdateV2Api* | **OpenApi2ClueWechatInstanceUpdatePost** | **Post** /open_api/2/clue/wechat_instance/update/ | 
*ClueWechatPoolListV2Api* | **OpenApi2ClueWechatPoolListGet** | **Get** /open_api/2/clue/wechat_pool/list/ | 
*CreativeAdMetricsV2Api* | **OpenApi2CreativeAdMetricsGet** | **Get** /open_api/2/creative/ad/metrics/ | 
*CreativeAutoGenerateConfigCreateV2Api* | **OpenApi2CreativeAutoGenerateConfigCreatePost** | **Post** /open_api/2/creative/auto_generate_config/create/ | 
*CreativeAutoGenerateConfigGetV2Api* | **OpenApi2CreativeAutoGenerateConfigGetGet** | **Get** /open_api/2/creative/auto_generate_config/get/ | 
*CreativeAutoGenerateConfigV2CreateV2Api* | **OpenApi2CreativeAutoGenerateConfigV2CreatePost** | **Post** /open_api/2/creative/auto_generate_config/v2/create/ | 
*CreativeDetailGetV30Api* | **OpenApiV30CreativeDetailGetGet** | **Get** /open_api/v3.0/creative/detail/get/ | 
*CreativeGetV2Api* | **OpenApi2CreativeGetGet** | **Get** /open_api/2/creative/get/ | 
*CreativeMaterialMetricsV2Api* | **OpenApi2CreativeMaterialMetricsGet** | **Get** /open_api/2/creative/material/metrics/ | 
*CreativeRejectReasonV2Api* | **OpenApi2CreativeRejectReasonGet** | **Get** /open_api/2/creative/reject_reason/ | 
*CreativeStrategyListV2Api* | **OpenApi2CreativeStrategyListGet** | **Get** /open_api/2/creative/strategy/list/ | 
*CreativeTemplateDetailGetV2Api* | **OpenApi2CreativeTemplateDetailGetGet** | **Get** /open_api/2/creative/template/detail/get/ | 
*CreativeTemplateListGetV2Api* | **OpenApi2CreativeTemplateListGetGet** | **Get** /open_api/2/creative/template/list/get/ | 
*CreativeTemplateTagsGetV2Api* | **OpenApi2CreativeTemplateTagsGetGet** | **Get** /open_api/2/creative/template/tags/get/ | 
*CustomerCenterAdvertiserListV2Api* | **OpenApi2CustomerCenterAdvertiserListGet** | **Get** /open_api/2/customer_center/advertiser/list/ | 
*CustomerCenterAdvertiserTransferableListV2Api* | **OpenApi2CustomerCenterAdvertiserTransferableListGet** | **Get** /open_api/2/customer_center/advertiser/transferable/list/ | 
*CustomerCenterFundTransferSeqCommitV2Api* | **OpenApi2CustomerCenterFundTransferSeqCommitPost** | **Post** /open_api/2/customer_center/fund/transfer_seq/commit/ | 
*CustomerCenterFundTransferSeqCreateV2Api* | **OpenApi2CustomerCenterFundTransferSeqCreatePost** | **Post** /open_api/2/customer_center/fund/transfer_seq/create/ | 
*DecorationCouponGetV30Api* | **OpenApiV30DecorationCouponGetGet** | **Get** /open_api/v3.0/decoration/coupon/get/ | 
*DmpBrandGetV2Api* | **OpenApi2DmpBrandGetGet** | **Get** /open_api/2/dmp/brand/get/ | 
*DmpCustomAudienceCopyV2Api* | **OpenApi2DmpCustomAudienceCopyPost** | **Post** /open_api/2/dmp/custom_audience/copy/ | 
*DmpCustomAudienceDeleteV2Api* | **OpenApi2DmpCustomAudienceDeletePost** | **Post** /open_api/2/dmp/custom_audience/delete/ | 
*DmpCustomAudiencePublishV2Api* | **OpenApi2DmpCustomAudiencePublishPost** | **Post** /open_api/2/dmp/custom_audience/publish/ | 
*DmpCustomAudiencePushV2V2Api* | **OpenApi2DmpCustomAudiencePushV2Post** | **Post** /open_api/2/dmp/custom_audience/push_v2/ | 
*DmpCustomAudienceReadV2Api* | **OpenApi2DmpCustomAudienceReadGet** | **Get** /open_api/2/dmp/custom_audience/read/ | 
*DmpCustomAudienceSelectV2Api* | **OpenApi2DmpCustomAudienceSelectGet** | **Get** /open_api/2/dmp/custom_audience/select/ | 
*DmpDataSourceCreateV2Api* | **OpenApi2DmpDataSourceCreatePost** | **Post** /open_api/2/dmp/data_source/create/ | 
*DmpDataSourceFileUploadV2Api* | **OpenApi2DmpDataSourceFileUploadPost** | **Post** /open_api/2/dmp/data_source/file/upload/ | 
*DmpDataSourceReadV2Api* | **OpenApi2DmpDataSourceReadGet** | **Get** /open_api/2/dmp/data_source/read/ | 
*DmpDataSourceUpdateV2Api* | **OpenApi2DmpDataSourceUpdatePost** | **Post** /open_api/2/dmp/data_source/update/ | 
*DouplusOrderListV30Api* | **OpenApiV30DouplusOrderListGet** | **Get** /open_api/v3.0/douplus/order/list/ | 
*DouplusOrderReportV30Api* | **OpenApiV30DouplusOrderReportGet** | **Get** /open_api/v3.0/douplus/order/report/ | 
*DpaAssetsDetailReadV2Api* | **OpenApi2DpaAssetsDetailReadGet** | **Get** /open_api/2/dpa/assets/detail/read/ | 
*DpaAssetsListV2Api* | **OpenApi2DpaAssetsListGet** | **Get** /open_api/2/dpa/assets/list/ | 
*DpaCategoryGetV2Api* | **OpenApi2DpaCategoryGetGet** | **Get** /open_api/2/dpa/category/get/ | 
*DpaDetailGetV2Api* | **OpenApi2DpaDetailGetGet** | **Get** /open_api/2/dpa/detail/get/ | 
*DpaDictGetV2Api* | **OpenApi2DpaDictGetGet** | **Get** /open_api/2/dpa/dict/get/ | 
*DpaMetaGetV2Api* | **OpenApi2DpaMetaGetGet** | **Get** /open_api/2/dpa/meta/get/ | 
*DpaProductAvailablesV2Api* | **OpenApi2DpaProductAvailablesGet** | **Get** /open_api/2/dpa/product/availables/ | 
*DpaProductCreateV2Api* | **OpenApi2DpaProductCreatePost** | **Post** /open_api/2/dpa/product/create/ | 
*DpaProductDeleteV2Api* | **OpenApi2DpaProductDeletePost** | **Post** /open_api/2/dpa/product/delete/ | 
*DpaProductDetailGetV2Api* | **OpenApi2DpaProductDetailGetGet** | **Get** /open_api/2/dpa/product/detail/get/ | 
*DpaProductStatusBatchUpdateV2Api* | **OpenApi2DpaProductStatusBatchUpdatePost** | **Post** /open_api/2/dpa/product_status/batch_update/ | 
*DpaProductUpdateV2Api* | **OpenApi2DpaProductUpdatePost** | **Post** /open_api/2/dpa/product/update/ | 
*DpaTemplateGetV2Api* | **OpenApi2DpaTemplateGetGet** | **Get** /open_api/2/dpa/template/get/ | 
*DpaVideoGetV2Api* | **OpenApi2DpaVideoGetGet** | **Get** /open_api/2/dpa/video/get/ | 
*EnterpriseBindListGetV10Api* | **OpenApiV10EnterpriseBindListGetGet** | **Get** /open_api/v1.0/enterprise/bind/list/get/ | 
*EnterpriseCommentDetailV10Api* | **OpenApiV10EnterpriseCommentDetailGet** | **Get** /open_api/v1.0/enterprise/comment/detail/ | 
*EnterpriseCommentListGetV10Api* | **OpenApiV10EnterpriseCommentListGetGet** | **Get** /open_api/v1.0/enterprise/comment/list/get/ | 
*EnterpriseCommentReplyListV10Api* | **OpenApiV10EnterpriseCommentReplyListGet** | **Get** /open_api/v1.0/enterprise/comment/reply/list/ | 
*EnterpriseCommentReplyV10Api* | **OpenApiV10EnterpriseCommentReplyPost** | **Post** /open_api/v1.0/enterprise/comment/reply/ | 
*EnterpriseFlowCategoryGetV10Api* | **OpenApiV10EnterpriseFlowCategoryGetGet** | **Get** /open_api/v1.0/enterprise/flow/category/get/ | 
*EnterpriseInfoV10Api* | **OpenApiV10EnterpriseInfoGet** | **Get** /open_api/v1.0/enterprise/info/ | 
*EnterpriseItemListV10Api* | **OpenApiV10EnterpriseItemListGet** | **Get** /open_api/v1.0/enterprise/item/list/ | 
*EnterpriseOperationLogGetV10Api* | **OpenApiV10EnterpriseOperationLogGetGet** | **Get** /open_api/v1.0/enterprise/operation/log/get/ | 
*EnterpriseOverviewDataGetV10Api* | **OpenApiV10EnterpriseOverviewDataGetGet** | **Get** /open_api/v1.0/enterprise/overview/data/get/ | 
*EnterpriseVideoInfoGetV10Api* | **OpenApiV10EnterpriseVideoInfoGetGet** | **Get** /open_api/v1.0/enterprise/video/info/get/ | 
*EventManagerAssetsCreateV2Api* | **OpenApi2EventManagerAssetsCreatePost** | **Post** /open_api/2/event_manager/assets/create/ | 
*EventManagerAvailableEventsGetV2Api* | **OpenApi2EventManagerAvailableEventsGetGet** | **Get** /open_api/2/event_manager/available_events/get/ | 
*EventManagerDeepBidTypeGetV30Api* | **OpenApiV30EventManagerDeepBidTypeGetGet** | **Get** /open_api/v3.0/event_manager/deep_bid_type/get/ | 
*EventManagerEventConfigsGetV2Api* | **OpenApi2EventManagerEventConfigsGetGet** | **Get** /open_api/2/event_manager/event_configs/get/ | 
*EventManagerEventsCreateV2Api* | **OpenApi2EventManagerEventsCreatePost** | **Post** /open_api/2/event_manager/events/create/ | 
*EventManagerOptimizedGoalGetV2V30Api* | **OpenApiV30EventManagerOptimizedGoalGetV2Get** | **Get** /open_api/v3.0/event_manager/optimized_goal/get_v2/ | 
*EventManagerShareCancelV30Api* | **OpenApiV30EventManagerShareCancelPost** | **Post** /open_api/v3.0/event_manager/share/cancel/ | 
*EventManagerShareGetV30Api* | **OpenApiV30EventManagerShareGetGet** | **Get** /open_api/v3.0/event_manager/share/get/ | 
*EventManagerShareV30Api* | **OpenApiV30EventManagerSharePost** | **Post** /open_api/v3.0/event_manager/share/ | 
*EventManagerTrackUrlCreateV2Api* | **OpenApi2EventManagerTrackUrlCreatePost** | **Post** /open_api/2/event_manager/track_url/create/ | 
*EventManagerTrackUrlGetV2Api* | **OpenApi2EventManagerTrackUrlGetGet** | **Get** /open_api/2/event_manager/track_url/get/ | 
*EventManagerTrackUrlUpdateV2Api* | **OpenApi2EventManagerTrackUrlUpdatePost** | **Post** /open_api/2/event_manager/track_url/update/ | 
*FileImageAdGetV2Api* | **OpenApi2FileImageAdGetGet** | **Get** /open_api/2/file/image/ad/get/ | 
*FileImageDeleteV30Api* | **OpenApiV30FileImageDeletePost** | **Post** /open_api/v3.0/file/image/delete/ | 
*FileImageGetV2Api* | **OpenApi2FileImageGetGet** | **Get** /open_api/2/file/image/get/ | 
*FileMaterialBindV2Api* | **OpenApi2FileMaterialBindPost** | **Post** /open_api/2/file/material/bind/ | 
*FileMaterialDetailV2Api* | **OpenApi2FileMaterialDetailGet** | **Get** /open_api/2/file/material/detail/ | 
*FileMaterialListV2Api* | **OpenApi2FileMaterialListGet** | **Get** /open_api/2/file/material/list/ | 
*FileQualityGetV30Api* | **OpenApiV30FileQualityGetGet** | **Get** /open_api/v3.0/file/quality/get/ | 
*FileQualitySubmitV30Api* | **OpenApiV30FileQualitySubmitPost** | **Post** /open_api/v3.0/file/quality/submit/ | 
*FileVideoAdGetV2Api* | **OpenApi2FileVideoAdGetGet** | **Get** /open_api/2/file/video/ad/get/ | 
*FileVideoAwemeGetV2Api* | **OpenApi2FileVideoAwemeGetGet** | **Get** /open_api/2/file/video/aweme/get/ | 
*FileVideoDeleteV2Api* | **OpenApi2FileVideoDeletePost** | **Post** /open_api/2/file/video/delete/ | 
*FileVideoEfficiencyGetV2Api* | **OpenApi2FileVideoEfficiencyGetGet** | **Get** /open_api/2/file/video/efficiency/get/ | 
*FileVideoGetV2Api* | **OpenApi2FileVideoGetGet** | **Get** /open_api/2/file/video/get/ | 
*FileVideoMaterialClearTaskCreateV2Api* | **OpenApi2FileVideoMaterialClearTaskCreatePost** | **Post** /open_api/2/file/video/material/clear_task/create/ | 
*FileVideoMaterialClearTaskGetV2Api* | **OpenApi2FileVideoMaterialClearTaskGetGet** | **Get** /open_api/2/file/video/material/clear_task/get/ | 
*FileVideoMaterialClearTaskResultGetV2Api* | **OpenApi2FileVideoMaterialClearTaskResultGetGet** | **Get** /open_api/2/file/video/material/clear_task_result/get/ | 
*FileVideoUpdateV2Api* | **OpenApi2FileVideoUpdatePost** | **Post** /open_api/2/file/video/update/ | 
*FundSharedWalletBalanceGetV2Api* | **OpenApi2FundSharedWalletBalanceGetGet** | **Get** /open_api/2/fund/shared_wallet_balance/get/ | 
*KeywordCreateV30Api* | **OpenApiV30KeywordCreatePost** | **Post** /open_api/v3.0/keyword/create/ | 
*KeywordDeleteV30Api* | **OpenApiV30KeywordDeletePost** | **Post** /open_api/v3.0/keyword/delete/ | 
*KeywordFeedadsSuggestV2Api* | **OpenApi2KeywordFeedadsSuggestGet** | **Get** /open_api/2/keyword_feedads/suggest/ | 
*KeywordGetV2Api* | **OpenApi2KeywordGetGet** | **Get** /open_api/2/keyword/get/ | 
*KeywordListV30Api* | **OpenApiV30KeywordListGet** | **Get** /open_api/v3.0/keyword/list/ | 
*KeywordUpdateV30Api* | **OpenApiV30KeywordUpdatePost** | **Post** /open_api/v3.0/keyword/update/ | 
*MajordomoAdvertiserSelectV2Api* | **OpenApi2MajordomoAdvertiserSelectGet** | **Get** /open_api/2/majordomo/advertiser/select/ | 
*MaterialStatusUpdateV30Api* | **OpenApiV30MaterialStatusUpdatePost** | **Post** /open_api/v3.0/material/status/update/ | 
*NativeAnchorCreateV30Api* | **OpenApiV30NativeAnchorCreatePost** | **Post** /open_api/v3.0/native_anchor/create/ | 
*NativeAnchorGetV30Api* | **OpenApiV30NativeAnchorGetGet** | **Get** /open_api/v3.0/native_anchor/get/ | 
*Oauth2AccessTokenApi* | **OpenApiOauth2AccessTokenPost** | **Post** /open_api/oauth2/access_token/ | 
*Oauth2AdvertiserGetApi* | **OpenApiOauth2AdvertiserGetGet** | **Get** /open_api/oauth2/advertiser/get/ | 
*Oauth2AppAccessTokenApi* | **OpenApiOauth2AppAccessTokenPost** | **Post** /open_api/oauth2/app_access_token/ | 
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
*PromotionBudgetUpdateV30Api* | **OpenApiV30PromotionBudgetUpdatePost** | **Post** /open_api/v3.0/promotion/budget/update/ | 
*PromotionCostProtectStatusGetV30Api* | **OpenApiV30PromotionCostProtectStatusGetGet** | **Get** /open_api/v3.0/promotion/cost_protect_status/get/ | 
*PromotionCreateV30Api* | **OpenApiV30PromotionCreatePost** | **Post** /open_api/v3.0/promotion/create/ | 
*PromotionDeepbidUpdateV30Api* | **OpenApiV30PromotionDeepbidUpdatePost** | **Post** /open_api/v3.0/promotion/deepbid/update/ | 
*PromotionDeleteV30Api* | **OpenApiV30PromotionDeletePost** | **Post** /open_api/v3.0/promotion/delete/ | 
*PromotionListV30Api* | **OpenApiV30PromotionListGet** | **Get** /open_api/v3.0/promotion/list/ | 
*PromotionRejectReasonGetV30Api* | **OpenApiV30PromotionRejectReasonGetGet** | **Get** /open_api/v3.0/promotion/reject_reason/get/ | 
*PromotionScheduleTimeUpdateV30Api* | **OpenApiV30PromotionScheduleTimeUpdatePost** | **Post** /open_api/v3.0/promotion/schedule_time/update/ | 
*PromotionStatusUpdateV30Api* | **OpenApiV30PromotionStatusUpdatePost** | **Post** /open_api/v3.0/promotion/status/update/ | 
*PromotionUpdateV30Api* | **OpenApiV30PromotionUpdatePost** | **Post** /open_api/v3.0/promotion/update/ | 
*QianchuanAdBidUpdateV10Api* | **OpenApiV10QianchuanAdBidUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/bid/update/ | 
*QianchuanAdBudgetUpdateV10Api* | **OpenApiV10QianchuanAdBudgetUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/budget/update/ | 
*QianchuanAdCreateV10Api* | **OpenApiV10QianchuanAdCreatePost** | **Post** /open_api/v1.0/qianchuan/ad/create/ | 
*QianchuanAdDetailGetV10Api* | **OpenApiV10QianchuanAdDetailGetGet** | **Get** /open_api/v1.0/qianchuan/ad/detail/get/ | 
*QianchuanAdGetV10Api* | **OpenApiV10QianchuanAdGetGet** | **Get** /open_api/v1.0/qianchuan/ad/get/ | 
*QianchuanAdKeywordsGetV10Api* | **OpenApiV10QianchuanAdKeywordsGetGet** | **Get** /open_api/v1.0/qianchuan/ad/keywords/get/ | 
*QianchuanAdKeywordsUpdateV10Api* | **OpenApiV10QianchuanAdKeywordsUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/keywords/update/ | 
*QianchuanAdPivativewordsGetV10Api* | **OpenApiV10QianchuanAdPivativewordsGetGet** | **Get** /open_api/v1.0/qianchuan/ad/pivativewords/get/ | 
*QianchuanAdPivativewordsUpdateV10Api* | **OpenApiV10QianchuanAdPivativewordsUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/pivativewords/update/ | 
*QianchuanAdQuotaGetV10Api* | **OpenApiV10QianchuanAdQuotaGetGet** | **Get** /open_api/v1.0/qianchuan/ad/quota/get/ | 
*QianchuanAdRecommendKeywordsGetV10Api* | **OpenApiV10QianchuanAdRecommendKeywordsGetGet** | **Get** /open_api/v1.0/qianchuan/ad/recommend_keywords/get | 
*QianchuanAdRegionUpdateV10Api* | **OpenApiV10QianchuanAdRegionUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/region/update/ | 
*QianchuanAdRejectReasonV10Api* | **OpenApiV10QianchuanAdRejectReasonGet** | **Get** /open_api/v1.0/qianchuan/ad/reject_reason/ | 
*QianchuanAdScheduleDateUpdateV10Api* | **OpenApiV10QianchuanAdScheduleDateUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/schedule_date/update/ | 
*QianchuanAdScheduleFixedRangeUpdateV10Api* | **OpenApiV10QianchuanAdScheduleFixedRangeUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/schedule_fixed_range/update/ | 
*QianchuanAdScheduleTimeUpdateV10Api* | **OpenApiV10QianchuanAdScheduleTimeUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/schedule_time/update/ | 
*QianchuanAdStatusUpdateV10Api* | **OpenApiV10QianchuanAdStatusUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/status/update/ | 
*QianchuanAdUpdateV10Api* | **OpenApiV10QianchuanAdUpdatePost** | **Post** /open_api/v1.0/qianchuan/ad/update/ | 
*QianchuanAdvertiserTypeGetV10Api* | **OpenApiV10QianchuanAdvertiserTypeGetGet** | **Get** /open_api/v1.0/qianchuan/advertiser/type/get/ | 
*QianchuanAudienceCreateByFileV10Api* | **OpenApiV10QianchuanAudienceCreateByFilePost** | **Post** /open_api/v1.0/qianchuan/audience/create_by_file/ | 
*QianchuanAudienceDeleteV10Api* | **OpenApiV10QianchuanAudienceDeletePost** | **Post** /open_api/v1.0/qianchuan/audience/delete/ | 
*QianchuanAudienceFilePartUploadV10Api* | **OpenApiV10QianchuanAudienceFilePartUploadPost** | **Post** /open_api/v1.0/qianchuan/audience_file/part_upload/ | 
*QianchuanAudienceFileUploadV10Api* | **OpenApiV10QianchuanAudienceFileUploadPost** | **Post** /open_api/v1.0/qianchuan/audience_file/upload/ | 
*QianchuanAudienceGroupGetV10Api* | **OpenApiV10QianchuanAudienceGroupGetGet** | **Get** /open_api/v1.0/qianchuan/audience_group/get/ | 
*QianchuanAudienceListGetV10Api* | **OpenApiV10QianchuanAudienceListGetGet** | **Get** /open_api/v1.0/qianchuan/audience_list/get/ | 
*QianchuanAudiencePushV10Api* | **OpenApiV10QianchuanAudiencePushPost** | **Post** /open_api/v1.0/qianchuan/audience/push/ | 
*QianchuanAwemeAuthorizedGetV10Api* | **OpenApiV10QianchuanAwemeAuthorizedGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/authorized/get/ | 
*QianchuanAwemeEstimateProfitV10Api* | **OpenApiV10QianchuanAwemeEstimateProfitGet** | **Get** /open_api/v1.0/qianchuan/aweme/estimate_profit/ | 
*QianchuanAwemeInterestActionInterestKeywordV10Api* | **OpenApiV10QianchuanAwemeInterestActionInterestKeywordGet** | **Get** /open_api/v1.0/qianchuan/aweme/interest_action/interest/keyword/ | 
*QianchuanAwemeOrderCreateV10Api* | **OpenApiV10QianchuanAwemeOrderCreatePost** | **Post** /open_api/v1.0/qianchuan/aweme/order/create/ | 
*QianchuanAwemeOrderDetailGetV10Api* | **OpenApiV10QianchuanAwemeOrderDetailGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/order/detail/get/ | 
*QianchuanAwemeOrderGetV10Api* | **OpenApiV10QianchuanAwemeOrderGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/order/get/ | 
*QianchuanAwemeOrderQuotaGetV10Api* | **OpenApiV10QianchuanAwemeOrderQuotaGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/order/quota/get/ | 
*QianchuanAwemeOrderTerminateV10Api* | **OpenApiV10QianchuanAwemeOrderTerminatePost** | **Post** /open_api/v1.0/qianchuan/aweme/order/terminate/ | 
*QianchuanAwemeProductAvailableGetV10Api* | **OpenApiV10QianchuanAwemeProductAvailableGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/product/available/get/ | 
*QianchuanAwemeReportOrderGetV10Api* | **OpenApiV10QianchuanAwemeReportOrderGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/report/order/get/ | 
*QianchuanAwemeSuggestBidV10Api* | **OpenApiV10QianchuanAwemeSuggestBidGet** | **Get** /open_api/v1.0/qianchuan/aweme/suggest_bid/ | 
*QianchuanAwemeSuggestRoiGoalV10Api* | **OpenApiV10QianchuanAwemeSuggestRoiGoalGet** | **Get** /open_api/v1.0/qianchuan/aweme/suggest/roi/goal/ | 
*QianchuanAwemeVideoGetV10Api* | **OpenApiV10QianchuanAwemeVideoGetGet** | **Get** /open_api/v1.0/qianchuan/aweme/video/get/ | 
*QianchuanBatchCampaignStatusUpdateV10Api* | **OpenApiV10QianchuanBatchCampaignStatusUpdatePost** | **Post** /open_api/v1.0/qianchuan/batch_campaign_status/update/ | 
*QianchuanCampaignCreateV10Api* | **OpenApiV10QianchuanCampaignCreatePost** | **Post** /open_api/v1.0/qianchuan/campaign/create/ | 
*QianchuanCampaignListGetV10Api* | **OpenApiV10QianchuanCampaignListGetGet** | **Get** /open_api/v1.0/qianchuan/campaign_list/get/ | 
*QianchuanCampaignUpdateV10Api* | **OpenApiV10QianchuanCampaignUpdatePost** | **Post** /open_api/v1.0/qianchuan/campaign/update/ | 
*QianchuanCreativeGetV10Api* | **OpenApiV10QianchuanCreativeGetGet** | **Get** /open_api/v1.0/qianchuan/creative/get/ | 
*QianchuanCreativeRejectReasonV10Api* | **OpenApiV10QianchuanCreativeRejectReasonGet** | **Get** /open_api/v1.0/qianchuan/creative/reject_reason/ | 
*QianchuanCreativeStatusUpdateV10Api* | **OpenApiV10QianchuanCreativeStatusUpdatePost** | **Post** /open_api/v1.0/qianchuan/creative/status/update/ | 
*QianchuanDmpAudiencesGetV10Api* | **OpenApiV10QianchuanDmpAudiencesGetGet** | **Get** /open_api/v1.0/qianchuan/dmp/audiences/get/ | 
*QianchuanEstimateEffectV10Api* | **OpenApiV10QianchuanEstimateEffectGet** | **Get** /open_api/v1.0/qianchuan/estimate/effect/ | 
*QianchuanFileImageDeleteV10Api* | **OpenApiV10QianchuanFileImageDeletePost** | **Post** /open_api/v1.0/qianchuan/file/image/delete/ | 
*QianchuanFileVideoAwemeGetV10Api* | **OpenApiV10QianchuanFileVideoAwemeGetGet** | **Get** /open_api/v1.0/qianchuan/file/video/aweme/get/ | 
*QianchuanFileVideoDeleteV10Api* | **OpenApiV10QianchuanFileVideoDeletePost** | **Post** /open_api/v1.0/qianchuan/file/video/delete/ | 
*QianchuanFileVideoEfficiencyGetV10Api* | **OpenApiV10QianchuanFileVideoEfficiencyGetGet** | **Get** /open_api/v1.0/qianchuan/file/video/efficiency/get/ | 
*QianchuanFileVideoOriginalGetV10Api* | **OpenApiV10QianchuanFileVideoOriginalGetGet** | **Get** /open_api/v1.0/qianchuan/file/video/original/get/ | 
*QianchuanFinanceDetailGetV10Api* | **OpenApiV10QianchuanFinanceDetailGetGet** | **Get** /open_api/v1.0/qianchuan/finance/detail/get/ | 
*QianchuanFinanceWalletGetV10Api* | **OpenApiV10QianchuanFinanceWalletGetGet** | **Get** /open_api/v1.0/qianchuan/finance/wallet/get/ | 
*QianchuanImageGetV10Api* | **OpenApiV10QianchuanImageGetGet** | **Get** /open_api/v1.0/qianchuan/image/get/ | 
*QianchuanKeywordCheckV10Api* | **OpenApiV10QianchuanKeywordCheckPost** | **Post** /open_api/v1.0/qianchuan/keyword/check/ | 
*QianchuanLqAdGetV10Api* | **OpenApiV10QianchuanLqAdGetGet** | **Get** /open_api/v1.0/qianchuan/lq_ad/get/ | 
*QianchuanOrientationPackageGetV10Api* | **OpenApiV10QianchuanOrientationPackageGetGet** | **Get** /open_api/v1.0/qianchuan/orientation_package/get/ | 
*QianchuanProductAnalyseCompareCreativeV10Api* | **OpenApiV10QianchuanProductAnalyseCompareCreativeGet** | **Get** /open_api/v1.0/qianchuan/product/analyse/compare_creative/ | 
*QianchuanProductAnalyseCompareStatsDataV10Api* | **OpenApiV10QianchuanProductAnalyseCompareStatsDataGet** | **Get** /open_api/v1.0/qianchuan/product/analyse/compare_stats_data/ | 
*QianchuanProductAnalyseListV10Api* | **OpenApiV10QianchuanProductAnalyseListGet** | **Get** /open_api/v1.0/qianchuan/product/analyse/list/ | 
*QianchuanProductAvailableGetV10Api* | **OpenApiV10QianchuanProductAvailableGetGet** | **Get** /open_api/v1.0/qianchuan/product/available/get/ | 
*QianchuanReportAdGetV10Api* | **OpenApiV10QianchuanReportAdGetGet** | **Get** /open_api/v1.0/qianchuan/report/ad/get/ | 
*QianchuanReportAdvertiserGetV10Api* | **OpenApiV10QianchuanReportAdvertiserGetGet** | **Get** /open_api/v1.0/qianchuan/report/advertiser/get/ | 
*QianchuanReportCreativeGetV10Api* | **OpenApiV10QianchuanReportCreativeGetGet** | **Get** /open_api/v1.0/qianchuan/report/creative/get/ | 
*QianchuanReportLiveGetV10Api* | **OpenApiV10QianchuanReportLiveGetGet** | **Get** /open_api/v1.0/qianchuan/report/live/get/ | 
*QianchuanReportLongTransferOrderGetV10Api* | **OpenApiV10QianchuanReportLongTransferOrderGetGet** | **Get** /open_api/v1.0/qianchuan/report/long_transfer/order/get/ | 
*QianchuanReportMaterialGetV10Api* | **OpenApiV10QianchuanReportMaterialGetGet** | **Get** /open_api/v1.0/qianchuan/report/material/get/ | 
*QianchuanReportSearchWordGetV10Api* | **OpenApiV10QianchuanReportSearchWordGetGet** | **Get** /open_api/v1.0/qianchuan/report/search_word/get/ | 
*QianchuanReportUniPromotionGetV10Api* | **OpenApiV10QianchuanReportUniPromotionGetGet** | **Get** /open_api/v1.0/qianchuan/report/uni_promotion/get/ | 
*QianchuanReportVideoUserLoseGetV10Api* | **OpenApiV10QianchuanReportVideoUserLoseGetGet** | **Get** /open_api/v1.0/qianchuan/report/video_user_lose/get/ | 
*QianchuanRoiGoalUpdateV10Api* | **OpenApiV10QianchuanRoiGoalUpdatePost** | **Post** /open_api/v1.0/qianchuan/roi/goal/update | 
*QianchuanShopAdvertiserListV10Api* | **OpenApiV10QianchuanShopAdvertiserListGet** | **Get** /open_api/v1.0/qianchuan/shop/advertiser/list/ | 
*QianchuanShopGetV10Api* | **OpenApiV10QianchuanShopGetGet** | **Get** /open_api/v1.0/qianchuan/shop/get/ | 
*QianchuanSuggestBidV10Api* | **OpenApiV10QianchuanSuggestBidGet** | **Get** /open_api/v1.0/qianchuan/suggest_bid/ | 
*QianchuanSuggestBudgetV10Api* | **OpenApiV10QianchuanSuggestBudgetGet** | **Get** /open_api/v1.0/qianchuan/suggest/budget/ | 
*QianchuanSuggestRoiGoalV10Api* | **OpenApiV10QianchuanSuggestRoiGoalGet** | **Get** /open_api/v1.0/qianchuan/suggest/roi/goal/ | 
*QianchuanTodayLiveRoomDetailGetV10Api* | **OpenApiV10QianchuanTodayLiveRoomDetailGetGet** | **Get** /open_api/v1.0/qianchuan/today_live/room/detail/get/ | 
*QianchuanTodayLiveRoomFlowPerformanceGetV10Api* | **OpenApiV10QianchuanTodayLiveRoomFlowPerformanceGetGet** | **Get** /open_api/v1.0/qianchuan/today_live/room/flow_performance/get/ | 
*QianchuanTodayLiveRoomGetV10Api* | **OpenApiV10QianchuanTodayLiveRoomGetGet** | **Get** /open_api/v1.0/qianchuan/today_live/room/get/ | 
*QianchuanTodayLiveRoomProductListGetV10Api* | **OpenApiV10QianchuanTodayLiveRoomProductListGetGet** | **Get** /open_api/v1.0/qianchuan/today_live/room/product_list/get/ | 
*QianchuanTodayLiveRoomUserGetV10Api* | **OpenApiV10QianchuanTodayLiveRoomUserGetGet** | **Get** /open_api/v1.0/qianchuan/today_live/room/user/get/ | 
*QianchuanToolsAllowCouponV10Api* | **OpenApiV10QianchuanToolsAllowCouponGet** | **Get** /open_api/v1.0/qianchuan/tools/allow_coupon/ | 
*QianchuanToolsAwemeAuthV10Api* | **OpenApiV10QianchuanToolsAwemeAuthPost** | **Post** /open_api/v1.0/qianchuan/tools/aweme_auth/ | 
*QianchuanToolsEstimateAudienceV10Api* | **OpenApiV10QianchuanToolsEstimateAudienceGet** | **Get** /open_api/v1.0/qianchuan/tools/estimate_audience/ | 
*QianchuanToolsGrayV10Api* | **OpenApiV10QianchuanToolsGrayGet** | **Get** /open_api/v1.0/qianchuan/tools/gray/ | 
*QianchuanToolsSmartBoostAdBoostReportGetV10Api* | **OpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGet** | **Get** /open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/report/get/ | 
*QianchuanToolsSmartBoostAdBoostSetV10Api* | **OpenApiV10QianchuanToolsSmartBoostAdBoostSetPost** | **Post** /open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/set/ | 
*QianchuanToolsSmartBoostAdBoostStatusGetV10Api* | **OpenApiV10QianchuanToolsSmartBoostAdBoostStatusGetGet** | **Get** /open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/status/get/ | 
*QianchuanToolsSmartBoostAdBoostVersionGetV10Api* | **OpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGet** | **Get** /open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/version/get/ | 
*QianchuanUniAwemeAuthorizedGetV10Api* | **OpenApiV10QianchuanUniAwemeAuthorizedGetGet** | **Get** /open_api/v1.0/qianchuan/uni_aweme/authorized/get/ | 
*QianchuanUniPromotionListV10Api* | **OpenApiV10QianchuanUniPromotionListGet** | **Get** /open_api/v1.0/qianchuan/uni_promotion/list/ | 
*QianchuanVideoGetV10Api* | **OpenApiV10QianchuanVideoGetGet** | **Get** /open_api/v1.0/qianchuan/video/get/ | 
*ReportAdGetV2Api* | **OpenApi2ReportAdGetGet** | **Get** /open_api/2/report/ad/get/ | 
*ReportAdvertiserGetV2Api* | **OpenApi2ReportAdvertiserGetGet** | **Get** /open_api/2/report/advertiser/get/ | 
*ReportAgentGetV2V2Api* | **OpenApi2ReportAgentGetV2Get** | **Get** /open_api/2/report/agent/get_v2/ | 
*ReportAudienceAgeV2Api* | **OpenApi2ReportAudienceAgeGet** | **Get** /open_api/2/report/audience/age/ | 
*ReportAudienceAwemeListV2Api* | **OpenApi2ReportAudienceAwemeListGet** | **Get** /open_api/2/report/audience/aweme/list/ | 
*ReportAudienceCityV2Api* | **OpenApi2ReportAudienceCityGet** | **Get** /open_api/2/report/audience/city/ | 
*ReportAudienceGenderV2Api* | **OpenApi2ReportAudienceGenderGet** | **Get** /open_api/2/report/audience/gender/ | 
*ReportAudienceInterestActionListV2Api* | **OpenApi2ReportAudienceInterestActionListGet** | **Get** /open_api/2/report/audience/interest_action/list/ | 
*ReportAudienceProvinceV2Api* | **OpenApi2ReportAudienceProvinceGet** | **Get** /open_api/2/report/audience/province/ | 
*ReportCampaignGetV2Api* | **OpenApi2ReportCampaignGetGet** | **Get** /open_api/2/report/campaign/get/ | 
*ReportCreativeGetV2Api* | **OpenApi2ReportCreativeGetGet** | **Get** /open_api/2/report/creative/get/ | 
*ReportCustomConfigGetV30Api* | **OpenApiV30ReportCustomConfigGetGet** | **Get** /open_api/v3.0/report/custom/config/get/ | 
*ReportCustomGetV30Api* | **OpenApiV30ReportCustomGetGet** | **Get** /open_api/v3.0/report/custom/get/ | 
*ReportLiveRoomAnalysisGetV2Api* | **OpenApi2ReportLiveRoomAnalysisGetGet** | **Get** /open_api/2/report/live_room/analysis/get/ | 
*ReportLiveRoomAttributeGetV2Api* | **OpenApi2ReportLiveRoomAttributeGetGet** | **Get** /open_api/2/report/live_room/attribute/get/ | 
*ReportLiveRoomAudiencePortraitGetV2Api* | **OpenApi2ReportLiveRoomAudiencePortraitGetGet** | **Get** /open_api/2/report/live_room/audience/portrait/get/ | 
*ReportLiveRoomFlowCategoryGetV2Api* | **OpenApi2ReportLiveRoomFlowCategoryGetGet** | **Get** /open_api/2/report/live_room/flow_category/get/ | 
*ReportLiveRoomProductGetV2Api* | **OpenApi2ReportLiveRoomProductGetGet** | **Get** /open_api/2/report/live_room/product/get/ | 
*ReportRtaExpGetV2Api* | **OpenApi2ReportRtaExpGetGet** | **Get** /open_api/2/report/rta_exp/get/ | 
*ReportRtaExpLocalDailyGetV30Api* | **OpenApiV30ReportRtaExpLocalDailyGetGet** | **Get** /open_api/v3.0/report/rta_exp_local_daily/get/ | 
*ReportRtaExpLocalHourlyGetV30Api* | **OpenApiV30ReportRtaExpLocalHourlyGetGet** | **Get** /open_api/v3.0/report/rta_exp_local_hourly/get/ | 
*ReportRubeexGetV2Api* | **OpenApi2ReportRubeexGetGet** | **Get** /open_api/2/report/rubeex/get/ | 
*ReportSitePageV2Api* | **OpenApi2ReportSitePageGet** | **Get** /open_api/2/report/site/page/ | 
*ReportVideoFrameGetV2Api* | **OpenApi2ReportVideoFrameGetGet** | **Get** /open_api/2/report/video/frame/get/ | 
*ServeMarketOrderGetV10Api* | **OpenApiV10ServeMarketOrderGetGet** | **Get** /open_api/v1.0/serve_market/order/get/ | 
*SpiTaskGetV2Api* | **OpenApi2SpiTaskGetGet** | **Get** /open_api/2/spi_task/get/ | 
*StarClueGetV2Api* | **OpenApi2StarClueGetGet** | **Get** /open_api/2/star/clue/get/ | 
*StarDemandListV2Api* | **OpenApi2StarDemandListGet** | **Get** /open_api/2/star/demand/list/ | 
*StarDemandOrderListV2Api* | **OpenApi2StarDemandOrderListGet** | **Get** /open_api/2/star/demand/order/list/ | 
*StarReportOrderOverviewGetV2Api* | **OpenApi2StarReportOrderOverviewGetGet** | **Get** /open_api/2/star/report/order_overview/get/ | 
*StarReportOrderUserDistributionGetV2Api* | **OpenApi2StarReportOrderUserDistributionGetGet** | **Get** /open_api/2/star/report/order_user_distribution/get/ | 
*SubscribeAccountsAddV30Api* | **OpenApiV30SubscribeAccountsAddPost** | **Post** /open_api/v3.0/subscribe/accounts/add/ | 
*SubscribeAccountsListV30Api* | **OpenApiV30SubscribeAccountsListGet** | **Get** /open_api/v3.0/subscribe/accounts/list/ | 
*SubscribeAccountsRemoveV30Api* | **OpenApiV30SubscribeAccountsRemovePost** | **Post** /open_api/v3.0/subscribe/accounts/remove/ | 
*SuggWordsV30Api* | **OpenApiV30SuggWordsPost** | **Post** /open_api/v3.0/sugg_words/ | 
*ToolQuickAppManagementQuickAppGetV2Api* | **OpenApi2ToolQuickAppManagementQuickAppGetGet** | **Get** /open_api/2/tool/quick_app_management/quick_app/get/ | 
*ToolsAbTestCreateV2Api* | **OpenApi2ToolsAbTestCreatePost** | **Post** /open_api/2/tools/ab_test/create/ | 
*ToolsAbTestDeleteV2Api* | **OpenApi2ToolsAbTestDeletePost** | **Post** /open_api/2/tools/ab_test/delete/ | 
*ToolsAbTestInfoGetV2Api* | **OpenApi2ToolsAbTestInfoGetGet** | **Get** /open_api/2/tools/ab_test_info/get/ | 
*ToolsAbTestListGetV2Api* | **OpenApi2ToolsAbTestListGetGet** | **Get** /open_api/2/tools/ab_test_list/get/ | 
*ToolsAbTestStopV2Api* | **OpenApi2ToolsAbTestStopPost** | **Post** /open_api/2/tools/ab_test/stop/ | 
*ToolsAbTestUpdateV2Api* | **OpenApi2ToolsAbTestUpdatePost** | **Post** /open_api/2/tools/ab_test/update/ | 
*ToolsActionTextGetV2Api* | **OpenApi2ToolsActionTextGetGet** | **Get** /open_api/2/tools/action_text/get/ | 
*ToolsAdPreviewQrcodeGetV30Api* | **OpenApiV30ToolsAdPreviewQrcodeGetGet** | **Get** /open_api/v3.0/tools/ad_preview/qrcode_get/ | 
*ToolsAdRaiseEstimateGetV2Api* | **OpenApi2ToolsAdRaiseEstimateGetGet** | **Get** /open_api/2/tools/ad_raise_estimate/get/ | 
*ToolsAdRaiseResultGetV2V2Api* | **OpenApi2ToolsAdRaiseResultGetV2Get** | **Get** /open_api/2/tools/ad_raise_result/get_v2/ | 
*ToolsAdRaiseStatusGetV2Api* | **OpenApi2ToolsAdRaiseStatusGetGet** | **Get** /open_api/2/tools/ad_raise_status/get/ | 
*ToolsAdRaiseVersionGetV2Api* | **OpenApi2ToolsAdRaiseVersionGetGet** | **Get** /open_api/2/tools/ad_raise_version/get/ | 
*ToolsAdStatExtraInfoGetV2Api* | **OpenApi2ToolsAdStatExtraInfoGetGet** | **Get** /open_api/2/tools/ad_stat_extra_info/get/ | 
*ToolsAdminInfoV2Api* | **OpenApi2ToolsAdminInfoGet** | **Get** /open_api/2/tools/admin/info/ | 
*ToolsAppManagementAndroidAppListV2Api* | **OpenApi2ToolsAppManagementAndroidAppListGet** | **Get** /open_api/2/tools/app_management/android_app/list/ | 
*ToolsAppManagementAndroidBasicPackageGetV2Api* | **OpenApi2ToolsAppManagementAndroidBasicPackageGetGet** | **Get** /open_api/2/tools/app_management/android_basic_package/get/ | 
*ToolsAppManagementAndroidBasicPackagePublishV2Api* | **OpenApi2ToolsAppManagementAndroidBasicPackagePublishPost** | **Post** /open_api/2/tools/app_management/android_basic_package/publish/ | 
*ToolsAppManagementAndroidBasicPackageUpdateV2Api* | **OpenApi2ToolsAppManagementAndroidBasicPackageUpdatePost** | **Post** /open_api/2/tools/app_management/android_basic_package/update/ | 
*ToolsAppManagementAppGetV2Api* | **OpenApi2ToolsAppManagementAppGetGet** | **Get** /open_api/2/tools/app_management/app/get/ | 
*ToolsAppManagementBookingGetV2Api* | **OpenApi2ToolsAppManagementBookingGetGet** | **Get** /open_api/2/tools/app_management/booking/get/ | 
*ToolsAppManagementBookingRecordsGetV2Api* | **OpenApi2ToolsAppManagementBookingRecordsGetGet** | **Get** /open_api/2/tools/app_management/booking_records/get/ | 
*ToolsAppManagementBpShareCancelV2Api* | **OpenApi2ToolsAppManagementBpShareCancelPost** | **Post** /open_api/2/tools/app_management/bp_share/cancel/ | 
*ToolsAppManagementBpShareV2Api* | **OpenApi2ToolsAppManagementBpSharePost** | **Post** /open_api/2/tools/app_management/bp_share/ | 
*ToolsAppManagementExtendPackageCreateV2Api* | **OpenApi2ToolsAppManagementExtendPackageCreatePost** | **Post** /open_api/2/tools/app_management/extend_package/create/ | 
*ToolsAppManagementExtendPackageCreateV2V2Api* | **OpenApi2ToolsAppManagementExtendPackageCreateV2Post** | **Post** /open_api/2/tools/app_management/extend_package/create_v2/ | 
*ToolsAppManagementExtendPackageListV2Api* | **OpenApi2ToolsAppManagementExtendPackageListGet** | **Get** /open_api/2/tools/app_management/extend_package/list/ | 
*ToolsAppManagementExtendPackageListV2V2Api* | **OpenApi2ToolsAppManagementExtendPackageListV2Get** | **Get** /open_api/2/tools/app_management/extend_package/list_v2/ | 
*ToolsAppManagementExtendPackageUpdateV2Api* | **OpenApi2ToolsAppManagementExtendPackageUpdatePost** | **Post** /open_api/2/tools/app_management/extend_package/update/ | 
*ToolsAppManagementIndustryInfoListV2Api* | **OpenApi2ToolsAppManagementIndustryInfoListGet** | **Get** /open_api/2/tools/app_management/industry_info/list/ | 
*ToolsAppManagementShareAccountListV2Api* | **OpenApi2ToolsAppManagementShareAccountListGet** | **Get** /open_api/2/tools/app_management/share_account/list/ | 
*ToolsAppManagementUpdateAuthorizationV2Api* | **OpenApi2ToolsAppManagementUpdateAuthorizationPost** | **Post** /open_api/2/tools/app_management/update/authorization/ | 
*ToolsAppManagementUploadTaskCreateV2Api* | **OpenApi2ToolsAppManagementUploadTaskCreatePost** | **Post** /open_api/2/tools/app_management/upload_task/create/ | 
*ToolsAppManagementUploadTaskListV2Api* | **OpenApi2ToolsAppManagementUploadTaskListGet** | **Get** /open_api/2/tools/app_management/upload_task/list/ | 
*ToolsAwemeAuthListV2Api* | **OpenApi2ToolsAwemeAuthListGet** | **Get** /open_api/2/tools/aweme_auth_list/ | 
*ToolsAwemeAuthorInfoGetV2Api* | **OpenApi2ToolsAwemeAuthorInfoGetGet** | **Get** /open_api/2/tools/aweme_author_info/get/ | 
*ToolsAwemeBannedCreateV30Api* | **OpenApiV30ToolsAwemeBannedCreatePost** | **Post** /open_api/v3.0/tools/aweme_banned/create/ | 
*ToolsAwemeBannedDeleteV30Api* | **OpenApiV30ToolsAwemeBannedDeletePost** | **Post** /open_api/v3.0/tools/aweme_banned/delete/ | 
*ToolsAwemeBannedListV30Api* | **OpenApiV30ToolsAwemeBannedListGet** | **Get** /open_api/v3.0/tools/aweme_banned/list/ | 
*ToolsAwemeCategoryTopAuthorGetV2Api* | **OpenApi2ToolsAwemeCategoryTopAuthorGetGet** | **Get** /open_api/2/tools/aweme_category_top_author/get/ | 
*ToolsAwemeInfoSearchV2Api* | **OpenApi2ToolsAwemeInfoSearchGet** | **Get** /open_api/2/tools/aweme_info_search/ | 
*ToolsAwemeMultiLevelCategoryGetV2Api* | **OpenApi2ToolsAwemeMultiLevelCategoryGetGet** | **Get** /open_api/2/tools/aweme_multi_level_category/get/ | 
*ToolsAwemeSimilarAuthorSearchV2Api* | **OpenApi2ToolsAwemeSimilarAuthorSearchGet** | **Get** /open_api/2/tools/aweme_similar_author_search/ | 
*ToolsBidSuggestV2Api* | **OpenApi2ToolsBidSuggestGet** | **Get** /open_api/2/tools/bid/suggest/ | 
*ToolsBidsSuggestV30Api* | **OpenApiV30ToolsBidsSuggestGet** | **Get** /open_api/v3.0/tools/bids/suggest/ | 
*ToolsBpAssetManagementShareCancelV30Api* | **OpenApiV30ToolsBpAssetManagementShareCancelPost** | **Post** /open_api/v3.0/tools/bp_asset_management/share/cancel/ | 
*ToolsBpAssetManagementShareGetV30Api* | **OpenApiV30ToolsBpAssetManagementShareGetGet** | **Get** /open_api/v3.0/tools/bp_asset_management/share/get/ | 
*ToolsBpAssetManagementShareV30Api* | **OpenApiV30ToolsBpAssetManagementSharePost** | **Post** /open_api/v3.0/tools/bp_asset_management/share/ | 
*ToolsClueCallbackV2Api* | **OpenApi2ToolsClueCallbackPost** | **Post** /open_api/2/tools/clue/callback/ | 
*ToolsClueFormDetailV2Api* | **OpenApi2ToolsClueFormDetailGet** | **Get** /open_api/2/tools/clue/form/detail/ | 
*ToolsClueFormGetV2Api* | **OpenApi2ToolsClueFormGetGet** | **Get** /open_api/2/tools/clue/form/get/ | 
*ToolsClueGetV2Api* | **OpenApi2ToolsClueGetGet** | **Get** /open_api/2/tools/clue/get/ | 
*ToolsClueSmartPhoneGetV2Api* | **OpenApi2ToolsClueSmartPhoneGetGet** | **Get** /open_api/2/tools/clue/smart_phone/get/ | 
*ToolsCommentGetV30Api* | **OpenApiV30ToolsCommentGetGet** | **Get** /open_api/v3.0/tools/comment/get/ | 
*ToolsCommentHideV30Api* | **OpenApiV30ToolsCommentHidePost** | **Post** /open_api/v3.0/tools/comment/hide/ | 
*ToolsCommentMid2itemIdV30Api* | **OpenApiV30ToolsCommentMid2itemIdGet** | **Get** /open_api/v3.0/tools/comment/mid2item_id/ | 
*ToolsCommentReplyGetV30Api* | **OpenApiV30ToolsCommentReplyGetGet** | **Get** /open_api/v3.0/tools/comment_reply/get/ | 
*ToolsCommentReplyV30Api* | **OpenApiV30ToolsCommentReplyPost** | **Post** /open_api/v3.0/tools/comment/reply/ | 
*ToolsCommentStickOnTopV30Api* | **OpenApiV30ToolsCommentStickOnTopPost** | **Post** /open_api/v3.0/tools/comment/stick_on_top/ | 
*ToolsCommentTermsBannedAddV30Api* | **OpenApiV30ToolsCommentTermsBannedAddPost** | **Post** /open_api/v3.0/tools/comment/terms_banned/add/ | 
*ToolsCommentTermsBannedDeleteV30Api* | **OpenApiV30ToolsCommentTermsBannedDeletePost** | **Post** /open_api/v3.0/tools/comment/terms_banned/delete/ | 
*ToolsCommentTermsBannedGetV30Api* | **OpenApiV30ToolsCommentTermsBannedGetGet** | **Get** /open_api/v3.0/tools/comment/terms_banned/get/ | 
*ToolsCommentTermsBannedUpdateV30Api* | **OpenApiV30ToolsCommentTermsBannedUpdatePost** | **Post** /open_api/v3.0/tools/comment/terms_banned/update/ | 
*ToolsCountryInfoV2Api* | **OpenApi2ToolsCountryInfoGet** | **Get** /open_api/2/tools/country/info/ | 
*ToolsCreativeWordSelectV2Api* | **OpenApi2ToolsCreativeWordSelectGet** | **Get** /open_api/2/tools/creative_word/select/ | 
*ToolsDiagnosisAdGetV2V2Api* | **OpenApi2ToolsDiagnosisAdGetV2Get** | **Get** /open_api/2/tools/diagnosis/ad/get_v2/ | 
*ToolsDiagnosisSuggestionGetV2Api* | **OpenApi2ToolsDiagnosisSuggestionGetGet** | **Get** /open_api/2/tools/diagnosis/suggestion/get/ | 
*ToolsDownloadPackageGetV2Api* | **OpenApi2ToolsDownloadPackageGetGet** | **Get** /open_api/2/tools/download/package/get/ | 
*ToolsDownloadPackageParseV2Api* | **OpenApi2ToolsDownloadPackageParsePost** | **Post** /open_api/2/tools/download/package/parse/ | 
*ToolsEstimateAudienceV2Api* | **OpenApi2ToolsEstimateAudienceGet** | **Get** /open_api/2/tools/estimate_audience/ | 
*ToolsEstimatedPriceGetV2Api* | **OpenApi2ToolsEstimatedPriceGetGet** | **Get** /open_api/2/tools/estimated_price/get/ | 
*ToolsEventAssetsGetV2Api* | **OpenApi2ToolsEventAssetsGetGet** | **Get** /open_api/2/tools/event/assets/get/ | 
*ToolsGrayGetV30Api* | **OpenApiV30ToolsGrayGetGet** | **Get** /open_api/v3.0/tools/gray/get/ | 
*ToolsIesAccountSearchV2Api* | **OpenApi2ToolsIesAccountSearchGet** | **Get** /open_api/2/tools/ies_account_search/ | 
*ToolsIndustryGetV2Api* | **OpenApi2ToolsIndustryGetGet** | **Get** /open_api/2/tools/industry/get/ | 
*ToolsInterestActionActionCategoryV2Api* | **OpenApi2ToolsInterestActionActionCategoryGet** | **Get** /open_api/2/tools/interest_action/action/category/ | 
*ToolsInterestActionActionKeywordV2Api* | **OpenApi2ToolsInterestActionActionKeywordGet** | **Get** /open_api/2/tools/interest_action/action/keyword/ | 
*ToolsInterestActionId2wordV2Api* | **OpenApi2ToolsInterestActionId2wordGet** | **Get** /open_api/2/tools/interest_action/id2word/ | 
*ToolsInterestActionInterestCategoryV2Api* | **OpenApi2ToolsInterestActionInterestCategoryGet** | **Get** /open_api/2/tools/interest_action/interest/category/ | 
*ToolsInterestActionInterestKeywordV2Api* | **OpenApi2ToolsInterestActionInterestKeywordGet** | **Get** /open_api/2/tools/interest_action/interest/keyword/ | 
*ToolsInterestActionKeywordSuggestV2Api* | **OpenApi2ToolsInterestActionKeywordSuggestGet** | **Get** /open_api/2/tools/interest_action/keyword/suggest/ | 
*ToolsKeyActionGetV2Api* | **OpenApi2ToolsKeyActionGetGet** | **Get** /open_api/2/tools/key_action/get/ | 
*ToolsKeywordsBidRatioCreateV30Api* | **OpenApiV30ToolsKeywordsBidRatioCreatePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/create/ | 
*ToolsKeywordsBidRatioDeleteV30Api* | **OpenApiV30ToolsKeywordsBidRatioDeletePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/delete/ | 
*ToolsKeywordsBidRatioGetV30Api* | **OpenApiV30ToolsKeywordsBidRatioGetGet** | **Get** /open_api/v3.0/tools/keywords_bid_ratio/get/ | 
*ToolsKeywordsBidRatioUpdateV30Api* | **OpenApiV30ToolsKeywordsBidRatioUpdatePost** | **Post** /open_api/v3.0/tools/keywords_bid_ratio/update/ | 
*ToolsKeywordsProjectInfoGetV30Api* | **OpenApiV30ToolsKeywordsProjectInfoGetGet** | **Get** /open_api/v3.0/tools/keywords_project_info/get/ | 
*ToolsLandingGroupCreateV2Api* | **OpenApi2ToolsLandingGroupCreatePost** | **Post** /open_api/2/tools/landing_group/create/ | 
*ToolsLandingGroupGetV2Api* | **OpenApi2ToolsLandingGroupGetGet** | **Get** /open_api/2/tools/landing_group/get/ | 
*ToolsLandingGroupSiteOptStatusUpdateV2Api* | **OpenApi2ToolsLandingGroupSiteOptStatusUpdatePost** | **Post** /open_api/2/tools/landing_group/site_opt_status/update/ | 
*ToolsLandingGroupUpdateV2Api* | **OpenApi2ToolsLandingGroupUpdatePost** | **Post** /open_api/2/tools/landing_group/update/ | 
*ToolsLiveAuthorizeListV2Api* | **OpenApi2ToolsLiveAuthorizeListGet** | **Get** /open_api/2/tools/live_authorize/list/ | 
*ToolsLogSearchV2Api* | **OpenApi2ToolsLogSearchGet** | **Get** /open_api/2/tools/log_search/ | 
*ToolsOrangeSiteGetV30Api* | **OpenApiV30ToolsOrangeSiteGetGet** | **Get** /open_api/v3.0/tools/orange_site/get/ | 
*ToolsPlayableCloudGameListV2Api* | **OpenApi2ToolsPlayableCloudGameListGet** | **Get** /open_api/2/tools/playable/cloud_game/list/ | 
*ToolsPlayableListGetV2Api* | **OpenApi2ToolsPlayableListGetGet** | **Get** /open_api/2/tools/playable_list/get/ | 
*ToolsPreAuditGetV2Api* | **OpenApi2ToolsPreAuditGetGet** | **Get** /open_api/2/tools/pre_audit/get/ | 
*ToolsPreAuditSendV2Api* | **OpenApi2ToolsPreAuditSendPost** | **Post** /open_api/2/tools/pre_audit/send/ | 
*ToolsPrivativeWordBatchGetV30Api* | **OpenApiV30ToolsPrivativeWordBatchGetPost** | **Post** /open_api/v3.0/tools/privative_word/batch_get/ | 
*ToolsPrivativeWordGetV2Api* | **OpenApi2ToolsPrivativeWordGetGet** | **Get** /open_api/2/tools/privative_word/get/ | 
*ToolsPrivativeWordProjectAddV30Api* | **OpenApiV30ToolsPrivativeWordProjectAddPost** | **Post** /open_api/v3.0/tools/privative_word/project/add/ | 
*ToolsPrivativeWordProjectUpdateV30Api* | **OpenApiV30ToolsPrivativeWordProjectUpdatePost** | **Post** /open_api/v3.0/tools/privative_word/project/update/ | 
*ToolsPrivativeWordPromotionAddV30Api* | **OpenApiV30ToolsPrivativeWordPromotionAddPost** | **Post** /open_api/v3.0/tools/privative_word/promotion/add/ | 
*ToolsPrivativeWordPromotionUpdateV30Api* | **OpenApiV30ToolsPrivativeWordPromotionUpdatePost** | **Post** /open_api/v3.0/tools/privative_word/promotion/update/ | 
*ToolsPromotionCardRecommendGetV2Api* | **OpenApi2ToolsPromotionCardRecommendGetGet** | **Get** /open_api/2/tools/promotion_card/recommend/get/ | 
*ToolsPromotionCardRecommendTitleGetV2Api* | **OpenApi2ToolsPromotionCardRecommendTitleGetGet** | **Get** /open_api/2/tools/promotion_card/recommend_title/get/ | 
*ToolsPromotionDiagnosisSuggestionAcceptV30Api* | **OpenApiV30ToolsPromotionDiagnosisSuggestionAcceptPost** | **Post** /open_api/v3.0/tools/promotion_diagnosis/suggestion/accept/ | 
*ToolsPromotionDiagnosisSuggestionGetV30Api* | **OpenApiV30ToolsPromotionDiagnosisSuggestionGetGet** | **Get** /open_api/v3.0/tools/promotion_diagnosis/suggestion/get/ | 
*ToolsPromotionRaiseSetV30Api* | **OpenApiV30ToolsPromotionRaiseSetPost** | **Post** /open_api/v3.0/tools/promotion_raise/set/ | 
*ToolsPromotionRaiseStatusCurrentIdsGetV30Api* | **OpenApiV30ToolsPromotionRaiseStatusCurrentIdsGetGet** | **Get** /open_api/v3.0/tools/promotion_raise_status_current_ids/get/ | 
*ToolsPromotionRaiseStatusGetV30Api* | **OpenApiV30ToolsPromotionRaiseStatusGetGet** | **Get** /open_api/v3.0/tools/promotion_raise_status/get/ | 
*ToolsPromotionRaiseStopV30Api* | **OpenApiV30ToolsPromotionRaiseStopPost** | **Post** /open_api/v3.0/tools/promotion_raise/stop/ | 
*ToolsPromotionRaiseVersionGetV30Api* | **OpenApiV30ToolsPromotionRaiseVersionGetGet** | **Get** /open_api/v3.0/tools/promotion_raise_version/get/ | 
*ToolsQuotaGetV2Api* | **OpenApi2ToolsQuotaGetGet** | **Get** /open_api/2/tools/quota/get/ | 
*ToolsRegionGetV2Api* | **OpenApi2ToolsRegionGetGet** | **Get** /open_api/2/tools/region/get/ | 
*ToolsRtaGetInfoV2Api* | **OpenApi2ToolsRtaGetInfoGet** | **Get** /open_api/2/tools/rta/get_info/ | 
*ToolsRtaGetV2Api* | **OpenApi2ToolsRtaGetGet** | **Get** /open_api/2/tools/rta/get/ | 
*ToolsRtaScopeGetV30Api* | **OpenApiV30ToolsRtaScopeGetGet** | **Get** /open_api/v3.0/tools/rta/scope/get/ | 
*ToolsRtaSetScopeV2Api* | **OpenApi2ToolsRtaSetScopePost** | **Post** /open_api/2/tools/rta/set_scope/ | 
*ToolsRtaStatusUpdateV2Api* | **OpenApi2ToolsRtaStatusUpdatePost** | **Post** /open_api/2/tools/rta/status_update/ | 
*ToolsRubeexGetV2Api* | **OpenApi2ToolsRubeexGetGet** | **Get** /open_api/2/tools/rubeex/get/ | 
*ToolsRubeexPlayableListV2Api* | **OpenApi2ToolsRubeexPlayableListGet** | **Get** /open_api/2/tools/rubeex_playable/list/ | 
*ToolsRubeexRemarkV2Api* | **OpenApi2ToolsRubeexRemarkGet** | **Get** /open_api/2/tools/rubeex/remark/ | 
*ToolsRubeexVersionGetV2Api* | **OpenApi2ToolsRubeexVersionGetGet** | **Get** /open_api/2/tools/rubeex/version/get/ | 
*ToolsSearchBidRatioGetV2Api* | **OpenApi2ToolsSearchBidRatioGetGet** | **Get** /open_api/2/tools/search_bid_ratio/get/ | 
*ToolsSiteCopyV2Api* | **OpenApi2ToolsSiteCopyPost** | **Post** /open_api/2/tools/site/copy/ | 
*ToolsSiteCreateV2Api* | **OpenApi2ToolsSiteCreatePost** | **Post** /open_api/2/tools/site/create/ | 
*ToolsSiteFormsListV2Api* | **OpenApi2ToolsSiteFormsListGet** | **Get** /open_api/2/tools/site/forms/list/ | 
*ToolsSiteGetV2Api* | **OpenApi2ToolsSiteGetGet** | **Get** /open_api/2/tools/site/get/ | 
*ToolsSiteHandselV2Api* | **OpenApi2ToolsSiteHandselPost** | **Post** /open_api/2/tools/site/handsel/ | 
*ToolsSitePreviewV2Api* | **OpenApi2ToolsSitePreviewGet** | **Get** /open_api/2/tools/site/preview/ | 
*ToolsSiteReadV2Api* | **OpenApi2ToolsSiteReadGet** | **Get** /open_api/2/tools/site/read/ | 
*ToolsSiteTemplateCreateV2Api* | **OpenApi2ToolsSiteTemplateCreatePost** | **Post** /open_api/2/tools/site_template/create/ | 
*ToolsSiteTemplateGetV2Api* | **OpenApi2ToolsSiteTemplateGetGet** | **Get** /open_api/2/tools/site_template/get/ | 
*ToolsSiteTemplatePicUrlGetV2Api* | **OpenApi2ToolsSiteTemplatePicUrlGetGet** | **Get** /open_api/2/tools/site_template/pic_url/get/ | 
*ToolsSiteTemplatePreviewV2Api* | **OpenApi2ToolsSiteTemplatePreviewGet** | **Get** /open_api/2/tools/site_template/preview/ | 
*ToolsSiteTemplateSiteCreateV2Api* | **OpenApi2ToolsSiteTemplateSiteCreatePost** | **Post** /open_api/2/tools/site_template/site/create/ | 
*ToolsSiteUpdateStatusV2Api* | **OpenApi2ToolsSiteUpdateStatusPost** | **Post** /open_api/2/tools/site/update_status/ | 
*ToolsSiteUpdateV2Api* | **OpenApi2ToolsSiteUpdatePost** | **Post** /open_api/2/tools/site/update/ | 
*ToolsSuggestBudgetGetV30Api* | **OpenApiV30ToolsSuggestBudgetGetGet** | **Get** /open_api/v3.0/tools/suggest_budget/get/ | 
*ToolsTaskRaiseCreateV2Api* | **OpenApi2ToolsTaskRaiseCreatePost** | **Post** /open_api/2/tools/task_raise/create/ | 
*ToolsTaskRaiseDataGetV2Api* | **OpenApi2ToolsTaskRaiseDataGetGet** | **Get** /open_api/2/tools/task_raise/data/get/ | 
*ToolsTaskRaiseGetV2Api* | **OpenApi2ToolsTaskRaiseGetGet** | **Get** /open_api/2/tools/task_raise/get/ | 
*ToolsTaskRaiseOptimizationIdsGetV2Api* | **OpenApi2ToolsTaskRaiseOptimizationIdsGetGet** | **Get** /open_api/2/tools/task_raise/optimization_ids/get/ | 
*ToolsTaskRaiseStatusStopV2Api* | **OpenApi2ToolsTaskRaiseStatusStopPost** | **Post** /open_api/2/tools/task_raise/status/stop/ | 
*ToolsThirdSiteCreateV2Api* | **OpenApi2ToolsThirdSiteCreatePost** | **Post** /open_api/2/tools/third_site/create/ | 
*ToolsThirdSiteDeleteV2Api* | **OpenApi2ToolsThirdSiteDeletePost** | **Post** /open_api/2/tools/third_site/delete/ | 
*ToolsThirdSiteGetV2Api* | **OpenApi2ToolsThirdSiteGetGet** | **Get** /open_api/2/tools/third_site/get/ | 
*ToolsThirdSitePreviewV2Api* | **OpenApi2ToolsThirdSitePreviewGet** | **Get** /open_api/2/tools/third_site/preview/ | 
*ToolsThirdSiteUpdateV2Api* | **OpenApi2ToolsThirdSiteUpdatePost** | **Post** /open_api/2/tools/third_site/update/ | 
*ToolsUnionFlowPackageCreateV2Api* | **OpenApi2ToolsUnionFlowPackageCreatePost** | **Post** /open_api/2/tools/union/flow_package/create/ | 
*ToolsUnionFlowPackageDeleteV2Api* | **OpenApi2ToolsUnionFlowPackageDeletePost** | **Post** /open_api/2/tools/union/flow_package/delete/ | 
*ToolsUnionFlowPackageGetV2Api* | **OpenApi2ToolsUnionFlowPackageGetGet** | **Get** /open_api/2/tools/union/flow_package/get/ | 
*ToolsUnionFlowPackagePromotionReportV30Api* | **OpenApiV30ToolsUnionFlowPackagePromotionReportGet** | **Get** /open_api/v3.0/tools/union/flow_package/promotion/report/ | 
*ToolsUnionFlowPackageReportV2Api* | **OpenApi2ToolsUnionFlowPackageReportGet** | **Get** /open_api/2/tools/union/flow_package/report/ | 
*ToolsUnionFlowPackageUpdateV2Api* | **OpenApi2ToolsUnionFlowPackageUpdatePost** | **Post** /open_api/2/tools/union/flow_package/update/ | 
*ToolsVideoCheckAvailableAnchorV2Api* | **OpenApi2ToolsVideoCheckAvailableAnchorGet** | **Get** /open_api/2/tools/video/check_available_anchor/ | 
*ToolsVideoCoverSuggestV2Api* | **OpenApi2ToolsVideoCoverSuggestGet** | **Get** /open_api/2/tools/video_cover/suggest/ | 
*ToolsWechatAppletCreateV30Api* | **OpenApiV30ToolsWechatAppletCreatePost** | **Post** /open_api/v3.0/tools/wechat_applet/create/ | 
*ToolsWechatAppletListV30Api* | **OpenApiV30ToolsWechatAppletListGet** | **Get** /open_api/v3.0/tools/wechat_applet/list/ | 
*ToolsWechatAppletUpdateV30Api* | **OpenApiV30ToolsWechatAppletUpdatePost** | **Post** /open_api/v3.0/tools/wechat_applet/update/ | 
*ToolsWechatGameCreateV30Api* | **OpenApiV30ToolsWechatGameCreatePost** | **Post** /open_api/v3.0/tools/wechat_game/create/ | 
*ToolsWechatGameListV30Api* | **OpenApiV30ToolsWechatGameListGet** | **Get** /open_api/v3.0/tools/wechat_game/list/ | 
*UserInfoV2Api* | **OpenApi2UserInfoGet** | **Get** /open_api/2/user/info/ | 


## 问题建议与反馈
如果您在使用SDK过程中有任何问题与建议，请随时登录[开发者官网](https://open.oceanengine.com/labels) ，点击右下角的"咨询"按钮，与我们的客服支持人员联系

## 后续计划
1. 丰富各类场景示例
2. 推出其他语言的SDK
