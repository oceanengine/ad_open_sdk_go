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

// QianchuanBrandAuthorizedGetV10ApiService QianchuanBrandAuthorizedGetV10Api service
type QianchuanBrandAuthorizedGetV10ApiService service

type ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanBrandAuthorizedGetV10ApiService
	advertiserId *int64
}

func (r *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest) Execute() (*QianchuanBrandAuthorizedGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanBrandAuthorizedGetGet Method for OpenApiV10QianchuanBrandAuthorizedGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest
*/
func (a *QianchuanBrandAuthorizedGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest {
	return &ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanBrandAuthorizedGetV10Response
func (a *QianchuanBrandAuthorizedGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanBrandAuthorizedGetGetRequest) (*QianchuanBrandAuthorizedGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanBrandAuthorizedGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/brand/authorized/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
