/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// SubscribeAccountsAddV30ApiService SubscribeAccountsAddV30Api service
type SubscribeAccountsAddV30ApiService service

type ApiOpenApiV30SubscribeAccountsAddPostRequest struct {
	ctx                            context.Context
	ApiService                     *SubscribeAccountsAddV30ApiService
	aPPAccessToken                 *string
	subscribeAccountsAddV30Request *SubscribeAccountsAddV30Request
}

func (r *ApiOpenApiV30SubscribeAccountsAddPostRequest) APPAccessToken(aPPAccessToken string) *ApiOpenApiV30SubscribeAccountsAddPostRequest {
	r.aPPAccessToken = &aPPAccessToken
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsAddPostRequest) SubscribeAccountsAddV30Request(subscribeAccountsAddV30Request SubscribeAccountsAddV30Request) *ApiOpenApiV30SubscribeAccountsAddPostRequest {
	r.subscribeAccountsAddV30Request = &subscribeAccountsAddV30Request
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsAddPostRequest) Execute() (*SubscribeAccountsAddV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30SubscribeAccountsAddPostRequest) AccessToken(accessToken string) *ApiOpenApiV30SubscribeAccountsAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsAddPostRequest) WithLog(enable bool) *ApiOpenApiV30SubscribeAccountsAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SubscribeAccountsAddPost Method for OpenApiV30SubscribeAccountsAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SubscribeAccountsAddPostRequest
*/
func (a *SubscribeAccountsAddV30ApiService) Post(ctx context.Context) *ApiOpenApiV30SubscribeAccountsAddPostRequest {
	return &ApiOpenApiV30SubscribeAccountsAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeAccountsAddV30Response
func (a *SubscribeAccountsAddV30ApiService) postExecute(r *ApiOpenApiV30SubscribeAccountsAddPostRequest) (*SubscribeAccountsAddV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *SubscribeAccountsAddV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/subscribe/accounts/add/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aPPAccessToken == nil {
		return localVarReturnValue, nil, ReportError("aPPAccessToken is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "APP-Access-Token", r.aPPAccessToken)
	// body params
	localVarPostBody = r.subscribeAccountsAddV30Request
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
