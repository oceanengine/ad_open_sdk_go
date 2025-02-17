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

// FileRebateCommonDownloadGetDownloadTaskListV2ApiService FileRebateCommonDownloadGetDownloadTaskListV2Api service
type FileRebateCommonDownloadGetDownloadTaskListV2ApiService service

type ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest struct {
	ctx        context.Context
	ApiService *FileRebateCommonDownloadGetDownloadTaskListV2ApiService
	agentId    *int64
	queryId    *string
}

func (r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) AgentId(agentId int64) *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest {
	r.agentId = &agentId
	return r
}

func (r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) QueryId(queryId string) *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest {
	r.queryId = &queryId
	return r
}

func (r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) Execute() (*FileRebateCommonDownloadGetDownloadTaskListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) WithLog(enable bool) *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileRebateCommonDownloadGetDownloadTaskListGet Method for OpenApi2FileRebateCommonDownloadGetDownloadTaskListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest
*/
func (a *FileRebateCommonDownloadGetDownloadTaskListV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest {
	return &ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileRebateCommonDownloadGetDownloadTaskListV2Response
func (a *FileRebateCommonDownloadGetDownloadTaskListV2ApiService) getExecute(r *ApiOpenApi2FileRebateCommonDownloadGetDownloadTaskListGetRequest) (*FileRebateCommonDownloadGetDownloadTaskListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileRebateCommonDownloadGetDownloadTaskListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/rebate/common_download/get_download_task_list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.queryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query_id", r.queryId)
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
