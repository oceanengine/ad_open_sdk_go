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

// ToolsLandingGroupCreateV2ApiService ToolsLandingGroupCreateV2Api service
type ToolsLandingGroupCreateV2ApiService service

type ApiOpenApi2ToolsLandingGroupCreatePostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsLandingGroupCreateV2ApiService
	xOrangeCaller                    *string
	toolsLandingGroupCreateV2Request *ToolsLandingGroupCreateV2Request
}

func (r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsLandingGroupCreatePostRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) ToolsLandingGroupCreateV2Request(toolsLandingGroupCreateV2Request ToolsLandingGroupCreateV2Request) *ApiOpenApi2ToolsLandingGroupCreatePostRequest {
	r.toolsLandingGroupCreateV2Request = &toolsLandingGroupCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) Execute() (*ToolsLandingGroupCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsLandingGroupCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsLandingGroupCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsLandingGroupCreatePost Method for OpenApi2ToolsLandingGroupCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsLandingGroupCreatePostRequest
*/
func (a *ToolsLandingGroupCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsLandingGroupCreatePostRequest {
	return &ApiOpenApi2ToolsLandingGroupCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsLandingGroupCreateV2Response
func (a *ToolsLandingGroupCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsLandingGroupCreatePostRequest) (*ToolsLandingGroupCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsLandingGroupCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/landing_group/create/"

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
	localVarPostBody = r.toolsLandingGroupCreateV2Request
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
