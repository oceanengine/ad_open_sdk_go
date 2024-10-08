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

// StardeliveryTaskAuthorVideoDetailV30ApiService StardeliveryTaskAuthorVideoDetailV30Api service
type StardeliveryTaskAuthorVideoDetailV30ApiService service

type ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest struct {
	ctx                context.Context
	ApiService         *StardeliveryTaskAuthorVideoDetailV30ApiService
	advertiserId       *int64
	starTaskId         *int64
	awemeId            *string
	starTaskVideoRange *StardeliveryTaskAuthorVideoDetailV30StarTaskVideoRange
	filtering          *StardeliveryTaskAuthorVideoDetailV30Filtering
	cursor             *int64
	count              *int64
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) StarTaskId(starTaskId int64) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.starTaskId = &starTaskId
	return r
}

// 抖音号id
func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) AwemeId(awemeId string) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.awemeId = &awemeId
	return r
}

// 返回的视频范围
func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) StarTaskVideoRange(starTaskVideoRange StardeliveryTaskAuthorVideoDetailV30StarTaskVideoRange) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.starTaskVideoRange = &starTaskVideoRange
	return r
}

// 过滤器
func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) Filtering(filtering StardeliveryTaskAuthorVideoDetailV30Filtering) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.filtering = &filtering
	return r
}

// 页码游标值
func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) Cursor(cursor int64) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.cursor = &cursor
	return r
}

// 页码游标值
func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) Count(count int64) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) Execute() (*StardeliveryTaskAuthorVideoDetailV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskAuthorVideoDetailGet Method for OpenApiV30StardeliveryTaskAuthorVideoDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest
*/
func (a *StardeliveryTaskAuthorVideoDetailV30ApiService) Get(ctx context.Context) *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest {
	return &ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskAuthorVideoDetailV30Response
func (a *StardeliveryTaskAuthorVideoDetailV30ApiService) getExecute(r *ApiOpenApiV30StardeliveryTaskAuthorVideoDetailGetRequest) (*StardeliveryTaskAuthorVideoDetailV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskAuthorVideoDetailV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task_author_video/detail/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_id", r.starTaskId)
	if r.starTaskVideoRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_video_range", r.starTaskVideoRange)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
