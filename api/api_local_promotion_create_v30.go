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

// LocalPromotionCreateV30ApiService LocalPromotionCreateV30Api service
type LocalPromotionCreateV30ApiService service

type ApiOpenApiV30LocalPromotionCreatePostRequest struct {
	ctx                            context.Context
	ApiService                     *LocalPromotionCreateV30ApiService
	localPromotionCreateV30Request *LocalPromotionCreateV30Request
}

func (r *ApiOpenApiV30LocalPromotionCreatePostRequest) LocalPromotionCreateV30Request(localPromotionCreateV30Request LocalPromotionCreateV30Request) *ApiOpenApiV30LocalPromotionCreatePostRequest {
	r.localPromotionCreateV30Request = &localPromotionCreateV30Request
	return r
}

func (r *ApiOpenApiV30LocalPromotionCreatePostRequest) Execute() (*LocalPromotionCreateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30LocalPromotionCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalPromotionCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalPromotionCreatePostRequest) WithLog(enable bool) *ApiOpenApiV30LocalPromotionCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalPromotionCreatePost Method for OpenApiV30LocalPromotionCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalPromotionCreatePostRequest
*/
func (a *LocalPromotionCreateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30LocalPromotionCreatePostRequest {
	return &ApiOpenApiV30LocalPromotionCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalPromotionCreateV30Response
func (a *LocalPromotionCreateV30ApiService) postExecute(r *ApiOpenApiV30LocalPromotionCreatePostRequest) (*LocalPromotionCreateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalPromotionCreateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/promotion/create/"

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
	localVarPostBody = r.localPromotionCreateV30Request
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