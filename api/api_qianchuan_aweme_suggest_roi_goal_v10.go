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

// QianchuanAwemeSuggestRoiGoalV10ApiService QianchuanAwemeSuggestRoiGoalV10Api service
type QianchuanAwemeSuggestRoiGoalV10ApiService service

type ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeSuggestRoiGoalV10ApiService
	advertiserId *int64
	awemeId      *int64
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) Execute() (*QianchuanAwemeSuggestRoiGoalV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeSuggestRoiGoalGet Method for OpenApiV10QianchuanAwemeSuggestRoiGoalGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest
*/
func (a *QianchuanAwemeSuggestRoiGoalV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest {
	return &ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeSuggestRoiGoalV10Response
func (a *QianchuanAwemeSuggestRoiGoalV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeSuggestRoiGoalGetRequest) (*QianchuanAwemeSuggestRoiGoalV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeSuggestRoiGoalV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/suggest/roi/goal/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
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
