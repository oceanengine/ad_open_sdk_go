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

// ToolsStarTaskTitleTopicGetV2ApiService ToolsStarTaskTitleTopicGetV2Api service
type ToolsStarTaskTitleTopicGetV2ApiService service

type ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest struct {
	ctx                     context.Context
	ApiService              *ToolsStarTaskTitleTopicGetV2ApiService
	advertiserId            *int64
	titleSpecifiesTopicItem *string
}

func (r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) TitleSpecifiesTopicItem(titleSpecifiesTopicItem string) *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest {
	r.titleSpecifiesTopicItem = &titleSpecifiesTopicItem
	return r
}

func (r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) Execute() (*ToolsStarTaskTitleTopicGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsStarTaskTitleTopicGetGet Method for OpenApi2ToolsStarTaskTitleTopicGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest
*/
func (a *ToolsStarTaskTitleTopicGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest {
	return &ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsStarTaskTitleTopicGetV2Response
func (a *ToolsStarTaskTitleTopicGetV2ApiService) getExecute(r *ApiOpenApi2ToolsStarTaskTitleTopicGetGetRequest) (*ToolsStarTaskTitleTopicGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsStarTaskTitleTopicGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/star_task/title_topic/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.titleSpecifiesTopicItem == nil {
		return localVarReturnValue, nil, ReportError("titleSpecifiesTopicItem is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "title_specifies_topic_item", r.titleSpecifiesTopicItem)
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
