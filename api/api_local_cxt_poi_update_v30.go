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

// LocalCxtPoiUpdateV30ApiService LocalCxtPoiUpdateV30Api service
type LocalCxtPoiUpdateV30ApiService service

type ApiOpenApiV30LocalCxtPoiUpdatePostRequest struct {
	ctx                         context.Context
	ApiService                  *LocalCxtPoiUpdateV30ApiService
	localCxtPoiUpdateV30Request *LocalCxtPoiUpdateV30Request
}

func (r *ApiOpenApiV30LocalCxtPoiUpdatePostRequest) LocalCxtPoiUpdateV30Request(localCxtPoiUpdateV30Request LocalCxtPoiUpdateV30Request) *ApiOpenApiV30LocalCxtPoiUpdatePostRequest {
	r.localCxtPoiUpdateV30Request = &localCxtPoiUpdateV30Request
	return r
}

func (r *ApiOpenApiV30LocalCxtPoiUpdatePostRequest) Execute() (*LocalCxtPoiUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30LocalCxtPoiUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalCxtPoiUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalCxtPoiUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30LocalCxtPoiUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalCxtPoiUpdatePost Method for OpenApiV30LocalCxtPoiUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalCxtPoiUpdatePostRequest
*/
func (a *LocalCxtPoiUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30LocalCxtPoiUpdatePostRequest {
	return &ApiOpenApiV30LocalCxtPoiUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalCxtPoiUpdateV30Response
func (a *LocalCxtPoiUpdateV30ApiService) postExecute(r *ApiOpenApiV30LocalCxtPoiUpdatePostRequest) (*LocalCxtPoiUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalCxtPoiUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/cxt/poi/update/"

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
	localVarPostBody = r.localCxtPoiUpdateV30Request
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
