/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
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

// ToolsInterestActionActionKeywordV2ApiService ToolsInterestActionActionKeywordV2Api service
type ToolsInterestActionActionKeywordV2ApiService service

type ApiOpenApi2ToolsInterestActionActionKeywordGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsInterestActionActionKeywordV2ApiService
	advertiserId *int64
	queryWords   *string
	actionDays   *ToolsInterestActionActionKeywordV2ActionDays
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) QueryWords(queryWords string) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	r.queryWords = &queryWords
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) ActionDays(actionDays ToolsInterestActionActionKeywordV2ActionDays) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) Execute() (*ToolsInterestActionActionKeywordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsInterestActionActionKeywordGet Method for OpenApi2ToolsInterestActionActionKeywordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsInterestActionActionKeywordGetRequest
*/
func (a *ToolsInterestActionActionKeywordV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest {
	return &ApiOpenApi2ToolsInterestActionActionKeywordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsInterestActionActionKeywordV2Response
func (a *ToolsInterestActionActionKeywordV2ApiService) getExecute(r *ApiOpenApi2ToolsInterestActionActionKeywordGetRequest) (*ToolsInterestActionActionKeywordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsInterestActionActionKeywordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/interest_action/action/keyword/"

	localVarHeaderParams := make(map[string]string)
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
	if r.actionDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_days", r.actionDays)
	}
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
