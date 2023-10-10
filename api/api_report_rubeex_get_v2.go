/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
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

// ReportRubeexGetV2ApiService ReportRubeexGetV2Api service
type ReportRubeexGetV2ApiService service

type ApiOpenApi2ReportRubeexGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportRubeexGetV2ApiService
	advertiserId *int64
	dimensions   *[]string
	filtering    *ReportRubeexGetV2Filtering
	metrics      *[]string
	order        *ReportRubeexGetV2Order
	page         *int64
	pageSize     *int64
	projectId    *int64
	statTimeDay  *[]string
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Dimensions(dimensions []string) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.dimensions = &dimensions
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Filtering(filtering ReportRubeexGetV2Filtering) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Metrics(metrics []string) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Order(order ReportRubeexGetV2Order) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.order = &order
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Page(page int64) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) ProjectId(projectId int64) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.projectId = &projectId
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) StatTimeDay(statTimeDay []string) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.statTimeDay = &statTimeDay
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) Execute() (*ReportRubeexGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportRubeexGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportRubeexGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportRubeexGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportRubeexGetGet Method for OpenApi2ReportRubeexGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportRubeexGetGetRequest
*/
func (a *ReportRubeexGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportRubeexGetGetRequest {
	return &ApiOpenApi2ReportRubeexGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportRubeexGetV2Response
func (a *ReportRubeexGetV2ApiService) getExecute(r *ApiOpenApi2ReportRubeexGetGetRequest) (*ReportRubeexGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportRubeexGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/rubeex/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.dimensions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dimensions", r.dimensions)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.metrics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.projectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
	}
	if r.statTimeDay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stat_time_day", r.statTimeDay)
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
