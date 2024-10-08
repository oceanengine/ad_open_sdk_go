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

// DpaProductStatusBatchUpdateV2ApiService DpaProductStatusBatchUpdateV2Api service
type DpaProductStatusBatchUpdateV2ApiService service

type ApiOpenApi2DpaProductStatusBatchUpdatePostRequest struct {
	ctx                                  context.Context
	ApiService                           *DpaProductStatusBatchUpdateV2ApiService
	dpaProductStatusBatchUpdateV2Request *DpaProductStatusBatchUpdateV2Request
}

func (r *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest) DpaProductStatusBatchUpdateV2Request(dpaProductStatusBatchUpdateV2Request DpaProductStatusBatchUpdateV2Request) *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest {
	r.dpaProductStatusBatchUpdateV2Request = &dpaProductStatusBatchUpdateV2Request
	return r
}

func (r *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest) Execute() (*DpaProductStatusBatchUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaProductStatusBatchUpdatePost Method for OpenApi2DpaProductStatusBatchUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaProductStatusBatchUpdatePostRequest
*/
func (a *DpaProductStatusBatchUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest {
	return &ApiOpenApi2DpaProductStatusBatchUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaProductStatusBatchUpdateV2Response
func (a *DpaProductStatusBatchUpdateV2ApiService) postExecute(r *ApiOpenApi2DpaProductStatusBatchUpdatePostRequest) (*DpaProductStatusBatchUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DpaProductStatusBatchUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/product_status/batch_update/"

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
	localVarPostBody = r.dpaProductStatusBatchUpdateV2Request
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
