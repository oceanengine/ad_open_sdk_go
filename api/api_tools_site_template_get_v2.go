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

// ToolsSiteTemplateGetV2ApiService ToolsSiteTemplateGetV2Api service
type ToolsSiteTemplateGetV2ApiService service

type ApiOpenApi2ToolsSiteTemplateGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsSiteTemplateGetV2ApiService
	advertiserId *int64
	filter       *ToolsSiteTemplateGetV2Filter
	page         *int64
	pageSize     *int64
}

// 广告主
func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) Filter(filter ToolsSiteTemplateGetV2Filter) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	r.filter = &filter
	return r
}

// 页码，默认值是&#x60;1&#x60;，范围page &gt;&#x3D;1
func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) Page(page int64) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值是&#x60;10&#x60;，范围&#x60;1~100&#x60;
func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) Execute() (*ToolsSiteTemplateGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteTemplateGetGet Method for OpenApi2ToolsSiteTemplateGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteTemplateGetGetRequest
*/
func (a *ToolsSiteTemplateGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSiteTemplateGetGetRequest {
	return &ApiOpenApi2ToolsSiteTemplateGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteTemplateGetV2Response
func (a *ToolsSiteTemplateGetV2ApiService) getExecute(r *ApiOpenApi2ToolsSiteTemplateGetGetRequest) (*ToolsSiteTemplateGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteTemplateGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site_template/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
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
