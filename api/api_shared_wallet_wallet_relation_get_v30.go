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

// SharedWalletWalletRelationGetV30ApiService SharedWalletWalletRelationGetV30Api service
type SharedWalletWalletRelationGetV30ApiService service

type ApiOpenApiV30SharedWalletWalletRelationGetGetRequest struct {
	ctx            context.Context
	ApiService     *SharedWalletWalletRelationGetV30ApiService
	accountId      *int64
	sharedWalletId *int64
	accountType    *SharedWalletWalletRelationGetV30AccountType
	page           *int32
	pageSize       *int32
}

// 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) AccountId(accountId int64) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.accountId = &accountId
	return r
}

// 共享钱包ID
func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) SharedWalletId(sharedWalletId int64) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.sharedWalletId = &sharedWalletId
	return r
}

// 账户类型
func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) AccountType(accountType SharedWalletWalletRelationGetV30AccountType) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.accountType = &accountType
	return r
}

// 页码；默认为1
func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) Page(page int32) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.page = &page
	return r
}

// 页面大小；默认为10
func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) Execute() (*SharedWalletWalletRelationGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SharedWalletWalletRelationGetGet Method for OpenApiV30SharedWalletWalletRelationGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SharedWalletWalletRelationGetGetRequest
*/
func (a *SharedWalletWalletRelationGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest {
	return &ApiOpenApiV30SharedWalletWalletRelationGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SharedWalletWalletRelationGetV30Response
func (a *SharedWalletWalletRelationGetV30ApiService) getExecute(r *ApiOpenApiV30SharedWalletWalletRelationGetGetRequest) (*SharedWalletWalletRelationGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SharedWalletWalletRelationGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/shared_wallet/wallet_relation/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.sharedWalletId == nil {
		return localVarReturnValue, nil, ReportError("sharedWalletId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "shared_wallet_id", r.sharedWalletId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
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