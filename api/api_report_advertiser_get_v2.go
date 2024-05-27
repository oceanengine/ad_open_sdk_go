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

// ReportAdvertiserGetV2ApiService ReportAdvertiserGetV2Api service
type ReportAdvertiserGetV2ApiService service

type ApiOpenApi2ReportAdvertiserGetGetRequest struct {
	ctx             context.Context
	ApiService      *ReportAdvertiserGetV2ApiService
	advertiserId    *int64
	endDate         **string
	fields          *[]string
	filtering       *ReportAdvertiserGetV2Filtering
	groupBy         *[]*ReportAdvertiserGetV2GroupBy
	orderField      *ReportAdvertiserGetV2OrderField
	orderType       *ReportAdvertiserGetV2OrderType
	page            *int64
	pageSize        *int64
	startDate       **string
	timeGranularity *ReportAdvertiserGetV2TimeGranularity
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) Filtering(filtering ReportAdvertiserGetV2Filtering) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) GroupBy(groupBy []*ReportAdvertiserGetV2GroupBy) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.groupBy = &groupBy
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) OrderField(orderField ReportAdvertiserGetV2OrderField) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) OrderType(orderType ReportAdvertiserGetV2OrderType) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) Page(page int64) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) TimeGranularity(timeGranularity ReportAdvertiserGetV2TimeGranularity) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) Execute() (*ReportAdvertiserGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportAdvertiserGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportAdvertiserGetGet Method for OpenApi2ReportAdvertiserGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportAdvertiserGetGetRequest
*/
func (a *ReportAdvertiserGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportAdvertiserGetGetRequest {
	return &ApiOpenApi2ReportAdvertiserGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportAdvertiserGetV2Response
func (a *ReportAdvertiserGetV2ApiService) getExecute(r *ApiOpenApi2ReportAdvertiserGetGetRequest) (*ReportAdvertiserGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportAdvertiserGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/advertiser/get/"

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
