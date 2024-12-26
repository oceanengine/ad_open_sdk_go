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

// LocalCxtReportAudienceGetV30ApiService LocalCxtReportAudienceGetV30Api service
type LocalCxtReportAudienceGetV30ApiService service

type ApiOpenApiV30LocalCxtReportAudienceGetGetRequest struct {
	ctx            context.Context
	ApiService     *LocalCxtReportAudienceGetV30ApiService
	localAccountId *int64
	startDate      *string
	endDate        *string
	metrics        *[]string
	dimension      *LocalCxtReportAudienceGetV30Dimension
}

func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) LocalAccountId(localAccountId int64) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.localAccountId = &localAccountId
	return r
}

// 查询起始日期，格式：yyyy-mm-dd
func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) StartDate(startDate string) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.startDate = &startDate
	return r
}

// 查询结束日期，格式：yyyy-mm-dd 注意：时间跨度不能超过365天
func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) EndDate(endDate string) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.endDate = &endDate
	return r
}

// 指标集
func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) Metrics(metrics []string) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) Dimension(dimension LocalCxtReportAudienceGetV30Dimension) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.dimension = &dimension
	return r
}

func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) Execute() (*LocalCxtReportAudienceGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) WithLog(enable bool) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalCxtReportAudienceGetGet Method for OpenApiV30LocalCxtReportAudienceGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalCxtReportAudienceGetGetRequest
*/
func (a *LocalCxtReportAudienceGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest {
	return &ApiOpenApiV30LocalCxtReportAudienceGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalCxtReportAudienceGetV30Response
func (a *LocalCxtReportAudienceGetV30ApiService) getExecute(r *ApiOpenApiV30LocalCxtReportAudienceGetGetRequest) (*LocalCxtReportAudienceGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalCxtReportAudienceGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/cxt/report/audience/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.localAccountId == nil {
		return localVarReturnValue, nil, ReportError("localAccountId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, ReportError("metrics is required and must be specified")
	}
	if len(*r.metrics) < 1 {
		return localVarReturnValue, nil, ReportError("metrics must have at least 1 elements")
	}
	if r.dimension == nil {
		return localVarReturnValue, nil, ReportError("dimension is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "local_account_id", r.localAccountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	parameterAddToHeaderOrQuery(localVarQueryParams, "dimension", r.dimension)
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
