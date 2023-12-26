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

// CreativeAutoGenerateConfigV2CreateV2ApiService CreativeAutoGenerateConfigV2CreateV2Api service
type CreativeAutoGenerateConfigV2CreateV2ApiService service

type ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest struct {
	ctx                                         context.Context
	ApiService                                  *CreativeAutoGenerateConfigV2CreateV2ApiService
	creativeAutoGenerateConfigV2CreateV2Request *CreativeAutoGenerateConfigV2CreateV2Request
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest) CreativeAutoGenerateConfigV2CreateV2Request(creativeAutoGenerateConfigV2CreateV2Request CreativeAutoGenerateConfigV2CreateV2Request) *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest {
	r.creativeAutoGenerateConfigV2CreateV2Request = &creativeAutoGenerateConfigV2CreateV2Request
	return r
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest) Execute() (*CreativeAutoGenerateConfigV2CreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest) WithLog(enable bool) *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeAutoGenerateConfigV2CreatePost Method for OpenApi2CreativeAutoGenerateConfigV2CreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest
*/
func (a *CreativeAutoGenerateConfigV2CreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest {
	return &ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeAutoGenerateConfigV2CreateV2Response
func (a *CreativeAutoGenerateConfigV2CreateV2ApiService) postExecute(r *ApiOpenApi2CreativeAutoGenerateConfigV2CreatePostRequest) (*CreativeAutoGenerateConfigV2CreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *CreativeAutoGenerateConfigV2CreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/auto_generate_config/v2/create/"

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
	localVarPostBody = r.creativeAutoGenerateConfigV2CreateV2Request
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
