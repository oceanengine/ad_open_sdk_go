/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// CgTransferQueryTransferBalanceV30ApiService CgTransferQueryTransferBalanceV30Api service
type CgTransferQueryTransferBalanceV30ApiService service

type ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest struct {
	ctx           context.Context
	ApiService    *CgTransferQueryTransferBalanceV30ApiService
	bizRequestNo  *string
	agentId       *int64
	accountIdList *[]int64
}

// 请求id，推荐uuid，方便请求链路对齐
func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) BizRequestNo(bizRequestNo string) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	r.bizRequestNo = &bizRequestNo
	return r
}

// 代理商账户id，用于鉴权
func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) AgentId(agentId int64) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	r.agentId = &agentId
	return r
}

// 查询账户id列表
func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) AccountIdList(accountIdList []int64) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	r.accountIdList = &accountIdList
	return r
}

func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) Execute() (*CgTransferQueryTransferBalanceV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferQueryTransferBalanceGet Method for OpenApiV30CgTransferQueryTransferBalanceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest
*/
func (a *CgTransferQueryTransferBalanceV30ApiService) Get(ctx context.Context) *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest {
	return &ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferQueryTransferBalanceV30Response
func (a *CgTransferQueryTransferBalanceV30ApiService) getExecute(r *ApiOpenApiV30CgTransferQueryTransferBalanceGetRequest) (*CgTransferQueryTransferBalanceV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferQueryTransferBalanceV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/query_transfer_balance/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bizRequestNo == nil {
		return localVarReturnValue, nil, ReportError("bizRequestNo is required and must be specified")
	}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.accountIdList == nil {
		return localVarReturnValue, nil, ReportError("accountIdList is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "biz_request_no", r.bizRequestNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id_list", r.accountIdList)
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
