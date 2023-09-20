/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
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

// ToolsLandingGroupUpdateV2ApiService ToolsLandingGroupUpdateV2Api service
type ToolsLandingGroupUpdateV2ApiService service

type ApiOpenApi2ToolsLandingGroupUpdatePostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsLandingGroupUpdateV2ApiService
	xOrangeCaller                    *string
	toolsLandingGroupUpdateV2Request *ToolsLandingGroupUpdateV2Request
}

func (r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsLandingGroupUpdatePostRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) ToolsLandingGroupUpdateV2Request(toolsLandingGroupUpdateV2Request ToolsLandingGroupUpdateV2Request) *ApiOpenApi2ToolsLandingGroupUpdatePostRequest {
	r.toolsLandingGroupUpdateV2Request = &toolsLandingGroupUpdateV2Request
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) Execute() (*ToolsLandingGroupUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsLandingGroupUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsLandingGroupUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsLandingGroupUpdatePost Method for OpenApi2ToolsLandingGroupUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsLandingGroupUpdatePostRequest
*/
func (a *ToolsLandingGroupUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsLandingGroupUpdatePostRequest {
	return &ApiOpenApi2ToolsLandingGroupUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsLandingGroupUpdateV2Response
func (a *ToolsLandingGroupUpdateV2ApiService) postExecute(r *ApiOpenApi2ToolsLandingGroupUpdatePostRequest) (*ToolsLandingGroupUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsLandingGroupUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/landing_group/update/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.xOrangeCaller != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Orange-Caller", r.xOrangeCaller)
	}
	// body params
	localVarPostBody = r.toolsLandingGroupUpdateV2Request
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
