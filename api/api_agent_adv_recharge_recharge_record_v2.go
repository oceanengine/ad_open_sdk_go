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

// AgentAdvRechargeRechargeRecordV2ApiService AgentAdvRechargeRechargeRecordV2Api service
type AgentAdvRechargeRechargeRecordV2ApiService service

type ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest struct {
	ctx        context.Context
	ApiService *AgentAdvRechargeRechargeRecordV2ApiService
	agentIds   *[]int64
	startTime  *string
	endTime    *string
	page       *int64
	pageSize   *int64
	filtering  *AgentAdvRechargeRechargeRecordV2Filtering
}

// 代理商id列表
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) AgentIds(agentIds []int64) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.agentIds = &agentIds
	return r
}

// 开始时间
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) StartTime(startTime string) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) EndTime(endTime string) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.endTime = &endTime
	return r
}

// 页码
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) Page(page int64) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.page = &page
	return r
}

// 每页大小
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) PageSize(pageSize int64) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.pageSize = &pageSize
	return r
}

// 过滤器
func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) Filtering(filtering AgentAdvRechargeRechargeRecordV2Filtering) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) Execute() (*AgentAdvRechargeRechargeRecordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) WithLog(enable bool) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentAdvRechargeRechargeRecordGet Method for OpenApi2AgentAdvRechargeRechargeRecordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest
*/
func (a *AgentAdvRechargeRechargeRecordV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest {
	return &ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentAdvRechargeRechargeRecordV2Response
func (a *AgentAdvRechargeRechargeRecordV2ApiService) getExecute(r *ApiOpenApi2AgentAdvRechargeRechargeRecordGetRequest) (*AgentAdvRechargeRechargeRecordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentAdvRechargeRechargeRecordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/adv/recharge/recharge_record/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentIds == nil {
		return localVarReturnValue, nil, ReportError("agentIds is required and must be specified")
	}
	if len(*r.agentIds) < 1 {
		return localVarReturnValue, nil, ReportError("agentIds must have at least 1 elements")
	}
	if len(*r.agentIds) > 20 {
		return localVarReturnValue, nil, ReportError("agentIds must have less than 20 elements")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_ids", r.agentIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
