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

// LocalCxtStatusUpdateV30ApiService LocalCxtStatusUpdateV30Api service
type LocalCxtStatusUpdateV30ApiService service

type ApiOpenApiV30LocalCxtStatusUpdatePostRequest struct {
	ctx                            context.Context
	ApiService                     *LocalCxtStatusUpdateV30ApiService
	localCxtStatusUpdateV30Request *LocalCxtStatusUpdateV30Request
}

func (r *ApiOpenApiV30LocalCxtStatusUpdatePostRequest) LocalCxtStatusUpdateV30Request(localCxtStatusUpdateV30Request LocalCxtStatusUpdateV30Request) *ApiOpenApiV30LocalCxtStatusUpdatePostRequest {
	r.localCxtStatusUpdateV30Request = &localCxtStatusUpdateV30Request
	return r
}

func (r *ApiOpenApiV30LocalCxtStatusUpdatePostRequest) Execute() (*LocalCxtStatusUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30LocalCxtStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalCxtStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalCxtStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30LocalCxtStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalCxtStatusUpdatePost Method for OpenApiV30LocalCxtStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalCxtStatusUpdatePostRequest
*/
func (a *LocalCxtStatusUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30LocalCxtStatusUpdatePostRequest {
	return &ApiOpenApiV30LocalCxtStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalCxtStatusUpdateV30Response
func (a *LocalCxtStatusUpdateV30ApiService) postExecute(r *ApiOpenApiV30LocalCxtStatusUpdatePostRequest) (*LocalCxtStatusUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalCxtStatusUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/cxt/status/update/"

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
	localVarPostBody = r.localCxtStatusUpdateV30Request
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