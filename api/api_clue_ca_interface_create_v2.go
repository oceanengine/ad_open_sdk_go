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

// ClueCaInterfaceCreateV2ApiService ClueCaInterfaceCreateV2Api service
type ClueCaInterfaceCreateV2ApiService service

type ApiOpenApi2ClueCaInterfaceCreatePostRequest struct {
	ctx                            context.Context
	ApiService                     *ClueCaInterfaceCreateV2ApiService
	clueCaInterfaceCreateV2Request *ClueCaInterfaceCreateV2Request
}

func (r *ApiOpenApi2ClueCaInterfaceCreatePostRequest) ClueCaInterfaceCreateV2Request(clueCaInterfaceCreateV2Request ClueCaInterfaceCreateV2Request) *ApiOpenApi2ClueCaInterfaceCreatePostRequest {
	r.clueCaInterfaceCreateV2Request = &clueCaInterfaceCreateV2Request
	return r
}

func (r *ApiOpenApi2ClueCaInterfaceCreatePostRequest) Execute() (*ClueCaInterfaceCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ClueCaInterfaceCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ClueCaInterfaceCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueCaInterfaceCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ClueCaInterfaceCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueCaInterfaceCreatePost Method for OpenApi2ClueCaInterfaceCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueCaInterfaceCreatePostRequest
*/
func (a *ClueCaInterfaceCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ClueCaInterfaceCreatePostRequest {
	return &ApiOpenApi2ClueCaInterfaceCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueCaInterfaceCreateV2Response
func (a *ClueCaInterfaceCreateV2ApiService) postExecute(r *ApiOpenApi2ClueCaInterfaceCreatePostRequest) (*ClueCaInterfaceCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueCaInterfaceCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/ca/interface/create/"

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
	localVarPostBody = r.clueCaInterfaceCreateV2Request
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
