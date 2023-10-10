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

// FileVideoMaterialClearTaskCreateV2ApiService FileVideoMaterialClearTaskCreateV2Api service
type FileVideoMaterialClearTaskCreateV2ApiService service

type ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest struct {
	ctx                                       context.Context
	ApiService                                *FileVideoMaterialClearTaskCreateV2ApiService
	fileVideoMaterialClearTaskCreateV2Request *FileVideoMaterialClearTaskCreateV2Request
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest) FileVideoMaterialClearTaskCreateV2Request(fileVideoMaterialClearTaskCreateV2Request FileVideoMaterialClearTaskCreateV2Request) *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest {
	r.fileVideoMaterialClearTaskCreateV2Request = &fileVideoMaterialClearTaskCreateV2Request
	return r
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest) Execute() (*FileVideoMaterialClearTaskCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest) WithLog(enable bool) *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoMaterialClearTaskCreatePost Method for OpenApi2FileVideoMaterialClearTaskCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest
*/
func (a *FileVideoMaterialClearTaskCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest {
	return &ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoMaterialClearTaskCreateV2Response
func (a *FileVideoMaterialClearTaskCreateV2ApiService) postExecute(r *ApiOpenApi2FileVideoMaterialClearTaskCreatePostRequest) (*FileVideoMaterialClearTaskCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *FileVideoMaterialClearTaskCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/material/clear_task/create/"

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
	localVarPostBody = r.fileVideoMaterialClearTaskCreateV2Request
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
