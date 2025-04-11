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

// QianchuanAdDetailGetV10ApiService QianchuanAdDetailGetV10Api service
type QianchuanAdDetailGetV10ApiService service

type ApiOpenApiV10QianchuanAdDetailGetGetRequest struct {
	ctx                context.Context
	ApiService         *QianchuanAdDetailGetV10ApiService
	advertiserId       *int64
	adId               *int64
	requestMaterialUrl *bool
	version            *string
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) RequestMaterialUrl(requestMaterialUrl bool) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	r.requestMaterialUrl = &requestMaterialUrl
	return r
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) Version(version string) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	r.version = &version
	return r
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) Execute() (*QianchuanAdDetailGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdDetailGetGet Method for OpenApiV10QianchuanAdDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdDetailGetGetRequest
*/
func (a *QianchuanAdDetailGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAdDetailGetGetRequest {
	return &ApiOpenApiV10QianchuanAdDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdDetailGetV10Response
func (a *QianchuanAdDetailGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAdDetailGetGetRequest) (*QianchuanAdDetailGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdDetailGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/detail/get/"

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
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if *r.adId < 1 {
		return localVarReturnValue, nil, ReportError("adId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	if r.requestMaterialUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "request_material_url", r.requestMaterialUrl)
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version)
	}
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
