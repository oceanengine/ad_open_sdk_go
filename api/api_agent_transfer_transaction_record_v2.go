/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// AgentTransferTransactionRecordV2ApiService AgentTransferTransactionRecordV2Api service
type AgentTransferTransactionRecordV2ApiService service

type ApiOpenApi2AgentTransferTransactionRecordGetRequest struct {
	ctx        context.Context
	ApiService *AgentTransferTransactionRecordV2ApiService
	agentId    *int64
	startDate  *string
	endDate    *string
	filtering  *AgentTransferTransactionRecordV2Filtering
	page       *int64
	pageSize   *int64
}

func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) AgentId(agentId int64) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.agentId = &agentId
	return r
}

// 开始时间，格式 yyyy-MM-dd，最远可以查询3年内的数据
func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) StartDate(startDate string) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.startDate = &startDate
	return r
}

// 截止时间，格式 yyyy-MM-dd
func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) EndDate(endDate string) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.endDate = &endDate
	return r
}

// 过滤器
func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) Filtering(filtering AgentTransferTransactionRecordV2Filtering) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.filtering = &filtering
	return r
}

// 页码
func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) Page(page int64) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) PageSize(pageSize int64) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) Execute() (*AgentTransferTransactionRecordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) AccessToken(accessToken string) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) WithLog(enable bool) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentTransferTransactionRecordGet Method for OpenApi2AgentTransferTransactionRecordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentTransferTransactionRecordGetRequest
*/
func (a *AgentTransferTransactionRecordV2ApiService) Get(ctx context.Context) *ApiOpenApi2AgentTransferTransactionRecordGetRequest {
	return &ApiOpenApi2AgentTransferTransactionRecordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentTransferTransactionRecordV2Response
func (a *AgentTransferTransactionRecordV2ApiService) getExecute(r *ApiOpenApi2AgentTransferTransactionRecordGetRequest) (*AgentTransferTransactionRecordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentTransferTransactionRecordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/transfer/transaction_record/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
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
