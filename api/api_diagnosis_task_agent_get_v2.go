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

// DiagnosisTaskAgentGetV2ApiService DiagnosisTaskAgentGetV2Api service
type DiagnosisTaskAgentGetV2ApiService service

type ApiOpenApi2DiagnosisTaskAgentGetGetRequest struct {
	ctx        context.Context
	ApiService *DiagnosisTaskAgentGetV2ApiService
	agentId    *int64
	taskIds    *[]int64
}

func (r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) AgentId(agentId int64) *ApiOpenApi2DiagnosisTaskAgentGetGetRequest {
	r.agentId = &agentId
	return r
}

func (r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) TaskIds(taskIds []int64) *ApiOpenApi2DiagnosisTaskAgentGetGetRequest {
	r.taskIds = &taskIds
	return r
}

func (r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) Execute() (*DiagnosisTaskAgentGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2DiagnosisTaskAgentGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) WithLog(enable bool) *ApiOpenApi2DiagnosisTaskAgentGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DiagnosisTaskAgentGetGet Method for OpenApi2DiagnosisTaskAgentGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DiagnosisTaskAgentGetGetRequest
*/
func (a *DiagnosisTaskAgentGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2DiagnosisTaskAgentGetGetRequest {
	return &ApiOpenApi2DiagnosisTaskAgentGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DiagnosisTaskAgentGetV2Response
func (a *DiagnosisTaskAgentGetV2ApiService) getExecute(r *ApiOpenApi2DiagnosisTaskAgentGetGetRequest) (*DiagnosisTaskAgentGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DiagnosisTaskAgentGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/diagnosis_task/agent/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.taskIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_ids", r.taskIds)
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
