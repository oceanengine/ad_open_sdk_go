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

// StarChallengeItemsDataV2ApiService StarChallengeItemsDataV2Api service
type StarChallengeItemsDataV2ApiService service

type ApiOpenApi2StarChallengeItemsDataGetRequest struct {
	ctx             context.Context
	ApiService      *StarChallengeItemsDataV2ApiService
	starId          *int64
	challengeTaskId *int64
	page            *int32
	limit           *int32
}

// 客户星图ID
func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) StarId(starId int64) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	r.starId = &starId
	return r
}

// 投稿任务ID
func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

// 页码
func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) Page(page int32) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	r.page = &page
	return r
}

// 分页大小，最大50
func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) Limit(limit int32) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) Execute() (*StarChallengeItemsDataV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeItemsDataGetRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeItemsDataGet Method for OpenApi2StarChallengeItemsDataGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeItemsDataGetRequest
*/
func (a *StarChallengeItemsDataV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarChallengeItemsDataGetRequest {
	return &ApiOpenApi2StarChallengeItemsDataGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeItemsDataV2Response
func (a *StarChallengeItemsDataV2ApiService) getExecute(r *ApiOpenApi2StarChallengeItemsDataGetRequest) (*StarChallengeItemsDataV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeItemsDataV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/items_data/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.challengeTaskId == nil {
		return localVarReturnValue, nil, ReportError("challengeTaskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
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
