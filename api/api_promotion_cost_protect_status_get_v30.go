/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
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

// PromotionCostProtectStatusGetV30ApiService PromotionCostProtectStatusGetV30Api service
type PromotionCostProtectStatusGetV30ApiService service

type ApiOpenApiV30PromotionCostProtectStatusGetGetRequest struct {
	ctx          context.Context
	ApiService   *PromotionCostProtectStatusGetV30ApiService
	advertiserId *int64
	promotionIds *[]int64
}

// 广告主id
func (r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告id
func (r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) PromotionIds(promotionIds []int64) *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest {
	r.promotionIds = &promotionIds
	return r
}

func (r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) Execute() (*PromotionCostProtectStatusGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) WithLog(enable bool) *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionCostProtectStatusGetGet Method for OpenApiV30PromotionCostProtectStatusGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionCostProtectStatusGetGetRequest
*/
func (a *PromotionCostProtectStatusGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest {
	return &ApiOpenApiV30PromotionCostProtectStatusGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionCostProtectStatusGetV30Response
func (a *PromotionCostProtectStatusGetV30ApiService) getExecute(r *ApiOpenApiV30PromotionCostProtectStatusGetGetRequest) (*PromotionCostProtectStatusGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *PromotionCostProtectStatusGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion/cost_protect_status/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.promotionIds == nil {
		return localVarReturnValue, nil, ReportError("promotionIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_ids", r.promotionIds)
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
