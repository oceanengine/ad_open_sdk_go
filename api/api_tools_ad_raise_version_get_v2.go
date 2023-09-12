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

// ToolsAdRaiseVersionGetV2ApiService ToolsAdRaiseVersionGetV2Api service
type ToolsAdRaiseVersionGetV2ApiService service

type ApiOpenApi2ToolsAdRaiseVersionGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdRaiseVersionGetV2ApiService
	adId         *int64
	advertiserId *int64
	pageSize     *int64
	page         *int64
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) Page(page int64) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) Execute() (*ToolsAdRaiseVersionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdRaiseVersionGetGet Method for OpenApi2ToolsAdRaiseVersionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdRaiseVersionGetGetRequest
*/
func (a *ToolsAdRaiseVersionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest {
	return &ApiOpenApi2ToolsAdRaiseVersionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdRaiseVersionGetV2Response
func (a *ToolsAdRaiseVersionGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAdRaiseVersionGetGetRequest) (*ToolsAdRaiseVersionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAdRaiseVersionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_raise_version/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if *r.adId < 1 {
		return localVarReturnValue, nil, ReportError("adId must be greater than 1")
	}
	if *r.adId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("adId must be less than 9223372036854775807")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
