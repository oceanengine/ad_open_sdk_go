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

// QianchuanUniAwemeAuthorizedGetV10ApiService QianchuanUniAwemeAuthorizedGetV10Api service
type QianchuanUniAwemeAuthorizedGetV10ApiService service

type ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanUniAwemeAuthorizedGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanUniAwemeAuthorizedGetV10Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤器
func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) Filtering(filtering QianchuanUniAwemeAuthorizedGetV10Filtering) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) Execute() (*QianchuanUniAwemeAuthorizedGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniAwemeAuthorizedGetGet Method for OpenApiV10QianchuanUniAwemeAuthorizedGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest
*/
func (a *QianchuanUniAwemeAuthorizedGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest {
	return &ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniAwemeAuthorizedGetV10Response
func (a *QianchuanUniAwemeAuthorizedGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanUniAwemeAuthorizedGetGetRequest) (*QianchuanUniAwemeAuthorizedGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniAwemeAuthorizedGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_aweme/authorized/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
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
