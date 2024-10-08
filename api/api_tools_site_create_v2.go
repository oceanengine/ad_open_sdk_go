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

// ToolsSiteCreateV2ApiService ToolsSiteCreateV2Api service
type ToolsSiteCreateV2ApiService service

type ApiOpenApi2ToolsSiteCreatePostRequest struct {
	ctx                      context.Context
	ApiService               *ToolsSiteCreateV2ApiService
	xOrangeCaller            *string
	toolsSiteCreateV2Request *ToolsSiteCreateV2Request
}

func (r *ApiOpenApi2ToolsSiteCreatePostRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSiteCreatePostRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsSiteCreatePostRequest) ToolsSiteCreateV2Request(toolsSiteCreateV2Request ToolsSiteCreateV2Request) *ApiOpenApi2ToolsSiteCreatePostRequest {
	r.toolsSiteCreateV2Request = &toolsSiteCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsSiteCreatePostRequest) Execute() (*ToolsSiteCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsSiteCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteCreatePost Method for OpenApi2ToolsSiteCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteCreatePostRequest
*/
func (a *ToolsSiteCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsSiteCreatePostRequest {
	return &ApiOpenApi2ToolsSiteCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteCreateV2Response
func (a *ToolsSiteCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsSiteCreatePostRequest) (*ToolsSiteCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/create/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	if r.xOrangeCaller != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-orange-caller", r.xOrangeCaller)
	}
	// body params
	localVarPostBody = r.toolsSiteCreateV2Request
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
