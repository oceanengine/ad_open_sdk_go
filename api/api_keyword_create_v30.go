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

// KeywordCreateV30ApiService KeywordCreateV30Api service
type KeywordCreateV30ApiService service

type ApiOpenApiV30KeywordCreatePostRequest struct {
	ctx                     context.Context
	ApiService              *KeywordCreateV30ApiService
	keywordCreateV30Request *KeywordCreateV30Request
}

func (r *ApiOpenApiV30KeywordCreatePostRequest) KeywordCreateV30Request(keywordCreateV30Request KeywordCreateV30Request) *ApiOpenApiV30KeywordCreatePostRequest {
	r.keywordCreateV30Request = &keywordCreateV30Request
	return r
}

func (r *ApiOpenApiV30KeywordCreatePostRequest) Execute() (*KeywordCreateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30KeywordCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30KeywordCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30KeywordCreatePostRequest) WithLog(enable bool) *ApiOpenApiV30KeywordCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30KeywordCreatePost Method for OpenApiV30KeywordCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30KeywordCreatePostRequest
*/
func (a *KeywordCreateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30KeywordCreatePostRequest {
	return &ApiOpenApiV30KeywordCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KeywordCreateV30Response
func (a *KeywordCreateV30ApiService) postExecute(r *ApiOpenApiV30KeywordCreatePostRequest) (*KeywordCreateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *KeywordCreateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/keyword/create/"

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
	localVarPostBody = r.keywordCreateV30Request
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
