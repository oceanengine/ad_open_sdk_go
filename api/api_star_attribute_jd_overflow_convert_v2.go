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

// StarAttributeJdOverflowConvertV2ApiService StarAttributeJdOverflowConvertV2Api service
type StarAttributeJdOverflowConvertV2ApiService service

type ApiOpenApi2StarAttributeJdOverflowConvertPostRequest struct {
	ctx                                     context.Context
	ApiService                              *StarAttributeJdOverflowConvertV2ApiService
	starAttributeJdOverflowConvertV2Request *StarAttributeJdOverflowConvertV2Request
}

func (r *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest) StarAttributeJdOverflowConvertV2Request(starAttributeJdOverflowConvertV2Request StarAttributeJdOverflowConvertV2Request) *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest {
	r.starAttributeJdOverflowConvertV2Request = &starAttributeJdOverflowConvertV2Request
	return r
}

func (r *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest) Execute() (*StarAttributeJdOverflowConvertV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest) WithLog(enable bool) *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarAttributeJdOverflowConvertPost Method for OpenApi2StarAttributeJdOverflowConvertPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarAttributeJdOverflowConvertPostRequest
*/
func (a *StarAttributeJdOverflowConvertV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest {
	return &ApiOpenApi2StarAttributeJdOverflowConvertPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarAttributeJdOverflowConvertV2Response
func (a *StarAttributeJdOverflowConvertV2ApiService) postExecute(r *ApiOpenApi2StarAttributeJdOverflowConvertPostRequest) (*StarAttributeJdOverflowConvertV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarAttributeJdOverflowConvertV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/attribute/jd_overflow_convert/"

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
	localVarPostBody = r.starAttributeJdOverflowConvertV2Request
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
