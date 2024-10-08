/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ReportBrandAdGetV30ApiService ReportBrandAdGetV30Api service
type ReportBrandAdGetV30ApiService service

type ApiOpenApiV30ReportBrandAdGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportBrandAdGetV30ApiService
	advertiserId *int64
	page         *int64
	size         *int64
	landingType  *ReportBrandAdGetV30LandingType
	pricingType  *ReportBrandAdGetV30PricingType
	adIds        *[]string
	startTime    *string
	endTime      *string
}

// 广告主ID
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 分页参数 默认从1开始
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) Page(page int64) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.page = &page
	return r
}

// 分页参数 默认每页最大值为100
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) Size(size int64) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.size = &size
	return r
}

// 推广目的
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) LandingType(landingType ReportBrandAdGetV30LandingType) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.landingType = &landingType
	return r
}

// 计费类型
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) PricingType(pricingType ReportBrandAdGetV30PricingType) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.pricingType = &pricingType
	return r
}

// 计划ID 每次请求上限100个
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) AdIds(adIds []string) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.adIds = &adIds
	return r
}

// 开始时间 格式2023-01-01
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) StartTime(startTime string) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间 格式2023-01-01
func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) EndTime(endTime string) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) Execute() (*ReportBrandAdGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportBrandAdGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportBrandAdGetGet Method for OpenApiV30ReportBrandAdGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportBrandAdGetGetRequest
*/
func (a *ReportBrandAdGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportBrandAdGetGetRequest {
	return &ApiOpenApiV30ReportBrandAdGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportBrandAdGetV30Response
func (a *ReportBrandAdGetV30ApiService) getExecute(r *ApiOpenApiV30ReportBrandAdGetGetRequest) (*ReportBrandAdGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportBrandAdGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/brand/ad/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if r.size == nil {
		return localVarReturnValue, nil, ReportError("size is required and must be specified")
	}
	if *r.size > 100 {
		return localVarReturnValue, nil, ReportError("size must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.landingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	}
	if r.pricingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pricing_type", r.pricingType)
	}
	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
