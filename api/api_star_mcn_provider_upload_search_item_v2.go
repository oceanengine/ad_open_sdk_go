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

// StarMcnProviderUploadSearchItemV2ApiService StarMcnProviderUploadSearchItemV2Api service
type StarMcnProviderUploadSearchItemV2ApiService service

type ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest struct {
	ctx                                      context.Context
	ApiService                               *StarMcnProviderUploadSearchItemV2ApiService
	starMcnProviderUploadSearchItemV2Request *StarMcnProviderUploadSearchItemV2Request
}

func (r *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest) StarMcnProviderUploadSearchItemV2Request(starMcnProviderUploadSearchItemV2Request StarMcnProviderUploadSearchItemV2Request) *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest {
	r.starMcnProviderUploadSearchItemV2Request = &starMcnProviderUploadSearchItemV2Request
	return r
}

func (r *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest) Execute() (*StarMcnProviderUploadSearchItemV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest) WithLog(enable bool) *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarMcnProviderUploadSearchItemPost Method for OpenApi2StarMcnProviderUploadSearchItemPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest
*/
func (a *StarMcnProviderUploadSearchItemV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest {
	return &ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarMcnProviderUploadSearchItemV2Response
func (a *StarMcnProviderUploadSearchItemV2ApiService) postExecute(r *ApiOpenApi2StarMcnProviderUploadSearchItemPostRequest) (*StarMcnProviderUploadSearchItemV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarMcnProviderUploadSearchItemV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/mcn/provider_upload_search_item/"

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
	localVarPostBody = r.starMcnProviderUploadSearchItemV2Request
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
