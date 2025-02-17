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

// ToolsEventAllAssetsDetailV2ApiService ToolsEventAllAssetsDetailV2Api service
type ToolsEventAllAssetsDetailV2ApiService service

type ApiOpenApi2ToolsEventAllAssetsDetailGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsEventAllAssetsDetailV2ApiService
	advertiserId *int64
	assetIds     *[]int64
}

// 广告主 id
func (r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 资产id列表，list长度最长50，当账户下不存在该资产id时不会返回，当资产共享失效时，不会返回详情
func (r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) AssetIds(assetIds []int64) *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest {
	r.assetIds = &assetIds
	return r
}

func (r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) Execute() (*ToolsEventAllAssetsDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsEventAllAssetsDetailGet Method for OpenApi2ToolsEventAllAssetsDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsEventAllAssetsDetailGetRequest
*/
func (a *ToolsEventAllAssetsDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest {
	return &ApiOpenApi2ToolsEventAllAssetsDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsEventAllAssetsDetailV2Response
func (a *ToolsEventAllAssetsDetailV2ApiService) getExecute(r *ApiOpenApi2ToolsEventAllAssetsDetailGetRequest) (*ToolsEventAllAssetsDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsEventAllAssetsDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/event/all_assets/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.assetIds == nil {
		return localVarReturnValue, nil, ReportError("assetIds is required and must be specified")
	}
	if len(*r.assetIds) < 1 {
		return localVarReturnValue, nil, ReportError("assetIds must have at least 1 elements")
	}
	if len(*r.assetIds) > 50 {
		return localVarReturnValue, nil, ReportError("assetIds must have less than 50 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_ids", r.assetIds)
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
