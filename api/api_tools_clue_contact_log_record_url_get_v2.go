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

// ToolsClueContactLogRecordUrlGetV2ApiService ToolsClueContactLogRecordUrlGetV2Api service
type ToolsClueContactLogRecordUrlGetV2ApiService service

type ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueContactLogRecordUrlGetV2ApiService
	advertiserId *int64
	clueId       *int64
	filter       *ToolsClueContactLogRecordUrlGetV2Filter
}

// 广告主id
func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 线索id
func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) ClueId(clueId int64) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	r.clueId = &clueId
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) Filter(filter ToolsClueContactLogRecordUrlGetV2Filter) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) Execute() (*ToolsClueContactLogRecordUrlGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueContactLogRecordUrlGetGet Method for OpenApi2ToolsClueContactLogRecordUrlGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest
*/
func (a *ToolsClueContactLogRecordUrlGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest {
	return &ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueContactLogRecordUrlGetV2Response
func (a *ToolsClueContactLogRecordUrlGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueContactLogRecordUrlGetGetRequest) (*ToolsClueContactLogRecordUrlGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueContactLogRecordUrlGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/contact_log/record_url/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.clueId == nil {
		return localVarReturnValue, nil, ReportError("clueId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "clue_id", r.clueId)
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
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
