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

// StarChallengeChooseTaskItemWithRewardV2ApiService StarChallengeChooseTaskItemWithRewardV2Api service
type StarChallengeChooseTaskItemWithRewardV2ApiService service

type ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest struct {
	ctx                                            context.Context
	ApiService                                     *StarChallengeChooseTaskItemWithRewardV2ApiService
	starChallengeChooseTaskItemWithRewardV2Request *StarChallengeChooseTaskItemWithRewardV2Request
}

func (r *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest) StarChallengeChooseTaskItemWithRewardV2Request(starChallengeChooseTaskItemWithRewardV2Request StarChallengeChooseTaskItemWithRewardV2Request) *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest {
	r.starChallengeChooseTaskItemWithRewardV2Request = &starChallengeChooseTaskItemWithRewardV2Request
	return r
}

func (r *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest) Execute() (*StarChallengeChooseTaskItemWithRewardV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeChooseTaskItemWithRewardPost Method for OpenApi2StarChallengeChooseTaskItemWithRewardPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest
*/
func (a *StarChallengeChooseTaskItemWithRewardV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest {
	return &ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeChooseTaskItemWithRewardV2Response
func (a *StarChallengeChooseTaskItemWithRewardV2ApiService) postExecute(r *ApiOpenApi2StarChallengeChooseTaskItemWithRewardPostRequest) (*StarChallengeChooseTaskItemWithRewardV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeChooseTaskItemWithRewardV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/choose_task_item_with_reward/"

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
	localVarPostBody = r.starChallengeChooseTaskItemWithRewardV2Request
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
