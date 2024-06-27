/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// ToolsTaskRaiseGetV2ApiService ToolsTaskRaiseGetV2Api service
type ToolsTaskRaiseGetV2ApiService service

type ApiOpenApi2ToolsTaskRaiseGetGetRequest struct {
	ctx             context.Context
	ApiService      *ToolsTaskRaiseGetV2ApiService
	advertiserId    *int64
	page            *int64
	pageSize        *int64
	platformVersion *ToolsTaskRaiseGetV2PlatformVersion
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) Page(page int64) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 2.0平台填2
func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) PlatformVersion(platformVersion ToolsTaskRaiseGetV2PlatformVersion) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	r.platformVersion = &platformVersion
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) Execute() (*ToolsTaskRaiseGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsTaskRaiseGetGet Method for OpenApi2ToolsTaskRaiseGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsTaskRaiseGetGetRequest
*/
func (a *ToolsTaskRaiseGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsTaskRaiseGetGetRequest {
	return &ApiOpenApi2ToolsTaskRaiseGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsTaskRaiseGetV2Response
func (a *ToolsTaskRaiseGetV2ApiService) getExecute(r *ApiOpenApi2ToolsTaskRaiseGetGetRequest) (*ToolsTaskRaiseGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsTaskRaiseGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/task_raise/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if *r.page > 99999 {
		return localVarReturnValue, nil, ReportError("page must be less than 99999")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 100 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	if r.platformVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform_version", r.platformVersion)
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
