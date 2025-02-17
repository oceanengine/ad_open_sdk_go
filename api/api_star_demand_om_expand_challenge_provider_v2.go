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

// StarDemandOmExpandChallengeProviderV2ApiService StarDemandOmExpandChallengeProviderV2Api service
type StarDemandOmExpandChallengeProviderV2ApiService service

type ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest struct {
	ctx                                          context.Context
	ApiService                                   *StarDemandOmExpandChallengeProviderV2ApiService
	starDemandOmExpandChallengeProviderV2Request *StarDemandOmExpandChallengeProviderV2Request
}

func (r *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest) StarDemandOmExpandChallengeProviderV2Request(starDemandOmExpandChallengeProviderV2Request StarDemandOmExpandChallengeProviderV2Request) *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest {
	r.starDemandOmExpandChallengeProviderV2Request = &starDemandOmExpandChallengeProviderV2Request
	return r
}

func (r *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest) Execute() (*StarDemandOmExpandChallengeProviderV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmExpandChallengeProviderPost Method for OpenApi2StarDemandOmExpandChallengeProviderPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest
*/
func (a *StarDemandOmExpandChallengeProviderV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest {
	return &ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmExpandChallengeProviderV2Response
func (a *StarDemandOmExpandChallengeProviderV2ApiService) postExecute(r *ApiOpenApi2StarDemandOmExpandChallengeProviderPostRequest) (*StarDemandOmExpandChallengeProviderV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmExpandChallengeProviderV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_expand_challenge_provider/"

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
	localVarPostBody = r.starDemandOmExpandChallengeProviderV2Request
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
