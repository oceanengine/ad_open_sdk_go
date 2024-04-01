/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

// AdConvertOptimizedTargetGetV30ApiService AdConvertOptimizedTargetGetV30Api service
type AdConvertOptimizedTargetGetV30ApiService service

type ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest struct {
	ctx                context.Context
	ApiService         *AdConvertOptimizedTargetGetV30ApiService
	advertiserId       *int64
	landingType        *AdConvertOptimizedTargetGetV30LandingType
	marketingPurpose   *AdConvertOptimizedTargetGetV30MarketingPurpose
	appType            *AdConvertOptimizedTargetGetV30AppType
	promotionContent   *AdConvertOptimizedTargetGetV30PromotionContent
	convertTracking    *AdConvertOptimizedTargetGetV30ConvertTracking
	adType             *AdConvertOptimizedTargetGetV30AdType
	deliveryRange      *AdConvertOptimizedTargetGetV30DeliveryRange
	externalUrl        *string
	androidPackageName *string
	itunesUrl          *string
	convertId          *int64
	page               *int64
	pageSize           *int64
}

// 广告主ID
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 推广目的
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) LandingType(landingType AdConvertOptimizedTargetGetV30LandingType) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.landingType = &landingType
	return r
}

// 营销目标
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) MarketingPurpose(marketingPurpose AdConvertOptimizedTargetGetV30MarketingPurpose) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.marketingPurpose = &marketingPurpose
	return r
}

// 应用类型
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) AppType(appType AdConvertOptimizedTargetGetV30AppType) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.appType = &appType
	return r
}

// 推广内容
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) PromotionContent(promotionContent AdConvertOptimizedTargetGetV30PromotionContent) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.promotionContent = &promotionContent
	return r
}

// 转化跟踪方式
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) ConvertTracking(convertTracking AdConvertOptimizedTargetGetV30ConvertTracking) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.convertTracking = &convertTracking
	return r
}

// 广告类型
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) AdType(adType AdConvertOptimizedTargetGetV30AdType) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.adType = &adType
	return r
}

// 投放范围
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) DeliveryRange(deliveryRange AdConvertOptimizedTargetGetV30DeliveryRange) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.deliveryRange = &deliveryRange
	return r
}

// 落地页链接
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) ExternalUrl(externalUrl string) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.externalUrl = &externalUrl
	return r
}

// 安卓应用包名
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) AndroidPackageName(androidPackageName string) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.androidPackageName = &androidPackageName
	return r
}

// IOS下载链接
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) ItunesUrl(itunesUrl string) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.itunesUrl = &itunesUrl
	return r
}

// 转化ID
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) ConvertId(convertId int64) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.convertId = &convertId
	return r
}

// 页数
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) Page(page int64) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) Execute() (*AdConvertOptimizedTargetGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) WithLog(enable bool) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdConvertOptimizedTargetGetGet Method for OpenApiV30AdConvertOptimizedTargetGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest
*/
func (a *AdConvertOptimizedTargetGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest {
	return &ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdConvertOptimizedTargetGetV30Response
func (a *AdConvertOptimizedTargetGetV30ApiService) getExecute(r *ApiOpenApiV30AdConvertOptimizedTargetGetGetRequest) (*AdConvertOptimizedTargetGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdConvertOptimizedTargetGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/ad_convert/optimized_target/get/"

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
	if r.marketingPurpose == nil {
		return localVarReturnValue, nil, ReportError("marketingPurpose is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_purpose", r.marketingPurpose)
	if r.appType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_type", r.appType)
	}
	if r.promotionContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_content", r.promotionContent)
	}
	if r.convertTracking != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_tracking", r.convertTracking)
	}
	if r.adType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_type", r.adType)
	}
	if r.deliveryRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_range", r.deliveryRange)
	}
	if r.externalUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_url", r.externalUrl)
	}
	if r.androidPackageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "android_package_name", r.androidPackageName)
	}
	if r.itunesUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itunes_url", r.itunesUrl)
	}
	if r.convertId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "convert_id", r.convertId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
