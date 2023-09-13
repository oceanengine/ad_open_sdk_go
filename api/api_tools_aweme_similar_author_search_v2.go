/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
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

// ToolsAwemeSimilarAuthorSearchV2ApiService ToolsAwemeSimilarAuthorSearchV2Api service
type ToolsAwemeSimilarAuthorSearchV2ApiService service

type ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAwemeSimilarAuthorSearchV2ApiService
	advertiserId *int64
	awemeId      *string
	behaviors    *[]*ToolsAwemeSimilarAuthorSearchV2Behaviors
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) AwemeId(awemeId string) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) Behaviors(behaviors []*ToolsAwemeSimilarAuthorSearchV2Behaviors) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	r.behaviors = &behaviors
	return r
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) Execute() (*ToolsAwemeSimilarAuthorSearchV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAwemeSimilarAuthorSearchGet Method for OpenApi2ToolsAwemeSimilarAuthorSearchGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest
*/
func (a *ToolsAwemeSimilarAuthorSearchV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest {
	return &ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeSimilarAuthorSearchV2Response
func (a *ToolsAwemeSimilarAuthorSearchV2ApiService) getExecute(r *ApiOpenApi2ToolsAwemeSimilarAuthorSearchGetRequest) (*ToolsAwemeSimilarAuthorSearchV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAwemeSimilarAuthorSearchV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aweme_similar_author_search/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	if r.behaviors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "behaviors", r.behaviors)
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
