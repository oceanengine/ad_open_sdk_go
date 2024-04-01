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

// ToolsClueContactLogOverviewQueryV2ApiService ToolsClueContactLogOverviewQueryV2Api service
type ToolsClueContactLogOverviewQueryV2ApiService service

type ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueContactLogOverviewQueryV2ApiService
	advertiserId *int64
	startTime    *int64
	endTime      *int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询开始时间。毫秒级unix时间戳。查询时间跨度不得超过一个月，只能查询近半年的数据。
func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) StartTime(startTime int64) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	r.startTime = &startTime
	return r
}

// 查询结束时间。毫秒级unix时间戳。查询时间跨度不得超过一个月，只能查询近半年的数据。
func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) EndTime(endTime int64) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) Execute() (*ToolsClueContactLogOverviewQueryV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueContactLogOverviewQueryGet Method for OpenApi2ToolsClueContactLogOverviewQueryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest
*/
func (a *ToolsClueContactLogOverviewQueryV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest {
	return &ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueContactLogOverviewQueryV2Response
func (a *ToolsClueContactLogOverviewQueryV2ApiService) getExecute(r *ApiOpenApi2ToolsClueContactLogOverviewQueryGetRequest) (*ToolsClueContactLogOverviewQueryV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueContactLogOverviewQueryV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/contact_log/overview/query/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
