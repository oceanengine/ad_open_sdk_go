/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// StarChallengeGetPushAdResultsV2ApiService StarChallengeGetPushAdResultsV2Api service
type StarChallengeGetPushAdResultsV2ApiService service

type ApiOpenApi2StarChallengeGetPushAdResultsGetRequest struct {
	ctx             context.Context
	ApiService      *StarChallengeGetPushAdResultsV2ApiService
	starId          *int64
	challengeTaskId *int64
	itemIds         *[]int64
}

// 客户星图ID
func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) StarId(starId int64) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	r.starId = &starId
	return r
}

// 投稿任务ID
func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

// 作品ID列表，单次最多50个
func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) ItemIds(itemIds []int64) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	r.itemIds = &itemIds
	return r
}

func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) Execute() (*StarChallengeGetPushAdResultsV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeGetPushAdResultsGet Method for OpenApi2StarChallengeGetPushAdResultsGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeGetPushAdResultsGetRequest
*/
func (a *StarChallengeGetPushAdResultsV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest {
	return &ApiOpenApi2StarChallengeGetPushAdResultsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeGetPushAdResultsV2Response
func (a *StarChallengeGetPushAdResultsV2ApiService) getExecute(r *ApiOpenApi2StarChallengeGetPushAdResultsGetRequest) (*StarChallengeGetPushAdResultsV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeGetPushAdResultsV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/get_push_ad_results/"

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
	if r.itemIds == nil {
		return localVarReturnValue, nil, ReportError("itemIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "item_ids", r.itemIds)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
