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

// AgentFundTransferSeqCommitV2ApiService AgentFundTransferSeqCommitV2Api service
type AgentFundTransferSeqCommitV2ApiService service

type ApiOpenApi2AgentFundTransferSeqCommitPostRequest struct {
	ctx                                 context.Context
	ApiService                          *AgentFundTransferSeqCommitV2ApiService
	agentFundTransferSeqCommitV2Request *AgentFundTransferSeqCommitV2Request
}

func (r *ApiOpenApi2AgentFundTransferSeqCommitPostRequest) AgentFundTransferSeqCommitV2Request(agentFundTransferSeqCommitV2Request AgentFundTransferSeqCommitV2Request) *ApiOpenApi2AgentFundTransferSeqCommitPostRequest {
	r.agentFundTransferSeqCommitV2Request = &agentFundTransferSeqCommitV2Request
	return r
}

func (r *ApiOpenApi2AgentFundTransferSeqCommitPostRequest) Execute() (*AgentFundTransferSeqCommitV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AgentFundTransferSeqCommitPostRequest) AccessToken(accessToken string) *ApiOpenApi2AgentFundTransferSeqCommitPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentFundTransferSeqCommitPostRequest) WithLog(enable bool) *ApiOpenApi2AgentFundTransferSeqCommitPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentFundTransferSeqCommitPost Method for OpenApi2AgentFundTransferSeqCommitPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentFundTransferSeqCommitPostRequest
*/
func (a *AgentFundTransferSeqCommitV2ApiService) Post(ctx context.Context) *ApiOpenApi2AgentFundTransferSeqCommitPostRequest {
	return &ApiOpenApi2AgentFundTransferSeqCommitPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentFundTransferSeqCommitV2Response
func (a *AgentFundTransferSeqCommitV2ApiService) postExecute(r *ApiOpenApi2AgentFundTransferSeqCommitPostRequest) (*AgentFundTransferSeqCommitV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentFundTransferSeqCommitV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/fund/transfer_seq/commit/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.agentFundTransferSeqCommitV2Request
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
