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

// YuntuBrandInfoGetV30ApiService YuntuBrandInfoGetV30Api service
type YuntuBrandInfoGetV30ApiService service

type ApiOpenApiV30YuntuBrandInfoGetGetRequest struct {
	ctx               context.Context
	ApiService        *YuntuBrandInfoGetV30ApiService
	yuntuBrandId      *int64
	serviceProviderId *int64
}

// 云图品牌ID
func (r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) YuntuBrandId(yuntuBrandId int64) *ApiOpenApiV30YuntuBrandInfoGetGetRequest {
	r.yuntuBrandId = &yuntuBrandId
	return r
}

// 服务商ID
func (r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) ServiceProviderId(serviceProviderId int64) *ApiOpenApiV30YuntuBrandInfoGetGetRequest {
	r.serviceProviderId = &serviceProviderId
	return r
}

func (r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) Execute() (*YuntuBrandInfoGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30YuntuBrandInfoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) WithLog(enable bool) *ApiOpenApiV30YuntuBrandInfoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30YuntuBrandInfoGetGet Method for OpenApiV30YuntuBrandInfoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30YuntuBrandInfoGetGetRequest
*/
func (a *YuntuBrandInfoGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30YuntuBrandInfoGetGetRequest {
	return &ApiOpenApiV30YuntuBrandInfoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return YuntuBrandInfoGetV30Response
func (a *YuntuBrandInfoGetV30ApiService) getExecute(r *ApiOpenApiV30YuntuBrandInfoGetGetRequest) (*YuntuBrandInfoGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *YuntuBrandInfoGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/yuntu/brand_info/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.yuntuBrandId == nil {
		return localVarReturnValue, nil, ReportError("yuntuBrandId is required and must be specified")
	}
	if r.serviceProviderId == nil {
		return localVarReturnValue, nil, ReportError("serviceProviderId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "yuntu_brand_id", r.yuntuBrandId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "service_provider_id", r.serviceProviderId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

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
