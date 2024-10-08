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

// ReportStardeliveryTaskVideoDataGetV30ApiService ReportStardeliveryTaskVideoDataGetV30Api service
type ReportStardeliveryTaskVideoDataGetV30ApiService service

type ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportStardeliveryTaskVideoDataGetV30ApiService
	advertiserId *int64
	starTaskId   *int64
	startDate    *string
	endDate      *string
	filtering    *ReportStardeliveryTaskVideoDataGetV30Filtering
	orderField   *string
	orderType    *ReportStardeliveryTaskVideoDataGetV30OrderType
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) StarTaskId(starTaskId int64) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.starTaskId = &starTaskId
	return r
}

// 查询起始日期
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) StartDate(startDate string) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.startDate = &startDate
	return r
}

// 查询结束日期
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) EndDate(endDate string) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.endDate = &endDate
	return r
}

// 过滤器
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) Filtering(filtering ReportStardeliveryTaskVideoDataGetV30Filtering) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.filtering = &filtering
	return r
}

// 排序字段
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) OrderField(orderField string) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) OrderType(orderType ReportStardeliveryTaskVideoDataGetV30OrderType) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.orderType = &orderType
	return r
}

// 页码
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) Page(page int32) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) Execute() (*ReportStardeliveryTaskVideoDataGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportStardeliveryTaskVideoDataGetGet Method for OpenApiV30ReportStardeliveryTaskVideoDataGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest
*/
func (a *ReportStardeliveryTaskVideoDataGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest {
	return &ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportStardeliveryTaskVideoDataGetV30Response
func (a *ReportStardeliveryTaskVideoDataGetV30ApiService) getExecute(r *ApiOpenApiV30ReportStardeliveryTaskVideoDataGetGetRequest) (*ReportStardeliveryTaskVideoDataGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportStardeliveryTaskVideoDataGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/stardelivery/task_video_data/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.starTaskId == nil {
		return localVarReturnValue, nil, ReportError("starTaskId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_id", r.starTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
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
