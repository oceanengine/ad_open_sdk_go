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

// ToolsThirdSiteCreateV2ApiService ToolsThirdSiteCreateV2Api service
type ToolsThirdSiteCreateV2ApiService service

type ApiOpenApi2ToolsThirdSiteCreatePostRequest struct {
	ctx                           context.Context
	ApiService                    *ToolsThirdSiteCreateV2ApiService
	toolsThirdSiteCreateV2Request *ToolsThirdSiteCreateV2Request
}

func (r *ApiOpenApi2ToolsThirdSiteCreatePostRequest) ToolsThirdSiteCreateV2Request(toolsThirdSiteCreateV2Request ToolsThirdSiteCreateV2Request) *ApiOpenApi2ToolsThirdSiteCreatePostRequest {
	r.toolsThirdSiteCreateV2Request = &toolsThirdSiteCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteCreatePostRequest) Execute() (*ToolsThirdSiteCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsThirdSiteCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsThirdSiteCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsThirdSiteCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsThirdSiteCreatePost Method for OpenApi2ToolsThirdSiteCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsThirdSiteCreatePostRequest
*/
func (a *ToolsThirdSiteCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsThirdSiteCreatePostRequest {
	return &ApiOpenApi2ToolsThirdSiteCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsThirdSiteCreateV2Response
func (a *ToolsThirdSiteCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsThirdSiteCreatePostRequest) (*ToolsThirdSiteCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsThirdSiteCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/third_site/create/"

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
	localVarPostBody = r.toolsThirdSiteCreateV2Request
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
