/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

// SharedWalletWalletBalanceGetV30ApiService SharedWalletWalletBalanceGetV30Api service
type SharedWalletWalletBalanceGetV30ApiService service

type ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest struct {
	ctx                  context.Context
	ApiService           *SharedWalletWalletBalanceGetV30ApiService
	accountId            *int64
	accountType          *SharedWalletWalletBalanceGetV30AccountType
	walletIdList         *[]int64
	walletBalanceFilters *SharedWalletWalletBalanceGetV30WalletBalanceFilters
}

// 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) AccountId(accountId int64) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型
func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) AccountType(accountType SharedWalletWalletBalanceGetV30AccountType) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	r.accountType = &accountType
	return r
}

// 资金共享钱包id列表，注意：传入的列表长度不可大于200
func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) WalletIdList(walletIdList []int64) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	r.walletIdList = &walletIdList
	return r
}

// 余额过滤条件
func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) WalletBalanceFilters(walletBalanceFilters SharedWalletWalletBalanceGetV30WalletBalanceFilters) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	r.walletBalanceFilters = &walletBalanceFilters
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) Execute() (*SharedWalletWalletBalanceGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SharedWalletWalletBalanceGetGet Method for OpenApiV30SharedWalletWalletBalanceGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest
*/
func (a *SharedWalletWalletBalanceGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest {
	return &ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SharedWalletWalletBalanceGetV30Response
func (a *SharedWalletWalletBalanceGetV30ApiService) getExecute(r *ApiOpenApiV30SharedWalletWalletBalanceGetGetRequest) (*SharedWalletWalletBalanceGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SharedWalletWalletBalanceGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/shared_wallet/wallet_balance/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.walletIdList == nil {
		return localVarReturnValue, nil, ReportError("walletIdList is required and must be specified")
	}
	if len(*r.walletIdList) < 1 {
		return localVarReturnValue, nil, ReportError("walletIdList must have at least 1 elements")
	}
	if len(*r.walletIdList) > 200 {
		return localVarReturnValue, nil, ReportError("walletIdList must have less than 200 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "wallet_id_list", r.walletIdList)
	if r.walletBalanceFilters != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wallet_balance_filters", r.walletBalanceFilters)
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
