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

// ToolsRubeexGetV2ApiService ToolsRubeexGetV2Api service
type ToolsRubeexGetV2ApiService service

type ApiOpenApi2ToolsRubeexGetGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsRubeexGetV2ApiService
	advertiserId  *int64
	filtering     *ToolsRubeexGetV2Filtering
	page          *int64
	pageSize      *int64
	platformName  *[]*ToolsRubeexGetV2PlatformName
	projectStatus *[]*ToolsRubeexGetV2ProjectStatus
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) Filtering(filtering ToolsRubeexGetV2Filtering) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) Page(page int64) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) PlatformName(platformName []*ToolsRubeexGetV2PlatformName) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.platformName = &platformName
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) ProjectStatus(projectStatus []*ToolsRubeexGetV2ProjectStatus) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.projectStatus = &projectStatus
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) Execute() (*ToolsRubeexGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRubeexGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRubeexGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRubeexGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRubeexGetGet Method for OpenApi2ToolsRubeexGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRubeexGetGetRequest
*/
func (a *ToolsRubeexGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRubeexGetGetRequest {
	return &ApiOpenApi2ToolsRubeexGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRubeexGetV2Response
func (a *ToolsRubeexGetV2ApiService) getExecute(r *ApiOpenApi2ToolsRubeexGetGetRequest) (*ToolsRubeexGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsRubeexGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rubeex/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.platformName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform_name", r.platformName)
	}
	if r.projectStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "project_status", r.projectStatus)
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
