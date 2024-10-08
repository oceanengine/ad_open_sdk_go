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

// StarReportDataTopicConfigV2ApiService StarReportDataTopicConfigV2Api service
type StarReportDataTopicConfigV2ApiService service

type ApiOpenApi2StarReportDataTopicConfigGetRequest struct {
	ctx        context.Context
	ApiService *StarReportDataTopicConfigV2ApiService
	starId     *int64
	workId     *int64
	demandId   *int64
	topics     *[]*StarReportDataTopicConfigV2Topics
}

// 发起请求的客户的starId
func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) StarId(starId int64) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	r.starId = &starId
	return r
}

// 交付作品Id，如：视频Id，直播间room_id.
func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) WorkId(workId int64) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	r.workId = &workId
	return r
}

// 任务Id
func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) DemandId(demandId int64) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	r.demandId = &demandId
	return r
}

// 数据主题: BASIC_DATA：基础信息、 FLOW_DATA：流量表现、 CONVERT_DATA：转化表现、 SEARCH_DATA：搜索表现、 RECOMMEND_DATA： 种草表现、 DY_SHOP_DATA：抖音进店、 USER_DISTRIBUTION_DATA：用户画像、 INDEX_SCORE_DATA： 指数得分、 COMMENT_DATA：评论数据 直播用户画像仅保留近90天且直播时长 &gt;&#x3D; 25 分钟直播数据
func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) Topics(topics []*StarReportDataTopicConfigV2Topics) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	r.topics = &topics
	return r
}

func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) Execute() (*StarReportDataTopicConfigV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportDataTopicConfigGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportDataTopicConfigGet Method for OpenApi2StarReportDataTopicConfigGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportDataTopicConfigGetRequest
*/
func (a *StarReportDataTopicConfigV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportDataTopicConfigGetRequest {
	return &ApiOpenApi2StarReportDataTopicConfigGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportDataTopicConfigV2Response
func (a *StarReportDataTopicConfigV2ApiService) getExecute(r *ApiOpenApi2StarReportDataTopicConfigGetRequest) (*StarReportDataTopicConfigV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarReportDataTopicConfigV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/data_topic_config/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "work_id", r.workId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "topics", r.topics)
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
