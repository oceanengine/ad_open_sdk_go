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

// QueryInvoiceV2ApiService QueryInvoiceV2Api service
type QueryInvoiceV2ApiService service

type ApiOpenApi2QueryInvoiceGetRequest struct {
	ctx        context.Context
	ApiService *QueryInvoiceV2ApiService
	agentId    *int64
	filtering  *QueryInvoiceV2Filtering
	pageSize   *int32
	page       *int64
}

func (r *ApiOpenApi2QueryInvoiceGetRequest) AgentId(agentId int64) *ApiOpenApi2QueryInvoiceGetRequest {
	r.agentId = &agentId
	return r
}

// 过滤器
func (r *ApiOpenApi2QueryInvoiceGetRequest) Filtering(filtering QueryInvoiceV2Filtering) *ApiOpenApi2QueryInvoiceGetRequest {
	r.filtering = &filtering
	return r
}

// 每页数量 最多100
func (r *ApiOpenApi2QueryInvoiceGetRequest) PageSize(pageSize int32) *ApiOpenApi2QueryInvoiceGetRequest {
	r.pageSize = &pageSize
	return r
}

// 页码
func (r *ApiOpenApi2QueryInvoiceGetRequest) Page(page int64) *ApiOpenApi2QueryInvoiceGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2QueryInvoiceGetRequest) Execute() (*QueryInvoiceV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2QueryInvoiceGetRequest) AccessToken(accessToken string) *ApiOpenApi2QueryInvoiceGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2QueryInvoiceGetRequest) WithLog(enable bool) *ApiOpenApi2QueryInvoiceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2QueryInvoiceGet Method for OpenApi2QueryInvoiceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2QueryInvoiceGetRequest
*/
func (a *QueryInvoiceV2ApiService) Get(ctx context.Context) *ApiOpenApi2QueryInvoiceGetRequest {
	return &ApiOpenApi2QueryInvoiceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryInvoiceV2Response
func (a *QueryInvoiceV2ApiService) getExecute(r *ApiOpenApi2QueryInvoiceGetRequest) (*QueryInvoiceV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QueryInvoiceV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/query/invoice/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
