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

// SuggWordsV30ApiService SuggWordsV30Api service
type SuggWordsV30ApiService service

type ApiOpenApiV30SuggWordsPostRequest struct {
	ctx                 context.Context
	ApiService          *SuggWordsV30ApiService
	suggWordsV30Request *SuggWordsV30Request
}

func (r *ApiOpenApiV30SuggWordsPostRequest) SuggWordsV30Request(suggWordsV30Request SuggWordsV30Request) *ApiOpenApiV30SuggWordsPostRequest {
	r.suggWordsV30Request = &suggWordsV30Request
	return r
}

func (r *ApiOpenApiV30SuggWordsPostRequest) Execute() (*SuggWordsV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30SuggWordsPostRequest) AccessToken(accessToken string) *ApiOpenApiV30SuggWordsPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SuggWordsPostRequest) WithLog(enable bool) *ApiOpenApiV30SuggWordsPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SuggWordsPost Method for OpenApiV30SuggWordsPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SuggWordsPostRequest
*/
func (a *SuggWordsV30ApiService) Post(ctx context.Context) *ApiOpenApiV30SuggWordsPostRequest {
	return &ApiOpenApiV30SuggWordsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SuggWordsV30Response
func (a *SuggWordsV30ApiService) postExecute(r *ApiOpenApiV30SuggWordsPostRequest) (*SuggWordsV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SuggWordsV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/sugg_words/"

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
	localVarPostBody = r.suggWordsV30Request
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
