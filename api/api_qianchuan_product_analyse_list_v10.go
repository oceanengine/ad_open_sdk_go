/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// QianchuanProductAnalyseListV10ApiService QianchuanProductAnalyseListV10Api service
type QianchuanProductAnalyseListV10ApiService service

type ApiOpenApiV10QianchuanProductAnalyseListGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanProductAnalyseListV10ApiService
	advertiserId *int64
	keyword      *string
	timeRange    *QianchuanProductAnalyseListV10TimeRange
	orderType    *QianchuanProductAnalyseListV10OrderType
	orderField   *string
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) Keyword(keyword string) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.keyword = &keyword
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) TimeRange(timeRange QianchuanProductAnalyseListV10TimeRange) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.timeRange = &timeRange
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) OrderType(orderType QianchuanProductAnalyseListV10OrderType) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) Page(page int32) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) Execute() (*QianchuanProductAnalyseListV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanProductAnalyseListGet Method for OpenApiV10QianchuanProductAnalyseListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanProductAnalyseListGetRequest
*/
func (a *QianchuanProductAnalyseListV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanProductAnalyseListGetRequest {
	return &ApiOpenApiV10QianchuanProductAnalyseListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanProductAnalyseListV10Response
func (a *QianchuanProductAnalyseListV10ApiService) getExecute(r *ApiOpenApiV10QianchuanProductAnalyseListGetRequest) (*QianchuanProductAnalyseListV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanProductAnalyseListV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/product/analyse/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword)
	}
	if r.timeRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_range", r.timeRange)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
