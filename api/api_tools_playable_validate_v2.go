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

// ToolsPlayableValidateV2ApiService ToolsPlayableValidateV2Api service
type ToolsPlayableValidateV2ApiService service

type ApiOpenApi2ToolsPlayableValidateGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsPlayableValidateV2ApiService
	advertiserId *int64
	playableId   *int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsPlayableValidateGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPlayableValidateGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 试玩素材ID
func (r *ApiOpenApi2ToolsPlayableValidateGetRequest) PlayableId(playableId int64) *ApiOpenApi2ToolsPlayableValidateGetRequest {
	r.playableId = &playableId
	return r
}

func (r *ApiOpenApi2ToolsPlayableValidateGetRequest) Execute() (*ToolsPlayableValidateV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPlayableValidateGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPlayableValidateGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPlayableValidateGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPlayableValidateGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPlayableValidateGet Method for OpenApi2ToolsPlayableValidateGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPlayableValidateGetRequest
*/
func (a *ToolsPlayableValidateV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPlayableValidateGetRequest {
	return &ApiOpenApi2ToolsPlayableValidateGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPlayableValidateV2Response
func (a *ToolsPlayableValidateV2ApiService) getExecute(r *ApiOpenApi2ToolsPlayableValidateGetRequest) (*ToolsPlayableValidateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPlayableValidateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/playable/validate/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.playableId == nil {
		return localVarReturnValue, nil, ReportError("playableId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "playable_id", r.playableId)
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
