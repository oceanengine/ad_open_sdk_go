/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// ToolsAdRaiseStatusGetV2ApiService ToolsAdRaiseStatusGetV2Api service
type ToolsAdRaiseStatusGetV2ApiService service

type ApiOpenApi2ToolsAdRaiseStatusGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdRaiseStatusGetV2ApiService
	adIds        *[]int64
	advertiserId *int64
}

func (r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) AdIds(adIds []int64) *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) Execute() (*ToolsAdRaiseStatusGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdRaiseStatusGetGet Method for OpenApi2ToolsAdRaiseStatusGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdRaiseStatusGetGetRequest
*/
func (a *ToolsAdRaiseStatusGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest {
	return &ApiOpenApi2ToolsAdRaiseStatusGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdRaiseStatusGetV2Response
func (a *ToolsAdRaiseStatusGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAdRaiseStatusGetGetRequest) (*ToolsAdRaiseStatusGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdRaiseStatusGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_raise_status/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
