/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// ToolsClueLiteContactGetV2ApiService ToolsClueLiteContactGetV2Api service
type ToolsClueLiteContactGetV2ApiService service

type ApiOpenApi2ToolsClueLiteContactGetPostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsClueLiteContactGetV2ApiService
	toolsClueLiteContactGetV2Request *ToolsClueLiteContactGetV2Request
}

func (r *ApiOpenApi2ToolsClueLiteContactGetPostRequest) ToolsClueLiteContactGetV2Request(toolsClueLiteContactGetV2Request ToolsClueLiteContactGetV2Request) *ApiOpenApi2ToolsClueLiteContactGetPostRequest {
	r.toolsClueLiteContactGetV2Request = &toolsClueLiteContactGetV2Request
	return r
}

func (r *ApiOpenApi2ToolsClueLiteContactGetPostRequest) Execute() (*ToolsClueLiteContactGetV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsClueLiteContactGetPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueLiteContactGetPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueLiteContactGetPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueLiteContactGetPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueLiteContactGetPost Method for OpenApi2ToolsClueLiteContactGetPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueLiteContactGetPostRequest
*/
func (a *ToolsClueLiteContactGetV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsClueLiteContactGetPostRequest {
	return &ApiOpenApi2ToolsClueLiteContactGetPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueLiteContactGetV2Response
func (a *ToolsClueLiteContactGetV2ApiService) postExecute(r *ApiOpenApi2ToolsClueLiteContactGetPostRequest) (*ToolsClueLiteContactGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueLiteContactGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/lite/contact/get/"

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolsClueLiteContactGetV2Request
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
