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

// StarCreateProjectV2ApiService StarCreateProjectV2Api service
type StarCreateProjectV2ApiService service

type ApiOpenApi2StarCreateProjectPostRequest struct {
	ctx                        context.Context
	ApiService                 *StarCreateProjectV2ApiService
	starCreateProjectV2Request *StarCreateProjectV2Request
}

func (r *ApiOpenApi2StarCreateProjectPostRequest) StarCreateProjectV2Request(starCreateProjectV2Request StarCreateProjectV2Request) *ApiOpenApi2StarCreateProjectPostRequest {
	r.starCreateProjectV2Request = &starCreateProjectV2Request
	return r
}

func (r *ApiOpenApi2StarCreateProjectPostRequest) Execute() (*StarCreateProjectV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarCreateProjectPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarCreateProjectPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarCreateProjectPostRequest) WithLog(enable bool) *ApiOpenApi2StarCreateProjectPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarCreateProjectPost Method for OpenApi2StarCreateProjectPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarCreateProjectPostRequest
*/
func (a *StarCreateProjectV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarCreateProjectPostRequest {
	return &ApiOpenApi2StarCreateProjectPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarCreateProjectV2Response
func (a *StarCreateProjectV2ApiService) postExecute(r *ApiOpenApi2StarCreateProjectPostRequest) (*StarCreateProjectV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarCreateProjectV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/create/project/"

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
	localVarPostBody = r.starCreateProjectV2Request
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
