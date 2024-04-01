/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

// ToolsPlayableListGetV2ApiService ToolsPlayableListGetV2Api service
type ToolsPlayableListGetV2ApiService service

type ApiOpenApi2ToolsPlayableListGetGetRequest struct {
	ctx            context.Context
	ApiService     *ToolsPlayableListGetV2ApiService
	advertiserId   *int64
	page           *int64
	pageSize       *int64
	playableSource *ToolsPlayableListGetV2PlayableSource
	playableUrl    *string
	status         *ToolsPlayableListGetV2Status
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) Page(page int64) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) PlayableSource(playableSource ToolsPlayableListGetV2PlayableSource) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.playableSource = &playableSource
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) PlayableUrl(playableUrl string) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.playableUrl = &playableUrl
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) Status(status ToolsPlayableListGetV2Status) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) Execute() (*ToolsPlayableListGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPlayableListGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPlayableListGetGet Method for OpenApi2ToolsPlayableListGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPlayableListGetGetRequest
*/
func (a *ToolsPlayableListGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPlayableListGetGetRequest {
	return &ApiOpenApi2ToolsPlayableListGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPlayableListGetV2Response
func (a *ToolsPlayableListGetV2ApiService) getExecute(r *ApiOpenApi2ToolsPlayableListGetGetRequest) (*ToolsPlayableListGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPlayableListGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/playable_list/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.playableSource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "playable_source", r.playableSource)
	}
	if r.playableUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "playable_url", r.playableUrl)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
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
