/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// QianchuanSuggestBidV10ApiService QianchuanSuggestBidV10Api service
type QianchuanSuggestBidV10ApiService service

type ApiOpenApiV10QianchuanSuggestBidGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanSuggestBidV10ApiService
	advertiserId   *int64
	marketingGoal  *QianchuanSuggestBidV10MarketingGoal
	externalAction *QianchuanSuggestBidV10ExternalAction
	awemeId        *int64
	productId      *int64
	campaignScene  *QianchuanSuggestBidV10CampaignScene
	ecomGuestType  *QianchuanSuggestBidV10EcomGuestType
	shopId         *int64
	brandId        *int64
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) MarketingGoal(marketingGoal QianchuanSuggestBidV10MarketingGoal) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) ExternalAction(externalAction QianchuanSuggestBidV10ExternalAction) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.productId = &productId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) CampaignScene(campaignScene QianchuanSuggestBidV10CampaignScene) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.campaignScene = &campaignScene
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) EcomGuestType(ecomGuestType QianchuanSuggestBidV10EcomGuestType) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.ecomGuestType = &ecomGuestType
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) ShopId(shopId int64) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.shopId = &shopId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) BrandId(brandId int64) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.brandId = &brandId
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) Execute() (*QianchuanSuggestBidV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestBidGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanSuggestBidGet Method for OpenApiV10QianchuanSuggestBidGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanSuggestBidGetRequest
*/
func (a *QianchuanSuggestBidV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanSuggestBidGetRequest {
	return &ApiOpenApiV10QianchuanSuggestBidGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanSuggestBidV10Response
func (a *QianchuanSuggestBidV10ApiService) getExecute(r *ApiOpenApiV10QianchuanSuggestBidGetRequest) (*QianchuanSuggestBidV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanSuggestBidV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/suggest_bid/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.externalAction == nil {
		return localVarReturnValue, nil, ReportError("externalAction is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
	}
	if r.campaignScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_scene", r.campaignScene)
	}
	if r.ecomGuestType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ecom_guest_type", r.ecomGuestType)
	}
	if r.shopId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shop_id", r.shopId)
	}
	if r.brandId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "brand_id", r.brandId)
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
