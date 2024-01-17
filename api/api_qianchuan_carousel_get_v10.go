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

// QianchuanCarouselGetV10ApiService QianchuanCarouselGetV10Api service
type QianchuanCarouselGetV10ApiService service

type ApiOpenApiV10QianchuanCarouselGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanCarouselGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanCarouselGetV10Filtering
	orderField   *string
	orderType    *QianchuanCarouselGetV10OrderType
	pageSize     *int64
	page         *int64
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) Filtering(filtering QianchuanCarouselGetV10Filtering) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) OrderType(orderType QianchuanCarouselGetV10OrderType) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) Execute() (*QianchuanCarouselGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanCarouselGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanCarouselGetGet Method for OpenApiV10QianchuanCarouselGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanCarouselGetGetRequest
*/
func (a *QianchuanCarouselGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanCarouselGetGetRequest {
	return &ApiOpenApiV10QianchuanCarouselGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanCarouselGetV10Response
func (a *QianchuanCarouselGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanCarouselGetGetRequest) (*QianchuanCarouselGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanCarouselGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/carousel/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
