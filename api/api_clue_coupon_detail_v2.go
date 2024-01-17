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

// ClueCouponDetailV2ApiService ClueCouponDetailV2Api service
type ClueCouponDetailV2ApiService service

type ApiOpenApi2ClueCouponDetailGetRequest struct {
	ctx          context.Context
	ApiService   *ClueCouponDetailV2ApiService
	activityId   *int64
	advertiserId *int64
}

func (r *ApiOpenApi2ClueCouponDetailGetRequest) ActivityId(activityId int64) *ApiOpenApi2ClueCouponDetailGetRequest {
	r.activityId = &activityId
	return r
}

func (r *ApiOpenApi2ClueCouponDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueCouponDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ClueCouponDetailGetRequest) Execute() (*ClueCouponDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueCouponDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueCouponDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueCouponDetailGetRequest) WithLog(enable bool) *ApiOpenApi2ClueCouponDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueCouponDetailGet Method for OpenApi2ClueCouponDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueCouponDetailGetRequest
*/
func (a *ClueCouponDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueCouponDetailGetRequest {
	return &ApiOpenApi2ClueCouponDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueCouponDetailV2Response
func (a *ClueCouponDetailV2ApiService) getExecute(r *ApiOpenApi2ClueCouponDetailGetRequest) (*ClueCouponDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ClueCouponDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/coupon/detail/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.activityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activity_id", r.activityId)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
