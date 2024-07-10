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

// QianchuanReportAdGetV10ApiService QianchuanReportAdGetV10Api service
type QianchuanReportAdGetV10ApiService service

type ApiOpenApiV10QianchuanReportAdGetGetRequest struct {
	ctx             context.Context
	ApiService      *QianchuanReportAdGetV10ApiService
	advertiserId    *int64
	startDate       *string
	endDate         *string
	fields          *[]string
	filtering       *QianchuanReportAdGetV10Filtering
	timeGranularity *QianchuanReportAdGetV10TimeGranularity
	orderField      *string
	orderType       *QianchuanReportAdGetV10OrderType
	page            *int32
	pageSize        *int32
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 开始时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) Filtering(filtering QianchuanReportAdGetV10Filtering) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) TimeGranularity(timeGranularity QianchuanReportAdGetV10TimeGranularity) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) OrderType(orderType QianchuanReportAdGetV10OrderType) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) Execute() (*QianchuanReportAdGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportAdGetGet Method for OpenApiV10QianchuanReportAdGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportAdGetGetRequest
*/
func (a *QianchuanReportAdGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportAdGetGetRequest {
	return &ApiOpenApiV10QianchuanReportAdGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportAdGetV10Response
func (a *QianchuanReportAdGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportAdGetGetRequest) (*QianchuanReportAdGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanReportAdGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/ad/get/"

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
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	if r.timeGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_granularity", r.timeGranularity)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
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
