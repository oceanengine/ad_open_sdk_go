/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// CreativeRejectReasonV2ApiService CreativeRejectReasonV2Api service
type CreativeRejectReasonV2ApiService service

type ApiOpenApi2CreativeRejectReasonGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeRejectReasonV2ApiService
	advertiserId *int64
	creativeIds  *[]int64
}

func (r *ApiOpenApi2CreativeRejectReasonGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeRejectReasonGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2CreativeRejectReasonGetRequest) CreativeIds(creativeIds []int64) *ApiOpenApi2CreativeRejectReasonGetRequest {
	r.creativeIds = &creativeIds
	return r
}

func (r *ApiOpenApi2CreativeRejectReasonGetRequest) Execute() (*CreativeRejectReasonV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeRejectReasonGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeRejectReasonGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeRejectReasonGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeRejectReasonGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeRejectReasonGet Method for OpenApi2CreativeRejectReasonGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeRejectReasonGetRequest
*/
func (a *CreativeRejectReasonV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeRejectReasonGetRequest {
	return &ApiOpenApi2CreativeRejectReasonGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeRejectReasonV2Response
func (a *CreativeRejectReasonV2ApiService) getExecute(r *ApiOpenApi2CreativeRejectReasonGetRequest) (*CreativeRejectReasonV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeRejectReasonV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/reject_reason/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.creativeIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creative_ids", r.creativeIds)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
