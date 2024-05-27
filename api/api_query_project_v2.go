/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// QueryProjectV2ApiService QueryProjectV2Api service
type QueryProjectV2ApiService service

type ApiOpenApi2QueryProjectGetRequest struct {
	ctx        context.Context
	ApiService *QueryProjectV2ApiService
	agentId    *int64
	count      *int32
	cursor     *int64
	filtering  *QueryProjectV2Filtering
}

// 代理商ID
func (r *ApiOpenApi2QueryProjectGetRequest) AgentId(agentId int64) *ApiOpenApi2QueryProjectGetRequest {
	r.agentId = &agentId
	return r
}

// 页码游标值（单页限制，最大1000，必填）
func (r *ApiOpenApi2QueryProjectGetRequest) Count(count int32) *ApiOpenApi2QueryProjectGetRequest {
	r.count = &count
	return r
}

// 页码游标值（根据cursor进行滚动查询，第一次传-1，后续传前一次返回列表中的最小值）
func (r *ApiOpenApi2QueryProjectGetRequest) Cursor(cursor int64) *ApiOpenApi2QueryProjectGetRequest {
	r.cursor = &cursor
	return r
}

// 过滤器
func (r *ApiOpenApi2QueryProjectGetRequest) Filtering(filtering QueryProjectV2Filtering) *ApiOpenApi2QueryProjectGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2QueryProjectGetRequest) Execute() (*QueryProjectV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2QueryProjectGetRequest) AccessToken(accessToken string) *ApiOpenApi2QueryProjectGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2QueryProjectGetRequest) WithLog(enable bool) *ApiOpenApi2QueryProjectGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2QueryProjectGet Method for OpenApi2QueryProjectGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2QueryProjectGetRequest
*/
func (a *QueryProjectV2ApiService) Get(ctx context.Context) *ApiOpenApi2QueryProjectGetRequest {
	return &ApiOpenApi2QueryProjectGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryProjectV2Response
func (a *QueryProjectV2ApiService) getExecute(r *ApiOpenApi2QueryProjectGetRequest) (*QueryProjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QueryProjectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/query/project/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.count == nil {
		return localVarReturnValue, nil, ReportError("count is required and must be specified")
	}
	if r.cursor == nil {
		return localVarReturnValue, nil, ReportError("cursor is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
	parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
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
