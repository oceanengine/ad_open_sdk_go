/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
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

// ReportCampaignGetV2ApiService ReportCampaignGetV2Api service
type ReportCampaignGetV2ApiService service

type ApiOpenApi2ReportCampaignGetGetRequest struct {
	ctx             context.Context
	ApiService      *ReportCampaignGetV2ApiService
	advertiserId    *int64
	endDate         **string
	fields          *[]string
	filtering       *ReportCampaignGetV2Filtering
	groupBy         *[]*ReportCampaignGetV2GroupBy
	orderField      *ReportCampaignGetV2OrderField
	orderType       *ReportCampaignGetV2OrderType
	page            *int64
	pageSize        *int64
	startDate       **string
	timeGranularity *ReportCampaignGetV2TimeGranularity
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) Filtering(filtering ReportCampaignGetV2Filtering) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) GroupBy(groupBy []*ReportCampaignGetV2GroupBy) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.groupBy = &groupBy
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) OrderField(orderField ReportCampaignGetV2OrderField) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) OrderType(orderType ReportCampaignGetV2OrderType) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) Page(page int64) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) TimeGranularity(timeGranularity ReportCampaignGetV2TimeGranularity) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) Execute() (*ReportCampaignGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportCampaignGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportCampaignGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportCampaignGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportCampaignGetGet Method for OpenApi2ReportCampaignGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportCampaignGetGetRequest
*/
func (a *ReportCampaignGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportCampaignGetGetRequest {
	return &ApiOpenApi2ReportCampaignGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportCampaignGetV2Response
func (a *ReportCampaignGetV2ApiService) getExecute(r *ApiOpenApi2ReportCampaignGetGetRequest) (*ReportCampaignGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportCampaignGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/campaign/get/"

	localVarHeaderParams := make(map[string]string)
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
