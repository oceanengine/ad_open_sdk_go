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

// ToolsAbTestUpdateV2ApiService ToolsAbTestUpdateV2Api service
type ToolsAbTestUpdateV2ApiService service

type ApiOpenApi2ToolsAbTestUpdatePostRequest struct {
	ctx                        context.Context
	ApiService                 *ToolsAbTestUpdateV2ApiService
	version                    string
	toolsAbTestUpdateV2Request *ToolsAbTestUpdateV2Request
}

func (r *ApiOpenApi2ToolsAbTestUpdatePostRequest) ToolsAbTestUpdateV2Request(toolsAbTestUpdateV2Request ToolsAbTestUpdateV2Request) *ApiOpenApi2ToolsAbTestUpdatePostRequest {
	r.toolsAbTestUpdateV2Request = &toolsAbTestUpdateV2Request
	return r
}

func (r *ApiOpenApi2ToolsAbTestUpdatePostRequest) Execute() (*ToolsAbTestUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAbTestUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAbTestUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAbTestUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAbTestUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAbTestUpdatePost Method for OpenApi2ToolsAbTestUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param version request version
	@return ApiOpenApi2ToolsAbTestUpdatePostRequest
*/
func (a *ToolsAbTestUpdateV2ApiService) Post(ctx context.Context, version string) *ApiOpenApi2ToolsAbTestUpdatePostRequest {
	return &ApiOpenApi2ToolsAbTestUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
		version:    version,
	}
}

// Execute executes the request
//
//	@return ToolsAbTestUpdateV2Response
func (a *ToolsAbTestUpdateV2ApiService) postExecute(r *ApiOpenApi2ToolsAbTestUpdatePostRequest) (*ToolsAbTestUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAbTestUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ab_test/update/"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)

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
	localVarPostBody = r.toolsAbTestUpdateV2Request
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
