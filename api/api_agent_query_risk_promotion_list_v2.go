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

// AgentQueryRiskPromotionListV2ApiService AgentQueryRiskPromotionListV2Api service
type AgentQueryRiskPromotionListV2ApiService service

type ApiOpenApi2AgentQueryRiskPromotionListGetRequest struct {
	ctx          context.Context
	ApiService   *AgentQueryRiskPromotionListV2ApiService
	agentId      *int64
	businessType *AgentQueryRiskPromotionListV2BusinessType
	startDate    *string
	endDate      *string
	cursor       *int64
	count        *int32
	filtering    *AgentQueryRiskPromotionListV2Filtering
}

// 代理商账户ID
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) AgentId(agentId int64) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.agentId = &agentId
	return r
}

// 业务线，默认AD业务线传-1
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) BusinessType(businessType AgentQueryRiskPromotionListV2BusinessType) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.businessType = &businessType
	return r
}

// 推送开始时间，比如：2024-03-01
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) StartDate(startDate string) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.startDate = &startDate
	return r
}

// 推送结束时间，比如：2024-03-01（最长跨度31天）
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) EndDate(endDate string) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.endDate = &endDate
	return r
}

// 页码游标值，初始从Long.MAX开始，后续传入返回的cursor值，不传值相当于page&#x3D;1，查询count条数据
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) Cursor(cursor int64) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.cursor = &cursor
	return r
}

// 页码游标值，最大支持500
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) Count(count int32) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.count = &count
	return r
}

// 过滤器
func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) Filtering(filtering AgentQueryRiskPromotionListV2Filtering) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) Execute() (*AgentQueryRiskPromotionListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) WithLog(enable bool) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentQueryRiskPromotionListGet Method for OpenApi2AgentQueryRiskPromotionListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentQueryRiskPromotionListGetRequest
*/
func (a *AgentQueryRiskPromotionListV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentQueryRiskPromotionListGetRequest {
	return &ApiOpenApi2AgentQueryRiskPromotionListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentQueryRiskPromotionListV2Response
func (a *AgentQueryRiskPromotionListV2ApiService) getExecute(r *ApiOpenApi2AgentQueryRiskPromotionListGetRequest) (*AgentQueryRiskPromotionListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentQueryRiskPromotionListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/query/risk_promotion_list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.businessType == nil {
		return localVarReturnValue, nil, ReportError("businessType is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "business_type", r.businessType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
