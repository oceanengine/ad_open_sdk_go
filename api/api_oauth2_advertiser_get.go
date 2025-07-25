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

// Oauth2AdvertiserGetApiService Oauth2AdvertiserGetApi service
type Oauth2AdvertiserGetApiService service

type ApiOpenApiOauth2AdvertiserGetGetRequest struct {
	ctx         context.Context
	ApiService  *Oauth2AdvertiserGetApiService
	accessToken *string
}

// 根据授权auth_code获取生成的AccessToken 授权页面使用相同账号授权对应同一个AccessToken，如使用多个不同的账号授权，则需要区分维护多个不同的AccessToken
func (r *ApiOpenApiOauth2AdvertiserGetGetRequest) AccessToken(accessToken string) *ApiOpenApiOauth2AdvertiserGetGetRequest {
	r.accessToken = &accessToken
	return r
}

func (r *ApiOpenApiOauth2AdvertiserGetGetRequest) Execute() (*Oauth2AdvertiserGetResponse, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiOauth2AdvertiserGetGetRequest) WithLog(enable bool) *ApiOpenApiOauth2AdvertiserGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiOauth2AdvertiserGetGet Method for OpenApiOauth2AdvertiserGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiOauth2AdvertiserGetGetRequest
*/
func (a *Oauth2AdvertiserGetApiService) Get(ctx context.Context) *ApiOpenApiOauth2AdvertiserGetGetRequest {
	return &ApiOpenApiOauth2AdvertiserGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Oauth2AdvertiserGetResponse
func (a *Oauth2AdvertiserGetApiService) getExecute(r *ApiOpenApiOauth2AdvertiserGetGetRequest) (*Oauth2AdvertiserGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *Oauth2AdvertiserGetResponse
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/oauth2/advertiser/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accessToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "access_token", r.accessToken)
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
