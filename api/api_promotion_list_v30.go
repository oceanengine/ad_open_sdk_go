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

// PromotionListV30ApiService PromotionListV30Api service
type PromotionListV30ApiService service

type ApiOpenApiV30PromotionListGetRequest struct {
	ctx                         context.Context
	ApiService                  *PromotionListV30ApiService
	advertiserId                *int64
	filtering                   *PromotionListV30Filtering
	fields                      *[]string
	includingMaterialAtrributes *PromotionListV30IncludingMaterialAtrributes
	page                        *int64
	pageSize                    *int64
	cursor                      *int64
	count                       *int64
}

func (r *ApiOpenApiV30PromotionListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30PromotionListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) Filtering(filtering PromotionListV30Filtering) *ApiOpenApiV30PromotionListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) Fields(fields []string) *ApiOpenApiV30PromotionListGetRequest {
	r.fields = &fields
	return r
}

// RETURN_CARRY_DATA 返回视频素材的搬运属性
func (r *ApiOpenApiV30PromotionListGetRequest) IncludingMaterialAtrributes(includingMaterialAtrributes PromotionListV30IncludingMaterialAtrributes) *ApiOpenApiV30PromotionListGetRequest {
	r.includingMaterialAtrributes = &includingMaterialAtrributes
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) Page(page int64) *ApiOpenApiV30PromotionListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30PromotionListGetRequest {
	r.pageSize = &pageSize
	return r
}

// 页码游标值，第一次拉取，无需入参 同时传入时，cursor优先级大于page 注：page+page_size与cursor+count为两种分页方式 cursor+count适用于获取数据记录数≥10000的场景
func (r *ApiOpenApiV30PromotionListGetRequest) Cursor(cursor int64) *ApiOpenApiV30PromotionListGetRequest {
	r.cursor = &cursor
	return r
}

// 页面数据量 注：page+page_size与cursor+count为两种分页方式 cursor+count适用于获取数据记录数≥10000的场景
func (r *ApiOpenApiV30PromotionListGetRequest) Count(count int64) *ApiOpenApiV30PromotionListGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) Execute() (*PromotionListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30PromotionListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionListGetRequest) WithLog(enable bool) *ApiOpenApiV30PromotionListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionListGet Method for OpenApiV30PromotionListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionListGetRequest
*/
func (a *PromotionListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30PromotionListGetRequest {
	return &ApiOpenApiV30PromotionListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionListV30Response
func (a *PromotionListV30ApiService) getExecute(r *ApiOpenApiV30PromotionListGetRequest) (*PromotionListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *PromotionListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.includingMaterialAtrributes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "including_material_atrributes", r.includingMaterialAtrributes)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
