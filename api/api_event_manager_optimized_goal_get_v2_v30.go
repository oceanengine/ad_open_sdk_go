/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// EventManagerOptimizedGoalGetV2V30ApiService EventManagerOptimizedGoalGetV2V30Api service
type EventManagerOptimizedGoalGetV2V30ApiService service

type ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest struct {
	ctx                context.Context
	ApiService         *EventManagerOptimizedGoalGetV2V30ApiService
	advertiserId       *int64
	landingType        *EventManagerOptimizedGoalGetV2V30LandingType
	adType             *EventManagerOptimizedGoalGetV2V30AdType
	assetType          *EventManagerOptimizedGoalGetV2V30AssetType
	assetId            *int64
	packageName        *string
	appType            *EventManagerOptimizedGoalGetV2V30AppType
	appPromotionType   *EventManagerOptimizedGoalGetV2V30AppPromotionType
	marketingGoal      *EventManagerOptimizedGoalGetV2V30MarketingGoal
	quickAppId         *int64
	deliveryMode       *EventManagerOptimizedGoalGetV2V30DeliveryMode
	miniProgramId      *string
	dpaAdtype          *EventManagerOptimizedGoalGetV2V30DpaAdtype
	microPromotionType *EventManagerOptimizedGoalGetV2V30MicroPromotionType
	microAppInstanceId *int64
	deliveryType       *EventManagerOptimizedGoalGetV2V30DeliveryType
}

// 广告主id
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 推广目的，允许值：APP 应用推广
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) LandingType(landingType EventManagerOptimizedGoalGetV2V30LandingType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.landingType = &landingType
	return r
}

// 广告类型，允许值：ALL
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AdType(adType EventManagerOptimizedGoalGetV2V30AdType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.adType = &adType
	return r
}

// 资产类型，允许值：APP 应用
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AssetType(assetType EventManagerOptimizedGoalGetV2V30AssetType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.assetType = &assetType
	return r
}

// 资产id
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AssetId(assetId int64) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.assetId = &assetId
	return r
}

// 应用包名称，当asset_type为应用APP时，若asset_id不为空，生效asset_id；若asset_id为空且package_name不为空时，生效package_name
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) PackageName(packageName string) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.packageName = &packageName
	return r
}

// 应用类型，当asset_type为应用APP时必填 可选值：ANDROID 、IOS
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AppType(appType EventManagerOptimizedGoalGetV2V30AppType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.appType = &appType
	return r
}

// 子目标，支持： DOWNLOAD（应用下载）、LAUNCH（应用调起）、RESERVE（预约下载）
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AppPromotionType(appPromotionType EventManagerOptimizedGoalGetV2V30AppPromotionType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.appPromotionType = &appPromotionType
	return r
}

// 营销场景， 允许值：VIDEO_AND_IMAGE 短视频/图片 LIVE直播
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) MarketingGoal(marketingGoal EventManagerOptimizedGoalGetV2V30MarketingGoal) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

// 快应用资产id
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) QuickAppId(quickAppId int64) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.quickAppId = &quickAppId
	return r
}

// 投放模式
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) DeliveryMode(deliveryMode EventManagerOptimizedGoalGetV2V30DeliveryMode) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.deliveryMode = &deliveryMode
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) MiniProgramId(miniProgramId string) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.miniProgramId = &miniProgramId
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) DpaAdtype(dpaAdtype EventManagerOptimizedGoalGetV2V30DpaAdtype) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.dpaAdtype = &dpaAdtype
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) MicroPromotionType(microPromotionType EventManagerOptimizedGoalGetV2V30MicroPromotionType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.microPromotionType = &microPromotionType
	return r
}

// 小程序资产id
func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) MicroAppInstanceId(microAppInstanceId int64) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.microAppInstanceId = &microAppInstanceId
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) DeliveryType(deliveryType EventManagerOptimizedGoalGetV2V30DeliveryType) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.deliveryType = &deliveryType
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) Execute() (*EventManagerOptimizedGoalGetV2V30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) AccessToken(accessToken string) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) WithLog(enable bool) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30EventManagerOptimizedGoalGetV2Get Method for OpenApiV30EventManagerOptimizedGoalGetV2Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest
*/
func (a *EventManagerOptimizedGoalGetV2V30ApiService) Get(ctx context.Context) *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest {
	return &ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerOptimizedGoalGetV2V30Response
func (a *EventManagerOptimizedGoalGetV2V30ApiService) getExecute(r *ApiOpenApiV30EventManagerOptimizedGoalGetV2GetRequest) (*EventManagerOptimizedGoalGetV2V30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EventManagerOptimizedGoalGetV2V30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/event_manager/optimized_goal/get_v2/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.landingType == nil {
		return localVarReturnValue, nil, ReportError("landingType is required and must be specified")
	}
	if r.adType == nil {
		return localVarReturnValue, nil, ReportError("adType is required and must be specified")
	}
	if r.assetType == nil {
		return localVarReturnValue, nil, ReportError("assetType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_type", r.adType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_type", r.assetType)
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
	}
	if r.packageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_name", r.packageName)
	}
	if r.appType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_type", r.appType)
	}
	if r.appPromotionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_promotion_type", r.appPromotionType)
	}
	if r.marketingGoal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	}
	if r.quickAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quick_app_id", r.quickAppId)
	}
	if r.deliveryMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_mode", r.deliveryMode)
	}
	if r.miniProgramId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mini_program_id", r.miniProgramId)
	}
	if r.dpaAdtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpa_adtype", r.dpaAdtype)
	}
	if r.microPromotionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "micro_promotion_type", r.microPromotionType)
	}
	if r.microAppInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "micro_app_instance_id", r.microAppInstanceId)
	}
	if r.deliveryType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_type", r.deliveryType)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
