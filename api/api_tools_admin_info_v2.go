/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsAdminInfoV2ApiService ToolsAdminInfoV2Api service
type ToolsAdminInfoV2ApiService service

type ApiOpenApi2ToolsAdminInfoGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdminInfoV2ApiService
	advertiserId *int64
	codes        *[]string
	language     *ToolsAdminInfoV2Language
	subDistrict  *ToolsAdminInfoV2SubDistrict
}

// 广告主id
func (r *ApiOpenApi2ToolsAdminInfoGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdminInfoGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 行政区域编码列表
func (r *ApiOpenApi2ToolsAdminInfoGetRequest) Codes(codes []string) *ApiOpenApi2ToolsAdminInfoGetRequest {
	r.codes = &codes
	return r
}

// 语言类型&lt;br&gt;&#x60;ZH_CN&#x60;表示常用名，如“北京”&lt;br&gt;&#x60;ZH_CN_GOV&#x60;表示官方全称，如“北京市”
func (r *ApiOpenApi2ToolsAdminInfoGetRequest) Language(language ToolsAdminInfoV2Language) *ApiOpenApi2ToolsAdminInfoGetRequest {
	r.language = &language
	return r
}

// 行政区域层级。&lt;br&gt;&#x60;NONE&#x60; 当前层级&lt;br&gt;&#x60;ONE_LEVEL&#x60;下一级区域&lt;br&gt;&#x60;TWO_LEVEL&#x60;下二级区域&lt;br&gt;&#x60;THREE_LEVEL&#x60;下三级区域&lt;br&gt;&#x60;FOUR_LEVEL&#x60;下四级区域
func (r *ApiOpenApi2ToolsAdminInfoGetRequest) SubDistrict(subDistrict ToolsAdminInfoV2SubDistrict) *ApiOpenApi2ToolsAdminInfoGetRequest {
	r.subDistrict = &subDistrict
	return r
}

func (r *ApiOpenApi2ToolsAdminInfoGetRequest) Execute() (*ToolsAdminInfoV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdminInfoGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdminInfoGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdminInfoGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdminInfoGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdminInfoGet Method for OpenApi2ToolsAdminInfoGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdminInfoGetRequest
*/
func (a *ToolsAdminInfoV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdminInfoGetRequest {
	return &ApiOpenApi2ToolsAdminInfoGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdminInfoV2Response
func (a *ToolsAdminInfoV2ApiService) getExecute(r *ApiOpenApi2ToolsAdminInfoGetRequest) (*ToolsAdminInfoV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdminInfoV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/admin/info/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.codes == nil {
		return localVarReturnValue, nil, ReportError("codes is required and must be specified")
	}
	if len(*r.codes) < 1 {
		return localVarReturnValue, nil, ReportError("codes must have at least 1 elements")
	}
	if r.language == nil {
		return localVarReturnValue, nil, ReportError("language is required and must be specified")
	}
	if r.subDistrict == nil {
		return localVarReturnValue, nil, ReportError("subDistrict is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "codes", r.codes)
	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language)
	parameterAddToHeaderOrQuery(localVarQueryParams, "sub_district", r.subDistrict)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
