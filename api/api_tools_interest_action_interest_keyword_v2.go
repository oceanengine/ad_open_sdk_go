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

// ToolsInterestActionInterestKeywordV2ApiService ToolsInterestActionInterestKeywordV2Api service
type ToolsInterestActionInterestKeywordV2ApiService service

type ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsInterestActionInterestKeywordV2ApiService
	advertiserId *int64
	queryWords   *string
}

func (r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) QueryWords(queryWords string) *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest {
	r.queryWords = &queryWords
	return r
}

func (r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) Execute() (*ToolsInterestActionInterestKeywordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsInterestActionInterestKeywordGet Method for OpenApi2ToolsInterestActionInterestKeywordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest
*/
func (a *ToolsInterestActionInterestKeywordV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest {
	return &ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsInterestActionInterestKeywordV2Response
func (a *ToolsInterestActionInterestKeywordV2ApiService) getExecute(r *ApiOpenApi2ToolsInterestActionInterestKeywordGetRequest) (*ToolsInterestActionInterestKeywordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsInterestActionInterestKeywordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/interest_action/interest/keyword/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.queryWords == nil {
		return localVarReturnValue, nil, ReportError("queryWords is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_words", r.queryWords)
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
