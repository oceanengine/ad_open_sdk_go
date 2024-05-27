/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanAdRecommendKeywordsGetV10ApiService QianchuanAdRecommendKeywordsGetV10Api service
type QianchuanAdRecommendKeywordsGetV10ApiService service

type ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAdRecommendKeywordsGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanAdRecommendKeywordsGetV10Filtering
	orderField   *QianchuanAdRecommendKeywordsGetV10OrderField
	orderType    *QianchuanAdRecommendKeywordsGetV10OrderType
	cacheId      *string
	page         *int32
	pageSize     *int32
}

// 千川广告账户id
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件 1. “search_word”、“product_id”、“aweme_id”，必须传入一个； 2. 优先按“search_word”过滤，其次按“product_id”过滤，最后按“aweme_id”过滤
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) Filtering(filtering QianchuanAdRecommendKeywordsGetV10Filtering) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.filtering = &filtering
	return r
}

// 排序字段： 推荐度（默认）DEFAULT、 相关性RELEVANCE、 月搜索量PV、 竞争程度COMPETITION
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) OrderField(orderField QianchuanAdRecommendKeywordsGetV10OrderField) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式，允许值： ASC 升序、 DESC 降序（默认）
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) OrderType(orderType QianchuanAdRecommendKeywordsGetV10OrderType) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.orderType = &orderType
	return r
}

// 选填入参。 请求唯一标识，能保证同样的过滤条件返回的数据是一致的。 使用过期后的cache_id请求该接口，不会报错，会当作全新的查询操作来处理。
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) CacheId(cacheId string) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.cacheId = &cacheId
	return r
}

// 页码，默认值：1
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，允许值：500, 1000，默认值：500
func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) Execute() (*QianchuanAdRecommendKeywordsGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdRecommendKeywordsGetGet Method for OpenApiV10QianchuanAdRecommendKeywordsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest
*/
func (a *QianchuanAdRecommendKeywordsGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest {
	return &ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdRecommendKeywordsGetV10Response
func (a *QianchuanAdRecommendKeywordsGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAdRecommendKeywordsGetGetRequest) (*QianchuanAdRecommendKeywordsGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdRecommendKeywordsGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/recommend_keywords/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.cacheId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cache_id", r.cacheId)
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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
