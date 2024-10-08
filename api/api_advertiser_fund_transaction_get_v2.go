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

// AdvertiserFundTransactionGetV2ApiService AdvertiserFundTransactionGetV2Api service
type AdvertiserFundTransactionGetV2ApiService service

type ApiOpenApi2AdvertiserFundTransactionGetGetRequest struct {
	ctx             context.Context
	ApiService      *AdvertiserFundTransactionGetV2ApiService
	advertiserId    *int64
	startDate       *string
	endDate         *string
	transactionType *AdvertiserFundTransactionGetV2TransactionType
	page            *int32
	pageSize        *int32
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) StartDate(startDate string) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) EndDate(endDate string) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) TransactionType(transactionType AdvertiserFundTransactionGetV2TransactionType) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.transactionType = &transactionType
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) Page(page int32) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) Execute() (*AdvertiserFundTransactionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) WithLog(enable bool) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdvertiserFundTransactionGetGet Method for OpenApi2AdvertiserFundTransactionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdvertiserFundTransactionGetGetRequest
*/
func (a *AdvertiserFundTransactionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2AdvertiserFundTransactionGetGetRequest {
	return &ApiOpenApi2AdvertiserFundTransactionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserFundTransactionGetV2Response
func (a *AdvertiserFundTransactionGetV2ApiService) getExecute(r *ApiOpenApi2AdvertiserFundTransactionGetGetRequest) (*AdvertiserFundTransactionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserFundTransactionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/advertiser/fund/transaction/get/"

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
	if r.transactionType == nil {
		return localVarReturnValue, nil, ReportError("transactionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "transaction_type", r.transactionType)
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
