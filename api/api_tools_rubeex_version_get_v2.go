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

// ToolsRubeexVersionGetV2ApiService ToolsRubeexVersionGetV2Api service
type ToolsRubeexVersionGetV2ApiService service

type ApiOpenApi2ToolsRubeexVersionGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsRubeexVersionGetV2ApiService
	advertiserId *int64
	projectId    *float64
}

// 广告主ID
func (r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsRubeexVersionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 作品id即试玩素材id，属于试玩素材唯一标识
func (r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) ProjectId(projectId float64) *ApiOpenApi2ToolsRubeexVersionGetGetRequest {
	r.projectId = &projectId
	return r
}

func (r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) Execute() (*ToolsRubeexVersionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRubeexVersionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRubeexVersionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRubeexVersionGetGet Method for OpenApi2ToolsRubeexVersionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRubeexVersionGetGetRequest
*/
func (a *ToolsRubeexVersionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRubeexVersionGetGetRequest {
	return &ApiOpenApi2ToolsRubeexVersionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRubeexVersionGetV2Response
func (a *ToolsRubeexVersionGetV2ApiService) getExecute(r *ApiOpenApi2ToolsRubeexVersionGetGetRequest) (*ToolsRubeexVersionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsRubeexVersionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rubeex/version/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > -9223372036854775616 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than -9223372036854775616")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, ReportError("projectId is required and must be specified")
	}
	if *r.projectId < 1 {
		return localVarReturnValue, nil, ReportError("projectId must be greater than 1")
	}
	if *r.projectId > 2147483647 {
		return localVarReturnValue, nil, ReportError("projectId must be less than 2147483647")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
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
