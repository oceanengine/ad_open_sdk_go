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

// StarTaskBindProjectV2ApiService StarTaskBindProjectV2Api service
type StarTaskBindProjectV2ApiService service

type ApiOpenApi2StarTaskBindProjectPostRequest struct {
	ctx                          context.Context
	ApiService                   *StarTaskBindProjectV2ApiService
	starTaskBindProjectV2Request *StarTaskBindProjectV2Request
}

func (r *ApiOpenApi2StarTaskBindProjectPostRequest) StarTaskBindProjectV2Request(starTaskBindProjectV2Request StarTaskBindProjectV2Request) *ApiOpenApi2StarTaskBindProjectPostRequest {
	r.starTaskBindProjectV2Request = &starTaskBindProjectV2Request
	return r
}

func (r *ApiOpenApi2StarTaskBindProjectPostRequest) Execute() (*StarTaskBindProjectV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarTaskBindProjectPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarTaskBindProjectPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarTaskBindProjectPostRequest) WithLog(enable bool) *ApiOpenApi2StarTaskBindProjectPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarTaskBindProjectPost Method for OpenApi2StarTaskBindProjectPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarTaskBindProjectPostRequest
*/
func (a *StarTaskBindProjectV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarTaskBindProjectPostRequest {
	return &ApiOpenApi2StarTaskBindProjectPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarTaskBindProjectV2Response
func (a *StarTaskBindProjectV2ApiService) postExecute(r *ApiOpenApi2StarTaskBindProjectPostRequest) (*StarTaskBindProjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarTaskBindProjectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/task/bind_project/"

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
	localVarPostBody = r.starTaskBindProjectV2Request
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
