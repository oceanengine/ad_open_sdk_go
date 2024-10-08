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

// QianchuanCreativeRejectReasonV10ApiService QianchuanCreativeRejectReasonV10Api service
type QianchuanCreativeRejectReasonV10ApiService service

type ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanCreativeRejectReasonV10ApiService
	advertiserId *int64
	creativeIds  *[]int64
}

func (r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) CreativeIds(creativeIds []int64) *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest {
	r.creativeIds = &creativeIds
	return r
}

func (r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) Execute() (*QianchuanCreativeRejectReasonV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanCreativeRejectReasonGet Method for OpenApiV10QianchuanCreativeRejectReasonGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest
*/
func (a *QianchuanCreativeRejectReasonV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest {
	return &ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanCreativeRejectReasonV10Response
func (a *QianchuanCreativeRejectReasonV10ApiService) getExecute(r *ApiOpenApiV10QianchuanCreativeRejectReasonGetRequest) (*QianchuanCreativeRejectReasonV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanCreativeRejectReasonV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/creative/reject_reason/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.creativeIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creative_ids", r.creativeIds)
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
