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

// StarComponentUpdateCommonComponentV2ApiService StarComponentUpdateCommonComponentV2Api service
type StarComponentUpdateCommonComponentV2ApiService service

type ApiOpenApi2StarComponentUpdateCommonComponentPostRequest struct {
	ctx                                         context.Context
	ApiService                                  *StarComponentUpdateCommonComponentV2ApiService
	starComponentUpdateCommonComponentV2Request *StarComponentUpdateCommonComponentV2Request
}

func (r *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest) StarComponentUpdateCommonComponentV2Request(starComponentUpdateCommonComponentV2Request StarComponentUpdateCommonComponentV2Request) *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest {
	r.starComponentUpdateCommonComponentV2Request = &starComponentUpdateCommonComponentV2Request
	return r
}

func (r *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest) Execute() (*StarComponentUpdateCommonComponentV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest) WithLog(enable bool) *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarComponentUpdateCommonComponentPost Method for OpenApi2StarComponentUpdateCommonComponentPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarComponentUpdateCommonComponentPostRequest
*/
func (a *StarComponentUpdateCommonComponentV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest {
	return &ApiOpenApi2StarComponentUpdateCommonComponentPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarComponentUpdateCommonComponentV2Response
func (a *StarComponentUpdateCommonComponentV2ApiService) postExecute(r *ApiOpenApi2StarComponentUpdateCommonComponentPostRequest) (*StarComponentUpdateCommonComponentV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarComponentUpdateCommonComponentV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/component/update_common_component/"

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
	localVarPostBody = r.starComponentUpdateCommonComponentV2Request
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
