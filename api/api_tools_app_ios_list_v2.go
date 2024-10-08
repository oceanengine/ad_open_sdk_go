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

// ToolsAppIosListV2ApiService ToolsAppIosListV2Api service
type ToolsAppIosListV2ApiService service

type ApiOpenApi2ToolsAppIosListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAppIosListV2ApiService
	advertiserId *int64
	filtering    *ToolsAppIosListV2Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAppIosListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) Filtering(filtering ToolsAppIosListV2Filtering) *ApiOpenApi2ToolsAppIosListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) Page(page int64) *ApiOpenApi2ToolsAppIosListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAppIosListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) Execute() (*ToolsAppIosListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppIosListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppIosListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppIosListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppIosListGet Method for OpenApi2ToolsAppIosListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppIosListGetRequest
*/
func (a *ToolsAppIosListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppIosListGetRequest {
	return &ApiOpenApi2ToolsAppIosListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppIosListV2Response
func (a *ToolsAppIosListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppIosListGetRequest) (*ToolsAppIosListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppIosListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app/ios/list/"

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
