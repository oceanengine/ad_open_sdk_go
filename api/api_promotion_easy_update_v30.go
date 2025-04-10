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

// PromotionEasyUpdateV30ApiService PromotionEasyUpdateV30Api service
type PromotionEasyUpdateV30ApiService service

type ApiOpenApiV30PromotionEasyUpdatePostRequest struct {
	ctx                           context.Context
	ApiService                    *PromotionEasyUpdateV30ApiService
	promotionEasyUpdateV30Request *PromotionEasyUpdateV30Request
}

func (r *ApiOpenApiV30PromotionEasyUpdatePostRequest) PromotionEasyUpdateV30Request(promotionEasyUpdateV30Request PromotionEasyUpdateV30Request) *ApiOpenApiV30PromotionEasyUpdatePostRequest {
	r.promotionEasyUpdateV30Request = &promotionEasyUpdateV30Request
	return r
}

func (r *ApiOpenApiV30PromotionEasyUpdatePostRequest) Execute() (*PromotionEasyUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30PromotionEasyUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionEasyUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionEasyUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30PromotionEasyUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionEasyUpdatePost Method for OpenApiV30PromotionEasyUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionEasyUpdatePostRequest
*/
func (a *PromotionEasyUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30PromotionEasyUpdatePostRequest {
	return &ApiOpenApiV30PromotionEasyUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionEasyUpdateV30Response
func (a *PromotionEasyUpdateV30ApiService) postExecute(r *ApiOpenApiV30PromotionEasyUpdatePostRequest) (*PromotionEasyUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *PromotionEasyUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion_easy/update/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.promotionEasyUpdateV30Request
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
