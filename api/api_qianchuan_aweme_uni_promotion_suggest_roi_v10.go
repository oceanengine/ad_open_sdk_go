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

// QianchuanAwemeUniPromotionSuggestRoiV10ApiService QianchuanAwemeUniPromotionSuggestRoiV10Api service
type QianchuanAwemeUniPromotionSuggestRoiV10ApiService service

type ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeUniPromotionSuggestRoiV10ApiService
	advertiserId *int64
	awemeId      *int64
	productId    *int64
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	r.productId = &productId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) Execute() (*QianchuanAwemeUniPromotionSuggestRoiV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeUniPromotionSuggestRoiGet Method for OpenApiV10QianchuanAwemeUniPromotionSuggestRoiGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest
*/
func (a *QianchuanAwemeUniPromotionSuggestRoiV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest {
	return &ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeUniPromotionSuggestRoiV10Response
func (a *QianchuanAwemeUniPromotionSuggestRoiV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestRoiGetRequest) (*QianchuanAwemeUniPromotionSuggestRoiV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeUniPromotionSuggestRoiV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/uni_promotion/suggest/roi/"

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
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if *r.awemeId < 1 {
		return localVarReturnValue, nil, ReportError("awemeId must be greater than 1")
	}
	if r.productId == nil {
		return localVarReturnValue, nil, ReportError("productId is required and must be specified")
	}
	if *r.productId < 1 {
		return localVarReturnValue, nil, ReportError("productId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
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
