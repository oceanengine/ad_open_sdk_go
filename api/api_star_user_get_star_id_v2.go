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

// StarUserGetStarIdV2ApiService StarUserGetStarIdV2Api service
type StarUserGetStarIdV2ApiService service

type ApiOpenApi2StarUserGetStarIdGetRequest struct {
	ctx           context.Context
	ApiService    *StarUserGetStarIdV2ApiService
	starId        *int64
	awemeAuthorId *int64
}

// 客户id
func (r *ApiOpenApi2StarUserGetStarIdGetRequest) StarId(starId int64) *ApiOpenApi2StarUserGetStarIdGetRequest {
	r.starId = &starId
	return r
}

// 达人抖音uid
func (r *ApiOpenApi2StarUserGetStarIdGetRequest) AwemeAuthorId(awemeAuthorId int64) *ApiOpenApi2StarUserGetStarIdGetRequest {
	r.awemeAuthorId = &awemeAuthorId
	return r
}

func (r *ApiOpenApi2StarUserGetStarIdGetRequest) Execute() (*StarUserGetStarIdV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarUserGetStarIdGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarUserGetStarIdGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarUserGetStarIdGetRequest) WithLog(enable bool) *ApiOpenApi2StarUserGetStarIdGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarUserGetStarIdGet Method for OpenApi2StarUserGetStarIdGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarUserGetStarIdGetRequest
*/
func (a *StarUserGetStarIdV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarUserGetStarIdGetRequest {
	return &ApiOpenApi2StarUserGetStarIdGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarUserGetStarIdV2Response
func (a *StarUserGetStarIdV2ApiService) getExecute(r *ApiOpenApi2StarUserGetStarIdGetRequest) (*StarUserGetStarIdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarUserGetStarIdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/user/get_star_id/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.awemeAuthorId == nil {
		return localVarReturnValue, nil, ReportError("awemeAuthorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_author_id", r.awemeAuthorId)
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
