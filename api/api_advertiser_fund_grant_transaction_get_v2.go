/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// AdvertiserFundGrantTransactionGetV2ApiService AdvertiserFundGrantTransactionGetV2Api service
type AdvertiserFundGrantTransactionGetV2ApiService service

type ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest struct {
	ctx             context.Context
	ApiService      *AdvertiserFundGrantTransactionGetV2ApiService
	advertiserId    *int64
	startTime       *string
	endTime         *string
	page            *string
	pageSize        *string
	transactionType *string
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) StartTime(startTime string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) EndTime(endTime string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) Page(page string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) PageSize(pageSize string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) TransactionType(transactionType string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.transactionType = &transactionType
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) Execute() (*AdvertiserFundGrantTransactionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) WithLog(enable bool) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdvertiserFundGrantTransactionGetGet Method for OpenApi2AdvertiserFundGrantTransactionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest
*/
func (a *AdvertiserFundGrantTransactionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest {
	return &ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserFundGrantTransactionGetV2Response
func (a *AdvertiserFundGrantTransactionGetV2ApiService) getExecute(r *ApiOpenApi2AdvertiserFundGrantTransactionGetGetRequest) (*AdvertiserFundGrantTransactionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserFundGrantTransactionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/advertiser/fund/grant_transaction/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	if r.transactionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transaction_type", r.transactionType)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
