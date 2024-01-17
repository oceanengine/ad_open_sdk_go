/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

// ToolsAppManagementAndroidBasicPackagePublishV2ApiService ToolsAppManagementAndroidBasicPackagePublishV2Api service
type ToolsAppManagementAndroidBasicPackagePublishV2ApiService service

type ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest struct {
	ctx                                                   context.Context
	ApiService                                            *ToolsAppManagementAndroidBasicPackagePublishV2ApiService
	toolsAppManagementAndroidBasicPackagePublishV2Request *ToolsAppManagementAndroidBasicPackagePublishV2Request
}

func (r *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest) ToolsAppManagementAndroidBasicPackagePublishV2Request(toolsAppManagementAndroidBasicPackagePublishV2Request ToolsAppManagementAndroidBasicPackagePublishV2Request) *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest {
	r.toolsAppManagementAndroidBasicPackagePublishV2Request = &toolsAppManagementAndroidBasicPackagePublishV2Request
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest) Execute() (*ToolsAppManagementAndroidBasicPackagePublishV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementAndroidBasicPackagePublishPost Method for OpenApi2ToolsAppManagementAndroidBasicPackagePublishPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest
*/
func (a *ToolsAppManagementAndroidBasicPackagePublishV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest {
	return &ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementAndroidBasicPackagePublishV2Response
func (a *ToolsAppManagementAndroidBasicPackagePublishV2ApiService) postExecute(r *ApiOpenApi2ToolsAppManagementAndroidBasicPackagePublishPostRequest) (*ToolsAppManagementAndroidBasicPackagePublishV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAppManagementAndroidBasicPackagePublishV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/android_basic_package/publish/"

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
	localVarPostBody = r.toolsAppManagementAndroidBasicPackagePublishV2Request
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
