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

// SharedWalletDailyStatGetV30ApiService SharedWalletDailyStatGetV30Api service
type SharedWalletDailyStatGetV30ApiService service

type ApiOpenApiV30SharedWalletDailyStatGetGetRequest struct {
	ctx            context.Context
	ApiService     *SharedWalletDailyStatGetV30ApiService
	accountId      *int64
	accountType    *SharedWalletDailyStatGetV30AccountType
	sharedWalletId *int64
	startDate      *string
	endDate        *string
	page           *int64
	pageSize       *int64
}

// 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) AccountId(accountId int64) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) AccountType(accountType SharedWalletDailyStatGetV30AccountType) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.accountType = &accountType
	return r
}

// 共享钱包id
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) SharedWalletId(sharedWalletId int64) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.sharedWalletId = &sharedWalletId
	return r
}

// 开始时间，格式YYYY-MM-DD，默认为当前年份的1月1日
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) StartDate(startDate string) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式YYYY-MM-DD，默认为今天，起止时间间隔不能超过一年。
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) EndDate(endDate string) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.endDate = &endDate
	return r
}

// 页码；默认为1；注意：page*page_size不可大于10000
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) Page(page int64) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.page = &page
	return r
}

// 页面大小；默认为10
func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) Execute() (*SharedWalletDailyStatGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SharedWalletDailyStatGetGet Method for OpenApiV30SharedWalletDailyStatGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SharedWalletDailyStatGetGetRequest
*/
func (a *SharedWalletDailyStatGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SharedWalletDailyStatGetGetRequest {
	return &ApiOpenApiV30SharedWalletDailyStatGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SharedWalletDailyStatGetV30Response
func (a *SharedWalletDailyStatGetV30ApiService) getExecute(r *ApiOpenApiV30SharedWalletDailyStatGetGetRequest) (*SharedWalletDailyStatGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SharedWalletDailyStatGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/shared_wallet/daily_stat/get/"

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
	if r.sharedWalletId == nil {
		return localVarReturnValue, nil, ReportError("sharedWalletId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "shared_wallet_id", r.sharedWalletId)
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
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
