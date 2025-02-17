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

// ReportAudienceInterestActionListV2ApiService ReportAudienceInterestActionListV2Api service
type ReportAudienceInterestActionListV2ApiService service

type ApiOpenApi2ReportAudienceInterestActionListGetRequest struct {
	ctx          context.Context
	ApiService   *ReportAudienceInterestActionListV2ApiService
	advertiserId *int64
	endDate      **string
	filtering    *ReportAudienceInterestActionListV2Filtering
	metrics      *[]string
	page         *int64
	pageSize     *int64
	startDate    **string
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) Filtering(filtering ReportAudienceInterestActionListV2Filtering) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) Metrics(metrics []string) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) Page(page int64) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) Execute() (*ReportAudienceInterestActionListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) WithLog(enable bool) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportAudienceInterestActionListGet Method for OpenApi2ReportAudienceInterestActionListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportAudienceInterestActionListGetRequest
*/
func (a *ReportAudienceInterestActionListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportAudienceInterestActionListGetRequest {
	return &ApiOpenApi2ReportAudienceInterestActionListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportAudienceInterestActionListV2Response
func (a *ReportAudienceInterestActionListV2ApiService) getExecute(r *ApiOpenApi2ReportAudienceInterestActionListGetRequest) (*ReportAudienceInterestActionListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportAudienceInterestActionListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/audience/interest_action/list/"

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
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.metrics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
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
