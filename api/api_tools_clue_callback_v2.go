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

// ToolsClueCallbackV2ApiService ToolsClueCallbackV2Api service
type ToolsClueCallbackV2ApiService service

type ApiOpenApi2ToolsClueCallbackPostRequest struct {
	ctx                        context.Context
	ApiService                 *ToolsClueCallbackV2ApiService
	toolsClueCallbackV2Request *ToolsClueCallbackV2Request
}

func (r *ApiOpenApi2ToolsClueCallbackPostRequest) ToolsClueCallbackV2Request(toolsClueCallbackV2Request ToolsClueCallbackV2Request) *ApiOpenApi2ToolsClueCallbackPostRequest {
	r.toolsClueCallbackV2Request = &toolsClueCallbackV2Request
	return r
}

func (r *ApiOpenApi2ToolsClueCallbackPostRequest) Execute() (*ToolsClueCallbackV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsClueCallbackPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueCallbackPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueCallbackPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueCallbackPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueCallbackPost Method for OpenApi2ToolsClueCallbackPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueCallbackPostRequest
*/
func (a *ToolsClueCallbackV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsClueCallbackPostRequest {
	return &ApiOpenApi2ToolsClueCallbackPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueCallbackV2Response
func (a *ToolsClueCallbackV2ApiService) postExecute(r *ApiOpenApi2ToolsClueCallbackPostRequest) (*ToolsClueCallbackV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueCallbackV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/callback/"

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

	// body params
	localVarPostBody = r.toolsClueCallbackV2Request
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
