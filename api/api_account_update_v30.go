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

// AccountUpdateV30ApiService AccountUpdateV30Api service
type AccountUpdateV30ApiService service

type ApiOpenApiV30AccountUpdatePostRequest struct {
	ctx                     context.Context
	ApiService              *AccountUpdateV30ApiService
	accountUpdateV30Request *AccountUpdateV30Request
}

func (r *ApiOpenApiV30AccountUpdatePostRequest) AccountUpdateV30Request(accountUpdateV30Request AccountUpdateV30Request) *ApiOpenApiV30AccountUpdatePostRequest {
	r.accountUpdateV30Request = &accountUpdateV30Request
	return r
}

func (r *ApiOpenApiV30AccountUpdatePostRequest) Execute() (*AccountUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AccountUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30AccountUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AccountUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30AccountUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AccountUpdatePost Method for OpenApiV30AccountUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AccountUpdatePostRequest
*/
func (a *AccountUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AccountUpdatePostRequest {
	return &ApiOpenApiV30AccountUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountUpdateV30Response
func (a *AccountUpdateV30ApiService) postExecute(r *ApiOpenApiV30AccountUpdatePostRequest) (*AccountUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AccountUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/account/update/"

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
	localVarPostBody = r.accountUpdateV30Request
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
