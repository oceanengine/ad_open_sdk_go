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

// ToolsSiteUpdateStatusV2ApiService ToolsSiteUpdateStatusV2Api service
type ToolsSiteUpdateStatusV2ApiService service

type ApiOpenApi2ToolsSiteUpdateStatusPostRequest struct {
	ctx                            context.Context
	ApiService                     *ToolsSiteUpdateStatusV2ApiService
	xOrangeCaller                  *string
	toolsSiteUpdateStatusV2Request *ToolsSiteUpdateStatusV2Request
}

func (r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSiteUpdateStatusPostRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) ToolsSiteUpdateStatusV2Request(toolsSiteUpdateStatusV2Request ToolsSiteUpdateStatusV2Request) *ApiOpenApi2ToolsSiteUpdateStatusPostRequest {
	r.toolsSiteUpdateStatusV2Request = &toolsSiteUpdateStatusV2Request
	return r
}

func (r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) Execute() (*ToolsSiteUpdateStatusV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteUpdateStatusPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteUpdateStatusPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteUpdateStatusPost Method for OpenApi2ToolsSiteUpdateStatusPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteUpdateStatusPostRequest
*/
func (a *ToolsSiteUpdateStatusV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsSiteUpdateStatusPostRequest {
	return &ApiOpenApi2ToolsSiteUpdateStatusPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteUpdateStatusV2Response
func (a *ToolsSiteUpdateStatusV2ApiService) postExecute(r *ApiOpenApi2ToolsSiteUpdateStatusPostRequest) (*ToolsSiteUpdateStatusV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteUpdateStatusV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/update_status/"

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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Orange-Caller", r.xOrangeCaller)
	}
	// body params
	localVarPostBody = r.toolsSiteUpdateStatusV2Request
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
