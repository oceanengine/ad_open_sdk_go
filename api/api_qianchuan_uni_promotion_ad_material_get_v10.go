/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

// QianchuanUniPromotionAdMaterialGetV10ApiService QianchuanUniPromotionAdMaterialGetV10Api service
type QianchuanUniPromotionAdMaterialGetV10ApiService service

type ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanUniPromotionAdMaterialGetV10ApiService
	advertiserId *int64
	adId         *int64
	filtering    *QianchuanUniPromotionAdMaterialGetV10Filtering
	page         *int32
	pageSize     *QianchuanUniPromotionAdMaterialGetV10PageSize
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) Filtering(filtering QianchuanUniPromotionAdMaterialGetV10Filtering) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.filtering = &filtering
	return r
}

// 页码
func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) PageSize(pageSize QianchuanUniPromotionAdMaterialGetV10PageSize) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) Execute() (*QianchuanUniPromotionAdMaterialGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniPromotionAdMaterialGetGet Method for OpenApiV10QianchuanUniPromotionAdMaterialGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest
*/
func (a *QianchuanUniPromotionAdMaterialGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest {
	return &ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniPromotionAdMaterialGetV10Response
func (a *QianchuanUniPromotionAdMaterialGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanUniPromotionAdMaterialGetGetRequest) (*QianchuanUniPromotionAdMaterialGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniPromotionAdMaterialGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_promotion/ad/material/get/"

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
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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