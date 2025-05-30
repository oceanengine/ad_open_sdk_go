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

// QianchuanAdMaterialSuggestionV10ApiService QianchuanAdMaterialSuggestionV10Api service
type QianchuanAdMaterialSuggestionV10ApiService service

type ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAdMaterialSuggestionV10ApiService
	advertiserId *int64
	adId         *int64
	materialIds  *[]int64
}

func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 请求计划id
func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	r.adId = &adId
	return r
}

// 请求素材id数组
func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) MaterialIds(materialIds []int64) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	r.materialIds = &materialIds
	return r
}

func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) Execute() (*QianchuanAdMaterialSuggestionV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdMaterialSuggestionGet Method for OpenApiV10QianchuanAdMaterialSuggestionGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest
*/
func (a *QianchuanAdMaterialSuggestionV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest {
	return &ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdMaterialSuggestionV10Response
func (a *QianchuanAdMaterialSuggestionV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAdMaterialSuggestionGetRequest) (*QianchuanAdMaterialSuggestionV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdMaterialSuggestionV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/material/suggestion/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if r.materialIds == nil {
		return localVarReturnValue, nil, ReportError("materialIds is required and must be specified")
	}
	if len(*r.materialIds) < 1 {
		return localVarReturnValue, nil, ReportError("materialIds must have at least 1 elements")
	}
	if len(*r.materialIds) > 50 {
		return localVarReturnValue, nil, ReportError("materialIds must have less than 50 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_ids", r.materialIds)
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
