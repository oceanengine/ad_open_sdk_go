/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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

// ToolsAppManagementUploadTaskCreateV2ApiService ToolsAppManagementUploadTaskCreateV2Api service
type ToolsAppManagementUploadTaskCreateV2ApiService service

type ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest struct {
	ctx                                         context.Context
	ApiService                                  *ToolsAppManagementUploadTaskCreateV2ApiService
	toolsAppManagementUploadTaskCreateV2Request *ToolsAppManagementUploadTaskCreateV2Request
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest) ToolsAppManagementUploadTaskCreateV2Request(toolsAppManagementUploadTaskCreateV2Request ToolsAppManagementUploadTaskCreateV2Request) *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest {
	r.toolsAppManagementUploadTaskCreateV2Request = &toolsAppManagementUploadTaskCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest) Execute() (*ToolsAppManagementUploadTaskCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementUploadTaskCreatePost Method for OpenApi2ToolsAppManagementUploadTaskCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest
*/
func (a *ToolsAppManagementUploadTaskCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest {
	return &ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementUploadTaskCreateV2Response
func (a *ToolsAppManagementUploadTaskCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsAppManagementUploadTaskCreatePostRequest) (*ToolsAppManagementUploadTaskCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAppManagementUploadTaskCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/upload_task/create/"

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
	localVarPostBody = r.toolsAppManagementUploadTaskCreateV2Request
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
