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

// LocalPromotionStatusUpdateV30ApiService LocalPromotionStatusUpdateV30Api service
type LocalPromotionStatusUpdateV30ApiService service

type ApiOpenApiV30LocalPromotionStatusUpdatePostRequest struct {
	ctx                                  context.Context
	ApiService                           *LocalPromotionStatusUpdateV30ApiService
	localPromotionStatusUpdateV30Request *LocalPromotionStatusUpdateV30Request
}

func (r *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest) LocalPromotionStatusUpdateV30Request(localPromotionStatusUpdateV30Request LocalPromotionStatusUpdateV30Request) *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest {
	r.localPromotionStatusUpdateV30Request = &localPromotionStatusUpdateV30Request
	return r
}

func (r *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest) Execute() (*LocalPromotionStatusUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalPromotionStatusUpdatePost Method for OpenApiV30LocalPromotionStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalPromotionStatusUpdatePostRequest
*/
func (a *LocalPromotionStatusUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest {
	return &ApiOpenApiV30LocalPromotionStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalPromotionStatusUpdateV30Response
func (a *LocalPromotionStatusUpdateV30ApiService) postExecute(r *ApiOpenApiV30LocalPromotionStatusUpdatePostRequest) (*LocalPromotionStatusUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalPromotionStatusUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/promotion/status/update/"

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
	localVarPostBody = r.localPromotionStatusUpdateV30Request
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