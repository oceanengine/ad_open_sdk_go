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

// StardeliveryTaskShareV30ApiService StardeliveryTaskShareV30Api service
type StardeliveryTaskShareV30ApiService service

type ApiOpenApiV30StardeliveryTaskSharePostRequest struct {
	ctx                             context.Context
	ApiService                      *StardeliveryTaskShareV30ApiService
	stardeliveryTaskShareV30Request *StardeliveryTaskShareV30Request
}

func (r *ApiOpenApiV30StardeliveryTaskSharePostRequest) StardeliveryTaskShareV30Request(stardeliveryTaskShareV30Request StardeliveryTaskShareV30Request) *ApiOpenApiV30StardeliveryTaskSharePostRequest {
	r.stardeliveryTaskShareV30Request = &stardeliveryTaskShareV30Request
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskSharePostRequest) Execute() (*StardeliveryTaskShareV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskSharePostRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskSharePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskSharePostRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskSharePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskSharePost Method for OpenApiV30StardeliveryTaskSharePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskSharePostRequest
*/
func (a *StardeliveryTaskShareV30ApiService) Post(ctx context.Context) *ApiOpenApiV30StardeliveryTaskSharePostRequest {
	return &ApiOpenApiV30StardeliveryTaskSharePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskShareV30Response
func (a *StardeliveryTaskShareV30ApiService) postExecute(r *ApiOpenApiV30StardeliveryTaskSharePostRequest) (*StardeliveryTaskShareV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskShareV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task/share/"

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
	localVarPostBody = r.stardeliveryTaskShareV30Request
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
