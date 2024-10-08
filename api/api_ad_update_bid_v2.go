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

// AdUpdateBidV2ApiService AdUpdateBidV2Api service
type AdUpdateBidV2ApiService service

type ApiOpenApi2AdUpdateBidPostRequest struct {
	ctx                  context.Context
	ApiService           *AdUpdateBidV2ApiService
	adUpdateBidV2Request *AdUpdateBidV2Request
}

func (r *ApiOpenApi2AdUpdateBidPostRequest) AdUpdateBidV2Request(adUpdateBidV2Request AdUpdateBidV2Request) *ApiOpenApi2AdUpdateBidPostRequest {
	r.adUpdateBidV2Request = &adUpdateBidV2Request
	return r
}

func (r *ApiOpenApi2AdUpdateBidPostRequest) Execute() (*AdUpdateBidV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AdUpdateBidPostRequest) AccessToken(accessToken string) *ApiOpenApi2AdUpdateBidPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdUpdateBidPostRequest) WithLog(enable bool) *ApiOpenApi2AdUpdateBidPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdUpdateBidPost Method for OpenApi2AdUpdateBidPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdUpdateBidPostRequest
*/
func (a *AdUpdateBidV2ApiService) Post(ctx context.Context) *ApiOpenApi2AdUpdateBidPostRequest {
	return &ApiOpenApi2AdUpdateBidPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdUpdateBidV2Response
func (a *AdUpdateBidV2ApiService) postExecute(r *ApiOpenApi2AdUpdateBidPostRequest) (*AdUpdateBidV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdUpdateBidV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/ad/update/bid/"

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
	localVarPostBody = r.adUpdateBidV2Request
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
