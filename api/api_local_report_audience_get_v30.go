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

// LocalReportAudienceGetV30ApiService LocalReportAudienceGetV30Api service
type LocalReportAudienceGetV30ApiService service

type ApiOpenApiV30LocalReportAudienceGetGetRequest struct {
	ctx               context.Context
	ApiService        *LocalReportAudienceGetV30ApiService
	localAccountId    *int64
	startDate         *string
	endDate           *string
	audienceDimension *LocalReportAudienceGetV30AudienceDimension
	dataDimension     *LocalReportAudienceGetV30DataDimension
	fields            *[]string
	orderType         *LocalReportAudienceGetV30OrderType
	orderField        *string
	filtering         *LocalReportAudienceGetV30Filtering
	page              *int64
	pageSize          *int64
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) LocalAccountId(localAccountId int64) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.localAccountId = &localAccountId
	return r
}

// 查询起始日期
func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) StartDate(startDate string) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.startDate = &startDate
	return r
}

// 查询结束日期
func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) EndDate(endDate string) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) AudienceDimension(audienceDimension LocalReportAudienceGetV30AudienceDimension) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.audienceDimension = &audienceDimension
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) DataDimension(dataDimension LocalReportAudienceGetV30DataDimension) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.dataDimension = &dataDimension
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) Fields(fields []string) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.fields = &fields
	return r
}

// 排序方式
func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) OrderType(orderType LocalReportAudienceGetV30OrderType) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.orderType = &orderType
	return r
}

// 排序字段
func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) OrderField(orderField string) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.orderField = &orderField
	return r
}

// 过滤器
func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) Filtering(filtering LocalReportAudienceGetV30Filtering) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) Page(page int64) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) Execute() (*LocalReportAudienceGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalReportAudienceGetGetRequest) WithLog(enable bool) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalReportAudienceGetGet Method for OpenApiV30LocalReportAudienceGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalReportAudienceGetGetRequest
*/
func (a *LocalReportAudienceGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30LocalReportAudienceGetGetRequest {
	return &ApiOpenApiV30LocalReportAudienceGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalReportAudienceGetV30Response
func (a *LocalReportAudienceGetV30ApiService) getExecute(r *ApiOpenApiV30LocalReportAudienceGetGetRequest) (*LocalReportAudienceGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalReportAudienceGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/report/audience/get/"

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
	if r.audienceDimension == nil {
		return localVarReturnValue, nil, ReportError("audienceDimension is required and must be specified")
	}
	if r.dataDimension == nil {
		return localVarReturnValue, nil, ReportError("dataDimension is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "local_account_id", r.localAccountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "audience_dimension", r.audienceDimension)
	parameterAddToHeaderOrQuery(localVarQueryParams, "data_dimension", r.dataDimension)
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
