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

// StarDemandOmGetChallengeItemsDataV2ApiService StarDemandOmGetChallengeItemsDataV2Api service
type StarDemandOmGetChallengeItemsDataV2ApiService service

type ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest struct {
	ctx             context.Context
	ApiService      *StarDemandOmGetChallengeItemsDataV2ApiService
	starId          *int64
	challengeTaskId *int64
	page            *int32
	limit           *int32
	developerId     *int64
}

// 客户星图ID
func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.starId = &starId
	return r
}

// 任务ID
func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

// page
func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) Page(page int32) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.page = &page
	return r
}

// limit  小于30
func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) Limit(limit int32) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) DeveloperId(developerId int64) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.developerId = &developerId
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) Execute() (*StarDemandOmGetChallengeItemsDataV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmGetChallengeItemsDataGet Method for OpenApi2StarDemandOmGetChallengeItemsDataGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest
*/
func (a *StarDemandOmGetChallengeItemsDataV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest {
	return &ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmGetChallengeItemsDataV2Response
func (a *StarDemandOmGetChallengeItemsDataV2ApiService) getExecute(r *ApiOpenApi2StarDemandOmGetChallengeItemsDataGetRequest) (*StarDemandOmGetChallengeItemsDataV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmGetChallengeItemsDataV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_get_challenge_items_data/"

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
	if *r.limit < 0 {
		return localVarReturnValue, nil, ReportError("limit must be greater than 0")
	}
	if *r.limit > 30 {
		return localVarReturnValue, nil, ReportError("limit must be less than 30")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
	if r.developerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer_id", r.developerId)
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
