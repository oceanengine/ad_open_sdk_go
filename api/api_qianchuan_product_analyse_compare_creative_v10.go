/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
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

// QianchuanProductAnalyseCompareCreativeV10ApiService QianchuanProductAnalyseCompareCreativeV10Api service
type QianchuanProductAnalyseCompareCreativeV10ApiService service

type ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanProductAnalyseCompareCreativeV10ApiService
	advertiserId *int64
	productId    *int64
	timeRange    *QianchuanProductAnalyseCompareCreativeV10TimeRange
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	r.productId = &productId
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) TimeRange(timeRange QianchuanProductAnalyseCompareCreativeV10TimeRange) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	r.timeRange = &timeRange
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) Execute() (*QianchuanProductAnalyseCompareCreativeV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanProductAnalyseCompareCreativeGet Method for OpenApiV10QianchuanProductAnalyseCompareCreativeGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest
*/
func (a *QianchuanProductAnalyseCompareCreativeV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest {
	return &ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanProductAnalyseCompareCreativeV10Response
func (a *QianchuanProductAnalyseCompareCreativeV10ApiService) getExecute(r *ApiOpenApiV10QianchuanProductAnalyseCompareCreativeGetRequest) (*QianchuanProductAnalyseCompareCreativeV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanProductAnalyseCompareCreativeV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/product/analyse/compare_creative/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.productId == nil {
		return localVarReturnValue, nil, ReportError("productId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
	if r.timeRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_range", r.timeRange)
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
