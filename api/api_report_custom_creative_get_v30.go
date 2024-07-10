/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// ReportCustomCreativeGetV30ApiService ReportCustomCreativeGetV30Api service
type ReportCustomCreativeGetV30ApiService service

type ApiOpenApiV30ReportCustomCreativeGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportCustomCreativeGetV30ApiService
	dimensions   *[]string
	advertiserId *int64
	metrics      *[]string
	filters      *[]*ReportCustomCreativeGetV30FiltersInner
	startTime    *string
	endTime      *string
	orderBy      *[]*ReportCustomCreativeGetV30OrderByInner
	page         *int32
	pageSize     *int32
}

// 维度列表
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) Dimensions(dimensions []string) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.dimensions = &dimensions
	return r
}

// 广告主id
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 指标列表
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) Metrics(metrics []string) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.metrics = &metrics
	return r
}

// 过滤条件
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) Filters(filters []*ReportCustomCreativeGetV30FiltersInner) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.filters = &filters
	return r
}

// 开始时间。单位毫秒
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) StartTime(startTime string) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间。单位毫秒
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) EndTime(endTime string) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.endTime = &endTime
	return r
}

// 排序
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) OrderBy(orderBy []*ReportCustomCreativeGetV30OrderByInner) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.orderBy = &orderBy
	return r
}

// 页码。默认为1
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) Page(page int32) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.page = &page
	return r
}

// 页面大小。默认为10，最大100
func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) Execute() (*ReportCustomCreativeGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportCustomCreativeGetGet Method for OpenApiV30ReportCustomCreativeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportCustomCreativeGetGetRequest
*/
func (a *ReportCustomCreativeGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportCustomCreativeGetGetRequest {
	return &ApiOpenApiV30ReportCustomCreativeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportCustomCreativeGetV30Response
func (a *ReportCustomCreativeGetV30ApiService) getExecute(r *ApiOpenApiV30ReportCustomCreativeGetGetRequest) (*ReportCustomCreativeGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportCustomCreativeGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/custom/creative/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dimensions == nil {
		return localVarReturnValue, nil, ReportError("dimensions is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, ReportError("metrics is required and must be specified")
	}
	if r.filters == nil {
		return localVarReturnValue, nil, ReportError("filters is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.orderBy == nil {
		return localVarReturnValue, nil, ReportError("orderBy is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "dimensions", r.dimensions)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy)
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
