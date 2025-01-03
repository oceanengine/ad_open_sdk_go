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

// ReportStardeliveryTaskDataGetV30ApiService ReportStardeliveryTaskDataGetV30Api service
type ReportStardeliveryTaskDataGetV30ApiService service

type ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest struct {
	ctx             context.Context
	ApiService      *ReportStardeliveryTaskDataGetV30ApiService
	advertiserId    *int64
	startDate       *string
	endDate         *string
	starTaskVersion *string
	filtering       *ReportStardeliveryTaskDataGetV30Filtering
	orderField      *string
	orderType       *ReportStardeliveryTaskDataGetV30OrderType
	page            *int32
	pageSize        *int32
}

func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询起始日期
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) StartDate(startDate string) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.startDate = &startDate
	return r
}

// 查询结束日期
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) EndDate(endDate string) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.endDate = &endDate
	return r
}

// 任务版本，枚举值： FROM_STAR_TO_BP FROM_AD 默认为FROM_AD
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) StarTaskVersion(starTaskVersion string) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.starTaskVersion = &starTaskVersion
	return r
}

// 过滤器
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) Filtering(filtering ReportStardeliveryTaskDataGetV30Filtering) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.filtering = &filtering
	return r
}

// 排序字段
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) OrderField(orderField string) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) OrderType(orderType ReportStardeliveryTaskDataGetV30OrderType) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) Page(page int32) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) Execute() (*ReportStardeliveryTaskDataGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportStardeliveryTaskDataGetGet Method for OpenApiV30ReportStardeliveryTaskDataGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest
*/
func (a *ReportStardeliveryTaskDataGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest {
	return &ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportStardeliveryTaskDataGetV30Response
func (a *ReportStardeliveryTaskDataGetV30ApiService) getExecute(r *ApiOpenApiV30ReportStardeliveryTaskDataGetGetRequest) (*ReportStardeliveryTaskDataGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportStardeliveryTaskDataGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/stardelivery/task_data/get/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.starTaskVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_version", r.starTaskVersion)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
