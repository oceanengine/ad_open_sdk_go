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
	"strings"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsAppManagementUploadTaskListV2ApiService ToolsAppManagementUploadTaskListV2Api service
type ToolsAppManagementUploadTaskListV2ApiService service

type ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest struct {
	ctx         context.Context
	ApiService  *ToolsAppManagementUploadTaskListV2ApiService
	accountId   *int64
	accountType *ToolsAppManagementUploadTaskListV2AccountType
	filtering   *ToolsAppManagementUploadTaskListV2Filtering
	version     string
}

// 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) AccountId(accountId int64) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型
func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) AccountType(accountType ToolsAppManagementUploadTaskListV2AccountType) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	r.accountType = &accountType
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) Filtering(filtering ToolsAppManagementUploadTaskListV2Filtering) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) Execute() (*ToolsAppManagementUploadTaskListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementUploadTaskListGet Method for OpenApi2ToolsAppManagementUploadTaskListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param version request version
	@return ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest
*/
func (a *ToolsAppManagementUploadTaskListV2ApiService) Get(ctx context.Context, version string) *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest {
	return &ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest{
		ApiService: a,
		ctx:        ctx,
		version:    version,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementUploadTaskListV2Response
func (a *ToolsAppManagementUploadTaskListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementUploadTaskListGetRequest) (*ToolsAppManagementUploadTaskListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementUploadTaskListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/upload_task/list/"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if *r.accountId < 1 {
		return localVarReturnValue, nil, ReportError("accountId must be greater than 1")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
