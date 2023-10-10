/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
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

// ToolsAwemeBannedDeleteV30ApiService ToolsAwemeBannedDeleteV30Api service
type ToolsAwemeBannedDeleteV30ApiService service

type ApiOpenApiV30ToolsAwemeBannedDeletePostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsAwemeBannedDeleteV30ApiService
	toolsAwemeBannedDeleteV30Request *ToolsAwemeBannedDeleteV30Request
}

func (r *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest) ToolsAwemeBannedDeleteV30Request(toolsAwemeBannedDeleteV30Request ToolsAwemeBannedDeleteV30Request) *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest {
	r.toolsAwemeBannedDeleteV30Request = &toolsAwemeBannedDeleteV30Request
	return r
}

func (r *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest) Execute() (*ToolsAwemeBannedDeleteV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsAwemeBannedDeletePost Method for OpenApiV30ToolsAwemeBannedDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsAwemeBannedDeletePostRequest
*/
func (a *ToolsAwemeBannedDeleteV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest {
	return &ApiOpenApiV30ToolsAwemeBannedDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeBannedDeleteV30Response
func (a *ToolsAwemeBannedDeleteV30ApiService) postExecute(r *ApiOpenApiV30ToolsAwemeBannedDeletePostRequest) (*ToolsAwemeBannedDeleteV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAwemeBannedDeleteV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/aweme_banned/delete/"

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
	localVarPostBody = r.toolsAwemeBannedDeleteV30Request
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
