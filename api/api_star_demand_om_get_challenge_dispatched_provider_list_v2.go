/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

// StarDemandOmGetChallengeDispatchedProviderListV2ApiService StarDemandOmGetChallengeDispatchedProviderListV2Api service
type StarDemandOmGetChallengeDispatchedProviderListV2ApiService service

type ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest struct {
	ctx             context.Context
	ApiService      *StarDemandOmGetChallengeDispatchedProviderListV2ApiService
	starId          *int64
	challengeTaskId *int64
	page            *int32
	limit           *int32
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) Page(page int32) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	r.page = &page
	return r
}

// 最大200
func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) Limit(limit int32) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) Execute() (*StarDemandOmGetChallengeDispatchedProviderListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmGetChallengeDispatchedProviderListGet Method for OpenApi2StarDemandOmGetChallengeDispatchedProviderListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest
*/
func (a *StarDemandOmGetChallengeDispatchedProviderListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest {
	return &ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmGetChallengeDispatchedProviderListV2Response
func (a *StarDemandOmGetChallengeDispatchedProviderListV2ApiService) getExecute(r *ApiOpenApi2StarDemandOmGetChallengeDispatchedProviderListGetRequest) (*StarDemandOmGetChallengeDispatchedProviderListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmGetChallengeDispatchedProviderListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_get_challenge_dispatched_provider_list/"

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
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, ReportError("limit is required and must be specified")
	}
	if *r.limit > 200 {
		return localVarReturnValue, nil, ReportError("limit must be less than 200")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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