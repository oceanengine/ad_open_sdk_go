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

// QianchuanSuggestBudgetV10ApiService QianchuanSuggestBudgetV10Api service
type QianchuanSuggestBudgetV10ApiService service

type ApiOpenApiV10QianchuanSuggestBudgetGetRequest struct {
	ctx                context.Context
	ApiService         *QianchuanSuggestBudgetV10ApiService
	advertiserId       *int64
	awemeId            *int64
	liveScheduleType   *QianchuanSuggestBudgetV10LiveScheduleType
	startTime          *string
	endTime            *string
	scheduleTime       *string
	scheduleFixedRange *int64
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) LiveScheduleType(liveScheduleType QianchuanSuggestBudgetV10LiveScheduleType) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.liveScheduleType = &liveScheduleType
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) ScheduleTime(scheduleTime string) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.scheduleTime = &scheduleTime
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) ScheduleFixedRange(scheduleFixedRange int64) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.scheduleFixedRange = &scheduleFixedRange
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) Execute() (*QianchuanSuggestBudgetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanSuggestBudgetGet Method for OpenApiV10QianchuanSuggestBudgetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanSuggestBudgetGetRequest
*/
func (a *QianchuanSuggestBudgetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanSuggestBudgetGetRequest {
	return &ApiOpenApiV10QianchuanSuggestBudgetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanSuggestBudgetV10Response
func (a *QianchuanSuggestBudgetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanSuggestBudgetGetRequest) (*QianchuanSuggestBudgetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanSuggestBudgetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/suggest/budget/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if r.liveScheduleType == nil {
		return localVarReturnValue, nil, ReportError("liveScheduleType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "live_schedule_type", r.liveScheduleType)
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.scheduleTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_time", r.scheduleTime)
	}
	if r.scheduleFixedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_fixed_range", r.scheduleFixedRange)
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
