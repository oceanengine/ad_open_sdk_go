/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// QianchuanAwemeReportOrderGetV10ApiService QianchuanAwemeReportOrderGetV10Api service
type QianchuanAwemeReportOrderGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeReportOrderGetV10ApiService
	advertiserId *int64
	startDate    *string
	endDate      *string
	filtering    *QianchuanAwemeReportOrderGetV10Filtering
	fields       *[]string
	orderField   *string
	orderType    *QianchuanAwemeReportOrderGetV10OrderType
	page         *int64
	pageSize     *QianchuanAwemeReportOrderGetV10PageSize
}

func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 开始时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式 2021-04-05，时间跨度不能超过180天
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) Filtering(filtering QianchuanAwemeReportOrderGetV10Filtering) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.filtering = &filtering
	return r
}

// 需要查询的消耗指标。不传递默认查询stat_cost
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.fields = &fields
	return r
}

// 排序字段，允许值参考数据指标，默认不传为stat_cost
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式，允许值： ASC 升序（默认）、DESC 降序
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) OrderType(orderType QianchuanAwemeReportOrderGetV10OrderType) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.orderType = &orderType
	return r
}

// 页码，默认值：1
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，允许值：10, 20, 50, 100, 200，默认值：10
func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) PageSize(pageSize QianchuanAwemeReportOrderGetV10PageSize) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) Execute() (*QianchuanAwemeReportOrderGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeReportOrderGetGet Method for OpenApiV10QianchuanAwemeReportOrderGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest
*/
func (a *QianchuanAwemeReportOrderGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeReportOrderGetV10Response
func (a *QianchuanAwemeReportOrderGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeReportOrderGetGetRequest) (*QianchuanAwemeReportOrderGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeReportOrderGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/report/order/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
