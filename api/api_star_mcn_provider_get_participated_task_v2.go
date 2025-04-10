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

// StarMcnProviderGetParticipatedTaskV2ApiService StarMcnProviderGetParticipatedTaskV2Api service
type StarMcnProviderGetParticipatedTaskV2ApiService service

type ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest struct {
	ctx                     context.Context
	ApiService              *StarMcnProviderGetParticipatedTaskV2ApiService
	starId                  *int64
	page                    *int32
	pageSize                *int32
	providerOrderTaskStatus *int32
	providerTaskCategory    *int32
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) StarId(starId int64) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) Page(page int32) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) PageSize(pageSize int32) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) ProviderOrderTaskStatus(providerOrderTaskStatus int32) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.providerOrderTaskStatus = &providerOrderTaskStatus
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) ProviderTaskCategory(providerTaskCategory int32) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.providerTaskCategory = &providerTaskCategory
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) Execute() (*StarMcnProviderGetParticipatedTaskV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) WithLog(enable bool) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarMcnProviderGetParticipatedTaskGet Method for OpenApi2StarMcnProviderGetParticipatedTaskGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest
*/
func (a *StarMcnProviderGetParticipatedTaskV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest {
	return &ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarMcnProviderGetParticipatedTaskV2Response
func (a *StarMcnProviderGetParticipatedTaskV2ApiService) getExecute(r *ApiOpenApi2StarMcnProviderGetParticipatedTaskGetRequest) (*StarMcnProviderGetParticipatedTaskV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarMcnProviderGetParticipatedTaskV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/mcn/provider_get_participated_task/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 50 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 50")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	if r.providerOrderTaskStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider_order_task_status", r.providerOrderTaskStatus)
	}
	if r.providerTaskCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider_task_category", r.providerTaskCategory)
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
