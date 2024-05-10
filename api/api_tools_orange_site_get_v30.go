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

// ToolsOrangeSiteGetV30ApiService ToolsOrangeSiteGetV30Api service
type ToolsOrangeSiteGetV30ApiService service

type ApiOpenApiV30ToolsOrangeSiteGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsOrangeSiteGetV30ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
	optimizeGoal *ToolsOrangeSiteGetV30OptimizeGoal
	status       *ToolsOrangeSiteGetV30Status
	filtering    *ToolsOrangeSiteGetV30Filtering
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) Page(page int32) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) OptimizeGoal(optimizeGoal ToolsOrangeSiteGetV30OptimizeGoal) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.optimizeGoal = &optimizeGoal
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) Status(status ToolsOrangeSiteGetV30Status) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) Filtering(filtering ToolsOrangeSiteGetV30Filtering) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) Execute() (*ToolsOrangeSiteGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsOrangeSiteGetGet Method for OpenApiV30ToolsOrangeSiteGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsOrangeSiteGetGetRequest
*/
func (a *ToolsOrangeSiteGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsOrangeSiteGetGetRequest {
	return &ApiOpenApiV30ToolsOrangeSiteGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsOrangeSiteGetV30Response
func (a *ToolsOrangeSiteGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsOrangeSiteGetGetRequest) (*ToolsOrangeSiteGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsOrangeSiteGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/orange_site/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if r.optimizeGoal == nil {
		return localVarReturnValue, nil, ReportError("optimizeGoal is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "optimize_goal", r.optimizeGoal)
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
