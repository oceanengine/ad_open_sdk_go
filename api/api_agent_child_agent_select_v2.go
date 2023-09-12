/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// AgentChildAgentSelectV2ApiService AgentChildAgentSelectV2Api service
type AgentChildAgentSelectV2ApiService service

type ApiOpenApi2AgentChildAgentSelectGetRequest struct {
	ctx          context.Context
	ApiService   *AgentChildAgentSelectV2ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) Page(page int32) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) PageSize(pageSize int32) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) Execute() (*AgentChildAgentSelectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentChildAgentSelectGetRequest) WithLog(enable bool) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentChildAgentSelectGet Method for OpenApi2AgentChildAgentSelectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentChildAgentSelectGetRequest
*/
func (a *AgentChildAgentSelectV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentChildAgentSelectGetRequest {
	return &ApiOpenApi2AgentChildAgentSelectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentChildAgentSelectV2Response
func (a *AgentChildAgentSelectV2ApiService) getExecute(r *ApiOpenApi2AgentChildAgentSelectGetRequest) (*AgentChildAgentSelectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *AgentChildAgentSelectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/child_agent/select/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
