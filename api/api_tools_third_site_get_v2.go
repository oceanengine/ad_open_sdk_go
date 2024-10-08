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

// ToolsThirdSiteGetV2ApiService ToolsThirdSiteGetV2Api service
type ToolsThirdSiteGetV2ApiService service

type ApiOpenApi2ToolsThirdSiteGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsThirdSiteGetV2ApiService
	advertiserId *int64
	filtering    *ToolsThirdSiteGetV2Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) Filtering(filtering ToolsThirdSiteGetV2Filtering) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) Page(page int64) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) Execute() (*ToolsThirdSiteGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsThirdSiteGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsThirdSiteGetGet Method for OpenApi2ToolsThirdSiteGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsThirdSiteGetGetRequest
*/
func (a *ToolsThirdSiteGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsThirdSiteGetGetRequest {
	return &ApiOpenApi2ToolsThirdSiteGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsThirdSiteGetV2Response
func (a *ToolsThirdSiteGetV2ApiService) getExecute(r *ApiOpenApi2ToolsThirdSiteGetGetRequest) (*ToolsThirdSiteGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsThirdSiteGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/third_site/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
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
