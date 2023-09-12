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

// ToolsRegionGetV2ApiService ToolsRegionGetV2Api service
type ToolsRegionGetV2ApiService service

type ApiOpenApi2ToolsRegionGetGetRequest struct {
	ctx         context.Context
	ApiService  *ToolsRegionGetV2ApiService
	regionLevel *ToolsRegionGetV2RegionLevel
	regionType  *ToolsRegionGetV2RegionType
}

func (r *ApiOpenApi2ToolsRegionGetGetRequest) RegionLevel(regionLevel ToolsRegionGetV2RegionLevel) *ApiOpenApi2ToolsRegionGetGetRequest {
	r.regionLevel = &regionLevel
	return r
}

func (r *ApiOpenApi2ToolsRegionGetGetRequest) RegionType(regionType ToolsRegionGetV2RegionType) *ApiOpenApi2ToolsRegionGetGetRequest {
	r.regionType = &regionType
	return r
}

func (r *ApiOpenApi2ToolsRegionGetGetRequest) Execute() (*ToolsRegionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRegionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRegionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRegionGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRegionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRegionGetGet Method for OpenApi2ToolsRegionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRegionGetGetRequest
*/
func (a *ToolsRegionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRegionGetGetRequest {
	return &ApiOpenApi2ToolsRegionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRegionGetV2Response
func (a *ToolsRegionGetV2ApiService) getExecute(r *ApiOpenApi2ToolsRegionGetGetRequest) (*ToolsRegionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsRegionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/region/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.regionLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region_level", r.regionLevel)
	}
	if r.regionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region_type", r.regionType)
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
