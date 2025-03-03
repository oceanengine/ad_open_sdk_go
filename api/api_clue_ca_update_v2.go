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

// ClueCaUpdateV2ApiService ClueCaUpdateV2Api service
type ClueCaUpdateV2ApiService service

type ApiOpenApi2ClueCaUpdatePostRequest struct {
	ctx                   context.Context
	ApiService            *ClueCaUpdateV2ApiService
	clueCaUpdateV2Request *ClueCaUpdateV2Request
}

func (r *ApiOpenApi2ClueCaUpdatePostRequest) ClueCaUpdateV2Request(clueCaUpdateV2Request ClueCaUpdateV2Request) *ApiOpenApi2ClueCaUpdatePostRequest {
	r.clueCaUpdateV2Request = &clueCaUpdateV2Request
	return r
}

func (r *ApiOpenApi2ClueCaUpdatePostRequest) Execute() (*ClueCaUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ClueCaUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ClueCaUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueCaUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2ClueCaUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueCaUpdatePost Method for OpenApi2ClueCaUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueCaUpdatePostRequest
*/
func (a *ClueCaUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ClueCaUpdatePostRequest {
	return &ApiOpenApi2ClueCaUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueCaUpdateV2Response
func (a *ClueCaUpdateV2ApiService) postExecute(r *ApiOpenApi2ClueCaUpdatePostRequest) (*ClueCaUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueCaUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/ca/update/"

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
	localVarPostBody = r.clueCaUpdateV2Request
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
