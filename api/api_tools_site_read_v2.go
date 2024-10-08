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

// ToolsSiteReadV2ApiService ToolsSiteReadV2Api service
type ToolsSiteReadV2ApiService service

type ApiOpenApi2ToolsSiteReadGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsSiteReadV2ApiService
	xOrangeCaller *string
	advertiserId  *string
	siteId        *string
}

func (r *ApiOpenApi2ToolsSiteReadGetRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSiteReadGetRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

// 广告主id，广告主id， 传的这个advertiser_id的数字的范围：1 &lt;&#x3D; advertiser_id &lt;&#x3D; MAX_INT64
func (r *ApiOpenApi2ToolsSiteReadGetRequest) AdvertiserId(advertiserId string) *ApiOpenApi2ToolsSiteReadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 橙子建站站点id
func (r *ApiOpenApi2ToolsSiteReadGetRequest) SiteId(siteId string) *ApiOpenApi2ToolsSiteReadGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApi2ToolsSiteReadGetRequest) Execute() (*ToolsSiteReadV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSiteReadGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteReadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteReadGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteReadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteReadGet Method for OpenApi2ToolsSiteReadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteReadGetRequest
*/
func (a *ToolsSiteReadV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSiteReadGetRequest {
	return &ApiOpenApi2ToolsSiteReadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteReadV2Response
func (a *ToolsSiteReadV2ApiService) getExecute(r *ApiOpenApi2ToolsSiteReadGetRequest) (*ToolsSiteReadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteReadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/read/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	if r.xOrangeCaller != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-orange-caller", r.xOrangeCaller)
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
