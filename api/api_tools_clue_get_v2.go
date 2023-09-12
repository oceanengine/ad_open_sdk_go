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

// ToolsClueGetV2ApiService ToolsClueGetV2Api service
type ToolsClueGetV2ApiService service

type ApiOpenApi2ToolsClueGetGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsClueGetV2ApiService
	advertiserIds *[]int64
	startTime     *string
	endTime       *string
	page          *int32
	pageSize      *int32
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) AdvertiserIds(advertiserIds []int64) *ApiOpenApi2ToolsClueGetGetRequest {
	r.advertiserIds = &advertiserIds
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) StartTime(startTime string) *ApiOpenApi2ToolsClueGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) EndTime(endTime string) *ApiOpenApi2ToolsClueGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) Page(page int32) *ApiOpenApi2ToolsClueGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsClueGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) Execute() (*ToolsClueGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueGetGet Method for OpenApi2ToolsClueGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueGetGetRequest
*/
func (a *ToolsClueGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueGetGetRequest {
	return &ApiOpenApi2ToolsClueGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueGetV2Response
func (a *ToolsClueGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueGetGetRequest) (*ToolsClueGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsClueGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_ids", r.advertiserIds)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
