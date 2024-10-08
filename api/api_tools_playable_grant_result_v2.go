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

// ToolsPlayableGrantResultV2ApiService ToolsPlayableGrantResultV2Api service
type ToolsPlayableGrantResultV2ApiService service

type ApiOpenApi2ToolsPlayableGrantResultGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsPlayableGrantResultV2ApiService
	advertiserId *int64
	taskIds      *[]int64
	startTime    *string
	endTime      *string
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 推送任务id，可根据推送的task_id进行检索，允许传入数量上限200个
func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) TaskIds(taskIds []int64) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.taskIds = &taskIds
	return r
}

// 根据任务创建时间进行过滤，过滤的起始时间，格式YYYY-mm-dd
func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) StartTime(startTime string) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.startTime = &startTime
	return r
}

// 根据任务的创建时间进行过滤，过滤的截止时间，格式YYYY-mm-dd
func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) EndTime(endTime string) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.endTime = &endTime
	return r
}

// 页码，默认1
func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) Page(page int64) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认20，最大200
func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) Execute() (*ToolsPlayableGrantResultV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPlayableGrantResultGet Method for OpenApi2ToolsPlayableGrantResultGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPlayableGrantResultGetRequest
*/
func (a *ToolsPlayableGrantResultV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPlayableGrantResultGetRequest {
	return &ApiOpenApi2ToolsPlayableGrantResultGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPlayableGrantResultV2Response
func (a *ToolsPlayableGrantResultV2ApiService) getExecute(r *ApiOpenApi2ToolsPlayableGrantResultGetRequest) (*ToolsPlayableGrantResultV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPlayableGrantResultV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/playable/grant/result/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.taskIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_ids", r.taskIds)
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
