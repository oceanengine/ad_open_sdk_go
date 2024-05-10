/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// Oauth2AppAccessTokenApiService Oauth2AppAccessTokenApi service
type Oauth2AppAccessTokenApiService service

type ApiOpenApiOauth2AppAccessTokenPostRequest struct {
	ctx                         context.Context
	ApiService                  *Oauth2AppAccessTokenApiService
	oauth2AppAccessTokenRequest *Oauth2AppAccessTokenRequest
}

func (r *ApiOpenApiOauth2AppAccessTokenPostRequest) Oauth2AppAccessTokenRequest(oauth2AppAccessTokenRequest Oauth2AppAccessTokenRequest) *ApiOpenApiOauth2AppAccessTokenPostRequest {
	r.oauth2AppAccessTokenRequest = &oauth2AppAccessTokenRequest
	return r
}

func (r *ApiOpenApiOauth2AppAccessTokenPostRequest) Execute() (*Oauth2AppAccessTokenResponse, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiOauth2AppAccessTokenPostRequest) WithLog(enable bool) *ApiOpenApiOauth2AppAccessTokenPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiOauth2AppAccessTokenPost Method for OpenApiOauth2AppAccessTokenPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiOauth2AppAccessTokenPostRequest
*/
func (a *Oauth2AppAccessTokenApiService) Post(ctx context.Context) *ApiOpenApiOauth2AppAccessTokenPostRequest {
	return &ApiOpenApiOauth2AppAccessTokenPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Oauth2AppAccessTokenResponse
func (a *Oauth2AppAccessTokenApiService) postExecute(r *ApiOpenApiOauth2AppAccessTokenPostRequest) (*Oauth2AppAccessTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *Oauth2AppAccessTokenResponse
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/oauth2/app_access_token/"

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.oauth2AppAccessTokenRequest
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
