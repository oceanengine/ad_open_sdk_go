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

// AdUpdateBudgetV2ApiService AdUpdateBudgetV2Api service
type AdUpdateBudgetV2ApiService service

type ApiOpenApi2AdUpdateBudgetPostRequest struct {
	ctx                     context.Context
	ApiService              *AdUpdateBudgetV2ApiService
	adUpdateBudgetV2Request *AdUpdateBudgetV2Request
}

func (r *ApiOpenApi2AdUpdateBudgetPostRequest) AdUpdateBudgetV2Request(adUpdateBudgetV2Request AdUpdateBudgetV2Request) *ApiOpenApi2AdUpdateBudgetPostRequest {
	r.adUpdateBudgetV2Request = &adUpdateBudgetV2Request
	return r
}

func (r *ApiOpenApi2AdUpdateBudgetPostRequest) Execute() (*AdUpdateBudgetV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AdUpdateBudgetPostRequest) AccessToken(accessToken string) *ApiOpenApi2AdUpdateBudgetPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdUpdateBudgetPostRequest) WithLog(enable bool) *ApiOpenApi2AdUpdateBudgetPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdUpdateBudgetPost Method for OpenApi2AdUpdateBudgetPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdUpdateBudgetPostRequest
*/
func (a *AdUpdateBudgetV2ApiService) Post(ctx context.Context) *ApiOpenApi2AdUpdateBudgetPostRequest {
	return &ApiOpenApi2AdUpdateBudgetPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdUpdateBudgetV2Response
func (a *AdUpdateBudgetV2ApiService) postExecute(r *ApiOpenApi2AdUpdateBudgetPostRequest) (*AdUpdateBudgetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdUpdateBudgetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/ad/update/budget/"

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
	localVarPostBody = r.adUpdateBudgetV2Request
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
