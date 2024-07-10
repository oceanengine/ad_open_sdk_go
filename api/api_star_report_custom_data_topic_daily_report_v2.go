/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// StarReportCustomDataTopicDailyReportV2ApiService StarReportCustomDataTopicDailyReportV2Api service
type StarReportCustomDataTopicDailyReportV2ApiService service

type ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest struct {
	ctx        context.Context
	ApiService *StarReportCustomDataTopicDailyReportV2ApiService
	starId     *int64
	workId     *int64
	demandId   *int64
	topics     *[]*StarReportCustomDataTopicDailyReportV2Topics
	startTime  *string
	endTime    *string
}

// 发起请求的客户的starId
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) StarId(starId int64) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.starId = &starId
	return r
}

// 交付作品Id：如视频Id、直播间room_id.
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) WorkId(workId int64) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.workId = &workId
	return r
}

// 任务id
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) DemandId(demandId int64) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.demandId = &demandId
	return r
}

// 数据主题: BASIC_DATA：基础信息、 FLOW_DATA：流量表现、 CONVERT_DATA：转化表现、 SEARCH_DATA：搜索表现、 RECOMMEND_DATA： 种草表现
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) Topics(topics []*StarReportCustomDataTopicDailyReportV2Topics) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.topics = &topics
	return r
}

// 起始时间: yyyy-MM-dd
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) StartTime(startTime string) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间: yyyy-MM-dd
func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) EndTime(endTime string) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) Execute() (*StarReportCustomDataTopicDailyReportV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportCustomDataTopicDailyReportGet Method for OpenApi2StarReportCustomDataTopicDailyReportGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest
*/
func (a *StarReportCustomDataTopicDailyReportV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest {
	return &ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportCustomDataTopicDailyReportV2Response
func (a *StarReportCustomDataTopicDailyReportV2ApiService) getExecute(r *ApiOpenApi2StarReportCustomDataTopicDailyReportGetRequest) (*StarReportCustomDataTopicDailyReportV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarReportCustomDataTopicDailyReportV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/custom_data_topic_daily_report/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.workId == nil {
		return localVarReturnValue, nil, ReportError("workId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}
	if r.topics == nil {
		return localVarReturnValue, nil, ReportError("topics is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "work_id", r.workId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "topics", r.topics)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
