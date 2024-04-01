/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

// CreativeProceduralCreativeUpdateV2ApiService CreativeProceduralCreativeUpdateV2Api service
type CreativeProceduralCreativeUpdateV2ApiService service

type ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest struct {
	ctx                                       context.Context
	ApiService                                *CreativeProceduralCreativeUpdateV2ApiService
	creativeProceduralCreativeUpdateV2Request *CreativeProceduralCreativeUpdateV2Request
}

func (r *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest) CreativeProceduralCreativeUpdateV2Request(creativeProceduralCreativeUpdateV2Request CreativeProceduralCreativeUpdateV2Request) *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest {
	r.creativeProceduralCreativeUpdateV2Request = &creativeProceduralCreativeUpdateV2Request
	return r
}

func (r *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest) Execute() (*CreativeProceduralCreativeUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeProceduralCreativeUpdatePost Method for OpenApi2CreativeProceduralCreativeUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest
*/
func (a *CreativeProceduralCreativeUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest {
	return &ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeProceduralCreativeUpdateV2Response
func (a *CreativeProceduralCreativeUpdateV2ApiService) postExecute(r *ApiOpenApi2CreativeProceduralCreativeUpdatePostRequest) (*CreativeProceduralCreativeUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeProceduralCreativeUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/procedural_creative/update/"

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
	localVarPostBody = r.creativeProceduralCreativeUpdateV2Request
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
