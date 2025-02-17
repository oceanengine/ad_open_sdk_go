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

// BrandAdCancelDeleteV30ApiService BrandAdCancelDeleteV30Api service
type BrandAdCancelDeleteV30ApiService service

type ApiOpenApiV30BrandAdCancelDeletePostRequest struct {
	ctx                           context.Context
	ApiService                    *BrandAdCancelDeleteV30ApiService
	brandAdCancelDeleteV30Request *BrandAdCancelDeleteV30Request
}

func (r *ApiOpenApiV30BrandAdCancelDeletePostRequest) BrandAdCancelDeleteV30Request(brandAdCancelDeleteV30Request BrandAdCancelDeleteV30Request) *ApiOpenApiV30BrandAdCancelDeletePostRequest {
	r.brandAdCancelDeleteV30Request = &brandAdCancelDeleteV30Request
	return r
}

func (r *ApiOpenApiV30BrandAdCancelDeletePostRequest) Execute() (*BrandAdCancelDeleteV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandAdCancelDeletePostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandAdCancelDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandAdCancelDeletePostRequest) WithLog(enable bool) *ApiOpenApiV30BrandAdCancelDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandAdCancelDeletePost Method for OpenApiV30BrandAdCancelDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandAdCancelDeletePostRequest
*/
func (a *BrandAdCancelDeleteV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandAdCancelDeletePostRequest {
	return &ApiOpenApiV30BrandAdCancelDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandAdCancelDeleteV30Response
func (a *BrandAdCancelDeleteV30ApiService) postExecute(r *ApiOpenApiV30BrandAdCancelDeletePostRequest) (*BrandAdCancelDeleteV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandAdCancelDeleteV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/ad/cancel_delete/"

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
	localVarPostBody = r.brandAdCancelDeleteV30Request
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
