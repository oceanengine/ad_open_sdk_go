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

// ToolsUnionFlowPackageDeleteV2ApiService ToolsUnionFlowPackageDeleteV2Api service
type ToolsUnionFlowPackageDeleteV2ApiService service

type ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest struct {
	ctx                                  context.Context
	ApiService                           *ToolsUnionFlowPackageDeleteV2ApiService
	toolsUnionFlowPackageDeleteV2Request *ToolsUnionFlowPackageDeleteV2Request
}

func (r *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest) ToolsUnionFlowPackageDeleteV2Request(toolsUnionFlowPackageDeleteV2Request ToolsUnionFlowPackageDeleteV2Request) *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest {
	r.toolsUnionFlowPackageDeleteV2Request = &toolsUnionFlowPackageDeleteV2Request
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest) Execute() (*ToolsUnionFlowPackageDeleteV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsUnionFlowPackageDeletePost Method for OpenApi2ToolsUnionFlowPackageDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest
*/
func (a *ToolsUnionFlowPackageDeleteV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest {
	return &ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsUnionFlowPackageDeleteV2Response
func (a *ToolsUnionFlowPackageDeleteV2ApiService) postExecute(r *ApiOpenApi2ToolsUnionFlowPackageDeletePostRequest) (*ToolsUnionFlowPackageDeleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsUnionFlowPackageDeleteV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/union/flow_package/delete/"

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
	localVarPostBody = r.toolsUnionFlowPackageDeleteV2Request
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
