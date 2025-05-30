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

// AgentChargeVerifyV2ApiService AgentChargeVerifyV2Api service
type AgentChargeVerifyV2ApiService service

type ApiOpenApi2AgentChargeVerifyGetRequest struct {
	ctx        context.Context
	ApiService *AgentChargeVerifyV2ApiService
	agentId    *int64
	chargeType *AgentChargeVerifyV2ChargeType
}

// 代理商ID
func (r *ApiOpenApi2AgentChargeVerifyGetRequest) AgentId(agentId int64) *ApiOpenApi2AgentChargeVerifyGetRequest {
	r.agentId = &agentId
	return r
}

// 充值类型
func (r *ApiOpenApi2AgentChargeVerifyGetRequest) ChargeType(chargeType AgentChargeVerifyV2ChargeType) *ApiOpenApi2AgentChargeVerifyGetRequest {
	r.chargeType = &chargeType
	return r
}

func (r *ApiOpenApi2AgentChargeVerifyGetRequest) Execute() (*AgentChargeVerifyV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentChargeVerifyGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentChargeVerifyGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentChargeVerifyGetRequest) WithLog(enable bool) *ApiOpenApi2AgentChargeVerifyGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentChargeVerifyGet Method for OpenApi2AgentChargeVerifyGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentChargeVerifyGetRequest
*/
func (a *AgentChargeVerifyV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentChargeVerifyGetRequest {
	return &ApiOpenApi2AgentChargeVerifyGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentChargeVerifyV2Response
func (a *AgentChargeVerifyV2ApiService) getExecute(r *ApiOpenApi2AgentChargeVerifyGetRequest) (*AgentChargeVerifyV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentChargeVerifyV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/charge/verify/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.chargeType == nil {
		return localVarReturnValue, nil, ReportError("chargeType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "charge_type", r.chargeType)
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
