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

// QianchuanAwemeUniPromotionSuggestV10ApiService QianchuanAwemeUniPromotionSuggestV10Api service
type QianchuanAwemeUniPromotionSuggestV10ApiService service

type ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeUniPromotionSuggestV10ApiService
	advertiserId *int64
	awemeId      *int64
	productId    *int64
	bidType      *QianchuanAwemeUniPromotionSuggestV10BidType
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	r.productId = &productId
	return r
}

// 出价类型
func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) BidType(bidType QianchuanAwemeUniPromotionSuggestV10BidType) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	r.bidType = &bidType
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) Execute() (*QianchuanAwemeUniPromotionSuggestV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeUniPromotionSuggestGet Method for OpenApiV10QianchuanAwemeUniPromotionSuggestGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest
*/
func (a *QianchuanAwemeUniPromotionSuggestV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest {
	return &ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeUniPromotionSuggestV10Response
func (a *QianchuanAwemeUniPromotionSuggestV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeUniPromotionSuggestGetRequest) (*QianchuanAwemeUniPromotionSuggestV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeUniPromotionSuggestV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/uni_promotion/suggest/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 0 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 0")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if *r.awemeId < 0 {
		return localVarReturnValue, nil, ReportError("awemeId must be greater than 0")
	}
	if r.productId == nil {
		return localVarReturnValue, nil, ReportError("productId is required and must be specified")
	}
	if *r.productId < 0 {
		return localVarReturnValue, nil, ReportError("productId must be greater than 0")
	}
	if r.bidType == nil {
		return localVarReturnValue, nil, ReportError("bidType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "bid_type", r.bidType)
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
