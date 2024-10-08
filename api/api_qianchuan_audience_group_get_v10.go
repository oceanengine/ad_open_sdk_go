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

// QianchuanAudienceGroupGetV10ApiService QianchuanAudienceGroupGetV10Api service
type QianchuanAudienceGroupGetV10ApiService service

type ApiOpenApiV10QianchuanAudienceGroupGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAudienceGroupGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanAudienceGroupGetV10Filtering
}

func (r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) Filtering(filtering QianchuanAudienceGroupGetV10Filtering) *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) Execute() (*QianchuanAudienceGroupGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAudienceGroupGetGet Method for OpenApiV10QianchuanAudienceGroupGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAudienceGroupGetGetRequest
*/
func (a *QianchuanAudienceGroupGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest {
	return &ApiOpenApiV10QianchuanAudienceGroupGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAudienceGroupGetV10Response
func (a *QianchuanAudienceGroupGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAudienceGroupGetGetRequest) (*QianchuanAudienceGroupGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAudienceGroupGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/audience_group/get/"

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
