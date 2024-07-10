/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// ToolsInterestActionKeywordSuggestV2ApiService ToolsInterestActionKeywordSuggestV2Api service
type ToolsInterestActionKeywordSuggestV2ApiService service

type ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsInterestActionKeywordSuggestV2ApiService
	advertiserId  *int64
	id            *int64
	tagType       *ToolsInterestActionKeywordSuggestV2TagType
	targetingType *ToolsInterestActionKeywordSuggestV2TargetingType
	actionDays    *ToolsInterestActionKeywordSuggestV2ActionDays
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) Id(id int64) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.id = &id
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) TagType(tagType ToolsInterestActionKeywordSuggestV2TagType) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.tagType = &tagType
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) TargetingType(targetingType ToolsInterestActionKeywordSuggestV2TargetingType) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.targetingType = &targetingType
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) ActionDays(actionDays ToolsInterestActionKeywordSuggestV2ActionDays) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) Execute() (*ToolsInterestActionKeywordSuggestV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsInterestActionKeywordSuggestGet Method for OpenApi2ToolsInterestActionKeywordSuggestGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest
*/
func (a *ToolsInterestActionKeywordSuggestV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest {
	return &ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsInterestActionKeywordSuggestV2Response
func (a *ToolsInterestActionKeywordSuggestV2ApiService) getExecute(r *ApiOpenApi2ToolsInterestActionKeywordSuggestGetRequest) (*ToolsInterestActionKeywordSuggestV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsInterestActionKeywordSuggestV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/interest_action/keyword/suggest/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.id == nil {
		return localVarReturnValue, nil, ReportError("id is required and must be specified")
	}
	if r.tagType == nil {
		return localVarReturnValue, nil, ReportError("tagType is required and must be specified")
	}
	if r.targetingType == nil {
		return localVarReturnValue, nil, ReportError("targetingType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id)
	parameterAddToHeaderOrQuery(localVarQueryParams, "tag_type", r.tagType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "targeting_type", r.targetingType)
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
