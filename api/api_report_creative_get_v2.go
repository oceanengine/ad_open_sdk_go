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

// ReportCreativeGetV2ApiService ReportCreativeGetV2Api service
type ReportCreativeGetV2ApiService service

type ApiOpenApi2ReportCreativeGetGetRequest struct {
	ctx             context.Context
	ApiService      *ReportCreativeGetV2ApiService
	advertiserId    *int64
	endDate         **string
	fields          *[]string
	filtering       *ReportCreativeGetV2Filtering
	groupBy         *[]*ReportCreativeGetV2GroupBy
	orderField      *ReportCreativeGetV2OrderField
	orderType       *ReportCreativeGetV2OrderType
	page            *int64
	pageSize        *int64
	startDate       **string
	timeGranularity *ReportCreativeGetV2TimeGranularity
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) Filtering(filtering ReportCreativeGetV2Filtering) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) GroupBy(groupBy []*ReportCreativeGetV2GroupBy) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.groupBy = &groupBy
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) OrderField(orderField ReportCreativeGetV2OrderField) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) OrderType(orderType ReportCreativeGetV2OrderType) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) Page(page int64) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) TimeGranularity(timeGranularity ReportCreativeGetV2TimeGranularity) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) Execute() (*ReportCreativeGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportCreativeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportCreativeGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportCreativeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportCreativeGetGet Method for OpenApi2ReportCreativeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportCreativeGetGetRequest
*/
func (a *ReportCreativeGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportCreativeGetGetRequest {
	return &ApiOpenApi2ReportCreativeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportCreativeGetV2Response
func (a *ReportCreativeGetV2ApiService) getExecute(r *ApiOpenApi2ReportCreativeGetGetRequest) (*ReportCreativeGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportCreativeGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/creative/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_by", r.groupBy)
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	}
	if r.timeGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_granularity", r.timeGranularity)
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
