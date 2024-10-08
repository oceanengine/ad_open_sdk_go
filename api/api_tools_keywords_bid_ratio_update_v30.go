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

// ToolsKeywordsBidRatioUpdateV30ApiService ToolsKeywordsBidRatioUpdateV30Api service
type ToolsKeywordsBidRatioUpdateV30ApiService service

type ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest struct {
	ctx                                   context.Context
	ApiService                            *ToolsKeywordsBidRatioUpdateV30ApiService
	toolsKeywordsBidRatioUpdateV30Request *ToolsKeywordsBidRatioUpdateV30Request
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest) ToolsKeywordsBidRatioUpdateV30Request(toolsKeywordsBidRatioUpdateV30Request ToolsKeywordsBidRatioUpdateV30Request) *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest {
	r.toolsKeywordsBidRatioUpdateV30Request = &toolsKeywordsBidRatioUpdateV30Request
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest) Execute() (*ToolsKeywordsBidRatioUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsKeywordsBidRatioUpdatePost Method for OpenApiV30ToolsKeywordsBidRatioUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest
*/
func (a *ToolsKeywordsBidRatioUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest {
	return &ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsKeywordsBidRatioUpdateV30Response
func (a *ToolsKeywordsBidRatioUpdateV30ApiService) postExecute(r *ApiOpenApiV30ToolsKeywordsBidRatioUpdatePostRequest) (*ToolsKeywordsBidRatioUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsKeywordsBidRatioUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/keywords_bid_ratio/update/"

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
	localVarPostBody = r.toolsKeywordsBidRatioUpdateV30Request
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
