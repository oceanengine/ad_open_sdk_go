/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// ToolsKeywordsProjectInfoGetV30ApiService ToolsKeywordsProjectInfoGetV30Api service
type ToolsKeywordsProjectInfoGetV30ApiService service

type ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest struct {
	ctx             context.Context
	ApiService      *ToolsKeywordsProjectInfoGetV30ApiService
	advertiserId    *int64
	promotionWordId *string
}

func (r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) PromotionWordId(promotionWordId string) *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest {
	r.promotionWordId = &promotionWordId
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) Execute() (*ToolsKeywordsProjectInfoGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsKeywordsProjectInfoGetGet Method for OpenApiV30ToolsKeywordsProjectInfoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest
*/
func (a *ToolsKeywordsProjectInfoGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest {
	return &ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsKeywordsProjectInfoGetV30Response
func (a *ToolsKeywordsProjectInfoGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsKeywordsProjectInfoGetGetRequest) (*ToolsKeywordsProjectInfoGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsKeywordsProjectInfoGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/keywords_project_info/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.promotionWordId == nil {
		return localVarReturnValue, nil, ReportError("promotionWordId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_word_id", r.promotionWordId)
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
