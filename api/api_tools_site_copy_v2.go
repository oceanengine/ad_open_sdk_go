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

// ToolsSiteCopyV2ApiService ToolsSiteCopyV2Api service
type ToolsSiteCopyV2ApiService service

type ApiOpenApi2ToolsSiteCopyPostRequest struct {
	ctx                    context.Context
	ApiService             *ToolsSiteCopyV2ApiService
	xOrangeCaller          *string
	toolsSiteCopyV2Request *ToolsSiteCopyV2Request
}

func (r *ApiOpenApi2ToolsSiteCopyPostRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSiteCopyPostRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsSiteCopyPostRequest) ToolsSiteCopyV2Request(toolsSiteCopyV2Request ToolsSiteCopyV2Request) *ApiOpenApi2ToolsSiteCopyPostRequest {
	r.toolsSiteCopyV2Request = &toolsSiteCopyV2Request
	return r
}

func (r *ApiOpenApi2ToolsSiteCopyPostRequest) Execute() (*ToolsSiteCopyV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsSiteCopyPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteCopyPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteCopyPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteCopyPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteCopyPost Method for OpenApi2ToolsSiteCopyPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteCopyPostRequest
*/
func (a *ToolsSiteCopyV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsSiteCopyPostRequest {
	return &ApiOpenApi2ToolsSiteCopyPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteCopyV2Response
func (a *ToolsSiteCopyV2ApiService) postExecute(r *ApiOpenApi2ToolsSiteCopyPostRequest) (*ToolsSiteCopyV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteCopyV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/copy/"

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
	localVarPostBody = r.toolsSiteCopyV2Request
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
