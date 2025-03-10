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

// SecurityAuditResultsV30ApiService SecurityAuditResultsV30Api service
type SecurityAuditResultsV30ApiService service

type ApiOpenApiV30SecurityAuditResultsGetRequest struct {
	ctx        context.Context
	ApiService *SecurityAuditResultsV30ApiService
	accountId  *int64
	objectId   *int64
}

// 代理商id
func (r *ApiOpenApiV30SecurityAuditResultsGetRequest) AccountId(accountId int64) *ApiOpenApiV30SecurityAuditResultsGetRequest {
	r.accountId = &accountId
	return r
}

// 送审对象ID
func (r *ApiOpenApiV30SecurityAuditResultsGetRequest) ObjectId(objectId int64) *ApiOpenApiV30SecurityAuditResultsGetRequest {
	r.objectId = &objectId
	return r
}

func (r *ApiOpenApiV30SecurityAuditResultsGetRequest) Execute() (*SecurityAuditResultsV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SecurityAuditResultsGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SecurityAuditResultsGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SecurityAuditResultsGetRequest) WithLog(enable bool) *ApiOpenApiV30SecurityAuditResultsGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SecurityAuditResultsGet Method for OpenApiV30SecurityAuditResultsGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SecurityAuditResultsGetRequest
*/
func (a *SecurityAuditResultsV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SecurityAuditResultsGetRequest {
	return &ApiOpenApiV30SecurityAuditResultsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityAuditResultsV30Response
func (a *SecurityAuditResultsV30ApiService) getExecute(r *ApiOpenApiV30SecurityAuditResultsGetRequest) (*SecurityAuditResultsV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SecurityAuditResultsV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/security/audit_results/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.objectId == nil {
		return localVarReturnValue, nil, ReportError("objectId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "object_id", r.objectId)
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
