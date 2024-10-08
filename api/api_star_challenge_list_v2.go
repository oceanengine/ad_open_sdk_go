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

// StarChallengeListV2ApiService StarChallengeListV2Api service
type StarChallengeListV2ApiService service

type ApiOpenApi2StarChallengeListGetRequest struct {
	ctx        context.Context
	ApiService *StarChallengeListV2ApiService
	starId     *int64
	page       *int32
	limit      *int32
}

// 客户星图ID
func (r *ApiOpenApi2StarChallengeListGetRequest) StarId(starId int64) *ApiOpenApi2StarChallengeListGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarChallengeListGetRequest) Page(page int32) *ApiOpenApi2StarChallengeListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarChallengeListGetRequest) Limit(limit int32) *ApiOpenApi2StarChallengeListGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarChallengeListGetRequest) Execute() (*StarChallengeListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarChallengeListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeListGetRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeListGet Method for OpenApi2StarChallengeListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeListGetRequest
*/
func (a *StarChallengeListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarChallengeListGetRequest {
	return &ApiOpenApi2StarChallengeListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeListV2Response
func (a *StarChallengeListV2ApiService) getExecute(r *ApiOpenApi2StarChallengeListGetRequest) (*StarChallengeListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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
