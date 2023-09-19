/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
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

// QianchuanReportSearchWordGetV10ApiService QianchuanReportSearchWordGetV10Api service
type QianchuanReportSearchWordGetV10ApiService service

type ApiOpenApiV10QianchuanReportSearchWordGetGetRequest struct {
	ctx             context.Context
	ApiService      *QianchuanReportSearchWordGetV10ApiService
	advertiserId    *int64
	startDate       *string
	endDate         *string
	fields          *[]string
	wordType        *QianchuanReportSearchWordGetV10WordType
	filtering       *QianchuanReportSearchWordGetV10Filtering
	timeGranularity *QianchuanReportSearchWordGetV10TimeGranularity
	orderField      *string
	orderType       *QianchuanReportSearchWordGetV10OrderType
	page            *int32
	pageSize        *int32
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 开始时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) WordType(wordType QianchuanReportSearchWordGetV10WordType) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.wordType = &wordType
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) Filtering(filtering QianchuanReportSearchWordGetV10Filtering) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) TimeGranularity(timeGranularity QianchuanReportSearchWordGetV10TimeGranularity) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) OrderType(orderType QianchuanReportSearchWordGetV10OrderType) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) Execute() (*QianchuanReportSearchWordGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportSearchWordGetGet Method for OpenApiV10QianchuanReportSearchWordGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportSearchWordGetGetRequest
*/
func (a *QianchuanReportSearchWordGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest {
	return &ApiOpenApiV10QianchuanReportSearchWordGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportSearchWordGetV10Response
func (a *QianchuanReportSearchWordGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportSearchWordGetGetRequest) (*QianchuanReportSearchWordGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanReportSearchWordGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/search_word/get/"

	localVarHeaderParams := make(map[string]string)
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
	if r.wordType == nil {
		return localVarReturnValue, nil, ReportError("wordType is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.timeGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_granularity", r.timeGranularity)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	parameterAddToHeaderOrQuery(localVarQueryParams, "word_type", r.wordType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
