/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// CreativeStrategyListV2ApiService CreativeStrategyListV2Api service
type CreativeStrategyListV2ApiService service

type ApiOpenApi2CreativeStrategyListGetRequest struct {
	ctx           context.Context
	ApiService    *CreativeStrategyListV2ApiService
	advertiserId  *int64
	strategyTypes *[]*CreativeStrategyListV2StrategyTypes
	page          *int32
	pageSize      *int32
}

func (r *ApiOpenApi2CreativeStrategyListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeStrategyListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 策略类型(支持多选)
func (r *ApiOpenApi2CreativeStrategyListGetRequest) StrategyTypes(strategyTypes []*CreativeStrategyListV2StrategyTypes) *ApiOpenApi2CreativeStrategyListGetRequest {
	r.strategyTypes = &strategyTypes
	return r
}

// 页码，从1开始, 默认1
func (r *ApiOpenApi2CreativeStrategyListGetRequest) Page(page int32) *ApiOpenApi2CreativeStrategyListGetRequest {
	r.page = &page
	return r
}

// 分页大小, 取值[1-100], 默认10
func (r *ApiOpenApi2CreativeStrategyListGetRequest) PageSize(pageSize int32) *ApiOpenApi2CreativeStrategyListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2CreativeStrategyListGetRequest) Execute() (*CreativeStrategyListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeStrategyListGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeStrategyListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeStrategyListGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeStrategyListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeStrategyListGet Method for OpenApi2CreativeStrategyListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeStrategyListGetRequest
*/
func (a *CreativeStrategyListV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeStrategyListGetRequest {
	return &ApiOpenApi2CreativeStrategyListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeStrategyListV2Response
func (a *CreativeStrategyListV2ApiService) getExecute(r *ApiOpenApi2CreativeStrategyListGetRequest) (*CreativeStrategyListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *CreativeStrategyListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/strategy/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.strategyTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "strategy_types", r.strategyTypes)
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
