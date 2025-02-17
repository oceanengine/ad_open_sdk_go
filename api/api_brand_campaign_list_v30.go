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

// BrandCampaignListV30ApiService BrandCampaignListV30Api service
type BrandCampaignListV30ApiService service

type ApiOpenApiV30BrandCampaignListGetRequest struct {
	ctx          context.Context
	ApiService   *BrandCampaignListV30ApiService
	advertiserId *int64
	filter       *BrandCampaignListV30Filter
	pageInfo     *BrandCampaignListV30PageInfo
}

// 广告主ID
func (r *ApiOpenApiV30BrandCampaignListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandCampaignListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件
func (r *ApiOpenApiV30BrandCampaignListGetRequest) Filter(filter BrandCampaignListV30Filter) *ApiOpenApiV30BrandCampaignListGetRequest {
	r.filter = &filter
	return r
}

// 分页信息
func (r *ApiOpenApiV30BrandCampaignListGetRequest) PageInfo(pageInfo BrandCampaignListV30PageInfo) *ApiOpenApiV30BrandCampaignListGetRequest {
	r.pageInfo = &pageInfo
	return r
}

func (r *ApiOpenApiV30BrandCampaignListGetRequest) Execute() (*BrandCampaignListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BrandCampaignListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandCampaignListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandCampaignListGetRequest) WithLog(enable bool) *ApiOpenApiV30BrandCampaignListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandCampaignListGet Method for OpenApiV30BrandCampaignListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandCampaignListGetRequest
*/
func (a *BrandCampaignListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BrandCampaignListGetRequest {
	return &ApiOpenApiV30BrandCampaignListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandCampaignListV30Response
func (a *BrandCampaignListV30ApiService) getExecute(r *ApiOpenApiV30BrandCampaignListGetRequest) (*BrandCampaignListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandCampaignListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/campaign/list/"

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
	if r.pageInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_info", r.pageInfo)
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
