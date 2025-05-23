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

// ToolsCountryInfoV2ApiService ToolsCountryInfoV2Api service
type ToolsCountryInfoV2ApiService service

type ApiOpenApi2ToolsCountryInfoGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCountryInfoV2ApiService
	advertiserId *int64
	language     *ToolsCountryInfoV2Language
}

// 广告主id
func (r *ApiOpenApi2ToolsCountryInfoGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsCountryInfoGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 语言类型&lt;br&gt;&#x60;ZH_CN&#x60;表示常用名，如“北京”&lt;br&gt;&#x60;ZH_CN_GOV&#x60;表示官方全称，如“北京市”
func (r *ApiOpenApi2ToolsCountryInfoGetRequest) Language(language ToolsCountryInfoV2Language) *ApiOpenApi2ToolsCountryInfoGetRequest {
	r.language = &language
	return r
}

func (r *ApiOpenApi2ToolsCountryInfoGetRequest) Execute() (*ToolsCountryInfoV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsCountryInfoGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsCountryInfoGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsCountryInfoGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsCountryInfoGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsCountryInfoGet Method for OpenApi2ToolsCountryInfoGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsCountryInfoGetRequest
*/
func (a *ToolsCountryInfoV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsCountryInfoGetRequest {
	return &ApiOpenApi2ToolsCountryInfoGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCountryInfoV2Response
func (a *ToolsCountryInfoV2ApiService) getExecute(r *ApiOpenApi2ToolsCountryInfoGetRequest) (*ToolsCountryInfoV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsCountryInfoV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/country/info/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.language == nil {
		return localVarReturnValue, nil, ReportError("language is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language)
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
