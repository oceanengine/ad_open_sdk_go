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

// ToolsVideoCoverSuggestV2ApiService ToolsVideoCoverSuggestV2Api service
type ToolsVideoCoverSuggestV2ApiService service

type ApiOpenApi2ToolsVideoCoverSuggestGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsVideoCoverSuggestV2ApiService
	advertiserId *int64
	videoId      *string
}

// 广告主ID
func (r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsVideoCoverSuggestGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 视频id
func (r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) VideoId(videoId string) *ApiOpenApi2ToolsVideoCoverSuggestGetRequest {
	r.videoId = &videoId
	return r
}

func (r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) Execute() (*ToolsVideoCoverSuggestV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsVideoCoverSuggestGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsVideoCoverSuggestGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsVideoCoverSuggestGet Method for OpenApi2ToolsVideoCoverSuggestGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsVideoCoverSuggestGetRequest
*/
func (a *ToolsVideoCoverSuggestV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsVideoCoverSuggestGetRequest {
	return &ApiOpenApi2ToolsVideoCoverSuggestGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsVideoCoverSuggestV2Response
func (a *ToolsVideoCoverSuggestV2ApiService) getExecute(r *ApiOpenApi2ToolsVideoCoverSuggestGetRequest) (*ToolsVideoCoverSuggestV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsVideoCoverSuggestV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/video_cover/suggest/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.videoId == nil {
		return localVarReturnValue, nil, ReportError("videoId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "video_id", r.videoId)
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
