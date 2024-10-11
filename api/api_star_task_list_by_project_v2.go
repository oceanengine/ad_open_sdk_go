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

// StarTaskListByProjectV2ApiService StarTaskListByProjectV2Api service
type StarTaskListByProjectV2ApiService service

type ApiOpenApi2StarTaskListByProjectGetRequest struct {
	ctx        context.Context
	ApiService *StarTaskListByProjectV2ApiService
	starId     *int64
	projectId  *int64
	page       *int64
	limit      *int64
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) StarId(starId int64) *ApiOpenApi2StarTaskListByProjectGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) ProjectId(projectId int64) *ApiOpenApi2StarTaskListByProjectGetRequest {
	r.projectId = &projectId
	return r
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) Page(page int64) *ApiOpenApi2StarTaskListByProjectGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) Limit(limit int64) *ApiOpenApi2StarTaskListByProjectGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) Execute() (*StarTaskListByProjectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarTaskListByProjectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarTaskListByProjectGetRequest) WithLog(enable bool) *ApiOpenApi2StarTaskListByProjectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarTaskListByProjectGet Method for OpenApi2StarTaskListByProjectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarTaskListByProjectGetRequest
*/
func (a *StarTaskListByProjectV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarTaskListByProjectGetRequest {
	return &ApiOpenApi2StarTaskListByProjectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarTaskListByProjectV2Response
func (a *StarTaskListByProjectV2ApiService) getExecute(r *ApiOpenApi2StarTaskListByProjectGetRequest) (*StarTaskListByProjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarTaskListByProjectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/task/list_by_project/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, ReportError("projectId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if *r.page > 10000 {
		return localVarReturnValue, nil, ReportError("page must be less than 10000")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, ReportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, ReportError("limit must be greater than 1")
	}
	if *r.limit > 50 {
		return localVarReturnValue, nil, ReportError("limit must be less than 50")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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