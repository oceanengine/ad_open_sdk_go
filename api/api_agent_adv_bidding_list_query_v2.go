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

// AgentAdvBiddingListQueryV2ApiService AgentAdvBiddingListQueryV2Api service
type AgentAdvBiddingListQueryV2ApiService service

type ApiOpenApi2AgentAdvBiddingListQueryGetRequest struct {
	ctx        context.Context
	ApiService *AgentAdvBiddingListQueryV2ApiService
	agentId    *int64
	startDate  *string
	endDate    *string
	filtering  *AgentAdvBiddingListQueryV2Filtering
	cursor     *int64
	cursorSize *int64
}

func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) AgentId(agentId int64) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.agentId = &agentId
	return r
}

// 开始时间。格式：YYYY-MM-DD。闭区间,可选日期范围是今天及以前（允许包括今天）。
func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) StartDate(startDate string) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间。格式：YYYY-MM-DD。闭区间,可选日期范围是今天及以前（允许包括今天）。
func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) EndDate(endDate string) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.endDate = &endDate
	return r
}

// 筛选条件。
func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) Filtering(filtering AgentAdvBiddingListQueryV2Filtering) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.filtering = &filtering
	return r
}

// 游标。第一次查询不填或传0。后续查询使用出参的cursor字段。page * page_size &gt; 10000 请使用此字段和cursor_size字段
func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) Cursor(cursor int64) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.cursor = &cursor
	return r
}

// 页面大小，即一页展示的数据数量。取值范围: 1-1000。
func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) CursorSize(cursorSize int64) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.cursorSize = &cursorSize
	return r
}

func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) Execute() (*AgentAdvBiddingListQueryV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) WithLog(enable bool) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentAdvBiddingListQueryGet Method for OpenApi2AgentAdvBiddingListQueryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentAdvBiddingListQueryGetRequest
*/
func (a *AgentAdvBiddingListQueryV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentAdvBiddingListQueryGetRequest {
	return &ApiOpenApi2AgentAdvBiddingListQueryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentAdvBiddingListQueryV2Response
func (a *AgentAdvBiddingListQueryV2ApiService) getExecute(r *ApiOpenApi2AgentAdvBiddingListQueryGetRequest) (*AgentAdvBiddingListQueryV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentAdvBiddingListQueryV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/adv/bidding/list/query/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.cursorSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor_size", r.cursorSize)
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
