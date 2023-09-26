/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

// QianchuanAwemeSuggestBidV10ApiService QianchuanAwemeSuggestBidV10Api service
type QianchuanAwemeSuggestBidV10ApiService service

type ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest struct {
	ctx             context.Context
	ApiService      *QianchuanAwemeSuggestBidV10ApiService
	advertiserId    *int64
	deliverySetting *QianchuanAwemeSuggestBidV10DeliverySetting
	audience        *QianchuanAwemeSuggestBidV10Audience
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 投放设置
func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) DeliverySetting(deliverySetting QianchuanAwemeSuggestBidV10DeliverySetting) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	r.deliverySetting = &deliverySetting
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) Audience(audience QianchuanAwemeSuggestBidV10Audience) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	r.audience = &audience
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) Execute() (*QianchuanAwemeSuggestBidV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeSuggestBidGet Method for OpenApiV10QianchuanAwemeSuggestBidGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest
*/
func (a *QianchuanAwemeSuggestBidV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest {
	return &ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeSuggestBidV10Response
func (a *QianchuanAwemeSuggestBidV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeSuggestBidGetRequest) (*QianchuanAwemeSuggestBidV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanAwemeSuggestBidV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/suggest_bid/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.deliverySetting == nil {
		return localVarReturnValue, nil, ReportError("deliverySetting is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_setting", r.deliverySetting)
	if r.audience != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience", r.audience)
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