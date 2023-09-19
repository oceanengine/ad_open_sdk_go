/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
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

// ToolsAppManagementExtendPackageCreateV2ApiService ToolsAppManagementExtendPackageCreateV2Api service
type ToolsAppManagementExtendPackageCreateV2ApiService service

type ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest struct {
	ctx                                            context.Context
	ApiService                                     *ToolsAppManagementExtendPackageCreateV2ApiService
	toolsAppManagementExtendPackageCreateV2Request *ToolsAppManagementExtendPackageCreateV2Request
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest) ToolsAppManagementExtendPackageCreateV2Request(toolsAppManagementExtendPackageCreateV2Request ToolsAppManagementExtendPackageCreateV2Request) *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest {
	r.toolsAppManagementExtendPackageCreateV2Request = &toolsAppManagementExtendPackageCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest) Execute() (*ToolsAppManagementExtendPackageCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementExtendPackageCreatePost Method for OpenApi2ToolsAppManagementExtendPackageCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest
*/
func (a *ToolsAppManagementExtendPackageCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest {
	return &ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementExtendPackageCreateV2Response
func (a *ToolsAppManagementExtendPackageCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsAppManagementExtendPackageCreatePostRequest) (*ToolsAppManagementExtendPackageCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAppManagementExtendPackageCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/extend_package/create/"

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
	localVarPostBody = r.toolsAppManagementExtendPackageCreateV2Request
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
