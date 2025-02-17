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

// Oauth2RenewTokenApiService Oauth2RenewTokenApi service
type Oauth2RenewTokenApiService service

type ApiOpenApiOauth2RenewTokenPostRequest struct {
	ctx                     context.Context
	ApiService              *Oauth2RenewTokenApiService
	oauth2RenewTokenRequest *Oauth2RenewTokenRequest
}

func (r *ApiOpenApiOauth2RenewTokenPostRequest) Oauth2RenewTokenRequest(oauth2RenewTokenRequest Oauth2RenewTokenRequest) *ApiOpenApiOauth2RenewTokenPostRequest {
	r.oauth2RenewTokenRequest = &oauth2RenewTokenRequest
	return r
}

func (r *ApiOpenApiOauth2RenewTokenPostRequest) Execute() (*Oauth2RenewTokenResponse, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiOauth2RenewTokenPostRequest) AccessToken(accessToken string) *ApiOpenApiOauth2RenewTokenPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiOauth2RenewTokenPostRequest) WithLog(enable bool) *ApiOpenApiOauth2RenewTokenPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiOauth2RenewTokenPost Method for OpenApiOauth2RenewTokenPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiOauth2RenewTokenPostRequest
*/
func (a *Oauth2RenewTokenApiService) Post(ctx context.Context) *ApiOpenApiOauth2RenewTokenPostRequest {
	return &ApiOpenApiOauth2RenewTokenPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Oauth2RenewTokenResponse
func (a *Oauth2RenewTokenApiService) postExecute(r *ApiOpenApiOauth2RenewTokenPostRequest) (*Oauth2RenewTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *Oauth2RenewTokenResponse
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/oauth2/renew_token/"

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
	localVarPostBody = r.oauth2RenewTokenRequest
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
