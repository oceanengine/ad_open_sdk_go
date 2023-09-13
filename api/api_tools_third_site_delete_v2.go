/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsThirdSiteDeleteV2ApiService ToolsThirdSiteDeleteV2Api service
type ToolsThirdSiteDeleteV2ApiService service

type ApiOpenApi2ToolsThirdSiteDeletePostRequest struct {
	ctx                           context.Context
	ApiService                    *ToolsThirdSiteDeleteV2ApiService
	toolsThirdSiteDeleteV2Request *ToolsThirdSiteDeleteV2Request
}

func (r *ApiOpenApi2ToolsThirdSiteDeletePostRequest) ToolsThirdSiteDeleteV2Request(toolsThirdSiteDeleteV2Request ToolsThirdSiteDeleteV2Request) *ApiOpenApi2ToolsThirdSiteDeletePostRequest {
	r.toolsThirdSiteDeleteV2Request = &toolsThirdSiteDeleteV2Request
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteDeletePostRequest) Execute() (*ToolsThirdSiteDeleteV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsThirdSiteDeletePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsThirdSiteDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteDeletePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsThirdSiteDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsThirdSiteDeletePost Method for OpenApi2ToolsThirdSiteDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsThirdSiteDeletePostRequest
*/
func (a *ToolsThirdSiteDeleteV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsThirdSiteDeletePostRequest {
	return &ApiOpenApi2ToolsThirdSiteDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsThirdSiteDeleteV2Response
func (a *ToolsThirdSiteDeleteV2ApiService) postExecute(r *ApiOpenApi2ToolsThirdSiteDeletePostRequest) (*ToolsThirdSiteDeleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsThirdSiteDeleteV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/third_site/delete/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolsThirdSiteDeleteV2Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
