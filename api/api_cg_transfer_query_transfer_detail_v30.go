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

// CgTransferQueryTransferDetailV30ApiService CgTransferQueryTransferDetailV30Api service
type CgTransferQueryTransferDetailV30ApiService service

type ApiOpenApiV30CgTransferQueryTransferDetailGetRequest struct {
	ctx                  context.Context
	ApiService           *CgTransferQueryTransferDetailV30ApiService
	bizRequestNo         *string
	agentId              *int64
	transferBizRequestNo *string
	transferSerial       *string
}

// 请求id，推荐uuid，方便请求链路对齐
func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) BizRequestNo(bizRequestNo string) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	r.bizRequestNo = &bizRequestNo
	return r
}

// 代理商账户id，用于鉴权
func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) AgentId(agentId int64) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	r.agentId = &agentId
	return r
}

// 发起转账的幂等id
func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) TransferBizRequestNo(transferBizRequestNo string) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	r.transferBizRequestNo = &transferBizRequestNo
	return r
}

// 转账单号，与transfer_biz_request_no两者传其一即可
func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) TransferSerial(transferSerial string) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	r.transferSerial = &transferSerial
	return r
}

func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) Execute() (*CgTransferQueryTransferDetailV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferQueryTransferDetailGet Method for OpenApiV30CgTransferQueryTransferDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferQueryTransferDetailGetRequest
*/
func (a *CgTransferQueryTransferDetailV30ApiService) Get(ctx context.Context) *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest {
	return &ApiOpenApiV30CgTransferQueryTransferDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferQueryTransferDetailV30Response
func (a *CgTransferQueryTransferDetailV30ApiService) getExecute(r *ApiOpenApiV30CgTransferQueryTransferDetailGetRequest) (*CgTransferQueryTransferDetailV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferQueryTransferDetailV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/query_transfer_detail/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "biz_request_no", r.bizRequestNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.transferBizRequestNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_biz_request_no", r.transferBizRequestNo)
	}
	if r.transferSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_serial", r.transferSerial)
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
