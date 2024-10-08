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

// StarChallengeExpandRangeV2ApiService StarChallengeExpandRangeV2Api service
type StarChallengeExpandRangeV2ApiService service

type ApiOpenApi2StarChallengeExpandRangePostRequest struct {
	ctx                               context.Context
	ApiService                        *StarChallengeExpandRangeV2ApiService
	starChallengeExpandRangeV2Request *StarChallengeExpandRangeV2Request
}

func (r *ApiOpenApi2StarChallengeExpandRangePostRequest) StarChallengeExpandRangeV2Request(starChallengeExpandRangeV2Request StarChallengeExpandRangeV2Request) *ApiOpenApi2StarChallengeExpandRangePostRequest {
	r.starChallengeExpandRangeV2Request = &starChallengeExpandRangeV2Request
	return r
}

func (r *ApiOpenApi2StarChallengeExpandRangePostRequest) Execute() (*StarChallengeExpandRangeV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarChallengeExpandRangePostRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeExpandRangePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeExpandRangePostRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeExpandRangePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeExpandRangePost Method for OpenApi2StarChallengeExpandRangePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeExpandRangePostRequest
*/
func (a *StarChallengeExpandRangeV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarChallengeExpandRangePostRequest {
	return &ApiOpenApi2StarChallengeExpandRangePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeExpandRangeV2Response
func (a *StarChallengeExpandRangeV2ApiService) postExecute(r *ApiOpenApi2StarChallengeExpandRangePostRequest) (*StarChallengeExpandRangeV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeExpandRangeV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/expand_range/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.starChallengeExpandRangeV2Request
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
