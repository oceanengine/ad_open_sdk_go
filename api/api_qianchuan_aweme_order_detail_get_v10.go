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

// QianchuanAwemeOrderDetailGetV10ApiService QianchuanAwemeOrderDetailGetV10Api service
type QianchuanAwemeOrderDetailGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeOrderDetailGetV10ApiService
	orderId      *int64
	advertiserId *int64
}

// 千川小店订单id
func (r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) OrderId(orderId int64) *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest {
	r.orderId = &orderId
	return r
}

// 千川广告主id
func (r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) Execute() (*QianchuanAwemeOrderDetailGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeOrderDetailGetGet Method for OpenApiV10QianchuanAwemeOrderDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest
*/
func (a *QianchuanAwemeOrderDetailGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeOrderDetailGetV10Response
func (a *QianchuanAwemeOrderDetailGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeOrderDetailGetGetRequest) (*QianchuanAwemeOrderDetailGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanAwemeOrderDetailGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/order/detail/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderId == nil {
		return localVarReturnValue, nil, ReportError("orderId is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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