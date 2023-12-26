/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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

// ToolsCommentGetV30ApiService ToolsCommentGetV30Api service
type ToolsCommentGetV30ApiService service

type ApiOpenApiV30ToolsCommentGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCommentGetV30ApiService
	advertiserId *int64
	startTime    *string
	endTime      *string
	orderField   *ToolsCommentGetV30OrderField
	orderType    *ToolsCommentGetV30OrderType
	filtering    *ToolsCommentGetV30Filtering
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询起始时间，格式：yyyy-MM-dd，时间跨度最大三个月，仅支持获取2021-01-01之后的评论
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) StartTime(startTime string) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.startTime = &startTime
	return r
}

// 查询截止时间，格式：yyyy-MM-dd，时间跨度最大三个月，仅支持获取2021-01-01之后的评论
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) EndTime(endTime string) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.endTime = &endTime
	return r
}

// 排序字段 允许值： REPLY_COUNT 按评论回复数量排序 LIKE_COUNT   按点赞数量排序 CREATE_TIME 按评论时间排序（默认值）
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) OrderField(orderField ToolsCommentGetV30OrderField) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序类型 允许值： ASC 升序 DESC 降序（默认值）
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) OrderType(orderType ToolsCommentGetV30OrderType) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.orderType = &orderType
	return r
}

// 过滤条件，若此字段不传，或传空则视为无限制条件
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) Filtering(filtering ToolsCommentGetV30Filtering) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.filtering = &filtering
	return r
}

// 页码 默认值: 1
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) Page(page int64) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.page = &page
	return r
}

// 页面大小 默认值: 10
func (r *ApiOpenApiV30ToolsCommentGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsCommentGetGetRequest) Execute() (*ToolsCommentGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsCommentGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsCommentGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsCommentGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsCommentGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsCommentGetGet Method for OpenApiV30ToolsCommentGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsCommentGetGetRequest
*/
func (a *ToolsCommentGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsCommentGetGetRequest {
	return &ApiOpenApiV30ToolsCommentGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCommentGetV30Response
func (a *ToolsCommentGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsCommentGetGetRequest) (*ToolsCommentGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsCommentGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/comment/get/"

	localVarHeaderParams := make(map[string]string)
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
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
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
