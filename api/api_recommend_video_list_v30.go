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

// RecommendVideoListV30ApiService RecommendVideoListV30Api service
type RecommendVideoListV30ApiService service

type ApiOpenApiV30RecommendVideoListPostRequest struct {
	ctx                          context.Context
	ApiService                   *RecommendVideoListV30ApiService
	recommendVideoListV30Request *RecommendVideoListV30Request
}

func (r *ApiOpenApiV30RecommendVideoListPostRequest) RecommendVideoListV30Request(recommendVideoListV30Request RecommendVideoListV30Request) *ApiOpenApiV30RecommendVideoListPostRequest {
	r.recommendVideoListV30Request = &recommendVideoListV30Request
	return r
}

func (r *ApiOpenApiV30RecommendVideoListPostRequest) Execute() (*RecommendVideoListV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30RecommendVideoListPostRequest) AccessToken(accessToken string) *ApiOpenApiV30RecommendVideoListPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30RecommendVideoListPostRequest) WithLog(enable bool) *ApiOpenApiV30RecommendVideoListPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30RecommendVideoListPost Method for OpenApiV30RecommendVideoListPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30RecommendVideoListPostRequest
*/
func (a *RecommendVideoListV30ApiService) Post(ctx context.Context) *ApiOpenApiV30RecommendVideoListPostRequest {
	return &ApiOpenApiV30RecommendVideoListPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RecommendVideoListV30Response
func (a *RecommendVideoListV30ApiService) postExecute(r *ApiOpenApiV30RecommendVideoListPostRequest) (*RecommendVideoListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *RecommendVideoListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/recommend/video/list/"

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
	localVarPostBody = r.recommendVideoListV30Request
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
