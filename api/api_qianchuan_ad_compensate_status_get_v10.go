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

// QianchuanAdCompensateStatusGetV10ApiService QianchuanAdCompensateStatusGetV10Api service
type QianchuanAdCompensateStatusGetV10ApiService service

type ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAdCompensateStatusGetV10ApiService
	advertiserId *int64
	adIds        *[]int64
}

func (r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 计划id列表
func (r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) AdIds(adIds []int64) *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) Execute() (*QianchuanAdCompensateStatusGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdCompensateStatusGetGet Method for OpenApiV10QianchuanAdCompensateStatusGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest
*/
func (a *QianchuanAdCompensateStatusGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest {
	return &ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdCompensateStatusGetV10Response
func (a *QianchuanAdCompensateStatusGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAdCompensateStatusGetGetRequest) (*QianchuanAdCompensateStatusGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdCompensateStatusGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/compensate_status/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adIds == nil {
		return localVarReturnValue, nil, ReportError("adIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
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
