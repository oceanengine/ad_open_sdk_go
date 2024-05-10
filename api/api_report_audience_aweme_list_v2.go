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

// ReportAudienceAwemeListV2ApiService ReportAudienceAwemeListV2Api service
type ReportAudienceAwemeListV2ApiService service

type ApiOpenApi2ReportAudienceAwemeListGetRequest struct {
	ctx          context.Context
	ApiService   *ReportAudienceAwemeListV2ApiService
	advertiserId *int64
	endDate      **string
	filtering    *ReportAudienceAwemeListV2Filtering
	metrics      *[]string
	page         *int64
	pageSize     *int64
	startDate    **string
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) Filtering(filtering ReportAudienceAwemeListV2Filtering) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) Metrics(metrics []string) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) Page(page int64) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) Execute() (*ReportAudienceAwemeListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportAudienceAwemeListGetRequest) WithLog(enable bool) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportAudienceAwemeListGet Method for OpenApi2ReportAudienceAwemeListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportAudienceAwemeListGetRequest
*/
func (a *ReportAudienceAwemeListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportAudienceAwemeListGetRequest {
	return &ApiOpenApi2ReportAudienceAwemeListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportAudienceAwemeListV2Response
func (a *ReportAudienceAwemeListV2ApiService) getExecute(r *ApiOpenApi2ReportAudienceAwemeListGetRequest) (*ReportAudienceAwemeListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportAudienceAwemeListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/audience/aweme/list/"

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
