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

// SharedWalletWalletInfoGetV30ApiService SharedWalletWalletInfoGetV30Api service
type SharedWalletWalletInfoGetV30ApiService service

type ApiOpenApiV30SharedWalletWalletInfoGetGetRequest struct {
	ctx          context.Context
	ApiService   *SharedWalletWalletInfoGetV30ApiService
	accountId    *int64
	walletIdList *[]int64
	accountType  *SharedWalletWalletInfoGetV30AccountType
}

// 账户ID，对于巨量广告、千川、本地推是广告主ID，对于代理商是代理商ID
func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) AccountId(accountId int64) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	r.accountId = &accountId
	return r
}

// 资金共享钱包id列表，注意：传入的列表长度不可大于200
func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) WalletIdList(walletIdList []int64) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	r.walletIdList = &walletIdList
	return r
}

// 账户类型
func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) AccountType(accountType SharedWalletWalletInfoGetV30AccountType) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) Execute() (*SharedWalletWalletInfoGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SharedWalletWalletInfoGetGet Method for OpenApiV30SharedWalletWalletInfoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SharedWalletWalletInfoGetGetRequest
*/
func (a *SharedWalletWalletInfoGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest {
	return &ApiOpenApiV30SharedWalletWalletInfoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SharedWalletWalletInfoGetV30Response
func (a *SharedWalletWalletInfoGetV30ApiService) getExecute(r *ApiOpenApiV30SharedWalletWalletInfoGetGetRequest) (*SharedWalletWalletInfoGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SharedWalletWalletInfoGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/shared_wallet/wallet_info/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
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
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "wallet_id_list", r.walletIdList)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
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
