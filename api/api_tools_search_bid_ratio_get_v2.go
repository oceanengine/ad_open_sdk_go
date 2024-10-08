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

// ToolsSearchBidRatioGetV2ApiService ToolsSearchBidRatioGetV2Api service
type ToolsSearchBidRatioGetV2ApiService service

type ApiOpenApi2ToolsSearchBidRatioGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsSearchBidRatioGetV2ApiService
	advertiserId *int64
	adId         *int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告ID，修改广告时需要传
func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) Execute() (*ToolsSearchBidRatioGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSearchBidRatioGetGet Method for OpenApi2ToolsSearchBidRatioGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSearchBidRatioGetGetRequest
*/
func (a *ToolsSearchBidRatioGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	return &ApiOpenApi2ToolsSearchBidRatioGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSearchBidRatioGetV2Response
func (a *ToolsSearchBidRatioGetV2ApiService) getExecute(r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) (*ToolsSearchBidRatioGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSearchBidRatioGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/search_bid_ratio/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
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
